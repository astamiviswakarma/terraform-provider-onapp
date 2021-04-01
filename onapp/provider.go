package onapp

import (
	"github.com/OnApp/onapp-sdk-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host to connect",
				DefaultFunc: schema.EnvDefaultFunc("ONAPP_HOSTNAME", ""),
			},
			"allowUnverifiedSSL": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Allow unverified ssl connection to host",
				DefaultFunc: schema.EnvDefaultFunc("ONAPP_ALLOW_UNVERIFIED_SSL", ""),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username",
				DefaultFunc: schema.EnvDefaultFunc("ONAPP_USERNAME", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access Token",
				DefaultFunc: schema.EnvDefaultFunc("ONAPP_TOKEN", ""),
			},
		},
		ConfigureFunc: configureFunc,
	}
}

func configureFunc(data *schema.ResourceData) (interface{}, error) {
	return onappgo.New(nil,
		onappgo.SetBaseURL(data.Get("hostname").(string)),
		onappgo.SetAllowUnverifiedSSL(data.Get("allowUnverifiedSSL").(bool)),
		onappgo.SetBasicAuth(data.Get("username").(string), data.Get("token").(string)),
	)
}
