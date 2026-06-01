package phone

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/spf13/cobra"
)

func newFileCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file",
		Short: "File management",
		Long: `Manage files — upload temporary files to GeeLark, upload files to cloud phones, and upload keybox files.

The "upload-temp" command combines two steps into one:
  1. Get upload URL via /open/v1/upload/getUrl
  2. Upload the local file to the returned uploadUrl (PUT)
  Returns the resourceUrl (valid for 3 days).

The "upload-to-phone" command supports both a direct URL (--url) and a local file (--file).
When --file is provided, it first uploads the file to GeeLark temp storage, then uploads to the cloud phone.`,
	}

	cmd.AddCommand(newUploadTempCmd(newClient))
	cmd.AddCommand(newUploadToPhoneCmd(newClient))
	cmd.AddCommand(newUploadStatusCmd(newClient))
	cmd.AddCommand(newKeyboxUploadDirectCmd(newClient))
	cmd.AddCommand(newKeyboxResultCmd(newClient))

	return cmd
}

var supportedFileTypes = map[string]bool{
	"jpg": true, "jpeg": true, "png": true, "gif": true,
	"bmp": true, "webp": true, "heif": true, "heic": true,
	"mp4": true, "webm": true, "xml": true, "apk": true, "xapk": true,
}

func getFileExt(path string) (string, error) {
	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(path)), ".")
	if ext == "" {
		return "", fmt.Errorf("cannot determine file type from extension: %s", path)
	}
	if !supportedFileTypes[ext] {
		return "", fmt.Errorf("unsupported file type: %s (supported: jpg, jpeg, png, gif, bmp, webp, heif, heic, mp4, webm, xml, apk, xapk)", ext)
	}
	return ext, nil
}

