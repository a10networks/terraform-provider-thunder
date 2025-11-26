package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProbeInfoStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_controller_probe_info_stats`: Statistics for the object probe-info\n\n__PLACEHOLDER__",
		ReadContext: resourceControllerProbeInfoStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_showtech_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DATA_SHOWTECH samples sent successfully to probe",
						},
						"data_showtech_failed": {
							Type: schema.TypeInt, Optional: true, Description: "DATA_SHOWTECH samples failed to send to probe",
						},
						"data_varlog_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DATA_VARLOG samples sent successfully to probe",
						},
						"data_varlog_failed": {
							Type: schema.TypeInt, Optional: true, Description: "DATA_VARLOG samples failed to send to probe",
						},
						"ssh_connection_failed": {
							Type: schema.TypeInt, Optional: true, Description: "SSH_CONNECTION failures (ACOS-A10C)",
						},
					},
				},
			},
		},
	}
}

func resourceControllerProbeInfoStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProbeInfoStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProbeInfoStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ControllerProbeInfoStatsStats := setObjectControllerProbeInfoStatsStats(res)
		d.Set("stats", ControllerProbeInfoStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectControllerProbeInfoStatsStats(ret edpt.DataControllerProbeInfoStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"data_showtech_sent":    ret.DtControllerProbeInfoStats.Stats.DataShowtechSent,
			"data_showtech_failed":  ret.DtControllerProbeInfoStats.Stats.DataShowtechFailed,
			"data_varlog_sent":      ret.DtControllerProbeInfoStats.Stats.DataVarlogSent,
			"data_varlog_failed":    ret.DtControllerProbeInfoStats.Stats.DataVarlogFailed,
			"ssh_connection_failed": ret.DtControllerProbeInfoStats.Stats.SshConnectionFailed,
		},
	}
}

func getObjectControllerProbeInfoStatsStats(d []interface{}) edpt.ControllerProbeInfoStatsStats {

	count1 := len(d)
	var ret edpt.ControllerProbeInfoStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DataShowtechSent = in["data_showtech_sent"].(int)
		ret.DataShowtechFailed = in["data_showtech_failed"].(int)
		ret.DataVarlogSent = in["data_varlog_sent"].(int)
		ret.DataVarlogFailed = in["data_varlog_failed"].(int)
		ret.SshConnectionFailed = in["ssh_connection_failed"].(int)
	}
	return ret
}

func dataToEndpointControllerProbeInfoStats(d *schema.ResourceData) edpt.ControllerProbeInfoStats {
	var ret edpt.ControllerProbeInfoStats

	ret.Stats = getObjectControllerProbeInfoStatsStats(d.Get("stats").([]interface{}))
	return ret
}
