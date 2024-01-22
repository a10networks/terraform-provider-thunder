package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_logging`: Logging template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateLoggingCreate,
		UpdateContext: resourceSlbTemplateLoggingUpdate,
		ReadContext:   resourceSlbTemplateLoggingRead,
		DeleteContext: resourceSlbTemplateLoggingDelete,

		Schema: map[string]*schema.Schema{
			"auto": {
				Type: schema.TypeString, Optional: true, Default: "auto", Description: "'auto': Configure auto NAT for logging, default is auto enabled;",
			},
			"format": {
				Type: schema.TypeString, Optional: true, Description: "Specify a format string for web logging (format string(less than 250 characters) for web logging)",
			},
			"keep_end": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Number of unmasked characters at the end (default: 0)",
			},
			"keep_start": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Number of unmasked characters at the beginning (default: 0)",
			},
			"local_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "1 to enable local logging (1 to enable local logging, default 0)",
			},
			"mask": {
				Type: schema.TypeString, Optional: true, Default: "X", Description: "Character to mask the matched pattern (default: X)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Logging Template Name",
			},
			"pcre_mask": {
				Type: schema.TypeString, Optional: true, Description: "Mask matched PCRE pattern in the log",
			},
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"pool_shared": {
				Type: schema.TypeString, Optional: true, Description: "Specify NAT pool or pool group",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a Service Group to the logging template (Service Group Name)",
			},
			"shared_partition_pool": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a NAT pool or pool group from shared partition",
			},
			"shared_partition_tcp_proxy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a TCP Proxy template from shared partition",
			},
			"tcp_proxy": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template Name",
			},
			"template_tcp_proxy_shared": {
				Type: schema.TypeString, Optional: true, Description: "TCP Proxy Template name",
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
func resourceSlbTemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateLogging(d *schema.ResourceData) edpt.SlbTemplateLogging {
	var ret edpt.SlbTemplateLogging
	ret.Inst.Auto = d.Get("auto").(string)
	ret.Inst.Format = d.Get("format").(string)
	ret.Inst.KeepEnd = d.Get("keep_end").(int)
	ret.Inst.KeepStart = d.Get("keep_start").(int)
	ret.Inst.LocalLogging = d.Get("local_logging").(int)
	ret.Inst.Mask = d.Get("mask").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PcreMask = d.Get("pcre_mask").(string)
	ret.Inst.Pool = d.Get("pool").(string)
	ret.Inst.PoolShared = d.Get("pool_shared").(string)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.SharedPartitionPool = d.Get("shared_partition_pool").(int)
	ret.Inst.SharedPartitionTcpProxyTemplate = d.Get("shared_partition_tcp_proxy_template").(int)
	ret.Inst.TcpProxy = d.Get("tcp_proxy").(string)
	ret.Inst.TemplateTcpProxyShared = d.Get("template_tcp_proxy_shared").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
