package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIcmpV6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_icmp_v6`: ICMPv6 template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIcmpV6Create,
		UpdateContext: resourceDdosZoneTemplateIcmpV6Update,
		ReadContext:   resourceDdosZoneTemplateIcmpV6Read,
		DeleteContext: resourceDdosZoneTemplateIcmpV6Delete,

		Schema: map[string]*schema.Schema{
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"icmp_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "sequence number",
						},
						"icmp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"icmp_filter_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of matching",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "filter using Berkeley packet filter syntax",
						},
						"icmp_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "list to take",
						},
						"icmp_filter_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src;",
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
			"filter_match_type": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Stop matching on drop/blacklist action; 'stop-on-first-match': Stop matching on first match;",
			},
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
						"icmp_type_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for this ICMP type",
						},
						"icmp_type_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Reject this ICMP type; 'blacklist-src': Blacklist-src this ICMP type; 'ignore': Ignore this ICMP type;",
						},
						"v6_src_rate_cfg": {
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
						"v6_src_code_cfg": {
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
						"src_code_other_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the rate with other code",
						},
						"src_code_other_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for  rate exceed",
						},
						"src_code_other_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
						},
						"v6_dst_rate_cfg": {
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
						"v6_dst_code_cfg": {
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
						"dst_code_other_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the rate with other code",
						},
						"dst_code_other_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for  rate exceed",
						},
						"dst_code_other_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
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
						"icmp_type_other_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for wildcard ICMP match",
						},
						"icmp_type_other_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Reject wildcard ICMP type; 'blacklist-src': Blacklist-src wildcard ICMP type; 'ignore': Ignore wildcard ICMP type;",
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
func resourceDdosZoneTemplateIcmpV6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV6Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV6Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIcmpV6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneTemplateIcmpV6FilterList(d []interface{}) []edpt.DdosZoneTemplateIcmpV6FilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV6FilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV6FilterList
		oi.IcmpFilterName = in["icmp_filter_name"].(string)
		oi.IcmpFilterSeq = in["icmp_filter_seq"].(int)
		oi.IcmpFilterRegex = in["icmp_filter_regex"].(string)
		oi.IcmpFilterInverseMatch = in["icmp_filter_inverse_match"].(int)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.IcmpFilterActionListName = in["icmp_filter_action_list_name"].(string)
		oi.IcmpFilterAction = in["icmp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosZoneTemplateIcmpV6TypeList(d []interface{}) []edpt.DdosZoneTemplateIcmpV6TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV6TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV6TypeList
		oi.TypeNumber = in["type_number"].(int)
		oi.IcmpTypeActionListName = in["icmp_type_action_list_name"].(string)
		oi.IcmpTypeAction = in["icmp_type_action"].(string)
		oi.V6SrcRateCfg = getObjectDdosZoneTemplateIcmpV6TypeListV6SrcRateCfg(in["v6_src_rate_cfg"].([]interface{}))
		oi.V6SrcCodeCfg = getSliceDdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg(in["v6_src_code_cfg"].([]interface{}))
		oi.SrcCodeOtherRate = in["src_code_other_rate"].(int)
		oi.SrcCodeOtherRateActionListName = in["src_code_other_rate_action_list_name"].(string)
		oi.SrcCodeOtherRateAction = in["src_code_other_rate_action"].(string)
		oi.V6DstRateCfg = getObjectDdosZoneTemplateIcmpV6TypeListV6DstRateCfg(in["v6_dst_rate_cfg"].([]interface{}))
		oi.V6DstCodeCfg = getSliceDdosZoneTemplateIcmpV6TypeListV6DstCodeCfg(in["v6_dst_code_cfg"].([]interface{}))
		oi.DstCodeOtherRate = in["dst_code_other_rate"].(int)
		oi.DstCodeOtherRateActionListName = in["dst_code_other_rate_action_list_name"].(string)
		oi.DstCodeOtherRateAction = in["dst_code_other_rate_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV6TypeListV6SrcRateCfg(d []interface{}) edpt.DdosZoneTemplateIcmpV6TypeListV6SrcRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV6TypeListV6SrcRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcTypeRate = in["src_type_rate"].(int)
		ret.SrcTypeRateActionListName = in["src_type_rate_action_list_name"].(string)
		ret.SrcTypeRateAction = in["src_type_rate_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg(d []interface{}) []edpt.DdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV6TypeListV6SrcCodeCfg
		oi.SrcCodeNumber = in["src_code_number"].(int)
		oi.SrcCodeRate = in["src_code_rate"].(int)
		oi.SrcCodeRateActionListName = in["src_code_rate_action_list_name"].(string)
		oi.SrcCodeRateAction = in["src_code_rate_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV6TypeListV6DstRateCfg(d []interface{}) edpt.DdosZoneTemplateIcmpV6TypeListV6DstRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV6TypeListV6DstRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstTypeRate = in["dst_type_rate"].(int)
		ret.DstTypeRateActionListName = in["dst_type_rate_action_list_name"].(string)
		ret.DstTypeRateAction = in["dst_type_rate_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateIcmpV6TypeListV6DstCodeCfg(d []interface{}) []edpt.DdosZoneTemplateIcmpV6TypeListV6DstCodeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV6TypeListV6DstCodeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV6TypeListV6DstCodeCfg
		oi.DstCodeNumber = in["dst_code_number"].(int)
		oi.DstCodeRate = in["dst_code_rate"].(int)
		oi.DstCodeRateActionListName = in["dst_code_rate_action_list_name"].(string)
		oi.DstCodeRateAction = in["dst_code_rate_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV6TypeOther312(d []interface{}) edpt.DdosZoneTemplateIcmpV6TypeOther312 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV6TypeOther312
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpTypeOtherActionListName = in["icmp_type_other_action_list_name"].(string)
		ret.IcmpTypeOtherAction = in["icmp_type_other_action"].(string)
		ret.Src = getObjectDdosZoneTemplateIcmpV6TypeOtherSrc313(in["src"].([]interface{}))
		ret.Dst = getObjectDdosZoneTemplateIcmpV6TypeOtherDst314(in["dst"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV6TypeOtherSrc313(d []interface{}) edpt.DdosZoneTemplateIcmpV6TypeOtherSrc313 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV6TypeOtherSrc313
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcTypeOtherRate = in["src_type_other_rate"].(int)
		ret.SrcTypeOtherRateActionListName = in["src_type_other_rate_action_list_name"].(string)
		ret.SrcTypeOtherRateAction = in["src_type_other_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV6TypeOtherDst314(d []interface{}) edpt.DdosZoneTemplateIcmpV6TypeOtherDst314 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV6TypeOtherDst314
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstTypeOtherRate = in["dst_type_other_rate"].(int)
		ret.DstTypeOtherRateActionListName = in["dst_type_other_rate_action_list_name"].(string)
		ret.DstTypeOtherRateAction = in["dst_type_other_rate_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateIcmpV6(d *schema.ResourceData) edpt.DdosZoneTemplateIcmpV6 {
	var ret edpt.DdosZoneTemplateIcmpV6
	ret.Inst.FilterList = getSliceDdosZoneTemplateIcmpV6FilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	ret.Inst.TypeList = getSliceDdosZoneTemplateIcmpV6TypeList(d.Get("type_list").([]interface{}))
	ret.Inst.TypeOther = getObjectDdosZoneTemplateIcmpV6TypeOther312(d.Get("type_other").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
