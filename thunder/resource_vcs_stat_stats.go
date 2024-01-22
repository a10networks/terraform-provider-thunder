package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsStatStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_vcs_stat_stats`: Statistics for the object stat\n\n__PLACEHOLDER__",
		ReadContext: resourceVcsStatStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"elect_recv_err": {
							Type: schema.TypeInt, Optional: true, Description: "Receive error counter of aVCS election",
						},
						"elect_send_err": {
							Type: schema.TypeInt, Optional: true, Description: "Send error counter of aVCS election",
						},
						"elect_recv_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Receive bytes counter of aVCS election",
						},
						"elect_send_byte": {
							Type: schema.TypeInt, Optional: true, Description: "Send bytes counter of aVCS election",
						},
						"elect_pdu_master_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Received vMaster-PDU counter of aVCS election",
						},
						"elect_pdu_master_cand_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Received MC-PDU counter of aVCS election",
						},
						"elect_pdu_slave_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Received vBlade-PDU counter of aVCS election",
						},
						"elect_pdu_master_take_over_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Received MTO-PDU counter of aVCS election",
						},
						"elect_pdu_unknown_recv": {
							Type: schema.TypeInt, Optional: true, Description: "Received Unknown-PDU counter of aVCS election",
						},
						"elect_pdu_master_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent vMaster-PDU counter of aVCS election",
						},
						"elect_pdu_master_cand_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent MC-PDU counter of aVCS election",
						},
						"elect_pdu_slave_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent vBlade-PDU counter of aVCS election",
						},
						"elect_pdu_master_take_over_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent MTO-PDU counter of aVCS election",
						},
						"elect_pdu_unknown_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sent Unknown-PDU counter of aVCS election",
						},
						"elect_pdu_inval": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid PDU counter of aVCS election",
						},
						"elect_pdu_hw_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "PDU HW mismatch counter of aVCS election",
						},
						"elect_pdu_cluster_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "PDU Chassis-ID mismatch counter of aVCS election",
						},
						"elect_pdu_dev_id_collision": {
							Type: schema.TypeInt, Optional: true, Description: "PDU Device-ID collision counter of aVCS election",
						},
						"elect_mc_discard_master": {
							Type: schema.TypeInt, Optional: true, Description: "MC discarded vMaster-PDU counter of aVCS election",
						},
						"elect_mc_replace_master": {
							Type: schema.TypeInt, Optional: true, Description: "MC replaced vMaster-PDU counter of aVCS election",
						},
						"elect_mc_dup_masterr": {
							Type: schema.TypeInt, Optional: true, Description: "MC duplicate vMaster-PDU counter of aVCS election",
						},
						"elect_mc_reset_timer_by_mc": {
							Type: schema.TypeInt, Optional: true, Description: "MC timers reset by MC-PDU counter of aVCS election",
						},
						"elect_mc_reset_timer_by_mto": {
							Type: schema.TypeInt, Optional: true, Description: "MC timers reset by MTO-PDU counter of aVCS election",
						},
						"elect_slave_dup_master": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade duplicate vMaster-PDU counter of aVCS election",
						},
						"elect_slave_discard_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade discard challenger counter of aVCS election",
						},
						"elect_slave_replace_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade replace challenger counter of aVCS election",
						},
						"elect_slave_dup_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade duplicate challenger counter of aVCS election",
						},
						"elect_slave_discard_neighbour": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade discard neighbour counter of aVCS election",
						},
						"elect_slave_too_many_neighbour": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade too many neighbours counter of aVCS election",
						},
						"elect_slave_dup_neighbour": {
							Type: schema.TypeInt, Optional: true, Description: "send vBlade duplicate neighbours of aVCS election",
						},
						"elect_master_discard_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster discard challenger counter of aVCS election",
						},
						"elect_master_new_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster new challenger counter of aVCS election",
						},
						"elect_master_replace_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster replace challenger counter of aVCS election",
						},
						"elect_master_dup_challenger": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster duplicate challenger counter of aVCS election",
						},
						"elect_master_discard_neighbour": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster discard neighbour counter of aVCS election",
						},
						"elect_master_too_many_neighbour": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster too many neighbours counter of aVCS election",
						},
						"elect_master_dup_neighbour": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster duplicate neighbours counter of aVCS election",
						},
						"elect_enter_master_cand_stat": {
							Type: schema.TypeInt, Optional: true, Description: "Enter MC counter of aVCS election",
						},
						"elect_enter_slave": {
							Type: schema.TypeInt, Optional: true, Description: "Enter vBlade counter of aVCS election",
						},
						"elect_enter_master": {
							Type: schema.TypeInt, Optional: true, Description: "Enter vMaster counter of aVCS election",
						},
						"elect_enter_master_take_over": {
							Type: schema.TypeInt, Optional: true, Description: "Enter MTO counter of aVCS election",
						},
						"elect_leave_master_cand": {
							Type: schema.TypeInt, Optional: true, Description: "Leave MC counter of aVCS election",
						},
						"elect_leave_slave": {
							Type: schema.TypeInt, Optional: true, Description: "Leave vBlade counter of aVCS election",
						},
						"elect_leave_master": {
							Type: schema.TypeInt, Optional: true, Description: "Leave vMaster counter of aVCS election",
						},
						"elect_leave_master_take_over": {
							Type: schema.TypeInt, Optional: true, Description: "Leave MTO counter of aVCS election",
						},
						"master_slave_start_err": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster Start vBlade Errors counter of aVCS election",
						},
						"master_slave_start": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster vBlades Started counter of aVCS election",
						},
						"master_slave_stop": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster vBlades stopped counter of aVCS election",
						},
						"master_cfg_upd": {
							Type: schema.TypeInt, Optional: true, Description: "Received vMaster Configuration Updates counter of aVCS election",
						},
						"master_cfg_upd_l_fail": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster Local Configuration Update Errors counter of aVCS election",
						},
						"master_cfg_upd_r_fail": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster Remote Configuration Update Errors counter of aVCS election",
						},
						"master_cfg_upd_notif_err": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster Configuration Update Notif Errors counter of aVCS election",
						},
						"master_cfg_upd_result_err": {
							Type: schema.TypeInt, Optional: true, Description: "vMaster Configuration Update Result Errors counter of aVCS election",
						},
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
						"slave_cfg_upd_fail": {
							Type: schema.TypeInt, Optional: true, Description: "vBlade Configuration Update Failures counter of aVCS election",
						},
						"daemon_n_elec_start": {
							Type: schema.TypeInt, Optional: true, Description: "times of aVCS election start",
						},
						"daemon_n_elec_stop": {
							Type: schema.TypeInt, Optional: true, Description: "times of aVCS election stop",
						},
						"daemon_recv_err": {
							Type: schema.TypeInt, Optional: true, Description: "counter of aVCS daemon receive error",
						},
						"daemon_send_err": {
							Type: schema.TypeInt, Optional: true, Description: "counter of aVCS daemon sent error",
						},
						"daemon_recv_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "bytes of aVCS daemon receive",
						},
						"daemon_sent_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "bytes of aVCS daemon sent",
						},
						"daemon_n_recv": {
							Type: schema.TypeInt, Optional: true, Description: "counter of aVCS daemon receive",
						},
						"daemon_n_sent": {
							Type: schema.TypeInt, Optional: true, Description: "counter of aVCS daemon sent",
						},
						"daemon_msg_inval": {
							Type: schema.TypeInt, Optional: true, Description: "counter of aVCS daemon invalid message",
						},
						"daemon_msg_handle_failure": {
							Type: schema.TypeInt, Optional: true, Description: "counter of aVCS daemon message handle failure",
						},
					},
				},
			},
		},
	}
}

func resourceVcsStatStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsStatStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsStatStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VcsStatStatsStats := setObjectVcsStatStatsStats(res)
		d.Set("stats", VcsStatStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVcsStatStatsStats(ret edpt.DataVcsStatStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"elect_recv_err":                  ret.DtVcsStatStats.Stats.Elect_recv_err,
			"elect_send_err":                  ret.DtVcsStatStats.Stats.Elect_send_err,
			"elect_recv_byte":                 ret.DtVcsStatStats.Stats.Elect_recv_byte,
			"elect_send_byte":                 ret.DtVcsStatStats.Stats.Elect_send_byte,
			"elect_pdu_master_recv":           ret.DtVcsStatStats.Stats.Elect_pdu_master_recv,
			"elect_pdu_master_cand_recv":      ret.DtVcsStatStats.Stats.Elect_pdu_master_cand_recv,
			"elect_pdu_slave_recv":            ret.DtVcsStatStats.Stats.Elect_pdu_slave_recv,
			"elect_pdu_master_take_over_recv": ret.DtVcsStatStats.Stats.Elect_pdu_master_take_over_recv,
			"elect_pdu_unknown_recv":          ret.DtVcsStatStats.Stats.Elect_pdu_unknown_recv,
			"elect_pdu_master_sent":           ret.DtVcsStatStats.Stats.Elect_pdu_master_sent,
			"elect_pdu_master_cand_sent":      ret.DtVcsStatStats.Stats.Elect_pdu_master_cand_sent,
			"elect_pdu_slave_sent":            ret.DtVcsStatStats.Stats.Elect_pdu_slave_sent,
			"elect_pdu_master_take_over_sent": ret.DtVcsStatStats.Stats.Elect_pdu_master_take_over_sent,
			"elect_pdu_unknown_sent":          ret.DtVcsStatStats.Stats.Elect_pdu_unknown_sent,
			"elect_pdu_inval":                 ret.DtVcsStatStats.Stats.Elect_pdu_inval,
			"elect_pdu_hw_mismatch":           ret.DtVcsStatStats.Stats.Elect_pdu_hw_mismatch,
			"elect_pdu_cluster_mismatch":      ret.DtVcsStatStats.Stats.Elect_pdu_cluster_mismatch,
			"elect_pdu_dev_id_collision":      ret.DtVcsStatStats.Stats.Elect_pdu_dev_id_collision,
			"elect_mc_discard_master":         ret.DtVcsStatStats.Stats.Elect_mc_discard_master,
			"elect_mc_replace_master":         ret.DtVcsStatStats.Stats.Elect_mc_replace_master,
			"elect_mc_dup_masterr":            ret.DtVcsStatStats.Stats.Elect_mc_dup_masterr,
			"elect_mc_reset_timer_by_mc":      ret.DtVcsStatStats.Stats.Elect_mc_reset_timer_by_mc,
			"elect_mc_reset_timer_by_mto":     ret.DtVcsStatStats.Stats.Elect_mc_reset_timer_by_mto,
			"elect_slave_dup_master":          ret.DtVcsStatStats.Stats.Elect_slave_dup_master,
			"elect_slave_discard_challenger":  ret.DtVcsStatStats.Stats.Elect_slave_discard_challenger,
			"elect_slave_replace_challenger":  ret.DtVcsStatStats.Stats.Elect_slave_replace_challenger,
			"elect_slave_dup_challenger":      ret.DtVcsStatStats.Stats.Elect_slave_dup_challenger,
			"elect_slave_discard_neighbour":   ret.DtVcsStatStats.Stats.Elect_slave_discard_neighbour,
			"elect_slave_too_many_neighbour":  ret.DtVcsStatStats.Stats.Elect_slave_too_many_neighbour,
			"elect_slave_dup_neighbour":       ret.DtVcsStatStats.Stats.Elect_slave_dup_neighbour,
			"elect_master_discard_challenger": ret.DtVcsStatStats.Stats.Elect_master_discard_challenger,
			"elect_master_new_challenger":     ret.DtVcsStatStats.Stats.Elect_master_new_challenger,
			"elect_master_replace_challenger": ret.DtVcsStatStats.Stats.Elect_master_replace_challenger,
			"elect_master_dup_challenger":     ret.DtVcsStatStats.Stats.Elect_master_dup_challenger,
			"elect_master_discard_neighbour":  ret.DtVcsStatStats.Stats.Elect_master_discard_neighbour,
			"elect_master_too_many_neighbour": ret.DtVcsStatStats.Stats.Elect_master_too_many_neighbour,
			"elect_master_dup_neighbour":      ret.DtVcsStatStats.Stats.Elect_master_dup_neighbour,
			"elect_enter_master_cand_stat":    ret.DtVcsStatStats.Stats.Elect_enter_master_cand_stat,
			"elect_enter_slave":               ret.DtVcsStatStats.Stats.Elect_enter_slave,
			"elect_enter_master":              ret.DtVcsStatStats.Stats.Elect_enter_master,
			"elect_enter_master_take_over":    ret.DtVcsStatStats.Stats.Elect_enter_master_take_over,
			"elect_leave_master_cand":         ret.DtVcsStatStats.Stats.Elect_leave_master_cand,
			"elect_leave_slave":               ret.DtVcsStatStats.Stats.Elect_leave_slave,
			"elect_leave_master":              ret.DtVcsStatStats.Stats.Elect_leave_master,
			"elect_leave_master_take_over":    ret.DtVcsStatStats.Stats.Elect_leave_master_take_over,
			"master_slave_start_err":          ret.DtVcsStatStats.Stats.Master_slave_start_err,
			"master_slave_start":              ret.DtVcsStatStats.Stats.Master_slave_start,
			"master_slave_stop":               ret.DtVcsStatStats.Stats.Master_slave_stop,
			"master_cfg_upd":                  ret.DtVcsStatStats.Stats.Master_cfg_upd,
			"master_cfg_upd_l_fail":           ret.DtVcsStatStats.Stats.Master_cfg_upd_l_fail,
			"master_cfg_upd_r_fail":           ret.DtVcsStatStats.Stats.Master_cfg_upd_r_fail,
			"master_cfg_upd_notif_err":        ret.DtVcsStatStats.Stats.Master_cfg_upd_notif_err,
			"master_cfg_upd_result_err":       ret.DtVcsStatStats.Stats.Master_cfg_upd_result_err,
			"slave_recv_err":                  ret.DtVcsStatStats.Stats.Slave_recv_err,
			"slave_send_err":                  ret.DtVcsStatStats.Stats.Slave_send_err,
			"slave_recv_bytes":                ret.DtVcsStatStats.Stats.Slave_recv_bytes,
			"slave_sent_bytes":                ret.DtVcsStatStats.Stats.Slave_sent_bytes,
			"slave_n_recv":                    ret.DtVcsStatStats.Stats.Slave_n_recv,
			"slave_n_sent":                    ret.DtVcsStatStats.Stats.Slave_n_sent,
			"slave_msg_inval":                 ret.DtVcsStatStats.Stats.Slave_msg_inval,
			"slave_keepalive":                 ret.DtVcsStatStats.Stats.Slave_keepalive,
			"slave_cfg_upd":                   ret.DtVcsStatStats.Stats.Slave_cfg_upd,
			"slave_cfg_upd_fail":              ret.DtVcsStatStats.Stats.Slave_cfg_upd_fail,
			"daemon_n_elec_start":             ret.DtVcsStatStats.Stats.Daemon_n_elec_start,
			"daemon_n_elec_stop":              ret.DtVcsStatStats.Stats.Daemon_n_elec_stop,
			"daemon_recv_err":                 ret.DtVcsStatStats.Stats.Daemon_recv_err,
			"daemon_send_err":                 ret.DtVcsStatStats.Stats.Daemon_send_err,
			"daemon_recv_bytes":               ret.DtVcsStatStats.Stats.Daemon_recv_bytes,
			"daemon_sent_bytes":               ret.DtVcsStatStats.Stats.Daemon_sent_bytes,
			"daemon_n_recv":                   ret.DtVcsStatStats.Stats.Daemon_n_recv,
			"daemon_n_sent":                   ret.DtVcsStatStats.Stats.Daemon_n_sent,
			"daemon_msg_inval":                ret.DtVcsStatStats.Stats.Daemon_msg_inval,
			"daemon_msg_handle_failure":       ret.DtVcsStatStats.Stats.Daemon_msg_handle_failure,
		},
	}
}

func getObjectVcsStatStatsStats(d []interface{}) edpt.VcsStatStatsStats {

	count1 := len(d)
	var ret edpt.VcsStatStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Elect_recv_err = in["elect_recv_err"].(int)
		ret.Elect_send_err = in["elect_send_err"].(int)
		ret.Elect_recv_byte = in["elect_recv_byte"].(int)
		ret.Elect_send_byte = in["elect_send_byte"].(int)
		ret.Elect_pdu_master_recv = in["elect_pdu_master_recv"].(int)
		ret.Elect_pdu_master_cand_recv = in["elect_pdu_master_cand_recv"].(int)
		ret.Elect_pdu_slave_recv = in["elect_pdu_slave_recv"].(int)
		ret.Elect_pdu_master_take_over_recv = in["elect_pdu_master_take_over_recv"].(int)
		ret.Elect_pdu_unknown_recv = in["elect_pdu_unknown_recv"].(int)
		ret.Elect_pdu_master_sent = in["elect_pdu_master_sent"].(int)
		ret.Elect_pdu_master_cand_sent = in["elect_pdu_master_cand_sent"].(int)
		ret.Elect_pdu_slave_sent = in["elect_pdu_slave_sent"].(int)
		ret.Elect_pdu_master_take_over_sent = in["elect_pdu_master_take_over_sent"].(int)
		ret.Elect_pdu_unknown_sent = in["elect_pdu_unknown_sent"].(int)
		ret.Elect_pdu_inval = in["elect_pdu_inval"].(int)
		ret.Elect_pdu_hw_mismatch = in["elect_pdu_hw_mismatch"].(int)
		ret.Elect_pdu_cluster_mismatch = in["elect_pdu_cluster_mismatch"].(int)
		ret.Elect_pdu_dev_id_collision = in["elect_pdu_dev_id_collision"].(int)
		ret.Elect_mc_discard_master = in["elect_mc_discard_master"].(int)
		ret.Elect_mc_replace_master = in["elect_mc_replace_master"].(int)
		ret.Elect_mc_dup_masterr = in["elect_mc_dup_masterr"].(int)
		ret.Elect_mc_reset_timer_by_mc = in["elect_mc_reset_timer_by_mc"].(int)
		ret.Elect_mc_reset_timer_by_mto = in["elect_mc_reset_timer_by_mto"].(int)
		ret.Elect_slave_dup_master = in["elect_slave_dup_master"].(int)
		ret.Elect_slave_discard_challenger = in["elect_slave_discard_challenger"].(int)
		ret.Elect_slave_replace_challenger = in["elect_slave_replace_challenger"].(int)
		ret.Elect_slave_dup_challenger = in["elect_slave_dup_challenger"].(int)
		ret.Elect_slave_discard_neighbour = in["elect_slave_discard_neighbour"].(int)
		ret.Elect_slave_too_many_neighbour = in["elect_slave_too_many_neighbour"].(int)
		ret.Elect_slave_dup_neighbour = in["elect_slave_dup_neighbour"].(int)
		ret.Elect_master_discard_challenger = in["elect_master_discard_challenger"].(int)
		ret.Elect_master_new_challenger = in["elect_master_new_challenger"].(int)
		ret.Elect_master_replace_challenger = in["elect_master_replace_challenger"].(int)
		ret.Elect_master_dup_challenger = in["elect_master_dup_challenger"].(int)
		ret.Elect_master_discard_neighbour = in["elect_master_discard_neighbour"].(int)
		ret.Elect_master_too_many_neighbour = in["elect_master_too_many_neighbour"].(int)
		ret.Elect_master_dup_neighbour = in["elect_master_dup_neighbour"].(int)
		ret.Elect_enter_master_cand_stat = in["elect_enter_master_cand_stat"].(int)
		ret.Elect_enter_slave = in["elect_enter_slave"].(int)
		ret.Elect_enter_master = in["elect_enter_master"].(int)
		ret.Elect_enter_master_take_over = in["elect_enter_master_take_over"].(int)
		ret.Elect_leave_master_cand = in["elect_leave_master_cand"].(int)
		ret.Elect_leave_slave = in["elect_leave_slave"].(int)
		ret.Elect_leave_master = in["elect_leave_master"].(int)
		ret.Elect_leave_master_take_over = in["elect_leave_master_take_over"].(int)
		ret.Master_slave_start_err = in["master_slave_start_err"].(int)
		ret.Master_slave_start = in["master_slave_start"].(int)
		ret.Master_slave_stop = in["master_slave_stop"].(int)
		ret.Master_cfg_upd = in["master_cfg_upd"].(int)
		ret.Master_cfg_upd_l_fail = in["master_cfg_upd_l_fail"].(int)
		ret.Master_cfg_upd_r_fail = in["master_cfg_upd_r_fail"].(int)
		ret.Master_cfg_upd_notif_err = in["master_cfg_upd_notif_err"].(int)
		ret.Master_cfg_upd_result_err = in["master_cfg_upd_result_err"].(int)
		ret.Slave_recv_err = in["slave_recv_err"].(int)
		ret.Slave_send_err = in["slave_send_err"].(int)
		ret.Slave_recv_bytes = in["slave_recv_bytes"].(int)
		ret.Slave_sent_bytes = in["slave_sent_bytes"].(int)
		ret.Slave_n_recv = in["slave_n_recv"].(int)
		ret.Slave_n_sent = in["slave_n_sent"].(int)
		ret.Slave_msg_inval = in["slave_msg_inval"].(int)
		ret.Slave_keepalive = in["slave_keepalive"].(int)
		ret.Slave_cfg_upd = in["slave_cfg_upd"].(int)
		ret.Slave_cfg_upd_fail = in["slave_cfg_upd_fail"].(int)
		ret.Daemon_n_elec_start = in["daemon_n_elec_start"].(int)
		ret.Daemon_n_elec_stop = in["daemon_n_elec_stop"].(int)
		ret.Daemon_recv_err = in["daemon_recv_err"].(int)
		ret.Daemon_send_err = in["daemon_send_err"].(int)
		ret.Daemon_recv_bytes = in["daemon_recv_bytes"].(int)
		ret.Daemon_sent_bytes = in["daemon_sent_bytes"].(int)
		ret.Daemon_n_recv = in["daemon_n_recv"].(int)
		ret.Daemon_n_sent = in["daemon_n_sent"].(int)
		ret.Daemon_msg_inval = in["daemon_msg_inval"].(int)
		ret.Daemon_msg_handle_failure = in["daemon_msg_handle_failure"].(int)
	}
	return ret
}

func dataToEndpointVcsStatStats(d *schema.ResourceData) edpt.VcsStatStats {
	var ret edpt.VcsStatStats

	ret.Stats = getObjectVcsStatStatsStats(d.Get("stats").([]interface{}))
	return ret
}
