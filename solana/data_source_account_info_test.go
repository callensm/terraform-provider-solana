package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testAccountInfoDataConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_account_info" "test" {
			public_key = "11111111111111111111111111111111"
		}
	`
)

func TestAccAccountInfoDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccountInfoDataConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccountInfoDataSucceeds("data.solana_account_info.test"),
				),
			},
		},
	})
}

func testAccountInfoDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Account Info Failure: %s not found", name)
		}

		if owner := val.Primary.Attributes["owner"]; owner != "NativeLoader1111111111111111111111111111111" {
			return fmt.Errorf("Account Info Failure: unexpected owner received (%s)", owner)
		}

		return nil
	}
}
