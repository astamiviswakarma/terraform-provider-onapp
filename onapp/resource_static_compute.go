package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceStaticCompute() *schema.Resource {
	return &schema.Resource{
		Create: resourceStaticComputeCreate,
		Read: 	resourceStaticComputeRead,
		Update: resourceStaticComputeUpdate,
		Delete: resourceStaticComputeDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceStaticComputeDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceStaticComputeUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceStaticComputeRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceStaticComputeCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}