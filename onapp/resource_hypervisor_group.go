package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceHypervisorGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceHypervisorGroupCreate,
		Read:   resourceHypervisorGroupRead,
		Update: resourceHypervisorGroupUpdate,
		Delete: resourceHypervisorGroupDelete,
		Schema: map[string]*schema.Schema{
			"cpu_flags_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"cpu_units": {
				Type:     schema.TypeInt,
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
			"failover_timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_vms_start_at_once": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"preconfigured_only": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"recovery_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"release_resource_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"run_sysprep": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"server_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vlan": {
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

func resourceHypervisorGroupDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceHypervisorGroupUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceHypervisorGroupRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceHypervisorGroupCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
