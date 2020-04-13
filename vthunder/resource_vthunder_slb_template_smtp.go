package vthunder

//vThunder resource SlbTemplateSMTP

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateSMTP() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateSMTPCreate,
		Update: resourceSlbTemplateSMTPUpdate,
		Read:   resourceSlbTemplateSMTPRead,
		Delete: resourceSlbTemplateSMTPDelete,
		Schema: map[string]*schema.Schema{
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"command_disable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"service_ready_msg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_domain_switching": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_string": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"service_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"switching_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
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
			"server_starttls_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"client_starttls_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateSMTPCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateSMTP (Inside resourceSlbTemplateSMTPCreate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSMTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSMTP --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateSMTP(client.Token, data, client.Host)

		return resourceSlbTemplateSMTPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateSMTPRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateSMTP (Inside resourceSlbTemplateSMTPRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateSMTP(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateSMTPUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateSMTP   (Inside resourceSlbTemplateSMTPUpdate) ")
		name := d.Get("name").(string)
		data := dataToTemplateSMTP(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateSMTP ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateSMTP(client.Token, name, data, client.Host)

		return resourceSlbTemplateSMTPRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateSMTPDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateSMTPDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateSMTP(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToTemplateSMTP(d *schema.ResourceData) go_vthunder.SMTP {
	var vc go_vthunder.SMTP
	var c go_vthunder.SMTPInstance

	clientDomainSwitchingCount := d.Get("client_domain_switching.#").(int)
	c.ServiceGroup = make([]go_vthunder.ClientDomainSwitching, 0, clientDomainSwitchingCount)
	for i := 0; i < clientDomainSwitchingCount; i++ {
		var cds go_vthunder.ClientDomainSwitching
		prefix := fmt.Sprintf("client_domain_switching.%d", i)
		cds.MatchString = d.Get(prefix + ".match_string").(string)
		cds.ServiceGroup = d.Get(prefix + "service_group").(string)
		cds.SwitchingType = d.Get(prefix + "switching_type").(string)
		c.ServiceGroup = append(c.ServiceGroup, cds)
	}
	commandDisableCount := d.Get("command_disable.#").(int)
	c.DisableType = make([]go_vthunder.CommandDisable, 0, commandDisableCount)
	for i := 0; i < commandDisableCount; i++ {
		var cd go_vthunder.CommandDisable
		prefix := fmt.Sprintf("command_disable.%d", i)
		cd.DisableType = d.Get(prefix + ".disable_type").(string)
		c.DisableType = append(c.DisableType, cd)
	}

	var temp go_vthunder.Template1
	temp.Logging = d.Get("template.0.logging").(string)
	c.Logging = temp
	c.Name = d.Get("name").(string)
	c.ServerDomain = d.Get("server_domain").(string)
	c.ServerStarttlsType = d.Get("server_starttls_type").(string)
	c.ClientStarttlsType = d.Get("client_starttls_type").(string)
	c.ServiceReadyMsg = d.Get("service_ready_msg").(string)
	c.UserTag = d.Get("user_tag").(string)

	vc.UUID = c
	return vc
}
