package solana

import (
	"context"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRecentPrioritizationFees() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://solana.com/docs/rpc/http/gethealth) Get recent prioritization fees for the network",

		Read: dataSourceRecentPrioritizationFeesRead,

		Schema: map[string]*schema.Schema{
			"accounts": {
				Type:        schema.TypeList,
				Description: "An array of account addresses as base-58 encoded strings.",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"recent_fees": {
				Type:        schema.TypeList,
				Description: "The results of the recent priority fee query.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fee": {
							Type:        schema.TypeInt,
							Description: "The fee of the recent transaction(s).",
							Computed:    true,
						},
						"slot": {
							Type:        schema.TypeInt,
							Description: "The slot of the provided priority fee.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRecentPrioritizationFeesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*providerConfig).rpcClient

	accounts, ok := d.GetOk("accounts")

	var addrs []solana.PublicKey
	var accountsListString string

	if ok {
		for _, acc := range accounts.([]string) {
			pk, err := solana.PublicKeyFromBase58(acc)
			if err != nil {
				return fmt.Errorf("RecentPrioritizationFees Failure: parsing accounts input %s", err)
			}

			addrs = append(addrs, pk)
			accountsListString += fmt.Sprintf("%s,", acc)
		}
	} else {
		accountsListString = "nil"
	}

	res, err := client.GetRecentPrioritizationFees(context.Background(), addrs)
	if err != nil {
		return err
	}

	accountsListString = strings.TrimRight(accountsListString, ",")

	var fees []map[string]any
	for _, r := range res {
		fees = append(fees, map[string]any{
			"fee":  r.PrioritizationFee,
			"slot": r.Slot,
		})
	}

	d.SetId(fmt.Sprintf("recent_prioritization_fees:%s", accountsListString))
	d.Set("recent_fees", fees)
	return nil
}
