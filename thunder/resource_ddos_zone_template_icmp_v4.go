package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateIcmpV4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_icmp_v4`: ICMPv4 template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateIcmpV4Create,
		UpdateContext: resourceDdosZoneTemplateIcmpV4Update,
		ReadContext:   resourceDdosZoneTemplateIcmpV4Read,
		DeleteContext: resourceDdosZoneTemplateIcmpV4Delete,

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
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
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
				Type: schema.TypeString, Required: true, Description: "DDOS ICMPv4 Template Name",
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
						"src_code_other_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the rate with other code",
						},
						"src_code_other_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for  rate exceed",
						},
						"src_code_other_rate_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for rate exceed (Default); 'blacklist-src': Blacklist-src for rate exceed; 'ignore': Do nothing for rate exceed;",
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
func resourceDdosZoneTemplateIcmpV4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateIcmpV4Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateIcmpV4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateIcmpV4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateIcmpV4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateIcmpV4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneTemplateIcmpV4FilterList(d []interface{}) []edpt.DdosZoneTemplateIcmpV4FilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV4FilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV4FilterList
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

func getSliceDdosZoneTemplateIcmpV4TypeList(d []interface{}) []edpt.DdosZoneTemplateIcmpV4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV4TypeList
		oi.TypeNumber = in["type_number"].(int)
		oi.IcmpTypeActionListName = in["icmp_type_action_list_name"].(string)
		oi.IcmpTypeAction = in["icmp_type_action"].(string)
		oi.V4SrcRateCfg = getObjectDdosZoneTemplateIcmpV4TypeListV4SrcRateCfg(in["v4_src_rate_cfg"].([]interface{}))
		oi.V4SrcCodeCfg = getSliceDdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg(in["v4_src_code_cfg"].([]interface{}))
		oi.SrcCodeOtherRate = in["src_code_other_rate"].(int)
		oi.SrcCodeOtherRateActionListName = in["src_code_other_rate_action_list_name"].(string)
		oi.SrcCodeOtherRateAction = in["src_code_other_rate_action"].(string)
		oi.V4DstRateCfg = getObjectDdosZoneTemplateIcmpV4TypeListV4DstRateCfg(in["v4_dst_rate_cfg"].([]interface{}))
		oi.V4DstCodeCfg = getSliceDdosZoneTemplateIcmpV4TypeListV4DstCodeCfg(in["v4_dst_code_cfg"].([]interface{}))
		oi.DstCodeOtherRate = in["dst_code_other_rate"].(int)
		oi.DstCodeOtherRateActionListName = in["dst_code_other_rate_action_list_name"].(string)
		oi.DstCodeOtherRateAction = in["dst_code_other_rate_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeListV4SrcRateCfg(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeListV4SrcRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeListV4SrcRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcTypeRate = in["src_type_rate"].(int)
		ret.SrcTypeRateActionListName = in["src_type_rate_action_list_name"].(string)
		ret.SrcTypeRateAction = in["src_type_rate_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg(d []interface{}) []edpt.DdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV4TypeListV4SrcCodeCfg
		oi.SrcCodeNumber = in["src_code_number"].(int)
		oi.SrcCodeRate = in["src_code_rate"].(int)
		oi.SrcCodeRateActionListName = in["src_code_rate_action_list_name"].(string)
		oi.SrcCodeRateAction = in["src_code_rate_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeListV4DstRateCfg(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeListV4DstRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeListV4DstRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstTypeRate = in["dst_type_rate"].(int)
		ret.DstTypeRateActionListName = in["dst_type_rate_action_list_name"].(string)
		ret.DstTypeRateAction = in["dst_type_rate_action"].(string)
	}
	return ret
}

func getSliceDdosZoneTemplateIcmpV4TypeListV4DstCodeCfg(d []interface{}) []edpt.DdosZoneTemplateIcmpV4TypeListV4DstCodeCfg {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateIcmpV4TypeListV4DstCodeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateIcmpV4TypeListV4DstCodeCfg
		oi.DstCodeNumber = in["dst_code_number"].(int)
		oi.DstCodeRate = in["dst_code_rate"].(int)
		oi.DstCodeRateActionListName = in["dst_code_rate_action_list_name"].(string)
		oi.DstCodeRateAction = in["dst_code_rate_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeOther309(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeOther309 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeOther309
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IcmpTypeOtherActionListName = in["icmp_type_other_action_list_name"].(string)
		ret.IcmpTypeOtherAction = in["icmp_type_other_action"].(string)
		ret.Src = getObjectDdosZoneTemplateIcmpV4TypeOtherSrc310(in["src"].([]interface{}))
		ret.Dst = getObjectDdosZoneTemplateIcmpV4TypeOtherDst311(in["dst"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeOtherSrc310(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeOtherSrc310 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeOtherSrc310
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcTypeOtherRate = in["src_type_other_rate"].(int)
		ret.SrcTypeOtherRateActionListName = in["src_type_other_rate_action_list_name"].(string)
		ret.SrcTypeOtherRateAction = in["src_type_other_rate_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateIcmpV4TypeOtherDst311(d []interface{}) edpt.DdosZoneTemplateIcmpV4TypeOtherDst311 {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateIcmpV4TypeOtherDst311
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DstTypeOtherRate = in["dst_type_other_rate"].(int)
		ret.DstTypeOtherRateActionListName = in["dst_type_other_rate_action_list_name"].(string)
		ret.DstTypeOtherRateAction = in["dst_type_other_rate_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateIcmpV4(d *schema.ResourceData) edpt.DdosZoneTemplateIcmpV4 {
	var ret edpt.DdosZoneTemplateIcmpV4
	ret.Inst.FilterList = getSliceDdosZoneTemplateIcmpV4FilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.IcmpTmplName = d.Get("icmp_tmpl_name").(string)
	ret.Inst.TypeList = getSliceDdosZoneTemplateIcmpV4TypeList(d.Get("type_list").([]interface{}))
	ret.Inst.TypeOther = getObjectDdosZoneTemplateIcmpV4TypeOther309(d.Get("type_other").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
