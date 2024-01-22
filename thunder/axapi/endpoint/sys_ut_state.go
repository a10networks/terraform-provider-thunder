

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SysUtState struct {
	Inst struct {

    Name string `json:"name"`
    NextStateList []SysUtStateNextStateList `json:"next-state-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"state"`
}


type SysUtStateNextStateList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    CaseList []SysUtStateNextStateListCaseList `json:"case-list"`
}


type SysUtStateNextStateListCaseList struct {
    CaseNumber int `json:"case-number"`
    Repeat int `json:"repeat"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    ActionList []SysUtStateNextStateListCaseListActionList `json:"action-list"`
}


type SysUtStateNextStateListCaseListActionList struct {
    Direction string `json:"direction"`
    Template string `json:"template"`
    Drop int `json:"drop"`
    Delay int `json:"delay"`
    Uuid string `json:"uuid"`
    L1 SysUtStateNextStateListCaseListActionListL1 `json:"l1"`
    L2 SysUtStateNextStateListCaseListActionListL2 `json:"l2"`
    L3 SysUtStateNextStateListCaseListActionListL3 `json:"l3"`
    Tcp SysUtStateNextStateListCaseListActionListTcp `json:"tcp"`
    Udp SysUtStateNextStateListCaseListActionListUdp `json:"udp"`
}


type SysUtStateNextStateListCaseListActionListL1 struct {
    EthList []SysUtStateNextStateListCaseListActionListL1EthList `json:"eth-list"`
    Trunk_list []SysUtStateNextStateListCaseListActionListL1Trunk_list `json:"trunk_list"`
    Length int `json:"length"`
    Value int `json:"value"`
    Auto int `json:"auto"`
    Uuid string `json:"uuid"`
}


type SysUtStateNextStateListCaseListActionListL1EthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type SysUtStateNextStateListCaseListActionListL1Trunk_list struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}


type SysUtStateNextStateListCaseListActionListL2 struct {
    Ethertype int `json:"ethertype"`
    Protocol string `json:"protocol" dval:"ipv4"`
    Value int `json:"value"`
    Vlan int `json:"vlan"`
    Uuid string `json:"uuid"`
    MacList []SysUtStateNextStateListCaseListActionListL2MacList `json:"mac-list"`
}


type SysUtStateNextStateListCaseListActionListL2MacList struct {
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


type SysUtStateNextStateListCaseListActionListL3 struct {
    Protocol int `json:"protocol"`
    Type string `json:"type"`
    Value int `json:"value"`
    Checksum string `json:"checksum" dval:"valid"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
    IpList []SysUtStateNextStateListCaseListActionListL3IpList `json:"ip-list"`
}


type SysUtStateNextStateListCaseListActionListL3IpList struct {
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


type SysUtStateNextStateListCaseListActionListTcp struct {
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
    Flags SysUtStateNextStateListCaseListActionListTcpFlags `json:"flags"`
    Options SysUtStateNextStateListCaseListActionListTcpOptions `json:"options"`
}


type SysUtStateNextStateListCaseListActionListTcpFlags struct {
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


type SysUtStateNextStateListCaseListActionListTcpOptions struct {
    Mss int `json:"mss"`
    Wscale int `json:"wscale"`
    SackType string `json:"sack-type"`
    TimeStampEnable int `json:"time-stamp-enable"`
    Nop int `json:"nop"`
    Uuid string `json:"uuid"`
}


type SysUtStateNextStateListCaseListActionListUdp struct {
    SrcPort int `json:"src-port"`
    DestPort int `json:"dest-port"`
    DestPortValue int `json:"dest-port-value"`
    NatPool string `json:"nat-pool"`
    Length int `json:"length"`
    Checksum string `json:"checksum" dval:"valid"`
    Uuid string `json:"uuid"`
}

func (p *SysUtState) GetId() string{
    return p.Inst.Name
}

func (p *SysUtState) getPath() string{
    return "sys-ut/state"
}

func (p *SysUtState) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtState::Post")
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

func (p *SysUtState) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtState::Get")
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
func (p *SysUtState) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtState::Put")
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

func (p *SysUtState) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SysUtState::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
