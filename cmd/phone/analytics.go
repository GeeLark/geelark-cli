package phone

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newAnalyticsCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "analytics",
		Short: "Analytics account management",
		Long:  "Manage analytics accounts (TikTok/YouTube/Instagram/Reddit) — list, add, update, delete, and query data.",
	}

	cmd.AddCommand(newAnalyticsAccountsListCmd(newClient))
	cmd.AddCommand(newAnalyticsAddAccountsCmd(newClient))
	cmd.AddCommand(newAnalyticsSimpleAddAccountCmd(newClient))
	cmd.AddCommand(newAnalyticsUpdateAccountCmd(newClient))
	cmd.AddCommand(newAnalyticsDeleteAccountCmd(newClient))
	cmd.AddCommand(newAnalyticsDataCmd(newClient))
	cmd.AddCommand(newAnalyticsTagsListCmd(newClient))
	cmd.AddCommand(newAnalyticsTagsCreateCmd(newClient))
	cmd.AddCommand(newAnalyticsTagsUpdateCmd(newClient))
	cmd.AddCommand(newAnalyticsTagsDeleteCmd(newClient))

	return cmd
}

func newAnalyticsAccountsListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize, channel int
	var account, userAccount string
	var channelSet bool

	cmd := &cobra.Command{
		Use:   "accounts-list",
		Short: "List analytics accounts",
		Long: `List analytics accounts with optional filters.
Channel: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit (omit for all)`,
		Example: `  geelark-cli phone analytics accounts-list --page 1 --page-size 10
  geelark-cli phone analytics accounts-list --channel 0 --account "tk_acc"
  geelark-cli phone analytics accounts-list --user-account "abc@gmail.com"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"page":     page,
				"pageSize": pageSize,
			}
			if account != "" {
				body["account"] = account
			}
			if channelSet {
				body["channel"] = channel
			}
			if userAccount != "" {
				body["userAccount"] = userAccount
			}
			result, err := c.PostAndPrint("/open/v1/analytics/accounts/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (1-100)")
	cmd.Flags().StringVar(&account, "account", "", "Account name filter")
	cmd.Flags().IntVar(&channel, "channel", -1, "Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit")
	cmd.Flags().StringVar(&userAccount, "user-account", "", "Operator account email")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		channelSet = cmd.Flags().Changed("channel")
		return nil
	}

	return cmd
}

func newAnalyticsAddAccountsCmd(newClient clientFactory) *cobra.Command {
	var channel int
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "add-accounts",
		Short: "Add analytics accounts (batch)",
		Long: `Add multiple analytics accounts via JSON. Max 200 per request.
For a single account, use 'simple-add-account'.
Channel: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit`,
		Example: `  geelark-cli phone analytics add-accounts --channel 0 --data "[{\"account\":\"acc1\",\"remark\":\"my note\"},{\"account\":\"acc2\"}]"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			var accountsData []interface{}
			if err := json.Unmarshal([]byte(dataJSON), &accountsData); err != nil {
				return fmt.Errorf("invalid --data JSON: %w", err)
			}
			body := map[string]interface{}{
				"channel":      channel,
				"accountsData": accountsData,
			}
			result, err := c.PostAndPrint("/open/v1/analytics/accounts/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&channel, "channel", 0, "Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit (required)")
	cmd.Flags().StringVar(&dataJSON, "data", "", `JSON array of accounts [{"account":"...","remark":"..."}] (required)`)
	_ = cmd.MarkFlagRequired("channel")
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newAnalyticsSimpleAddAccountCmd(newClient clientFactory) *cobra.Command {
	var channel int
	var account, remark, tagIds string

	cmd := &cobra.Command{
		Use:   "simple-add-account",
		Short: "Quick add a single analytics account",
		Long: `Simplified command to add a single analytics account with flat flags.
Channel: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit`,
		Example: `  geelark-cli phone analytics simple-add-account --channel 0 --account "myAccount"
  geelark-cli phone analytics simple-add-account --channel 1 --account "ytAcc" --remark "my note" --tag-ids "tag1,tag2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			item := map[string]interface{}{
				"account": account,
			}
			if remark != "" {
				item["remark"] = remark
			}
			if tagIds != "" {
				item["tagIds"] = strings.Split(tagIds, ",")
			}
			body := map[string]interface{}{
				"channel":      channel,
				"accountsData": []interface{}{item},
			}
			result, err := c.PostAndPrint("/open/v1/analytics/accounts/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&channel, "channel", 0, "Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit (required)")
	cmd.Flags().StringVar(&account, "account", "", "Account name, max 64 chars (required)")
	cmd.Flags().StringVar(&remark, "remark", "", "Remark/note")
	cmd.Flags().StringVar(&tagIds, "tag-ids", "", "Comma-separated tag IDs, max 20 after deduplication")
	_ = cmd.MarkFlagRequired("channel")
	_ = cmd.MarkFlagRequired("account")

	return cmd
}

func newAnalyticsUpdateAccountCmd(newClient clientFactory) *cobra.Command {
	var id, account, remark, tagIds string
	var channel int
	var channelSet, tagIdsSet bool

	cmd := &cobra.Command{
		Use:   "update-account",
		Short: "Update an analytics account",
		Long: `Update an analytics account by ID.
Channel: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit`,
		Example: `  geelark-cli phone analytics update-account --id "565523829426802069" --account "newName"
  geelark-cli phone analytics update-account --id "id" --remark "new remark" --channel 1
  geelark-cli phone analytics update-account --id "id" --tag-ids "tag1,tag2"
  geelark-cli phone analytics update-account --id "id" --tag-ids ""`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"id": id,
			}
			if account != "" {
				body["account"] = account
			}
			if channelSet {
				body["channel"] = channel
			}
			if remark != "" {
				body["remark"] = remark
			}
			if tagIdsSet {
				if tagIds != "" {
					body["tagIds"] = strings.Split(tagIds, ",")
				} else {
					body["tagIds"] = []string{}
				}
			}
			result, err := c.PostAndPrint("/open/v1/analytics/accounts/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Account ID (required)")
	cmd.Flags().StringVar(&account, "account", "", "New platform account, max 64 chars")
	cmd.Flags().IntVar(&channel, "channel", -1, "Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit")
	cmd.Flags().StringVar(&remark, "remark", "", "New remark")
	cmd.Flags().StringVar(&tagIds, "tag-ids", "", "Comma-separated tag IDs (max 20); pass empty string to clear tags")
	_ = cmd.MarkFlagRequired("id")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		channelSet = cmd.Flags().Changed("channel")
		tagIdsSet = cmd.Flags().Changed("tag-ids")
		return nil
	}

	return cmd
}

func newAnalyticsDeleteAccountCmd(newClient clientFactory) *cobra.Command {
	var channel int
	var account string

	cmd := &cobra.Command{
		Use:   "delete-account",
		Short: "Delete an analytics account",
		Long: `Delete an analytics account by channel and account name.
Channel: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit`,
		Example: `  geelark-cli phone analytics delete-account --channel 0 --account "myAccount"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"channel": channel,
				"account": account,
			}
			result, err := c.PostAndPrint("/open/v1/analytics/accounts/delete", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&channel, "channel", 0, "Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit (required)")
	cmd.Flags().StringVar(&account, "account", "", "Account name (required)")
	_ = cmd.MarkFlagRequired("channel")
	_ = cmd.MarkFlagRequired("account")

	return cmd
}

func newAnalyticsDataCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize, channel, dataDate int
	var account, createdId string
	var channelSet, dataDateSet bool

	cmd := &cobra.Command{
		Use:   "data",
		Short: "Get analytics account data",
		Long: `Query analytics account data (play count, follower count, digg, comments, etc.).
Channel: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit. Requires Pro plan.`,
		Example: `  geelark-cli phone analytics data --page 1 --page-size 10
  geelark-cli phone analytics data --channel 0 --account "tk_acc" --data-date 1764137986
  geelark-cli phone analytics data --created-id "user_id" --channel 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"page":     page,
				"pageSize": pageSize,
			}
			if account != "" {
				body["account"] = account
			}
			if dataDateSet {
				body["dataDate"] = dataDate
			}
			if createdId != "" {
				body["createdId"] = createdId
			}
			if channelSet {
				body["channel"] = channel
			}
			result, err := c.PostAndPrint("/open/v1/analytics/data", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Page size (1-100)")
	cmd.Flags().StringVar(&account, "account", "", "Account name filter")
	cmd.Flags().IntVar(&dataDate, "data-date", 0, "Search date timestamp (seconds)")
	cmd.Flags().StringVar(&createdId, "created-id", "", "User ID who added the account")
	cmd.Flags().IntVar(&channel, "channel", -1, "Platform: 0=TikTok, 1=YouTube, 2=Instagram, 4=Reddit (omit for all)")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		channelSet = cmd.Flags().Changed("channel")
		dataDateSet = cmd.Flags().Changed("data-date")
		return nil
	}

	return cmd
}

func newAnalyticsTagsListCmd(newClient clientFactory) *cobra.Command {
	var name string
	var page, pageSize int

	cmd := &cobra.Command{
		Use:   "tags-list",
		Short: "List analytics tags",
		Example: `  geelark-cli phone analytics tags-list --page 1 --page-size 200
  geelark-cli phone analytics tags-list --name "Important"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"page":     page,
				"pageSize": pageSize,
			}
			if name != "" {
				body["name"] = name
			}
			result, err := c.PostAndPrint("/open/v1/analytics/tags/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Tag name (fuzzy search)")
	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 200, "Page size (max 200)")

	return cmd
}

