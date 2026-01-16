package solana

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testRecentPrioritizationFeesConfig = `
        provider "solana" {
            endpoint = "https://api.testnet.solana.com"
        }

        data "solana_recent_prioritization_fees" "test" {
            # accounts = []
        }
    `
)

func TestAccRecentPrioritizationFeesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testRecentPrioritizationFeesConfig,
				Check: resource.ComposeTestCheckFunc(
					testRecentPrioritizationFeesSucceeds("data.solana_recent_prioritization_fees.test"),
				),
			},
		},
	})
}

func testRecentPrioritizationFeesSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("RecentPrioritizationFees Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("RecentPrioritizationFees Failure: ID was not set")
		}

		amount := val.Primary.Attributes["recent_fees.#"]
		if amt, _ := strconv.Atoi(amount); amt == 0 {
			return fmt.Errorf("RecentPrioritizationFees Failure: empty recent fees results")
		}

		return nil
	}
}
