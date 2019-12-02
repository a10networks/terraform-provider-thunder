package vthunder

//vThunder resource Vrrp common

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceVrrpSessionSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceVrrpSessionSyncCreate,
		Update: resourceVrrpSessionSyncUpdate,
		Read:   resourceVrrpSessionSyncRead,
		Delete: resourceVrrpSessionSyncDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceVrrpSessionSyncCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating VrrpSessionSync (Inside resourceVrrpSessionSyncCreate)")

	if client.Host != "" {
		vc := dataToVrrpSessionSync(d)
		go_vthunder.PostVrrpSessionSync(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceVrrpSessionSyncRead(d, meta)
	}
	return nil
}

func resourceVrrpSessionSyncRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading VrrpSessionSync (Inside resourceVrrpSessionSyncRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		vc, err := go_vthunder.GetVrrpSessionSync(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No VrrpSessionSync found")
			//d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceVrrpSessionSyncUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceVrrpSessionSyncRead(d, meta)
}

func resourceVrrpSessionSyncDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceVrrpSessionSyncRead(d, meta)
}

//Utility method to instantiate DnsPrimary Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToVrrpSessionSync(d *schema.ResourceData) go_vthunder.SessionSync {
	var vc go_vthunder.SessionSync

	var c go_vthunder.SessionSyncInstance

	c.Action = d.Get("action").(string)

	vc.Action = c

	return vc
}
