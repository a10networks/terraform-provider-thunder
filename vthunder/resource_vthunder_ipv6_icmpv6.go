package vthunder

//vThunder resource Ipv6Icmpv6

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpv6Icmpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpv6Icmpv6Create,
		Update: resourceIpv6Icmpv6Update,
		Read:   resourceIpv6Icmpv6Read,
		Delete: resourceIpv6Icmpv6Delete,
		Schema: map[string]*schema.Schema{
			"redirect": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"unreachable": {
				Type:        schema.TypeInt,
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

func resourceIpv6Icmpv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Ipv6Icmpv6 (Inside resourceIpv6Icmpv6Create) ")

		data := dataToIpv6Icmpv6(d)
		logger.Println("[INFO] received V from method data to Ipv6Icmpv6 --")
		d.SetId("1")
		go_vthunder.PostIpv6Icmpv6(client.Token, data, client.Host)

		return resourceIpv6Icmpv6Read(d, meta)

	}
	return nil
}

func resourceIpv6Icmpv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Ipv6Icmpv6 (Inside resourceIpv6Icmpv6Read)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpv6Icmpv6(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpv6Icmpv6Update(d *schema.ResourceData, meta interface{}) error {

	return resourceIpv6Icmpv6Read(d, meta)
}

func resourceIpv6Icmpv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpv6Icmpv6Read(d, meta)
}
func dataToIpv6Icmpv6(d *schema.ResourceData) go_vthunder.Icmpv6 {
	var vc go_vthunder.Icmpv6
	var c go_vthunder.Icmpv6Instance
	c.Redirect = d.Get("redirect").(int)
	c.Unreachable = d.Get("unreachable").(int)

	vc.UUID = c
	return vc
}
