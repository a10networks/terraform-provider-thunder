package thunder

//Thunder resource SlbTemplateExternalService

import (
	"context"
	"fmt"
	"log"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateExternalService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateExternalServiceCreate,
		UpdateContext: resourceSlbTemplateExternalServiceUpdate,
		ReadContext:   resourceSlbTemplateExternalServiceRead,
		DeleteContext: resourceSlbTemplateExternalServiceDelete,
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

func resourceSlbTemplateExternalServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SlbTemplateExternalService (Inside resourceSlbTemplateExternalServiceCreate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateExternalService(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateExternalService --")
		d.SetId(name)
		err := go_thunder.PostSlbTemplateExternalService(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateExternalServiceRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateExternalServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SlbTemplateExternalService (Inside resourceSlbTemplateExternalServiceRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetSlbTemplateExternalService(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSlbTemplateExternalServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying SlbTemplateExternalService   (Inside resourceSlbTemplateExternalServiceUpdate) ")
		name := d.Get("name").(string)
		data := dataToSlbTemplateExternalService(d)
		logger.Println("[INFO] received formatted data from method data to SlbTemplateExternalService ")
		d.SetId(name)
		err := go_thunder.PutSlbTemplateExternalService(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSlbTemplateExternalServiceRead(ctx, d, meta)

	}
	return diags
}

func resourceSlbTemplateExternalServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceSlbTemplateExternalServiceDelete) " + name)
		err := go_thunder.DeleteSlbTemplateExternalService(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToSlbTemplateExternalService(d *schema.ResourceData) go_thunder.ExternalService {
	var vc go_thunder.ExternalService
	var c go_thunder.ExternalServiceInstance

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
	c.RequestHeaderForward = make([]go_thunder.RequestHeaderForwardList, 0, RequestHeaderForwardListCount)
	for i := 0; i < RequestHeaderForwardListCount; i++ {
		var rhf go_thunder.RequestHeaderForwardList
		prefix := fmt.Sprintf("request_header_forward_list.%d", i)
		rhf.RequestHeaderForward = d.Get(prefix + ".request_header_forward").(string)
		c.RequestHeaderForward = append(c.RequestHeaderForward, rhf)
	}
	BypassIPCfgCount := d.Get("bypass_ip_cfg.#").(int)
	c.BypassIP = make([]go_thunder.BypassIPCfg1, 0, BypassIPCfgCount)
	for i := 0; i < BypassIPCfgCount; i++ {
		var bicc go_thunder.BypassIPCfg1
		prefix := fmt.Sprintf("bypass_ip_cfg.%d", i)
		bicc.BypassIP = d.Get(prefix + ".bypass_ip").(string)
		bicc.Mask = d.Get(prefix + "mask").(string)
		c.BypassIP = append(c.BypassIP, bicc)
	}

	vc.UUID = c
	return vc
}
