package vthunder

// vThunder resource Slb HealthGateway

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbHealthGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbHealthGatewayCreate,
		Update: resourceSlbHealthGatewayUpdate,
		Read:   resourceSlbHealthGatewayRead,
		Delete: resourceSlbHealthGatewayDelete,

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

func resourceSlbHealthGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	logger.Println("[INFO] Creating health_gateway (Inside resourceSlbHealthGatewayCreate)")

	if client.Host != "" {
		vc := dataToSlbHealthGateway(d)
		d.SetId("1")
		go_vthunder.PostSlbHealthGateway(client.Token, vc, client.Host)

		return resourceSlbHealthGatewayRead(d, meta)
	}
	return nil
}

func resourceSlbHealthGatewayRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading health_gateway (Inside resourceSlbHealthGatewayRead)")

	client := meta.(vThunder)

	if client.Host != "" {

		name := d.Id()

		vc, err := go_vthunder.GetSlbHealthGateway(client.Token, client.Host)

		if vc == nil {
			logger.Println("[INFO] No health_gateway found" + name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceSlbHealthGatewayUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHealthGatewayRead(d, meta)
}

func resourceSlbHealthGatewayDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbHealthGatewayRead(d, meta)
}

//Utility method to instantiate HealthGateway Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToSlbHealthGateway(d *schema.ResourceData) go_vthunder.HealthGateway {
	var vc go_vthunder.HealthGateway

	var c go_vthunder.HealthGatewayInstance

	se_count := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable32, 0, se_count)

	for i := 0; i < se_count; i++ {
		var se go_vthunder.SamplingEnable32
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		se.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, se)
	}

	vc.UUID = c

	return vc
}
