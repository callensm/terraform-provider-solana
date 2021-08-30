package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testAddressSignaturesConfig = `
		provider "solana" {
			cluster = "testnet"
		}

		data "solana_address_signatures" "sigs" {
			address = "11111111111111111111111111111111"

            search_options {
                limit = 5
            }
		}
	`
)

func TestAccAddressSignaturesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAddressSignaturesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAddressSignaturesDataSucceeds("data.solana_address_signatures.sigs"),
				),
			},
		},
	})
}

func testAddressSignaturesDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Address Signatures Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Address Signatures Failure: ID was not set")
		}

		return nil
	}
}
