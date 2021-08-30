package solana

import (
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type providerConfig struct {
	rpcClient *rpc.Client
}

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The RPC endpoint for the target Solana cluster",
			},
		},

		ConfigureFunc: initializeProvider,

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"solana_account_info": dataSourceAccountInfo(),
			"solana_balance":      dataSourceBalance(),
		},
	}
}

func initializeProvider(d *schema.ResourceData) (interface{}, error) {
	return &providerConfig{
		rpcClient: rpc.NewClient(d.Get("endpoint").(string)),
	}, nil
}
