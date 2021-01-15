package thunder

//Thunder resource RouterBgpNetworkSynchronization

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgpNetworkSynchronization() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpNetworkSynchronizationCreate,
		Update: resourceRouterBgpNetworkSynchronizationUpdate,
		Read:   resourceRouterBgpNetworkSynchronizationRead,
		Delete: resourceRouterBgpNetworkSynchronizationDelete,
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

func resourceRouterBgpNetworkSynchronizationCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNetworkSynchronization (Inside resourceRouterBgpNetworkSynchronizationCreate) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNetworkSynchronization(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNetworkSynchronization --")
		d.SetId(strconv.Itoa(name1))
		go_thunder.PostRouterBgpNetworkSynchronization(client.Token, name, data, client.Host)

		return resourceRouterBgpNetworkSynchronizationRead(d, meta)

	}
	return nil
}

func resourceRouterBgpNetworkSynchronizationRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading RouterBgpNetworkSynchronization (Inside resourceRouterBgpNetworkSynchronizationRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNetworkSynchronization(client.Token, name1, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return err
	}
	return nil
}

func resourceRouterBgpNetworkSynchronizationUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpNetworkSynchronizationRead(d, meta)
}

func resourceRouterBgpNetworkSynchronizationDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceRouterBgpNetworkSynchronizationRead(d, meta)
}
func dataToRouterBgpNetworkSynchronization(d *schema.ResourceData) go_thunder.RouterBgpNetworkSynchronization {
	var vc go_thunder.RouterBgpNetworkSynchronization
	var c go_thunder.RouterBgpNetworkSynchronizationInstance
	c.NetworkSynchronization = d.Get("network_synchronization").(int)

	vc.NetworkSynchronization = c
	return vc
}
