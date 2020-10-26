package thunder

//Thunder resource IpNatGlobal

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpNatGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpNatGlobalCreate,
		Update: resourceIpNatGlobalUpdate,
		Read:   resourceIpNatGlobalRead,
		Delete: resourceIpNatGlobalDelete,
		Schema: map[string]*schema.Schema{
			"reset_idle_tcp_conn": {
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

func resourceIpNatGlobalCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpNatGlobal (Inside resourceIpNatGlobalCreate) ")

		data := dataToIpNatGlobal(d)
		logger.Println("[INFO] received V from method data to IpNatGlobal --")
		d.SetId("1")
		go_thunder.PostIpNatGlobal(client.Token, data, client.Host)

		return resourceIpNatGlobalRead(d, meta)

	}
	return nil
}

func resourceIpNatGlobalRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IpNatGlobal (Inside resourceIpNatGlobalRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpNatGlobal(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpNatGlobalUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpNatGlobalRead(d, meta)
}

func resourceIpNatGlobalDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpNatGlobalRead(d, meta)
}
func dataToIpNatGlobal(d *schema.ResourceData) go_thunder.NatGlobal {
	var vc go_thunder.NatGlobal
	var c go_thunder.NatGlobalInstance

	c.ResetIdleTcpConn = d.Get("reset_idle_tcp_conn").(int)

	vc.UUID = c
	return vc
}
