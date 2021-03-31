package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceSoftwareLicense() *schema.Resource {
	return &schema.Resource{
		Create: resourceSoftwareLicenseCreate,
		Read: 	resourceSoftwareLicenseRead,
		Update: resourceSoftwareLicenseUpdate,
		Delete: resourceSoftwareLicenseDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceSoftwareLicenseDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceSoftwareLicenseUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceSoftwareLicenseRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceSoftwareLicenseCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}