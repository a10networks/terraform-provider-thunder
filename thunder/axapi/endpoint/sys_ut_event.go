

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SysUtEvent struct {
	Inst struct {

    ActionList []SysUtEventActionList `json:"action-list"`
    EventNumber int `json:"event-number"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"event"`
}


type SysUtEventActionList struct {
    Direction string `json:"direction"`
    Template string `json:"template"`
    Drop int `json:"drop"`
    Delay int `json:"delay"`
    Uuid string `json:"uuid"`
    L1 SysUtEventActionListL1 `json:"l1"`
    L2 SysUtEventActionListL2 `json:"l2"`
    L3 SysUtEventActionListL3 `json:"l3"`
    Tcp SysUtEventActionListTcp `json:"tcp"`
    Udp SysUtEventActionListUdp `json:"udp"`
    IgnoreValidation SysUtEventActionListIgnoreValidation `json:"ignore-validation"`
}


type SysUtEventActionListL1 struct {
    EthList []SysUtEventActionListL1EthList `json:"eth-list"`
    Trunk_list []SysUtEventActionListL1Trunk_list `json:"trunk_list"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionListL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtEventActionListL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtEventActionListL2 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtEventActionListL2MacList `json:"mac-list"`
}


type SysUtEventActionListL2MacList struct {
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


type SysUtEventActionListL3 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtEventActionListL3IpList `json:"ip-list"`
}


type SysUtEventActionListL3IpList struct {
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


type SysUtEventActionListTcp struct {
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
    Flags SysUtEventActionListTcpFlags `json:"flags"`
    Options SysUtEventActionListTcpOptions `json:"options"`
}


type SysUtEventActionListTcpFlags struct {
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


type SysUtEventActionListTcpOptions struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionListUdp struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}


type SysUtEventActionListIgnoreValidation struct {
    L1 int `json:"l1"`
    L2 int `json:"l2"`
    L3 int `json:"l3"`
    L4 int `json:"l4"`
    All int `json:"all"`
    Uuid string `json:"uuid"`
}

func (p *SysUtEvent) GetId() string{
    return strconv.Itoa(p.Inst.EventNumber)
}

func (p *SysUtEvent) getPath() string{
    return "sys-ut/event"
}

func (p *SysUtEvent) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEvent::Post")
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

func (p *SysUtEvent) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEvent::Get")
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
func (p *SysUtEvent) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEvent::Put")
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

func (p *SysUtEvent) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtEvent::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
