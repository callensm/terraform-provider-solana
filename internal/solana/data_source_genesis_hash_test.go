package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testGenesisHashConfig = `
        provider "solana" {
            endpoint = "https://api.testnet.solana.com"
        }

        data "solana_genesis_hash" "test" {}
    `
)

func TestAccGenesisHashDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testGenesisHashConfig,
				Check: resource.ComposeTestCheckFunc(
					testGenesisHashSucceeds("data.solana_genesis_hash.test"),
				),
			},
		},
	})
}

func testGenesisHashSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Genesis Hash Failure: %s was not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Genesis Hash Failure: ID was not set")
		}

		if val.Primary.Attributes["block_hash"] != "4uhcVJyU9pJkvQyS88uRDiswHXSCkY3zQawwpjk2NsNY" {
			return fmt.Errorf("Genesis Hash Failure: incorrect genesis block hash found")
		}

		return nil
	}
}
