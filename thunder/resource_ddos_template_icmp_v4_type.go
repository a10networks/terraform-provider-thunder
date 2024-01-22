package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateIcmpV4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_icmp_v4_type`: Specify ICMP type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateIcmpV4TypeCreate,
		UpdateContext: resourceDdosTemplateIcmpV4TypeUpdate,
		ReadContext:   resourceDdosTemplateIcmpV4TypeRead,
		DeleteContext: resourceDdosTemplateIcmpV4TypeDelete,

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
func resourceDdosTemplateIcmpV4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateIcmpV4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateIcmpV4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateIcmpV4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateIcmpV4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateIcmpV4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateIcmpV4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosTemplateIcmpV4TypeCode(d []interface{}) []edpt.DdosTemplateIcmpV4TypeCode {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateIcmpV4TypeCode, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateIcmpV4TypeCode
		oi.CodeNumber = in["code_number"].(int)
		oi.CodeRate = in["code_rate"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateIcmpV4TypeCodeOther(d []interface{}) edpt.DdosTemplateIcmpV4TypeCodeOther {

	count1 := len(d)
	var ret edpt.DdosTemplateIcmpV4TypeCodeOther
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CodeOtherRate = in["code_other_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosTemplateIcmpV4Type(d *schema.ResourceData) edpt.DdosTemplateIcmpV4Type {
	var ret edpt.DdosTemplateIcmpV4Type
	ret.Inst.Code = getSliceDdosTemplateIcmpV4TypeCode(d.Get("code").([]interface{}))
	ret.Inst.CodeOther = getObjectDdosTemplateIcmpV4TypeCodeOther(d.Get("code_other").([]interface{}))
	ret.Inst.TypeDeny = d.Get("type_deny").(int)
	ret.Inst.TypeNumber = d.Get("type_number").(int)
	ret.Inst.TypeRate = d.Get("type_rate").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
