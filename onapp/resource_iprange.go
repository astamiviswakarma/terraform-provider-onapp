package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceIPRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPRangeCreate,
		Read: 	resourceIPRangeRead,
		Update: resourceIPRangeUpdate,
		Delete: resourceIPRangeDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceIPRangeDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIPRangeUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIPRangeRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceIPRangeCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}