package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbConnectionReuseStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_connection_reuse_stats`: Statistics for the object connection-reuse\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbConnectionReuseStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"current_open": {
							Type: schema.TypeInt, Optional: true, Description: "Open persist",
						},
						"current_active": {
							Type: schema.TypeInt, Optional: true, Description: "Active persist",
						},
						"nbind": {
							Type: schema.TypeInt, Optional: true, Description: "Total bind",
						},
						"nunbind": {
							Type: schema.TypeInt, Optional: true, Description: "Total unbind",
						},
						"nestab": {
							Type: schema.TypeInt, Optional: true, Description: "Total established",
						},
						"ntermi": {
							Type: schema.TypeInt, Optional: true, Description: "Total terminated",
						},
						"ntermi_err": {
							Type: schema.TypeInt, Optional: true, Description: "Total terminated by err",
						},
						"delay_unbind": {
							Type: schema.TypeInt, Optional: true, Description: "Delayed unbind",
						},
						"long_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Long resp",
						},
						"miss_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Missed resp",
						},
						"unbound_data_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Unbound data rcvd",
						},
						"pause_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Pause request",
						},
						"pause_conn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Pause request fail",
						},
						"resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume request",
						},
						"not_remove_from_rport": {
							Type: schema.TypeInt, Optional: true, Description: "Not remove from list",
						},
					},
				},
			},
		},
	}
}

func resourceSlbConnectionReuseStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbConnectionReuseStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbConnectionReuseStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbConnectionReuseStatsStats := setObjectSlbConnectionReuseStatsStats(res)
		d.Set("stats", SlbConnectionReuseStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbConnectionReuseStatsStats(ret edpt.DataSlbConnectionReuseStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"current_open":          ret.DtSlbConnectionReuseStats.Stats.Current_open,
			"current_active":        ret.DtSlbConnectionReuseStats.Stats.Current_active,
			"nbind":                 ret.DtSlbConnectionReuseStats.Stats.Nbind,
			"nunbind":               ret.DtSlbConnectionReuseStats.Stats.Nunbind,
			"nestab":                ret.DtSlbConnectionReuseStats.Stats.Nestab,
			"ntermi":                ret.DtSlbConnectionReuseStats.Stats.Ntermi,
			"ntermi_err":            ret.DtSlbConnectionReuseStats.Stats.Ntermi_err,
			"delay_unbind":          ret.DtSlbConnectionReuseStats.Stats.Delay_unbind,
			"long_resp":             ret.DtSlbConnectionReuseStats.Stats.Long_resp,
			"miss_resp":             ret.DtSlbConnectionReuseStats.Stats.Miss_resp,
			"unbound_data_rcv":      ret.DtSlbConnectionReuseStats.Stats.Unbound_data_rcv,
			"pause_conn":            ret.DtSlbConnectionReuseStats.Stats.Pause_conn,
			"pause_conn_fail":       ret.DtSlbConnectionReuseStats.Stats.Pause_conn_fail,
			"resume_conn":           ret.DtSlbConnectionReuseStats.Stats.Resume_conn,
			"not_remove_from_rport": ret.DtSlbConnectionReuseStats.Stats.Not_remove_from_rport,
		},
	}
}

func getObjectSlbConnectionReuseStatsStats(d []interface{}) edpt.SlbConnectionReuseStatsStats {

	count1 := len(d)
	var ret edpt.SlbConnectionReuseStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Current_open = in["current_open"].(int)
		ret.Current_active = in["current_active"].(int)
		ret.Nbind = in["nbind"].(int)
		ret.Nunbind = in["nunbind"].(int)
		ret.Nestab = in["nestab"].(int)
		ret.Ntermi = in["ntermi"].(int)
		ret.Ntermi_err = in["ntermi_err"].(int)
		ret.Delay_unbind = in["delay_unbind"].(int)
		ret.Long_resp = in["long_resp"].(int)
		ret.Miss_resp = in["miss_resp"].(int)
		ret.Unbound_data_rcv = in["unbound_data_rcv"].(int)
		ret.Pause_conn = in["pause_conn"].(int)
		ret.Pause_conn_fail = in["pause_conn_fail"].(int)
		ret.Resume_conn = in["resume_conn"].(int)
		ret.Not_remove_from_rport = in["not_remove_from_rport"].(int)
	}
	return ret
}

func dataToEndpointSlbConnectionReuseStats(d *schema.ResourceData) edpt.SlbConnectionReuseStats {
	var ret edpt.SlbConnectionReuseStats

	ret.Stats = getObjectSlbConnectionReuseStatsStats(d.Get("stats").([]interface{}))
	return ret
}
