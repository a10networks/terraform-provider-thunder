package vthunder

// vThunder resource Slb ConnectionReuse

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbConnectionReuse() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbConnectionReuseCreate,
		Update: resourceSlbConnectionReuseUpdate,
		Read:   resourceSlbConnectionReuseRead,
		Delete: resourceSlbConnectionReuseDelete,

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

func resourceSlbConnectionReuseCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating connection-reuse (Inside resourceSlbConnectionReuseCreate)")

	if client.Host != "" {
		vc := dataToSlbConnectionReuse(d)
		d.SetId("1")
		go_vthunder.PostSlbConnectionReuse(client.Token, vc, client.Host)

		return resourceSlbConnectionReuseRead(d, meta)
	}
	return nil
}

func resourceSlbConnectionReuseRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading connection-reuse (Inside resourceSlbConnectionReuseRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbConnectionReuse(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No connection-reuse found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbConnectionReuseUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbConnectionReuseRead(d, meta)
}

func resourceSlbConnectionReuseDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbConnectionReuseRead(d, meta)
}

//Utility method to instantiate ConnectionReuse Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbConnectionReuse(d *schema.ResourceData) go_vthunder.ConnectionReuse {
	var vc go_vthunder.ConnectionReuse

	var c go_vthunder.SlbConnectionReuseInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable7, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable7
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
