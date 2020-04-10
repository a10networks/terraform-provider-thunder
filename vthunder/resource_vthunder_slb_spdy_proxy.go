package vthunder

//vThunder resource SlbSpdyProxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSpdyProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSpdyProxyCreate,
		Update: resourceSlbSpdyProxyUpdate,
		Read:   resourceSlbSpdyProxyRead,
		Delete: resourceSlbSpdyProxyDelete,
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

func resourceSlbSpdyProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSpdyProxy (Inside resourceSlbSpdyProxyCreate) ")

		data := dataToSlbSpdyProxy(d)
		logger.Println("[INFO] received V from method data to SlbSpdyProxy --")
		d.SetId("1")
		go_vthunder.PostSlbSpdyProxy(client.Token, data, client.Host)

		return resourceSlbSpdyProxyRead(d, meta)

	}
	return nil
}

func resourceSlbSpdyProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbSpdyProxy (Inside resourceSlbSpdyProxyRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbSpdyProxy(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}
func resourceSlbSpdyProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSpdyProxyRead(d, meta)
}

func resourceSlbSpdyProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSpdyProxyRead(d, meta)
}

func dataToSlbSpdyProxy(d *schema.ResourceData) go_vthunder.SpdyProxy {
	var vc go_vthunder.SpdyProxy
	var c go_vthunder.SpdyProxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable14, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable14
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
