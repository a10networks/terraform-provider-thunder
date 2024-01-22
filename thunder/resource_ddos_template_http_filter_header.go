package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateHttpFilterHeader() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_http_filter_header`: HTTP Header Regex Filter Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateHttpFilterHeaderCreate,
		UpdateContext: resourceDdosTemplateHttpFilterHeaderUpdate,
		ReadContext:   resourceDdosTemplateHttpFilterHeaderRead,
		DeleteContext: resourceDdosTemplateHttpFilterHeaderDelete,

		Schema: map[string]*schema.Schema{
			"http_filter_header_blacklist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Also blacklist the source when action is taken",
			},
			"http_filter_header_count_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take no action and continue processing the next filter",
			},
			"http_filter_header_regex": {
				Type: schema.TypeString, Optional: true, Description: "Regex Expression",
			},
			"http_filter_header_seq": {
				Type: schema.TypeInt, Required: true, Description: "Sequence number",
			},
			"http_filter_header_unmatched": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
			},
			"http_filter_header_whitelist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Whitelist the source after filter passes, packets are dropped until then",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosTemplateHttpFilterHeaderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpFilterHeaderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttpFilterHeader(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateHttpFilterHeaderRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateHttpFilterHeaderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpFilterHeaderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttpFilterHeader(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateHttpFilterHeaderRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateHttpFilterHeaderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpFilterHeaderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttpFilterHeader(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateHttpFilterHeaderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateHttpFilterHeaderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateHttpFilterHeader(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateHttpFilterHeader(d *schema.ResourceData) edpt.DdosTemplateHttpFilterHeader {
	var ret edpt.DdosTemplateHttpFilterHeader
	ret.Inst.HttpFilterHeaderBlacklist = d.Get("http_filter_header_blacklist").(int)
	ret.Inst.HttpFilterHeaderCountOnly = d.Get("http_filter_header_count_only").(int)
	ret.Inst.HttpFilterHeaderRegex = d.Get("http_filter_header_regex").(string)
	ret.Inst.HttpFilterHeaderSeq = d.Get("http_filter_header_seq").(int)
	ret.Inst.HttpFilterHeaderUnmatched = d.Get("http_filter_header_unmatched").(int)
	ret.Inst.HttpFilterHeaderWhitelist = d.Get("http_filter_header_whitelist").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.HttpTmplName = d.Get("http_tmpl_name").(string)
	return ret
}
