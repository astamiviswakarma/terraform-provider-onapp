package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualMachineCreate,
		Read:   resourceVirtualMachineRead,
		Update: resourceVirtualMachineUpdate,
		Delete: resourceVirtualMachineDelete,
		Schema: map[string]*schema.Schema{
			"custom_recipe_variables_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"acceleration_allowed": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"admin_note": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_shares": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cpu_sockets": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpus": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_store_group_primary_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"data_store_group_swap_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"enable_autoscale": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hypervisor_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"hypervisor_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"initial_root_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initial_root_password_encryption_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_package_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"label": {
				Type:     schema.TypeString,
				Required: true,
			},
			"licensing_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"licensing_server_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"licensing_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"location_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"network_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"primary_disk_min_iops": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"primary_disk_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"primary_network_group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"rate_limit": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"recipe_joins_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"required_automatic_backup": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"required_ipaddress_assignment": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"required_virtual_machine_build": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"required_virtual_machine_startup": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"selected_ipaddress": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_addon_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional: true,
			},
			"swap_disk_min_iops": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"swap_disk_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"template_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_of_format": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virsh_console": {
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

func resourceVirtualMachineDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceVirtualMachineUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceVirtualMachineRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceVirtualMachineCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}
