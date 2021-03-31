package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceDataStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataStoreCreate,
		Read: 	resourceDataStoreRead,
		Update: resourceDataStoreUpdate,
		Delete: resourceDataStoreDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceDataStoreDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}