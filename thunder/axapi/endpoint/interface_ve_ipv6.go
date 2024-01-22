

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeIpv6 struct {
	Inst struct {

    AddressList []InterfaceVeIpv6AddressList `json:"address-list"`
    Inbound int `json:"inbound"`
    Inside int `json:"inside"`
    Ipv6Enable int `json:"ipv6-enable"`
    Ospf InterfaceVeIpv6Ospf944 `json:"ospf"`
    Outside int `json:"outside"`
    Rip InterfaceVeIpv6Rip954 `json:"rip"`
    Router InterfaceVeIpv6Router956 `json:"router"`
    RouterAdver InterfaceVeIpv6RouterAdver `json:"router-adver"`
    StatefulFirewall InterfaceVeIpv6StatefulFirewall961 `json:"stateful-firewall"`
    TtlIgnore int `json:"ttl-ignore"`
    Uuid string `json:"uuid"`
    V6AclName string `json:"v6-acl-name"`
    Ifnum string 

	} `json:"ipv6"`
}


type InterfaceVeIpv6AddressList struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
}


type InterfaceVeIpv6Ospf944 struct {
    NetworkList []InterfaceVeIpv6OspfNetworkList945 `json:"network-list"`
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
    CostCfg []InterfaceVeIpv6OspfCostCfg946 `json:"cost-cfg"`
    DeadIntervalCfg []InterfaceVeIpv6OspfDeadIntervalCfg947 `json:"dead-interval-cfg"`
    HelloIntervalCfg []InterfaceVeIpv6OspfHelloIntervalCfg948 `json:"hello-interval-cfg"`
    MtuIgnoreCfg []InterfaceVeIpv6OspfMtuIgnoreCfg949 `json:"mtu-ignore-cfg"`
    NeighborCfg []InterfaceVeIpv6OspfNeighborCfg950 `json:"neighbor-cfg"`
    PriorityCfg []InterfaceVeIpv6OspfPriorityCfg951 `json:"priority-cfg"`
    RetransmitIntervalCfg []InterfaceVeIpv6OspfRetransmitIntervalCfg952 `json:"retransmit-interval-cfg"`
    TransmitDelayCfg []InterfaceVeIpv6OspfTransmitDelayCfg953 `json:"transmit-delay-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6OspfNetworkList945 struct {
    BroadcastType string `json:"broadcast-type"`
    P2mpNbma int `json:"p2mp-nbma"`
    NetworkInstanceId int `json:"network-instance-id"`
}


type InterfaceVeIpv6OspfCostCfg946 struct {
    Cost int `json:"cost"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfDeadIntervalCfg947 struct {
    DeadInterval int `json:"dead-interval" dval:"40"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfHelloIntervalCfg948 struct {
    HelloInterval int `json:"hello-interval" dval:"10"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfMtuIgnoreCfg949 struct {
    MtuIgnore int `json:"mtu-ignore"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfNeighborCfg950 struct {
    Neighbor string `json:"neighbor" dval:"::"`
    NeigInst int `json:"neig-inst"`
    NeighborCost int `json:"neighbor-cost"`
    NeighborPollInterval int `json:"neighbor-poll-interval"`
    NeighborPriority int `json:"neighbor-priority"`
}


type InterfaceVeIpv6OspfPriorityCfg951 struct {
    Priority int `json:"priority" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfRetransmitIntervalCfg952 struct {
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6OspfTransmitDelayCfg953 struct {
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6Rip954 struct {
    SplitHorizonCfg InterfaceVeIpv6RipSplitHorizonCfg955 `json:"split-horizon-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RipSplitHorizonCfg955 struct {
    State string `json:"state" dval:"poisoned"`
}


type InterfaceVeIpv6Router956 struct {
    Ripng InterfaceVeIpv6RouterRipng957 `json:"ripng"`
    Ospf InterfaceVeIpv6RouterOspf958 `json:"ospf"`
    Isis InterfaceVeIpv6RouterIsis960 `json:"isis"`
}


type InterfaceVeIpv6RouterRipng957 struct {
    Rip int `json:"rip"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RouterOspf958 struct {
    AreaList []InterfaceVeIpv6RouterOspfAreaList959 `json:"area-list"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RouterOspfAreaList959 struct {
    AreaIdNum int `json:"area-id-num"`
    AreaIdAddr string `json:"area-id-addr"`
    Tag string `json:"tag"`
    InstanceId int `json:"instance-id"`
}


type InterfaceVeIpv6RouterIsis960 struct {
    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
}


type InterfaceVeIpv6RouterAdver struct {
    Action string `json:"action" dval:"disable"`
    DefaultLifetime int `json:"default-lifetime" dval:"1800"`
    HopLimit int `json:"hop-limit" dval:"255"`
    MaxInterval int `json:"max-interval" dval:"600"`
    MinInterval int `json:"min-interval" dval:"200"`
    RateLimit int `json:"rate-limit" dval:"100000"`
    ReachableTime int `json:"reachable-time"`
    RetransmitTimer int `json:"retransmit-timer"`
    AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
    AdverMtu int `json:"adver-mtu"`
    PrefixList []InterfaceVeIpv6RouterAdverPrefixList `json:"prefix-list"`
    ManagedConfigAction string `json:"managed-config-action" dval:"disable"`
    OtherConfigAction string `json:"other-config-action" dval:"disable"`
    AdverVrid int `json:"adver-vrid"`
    UseFloatingIp int `json:"use-floating-ip"`
    FloatingIp string `json:"floating-ip"`
    AdverVridDefault int `json:"adver-vrid-default"`
    UseFloatingIpDefaultVrid int `json:"use-floating-ip-default-vrid"`
    FloatingIpDefaultVrid string `json:"floating-ip-default-vrid"`
}


type InterfaceVeIpv6RouterAdverPrefixList struct {
    Prefix string `json:"prefix"`
    NotAutonomous int `json:"not-autonomous"`
    NotOnLink int `json:"not-on-link"`
    PreferredLifetime int `json:"preferred-lifetime" dval:"604800"`
    ValidLifetime int `json:"valid-lifetime" dval:"2592000"`
}


type InterfaceVeIpv6StatefulFirewall961 struct {
    Inside int `json:"inside"`
    ClassList string `json:"class-list"`
    Outside int `json:"outside"`
    AccessList int `json:"access-list"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceVeIpv6) GetId() string{
    return "1"
}

func (p *InterfaceVeIpv6) getPath() string{
    return "interface/ve/" +p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceVeIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpv6::Post")
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

func (p *InterfaceVeIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpv6::Get")
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
func (p *InterfaceVeIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpv6::Put")
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

func (p *InterfaceVeIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceVeIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
