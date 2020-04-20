package vthunder

//vThunder resource SlbSSLForwardProxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSSLForwardProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSSLForwardProxyCreate,
		Update: resourceSlbSSLForwardProxyUpdate,
		Read:   resourceSlbSSLForwardProxyRead,
		Delete: resourceSlbSSLForwardProxyDelete,
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

func resourceSlbSSLForwardProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSSLForwardProxy (Inside resourceSlbSSLForwardProxyCreate) ")

		data := dataToSlbSSLForwardProxy(d)
		logger.Println("[INFO] received V from method data to SlbSSLForwardProxy --")
		d.SetId("1")
		go_vthunder.PostSlbSSLForwardProxy(client.Token, data, client.Host)

		return resourceSlbSSLForwardProxyRead(d, meta)

	}
	return nil
}

func resourceSlbSSLForwardProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSSLForwardProxyRead(d, meta)
}

func resourceSlbSSLForwardProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSSLForwardProxyRead(d, meta)
}

func resourceSlbSSLForwardProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbSSLForwardProxy (Inside resourceSlbSSLForwardProxyRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetSlbSSLForwardProxy(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func dataToSlbSSLForwardProxy(d *schema.ResourceData) go_vthunder.SSLForwardProxy {
	var vc go_vthunder.SSLForwardProxy
	var c go_vthunder.SSLForwardProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable38, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable38
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
