package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func resourceAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessControlCreate,
		Read:   resourceAccessControlRead,
		Update: resourceAccessControlUpdate,
		Delete: resourceAccessControlDelete,
		Schema: map[string]*schema.Schema{
			"bucket_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limits": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceAccessControlDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceAccessControlUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	log.Println("onapp.resourceAccessControlUpdate")
	log.Println("No Op!")

	return nil
}

func resourceAccessControlRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceAccessControlCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
