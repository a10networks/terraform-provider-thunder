

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoTcpUdpOper struct {
    
    IpFilteringPolicyOper DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstZoneIpProtoProtoTcpUdpOperOper `json:"oper"`
    Protocol string `json:"protocol"`
    ZoneName string 

}
type DataDdosDstZoneIpProtoProtoTcpUdpOper struct {
    DtDdosDstZoneIpProtoProtoTcpUdpOper DdosDstZoneIpProtoProtoTcpUdpOper `json:"proto-tcp-udp"`
}


type DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOper struct {
    Oper DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstZoneIpProtoProtoTcpUdpOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstZoneIpProtoProtoTcpUdpOperOper struct {
    Ddos_entry_list []DdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    Sources int `json:"sources"`
    OverflowPolicy int `json:"overflow-policy"`
    SourcesAllEntries int `json:"sources-all-entries"`
    ClassList string `json:"class-list"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    Ipv6 string `json:"ipv6"`
    Exceeded int `json:"exceeded"`
    BlackListed int `json:"black-listed"`
    WhiteListed int `json:"white-listed"`
    Authenticated int `json:"authenticated"`
    Level int `json:"level"`
    AppStat int `json:"app-stat"`
    Indicators int `json:"indicators"`
    IndicatorDetail int `json:"indicator-detail"`
    HwBlacklisted int `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstZoneIpProtoProtoTcpUdpOperOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    BwState string `json:"bw-state"`
    Is_auth_passed string `json:"is_auth_passed"`
    Level int `json:"level"`
    BlReasoningRcode string `json:"bl-reasoning-rcode"`
    BlReasoningTimestamp string `json:"bl-reasoning-timestamp"`
    CurrentConnections string `json:"current-connections"`
    IsConnectionsExceed int `json:"is-connections-exceed"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    IsConnectionRateExceed int `json:"is-connection-rate-exceed"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    IsPacketRateExceed int `json:"is-packet-rate-exceed"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    IsKbitRateExceed int `json:"is-kBit-rate-exceed"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    IsFragPacketRateExceed int `json:"is-frag-packet-rate-exceed"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    Age int `json:"age"`
    LockupTime int `json:"lockup-time"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId int `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}

func (p *DdosDstZoneIpProtoProtoTcpUdpOper) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoTcpUdpOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-tcp-udp/"+p.Protocol+"/oper"
}

func (p *DdosDstZoneIpProtoProtoTcpUdpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoTcpUdpOper,error) {
logger.Println("DdosDstZoneIpProtoProtoTcpUdpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoTcpUdpOper
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
