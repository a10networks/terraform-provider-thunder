package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateSipFilterHeader() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_sip_filter_header`: SIP Header Regex Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateSipFilterHeaderCreate,
		UpdateContext: resourceDdosTemplateSipFilterHeaderUpdate,
		ReadContext:   resourceDdosTemplateSipFilterHeaderRead,
		DeleteContext: resourceDdosTemplateSipFilterHeaderDelete,

		Schema: map[string]*schema.Schema{
			"sip_filter_header_blacklist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Also blacklist the source when action is taken",
			},
			"sip_filter_header_count_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take no action and continue processing the next filter",
			},
			"sip_filter_header_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"sip_filter_header_seq": {
				Type: schema.TypeInt, Required: true, Description: "Sequence number",
			},
			"sip_filter_header_unmatched": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
			},
			"sip_filter_header_whitelist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Whitelist the source after filter passes, packets are dropped until then",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosTemplateSipFilterHeaderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipFilterHeaderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipFilterHeader(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSipFilterHeaderRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateSipFilterHeaderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipFilterHeaderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipFilterHeader(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateSipFilterHeaderRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateSipFilterHeaderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipFilterHeaderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipFilterHeader(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateSipFilterHeaderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateSipFilterHeaderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateSipFilterHeader(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateSipFilterHeader(d *schema.ResourceData) edpt.DdosTemplateSipFilterHeader {
	var ret edpt.DdosTemplateSipFilterHeader
	ret.Inst.SipFilterHeaderBlacklist = d.Get("sip_filter_header_blacklist").(int)
	ret.Inst.SipFilterHeaderCountOnly = d.Get("sip_filter_header_count_only").(int)
	ret.Inst.SipFilterHeaderRegex = d.Get("sip_filter_header_regex").(string)
	ret.Inst.SipFilterHeaderSeq = d.Get("sip_filter_header_seq").(int)
	ret.Inst.SipFilterHeaderUnmatched = d.Get("sip_filter_header_unmatched").(int)
	ret.Inst.SipFilterHeaderWhitelist = d.Get("sip_filter_header_whitelist").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SipTmplName = d.Get("sip_tmpl_name").(string)
	return ret
}
