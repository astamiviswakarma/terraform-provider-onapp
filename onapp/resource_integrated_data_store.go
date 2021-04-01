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
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"replicas": {
				Type: schema.TypeString,
				Required: true,
			},
			"stripes": {
				Type: schema.TypeString,
				Required: true,
			},
			"overcommit": {
				Type: schema.TypeString,
				Required: true,
			},
			"node_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
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