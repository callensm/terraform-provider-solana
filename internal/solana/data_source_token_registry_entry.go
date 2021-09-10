package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/tokenregistry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTokenRegistryEntry() *schema.Resource {
	return &schema.Resource{
		Description: "Provides the metadata for a Solana token registry entry based on the given minting address.",

		Read: dataSourceTokenRegistryEntryRead,

		Schema: map[string]*schema.Schema{
			"mint_address": {
				Type:        schema.TypeString,
				Description: "The public key of the mint account address as a base-58 string.",
				Required:    true,
			},
			"initialized": {
				Type:        schema.TypeBool,
				Description: "Flag that indicates in the token entry has been initialized.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the token.",
				Computed:    true,
			},
			"authority": {
				Type:        schema.TypeString,
				Description: "The public key of the registration authority as a base-58 string.",
				Computed:    true,
			},
		},
	}
}

func dataSourceTokenRegistryEntryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	mintAddressKey, err := solana.PublicKeyFromBase58(d.Get("mint_address").(string))
	if err != nil {
		return err
	}

	entry, err := tokenregistry.GetTokenRegistryEntry(context.Background(), client, mintAddressKey)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("token_registry_entry:%s", mintAddressKey.String()))
	d.Set("initialized", entry.IsInitialized)
	d.Set("name", entry.Name.String())
	d.Set("authority", entry.RegistrationAuthority.String())
	return nil
}
