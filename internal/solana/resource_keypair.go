package solana

import (
	"crypto/sha1"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/gagliardetto/solana-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceKeypair() *schema.Resource {
	return &schema.Resource{
		Description: "Genereate a random or grinded keypair compatible with the `solana-keygen` CLI.",

		Create: resourceKeypairCreate,
		Read:   resourceKeypairRead,
		Delete: resourceKeypairDelete,

		Schema: map[string]*schema.Schema{
			"grind": {
				Type:         schema.TypeSet,
				Description:  "The grind options for the new keypair.",
				Optional:     true,
				ExactlyOneOf: []string{"grind", "random"},
				ForceNew:     true,
				MaxItems:     1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitive": {
							Type:        schema.TypeBool,
							Description: "Case sensitivity for the prefix and suffix parameters.",
							Optional:    true,
							Default:     false,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "The desired prefix of the public key.",
							Optional:    true,
						},
						"suffix": {
							Type:        schema.TypeString,
							Description: "The desired suffix of the public key.",
							Optional:    true,
						},
						"threads": {
							Type:         schema.TypeInt,
							Description:  "The number of threads to grind for the keypair (between 1-15).",
							Optional:     true,
							Default:      10,
							ForceNew:     true,
							ValidateFunc: validation.IntBetween(1, 15),
						},
					},
				},
			},
			"random": {
				Type:         schema.TypeBool,
				Description:  "Whether to generate a random new keypair.",
				Optional:     true,
				Default:      false,
				ExactlyOneOf: []string{"grind", "random"},
				ForceNew:     true,
			},
			"output_path": {
				Type:         schema.TypeString,
				Description:  "The file path to output the keypair into.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.All(validation.StringIsNotEmpty, validation.StringMatch(regexp.MustCompile(`.+\.json`), "Output file must be a JSON extension.")),
			},
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

func resourceKeypairCreate(d *schema.ResourceData, meta any) error {
	var key solana.PrivateKey
	var err error

	if rnd := d.Get("random").(bool); rnd {
		key, err = solana.NewRandomPrivateKey()
		if err != nil {
			return err
		}
	} else {
		opts := d.Get("grind").(*schema.Set).List()[0].(map[string]any)
		threads := opts["threads"].(int)

		var wg sync.WaitGroup

		keyChan := make(chan solana.PrivateKey, 1)
		defer close(keyChan)

		stopChan := make(chan bool, threads)
		defer close(stopChan)

		for range threads {
			wg.Go(func() {
				grindKeypairWithOptions(opts["prefix"].(string), opts["suffix"].(string), opts["case_sensitive"].(bool), keyChan, stopChan)
			})
		}

		key = <-keyChan
		for range threads {
			stopChan <- true
		}

		wg.Wait()
	}

	err = os.WriteFile(d.Get("output_path").(string), []byte(bytesAsJSONArray(key)), os.ModePerm)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("keypair:%x", sha1.Sum([]byte(key.PublicKey().String()))))
	d.Set("public_key", key.PublicKey().String())
	d.Set("private_key", key.String())
	return nil
}

func resourceKeypairRead(d *schema.ResourceData, meta any) error {
	key, err := solana.PrivateKeyFromSolanaKeygenFile(d.Get("output_path").(string))
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("keypair:%x", sha1.Sum([]byte(key.PublicKey().String()))))
	d.Set("public_key", key.PublicKey().String())
	d.Set("private_key", key.String())
	return nil
}

func resourceKeypairDelete(d *schema.ResourceData, meta any) error {
	err := os.Remove(d.Get("output_path").(string))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func grindKeypairWithOptions(prefix, suffix string, caseSensitive bool, keyChan chan<- solana.PrivateKey, stopChan <-chan bool) {
	for {
		select {
		case <-stopChan:
			return
		default:
			keypair, err := solana.NewRandomPrivateKey()
			if err != nil {
				continue
			}

			pubkey := keypair.PublicKey().String()

			if caseSensitive {
				if prefix != "" && !strings.HasPrefix(pubkey, prefix) {
					continue
				}

				if suffix != "" && !strings.HasSuffix(pubkey, suffix) {
					continue
				}
			} else {
				lowerPubkey := strings.ToLower(pubkey)

				if prefix != "" && !strings.HasPrefix(lowerPubkey, strings.ToLower(prefix)) {
					continue
				}

				if suffix != "" && !strings.HasPrefix(lowerPubkey, strings.ToLower(suffix)) {
					continue
				}
			}

			keyChan <- keypair
			return
		}
	}
}

func bytesAsJSONArray(x []byte) string {
	bytesAsString := make([]string, len(x))
	for i, y := range x {
		bytesAsString[i] = fmt.Sprintf("%v", y)
	}
	return fmt.Sprintf("[%s]", strings.Join(bytesAsString, ","))
}
