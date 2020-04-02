package vthunder

//vThunder resource SlbTemplateExternalService

import (
	"fmt"
	"log"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSlbTemplateExternalService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSlbTemplateExternalServiceCreate,
		Update: resourceSlbTemplateExternalServiceUpdate,
		Read:   resourceSlbTemplateExternalServiceRead,
		Delete: resourceSlbTemplateExternalServiceDelete,
		Schema: map[string]*schema.Schema{
			"template_tcp_proxy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_persist_source_ip_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_tcp_proxy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"service_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bypass_ip_cfg": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"mask": {
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
			"request_header_forward_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_header_forward": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"tcp_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"source_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_persist_source_ip_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"failure_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSlbTemplateExternalServiceCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateExternalService (Inside resourceSlbTemplateExternalServiceCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateExternalService(d)
		logger.Println("[INFO] received V from method data to SlbTemplateExternalService --")
		d.SetId(name)
		go_vthunder.PostSlbTemplateExternalService(client.Token, data, client.Host)

		return resourceSlbTemplateExternalServiceRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateExternalServiceRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading SlbTemplateExternalService (Inside resourceSlbTemplateExternalServiceRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetSlbTemplateExternalService(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceSlbTemplateExternalServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateExternalService   (Inside resourceSlbTemplateExternalServiceUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateExternalService(d)
		logger.Println("[INFO] received V from method data to SlbTemplateExternalService ")
		d.SetId(name)
		go_vthunder.PutSlbTemplateExternalService(client.Token, name, data, client.Host)

		return resourceSlbTemplateExternalServiceRead(d, meta)

	}
	return nil
}

func resourceSlbTemplateExternalServiceDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateExternalServiceDelete) " + name)
		err := go_vthunder.DeleteSlbTemplateExternalService(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateExternalService(d *schema.ResourceData) go_vthunder.ExternalService {
	var vc go_vthunder.ExternalService
	var c go_vthunder.ExternalServiceInstance

	c.Name = d.Get("name").(string)
	c.SharedPartitionPersistSourceIPTemplate = d.Get("shared_partition_persist_source_ip_template").(int)
	c.Type = d.Get("type").(string)
	c.SourceIP = d.Get("source_ip").(string)
	c.TemplateTCPProxyShared = d.Get("template_tcp_proxy_shared").(string)
	c.UserTag = d.Get("user_tag").(string)
	c.SharedPartitionTCPProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	c.Action = d.Get("action").(string)
	c.ServiceGroup = d.Get("service_group").(string)
	c.FailureAction = d.Get("failure_action").(string)
	c.Timeout = d.Get("timeout").(int)
	c.TCPProxy = d.Get("tcp_proxy").(string)
	c.TemplatePersistSourceIPShared = d.Get("template_persist_source_ip_shared").(string)
	RequestHeaderForwardListCount := d.Get("request_header_forward_list.#").(int)
	c.RequestHeaderForward = make([]go_vthunder.RequestHeaderForwardList, 0, RequestHeaderForwardListCount)
	for i := 0; i < RequestHeaderForwardListCount; i++ {
		var rhf go_vthunder.RequestHeaderForwardList
		prefix := fmt.Sprintf("request_header_forward_list.%d", i)
		rhf.RequestHeaderForward = d.Get(prefix + ".request_header_forward").(string)
		c.RequestHeaderForward = append(c.RequestHeaderForward, rhf)
	}
	BypassIPCfgCount := d.Get("bypass_ip_cfg.#").(int)
	c.BypassIP = make([]go_vthunder.BypassIPCfg1, 0, BypassIPCfgCount)
	for i := 0; i < BypassIPCfgCount; i++ {
		var bicc go_vthunder.BypassIPCfg1
		prefix := fmt.Sprintf("bypass_ip_cfg.%d", i)
		bicc.BypassIP = d.Get(prefix + ".bypass_ip").(string)
		bicc.Mask = d.Get(prefix + "mask").(string)
		c.BypassIP = append(c.BypassIP, bicc)
	}

	vc.UUID = c
	return vc
}
