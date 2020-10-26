package thunder

//Thunder resource FwApplyChanges

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwApplyChanges() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwApplyChangesCreate,
		Update: resourceFwApplyChangesUpdate,
		Read:   resourceFwApplyChangesRead,
		Delete: resourceFwApplyChangesDelete,
		Schema: map[string]*schema.Schema{
			"forced": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwApplyChangesCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwApplyChanges (Inside resourceFwApplyChangesCreate) ")

		data := dataToFwApplyChanges(d)
		logger.Println("[INFO] received formatted data from method data to FwApplyChanges --")
		d.SetId(strconv.Itoa('1'))
		go_thunder.PostFwApplyChanges(client.Token, data, client.Host)

		return resourceFwApplyChangesRead(d, meta)

	}
	return nil
}

func resourceFwApplyChangesRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading FwApplyChanges (Inside resourceFwApplyChangesRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetFwApplyChanges(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwApplyChangesUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwApplyChangesRead(d, meta)
}

func resourceFwApplyChangesDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwApplyChangesRead(d, meta)
}
func dataToFwApplyChanges(d *schema.ResourceData) go_thunder.FwApplyChanges {
	var vc go_thunder.FwApplyChanges
	var c go_thunder.FwApplyChangesInstance
	c.Forced = d.Get("forced").(int)

	vc.Forced = c
	return vc
}
