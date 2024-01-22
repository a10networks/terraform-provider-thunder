

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifIpv6Ospf struct {
	Inst struct {

    Bfd int `json:"bfd"`
    CostCfg []InterfaceLifIpv6OspfCostCfg `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceLifIpv6OspfDeadIntervalCfg `json:"dead-interval-cfg"`
    Disable int `json:"disable"`
    HelloIntervalCfg []InterfaceLifIpv6OspfHelloIntervalCfg `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceLifIpv6OspfMtuIgnoreCfg `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceLifIpv6OspfNeighborCfg `json:"neighbor-cfg"`
    NetworkList []InterfaceLifIpv6OspfNetworkList `json:"network-list"`
    PriorityCfg []InterfaceLifIpv6OspfPriorityCfg `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceLifIpv6OspfTransmitDelayCfg `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"ospf"`
}


type InterfaceLifIpv6OspfCostCfg struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfDeadIntervalCfg struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfHelloIntervalCfg struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfMtuIgnoreCfg struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfNeighborCfg struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceLifIpv6OspfNetworkList struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceLifIpv6OspfPriorityCfg struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfRetransmitIntervalCfg struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceLifIpv6OspfTransmitDelayCfg struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}

func (p *InterfaceLifIpv6Ospf) GetId() string{
    return "1"
}

func (p *InterfaceLifIpv6Ospf) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/ipv6/ospf"
}

func (p *InterfaceLifIpv6Ospf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6Ospf::Post")
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

func (p *InterfaceLifIpv6Ospf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6Ospf::Get")
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
func (p *InterfaceLifIpv6Ospf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6Ospf::Put")
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

func (p *InterfaceLifIpv6Ospf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6Ospf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
