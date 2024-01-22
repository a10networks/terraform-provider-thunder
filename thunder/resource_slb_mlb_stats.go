package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbMlbStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_mlb_stats`: Statistics for the object mlb\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbMlbStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_msg_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Client message sent",
						},
						"server_msg_received": {
							Type: schema.TypeInt, Optional: true, Description: "Server message received",
						},
						"server_conn_created": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection created",
						},
						"server_conn_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection reset",
						},
						"server_conn_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection failed",
						},
						"server_conn_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Server connection closed",
						},
						"client_conn_created": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection created",
						},
						"client_conn_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection closed",
						},
						"client_conn_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Client connection not found",
						},
						"msg_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Message dropped",
						},
						"msg_rerouted": {
							Type: schema.TypeInt, Optional: true, Description: "Message rerouted",
						},
						"mlb_dcmsg_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Dcmsg sent",
						},
						"mlb_dcmsg_received": {
							Type: schema.TypeInt, Optional: true, Description: "Dcmsg received",
						},
						"mlb_dcmsg_error": {
							Type: schema.TypeInt, Optional: true, Description: "Dcmsg error",
						},
						"mlb_dcmsg_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Dcmsg alloc",
						},
						"mlb_dcmsg_free": {
							Type: schema.TypeInt, Optional: true, Description: "Dcmsg free",
						},
						"mlb_server_probe": {
							Type: schema.TypeInt, Optional: true, Description: "Server probe",
						},
						"mlb_server_down": {
							Type: schema.TypeInt, Optional: true, Description: "Server down",
						},
					},
				},
			},
		},
	}
}

func resourceSlbMlbStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbMlbStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbMlbStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbMlbStatsStats := setObjectSlbMlbStatsStats(res)
		d.Set("stats", SlbMlbStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbMlbStatsStats(ret edpt.DataSlbMlbStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"client_msg_sent":       ret.DtSlbMlbStats.Stats.Client_msg_sent,
			"server_msg_received":   ret.DtSlbMlbStats.Stats.Server_msg_received,
			"server_conn_created":   ret.DtSlbMlbStats.Stats.Server_conn_created,
			"server_conn_rst":       ret.DtSlbMlbStats.Stats.Server_conn_rst,
			"server_conn_failed":    ret.DtSlbMlbStats.Stats.Server_conn_failed,
			"server_conn_closed":    ret.DtSlbMlbStats.Stats.Server_conn_closed,
			"client_conn_created":   ret.DtSlbMlbStats.Stats.Client_conn_created,
			"client_conn_closed":    ret.DtSlbMlbStats.Stats.Client_conn_closed,
			"client_conn_not_found": ret.DtSlbMlbStats.Stats.Client_conn_not_found,
			"msg_dropped":           ret.DtSlbMlbStats.Stats.Msg_dropped,
			"msg_rerouted":          ret.DtSlbMlbStats.Stats.Msg_rerouted,
			"mlb_dcmsg_sent":        ret.DtSlbMlbStats.Stats.Mlb_dcmsg_sent,
			"mlb_dcmsg_received":    ret.DtSlbMlbStats.Stats.Mlb_dcmsg_received,
			"mlb_dcmsg_error":       ret.DtSlbMlbStats.Stats.Mlb_dcmsg_error,
			"mlb_dcmsg_alloc":       ret.DtSlbMlbStats.Stats.Mlb_dcmsg_alloc,
			"mlb_dcmsg_free":        ret.DtSlbMlbStats.Stats.Mlb_dcmsg_free,
			"mlb_server_probe":      ret.DtSlbMlbStats.Stats.Mlb_server_probe,
			"mlb_server_down":       ret.DtSlbMlbStats.Stats.Mlb_server_down,
		},
	}
}

func getObjectSlbMlbStatsStats(d []interface{}) edpt.SlbMlbStatsStats {

	count1 := len(d)
	var ret edpt.SlbMlbStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Client_msg_sent = in["client_msg_sent"].(int)
		ret.Server_msg_received = in["server_msg_received"].(int)
		ret.Server_conn_created = in["server_conn_created"].(int)
		ret.Server_conn_rst = in["server_conn_rst"].(int)
		ret.Server_conn_failed = in["server_conn_failed"].(int)
		ret.Server_conn_closed = in["server_conn_closed"].(int)
		ret.Client_conn_created = in["client_conn_created"].(int)
		ret.Client_conn_closed = in["client_conn_closed"].(int)
		ret.Client_conn_not_found = in["client_conn_not_found"].(int)
		ret.Msg_dropped = in["msg_dropped"].(int)
		ret.Msg_rerouted = in["msg_rerouted"].(int)
		ret.Mlb_dcmsg_sent = in["mlb_dcmsg_sent"].(int)
		ret.Mlb_dcmsg_received = in["mlb_dcmsg_received"].(int)
		ret.Mlb_dcmsg_error = in["mlb_dcmsg_error"].(int)
		ret.Mlb_dcmsg_alloc = in["mlb_dcmsg_alloc"].(int)
		ret.Mlb_dcmsg_free = in["mlb_dcmsg_free"].(int)
		ret.Mlb_server_probe = in["mlb_server_probe"].(int)
		ret.Mlb_server_down = in["mlb_server_down"].(int)
	}
	return ret
}

func dataToEndpointSlbMlbStats(d *schema.ResourceData) edpt.SlbMlbStats {
	var ret edpt.SlbMlbStats

	ret.Stats = getObjectSlbMlbStatsStats(d.Get("stats").([]interface{}))
	return ret
}
