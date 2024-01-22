package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAuthorizationPolicyAttribute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_aam_authorization_policy_attribute`: Authorization-policy attribute configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceAamAuthorizationPolicyAttributeCreate,
		UpdateContext: resourceAamAuthorizationPolicyAttributeUpdate,
		ReadContext:   resourceAamAuthorizationPolicyAttributeRead,
		DeleteContext: resourceAamAuthorizationPolicyAttributeDelete,

		Schema: map[string]*schema.Schema{
			"a10_ax_auth_uri": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Custom-defined attribute",
			},
			"a10_dynamic_defined": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The value of this attribute will depend on AX configuration instead of user configuration",
			},
			"any": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Matched when attribute is present (with any value).",
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
			"attr_num": {
				Type: schema.TypeInt, Required: true, Description: "Set attribute ID for authorization policy",
			},
			"attr_number": {
				Type: schema.TypeString, Optional: true, Description: "'equal': Operation type is equal; 'not-equal': Operation type is not equal; 'less-than': Operation type is less-than; 'more-than': Operation type is more-than; 'less-than-equal-to': Operation type is less-than-equal-to; 'more-than-equal-to': Operation type is more-thatn-equal-to;",
			},
			"attr_number_val": {
				Type: schema.TypeString, Optional: true, Description: "Set attribute value",
			},
			"attr_str": {
				Type: schema.TypeString, Optional: true, Description: "'match': Operation type is match; 'sub-string': Operation type is sub-string;",
			},
			"attr_str_val": {
				Type: schema.TypeString, Optional: true, Description: "Set attribute value",
			},
			"attr_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify attribute type",
			},
			"attribute_name": {
				Type: schema.TypeString, Optional: true, Description: "Specify attribute name",
			},
			"custom_attr_str": {
				Type: schema.TypeString, Optional: true, Description: "'match': Operation type is match; 'sub-string': Operation type is sub-string;",
			},
			"custom_attr_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify attribute type",
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
			"string_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Attribute type is string",
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
func resourceAamAuthorizationPolicyAttributeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyAttributeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyAttribute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthorizationPolicyAttributeRead(ctx, d, meta)
	}
	return diags
}

func resourceAamAuthorizationPolicyAttributeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyAttributeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyAttribute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAamAuthorizationPolicyAttributeRead(ctx, d, meta)
	}
	return diags
}
func resourceAamAuthorizationPolicyAttributeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyAttributeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyAttribute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAamAuthorizationPolicyAttributeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAuthorizationPolicyAttributeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAuthorizationPolicyAttribute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAamAuthorizationPolicyAttribute(d *schema.ResourceData) edpt.AamAuthorizationPolicyAttribute {
	var ret edpt.AamAuthorizationPolicyAttribute
	ret.Inst.A10AxAuthUri = d.Get("a10_ax_auth_uri").(int)
	ret.Inst.A10DynamicDefined = d.Get("a10_dynamic_defined").(int)
	ret.Inst.Any = d.Get("any").(int)
	ret.Inst.AttrInt = d.Get("attr_int").(string)
	ret.Inst.AttrIntVal = d.Get("attr_int_val").(int)
	ret.Inst.AttrIp = d.Get("attr_ip").(string)
	ret.Inst.AttrIpv4 = d.Get("attr_ipv4").(string)
	ret.Inst.AttrNum = d.Get("attr_num").(int)
	ret.Inst.AttrNumber = d.Get("attr_number").(string)
	ret.Inst.AttrNumberVal = d.Get("attr_number_val").(string)
	ret.Inst.AttrStr = d.Get("attr_str").(string)
	ret.Inst.AttrStrVal = d.Get("attr_str_val").(string)
	ret.Inst.AttrType = d.Get("attr_type").(int)
	ret.Inst.AttributeName = d.Get("attribute_name").(string)
	ret.Inst.CustomAttrStr = d.Get("custom_attr_str").(string)
	ret.Inst.CustomAttrType = d.Get("custom_attr_type").(int)
	ret.Inst.IntegerType = d.Get("integer_type").(int)
	ret.Inst.IpType = d.Get("ip_type").(int)
	ret.Inst.NumberType = d.Get("number_type").(int)
	ret.Inst.StringType = d.Get("string_type").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
