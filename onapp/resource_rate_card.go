package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceRateCard() *schema.Resource {
	return &schema.Resource{
		Create: resourceRateCardCreate,
		Read: 	resourceRateCardRead,
		Update: resourceRateCardUpdate,
		Delete: resourceRateCardDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
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