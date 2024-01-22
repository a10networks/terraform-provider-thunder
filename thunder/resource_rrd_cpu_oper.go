package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_cpu_oper`: Operational Status for the object cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdCpuOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"end_time": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cpu_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cpu_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cpu_type": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cpu_usage": {
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

func resourceRrdCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdCpuOperOper := setObjectRrdCpuOperOper(res)
		d.Set("oper", RrdCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdCpuOperOper(ret edpt.DataRrdCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time": ret.DtRrdCpuOper.Oper.StartTime,
			"end_time":   ret.DtRrdCpuOper.Oper.EndTime,
			"cpu_usage":  setSliceRrdCpuOperOperCpuUsage(ret.DtRrdCpuOper.Oper.CpuUsage),
		},
	}
}

func setSliceRrdCpuOperOperCpuUsage(d []edpt.RrdCpuOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["cpu_id"] = item.CpuId
		in["cpu_type"] = item.CpuType
		in["cpu_usage"] = item.CpuUsage
		result = append(result, in)
	}
	return result
}

func getObjectRrdCpuOperOper(d []interface{}) edpt.RrdCpuOperOper {

	count1 := len(d)
	var ret edpt.RrdCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.CpuUsage = getSliceRrdCpuOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceRrdCpuOperOperCpuUsage(d []interface{}) []edpt.RrdCpuOperOperCpuUsage {

	count1 := len(d)
	ret := make([]edpt.RrdCpuOperOperCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdCpuOperOperCpuUsage
		oi.Time = in["time"].(int)
		oi.CpuId = in["cpu_id"].(int)
		oi.CpuType = in["cpu_type"].(int)
		oi.CpuUsage = in["cpu_usage"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdCpuOper(d *schema.ResourceData) edpt.RrdCpuOper {
	var ret edpt.RrdCpuOper

	ret.Oper = getObjectRrdCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
