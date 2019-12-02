package vthunder

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceConfigureSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigureSyncCreate,
		Update: resourceConfigureSyncUpdate,
		Read:   resourceConfigureSyncRead,
		Delete: resourceConfigureSyncDelete,

		Schema: map[string]*schema.Schema{
			"shared": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"partition_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"usr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"auto_authentication": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all_partitions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"pwd_enc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pwd": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceConfigureSyncCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating ConfigureSync (Inside resourceConfigureSyncCreate)")

	if client.Host != "" {
		vc := dataToConfigureSync(d)
		go_vthunder.PostConfigureSync(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceConfigureSyncRead(d, meta)
	}
	return nil
}

func resourceConfigureSyncRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading ConfigureSync (Inside resourceConfigureSyncRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		vc, err := go_vthunder.GetConfigureSync(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No ConfigureSync found")
			//d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceConfigureSyncUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceConfigureSyncRead(d, meta)
}

func resourceConfigureSyncDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceConfigureSyncRead(d, meta)
}

//Utility method to instantiate Glm Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToConfigureSync(d *schema.ResourceData) go_vthunder.Sync {
	var vc go_vthunder.Sync

	var c go_vthunder.SyncInstance

	c.AllPartitions = d.Get("all_partitions").(int)
	c.PrivateKey = d.Get("private_key").(string)

	vc.AllPartitions = c

	return vc
}
