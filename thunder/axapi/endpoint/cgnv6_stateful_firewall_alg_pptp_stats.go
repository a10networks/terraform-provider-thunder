

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallAlgPptpStats struct {
    
    Stats Cgnv6StatefulFirewallAlgPptpStatsStats `json:"stats"`

}
type DataCgnv6StatefulFirewallAlgPptpStats struct {
    DtCgnv6StatefulFirewallAlgPptpStats Cgnv6StatefulFirewallAlgPptpStats `json:"pptp"`
}


type Cgnv6StatefulFirewallAlgPptpStatsStats struct {
    CallsEstablished int `json:"calls-established"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    GreSessionCreated int `json:"gre-session-created"`
    GreSessionFreed int `json:"gre-session-freed"`
}

func (p *Cgnv6StatefulFirewallAlgPptpStats) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallAlgPptpStats) getPath() string{
    return "cgnv6/stateful-firewall/alg/pptp/stats"
}

func (p *Cgnv6StatefulFirewallAlgPptpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6StatefulFirewallAlgPptpStats,error) {
logger.Println("Cgnv6StatefulFirewallAlgPptpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6StatefulFirewallAlgPptpStats
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
