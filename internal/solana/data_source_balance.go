package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceBalance() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getbalance) Provides the balance of the account of the provided public key.",

		Read: dataSourceBalanceRead,

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:         schema.TypeString,
				Description:  "Base-58 encoded public key of the account to query.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(32, 44),
			},
			"balance": {
				Type:        schema.TypeInt, // FIXME:
				Description: "The balance of the queried account.",
				Computed:    true,
			},
		},
	}
}

func dataSourceBalanceRead(d *schema.ResourceData, meta any) error {
	client := meta.(*providerConfig).rpcClient

	pub := d.Get("public_key").(string)
	address, err := solana.PublicKeyFromBase58(pub)
	if err != nil {
		return err
	}

	res, err := client.GetBalance(context.Background(), address, rpc.CommitmentRecent) // FIXME:
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("balance:%s", pub))
	d.Set("balance", uint64(res.Value)) // FIXME:
	return nil
}
