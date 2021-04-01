package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceRateCard() *schema.Resource {
	return &schema.Resource{
		Create: resourceRateCardCreate,
		Read:   resourceRateCardRead,
		Update: resourceRateCardUpdate,
		Delete: resourceRateCardDelete,
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
			"timing_strategy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"apply_to_all_resources_in_the_bucket": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"prices": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceRateCardDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceRateCardUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceRateCardRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceRateCardCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
