package vthunder

// vThunder resource Slb HealthStat

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbHealthStat() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbHealthStatCreate,
		Update: resourceSlbHealthStatUpdate,
		Read:   resourceSlbHealthStatRead,
		Delete: resourceSlbHealthStatDelete,

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

func resourceSlbHealthStatCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating health-stat (Inside resourceSlbHealthStatCreate)")

	if client.Host != "" {
		vc := dataToSlbHealthStat(d)
		d.SetId("1")
		go_vthunder.PostSlbHealthStat(client.Token, vc, client.Host)

		return resourceSlbHealthStatRead(d, meta)
	}
	return nil
}

func resourceSlbHealthStatRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading health-stat (Inside resourceSlbHealthStatRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbHealthStat(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No health-stat found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbHealthStatUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHealthStatRead(d, meta)
}

func resourceSlbHealthStatDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHealthStatRead(d, meta)
}

//Utility method to instantiate HealthStat Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHealthStat(d *schema.ResourceData) go_vthunder.HealthStat {
	var vc go_vthunder.HealthStat

	var c go_vthunder.HealthStatInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable33, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable33
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
