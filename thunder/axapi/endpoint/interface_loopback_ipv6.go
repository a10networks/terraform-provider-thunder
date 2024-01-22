

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackIpv6 struct {
	Inst struct {

    AddressList []InterfaceLoopbackIpv6AddressList `json:"address-list"`
    Ipv6Enable int `json:"ipv6-enable"`
    Ospf InterfaceLoopbackIpv6Ospf669 `json:"ospf"`
    Rip InterfaceLoopbackIpv6Rip677 `json:"rip"`
    Router InterfaceLoopbackIpv6Router679 `json:"router"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ipv6"`
}


type InterfaceLoopbackIpv6AddressList struct {
    Ipv6Addr string `json:"ipv6-addr"`
    Anycast int `json:"anycast"`
    LinkLocal int `json:"link-local"`
}


type InterfaceLoopbackIpv6Ospf669 struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceLoopbackIpv6OspfCostCfg670 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceLoopbackIpv6OspfDeadIntervalCfg671 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceLoopbackIpv6OspfHelloIntervalCfg672 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceLoopbackIpv6OspfMtuIgnoreCfg673 `json:"mtu-ignore-cfg"`
    PriorityCfg []InterfaceLoopbackIpv6OspfPriorityCfg674 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceLoopbackIpv6OspfRetransmitIntervalCfg675 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceLoopbackIpv6OspfTransmitDelayCfg676 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6OspfCostCfg670 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfDeadIntervalCfg671 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfHelloIntervalCfg672 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfMtuIgnoreCfg673 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfPriorityCfg674 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfRetransmitIntervalCfg675 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6OspfTransmitDelayCfg676 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6Rip677 struct {
    SplitHorizonCfg InterfaceLoopbackIpv6RipSplitHorizonCfg678 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6RipSplitHorizonCfg678 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceLoopbackIpv6Router679 struct {
    Ripng InterfaceLoopbackIpv6RouterRipng680 `json:"ripng"`
    Ospf InterfaceLoopbackIpv6RouterOspf681 `json:"ospf"`
    Isis InterfaceLoopbackIpv6RouterIsis683 `json:"isis"`
}


type InterfaceLoopbackIpv6RouterRipng680 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6RouterOspf681 struct {
    AreaList []InterfaceLoopbackIpv6RouterOspfAreaList682 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceLoopbackIpv6RouterOspfAreaList682 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLoopbackIpv6RouterIsis683 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceLoopbackIpv6) GetId() string{
    return "1"
}

func (p *InterfaceLoopbackIpv6) getPath() string{
    return "interface/loopback/" +p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceLoopbackIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpv6::Post")
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

func (p *InterfaceLoopbackIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpv6::Get")
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
func (p *InterfaceLoopbackIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpv6::Put")
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

func (p *InterfaceLoopbackIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
