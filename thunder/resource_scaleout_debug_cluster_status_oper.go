package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugClusterStatusOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_cluster_status_oper`: Operational Status for the object status\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugClusterStatusOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scaleout_device_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"scaleout_priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"scaleout_current_role": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"scaleout_enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"g_scaleout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"min_node_num": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cluster_disc_timer_running": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"explicitly_stop_service": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"maintain_mode_configured": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"device_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vcs_device_id": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"vcs_priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"follow_vcs": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"active_device_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"active_device_count_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"scaleout_device_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"scaleout_device_group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"index": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"members": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"members_low64": {
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

func resourceScaleoutDebugClusterStatusOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugClusterStatusOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugClusterStatusOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugClusterStatusOperOper := setObjectScaleoutDebugClusterStatusOperOper(res)
		d.Set("oper", ScaleoutDebugClusterStatusOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugClusterStatusOperOper(ret edpt.DataScaleoutDebugClusterStatusOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scaleout_device_id":         ret.DtScaleoutDebugClusterStatusOper.Oper.Scaleout_device_id,
			"scaleout_priority":          ret.DtScaleoutDebugClusterStatusOper.Oper.Scaleout_priority,
			"scaleout_current_role":      ret.DtScaleoutDebugClusterStatusOper.Oper.Scaleout_current_role,
			"scaleout_enabled":           ret.DtScaleoutDebugClusterStatusOper.Oper.Scaleout_enabled,
			"g_scaleout":                 ret.DtScaleoutDebugClusterStatusOper.Oper.G_scaleout,
			"min_node_num":               ret.DtScaleoutDebugClusterStatusOper.Oper.Min_node_num,
			"cluster_disc_timer_running": ret.DtScaleoutDebugClusterStatusOper.Oper.Cluster_disc_timer_running,
			"explicitly_stop_service":    ret.DtScaleoutDebugClusterStatusOper.Oper.Explicitly_stop_service,
			"maintain_mode_configured":   ret.DtScaleoutDebugClusterStatusOper.Oper.Maintain_mode_configured,
			"device_id":                  ret.DtScaleoutDebugClusterStatusOper.Oper.Device_id,
			"priority":                   ret.DtScaleoutDebugClusterStatusOper.Oper.Priority,
			"vcs_device_id":              ret.DtScaleoutDebugClusterStatusOper.Oper.Vcs_device_id,
			"vcs_priority":               ret.DtScaleoutDebugClusterStatusOper.Oper.Vcs_priority,
			"follow_vcs":                 ret.DtScaleoutDebugClusterStatusOper.Oper.Follow_vcs,
			"active_device_count":        ret.DtScaleoutDebugClusterStatusOper.Oper.Active_device_count,
			"active_device_count_list":   setSliceScaleoutDebugClusterStatusOperOperActive_device_count_list(ret.DtScaleoutDebugClusterStatusOper.Oper.Active_device_count_list),
			"scaleout_device_list":       setSliceScaleoutDebugClusterStatusOperOperScaleout_device_list(ret.DtScaleoutDebugClusterStatusOper.Oper.Scaleout_device_list),
			"scaleout_device_group_list": setSliceScaleoutDebugClusterStatusOperOperScaleout_device_group_list(ret.DtScaleoutDebugClusterStatusOper.Oper.Scaleout_device_group_list),
		},
	}
}

func setSliceScaleoutDebugClusterStatusOperOperActive_device_count_list(d []edpt.ScaleoutDebugClusterStatusOperOperActive_device_count_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.Device_id
		in["state"] = item.State
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugClusterStatusOperOperScaleout_device_list(d []edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["device_id"] = item.Device_id
		in["state"] = item.State
		result = append(result, in)
	}
	return result
}

func setSliceScaleoutDebugClusterStatusOperOperScaleout_device_group_list(d []edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_group_list) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["index"] = item.Index
		in["members"] = item.Members
		in["members_low64"] = item.Members_low64
		result = append(result, in)
	}
	return result
}

func getObjectScaleoutDebugClusterStatusOperOper(d []interface{}) edpt.ScaleoutDebugClusterStatusOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugClusterStatusOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Scaleout_device_id = in["scaleout_device_id"].(int)
		ret.Scaleout_priority = in["scaleout_priority"].(int)
		ret.Scaleout_current_role = in["scaleout_current_role"].(int)
		ret.Scaleout_enabled = in["scaleout_enabled"].(int)
		ret.G_scaleout = in["g_scaleout"].(int)
		ret.Min_node_num = in["min_node_num"].(int)
		ret.Cluster_disc_timer_running = in["cluster_disc_timer_running"].(int)
		ret.Explicitly_stop_service = in["explicitly_stop_service"].(int)
		ret.Maintain_mode_configured = in["maintain_mode_configured"].(int)
		ret.Device_id = in["device_id"].(int)
		ret.Priority = in["priority"].(int)
		ret.Vcs_device_id = in["vcs_device_id"].(int)
		ret.Vcs_priority = in["vcs_priority"].(int)
		ret.Follow_vcs = in["follow_vcs"].(int)
		ret.Active_device_count = in["active_device_count"].(int)
		ret.Active_device_count_list = getSliceScaleoutDebugClusterStatusOperOperActive_device_count_list(in["active_device_count_list"].([]interface{}))
		ret.Scaleout_device_list = getSliceScaleoutDebugClusterStatusOperOperScaleout_device_list(in["scaleout_device_list"].([]interface{}))
		ret.Scaleout_device_group_list = getSliceScaleoutDebugClusterStatusOperOperScaleout_device_group_list(in["scaleout_device_group_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutDebugClusterStatusOperOperActive_device_count_list(d []interface{}) []edpt.ScaleoutDebugClusterStatusOperOperActive_device_count_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugClusterStatusOperOperActive_device_count_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugClusterStatusOperOperActive_device_count_list
		oi.Device_id = in["device_id"].(int)
		oi.State = in["state"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugClusterStatusOperOperScaleout_device_list(d []interface{}) []edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_list
		oi.Device_id = in["device_id"].(int)
		oi.State = in["state"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutDebugClusterStatusOperOperScaleout_device_group_list(d []interface{}) []edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_group_list {

	count1 := len(d)
	ret := make([]edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_group_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutDebugClusterStatusOperOperScaleout_device_group_list
		oi.Index = in["index"].(int)
		oi.Members = in["members"].(int)
		oi.Members_low64 = in["members_low64"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutDebugClusterStatusOper(d *schema.ResourceData) edpt.ScaleoutDebugClusterStatusOper {
	var ret edpt.ScaleoutDebugClusterStatusOper

	ret.Oper = getObjectScaleoutDebugClusterStatusOperOper(d.Get("oper").([]interface{}))
	return ret
}
