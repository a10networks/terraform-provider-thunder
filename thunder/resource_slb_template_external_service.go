package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateExternalService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_external_service`: External service template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateExternalServiceCreate,
		UpdateContext: resourceSlbTemplateExternalServiceUpdate,
		ReadContext:   resourceSlbTemplateExternalServiceRead,
		DeleteContext: resourceSlbTemplateExternalServiceDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "continue", Description: "'continue': Continue; 'drop': Drop; 'reset': Reset;",
			},
			"bypass_ip_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_ip": {
							Type: schema.TypeString, Optional: true, Description: "ip address to bypass external service",
						},
						"mask": {
							Type: schema.TypeString, Optional: true, Description: "IP prefix mask",
						},
					},
				},
			},
			"failure_action": {
				Type: schema.TypeString, Optional: true, Default: "continue", Description: "'continue': Continue; 'drop': Drop; 'reset': Reset;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "External Service Template Name",
			},
			"request_header_forward_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_header_forward": {
							Type: schema.TypeString, Optional: true, Description: "Request header to be forwarded to external service (Header Name)",
						},
					},
				},
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to the template (Service Group Name)",
			},
			"shared_partition_persist_source_ip_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a persist source ip template from shared partition",
			},
			"shared_partition_tcp_proxy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a TCP Proxy template from shared partition",
			},
			"source_ip": {
				Type: schema.TypeString, Optional: true, Description: "Source IP persistence template (Source IP persistence template name)",
			},
			"tcp_proxy": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template Name",
			},
			"template_persist_source_ip_shared": {
				Type: schema.TypeString, Optional: true, Description: "Source IP Persistence Template Name",
			},
			"template_tcp_proxy_shared": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template name",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Timeout value 1 - 200 in units of 200ms, default is 5 (default is 1000ms) (1 - 200 in units of 200ms, default is 5 (1000ms))",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Default: "url-filter", Description: "'skyfire-icap': Skyfire ICAP service; 'url-filter': URL filtering service;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTemplateExternalServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateExternalServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateExternalService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateExternalServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateExternalServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateExternalServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateExternalService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateExternalServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateExternalServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateExternalServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateExternalService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateExternalServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateExternalServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateExternalService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateExternalServiceBypassIpCfg(d []interface{}) []edpt.SlbTemplateExternalServiceBypassIpCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateExternalServiceBypassIpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateExternalServiceBypassIpCfg
		oi.BypassIp = in["bypass_ip"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateExternalServiceRequestHeaderForwardList(d []interface{}) []edpt.SlbTemplateExternalServiceRequestHeaderForwardList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateExternalServiceRequestHeaderForwardList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateExternalServiceRequestHeaderForwardList
		oi.RequestHeaderForward = in["request_header_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateExternalService(d *schema.ResourceData) edpt.SlbTemplateExternalService {
	var ret edpt.SlbTemplateExternalService
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.BypassIpCfg = getSliceSlbTemplateExternalServiceBypassIpCfg(d.Get("bypass_ip_cfg").([]interface{}))
	ret.Inst.FailureAction = d.Get("failure_action").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RequestHeaderForwardList = getSliceSlbTemplateExternalServiceRequestHeaderForwardList(d.Get("request_header_forward_list").([]interface{}))
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.SharedPartitionPersistSourceIpTemplate = d.Get("shared_partition_persist_source_ip_template").(int)
	ret.Inst.SharedPartitionTcpProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	ret.Inst.SourceIp = d.Get("source_ip").(string)
	ret.Inst.TcpProxy = d.Get("tcp_proxy").(string)
	ret.Inst.TemplatePersistSourceIpShared = d.Get("template_persist_source_ip_shared").(string)
	ret.Inst.TemplateTcpProxyShared = d.Get("template_tcp_proxy_shared").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
