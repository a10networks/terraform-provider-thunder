package vthunder

//vThunder resource FwResourceUsage

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwResourceUsage() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwResourceUsageCreate,
		Update: resourceFwResourceUsageUpdate,
		Read:   resourceFwResourceUsageRead,
		Delete: resourceFwResourceUsageDelete,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwResourceUsageCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwResourceUsage (Inside resourceFwResourceUsageCreate) ")

		data := dataToFwResourceUsage(d)
		logger.Println("[INFO] received formatted data from method data to FwResourceUsage --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwResourceUsage(client.Token, data, client.Host)

		return resourceFwResourceUsageRead(d, meta)

	}
	return nil
}

func resourceFwResourceUsageRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwResourceUsage (Inside resourceFwResourceUsageRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwResourceUsage(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwResourceUsageUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwResourceUsageRead(d, meta)
}

func resourceFwResourceUsageDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwResourceUsageRead(d, meta)
}
func dataToFwResourceUsage(d *schema.ResourceData) go_vthunder.FwResourceUsage {
	var vc go_vthunder.FwResourceUsage
	var c go_vthunder.FwResourceUsageInstance

	vc.UUID = c
	return vc
}
