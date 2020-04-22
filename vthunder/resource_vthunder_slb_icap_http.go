package vthunder

//vThunder resource SlbIcapHTTP

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbIcapHTTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbIcapHTTPCreate,
		Update: resourceSlbIcapHTTPUpdate,
		Read:   resourceSlbIcapHTTPRead,
		Delete: resourceSlbIcapHTTPDelete,
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

func resourceSlbIcapHTTPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbIcapHTTP (Inside resourceSlbIcapHTTPCreate) ")

		data := dataToSlbIcapHTTP(d)
		logger.Println("[INFO] received V from method data to SlbIcapHTTP --")
		d.SetId("1")
		go_vthunder.PostSlbIcapHTTP(client.Token, data, client.Host)

		return resourceSlbIcapHTTPRead(d, meta)

	}
	return nil
}

func resourceSlbIcapHTTPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbIcapHTTP (Inside resourceSlbIcapHTTPRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetSlbIcapHTTP(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}
func resourceSlbIcapHTTPUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbIcapHTTPRead(d, meta)
}

func resourceSlbIcapHTTPDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbIcapHTTPRead(d, meta)
}
func dataToSlbIcapHTTP(d *schema.ResourceData) go_vthunder.IcapHTTP {
	var vc go_vthunder.IcapHTTP
	var c go_vthunder.IcapHTTPInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable40, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable40
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
