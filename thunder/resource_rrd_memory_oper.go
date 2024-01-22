package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRrdMemoryOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rrd_memory_oper`: Operational Status for the object memory\n\n__PLACEHOLDER__",
		ReadContext: resourceRrdMemoryOperRead,

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
						"mem_usage": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"mem_usage": {
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

func resourceRrdMemoryOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRrdMemoryOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRrdMemoryOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RrdMemoryOperOper := setObjectRrdMemoryOperOper(res)
		d.Set("oper", RrdMemoryOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRrdMemoryOperOper(ret edpt.DataRrdMemoryOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_time": ret.DtRrdMemoryOper.Oper.StartTime,
			"end_time":   ret.DtRrdMemoryOper.Oper.EndTime,
			"mem_usage":  setSliceRrdMemoryOperOperMemUsage(ret.DtRrdMemoryOper.Oper.MemUsage),
		},
	}
}

func setSliceRrdMemoryOperOperMemUsage(d []edpt.RrdMemoryOperOperMemUsage) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["time"] = item.Time
		in["mem_usage"] = item.MemUsage
		result = append(result, in)
	}
	return result
}

func getObjectRrdMemoryOperOper(d []interface{}) edpt.RrdMemoryOperOper {

	count1 := len(d)
	var ret edpt.RrdMemoryOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartTime = in["start_time"].(int)
		ret.EndTime = in["end_time"].(int)
		ret.MemUsage = getSliceRrdMemoryOperOperMemUsage(in["mem_usage"].([]interface{}))
	}
	return ret
}

func getSliceRrdMemoryOperOperMemUsage(d []interface{}) []edpt.RrdMemoryOperOperMemUsage {

	count1 := len(d)
	ret := make([]edpt.RrdMemoryOperOperMemUsage, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RrdMemoryOperOperMemUsage
		oi.Time = in["time"].(int)
		oi.MemUsage = in["mem_usage"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRrdMemoryOper(d *schema.ResourceData) edpt.RrdMemoryOper {
	var ret edpt.RrdMemoryOper

	ret.Oper = getObjectRrdMemoryOperOper(d.Get("oper").([]interface{}))
	return ret
}
