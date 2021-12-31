package solana

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	testRandomKeypairConfig = `
		provider "solana" {
			endpoint = "https://api.testnet.solana.com"
		}

		resource "solana_keypair" "test" {
            output_path = "${path.module}/random.json"
            random      = true
        }
	`

	testGrindedKeypairConfig = `
        provider "solana" {
            cluster = "testnet"
        }

        resource "solana_keypair" "test" {
            output_path = "${path.module}/grinded.json"

            grind {
                prefix = "abc"
            }
        }
    `
)

func TestAccKeypairResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testRandomKeypairConfig,
				Check: resource.ComposeTestCheckFunc(
					testKeypairSucceeds("solana_keypair.test"),
				),
			},
			{
				Config: testGrindedKeypairConfig,
				Check: resource.ComposeTestCheckFunc(
					testKeypairSucceeds("solana_keypair.test"),
				),
			},
		},
	})
}

func testKeypairSucceeds(name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		val, ok := state.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Keypair Failure: %s not found", name)
		}

		if val.Primary.ID == "" {
			return fmt.Errorf("Keypair Failure: ID was not set")
		}

		if val.Primary.Attributes["public_key"] == "" || val.Primary.Attributes["private_key"] == "" {
			return fmt.Errorf("Keypair Failure: `public_key` and/or `private_key` were not set")
		}

		return nil
	}
}
