package thunder

//Thunder resource SnmpServerEnableTrapsSnmp

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSnmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsSnmpCreate,
		Update: resourceSnmpServerEnableTrapsSnmpUpdate,
		Read:   resourceSnmpServerEnableTrapsSnmpRead,
		Delete: resourceSnmpServerEnableTrapsSnmpDelete,
		Schema: map[string]*schema.Schema{
			"linkup": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"linkdown": {
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

func resourceSnmpServerEnableTrapsSnmpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSnmp (Inside resourceSnmpServerEnableTrapsSnmpCreate) ")

		data := dataToSnmpServerEnableTrapsSnmp(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSnmp --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsSnmp(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsSnmpRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsSnmpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSnmp (Inside resourceSnmpServerEnableTrapsSnmpRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSnmp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsSnmpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSnmpRead(d, meta)
}

func resourceSnmpServerEnableTrapsSnmpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSnmpRead(d, meta)
}
func dataToSnmpServerEnableTrapsSnmp(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSnmp {
	var vc go_thunder.SnmpServerEnableTrapsSnmp
	var c go_thunder.SnmpServerEnableTrapsSnmpInstance
	c.Linkup = d.Get("linkup").(int)
	c.All = d.Get("all").(int)
	c.Linkdown = d.Get("linkdown").(int)

	vc.Linkup = c
	return vc
}
