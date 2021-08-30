package solana

import (
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

var (
	clusterNameOptions = []string{"localnet", "testnet", "devnet", "mainnet-beta", "serumnet"}
)

type providerConfig struct {
	rpcClient *rpc.Client
}

func New() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: initializeProvider,

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Name of the Solana cluster to target",
				ConflictsWith: []string{"endpoint"},
				ValidateFunc:  validateProviderClusterName,
			},
			"endpoint": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "The RPC endpoint for the target Solana cluster",
				ConflictsWith: []string{"cluster"},
			},
		},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"solana_account":             dataSourceAccount(),
			"solana_balance":             dataSourceBalance(),
			"solana_epoch":               dataSourceEpoch(),
			"solana_recent_blockhash":    dataSourceRecentBlockhash(),
			"solana_rent_exemption_cost": dataSourceRentExemptionCost(),
			"solana_supply":              dataSourceSupply(),
			"solana_transaction":         dataSourceTransaction(),
			"solana_version":             dataSourceVersion(),
		},
	}
}

func initializeProvider(d *schema.ResourceData) (interface{}, error) {
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
		rpcClient: rpc.NewClient(endpoint),
	}, nil
}

func validateProviderClusterName(val interface{}, k string) (warnings []string, errs []error) {
	for _, name := range clusterNameOptions {
		if name == val.(string) {
			return
		}
	}

	errs = append(errs, fmt.Errorf("you received cluster name (%s) is not a valid option", val.(string)))
	return
}