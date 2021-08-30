package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRecentBlockhash() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieves a recent block hash from the ledger and the associated cost in lamports per signature on a new transaction for that block",

		Read: dataSourceRecentBlockhashRead,

		Schema: map[string]*schema.Schema{
			"blockhash": {
				Type:        schema.TypeString,
				Description: "Base-58 encoded hash string of the block",
				Computed:    true,
			},
			"lamports_per_signature": {
				Type:        schema.TypeInt,
				Description: "The lamports cost per signature of the block",
				Computed:    true,
			},
		},
	}
}

func dataSourceRecentBlockhashRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	res, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentRecent)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("recent_blockhash:%s", res.Value.Blockhash.String()))
	d.Set("blockhash", res.Value.Blockhash.String())
	d.Set("lamports_per_signature", uint64(res.Value.FeeCalculator.LamportsPerSignature))
	return nil
}
