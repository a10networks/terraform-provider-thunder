package vthunder

//vThunder resource IpDnsSuffix

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
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
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpDnsSuffix (Inside resourceIpDnsSuffixCreate) ")

		data := dataToIpDnsSuffix(d)
		logger.Println("[INFO] received V from method data to IpDnsSuffix --")
		d.SetId("1")
		go_vthunder.PostIpDnsSuffix(client.Token, data, client.Host)

		return resourceIpDnsSuffixRead(d, meta)

	}
	return nil
}

func resourceIpDnsSuffixRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IpDnsSuffix (Inside resourceIpDnsSuffixRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpDnsSuffix(client.Token, client.Host)
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
func dataToIpDnsSuffix(d *schema.ResourceData) go_vthunder.DnsSuffix {
	var vc go_vthunder.DnsSuffix
	var c go_vthunder.DnsSuffixInstance
	c.DomainName = d.Get("domain_name").(string)

	vc.UUID = c
	return vc
}
