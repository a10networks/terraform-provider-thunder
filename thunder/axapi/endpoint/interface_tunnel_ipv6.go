

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnelIpv6 struct {
	Inst struct {

    AddressCfg []InterfaceTunnelIpv6AddressCfg `json:"address-cfg"`
    Inside int `json:"inside"`
    Ipv6Enable int `json:"ipv6-enable"`
    Ospf InterfaceTunnelIpv6Ospf872 `json:"ospf"`
    Outside int `json:"outside"`
    Router InterfaceTunnelIpv6Router882 `json:"router"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ipv6"`
}


type InterfaceTunnelIpv6AddressCfg struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceTunnelIpv6Ospf872 struct {
    NetworkList []InterfaceTunnelIpv6OspfNetworkList873 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceTunnelIpv6OspfCostCfg874 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceTunnelIpv6OspfDeadIntervalCfg875 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceTunnelIpv6OspfHelloIntervalCfg876 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceTunnelIpv6OspfMtuIgnoreCfg877 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceTunnelIpv6OspfNeighborCfg878 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceTunnelIpv6OspfPriorityCfg879 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceTunnelIpv6OspfRetransmitIntervalCfg880 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceTunnelIpv6OspfTransmitDelayCfg881 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpv6OspfNetworkList873 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceTunnelIpv6OspfCostCfg874 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfDeadIntervalCfg875 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfHelloIntervalCfg876 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfMtuIgnoreCfg877 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfNeighborCfg878 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceTunnelIpv6OspfPriorityCfg879 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfRetransmitIntervalCfg880 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6OspfTransmitDelayCfg881 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceTunnelIpv6Router882 struct {
    Ripng InterfaceTunnelIpv6RouterRipng883 `json:"ripng"`
    Ospf InterfaceTunnelIpv6RouterOspf884 `json:"ospf"`
}


type InterfaceTunnelIpv6RouterRipng883 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpv6RouterOspf884 struct {
    AreaList []InterfaceTunnelIpv6RouterOspfAreaList885 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceTunnelIpv6RouterOspfAreaList885 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}

func (p *InterfaceTunnelIpv6) GetId() string{
    return "1"
}

func (p *InterfaceTunnelIpv6) getPath() string{
    return "interface/tunnel/" +p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceTunnelIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpv6::Post")
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

func (p *InterfaceTunnelIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpv6::Get")
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
func (p *InterfaceTunnelIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpv6::Put")
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

func (p *InterfaceTunnelIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
