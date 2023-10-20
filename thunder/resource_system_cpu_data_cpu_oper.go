package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuDataCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_data_cpu_oper`: Operational Status for the object data-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuDataCpuOperRead,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number_of_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"number_of_data_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"1_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"5_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"10_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"30_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"60_sec": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcpu_str": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSystemCpuDataCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuDataCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuDataCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		items := setObjectSystemCpuDataCpuOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", items)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuDataCpuOperOper(res edpt.SystemCpuDataCpus) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["number_of_cpu"] = res.DataSystemCpuDataCpu.Oper.NumberOfCpu
	in["number_of_data_cpu"] = res.DataSystemCpuDataCpu.Oper.NumberOfDataCpu
	in["cpu_usage"] = setSliceSystemCpuDataCpuOperOperCpuUsage(res.DataSystemCpuDataCpu.Oper.CpuUsage)
	result = append(result, in)
	return result
}

func setSliceSystemCpuDataCpuOperOperCpuUsage(d []edpt.SystemCpuDataCpuOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_id"] = item.CpuId
		in["1_sec"] = item.Sec1
		in["5_sec"] = item.Sec5
		in["10_sec"] = item.Sec10
		in["30_sec"] = item.Sec30
		in["60_sec"] = item.Sec60
		in["dcpu_str"] = item.DcpuStr
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuDataCpuOperOper(d []interface{}) edpt.SystemCpuDataCpuOperOper {
	count := len(d)
	var ret edpt.SystemCpuDataCpuOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.NumberOfCpu = in["number_of_cpu"].(int)
		ret.NumberOfDataCpu = in["number_of_data_cpu"].(int)
		ret.CpuUsage = getSliceSystemCpuDataCpuOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuDataCpuOperOperCpuUsage(d []interface{}) []edpt.SystemCpuDataCpuOperOperCpuUsage {
	count := len(d)
	ret := make([]edpt.SystemCpuDataCpuOperOperCpuUsage, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuDataCpuOperOperCpuUsage
		oi.CpuId = in["cpu_id"].(int)
		oi.Sec1 = in["1_sec"].(int)
		oi.Sec5 = in["5_sec"].(int)
		oi.Sec10 = in["10_sec"].(int)
		oi.Sec30 = in["30_sec"].(int)
		oi.Sec60 = in["60_sec"].(int)
		oi.DcpuStr = in["dcpu_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuDataCpuOper(d *schema.ResourceData) edpt.SystemCpuDataCpuOper {
	var ret edpt.SystemCpuDataCpuOper
	ret.Oper = getObjectSystemCpuDataCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
