package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testAccountDataConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_account" "test" {
			public_key = "11111111111111111111111111111111"
		}
	`
)

func TestAccAccountDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccountDataConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccountDataSucceeds("data.solana_account.test"),
				),
			},
		},
	})
}

func testAccountDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Account Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Account Failure: ID was not set")
		}

		if owner := val.Primary.Attributes["owner"]; owner != "NativeLoader1111111111111111111111111111111" {
			return fmt.Errorf("Account Failure: unexpected owner received (%s)", owner)
		}

		return nil
	}
}
