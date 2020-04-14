package vthunder

//vThunder resource TemplateClientSsh

import (
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTemplateClientSsh() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateClientSshCreate,
		Update: resourceTemplateClientSshUpdate,
		Read:   resourceTemplateClientSshRead,
		Delete: resourceTemplateClientSshDelete,
		Schema: map[string]*schema.Schema{
			"forward_proxy_hostkey": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"encrypted": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
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
			"passphrase": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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

func resourceTemplateClientSshCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating TemplateClientSsh (Inside resourceTemplateClientSshCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateClientSsh(d)
		logger.Println("[INFO] received formatted data from method data to TemplateClientSsh --")
		d.SetId(name)
		go_vthunder.PostTemplateClientSsh(client.Token, data, client.Host)

		return resourceTemplateClientSshRead(d, meta)

	}
	return nil
}

func resourceTemplateClientSshRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading TemplateClientSsh (Inside resourceTemplateClientSshRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetTemplateClientSsh(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceTemplateClientSshUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying TemplateClientSsh   (Inside resourceTemplateClientSshUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateClientSsh(d)
		logger.Println("[INFO] received V from method data to TemplateClientSsh ")
		d.SetId(name)
		go_vthunder.PutTemplateClientSsh(client.Token, name, data, client.Host)

		return resourceTemplateClientSshRead(d, meta)

	}
	return nil
}

func resourceTemplateClientSshDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceTemplateClientSshDelete) " + name)
		err := go_vthunder.DeleteTemplateClientSsh(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateClientSsh(d *schema.ResourceData) go_vthunder.ClientSSH {
	var vc go_vthunder.ClientSSH
	var c go_vthunder.ClientSSHInstance

	c.Encrypted = d.Get("encrypted").(string)
	c.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	c.ForwardProxyHostkey = d.Get("forward_proxy_hostkey").(string)
	c.Name = d.Get("name").(string)
	c.Passphrase = d.Get("passphrase").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
