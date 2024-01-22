package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosResourceTrackingCpuOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_resource_tracking_cpu_oper`: Operational Status for the object cpu\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosResourceTrackingCpuOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"if_show_error_str": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"error_str": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"timestamps": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timestamp": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"average_cpu_percent": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"entries": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"entry": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"address": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"relative_cpu_percent": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
												"absolute_cpu_percent": {
													Type: schema.TypeString, Optional: true, Description: "",
												},
											},
										},
									},
								},
							},
						},
						"max_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceDdosResourceTrackingCpuOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosResourceTrackingCpuOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosResourceTrackingCpuOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosResourceTrackingCpuOperOper := setObjectDdosResourceTrackingCpuOperOper(res)
		d.Set("oper", DdosResourceTrackingCpuOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosResourceTrackingCpuOperOper(ret edpt.DataDdosResourceTrackingCpuOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"if_show_error_str": ret.DtDdosResourceTrackingCpuOper.Oper.IfShowErrorStr,
			"error_str":         ret.DtDdosResourceTrackingCpuOper.Oper.ErrorStr,
			"timestamps":        setSliceDdosResourceTrackingCpuOperOperTimestamps(ret.DtDdosResourceTrackingCpuOper.Oper.Timestamps),
			"max_count":         ret.DtDdosResourceTrackingCpuOper.Oper.MaxCount,
		},
	}
}

func setSliceDdosResourceTrackingCpuOperOperTimestamps(d []edpt.DdosResourceTrackingCpuOperOperTimestamps) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["timestamp"] = item.Timestamp
		in["average_cpu_percent"] = item.AverageCpuPercent
		in["entries"] = setSliceDdosResourceTrackingCpuOperOperTimestampsEntries(item.Entries)
		result = append(result, in)
	}
	return result
}

func setSliceDdosResourceTrackingCpuOperOperTimestampsEntries(d []edpt.DdosResourceTrackingCpuOperOperTimestampsEntries) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["entry"] = item.Entry
		in["address"] = item.Address
		in["relative_cpu_percent"] = item.RelativeCpuPercent
		in["absolute_cpu_percent"] = item.AbsoluteCpuPercent
		result = append(result, in)
	}
	return result
}

func getObjectDdosResourceTrackingCpuOperOper(d []interface{}) edpt.DdosResourceTrackingCpuOperOper {

	count1 := len(d)
	var ret edpt.DdosResourceTrackingCpuOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IfShowErrorStr = in["if_show_error_str"].(int)
		ret.ErrorStr = in["error_str"].(string)
		ret.Timestamps = getSliceDdosResourceTrackingCpuOperOperTimestamps(in["timestamps"].([]interface{}))
		ret.MaxCount = in["max_count"].(int)
	}
	return ret
}

func getSliceDdosResourceTrackingCpuOperOperTimestamps(d []interface{}) []edpt.DdosResourceTrackingCpuOperOperTimestamps {

	count1 := len(d)
	ret := make([]edpt.DdosResourceTrackingCpuOperOperTimestamps, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosResourceTrackingCpuOperOperTimestamps
		oi.Timestamp = in["timestamp"].(string)
		oi.AverageCpuPercent = in["average_cpu_percent"].(string)
		oi.Entries = getSliceDdosResourceTrackingCpuOperOperTimestampsEntries(in["entries"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosResourceTrackingCpuOperOperTimestampsEntries(d []interface{}) []edpt.DdosResourceTrackingCpuOperOperTimestampsEntries {

	count1 := len(d)
	ret := make([]edpt.DdosResourceTrackingCpuOperOperTimestampsEntries, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosResourceTrackingCpuOperOperTimestampsEntries
		oi.Entry = in["entry"].(string)
		oi.Address = in["address"].(string)
		oi.RelativeCpuPercent = in["relative_cpu_percent"].(string)
		oi.AbsoluteCpuPercent = in["absolute_cpu_percent"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosResourceTrackingCpuOper(d *schema.ResourceData) edpt.DdosResourceTrackingCpuOper {
	var ret edpt.DdosResourceTrackingCpuOper

	ret.Oper = getObjectDdosResourceTrackingCpuOperOper(d.Get("oper").([]interface{}))
	return ret
}
