package vthunder

//vThunder resource FwAlgSip

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgSip() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgSipCreate,
		Update: resourceFwAlgSipUpdate,
		Read:   resourceFwAlgSipRead,
		Delete: resourceFwAlgSipDelete,
		Schema: map[string]*schema.Schema{
			"default_port_disable": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
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

func resourceFwAlgSipCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgSip (Inside resourceFwAlgSipCreate) ")

		data := dataToFwAlgSip(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgSip --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwAlgSip(client.Token, data, client.Host)

		return resourceFwAlgSipRead(d, meta)

	}
	return nil
}

func resourceFwAlgSipRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwAlgSip (Inside resourceFwAlgSipRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwAlgSip(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgSipUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgSipRead(d, meta)
}

func resourceFwAlgSipDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgSipRead(d, meta)
}
func dataToFwAlgSip(d *schema.ResourceData) go_vthunder.FwAlgSip {
	var vc go_vthunder.FwAlgSip
	var c go_vthunder.FwAlgSipInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.FwAlgSipSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.FwAlgSipSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.DefaultPortDisable = c
	return vc
}
