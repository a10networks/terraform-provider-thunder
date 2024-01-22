package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsVbladesStatStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vcs_vblades_stat_stats`: Statistics for the object stat\n\n__PLACEHOLDER__",
		ReadContext: resourceVcsVbladesStatStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"slave_recv_err": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Receive Errors counter of aVCS election",
						},
						"slave_send_err": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Send Errors counter of aVCS election",
						},
						"slave_recv_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Received Bytes counter of aVCS election",
						},
						"slave_sent_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Sent Bytes counter of aVCS election",
						},
						"slave_n_recv": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Received Messages counter of aVCS election",
						},
						"slave_n_sent": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Sent Messages counter of aVCS election",
						},
						"slave_msg_inval": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Invalid Messages counter of aVCS election",
						},
						"slave_keepalive": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Received Keepalives counter of aVCS election",
						},
						"slave_cfg_upd": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Received Configuration Updates counter of aVCS election",
						},
						"slave_cfg_upd_l1_fail": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Local Configuration Update Errors (1) counter of aVCS election",
						},
						"slave_cfg_upd_r_fail": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Remote Configuration Update Errors counter of aVCS election",
						},
						"slave_cfg_upd_l2_fail": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Local Configuration Update Errors (2) counter of aVCS election",
						},
						"slave_cfg_upd_notif_err": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Configuration Update Notif Errors counter of aVCS election",
						},
						"slave_cfg_upd_result_err": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Configuration Update Result Errors counter of aVCS election",
						},
					},
				},
			},
			"vblade_id": {
				Type: schema.TypeInt, Required: true, Description: "vBlade-id",
			},
		},
	}
}

func resourceVcsVbladesStatStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsVbladesStatStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsVbladesStatStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VcsVbladesStatStatsStats := setObjectVcsVbladesStatStatsStats(res)
		d.Set("stats", VcsVbladesStatStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVcsVbladesStatStatsStats(ret edpt.DataVcsVbladesStatStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"slave_recv_err":           ret.DtVcsVbladesStatStats.Stats.Slave_recv_err,
			"slave_send_err":           ret.DtVcsVbladesStatStats.Stats.Slave_send_err,
			"slave_recv_bytes":         ret.DtVcsVbladesStatStats.Stats.Slave_recv_bytes,
			"slave_sent_bytes":         ret.DtVcsVbladesStatStats.Stats.Slave_sent_bytes,
			"slave_n_recv":             ret.DtVcsVbladesStatStats.Stats.Slave_n_recv,
			"slave_n_sent":             ret.DtVcsVbladesStatStats.Stats.Slave_n_sent,
			"slave_msg_inval":          ret.DtVcsVbladesStatStats.Stats.Slave_msg_inval,
			"slave_keepalive":          ret.DtVcsVbladesStatStats.Stats.Slave_keepalive,
			"slave_cfg_upd":            ret.DtVcsVbladesStatStats.Stats.Slave_cfg_upd,
			"slave_cfg_upd_l1_fail":    ret.DtVcsVbladesStatStats.Stats.Slave_cfg_upd_l1_fail,
			"slave_cfg_upd_r_fail":     ret.DtVcsVbladesStatStats.Stats.Slave_cfg_upd_r_fail,
			"slave_cfg_upd_l2_fail":    ret.DtVcsVbladesStatStats.Stats.Slave_cfg_upd_l2_fail,
			"slave_cfg_upd_notif_err":  ret.DtVcsVbladesStatStats.Stats.Slave_cfg_upd_notif_err,
			"slave_cfg_upd_result_err": ret.DtVcsVbladesStatStats.Stats.Slave_cfg_upd_result_err,
		},
	}
}

func getObjectVcsVbladesStatStatsStats(d []interface{}) edpt.VcsVbladesStatStatsStats {

	count1 := len(d)
	var ret edpt.VcsVbladesStatStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Slave_recv_err = in["slave_recv_err"].(int)
		ret.Slave_send_err = in["slave_send_err"].(int)
		ret.Slave_recv_bytes = in["slave_recv_bytes"].(int)
		ret.Slave_sent_bytes = in["slave_sent_bytes"].(int)
		ret.Slave_n_recv = in["slave_n_recv"].(int)
		ret.Slave_n_sent = in["slave_n_sent"].(int)
		ret.Slave_msg_inval = in["slave_msg_inval"].(int)
		ret.Slave_keepalive = in["slave_keepalive"].(int)
		ret.Slave_cfg_upd = in["slave_cfg_upd"].(int)
		ret.Slave_cfg_upd_l1_fail = in["slave_cfg_upd_l1_fail"].(int)
		ret.Slave_cfg_upd_r_fail = in["slave_cfg_upd_r_fail"].(int)
		ret.Slave_cfg_upd_l2_fail = in["slave_cfg_upd_l2_fail"].(int)
		ret.Slave_cfg_upd_notif_err = in["slave_cfg_upd_notif_err"].(int)
		ret.Slave_cfg_upd_result_err = in["slave_cfg_upd_result_err"].(int)
	}
	return ret
}

func dataToEndpointVcsVbladesStatStats(d *schema.ResourceData) edpt.VcsVbladesStatStats {
	var ret edpt.VcsVbladesStatStats

	ret.Stats = getObjectVcsVbladesStatStatsStats(d.Get("stats").([]interface{}))

	ret.VbladeId = d.Get("vblade_id").(int)
	return ret
}
