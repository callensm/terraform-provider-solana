package solana

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSlotLeaders() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://solana.com/docs/rpc/http/getslotleaders) Get the list of current slot leader addresses.",

		Read: dataSourceSlotLeadersRead,

		Schema: map[string]*schema.Schema{
			"limit": {
				Type:         schema.TypeInt,
				Description:  "The number of slot leaders to get from the starting slot (between 1 - 5,000).",
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 5000),
			},
			"start_slot": {
				Type:         schema.TypeInt,
				Description:  "The starting slot number for the query.",
				Required:     true,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"leaders": {
				Type:        schema.TypeList,
				Description: "The list of slot leader addresses found.",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceSlotLeadersRead(d *schema.ResourceData, meta any) error {
	client := meta.(*providerConfig).rpcClient

	limit := d.Get("limit").(int)
	startSlot := d.Get("start_slot").(int)

	res, err := client.GetSlotLeaders(context.Background(), uint64(startSlot), uint64(limit))
	if err != nil {
		return err
	}

	var addrs []string
	for _, pk := range res {
		addrs = append(addrs, pk.String())
	}

	d.SetId(fmt.Sprintf("slot_leaders:%d:%d", startSlot, limit))
	d.Set("leaders", addrs)
	return nil
}
