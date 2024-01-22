package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuDataCpuOverallOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_data_cpu_overall_oper`: Operational Status for the object data-cpu-overall\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuDataCpuOverallOperRead,

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
									"sec1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sec5": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sec10": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sec30": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sec60": {
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

func resourceSystemCpuDataCpuOverallOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuDataCpuOverallOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuDataCpuOverallOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCpuDataCpuOverallOperOper := setObjectSystemCpuDataCpuOverallOperOper(res)
		d.Set("oper", SystemCpuDataCpuOverallOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuDataCpuOverallOperOper(ret edpt.DataSystemCpuDataCpuOverallOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"number_of_cpu":      ret.DtSystemCpuDataCpuOverallOper.Oper.NumberOfCpu,
			"number_of_data_cpu": ret.DtSystemCpuDataCpuOverallOper.Oper.NumberOfDataCpu,
			"cpu_usage":          setSliceSystemCpuDataCpuOverallOperOperCpuUsage(ret.DtSystemCpuDataCpuOverallOper.Oper.CpuUsage),
		},
	}
}

func setSliceSystemCpuDataCpuOverallOperOperCpuUsage(d []edpt.SystemCpuDataCpuOverallOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_id"] = item.CpuId
		in["sec1"] = item.Sec1
		in["sec5"] = item.Sec5
		in["sec10"] = item.Sec10
		in["sec30"] = item.Sec30
		in["sec60"] = item.Sec60
		in["dcpu_str"] = item.DcpuStr
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuDataCpuOverallOperOper(d []interface{}) edpt.SystemCpuDataCpuOverallOperOper {

	count1 := len(d)
	var ret edpt.SystemCpuDataCpuOverallOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NumberOfCpu = in["number_of_cpu"].(int)
		ret.NumberOfDataCpu = in["number_of_data_cpu"].(int)
		ret.CpuUsage = getSliceSystemCpuDataCpuOverallOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuDataCpuOverallOperOperCpuUsage(d []interface{}) []edpt.SystemCpuDataCpuOverallOperOperCpuUsage {

	count1 := len(d)
	ret := make([]edpt.SystemCpuDataCpuOverallOperOperCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuDataCpuOverallOperOperCpuUsage
		oi.CpuId = in["cpu_id"].(int)
		oi.Sec1 = in["sec1"].(int)
		oi.Sec5 = in["sec5"].(int)
		oi.Sec10 = in["sec10"].(int)
		oi.Sec30 = in["sec30"].(int)
		oi.Sec60 = in["sec60"].(int)
		oi.DcpuStr = in["dcpu_str"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuDataCpuOverallOper(d *schema.ResourceData) edpt.SystemCpuDataCpuOverallOper {
	var ret edpt.SystemCpuDataCpuOverallOper

	ret.Oper = getObjectSystemCpuDataCpuOverallOperOper(d.Get("oper").([]interface{}))
	return ret
}
