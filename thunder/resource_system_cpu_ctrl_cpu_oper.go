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
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCpuCtrlCpuOperOper := setObjectSystemCpuCtrlCpuOperOper(res)
		d.Set("oper", SystemCpuCtrlCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuCtrlCpuOperOper(ret edpt.DataSystemCpuCtrlCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"current_time":  ret.DtSystemCpuCtrlCpuOper.Oper.CurrentTime,
			"number_of_cpu": ret.DtSystemCpuCtrlCpuOper.Oper.NumberOfCpu,
			"cpu_usage":     setSliceSystemCpuCtrlCpuOperOperCpuUsage(ret.DtSystemCpuCtrlCpuOper.Oper.CpuUsage),
		},
	}
}

func setSliceSystemCpuCtrlCpuOperOperCpuUsage(d []edpt.SystemCpuCtrlCpuOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_id"] = item.CpuId
		in["sec1"] = item.Sec1
		in["sec5"] = item.Sec5
		in["sec10"] = item.Sec10
		in["sec30"] = item.Sec30
		in["sec60"] = item.Sec60
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuCtrlCpuOperOper(d []interface{}) edpt.SystemCpuCtrlCpuOperOper {

	count1 := len(d)
	var ret edpt.SystemCpuCtrlCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CurrentTime = in["current_time"].(string)
		ret.NumberOfCpu = in["number_of_cpu"].(int)
		ret.CpuUsage = getSliceSystemCpuCtrlCpuOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuCtrlCpuOperOperCpuUsage(d []interface{}) []edpt.SystemCpuCtrlCpuOperOperCpuUsage {

	count1 := len(d)
	ret := make([]edpt.SystemCpuCtrlCpuOperOperCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuCtrlCpuOperOperCpuUsage
		oi.CpuId = in["cpu_id"].(int)
		oi.Sec1 = in["sec1"].(int)
		oi.Sec5 = in["sec5"].(int)
		oi.Sec10 = in["sec10"].(int)
		oi.Sec30 = in["sec30"].(int)
		oi.Sec60 = in["sec60"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuCtrlCpuOper(d *schema.ResourceData) edpt.SystemCpuCtrlCpuOper {
	var ret edpt.SystemCpuCtrlCpuOper

	ret.Oper = getObjectSystemCpuCtrlCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
