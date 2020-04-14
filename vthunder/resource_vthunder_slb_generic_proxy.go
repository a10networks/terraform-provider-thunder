package vthunder

// vThunder resource Slb GenericProxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbGenericProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbGenericProxyCreate,
		Update: resourceSlbGenericProxyUpdate,
		Read:   resourceSlbGenericProxyRead,
		Delete: resourceSlbGenericProxyDelete,

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

func resourceSlbGenericProxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating generic-proxy (Inside resourceSlbGenericProxyCreate)")

	if client.Host != "" {
		vc := dataToSlbGenericProxy(d)
		d.SetId("1")
		go_vthunder.PostSlbGenericProxy(client.Token, vc, client.Host)

		return resourceSlbGenericProxyRead(d, meta)
	}
	return nil
}

func resourceSlbGenericProxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading generic-proxy (Inside resourceSlbGenericProxyRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbGenericProxy(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No generic-proxy found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbGenericProxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbGenericProxyRead(d, meta)
}

func resourceSlbGenericProxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbGenericProxyRead(d, meta)
}

//Utility method to instantiate GenericProxy Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbGenericProxy(d *schema.ResourceData) go_vthunder.GenericProxy {
	var vc go_vthunder.GenericProxy

	var c go_vthunder.GenericProxyInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable23, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable23
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
