package vthunder

//vThunder resource TemplateConnectionReuse

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateConnectionReuse() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateConnectionReuseCreate,
		Update: resourceTemplateConnectionReuseUpdate,
		Read:   resourceTemplateConnectionReuseRead,
		Delete: resourceTemplateConnectionReuseDelete,
		Schema: map[string]*schema.Schema{
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"keep_alive_conn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"num_conn_per_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"preopen": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"limit_per_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceTemplateConnectionReuseCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateConnectionReuse (Inside resourceTemplateConnectionReuseCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateConnectionReuse(d)
		logger.Println("[INFO] received V from method data to TemplateConnectionReuse --")
		d.SetId(name)
		go_vthunder.PostTemplateConnectionReuse(client.Token, data, client.Host)

		return resourceTemplateConnectionReuseRead(d, meta)

	}
	return nil
}

func resourceTemplateConnectionReuseRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateConnectionReuse (Inside resourceTemplateConnectionReuseRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateConnectionReuse(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateConnectionReuseUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateConnectionReuse   (Inside resourceTemplateConnectionReuseUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateConnectionReuse(d)
		logger.Println("[INFO] received V from method data to TemplateConnectionReuse ")
		d.SetId(name)
		go_vthunder.PutTemplateConnectionReuse(client.Token, name, data, client.Host)

		return resourceTemplateConnectionReuseRead(d, meta)

	}
	return nil
}

func resourceTemplateConnectionReuseDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateConnectionReuseDelete) " + name)
		err := go_vthunder.DeleteTemplateConnectionReuse(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateConnectionReuse(d *schema.ResourceData) go_vthunder.Connection_Reuse {
	var vc go_vthunder.Connection_Reuse
	var c go_vthunder.ConnectionReuseInstance
	c.KeepAliveConn = d.Get("keep_alive_conn").(int)
	c.LimitPerServer = d.Get("limit_per_server").(int)
	c.Name = d.Get("name").(string)
	c.NumConnPerPort = d.Get("num_conn_per_port").(int)
	c.Preopen = d.Get("preopen").(int)
	c.Timeout = d.Get("timeout").(int)
	c.UserTag = d.Get("user_tag").(string)
	vc.Name = c
	return vc
}
