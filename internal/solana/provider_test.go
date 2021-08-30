package solana

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

var (
	testAccProviders = map[string]*schema.Provider{
		"solana": New(),
	}
)

func TestProviderValidation(t *testing.T) {
	p := testAccProviders["solana"]
	err := p.InternalValidate()
	assert.NoError(t, err)
}
