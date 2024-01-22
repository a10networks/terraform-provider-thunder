

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFtpCtlStats struct {
    
    Stats SlbFtpCtlStatsStats `json:"stats"`

}
type DataSlbFtpCtlStats struct {
    DtSlbFtpCtlStats SlbFtpCtlStats `json:"ftp-ctl"`
}


type SlbFtpCtlStatsStats struct {
    Sessions_num int `json:"sessions_num"`
    Alg_pkts_num int `json:"alg_pkts_num"`
    Alg_pkts_xmitted_num int `json:"alg_pkts_xmitted_num"`
    Alg_port_helper_created int `json:"alg_port_helper_created"`
    Alg_pasv_helper_created int `json:"alg_pasv_helper_created"`
    Alg_port_helper_freed_unused int `json:"alg_port_helper_freed_unused"`
    Alg_pasv_helper_freed_unused int `json:"alg_pasv_helper_freed_unused"`
    Alg_port_helper_nat_free int `json:"alg_port_helper_nat_free"`
}

func (p *SlbFtpCtlStats) GetId() string{
    return "1"
}

func (p *SlbFtpCtlStats) getPath() string{
    return "slb/ftp-ctl/stats"
}

func (p *SlbFtpCtlStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFtpCtlStats,error) {
logger.Println("SlbFtpCtlStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFtpCtlStats
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
