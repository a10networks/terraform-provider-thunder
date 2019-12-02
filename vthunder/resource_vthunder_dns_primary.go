package vthunder

//vThunder resource Vrrp common

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceDnsPrimary() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsPrimaryCreate,
		Update: resourceDnsPrimaryUpdate,
		Read:   resourceDnsPrimaryRead,
		Delete: resourceDnsPrimaryDelete,

		Schema: map[string]*schema.Schema{
			"ip_v6_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_v4_addr": {
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

func resourceDnsPrimaryCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating DnsPrimary (Inside resourceDnsPrimaryCreate)")

	if client.Host != "" {
		vc := dataToDnsPrimary(d)
		go_vthunder.PostDnsPrimary(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceDnsPrimaryRead(d, meta)
	}
	return nil
}

func resourceDnsPrimaryRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DnsPrimary (Inside resourceDnsPrimaryRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		vc, err := go_vthunder.GetDnsPrimary(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No DnsPrimary found")
			//d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceDnsPrimaryUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceDnsPrimaryRead(d, meta)
}

func resourceDnsPrimaryDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceDnsPrimaryRead(d, meta)
}

//Utility method to instantiate DnsPrimary Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToDnsPrimary(d *schema.ResourceData) go_vthunder.Primary {
	var vc go_vthunder.Primary

	var c go_vthunder.PrimaryInstance

	c.IPV4Addr = d.Get("ip_v4_addr").(string)

	vc.IPV4Addr = c

	return vc
}
