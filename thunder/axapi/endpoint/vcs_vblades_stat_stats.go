

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VcsVbladesStatStats struct {
    
    Stats VcsVbladesStatStatsStats `json:"stats"`
    VbladeId int `json:"vblade-id"`

}
type DataVcsVbladesStatStats struct {
    DtVcsVbladesStatStats VcsVbladesStatStats `json:"stat"`
}


type VcsVbladesStatStatsStats struct {
    Slave_recv_err int `json:"slave_recv_err"`
    Slave_send_err int `json:"slave_send_err"`
    Slave_recv_bytes int `json:"slave_recv_bytes"`
    Slave_sent_bytes int `json:"slave_sent_bytes"`
    Slave_n_recv int `json:"slave_n_recv"`
    Slave_n_sent int `json:"slave_n_sent"`
    Slave_msg_inval int `json:"slave_msg_inval"`
    Slave_keepalive int `json:"slave_keepalive"`
    Slave_cfg_upd int `json:"slave_cfg_upd"`
    Slave_cfg_upd_l1_fail int `json:"slave_cfg_upd_l1_fail"`
    Slave_cfg_upd_r_fail int `json:"slave_cfg_upd_r_fail"`
    Slave_cfg_upd_l2_fail int `json:"slave_cfg_upd_l2_fail"`
    Slave_cfg_upd_notif_err int `json:"slave_cfg_upd_notif_err"`
    Slave_cfg_upd_result_err int `json:"slave_cfg_upd_result_err"`
}

func (p *VcsVbladesStatStats) GetId() string{
    return "1"
}

func (p *VcsVbladesStatStats) getPath() string{
    
    return "vcs-vblades/stat/" +strconv.Itoa(p.VbladeId)+"/stats"
}

func (p *VcsVbladesStatStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVcsVbladesStatStats,error) {
logger.Println("VcsVbladesStatStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVcsVbladesStatStats
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
