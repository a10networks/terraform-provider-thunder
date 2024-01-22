package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLimitEntryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_limit_entry_oper`: Operational Status for the object limit-entry\n\n__PLACEHOLDER__",
		ReadContext: resourceFwLimitEntryOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"prefix_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rule_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"curr_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"limit_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"prefix6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"prefix4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"prefix_len6": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"prefix_len4": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceFwLimitEntryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLimitEntryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLimitEntryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwLimitEntryOperOper := setObjectFwLimitEntryOperOper(res)
		d.Set("oper", FwLimitEntryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwLimitEntryOperOper(ret edpt.DataFwLimitEntryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"limit_entry_list":  setSliceFwLimitEntryOperOperLimitEntryList(ret.DtFwLimitEntryOper.Oper.LimitEntryList),
			"limit_entry_count": ret.DtFwLimitEntryOper.Oper.LimitEntryCount,
			"prefix6":           ret.DtFwLimitEntryOper.Oper.Prefix6,
			"prefix4":           ret.DtFwLimitEntryOper.Oper.Prefix4,
			"prefix_len6":       ret.DtFwLimitEntryOper.Oper.PrefixLen6,
			"prefix_len4":       ret.DtFwLimitEntryOper.Oper.PrefixLen4,
		},
	}
}

func setSliceFwLimitEntryOperOperLimitEntryList(d []edpt.FwLimitEntryOperOperLimitEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["prefix_len"] = item.PrefixLen
		in["rule_name"] = item.RuleName
		in["curr_count"] = item.CurrCount
		in["max_count"] = item.MaxCount
		in["type"] = item.Type
		result = append(result, in)
	}
	return result
}

func getObjectFwLimitEntryOperOper(d []interface{}) edpt.FwLimitEntryOperOper {

	count1 := len(d)
	var ret edpt.FwLimitEntryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LimitEntryList = getSliceFwLimitEntryOperOperLimitEntryList(in["limit_entry_list"].([]interface{}))
		ret.LimitEntryCount = in["limit_entry_count"].(int)
		ret.Prefix6 = in["prefix6"].(string)
		ret.Prefix4 = in["prefix4"].(string)
		ret.PrefixLen6 = in["prefix_len6"].(int)
		ret.PrefixLen4 = in["prefix_len4"].(int)
	}
	return ret
}

func getSliceFwLimitEntryOperOperLimitEntryList(d []interface{}) []edpt.FwLimitEntryOperOperLimitEntryList {

	count1 := len(d)
	ret := make([]edpt.FwLimitEntryOperOperLimitEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwLimitEntryOperOperLimitEntryList
		oi.Address = in["address"].(string)
		oi.PrefixLen = in["prefix_len"].(int)
		oi.RuleName = in["rule_name"].(string)
		oi.CurrCount = in["curr_count"].(int)
		oi.MaxCount = in["max_count"].(int)
		oi.Type = in["type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwLimitEntryOper(d *schema.ResourceData) edpt.FwLimitEntryOper {
	var ret edpt.FwLimitEntryOper

	ret.Oper = getObjectFwLimitEntryOperOper(d.Get("oper").([]interface{}))
	return ret
}
