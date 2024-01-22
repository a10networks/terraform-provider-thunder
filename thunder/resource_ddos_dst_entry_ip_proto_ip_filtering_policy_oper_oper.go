package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryIpProtoIpFilteringPolicyOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_ip_proto_ip_filtering_policy_oper_oper`: Operational Status for the object ip-filtering-policy-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryIpProtoIpFilteringPolicyOperOperRead,

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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryIpProtoIpFilteringPolicyOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryIpProtoIpFilteringPolicyOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryIpProtoIpFilteringPolicyOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryIpProtoIpFilteringPolicyOperOperOper := setObjectDdosDstEntryIpProtoIpFilteringPolicyOperOperOper(res)
		d.Set("oper", DdosDstEntryIpProtoIpFilteringPolicyOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryIpProtoIpFilteringPolicyOperOperOper(ret edpt.DataDdosDstEntryIpProtoIpFilteringPolicyOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList(ret.DtDdosDstEntryIpProtoIpFilteringPolicyOperOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList(d []edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstEntryIpProtoIpFilteringPolicyOperOperOper(d []interface{}) edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList(d []interface{}) []edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryIpProtoIpFilteringPolicyOperOper(d *schema.ResourceData) edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOper {
	var ret edpt.DdosDstEntryIpProtoIpFilteringPolicyOperOper

	ret.Oper = getObjectDdosDstEntryIpProtoIpFilteringPolicyOperOperOper(d.Get("oper").([]interface{}))

	ret.PortNum = d.Get("port_num").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
