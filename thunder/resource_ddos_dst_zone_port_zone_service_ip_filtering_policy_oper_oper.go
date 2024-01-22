package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceIpFilteringPolicyOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_zone_service_ip_filtering_policy_oper_oper`: Operational Status for the object ip-filtering-policy-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperRead,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}

func resourceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceIpFilteringPolicyOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper := setObjectDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper(res)
		d.Set("oper", DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper(ret edpt.DataDdosDstZonePortZoneServiceIpFilteringPolicyOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList(ret.DtDdosDstZonePortZoneServiceIpFilteringPolicyOperOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList(d []edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper(d []interface{}) edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceIpFilteringPolicyOperOper(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOper {
	var ret edpt.DdosDstZonePortZoneServiceIpFilteringPolicyOperOper

	ret.Oper = getObjectDdosDstZonePortZoneServiceIpFilteringPolicyOperOperOper(d.Get("oper").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortNum = d.Get("port_num").(string)

	ret.Protocol = d.Get("protocol").(string)
	return ret
}
