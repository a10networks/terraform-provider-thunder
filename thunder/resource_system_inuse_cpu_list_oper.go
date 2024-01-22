package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemInuseCpuListOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_inuse_cpu_list_oper`: Operational Status for the object inuse-cpu-list\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemInuseCpuListOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_inuse_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_num": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"numa_node": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"cpu_id": {
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

func resourceSystemInuseCpuListOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemInuseCpuListOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemInuseCpuListOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemInuseCpuListOperOper := setObjectSystemInuseCpuListOperOper(res)
		d.Set("oper", SystemInuseCpuListOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemInuseCpuListOperOper(ret edpt.DataSystemInuseCpuListOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"system_inuse_cpu_list": setSliceSystemInuseCpuListOperOperSystemInuseCpuList(ret.DtSystemInuseCpuListOper.Oper.SystemInuseCpuList),
		},
	}
}

func setSliceSystemInuseCpuListOperOperSystemInuseCpuList(d []edpt.SystemInuseCpuListOperOperSystemInuseCpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["cpu_num"] = item.CpuNum
		in["numa_node"] = item.NumaNode
		in["cpu_id"] = item.CpuId
		in["ht_cpu"] = item.HtCpu
		in["status"] = item.Status
		result = append(result, in)
	}
	return result
}

func getObjectSystemInuseCpuListOperOper(d []interface{}) edpt.SystemInuseCpuListOperOper {

	count1 := len(d)
	var ret edpt.SystemInuseCpuListOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SystemInuseCpuList = getSliceSystemInuseCpuListOperOperSystemInuseCpuList(in["system_inuse_cpu_list"].([]interface{}))
	}
	return ret
}

func getSliceSystemInuseCpuListOperOperSystemInuseCpuList(d []interface{}) []edpt.SystemInuseCpuListOperOperSystemInuseCpuList {

	count1 := len(d)
	ret := make([]edpt.SystemInuseCpuListOperOperSystemInuseCpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemInuseCpuListOperOperSystemInuseCpuList
		oi.CpuNum = in["cpu_num"].(string)
		oi.NumaNode = in["numa_node"].(string)
		oi.CpuId = in["cpu_id"].(string)
		oi.HtCpu = in["ht_cpu"].(string)
		oi.Status = in["status"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemInuseCpuListOper(d *schema.ResourceData) edpt.SystemInuseCpuListOper {
	var ret edpt.SystemInuseCpuListOper

	ret.Oper = getObjectSystemInuseCpuListOperOper(d.Get("oper").([]interface{}))
	return ret
}
