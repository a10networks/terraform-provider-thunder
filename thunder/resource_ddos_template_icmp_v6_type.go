package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateIcmpV6Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_icmp_v6_type`: Specify ICMP type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateIcmpV6TypeCreate,
		UpdateContext: resourceDdosTemplateIcmpV6TypeUpdate,
		ReadContext:   resourceDdosTemplateIcmpV6TypeRead,
		DeleteContext: resourceDdosTemplateIcmpV6TypeDelete,

		Schema: map[string]*schema.Schema{
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
			"type_deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reject this ICMP type",
			},
			"type_number": {
				Type: schema.TypeInt, Required: true, Description: "Specify ICMP type number",
			},
			"type_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the whole rate with this type",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"icmp_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "IcmpTmplName",
			},
		},
	}
}
func resourceDdosTemplateIcmpV6TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV6TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateIcmpV6TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV6TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateIcmpV6TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateIcmpV6TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV6TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV6Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosTemplateIcmpV6TypeCode(d []interface{}) []edpt.DdosTemplateIcmpV6TypeCode {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateIcmpV6TypeCode, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateIcmpV6TypeCode
		oi.CodeNumber = in["code_number"].(int)
		oi.CodeRate = in["code_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateIcmpV6TypeCodeOther(d []interface{}) edpt.DdosTemplateIcmpV6TypeCodeOther {

	count1 := len(d)
	var ret edpt.DdosTemplateIcmpV6TypeCodeOther
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CodeOtherRate = in["code_other_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosTemplateIcmpV6Type(d *schema.ResourceData) edpt.DdosTemplateIcmpV6Type {
	var ret edpt.DdosTemplateIcmpV6Type
	ret.Inst.Code = getSliceDdosTemplateIcmpV6TypeCode(d.Get("code").([]interface{}))
	ret.Inst.CodeOther = getObjectDdosTemplateIcmpV6TypeCodeOther(d.Get("code_other").([]interface{}))
	ret.Inst.TypeDeny = d.Get("type_deny").(int)
	ret.Inst.TypeNumber = d.Get("type_number").(int)
	ret.Inst.TypeRate = d.Get("type_rate").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
