package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_other_ip_filtering_policy_oper_oper`: Operational Status for the object ip-filtering-policy-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperRead,

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
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "PortOther",
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

func resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper := setObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper(ret edpt.DataDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList(ret.DtDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList(d []edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper {
	var ret edpt.DdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceOtherIpFilteringPolicyOperOperOper(d.Get("oper").([]interface{}))

	ret.PortOther = d.Get("port_other").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
