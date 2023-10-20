package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuCtrlCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_ctrl_cpu_oper`: Operational Status for the object ctrl-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuCtrlCpuOperRead,
		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"number_of_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "", Default: 0,
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceSystemCpuCtrlCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuCtrlCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuCtrlCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		items := setObjectSystemCpuCtrlCpuOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", items)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func setObjectSystemCpuCtrlCpuOperOper(res edpt.SystemCpuCtrlCpus) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})
	in["current_time"] = res.DataSystemCpuCtrlCpu.Oper.CurrentTime
	in["number_of_cpu"] = res.DataSystemCpuCtrlCpu.Oper.NumberOfCpu
	in["cpu_usage"] = setSliceSystemCpuCtrlCpuOperOperCpuUsage(res.DataSystemCpuCtrlCpu.Oper.CpuUsage)
	result = append(result, in)
	return result
}

func setSliceSystemCpuCtrlCpuOperOperCpuUsage(d []edpt.SystemCpuCtrlCpuOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_id"] = item.CpuId
		in["1_sec"] = item.Sec1
		in["5_sec"] = item.Sec5
		in["10_sec"] = item.Sec10
		in["30_sec"] = item.Sec30
		in["60_sec"] = item.Sec60
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuCtrlCpuOperOper(d []interface{}) edpt.SystemCpuCtrlCpuOperOper {
	count := len(d)
	var ret edpt.SystemCpuCtrlCpuOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrentTime = in["current_time"].(string)
		ret.NumberOfCpu = in["number_of_cpu"].(int)
		ret.CpuUsage = getSliceSystemCpuCtrlCpuOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuCtrlCpuOperOperCpuUsage(d []interface{}) []edpt.SystemCpuCtrlCpuOperOperCpuUsage {
	count := len(d)
	ret := make([]edpt.SystemCpuCtrlCpuOperOperCpuUsage, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuCtrlCpuOperOperCpuUsage
		oi.CpuId = in["cpu_id"].(int)
		oi.Sec1 = in["1_sec"].(int)
		oi.Sec5 = in["5_sec"].(int)
		oi.Sec10 = in["10_sec"].(int)
		oi.Sec30 = in["30_sec"].(int)
		oi.Sec60 = in["60_sec"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuCtrlCpuOper(d *schema.ResourceData) edpt.SystemCpuCtrlCpuOper {
	var ret edpt.SystemCpuCtrlCpuOper
	ret.Oper = getObjectSystemCpuCtrlCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
