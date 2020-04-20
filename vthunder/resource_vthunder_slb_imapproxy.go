package vthunder

//vThunder resource SlbImapproxy

import (
	"fmt"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbImapproxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbImapproxyCreate,
		Update: resourceSlbImapproxyUpdate,
		Read:   resourceSlbImapproxyRead,
		Delete: resourceSlbImapproxyDelete,
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

func resourceSlbImapproxyCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbImapproxy (Inside resourceSlbImapproxyCreate) ")

		data := dataToSlbImapproxy(d)
		logger.Println("[INFO] received V from method data to SlbImapproxy --")
		d.SetId("1")
		go_vthunder.PostSlbImapproxy(client.Token, data, client.Host)

		return resourceSlbImapproxyRead(d, meta)

	}
	return nil
}

func resourceSlbImapproxyRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbImapproxy (Inside resourceSlbImapproxyRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetSlbImapproxy(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbImapproxyUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbImapproxyRead(d, meta)
}

func resourceSlbImapproxyDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbImapproxyRead(d, meta)
}

func dataToSlbImapproxy(d *schema.ResourceData) go_vthunder.Imapproxy {
	var vc go_vthunder.Imapproxy
	var c go_vthunder.ImapproxyInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable41, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable41
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}
	vc.Counters1 = c
	return vc
}
