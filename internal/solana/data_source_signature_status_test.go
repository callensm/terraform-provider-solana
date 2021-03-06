package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testSignatureStatusConfig = `
		provider "solana" {
			cluster = "testnet"
		}

		data "solana_signature_status" "test" {
            signature = "5Pe5B4239HixXdHfzjj7ffqHVN7rGnmEfbGzLZhyycinfwb3LciRuaxPUiyRJFQuKo4h8yg65es8L7B2dMZcT3xC"
            search_transaction_history = true
        }
	`
)

func TestAccSignatureStatusDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testSignatureStatusConfig,
				Check: resource.ComposeTestCheckFunc(
					testSignatureStatusSucceeds("data.solana_signature_status.test"),
				),
			},
		},
	})
}

func testSignatureStatusSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Signature Status Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Signature Status Failure: ID was not set")
		}

		return nil
	}
}
