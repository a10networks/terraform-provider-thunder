package vthunder

// vThunder resource Slb HTTP2

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbHTTP2() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbHTTP2Create,
		Update: resourceSlbHTTP2Update,
		Read:   resourceSlbHTTP2Read,
		Delete: resourceSlbHTTP2Delete,

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

func resourceSlbHTTP2Create(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating HTTP2 (Inside resourceSlbHTTP2Create)")

	if client.Host != "" {
		vc := dataToSlbHTTP2(d)
		d.SetId("1")
		go_vthunder.PostSlbHTTP2(client.Token, vc, client.Host)

		return resourceSlbHTTP2Read(d, meta)
	}
	return nil
}

func resourceSlbHTTP2Read(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading HTTP2 (Inside resourceSlbHTTP2Read)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbHTTP2(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No HTTP2 found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbHTTP2Update(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHTTP2Read(d, meta)
}

func resourceSlbHTTP2Delete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHTTP2Read(d, meta)
}

//Utility method to instantiate HTTP2 Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHTTP2(d *schema.ResourceData) go_vthunder.HTTP2 {
	var vc go_vthunder.HTTP2

	var c go_vthunder.HTTP2Instance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable34, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable34
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
