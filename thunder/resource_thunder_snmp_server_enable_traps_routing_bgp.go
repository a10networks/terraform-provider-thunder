package thunder

//Thunder resource SnmpServerEnableTrapsRoutingBgp

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerEnableTrapsRoutingBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsRoutingBgpCreate,
		Update: resourceSnmpServerEnableTrapsRoutingBgpUpdate,
		Read:   resourceSnmpServerEnableTrapsRoutingBgpRead,
		Delete: resourceSnmpServerEnableTrapsRoutingBgpDelete,
		Schema: map[string]*schema.Schema{
			"bgp_established_notification": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"bgp_backward_trans_notification": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsRoutingBgpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsRoutingBgp (Inside resourceSnmpServerEnableTrapsRoutingBgpCreate) ")

		data := dataToSnmpServerEnableTrapsRoutingBgp(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsRoutingBgp --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsRoutingBgp(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsRoutingBgpRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsRoutingBgpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsRoutingBgp (Inside resourceSnmpServerEnableTrapsRoutingBgpRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsRoutingBgp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsRoutingBgpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsRoutingBgpRead(d, meta)
}

func resourceSnmpServerEnableTrapsRoutingBgpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsRoutingBgpRead(d, meta)
}
func dataToSnmpServerEnableTrapsRoutingBgp(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsRoutingBgp {
	var vc go_thunder.SnmpServerEnableTrapsRoutingBgp
	var c go_thunder.SnmpServerEnableTrapsRoutingBgpInstance
	c.BgpEstablishedNotification = d.Get("bgp_established_notification").(int)
	c.BgpBackwardTransNotification = d.Get("bgp_backward_trans_notification").(int)

	vc.BgpEstablishedNotification = c
	return vc
}
