package main

import (
	"github.com/callensm/terraform-provider-solana/solana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	version string
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: solana.New,
	})
}
