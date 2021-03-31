package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceNetworkJoin() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkJoinCreate,
		Read: 	resourceNetworkJoinRead,
		Update: resourceNetworkJoinUpdate,
		Delete: resourceNetworkJoinDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceNetworkJoinDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceNetworkJoinUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceNetworkJoinRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceNetworkJoinCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}