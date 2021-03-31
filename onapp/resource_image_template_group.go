package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceImageTemplateGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageTemplateGroupCreate,
		Read: 	resourceImageTemplateGroupRead,
		Update: resourceImageTemplateGroupUpdate,
		Delete: resourceImageTemplateGroupDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceImageTemplateGroupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceImageTemplateGroupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceImageTemplateGroupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceImageTemplateGroupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}