package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBackupResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupResourceCreate,
		Read: 	resourceBackupResourceRead,
		Update: resourceBackupResourceUpdate,
		Delete: resourceBackupResourceDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceBackupResourceDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupResourceUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupResourceRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupResourceCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}