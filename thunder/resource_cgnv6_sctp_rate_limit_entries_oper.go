package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SctpRateLimitEntriesOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_sctp_rate_limit_entries_oper`: Operational Status for the object rate-limit-entries\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6SctpRateLimitEntriesOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit_entries_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"direction": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rate_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6SctpRateLimitEntriesOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitEntriesOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitEntriesOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6SctpRateLimitEntriesOperOper := setObjectCgnv6SctpRateLimitEntriesOperOper(res)
		d.Set("oper", Cgnv6SctpRateLimitEntriesOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6SctpRateLimitEntriesOperOper(ret edpt.DataCgnv6SctpRateLimitEntriesOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rate_limit_entries_list": setSliceCgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList(ret.DtCgnv6SctpRateLimitEntriesOper.Oper.RateLimitEntriesList),
			"entry_count":             ret.DtCgnv6SctpRateLimitEntriesOper.Oper.EntryCount,
		},
	}
}

func setSliceCgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList(d []edpt.Cgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["address"] = item.Address
		in["direction"] = item.Direction
		in["pps"] = item.Pps
		in["rate_limit"] = item.RateLimit
		result = append(result, in)
	}
	return result
}

func getObjectCgnv6SctpRateLimitEntriesOperOper(d []interface{}) edpt.Cgnv6SctpRateLimitEntriesOperOper {

	count1 := len(d)
	var ret edpt.Cgnv6SctpRateLimitEntriesOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimitEntriesList = getSliceCgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList(in["rate_limit_entries_list"].([]interface{}))
		ret.EntryCount = in["entry_count"].(int)
	}
	return ret
}

func getSliceCgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList(d []interface{}) []edpt.Cgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6SctpRateLimitEntriesOperOperRateLimitEntriesList
		oi.Address = in["address"].(string)
		oi.Direction = in["direction"].(string)
		oi.Pps = in["pps"].(int)
		oi.RateLimit = in["rate_limit"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6SctpRateLimitEntriesOper(d *schema.ResourceData) edpt.Cgnv6SctpRateLimitEntriesOper {
	var ret edpt.Cgnv6SctpRateLimitEntriesOper

	ret.Oper = getObjectCgnv6SctpRateLimitEntriesOperOper(d.Get("oper").([]interface{}))
	return ret
}
