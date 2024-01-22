package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_http_http_protocol_check_header_filter_rule`: Header filter rules\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleCreate,
		UpdateContext: resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleUpdate,
		ReadContext:   resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleRead,
		DeleteContext: resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleDelete,

		Schema: map[string]*schema.Schema{
			"action_value": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request;",
			},
			"header_name_value": {
				Type: schema.TypeString, Optional: true, Description: "Header name value",
			},
			"header_value_value": {
				Type: schema.TypeString, Optional: true, Description: "Header value",
			},
			"match_type_value": {
				Type: schema.TypeString, Optional: true, Description: "'full-text': Full text match; 'pcre': PCRE match;",
			},
			"seq_num": {
				Type: schema.TypeInt, Required: true, Description: "Specify a sequence number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheckHeaderFilterRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheckHeaderFilterRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheckHeaderFilterRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheckHeaderFilterRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateHttpHttpProtocolCheckHeaderFilterRule(d *schema.ResourceData) edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRule {
	var ret edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRule
	ret.Inst.ActionValue = d.Get("action_value").(string)
	ret.Inst.HeaderNameValue = d.Get("header_name_value").(string)
	ret.Inst.HeaderValueValue = d.Get("header_value_value").(string)
	ret.Inst.MatchTypeValue = d.Get("match_type_value").(string)
	ret.Inst.SeqNum = d.Get("seq_num").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
