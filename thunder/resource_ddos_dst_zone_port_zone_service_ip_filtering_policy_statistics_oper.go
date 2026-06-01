package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_ip_filtering_policy_statistics_oper`: Operational Status for the object ip-filtering-policy-statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperRead,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper := setObjectDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper(ret edpt.DataDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList(ret.DtDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList(d []edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList) []map[string]interface{} {
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

func getObjectDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		oi.Blacklisted_src_count = in["blacklisted_src_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper {
	var ret edpt.DdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceIpFilteringPolicyStatisticsOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
