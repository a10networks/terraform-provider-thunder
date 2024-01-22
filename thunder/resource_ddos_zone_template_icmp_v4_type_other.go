package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIcmpV4TypeOther() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_icmp_v4_type_other`: Specify other type\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIcmpV4TypeOtherCreate,
		UpdateContext: resourceDdosZoneTemplateIcmpV4TypeOtherUpdate,
		ReadContext:   resourceDdosZoneTemplateIcmpV4TypeOtherRead,
		DeleteContext: resourceDdosZoneTemplateIcmpV4TypeOtherDelete,

		Schema: map[string]*schema.Schema{
			"dst": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_type_other_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the whole dst rate for wildcard ICMP type",
						},
						"dst_type_other_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for rate exceed",
						},
						"dst_type_other_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
					},
				},
			},
			"icmp_type_other_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Reject wildcard ICMP type; 'blacklist-src': Blacklist-src wildcard ICMP type; 'ignore': Ignore wildcard ICMP type;",
			},
			"icmp_type_other_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for wildcard ICMP match",
			},
			"src": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_type_other_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the whole src rate for wildcard ICMP type",
						},
						"src_type_other_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for rate exceed",
						},
						"src_type_other_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
					},
				},
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
func resourceDdosZoneTemplateIcmpV4TypeOtherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeOtherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4TypeOther(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4TypeOtherRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4TypeOtherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeOtherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4TypeOther(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4TypeOtherRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIcmpV4TypeOtherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeOtherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4TypeOther(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4TypeOtherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4TypeOtherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4TypeOther(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosZoneTemplateIcmpV4TypeOtherDst(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeOtherDst {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeOtherDst
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstTypeOtherRate = in["dst_type_other_rate"].(int)
		ret.DstTypeOtherRateActionListName = in["dst_type_other_rate_action_list_name"].(string)
		ret.DstTypeOtherRateAction = in["dst_type_other_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeOtherSrc(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeOtherSrc {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeOtherSrc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcTypeOtherRate = in["src_type_other_rate"].(int)
		ret.SrcTypeOtherRateActionListName = in["src_type_other_rate_action_list_name"].(string)
		ret.SrcTypeOtherRateAction = in["src_type_other_rate_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateIcmpV4TypeOther(d *schema.ResourceData) edpt.DdosZoneTemplateIcmpV4TypeOther {
	var ret edpt.DdosZoneTemplateIcmpV4TypeOther
	ret.Inst.Dst = getObjectDdosZoneTemplateIcmpV4TypeOtherDst(d.Get("dst").([]interface{}))
	ret.Inst.IcmpTypeOtherAction = d.Get("icmp_type_other_action").(string)
	ret.Inst.IcmpTypeOtherActionListName = d.Get("icmp_type_other_action_list_name").(string)
	ret.Inst.Src = getObjectDdosZoneTemplateIcmpV4TypeOtherSrc(d.Get("src").([]interface{}))
	//omit uuid
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	return ret
}
