package thunder

//Thunder resource SnmpServerDisableTraps

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerDisableTraps() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerDisableTrapsCreate,
		Update: resourceSnmpServerDisableTrapsUpdate,
		Read:   resourceSnmpServerDisableTrapsRead,
		Delete: resourceSnmpServerDisableTrapsDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"vrrp_a": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"snmp": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gslb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"slb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerDisableTrapsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerDisableTraps (Inside resourceSnmpServerDisableTrapsCreate) ")

		data := dataToSnmpServerDisableTraps(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerDisableTraps --")
		d.SetId("1")
		go_thunder.PostSnmpServerDisableTraps(client.Token, data, client.Host)

		return resourceSnmpServerDisableTrapsRead(d, meta)

	}
	return nil
}

func resourceSnmpServerDisableTrapsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerDisableTraps (Inside resourceSnmpServerDisableTrapsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerDisableTraps(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerDisableTrapsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerDisableTrapsRead(d, meta)
}

func resourceSnmpServerDisableTrapsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerDisableTrapsRead(d, meta)
}
func dataToSnmpServerDisableTraps(d *schema.ResourceData) go_thunder.SnmpServerDisableTraps {
	var vc go_thunder.SnmpServerDisableTraps
	var c go_thunder.SnmpServerDisableTrapsInstance
	c.All = d.Get("all").(int)
	c.SlbChange = d.Get("slb_change").(int)
	c.VrrpA = d.Get("vrrp_a").(int)
	c.Snmp = d.Get("snmp").(int)
	c.Gslb = d.Get("gslb").(int)
	c.Slb = d.Get("slb").(int)

	vc.All = c
	return vc
}