func uploadTempFile(c *client.Client, filePath string) (string, error) {
	ext, err := getFileExt(filePath)
	if err != nil {
		return "", err
	}

	resp, err := c.Post("/open/v1/upload/getUrl", map[string]interface{}{"fileType": ext})
	if err != nil {
		return "", fmt.Errorf("failed to get upload URL: %w", err)
	}
	if resp.Code != 0 {
		return "", fmt.Errorf("getUrl API error: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	var uploadData struct {
		UploadUrl   string `json:"uploadUrl"`
		ResourceUrl string `json:"resourceUrl"`
	}
	if err := json.Unmarshal(resp.Data, &uploadData); err != nil {
		return "", fmt.Errorf("failed to parse upload URL response: %w", err)
	}

	fmt.Printf("Uploading file %s ...\n", filePath)
	if err := c.PutFile(uploadData.UploadUrl, filePath); err != nil {
		return "", fmt.Errorf("file upload failed: %w", err)
	}
	fmt.Println("File uploaded successfully.")

	return uploadData.ResourceUrl, nil
}

func newUploadTempCmd(newClient clientFactory) *cobra.Command {
	var file string
	cmd := &cobra.Command{
		Use:   "upload-temp",
		Short: "Upload a local file to GeeLark temporary storage",
		Long: `Upload a local file to GeeLark temporary storage (valid for 3 days).
For longer storage, use "library material-create" instead.

This command combines:
  1. Get upload URL via /open/v1/upload/getUrl
  2. Upload the local file to the returned uploadUrl (PUT)
Returns the resourceUrl for accessing the file.`,
		Example: `  geelark-cli phone file upload-temp --file ./video.mp4
  geelark-cli phone file upload-temp --file ./photo.jpg`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			resourceUrl, err := uploadTempFile(c, file)
			if err != nil {
				return err
			}

			fmt.Printf("resourceUrl: %s\n", resourceUrl)
			fmt.Println("Note: temporary files expire in 3 days. Use 'library material-create' for longer storage (30 days).")
			return nil
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Local file path to upload (required)")
	_ = cmd.MarkFlagRequired("file")
	return cmd
}

func newUploadToPhoneCmd(newClient clientFactory) *cobra.Command {
	var id, fileURL, file string
	cmd := &cobra.Command{
		Use:   "upload-to-phone",
		Short: "Upload a file to a cloud phone",
		Long: `Upload a file to the "Downloads" folder of a cloud phone.
The cloud phone must be running before uploading.

Supports two modes:
  --url:  Provide a file URL directly (e.g. from library or temp upload)
  --file: Provide a local file path (will be uploaded to GeeLark temp storage first)`,
		Example: `  geelark-cli phone file upload-to-phone --id "557536075321468390" --url "https://material.geelark.cn/xxx.mp4"
  geelark-cli phone file upload-to-phone --id "557536075321468390" --file ./local-video.mp4`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if fileURL == "" && file == "" {
				return fmt.Errorf("either --url or --file must be provided")
			}
			if fileURL != "" && file != "" {
				return fmt.Errorf("--url and --file are mutually exclusive, please provide only one")
			}

			resourceUrl := fileURL
			if file != "" {
				resourceUrl, err = uploadTempFile(c, file)
				if err != nil {
					return err
				}
			}

			result, err := c.PostAndPrint("/open/v1/phone/uploadFile", map[string]interface{}{
				"id":      id,
				"fileUrl": resourceUrl,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&fileURL, "url", "", "File URL to upload")
	cmd.Flags().StringVar(&file, "file", "", "Local file path (will be uploaded to temp storage first)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newUploadStatusCmd(newClient clientFactory) *cobra.Command {
	var taskID string
	cmd := &cobra.Command{
		Use:     "upload-status",
		Short:   "Query the upload status of a file to the cloud phone",
		Long:    "Query the upload status within 1 hour of initiating the upload task.",
		Example: `  geelark-cli phone file upload-status --task-id "1850726441252569088"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/phone/uploadFile/result", map[string]interface{}{
				"taskId": taskID,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&taskID, "task-id", "", "Upload task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}

func newKeyboxUploadDirectCmd(newClient clientFactory) *cobra.Command {
	var id, fileURL, file string
	cmd := &cobra.Command{
		Use:   "keybox-upload",
		Short: "Upload a keybox file to a cloud phone",
		Long: `Upload a keybox file to pass Google's integrity verification.
Currently only supported on Android 12/13/15.

Supports two modes:
  --url:  Provide a file URL directly
  --file: Provide a local file path (will be uploaded to GeeLark temp storage first)`,
		Example: `  geelark-cli phone file keybox-upload --id "557536075321468390" --url "https://material.geelark.cn/xxx.xml"
  geelark-cli phone file keybox-upload --id "557536075321468390" --file ./keybox.xml`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			if fileURL == "" && file == "" {
				return fmt.Errorf("either --url or --file must be provided")
			}
			if fileURL != "" && file != "" {
				return fmt.Errorf("--url and --file are mutually exclusive, please provide only one")
			}

			resourceUrl := fileURL
			if file != "" {
				resourceUrl, err = uploadTempFile(c, file)
				if err != nil {
					return err
				}
			}

			result, err := c.PostAndPrint("/open/v1/phone/keyboxUpload", map[string]interface{}{
				"id":      id,
				"fileUrl": resourceUrl,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (required)")
	cmd.Flags().StringVar(&fileURL, "url", "", "Keybox file URL")
	cmd.Flags().StringVar(&file, "file", "", "Local keybox file path (will be uploaded to temp storage first)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}

func newKeyboxResultCmd(newClient clientFactory) *cobra.Command {
	var taskID string
	cmd := &cobra.Command{
		Use:     "keybox-result",
		Short:   "Query the keybox upload task result",
		Example: `  geelark-cli phone file keybox-result --task-id "1850726441252569088"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/phone/keyboxUpload/result", map[string]interface{}{
				"id": taskID,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&taskID, "task-id", "", "Keybox upload task ID (required)")
	_ = cmd.MarkFlagRequired("task-id")
	return cmd
}
