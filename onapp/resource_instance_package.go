package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceInstancePackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstancePackageCreate,
		Read: 	resourceInstancePackageRead,
		Update: resourceInstancePackageUpdate,
		Delete: resourceInstancePackageDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceInstancePackageDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceInstancePackageUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceInstancePackageRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceInstancePackageCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}