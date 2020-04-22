package vthunder

//vThunder resource SlL4

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlL4() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbL4Create,
		Update: resourceSlbL4Update,
		Read:   resourceSlbL4Read,
		Delete: resourceSlbL4Delete,
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

func resourceSlbL4Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlL4 (Inside resourceSlL4Create) ")

		data := dataToSlL4(d)
		logger.Println("[INFO] received V from method data to SlL4 --")
		d.SetId("1")
		go_vthunder.PostSlbL4(client.Token, data, client.Host)

		return resourceSlbL4Read(d, meta)

	}
	return nil
}

func resourceSlbL4Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlL4 (Inside resourceSlL4Read)")

	if client.Host != "" {
		data, err := go_vthunder.GetSlbL4(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbL4Update(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbL4Read(d, meta)
}

func resourceSlbL4Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbL4Read(d, meta)
}

func dataToSlL4(d *schema.ResourceData) go_vthunder.L4 {
	var vc go_vthunder.L4
	var c go_vthunder.L4Instance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable42, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable42
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
