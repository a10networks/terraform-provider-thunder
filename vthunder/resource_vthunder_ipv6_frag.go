package vthunder

//vThunder resource Ipv6Frag

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpv6Frag() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpv6FragCreate,
		Update: resourceIpv6FragUpdate,
		Read:   resourceIpv6FragRead,
		Delete: resourceIpv6FragDelete,
		Schema: map[string]*schema.Schema{
			"frag_timeout": {
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
		},
	}
}

func resourceIpv6FragCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Ipv6Frag (Inside resourceIpv6FragCreate) ")

		data := dataToIpv6Frag(d)
		logger.Println("[INFO] received V from method data to Ipv6Frag --")
		d.SetId("1")
		go_vthunder.PostIpv6Frag(client.Token, data, client.Host)

		return resourceIpv6FragRead(d, meta)

	}
	return nil
}

func resourceIpv6FragRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Ipv6Frag (Inside resourceIpv6FragRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpv6Frag(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpv6FragUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpv6FragRead(d, meta)
}

func resourceIpv6FragDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpv6FragRead(d, meta)
}
func dataToIpv6Frag(d *schema.ResourceData) go_vthunder.IPv6Frag {
	var vc go_vthunder.IPv6Frag
	var c go_vthunder.IPv6FragInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable49, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable49
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.FragTimeout = d.Get("frag_timeout").(int)

	vc.UUID = c
	return vc
}
