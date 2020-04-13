package vthunder

//vThunder resource SlbTemplateServerSSH

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateServerSSH() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateServerSSHCreate,
		Update: resourceSlbTemplateServerSSHUpdate,
		Read:   resourceSlbTemplateServerSSHRead,
		Delete: resourceSlbTemplateServerSSHDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
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
			"forward_proxy_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateServerSSHCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateServerSSH (Inside resourceSlbTemplateServerSSHCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServerSSH(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSSH --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateServerSSH(client.Token, data, client.Host)

		return resourceSlbTemplateServerSSHRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateServerSSHRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateServerSSH (Inside resourceSlbTemplateServerSSHRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateServerSSH(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateServerSSHUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateServerSSH   (Inside resourceSlbTemplateServerSSHUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateServerSSH(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateServerSSH ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateServerSSH(client.Token, name, data, client.Host)

		return resourceSlbTemplateServerSSHRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateServerSSHDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateServerSSHDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateServerSSH(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateServerSSH(d *schema.ResourceData) go_vthunder.ServerSSH {
	var vc go_vthunder.ServerSSH
	var c go_vthunder.ServerSSHInstance

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.SamplingEnable2, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj1 go_vthunder.SamplingEnable2
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj1.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj1)
	}

	c.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
