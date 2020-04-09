package vthunder

// vThunder resource Slb CrlSrcip

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbCrlSrcip() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbCrlSrcipCreate,
		Update: resourceSlbCrlSrcipUpdate,
		Read:   resourceSlbCrlSrcipRead,
		Delete: resourceSlbCrlSrcipDelete,

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

func resourceSlbCrlSrcipCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating crl-srcip (Inside resourceSlbCrlSrcipCreate)")

	if client.Host != "" {
		vc := dataToSlbCrlSrcip(d)
		d.SetId("1")
		go_vthunder.PostSlbCrlSrcip(client.Token, vc, client.Host)

		return resourceSlbCrlSrcipRead(d, meta)
	}
	return nil
}

func resourceSlbCrlSrcipRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading crl-srcip (Inside resourceSlbCrlSrcipRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbCrlSrcip(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No crl-srcip found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbCrlSrcipUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbCrlSrcipRead(d, meta)
}

func resourceSlbCrlSrcipDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbCrlSrcipRead(d, meta)
}

//Utility method to instantiate CrlSrcip Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbCrlSrcip(d *schema.ResourceData) go_vthunder.CrlSrcip {
	var vc go_vthunder.CrlSrcip

	var c go_vthunder.CrlSrcipInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable8, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable8
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
