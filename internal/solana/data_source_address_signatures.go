package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceAddressSignatures() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getsignaturesforaddress) Returns confirmed signatures for transactions involving an address backwards in time from the provided signature or most recent confirmed block.",

		Read: dataSourceAddressSignaturesRead,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:         schema.TypeString,
				Description:  "Account address as a base-58 encoded string.",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(32, 44),
			},
			"search_options": {
				Type:        schema.TypeList,
				Description: "Optional parameters to define a custom search time box for the signatures.",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit": {
							Type:         schema.TypeInt,
							Description:  "Maximum transaction signatures to return (between 1 and 1,000).",
							Optional:     true,
							Default:      1000,
							ValidateFunc: validation.IntBetween(1, 1000),
						},
						"before": {
							Type:        schema.TypeString,
							Description: "Start searching backwards from this transaction signature. If not provided the search starts from the top of the highest max confirmed block.",
							Optional:    true,
						},
						"after": {
							Type:        schema.TypeString,
							Description: "Search until this transaction signature or the response limit is reached.",
							Optional:    true,
						},
					},
				},
			},
			"results": {
				Type:        schema.TypeList,
				Description: "A list of transaction signature information, ordered from newest to oldest transaction.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature": {
							Type:        schema.TypeString,
							Description: "Transaction signature as a base-58 encoded string.",
							Computed:    true,
						},
						"slot": {
							Type:        schema.TypeInt,
							Description: "The slot that contains the block with the transaction.",
							Computed:    true,
						},
						"memo": {
							Type:        schema.TypeString,
							Description: "Memo associated with the transaction.",
							Computed:    true,
						},
						"block_time": {
							Type:        schema.TypeInt,
							Description: "Estimated production time as a Unix timestamp of when the transaction was processed.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAddressSignaturesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	rpcOpts := &rpc.GetSignaturesForAddressOpts{}
	if val, ok := d.GetOk("search_options"); ok {
		searchOptions := val.([]interface{})[0].(map[string]interface{})

		if limit, ok := searchOptions["limit"]; ok {
			tmp := limit.(int)
			rpcOpts.Limit = &tmp
		}

		if before, ok := searchOptions["before"]; ok && before.(string) != "" {
			beforeSig, err := solana.SignatureFromBase58(before.(string))
			if err != nil {
				return err
			}

			rpcOpts.Before = beforeSig
		}

		if after, ok := searchOptions["after"]; ok && after.(string) != "" {
			afterSig, err := solana.SignatureFromBase58(after.(string))
			if err != nil {
				return err
			}

			rpcOpts.Before = afterSig
		}
	}

	account, err := solana.PublicKeyFromBase58(d.Get("address").(string))
	if err != nil {
		return err
	}

	res, err := client.GetSignaturesForAddressWithOpts(context.Background(), account, rpcOpts)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("address_signatures:%s", account.String()))

	var results []map[string]interface{}
	for _, sig := range res {
		r := map[string]interface{}{
			"signature":  sig.Signature.String(),
			"slot":       uint64(sig.Slot),
			"block_time": int64(sig.BlockTime),
		}

		if sig.Memo != nil {
			r["memo"] = *sig.Memo
		}

		results = append(results, r)
	}

	d.Set("results", results)
	return nil
}
