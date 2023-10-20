package thunder

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"util"
)

func resourceSlbTemplateDnsLogging() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSlbTemplateDnsLoggingCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingRead,
		DeleteContext: resourceSlbTemplateDnsLoggingDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "DNS Logging Template Name",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS Logging template",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"dns_logging_type": {
				Type: schema.TypeString, Optional: true, Description: "'query': DNS Query Logging; ",
				ValidateFunc: validation.StringInSlice([]string{"query"}, false),
			},
			"dns_logging_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'both': Log DNS over tcp and udp; 'tcp': Log DNS over tcp; 'udp': Log DNS over udp; ",
				ValidateFunc: validation.StringInSlice([]string{"both", "tcp", "udp"}, false),
			},
			"dns_logging_request_section": {
				Type: schema.TypeString, Optional: true, Description: "'all': Log DNS header and question section; 'header': Log DNS header information; 'question': Log DNS question section; ",
				ValidateFunc: validation.StringInSlice([]string{"all", "header", "question"}, false),
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceSlbTemplateDnsLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceSlbTemplateDnsLoggingCreate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLogging(d)
		d.SetId(obj.GetId())
		err := PostEndpointSlbTemplateDnsLogging(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceSlbTemplateDnsLoggingRead()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		_, err := GetEndpointSlbTemplateDnsLogging(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceSlbTemplateDnsLoggingUpdate()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLogging(d)
		err := PutEndpointSlbTemplateDnsLogging(client.Token, obj, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("resourceSlbTemplateDnsLoggingDelete()")
	client := meta.(Thunder)
	var diags diag.Diagnostics
	if client.Host != "" {
		err := DeleteEndpointSlbTemplateDnsLogging(client.Token, client.Host, d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsLogging(d *schema.ResourceData) EndpointSlbTemplateDnsLogging {
	var ret EndpointSlbTemplateDnsLogging
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DnsLoggingType = d.Get("dns_logging_type").(string)
	ret.Inst.DnsLoggingProtocol = d.Get("dns_logging_protocol").(string)
	ret.Inst.DnsLoggingRequestSection = d.Get("dns_logging_request_section").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
