package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_proxy_stats`: Statistics for the object proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_event": {
							Type: schema.TypeInt, Optional: true, Description: "TCP stack event",
						},
						"est_event": {
							Type: schema.TypeInt, Optional: true, Description: "Connection established",
						},
						"data_event": {
							Type: schema.TypeInt, Optional: true, Description: "Data received",
						},
						"client_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Client FIN",
						},
						"server_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Server FIN",
						},
						"wbuf_event": {
							Type: schema.TypeInt, Optional: true, Description: "Ready to send data",
						},
						"err_event": {
							Type: schema.TypeInt, Optional: true, Description: "Error occured",
						},
						"no_mem": {
							Type: schema.TypeInt, Optional: true, Description: "No memory",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Client RST",
						},
						"server_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Server RST",
						},
						"queue_depth_over_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Queue depth over limit",
						},
						"event_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Event failed",
						},
						"conn_not_exist": {
							Type: schema.TypeInt, Optional: true, Description: "Conn not exist",
						},
						"service_alloc_cb": {
							Type: schema.TypeInt, Optional: true, Description: "Service alloc callback",
						},
						"service_alloc_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Service alloc callback failed",
						},
						"service_free_cb": {
							Type: schema.TypeInt, Optional: true, Description: "Service free callback",
						},
						"service_free_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Service free callback failed",
						},
						"est_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "App EST callback failed",
						},
						"data_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "App DATA callback failed",
						},
						"wbuf_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "App WBUF callback failed",
						},
						"err_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "App ERR callback failed",
						},
						"start_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Start server conn",
						},
						"start_server_conn_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Success",
						},
						"start_server_conn_no_route": {
							Type: schema.TypeInt, Optional: true, Description: "No route to server",
						},
						"start_server_conn_fail_mem": {
							Type: schema.TypeInt, Optional: true, Description: "No memory",
						},
						"start_server_conn_fail_snat": {
							Type: schema.TypeInt, Optional: true, Description: "Failed Source NAT",
						},
						"start_server_conn_fail_persist": {
							Type: schema.TypeInt, Optional: true, Description: "Fail Persistence",
						},
						"start_server_conn_fail_server": {
							Type: schema.TypeInt, Optional: true, Description: "Fail Server issue",
						},
						"start_server_conn_fail_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Fail Tuple Issue",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Line too long",
						},
					},
				},
			},
		},
	}
}

func resourceSlbProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbProxyStatsStats := setObjectSlbProxyStatsStats(res)
		d.Set("stats", SlbProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbProxyStatsStats(ret edpt.DataSlbProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_event":                      ret.DtSlbProxyStats.Stats.Tcp_event,
			"est_event":                      ret.DtSlbProxyStats.Stats.Est_event,
			"data_event":                     ret.DtSlbProxyStats.Stats.Data_event,
			"client_fin":                     ret.DtSlbProxyStats.Stats.Client_fin,
			"server_fin":                     ret.DtSlbProxyStats.Stats.Server_fin,
			"wbuf_event":                     ret.DtSlbProxyStats.Stats.Wbuf_event,
			"err_event":                      ret.DtSlbProxyStats.Stats.Err_event,
			"no_mem":                         ret.DtSlbProxyStats.Stats.No_mem,
			"client_rst":                     ret.DtSlbProxyStats.Stats.Client_rst,
			"server_rst":                     ret.DtSlbProxyStats.Stats.Server_rst,
			"queue_depth_over_limit":         ret.DtSlbProxyStats.Stats.Queue_depth_over_limit,
			"event_failed":                   ret.DtSlbProxyStats.Stats.Event_failed,
			"conn_not_exist":                 ret.DtSlbProxyStats.Stats.Conn_not_exist,
			"service_alloc_cb":               ret.DtSlbProxyStats.Stats.Service_alloc_cb,
			"service_alloc_cb_failed":        ret.DtSlbProxyStats.Stats.Service_alloc_cb_failed,
			"service_free_cb":                ret.DtSlbProxyStats.Stats.Service_free_cb,
			"service_free_cb_failed":         ret.DtSlbProxyStats.Stats.Service_free_cb_failed,
			"est_cb_failed":                  ret.DtSlbProxyStats.Stats.Est_cb_failed,
			"data_cb_failed":                 ret.DtSlbProxyStats.Stats.Data_cb_failed,
			"wbuf_cb_failed":                 ret.DtSlbProxyStats.Stats.Wbuf_cb_failed,
			"err_cb_failed":                  ret.DtSlbProxyStats.Stats.Err_cb_failed,
			"start_server_conn":              ret.DtSlbProxyStats.Stats.Start_server_conn,
			"start_server_conn_succ":         ret.DtSlbProxyStats.Stats.Start_server_conn_succ,
			"start_server_conn_no_route":     ret.DtSlbProxyStats.Stats.Start_server_conn_no_route,
			"start_server_conn_fail_mem":     ret.DtSlbProxyStats.Stats.Start_server_conn_fail_mem,
			"start_server_conn_fail_snat":    ret.DtSlbProxyStats.Stats.Start_server_conn_fail_snat,
			"start_server_conn_fail_persist": ret.DtSlbProxyStats.Stats.Start_server_conn_fail_persist,
			"start_server_conn_fail_server":  ret.DtSlbProxyStats.Stats.Start_server_conn_fail_server,
			"start_server_conn_fail_tuple":   ret.DtSlbProxyStats.Stats.Start_server_conn_fail_tuple,
			"line_too_long":                  ret.DtSlbProxyStats.Stats.Line_too_long,
		},
	}
}

func getObjectSlbProxyStatsStats(d []interface{}) edpt.SlbProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp_event = in["tcp_event"].(int)
		ret.Est_event = in["est_event"].(int)
		ret.Data_event = in["data_event"].(int)
		ret.Client_fin = in["client_fin"].(int)
		ret.Server_fin = in["server_fin"].(int)
		ret.Wbuf_event = in["wbuf_event"].(int)
		ret.Err_event = in["err_event"].(int)
		ret.No_mem = in["no_mem"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Server_rst = in["server_rst"].(int)
		ret.Queue_depth_over_limit = in["queue_depth_over_limit"].(int)
		ret.Event_failed = in["event_failed"].(int)
		ret.Conn_not_exist = in["conn_not_exist"].(int)
		ret.Service_alloc_cb = in["service_alloc_cb"].(int)
		ret.Service_alloc_cb_failed = in["service_alloc_cb_failed"].(int)
		ret.Service_free_cb = in["service_free_cb"].(int)
		ret.Service_free_cb_failed = in["service_free_cb_failed"].(int)
		ret.Est_cb_failed = in["est_cb_failed"].(int)
		ret.Data_cb_failed = in["data_cb_failed"].(int)
		ret.Wbuf_cb_failed = in["wbuf_cb_failed"].(int)
		ret.Err_cb_failed = in["err_cb_failed"].(int)
		ret.Start_server_conn = in["start_server_conn"].(int)
		ret.Start_server_conn_succ = in["start_server_conn_succ"].(int)
		ret.Start_server_conn_no_route = in["start_server_conn_no_route"].(int)
		ret.Start_server_conn_fail_mem = in["start_server_conn_fail_mem"].(int)
		ret.Start_server_conn_fail_snat = in["start_server_conn_fail_snat"].(int)
		ret.Start_server_conn_fail_persist = in["start_server_conn_fail_persist"].(int)
		ret.Start_server_conn_fail_server = in["start_server_conn_fail_server"].(int)
		ret.Start_server_conn_fail_tuple = in["start_server_conn_fail_tuple"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
	}
	return ret
}

func dataToEndpointSlbProxyStats(d *schema.ResourceData) edpt.SlbProxyStats {
	var ret edpt.SlbProxyStats

	ret.Stats = getObjectSlbProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
