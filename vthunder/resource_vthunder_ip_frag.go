package vthunder

//vThunder resource IpFrag

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpFrag() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpFragCreate,
		Update: resourceIpFragUpdate,
		Read:   resourceIpFragRead,
		Delete: resourceIpFragDelete,
		Schema: map[string]*schema.Schema{
			"cpu_threshold": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"high": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"low": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"max_packets_per_reassembly": {
				Type:        schema.TypeInt,
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
			"buff": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"max_reassembly_sessions": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpFragCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpFrag (Inside resourceIpFragCreate) ")

		data := dataToIpFrag(d)
		logger.Println("[INFO] received V from method data to IpFrag --")
		d.SetId("1")
		go_vthunder.PostIpFrag(client.Token, data, client.Host)

		return resourceIpFragRead(d, meta)

	}
	return nil
}

func resourceIpFragRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IpFrag (Inside resourceIpFragRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpFrag(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpFragUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpFragRead(d, meta)
}

func resourceIpFragDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpFragRead(d, meta)
}
func dataToIpFrag(d *schema.ResourceData) go_vthunder.Frag {
	var vc go_vthunder.Frag
	var c go_vthunder.FragInstance

	c.MaxReassemblySessions = d.Get("max_reassembly_sessions").(int)
	c.MaxPacketsPerReassembly = d.Get("max_packets_per_reassembly").(int)

	var obj1 go_vthunder.CPUThreshold
	prefix := "cpu_threshold.0."
	obj1.Low = d.Get(prefix + "low").(int)
	obj1.High = d.Get(prefix + "high").(int)
	c.High = obj1

	c.Timeout = d.Get("timeout").(int)
	c.Buff = d.Get("buff").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable47, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj2 go_vthunder.SamplingEnable47
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj2)
	}
	vc.UUID = c
	return vc
}
