package thunder

//Thunder resource SlbTransperentTcpTemplate

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
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
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTransperentTcpTemplate (Inside resourceSlbTransperentTcpTemplateCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTransperentTcpTemplate(d)
		logger.Println("[INFO] received formatted data from method data to SlbTransperentTcpTemplate --")
		d.SetId(name)
		go_thunder.PostSlbTransperentTcpTemplate(client.Token, data, client.Host)

		return resourceSlbTransperentTcpTemplateRead(d, meta)

	}
	return nil
}

func resourceSlbTransperentTcpTemplateRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTransperentTcpTemplate (Inside resourceSlbTransperentTcpTemplateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTransperentTcpTemplate(client.Token, name, client.Host)
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

func dataToSlbTransperentTcpTemplate(d *schema.ResourceData) go_thunder.TransperentTcpTemplate {
	var vc go_thunder.TransperentTcpTemplate
	var c go_thunder.TransperentTcpTemplateInstance
	c.Name = d.Get("name").(string)

	vc.Name = c
	return vc
}
