package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceUserWhiteList() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserWhiteListCreate,
		Read: 	resourceUserWhiteListRead,
		Update: resourceUserWhiteListUpdate,
		Delete: resourceUserWhiteListDelete,
		Schema: map[string]*schema.Schema{
			// can be empty
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ip": {
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

func resourceUserWhiteListDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceUserWhiteListUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceUserWhiteListRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceUserWhiteListCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}