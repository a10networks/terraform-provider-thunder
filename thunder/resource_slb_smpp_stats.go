package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSmppStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_smpp_stats`: Statistics for the object smpp\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSmppStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_proxy_current": {
							Type: schema.TypeInt, Optional: true, Description: "Curr SMPP Proxy",
						},
						"msg_proxy_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total SMPP Proxy",
						},
						"msg_proxy_client_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Client message rcvd",
						},
						"msg_proxy_client_send_success": {
							Type: schema.TypeInt, Optional: true, Description: "Sent to server",
						},
						"msg_proxy_client_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "Incomplete",
						},
						"msg_proxy_client_drop": {
							Type: schema.TypeInt, Optional: true, Description: "AX responds directly",
						},
						"msg_proxy_client_connection": {
							Type: schema.TypeInt, Optional: true, Description: "Connecting server",
						},
						"msg_proxy_client_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SMPP messages received from client but failed to forward to server",
						},
						"msg_proxy_server_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Server message rcvd",
						},
						"msg_proxy_server_send_success": {
							Type: schema.TypeInt, Optional: true, Description: "Sent to client",
						},
						"msg_proxy_server_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "Incomplete",
						},
						"msg_proxy_server_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Number of the packet AX drop",
						},
						"msg_proxy_server_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SMPP messages received from server but failed to forward to client",
						},
						"msg_proxy_create_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Server conn created",
						},
						"msg_proxy_start_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server connection created successfully",
						},
						"msg_proxy_fail_start_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server connection created failed",
						},
						"ax_response_directly": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packet which AX responds directly",
						},
						"select_client_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Client conn selection",
						},
						"select_client_by_req": {
							Type: schema.TypeInt, Optional: true, Description: "Select by request",
						},
						"select_client_from_list": {
							Type: schema.TypeInt, Optional: true, Description: "Select by roundbin",
						},
						"select_client_by_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Select by conn",
						},
						"select_client_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Select failed",
						},
						"select_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Server conn selection",
						},
						"select_server_by_req": {
							Type: schema.TypeInt, Optional: true, Description: "Select by request",
						},
						"select_server_from_list": {
							Type: schema.TypeInt, Optional: true, Description: "Select by roundbin",
						},
						"select_server_by_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Select server conn by client conn",
						},
						"select_server_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Fail to select server conn",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSmppStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSmppStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSmppStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSmppStatsStats := setObjectSlbSmppStatsStats(res)
		d.Set("stats", SlbSmppStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSmppStatsStats(ret edpt.DataSlbSmppStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"msg_proxy_current":                ret.DtSlbSmppStats.Stats.Msg_proxy_current,
			"msg_proxy_total":                  ret.DtSlbSmppStats.Stats.Msg_proxy_total,
			"msg_proxy_client_recv":            ret.DtSlbSmppStats.Stats.Msg_proxy_client_recv,
			"msg_proxy_client_send_success":    ret.DtSlbSmppStats.Stats.Msg_proxy_client_send_success,
			"msg_proxy_client_incomplete":      ret.DtSlbSmppStats.Stats.Msg_proxy_client_incomplete,
			"msg_proxy_client_drop":            ret.DtSlbSmppStats.Stats.Msg_proxy_client_drop,
			"msg_proxy_client_connection":      ret.DtSlbSmppStats.Stats.Msg_proxy_client_connection,
			"msg_proxy_client_fail":            ret.DtSlbSmppStats.Stats.Msg_proxy_client_fail,
			"msg_proxy_server_recv":            ret.DtSlbSmppStats.Stats.Msg_proxy_server_recv,
			"msg_proxy_server_send_success":    ret.DtSlbSmppStats.Stats.Msg_proxy_server_send_success,
			"msg_proxy_server_incomplete":      ret.DtSlbSmppStats.Stats.Msg_proxy_server_incomplete,
			"msg_proxy_server_drop":            ret.DtSlbSmppStats.Stats.Msg_proxy_server_drop,
			"msg_proxy_server_fail":            ret.DtSlbSmppStats.Stats.Msg_proxy_server_fail,
			"msg_proxy_create_server_conn":     ret.DtSlbSmppStats.Stats.Msg_proxy_create_server_conn,
			"msg_proxy_start_server_conn":      ret.DtSlbSmppStats.Stats.Msg_proxy_start_server_conn,
			"msg_proxy_fail_start_server_conn": ret.DtSlbSmppStats.Stats.Msg_proxy_fail_start_server_conn,
			"ax_response_directly":             ret.DtSlbSmppStats.Stats.Ax_response_directly,
			"select_client_conn":               ret.DtSlbSmppStats.Stats.Select_client_conn,
			"select_client_by_req":             ret.DtSlbSmppStats.Stats.Select_client_by_req,
			"select_client_from_list":          ret.DtSlbSmppStats.Stats.Select_client_from_list,
			"select_client_by_conn":            ret.DtSlbSmppStats.Stats.Select_client_by_conn,
			"select_client_fail":               ret.DtSlbSmppStats.Stats.Select_client_fail,
			"select_server_conn":               ret.DtSlbSmppStats.Stats.Select_server_conn,
			"select_server_by_req":             ret.DtSlbSmppStats.Stats.Select_server_by_req,
			"select_server_from_list":          ret.DtSlbSmppStats.Stats.Select_server_from_list,
			"select_server_by_conn":            ret.DtSlbSmppStats.Stats.Select_server_by_conn,
			"select_server_fail":               ret.DtSlbSmppStats.Stats.Select_server_fail,
		},
	}
}

func getObjectSlbSmppStatsStats(d []interface{}) edpt.SlbSmppStatsStats {

	count1 := len(d)
	var ret edpt.SlbSmppStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Msg_proxy_current = in["msg_proxy_current"].(int)
		ret.Msg_proxy_total = in["msg_proxy_total"].(int)
		ret.Msg_proxy_client_recv = in["msg_proxy_client_recv"].(int)
		ret.Msg_proxy_client_send_success = in["msg_proxy_client_send_success"].(int)
		ret.Msg_proxy_client_incomplete = in["msg_proxy_client_incomplete"].(int)
		ret.Msg_proxy_client_drop = in["msg_proxy_client_drop"].(int)
		ret.Msg_proxy_client_connection = in["msg_proxy_client_connection"].(int)
		ret.Msg_proxy_client_fail = in["msg_proxy_client_fail"].(int)
		ret.Msg_proxy_server_recv = in["msg_proxy_server_recv"].(int)
		ret.Msg_proxy_server_send_success = in["msg_proxy_server_send_success"].(int)
		ret.Msg_proxy_server_incomplete = in["msg_proxy_server_incomplete"].(int)
		ret.Msg_proxy_server_drop = in["msg_proxy_server_drop"].(int)
		ret.Msg_proxy_server_fail = in["msg_proxy_server_fail"].(int)
		ret.Msg_proxy_create_server_conn = in["msg_proxy_create_server_conn"].(int)
		ret.Msg_proxy_start_server_conn = in["msg_proxy_start_server_conn"].(int)
		ret.Msg_proxy_fail_start_server_conn = in["msg_proxy_fail_start_server_conn"].(int)
		ret.Ax_response_directly = in["ax_response_directly"].(int)
		ret.Select_client_conn = in["select_client_conn"].(int)
		ret.Select_client_by_req = in["select_client_by_req"].(int)
		ret.Select_client_from_list = in["select_client_from_list"].(int)
		ret.Select_client_by_conn = in["select_client_by_conn"].(int)
		ret.Select_client_fail = in["select_client_fail"].(int)
		ret.Select_server_conn = in["select_server_conn"].(int)
		ret.Select_server_by_req = in["select_server_by_req"].(int)
		ret.Select_server_from_list = in["select_server_from_list"].(int)
		ret.Select_server_by_conn = in["select_server_by_conn"].(int)
		ret.Select_server_fail = in["select_server_fail"].(int)
	}
	return ret
}

func dataToEndpointSlbSmppStats(d *schema.ResourceData) edpt.SlbSmppStats {
	var ret edpt.SlbSmppStats

	ret.Stats = getObjectSlbSmppStatsStats(d.Get("stats").([]interface{}))
	return ret
}
