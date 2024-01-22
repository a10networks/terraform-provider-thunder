

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsStatStats struct {
    
    Stats VcsStatStatsStats `json:"stats"`

}
type DataVcsStatStats struct {
    DtVcsStatStats VcsStatStats `json:"stat"`
}


type VcsStatStatsStats struct {
    Elect_recv_err int `json:"elect_recv_err"`
    Elect_send_err int `json:"elect_send_err"`
    Elect_recv_byte int `json:"elect_recv_byte"`
    Elect_send_byte int `json:"elect_send_byte"`
    Elect_pdu_master_recv int `json:"elect_pdu_master_recv"`
    Elect_pdu_master_cand_recv int `json:"elect_pdu_master_cand_recv"`
    Elect_pdu_slave_recv int `json:"elect_pdu_slave_recv"`
    Elect_pdu_master_take_over_recv int `json:"elect_pdu_master_take_over_recv"`
    Elect_pdu_unknown_recv int `json:"elect_pdu_unknown_recv"`
    Elect_pdu_master_sent int `json:"elect_pdu_master_sent"`
    Elect_pdu_master_cand_sent int `json:"elect_pdu_master_cand_sent"`
    Elect_pdu_slave_sent int `json:"elect_pdu_slave_sent"`
    Elect_pdu_master_take_over_sent int `json:"elect_pdu_master_take_over_sent"`
    Elect_pdu_unknown_sent int `json:"elect_pdu_unknown_sent"`
    Elect_pdu_inval int `json:"elect_pdu_inval"`
    Elect_pdu_hw_mismatch int `json:"elect_pdu_hw_mismatch"`
    Elect_pdu_cluster_mismatch int `json:"elect_pdu_cluster_mismatch"`
    Elect_pdu_dev_id_collision int `json:"elect_pdu_dev_id_collision"`
    Elect_mc_discard_master int `json:"elect_mc_discard_master"`
    Elect_mc_replace_master int `json:"elect_mc_replace_master"`
    Elect_mc_dup_masterr int `json:"elect_mc_dup_masterr"`
    Elect_mc_reset_timer_by_mc int `json:"elect_mc_reset_timer_by_mc"`
    Elect_mc_reset_timer_by_mto int `json:"elect_mc_reset_timer_by_mto"`
    Elect_slave_dup_master int `json:"elect_slave_dup_master"`
    Elect_slave_discard_challenger int `json:"elect_slave_discard_challenger"`
    Elect_slave_replace_challenger int `json:"elect_slave_replace_challenger"`
    Elect_slave_dup_challenger int `json:"elect_slave_dup_challenger"`
    Elect_slave_discard_neighbour int `json:"elect_slave_discard_neighbour"`
    Elect_slave_too_many_neighbour int `json:"elect_slave_too_many_neighbour"`
    Elect_slave_dup_neighbour int `json:"elect_slave_dup_neighbour"`
    Elect_master_discard_challenger int `json:"elect_master_discard_challenger"`
    Elect_master_new_challenger int `json:"elect_master_new_challenger"`
    Elect_master_replace_challenger int `json:"elect_master_replace_challenger"`
    Elect_master_dup_challenger int `json:"elect_master_dup_challenger"`
    Elect_master_discard_neighbour int `json:"elect_master_discard_neighbour"`
    Elect_master_too_many_neighbour int `json:"elect_master_too_many_neighbour"`
    Elect_master_dup_neighbour int `json:"elect_master_dup_neighbour"`
    Elect_enter_master_cand_stat int `json:"elect_enter_master_cand_stat"`
    Elect_enter_slave int `json:"elect_enter_slave"`
    Elect_enter_master int `json:"elect_enter_master"`
    Elect_enter_master_take_over int `json:"elect_enter_master_take_over"`
    Elect_leave_master_cand int `json:"elect_leave_master_cand"`
    Elect_leave_slave int `json:"elect_leave_slave"`
    Elect_leave_master int `json:"elect_leave_master"`
    Elect_leave_master_take_over int `json:"elect_leave_master_take_over"`
    Master_slave_start_err int `json:"master_slave_start_err"`
    Master_slave_start int `json:"master_slave_start"`
    Master_slave_stop int `json:"master_slave_stop"`
    Master_cfg_upd int `json:"master_cfg_upd"`
    Master_cfg_upd_l_fail int `json:"master_cfg_upd_l_fail"`
    Master_cfg_upd_r_fail int `json:"master_cfg_upd_r_fail"`
    Master_cfg_upd_notif_err int `json:"master_cfg_upd_notif_err"`
    Master_cfg_upd_result_err int `json:"master_cfg_upd_result_err"`
    Slave_recv_err int `json:"slave_recv_err"`
    Slave_send_err int `json:"slave_send_err"`
    Slave_recv_bytes int `json:"slave_recv_bytes"`
    Slave_sent_bytes int `json:"slave_sent_bytes"`
    Slave_n_recv int `json:"slave_n_recv"`
    Slave_n_sent int `json:"slave_n_sent"`
    Slave_msg_inval int `json:"slave_msg_inval"`
    Slave_keepalive int `json:"slave_keepalive"`
    Slave_cfg_upd int `json:"slave_cfg_upd"`
    Slave_cfg_upd_fail int `json:"slave_cfg_upd_fail"`
    Daemon_n_elec_start int `json:"daemon_n_elec_start"`
    Daemon_n_elec_stop int `json:"daemon_n_elec_stop"`
    Daemon_recv_err int `json:"daemon_recv_err"`
    Daemon_send_err int `json:"daemon_send_err"`
    Daemon_recv_bytes int `json:"daemon_recv_bytes"`
    Daemon_sent_bytes int `json:"daemon_sent_bytes"`
    Daemon_n_recv int `json:"daemon_n_recv"`
    Daemon_n_sent int `json:"daemon_n_sent"`
    Daemon_msg_inval int `json:"daemon_msg_inval"`
    Daemon_msg_handle_failure int `json:"daemon_msg_handle_failure"`
}

func (p *VcsStatStats) GetId() string{
    return "1"
}

func (p *VcsStatStats) getPath() string{
    return "vcs/stat/stats"
}

func (p *VcsStatStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVcsStatStats,error) {
logger.Println("VcsStatStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVcsStatStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
