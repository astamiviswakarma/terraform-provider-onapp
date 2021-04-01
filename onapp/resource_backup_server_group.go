package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBackupServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupServerGroupCreate,
		Read:   resourceBackupServerGroupRead,
		Update: resourceBackupServerGroupUpdate,
		Delete: resourceBackupServerGroupDelete,
		Schema: map[string]*schema.Schema{
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"server_type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceBackupServerGroupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerGroupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerGroupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerGroupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
