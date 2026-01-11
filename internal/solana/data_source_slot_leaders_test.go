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
            start_slot = 380950223
        }
    `
)

var expectedLeaders = [3]string{"dzkLmvHKsScz186ZAvnXMyHCeRznFrKyX7pzbozbz4T", "4eDS5fB3keUA5qxwphZaCDGUqfrmZSry1E2t9bjzUPHS", "4eDS5fB3keUA5qxwphZaCDGUqfrmZSry1E2t9bjzUPHS"}

func TestAccSlotLeaders(t *testing.T) {
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

		l0 := val.Primary.Attributes["leaders.0"]
		l1 := val.Primary.Attributes["leaders.1"]
		l2 := val.Primary.Attributes["leaders.2"]

		if l0 != expectedLeaders[0] || l1 != expectedLeaders[1] || l2 != expectedLeaders[2] {
			return fmt.Errorf("Slot Leaders Failure: leader address arrays not do match")
		}

		return nil
	}
}
