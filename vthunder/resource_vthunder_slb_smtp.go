package vthunder

//vThunder resource SlbSmtp

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbSmtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbSmtpCreate,
		Update: resourceSlbSmtpUpdate,
		Read:   resourceSlbSmtpRead,
		Delete: resourceSlbSmtpDelete,
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

func resourceSlbSmtpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbSmtp (Inside resourceSlbSmtpCreate) ")

		data := dataToSlbSmtp(d)
		logger.Println("[INFO] received formatted data from method data to SlbSmtp --")
		d.SetId("1")
		go_vthunder.PostSlbSmtp(client.Token, data, client.Host)

		return resourceSlbSmtpRead(d, meta)

	}
	return nil
}

func resourceSlbSmtpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbSmtp (Inside resourceSlbSmtpRead)")

	if client.Host != "" {

		data, err := go_vthunder.GetSlbSmtp(client.Token, client.Host)
		if data == nil {

			return nil
		}
		return err
	}
	return nil
}

func resourceSlbSmtpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSmtpRead(d, meta)
}

func resourceSlbSmtpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbSmtpRead(d, meta)
}
func dataToSlbSmtp(d *schema.ResourceData) go_vthunder.SlbSmtp {
	var vc go_vthunder.SlbSmtp
	var c go_vthunder.SlbSmtpInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable12, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable12
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		obj1.Counters1 = d.Get(prefix + ".counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	vc.Counters1 = c
	return vc
}
