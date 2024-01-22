

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6PcpStats struct {
    
    Stats Cgnv6PcpStatsStats `json:"stats"`

}
type DataCgnv6PcpStats struct {
    DtCgnv6PcpStats Cgnv6PcpStats `json:"pcp"`
}


type Cgnv6PcpStatsStats struct {
    PacketsRcv int `json:"packets-rcv"`
    LsnMapProcessSuccess int `json:"lsn-map-process-success"`
    DsliteMapProcessSuccess int `json:"dslite-map-process-success"`
    Nat64MapProcessSuccess int `json:"nat64-map-process-success"`
    LsnPeerProcessSuccess int `json:"lsn-peer-process-success"`
    DslitePeerProcessSuccess int `json:"dslite-peer-process-success"`
    Nat64PeerProcessSuccess int `json:"nat64-peer-process-success"`
    LsnAnnounceProcessSuccess int `json:"lsn-announce-process-success"`
    DsliteAnnounceProcessSuccess int `json:"dslite-announce-process-success"`
    Nat64AnnounceProcessSuccess int `json:"nat64-announce-process-success"`
    PktNotRequestDrop int `json:"pkt-not-request-drop"`
    PktTooShortDrop int `json:"pkt-too-short-drop"`
    NorouteDrop int `json:"noroute-drop"`
    UnsupportedVersion int `json:"unsupported-version"`
    NotAuthorized int `json:"not-authorized"`
    MalformRequest int `json:"malform-request"`
    UnsuppOpcode int `json:"unsupp-opcode"`
    UnsuppOption int `json:"unsupp-option"`
    MalformOption int `json:"malform-option"`
    NoResources int `json:"no-resources"`
    UnsuppProtocol int `json:"unsupp-protocol"`
    UserQuotaExceeded int `json:"user-quota-exceeded"`
    CannotProvideSuggest int `json:"cannot-provide-suggest"`
    AddressMismatch int `json:"address-mismatch"`
    ExcessiveRemotePeers int `json:"excessive-remote-peers"`
    PktNotFromNatInside int `json:"pkt-not-from-nat-inside"`
    L4ProcessError int `json:"l4-process-error"`
    InternalErrorDrop int `json:"internal-error-drop"`
    Unsol_ance_sent_succ int `json:"unsol_ance_sent_succ"`
    Unsol_ance_sent_fail int `json:"unsol_ance_sent_fail"`
    Ha_sync_epoch_sent int `json:"ha_sync_epoch_sent"`
    Ha_sync_epoch_rcv int `json:"ha_sync_epoch_rcv"`
}

func (p *Cgnv6PcpStats) GetId() string{
    return "1"
}

func (p *Cgnv6PcpStats) getPath() string{
    return "cgnv6/pcp/stats"
}

func (p *Cgnv6PcpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6PcpStats,error) {
logger.Println("Cgnv6PcpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6PcpStats
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
