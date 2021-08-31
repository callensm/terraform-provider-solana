package solana

import (
	"context"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSignatureStatus() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getsignaturestatuses) Provides the status for the inputted signature.",

		Read: dataSourceSignatureStatusRead,

		Schema: map[string]*schema.Schema{
			"signature": {
				Type:        schema.TypeString,
				Description: "The transaction signature to confirm status as a base-58 encoded string.",
				Required:    true,
			},
			"search_transaction_history": {
				Type:        schema.TypeBool,
				Description: "If `true`, a Solana node will search its ledger cache for the signature if not found in the recent status cache. Defaults to `false`.",
				Optional:    true,
				Default:     false,
			},
			"slot": {
				Type:        schema.TypeInt,
				Description: "The slot the transaction was processed.",
				Computed:    true,
			},
			"confirmations": {
				Type:        schema.TypeInt,
				Description: "Number of blocks since the signature confirmation and finalized by a supermajority of the cluster. If `-1` then it is the root.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "The transaction's cluster confirmation status (`processed`, `confirmed`, or `finalized`).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSignatureStatusRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	sig, err := solana.SignatureFromBase58(d.Get("signature").(string))
	if err != nil {
		return err
	}

	res, err := client.GetSignatureStatuses(context.Background(), d.Get("search_transaction_history").(bool), sig)
	if err != nil {
		return err
	}

	if res == nil || len(res.Value) == 0 || res.Value[0] == nil {
		return errors.New("no status results found for inputted signature")
	}

	d.SetId(fmt.Sprintf("signature_status:%s", sig.String()))
	d.Set("slot", uint64(res.Value[0].Slot))

	if confs := res.Value[0].Confirmations; confs != nil {
		d.Set("confirmations", uint64(*confs))
	} else {
		d.Set("confirmations", -1)
	}

	d.Set("status", string(res.Value[0].ConfirmationStatus))

	return nil
}
