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
            signature = "MowmH2w4va9Q4v5Bo3JummYQQXdVsqThgeSzx3tK7edNrJXsAy2JwbzDLno3s8NKWVUvsWu3b7wuaUPLLuDdyTV"
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
