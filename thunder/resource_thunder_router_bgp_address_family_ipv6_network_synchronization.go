package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NetworkSynchronization

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronization() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate,
		Update: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate,
		Read:   resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead,
		Delete: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete,
		Schema: map[string]*schema.Schema{
			"network_synchronization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NetworkSynchronization (Inside resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate) ")
		name1 := d.Get("as_number").(int)
		logger.Println(name1)
		name := strconv.Itoa(name1)

		logger.Println(name)
		data := dataToRouterBgpAddressFamilyIpv6NetworkSynchronization(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NetworkSynchronization --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostRouterBgpAddressFamilyIpv6NetworkSynchronization(client.Token, name, data, client.Host)

		return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(d, meta)

	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NetworkSynchronization (Inside resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NetworkSynchronization(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(d, meta)
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(d, meta)
}
func dataToRouterBgpAddressFamilyIpv6NetworkSynchronization(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NetworkSynchronization {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NetworkSynchronization
	var c go_thunder.RouterBgpAddressFamilyIpv6NetworkSynchronizationInstance
	c.NetworkSynchronization = d.Get("network_synchronization").(int)

	vc.NetworkSynchronization = c
	return vc
}
