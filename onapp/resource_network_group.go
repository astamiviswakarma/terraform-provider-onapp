package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceNetworkGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkGroupCreate,
		Read: 	resourceNetworkGroupRead,
		Update: resourceNetworkGroupUpdate,
		Delete: resourceNetworkGroupDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
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