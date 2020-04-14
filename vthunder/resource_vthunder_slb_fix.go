package vthunder

// vThunder resource Slb Fix

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbFix() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbFixCreate,
		Update: resourceSlbFixUpdate,
		Read:   resourceSlbFixRead,
		Delete: resourceSlbFixDelete,

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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceSlbFixCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating fix (Inside resourceSlbFixCreate)")

	if client.Host != "" {
		vc := dataToSlbFix(d)
		d.SetId("1")
		go_vthunder.PostSlbFix(client.Token, vc, client.Host)

		return resourceSlbFixRead(d, meta)
	}
	return nil
}

func resourceSlbFixRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading fix (Inside resourceSlbFixRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbFix(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No fix found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbFixUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFixRead(d, meta)
}

func resourceSlbFixDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFixRead(d, meta)
}

//Utility method to instantiate Fix Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFix(d *schema.ResourceData) go_vthunder.SlbFix {
	var vc go_vthunder.SlbFix

	var c go_vthunder.SlbFixInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable19, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable19
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
