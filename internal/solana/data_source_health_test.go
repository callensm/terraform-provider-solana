package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testHealthConfig = `
        provider "solana" {
            endpoint = "https://api.testnet.solana.com"
        }

        data "solana_health" "test" {}
    `
)

func TestAccHealthDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testHealthConfig,
				Check: resource.ComposeTestCheckFunc(
					testHealthDataSucceeds("data.solana_health.test"),
				),
			},
		},
	})
}

func testHealthDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Health Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Health Failure: ID was not set")
		}

		if healthy := val.Primary.Attributes["healthy"]; healthy != "true" {
			return fmt.Errorf("Health Failure: expected healthy to be true")
		}

		return nil
	}
}