func newAnalyticsTagsCreateCmd(newClient clientFactory) *cobra.Command {
	var name string
	var color int

	cmd := &cobra.Command{
		Use:     "tags-create",
		Short:   "Create an analytics tag",
		Example: `  geelark-cli phone analytics tags-create --name "Important Account" --color 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"name": name,
			}
			if cmd.Flags().Changed("color") {
				body["color"] = color
			}
			result, err := c.PostAndPrint("/open/v1/analytics/tags/create", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Tag name, max 100 chars, must be unique (required)")
	cmd.Flags().IntVar(&color, "color", 0, "Tag color value (default 0)")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func newAnalyticsTagsUpdateCmd(newClient clientFactory) *cobra.Command {
	var id, name string
	var color int

	cmd := &cobra.Command{
		Use:     "tags-update",
		Short:   "Update an analytics tag",
		Example: `  geelark-cli phone analytics tags-update --id "tag_id" --name "Core Account" --color 2`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"id":   id,
				"name": name,
			}
			if cmd.Flags().Changed("color") {
				body["color"] = color
			}
			result, err := c.PostAndPrint("/open/v1/analytics/tags/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Tag ID (required)")
	cmd.Flags().StringVar(&name, "name", "", "Tag name, max 100 chars (required)")
	cmd.Flags().IntVar(&color, "color", 0, "Tag color value (default 0)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func newAnalyticsTagsDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "tags-delete",
		Short:   "Delete analytics tags",
		Example: `  geelark-cli phone analytics tags-delete --ids "tag1,tag2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}
			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}
			result, err := c.PostAndPrint("/open/v1/analytics/tags/del", body)
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
