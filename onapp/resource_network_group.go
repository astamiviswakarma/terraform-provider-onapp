package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceNetworkGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkGroupCreate,
		Read:   resourceNetworkGroupRead,
		Update: resourceNetworkGroupUpdate,
		Delete: resourceNetworkGroupDelete,
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

func resourceNetworkGroupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceNetworkGroupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceNetworkGroupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceNetworkGroupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
