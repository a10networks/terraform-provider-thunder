package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbL7sessionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_l7session_stats`: Statistics for the object l7session\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbL7sessionStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_server_conn_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Start Server Conn Success",
						},
						"conn_not_exist": {
							Type: schema.TypeInt, Optional: true, Description: "Conn does not exist",
						},
						"data_event": {
							Type: schema.TypeInt, Optional: true, Description: "Data event from TCP",
						},
						"client_fin": {
							Type: schema.TypeInt, Optional: true, Description: "FIN from client",
						},
						"server_fin": {
							Type: schema.TypeInt, Optional: true, Description: "FIN from server",
						},
						"wbuf_event": {
							Type: schema.TypeInt, Optional: true, Description: "Wbuf event from TCP",
						},
						"wbuf_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Wbuf event callback failed",
						},
						"err_event": {
							Type: schema.TypeInt, Optional: true, Description: "Err event from TCP",
						},
						"err_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Err event callback failed",
						},
						"server_conn_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection failed",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "RST from client",
						},
						"server_rst": {
							Type: schema.TypeInt, Optional: true, Description: "RST from server",
						},
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Curr proxy conn",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total proxy conn",
						},
						"server_select_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Server selection fail",
						},
						"data_cb_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Data event callback fail",
						},
						"hps_fwdreq_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fwd req fail",
						},
						"udp_data_event": {
							Type: schema.TypeInt, Optional: true, Description: "Data event from UDP",
						},
					},
				},
			},
		},
	}
}

func resourceSlbL7sessionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL7sessionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL7sessionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbL7sessionStatsStats := setObjectSlbL7sessionStatsStats(res)
		d.Set("stats", SlbL7sessionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbL7sessionStatsStats(ret edpt.DataSlbL7sessionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"start_server_conn_succ": ret.DtSlbL7sessionStats.Stats.Start_server_conn_succ,
			"conn_not_exist":         ret.DtSlbL7sessionStats.Stats.Conn_not_exist,
			"data_event":             ret.DtSlbL7sessionStats.Stats.Data_event,
			"client_fin":             ret.DtSlbL7sessionStats.Stats.Client_fin,
			"server_fin":             ret.DtSlbL7sessionStats.Stats.Server_fin,
			"wbuf_event":             ret.DtSlbL7sessionStats.Stats.Wbuf_event,
			"wbuf_cb_failed":         ret.DtSlbL7sessionStats.Stats.Wbuf_cb_failed,
			"err_event":              ret.DtSlbL7sessionStats.Stats.Err_event,
			"err_cb_failed":          ret.DtSlbL7sessionStats.Stats.Err_cb_failed,
			"server_conn_failed":     ret.DtSlbL7sessionStats.Stats.Server_conn_failed,
			"client_rst":             ret.DtSlbL7sessionStats.Stats.Client_rst,
			"server_rst":             ret.DtSlbL7sessionStats.Stats.Server_rst,
			"curr_proxy":             ret.DtSlbL7sessionStats.Stats.Curr_proxy,
			"total_proxy":            ret.DtSlbL7sessionStats.Stats.Total_proxy,
			"server_select_fail":     ret.DtSlbL7sessionStats.Stats.Server_select_fail,
			"data_cb_failed":         ret.DtSlbL7sessionStats.Stats.Data_cb_failed,
			"hps_fwdreq_fail":        ret.DtSlbL7sessionStats.Stats.Hps_fwdreq_fail,
			"udp_data_event":         ret.DtSlbL7sessionStats.Stats.Udp_data_event,
		},
	}
}

func getObjectSlbL7sessionStatsStats(d []interface{}) edpt.SlbL7sessionStatsStats {

	count1 := len(d)
	var ret edpt.SlbL7sessionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Start_server_conn_succ = in["start_server_conn_succ"].(int)
		ret.Conn_not_exist = in["conn_not_exist"].(int)
		ret.Data_event = in["data_event"].(int)
		ret.Client_fin = in["client_fin"].(int)
		ret.Server_fin = in["server_fin"].(int)
		ret.Wbuf_event = in["wbuf_event"].(int)
		ret.Wbuf_cb_failed = in["wbuf_cb_failed"].(int)
		ret.Err_event = in["err_event"].(int)
		ret.Err_cb_failed = in["err_cb_failed"].(int)
		ret.Server_conn_failed = in["server_conn_failed"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Server_rst = in["server_rst"].(int)
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Server_select_fail = in["server_select_fail"].(int)
		ret.Data_cb_failed = in["data_cb_failed"].(int)
		ret.Hps_fwdreq_fail = in["hps_fwdreq_fail"].(int)
		ret.Udp_data_event = in["udp_data_event"].(int)
	}
	return ret
}

func dataToEndpointSlbL7sessionStats(d *schema.ResourceData) edpt.SlbL7sessionStats {
	var ret edpt.SlbL7sessionStats

	ret.Stats = getObjectSlbL7sessionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
