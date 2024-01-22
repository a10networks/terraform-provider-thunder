

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatAlgPptpStats struct {
    
    Stats IpNatAlgPptpStatsStats `json:"stats"`

}
type DataIpNatAlgPptpStats struct {
    DtIpNatAlgPptpStats IpNatAlgPptpStats `json:"pptp"`
}


type IpNatAlgPptpStatsStats struct {
    CurrentSmpSessions int `json:"current-smp-sessions"`
    CurrentGreSessions int `json:"current-gre-sessions"`
    SmpSessionCreationFailure int `json:"smp-session-creation-failure"`
    TruncatedPnsMessage int `json:"truncated-pns-message"`
    TruncatedPacMessage int `json:"truncated-pac-message"`
    MismatchedPnsCallId int `json:"mismatched-pns-call-id"`
    MismatchedPacCallId int `json:"mismatched-pac-call-id"`
    RetransmittedPnsMessage int `json:"retransmitted-pns-message"`
    RetransmittedPacMessage int `json:"retransmitted-pac-message"`
    TruncatedGrePacket int `json:"truncated-gre-packet"`
    UnknownGreVersion int `json:"unknown-gre-version"`
    NoMatchingGreSession int `json:"no-matching-gre-session"`
}

func (p *IpNatAlgPptpStats) GetId() string{
    return "1"
}

func (p *IpNatAlgPptpStats) getPath() string{
    return "ip/nat/alg/pptp/stats"
}

func (p *IpNatAlgPptpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpNatAlgPptpStats,error) {
logger.Println("IpNatAlgPptpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpNatAlgPptpStats
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
