package thunder

//Thunder resource SnmpServerEngineID

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEngineID() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEngineIDCreate,
		Update: resourceSnmpServerEngineIDUpdate,
		Read:   resourceSnmpServerEngineIDRead,
		Delete: resourceSnmpServerEngineIDDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"eng_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEngineIDCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEngineID (Inside resourceSnmpServerEngineIDCreate) ")

		data := dataToSnmpServerEngineID(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEngineID --")
		d.SetId("1")
		go_thunder.PostSnmpServerEngineID(client.Token, data, client.Host)

		return resourceSnmpServerEngineIDRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEngineIDRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEngineID (Inside resourceSnmpServerEngineIDRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEngineID(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEngineIDUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEngineIDRead(d, meta)
}

func resourceSnmpServerEngineIDDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEngineIDRead(d, meta)
}
func dataToSnmpServerEngineID(d *schema.ResourceData) go_thunder.SnmpServerEngineID {
	var vc go_thunder.SnmpServerEngineID
	var c go_thunder.SnmpServerEngineIDInstance
	c.EngID = d.Get("eng_id").(string)

	vc.UUID = c
	return vc
}
