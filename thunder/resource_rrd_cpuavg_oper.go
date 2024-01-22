package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdCpuavgOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_cpuavg_oper`: Operational Status for the object cpuavg\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdCpuavgOperRead,

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
									"cpu_num": {
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

func resourceRrdCpuavgOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdCpuavgOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdCpuavgOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdCpuavgOperOper := setObjectRrdCpuavgOperOper(res)
		d.Set("oper", RrdCpuavgOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdCpuavgOperOper(ret edpt.DataRrdCpuavgOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time": ret.DtRrdCpuavgOper.Oper.StartTime,
			"end_time":   ret.DtRrdCpuavgOper.Oper.EndTime,
			"cpu_usage":  setSliceRrdCpuavgOperOperCpuUsage(ret.DtRrdCpuavgOper.Oper.CpuUsage),
		},
	}
}

func setSliceRrdCpuavgOperOperCpuUsage(d []edpt.RrdCpuavgOperOperCpuUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["cpu_num"] = item.CpuNum
		in["cpu_type"] = item.CpuType
		in["cpu_usage"] = item.CpuUsage
		result = append(result, in)
	}
	return result
}

func getObjectRrdCpuavgOperOper(d []interface{}) edpt.RrdCpuavgOperOper {

	count1 := len(d)
	var ret edpt.RrdCpuavgOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.CpuUsage = getSliceRrdCpuavgOperOperCpuUsage(in["cpu_usage"].([]interface{}))
	}
	return ret
}

func getSliceRrdCpuavgOperOperCpuUsage(d []interface{}) []edpt.RrdCpuavgOperOperCpuUsage {

	count1 := len(d)
	ret := make([]edpt.RrdCpuavgOperOperCpuUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdCpuavgOperOperCpuUsage
		oi.Time = in["time"].(int)
		oi.CpuNum = in["cpu_num"].(int)
		oi.CpuType = in["cpu_type"].(int)
		oi.CpuUsage = in["cpu_usage"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdCpuavgOper(d *schema.ResourceData) edpt.RrdCpuavgOper {
	var ret edpt.RrdCpuavgOper

	ret.Oper = getObjectRrdCpuavgOperOper(d.Get("oper").([]interface{}))
	return ret
}
