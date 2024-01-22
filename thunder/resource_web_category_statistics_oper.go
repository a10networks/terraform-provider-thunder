package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryStatisticsOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_statistics_oper`: Operational Status for the object statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryStatisticsOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_dplane_threads": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"num_lookup_threads": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"per_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"req_queue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_processed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"req_lookup_processed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_req_queue": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_req_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_req_processed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"total_req_lookup_processed": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"clear_cache": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryStatisticsOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryStatisticsOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryStatisticsOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryStatisticsOperOper := setObjectWebCategoryStatisticsOperOper(res)
		d.Set("oper", WebCategoryStatisticsOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryStatisticsOperOper(ret edpt.DataWebCategoryStatisticsOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_dplane_threads":         ret.DtWebCategoryStatisticsOper.Oper.NumDplaneThreads,
			"num_lookup_threads":         ret.DtWebCategoryStatisticsOper.Oper.NumLookupThreads,
			"per_cpu_list":               setSliceWebCategoryStatisticsOperOperPerCpuList(ret.DtWebCategoryStatisticsOper.Oper.PerCpuList),
			"total_req_queue":            ret.DtWebCategoryStatisticsOper.Oper.TotalReqQueue,
			"total_req_dropped":          ret.DtWebCategoryStatisticsOper.Oper.TotalReqDropped,
			"total_req_processed":        ret.DtWebCategoryStatisticsOper.Oper.TotalReqProcessed,
			"total_req_lookup_processed": ret.DtWebCategoryStatisticsOper.Oper.TotalReqLookupProcessed,
			"clear_cache":                ret.DtWebCategoryStatisticsOper.Oper.ClearCache,
		},
	}
}

func setSliceWebCategoryStatisticsOperOperPerCpuList(d []edpt.WebCategoryStatisticsOperOperPerCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["req_queue"] = item.ReqQueue
		in["req_dropped"] = item.ReqDropped
		in["req_processed"] = item.ReqProcessed
		in["req_lookup_processed"] = item.ReqLookupProcessed
		result = append(result, in)
	}
	return result
}

func getObjectWebCategoryStatisticsOperOper(d []interface{}) edpt.WebCategoryStatisticsOperOper {

	count1 := len(d)
	var ret edpt.WebCategoryStatisticsOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumDplaneThreads = in["num_dplane_threads"].(int)
		ret.NumLookupThreads = in["num_lookup_threads"].(int)
		ret.PerCpuList = getSliceWebCategoryStatisticsOperOperPerCpuList(in["per_cpu_list"].([]interface{}))
		ret.TotalReqQueue = in["total_req_queue"].(int)
		ret.TotalReqDropped = in["total_req_dropped"].(int)
		ret.TotalReqProcessed = in["total_req_processed"].(int)
		ret.TotalReqLookupProcessed = in["total_req_lookup_processed"].(int)
		ret.ClearCache = in["clear_cache"].(string)
	}
	return ret
}

func getSliceWebCategoryStatisticsOperOperPerCpuList(d []interface{}) []edpt.WebCategoryStatisticsOperOperPerCpuList {

	count1 := len(d)
	ret := make([]edpt.WebCategoryStatisticsOperOperPerCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.WebCategoryStatisticsOperOperPerCpuList
		oi.ReqQueue = in["req_queue"].(int)
		oi.ReqDropped = in["req_dropped"].(int)
		oi.ReqProcessed = in["req_processed"].(int)
		oi.ReqLookupProcessed = in["req_lookup_processed"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointWebCategoryStatisticsOper(d *schema.ResourceData) edpt.WebCategoryStatisticsOper {
	var ret edpt.WebCategoryStatisticsOper

	ret.Oper = getObjectWebCategoryStatisticsOperOper(d.Get("oper").([]interface{}))
	return ret
}
