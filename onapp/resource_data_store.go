package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceDataStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataStoreCreate,
		Read:   resourceDataStoreRead,
		Update: resourceDataStoreUpdate,
		Delete: resourceDataStoreDelete,
		Schema: map[string]*schema.Schema{
			"label": {
				Type: schema.TypeString,
				Required: true,
			},
			"data_store_group_id": {
				Type: schema.TypeInt,
				Required: true,
			},
			"local_hypervisor_id": {
				Type: schema.TypeInt,
				Required: true,
			},
			"ip": {
				Type: schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"data_store_size": {
				Type: schema.TypeInt,
				Required: true,
			},
			"data_store_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"iscsi_ip": {
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
