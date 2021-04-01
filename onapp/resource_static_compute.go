package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceStaticCompute() *schema.Resource {
	return &schema.Resource{
		Create: resourceStaticComputeCreate,
		Read:   resourceStaticComputeRead,
		Update: resourceStaticComputeUpdate,
		Delete: resourceStaticComputeDelete,
		Schema: map[string]*schema.Schema{
			"backup_ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_units": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"disable_failover": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"failover_recipe_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hypervisor_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hypervisor_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"power_cycle_command": {
				Type:     schema.TypeString,
				Required: true,
			},
			"segregation_os_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"static_integrated_storage": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceStaticComputeDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceStaticComputeUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceStaticComputeRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceStaticComputeCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
