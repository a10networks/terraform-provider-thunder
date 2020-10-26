package thunder

//Thunder resource Vrrp common

import (
	"github.com/go_thunder/thunder"
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
	client := meta.(Thunder)

	logger.Println("[INFO] Creating DnsPrimary (Inside resourceDnsPrimaryCreate)")

	if client.Host != "" {
		vc := dataToDnsPrimary(d)
		go_thunder.PostDnsPrimary(client.Token, vc, client.Host)
		d.SetId("1")

		return resourceDnsPrimaryRead(d, meta)
	}
	return nil
}

func resourceDnsPrimaryRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DnsPrimary (Inside resourceDnsPrimaryRead)")

	client := meta.(Thunder)

	if client.Host != "" {

		vc, err := go_thunder.GetDnsPrimary(client.Token, client.Host)

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

func dataToDnsPrimary(d *schema.ResourceData) go_thunder.Primary {
	var vc go_thunder.Primary

	var c go_thunder.PrimaryInstance

	c.IPV4Addr = d.Get("ip_v4_addr").(string)

	vc.IPV4Addr = c

	return vc
}
