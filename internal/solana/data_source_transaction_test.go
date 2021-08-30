package solana

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	// FIXME:
	testTransactionConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
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
					testSupplySucceeds("data.solana_transaction.tx"),
				),
			},
		},
	})
}

func testTransactionSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// TODO:
		return nil
	}
}
