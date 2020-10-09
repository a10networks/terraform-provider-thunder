package thunder

//Thunder resource FwFullConeSession

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwFullConeSession() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwFullConeSessionCreate,
		Update: resourceFwFullConeSessionUpdate,
		Read:   resourceFwFullConeSessionRead,
		Delete: resourceFwFullConeSessionDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwFullConeSessionCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwFullConeSession (Inside resourceFwFullConeSessionCreate) ")

		data := dataToFwFullConeSession(d)
		logger.Println("[INFO] received formatted data from method data to FwFullConeSession --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwFullConeSession(client.Token, data, client.Host)

		return resourceFwFullConeSessionRead(d, meta)

	}
	return nil
}

func resourceFwFullConeSessionRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwFullConeSession (Inside resourceFwFullConeSessionRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwFullConeSession(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwFullConeSessionUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwFullConeSessionRead(d, meta)
}

func resourceFwFullConeSessionDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwFullConeSessionRead(d, meta)
}
func dataToFwFullConeSession(d *schema.ResourceData) go_thunder.FwFullConeSession {
	var vc go_thunder.FwFullConeSession
	var c go_thunder.FwFullConeSessionInstance

	vc.UUID = c
	return vc
}
