package vthunder

//vThunder resource SlbSip

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSip() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSipCreate,
		Update: resourceSlbSipUpdate,
		Read:   resourceSlbSipRead,
		Delete: resourceSlbSipDelete,
		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSipCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSip (Inside resourceSlbSipCreate) ")

		data := dataToSlbSip(d)
		logger.Println("[INFO] received V from method data to SlbSip --")
		d.SetId("1")
		go_vthunder.PostSlbSip(client.Token, data, client.Host)

		return resourceSlbSipRead(d, meta)

	}
	return nil
}

func resourceSlbSipRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbSip (Inside resourceSlbSipRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbSip(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSipUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSipRead(d, meta)
}

func resourceSlbSipDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSipRead(d, meta)
}

func dataToSlbSip(d *schema.ResourceData) go_vthunder.SlbSip {
	var vc go_vthunder.SlbSip
	var c go_vthunder.SlbSipInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable10, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable10
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
