package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceDataStoreGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataStoreGroupCreate,
		Read:   resourceDataStoreGroupRead,
		Update: resourceDataStoreGroupUpdate,
		Delete: resourceDataStoreGroupDelete,
		Schema: map[string]*schema.Schema{
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"preconfigured_only": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"server_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceDataStoreGroupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreGroupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreGroupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreGroupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
