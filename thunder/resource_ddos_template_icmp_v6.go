package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateIcmpV6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_icmp_v6`: ICMPv6 template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateIcmpV6Create,
		UpdateContext: resourceDdosTemplateIcmpV6Update,
		ReadContext:   resourceDdosTemplateIcmpV6Read,
		DeleteContext: resourceDdosTemplateIcmpV6Delete,

		Schema: map[string]*schema.Schema{
			"icmp_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "DDOS ICMPv6 Template Name",
			},
			"type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type_number": {
							Type: schema.TypeInt, Required: true, Description: "Specify ICMP type number",
						},
						"type_deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reject this ICMP type",
						},
						"type_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the whole rate with this type",
						},
						"code": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code_number": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP code",
									},
									"code_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Specify the rate with the code",
									},
								},
							},
						},
						"code_other": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code_other_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Specify rate with other code",
									},
								},
							},
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
			"type_other": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type_other_deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Deny all other type",
						},
						"type_other_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify rate with other type",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
func resourceDdosTemplateIcmpV6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV6Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateIcmpV6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV6Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateIcmpV6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateIcmpV6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosTemplateIcmpV6TypeList(d []interface{}) []edpt.DdosTemplateIcmpV6TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateIcmpV6TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateIcmpV6TypeList
		oi.TypeNumber = in["type_number"].(int)
		oi.TypeDeny = in["type_deny"].(int)
		oi.TypeRate = in["type_rate"].(int)
		oi.Code = getSliceDdosTemplateIcmpV6TypeListCode(in["code"].([]interface{}))
		oi.CodeOther = getObjectDdosTemplateIcmpV6TypeListCodeOther(in["code_other"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateIcmpV6TypeListCode(d []interface{}) []edpt.DdosTemplateIcmpV6TypeListCode {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateIcmpV6TypeListCode, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateIcmpV6TypeListCode
		oi.CodeNumber = in["code_number"].(int)
		oi.CodeRate = in["code_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateIcmpV6TypeListCodeOther(d []interface{}) edpt.DdosTemplateIcmpV6TypeListCodeOther {

	count1 := len(d)
	var ret edpt.DdosTemplateIcmpV6TypeListCodeOther
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CodeOtherRate = in["code_other_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateIcmpV6TypeOther297(d []interface{}) edpt.DdosTemplateIcmpV6TypeOther297 {

	count1 := len(d)
	var ret edpt.DdosTemplateIcmpV6TypeOther297
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TypeOtherDeny = in["type_other_deny"].(int)
		ret.TypeOtherRate = in["type_other_rate"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosTemplateIcmpV6(d *schema.ResourceData) edpt.DdosTemplateIcmpV6 {
	var ret edpt.DdosTemplateIcmpV6
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	ret.Inst.TypeList = getSliceDdosTemplateIcmpV6TypeList(d.Get("type_list").([]interface{}))
	ret.Inst.TypeOther = getObjectDdosTemplateIcmpV6TypeOther297(d.Get("type_other").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
