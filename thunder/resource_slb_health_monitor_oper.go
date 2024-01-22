package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthMonitorOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_monitor_oper`: Operational Status for the object health-monitor\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthMonitorOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"health_monitor_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"interval": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"retries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"up_retries": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"method": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"ssl_refresh": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pin_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pin_process_index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"all_partitions": {
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

func resourceSlbHealthMonitorOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthMonitorOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthMonitorOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthMonitorOperOper := setObjectSlbHealthMonitorOperOper(res)
		d.Set("oper", SlbHealthMonitorOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthMonitorOperOper(ret edpt.DataSlbHealthMonitorOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"health_monitor_list": setSliceSlbHealthMonitorOperOperHealthMonitorList(ret.DtSlbHealthMonitorOper.Oper.HealthMonitorList),
		},
	}
}

func setSliceSlbHealthMonitorOperOperHealthMonitorList(d []edpt.SlbHealthMonitorOperOperHealthMonitorList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["interval"] = item.Interval
		in["retries"] = item.Retries
		in["timeout"] = item.Timeout
		in["up_retries"] = item.UpRetries
		in["method"] = item.Method
		in["status"] = item.Status
		in["partition"] = item.Partition
		in["ssl_refresh"] = item.SslRefresh
		in["pin_id"] = item.PinId
		in["pin_process_index"] = item.PinProcessIndex
		in["all_partitions"] = item.AllPartitions
		result = append(result, in)
	}
	return result
}

func getObjectSlbHealthMonitorOperOper(d []interface{}) edpt.SlbHealthMonitorOperOper {

	count1 := len(d)
	var ret edpt.SlbHealthMonitorOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HealthMonitorList = getSliceSlbHealthMonitorOperOperHealthMonitorList(in["health_monitor_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbHealthMonitorOperOperHealthMonitorList(d []interface{}) []edpt.SlbHealthMonitorOperOperHealthMonitorList {

	count1 := len(d)
	ret := make([]edpt.SlbHealthMonitorOperOperHealthMonitorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHealthMonitorOperOperHealthMonitorList
		oi.Name = in["name"].(string)
		oi.Interval = in["interval"].(int)
		oi.Retries = in["retries"].(int)
		oi.Timeout = in["timeout"].(int)
		oi.UpRetries = in["up_retries"].(int)
		oi.Method = in["method"].(string)
		oi.Status = in["status"].(string)
		oi.Partition = in["partition"].(string)
		oi.SslRefresh = in["ssl_refresh"].(int)
		oi.PinId = in["pin_id"].(int)
		oi.PinProcessIndex = in["pin_process_index"].(int)
		oi.AllPartitions = in["all_partitions"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHealthMonitorOper(d *schema.ResourceData) edpt.SlbHealthMonitorOper {
	var ret edpt.SlbHealthMonitorOper

	ret.Oper = getObjectSlbHealthMonitorOperOper(d.Get("oper").([]interface{}))
	return ret
}
