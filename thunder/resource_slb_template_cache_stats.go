package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCacheStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_cache_stats`: Statistics for the object cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateCacheStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify cache template name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Cache hits",
						},
						"miss": {
							Type: schema.TypeInt, Optional: true, Description: "Cache misses",
						},
						"bytes_served": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes served from cache",
						},
						"total_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total requests received",
						},
						"caching_req": {
							Type: schema.TypeInt, Optional: true, Description: "Total requests to cache",
						},
						"nc_req_header": {
							Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheNcReqHeader, help nc_req_header",
						},
						"nc_res_header": {
							Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheNcResHeader, help nc_res_header",
						},
						"rv_success": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rv_failure": {
							Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheRvFailure, help rv_failure",
						},
						"ims_request": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nm_response": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_type_cl": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_type_ce": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_type_304": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_type_other": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_no_compress": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_gzip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_deflate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_other": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"nocache_match": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"match": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"invalidate_match": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"content_toobig": {
							Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheContentToobig, help content_toobig",
						},
						"content_toosmall": {
							Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheContentToosmall, help content_toosmall",
						},
						"entry_create_failures": {
							Type: schema.TypeInt, Optional: true, Description: "slbTemplateCacheEntryCreateFailures, help entry_create_failures",
						},
						"mem_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entry_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"replaced_entry": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"aging_entry": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cleaned_entry": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_type_stream": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"header_save_error": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"rsp_br": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplateCacheStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCacheStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCacheStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateCacheStatsStats := setObjectSlbTemplateCacheStatsStats(res)
		d.Set("stats", SlbTemplateCacheStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateCacheStatsStats(ret edpt.DataSlbTemplateCacheStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits":                  ret.DtSlbTemplateCacheStats.Stats.Hits,
			"miss":                  ret.DtSlbTemplateCacheStats.Stats.Miss,
			"bytes_served":          ret.DtSlbTemplateCacheStats.Stats.Bytes_served,
			"total_req":             ret.DtSlbTemplateCacheStats.Stats.Total_req,
			"caching_req":           ret.DtSlbTemplateCacheStats.Stats.Caching_req,
			"nc_req_header":         ret.DtSlbTemplateCacheStats.Stats.Nc_req_header,
			"nc_res_header":         ret.DtSlbTemplateCacheStats.Stats.Nc_res_header,
			"rv_success":            ret.DtSlbTemplateCacheStats.Stats.Rv_success,
			"rv_failure":            ret.DtSlbTemplateCacheStats.Stats.Rv_failure,
			"ims_request":           ret.DtSlbTemplateCacheStats.Stats.Ims_request,
			"nm_response":           ret.DtSlbTemplateCacheStats.Stats.Nm_response,
			"rsp_type_cl":           ret.DtSlbTemplateCacheStats.Stats.Rsp_type_cl,
			"rsp_type_ce":           ret.DtSlbTemplateCacheStats.Stats.Rsp_type_ce,
			"rsp_type_304":          ret.DtSlbTemplateCacheStats.Stats.Rsp_type_304,
			"rsp_type_other":        ret.DtSlbTemplateCacheStats.Stats.Rsp_type_other,
			"rsp_no_compress":       ret.DtSlbTemplateCacheStats.Stats.Rsp_no_compress,
			"rsp_gzip":              ret.DtSlbTemplateCacheStats.Stats.Rsp_gzip,
			"rsp_deflate":           ret.DtSlbTemplateCacheStats.Stats.Rsp_deflate,
			"rsp_other":             ret.DtSlbTemplateCacheStats.Stats.Rsp_other,
			"nocache_match":         ret.DtSlbTemplateCacheStats.Stats.Nocache_match,
			"match":                 ret.DtSlbTemplateCacheStats.Stats.Match,
			"invalidate_match":      ret.DtSlbTemplateCacheStats.Stats.Invalidate_match,
			"content_toobig":        ret.DtSlbTemplateCacheStats.Stats.Content_toobig,
			"content_toosmall":      ret.DtSlbTemplateCacheStats.Stats.Content_toosmall,
			"entry_create_failures": ret.DtSlbTemplateCacheStats.Stats.Entry_create_failures,
			"mem_size":              ret.DtSlbTemplateCacheStats.Stats.Mem_size,
			"entry_num":             ret.DtSlbTemplateCacheStats.Stats.Entry_num,
			"replaced_entry":        ret.DtSlbTemplateCacheStats.Stats.Replaced_entry,
			"aging_entry":           ret.DtSlbTemplateCacheStats.Stats.Aging_entry,
			"cleaned_entry":         ret.DtSlbTemplateCacheStats.Stats.Cleaned_entry,
			"rsp_type_stream":       ret.DtSlbTemplateCacheStats.Stats.Rsp_type_stream,
			"header_save_error":     ret.DtSlbTemplateCacheStats.Stats.Header_save_error,
			"rsp_br":                ret.DtSlbTemplateCacheStats.Stats.Rsp_br,
		},
	}
}

func getObjectSlbTemplateCacheStatsStats(d []interface{}) edpt.SlbTemplateCacheStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplateCacheStatsStats
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
		ret.Header_save_error = in["header_save_error"].(int)
		ret.Rsp_br = in["rsp_br"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateCacheStats(d *schema.ResourceData) edpt.SlbTemplateCacheStats {
	var ret edpt.SlbTemplateCacheStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplateCacheStatsStats(d.Get("stats").([]interface{}))
	return ret
}
