package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHwCompressOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_hw_compress_oper`: Operational Status for the object hw-compress\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHwCompressOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"request_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"submit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"response_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"failure_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"failure_code": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ring_full_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_outstanding_request_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_outstanding_submit_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hw_compress_disabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHwCompressOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHwCompressOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHwCompressOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHwCompressOperOper := setObjectSlbHwCompressOperOper(res)
		d.Set("oper", SlbHwCompressOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHwCompressOperOper(ret edpt.DataSlbHwCompressOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l4_cpu_list":          setSliceSlbHwCompressOperOperL4CpuList(ret.DtSlbHwCompressOper.Oper.L4CpuList),
			"cpu_count":            ret.DtSlbHwCompressOper.Oper.CpuCount,
			"hw_compress_disabled": ret.DtSlbHwCompressOper.Oper.HwCompressDisabled,
		},
	}
}

func setSliceSlbHwCompressOperOperL4CpuList(d []edpt.SlbHwCompressOperOperL4CpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["request_count"] = item.Request_count
		in["submit_count"] = item.Submit_count
		in["response_count"] = item.Response_count
		in["failure_count"] = item.Failure_count
		in["failure_code"] = item.Failure_code
		in["ring_full_count"] = item.Ring_full_count
		in["max_outstanding_request_count"] = item.Max_outstanding_request_count
		in["max_outstanding_submit_count"] = item.Max_outstanding_submit_count
		result = append(result, in)
	}
	return result
}

func getObjectSlbHwCompressOperOper(d []interface{}) edpt.SlbHwCompressOperOper {

	count1 := len(d)
	var ret edpt.SlbHwCompressOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4CpuList = getSliceSlbHwCompressOperOperL4CpuList(in["l4_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
		ret.HwCompressDisabled = in["hw_compress_disabled"].(int)
	}
	return ret
}

func getSliceSlbHwCompressOperOperL4CpuList(d []interface{}) []edpt.SlbHwCompressOperOperL4CpuList {

	count1 := len(d)
	ret := make([]edpt.SlbHwCompressOperOperL4CpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHwCompressOperOperL4CpuList
		oi.Request_count = in["request_count"].(int)
		oi.Submit_count = in["submit_count"].(int)
		oi.Response_count = in["response_count"].(int)
		oi.Failure_count = in["failure_count"].(int)
		oi.Failure_code = in["failure_code"].(int)
		oi.Ring_full_count = in["ring_full_count"].(int)
		oi.Max_outstanding_request_count = in["max_outstanding_request_count"].(int)
		oi.Max_outstanding_submit_count = in["max_outstanding_submit_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHwCompressOper(d *schema.ResourceData) edpt.SlbHwCompressOper {
	var ret edpt.SlbHwCompressOper

	ret.Oper = getObjectSlbHwCompressOperOper(d.Get("oper").([]interface{}))
	return ret
}
