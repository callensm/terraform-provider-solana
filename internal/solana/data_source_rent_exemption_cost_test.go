package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testRentExemptionCostConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_rent_exemption_cost" "test" {
            data_length = 50
        }
	`
)

func TestAccRentExemptionCostSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testRentExemptionCostConfig,
				Check: resource.ComposeTestCheckFunc(
					testRentExemptionCostSucceeds("data.solana_rent_exemption_cost.test"),
				),
			},
		},
	})
}

func testRentExemptionCostSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Rent Exemption Cost Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Rent Exemption Cost Failure: ID was not set")
		}

		return nil
	}
}
