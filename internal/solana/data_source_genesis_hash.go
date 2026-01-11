package solana

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGenesisHash() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://solana.com/docs/rpc/http/getgenesishash) Get the genesis hash of the configured RPC endpoint.",

		Read: dataSourceGenesisHashRead,

		Schema: map[string]*schema.Schema{
			"block_hash": {
				Type:        schema.TypeString,
				Description: "The genesis block hash.",
				Computed:    true,
			},
		},
	}
}

func dataSourceGenesisHashRead(d *schema.ResourceData, meta any) error {
	client := meta.(*providerConfig).rpcClient

	res, err := client.GetGenesisHash(context.Background())
	if err != nil {
		return err
	}

	d.SetId("genesis_hash")
	d.Set("block_hash", res.String())
	return nil
}
