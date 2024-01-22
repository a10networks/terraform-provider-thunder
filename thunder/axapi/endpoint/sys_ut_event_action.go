

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtEventAction struct {
	Inst struct {

    Delay int `json:"delay"`
    Direction string `json:"direction"`
    Drop int `json:"drop"`
    IgnoreValidation SysUtEventActionIgnoreValidation1528 `json:"ignore-validation"`
    L1 SysUtEventActionL11529 `json:"l1"`
    L2 SysUtEventActionL21532 `json:"l2"`
    L3 SysUtEventActionL31534 `json:"l3"`
    Tcp SysUtEventActionTcp1536 `json:"tcp"`
    Template string `json:"template"`
    Udp SysUtEventActionUdp1539 `json:"udp"`
    Uuid string `json:"uuid"`
    EventNumber string 

	} `json:"action"`
}


type SysUtEventActionIgnoreValidation1528 struct {
    L1 int `json:"l1"`
    L2 int `json:"l2"`
    L3 int `json:"l3"`
    L4 int `json:"l4"`
    All int `json:"all"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionL11529 struct {
    EthList []SysUtEventActionL1EthList1530 `json:"eth-list"`
    Trunk_list []SysUtEventActionL1Trunk_list1531 `json:"trunk_list"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionL1EthList1530 struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtEventActionL1Trunk_list1531 struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtEventActionL21532 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtEventActionL2MacList1533 `json:"mac-list"`
}


type SysUtEventActionL2MacList1533 struct {
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


type SysUtEventActionL31534 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtEventActionL3IpList1535 `json:"ip-list"`
}


type SysUtEventActionL3IpList1535 struct {
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


type SysUtEventActionTcp1536 struct {
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
    Flags SysUtEventActionTcpFlags1537 `json:"flags"`
    Options SysUtEventActionTcpOptions1538 `json:"options"`
}


type SysUtEventActionTcpFlags1537 struct {
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


type SysUtEventActionTcpOptions1538 struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionUdp1539 struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}

func (p *SysUtEventAction) GetId() string{
    return p.Inst.Direction
}

func (p *SysUtEventAction) getPath() string{
    return "sys-ut/event/" +p.Inst.EventNumber + "/action"
}

func (p *SysUtEventAction) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventAction::Post")
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

func (p *SysUtEventAction) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventAction::Get")
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
func (p *SysUtEventAction) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventAction::Put")
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

func (p *SysUtEventAction) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEventAction::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
