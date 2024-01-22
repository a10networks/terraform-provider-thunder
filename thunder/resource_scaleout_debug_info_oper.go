package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutDebugInfoOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_scaleout_debug_info_oper`: Operational Status for the object info\n\n__PLACEHOLDER__",
		ReadContext: resourceScaleoutDebugInfoOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scaleout_enabled": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"db_process_running": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"g_scaleout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"scaleout_current_role": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"traffic_map_update": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"explicitly_stop_service": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cluster_disc_timer_running": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"pending_scaleout_exit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"perform_tracking_work": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"node_disable_in_prog": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"node_enable_in_prog": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"trigger_cluster_exit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"trigger_enable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"trigger_disable": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cluster_discovery_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cluster_discovery_start_timestamp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"db_operation_max_retry": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"so_device_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"active_device_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceScaleoutDebugInfoOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutDebugInfoOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutDebugInfoOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ScaleoutDebugInfoOperOper := setObjectScaleoutDebugInfoOperOper(res)
		d.Set("oper", ScaleoutDebugInfoOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectScaleoutDebugInfoOperOper(ret edpt.DataScaleoutDebugInfoOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"scaleout_enabled":                  ret.DtScaleoutDebugInfoOper.Oper.Scaleout_enabled,
			"db_process_running":                ret.DtScaleoutDebugInfoOper.Oper.Db_process_running,
			"g_scaleout":                        ret.DtScaleoutDebugInfoOper.Oper.G_scaleout,
			"scaleout_current_role":             ret.DtScaleoutDebugInfoOper.Oper.Scaleout_current_role,
			"traffic_map_update":                ret.DtScaleoutDebugInfoOper.Oper.Traffic_map_update,
			"explicitly_stop_service":           ret.DtScaleoutDebugInfoOper.Oper.Explicitly_stop_service,
			"cluster_disc_timer_running":        ret.DtScaleoutDebugInfoOper.Oper.Cluster_disc_timer_running,
			"pending_scaleout_exit":             ret.DtScaleoutDebugInfoOper.Oper.Pending_scaleout_exit,
			"perform_tracking_work":             ret.DtScaleoutDebugInfoOper.Oper.Perform_tracking_work,
			"node_disable_in_prog":              ret.DtScaleoutDebugInfoOper.Oper.Node_disable_in_prog,
			"node_enable_in_prog":               ret.DtScaleoutDebugInfoOper.Oper.Node_enable_in_prog,
			"trigger_cluster_exit":              ret.DtScaleoutDebugInfoOper.Oper.Trigger_cluster_exit,
			"trigger_enable":                    ret.DtScaleoutDebugInfoOper.Oper.Trigger_enable,
			"trigger_disable":                   ret.DtScaleoutDebugInfoOper.Oper.Trigger_disable,
			"cluster_discovery_timeout":         ret.DtScaleoutDebugInfoOper.Oper.Cluster_discovery_timeout,
			"cluster_discovery_start_timestamp": ret.DtScaleoutDebugInfoOper.Oper.Cluster_discovery_start_timestamp,
			"db_operation_max_retry":            ret.DtScaleoutDebugInfoOper.Oper.Db_operation_max_retry,
			"so_device_count":                   ret.DtScaleoutDebugInfoOper.Oper.So_device_count,
			"active_device_count":               ret.DtScaleoutDebugInfoOper.Oper.Active_device_count,
		},
	}
}

func getObjectScaleoutDebugInfoOperOper(d []interface{}) edpt.ScaleoutDebugInfoOperOper {

	count1 := len(d)
	var ret edpt.ScaleoutDebugInfoOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Scaleout_enabled = in["scaleout_enabled"].(int)
		ret.Db_process_running = in["db_process_running"].(int)
		ret.G_scaleout = in["g_scaleout"].(int)
		ret.Scaleout_current_role = in["scaleout_current_role"].(int)
		ret.Traffic_map_update = in["traffic_map_update"].(int)
		ret.Explicitly_stop_service = in["explicitly_stop_service"].(int)
		ret.Cluster_disc_timer_running = in["cluster_disc_timer_running"].(int)
		ret.Pending_scaleout_exit = in["pending_scaleout_exit"].(int)
		ret.Perform_tracking_work = in["perform_tracking_work"].(int)
		ret.Node_disable_in_prog = in["node_disable_in_prog"].(int)
		ret.Node_enable_in_prog = in["node_enable_in_prog"].(int)
		ret.Trigger_cluster_exit = in["trigger_cluster_exit"].(int)
		ret.Trigger_enable = in["trigger_enable"].(int)
		ret.Trigger_disable = in["trigger_disable"].(int)
		ret.Cluster_discovery_timeout = in["cluster_discovery_timeout"].(int)
		ret.Cluster_discovery_start_timestamp = in["cluster_discovery_start_timestamp"].(int)
		ret.Db_operation_max_retry = in["db_operation_max_retry"].(int)
		ret.So_device_count = in["so_device_count"].(int)
		ret.Active_device_count = in["active_device_count"].(int)
	}
	return ret
}

func dataToEndpointScaleoutDebugInfoOper(d *schema.ResourceData) edpt.ScaleoutDebugInfoOper {
	var ret edpt.ScaleoutDebugInfoOper

	ret.Oper = getObjectScaleoutDebugInfoOperOper(d.Get("oper").([]interface{}))
	return ret
}
