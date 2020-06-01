package vthunder

//vThunder resource Ipv6NatIcmpv6

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpv6NatIcmpv6() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpv6NatIcmpv6Create,
		Update: resourceIpv6NatIcmpv6Update,
		Read:   resourceIpv6NatIcmpv6Read,
		Delete: resourceIpv6NatIcmpv6Delete,
		Schema: map[string]*schema.Schema{
			"respond_to_ping": {
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

func resourceIpv6NatIcmpv6Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Ipv6NatIcmpv6 (Inside resourceIpv6NatIcmpv6Create) ")

		data := dataToIpv6NatIcmpv6(d)
		logger.Println("[INFO] received V from method data to Ipv6NatIcmpv6 --")
		d.SetId("1")
		go_vthunder.PostIpv6NatIcmpv6(client.Token, data, client.Host)

		return resourceIpv6NatIcmpv6Read(d, meta)

	}
	return nil
}

func resourceIpv6NatIcmpv6Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Ipv6NatIcmpv6 (Inside resourceIpv6NatIcmpv6Read)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpv6NatIcmpv6(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpv6NatIcmpv6Update(d *schema.ResourceData, meta interface{}) error {

	return resourceIpv6NatIcmpv6Read(d, meta)
}

func resourceIpv6NatIcmpv6Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpv6NatIcmpv6Read(d, meta)
}
func dataToIpv6NatIcmpv6(d *schema.ResourceData) go_vthunder.NatIcmpv6 {
	var vc go_vthunder.NatIcmpv6
	var c go_vthunder.NatIcmpv6Instance
	c.RespondToPing = d.Get("respond_to_ping").(int)

	vc.UUID = c
	return vc
}
