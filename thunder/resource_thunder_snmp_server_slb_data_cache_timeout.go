package thunder

//Thunder resource SnmpServerSlbDataCacheTimeout

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerSlbDataCacheTimeout() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerSlbDataCacheTimeoutCreate,
		Update: resourceSnmpServerSlbDataCacheTimeoutUpdate,
		Read:   resourceSnmpServerSlbDataCacheTimeoutRead,
		Delete: resourceSnmpServerSlbDataCacheTimeoutDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"slblimit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerSlbDataCacheTimeoutCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerSlbDataCacheTimeout (Inside resourceSnmpServerSlbDataCacheTimeoutCreate) ")

		data := dataToSnmpServerSlbDataCacheTimeout(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerSlbDataCacheTimeout --")
		d.SetId("1")
		go_thunder.PostSnmpServerSlbDataCacheTimeout(client.Token, data, client.Host)

		return resourceSnmpServerSlbDataCacheTimeoutRead(d, meta)

	}
	return nil
}

func resourceSnmpServerSlbDataCacheTimeoutRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerSlbDataCacheTimeout (Inside resourceSnmpServerSlbDataCacheTimeoutRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerSlbDataCacheTimeout(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerSlbDataCacheTimeoutUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerSlbDataCacheTimeoutRead(d, meta)
}

func resourceSnmpServerSlbDataCacheTimeoutDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerSlbDataCacheTimeoutRead(d, meta)
}
func dataToSnmpServerSlbDataCacheTimeout(d *schema.ResourceData) go_thunder.SnmpServerSlbDataCacheTimeout {
	var vc go_thunder.SnmpServerSlbDataCacheTimeout
	var c go_thunder.SnmpServerSlbDataCacheTimeoutInstance
	c.Slblimit = d.Get("slblimit").(int)

	vc.UUID = c
	return vc
}
