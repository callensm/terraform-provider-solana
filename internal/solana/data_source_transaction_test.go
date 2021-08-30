package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testTransactionConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

        data "solana_transaction" "tx" {
            signature = "4qoJbgVoPRLKe7bRTD9aMkh9q9spkcyZyZxQyeihGM26uLd7AeyahYhMWjnGwm2BDCsi7a9LuLwQr8iGV2gwATzK"
            encoding  = "json"
        }
	`
)

func TestAccTransactionDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testTransactionConfig,
				Check: resource.ComposeTestCheckFunc(
					testTransactionSucceeds("data.solana_transaction.tx"),
				),
			},
		},
	})
}

func testTransactionSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Transaction Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Transaction Failure: ID was not set")
		}

		return nil
	}
}
