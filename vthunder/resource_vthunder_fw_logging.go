package vthunder

//vThunder resource FwLogging

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwLogging() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwLoggingCreate,
		Update: resourceFwLoggingUpdate,
		Read:   resourceFwLoggingRead,
		Delete: resourceFwLoggingDelete,
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
			"name": {
				Type:        schema.TypeString,
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

func resourceFwLoggingCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwLogging (Inside resourceFwLoggingCreate) ")

		data := dataToFwLogging(d)
		logger.Println("[INFO] received formatted data from method data to FwLogging --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwLogging(client.Token, data, client.Host)

		return resourceFwLoggingRead(d, meta)

	}
	return nil
}

func resourceFwLoggingRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwLogging (Inside resourceFwLoggingRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwLogging(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwLoggingUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwLoggingRead(d, meta)
}

func resourceFwLoggingDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwLoggingRead(d, meta)
}
func dataToFwLogging(d *schema.ResourceData) go_vthunder.FwLogging {
	var vc go_vthunder.FwLogging
	var c go_vthunder.FwLoggingInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.FwLoggingSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.FwLoggingSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.Name = d.Get("name").(string)

	vc.SamplingEnable = c
	return vc
}
