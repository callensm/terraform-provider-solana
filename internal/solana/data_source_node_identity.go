package solana

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodeIdentity() *schema.Resource {
	return &schema.Resource{
		Description: "[(JSON RPC)](https://docs.solana.com/developing/clients/jsonrpc-api#getidentity) Provides the identity public key for the current node.",

		Read: dataSourceNodeIdentityRead,

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:        schema.TypeString,
				Description: "The identity public key of the current node as a base-58 encoded string.",
				Computed:    true,
			},
		},
	}
}

func dataSourceNodeIdentityRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	res, err := client.GetIdentity(context.Background())
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("node_identity:%s", res.Identity.String()))
	d.Set("public_key", res.Identity.String())
	return nil
}
