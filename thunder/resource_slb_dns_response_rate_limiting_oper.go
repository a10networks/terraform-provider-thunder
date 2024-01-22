package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsResponseRateLimitingOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dns_response_rate_limiting_oper`: Operational Status for the object dns-response-rate-limiting\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDnsResponseRateLimitingOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter_address_v4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter_address_v6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter_fqdn": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"filter_debug": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entry_fqdn": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entry_hit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"entry_age": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"entry_response_credit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"entry_action": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entry_over_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"dnsrrl_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total_created": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_inserted": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_withdrew": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_ready_to_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_logs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_overflow_entry_hits": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_refill": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_credit_exceeded": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"other_thread_refill": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_entry_create_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_entry_create_oom": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_entry_ext_create_oom": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_entry_insert_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_vport_fail_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbDnsResponseRateLimitingOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsResponseRateLimitingOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsResponseRateLimitingOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDnsResponseRateLimitingOperOper := setObjectSlbDnsResponseRateLimitingOperOper(res)
		d.Set("oper", SlbDnsResponseRateLimitingOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDnsResponseRateLimitingOperOper(ret edpt.DataSlbDnsResponseRateLimitingOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter_type":       ret.DtSlbDnsResponseRateLimitingOper.Oper.Filter_type,
			"filter_address_v4": ret.DtSlbDnsResponseRateLimitingOper.Oper.Filter_address_v4,
			"filter_address_v6": ret.DtSlbDnsResponseRateLimitingOper.Oper.Filter_address_v6,
			"filter_fqdn":       ret.DtSlbDnsResponseRateLimitingOper.Oper.Filter_fqdn,
			"filter_debug":      ret.DtSlbDnsResponseRateLimitingOper.Oper.Filter_debug,
			"entry_list":        setSliceSlbDnsResponseRateLimitingOperOperEntryList(ret.DtSlbDnsResponseRateLimitingOper.Oper.EntryList),
			"dnsrrl_cpu_list":   setSliceSlbDnsResponseRateLimitingOperOperDnsrrlCpuList(ret.DtSlbDnsResponseRateLimitingOper.Oper.DnsrrlCpuList),
			"cpu_count":         ret.DtSlbDnsResponseRateLimitingOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbDnsResponseRateLimitingOperOperEntryList(d []edpt.SlbDnsResponseRateLimitingOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entry_address"] = item.EntryAddress
		in["entry_fqdn"] = item.EntryFqdn
		in["entry_hit_count"] = item.EntryHitCount
		in["entry_age"] = item.EntryAge
		in["entry_response_credit"] = item.EntryResponseCredit
		in["entry_action"] = item.EntryAction
		in["entry_over_limit"] = item.EntryOverLimit
		result = append(result, in)
	}
	return result
}

func setSliceSlbDnsResponseRateLimitingOperOperDnsrrlCpuList(d []edpt.SlbDnsResponseRateLimitingOperOperDnsrrlCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["total_created"] = item.Total_created
		in["total_inserted"] = item.Total_inserted
		in["total_withdrew"] = item.Total_withdrew
		in["total_ready_to_free"] = item.Total_ready_to_free
		in["total_freed"] = item.Total_freed
		in["total_logs"] = item.Total_logs
		in["total_overflow_entry_hits"] = item.Total_overflow_entry_hits
		in["total_refill"] = item.Total_refill
		in["total_credit_exceeded"] = item.Total_credit_exceeded
		in["other_thread_refill"] = item.Other_thread_refill
		in["err_entry_create_failed"] = item.Err_entry_create_failed
		in["err_entry_create_oom"] = item.Err_entry_create_oom
		in["err_entry_ext_create_oom"] = item.Err_entry_ext_create_oom
		in["err_entry_insert_failed"] = item.Err_entry_insert_failed
		in["err_vport_fail_match"] = item.Err_vport_fail_match
		result = append(result, in)
	}
	return result
}

func getObjectSlbDnsResponseRateLimitingOperOper(d []interface{}) edpt.SlbDnsResponseRateLimitingOperOper {

	count1 := len(d)
	var ret edpt.SlbDnsResponseRateLimitingOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Filter_type = in["filter_type"].(string)
		ret.Filter_address_v4 = in["filter_address_v4"].(string)
		ret.Filter_address_v6 = in["filter_address_v6"].(string)
		ret.Filter_fqdn = in["filter_fqdn"].(string)
		ret.Filter_debug = in["filter_debug"].(int)
		ret.EntryList = getSliceSlbDnsResponseRateLimitingOperOperEntryList(in["entry_list"].([]interface{}))
		ret.DnsrrlCpuList = getSliceSlbDnsResponseRateLimitingOperOperDnsrrlCpuList(in["dnsrrl_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbDnsResponseRateLimitingOperOperEntryList(d []interface{}) []edpt.SlbDnsResponseRateLimitingOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.SlbDnsResponseRateLimitingOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsResponseRateLimitingOperOperEntryList
		oi.EntryAddress = in["entry_address"].(string)
		oi.EntryFqdn = in["entry_fqdn"].(string)
		oi.EntryHitCount = in["entry_hit_count"].(int)
		oi.EntryAge = in["entry_age"].(int)
		oi.EntryResponseCredit = in["entry_response_credit"].(int)
		oi.EntryAction = in["entry_action"].(string)
		oi.EntryOverLimit = in["entry_over_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbDnsResponseRateLimitingOperOperDnsrrlCpuList(d []interface{}) []edpt.SlbDnsResponseRateLimitingOperOperDnsrrlCpuList {

	count1 := len(d)
	ret := make([]edpt.SlbDnsResponseRateLimitingOperOperDnsrrlCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbDnsResponseRateLimitingOperOperDnsrrlCpuList
		oi.Total_created = in["total_created"].(int)
		oi.Total_inserted = in["total_inserted"].(int)
		oi.Total_withdrew = in["total_withdrew"].(int)
		oi.Total_ready_to_free = in["total_ready_to_free"].(int)
		oi.Total_freed = in["total_freed"].(int)
		oi.Total_logs = in["total_logs"].(int)
		oi.Total_overflow_entry_hits = in["total_overflow_entry_hits"].(int)
		oi.Total_refill = in["total_refill"].(int)
		oi.Total_credit_exceeded = in["total_credit_exceeded"].(int)
		oi.Other_thread_refill = in["other_thread_refill"].(int)
		oi.Err_entry_create_failed = in["err_entry_create_failed"].(int)
		oi.Err_entry_create_oom = in["err_entry_create_oom"].(int)
		oi.Err_entry_ext_create_oom = in["err_entry_ext_create_oom"].(int)
		oi.Err_entry_insert_failed = in["err_entry_insert_failed"].(int)
		oi.Err_vport_fail_match = in["err_vport_fail_match"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbDnsResponseRateLimitingOper(d *schema.ResourceData) edpt.SlbDnsResponseRateLimitingOper {
	var ret edpt.SlbDnsResponseRateLimitingOper

	ret.Oper = getObjectSlbDnsResponseRateLimitingOperOper(d.Get("oper").([]interface{}))
	return ret
}
