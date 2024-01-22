package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateRespmodIcap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_respmod_icap`: RESPMOD ICAP template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateRespmodIcapCreate,
		UpdateContext: resourceSlbTemplateRespmodIcapUpdate,
		ReadContext:   resourceSlbTemplateRespmodIcapRead,
		DeleteContext: resourceSlbTemplateRespmodIcapDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "continue", Description: "'continue': Continue; 'drop': Drop; 'reset': Reset;",
			},
			"bypass_ip_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_ip": {
							Type: schema.TypeString, Optional: true, Description: "ip address to bypass respmod-icap service",
						},
						"mask": {
							Type: schema.TypeString, Optional: true, Description: "IP prefix mask",
						},
					},
				},
			},
			"disable_http_server_reset": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Don't reset http server",
			},
			"fail_close": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "When template sg is down mark vport down",
			},
			"failure_action": {
				Type: schema.TypeString, Optional: true, Default: "continue", Description: "'continue': Continue; 'drop': Drop; 'reset': Reset;",
			},
			"include_protocol_in_uri": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include protocol and port in HTTP URI",
			},
			"log_only_allowed_method": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only log allowed HTTP method",
			},
			"logging": {
				Type: schema.TypeString, Optional: true, Description: "logging template (Logging template name)",
			},
			"min_payload_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "min-payload-size value 0 - 65535, default is 0",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Reqmod ICAP Template Name",
			},
			"preview": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Preview value 1 - 32768, default is 32768",
			},
			"server_ssl": {
				Type: schema.TypeString, Optional: true, Description: "Server SSL template (Server SSL template name)",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to the template (Service Group Name)",
			},
			"service_url": {
				Type: schema.TypeString, Optional: true, Description: "URL to send to ICAP server (Service URL Name)",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"x_auth_url": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use URL format for authentication",
			},
		},
	}
}
func resourceSlbTemplateRespmodIcapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateRespmodIcapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateRespmodIcap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateRespmodIcapRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateRespmodIcapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateRespmodIcapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateRespmodIcap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateRespmodIcapRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateRespmodIcapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateRespmodIcapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateRespmodIcap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateRespmodIcapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateRespmodIcapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateRespmodIcap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateRespmodIcapBypassIpCfg(d []interface{}) []edpt.SlbTemplateRespmodIcapBypassIpCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateRespmodIcapBypassIpCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateRespmodIcapBypassIpCfg
		oi.BypassIp = in["bypass_ip"].(string)
		oi.Mask = in["mask"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateRespmodIcap(d *schema.ResourceData) edpt.SlbTemplateRespmodIcap {
	var ret edpt.SlbTemplateRespmodIcap
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.BypassIpCfg = getSliceSlbTemplateRespmodIcapBypassIpCfg(d.Get("bypass_ip_cfg").([]interface{}))
	ret.Inst.DisableHttpServerReset = d.Get("disable_http_server_reset").(int)
	ret.Inst.FailClose = d.Get("fail_close").(int)
	ret.Inst.FailureAction = d.Get("failure_action").(string)
	ret.Inst.IncludeProtocolInUri = d.Get("include_protocol_in_uri").(int)
	ret.Inst.LogOnlyAllowedMethod = d.Get("log_only_allowed_method").(int)
	ret.Inst.Logging = d.Get("logging").(string)
	ret.Inst.MinPayloadSize = d.Get("min_payload_size").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Preview = d.Get("preview").(int)
	ret.Inst.ServerSsl = d.Get("server_ssl").(string)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.ServiceUrl = d.Get("service_url").(string)
	ret.Inst.SharedPartitionPersistSourceIpTemplate = d.Get("shared_partition_persist_source_ip_template").(int)
	ret.Inst.SharedPartitionTcpProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	ret.Inst.SourceIp = d.Get("source_ip").(string)
	ret.Inst.TcpProxy = d.Get("tcp_proxy").(string)
	ret.Inst.TemplatePersistSourceIpShared = d.Get("template_persist_source_ip_shared").(string)
	ret.Inst.TemplateTcpProxyShared = d.Get("template_tcp_proxy_shared").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.XAuthUrl = d.Get("x_auth_url").(int)
	return ret
}
