package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceCloudbootIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudbootIPAddressCreate,
		Read: 	resourceCloudbootIPAddressRead,
		Update: resourceCloudbootIPAddressUpdate,
		Delete: resourceCloudbootIPAddressDelete,
		Schema: map[string]*schema.Schema{
			"address": {
				Type: schema.TypeString,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceCloudbootIPAddressDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootIPAddressUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootIPAddressRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootIPAddressCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}