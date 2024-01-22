package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSipStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_sip_stats`: Statistics for the object sip\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSipStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_proxy_current": {
							Type: schema.TypeInt, Optional: true, Description: "Number of current sip proxy connections",
						},
						"msg_proxy_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of sip proxy connections",
						},
						"msg_proxy_client_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SIP messages received from client",
						},
						"msg_proxy_client_send_success": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SIP messages received from client and forwarded to server",
						},
						"msg_proxy_client_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packet which contains incomplete message",
						},
						"msg_proxy_client_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Number of AX drop",
						},
						"msg_proxy_client_connection": {
							Type: schema.TypeInt, Optional: true, Description: "Connecting server",
						},
						"msg_proxy_client_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SIP messages received from client but failed to forward to server",
						},
						"msg_proxy_server_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SIP messages received from server",
						},
						"msg_proxy_server_send_success": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SIP messages received from server and forwarded to client",
						},
						"msg_proxy_server_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "Number of packet which contains incomplete message",
						},
						"msg_proxy_server_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Number of AX drop",
						},
						"msg_proxy_server_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of SIP messages received from server but failed to forward to client",
						},
						"msg_proxy_create_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server connection system tries to create",
						},
						"msg_proxy_start_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server connection created successfully",
						},
						"msg_proxy_fail_start_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Number of server connection create failed",
						},
						"session_created": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Session created",
						},
						"session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "SIP Session freed",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSipStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSipStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSipStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSipStatsStats := setObjectSlbSipStatsStats(res)
		d.Set("stats", SlbSipStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSipStatsStats(ret edpt.DataSlbSipStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"msg_proxy_current":                ret.DtSlbSipStats.Stats.Msg_proxy_current,
			"msg_proxy_total":                  ret.DtSlbSipStats.Stats.Msg_proxy_total,
			"msg_proxy_client_recv":            ret.DtSlbSipStats.Stats.Msg_proxy_client_recv,
			"msg_proxy_client_send_success":    ret.DtSlbSipStats.Stats.Msg_proxy_client_send_success,
			"msg_proxy_client_incomplete":      ret.DtSlbSipStats.Stats.Msg_proxy_client_incomplete,
			"msg_proxy_client_drop":            ret.DtSlbSipStats.Stats.Msg_proxy_client_drop,
			"msg_proxy_client_connection":      ret.DtSlbSipStats.Stats.Msg_proxy_client_connection,
			"msg_proxy_client_fail":            ret.DtSlbSipStats.Stats.Msg_proxy_client_fail,
			"msg_proxy_server_recv":            ret.DtSlbSipStats.Stats.Msg_proxy_server_recv,
			"msg_proxy_server_send_success":    ret.DtSlbSipStats.Stats.Msg_proxy_server_send_success,
			"msg_proxy_server_incomplete":      ret.DtSlbSipStats.Stats.Msg_proxy_server_incomplete,
			"msg_proxy_server_drop":            ret.DtSlbSipStats.Stats.Msg_proxy_server_drop,
			"msg_proxy_server_fail":            ret.DtSlbSipStats.Stats.Msg_proxy_server_fail,
			"msg_proxy_create_server_conn":     ret.DtSlbSipStats.Stats.Msg_proxy_create_server_conn,
			"msg_proxy_start_server_conn":      ret.DtSlbSipStats.Stats.Msg_proxy_start_server_conn,
			"msg_proxy_fail_start_server_conn": ret.DtSlbSipStats.Stats.Msg_proxy_fail_start_server_conn,
			"session_created":                  ret.DtSlbSipStats.Stats.Session_created,
			"session_freed":                    ret.DtSlbSipStats.Stats.Session_freed,
		},
	}
}

func getObjectSlbSipStatsStats(d []interface{}) edpt.SlbSipStatsStats {

	count1 := len(d)
	var ret edpt.SlbSipStatsStats
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
		ret.Session_created = in["session_created"].(int)
		ret.Session_freed = in["session_freed"].(int)
	}
	return ret
}

func dataToEndpointSlbSipStats(d *schema.ResourceData) edpt.SlbSipStats {
	var ret edpt.SlbSipStats

	ret.Stats = getObjectSlbSipStatsStats(d.Get("stats").([]interface{}))
	return ret
}
