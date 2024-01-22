package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthenticationRelayFormBasedInstance() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authentication_relay_form_based_instance`: Form-based Authentication Relay\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthenticationRelayFormBasedInstanceCreate,
		UpdateContext: resourceAamAuthenticationRelayFormBasedInstanceUpdate,
		ReadContext:   resourceAamAuthenticationRelayFormBasedInstanceRead,
		DeleteContext: resourceAamAuthenticationRelayFormBasedInstanceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify form-based authentication relay name",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"request_uri_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_type": {
							Type: schema.TypeString, Required: true, Description: "'equals': URI exactly matches the string; 'contains': URI string contains another sub string; 'starts-with': URI string starts with sub string; 'ends-with': URI string ends with sub string;",
						},
						"uri": {
							Type: schema.TypeString, Required: true, Description: "Specify request URI",
						},
						"user_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify username variable name",
						},
						"password_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify password variable name",
						},
						"domain_variable": {
							Type: schema.TypeString, Optional: true, Description: "Specify domain variable name",
						},
						"other_variables": {
							Type: schema.TypeString, Optional: true, Description: "Specify other variables (n1=v1&n2=v2) in form relay",
						},
						"max_packet_collect_size": {
							Type: schema.TypeInt, Optional: true, Default: 1048576, Description: "Specify the max packet collection size in bytes, default is 1MB",
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
						"action_uri": {
							Type: schema.TypeString, Optional: true, Description: "Specify the action-URI",
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
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'request': Request; 'invalid_srv_rsp': Invalid Server Response; 'post_fail': POST Failed; 'invalid_cred': Invalid Credential; 'bad_req': Bad Request; 'not_fnd': Not Found; 'error': Internal Server Error; 'other_error': Other Error;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAamAuthenticationRelayFormBasedInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstance(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayFormBasedInstanceRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthenticationRelayFormBasedInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstance(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthenticationRelayFormBasedInstanceRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthenticationRelayFormBasedInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstance(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthenticationRelayFormBasedInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthenticationRelayFormBasedInstanceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthenticationRelayFormBasedInstance(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthenticationRelayFormBasedInstanceRequestUriList(d []interface{}) []edpt.AamAuthenticationRelayFormBasedInstanceRequestUriList {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayFormBasedInstanceRequestUriList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayFormBasedInstanceRequestUriList
		oi.MatchType = in["match_type"].(string)
		oi.Uri = in["uri"].(string)
		oi.UserVariable = in["user_variable"].(string)
		oi.PasswordVariable = in["password_variable"].(string)
		oi.DomainVariable = in["domain_variable"].(string)
		oi.OtherVariables = in["other_variables"].(string)
		oi.MaxPacketCollectSize = in["max_packet_collect_size"].(int)
		oi.Cookie = getObjectAamAuthenticationRelayFormBasedInstanceRequestUriListCookie(in["cookie"].([]interface{}))
		oi.ActionUri = in["action_uri"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAamAuthenticationRelayFormBasedInstanceRequestUriListCookie(d []interface{}) edpt.AamAuthenticationRelayFormBasedInstanceRequestUriListCookie {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayFormBasedInstanceRequestUriListCookie
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CookieValue = getObjectAamAuthenticationRelayFormBasedInstanceRequestUriListCookieCookieValue(in["cookie_value"].([]interface{}))
	}
	return ret
}

func getObjectAamAuthenticationRelayFormBasedInstanceRequestUriListCookieCookieValue(d []interface{}) edpt.AamAuthenticationRelayFormBasedInstanceRequestUriListCookieCookieValue {

	count1 := len(d)
	var ret edpt.AamAuthenticationRelayFormBasedInstanceRequestUriListCookieCookieValue
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CookieValue = in["cookie_value"].(string)
	}
	return ret
}

func getSliceAamAuthenticationRelayFormBasedInstanceSamplingEnable(d []interface{}) []edpt.AamAuthenticationRelayFormBasedInstanceSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.AamAuthenticationRelayFormBasedInstanceSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthenticationRelayFormBasedInstanceSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthenticationRelayFormBasedInstance(d *schema.ResourceData) edpt.AamAuthenticationRelayFormBasedInstance {
	var ret edpt.AamAuthenticationRelayFormBasedInstance
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.RequestUriList = getSliceAamAuthenticationRelayFormBasedInstanceRequestUriList(d.Get("request_uri_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceAamAuthenticationRelayFormBasedInstanceSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
