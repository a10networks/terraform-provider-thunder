package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRcCacheGlobalOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rc_cache_global_oper`: Operational Status for the object rc-cache-global\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRcCacheGlobalOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"url": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bytes": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"type": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"expires": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"host1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"url1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bytes1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"response_hdr_len": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status_code": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"etag": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cache_control": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"date": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"last_modified": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"time_elapsed": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"age": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"expires1": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"hits": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"misses": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"concurrent_readers": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"content_encoding": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"http_version": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"response_chunked_encoding": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"weak_etag": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"full_response_cache": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"http_request_method": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vserver_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"vport": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"memory_configured": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"memory_used": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"memory_used_locally": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"percent_used": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"cache_hits": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cache_miss": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"memory_used": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hit_ratio": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ratio304_": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bytes_served": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"total_request": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"no_cache_requests": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cacheable_requests": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ims_requests": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_server_304_not_modified": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_server_200_ok_chunk": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_server_no_cache_response": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_server_200_ok_cont": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_server_other": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_cache_304_not_modified": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_cache_200_ok_gzip": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_cache_other": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_cache_200_ok_no_comp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resp_cache_200_ok_deflate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entries_cached": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entries_aged": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entries_create_fail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entries_replaced": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"entries_cleaned": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"revalidation_success": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"revalidation_failure": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_uri_nocache": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"polic_uri_invalidate": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_content_small": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_uri_cache": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"policy_content_big": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"virtual_server": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"virtual_port": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"display_detail": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"uri_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"replacement_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"one_256th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_128th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_64th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_32th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_16th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_8th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_4th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_2th": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"two": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"four": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"eight": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sixteen": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"thirty_two": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sixty_four": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"one_twenty_eight": {
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

func resourceSlbRcCacheGlobalOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRcCacheGlobalOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRcCacheGlobalOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRcCacheGlobalOperOper := setObjectSlbRcCacheGlobalOperOper(res)
		d.Set("oper", SlbRcCacheGlobalOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRcCacheGlobalOperOper(ret edpt.DataSlbRcCacheGlobalOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"entry_list":                    setSliceSlbRcCacheGlobalOperOperEntryList(ret.DtSlbRcCacheGlobalOper.Oper.EntryList),
			"cache_hits":                    ret.DtSlbRcCacheGlobalOper.Oper.Cache_hits,
			"cache_miss":                    ret.DtSlbRcCacheGlobalOper.Oper.Cache_miss,
			"memory_used":                   ret.DtSlbRcCacheGlobalOper.Oper.Memory_used,
			"hit_ratio":                     ret.DtSlbRcCacheGlobalOper.Oper.Hit_ratio,
			"ratio304_":                     ret.DtSlbRcCacheGlobalOper.Oper.Ratio304_,
			"bytes_served":                  ret.DtSlbRcCacheGlobalOper.Oper.Bytes_served,
			"total_request":                 ret.DtSlbRcCacheGlobalOper.Oper.Total_request,
			"no_cache_requests":             ret.DtSlbRcCacheGlobalOper.Oper.No_cache_requests,
			"cacheable_requests":            ret.DtSlbRcCacheGlobalOper.Oper.Cacheable_requests,
			"ims_requests":                  ret.DtSlbRcCacheGlobalOper.Oper.Ims_requests,
			"resp_server_304_not_modified":  ret.DtSlbRcCacheGlobalOper.Oper.Resp_server_304_not_modified,
			"resp_server_200_ok_chunk":      ret.DtSlbRcCacheGlobalOper.Oper.Resp_server_200_ok_chunk,
			"resp_server_no_cache_response": ret.DtSlbRcCacheGlobalOper.Oper.Resp_server_no_cache_response,
			"resp_server_200_ok_cont":       ret.DtSlbRcCacheGlobalOper.Oper.Resp_server_200_ok_cont,
			"resp_server_other":             ret.DtSlbRcCacheGlobalOper.Oper.Resp_server_other,
			"resp_cache_304_not_modified":   ret.DtSlbRcCacheGlobalOper.Oper.Resp_cache_304_not_modified,
			"resp_cache_200_ok_gzip":        ret.DtSlbRcCacheGlobalOper.Oper.Resp_cache_200_ok_gzip,
			"resp_cache_other":              ret.DtSlbRcCacheGlobalOper.Oper.Resp_cache_other,
			"resp_cache_200_ok_no_comp":     ret.DtSlbRcCacheGlobalOper.Oper.Resp_cache_200_ok_no_comp,
			"resp_cache_200_ok_deflate":     ret.DtSlbRcCacheGlobalOper.Oper.Resp_cache_200_ok_deflate,
			"entries_cached":                ret.DtSlbRcCacheGlobalOper.Oper.Entries_cached,
			"entries_aged":                  ret.DtSlbRcCacheGlobalOper.Oper.Entries_aged,
			"entries_create_fail":           ret.DtSlbRcCacheGlobalOper.Oper.Entries_create_fail,
			"entries_replaced":              ret.DtSlbRcCacheGlobalOper.Oper.Entries_replaced,
			"entries_cleaned":               ret.DtSlbRcCacheGlobalOper.Oper.Entries_cleaned,
			"revalidation_success":          ret.DtSlbRcCacheGlobalOper.Oper.Revalidation_success,
			"revalidation_failure":          ret.DtSlbRcCacheGlobalOper.Oper.Revalidation_failure,
			"policy_uri_nocache":            ret.DtSlbRcCacheGlobalOper.Oper.Policy_uri_nocache,
			"polic_uri_invalidate":          ret.DtSlbRcCacheGlobalOper.Oper.Polic_uri_invalidate,
			"policy_content_small":          ret.DtSlbRcCacheGlobalOper.Oper.Policy_content_small,
			"policy_uri_cache":              ret.DtSlbRcCacheGlobalOper.Oper.Policy_uri_cache,
			"policy_content_big":            ret.DtSlbRcCacheGlobalOper.Oper.Policy_content_big,
			"virtual_server":                ret.DtSlbRcCacheGlobalOper.Oper.Virtual_server,
			"virtual_port":                  ret.DtSlbRcCacheGlobalOper.Oper.Virtual_port,
			"display_detail":                ret.DtSlbRcCacheGlobalOper.Oper.Display_detail,
			"uri_name":                      ret.DtSlbRcCacheGlobalOper.Oper.Uri_name,
			"replacement_list":              setSliceSlbRcCacheGlobalOperOperReplacementList(ret.DtSlbRcCacheGlobalOper.Oper.ReplacementList),
		},
	}
}

func setSliceSlbRcCacheGlobalOperOperEntryList(d []edpt.SlbRcCacheGlobalOperOperEntryList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["host"] = item.Host
		in["url"] = item.Url
		in["bytes"] = item.Bytes
		in["type"] = item.Type
		in["status"] = item.Status
		in["expires"] = item.Expires
		in["host1"] = item.Host1
		in["url1"] = item.Url1
		in["bytes1"] = item.Bytes1
		in["response_hdr_len"] = item.Response_hdr_len
		in["status_code"] = item.Status_code
		in["etag"] = item.Etag
		in["cache_control"] = item.Cache_control
		in["date"] = item.Date
		in["last_modified"] = item.Last_modified
		in["time_elapsed"] = item.Time_elapsed
		in["age"] = item.Age
		in["expires1"] = item.Expires1
		in["hits"] = item.Hits
		in["misses"] = item.Misses
		in["concurrent_readers"] = item.Concurrent_readers
		in["content_encoding"] = item.Content_encoding
		in["http_version"] = item.Http_version
		in["response_chunked_encoding"] = item.Response_chunked_encoding
		in["weak_etag"] = item.Weak_etag
		in["full_response_cache"] = item.Full_response_cache
		in["http_request_method"] = item.Http_request_method
		in["vserver_name"] = item.Vserver_name
		in["vport"] = item.Vport
		in["memory_configured"] = item.Memory_configured
		in["memory_used"] = item.Memory_used
		in["memory_used_locally"] = item.Memory_used_locally
		in["percent_used"] = item.Percent_used
		in["partition"] = item.Partition
		result = append(result, in)
	}
	return result
}

func setSliceSlbRcCacheGlobalOperOperReplacementList(d []edpt.SlbRcCacheGlobalOperOperReplacementList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["one_256th"] = item.One256th
		in["one_128th"] = item.One128th
		in["one_64th"] = item.One64th
		in["one_32th"] = item.One32th
		in["one_16th"] = item.One16th
		in["one_8th"] = item.One8th
		in["one_4th"] = item.One4th
		in["one_2th"] = item.One2th
		in["one"] = item.One
		in["two"] = item.Two
		in["four"] = item.Four
		in["eight"] = item.Eight
		in["sixteen"] = item.Sixteen
		in["thirty_two"] = item.ThirtyTwo
		in["sixty_four"] = item.SixtyFour
		in["one_twenty_eight"] = item.OneTwentyEight
		result = append(result, in)
	}
	return result
}

func getObjectSlbRcCacheGlobalOperOper(d []interface{}) edpt.SlbRcCacheGlobalOperOper {

	count1 := len(d)
	var ret edpt.SlbRcCacheGlobalOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EntryList = getSliceSlbRcCacheGlobalOperOperEntryList(in["entry_list"].([]interface{}))
		ret.Cache_hits = in["cache_hits"].(int)
		ret.Cache_miss = in["cache_miss"].(int)
		ret.Memory_used = in["memory_used"].(string)
		ret.Hit_ratio = in["hit_ratio"].(string)
		ret.Ratio304_ = in["ratio304_"].(string)
		ret.Bytes_served = in["bytes_served"].(string)
		ret.Total_request = in["total_request"].(int)
		ret.No_cache_requests = in["no_cache_requests"].(int)
		ret.Cacheable_requests = in["cacheable_requests"].(int)
		ret.Ims_requests = in["ims_requests"].(int)
		ret.Resp_server_304_not_modified = in["resp_server_304_not_modified"].(int)
		ret.Resp_server_200_ok_chunk = in["resp_server_200_ok_chunk"].(int)
		ret.Resp_server_no_cache_response = in["resp_server_no_cache_response"].(int)
		ret.Resp_server_200_ok_cont = in["resp_server_200_ok_cont"].(int)
		ret.Resp_server_other = in["resp_server_other"].(int)
		ret.Resp_cache_304_not_modified = in["resp_cache_304_not_modified"].(int)
		ret.Resp_cache_200_ok_gzip = in["resp_cache_200_ok_gzip"].(int)
		ret.Resp_cache_other = in["resp_cache_other"].(int)
		ret.Resp_cache_200_ok_no_comp = in["resp_cache_200_ok_no_comp"].(int)
		ret.Resp_cache_200_ok_deflate = in["resp_cache_200_ok_deflate"].(int)
		ret.Entries_cached = in["entries_cached"].(int)
		ret.Entries_aged = in["entries_aged"].(int)
		ret.Entries_create_fail = in["entries_create_fail"].(int)
		ret.Entries_replaced = in["entries_replaced"].(int)
		ret.Entries_cleaned = in["entries_cleaned"].(int)
		ret.Revalidation_success = in["revalidation_success"].(int)
		ret.Revalidation_failure = in["revalidation_failure"].(int)
		ret.Policy_uri_nocache = in["policy_uri_nocache"].(int)
		ret.Polic_uri_invalidate = in["polic_uri_invalidate"].(int)
		ret.Policy_content_small = in["policy_content_small"].(int)
		ret.Policy_uri_cache = in["policy_uri_cache"].(int)
		ret.Policy_content_big = in["policy_content_big"].(int)
		ret.Virtual_server = in["virtual_server"].(string)
		ret.Virtual_port = in["virtual_port"].(int)
		ret.Display_detail = in["display_detail"].(int)
		ret.Uri_name = in["uri_name"].(string)
		ret.ReplacementList = getSliceSlbRcCacheGlobalOperOperReplacementList(in["replacement_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbRcCacheGlobalOperOperEntryList(d []interface{}) []edpt.SlbRcCacheGlobalOperOperEntryList {

	count1 := len(d)
	ret := make([]edpt.SlbRcCacheGlobalOperOperEntryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRcCacheGlobalOperOperEntryList
		oi.Host = in["host"].(string)
		oi.Url = in["url"].(string)
		oi.Bytes = in["bytes"].(int)
		oi.Type = in["type"].(string)
		oi.Status = in["status"].(string)
		oi.Expires = in["expires"].(string)
		oi.Host1 = in["host1"].(string)
		oi.Url1 = in["url1"].(string)
		oi.Bytes1 = in["bytes1"].(string)
		oi.Response_hdr_len = in["response_hdr_len"].(string)
		oi.Status_code = in["status_code"].(string)
		oi.Etag = in["etag"].(string)
		oi.Cache_control = in["cache_control"].(string)
		oi.Date = in["date"].(string)
		oi.Last_modified = in["last_modified"].(string)
		oi.Time_elapsed = in["time_elapsed"].(string)
		oi.Age = in["age"].(string)
		oi.Expires1 = in["expires1"].(string)
		oi.Hits = in["hits"].(string)
		oi.Misses = in["misses"].(string)
		oi.Concurrent_readers = in["concurrent_readers"].(string)
		oi.Content_encoding = in["content_encoding"].(string)
		oi.Http_version = in["http_version"].(string)
		oi.Response_chunked_encoding = in["response_chunked_encoding"].(string)
		oi.Weak_etag = in["weak_etag"].(string)
		oi.Full_response_cache = in["full_response_cache"].(string)
		oi.Http_request_method = in["http_request_method"].(string)
		oi.Vserver_name = in["vserver_name"].(string)
		oi.Vport = in["vport"].(string)
		oi.Memory_configured = in["memory_configured"].(string)
		oi.Memory_used = in["memory_used"].(string)
		oi.Memory_used_locally = in["memory_used_locally"].(string)
		oi.Percent_used = in["percent_used"].(string)
		oi.Partition = in["partition"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbRcCacheGlobalOperOperReplacementList(d []interface{}) []edpt.SlbRcCacheGlobalOperOperReplacementList {

	count1 := len(d)
	ret := make([]edpt.SlbRcCacheGlobalOperOperReplacementList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbRcCacheGlobalOperOperReplacementList
		oi.One256th = in["one_256th"].(int)
		oi.One128th = in["one_128th"].(int)
		oi.One64th = in["one_64th"].(int)
		oi.One32th = in["one_32th"].(int)
		oi.One16th = in["one_16th"].(int)
		oi.One8th = in["one_8th"].(int)
		oi.One4th = in["one_4th"].(int)
		oi.One2th = in["one_2th"].(int)
		oi.One = in["one"].(int)
		oi.Two = in["two"].(int)
		oi.Four = in["four"].(int)
		oi.Eight = in["eight"].(int)
		oi.Sixteen = in["sixteen"].(int)
		oi.ThirtyTwo = in["thirty_two"].(int)
		oi.SixtyFour = in["sixty_four"].(int)
		oi.OneTwentyEight = in["one_twenty_eight"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbRcCacheGlobalOper(d *schema.ResourceData) edpt.SlbRcCacheGlobalOper {
	var ret edpt.SlbRcCacheGlobalOper

	ret.Oper = getObjectSlbRcCacheGlobalOperOper(d.Get("oper").([]interface{}))
	return ret
}
