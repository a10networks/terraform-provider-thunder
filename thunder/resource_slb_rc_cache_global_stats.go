package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRcCacheGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rc_cache_global_stats`: Statistics for the object rc-cache-global\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRcCacheGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Cache Hits",
						},
						"miss": {
							Type: schema.TypeInt, Optional: true, Description: "Cache Misses",
						},
						"bytes_served": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes Served",
						},
						"total_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total Requests",
						},
						"caching_req": {
							Type: schema.TypeInt, Optional: true, Description: "Cacheable Requests",
						},
						"nc_req_header": {
							Type: schema.TypeInt, Optional: true, Description: "No-cache Request",
						},
						"nc_res_header": {
							Type: schema.TypeInt, Optional: true, Description: "Not cacheable",
						},
						"rv_success": {
							Type: schema.TypeInt, Optional: true, Description: "Revalidation Successes",
						},
						"rv_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Revalidation Failures",
						},
						"ims_request": {
							Type: schema.TypeInt, Optional: true, Description: "IMS Requests",
						},
						"nm_response": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from cache 304 Not Modified",
						},
						"rsp_type_cl": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from server 200 OK - Cont Len",
						},
						"rsp_type_ce": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from server 200 OK - Chnk Enc",
						},
						"rsp_type_304": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from server 304 Not Modified",
						},
						"rsp_type_other": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from server 200 OK - Other",
						},
						"rsp_no_compress": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from cache 200 OK - No Comp",
						},
						"rsp_gzip": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from cache 200 OK - Gzip",
						},
						"rsp_deflate": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from cache 200 OK - Deflate",
						},
						"rsp_other": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from cache Other",
						},
						"nocache_match": {
							Type: schema.TypeInt, Optional: true, Description: "Policy URI nocache",
						},
						"match": {
							Type: schema.TypeInt, Optional: true, Description: "Policy URI cache",
						},
						"invalidate_match": {
							Type: schema.TypeInt, Optional: true, Description: "Policy URI invalidate",
						},
						"content_toobig": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Content Too Big",
						},
						"content_toosmall": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Content Too Small",
						},
						"entry_create_failures": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Create failures",
						},
						"mem_size": {
							Type: schema.TypeInt, Optional: true, Description: "Memory Used",
						},
						"entry_num": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Cached",
						},
						"replaced_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Replaced",
						},
						"aging_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Aged Out",
						},
						"cleaned_entry": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Cleaned",
						},
						"rsp_type_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from http2 server 200 OK - Stream",
						},
						"rsp_br": {
							Type: schema.TypeInt, Optional: true, Description: "Responses from cache 200 OK - Brotli",
						},
					},
				},
			},
		},
	}
}

func resourceSlbRcCacheGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRcCacheGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRcCacheGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRcCacheGlobalStatsStats := setObjectSlbRcCacheGlobalStatsStats(res)
		d.Set("stats", SlbRcCacheGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRcCacheGlobalStatsStats(ret edpt.DataSlbRcCacheGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits":                  ret.DtSlbRcCacheGlobalStats.Stats.Hits,
			"miss":                  ret.DtSlbRcCacheGlobalStats.Stats.Miss,
			"bytes_served":          ret.DtSlbRcCacheGlobalStats.Stats.Bytes_served,
			"total_req":             ret.DtSlbRcCacheGlobalStats.Stats.Total_req,
			"caching_req":           ret.DtSlbRcCacheGlobalStats.Stats.Caching_req,
			"nc_req_header":         ret.DtSlbRcCacheGlobalStats.Stats.Nc_req_header,
			"nc_res_header":         ret.DtSlbRcCacheGlobalStats.Stats.Nc_res_header,
			"rv_success":            ret.DtSlbRcCacheGlobalStats.Stats.Rv_success,
			"rv_failure":            ret.DtSlbRcCacheGlobalStats.Stats.Rv_failure,
			"ims_request":           ret.DtSlbRcCacheGlobalStats.Stats.Ims_request,
			"nm_response":           ret.DtSlbRcCacheGlobalStats.Stats.Nm_response,
			"rsp_type_cl":           ret.DtSlbRcCacheGlobalStats.Stats.Rsp_type_cl,
			"rsp_type_ce":           ret.DtSlbRcCacheGlobalStats.Stats.Rsp_type_ce,
			"rsp_type_304":          ret.DtSlbRcCacheGlobalStats.Stats.Rsp_type_304,
			"rsp_type_other":        ret.DtSlbRcCacheGlobalStats.Stats.Rsp_type_other,
			"rsp_no_compress":       ret.DtSlbRcCacheGlobalStats.Stats.Rsp_no_compress,
			"rsp_gzip":              ret.DtSlbRcCacheGlobalStats.Stats.Rsp_gzip,
			"rsp_deflate":           ret.DtSlbRcCacheGlobalStats.Stats.Rsp_deflate,
			"rsp_other":             ret.DtSlbRcCacheGlobalStats.Stats.Rsp_other,
			"nocache_match":         ret.DtSlbRcCacheGlobalStats.Stats.Nocache_match,
			"match":                 ret.DtSlbRcCacheGlobalStats.Stats.Match,
			"invalidate_match":      ret.DtSlbRcCacheGlobalStats.Stats.Invalidate_match,
			"content_toobig":        ret.DtSlbRcCacheGlobalStats.Stats.Content_toobig,
			"content_toosmall":      ret.DtSlbRcCacheGlobalStats.Stats.Content_toosmall,
			"entry_create_failures": ret.DtSlbRcCacheGlobalStats.Stats.Entry_create_failures,
			"mem_size":              ret.DtSlbRcCacheGlobalStats.Stats.Mem_size,
			"entry_num":             ret.DtSlbRcCacheGlobalStats.Stats.Entry_num,
			"replaced_entry":        ret.DtSlbRcCacheGlobalStats.Stats.Replaced_entry,
			"aging_entry":           ret.DtSlbRcCacheGlobalStats.Stats.Aging_entry,
			"cleaned_entry":         ret.DtSlbRcCacheGlobalStats.Stats.Cleaned_entry,
			"rsp_type_stream":       ret.DtSlbRcCacheGlobalStats.Stats.Rsp_type_stream,
			"rsp_br":                ret.DtSlbRcCacheGlobalStats.Stats.Rsp_br,
		},
	}
}

func getObjectSlbRcCacheGlobalStatsStats(d []interface{}) edpt.SlbRcCacheGlobalStatsStats {

	count1 := len(d)
	var ret edpt.SlbRcCacheGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
		ret.Miss = in["miss"].(int)
		ret.Bytes_served = in["bytes_served"].(int)
		ret.Total_req = in["total_req"].(int)
		ret.Caching_req = in["caching_req"].(int)
		ret.Nc_req_header = in["nc_req_header"].(int)
		ret.Nc_res_header = in["nc_res_header"].(int)
		ret.Rv_success = in["rv_success"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Ims_request = in["ims_request"].(int)
		ret.Nm_response = in["nm_response"].(int)
		ret.Rsp_type_cl = in["rsp_type_cl"].(int)
		ret.Rsp_type_ce = in["rsp_type_ce"].(int)
		ret.Rsp_type_304 = in["rsp_type_304"].(int)
		ret.Rsp_type_other = in["rsp_type_other"].(int)
		ret.Rsp_no_compress = in["rsp_no_compress"].(int)
		ret.Rsp_gzip = in["rsp_gzip"].(int)
		ret.Rsp_deflate = in["rsp_deflate"].(int)
		ret.Rsp_other = in["rsp_other"].(int)
		ret.Nocache_match = in["nocache_match"].(int)
		ret.Match = in["match"].(int)
		ret.Invalidate_match = in["invalidate_match"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		ret.Mem_size = in["mem_size"].(int)
		ret.Entry_num = in["entry_num"].(int)
		ret.Replaced_entry = in["replaced_entry"].(int)
		ret.Aging_entry = in["aging_entry"].(int)
		ret.Cleaned_entry = in["cleaned_entry"].(int)
		ret.Rsp_type_stream = in["rsp_type_stream"].(int)
		ret.Rsp_br = in["rsp_br"].(int)
	}
	return ret
}

func dataToEndpointSlbRcCacheGlobalStats(d *schema.ResourceData) edpt.SlbRcCacheGlobalStats {
	var ret edpt.SlbRcCacheGlobalStats

	ret.Stats = getObjectSlbRcCacheGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
