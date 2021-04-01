package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupCreate,
		Read:   resourceBackupRead,
		Update: resourceBackupUpdate,
		Delete: resourceBackupDelete,
		Schema: map[string]*schema.Schema{
			"disk_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"note": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_windows_backup": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"virtual_machine_id": { // Additional fie42ld to determine Virtual Machine to create disk backup
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

func resourceBackupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceBackupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
