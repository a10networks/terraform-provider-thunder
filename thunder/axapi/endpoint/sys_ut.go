

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUt struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    Common SysUtCommon1569 `json:"common"`
    EventList []SysUtEventList `json:"event-list"`
    RunTest SysUtRunTest1570 `json:"run-test"`
    SecondaryName string `json:"secondary-name"`
    StateList []SysUtStateList `json:"state-list"`
    TemplateList []SysUtTemplateList `json:"template-list"`
    TestName string `json:"test-name"`
    Uuid string `json:"uuid"`

	} `json:"sys-ut"`
}


type SysUtCommon1569 struct {
    ProceedOnError int `json:"proceed-on-error"`
    Delay int `json:"delay"`
    Uuid string `json:"uuid"`
}


type SysUtEventList struct {
    EventNumber int `json:"event-number"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    ActionList []SysUtEventListActionList `json:"action-list"`
}


type SysUtEventListActionList struct {
    Direction string `json:"direction"`
    Template string `json:"template"`
    Drop int `json:"drop"`
    Delay int `json:"delay"`
    Uuid string `json:"uuid"`
    L1 SysUtEventListActionListL1 `json:"l1"`
    L2 SysUtEventListActionListL2 `json:"l2"`
    L3 SysUtEventListActionListL3 `json:"l3"`
    Tcp SysUtEventListActionListTcp `json:"tcp"`
    Udp SysUtEventListActionListUdp `json:"udp"`
    IgnoreValidation SysUtEventListActionListIgnoreValidation `json:"ignore-validation"`
}


type SysUtEventListActionListL1 struct {
    EthList []SysUtEventListActionListL1EthList `json:"eth-list"`
    Trunk_list []SysUtEventListActionListL1Trunk_list `json:"trunk_list"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtEventListActionListL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtEventListActionListL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtEventListActionListL2 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtEventListActionListL2MacList `json:"mac-list"`
}


type SysUtEventListActionListL2MacList struct {
    SrcDst string `json:"src-dst"`
    AddressType string `json:"address-type"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Value string `json:"value"`
    Uuid string `json:"uuid"`
}


type SysUtEventListActionListL3 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtEventListActionListL3IpList `json:"ip-list"`
}


type SysUtEventListActionListL3IpList struct {
    SrcDst string `json:"src-dst"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    NatPool string `json:"nat-pool"`
    VirtualServer string `json:"virtual-server"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
}


type SysUtEventListActionListTcp struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    SeqNumber string `json:"seq-number" dval:"valid"`
    AckSeqNumber string `json:"ack-seq-number" dval:"valid"`
    Checksum string `json:"checksum" dval:"valid"`
    Urgent string `json:"urgent" dval:"valid"`
    Window string `json:"window" dval:"valid"`
    Uuid string `json:"uuid"`
    Flags SysUtEventListActionListTcpFlags `json:"flags"`
    Options SysUtEventListActionListTcpOptions `json:"options"`
}


type SysUtEventListActionListTcpFlags struct {
    Syn int `json:"syn"`
    Ack int `json:"ack"`
    Fin int `json:"fin"`
    Rst int `json:"rst"`
    Psh int `json:"psh"`
    Ece int `json:"ece"`
    Urg int `json:"urg"`
    Cwr int `json:"cwr"`
    Uuid string `json:"uuid"`
}


type SysUtEventListActionListTcpOptions struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtEventListActionListUdp struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}


type SysUtEventListActionListIgnoreValidation struct {
    L1 int `json:"l1"`
    L2 int `json:"l2"`
    L3 int `json:"l3"`
    L4 int `json:"l4"`
    All int `json:"all"`
    Uuid string `json:"uuid"`
}


type SysUtRunTest1570 struct {
    Mode string `json:"mode"`
}


type SysUtStateList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    NextStateList []SysUtStateListNextStateList `json:"next-state-list"`
}


type SysUtStateListNextStateList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    CaseList []SysUtStateListNextStateListCaseList `json:"case-list"`
}


type SysUtStateListNextStateListCaseList struct {
    CaseNumber int `json:"case-number"`
    Repeat int `json:"repeat"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    ActionList []SysUtStateListNextStateListCaseListActionList `json:"action-list"`
}


type SysUtStateListNextStateListCaseListActionList struct {
    Direction string `json:"direction"`
    Template string `json:"template"`
    Drop int `json:"drop"`
    Delay int `json:"delay"`
    Uuid string `json:"uuid"`
    L1 SysUtStateListNextStateListCaseListActionListL1 `json:"l1"`
    L2 SysUtStateListNextStateListCaseListActionListL2 `json:"l2"`
    L3 SysUtStateListNextStateListCaseListActionListL3 `json:"l3"`
    Tcp SysUtStateListNextStateListCaseListActionListTcp `json:"tcp"`
    Udp SysUtStateListNextStateListCaseListActionListUdp `json:"udp"`
}


