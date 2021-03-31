package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceIPNet() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPNetCreate,
		Read: 	resourceIPNetRead,
		Update: resourceIPNetUpdate,
		Delete: resourceIPNetDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceIPNetDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIPNetUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIPNetRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIPNetCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}