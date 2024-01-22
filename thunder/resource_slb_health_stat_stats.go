package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHealthStatStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_health_stat_stats`: Statistics for the object health-stat\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHealthStatStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"num_burst": {
							Type: schema.TypeInt, Optional: true, Description: "Number of burst",
						},
						"max_jiffie": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum number of jiffies",
						},
						"min_jiffie": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum number of jiffies",
						},
						"avg_jiffie": {
							Type: schema.TypeInt, Optional: true, Description: "Average number of jiffies",
						},
						"open_socket": {
							Type: schema.TypeInt, Optional: true, Description: "Number of open sockets",
						},
						"open_socket_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Number of failed open sockets",
						},
						"close_socket": {
							Type: schema.TypeInt, Optional: true, Description: "Number of closed sockets",
						},
						"connect_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Number of failed connections",
						},
						"send_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packets sent",
						},
						"send_packet_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packet send failures",
						},
						"recv_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Number of received packets",
						},
						"recv_packet_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Number of failed packet receives",
						},
						"retry_times": {
							Type: schema.TypeInt, Optional: true, Description: "Retry times",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Timouet value",
						},
						"unexpected_error": {
							Type: schema.TypeInt, Optional: true, Description: "Number of unexpected errors",
						},
						"conn_imdt_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Number of connection immediete success",
						},
						"sock_close_before_17": {
							Type: schema.TypeInt, Optional: true, Description: "Number of sockets closed before l7",
						},
						"sock_close_without_notify": {
							Type: schema.TypeInt, Optional: true, Description: "Number of sockets closed without notify",
						},
						"curr_health_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Current health rate",
						},
						"ext_health_rate": {
							Type: schema.TypeInt, Optional: true, Description: "External health rate",
						},
						"ext_health_rate_val": {
							Type: schema.TypeInt, Optional: true, Description: "External health rate value",
						},
						"total_number": {
							Type: schema.TypeInt, Optional: true, Description: "Total number",
						},
						"status_up": {
							Type: schema.TypeInt, Optional: true, Description: "Number of status ups",
						},
						"status_down": {
							Type: schema.TypeInt, Optional: true, Description: "Number of status downs",
						},
						"status_unkn": {
							Type: schema.TypeInt, Optional: true, Description: "Number of status unknowns",
						},
						"status_other": {
							Type: schema.TypeInt, Optional: true, Description: "Number of other status",
						},
						"running_time": {
							Type: schema.TypeInt, Optional: true, Description: "Running time",
						},
						"config_health_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Config health rate",
						},
						"ssl_post_handshake_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Number of ssl post handshake packets before client sends request",
						},
						"timeout_with_packet": {
							Type: schema.TypeInt, Optional: true, Description: "Number of pin timeouts while socket has packets",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHealthStatStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHealthStatStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHealthStatStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHealthStatStatsStats := setObjectSlbHealthStatStatsStats(res)
		d.Set("stats", SlbHealthStatStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHealthStatStatsStats(ret edpt.DataSlbHealthStatStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"num_burst":                 ret.DtSlbHealthStatStats.Stats.Num_burst,
			"max_jiffie":                ret.DtSlbHealthStatStats.Stats.Max_jiffie,
			"min_jiffie":                ret.DtSlbHealthStatStats.Stats.Min_jiffie,
			"avg_jiffie":                ret.DtSlbHealthStatStats.Stats.Avg_jiffie,
			"open_socket":               ret.DtSlbHealthStatStats.Stats.Open_socket,
			"open_socket_failed":        ret.DtSlbHealthStatStats.Stats.Open_socket_failed,
			"close_socket":              ret.DtSlbHealthStatStats.Stats.Close_socket,
			"connect_failed":            ret.DtSlbHealthStatStats.Stats.Connect_failed,
			"send_packet":               ret.DtSlbHealthStatStats.Stats.Send_packet,
			"send_packet_failed":        ret.DtSlbHealthStatStats.Stats.Send_packet_failed,
			"recv_packet":               ret.DtSlbHealthStatStats.Stats.Recv_packet,
			"recv_packet_failed":        ret.DtSlbHealthStatStats.Stats.Recv_packet_failed,
			"retry_times":               ret.DtSlbHealthStatStats.Stats.Retry_times,
			"timeout":                   ret.DtSlbHealthStatStats.Stats.Timeout,
			"unexpected_error":          ret.DtSlbHealthStatStats.Stats.Unexpected_error,
			"conn_imdt_succ":            ret.DtSlbHealthStatStats.Stats.Conn_imdt_succ,
			"sock_close_before_17":      ret.DtSlbHealthStatStats.Stats.Sock_close_before_17,
			"sock_close_without_notify": ret.DtSlbHealthStatStats.Stats.Sock_close_without_notify,
			"curr_health_rate":          ret.DtSlbHealthStatStats.Stats.Curr_health_rate,
			"ext_health_rate":           ret.DtSlbHealthStatStats.Stats.Ext_health_rate,
			"ext_health_rate_val":       ret.DtSlbHealthStatStats.Stats.Ext_health_rate_val,
			"total_number":              ret.DtSlbHealthStatStats.Stats.Total_number,
			"status_up":                 ret.DtSlbHealthStatStats.Stats.Status_up,
			"status_down":               ret.DtSlbHealthStatStats.Stats.Status_down,
			"status_unkn":               ret.DtSlbHealthStatStats.Stats.Status_unkn,
			"status_other":              ret.DtSlbHealthStatStats.Stats.Status_other,
			"running_time":              ret.DtSlbHealthStatStats.Stats.Running_time,
			"config_health_rate":        ret.DtSlbHealthStatStats.Stats.Config_health_rate,
			"ssl_post_handshake_packet": ret.DtSlbHealthStatStats.Stats.Ssl_post_handshake_packet,
			"timeout_with_packet":       ret.DtSlbHealthStatStats.Stats.Timeout_with_packet,
		},
	}
}

func getObjectSlbHealthStatStatsStats(d []interface{}) edpt.SlbHealthStatStatsStats {

	count1 := len(d)
	var ret edpt.SlbHealthStatStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Num_burst = in["num_burst"].(int)
		ret.Max_jiffie = in["max_jiffie"].(int)
		ret.Min_jiffie = in["min_jiffie"].(int)
		ret.Avg_jiffie = in["avg_jiffie"].(int)
		ret.Open_socket = in["open_socket"].(int)
		ret.Open_socket_failed = in["open_socket_failed"].(int)
		ret.Close_socket = in["close_socket"].(int)
		ret.Connect_failed = in["connect_failed"].(int)
		ret.Send_packet = in["send_packet"].(int)
		ret.Send_packet_failed = in["send_packet_failed"].(int)
		ret.Recv_packet = in["recv_packet"].(int)
		ret.Recv_packet_failed = in["recv_packet_failed"].(int)
		ret.Retry_times = in["retry_times"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Unexpected_error = in["unexpected_error"].(int)
		ret.Conn_imdt_succ = in["conn_imdt_succ"].(int)
		ret.Sock_close_before_17 = in["sock_close_before_17"].(int)
		ret.Sock_close_without_notify = in["sock_close_without_notify"].(int)
		ret.Curr_health_rate = in["curr_health_rate"].(int)
		ret.Ext_health_rate = in["ext_health_rate"].(int)
		ret.Ext_health_rate_val = in["ext_health_rate_val"].(int)
		ret.Total_number = in["total_number"].(int)
		ret.Status_up = in["status_up"].(int)
		ret.Status_down = in["status_down"].(int)
		ret.Status_unkn = in["status_unkn"].(int)
		ret.Status_other = in["status_other"].(int)
		ret.Running_time = in["running_time"].(int)
		ret.Config_health_rate = in["config_health_rate"].(int)
		ret.Ssl_post_handshake_packet = in["ssl_post_handshake_packet"].(int)
		ret.Timeout_with_packet = in["timeout_with_packet"].(int)
	}
	return ret
}

func dataToEndpointSlbHealthStatStats(d *schema.ResourceData) edpt.SlbHealthStatStats {
	var ret edpt.SlbHealthStatStats

	ret.Stats = getObjectSlbHealthStatStatsStats(d.Get("stats").([]interface{}))
	return ret
}
