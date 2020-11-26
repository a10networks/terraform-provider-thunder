package thunder

//Thunder resource SnmpServerContact

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerContact() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerContactCreate,
		Update: resourceSnmpServerContactUpdate,
		Read:   resourceSnmpServerContactRead,
		Delete: resourceSnmpServerContactDelete,
		Schema: map[string]*schema.Schema{
			"contact_name": {
				Type:        schema.TypeString,
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

func resourceSnmpServerContactCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerContact (Inside resourceSnmpServerContactCreate) ")

		data := dataToSnmpServerContact(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerContact --")
		d.SetId("1")
		go_thunder.PostSnmpServerContact(client.Token, data, client.Host)

		return resourceSnmpServerContactRead(d, meta)

	}
	return nil
}

func resourceSnmpServerContactRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerContact (Inside resourceSnmpServerContactRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerContact(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerContactUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerContactRead(d, meta)
}

func resourceSnmpServerContactDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerContactRead(d, meta)
}
func dataToSnmpServerContact(d *schema.ResourceData) go_thunder.SnmpServerContact {
	var vc go_thunder.SnmpServerContact
	var c go_thunder.SnmpServerContactInstance
	c.ContactName = d.Get("contact_name").(string)

	vc.ContactName = c
	return vc
}
