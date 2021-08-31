package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceTokenSupply() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#gettokensupply) Provides details about the total supply of an SPL Token type.",

		Read: dataSourceTokenSupplyRead,

		Schema: map[string]*schema.Schema{
			"minter_public_key": {
				Type:         schema.TypeString,
				Description:  "Public key of the token Mint to query as a base-58 encoded string.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(32, 44),
			},
			"amount": {
				Type:        schema.TypeString,
				Description: "The raw total token supply without decimals.",
				Computed:    true,
			},
			"decimals": {
				Type:        schema.TypeInt,
				Description: "Number of base-10 digits to the right of the decimal place.",
				Computed:    true,
			},
			"ui_amount": {
				Type:        schema.TypeString,
				Description: "The total token supply using mint-prescribed decimals",
				Computed:    true,
			},
		},
	}
}

func dataSourceTokenSupplyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	pk, err := solana.PublicKeyFromBase58(d.Get("minter_public_key").(string))
	if err != nil {
		return err
	}

	res, err := client.GetTokenSupply(context.Background(), pk, rpc.CommitmentType(rpc.ConfirmationStatusFinalized))
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("token_supply:%s", pk.String()))
	d.Set("amount", res.Value.Amount)
	d.Set("decimals", res.Value.Decimals)
	d.Set("ui_amount", res.Value.UiAmountString)
	return nil
}
