package vthunder

//vThunder resource SlbRcCacheGlobal

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbRcCacheGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbRcCacheGlobalCreate,
		Update: resourceSlbRcCacheGlobalUpdate,
		Read:   resourceSlbRcCacheGlobalRead,
		Delete: resourceSlbRcCacheGlobalDelete,
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

func resourceSlbRcCacheGlobalCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbRcCacheGlobal (Inside resourceSlbRcCacheGlobalCreate) ")

		data := dataToSlbRcCacheGlobal(d)
		logger.Println("[INFO] received formatted data from method data to SlbRcCacheGlobal --")
		d.SetId("1")
		go_vthunder.PostSlbRcCacheGlobal(client.Token, data, client.Host)

		return resourceSlbRcCacheGlobalRead(d, meta)

	}
	return nil
}
func resourceSlbRcCacheGlobalUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbRcCacheGlobalRead(d, meta)
}

func resourceSlbRcCacheGlobalDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbRcCacheGlobalRead(d, meta)
}

func resourceSlbRcCacheGlobalRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbRcCacheGlobal (Inside resourceSlbRcCacheGlobalRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbRcCacheGlobal(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func dataToSlbRcCacheGlobal(d *schema.ResourceData) go_vthunder.RcCacheGlobal {
	var vc go_vthunder.RcCacheGlobal
	var c go_vthunder.RcCacheGlobalInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable31, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable31
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
