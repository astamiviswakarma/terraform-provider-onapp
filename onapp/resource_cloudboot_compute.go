package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceCloudbootCompute() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudbootComputeCreate,
		Read: 	resourceCloudbootComputeRead,
		Update: resourceCloudbootComputeUpdate,
		Delete: resourceCloudbootComputeDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceCloudbootComputeDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootComputeUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootComputeRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootComputeCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}