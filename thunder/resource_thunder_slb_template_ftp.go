package thunder

//Thunder resource SlbTemplateFTP

import (
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateFTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateFTPCreate,
		Update: resourceSlbTemplateFTPUpdate,
		Read:   resourceSlbTemplateFTPRead,
		Delete: resourceSlbTemplateFTPDelete,
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

func resourceSlbTemplateFTPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateFTP (Inside resourceSlbTemplateFTPCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateFTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateFTP --")
		d.SetId(name)
		go_thunder.PostSlbTemplateFTP(client.Token, data, client.Host)

		return resourceSlbTemplateFTPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateFTPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading SlbTemplateFTP (Inside resourceSlbTemplateFTPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateFTP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateFTPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateFTP   (Inside resourceSlbTemplateFTPUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateFTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateFTP ")
		d.SetId(name)
		go_thunder.PutSlbTemplateFTP(client.Token, name, data, client.Host)

		return resourceSlbTemplateFTPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateFTPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateFTPDelete) " + name)
		err := go_thunder.DeleteSlbTemplateFTP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateFTP(d *schema.ResourceData) go_thunder.FTP {
	var vc go_thunder.FTP
	var c go_thunder.FTPInstance
	c.UserTag = d.Get("user_tag").(string)
	c.To = d.Get("to").(int)
	c.ActiveModePort = d.Get("active_mode_port").(int)
	c.ActiveModePortVal = d.Get("active_mode_port_val").(int)
	c.Any = d.Get("any").(int)
	c.Name = d.Get("name").(string)

	vc.UUID = c
	return vc
}
