

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifIpv6 struct {
	Inst struct {

    AddressList []InterfaceLifIpv6AddressList `json:"address-list"`
    Inside int `json:"inside"`
    Ipv6Enable int `json:"ipv6-enable"`
    Ospf InterfaceLifIpv6Ospf578 `json:"ospf"`
    Outside int `json:"outside"`
    Router InterfaceLifIpv6Router588 `json:"router"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"ipv6"`
}


type InterfaceLifIpv6AddressList struct {
    Ipv6Addr string `json:"ipv6-addr"`
    Anycast int `json:"anycast"`
    LinkLocal int `json:"link-local"`
}


type InterfaceLifIpv6Ospf578 struct {
    NetworkList []InterfaceLifIpv6OspfNetworkList579 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceLifIpv6OspfCostCfg580 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceLifIpv6OspfDeadIntervalCfg581 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceLifIpv6OspfHelloIntervalCfg582 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceLifIpv6OspfMtuIgnoreCfg583 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceLifIpv6OspfNeighborCfg584 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceLifIpv6OspfPriorityCfg585 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg586 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceLifIpv6OspfTransmitDelayCfg587 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6OspfNetworkList579 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceLifIpv6OspfCostCfg580 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfDeadIntervalCfg581 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfHelloIntervalCfg582 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfMtuIgnoreCfg583 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfNeighborCfg584 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceLifIpv6OspfPriorityCfg585 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfRetransmitIntervalCfg586 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfTransmitDelayCfg587 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6Router588 struct {
    Ripng InterfaceLifIpv6RouterRipng589 `json:"ripng"`
    Ospf InterfaceLifIpv6RouterOspf590 `json:"ospf"`
    Isis InterfaceLifIpv6RouterIsis592 `json:"isis"`
}


type InterfaceLifIpv6RouterRipng589 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6RouterOspf590 struct {
    AreaList []InterfaceLifIpv6RouterOspfAreaList591 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceLifIpv6RouterOspfAreaList591 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6RouterIsis592 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceLifIpv6) GetId() string{
    return "1"
}

func (p *InterfaceLifIpv6) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/ipv6"
}

func (p *InterfaceLifIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6::Post")
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

func (p *InterfaceLifIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6::Get")
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
func (p *InterfaceLifIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6::Put")
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

func (p *InterfaceLifIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
