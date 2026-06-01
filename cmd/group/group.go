package group

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

// NewCmd creates the group command group.
func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group",
		Short: "Group management",
		Long:  "Manage cloud phone groups — list, create, update, and delete.",
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
	var ids, names string

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List groups",
		Long:    "Retrieve group information with optional filters.",
		Example: `  geelark-cli group list --page 1 --page-size 10`,
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

			result, err := c.PostAndPrint("/open/v1/group/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 100)")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated group IDs")
	cmd.Flags().StringVar(&names, "names", "", "Comma-separated group names")

	return cmd
}

func newCreateCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create groups",
		Long:    "Create groups with specified name and optional remark. Supports batch creation.",
		Example: `  geelark-cli group create --data "[{\"name\":\"myGroup\",\"remark\":\"my remark\"}]"`,
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

			result, err := c.PostAndPrint("/open/v1/group/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of groups [{"name":"...", "remark":"..."}] (required)`)
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newUpdateCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:     "update",
		Short:   "Update groups",
		Long:    "Modify group information including name and remark.",
		Example: `  geelark-cli group update --data "[{\"id\":\"group_id\",\"name\":\"newName\",\"remark\":\"newRemark\"}]"`,
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

			result, err := c.PostAndPrint("/open/v1/group/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of groups [{"id":"...","name":"...","remark":"..."}] (required)`)
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newSimpleCreateCmd(newClient clientFactory) *cobra.Command {
	var name, remark string

	cmd := &cobra.Command{
		Use:   "simple-create",
		Short: "Quick create a single group",
		Long:  "Simplified command to create a single group with flat flags. Use 'create' for batch creation.",
		Example: `  geelark-cli group simple-create --name "myGroup"
  geelark-cli group simple-create --name "myGroup" --remark "my remark"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			item := map[string]interface{}{
				"name": name,
			}
			if remark != "" {
				item["remark"] = remark
			}

			body := map[string]interface{}{
				"list": []interface{}{item},
			}

			result, err := c.PostAndPrint("/open/v1/group/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Group name, up to 50 characters (required)")
	cmd.Flags().StringVar(&remark, "remark", "", "Group remark, up to 500 characters")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func newSimpleUpdateCmd(newClient clientFactory) *cobra.Command {
	var id, name, remark string

	cmd := &cobra.Command{
		Use:   "simple-update",
		Short: "Quick update a single group",
		Long:  "Simplified command to update a single group with flat flags. Use 'update' for batch updates.",
		Example: `  geelark-cli group simple-update --id "group_id" --name "newName"
  geelark-cli group simple-update --id "group_id" --name "newName" --remark "newRemark"
  geelark-cli group simple-update --id "group_id" --remark "newRemark"`,
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
			if remark != "" {
				item["remark"] = remark
			}

			body := map[string]interface{}{
				"list": []interface{}{item},
			}

			result, err := c.PostAndPrint("/open/v1/group/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Group ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "New group name, up to 50 characters")
	cmd.Flags().StringVar(&remark, "remark", "", "New group remark, up to 500 characters")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete groups",
		Long:    "Delete groups by IDs.",
		Example: `  geelark-cli group delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/group/delete", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated group IDs (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}
