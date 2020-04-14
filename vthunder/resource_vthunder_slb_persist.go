package vthunder

//vThunder resource Slbpersist

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbpersist() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbpersistCreate,
		Update: resourceSlbpersistUpdate,
		Read:   resourceSlbpersistRead,
		Delete: resourceSlbpersistDelete,
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

func resourceSlbpersistCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating Slbpersist (Inside resourceSlbpersistCreate) ")

		data := dataToSlbpersist(d)
		logger.Println("[INFO] received formatted data from method data to Slbpersist --")
		d.SetId("1")
		go_vthunder.PostSlbpersist(client.Token, data, client.Host)

		return resourceSlbpersistRead(d, meta)

	}
	return nil
}

func resourceSlbpersistRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading Slbpersist (Inside resourceSlbpersistRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbpersist(client.Token, client.Host)
		if data == nil {

			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbpersistUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbpersistRead(d, meta)
}

func resourceSlbpersistDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbpersistRead(d, meta)
}
func dataToSlbpersist(d *schema.ResourceData) go_vthunder.SlbPersist {
	var vc go_vthunder.SlbPersist
	var c go_vthunder.SlbPersistInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable26, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable26
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
