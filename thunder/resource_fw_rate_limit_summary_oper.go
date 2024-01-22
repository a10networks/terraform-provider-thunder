package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwRateLimitSummaryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_rate_limit_summary_oper`: Operational Status for the object summary\n\n__PLACEHOLDER__",
		ReadContext: resourceFwRateLimitSummaryOperRead,

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
	}
}

func resourceFwRateLimitSummaryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitSummaryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimitSummaryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwRateLimitSummaryOperOper := setObjectFwRateLimitSummaryOperOper(res)
		d.Set("oper", FwRateLimitSummaryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwRateLimitSummaryOperOper(ret edpt.DataFwRateLimitSummaryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"mem_reserved":                                 ret.DtFwRateLimitSummaryOper.Oper.Mem_reserved,
			"mem_used":                                     ret.DtFwRateLimitSummaryOper.Oper.Mem_used,
			"alloc_failures":                               ret.DtFwRateLimitSummaryOper.Oper.Alloc_failures,
			"total_num_entries":                            ret.DtFwRateLimitSummaryOper.Oper.Total_num_entries,
			"total_entries_scope_aggregate":                ret.DtFwRateLimitSummaryOper.Oper.Total_entries_scope_aggregate,
			"total_entries_scope_subscriber_ip":            ret.DtFwRateLimitSummaryOper.Oper.Total_entries_scope_subscriber_ip,
			"total_entries_scope_subscriber_prefix":        ret.DtFwRateLimitSummaryOper.Oper.Total_entries_scope_subscriber_prefix,
			"total_entries_scope_parent":                   ret.DtFwRateLimitSummaryOper.Oper.Total_entries_scope_parent,
			"total_entries_scope_parent_subscriber_ip":     ret.DtFwRateLimitSummaryOper.Oper.Total_entries_scope_parent_subscriber_ip,
			"total_entries_scope_parent_subscriber_prefix": ret.DtFwRateLimitSummaryOper.Oper.Total_entries_scope_parent_subscriberPrefix,
		},
	}
}

func getObjectFwRateLimitSummaryOperOper(d []interface{}) edpt.FwRateLimitSummaryOperOper {

	count1 := len(d)
	var ret edpt.FwRateLimitSummaryOperOper
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

func dataToEndpointFwRateLimitSummaryOper(d *schema.ResourceData) edpt.FwRateLimitSummaryOper {
	var ret edpt.FwRateLimitSummaryOper

	ret.Oper = getObjectFwRateLimitSummaryOperOper(d.Get("oper").([]interface{}))
	return ret
}
