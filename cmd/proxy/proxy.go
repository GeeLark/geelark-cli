package proxy

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

// NewCmd creates the proxy command group.
func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Proxy management",
		Long:  "Manage proxies — list, add, update, delete, and check proxies.",
	}

	cmd.AddCommand(newListCmd(newClient))
	cmd.AddCommand(newAddCmd(newClient))
	cmd.AddCommand(newSimpleAddCmd(newClient))
	cmd.AddCommand(newUpdateCmd(newClient))
	cmd.AddCommand(newSimpleUpdateCmd(newClient))
	cmd.AddCommand(newDeleteCmd(newClient))
	cmd.AddCommand(newCheckCmd(newClient))

	return cmd
}

func newListCmd(newClient clientFactory) *cobra.Command {
	var page, pageSize int
	var ids string

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List all proxies",
		Long:    "Get all proxies with pagination.",
		Example: `  geelark-cli proxy list --page 1 --page-size 10`,
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

			result, err := c.PostAndPrint("/open/v1/proxy/list", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&page, "page", 1, "Page number")
	cmd.Flags().IntVar(&pageSize, "page-size", 10, "Number of items per page (max 100)")
	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated proxy IDs")

	return cmd
}

func newAddCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add proxies (batch)",
		Long: `Add one or more proxies in batch via JSON. Duplicate proxies will not be added.
For a single proxy, prefer 'simple-add' which uses flat flags.

proxyQueryChannel: 1=IPApi, 2=IP2Location (default 2)`,
		Example: `  geelark-cli proxy add --data "[{\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\",\"proxyQueryChannel\":2}]"
  geelark-cli proxy add --data "[{\"scheme\":\"http\",\"server\":\"1.2.3.4\",\"port\":8080,\"proxyQueryChannel\":1},{\"scheme\":\"socks5\",\"server\":\"5.6.7.8\",\"port\":1080,\"proxyQueryChannel\":2}]"`,
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

			result, err := c.PostAndPrint("/open/v1/proxy/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON array of proxy items [{scheme,server,port,username,password,proxyQueryChannel}] (required)")
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newSimpleAddCmd(newClient clientFactory) *cobra.Command {
	var scheme, server, username, password string
	var port, proxyQueryChannel int
	var pqcSet bool

	cmd := &cobra.Command{
		Use:   "simple-add",
		Short: "Quick add a single proxy",
		Long: `Simplified command to add a single proxy with flat flags.
Use 'add' for batch creation.`,
		Example: `  geelark-cli proxy simple-add --scheme socks5 --server 192.3.8.1 --port 8000
  geelark-cli proxy simple-add --scheme socks5 --server 192.3.8.1 --port 8000 --username admin --password admin
  geelark-cli proxy simple-add --scheme http --server 1.2.3.4 --port 8080 --proxy-query-channel 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			item := map[string]interface{}{
				"scheme": scheme,
				"server": server,
				"port":   port,
			}
			if username != "" {
				item["username"] = username
			}
			if password != "" {
				item["password"] = password
			}
			if pqcSet {
				item["proxyQueryChannel"] = proxyQueryChannel
			}

			body := map[string]interface{}{
				"list": []interface{}{item},
			}

			result, err := c.PostAndPrint("/open/v1/proxy/add", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&scheme, "scheme", "socks5", "Proxy type: socks5 (default), http, https")
	cmd.Flags().StringVar(&server, "server", "", "Proxy server address (required)")
	cmd.Flags().IntVar(&port, "port", 0, "Proxy port (required)")
	cmd.Flags().StringVar(&username, "username", "", "Proxy username")
	cmd.Flags().StringVar(&password, "password", "", "Proxy password")
	cmd.Flags().IntVar(&proxyQueryChannel, "proxy-query-channel", 2, "Detection channel: 1=IPApi, 2=IP2Location")
	_ = cmd.MarkFlagRequired("server")
	_ = cmd.MarkFlagRequired("port")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		pqcSet = cmd.Flags().Changed("proxy-query-channel")
		return nil
	}

	return cmd
}

func newUpdateCmd(newClient clientFactory) *cobra.Command {
	var dataJSON string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update proxies (batch)",
		Long: `Update one or more proxies in batch via JSON.
For a single proxy, prefer 'simple-update' which uses flat flags.

proxyQueryChannel: 1=IPApi, 2=IP2Location (default keeps original)`,
		Example: `  geelark-cli proxy update --data "[{\"id\":\"proxy_id\",\"scheme\":\"socks5\",\"server\":\"192.3.8.1\",\"port\":8000,\"username\":\"admin\",\"password\":\"admin\",\"proxyQueryChannel\":2}]"
  geelark-cli proxy update --data "[{\"id\":\"id1\",\"scheme\":\"http\",\"server\":\"1.2.3.4\",\"port\":8080,\"proxyQueryChannel\":1}]"`,
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

			result, err := c.PostAndPrint("/open/v1/proxy/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&dataJSON, "data", "", "JSON array of proxy update items [{id,scheme,server,port,username,password,proxyQueryChannel}] (required)")
	_ = cmd.MarkFlagRequired("data")

	return cmd
}

func newSimpleUpdateCmd(newClient clientFactory) *cobra.Command {
	var id, scheme, server, username, password string
	var port, proxyQueryChannel int
	var pqcSet bool

	cmd := &cobra.Command{
		Use:   "simple-update",
		Short: "Quick update a single proxy",
		Long: `Simplified command to update a single proxy with flat flags.
Use 'update' for batch updates.`,
		Example: `  geelark-cli proxy simple-update --id "proxy_id" --scheme socks5 --server 192.3.8.1 --port 8000
  geelark-cli proxy simple-update --id "proxy_id" --scheme socks5 --server 192.3.8.1 --port 8000 --username admin --password admin --proxy-query-channel 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			item := map[string]interface{}{
				"id":     id,
				"scheme": scheme,
				"server": server,
				"port":   port,
			}
			if username != "" {
				item["username"] = username
			}
			if password != "" {
				item["password"] = password
			}
			if pqcSet {
				item["proxyQueryChannel"] = proxyQueryChannel
			}

			body := map[string]interface{}{
				"list": []interface{}{item},
			}

			result, err := c.PostAndPrint("/open/v1/proxy/update", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Proxy ID (required)")
	cmd.Flags().StringVar(&scheme, "scheme", "socks5", "Proxy type: socks5 (default), http, https")
	cmd.Flags().StringVar(&server, "server", "", "Proxy server address (required)")
	cmd.Flags().IntVar(&port, "port", 0, "Proxy port (required)")
	cmd.Flags().StringVar(&username, "username", "", "Proxy username")
	cmd.Flags().StringVar(&password, "password", "", "Proxy password")
	cmd.Flags().IntVar(&proxyQueryChannel, "proxy-query-channel", 2, "Detection channel: 1=IPApi, 2=IP2Location")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("server")
	_ = cmd.MarkFlagRequired("port")

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		pqcSet = cmd.Flags().Changed("proxy-query-channel")
		return nil
	}

	return cmd
}

func newDeleteCmd(newClient clientFactory) *cobra.Command {
	var ids string

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete proxies",
		Long:    "Delete proxies by IDs.",
		Example: `  geelark-cli proxy delete --ids "id1,id2"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"ids": strings.Split(ids, ","),
			}

			result, err := c.PostAndPrint("/open/v1/proxy/delete", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&ids, "ids", "", "Comma-separated proxy IDs (required)")
	_ = cmd.MarkFlagRequired("ids")

	return cmd
}

func newCheckCmd(newClient clientFactory) *cobra.Command {
	var proxyType, server, username, password, channel string
	var port int

	cmd := &cobra.Command{
		Use:     "check",
		Short:   "Check/detect a proxy",
		Long:    "Detect a proxy to verify its connectivity and get outbound IP info.",
		Example: `  geelark-cli proxy check --type socks5 --server 185.162.130.86 --port 10000 --channel IP2Location`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"proxyQueryChannel": channel,
				"proxyType":         proxyType,
				"server":            server,
				"port":              port,
			}
			if username != "" {
				body["username"] = username
			}
			if password != "" {
				body["password"] = password
			}

			result, err := c.PostAndPrint("/open/v1/proxy/check", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&channel, "channel", "IP2Location", "IP lookup source: IP-API or IP2Location")
	cmd.Flags().StringVar(&proxyType, "type", "socks5", "Proxy type: socks5, http, https")
	cmd.Flags().StringVar(&server, "server", "", "Proxy server address (required)")
	cmd.Flags().IntVar(&port, "port", 0, "Proxy port (required)")
	cmd.Flags().StringVar(&username, "username", "", "Proxy username")
	cmd.Flags().StringVar(&password, "password", "", "Proxy password")
	_ = cmd.MarkFlagRequired("server")
	_ = cmd.MarkFlagRequired("port")

	return cmd
}
