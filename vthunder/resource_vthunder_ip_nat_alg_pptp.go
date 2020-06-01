package vthunder

//vThunder resource IpNatAlgPptp

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpNatAlgPptp() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpNatAlgPptpCreate,
		Update: resourceIpNatAlgPptpUpdate,
		Read:   resourceIpNatAlgPptpRead,
		Delete: resourceIpNatAlgPptpDelete,
		Schema: map[string]*schema.Schema{
			"pptp": {
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

func resourceIpNatAlgPptpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpNatAlgPptp (Inside resourceIpNatAlgPptpCreate) ")

		data := dataToIpNatAlgPptp(d)
		logger.Println("[INFO] received V from method data to IpNatAlgPptp --")
		d.SetId("1")
		go_vthunder.PostIpNatAlgPptp(client.Token, data, client.Host)

		return resourceIpNatAlgPptpRead(d, meta)

	}
	return nil
}

func resourceIpNatAlgPptpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IpNatAlgPptp (Inside resourceIpNatAlgPptpRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetIpNatAlgPptp(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpNatAlgPptpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpNatAlgPptpRead(d, meta)
}

func resourceIpNatAlgPptpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpNatAlgPptpRead(d, meta)
}
func dataToIpNatAlgPptp(d *schema.ResourceData) go_vthunder.NatAlgPptp {
	var vc go_vthunder.NatAlgPptp
	var c go_vthunder.NatAlgPptpInstance
	c.Pptp = d.Get("pptp").(string)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable48, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable48
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}
	vc.UUID = c
	return vc
}
