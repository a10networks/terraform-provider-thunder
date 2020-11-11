package thunder

//Thunder resource SnmpServerEnableTrapsGslb

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsGslb() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsGslbCreate,
		Update: resourceSnmpServerEnableTrapsGslbUpdate,
		Read:   resourceSnmpServerEnableTrapsGslbRead,
		Delete: resourceSnmpServerEnableTrapsGslbDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"group": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"zone": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"site": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsGslbCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsGslb (Inside resourceSnmpServerEnableTrapsGslbCreate) ")

		data := dataToSnmpServerEnableTrapsGslb(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsGslb --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsGslb(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsGslbRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsGslbRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsGslb (Inside resourceSnmpServerEnableTrapsGslbRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsGslb(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsGslbUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsGslbRead(d, meta)
}

func resourceSnmpServerEnableTrapsGslbDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsGslbRead(d, meta)
}
func dataToSnmpServerEnableTrapsGslb(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsGslb {
	var vc go_thunder.SnmpServerEnableTrapsGslb
	var c go_thunder.SnmpServerEnableTrapsGslbInstance
	c.All = d.Get("all").(int)
	c.Group = d.Get("group").(int)
	c.Zone = d.Get("zone").(int)
	c.Site = d.Get("site").(int)
	c.ServiceIP = d.Get("service_ip").(int)

	vc.All = c
	return vc
}
