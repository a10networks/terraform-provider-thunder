package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_number_ip_filtering_policy_statistics_oper`: Operational Status for the object ip-filtering-policy-statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"blacklisted_src_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper := setObjectDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList(ret.DtDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList(d []edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		in["blacklisted_src_count"] = item.Blacklisted_src_count
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		oi.Blacklisted_src_count = in["blacklisted_src_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper {
	var ret edpt.DdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNumberIpFilteringPolicyStatisticsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
