package vthunder

//vThunder resource FwTopKRules

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwTopKRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwTopKRulesCreate,
		Update: resourceFwTopKRulesUpdate,
		Read:   resourceFwTopKRulesRead,
		Delete: resourceFwTopKRulesDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwTopKRulesCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwTopKRules (Inside resourceFwTopKRulesCreate) ")

		data := dataToFwTopKRules(d)
		logger.Println("[INFO] received formatted data from method data to FwTopKRules --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwTopKRules(client.Token, data, client.Host)

		return resourceFwTopKRulesRead(d, meta)

	}
	return nil
}

func resourceFwTopKRulesRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwTopKRules (Inside resourceFwTopKRulesRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwTopKRules(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwTopKRulesUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTopKRulesRead(d, meta)
}

func resourceFwTopKRulesDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwTopKRulesRead(d, meta)
}
func dataToFwTopKRules(d *schema.ResourceData) go_vthunder.FwTopKRules {
	var vc go_vthunder.FwTopKRules
	var c go_vthunder.FwTopKRulesInstance

	vc.UUID = c
	return vc
}
