package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testRecentBlockhashDataConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_recent_blockhash" "test" {}
	`
)

func TestAccRecentBlockhashDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testRecentBlockhashDataConfig,
				Check: resource.ComposeTestCheckFunc(
					testRecentBlockhashDataSucceeds("data.solana_recent_blockhash.test"),
				),
			},
		},
	})
}

func testRecentBlockhashDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Recent Blockhash Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Recent Blockhash Failure: ID was not set")
		}

		return nil
	}
}
