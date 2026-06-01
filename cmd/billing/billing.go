package billing

import (
	"fmt"

	"github.com/geelark-tech/geelark-cli/internal/client"
	"github.com/spf13/cobra"
)

type clientFactory func() (*client.Client, error)

// NewCmd creates the billing command group.
func NewCmd(newClient clientFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "billing",
		Short: "Billing management",
		Long:  "Query account balance, transaction details, plan info, and manage subscriptions.",
	}

	cmd.AddCommand(newBalanceCmd(newClient))
	cmd.AddCommand(newTransactionDetailCmd(newClient))
	cmd.AddCommand(newPlanListCmd(newClient))
	cmd.AddCommand(newPlanInfoCmd(newClient))
	cmd.AddCommand(newPlanUpgradeCmd(newClient))
	cmd.AddCommand(newPlanRenewCmd(newClient))
	cmd.AddCommand(newBuyTimeAddOnCmd(newClient))

	return cmd
}

func newBalanceCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "balance",
		Short:   "Query account balance",
		Long:    "Query account balance. Rate limit: 10 requests per minute.",
		Example: `  geelark-cli billing balance`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			result, err := c.PostAndPrint("/open/v1/pay/wallet", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newTransactionDetailCmd(newClient clientFactory) *cobra.Command {
	var id, lastFlowId string
	var startAt, endAt, limit int

	cmd := &cobra.Command{
		Use:   "transaction-detail",
		Short: "Query billing transaction details",
		Long: `Query billing transaction details. Supports filtering by cloud phone ID, time range, and pagination.
Only supports searching data within the last 3 days.`,
		Example: `  geelark-cli billing transaction-detail
  geelark-cli billing transaction-detail --id "phone_id" --limit 20
  geelark-cli billing transaction-detail --start-at 1774593838 --end-at 1774593840
  geelark-cli billing transaction-detail --last-flow-id "flow_id" --limit 100`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{}
			if id != "" {
				body["id"] = id
			}
			if startAt > 0 {
				body["startAt"] = startAt
			}
			if endAt > 0 {
				body["endAt"] = endAt
			}
			if limit > 0 {
				body["limit"] = limit
			}
			if lastFlowId != "" {
				body["lastFlowId"] = lastFlowId
			}

			result, err := c.PostAndPrint("/open/v1/billing/transaction/detail", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Cloud phone ID (optional, returns all if not specified)")
	cmd.Flags().IntVar(&startAt, "start-at", 0, "Filter start time, second-level timestamp (last 3 days only)")
	cmd.Flags().IntVar(&endAt, "end-at", 0, "Filter end time, second-level timestamp (last 3 days only)")
	cmd.Flags().IntVar(&limit, "limit", 100, "Number of records to return (default 100, max 1000)")
	cmd.Flags().StringVar(&lastFlowId, "last-flow-id", "", "Last flow ID from previous request for pagination")

	return cmd
}

func newPlanListCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "plan-list",
		Short:   "Get all available plans",
		Long:    "Get the list of all available subscription plans with pricing and limits.",
		Example: `  geelark-cli billing plan-list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			result, err := c.PostAndPrint("/open/v1/pay/profiles/list", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newPlanInfoCmd(newClient clientFactory) *cobra.Command {
	return &cobra.Command{
		Use:     "plan-info",
		Short:   "Get current subscription plan info",
		Long:    "Get the current subscription plan information. Rate limit: 1 request per minute.",
		Example: `  geelark-cli billing plan-info`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			result, err := c.PostAndPrint("/open/v1/pay/plan/info", nil)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func newPlanUpgradeCmd(newClient clientFactory) *cobra.Command {
	var profilesId, promoCode string
	var parallelsNum, monthlyRentalNum, days int

	cmd := &cobra.Command{
		Use:   "plan-upgrade",
		Short: "Upgrade subscription plan",
		Long: `Upgrade the subscription plan. Only supports upgrading (not downgrading).
Use 'plan-list' to get available profilesId values.
Ensure sufficient balance before upgrading.`,
		Example: `  geelark-cli billing plan-upgrade --profiles-id "497540679501610040" --parallels-num 1 --monthly-rental-num 1
  geelark-cli billing plan-upgrade --profiles-id "512719311391949750" --parallels-num 3 --monthly-rental-num 5 --days 30
  geelark-cli billing plan-upgrade --profiles-id "id" --parallels-num 1 --monthly-rental-num 1 --promo-code "GeeLark666"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"profilesId":       profilesId,
				"parallelsNum":     parallelsNum,
				"monthlyRentalNum": monthlyRentalNum,
			}
			if days > 0 {
				body["days"] = days
			}
			if promoCode != "" {
				body["promoCode"] = promoCode
			}

			result, err := c.PostAndPrint("/open/v1/pay/plan/upgrade", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().StringVar(&profilesId, "profiles-id", "", "Plan profiles ID from plan-list (required)")
	cmd.Flags().IntVar(&parallelsNum, "parallels-num", 1, "Parallels number (>= current plan)")
	cmd.Flags().IntVar(&monthlyRentalNum, "monthly-rental-num", 0, "Monthly rental number (>= current plan)")
	cmd.Flags().IntVar(&days, "days", 0, "Renewal duration: 30/90/180/360 days (required when plan expired)")
	cmd.Flags().StringVar(&promoCode, "promo-code", "", "Promo code (optional)")
	_ = cmd.MarkFlagRequired("profiles-id")

	return cmd
}

func newPlanRenewCmd(newClient clientFactory) *cobra.Command {
	var days int
	var promoCode string

	cmd := &cobra.Command{
		Use:   "plan-renew",
		Short: "Renew subscription plan",
		Long: `Renew the current subscription plan.
Ensure sufficient balance before renewing.`,
		Example: `  geelark-cli billing plan-renew --days 30
  geelark-cli billing plan-renew --days 90 --promo-code "GeeLark666"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"days": days,
			}
			if promoCode != "" {
				body["promoCode"] = promoCode
			}

			result, err := c.PostAndPrint("/open/v1/pay/plan/continue", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&days, "days", 30, "Renewal duration: 30, 90, 180, or 360 days (required)")
	cmd.Flags().StringVar(&promoCode, "promo-code", "", "Promo code (optional)")
	_ = cmd.MarkFlagRequired("days")

	return cmd
}

func newBuyTimeAddOnCmd(newClient clientFactory) *cobra.Command {
	var minutes int
	var promoCode string

	cmd := &cobra.Command{
		Use:   "buy-time-addon",
		Short: "Buy time add-on minutes",
		Long: `Buy time add-on minutes. Ensure sufficient balance.
Available minutes: 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000, 1000000, 2000000, 5000000, 10000000`,
		Example: `  geelark-cli billing buy-time-addon --minutes 2000
  geelark-cli billing buy-time-addon --minutes 10000 --promo-code "GeeLark666"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newClient()
			if err != nil {
				return err
			}

			body := map[string]interface{}{
				"minutes": minutes,
			}
			if promoCode != "" {
				body["promoCode"] = promoCode
			}

			result, err := c.PostAndPrint("/open/v1/pay/timeAddOn/buy", body)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.Flags().IntVar(&minutes, "minutes", 0, "Minutes to buy: 2000/5000/10000/20000/50000/100000/200000/500000/1000000/2000000/5000000/10000000 (required)")
	cmd.Flags().StringVar(&promoCode, "promo-code", "", "Promo code (optional)")
	_ = cmd.MarkFlagRequired("minutes")

	return cmd
}
