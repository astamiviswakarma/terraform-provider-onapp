package main

import (
	"github.com/astamiviswakarma/terraform-provider-onapp/onapp"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)


func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onapp.Provider,
	})
}