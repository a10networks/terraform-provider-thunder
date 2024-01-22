

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlgPptpStats struct {
    
    Stats FwAlgPptpStatsStats `json:"stats"`

}
type DataFwAlgPptpStats struct {
    DtFwAlgPptpStats FwAlgPptpStats `json:"pptp"`
}


type FwAlgPptpStatsStats struct {
    CallsEstablished int `json:"calls-established"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    GreSessionCreated int `json:"gre-session-created"`
    GreSessionFreed int `json:"gre-session-freed"`
}

func (p *FwAlgPptpStats) GetId() string{
    return "1"
}

func (p *FwAlgPptpStats) getPath() string{
    return "fw/alg/pptp/stats"
}

func (p *FwAlgPptpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAlgPptpStats,error) {
logger.Println("FwAlgPptpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAlgPptpStats
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
