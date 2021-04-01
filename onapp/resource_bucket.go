package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceBucketCreate,
		Read:   resourceBucketRead,
		Update: resourceBucketUpdate,
		Delete: resourceBucketDelete,
		Schema: map[string]*schema.Schema{
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"currency_code": {
				Type:     schema.TypeString,
				Required: true,
			},
			"monthly_price": {
				Type: schema.TypeFloat,
				Required: true,
			},
			"allows_kms": {
				Type: schema.TypeBool,
				Required: true,
			},
			"allows_mak": {
				Type: schema.TypeBool,
				Required: true,
			},
			"allows_own": {
				Type: schema.TypeBool,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceBucketDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBucketUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBucketRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBucketCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
