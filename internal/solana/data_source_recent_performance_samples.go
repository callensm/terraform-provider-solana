package solana

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRecentPerformanceSamples() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://solana.com/docs/rpc/http/getrecentperformancesamples) Get a list of recent performance samples, in reverse slot order.",

		Read: dataSourceRecentPerformanceSamplesRead,

		Schema: map[string]*schema.Schema{
			"amount": {
				Type:         schema.TypeInt,
				Description:  "The amount of samples to return (default: 2, max: 720).",
				Optional:     true,
				Default:      2,
				ForceNew:     true,
				ValidateFunc: validation.IntBetween(1, 720),
			},
			"samples": {
				Type:        schema.TypeList,
				Description: "The list of performance samples.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_transactions": {
							Type:        schema.TypeInt,
							Description: "The number of transactions processed in the period.",
							Computed:    true,
						},
						"num_slots": {
							Type:        schema.TypeInt,
							Description: "The number of slots processed in the period.",
							Computed:    true,
						},
						"sample_period_seconds": {
							Type:        schema.TypeInt,
							Description: "The period in seconds over which this sample was taken.",
							Computed:    true,
						},
						"slot": {
							Type:        schema.TypeInt,
							Description: "The slot of the current sample.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRecentPerformanceSamplesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*providerConfig).rpcClient

	amount := d.Get("amount").(int)
	uAmount := uint(amount)

	res, err := client.GetRecentPerformanceSamples(context.Background(), &uAmount)
	if err != nil {
		return err
	}

	var samples []map[string]any
	for _, s := range res {
		samples = append(samples, map[string]any{
			"num_transactions":      s.NumTransactions,
			"num_slots":             s.NumSlots,
			"sample_period_seconds": s.SamplePeriodSecs,
			"slot":                  s.Slot,
		})
	}

	d.SetId(fmt.Sprintf("recent_performance_samples:%d:%d", amount, time.Now().Unix()))
	d.Set("samples", samples)
	return nil
}
