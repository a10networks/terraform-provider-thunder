package vthunder

//vThunder resource FwApplyChanges

import (
	"github.com/go_vthunder/vthunder"
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
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwApplyChanges (Inside resourceFwApplyChangesCreate) ")

		data := dataToFwApplyChanges(d)
		logger.Println("[INFO] received formatted data from method data to FwApplyChanges --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwApplyChanges(client.Token, data, client.Host)

		return resourceFwApplyChangesRead(d, meta)

	}
	return nil
}

func resourceFwApplyChangesRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwApplyChanges (Inside resourceFwApplyChangesRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwApplyChanges(client.Token, client.Host)
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
func dataToFwApplyChanges(d *schema.ResourceData) go_vthunder.FwApplyChanges {
	var vc go_vthunder.FwApplyChanges
	var c go_vthunder.FwApplyChangesInstance
	c.Forced = d.Get("forced").(int)

	vc.Forced = c
	return vc
}
