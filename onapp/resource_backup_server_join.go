package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBackupServerJoin() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupServerJoinCreate,
		Read: 	resourceBackupServerJoinRead,
		Update: resourceBackupServerJoinUpdate,
		Delete: resourceBackupServerJoinDelete,
		Schema: map[string]*schema.Schema{
			"backup_server_id": {
				Type: schema.TypeInt,
				Required: true,
			},
			"target_join_id": {
				Type: schema.TypeInt,
				Required: true,
			},
			"target_join_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceBackupServerJoinDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerJoinUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerJoinRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupServerJoinCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}