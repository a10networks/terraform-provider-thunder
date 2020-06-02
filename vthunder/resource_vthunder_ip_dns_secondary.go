package vthunder

//vThunder resource IpDnsSecondary

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpDnsSecondary() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpDnsSecondaryCreate,
		Update: resourceIpDnsSecondaryUpdate,
		Read:   resourceIpDnsSecondaryRead,
		Delete: resourceIpDnsSecondaryDelete,
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

func resourceIpDnsSecondaryCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsSecondary (Inside resourceIpDnsSecondaryCreate) ")

		data := dataToIpDnsSecondary(d)
		logger.Println("[INFO] received V from method data to IpDnsSecondary --")
		d.SetId("1")
		go_vthunder.PostIpDnsSecondary(client.Token, data, client.Host)

		return resourceIpDnsSecondaryRead(d, meta)

	}
	return nil
}

func resourceIpDnsSecondaryRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IpDnsSecondary (Inside resourceIpDnsSecondaryRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpDnsSecondary(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpDnsSecondaryUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpDnsSecondaryRead(d, meta)
}

func resourceIpDnsSecondaryDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpDnsSecondaryRead(d, meta)
}
func dataToIpDnsSecondary(d *schema.ResourceData) go_vthunder.DnsSecondary {
	var vc go_vthunder.DnsSecondary
	var c go_vthunder.DnsSecondaryInstance
	c.IPV4Addr = d.Get("ip_v4_addr").(string)
	c.IPV6Addr = d.Get("ip_v6_addr").(string)

	vc.UUID = c
	return vc
}
