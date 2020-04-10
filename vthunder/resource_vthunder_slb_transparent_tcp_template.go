package vthunder

//vThunder resource SlbTransperentTcpTemplate

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTransperentTcpTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTransperentTcpTemplateCreate,
		Update: resourceSlbTransperentTcpTemplateUpdate,
		Read:   resourceSlbTransperentTcpTemplateRead,
		Delete: resourceSlbTransperentTcpTemplateDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTransperentTcpTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTransperentTcpTemplate (Inside resourceSlbTransperentTcpTemplateCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTransperentTcpTemplate(d)
		logger.Println("[INFO] received V from method data to SlbTransperentTcpTemplate --")
		d.SetId(name)
		go_vthunder.PostSlbTransperentTcpTemplate(client.Token, data, client.Host)

		return resourceSlbTransperentTcpTemplateRead(d, meta)

	}
	return nil
}

func resourceSlbTransperentTcpTemplateRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTransperentTcpTemplate (Inside resourceSlbTransperentTcpTemplateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTransperentTcpTemplate(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTransperentTcpTemplateUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbTransperentTcpTemplateRead(d, meta)
}

func resourceSlbTransperentTcpTemplateDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbTransperentTcpTemplateRead(d, meta)
}

func dataToSlbTransperentTcpTemplate(d *schema.ResourceData) go_vthunder.TransperentTcpTemplate {
	var vc go_vthunder.TransperentTcpTemplate
	var c go_vthunder.TransperentTcpTemplateInstance
	c.Name = d.Get("name").(string)

	vc.Name = c
	return vc
}
