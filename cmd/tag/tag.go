package tag

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

// NewCmd creates the tag command group.
func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag",
		Short: "Tag management",
		Long:  "Manage cloud phone tags — list, create, update, and delete.",
	}

	cmd.AddCommand(newListCmd(newClient))
	cmd.AddCommand(newCreateCmd(newClient))
	cmd.AddCommand(newSimpleCreateCmd(newClient))
	cmd.AddCommand(newUpdateCmd(newClient))
	cmd.AddCommand(newSimpleUpdateCmd(newClient))
	cmd.AddCommand(newDeleteCmd(newClient))

	return cmd
}

func newListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var ids, names, colors string

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List tags",
		Long:    "Retrieve tag information with optional filters.",
		Example: `  geelark-cli tag list --page 1 --page-size 10`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"page":     page,
				"pageSize": pageSize,
			}
			if ids != "" {
				body["ids"] = strings.Split(ids, ",")
			}
			if names != "" {
				body["names"] = strings.Split(names, ",")
			}
			if colors != "" {
				body["colors"] = strings.Split(colors, ",")
			}

			result, err := c.PostAndPrint("/open/v1/tag/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 100)")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated tag IDs")
	cmd.Flags().StringVar(&names, "names", "", "Comma-separated tag names")
	cmd.Flags().StringVar(&colors, "colors", "", "Comma-separated tag colors (white,red,blue,green,yellow,purple)")

	return cmd
}

func newCreateCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create tags",
		Long: `Create tags with specified name and optional color. Supports batch creation.
Colors: white, red, blue, green, yellow, purple (default: white)`,
		Example: `  geelark-cli tag create --data "[{\"name\":\"myTag\",\"color\":\"red\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			var list []interface{}
			if err := json.Unmarshal([]byte(dataJSON), &list); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}

			body := map[string]interface{}{
				"list": list,
			}

			result, err := c.PostAndPrint("/open/v1/tag/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of tags [{"name":"...","color":"..."}] (required)`)
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newUpdateCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:     "update",
		Short:   "Update tags",
		Long:    "Modify tag information including name and color.",
		Example: `  geelark-cli tag update --data "[{\"id\":\"tag_id\",\"name\":\"newName\",\"color\":\"blue\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			var list []interface{}
			if err := json.Unmarshal([]byte(dataJSON), &list); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}

			body := map[string]interface{}{
				"list": list,
			}

			result, err := c.PostAndPrint("/open/v1/tag/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of tags [{"id":"...","name":"...","color":"..."}] (required)`)
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newSimpleCreateCmd(newClient clientFactory) *cobra.Command {
	var name, color string

	cmd := &cobra.Command{
		Use:   "simple-create",
		Short: "Quick create a single tag",
		Long: `Simplified command to create a single tag with flat flags.
Use 'create' for batch creation.
Colors: white (default), red, blue, green, yellow, purple`,
		Example: `  geelark-cli tag simple-create --name "myTag"
  geelark-cli tag simple-create --name "myTag" --color red`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			item := map[string]interface{}{
				"name": name,
			}
			if color != "" {
				item["color"] = color
			}

			body := map[string]interface{}{
				"list": []interface{}{item},
			}

			result, err := c.PostAndPrint("/open/v1/tag/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Tag name, up to 30 characters (required)")
	cmd.Flags().StringVar(&color, "color", "", "Tag color: white (default), red, blue, green, yellow, purple")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func newSimpleUpdateCmd(newClient clientFactory) *cobra.Command {
	var id, name, color string

	cmd := &cobra.Command{
		Use:   "simple-update",
		Short: "Quick update a single tag",
		Long: `Simplified command to update a single tag with flat flags.
Use 'update' for batch updates.
Colors: white, red, blue, green, yellow, purple`,
		Example: `  geelark-cli tag simple-update --id "tag_id" --name "newName"
  geelark-cli tag simple-update --id "tag_id" --color "blue"
  geelark-cli tag simple-update --id "tag_id" --name "newName" --color "blue"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			item := map[string]interface{}{
				"id": id,
			}
			if name != "" {
				item["name"] = name
			}
			if color != "" {
				item["color"] = color
			}

			body := map[string]interface{}{
				"list": []interface{}{item},
			}

			result, err := c.PostAndPrint("/open/v1/tag/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Tag ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "New tag name, up to 30 characters")
	cmd.Flags().StringVar(&color, "color", "", "New tag color: white, red, blue, green, yellow, purple")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete tags",
		Long:    "Delete tags by IDs.",
		Example: `  geelark-cli tag delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/tag/delete", body)
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
