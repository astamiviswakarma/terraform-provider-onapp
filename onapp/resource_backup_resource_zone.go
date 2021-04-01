package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBackupResourceZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupResourceZoneCreate,
		Read: 	resourceBackupResourceZoneRead,
		Update: resourceBackupResourceZoneUpdate,
		Delete: resourceBackupResourceZoneDelete,
		Schema: map[string]*schema.Schema{
			"label": {
				Type: schema.TypeString,
				Required: true,
			},
			"location_group_id": {
				Type: schema.TypeInt,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceBackupResourceZoneDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupResourceZoneUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupResourceZoneRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupResourceZoneCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}