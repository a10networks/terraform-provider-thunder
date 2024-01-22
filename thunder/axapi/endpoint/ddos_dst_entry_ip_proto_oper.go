

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryIpProtoOper struct {
    
    IpFilteringPolicyOper DdosDstEntryIpProtoOperIpFilteringPolicyOper `json:"ip-filtering-policy-oper"`
    Oper DdosDstEntryIpProtoOperOper `json:"oper"`
    PortNum int `json:"port-num"`
    DstEntryName string 

}
type DataDdosDstEntryIpProtoOper struct {
    DtDdosDstEntryIpProtoOper DdosDstEntryIpProtoOper `json:"ip-proto"`
}


type DdosDstEntryIpProtoOperIpFilteringPolicyOper struct {
    Oper DdosDstEntryIpProtoOperIpFilteringPolicyOperOper `json:"oper"`
}


type DdosDstEntryIpProtoOperIpFilteringPolicyOperOper struct {
    RuleList []DdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList `json:"rule-list"`
}


type DdosDstEntryIpProtoOperIpFilteringPolicyOperOperRuleList struct {
    Seq int `json:"seq"`
    Hits int `json:"hits"`
}


type DdosDstEntryIpProtoOperOper struct {
    Ddos_entry_list []DdosDstEntryIpProtoOperOperDdos_entry_list `json:"ddos_entry_list"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    ReportingStatus int `json:"reporting-status"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    AllIpProtos int `json:"all-ip-protos"`
    PortProtocol string `json:"port-protocol"`
    AppStat int `json:"app-stat"`
    SflowSourceId int `json:"sflow-source-id"`
    HwBlacklisted string `json:"hw-blacklisted"`
    SuffixRequestRate int `json:"suffix-request-rate"`
    DomainName string `json:"domain-name"`
}


type DdosDstEntryIpProtoOperOperDdos_entry_list struct {
    DstAddressStr string `json:"dst-address-str"`
    SrcAddressStr string `json:"src-address-str"`
    PortStr string `json:"port-str"`
    StateStr string `json:"state-str"`
    LevelStr string `json:"level-str"`
    CurrentConnections string `json:"current-connections"`
    ConnectionLimit string `json:"connection-limit"`
    CurrentConnectionRate string `json:"current-connection-rate"`
    ConnectionRateLimit string `json:"connection-rate-limit"`
    CurrentPacketRate string `json:"current-packet-rate"`
    PacketRateLimit string `json:"packet-rate-limit"`
    CurrentKbitRate string `json:"current-kBit-rate"`
    KbitRateLimit string `json:"kBit-rate-limit"`
    CurrentFragPacketRate string `json:"current-frag-packet-rate"`
    FragPacketRateLimit string `json:"frag-packet-rate-limit"`
    CurrentAppStat1 string `json:"current-app-stat1"`
    AppStat1Limit string `json:"app-stat1-limit"`
    CurrentAppStat2 string `json:"current-app-stat2"`
    AppStat2Limit string `json:"app-stat2-limit"`
    CurrentAppStat3 string `json:"current-app-stat3"`
    AppStat3Limit string `json:"app-stat3-limit"`
    CurrentAppStat4 string `json:"current-app-stat4"`
    AppStat4Limit string `json:"app-stat4-limit"`
    CurrentAppStat5 string `json:"current-app-stat5"`
    AppStat5Limit string `json:"app-stat5-limit"`
    CurrentAppStat6 string `json:"current-app-stat6"`
    AppStat6Limit string `json:"app-stat6-limit"`
    CurrentAppStat7 string `json:"current-app-stat7"`
    AppStat7Limit string `json:"app-stat7-limit"`
    CurrentAppStat8 string `json:"current-app-stat8"`
    AppStat8Limit string `json:"app-stat8-limit"`
    AgeStr string `json:"age-str"`
    LockupTimeStr string `json:"lockup-time-str"`
    DynamicEntryCount string `json:"dynamic-entry-count"`
    DynamicEntryLimit string `json:"dynamic-entry-limit"`
    SflowSourceId string `json:"sflow-source-id"`
    DebugStr string `json:"debug-str"`
}

func (p *DdosDstEntryIpProtoOper) GetId() string{
    return "1"
}

func (p *DdosDstEntryIpProtoOper) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/ip-proto/" +strconv.Itoa(p.PortNum)+"/oper"
}

func (p *DdosDstEntryIpProtoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntryIpProtoOper,error) {
logger.Println("DdosDstEntryIpProtoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntryIpProtoOper
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
