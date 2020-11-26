package thunder

//Thunder resource SnmpServerEnableTrapsVrrpA

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsVrrpA() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsVrrpACreate,
		Update: resourceSnmpServerEnableTrapsVrrpAUpdate,
		Read:   resourceSnmpServerEnableTrapsVrrpARead,
		Delete: resourceSnmpServerEnableTrapsVrrpADelete,
		Schema: map[string]*schema.Schema{
			"active": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"standby": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all": {
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

func resourceSnmpServerEnableTrapsVrrpACreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsVrrpA (Inside resourceSnmpServerEnableTrapsVrrpACreate) ")

		data := dataToSnmpServerEnableTrapsVrrpA(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsVrrpA --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsVrrpA(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsVrrpARead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsVrrpARead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsVrrpA (Inside resourceSnmpServerEnableTrapsVrrpARead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsVrrpA(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsVrrpAUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsVrrpARead(d, meta)
}

func resourceSnmpServerEnableTrapsVrrpADelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsVrrpARead(d, meta)
}
func dataToSnmpServerEnableTrapsVrrpA(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsVrrpA {
	var vc go_thunder.SnmpServerEnableTrapsVrrpA
	var c go_thunder.SnmpServerEnableTrapsVrrpAInstance
	c.Active = d.Get("active").(int)
	c.Standby = d.Get("standby").(int)
	c.All = d.Get("all").(int)

	vc.Active = c
	return vc
}
