package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateSipMalformedSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_sip_malformed_sip`: Configure and enable malformed SIP check\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateSipMalformedSipCreate,
		UpdateContext: resourceDdosTemplateSipMalformedSipUpdate,
		ReadContext:   resourceDdosTemplateSipMalformedSipRead,
		DeleteContext: resourceDdosTemplateSipMalformedSipDelete,

		Schema: map[string]*schema.Schema{
			"malformed_sip_call_id_max_length": {
				Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum call-id length. Default value is 32511",
			},
			"malformed_sip_check": {
				Type: schema.TypeString, Optional: true, Description: "'enable-check': Enable malformed SIP parameters;",
			},
			"malformed_sip_max_header_name_length": {
				Type: schema.TypeInt, Optional: true, Default: 63, Description: "Set the maximum header name length. Default value is 63",
			},
			"malformed_sip_max_header_value_length": {
				Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum header value length. Default value is 32511",
			},
			"malformed_sip_max_line_size": {
				Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum line size. Default value is 32511",
			},
			"malformed_sip_max_uri_length": {
				Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maximum uri size. Default value is 32511",
			},
			"malformed_sip_sdp_max_length": {
				Type: schema.TypeInt, Optional: true, Default: 32511, Description: "Set the maxinum SDP content length. Default value is 32511",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"sip_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "SipTmplName",
			},
		},
	}
}
func resourceDdosTemplateSipMalformedSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipMalformedSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipMalformedSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSipMalformedSipRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateSipMalformedSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipMalformedSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipMalformedSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSipMalformedSipRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateSipMalformedSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipMalformedSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipMalformedSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateSipMalformedSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipMalformedSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipMalformedSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateSipMalformedSip(d *schema.ResourceData) edpt.DdosTemplateSipMalformedSip {
	var ret edpt.DdosTemplateSipMalformedSip
	ret.Inst.MalformedSipCallIdMaxLength = d.Get("malformed_sip_call_id_max_length").(int)
	ret.Inst.MalformedSipCheck = d.Get("malformed_sip_check").(string)
	ret.Inst.MalformedSipMaxHeaderNameLength = d.Get("malformed_sip_max_header_name_length").(int)
	ret.Inst.MalformedSipMaxHeaderValueLength = d.Get("malformed_sip_max_header_value_length").(int)
	ret.Inst.MalformedSipMaxLineSize = d.Get("malformed_sip_max_line_size").(int)
	ret.Inst.MalformedSipMaxUriLength = d.Get("malformed_sip_max_uri_length").(int)
	ret.Inst.MalformedSipSdpMaxLength = d.Get("malformed_sip_sdp_max_length").(int)
	//omit uuid
	ret.Inst.SipTmplName = d.Get("sip_tmpl_name").(string)
	return ret
}
