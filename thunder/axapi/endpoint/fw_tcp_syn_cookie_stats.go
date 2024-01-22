

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwTcpSynCookieStats struct {
    
    Stats FwTcpSynCookieStatsStats `json:"stats"`

}
type DataFwTcpSynCookieStats struct {
    DtFwTcpSynCookieStats FwTcpSynCookieStats `json:"syn-cookie"`
}


type FwTcpSynCookieStatsStats struct {
    Syn_ack_sent int `json:"syn_ack_sent"`
    Verification_passed int `json:"verification_passed"`
    Verification_failed int `json:"verification_failed"`
}

func (p *FwTcpSynCookieStats) GetId() string{
    return "1"
}

func (p *FwTcpSynCookieStats) getPath() string{
    return "fw/tcp/syn-cookie/stats"
}

func (p *FwTcpSynCookieStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwTcpSynCookieStats,error) {
logger.Println("FwTcpSynCookieStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwTcpSynCookieStats
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
