package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugTrafficCaptureOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_debug_traffic_capture_oper`: Operational Status for the object debug-traffic-capture\n\n__PLACEHOLDER__",
		ReadContext: resourceDebugTrafficCaptureOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_info": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"status": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end_time": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"end_reason": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"pkt_count": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pkt_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"disk_size": {
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

func resourceDebugTrafficCaptureOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugTrafficCaptureOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugTrafficCaptureOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DebugTrafficCaptureOperOper := setObjectDebugTrafficCaptureOperOper(res)
		d.Set("oper", DebugTrafficCaptureOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDebugTrafficCaptureOperOper(ret edpt.DataDebugTrafficCaptureOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"status_info": setSliceDebugTrafficCaptureOperOperStatusInfo(ret.DtDebugTrafficCaptureOper.Oper.StatusInfo),
		},
	}
}

func setSliceDebugTrafficCaptureOperOperStatusInfo(d []edpt.DebugTrafficCaptureOperOperStatusInfo) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["status"] = item.Status
		in["end_time"] = item.End_time
		in["end_reason"] = item.End_reason
		in["pkt_count"] = item.Pkt_count
		in["pkt_dropped"] = item.Pkt_dropped
		in["disk_size"] = item.Disk_size
		result = append(result, in)
	}
	return result
}

func getObjectDebugTrafficCaptureOperOper(d []interface{}) edpt.DebugTrafficCaptureOperOper {

	count1 := len(d)
	var ret edpt.DebugTrafficCaptureOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatusInfo = getSliceDebugTrafficCaptureOperOperStatusInfo(in["status_info"].([]interface{}))
	}
	return ret
}

func getSliceDebugTrafficCaptureOperOperStatusInfo(d []interface{}) []edpt.DebugTrafficCaptureOperOperStatusInfo {

	count1 := len(d)
	ret := make([]edpt.DebugTrafficCaptureOperOperStatusInfo, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DebugTrafficCaptureOperOperStatusInfo
		oi.Name = in["name"].(string)
		oi.Status = in["status"].(string)
		oi.End_time = in["end_time"].(string)
		oi.End_reason = in["end_reason"].(string)
		oi.Pkt_count = in["pkt_count"].(int)
		oi.Pkt_dropped = in["pkt_dropped"].(int)
		oi.Disk_size = in["disk_size"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDebugTrafficCaptureOper(d *schema.ResourceData) edpt.DebugTrafficCaptureOper {
	var ret edpt.DebugTrafficCaptureOper

	ret.Oper = getObjectDebugTrafficCaptureOperOper(d.Get("oper").([]interface{}))
	return ret
}
