package onapp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"time"
)

func resourceDataStoreJoin() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataStoreJoinCreate,
		Read: 	resourceDataStoreJoinRead,
		Update: resourceDataStoreJoinUpdate,
		Delete: resourceDataStoreJoinDelete,
		Schema: map[string]*schema.Schema{
			//@TODO: implement this
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
			Delete: schema.DefaultTimeout(20 * time.Minute), //@TODO: reconfigure this
		},
	}
}

func resourceDataStoreJoinDelete(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreJoinUpdate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreJoinRead(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}

func resourceDataStoreJoinCreate(data *schema.ResourceData, i interface{}) error {
	//@TODO: implement this
	return nil
}