package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testEpochDataConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_epoch" "test" {}
	`
)

func TestAccEpochDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testEpochDataConfig,
				Check: resource.ComposeTestCheckFunc(
					testEpochDataSucceeds("data.solana_epoch.test"),
				),
			},
		},
	})
}

func testEpochDataSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Epoch Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Epoch Failure: ID was not set")
		}

		return nil
	}
}
