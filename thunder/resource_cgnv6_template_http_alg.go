package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateHttpAlg() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_http_alg`: HTTP-ALG Template\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateHttpAlgCreate,
		UpdateContext: resourceCgnv6TemplateHttpAlgUpdate,
		ReadContext:   resourceCgnv6TemplateHttpAlgRead,
		DeleteContext: resourceCgnv6TemplateHttpAlgDelete,

		Schema: map[string]*schema.Schema{
			"header_name_client_ip": {
				Type: schema.TypeString, Optional: true, Default: "X-Forwarded-For", Description: "Header name (default: X-Forwarded-For)",
			},
			"header_name_msisdn": {
				Type: schema.TypeString, Optional: true, Default: "X-MSISDN", Description: "Header name (default: X-MSISDN)",
			},
			"include_tunnel_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the tunnel IP (applies to DS-Lite and 6RD-NAT64 sessions)",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Default: "append", Description: "'append': Append if there is already a header (default); 'replace': Replace if there is already a header;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "HTTP-ALG template name",
			},
			"request_insert_client_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert Client IP into HTTP request",
			},
			"request_insert_msisdn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert MSISDN into HTTP request",
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
func resourceCgnv6TemplateHttpAlgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateHttpAlgCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateHttpAlg(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateHttpAlgRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateHttpAlgUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateHttpAlgUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateHttpAlg(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateHttpAlgRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateHttpAlgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateHttpAlgDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateHttpAlg(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateHttpAlgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateHttpAlgRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateHttpAlg(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6TemplateHttpAlg(d *schema.ResourceData) edpt.Cgnv6TemplateHttpAlg {
	var ret edpt.Cgnv6TemplateHttpAlg
	ret.Inst.HeaderNameClientIp = d.Get("header_name_client_ip").(string)
	ret.Inst.HeaderNameMsisdn = d.Get("header_name_msisdn").(string)
	ret.Inst.IncludeTunnelIp = d.Get("include_tunnel_ip").(int)
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RequestInsertClientIp = d.Get("request_insert_client_ip").(int)
	ret.Inst.RequestInsertMsisdn = d.Get("request_insert_msisdn").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
