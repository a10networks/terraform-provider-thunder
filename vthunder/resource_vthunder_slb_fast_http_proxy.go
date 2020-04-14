package vthunder

// vThunder resource Slb FastHttpProxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbFastHttpProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbFastHttpProxyCreate,
		Update: resourceSlbFastHttpProxyUpdate,
		Read:   resourceSlbFastHttpProxyRead,
		Delete: resourceSlbFastHttpProxyDelete,

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
						"counters2": {
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

func resourceSlbFastHttpProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating fast-http-proxy (Inside resourceSlbFastHttpProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbFastHttpProxy(d)
		d.SetId("1")
		go_vthunder.PostSlbFastHttpProxy(client.Token, vc, client.Host)

		return resourceSlbFastHttpProxyRead(d, meta)
	}
	return nil
}

func resourceSlbFastHttpProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading fast-http-proxy (Inside resourceSlbFastHttpProxyRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbFastHttpProxy(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No fast-http-proxy found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbFastHttpProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFastHttpProxyRead(d, meta)
}

func resourceSlbFastHttpProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbFastHttpProxyRead(d, meta)
}

//Utility method to instantiate FastHttpProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbFastHttpProxy(d *schema.ResourceData) go_vthunder.FastHttpProxy {
	var vc go_vthunder.FastHttpProxy

	var c go_vthunder.FastHttpProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable18, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable18
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		se.Counters2 = d.Get(prefix + ".counters2").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
