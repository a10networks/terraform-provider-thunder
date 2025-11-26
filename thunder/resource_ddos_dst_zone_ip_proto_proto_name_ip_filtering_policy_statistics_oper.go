package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_name_ip_filtering_policy_statistics_oper`: Operational Status for the object ip-filtering-policy-statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper := setObjectDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper(ret edpt.DataDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList(ret.DtDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList(d []edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList) []map[string]interface{} {
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

func getObjectDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		oi.Blacklisted_src_count = in["blacklisted_src_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper {
	var ret edpt.DdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoNameIpFilteringPolicyStatisticsOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
