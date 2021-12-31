package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testTokenRegistryEntryConfig = `
		provider "solana" {
			cluster = "mainnet-beta"
		}

		data "solana_token_registry_entry" "entry" {
            mint_address = "SRMuApVNdxXokk5GT7XD5cUUgXMBCoAz2LHeuAoKWRt"
        }
	`
)

func TestAccTokenRegistryEntryDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testTokenRegistryEntryConfig,
				Check: resource.ComposeTestCheckFunc(
					testTokenRegistryEntrySucceeds("data.solana_token_registry_entry.entry"),
				),
			},
		},
	})
}

func testTokenRegistryEntrySucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Version Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Version Failure: ID was not set")
		}

		return nil
	}
}
