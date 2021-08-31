package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testSupplyConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_supply" "test" {}
	`
)

func TestAccSupplyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testSupplyConfig,
				Check: resource.ComposeTestCheckFunc(
					testSupplySucceeds("data.solana_supply.test"),
				),
			},
		},
	})
}

func testSupplySucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Supply Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Supply Failure: ID was not set")
		}

		return nil
	}
}
