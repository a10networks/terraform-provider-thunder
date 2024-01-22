package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthorizationPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authorization_policy`: Authorization-policy configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthorizationPolicyCreate,
		UpdateContext: resourceAamAuthorizationPolicyUpdate,
		ReadContext:   resourceAamAuthorizationPolicyRead,
		DeleteContext: resourceAamAuthorizationPolicyDelete,

		Schema: map[string]*schema.Schema{
			"attribute_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_num": {
							Type: schema.TypeInt, Required: true, Description: "Set attribute ID for authorization policy",
						},
						"attribute_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify attribute name",
						},
						"any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Matched when attribute is present (with any value).",
						},
						"attr_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify attribute type",
						},
						"string_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Attribute type is string",
						},
						"integer_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Attribute type is integer",
						},
						"ip_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP address is transformed into network byte order",
						},
						"number_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Attribute type is decimal number",
						},
						"attr_str": {
							Type: schema.TypeString, Optional: true, Description: "'match': Operation type is match; 'sub-string': Operation type is sub-string;",
						},
						"attr_str_val": {
							Type: schema.TypeString, Optional: true, Description: "Set attribute value",
						},
						"attr_int": {
							Type: schema.TypeString, Optional: true, Description: "'equal': Operation type is equal; 'not-equal': Operation type is not equal; 'less-than': Operation type is less-than; 'more-than': Operation type is more-than; 'less-than-equal-to': Operation type is less-than-equal-to; 'more-than-equal-to': Operation type is more-thatn-equal-to;",
						},
						"attr_int_val": {
							Type: schema.TypeInt, Optional: true, Description: "Set attribute value",
						},
						"attr_ip": {
							Type: schema.TypeString, Optional: true, Description: "'equal': Operation type is equal; 'not-equal': Operation type is not-equal;",
						},
						"attr_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 address",
						},
						"attr_number": {
							Type: schema.TypeString, Optional: true, Description: "'equal': Operation type is equal; 'not-equal': Operation type is not equal; 'less-than': Operation type is less-than; 'more-than': Operation type is more-than; 'less-than-equal-to': Operation type is less-than-equal-to; 'more-than-equal-to': Operation type is more-thatn-equal-to;",
						},
						"attr_number_val": {
							Type: schema.TypeString, Optional: true, Description: "Set attribute value",
						},
						"a10_ax_auth_uri": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Custom-defined attribute",
						},
						"custom_attr_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify attribute type",
						},
						"custom_attr_str": {
							Type: schema.TypeString, Optional: true, Description: "'match': Operation type is match; 'sub-string': Operation type is sub-string;",
						},
						"a10_dynamic_defined": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "The value of this attribute will depend on AX configuration instead of user configuration",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"attribute_rule": {
				Type: schema.TypeString, Optional: true, Description: "Define attribute rule for authorization policy",
			},
			"extended_filter": {
				Type: schema.TypeString, Optional: true, Description: "Extended search filter. EX: Check whether user belongs to a nested group. (memberOf:1.2.840.113556.1.4.1941:=$GROUP-DN)",
			},
			"forward_policy_authorize_only": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "This policy only provides server info for forward policy feature",
			},
			"jwt_authorization": {
				Type: schema.TypeString, Optional: true, Description: "Specify JWT authorization template (Specify JWT authorization template name)",
			},
			"jwt_claim_map_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_num": {
							Type: schema.TypeInt, Required: true, Description: "Spcify attribute ID for claim mapping",
						},
						"claim": {
							Type: schema.TypeString, Optional: true, Description: "Specify JWT claim name to map to.",
						},
						"type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify claim type",
						},
						"string_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Claim type is string",
						},
						"number_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Claim type is number",
						},
						"boolean_type": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Claim type is boolean",
						},
						"str_val": {
							Type: schema.TypeString, Optional: true, Description: "Specify JWT claim value.",
						},
						"num_val": {
							Type: schema.TypeInt, Optional: true, Description: "Specify JWT claim value.",
						},
						"bool_val": {
							Type: schema.TypeString, Optional: true, Description: "'true': True; 'false': False;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify authorization policy name",
			},
			"server": {
				Type: schema.TypeString, Optional: true, Description: "Specify a LDAP or RADIUS server for authorization (Specify a LDAP or RADIUS server name)",
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Specify an authentication service group for authorization (Specify authentication service group name)",
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
func resourceAamAuthorizationPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthorizationPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthorizationPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthorizationPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthorizationPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthorizationPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAamAuthorizationPolicyAttributeList(d []interface{}) []edpt.AamAuthorizationPolicyAttributeList {

	count1 := len(d)
	ret := make([]edpt.AamAuthorizationPolicyAttributeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthorizationPolicyAttributeList
		oi.AttrNum = in["attr_num"].(int)
		oi.AttributeName = in["attribute_name"].(string)
		oi.Any = in["any"].(int)
		oi.AttrType = in["attr_type"].(int)
		oi.StringType = in["string_type"].(int)
		oi.IntegerType = in["integer_type"].(int)
		oi.IpType = in["ip_type"].(int)
		oi.NumberType = in["number_type"].(int)
		oi.AttrStr = in["attr_str"].(string)
		oi.AttrStrVal = in["attr_str_val"].(string)
		oi.AttrInt = in["attr_int"].(string)
		oi.AttrIntVal = in["attr_int_val"].(int)
		oi.AttrIp = in["attr_ip"].(string)
		oi.AttrIpv4 = in["attr_ipv4"].(string)
		oi.AttrNumber = in["attr_number"].(string)
		oi.AttrNumberVal = in["attr_number_val"].(string)
		oi.A10AxAuthUri = in["a10_ax_auth_uri"].(int)
		oi.CustomAttrType = in["custom_attr_type"].(int)
		oi.CustomAttrStr = in["custom_attr_str"].(string)
		oi.A10DynamicDefined = in["a10_dynamic_defined"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceAamAuthorizationPolicyJwtClaimMapList(d []interface{}) []edpt.AamAuthorizationPolicyJwtClaimMapList {

	count1 := len(d)
	ret := make([]edpt.AamAuthorizationPolicyJwtClaimMapList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AamAuthorizationPolicyJwtClaimMapList
		oi.AttrNum = in["attr_num"].(int)
		oi.Claim = in["claim"].(string)
		oi.Type = in["type"].(int)
		oi.StringType = in["string_type"].(int)
		oi.NumberType = in["number_type"].(int)
		oi.BooleanType = in["boolean_type"].(int)
		oi.StrVal = in["str_val"].(string)
		oi.NumVal = in["num_val"].(int)
		oi.BoolVal = in["bool_val"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAamAuthorizationPolicy(d *schema.ResourceData) edpt.AamAuthorizationPolicy {
	var ret edpt.AamAuthorizationPolicy
	ret.Inst.AttributeList = getSliceAamAuthorizationPolicyAttributeList(d.Get("attribute_list").([]interface{}))
	ret.Inst.AttributeRule = d.Get("attribute_rule").(string)
	ret.Inst.ExtendedFilter = d.Get("extended_filter").(string)
	ret.Inst.ForwardPolicyAuthorizeOnly = d.Get("forward_policy_authorize_only").(int)
	ret.Inst.JwtAuthorization = d.Get("jwt_authorization").(string)
	ret.Inst.JwtClaimMapList = getSliceAamAuthorizationPolicyJwtClaimMapList(d.Get("jwt_claim_map_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Server = d.Get("server").(string)
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
