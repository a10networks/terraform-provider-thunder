package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuVserverDataCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_vserver_data_cpu_oper`: Operational Status for the object vserver-data-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuVserverDataCpuOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vserver_cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vserver_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dcpu_str": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSystemCpuVserverDataCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuVserverDataCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuVserverDataCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCpuVserverDataCpuOperOper := setObjectSystemCpuVserverDataCpuOperOper(res)
		d.Set("oper", SystemCpuVserverDataCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuVserverDataCpuOperOper(ret edpt.DataSystemCpuVserverDataCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vserver_cpu_usage": setSliceSystemCpuVserverDataCpuOperOperVserverCpuUsage(ret.DtSystemCpuVserverDataCpuOper.Oper.VserverCpuUsage),
		},
	}
}

func setSliceSystemCpuVserverDataCpuOperOperVserverCpuUsage(d []edpt.SystemCpuVserverDataCpuOperOperVserverCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vserver_name"] = item.VserverName
		in["cpu_id"] = item.CpuId
		in["dcpu_str"] = item.DcpuStr
		in["sec1"] = item.Sec1
		in["sec5"] = item.Sec5
		in["sec10"] = item.Sec10
		in["sec30"] = item.Sec30
		in["sec60"] = item.Sec60
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuVserverDataCpuOperOper(d []interface{}) edpt.SystemCpuVserverDataCpuOperOper {

	count1 := len(d)
	var ret edpt.SystemCpuVserverDataCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VserverCpuUsage = getSliceSystemCpuVserverDataCpuOperOperVserverCpuUsage(in["vserver_cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuVserverDataCpuOperOperVserverCpuUsage(d []interface{}) []edpt.SystemCpuVserverDataCpuOperOperVserverCpuUsage {

	count1 := len(d)
	ret := make([]edpt.SystemCpuVserverDataCpuOperOperVserverCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuVserverDataCpuOperOperVserverCpuUsage
		oi.VserverName = in["vserver_name"].(string)
		oi.CpuId = in["cpu_id"].(int)
		oi.DcpuStr = in["dcpu_str"].(string)
		oi.Sec1 = in["sec1"].(int)
		oi.Sec5 = in["sec5"].(int)
		oi.Sec10 = in["sec10"].(int)
		oi.Sec30 = in["sec30"].(int)
		oi.Sec60 = in["sec60"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuVserverDataCpuOper(d *schema.ResourceData) edpt.SystemCpuVserverDataCpuOper {
	var ret edpt.SystemCpuVserverDataCpuOper

	ret.Oper = getObjectSystemCpuVserverDataCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
