package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSupply() *schema.Resource {
	return &schema.Resource{
		Description: "Provides information about the current supply on the network",
		Read:        dataSourceSupplyRead,
		Schema: map[string]*schema.Schema{
			"total": {
				Type:        schema.TypeInt,
				Description: "Total supply of lamports",
				Computed:    true,
			},
			"circulating": {
				Type:        schema.TypeInt,
				Description: "Circulating supply in lamports",
				Computed:    true,
			},
			"non_circulating": {
				Type:        schema.TypeInt,
				Description: "Non-circulating supply in lamports",
				Computed:    true,
			},
		},
	}
}

func dataSourceSupplyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	res, err := client.GetSupply(context.Background(), rpc.CommitmentRecent)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("supply:%d", res.Value.Total))
	d.Set("total", res.Value.Total)
	d.Set("circulating", res.Value.Circulating)
	d.Set("non_circulating", res.Value.NonCirculating)
	return nil
}
