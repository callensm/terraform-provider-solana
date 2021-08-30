package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testVersionConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		data "solana_version" "test" {}
	`
)

func TestAccVersionSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testVersionConfig,
				Check: resource.ComposeTestCheckFunc(
					testVersionSucceeds("data.solana_version.test"),
				),
			},
		},
	})
}

func testVersionSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Version Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Version Failure: ID was not set")
		}

		return nil
	}
}
