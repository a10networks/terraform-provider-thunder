package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHwCompressStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_hw_compress_stats`: Statistics for the object hw-compress\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHwCompressStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total request count",
						},
						"submit_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total submit count",
						},
						"response_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total response count",
						},
						"failure_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total failure count",
						},
						"failure_code": {
							Type: schema.TypeInt, Optional: true, Description: "Last failure code",
						},
						"ring_full_count": {
							Type: schema.TypeInt, Optional: true, Description: "Compression queue full",
						},
						"max_outstanding_request_count": {
							Type: schema.TypeInt, Optional: true, Description: "Max queued request count",
						},
						"max_outstanding_submit_count": {
							Type: schema.TypeInt, Optional: true, Description: "Max queued submit count",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHwCompressStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHwCompressStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHwCompressStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHwCompressStatsStats := setObjectSlbHwCompressStatsStats(res)
		d.Set("stats", SlbHwCompressStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHwCompressStatsStats(ret edpt.DataSlbHwCompressStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_count":                 ret.DtSlbHwCompressStats.Stats.Request_count,
			"submit_count":                  ret.DtSlbHwCompressStats.Stats.Submit_count,
			"response_count":                ret.DtSlbHwCompressStats.Stats.Response_count,
			"failure_count":                 ret.DtSlbHwCompressStats.Stats.Failure_count,
			"failure_code":                  ret.DtSlbHwCompressStats.Stats.Failure_code,
			"ring_full_count":               ret.DtSlbHwCompressStats.Stats.Ring_full_count,
			"max_outstanding_request_count": ret.DtSlbHwCompressStats.Stats.Max_outstanding_request_count,
			"max_outstanding_submit_count":  ret.DtSlbHwCompressStats.Stats.Max_outstanding_submit_count,
		},
	}
}

func getObjectSlbHwCompressStatsStats(d []interface{}) edpt.SlbHwCompressStatsStats {

	count1 := len(d)
	var ret edpt.SlbHwCompressStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Request_count = in["request_count"].(int)
		ret.Submit_count = in["submit_count"].(int)
		ret.Response_count = in["response_count"].(int)
		ret.Failure_count = in["failure_count"].(int)
		ret.Failure_code = in["failure_code"].(int)
		ret.Ring_full_count = in["ring_full_count"].(int)
		ret.Max_outstanding_request_count = in["max_outstanding_request_count"].(int)
		ret.Max_outstanding_submit_count = in["max_outstanding_submit_count"].(int)
	}
	return ret
}

func dataToEndpointSlbHwCompressStats(d *schema.ResourceData) edpt.SlbHwCompressStats {
	var ret edpt.SlbHwCompressStats

	ret.Stats = getObjectSlbHwCompressStatsStats(d.Get("stats").([]interface{}))
	return ret
}
