package phone

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func newLibraryCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "library",
		Short: "Library (material) management",
		Long: `Manage GeeLark Library — upload, search, delete materials and manage tags.

The "material-create" command combines three steps into one:
  1. Get upload URL via /open/v1/material/getUploadUrl
  2. Upload the local file to the returned uploadUrl (PUT)
  3. Create the material record via /open/v1/material/create using the resourceUrl`,
	}

	cmd.AddCommand(newMaterialCreateCmd(newClient))
	cmd.AddCommand(newMaterialSearchCmd(newClient))
	cmd.AddCommand(newMaterialDeleteCmd(newClient))
	cmd.AddCommand(newMaterialTagCreateCmd(newClient))
	cmd.AddCommand(newMaterialTagSearchCmd(newClient))
	cmd.AddCommand(newMaterialTagDeleteCmd(newClient))
	cmd.AddCommand(newMaterialTagSetCmd(newClient))

	return cmd
}

func newMaterialCreateCmd(newClient clientFactory) *cobra.Command {
	var file, fileName, tagsId string
	cmd := &cobra.Command{
		Use:   "material-create",
		Short: "Upload a file and create a material in the Library",
		Long: `Upload a local file to the Library and create a material record.
This command combines three API calls into one:
  1. Get upload URL (getUploadUrl)
  2. Upload the file to the URL (PUT)
  3. Create the material (create)`,
		Example: `  geelark-cli phone library material-create --file ./photo.jpg
  geelark-cli phone library material-create --file ./video.mp4 --tags-id "tag1,tag2"
  geelark-cli phone library material-create --file ./data.xml --file-name "my-data.xml"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			ext, err := getFileExt(file)
			if err != nil {
				return err
			}

			resp, err := c.Post("/open/v1/material/getUploadUrl", map[string]interface{}{"fileType": ext})
			if err != nil {
				return fmt.Errorf("failed to get upload URL: %w", err)
			}
			if resp.Code != 0 {
				return fmt.Errorf("getUploadUrl API error: code=%d, msg=%s", resp.Code, resp.Msg)
			}

			var uploadData struct {
				UploadUrl   string `json:"uploadUrl"`
				ResourceUrl string `json:"resourceUrl"`
			}
			if err := json.Unmarshal(resp.Data, &uploadData); err != nil {
				return fmt.Errorf("failed to parse upload URL response: %w", err)
			}

			if err := c.PutFile(uploadData.UploadUrl, file); err != nil {
				return fmt.Errorf("file upload failed: %w", err)
			}

			createBody := map[string]interface{}{"url": uploadData.ResourceUrl}
			if fileName != "" {
				createBody["fileName"] = fileName
			} else {
				createBody["fileName"] = filepath.Base(file)
			}
			if tagsId != "" {
				createBody["tagsId"] = strings.Split(tagsId, ",")
			}

			result, err := c.PostAndPrint("/open/v1/material/create", createBody)
			if err != nil {
				return fmt.Errorf("failed to create material: %w", err)
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&file, "file", "", "Local file path to upload (required)")
	cmd.Flags().StringVar(&fileName, "file-name", "", "Material name (defaults to file base name, max 200 chars)")
	cmd.Flags().StringVar(&tagsId, "tags-id", "", "Comma-separated tag IDs")
	_ = cmd.MarkFlagRequired("file")
	return cmd
}

func newMaterialSearchCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var fileName, tagsId string
	var source int
	var fileType, ids string
	var sourceSet bool
	cmd := &cobra.Command{
		Use:     "material-search",
		Short:   "Search materials in the Library",
		Example: `  geelark-cli phone library material-search --page 1 --page-size 50`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if page > 0 {
				body["page"] = page
			}
			if pageSize > 0 {
				body["pageSize"] = pageSize
			}
			if fileName != "" {
				body["fileName"] = fileName
			}
			if tagsId != "" {
				body["tagIds"] = strings.Split(tagsId, ",")
			}
			if sourceSet {
				body["source"] = source
			}
			if fileType != "" {
				var ft []interface{}
				for _, s := range strings.Split(fileType, ",") {
					var v int
					fmt.Sscanf(s, "%d", &v)
					ft = append(ft, v)
				}
				body["fileType"] = ft
			}
			if ids != "" {
				body["ids"] = strings.Split(ids, ",")
			}
			result, err := c.PostAndPrint("/open/v1/material/search", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 50, "Page size (max 200)")
	cmd.Flags().StringVar(&fileName, "file-name", "", "Search by file name")
	cmd.Flags().StringVar(&tagsId, "tags-id", "", "Comma-separated tag IDs")
	cmd.Flags().IntVar(&source, "source", -1, "Source: 0=upload, 1=AI Edit, 2=Baidu Cloud, 3=GhostCut, 4=GoogleDrive, 5=Image to video")
	cmd.Flags().StringVar(&fileType, "file-type", "", "Comma-separated file types: 1=image, 2=video")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated material IDs (max 100)")
	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		sourceSet = cmd.Flags().Changed("source")
		return nil
	}
	return cmd
}

func newMaterialDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string
	cmd := &cobra.Command{
		Use:     "material-delete",
		Short:   "Delete materials from the Library",
		Example: `  geelark-cli phone library material-delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/material/del", map[string]interface{}{"ids": strings.Split(ids, ",")})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated material IDs (required)")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newMaterialTagCreateCmd(newClient clientFactory) *cobra.Command {
	var name string
	var color int
	var colorSet bool
	cmd := &cobra.Command{
		Use:     "tag-create",
		Short:   "Create a material tag",
		Example: `  geelark-cli phone library tag-create --name "my-tag" --color 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"name": name}
			if colorSet {
				body["color"] = color
			}
			result, err := c.PostAndPrint("/open/v1/material/tag/create", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Tag name, max 30 chars (required)")
	cmd.Flags().IntVar(&color, "color", 0, "Tag color: 0=White, 1=Red, 2=Blue, 3=Green, 4=Yellow, 5=Purple")
	_ = cmd.MarkFlagRequired("name")
	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		colorSet = cmd.Flags().Changed("color")
		return nil
	}
	return cmd
}

func newMaterialTagSearchCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var name string
	cmd := &cobra.Command{
		Use:     "tag-search",
		Short:   "Search material tags",
		Example: `  geelark-cli phone library tag-search --page 1 --page-size 50`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{}
			if page > 0 {
				body["page"] = page
			}
			if pageSize > 0 {
				body["pageSize"] = pageSize
			}
			if name != "" {
				body["name"] = name
			}
			result, err := c.PostAndPrint("/open/v1/material/tag/search", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 50, "Page size (max 200)")
	cmd.Flags().StringVar(&name, "name", "", "Search by tag name")
	return cmd
}

func newMaterialTagDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string
	cmd := &cobra.Command{
		Use:     "tag-delete",
		Short:   "Delete material tags",
		Example: `  geelark-cli phone library tag-delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			result, err := c.PostAndPrint("/open/v1/material/tag/del", map[string]interface{}{"ids": strings.Split(ids, ",")})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated tag IDs (required)")
	_ = cmd.MarkFlagRequired("ids")
	return cmd
}

func newMaterialTagSetCmd(newClient clientFactory) *cobra.Command {
	var materialsId, tagsId string
	cmd := &cobra.Command{
		Use:   "tag-set",
		Short: "Set tags on materials",
		Long: `Set tags on materials. Each tagging operation will delete all existing tags
and only retain the newly set tags.`,
		Example: `  geelark-cli phone library tag-set --materials-id "mat1,mat2" --tags-id "tag1,tag2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{"materialsId": strings.Split(materialsId, ",")}
			if tagsId != "" {
				body["tagsId"] = strings.Split(tagsId, ",")
			}
			result, err := c.PostAndPrint("/open/v1/material/tag/set", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&materialsId, "materials-id", "", "Comma-separated material IDs (required)")
	cmd.Flags().StringVar(&tagsId, "tags-id", "", "Comma-separated tag IDs (will replace all existing tags)")
	_ = cmd.MarkFlagRequired("materials-id")
	return cmd
}
