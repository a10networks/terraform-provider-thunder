

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgPptpStats struct {
    
    Stats Cgnv6LsnAlgPptpStatsStats `json:"stats"`

}
type DataCgnv6LsnAlgPptpStats struct {
    DtCgnv6LsnAlgPptpStats Cgnv6LsnAlgPptpStats `json:"pptp"`
}


type Cgnv6LsnAlgPptpStatsStats struct {
    CallsEstablished int `json:"calls-established"`
    MismatchedPnsCallId int `json:"mismatched-pns-call-id"`
    GreSessionsCreated int `json:"gre-sessions-created"`
    GreSessionsFreed int `json:"gre-sessions-freed"`
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
}

func (p *Cgnv6LsnAlgPptpStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgPptpStats) getPath() string{
    return "cgnv6/lsn/alg/pptp/stats"
}

func (p *Cgnv6LsnAlgPptpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnAlgPptpStats,error) {
logger.Println("Cgnv6LsnAlgPptpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnAlgPptpStats
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
