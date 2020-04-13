package vthunder

// vThunder resource Slb DNSResponseRateLimiting

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbDNSResponseRateLimiting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbDNSResponseRateLimitingCreate,
		Update: resourceSlbDNSResponseRateLimitingUpdate,
		Read:   resourceSlbDNSResponseRateLimitingRead,
		Delete: resourceSlbDNSResponseRateLimitingDelete,

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

func resourceSlbDNSResponseRateLimitingCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating DNS Response Rate Limiting (Inside resourceSlbDNSResponseRateLimitingCreate)")

	if client.Host != "" {
		vc := dataToSlbDNSResponseRateLimiting(d)
		d.SetId("1")
		go_vthunder.PostSlbDNSResponseRateLimiting(client.Token, vc, client.Host)

		return resourceSlbDNSResponseRateLimitingRead(d, meta)
	}
	return nil
}

func resourceSlbDNSResponseRateLimitingRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DNS Response Rate Limiting (Inside resourceSlbDNSResponseRateLimitingRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbDNSResponseRateLimiting(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No DNS Response Rate Limiting found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbDNSResponseRateLimitingUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbDNSResponseRateLimitingRead(d, meta)
}

func resourceSlbDNSResponseRateLimitingDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbDNSResponseRateLimitingRead(d, meta)
}

//Utility method to instantiate DNSResponseRateLimiting Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbDNSResponseRateLimiting(d *schema.ResourceData) go_vthunder.SlbDNSResponseRateLimiting {
	var vc go_vthunder.SlbDNSResponseRateLimiting

	var c go_vthunder.SlbDNSResponseRateLimitingInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable17, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable17
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
