package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBalance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBalanceRead,

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:        schema.TypeString,
				Description: "Base-58 encoded public key of the account to query",
				Required:    true,
			},
			"balance": {
				Type:        schema.TypeInt, // FIXME:
				Description: "The balance of the queried account",
				Computed:    true,
			},
		},
	}
}

func dataSourceBalanceRead(d *schema.ResourceData, meta interface{}) error {
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
