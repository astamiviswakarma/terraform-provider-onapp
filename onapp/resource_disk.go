package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceDisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiskCreate,
		Read:   resourceDiskRead,
		Update: resourceDiskUpdate,
		Delete: resourceDiskDelete,
		Schema: map[string]*schema.Schema{
			"add_to_linux_fstab": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"add_to_freebsd_fstab": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"primary": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"disk_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"file_system": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_store_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"require_format_disk": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"mount_point": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hot_attach": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"min_iops": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"mounted": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Additional field to determine Virtual Machine to create disk
			"virtual_machine_id": {
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

func resourceDiskDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDiskUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDiskRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDiskCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
