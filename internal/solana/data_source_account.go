package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getaccountinfo) Provides all information associated with the account of the provided public key",

		Read: dataSourceAccountRead,

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:         schema.TypeString,
				Description:  "Base-58 encoded public key of the account to query.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(32, 44),
			},
			"encoding": {
				Type:         schema.TypeString,
				Description:  "Desired encoding for returned transaction data (`json`, `jsonParsed`, `base58`, `base64`). Defaults to `base64`.",
				Optional:     true,
				Default:      solana.EncodingBase64,
				ValidateFunc: validation.StringInSlice(dataEncodingOptions, false),
			},
			"data": {
				Type:        schema.TypeString,
				Description: "Data associated with the account, either as encoded binary data or JSON depending on the value of `encoding`.",
				Computed:    true,
			},
			"lamports": {
				Type:        schema.TypeInt,
				Description: "Number of lamports assigned to the account.",
				Computed:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "Base-58 encoded public key of the program the account is assigned to.",
				Computed:    true,
			},
			"executable": {
				Type:        schema.TypeBool,
				Description: "Indicates whether the account contains a program.",
				Computed:    true,
			},
			"rent_epoch": {
				Type:        schema.TypeInt,
				Description: "The epoch at which this account will next owe rent.",
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

	encoding := solana.EncodingType(d.Get("encoding").(string))
	res, err := client.GetAccountInfoWithOpts(context.Background(), address, &rpc.GetAccountInfoOpts{
		Encoding: encoding,
	})
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("account:%s", pub))
	d.Set("lamports", uint64(res.Value.Lamports))
	d.Set("owner", res.Value.Owner.String())
	d.Set("executable", res.Value.Executable)
	d.Set("rent_epoch", uint64(res.Value.RentEpoch))

	if encoding == solana.EncodingJSONParsed {
		d.Set("data", string(res.Value.Data.GetRawJSON()))
	} else {
		d.Set("data", string(res.Value.Data.GetBinary()))
	}

	return nil
}
