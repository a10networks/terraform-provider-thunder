package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateIcmpV4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_icmp_v4`: ICMPv4 template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateIcmpV4Create,
		UpdateContext: resourceDdosTemplateIcmpV4Update,
		ReadContext:   resourceDdosTemplateIcmpV4Read,
		DeleteContext: resourceDdosTemplateIcmpV4Delete,

		Schema: map[string]*schema.Schema{
			"icmp_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "DDOS ICMPv4 Template Name",
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
func resourceDdosTemplateIcmpV4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV4Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateIcmpV4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV4Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateIcmpV4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateIcmpV4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosTemplateIcmpV4TypeList(d []interface{}) []edpt.DdosTemplateIcmpV4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateIcmpV4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateIcmpV4TypeList
		oi.TypeNumber = in["type_number"].(int)
		oi.TypeDeny = in["type_deny"].(int)
		oi.TypeRate = in["type_rate"].(int)
		oi.Code = getSliceDdosTemplateIcmpV4TypeListCode(in["code"].([]interface{}))
		oi.CodeOther = getObjectDdosTemplateIcmpV4TypeListCodeOther(in["code_other"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosTemplateIcmpV4TypeListCode(d []interface{}) []edpt.DdosTemplateIcmpV4TypeListCode {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateIcmpV4TypeListCode, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateIcmpV4TypeListCode
		oi.CodeNumber = in["code_number"].(int)
		oi.CodeRate = in["code_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateIcmpV4TypeListCodeOther(d []interface{}) edpt.DdosTemplateIcmpV4TypeListCodeOther {

	count1 := len(d)
	var ret edpt.DdosTemplateIcmpV4TypeListCodeOther
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CodeOtherRate = in["code_other_rate"].(int)
	}
	return ret
}

func getObjectDdosTemplateIcmpV4TypeOther296(d []interface{}) edpt.DdosTemplateIcmpV4TypeOther296 {

	count1 := len(d)
	var ret edpt.DdosTemplateIcmpV4TypeOther296
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TypeOtherDeny = in["type_other_deny"].(int)
		ret.TypeOtherRate = in["type_other_rate"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosTemplateIcmpV4(d *schema.ResourceData) edpt.DdosTemplateIcmpV4 {
	var ret edpt.DdosTemplateIcmpV4
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	ret.Inst.TypeList = getSliceDdosTemplateIcmpV4TypeList(d.Get("type_list").([]interface{}))
	ret.Inst.TypeOther = getObjectDdosTemplateIcmpV4TypeOther296(d.Get("type_other").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
