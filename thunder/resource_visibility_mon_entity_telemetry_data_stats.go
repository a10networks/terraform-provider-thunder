package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonEntityTelemetryDataStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_mon_entity_telemetry_data_stats`: Statistics for the object mon-entity-telemetry-data\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonEntityTelemetryDataStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"in_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN pkts",
						},
						"out_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT pkts",
						},
						"in_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN bytes",
						},
						"out_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT bytes",
						},
						"errors": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric ERRORS",
						},
						"in_small_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN SMALL pkt",
						},
						"in_frag": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN frag",
						},
						"out_small_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT SMALL pkt",
						},
						"out_frag": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric OUT frag",
						},
						"new_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric New Sessions",
						},
						"avg_data_cpu_util": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric Average data CPU utilization",
						},
						"outside_intf_util": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric Outside interface utilization",
						},
						"concurrent_conn": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"in_bytes_per_out_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric IN bytes per OUT bytes",
						},
						"drop_pkts_per_pkts": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric Drop pkts per pkts",
						},
						"tcp_in_syn": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN syn",
						},
						"tcp_out_syn": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT syn",
						},
						"tcp_in_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN fin",
						},
						"tcp_out_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT fin",
						},
						"tcp_in_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN payload",
						},
						"tcp_out_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT payload",
						},
						"tcp_in_rexmit": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN rexmit",
						},
						"tcp_out_rexmit": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT rexmit",
						},
						"tcp_in_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN rst",
						},
						"tcp_out_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT rst",
						},
						"tcp_in_empty_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP_IN EMPTY ack",
						},
						"tcp_out_empty_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT EMPTY ack",
						},
						"tcp_in_zero_wnd": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP IN ZERO wnd",
						},
						"tcp_out_zero_wnd": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP OUT ZERO wnd",
						},
						"tcp_conn_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP connection miss",
						},
						"tcp_fwd_syn_per_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Monitored Entity telemetry Metric TCP FWD SYN per FIN",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityMonEntityTelemetryDataStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonEntityTelemetryDataStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonEntityTelemetryDataStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonEntityTelemetryDataStatsStats := setObjectVisibilityMonEntityTelemetryDataStatsStats(res)
		d.Set("stats", VisibilityMonEntityTelemetryDataStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonEntityTelemetryDataStatsStats(ret edpt.DataVisibilityMonEntityTelemetryDataStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"in_pkts":                ret.DtVisibilityMonEntityTelemetryDataStats.Stats.In_pkts,
			"out_pkts":               ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Out_pkts,
			"in_bytes":               ret.DtVisibilityMonEntityTelemetryDataStats.Stats.In_bytes,
			"out_bytes":              ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Out_bytes,
			"errors":                 ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Errors,
			"in_small_pkt":           ret.DtVisibilityMonEntityTelemetryDataStats.Stats.In_small_pkt,
			"in_frag":                ret.DtVisibilityMonEntityTelemetryDataStats.Stats.In_frag,
			"out_small_pkt":          ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Out_small_pkt,
			"out_frag":               ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Out_frag,
			"new_conn":               ret.DtVisibilityMonEntityTelemetryDataStats.Stats.NewConn,
			"avg_data_cpu_util":      ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Avg_data_cpu_util,
			"outside_intf_util":      ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Outside_intf_util,
			"concurrent_conn":        ret.DtVisibilityMonEntityTelemetryDataStats.Stats.ConcurrentConn,
			"in_bytes_per_out_bytes": ret.DtVisibilityMonEntityTelemetryDataStats.Stats.In_bytes_per_out_bytes,
			"drop_pkts_per_pkts":     ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Drop_pkts_per_pkts,
			"tcp_in_syn":             ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_syn,
			"tcp_out_syn":            ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_syn,
			"tcp_in_fin":             ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_fin,
			"tcp_out_fin":            ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_fin,
			"tcp_in_payload":         ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_payload,
			"tcp_out_payload":        ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_payload,
			"tcp_in_rexmit":          ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_rexmit,
			"tcp_out_rexmit":         ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_rexmit,
			"tcp_in_rst":             ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_rst,
			"tcp_out_rst":            ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_rst,
			"tcp_in_empty_ack":       ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_empty_ack,
			"tcp_out_empty_ack":      ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_empty_ack,
			"tcp_in_zero_wnd":        ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_in_zero_wnd,
			"tcp_out_zero_wnd":       ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_out_zero_wnd,
			"tcp_conn_miss":          ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_conn_miss,
			"tcp_fwd_syn_per_fin":    ret.DtVisibilityMonEntityTelemetryDataStats.Stats.Tcp_fwd_syn_per_fin,
		},
	}
}

func getObjectVisibilityMonEntityTelemetryDataStatsStats(d []interface{}) edpt.VisibilityMonEntityTelemetryDataStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityMonEntityTelemetryDataStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.In_pkts = in["in_pkts"].(int)
		ret.Out_pkts = in["out_pkts"].(int)
		ret.In_bytes = in["in_bytes"].(int)
		ret.Out_bytes = in["out_bytes"].(int)
		ret.Errors = in["errors"].(int)
		ret.In_small_pkt = in["in_small_pkt"].(int)
		ret.In_frag = in["in_frag"].(int)
		ret.Out_small_pkt = in["out_small_pkt"].(int)
		ret.Out_frag = in["out_frag"].(int)
		ret.NewConn = in["new_conn"].(int)
		ret.Avg_data_cpu_util = in["avg_data_cpu_util"].(int)
		ret.Outside_intf_util = in["outside_intf_util"].(int)
		ret.ConcurrentConn = in["concurrent_conn"].(int)
		ret.In_bytes_per_out_bytes = in["in_bytes_per_out_bytes"].(int)
		ret.Drop_pkts_per_pkts = in["drop_pkts_per_pkts"].(int)
		ret.Tcp_in_syn = in["tcp_in_syn"].(int)
		ret.Tcp_out_syn = in["tcp_out_syn"].(int)
		ret.Tcp_in_fin = in["tcp_in_fin"].(int)
		ret.Tcp_out_fin = in["tcp_out_fin"].(int)
		ret.Tcp_in_payload = in["tcp_in_payload"].(int)
		ret.Tcp_out_payload = in["tcp_out_payload"].(int)
		ret.Tcp_in_rexmit = in["tcp_in_rexmit"].(int)
		ret.Tcp_out_rexmit = in["tcp_out_rexmit"].(int)
		ret.Tcp_in_rst = in["tcp_in_rst"].(int)
		ret.Tcp_out_rst = in["tcp_out_rst"].(int)
		ret.Tcp_in_empty_ack = in["tcp_in_empty_ack"].(int)
		ret.Tcp_out_empty_ack = in["tcp_out_empty_ack"].(int)
		ret.Tcp_in_zero_wnd = in["tcp_in_zero_wnd"].(int)
		ret.Tcp_out_zero_wnd = in["tcp_out_zero_wnd"].(int)
		ret.Tcp_conn_miss = in["tcp_conn_miss"].(int)
		ret.Tcp_fwd_syn_per_fin = in["tcp_fwd_syn_per_fin"].(int)
	}
	return ret
}

func dataToEndpointVisibilityMonEntityTelemetryDataStats(d *schema.ResourceData) edpt.VisibilityMonEntityTelemetryDataStats {
	var ret edpt.VisibilityMonEntityTelemetryDataStats

	ret.Stats = getObjectVisibilityMonEntityTelemetryDataStatsStats(d.Get("stats").([]interface{}))
	return ret
}
