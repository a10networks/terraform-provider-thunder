package vthunder

// vThunder resource Slb DNSCache

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbDNSCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbDNSCacheCreate,
		Update: resourceSlbDNSCacheUpdate,
		Read:   resourceSlbDNSCacheRead,
		Delete: resourceSlbDNSCacheDelete,

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

func resourceSlbDNSCacheCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating DNSCache (Inside resourceSlbDNSCacheCreate)")

	if client.Host != "" {
		vc := dataToSlbDNSCache(d)
		d.SetId("1")
		go_vthunder.PostSlbDNSCache(client.Token, vc, client.Host)

		return resourceSlbDNSCacheRead(d, meta)
	}
	return nil
}

func resourceSlbDNSCacheRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading DNSCache (Inside resourceSlbDNSCacheRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbDNSCache(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No DNSCache found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbDNSCacheUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbDNSCacheRead(d, meta)
}

func resourceSlbDNSCacheDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbDNSCacheRead(d, meta)
}

//Utility method to instantiate DNSCache Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbDNSCache(d *schema.ResourceData) go_vthunder.DNSCache {
	var vc go_vthunder.DNSCache

	var c go_vthunder.DNSCacheInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable16, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable16
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
