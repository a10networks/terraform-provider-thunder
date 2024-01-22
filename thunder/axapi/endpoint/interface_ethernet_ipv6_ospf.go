

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetIpv6Ospf struct {
	Inst struct {

    Bfd int `json:"bfd"`
    CostCfg []InterfaceEthernetIpv6OspfCostCfg `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceEthernetIpv6OspfDeadIntervalCfg `json:"dead-interval-cfg"`
    Disable int `json:"disable"`
    HelloIntervalCfg []InterfaceEthernetIpv6OspfHelloIntervalCfg `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceEthernetIpv6OspfMtuIgnoreCfg `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceEthernetIpv6OspfNeighborCfg `json:"neighbor-cfg"`
    NetworkList []InterfaceEthernetIpv6OspfNetworkList `json:"network-list"`
    PriorityCfg []InterfaceEthernetIpv6OspfPriorityCfg `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceEthernetIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceEthernetIpv6OspfTransmitDelayCfg `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ospf"`
}


type InterfaceEthernetIpv6OspfCostCfg struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfDeadIntervalCfg struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfHelloIntervalCfg struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfMtuIgnoreCfg struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfNeighborCfg struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceEthernetIpv6OspfNetworkList struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceEthernetIpv6OspfPriorityCfg struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfRetransmitIntervalCfg struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfTransmitDelayCfg struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}

func (p *InterfaceEthernetIpv6Ospf) GetId() string{
    return "1"
}

func (p *InterfaceEthernetIpv6Ospf) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/ipv6/ospf"
}

func (p *InterfaceEthernetIpv6Ospf) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6Ospf::Post")
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

func (p *InterfaceEthernetIpv6Ospf) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6Ospf::Get")
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
func (p *InterfaceEthernetIpv6Ospf) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6Ospf::Put")
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

func (p *InterfaceEthernetIpv6Ospf) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6Ospf::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
