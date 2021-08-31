package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testNodeIdentityConfig = `
		provider "solana" {
			cluster = "testnet"
		}

		data "solana_node_identity" "test" {}
	`
)

func TestAccNodeIdentityDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testNodeIdentityConfig,
				Check: resource.ComposeTestCheckFunc(
					testNodeIdentityDataSucceeds("data.solana_node_identity.test"),
				),
			},
		},
	})
}

func testNodeIdentityDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Node Identity Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Node Identity Failure: ID was not set")
		}

		if val.Primary.Attributes["public_key"] == "" {
			return fmt.Errorf("Node Identity Failure: node public_key was not properly outputted")
		}

		return nil
	}
}
