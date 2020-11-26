package thunder

//Thunder resource SnmpServerEnableTrapsNetwork

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerEnableTrapsNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsNetworkCreate,
		Update: resourceSnmpServerEnableTrapsNetworkUpdate,
		Read:   resourceSnmpServerEnableTrapsNetworkRead,
		Delete: resourceSnmpServerEnableTrapsNetworkDelete,
		Schema: map[string]*schema.Schema{
			"trunk_port_threshold": {
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

func resourceSnmpServerEnableTrapsNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsNetwork (Inside resourceSnmpServerEnableTrapsNetworkCreate) ")

		data := dataToSnmpServerEnableTrapsNetwork(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsNetwork --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsNetwork(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsNetworkRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsNetworkRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsNetwork (Inside resourceSnmpServerEnableTrapsNetworkRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsNetwork(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsNetworkUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsNetworkRead(d, meta)
}

func resourceSnmpServerEnableTrapsNetworkDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsNetworkRead(d, meta)
}
func dataToSnmpServerEnableTrapsNetwork(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsNetwork {
	var vc go_thunder.SnmpServerEnableTrapsNetwork
	var c go_thunder.SnmpServerEnableTrapsNetworkInstance
	c.TrunkPortThreshold = d.Get("trunk_port_threshold").(int)

	vc.TrunkPortThreshold = c
	return vc
}
