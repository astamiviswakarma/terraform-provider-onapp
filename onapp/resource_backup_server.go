package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBackupServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupServerCreate,
		Read: 	resourceBackupServerRead,
		Update: resourceBackupServerUpdate,
		Delete: resourceBackupServerDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceBackupServerDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}