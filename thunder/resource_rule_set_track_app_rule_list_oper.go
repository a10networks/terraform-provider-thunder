package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetTrackAppRuleListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_track_app_rule_list_oper`: Operational Status for the object track-app-rule-list\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetTrackAppRuleListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceRuleSetTrackAppRuleListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetTrackAppRuleListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetTrackAppRuleListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetTrackAppRuleListOperOper := setObjectRuleSetTrackAppRuleListOperOper(res)
		d.Set("oper", RuleSetTrackAppRuleListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetTrackAppRuleListOperOper(ret edpt.DataRuleSetTrackAppRuleListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rule_list": setSliceRuleSetTrackAppRuleListOperOperRuleList(ret.DtRuleSetTrackAppRuleListOper.Oper.RuleList),
		},
	}
}

func setSliceRuleSetTrackAppRuleListOperOperRuleList(d []edpt.RuleSetTrackAppRuleListOperOperRuleList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		result = append(result, in)
	}
	return result
}

func getObjectRuleSetTrackAppRuleListOperOper(d []interface{}) edpt.RuleSetTrackAppRuleListOperOper {

	count1 := len(d)
	var ret edpt.RuleSetTrackAppRuleListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleList = getSliceRuleSetTrackAppRuleListOperOperRuleList(in["rule_list"].([]interface{}))
	}
	return ret
}

func getSliceRuleSetTrackAppRuleListOperOperRuleList(d []interface{}) []edpt.RuleSetTrackAppRuleListOperOperRuleList {

	count1 := len(d)
	ret := make([]edpt.RuleSetTrackAppRuleListOperOperRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetTrackAppRuleListOperOperRuleList
		oi.Name = in["name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetTrackAppRuleListOper(d *schema.ResourceData) edpt.RuleSetTrackAppRuleListOper {
	var ret edpt.RuleSetTrackAppRuleListOper

	ret.Oper = getObjectRuleSetTrackAppRuleListOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
