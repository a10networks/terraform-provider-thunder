package thunder

//Thunder resource IpDnsSuffix

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpDnsSuffix() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpDnsSuffixCreate,
		Update: resourceIpDnsSuffixUpdate,
		Read:   resourceIpDnsSuffixRead,
		Delete: resourceIpDnsSuffixDelete,
		Schema: map[string]*schema.Schema{
			"domain_name": {
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

func resourceIpDnsSuffixCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsSuffix (Inside resourceIpDnsSuffixCreate) ")

		data := dataToIpDnsSuffix(d)
		logger.Println("[INFO] received V from method data to IpDnsSuffix --")
		d.SetId("1")
		go_thunder.PostIpDnsSuffix(client.Token, data, client.Host)

		return resourceIpDnsSuffixRead(d, meta)

	}
	return nil
}

func resourceIpDnsSuffixRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpDnsSuffix (Inside resourceIpDnsSuffixRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpDnsSuffix(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpDnsSuffixUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpDnsSuffixRead(d, meta)
}

func resourceIpDnsSuffixDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpDnsSuffixRead(d, meta)
}
func dataToIpDnsSuffix(d *schema.ResourceData) go_thunder.DnsSuffix {
	var vc go_thunder.DnsSuffix
	var c go_thunder.DnsSuffixInstance
	c.DomainName = d.Get("domain_name").(string)

	vc.UUID = c
	return vc
}
