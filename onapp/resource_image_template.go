package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceImageTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageTemplateCreate,
		Read: 	resourceImageTemplateRead,
		Update: resourceImageTemplateUpdate,
		Delete: resourceImageTemplateDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceImageTemplateDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceImageTemplateUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceImageTemplateRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceImageTemplateCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}