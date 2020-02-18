package vthunder

//vThunder resource TemplateFTP

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateFTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateFTPCreate,
		Update: resourceTemplateFTPUpdate,
		Read:   resourceTemplateFTPRead,
		Delete: resourceTemplateFTPDelete,
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
			"active_mode_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"to": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"any": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"active_mode_port_val": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateFTPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateFTP (Inside resourceTemplateFTPCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateFTP(d)
		logger.Println("[INFO] received V from method data to TemplateFTP --")
		d.SetId(name)
		go_vthunder.PostTemplateFTP(client.Token, data, client.Host)

		return resourceTemplateFTPRead(d, meta)

	}
	return nil
}

func resourceTemplateFTPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateFTP (Inside resourceTemplateFTPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateFTP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateFTPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateFTP   (Inside resourceTemplateFTPUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateFTP(d)
		logger.Println("[INFO] received V from method data to TemplateFTP ")
		d.SetId(name)
		go_vthunder.PutTemplateFTP(client.Token, name, data, client.Host)

		return resourceTemplateFTPRead(d, meta)

	}
	return nil
}

func resourceTemplateFTPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateFTPDelete) " + name)
		err := go_vthunder.DeleteTemplateFTP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateFTP(d *schema.ResourceData) go_vthunder.FTP {
	var vc go_vthunder.FTP
	var c go_vthunder.FtpInstance

	c.Name = d.Get("name").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.ActiveModePort = d.Get("active_mode_port").(int)
	c.ActiveModePortVal = d.Get("active_mode_port_val").(int)
	c.Any = d.Get("any").(int)
	c.To = d.Get("to").(int)

	vc.UUID = c
	return vc
}
