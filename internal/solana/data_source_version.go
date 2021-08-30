package solana

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVersion() *schema.Resource {
	return &schema.Resource{
		Description: "Provides the current Solana software version running on the network node",

		Read: dataSourceVersionRead,

		Schema: map[string]*schema.Schema{
			"solana_core": {
				Type:        schema.TypeString,
				Description: "Software version of `solana-core` running on the node",
				Computed:    true,
			},
			"feature_set": {
				Type:        schema.TypeInt,
				Description: "A unique identifier of the current software's feature set",
				Computed:    true,
			},
		},
	}
}

func dataSourceVersionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	res, err := client.GetVersion(context.Background())
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("version:%s", res.SolanaCore))
	d.Set("solana_core", res.SolanaCore)
	d.Set("feature_set", int(res.FeatureSet))
	return nil
}
