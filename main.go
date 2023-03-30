package main

import (
	"github.com/AsahiPeak/terraform-provider-rizhiyi/rizhiyi"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: rizhiyi.Provider,
	}
	plugin.Serve(&opts)
}
