package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_ip_proto_proto_tcp_udp_ip_filtering_policy_oper_oper`: Operational Status for the object ip-filtering-policy-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperRead,

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
								},
							},
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper := setObjectDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper(res)
		d.Set("oper", DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper(ret edpt.DataDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList(ret.DtDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList(d []edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper(d []interface{}) edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper {
	var ret edpt.DdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOper

	ret.Oper = getObjectDdosDstZoneIpProtoProtoTcpUdpIpFilteringPolicyOperOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
