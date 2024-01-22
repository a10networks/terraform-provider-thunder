package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetApplicationOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_application_oper`: Operational Status for the object application\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetApplicationOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category_stat": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"app_stat": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"protocol": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rule": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"rule_set_only": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"stat_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"category": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"type": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"conns": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"bytes": {
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceRuleSetApplicationOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetApplicationOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetApplicationOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetApplicationOperOper := setObjectRuleSetApplicationOperOper(res)
		d.Set("oper", RuleSetApplicationOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetApplicationOperOper(ret edpt.DataRuleSetApplicationOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"category_stat": ret.DtRuleSetApplicationOper.Oper.CategoryStat,
			"app_stat":      ret.DtRuleSetApplicationOper.Oper.AppStat,
			"protocol":      ret.DtRuleSetApplicationOper.Oper.Protocol,
			"rule":          ret.DtRuleSetApplicationOper.Oper.Rule,
			"rule_set_only": ret.DtRuleSetApplicationOper.Oper.RuleSetOnly,
			"rule_list":     setSliceRuleSetApplicationOperOperRuleList(ret.DtRuleSetApplicationOper.Oper.RuleList),
		},
	}
}

func setSliceRuleSetApplicationOperOperRuleList(d []edpt.RuleSetApplicationOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stat_list"] = setSliceRuleSetApplicationOperOperRuleListStatList(item.StatList)
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetApplicationOperOperRuleListStatList(d []edpt.RuleSetApplicationOperOperRuleListStatList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["category"] = item.Category
		in["type"] = item.Type
		in["conns"] = item.Conns
		in["bytes"] = item.Bytes
		result = append(result, in)
	}
	return result
}

func getObjectRuleSetApplicationOperOper(d []interface{}) edpt.RuleSetApplicationOperOper {

	count1 := len(d)
	var ret edpt.RuleSetApplicationOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CategoryStat = in["category_stat"].(string)
		ret.AppStat = in["app_stat"].(string)
		ret.Protocol = in["protocol"].(int)
		ret.Rule = in["rule"].(string)
		ret.RuleSetOnly = in["rule_set_only"].(int)
		ret.RuleList = getSliceRuleSetApplicationOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceRuleSetApplicationOperOperRuleList(d []interface{}) []edpt.RuleSetApplicationOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetApplicationOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetApplicationOperOperRuleList
		oi.Name = in["name"].(string)
		oi.StatList = getSliceRuleSetApplicationOperOperRuleListStatList(in["stat_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetApplicationOperOperRuleListStatList(d []interface{}) []edpt.RuleSetApplicationOperOperRuleListStatList {

	count1 := len(d)
	ret := make([]edpt.RuleSetApplicationOperOperRuleListStatList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetApplicationOperOperRuleListStatList
		oi.Name = in["name"].(string)
		oi.Category = in["category"].(string)
		oi.Type = in["type"].(string)
		oi.Conns = in["conns"].(int)
		oi.Bytes = in["bytes"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetApplicationOper(d *schema.ResourceData) edpt.RuleSetApplicationOper {
	var ret edpt.RuleSetApplicationOper

	ret.Oper = getObjectRuleSetApplicationOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
