package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_ip_filtering_policy_statistics_oper`: Operational Status for the object ip-filtering-policy-statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperRead,

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
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "PortOther",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper := setObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList(ret.DtDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList(d []edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		oi.Blacklisted_src_count = in["blacklisted_src_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyStatisticsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortOther = d.Get("port_other").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
