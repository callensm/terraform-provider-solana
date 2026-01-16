package solana

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHealth() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://solana.com/docs/rpc/http/gethealth) Get the health of the Solana network",

		Read: dataSourceHealthRead,

		Schema: map[string]*schema.Schema{
			"healthy": {
				Type:        schema.TypeBool,
				Description: "Whether or not the network is healthy.",
				Computed:    true,
			},
		},
	}
}

func dataSourceHealthRead(d *schema.ResourceData, meta any) error {
	client := meta.(*providerConfig).rpcClient

	status, err := client.GetHealth(context.Background())
	if err != nil {
		return err
	}

	d.SetId("health")
	d.Set("healthy", status == "ok")
	return nil
}