type SysUtStateListNextStateListCaseListActionListL1 struct {
    EthList []SysUtStateListNextStateListCaseListActionListL1EthList `json:"eth-list"`
    Trunk_list []SysUtStateListNextStateListCaseListActionListL1Trunk_list `json:"trunk_list"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtStateListNextStateListCaseListActionListL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtStateListNextStateListCaseListActionListL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtStateListNextStateListCaseListActionListL2 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtStateListNextStateListCaseListActionListL2MacList `json:"mac-list"`
}


type SysUtStateListNextStateListCaseListActionListL2MacList struct {
    SrcDst string `json:"src-dst"`
    AddressType string `json:"address-type"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Value string `json:"value"`
    Uuid string `json:"uuid"`
}


type SysUtStateListNextStateListCaseListActionListL3 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtStateListNextStateListCaseListActionListL3IpList `json:"ip-list"`
}


type SysUtStateListNextStateListCaseListActionListL3IpList struct {
    SrcDst string `json:"src-dst"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
}


type SysUtStateListNextStateListCaseListActionListTcp struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    SeqNumber string `json:"seq-number" dval:"valid"`
    AckSeqNumber string `json:"ack-seq-number" dval:"valid"`
    Checksum string `json:"checksum" dval:"valid"`
    Urgent string `json:"urgent" dval:"valid"`
    Window string `json:"window" dval:"valid"`
    Uuid string `json:"uuid"`
    Flags SysUtStateListNextStateListCaseListActionListTcpFlags `json:"flags"`
    Options SysUtStateListNextStateListCaseListActionListTcpOptions `json:"options"`
}


type SysUtStateListNextStateListCaseListActionListTcpFlags struct {
    Syn int `json:"syn"`
    Ack int `json:"ack"`
    Fin int `json:"fin"`
    Rst int `json:"rst"`
    Psh int `json:"psh"`
    Ece int `json:"ece"`
    Urg int `json:"urg"`
    Cwr int `json:"cwr"`
    Uuid string `json:"uuid"`
}


type SysUtStateListNextStateListCaseListActionListTcpOptions struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtStateListNextStateListCaseListActionListUdp struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    IgnoreValidation SysUtTemplateListIgnoreValidation `json:"ignore-validation"`
    L1 SysUtTemplateListL1 `json:"l1"`
    L2 SysUtTemplateListL2 `json:"l2"`
    L3 SysUtTemplateListL3 `json:"l3"`
    Tcp SysUtTemplateListTcp `json:"tcp"`
    Udp SysUtTemplateListUdp `json:"udp"`
}


type SysUtTemplateListIgnoreValidation struct {
    L1 int `json:"l1"`
    L2 int `json:"l2"`
    L3 int `json:"l3"`
    L4 int `json:"l4"`
    All int `json:"all"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListL1 struct {
    EthList []SysUtTemplateListL1EthList `json:"eth-list"`
    Trunk_list []SysUtTemplateListL1Trunk_list `json:"trunk_list"`
    Drop int `json:"drop"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtTemplateListL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtTemplateListL2 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtTemplateListL2MacList `json:"mac-list"`
}


type SysUtTemplateListL2MacList struct {
    SrcDst string `json:"src-dst"`
    AddressType string `json:"address-type"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Value string `json:"value"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListL3 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtTemplateListL3IpList `json:"ip-list"`
}


type SysUtTemplateListL3IpList struct {
    SrcDst string `json:"src-dst"`
    Ipv4StartAddress string `json:"ipv4-start-address"`
    Ipv4EndAddress string `json:"ipv4-end-address"`
    Ipv6StartAddress string `json:"ipv6-start-address"`
    Ipv6EndAddress string `json:"ipv6-end-address"`
    VirtualServer string `json:"virtual-server"`
    NatPool string `json:"nat-pool"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListTcp struct {
    SrcPortRange []SysUtTemplateListTcpSrcPortRange `json:"src-port-range"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    SeqNumber string `json:"seq-number" dval:"valid"`
    AckSeqNumber string `json:"ack-seq-number" dval:"valid"`
    Checksum string `json:"checksum" dval:"valid"`
    Urgent string `json:"urgent" dval:"valid"`
    Window string `json:"window" dval:"valid"`
    Uuid string `json:"uuid"`
    Flags SysUtTemplateListTcpFlags `json:"flags"`
    Options SysUtTemplateListTcpOptions `json:"options"`
}


type SysUtTemplateListTcpSrcPortRange struct {
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
}


type SysUtTemplateListTcpFlags struct {
    Syn int `json:"syn"`
    Ack int `json:"ack"`
    Fin int `json:"fin"`
    Rst int `json:"rst"`
    Psh int `json:"psh"`
    Ece int `json:"ece"`
    Urg int `json:"urg"`
    Cwr int `json:"cwr"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListTcpOptions struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListUdp struct {
    SrcPortRange []SysUtTemplateListUdpSrcPortRange `json:"src-port-range"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateListUdpSrcPortRange struct {
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
}

func (p *SysUt) GetId() string{
    return "1"
}

func (p *SysUt) getPath() string{
    return "sys-ut"
}

func (p *SysUt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUt::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *SysUt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUt::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SysUt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUt::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SysUt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
