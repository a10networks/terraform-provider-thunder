package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwTopKRulesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_top_k_rules_oper`: Operational Status for the object top-k-rules\n\n__PLACEHOLDER__",
		ReadContext: resourceFwTopKRulesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_topk_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"topk_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rule_name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"metric_value": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceFwTopKRulesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwTopKRulesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwTopKRulesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwTopKRulesOperOper := setObjectFwTopKRulesOperOper(res)
		d.Set("oper", FwTopKRulesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwTopKRulesOperOper(ret edpt.DataFwTopKRulesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"metric_topk_list": setSliceFwTopKRulesOperOperMetricTopkList(ret.DtFwTopKRulesOper.Oper.MetricTopkList),
		},
	}
}

func setSliceFwTopKRulesOperOperMetricTopkList(d []edpt.FwTopKRulesOperOperMetricTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["metric_name"] = item.MetricName
		in["topk_list"] = setSliceFwTopKRulesOperOperMetricTopkListTopkList(item.TopkList)
		result = append(result, in)
	}
	return result
}

func setSliceFwTopKRulesOperOperMetricTopkListTopkList(d []edpt.FwTopKRulesOperOperMetricTopkListTopkList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rule_name"] = item.RuleName
		in["metric_value"] = item.MetricValue
		result = append(result, in)
	}
	return result
}

func getObjectFwTopKRulesOperOper(d []interface{}) edpt.FwTopKRulesOperOper {

	count1 := len(d)
	var ret edpt.FwTopKRulesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MetricTopkList = getSliceFwTopKRulesOperOperMetricTopkList(in["metric_topk_list"].([]interface{}))
	}
	return ret
}

func getSliceFwTopKRulesOperOperMetricTopkList(d []interface{}) []edpt.FwTopKRulesOperOperMetricTopkList {

	count1 := len(d)
	ret := make([]edpt.FwTopKRulesOperOperMetricTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTopKRulesOperOperMetricTopkList
		oi.MetricName = in["metric_name"].(string)
		oi.TopkList = getSliceFwTopKRulesOperOperMetricTopkListTopkList(in["topk_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwTopKRulesOperOperMetricTopkListTopkList(d []interface{}) []edpt.FwTopKRulesOperOperMetricTopkListTopkList {

	count1 := len(d)
	ret := make([]edpt.FwTopKRulesOperOperMetricTopkListTopkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwTopKRulesOperOperMetricTopkListTopkList
		oi.RuleName = in["rule_name"].(string)
		oi.MetricValue = in["metric_value"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwTopKRulesOper(d *schema.ResourceData) edpt.FwTopKRulesOper {
	var ret edpt.FwTopKRulesOper

	ret.Oper = getObjectFwTopKRulesOperOper(d.Get("oper").([]interface{}))
	return ret
}
