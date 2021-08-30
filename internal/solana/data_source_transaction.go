package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	transactionDataEncodingOptions = []string{
		string(solana.EncodingBase58),
		string(solana.EncodingBase64),
		string(solana.EncodingJSON),
		string(solana.EncodingJSONParsed),
	}
)

func dataSourceTransaction() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#gettransaction) Provides the details for a confirmed transaction.",

		Read: dataSourceTransactionRead,

		Schema: map[string]*schema.Schema{
			"signature": {
				Type:        schema.TypeString,
				Description: "Transaction signature as a base-58 encoded string.",
				Required:    true,
			},
			"encoding": {
				Type:         schema.TypeString,
				Description:  "Desired encoding for returned transaction data (`json`, `jsonParsed`, `base58`, `base64`). Defaults to `base64`.",
				Optional:     true,
				Default:      solana.EncodingBase64,
				ValidateFunc: validation.StringInSlice(transactionDataEncodingOptions, false),
			},
			"slot": {
				Type:        schema.TypeInt,
				Description: "The slot in which the transaction was processed.",
				Computed:    true,
			},
			"block_time": {
				Type:        schema.TypeInt,
				Description: "The estimated production time as a Unix timestamp of when the transaction was processed.",
				Computed:    true,
			},
		},
	}
}

func dataSourceTransactionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	sig, err := solana.SignatureFromBase58(d.Get("signature").(string))
	if err != nil {
		return err
	}

	res, err := client.GetTransaction(context.Background(), sig, &rpc.GetTransactionOpts{
		Encoding: solana.EncodingType(d.Get("encoding").(string)),
	})

	d.SetId(fmt.Sprintf("transaction:%s_%d", sig.String(), res.Slot))
	d.Set("slot", uint64(res.Slot))
	d.Set("block_time", int64(res.BlockTime))
	return nil
}

func validateTransactionDataEncoding(val interface{}, k string) (warnings []string, errs []error) {
	for _, encoding := range transactionDataEncodingOptions {
		if string(encoding) == val.(string) {
			return
		}
	}

	errs = append(errs, fmt.Errorf("transaction data encoding input (%s) is not a valid option", val.(string)))
	return
}
