package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuMapOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_map_oper`: Operational Status for the object cpu-map\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuMapOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_cpu_map": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_idx": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cpu_type": {
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

func resourceSystemCpuMapOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuMapOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuMapOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCpuMapOperOper := setObjectSystemCpuMapOperOper(res)
		d.Set("oper", SystemCpuMapOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuMapOperOper(ret edpt.DataSystemCpuMapOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_cpu_map": setSliceSystemCpuMapOperOperSystemCpuMap(ret.DtSystemCpuMapOper.Oper.SystemCpuMap),
		},
	}
}

func setSliceSystemCpuMapOperOperSystemCpuMap(d []edpt.SystemCpuMapOperOperSystemCpuMap) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_idx"] = item.CpuIdx
		in["cpu_type"] = item.CpuType
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuMapOperOper(d []interface{}) edpt.SystemCpuMapOperOper {

	count1 := len(d)
	var ret edpt.SystemCpuMapOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemCpuMap = getSliceSystemCpuMapOperOperSystemCpuMap(in["system_cpu_map"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuMapOperOperSystemCpuMap(d []interface{}) []edpt.SystemCpuMapOperOperSystemCpuMap {

	count1 := len(d)
	ret := make([]edpt.SystemCpuMapOperOperSystemCpuMap, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuMapOperOperSystemCpuMap
		oi.CpuIdx = in["cpu_idx"].(string)
		oi.CpuType = in["cpu_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuMapOper(d *schema.ResourceData) edpt.SystemCpuMapOper {
	var ret edpt.SystemCpuMapOper

	ret.Oper = getObjectSystemCpuMapOperOper(d.Get("oper").([]interface{}))
	return ret
}
