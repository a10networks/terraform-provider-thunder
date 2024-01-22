

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtTemplate struct {
	Inst struct {

    IgnoreValidation SysUtTemplateIgnoreValidation1555 `json:"ignore-validation"`
    L1 SysUtTemplateL11556 `json:"l1"`
    L2 SysUtTemplateL21559 `json:"l2"`
    L3 SysUtTemplateL31561 `json:"l3"`
    Name string `json:"name"`
    Tcp SysUtTemplateTcp1563 `json:"tcp"`
    Udp SysUtTemplateUdp1567 `json:"udp"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"template"`
}


type SysUtTemplateIgnoreValidation1555 struct {
    L1 int `json:"l1"`
    L2 int `json:"l2"`
    L3 int `json:"l3"`
    L4 int `json:"l4"`
    All int `json:"all"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateL11556 struct {
    EthList []SysUtTemplateL1EthList1557 `json:"eth-list"`
    Trunk_list []SysUtTemplateL1Trunk_list1558 `json:"trunk_list"`
    Drop int `json:"drop"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateL1EthList1557 struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtTemplateL1Trunk_list1558 struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtTemplateL21559 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtTemplateL2MacList1560 `json:"mac-list"`
}


type SysUtTemplateL2MacList1560 struct {
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


type SysUtTemplateL31561 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtTemplateL3IpList1562 `json:"ip-list"`
}


type SysUtTemplateL3IpList1562 struct {
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


type SysUtTemplateTcp1563 struct {
    SrcPortRange []SysUtTemplateTcpSrcPortRange1564 `json:"src-port-range"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    SeqNumber string `json:"seq-number" dval:"valid"`
    AckSeqNumber string `json:"ack-seq-number" dval:"valid"`
    Checksum string `json:"checksum" dval:"valid"`
    Urgent string `json:"urgent" dval:"valid"`
    Window string `json:"window" dval:"valid"`
    Uuid string `json:"uuid"`
    Flags SysUtTemplateTcpFlags1565 `json:"flags"`
    Options SysUtTemplateTcpOptions1566 `json:"options"`
}


type SysUtTemplateTcpSrcPortRange1564 struct {
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
}


type SysUtTemplateTcpFlags1565 struct {
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


type SysUtTemplateTcpOptions1566 struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateUdp1567 struct {
    SrcPortRange []SysUtTemplateUdpSrcPortRange1568 `json:"src-port-range"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}


type SysUtTemplateUdpSrcPortRange1568 struct {
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
}

func (p *SysUtTemplate) GetId() string{
    return p.Inst.Name
}

func (p *SysUtTemplate) getPath() string{
    return "sys-ut/template"
}

func (p *SysUtTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplate::Post")
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

func (p *SysUtTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplate::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *SysUtTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplate::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *SysUtTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
