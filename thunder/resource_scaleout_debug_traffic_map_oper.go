package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugTrafficMapOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_traffic_map_oper`: Operational Status for the object traffic-map\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugTrafficMapOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cmd": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"buffer_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buckets_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_group": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"active_device": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"standby_device": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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

func resourceScaleoutDebugTrafficMapOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugTrafficMapOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugTrafficMapOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugTrafficMapOperOper := setObjectScaleoutDebugTrafficMapOperOper(res)
		d.Set("oper", ScaleoutDebugTrafficMapOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugTrafficMapOperOper(ret edpt.DataScaleoutDebugTrafficMapOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"device_group_list": setSliceScaleoutDebugTrafficMapOperOperDeviceGroupList(ret.DtScaleoutDebugTrafficMapOper.Oper.DeviceGroupList),
		},
	}
}

func setSliceScaleoutDebugTrafficMapOperOperDeviceGroupList(d []edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["rc"] = item.Rc
		in["cmd"] = item.Cmd
		in["buffer_len"] = item.Buffer_len
		in["buckets_list"] = setSliceScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList(item.BucketsList)
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList(d []edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["user_group"] = item.User_group
		in["active_device"] = item.Active_device
		in["standby_device"] = item.Standby_device
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugTrafficMapOperOper(d []interface{}) edpt.ScaleoutDebugTrafficMapOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugTrafficMapOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DeviceGroupList = getSliceScaleoutDebugTrafficMapOperOperDeviceGroupList(in["device_group_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugTrafficMapOperOperDeviceGroupList(d []interface{}) []edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupList
		oi.Rc = in["rc"].(int)
		oi.Cmd = in["cmd"].(string)
		oi.Buffer_len = in["buffer_len"].(int)
		oi.BucketsList = getSliceScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList(in["buckets_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList(d []interface{}) []edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugTrafficMapOperOperDeviceGroupListBucketsList
		oi.User_group = in["user_group"].(int)
		oi.Active_device = in["active_device"].(int)
		oi.Standby_device = in["standby_device"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugTrafficMapOper(d *schema.ResourceData) edpt.ScaleoutDebugTrafficMapOper {
	var ret edpt.ScaleoutDebugTrafficMapOper

	ret.Oper = getObjectScaleoutDebugTrafficMapOperOper(d.Get("oper").([]interface{}))
	return ret
}
