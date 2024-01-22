package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuVportDataCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_vport_data_cpu_oper`: Operational Status for the object vport-data-cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuVportDataCpuOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vport_cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vserver_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"portnumber": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"protocol": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"vport_type": {
										Type: schema.TypeString, Optional: true, Description: "",
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

func resourceSystemCpuVportDataCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuVportDataCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuVportDataCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCpuVportDataCpuOperOper := setObjectSystemCpuVportDataCpuOperOper(res)
		d.Set("oper", SystemCpuVportDataCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuVportDataCpuOperOper(ret edpt.DataSystemCpuVportDataCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"vport_cpu_usage": setSliceSystemCpuVportDataCpuOperOperVportCpuUsage(ret.DtSystemCpuVportDataCpuOper.Oper.VportCpuUsage),
		},
	}
}

func setSliceSystemCpuVportDataCpuOperOperVportCpuUsage(d []edpt.SystemCpuVportDataCpuOperOperVportCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["vserver_name"] = item.VserverName
		in["portnumber"] = item.Portnumber
		in["protocol"] = item.Protocol
		in["cpu_id"] = item.CpuId
		in["vport_type"] = item.VportType
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

func getObjectSystemCpuVportDataCpuOperOper(d []interface{}) edpt.SystemCpuVportDataCpuOperOper {

	count1 := len(d)
	var ret edpt.SystemCpuVportDataCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.VportCpuUsage = getSliceSystemCpuVportDataCpuOperOperVportCpuUsage(in["vport_cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuVportDataCpuOperOperVportCpuUsage(d []interface{}) []edpt.SystemCpuVportDataCpuOperOperVportCpuUsage {

	count1 := len(d)
	ret := make([]edpt.SystemCpuVportDataCpuOperOperVportCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuVportDataCpuOperOperVportCpuUsage
		oi.VserverName = in["vserver_name"].(string)
		oi.Portnumber = in["portnumber"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.CpuId = in["cpu_id"].(int)
		oi.VportType = in["vport_type"].(string)
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

func dataToEndpointSystemCpuVportDataCpuOper(d *schema.ResourceData) edpt.SystemCpuVportDataCpuOper {
	var ret edpt.SystemCpuVportDataCpuOper

	ret.Oper = getObjectSystemCpuVportDataCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
