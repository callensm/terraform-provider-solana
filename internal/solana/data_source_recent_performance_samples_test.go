package solana

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testRecentPerformanceSamplesConfig = `
        provider "solana" {
            endpoint = "https://api.testnet.solana.com"
        }

        data "solana_recent_performance_samples" "test" {
            amount = 3
        }
    `
)

func TestAccRecentPerformanceSamplesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testRecentPerformanceSamplesConfig,
				Check: resource.ComposeTestCheckFunc(
					testRecentPerformanceSamplesSucceeds("data.solana_recent_performance_samples.test"),
				),
			},
		},
	})
}

func testRecentPerformanceSamplesSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("RecentPerformanceSamples Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("RecentPerformanceSamples Failure: ID was not set")
		}

		amount := val.Primary.Attributes["samples.#"]
		if amt, _ := strconv.Atoi(amount); amt != 3 {
			return fmt.Errorf("RecentPerformanceSamples Failure: empty or unexpected amount of samples")
		}

		return nil
	}
}
