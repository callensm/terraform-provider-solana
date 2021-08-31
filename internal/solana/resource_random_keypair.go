package solana

import (
	"crypto/sha1"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRandomKeypair() *schema.Resource {
	return &schema.Resource{
		Description: "",

		Create: resourceRandomKeypairCreate,
		Read:   resourceRandomKeypairRead,
		Delete: resourceRandomKeypairDelete,

		Schema: map[string]*schema.Schema{
			"public_key": {
				Type:        schema.TypeString,
				Description: "The randomly generated public key.",
				Computed:    true,
			},
			"private_key": {
				Type:        schema.TypeString,
				Description: "The randomly generated private key.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func resourceRandomKeypairCreate(d *schema.ResourceData, meta interface{}) error {
	pub, priv, err := solana.NewRandomPrivateKey()
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("random_keypair:%x", sha1.Sum([]byte(pub.String()))))
	d.Set("public_key", pub.String())
	d.Set("private_key", priv.String())
	return nil
}

func resourceRandomKeypairRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRandomKeypairDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
