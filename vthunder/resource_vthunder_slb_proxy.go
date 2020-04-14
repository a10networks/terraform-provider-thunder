package vthunder

//vThunder resource SlbProxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbProxyCreate,
		Update: resourceSlbProxyUpdate,
		Read:   resourceSlbProxyRead,
		Delete: resourceSlbProxyDelete,
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

func resourceSlbProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbProxy (Inside resourceSlbProxyCreate) ")

		data := dataToSlbProxy(d)
		logger.Println("[INFO] received formatted data from method data to SlbProxy --")
		d.SetId("1")
		go_vthunder.PostSlbProxy(client.Token, data, client.Host)

		return resourceSlbProxyRead(d, meta)

	}
	return nil
}

func resourceSlbProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbProxy (Inside resourceSlbProxyRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbProxy(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbProxyRead(d, meta)
}

func resourceSlbProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbProxyRead(d, meta)
}

func dataToSlbProxy(d *schema.ResourceData) go_vthunder.Proxy {
	var vc go_vthunder.Proxy
	var c go_vthunder.ProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable29, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable29
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
