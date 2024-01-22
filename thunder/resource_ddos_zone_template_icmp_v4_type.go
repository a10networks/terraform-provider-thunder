package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIcmpV4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_icmp_v4_type`: Specify ICMP type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIcmpV4TypeCreate,
		UpdateContext: resourceDdosZoneTemplateIcmpV4TypeUpdate,
		ReadContext:   resourceDdosZoneTemplateIcmpV4TypeRead,
		DeleteContext: resourceDdosZoneTemplateIcmpV4TypeDelete,

		Schema: map[string]*schema.Schema{
			"dst_code_other_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the rate with other code",
			},
			"dst_code_other_rate_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
			},
			"dst_code_other_rate_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for  rate exceed",
			},
			"icmp_type_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Reject this ICMP type; 'blacklist-src': Blacklist-src this ICMP type; 'ignore': Ignore this ICMP type;",
			},
			"icmp_type_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for this ICMP type",
			},
			"src_code_other_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the rate with other code",
			},
			"src_code_other_rate_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
			},
			"src_code_other_rate_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for  rate exceed",
			},
			"type_number": {
				Type: schema.TypeInt, Required: true, Description: "Specify ICMP type number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_dst_code_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_code_number": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP code for this dst rate",
						},
						"dst_code_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the rate with the code",
						},
						"dst_code_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for rate exceed",
						},
						"dst_code_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
					},
				},
			},
			"v4_dst_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_type_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the whole dst rate for this type",
						},
						"dst_type_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for rate exceed",
						},
						"dst_type_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
					},
				},
			},
			"v4_src_code_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_code_number": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP code for this src rate",
						},
						"src_code_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the rate with the code",
						},
						"src_code_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for rate exceed",
						},
						"src_code_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
					},
				},
			},
			"v4_src_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_type_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the whole src rate for this type",
						},
						"src_type_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for rate exceed",
						},
						"src_type_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
					},
				},
			},
			"icmp_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "IcmpTmplName",
			},
		},
	}
}
func resourceDdosZoneTemplateIcmpV4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIcmpV4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneTemplateIcmpV4TypeV4DstCodeCfg(d []interface{}) []edpt.DdosZoneTemplateIcmpV4TypeV4DstCodeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV4TypeV4DstCodeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV4TypeV4DstCodeCfg
		oi.DstCodeNumber = in["dst_code_number"].(int)
		oi.DstCodeRate = in["dst_code_rate"].(int)
		oi.DstCodeRateActionListName = in["dst_code_rate_action_list_name"].(string)
		oi.DstCodeRateAction = in["dst_code_rate_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeV4DstRateCfg(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeV4DstRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeV4DstRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstTypeRate = in["dst_type_rate"].(int)
		ret.DstTypeRateActionListName = in["dst_type_rate_action_list_name"].(string)
		ret.DstTypeRateAction = in["dst_type_rate_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateIcmpV4TypeV4SrcCodeCfg(d []interface{}) []edpt.DdosZoneTemplateIcmpV4TypeV4SrcCodeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV4TypeV4SrcCodeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV4TypeV4SrcCodeCfg
		oi.SrcCodeNumber = in["src_code_number"].(int)
		oi.SrcCodeRate = in["src_code_rate"].(int)
		oi.SrcCodeRateActionListName = in["src_code_rate_action_list_name"].(string)
		oi.SrcCodeRateAction = in["src_code_rate_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeV4SrcRateCfg(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeV4SrcRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeV4SrcRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcTypeRate = in["src_type_rate"].(int)
		ret.SrcTypeRateActionListName = in["src_type_rate_action_list_name"].(string)
		ret.SrcTypeRateAction = in["src_type_rate_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateIcmpV4Type(d *schema.ResourceData) edpt.DdosZoneTemplateIcmpV4Type {
	var ret edpt.DdosZoneTemplateIcmpV4Type
	ret.Inst.DstCodeOtherRate = d.Get("dst_code_other_rate").(int)
	ret.Inst.DstCodeOtherRateAction = d.Get("dst_code_other_rate_action").(string)
	ret.Inst.DstCodeOtherRateActionListName = d.Get("dst_code_other_rate_action_list_name").(string)
	ret.Inst.IcmpTypeAction = d.Get("icmp_type_action").(string)
	ret.Inst.IcmpTypeActionListName = d.Get("icmp_type_action_list_name").(string)
	ret.Inst.SrcCodeOtherRate = d.Get("src_code_other_rate").(int)
	ret.Inst.SrcCodeOtherRateAction = d.Get("src_code_other_rate_action").(string)
	ret.Inst.SrcCodeOtherRateActionListName = d.Get("src_code_other_rate_action_list_name").(string)
	ret.Inst.TypeNumber = d.Get("type_number").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.V4DstCodeCfg = getSliceDdosZoneTemplateIcmpV4TypeV4DstCodeCfg(d.Get("v4_dst_code_cfg").([]interface{}))
	ret.Inst.V4DstRateCfg = getObjectDdosZoneTemplateIcmpV4TypeV4DstRateCfg(d.Get("v4_dst_rate_cfg").([]interface{}))
	ret.Inst.V4SrcCodeCfg = getSliceDdosZoneTemplateIcmpV4TypeV4SrcCodeCfg(d.Get("v4_src_code_cfg").([]interface{}))
	ret.Inst.V4SrcRateCfg = getObjectDdosZoneTemplateIcmpV4TypeV4SrcRateCfg(d.Get("v4_src_rate_cfg").([]interface{}))
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
