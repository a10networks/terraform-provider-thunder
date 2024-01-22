package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryL4TypeIpFilteringPolicyOperOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_entry_l4_type_ip_filtering_policy_oper_oper`: Operational Status for the object ip-filtering-policy-oper\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstEntryL4TypeIpFilteringPolicyOperOperRead,

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
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}

func resourceDdosDstEntryL4TypeIpFilteringPolicyOperOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryL4TypeIpFilteringPolicyOperOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryL4TypeIpFilteringPolicyOperOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstEntryL4TypeIpFilteringPolicyOperOperOper := setObjectDdosDstEntryL4TypeIpFilteringPolicyOperOperOper(res)
		d.Set("oper", DdosDstEntryL4TypeIpFilteringPolicyOperOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstEntryL4TypeIpFilteringPolicyOperOperOper(ret edpt.DataDdosDstEntryL4TypeIpFilteringPolicyOperOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceDdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList(ret.DtDdosDstEntryL4TypeIpFilteringPolicyOperOper.Oper.RuleList),
		},
	}
}

func setSliceDdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList(d []edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["seq"] = item.Seq
		in["hits"] = item.Hits
		result = append(result, in)
	}
	return result
}

func getObjectDdosDstEntryL4TypeIpFilteringPolicyOperOperOper(d []interface{}) edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOperOper {

	count1 := len(d)
	var ret edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceDdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList(d []interface{}) []edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOperOperRuleList
		oi.Seq = in["seq"].(int)
		oi.Hits = in["hits"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstEntryL4TypeIpFilteringPolicyOperOper(d *schema.ResourceData) edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOper {
	var ret edpt.DdosDstEntryL4TypeIpFilteringPolicyOperOper

	ret.Oper = getObjectDdosDstEntryL4TypeIpFilteringPolicyOperOperOper(d.Get("oper").([]interface{}))

	ret.Protocol = d.Get("protocol").(string)

	ret.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
