

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetIpv6 struct {
	Inst struct {

    AccessListCfg InterfaceEthernetIpv6AccessListCfg `json:"access-list-cfg"`
    AddressList []InterfaceEthernetIpv6AddressList `json:"address-list"`
    Inside int `json:"inside"`
    Ipv6Enable int `json:"ipv6-enable"`
    Ospf InterfaceEthernetIpv6Ospf466 `json:"ospf"`
    Outside int `json:"outside"`
    Rip InterfaceEthernetIpv6Rip476 `json:"rip"`
    Router InterfaceEthernetIpv6Router478 `json:"router"`
    RouterAdver InterfaceEthernetIpv6RouterAdver `json:"router-adver"`
    StatefulFirewall InterfaceEthernetIpv6StatefulFirewall483 `json:"stateful-firewall"`
    TtlIgnore int `json:"ttl-ignore"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ipv6"`
}


type InterfaceEthernetIpv6AccessListCfg struct {
    V6AclName string `json:"v6-acl-name"`
    Inbound int `json:"inbound"`
}


type InterfaceEthernetIpv6AddressList struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceEthernetIpv6Ospf466 struct {
    NetworkList []InterfaceEthernetIpv6OspfNetworkList467 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceEthernetIpv6OspfCostCfg468 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceEthernetIpv6OspfDeadIntervalCfg469 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceEthernetIpv6OspfHelloIntervalCfg470 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceEthernetIpv6OspfMtuIgnoreCfg471 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceEthernetIpv6OspfNeighborCfg472 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceEthernetIpv6OspfPriorityCfg473 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceEthernetIpv6OspfRetransmitIntervalCfg474 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceEthernetIpv6OspfTransmitDelayCfg475 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6OspfNetworkList467 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceEthernetIpv6OspfCostCfg468 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfDeadIntervalCfg469 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfHelloIntervalCfg470 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfMtuIgnoreCfg471 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfNeighborCfg472 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceEthernetIpv6OspfPriorityCfg473 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfRetransmitIntervalCfg474 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6OspfTransmitDelayCfg475 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6Rip476 struct {
    SplitHorizonCfg InterfaceEthernetIpv6RipSplitHorizonCfg477 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RipSplitHorizonCfg477 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceEthernetIpv6Router478 struct {
    Ripng InterfaceEthernetIpv6RouterRipng479 `json:"ripng"`
    Ospf InterfaceEthernetIpv6RouterOspf480 `json:"ospf"`
    Isis InterfaceEthernetIpv6RouterIsis482 `json:"isis"`
}


type InterfaceEthernetIpv6RouterRipng479 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RouterOspf480 struct {
    AreaList []InterfaceEthernetIpv6RouterOspfAreaList481 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RouterOspfAreaList481 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceEthernetIpv6RouterIsis482 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceEthernetIpv6RouterAdver struct {
    Action string `json:"action" dval:"disable"`
    HopLimit int `json:"hop-limit" dval:"255"`
    MaxInterval int `json:"max-interval" dval:"600"`
    MinInterval int `json:"min-interval" dval:"200"`
    DefaultLifetime int `json:"default-lifetime" dval:"1800"`
    RateLimit int `json:"rate-limit" dval:"100000"`
    ReachableTime int `json:"reachable-time"`
    RetransmitTimer int `json:"retransmit-timer"`
    AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
    AdverMtu int `json:"adver-mtu"`
    PrefixList []InterfaceEthernetIpv6RouterAdverPrefixList `json:"prefix-list"`
    ManagedConfigAction string `json:"managed-config-action" dval:"disable"`
    OtherConfigAction string `json:"other-config-action" dval:"disable"`
    AdverVrid int `json:"adver-vrid"`
    UseFloatingIp int `json:"use-floating-ip"`
    FloatingIp string `json:"floating-ip"`
    AdverVridDefault int `json:"adver-vrid-default"`
    UseFloatingIpDefaultVrid int `json:"use-floating-ip-default-vrid"`
    FloatingIpDefaultVrid string `json:"floating-ip-default-vrid"`
}


type InterfaceEthernetIpv6RouterAdverPrefixList struct {
    Prefix string `json:"prefix"`
    NotAutonomous int `json:"not-autonomous"`
    NotOnLink int `json:"not-on-link"`
    PreferredLifetime int `json:"preferred-lifetime" dval:"604800"`
    ValidLifetime int `json:"valid-lifetime" dval:"2592000"`
}


type InterfaceEthernetIpv6StatefulFirewall483 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceEthernetIpv6) GetId() string{
    return "1"
}

func (p *InterfaceEthernetIpv6) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceEthernetIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6::Post")
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

func (p *InterfaceEthernetIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6::Get")
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
func (p *InterfaceEthernetIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6::Put")
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

func (p *InterfaceEthernetIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
