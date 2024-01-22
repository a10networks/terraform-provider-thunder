package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwRateLimitOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_rate_limit_oper`: Operational Status for the object rate-limit\n\n__PLACEHOLDER__",
		ReadContext: resourceFwRateLimitOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit_list": {
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
									"template_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cps_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cps_allowed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"uplink_traffic_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"uplink_traffic_allowed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"downlink_traffic_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"downlink_traffic_allowed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_traffic_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_traffic_allowed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"drop_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"v4_address": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"v4_netmask": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"v6_prefix": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"template_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"summary": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oper": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mem_reserved": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mem_used": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_failures": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_num_entries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_entries_scope_aggregate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_entries_scope_subscriber_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_entries_scope_subscriber_prefix": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_entries_scope_parent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_entries_scope_parent_subscriber_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_entries_scope_parent_subscriber_prefix": {
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

func resourceFwRateLimitOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimitOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwRateLimitOperOper := setObjectFwRateLimitOperOper(res)
		d.Set("oper", FwRateLimitOperOper)
		FwRateLimitOperSummary := setObjectFwRateLimitOperSummary(res)
		d.Set("summary", FwRateLimitOperSummary)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwRateLimitOperOper(ret edpt.DataFwRateLimitOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rate_limit_list": setSliceFwRateLimitOperOperRateLimitList(ret.DtFwRateLimitOper.Oper.RateLimitList),
			"v4_address":      ret.DtFwRateLimitOper.Oper.V4Address,
			"v4_netmask":      ret.DtFwRateLimitOper.Oper.V4Netmask,
			"v6_prefix":       ret.DtFwRateLimitOper.Oper.V6Prefix,
			"template_id":     ret.DtFwRateLimitOper.Oper.TemplateId,
		},
	}
}

func setSliceFwRateLimitOperOperRateLimitList(d []edpt.FwRateLimitOperOperRateLimitList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["prefix_len"] = item.PrefixLen
		in["rule_name"] = item.RuleName
		in["template_id"] = item.TemplateId
		in["type"] = item.Type
		in["cps_received"] = item.CpsReceived
		in["cps_allowed"] = item.CpsAllowed
		in["uplink_traffic_received"] = item.UplinkTrafficReceived
		in["uplink_traffic_allowed"] = item.UplinkTrafficAllowed
		in["downlink_traffic_received"] = item.DownlinkTrafficReceived
		in["downlink_traffic_allowed"] = item.DownlinkTrafficAllowed
		in["total_traffic_received"] = item.TotalTrafficReceived
		in["total_traffic_allowed"] = item.TotalTrafficAllowed
		in["drop_count"] = item.DropCount
		result = append(result, in)
	}
	return result
}

func setObjectFwRateLimitOperSummary(ret edpt.DataFwRateLimitOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"oper": setObjectFwRateLimitOperSummaryOper(ret.DtFwRateLimitOper.Summary.Oper),
		},
	}
}

func setObjectFwRateLimitOperSummaryOper(d edpt.FwRateLimitOperSummaryOper) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["mem_reserved"] = d.Mem_reserved

	in["mem_used"] = d.Mem_used

	in["alloc_failures"] = d.Alloc_failures

	in["total_num_entries"] = d.Total_num_entries

	in["total_entries_scope_aggregate"] = d.Total_entries_scope_aggregate

	in["total_entries_scope_subscriber_ip"] = d.Total_entries_scope_subscriber_ip

	in["total_entries_scope_subscriber_prefix"] = d.Total_entries_scope_subscriber_prefix

	in["total_entries_scope_parent"] = d.Total_entries_scope_parent

	in["total_entries_scope_parent_subscriber_ip"] = d.Total_entries_scope_parent_subscriber_ip

	in["total_entries_scope_parent_subscriber_prefix"] = d.Total_entries_scope_parent_subscriberPrefix
	result = append(result, in)
	return result
}

func getObjectFwRateLimitOperOper(d []interface{}) edpt.FwRateLimitOperOper {

	count1 := len(d)
	var ret edpt.FwRateLimitOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimitList = getSliceFwRateLimitOperOperRateLimitList(in["rate_limit_list"].([]interface{}))
		ret.V4Address = in["v4_address"].(string)
		ret.V4Netmask = in["v4_netmask"].(string)
		ret.V6Prefix = in["v6_prefix"].(string)
		ret.TemplateId = in["template_id"].(int)
	}
	return ret
}

func getSliceFwRateLimitOperOperRateLimitList(d []interface{}) []edpt.FwRateLimitOperOperRateLimitList {

	count1 := len(d)
	ret := make([]edpt.FwRateLimitOperOperRateLimitList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwRateLimitOperOperRateLimitList
		oi.Address = in["address"].(string)
		oi.PrefixLen = in["prefix_len"].(int)
		oi.RuleName = in["rule_name"].(string)
		oi.TemplateId = in["template_id"].(int)
		oi.Type = in["type"].(string)
		oi.CpsReceived = in["cps_received"].(int)
		oi.CpsAllowed = in["cps_allowed"].(int)
		oi.UplinkTrafficReceived = in["uplink_traffic_received"].(int)
		oi.UplinkTrafficAllowed = in["uplink_traffic_allowed"].(int)
		oi.DownlinkTrafficReceived = in["downlink_traffic_received"].(int)
		oi.DownlinkTrafficAllowed = in["downlink_traffic_allowed"].(int)
		oi.TotalTrafficReceived = in["total_traffic_received"].(int)
		oi.TotalTrafficAllowed = in["total_traffic_allowed"].(int)
		oi.DropCount = in["drop_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFwRateLimitOperSummary(d []interface{}) edpt.FwRateLimitOperSummary {

	count1 := len(d)
	var ret edpt.FwRateLimitOperSummary
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Oper = getObjectFwRateLimitOperSummaryOper(in["oper"].([]interface{}))
	}
	return ret
}

func getObjectFwRateLimitOperSummaryOper(d []interface{}) edpt.FwRateLimitOperSummaryOper {

	count1 := len(d)
	var ret edpt.FwRateLimitOperSummaryOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mem_reserved = in["mem_reserved"].(int)
		ret.Mem_used = in["mem_used"].(int)
		ret.Alloc_failures = in["alloc_failures"].(int)
		ret.Total_num_entries = in["total_num_entries"].(int)
		ret.Total_entries_scope_aggregate = in["total_entries_scope_aggregate"].(int)
		ret.Total_entries_scope_subscriber_ip = in["total_entries_scope_subscriber_ip"].(int)
		ret.Total_entries_scope_subscriber_prefix = in["total_entries_scope_subscriber_prefix"].(int)
		ret.Total_entries_scope_parent = in["total_entries_scope_parent"].(int)
		ret.Total_entries_scope_parent_subscriber_ip = in["total_entries_scope_parent_subscriber_ip"].(int)
		ret.Total_entries_scope_parent_subscriberPrefix = in["total_entries_scope_parent_subscriber_prefix"].(int)
	}
	return ret
}

func dataToEndpointFwRateLimitOper(d *schema.ResourceData) edpt.FwRateLimitOper {
	var ret edpt.FwRateLimitOper

	ret.Oper = getObjectFwRateLimitOperOper(d.Get("oper").([]interface{}))

	ret.Summary = getObjectFwRateLimitOperSummary(d.Get("summary").([]interface{}))
	return ret
}
