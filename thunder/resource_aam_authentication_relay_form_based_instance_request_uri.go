package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayFormBasedInstanceRequestUri() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_form_based_instance_request_uri`: URI of authentication web page\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayFormBasedInstanceRequestUriCreate,
		UpdateContext: resourceAamAuthenticationRelayFormBasedInstanceRequestUriUpdate,
		ReadContext:   resourceAamAuthenticationRelayFormBasedInstanceRequestUriRead,
		DeleteContext: resourceAamAuthenticationRelayFormBasedInstanceRequestUriDelete,

		Schema: map[string]*schema.Schema{
			"action_uri": {
				Type: schema.TypeString, Optional: true, Description: "Specify the action-URI",
			},
			"cookie": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cookie_value": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cookie_value": {
										Type: schema.TypeString, Optional: true, Description: "Specify cookie in POST packet",
									},
								},
							},
						},
					},
				},
			},
			"domain_variable": {
				Type: schema.TypeString, Optional: true, Description: "Specify domain variable name",
			},
			"match_type": {
				Type: schema.TypeString, Required: true, Description: "'equals': URI exactly matches the string; 'contains': URI string contains another sub string; 'starts-with': URI string starts with sub string; 'ends-with': URI string ends with sub string;",
			},
			"max_packet_collect_size": {
				Type: schema.TypeInt, Optional: true, Default: 1048576, Description: "Specify the max packet collection size in bytes, default is 1MB",
			},
			"other_variables": {
				Type: schema.TypeString, Optional: true, Description: "Specify other variables (n1=v1&n2=v2) in form relay",
			},
			"password_variable": {
				Type: schema.TypeString, Optional: true, Description: "Specify password variable name",
			},
			"uri": {
				Type: schema.TypeString, Required: true, Description: "Specify request URI",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"user_variable": {
				Type: schema.TypeString, Optional: true, Description: "Specify username variable name",
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
func resourceAamAuthenticationRelayFormBasedInstanceRequestUriCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceRequestUriCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstanceRequestUri(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayFormBasedInstanceRequestUriRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayFormBasedInstanceRequestUriUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceRequestUriUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstanceRequestUri(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayFormBasedInstanceRequestUriRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayFormBasedInstanceRequestUriDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceRequestUriDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstanceRequestUri(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayFormBasedInstanceRequestUriRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceRequestUriRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstanceRequestUri(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAamAuthenticationRelayFormBasedInstanceRequestUriCookie(d []interface{}) edpt.AamAuthenticationRelayFormBasedInstanceRequestUriCookie {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayFormBasedInstanceRequestUriCookie
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CookieValue = getObjectAamAuthenticationRelayFormBasedInstanceRequestUriCookieCookieValue(in["cookie_value"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationRelayFormBasedInstanceRequestUriCookieCookieValue(d []interface{}) edpt.AamAuthenticationRelayFormBasedInstanceRequestUriCookieCookieValue {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayFormBasedInstanceRequestUriCookieCookieValue
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CookieValue = in["cookie_value"].(string)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayFormBasedInstanceRequestUri(d *schema.ResourceData) edpt.AamAuthenticationRelayFormBasedInstanceRequestUri {
	var ret edpt.AamAuthenticationRelayFormBasedInstanceRequestUri
	ret.Inst.ActionUri = d.Get("action_uri").(string)
	ret.Inst.Cookie = getObjectAamAuthenticationRelayFormBasedInstanceRequestUriCookie(d.Get("cookie").([]interface{}))
	ret.Inst.DomainVariable = d.Get("domain_variable").(string)
	ret.Inst.MatchType = d.Get("match_type").(string)
	ret.Inst.MaxPacketCollectSize = d.Get("max_packet_collect_size").(int)
	ret.Inst.OtherVariables = d.Get("other_variables").(string)
	ret.Inst.PasswordVariable = d.Get("password_variable").(string)
	ret.Inst.Uri = d.Get("uri").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	ret.Inst.UserVariable = d.Get("user_variable").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
