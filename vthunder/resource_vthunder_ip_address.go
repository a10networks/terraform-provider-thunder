package vthunder

//vThunder resource IPAddress

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPAddressCreate,
		Update: resourceIPAddressUpdate,
		Read:   resourceIPAddressRead,
		Delete: resourceIPAddressDelete,
		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_mask": {
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

func resourceIPAddressCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IPAddress (Inside resourceIPAddressCreate) ")
		data := dataToIPAddress(d)
		logger.Println("[INFO] received V from method data to IPAddress --")
		d.SetId("1")
		go_vthunder.PostIPAddress(client.Token, data, client.Host)

		return resourceIPAddressRead(d, meta)

	}
	return nil
}

func resourceIPAddressRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IPAddress (Inside resourceIPAddressRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetIPAddress(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIPAddressUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIPAddressRead(d, meta)
}

func resourceIPAddressDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIPAddressRead(d, meta)
}
func dataToIPAddress(d *schema.ResourceData) go_vthunder.IPAddress {
	var vc go_vthunder.IPAddress
	var c go_vthunder.IPAddressInstance
	c.IPAddr = d.Get("ip_addr").(string)
	c.IPMask = d.Get("ip_mask").(string)

	vc.UUID = c
	return vc
}
