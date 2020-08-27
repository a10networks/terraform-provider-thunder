package vthunder

//vThunder resource FwAlgRtsp

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgRtsp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgRtspCreate,
		Update: resourceFwAlgRtspUpdate,
		Read:   resourceFwAlgRtspRead,
		Delete: resourceFwAlgRtspDelete,
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

func resourceFwAlgRtspCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgRtsp (Inside resourceFwAlgRtspCreate) ")

		data := dataToFwAlgRtsp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgRtsp --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwAlgRtsp(client.Token, data, client.Host)

		return resourceFwAlgRtspRead(d, meta)

	}
	return nil
}

func resourceFwAlgRtspRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwAlgRtsp (Inside resourceFwAlgRtspRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwAlgRtsp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgRtspUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgRtspRead(d, meta)
}

func resourceFwAlgRtspDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgRtspRead(d, meta)
}
func dataToFwAlgRtsp(d *schema.ResourceData) go_vthunder.FwAlgRtsp {
	var vc go_vthunder.FwAlgRtsp
	var c go_vthunder.FwAlgRtspInstance
	c.DefaultPortDisable = d.Get("default_port_disable").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.FwAlgRtspSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.FwAlgRtspSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.DefaultPortDisable = c
	return vc
}
