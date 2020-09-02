package vthunder

//vThunder resource FwLocalLog

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwLocalLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwLocalLogCreate,
		Update: resourceFwLocalLogUpdate,
		Read:   resourceFwLocalLogRead,
		Delete: resourceFwLocalLogDelete,
		Schema: map[string]*schema.Schema{
			"local_logging": {
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

func resourceFwLocalLogCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwLocalLog (Inside resourceFwLocalLogCreate) ")

		data := dataToFwLocalLog(d)
		logger.Println("[INFO] received formatted data from method data to FwLocalLog --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwLocalLog(client.Token, data, client.Host)

		return resourceFwLocalLogRead(d, meta)

	}
	return nil
}

func resourceFwLocalLogRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwLocalLog (Inside resourceFwLocalLogRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwLocalLog(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwLocalLogUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwLocalLogRead(d, meta)
}

func resourceFwLocalLogDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwLocalLogRead(d, meta)
}
func dataToFwLocalLog(d *schema.ResourceData) go_vthunder.FwLocalLog {
	var vc go_vthunder.FwLocalLog
	var c go_vthunder.FwLocalLogInstance
	c.LocalLogging = d.Get("local_logging").(int)

	vc.LocalLogging = c
	return vc
}
