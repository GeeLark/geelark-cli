package phone

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func newOEMCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oem",
		Short: "OEM / White Label customization",
		Long: `Customize brand name, logo, sidebar entrance, QR code domain, etc.
Requires a plan with 50+ profiles.`,
	}

	cmd.AddCommand(newOEMCustomizationCmd(newClient))

	return cmd
}

func newOEMCustomizationCmd(newClient clientFactory) *cobra.Command {
	var title, logo, mirrorURL, toolBarJSON string
	var hideHeader, hideSidebar bool
	var hideHeaderSet, hideSidebarSet bool

	cmd := &cobra.Command{
		Use:   "customization",
		Short: "Customize OEM/white label settings",
		Long: `Customize the cloud phone OEM/white label settings.
- title: Brand title (max 64 bytes)
- logo: Logo URL (max 255 bytes)
- hide-header: Hide header at the top of the cloud phone
- hide-sidebar: Hide the sidebar
- mirror-url: QR code/Mirror entrance URL (max 255 chars)
- toolbar: JSON array of toolbar settings.
  Items: networkQuality, rotate, screenshot, upload, library, volumeUp, volumeDown,
         speedUp, detection, quality, restart, appStore, qcode, export, timing,
         liveStreaming, clear, teamApp`,
		Example: `  geelark-cli phone oem customization --title "MyBrand" --logo "https://example.com/logo.png"
  geelark-cli phone oem customization --hide-header --mirror-url "https://www.abcd.com/mirror/url"
  geelark-cli phone oem customization --hide-sidebar
  geelark-cli phone oem customization --toolbar "[{\"toolBar\":\"networkQuality\",\"visible\":false},{\"toolBar\":\"rotate\",\"visible\":false}]"
  geelark-cli phone oem customization --title "X" --toolbar "[{\"toolBar\":\"screenshot\",\"visible\":true,\"iconUrl\":\"https://example.com/icon.svg\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{}
			if title != "" {
				body["title"] = title
			}
			if logo != "" {
				body["logo"] = logo
			}
			if hideHeaderSet {
				body["hideHeader"] = hideHeader
			}
			if hideSidebarSet {
				body["hideSidebar"] = hideSidebar
			}
			if mirrorURL != "" {
				body["mirrorUrl"] = mirrorURL
			}
			if toolBarJSON != "" {
				var tb []interface{}
				if err := json.Unmarshal([]byte(toolBarJSON), &tb); err != nil {
					return fmt.Errorf("invalid --toolbar JSON: %w", err)
				}
				body["toolBarSettings"] = tb
			}

			result, err := c.PostAndPrint("/open/v1/phone/customization", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "", "Brand title, max 64 bytes")
	cmd.Flags().StringVar(&logo, "logo", "", "Logo URL, max 255 bytes")
	cmd.Flags().BoolVar(&hideHeader, "hide-header", false, "Hide the header at the top of the cloud phone")
	cmd.Flags().BoolVar(&hideSidebar, "hide-sidebar", false, "Hide the sidebar")
	cmd.Flags().StringVar(&mirrorURL, "mirror-url", "", "QR code/Mirror entrance URL, max 255 chars")
	cmd.Flags().StringVar(&toolBarJSON, "toolbar", "", `JSON array of toolbar settings [{"toolBar":"...","visible":bool,"iconUrl":"..."}]`)

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		hideHeaderSet = cmd.Flags().Changed("hide-header")
		hideSidebarSet = cmd.Flags().Changed("hide-sidebar")
		return nil
	}

	return cmd
}
