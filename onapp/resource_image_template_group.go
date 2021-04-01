package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceImageTemplateGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceImageTemplateGroupCreate,
		Read:   resourceImageTemplateGroupRead,
		Update: resourceImageTemplateGroupUpdate,
		Delete: resourceImageTemplateGroupDelete,
		Schema: map[string]*schema.Schema{
			"kms": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"kms_host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kms_port": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kms_server_label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mak": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"own": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
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
