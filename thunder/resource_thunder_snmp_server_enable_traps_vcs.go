package thunder

//Thunder resource SnmpServerEnableTrapsVcs

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsVcs() *schema.Resource {
	return &schema.Resource{
		Create: resourceSnmpServerEnableTrapsVcsCreate,
		Update: resourceSnmpServerEnableTrapsVcsUpdate,
		Read:   resourceSnmpServerEnableTrapsVcsRead,
		Delete: resourceSnmpServerEnableTrapsVcsDelete,
		Schema: map[string]*schema.Schema{
			"state_change": {
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

func resourceSnmpServerEnableTrapsVcsCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsVcs (Inside resourceSnmpServerEnableTrapsVcsCreate) ")

		data := dataToSnmpServerEnableTrapsVcs(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsVcs --")
		d.SetId("1")
		go_thunder.PostSnmpServerEnableTrapsVcs(client.Token, data, client.Host)

		return resourceSnmpServerEnableTrapsVcsRead(d, meta)

	}
	return nil
}

func resourceSnmpServerEnableTrapsVcsRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SnmpServerEnableTrapsVcs (Inside resourceSnmpServerEnableTrapsVcsRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsVcs(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSnmpServerEnableTrapsVcsUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsVcsRead(d, meta)
}

func resourceSnmpServerEnableTrapsVcsDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSnmpServerEnableTrapsVcsRead(d, meta)
}
func dataToSnmpServerEnableTrapsVcs(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsVcs {
	var vc go_thunder.SnmpServerEnableTrapsVcs
	var c go_thunder.SnmpServerEnableTrapsVcsInstance
	c.StateChange = d.Get("state_change").(int)

	vc.StateChange = c
	return vc
}
