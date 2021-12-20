package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRentExemptionCost() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getminimumbalanceforrentexemption) Provides the minimum balance required to make an account rent exempt in lamports.",

		Read: dataSourceRentExemptionCostRead,

		Schema: map[string]*schema.Schema{
			"data_length": {
				Type:        schema.TypeInt,
				Description: "The length of the data stored in the account.",
				Required:    true,
			},
			"lamports": {
				Type:        schema.TypeInt,
				Description: "The calculated minimum cost of rent exemption for the account size.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRentExemptionCostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	dataLength := d.Get("data_length").(uint64)
	res, err := client.GetMinimumBalanceForRentExemption(context.Background(), dataLength, rpc.CommitmentRecent)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("rent_exemption_cost:%d", dataLength))
	d.Set("lamports", res)
	return nil
}
