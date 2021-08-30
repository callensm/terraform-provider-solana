package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testBalanceDataConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_balance" "test" {
			public_key = "11111111111111111111111111111111"
		}
	`
)

func TestAccBalanceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testBalanceDataConfig,
				Check: resource.ComposeTestCheckFunc(
					testBalanceDataSucceeds("data.solana_balance.test"),
				),
			},
		},
	})
}

func testBalanceDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Balance Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Balance Failure: ID was not set")
		}

		if bal := val.Primary.Attributes["balance"]; bal != "1" {
			return fmt.Errorf("Balance Failure: incorrect testnet balance received (%s)", bal)
		}

		return nil
	}
}
