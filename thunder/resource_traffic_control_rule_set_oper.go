package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_traffic_control_rule_set_oper`: Operational Status for the object rule-set\n\n__PLACEHOLDER__",
		ReadContext: resourceTrafficControlRuleSetOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule set name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"policy_rule_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rule_stats": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_hitcount": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
					},
				},
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Rule name",
						},
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hitcount": {
										Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceTrafficControlRuleSetOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TrafficControlRuleSetOperOper := setObjectTrafficControlRuleSetOperOper(res)
		d.Set("oper", TrafficControlRuleSetOperOper)
		TrafficControlRuleSetOperRuleList := setSliceTrafficControlRuleSetOperRuleList(res)
		d.Set("rule_list", TrafficControlRuleSetOperRuleList)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTrafficControlRuleSetOperOper(ret edpt.DataTrafficControlRuleSetOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"policy_status":     ret.DtTrafficControlRuleSetOper.Oper.PolicyStatus,
			"policy_rule_count": ret.DtTrafficControlRuleSetOper.Oper.PolicyRuleCount,
			"rule_stats":        setSliceTrafficControlRuleSetOperOperRuleStats(ret.DtTrafficControlRuleSetOper.Oper.RuleStats),
		},
	}
}

func setSliceTrafficControlRuleSetOperOperRuleStats(d []edpt.TrafficControlRuleSetOperOperRuleStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rule_name"] = item.RuleName
		in["rule_status"] = item.RuleStatus
		in["rule_hitcount"] = item.RuleHitcount
		result = append(result, in)
	}
	return result
}

func setSliceTrafficControlRuleSetOperRuleList(d edpt.DataTrafficControlRuleSetOper) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtTrafficControlRuleSetOper.RuleList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["oper"] = setObjectTrafficControlRuleSetOperRuleListOper(item.Oper)
		result = append(result, in)
	}
	return result
}

func setObjectTrafficControlRuleSetOperRuleListOper(d edpt.TrafficControlRuleSetOperRuleListOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["status"] = d.Status

	in["hitcount"] = d.Hitcount
	result = append(result, in)
	return result
}

func getObjectTrafficControlRuleSetOperOper(d []interface{}) edpt.TrafficControlRuleSetOperOper {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PolicyStatus = in["policy_status"].(string)
		ret.PolicyRuleCount = in["policy_rule_count"].(int)
		ret.RuleStats = getSliceTrafficControlRuleSetOperOperRuleStats(in["rule_stats"].([]interface{}))
	}
	return ret
}

func getSliceTrafficControlRuleSetOperOperRuleStats(d []interface{}) []edpt.TrafficControlRuleSetOperOperRuleStats {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetOperOperRuleStats, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetOperOperRuleStats
		oi.RuleName = in["rule_name"].(string)
		oi.RuleStatus = in["rule_status"].(string)
		oi.RuleHitcount = in["rule_hitcount"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTrafficControlRuleSetOperRuleList(d []interface{}) []edpt.TrafficControlRuleSetOperRuleList {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetOperRuleList
		oi.Name = in["name"].(string)
		oi.Oper = getObjectTrafficControlRuleSetOperRuleListOper(in["oper"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectTrafficControlRuleSetOperRuleListOper(d []interface{}) edpt.TrafficControlRuleSetOperRuleListOper {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetOperRuleListOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Status = in["status"].(string)
		ret.Hitcount = in["hitcount"].(int)
	}
	return ret
}

func dataToEndpointTrafficControlRuleSetOper(d *schema.ResourceData) edpt.TrafficControlRuleSetOper {
	var ret edpt.TrafficControlRuleSetOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectTrafficControlRuleSetOperOper(d.Get("oper").([]interface{}))

	ret.RuleList = getSliceTrafficControlRuleSetOperRuleList(d.Get("rule_list").([]interface{}))
	return ret
}
