package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testSlotLeadersConfig = `
        provider "solana" {
            endpoint = "https://api.testnet.solana.com"
        }

        data "solana_slot_leaders" "test" {
            limit      = 3
            start_slot = 384858752
        }
    `
)

func TestAccSlotLeadersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testSlotLeadersConfig,
				Check: resource.ComposeTestCheckFunc(
					testSlotLeaders("data.solana_slot_leaders.test"),
				),
			},
		},
	})
}

func testSlotLeaders(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Slot Leaders Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Slot Leaders Failure: ID was not set")
		}

		amt := val.Primary.Attributes["leaders.#"]

		if amt != "3" {
			return fmt.Errorf("Slot Leaders Failure: leader address arrays not do match")
		}

		return nil
	}
}
