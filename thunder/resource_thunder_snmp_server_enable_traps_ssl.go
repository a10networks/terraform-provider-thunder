package thunder

//Thunder resource SnmpServerEnableTrapsSsl

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSsl() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsSslCreate,
		Update: resourceSnmpServerEnableTrapsSslUpdate,
		Read:   resourceSnmpServerEnableTrapsSslRead,
		Delete: resourceSnmpServerEnableTrapsSslDelete,
		Schema: map[string]*schema.Schema{
			"server_certificate_error": {
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

func resourceSnmpServerEnableTrapsSslCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSsl (Inside resourceSnmpServerEnableTrapsSslCreate) ")

		data := dataToSnmpServerEnableTrapsSsl(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSsl --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsSsl(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsSslRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsSslRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSsl (Inside resourceSnmpServerEnableTrapsSslRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSsl(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsSslUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSslRead(d, meta)
}

func resourceSnmpServerEnableTrapsSslDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsSslRead(d, meta)
}
func dataToSnmpServerEnableTrapsSsl(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSsl {
	var vc go_thunder.SnmpServerEnableTrapsSsl
	var c go_thunder.SnmpServerEnableTrapsSslInstance
	c.ServerCertificateError = d.Get("server_certificate_error").(int)

	vc.ServerCertificateError = c
	return vc
}
