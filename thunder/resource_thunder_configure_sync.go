package thunder

import (
	"github.com/go_thunder/thunder"
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
	client := meta.(Thunder)

	logger.Println("[INFO] Creating ConfigureSync (Inside resourceConfigureSyncCreate)")

	if client.Host != "" {
		vc := dataToConfigureSync(d)
		go_thunder.PostConfigureSync(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceConfigureSyncRead(d, meta)
	}
	return nil
}

func resourceConfigureSyncRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading ConfigureSync (Inside resourceConfigureSyncRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		vc, err := go_thunder.GetConfigureSync(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No ConfigureSync found")

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

func dataToConfigureSync(d *schema.ResourceData) go_thunder.Sync {
	var vc go_thunder.Sync

	var c go_thunder.SyncInstance

	c.AllPartitions = d.Get("all_partitions").(int)
	c.PrivateKey = d.Get("private_key").(string)
	c.PartitionName = d.Get("partition_name").(string)
	c.Pwd = d.Get("pwd").(string)
	c.AutoAuthentication = d.Get("auto_authentication").(int)
	c.Address = d.Get("address").(string)
	c.Shared = d.Get("shared").(int)
	c.Type = d.Get("type").(string)
	c.PwdEnc = d.Get("pwd_enc").(string)
	c.Usr = d.Get("usr").(string)

	vc.AllPartitions = c

	return vc
}
