package solana

import (
	"errors"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

var (
	clusterNameOptions  = []string{"localnet", "testnet", "devnet", "mainnet-beta", "serumnet"}
	dataEncodingOptions = []string{
		string(solana.EncodingBase58),
		string(solana.EncodingBase64),
		string(solana.EncodingJSONParsed),
	}
)

type providerConfig struct {
	rpcClient *rpc.Client
}

// New compiles and returns a new instance of the Solana provider
func New() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: initializeProvider,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Name of the Solana cluster to target. This field is mutually exclusive with `endpoint`.",
				ConflictsWith: []string{"endpoint"},
				ValidateFunc:  validation.StringInSlice(clusterNameOptions, false),
			},
			"endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The RPC endpoint for the target Solana cluster. This field is mutually exclusive with `cluster`.",
				ConflictsWith: []string{"cluster"},
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"solana_associated_token_account": resourceAssociatedTokenAccount(),
			"solana_keypair":                  resourceKeypair(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"solana_account":                    dataSourceAccount(),
			"solana_address_signatures":         dataSourceAddressSignatures(),
			"solana_balance":                    dataSourceBalance(),
			"solana_epoch":                      dataSourceEpoch(),
			"solana_genesis_hash":               dataSourceGenesisHash(),
			"solana_health":                     dataSourceHealth(),
			"solana_node_identity":              dataSourceNodeIdentity(),
			"solana_recent_performance_samples": dataSourceRecentPerformanceSamples(),
			"solana_recent_prioritization_fees": dataSourceRecentPrioritizationFees(),
			"solana_rent_exemption_cost":        dataSourceRentExemptionCost(),
			"solana_signature_status":           dataSourceSignatureStatus(),
			"solana_slot_leaders":               dataSourceSlotLeaders(),
			"solana_supply":                     dataSourceSupply(),
			"solana_token_supply":               dataSourceTokenSupply(),
			"solana_transaction":                dataSourceTransaction(),
			"solana_version":                    dataSourceVersion(),
		},
	}
}

func initializeProvider(d *schema.ResourceData) (any, error) {
	var endpoint string

	if clusterArg, ok := d.GetOk("cluster"); ok {
		endpoint = clusterEndpoints[clusterArg.(string)]
	} else if endpointArg, ok := d.GetOk("endpoint"); ok {
		endpoint = endpointArg.(string)
	}

	if endpoint == "" {
		return nil, errors.New("the received provider inputs were not sufficient to derive a cluster endpoint")
	}

	return &providerConfig{
		rpcClient: rpc.New(endpoint),
	}, nil
}
