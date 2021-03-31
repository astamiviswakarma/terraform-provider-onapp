package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceHypervisorGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceHypervisorGroupCreate,
		Read: 	resourceHypervisorGroupRead,
		Update: resourceHypervisorGroupUpdate,
		Delete: resourceHypervisorGroupDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceHypervisorGroupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceHypervisorGroupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceHypervisorGroupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceHypervisorGroupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}