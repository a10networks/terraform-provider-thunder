package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortRangeIpFilteringPolicyOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_port_range_ip_filtering_policy_oper_oper`: Operational Status for the object ip-filtering-policy-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZonePortRangeIpFilteringPolicyOperOperRead,

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
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_range_end": {
				Type: schema.TypeString, Required: true, Description: "PortRangeEnd",
			},
			"port_range_start": {
				Type: schema.TypeString, Required: true, Description: "PortRangeStart",
			},
		},
	}
}

func resourceDdosDstZonePortRangeIpFilteringPolicyOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortRangeIpFilteringPolicyOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortRangeIpFilteringPolicyOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZonePortRangeIpFilteringPolicyOperOperOper := setObjectDdosDstZonePortRangeIpFilteringPolicyOperOperOper(res)
		d.Set("oper", DdosDstZonePortRangeIpFilteringPolicyOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZonePortRangeIpFilteringPolicyOperOperOper(ret edpt.DataDdosDstZonePortRangeIpFilteringPolicyOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList(ret.DtDdosDstZonePortRangeIpFilteringPolicyOperOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList(d []edpt.DdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstZonePortRangeIpFilteringPolicyOperOperOper(d []interface{}) edpt.DdosDstZonePortRangeIpFilteringPolicyOperOperOper {

	count1 := len(d)
	var ret edpt.DdosDstZonePortRangeIpFilteringPolicyOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList(d []interface{}) []edpt.DdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortRangeIpFilteringPolicyOperOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZonePortRangeIpFilteringPolicyOperOper(d *schema.ResourceData) edpt.DdosDstZonePortRangeIpFilteringPolicyOperOper {
	var ret edpt.DdosDstZonePortRangeIpFilteringPolicyOperOper

	ret.Oper = getObjectDdosDstZonePortRangeIpFilteringPolicyOperOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.ZoneName = d.Get("zone_name").(string)

	ret.PortRangeEnd = d.Get("port_range_end").(string)

	ret.PortRangeStart = d.Get("port_range_start").(string)
	return ret
}
