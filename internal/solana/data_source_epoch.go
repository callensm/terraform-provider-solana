package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpoch() *schema.Resource {
	return &schema.Resource{
		Description: "Provides all of the relevant information about the current epoch",
		Read:        dataSourceEpochRead,
		Schema: map[string]*schema.Schema{
			"absolute_slot": {
				Type:        schema.TypeInt,
				Description: "The current absolute slot in the epoch",
				Computed:    true,
			},
			"block_height": {
				Type:        schema.TypeInt,
				Description: "The current block height",
				Computed:    true,
			},
			"epoch": {
				Type:        schema.TypeInt,
				Description: "The current epoch count",
				Computed:    true,
			},
			"slot_index": {
				Type:        schema.TypeInt,
				Description: "The current slot relative to the start of the current epoch",
				Computed:    true,
			},
			"slots_in_epoch": {
				Type:        schema.TypeInt,
				Description: "The number of slots in this epoch",
				Computed:    true,
			},
		},
	}
}

func dataSourceEpochRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	res, err := client.GetEpochInfo(context.Background(), rpc.CommitmentRecent)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("epoch:%d", uint64(res.Epoch)))
	d.Set("absolute_slot", uint64(res.AbsoluteSlot))
	d.Set("block_height", uint64(res.TransactionCount))
	d.Set("epoch", uint64(res.Epoch))
	d.Set("slot_index", uint64(res.SlotIndex))
	d.Set("slots_in_epoch", uint64(res.SlotsInEpoch))
	return nil
}
