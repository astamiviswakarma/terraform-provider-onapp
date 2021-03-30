package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Endpoint hostname",
			},
		},
		//ConfigureContextFunc: configureContextFunc,
	}
}

//func configureContextFunc(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
//	return
//}