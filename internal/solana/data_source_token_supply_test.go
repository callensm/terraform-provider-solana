package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testTokenSupplyConfig = `
		provider "solana" {
			cluster = "testnet"
		}

		data "solana_token_supply" "test" {
            minter_public_key = "F1oxJ4NGgKr6guoYBvZpfEfYoJSEWWmFZjHnQ7TS2vXc"
        }
	`
)

func TestAccTokenSupplyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testTokenSupplyConfig,
				Check: resource.ComposeTestCheckFunc(
					testTokenSupplySucceeds("data.solana_token_supply.test"),
				),
			},
		},
	})
}

func testTokenSupplySucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Token Supply Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Token Supply Failure: ID was not set")
		}

		if val.Primary.Attributes["ui_amount"] == "" {
			return fmt.Errorf("Token Supply Failure: ui_amount was empty")
		}

		return nil
	}
}
