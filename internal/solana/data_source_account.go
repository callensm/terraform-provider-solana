package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Description: "Provides all information associated with the account of the provided public key",

		Read: dataSourceAccountRead,

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:        schema.TypeString,
				Description: "Base-58 encoded public key of the account to query",
				Required:    true,
			},
			"lamports": {
				Type:        schema.TypeInt,
				Description: "Number of lamports assigned to the account",
				Computed:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "Base-58 encoded public key of the program the account is assigned to",
				Computed:    true,
			},
			"executable": {
				Type:        schema.TypeBool,
				Description: "Indicates whether the account contains a program",
				Computed:    true,
			},
			"rent_epoch": {
				Type:        schema.TypeInt,
				Description: "The epoch at which this account will next owe rent",
				Computed:    true,
			},
		},
	}
}

func dataSourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	pub := d.Get("public_key").(string)
	address, err := solana.PublicKeyFromBase58(pub)
	if err != nil {
		return err
	}

	res, err := client.GetAccountInfo(context.Background(), address)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("account:%s", pub))
	d.Set("lamports", uint64(res.Value.Lamports))
	d.Set("owner", res.Value.Owner.String())
	d.Set("executable", res.Value.Executable)
	d.Set("rent_epoch", uint64(res.Value.RentEpoch))
	return nil
}
