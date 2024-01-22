

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv6Redistribute struct {
	Inst struct {

    ConnectedCfg RouterBgpAddressFamilyIpv6RedistributeConnectedCfg `json:"connected-cfg"`
    FloatingIpCfg RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg `json:"floating-ip-cfg"`
    IpNatCfg RouterBgpAddressFamilyIpv6RedistributeIpNatCfg `json:"ip-nat-cfg"`
    IpNatListCfg RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg `json:"ip-nat-list-cfg"`
    IsisCfg RouterBgpAddressFamilyIpv6RedistributeIsisCfg `json:"isis-cfg"`
    Lw4o6Cfg RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg `json:"lw4o6-cfg"`
    NatMapCfg RouterBgpAddressFamilyIpv6RedistributeNatMapCfg `json:"nat-map-cfg"`
    Nat64Cfg RouterBgpAddressFamilyIpv6RedistributeNat64Cfg `json:"nat64-cfg"`
    OspfCfg RouterBgpAddressFamilyIpv6RedistributeOspfCfg `json:"ospf-cfg"`
    RipCfg RouterBgpAddressFamilyIpv6RedistributeRipCfg `json:"rip-cfg"`
    StaticCfg RouterBgpAddressFamilyIpv6RedistributeStaticCfg `json:"static-cfg"`
    StaticNatCfg RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg `json:"static-nat-cfg"`
    Uuid string `json:"uuid"`
    Vip RouterBgpAddressFamilyIpv6RedistributeVip `json:"vip"`
    AsNumber string 

	} `json:"redistribute"`
}


type RouterBgpAddressFamilyIpv6RedistributeConnectedCfg struct {
    Connected int `json:"connected"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg struct {
    FloatingIp int `json:"floating-ip"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeIpNatCfg struct {
    IpNat int `json:"ip-nat"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg struct {
    IpNatList int `json:"ip-nat-list"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeIsisCfg struct {
    Isis int `json:"isis"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg struct {
    Lw4o6 int `json:"lw4o6"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeNatMapCfg struct {
    NatMap int `json:"nat-map"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeNat64Cfg struct {
    Nat64 int `json:"nat64"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeOspfCfg struct {
    Ospf int `json:"ospf"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeRipCfg struct {
    Rip int `json:"rip"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeStaticCfg struct {
    Static int `json:"static"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg struct {
    StaticNat int `json:"static-nat"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeVip struct {
    OnlyFlaggedCfg RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg `json:"only-flagged-cfg"`
    OnlyNotFlaggedCfg RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg `json:"only-not-flagged-cfg"`
}


type RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg struct {
    OnlyFlagged int `json:"only-flagged"`
    RouteMap string `json:"route-map"`
}


type RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg struct {
    OnlyNotFlagged int `json:"only-not-flagged"`
    RouteMap string `json:"route-map"`
}

func (p *RouterBgpAddressFamilyIpv6Redistribute) GetId() string{
    return "1"
}

func (p *RouterBgpAddressFamilyIpv6Redistribute) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv6/redistribute"
}

func (p *RouterBgpAddressFamilyIpv6Redistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6Redistribute::Post")
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

func (p *RouterBgpAddressFamilyIpv6Redistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6Redistribute::Get")
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
func (p *RouterBgpAddressFamilyIpv6Redistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6Redistribute::Put")
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

func (p *RouterBgpAddressFamilyIpv6Redistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6Redistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
