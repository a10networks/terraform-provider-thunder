

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosProtectStats struct {
    
    Stats DdosProtectStatsStats `json:"stats"`

}
type DataDdosProtectStats struct {
    DtDdosProtectStats DdosProtectStats `json:"protect"`
}


type DdosProtectStatsStats struct {
    Rr_sess_lookup int `json:"rr_sess_lookup"`
    Dcmsg_pkt int `json:"dcmsg_pkt"`
    Runtime_enable int `json:"runtime_enable"`
    Runtime_disable int `json:"runtime_disable"`
}

func (p *DdosProtectStats) GetId() string{
    return "1"
}

func (p *DdosProtectStats) getPath() string{
    return "ddos/protect/stats"
}

func (p *DdosProtectStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosProtectStats,error) {
logger.Println("DdosProtectStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosProtectStats
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
