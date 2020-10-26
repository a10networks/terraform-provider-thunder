package thunder

//Thunder resource IpDnsPrimary

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpDnsPrimary() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpDnsPrimaryCreate,
		Update: resourceIpDnsPrimaryUpdate,
		Read:   resourceIpDnsPrimaryRead,
		Delete: resourceIpDnsPrimaryDelete,
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

func resourceIpDnsPrimaryCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsPrimary (Inside resourceIpDnsPrimaryCreate) ")

		data := dataToIpDnsPrimary(d)
		logger.Println("[INFO] received V from method data to IpDnsPrimary --")
		d.SetId("1")
		go_thunder.PostIpDnsPrimary(client.Token, data, client.Host)

		return resourceIpDnsPrimaryRead(d, meta)

	}
	return nil
}

func resourceIpDnsPrimaryRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpDnsPrimary (Inside resourceIpDnsPrimaryRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpDnsPrimary(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpDnsPrimaryUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpDnsPrimaryRead(d, meta)
}

func resourceIpDnsPrimaryDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpDnsPrimaryRead(d, meta)
}
func dataToIpDnsPrimary(d *schema.ResourceData) go_thunder.DnsPrimary {
	var vc go_thunder.DnsPrimary
	var c go_thunder.DnsPrimaryInstance
	c.IPV4Addr = d.Get("ip_v4_addr").(string)
	c.IPV6Addr = d.Get("ip_v6_addr").(string)

	vc.UUID = c
	return vc
}
