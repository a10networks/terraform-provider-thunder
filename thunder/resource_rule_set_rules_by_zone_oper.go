package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRulesByZoneOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_rules_by_zone_oper`: Operational Status for the object rules-by-zone\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetRulesByZoneOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"to": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"rule_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"source_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"source": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"dest_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dest": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"service_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"service": {
																Type: schema.TypeString, Optional: true, Description: "",
															},
														},
													},
												},
												"dscp_list": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dscp": {
																Type: schema.TypeString, Optional: true, Description: "",
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
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceRuleSetRulesByZoneOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRulesByZoneOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRulesByZoneOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetRulesByZoneOperOper := setObjectRuleSetRulesByZoneOperOper(res)
		d.Set("oper", RuleSetRulesByZoneOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetRulesByZoneOperOper(ret edpt.DataRuleSetRulesByZoneOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"group_list": setSliceRuleSetRulesByZoneOperOperGroupList(ret.DtRuleSetRulesByZoneOper.Oper.GroupList),
		},
	}
}

func setSliceRuleSetRulesByZoneOperOperGroupList(d []edpt.RuleSetRulesByZoneOperOperGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["from"] = item.From
		in["to"] = item.To
		in["rule_list"] = setSliceRuleSetRulesByZoneOperOperGroupListRuleList(item.RuleList)
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetRulesByZoneOperOperGroupListRuleList(d []edpt.RuleSetRulesByZoneOperOperGroupListRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["action"] = item.Action
		in["source_list"] = setSliceRuleSetRulesByZoneOperOperGroupListRuleListSourceList(item.SourceList)
		in["dest_list"] = setSliceRuleSetRulesByZoneOperOperGroupListRuleListDestList(item.DestList)
		in["service_list"] = setSliceRuleSetRulesByZoneOperOperGroupListRuleListServiceList(item.ServiceList)
		in["dscp_list"] = setSliceRuleSetRulesByZoneOperOperGroupListRuleListDscpList(item.DscpList)
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetRulesByZoneOperOperGroupListRuleListSourceList(d []edpt.RuleSetRulesByZoneOperOperGroupListRuleListSourceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["source"] = item.Source
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetRulesByZoneOperOperGroupListRuleListDestList(d []edpt.RuleSetRulesByZoneOperOperGroupListRuleListDestList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dest"] = item.Dest
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetRulesByZoneOperOperGroupListRuleListServiceList(d []edpt.RuleSetRulesByZoneOperOperGroupListRuleListServiceList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["service"] = item.Service
		result = append(result, in)
	}
	return result
}

func setSliceRuleSetRulesByZoneOperOperGroupListRuleListDscpList(d []edpt.RuleSetRulesByZoneOperOperGroupListRuleListDscpList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["dscp"] = item.Dscp
		result = append(result, in)
	}
	return result
}

func getObjectRuleSetRulesByZoneOperOper(d []interface{}) edpt.RuleSetRulesByZoneOperOper {

	count1 := len(d)
	var ret edpt.RuleSetRulesByZoneOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GroupList = getSliceRuleSetRulesByZoneOperOperGroupList(in["group_list"].([]interface{}))
	}
	return ret
}

func getSliceRuleSetRulesByZoneOperOperGroupList(d []interface{}) []edpt.RuleSetRulesByZoneOperOperGroupList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneOperOperGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneOperOperGroupList
		oi.From = in["from"].(string)
		oi.To = in["to"].(string)
		oi.RuleList = getSliceRuleSetRulesByZoneOperOperGroupListRuleList(in["rule_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRulesByZoneOperOperGroupListRuleList(d []interface{}) []edpt.RuleSetRulesByZoneOperOperGroupListRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneOperOperGroupListRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneOperOperGroupListRuleList
		oi.Name = in["name"].(string)
		oi.Action = in["action"].(string)
		oi.SourceList = getSliceRuleSetRulesByZoneOperOperGroupListRuleListSourceList(in["source_list"].([]interface{}))
		oi.DestList = getSliceRuleSetRulesByZoneOperOperGroupListRuleListDestList(in["dest_list"].([]interface{}))
		oi.ServiceList = getSliceRuleSetRulesByZoneOperOperGroupListRuleListServiceList(in["service_list"].([]interface{}))
		oi.DscpList = getSliceRuleSetRulesByZoneOperOperGroupListRuleListDscpList(in["dscp_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRulesByZoneOperOperGroupListRuleListSourceList(d []interface{}) []edpt.RuleSetRulesByZoneOperOperGroupListRuleListSourceList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneOperOperGroupListRuleListSourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneOperOperGroupListRuleListSourceList
		oi.Source = in["source"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRulesByZoneOperOperGroupListRuleListDestList(d []interface{}) []edpt.RuleSetRulesByZoneOperOperGroupListRuleListDestList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneOperOperGroupListRuleListDestList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneOperOperGroupListRuleListDestList
		oi.Dest = in["dest"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRulesByZoneOperOperGroupListRuleListServiceList(d []interface{}) []edpt.RuleSetRulesByZoneOperOperGroupListRuleListServiceList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneOperOperGroupListRuleListServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneOperOperGroupListRuleListServiceList
		oi.Service = in["service"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRulesByZoneOperOperGroupListRuleListDscpList(d []interface{}) []edpt.RuleSetRulesByZoneOperOperGroupListRuleListDscpList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneOperOperGroupListRuleListDscpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneOperOperGroupListRuleListDscpList
		oi.Dscp = in["dscp"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetRulesByZoneOper(d *schema.ResourceData) edpt.RuleSetRulesByZoneOper {
	var ret edpt.RuleSetRulesByZoneOper

	ret.Oper = getObjectRuleSetRulesByZoneOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
