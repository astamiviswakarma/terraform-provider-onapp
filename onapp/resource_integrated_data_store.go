package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceIntegratedDataStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntegratedDataStoreCreate,
		Read: 	resourceIntegratedDataStoreRead,
		Update: resourceIntegratedDataStoreUpdate,
		Delete: resourceIntegratedDataStoreDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceIntegratedDataStoreDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIntegratedDataStoreUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIntegratedDataStoreRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIntegratedDataStoreCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}