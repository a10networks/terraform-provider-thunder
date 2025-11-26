package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dns_logging`: DNS Logging template\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDnsLoggingCreate,
		UpdateContext: resourceDdosDnsLoggingUpdate,
		ReadContext:   resourceDdosDnsLoggingRead,
		DeleteContext: resourceDdosDnsLoggingDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS Logging template",
			},
			"dns_logging_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'both': Log DNS over tcp and udp; 'tcp': Log DNS over tcp; 'udp': Log DNS over udp;",
			},
			"dns_logging_request_section": {
				Type: schema.TypeString, Optional: true, Description: "'all': Log DNS header and question section; 'header': Log DNS header information; 'question': Log DNS question section;",
			},
			"dns_logging_type": {
				Type: schema.TypeString, Optional: true, Description: "'query': DNS Query Logging;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Logging Template Name",
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
func resourceDdosDnsLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDnsLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDnsLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDnsLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDnsLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDnsLogging(d *schema.ResourceData) edpt.DdosDnsLogging {
	var ret edpt.DdosDnsLogging
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DnsLoggingProtocol = d.Get("dns_logging_protocol").(string)
	ret.Inst.DnsLoggingRequestSection = d.Get("dns_logging_request_section").(string)
	ret.Inst.DnsLoggingType = d.Get("dns_logging_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
