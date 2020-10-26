package thunder

//Thunder resource SlbTransparentAclTemplate

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTransparentAclTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTransparentAclTemplateCreate,
		Update: resourceSlbTransparentAclTemplateUpdate,
		Read:   resourceSlbTransparentAclTemplateRead,
		Delete: resourceSlbTransparentAclTemplateDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTransparentAclTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTransparentAclTemplate (Inside resourceSlbTransparentAclTemplateCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTransparentAclTemplate(d)
		logger.Println("[INFO] received formatted data from method data to SlbTransparentAclTemplate --")
		d.SetId(name)
		go_thunder.PostSlbTransparentAclTemplate(client.Token, data, client.Host)

		return resourceSlbTransparentAclTemplateRead(d, meta)

	}
	return nil
}

func resourceSlbTransparentAclTemplateRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTransparentAclTemplate (Inside resourceSlbTransparentAclTemplateRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTransparentAclTemplate(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTransparentAclTemplateUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbTransparentAclTemplateRead(d, meta)
}

func resourceSlbTransparentAclTemplateDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSlbTransparentAclTemplateRead(d, meta)
}

func dataToSlbTransparentAclTemplate(d *schema.ResourceData) go_thunder.TransparentAclTemplate {
	var vc go_thunder.TransparentAclTemplate
	var c go_thunder.TransparentAclTemplateInstance
	c.Name = d.Get("name").(string)

	vc.Name = c
	return vc
}
