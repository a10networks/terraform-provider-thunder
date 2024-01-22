package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateHttpMalformedHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_http_malformed_http`: Configure and enable malformed HTTP check\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateHttpMalformedHttpCreate,
		UpdateContext: resourceDdosZoneTemplateHttpMalformedHttpUpdate,
		ReadContext:   resourceDdosZoneTemplateHttpMalformedHttpRead,
		DeleteContext: resourceDdosZoneTemplateHttpMalformedHttpDelete,

		Schema: map[string]*schema.Schema{
			"malformed_http": {
				Type: schema.TypeString, Required: true, Description: "'check': Configure malformed HTTP parameters;",
			},
			"malformed_http_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'reset': Reset client connection; 'blacklist-src': Blacklist-src;",
			},
			"malformed_http_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
			},
			"malformed_http_bad_chunk_mon_enabled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enabling bad chunk monitoring. Default is disabled",
			},
			"malformed_http_max_content_length": {
				Type: schema.TypeInt, Optional: true, Default: 4294967295, Description: "Set the maxinum content-length header. Default value is 4294967295 bytes",
			},
			"malformed_http_max_header_name_size": {
				Type: schema.TypeInt, Optional: true, Default: 64, Description: "Set the maxinum header name length. Default value is 64.",
			},
			"malformed_http_max_line_size": {
				Type: schema.TypeInt, Optional: true, Default: 32512, Description: "Set the maximum line size. Default value is 32512",
			},
			"malformed_http_max_num_headers": {
				Type: schema.TypeInt, Optional: true, Default: 90, Description: "Set the maximum number of headers. Default value is 90",
			},
			"malformed_http_max_req_line_size": {
				Type: schema.TypeInt, Optional: true, Default: 32512, Description: "Set the maximum request line size. Default value is 32512",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"http_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "HttpTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateHttpMalformedHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpMalformedHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpMalformedHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateHttpMalformedHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateHttpMalformedHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpMalformedHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpMalformedHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateHttpMalformedHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateHttpMalformedHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpMalformedHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpMalformedHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateHttpMalformedHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateHttpMalformedHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateHttpMalformedHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateHttpMalformedHttp(d *schema.ResourceData) edpt.DdosZoneTemplateHttpMalformedHttp {
	var ret edpt.DdosZoneTemplateHttpMalformedHttp
	ret.Inst.MalformedHttp = d.Get("malformed_http").(string)
	ret.Inst.MalformedHttpAction = d.Get("malformed_http_action").(string)
	ret.Inst.MalformedHttpActionListName = d.Get("malformed_http_action_list_name").(string)
	ret.Inst.MalformedHttpBadChunkMonEnabled = d.Get("malformed_http_bad_chunk_mon_enabled").(int)
	ret.Inst.MalformedHttpMaxContentLength = d.Get("malformed_http_max_content_length").(int)
	ret.Inst.MalformedHttpMaxHeaderNameSize = d.Get("malformed_http_max_header_name_size").(int)
	ret.Inst.MalformedHttpMaxLineSize = d.Get("malformed_http_max_line_size").(int)
	ret.Inst.MalformedHttpMaxNumHeaders = d.Get("malformed_http_max_num_headers").(int)
	ret.Inst.MalformedHttpMaxReqLineSize = d.Get("malformed_http_max_req_line_size").(int)
	//omit uuid
	ret.Inst.HttpTmplName = d.Get("http_tmpl_name").(string)
	return ret
}
