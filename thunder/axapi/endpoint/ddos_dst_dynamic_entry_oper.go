

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntryOper struct {
    
    Oper DdosDstDynamicEntryOperOper `json:"oper"`

}
type DataDdosDstDynamicEntryOper struct {
    DtDdosDstDynamicEntryOper DdosDstDynamicEntryOper `json:"dynamic-entry"`
}


type DdosDstDynamicEntryOperOper struct {
    Ddos_entry_list []DdosDstDynamicEntryOperOperDdos_entry_list `json:"ddos_entry_list"`
    Ip_conn_total int `json:"ip_conn_total"`
    Ipv6_conn_total int `json:"ipv6_conn_total"`
    EntryDisplayedCount int `json:"entry-displayed-count"`
    ServiceDisplayedCount int `json:"service-displayed-count"`
    Ipv6 string `json:"ipv6"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
    OverflowPolicy string `json:"overflow-policy"`
    IpProtoNum int `json:"ip-proto-num"`
    L4TypeStr string `json:"l4-type-str"`
    PortNum int `json:"port-num"`
    PortRangeStart int `json:"port-range-start"`
    PortRangeEnd int `json:"port-range-end"`
    SrcPortNum int `json:"src-port-num"`
    SrcPortRangeStart int `json:"src-port-range-start"`
    SrcPortRangeEnd int `json:"src-port-range-end"`
    Protocol string `json:"protocol"`
    SportProtocol string `json:"sport-protocol"`
    AppStat int `json:"app-stat"`
    AllEntries int `json:"all-entries"`
    AllIpProtos int `json:"all-ip-protos"`
    AllL4Types int `json:"all-l4-types"`
    AllPorts int `json:"all-ports"`
    AllSrcPorts int `json:"all-src-ports"`
    BlackHoled int `json:"black-holed"`
    Exceeded int `json:"exceeded"`
    MaxCount int `json:"max-count"`
    ResourceUsage int `json:"resource-usage"`
    HwBlacklisted int `json:"hw-blacklisted"`
}


type DdosDstDynamicEntryOperOperDdos_entry_list struct {
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

func (p *DdosDstDynamicEntryOper) GetId() string{
    return "1"
}

func (p *DdosDstDynamicEntryOper) getPath() string{
    return "ddos/dst/dynamic-entry/oper"
}

func (p *DdosDstDynamicEntryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstDynamicEntryOper,error) {
logger.Println("DdosDstDynamicEntryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstDynamicEntryOper
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
