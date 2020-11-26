package thunder

//Thunder resource SnmpServerLocation

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSnmpServerLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerLocationCreate,
		Update: resourceSnmpServerLocationUpdate,
		Read:   resourceSnmpServerLocationRead,
		Delete: resourceSnmpServerLocationDelete,
		Schema: map[string]*schema.Schema{
			"loc": {
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

func resourceSnmpServerLocationCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerLocation (Inside resourceSnmpServerLocationCreate) ")

		data := dataToSnmpServerLocation(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerLocation --")
		d.SetId("1")
		go_thunder.PostSnmpServerLocation(client.Token, data, client.Host)

		return resourceSnmpServerLocationRead(d, meta)

	}
	return nil
}

func resourceSnmpServerLocationRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerLocation (Inside resourceSnmpServerLocationRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerLocation(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerLocationUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerLocationRead(d, meta)
}

func resourceSnmpServerLocationDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerLocationRead(d, meta)
}
func dataToSnmpServerLocation(d *schema.ResourceData) go_thunder.SnmpServerLocation {
	var vc go_thunder.SnmpServerLocation
	var c go_thunder.SnmpServerLocationInstance
	c.Loc = d.Get("loc").(string)

	vc.Loc = c
	return vc
}
