package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceCloudbootCompute() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudbootComputeCreate,
		Read:   resourceCloudbootComputeRead,
		Update: resourceCloudbootComputeUpdate,
		Delete: resourceCloudbootComputeDelete,
		Schema: map[string]*schema.Schema{
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pxe_ip_address_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hypervisor_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hypervisor_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"segregation_os_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"backup_ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"collect_stats": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"disable_failover": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"format_disks": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"passthrough_disks": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"storage_controller_memory_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"static_integrated_storage": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"disks_per_storage_controller": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cloud_boot_os": {
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_config": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mac": {
				Type:     schema.TypeString,
				Required: true,
			},
			"storage                     *storage `json:"storage,omitempty"`
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceCloudbootComputeDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootComputeUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootComputeRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceCloudbootComputeCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
