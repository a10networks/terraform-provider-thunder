package thunder

//Thunder resource IpNatIcmp

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpNatIcmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpNatIcmpCreate,
		Update: resourceIpNatIcmpUpdate,
		Read:   resourceIpNatIcmpRead,
		Delete: resourceIpNatIcmpDelete,
		Schema: map[string]*schema.Schema{
			"respond_to_ping": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"always_source_nat_errors": {
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

func resourceIpNatIcmpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpNatIcmp (Inside resourceIpNatIcmpCreate) ")

		data := dataToIpNatIcmp(d)
		logger.Println("[INFO] received V from method data to IpNatIcmp --")
		d.SetId("1")
		go_thunder.PostIpNatIcmp(client.Token, data, client.Host)

		return resourceIpNatIcmpRead(d, meta)

	}
	return nil
}

func resourceIpNatIcmpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpNatIcmp (Inside resourceIpNatIcmpRead)")

	if client.Host != "" {

		data, err := go_thunder.GetIpNatIcmp(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpNatIcmpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpNatIcmpRead(d, meta)
}

func resourceIpNatIcmpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpNatIcmpRead(d, meta)
}
func dataToIpNatIcmp(d *schema.ResourceData) go_thunder.NatIcmp {
	var vc go_thunder.NatIcmp
	var c go_thunder.NatIcmpInstance
	c.RespondToPing = d.Get("respond_to_ping").(int)
	c.AlwaysSourceNatErrors = d.Get("always_source_nat_errors").(int)

	vc.UUID = c
	return vc
}
