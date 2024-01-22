package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCpuListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_cpu_list_oper`: Operational Status for the object cpu-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemCpuListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_idx": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cpu_num": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"numa_node": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ht_cpu": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
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

func resourceSystemCpuListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemCpuListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemCpuListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemCpuListOperOper := setObjectSystemCpuListOperOper(res)
		d.Set("oper", SystemCpuListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemCpuListOperOper(ret edpt.DataSystemCpuListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_cpu_list": setSliceSystemCpuListOperOperSystemCpuList(ret.DtSystemCpuListOper.Oper.SystemCpuList),
		},
	}
}

func setSliceSystemCpuListOperOperSystemCpuList(d []edpt.SystemCpuListOperOperSystemCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_idx"] = item.CpuIdx
		in["cpu_num"] = item.CpuNum
		in["numa_node"] = item.NumaNode
		in["ht_cpu"] = item.HtCpu
		in["status"] = item.Status
		result = append(result, in)
	}
	return result
}

func getObjectSystemCpuListOperOper(d []interface{}) edpt.SystemCpuListOperOper {

	count1 := len(d)
	var ret edpt.SystemCpuListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemCpuList = getSliceSystemCpuListOperOperSystemCpuList(in["system_cpu_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemCpuListOperOperSystemCpuList(d []interface{}) []edpt.SystemCpuListOperOperSystemCpuList {

	count1 := len(d)
	ret := make([]edpt.SystemCpuListOperOperSystemCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemCpuListOperOperSystemCpuList
		oi.CpuIdx = in["cpu_idx"].(string)
		oi.CpuNum = in["cpu_num"].(string)
		oi.NumaNode = in["numa_node"].(string)
		oi.HtCpu = in["ht_cpu"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemCpuListOper(d *schema.ResourceData) edpt.SystemCpuListOper {
	var ret edpt.SystemCpuListOper

	ret.Oper = getObjectSystemCpuListOperOper(d.Get("oper").([]interface{}))
	return ret
}
