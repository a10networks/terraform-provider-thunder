package thunder

//Thunder resource IpIcmp

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpIcmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpIcmpCreate,
		Update: resourceIpIcmpUpdate,
		Read:   resourceIpIcmpRead,
		Delete: resourceIpIcmpDelete,
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

func resourceIpIcmpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpIcmp (Inside resourceIpIcmpCreate) ")

		data := dataToIpIcmp(d)
		logger.Println("[INFO] received V from method data to IpIcmp --")
		d.SetId("1")
		go_thunder.PostIpIcmp(client.Token, data, client.Host)

		return resourceIpIcmpRead(d, meta)

	}
	return nil
}

func resourceIpIcmpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpIcmp (Inside resourceIpIcmpRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpIcmp(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpIcmpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpIcmpRead(d, meta)
}

func resourceIpIcmpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpIcmpRead(d, meta)
}
func dataToIpIcmp(d *schema.ResourceData) go_thunder.Icmp {
	var vc go_thunder.Icmp
	var c go_thunder.IcmpInstance
	c.Redirect = d.Get("redirect").(int)
	c.Unreachable = d.Get("unreachable").(int)

	vc.UUID = c
	return vc
}
