package vthunder

//vThunder resource FwApp

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAppCreate,
		Update: resourceFwAppUpdate,
		Read:   resourceFwAppRead,
		Delete: resourceFwAppDelete,
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

func resourceFwAppCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwApp (Inside resourceFwAppCreate) ")

		data := dataToFwApp(d)
		logger.Println("[INFO] received formatted data from method data to FwApp --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwApp(client.Token, data, client.Host)

		return resourceFwAppRead(d, meta)

	}
	return nil
}

func resourceFwAppRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwApp (Inside resourceFwAppRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwApp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAppUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAppRead(d, meta)
}

func resourceFwAppDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAppRead(d, meta)
}
func dataToFwApp(d *schema.ResourceData) go_vthunder.FwApp {
	var vc go_vthunder.FwApp
	var c go_vthunder.FwAppInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.FwAppSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.FwAppSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.SamplingEnable = c
	return vc
}
