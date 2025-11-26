package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_topn_oper`: Operational Status for the object topn\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityTopnOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"metric": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"memory_usage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"topn_duration": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"metric_topn_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"object_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"object_val": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"total_memory": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"used_memory": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityTopnOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityTopnOperOper := setObjectVisibilityTopnOperOper(res)
		d.Set("oper", VisibilityTopnOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityTopnOperOper(ret edpt.DataVisibilityTopnOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"class":            ret.DtVisibilityTopnOper.Oper.Class,
			"metric":           ret.DtVisibilityTopnOper.Oper.Metric,
			"memory_usage":     ret.DtVisibilityTopnOper.Oper.MemoryUsage,
			"topn_duration":    ret.DtVisibilityTopnOper.Oper.TopnDuration,
			"metric_topn_list": setSliceVisibilityTopnOperOperMetricTopnList(ret.DtVisibilityTopnOper.Oper.MetricTopnList),
			"total_memory":     ret.DtVisibilityTopnOper.Oper.TotalMemory,
			"used_memory":      ret.DtVisibilityTopnOper.Oper.UsedMemory,
		},
	}
}

func setSliceVisibilityTopnOperOperMetricTopnList(d []edpt.VisibilityTopnOperOperMetricTopnList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["object_name"] = item.ObjectName
		in["object_val"] = item.ObjectVal
		result = append(result, in)
	}
	return result
}

func getObjectVisibilityTopnOperOper(d []interface{}) edpt.VisibilityTopnOperOper {

	count1 := len(d)
	var ret edpt.VisibilityTopnOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Class = in["class"].(string)
		ret.Metric = in["metric"].(string)
		ret.MemoryUsage = in["memory_usage"].(int)
		ret.TopnDuration = in["topn_duration"].(string)
		ret.MetricTopnList = getSliceVisibilityTopnOperOperMetricTopnList(in["metric_topn_list"].([]interface{}))
		ret.TotalMemory = in["total_memory"].(string)
		ret.UsedMemory = in["used_memory"].(string)
	}
	return ret
}

func getSliceVisibilityTopnOperOperMetricTopnList(d []interface{}) []edpt.VisibilityTopnOperOperMetricTopnList {

	count1 := len(d)
	ret := make([]edpt.VisibilityTopnOperOperMetricTopnList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityTopnOperOperMetricTopnList
		oi.ObjectName = in["object_name"].(string)
		oi.ObjectVal = in["object_val"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityTopnOper(d *schema.ResourceData) edpt.VisibilityTopnOper {
	var ret edpt.VisibilityTopnOper

	ret.Oper = getObjectVisibilityTopnOperOper(d.Get("oper").([]interface{}))
	return ret
}
