package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateHttpHttpProtocolCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_http_http_protocol_check`: HTTP protocol compliance check\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateHttpHttpProtocolCheckCreate,
		UpdateContext: resourceSlbTemplateHttpHttpProtocolCheckUpdate,
		ReadContext:   resourceSlbTemplateHttpHttpProtocolCheckRead,
		DeleteContext: resourceSlbTemplateHttpHttpProtocolCheckDelete,

		Schema: map[string]*schema.Schema{
			"get_and_payload": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"h2up_content_length_alias": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"h2up_with_host_and_auth": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"h2up_with_transfer_encoding": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"header_filter_rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": {
							Type: schema.TypeInt, Required: true, Description: "Specify a sequence number",
						},
						"match_type_value": {
							Type: schema.TypeString, Optional: true, Description: "'full-text': Full text match; 'pcre': PCRE match;",
						},
						"header_name_value": {
							Type: schema.TypeString, Optional: true, Description: "Header name value",
						},
						"header_value_value": {
							Type: schema.TypeString, Optional: true, Description: "Header value",
						},
						"action_value": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"malformed_h2up_header_value": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"malformed_h2up_scheme_value": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"multiple_content_length": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"multiple_transfer_encoding": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and send 400 to the client side;",
			},
			"transfer_encoding_and_content_length": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop the request and Send 400 to the client side;",
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
func resourceSlbTemplateHttpHttpProtocolCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpHttpProtocolCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateHttpHttpProtocolCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateHttpHttpProtocolCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateHttpHttpProtocolCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateHttpHttpProtocolCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateHttpHttpProtocolCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateHttpHttpProtocolCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList(d []interface{}) []edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList
		oi.SeqNum = in["seq_num"].(int)
		oi.MatchTypeValue = in["match_type_value"].(string)
		oi.HeaderNameValue = in["header_name_value"].(string)
		oi.HeaderValueValue = in["header_value_value"].(string)
		oi.ActionValue = in["action_value"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateHttpHttpProtocolCheck(d *schema.ResourceData) edpt.SlbTemplateHttpHttpProtocolCheck {
	var ret edpt.SlbTemplateHttpHttpProtocolCheck
	ret.Inst.GetAndPayload = d.Get("get_and_payload").(string)
	ret.Inst.H2upContentLengthAlias = d.Get("h2up_content_length_alias").(string)
	ret.Inst.H2upWithHostAndAuth = d.Get("h2up_with_host_and_auth").(string)
	ret.Inst.H2upWithTransferEncoding = d.Get("h2up_with_transfer_encoding").(string)
	ret.Inst.HeaderFilterRuleList = getSliceSlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList(d.Get("header_filter_rule_list").([]interface{}))
	ret.Inst.MalformedH2upHeaderValue = d.Get("malformed_h2up_header_value").(string)
	ret.Inst.MalformedH2upSchemeValue = d.Get("malformed_h2up_scheme_value").(string)
	ret.Inst.MultipleContentLength = d.Get("multiple_content_length").(string)
	ret.Inst.MultipleTransferEncoding = d.Get("multiple_transfer_encoding").(string)
	ret.Inst.TransferEncodingAndContentLength = d.Get("transfer_encoding_and_content_length").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
