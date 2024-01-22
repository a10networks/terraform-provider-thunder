

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpRedistribute struct {
	Inst struct {

    ConnectedCfg RouterBgpRedistributeConnectedCfg `json:"connected-cfg"`
    FloatingIpCfg RouterBgpRedistributeFloatingIpCfg `json:"floating-ip-cfg"`
    IpNatCfg RouterBgpRedistributeIpNatCfg `json:"ip-nat-cfg"`
    IpNatListCfg RouterBgpRedistributeIpNatListCfg `json:"ip-nat-list-cfg"`
    IsisCfg RouterBgpRedistributeIsisCfg `json:"isis-cfg"`
    Lw4o6Cfg RouterBgpRedistributeLw4o6Cfg `json:"lw4o6-cfg"`
    NatMapCfg RouterBgpRedistributeNatMapCfg `json:"nat-map-cfg"`
    OspfCfg RouterBgpRedistributeOspfCfg `json:"ospf-cfg"`
    RipCfg RouterBgpRedistributeRipCfg `json:"rip-cfg"`
    StaticCfg RouterBgpRedistributeStaticCfg `json:"static-cfg"`
    StaticNatCfg RouterBgpRedistributeStaticNatCfg `json:"static-nat-cfg"`
    Uuid string `json:"uuid"`
    Vip RouterBgpRedistributeVip `json:"vip"`
    AsNumber string 

	} `json:"redistribute"`
}


type RouterBgpRedistributeConnectedCfg struct {
    Connected int `json:"connected"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeFloatingIpCfg struct {
    FloatingIp int `json:"floating-ip"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeIpNatCfg struct {
    IpNat int `json:"ip-nat"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeIpNatListCfg struct {
    IpNatList int `json:"ip-nat-list"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeIsisCfg struct {
    Isis int `json:"isis"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeLw4o6Cfg struct {
    Lw4o6 int `json:"lw4o6"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeNatMapCfg struct {
    NatMap int `json:"nat-map"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeOspfCfg struct {
    Ospf int `json:"ospf"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeRipCfg struct {
    Rip int `json:"rip"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeStaticCfg struct {
    Static int `json:"static"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeStaticNatCfg struct {
    StaticNat int `json:"static-nat"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeVip struct {
    OnlyFlaggedCfg RouterBgpRedistributeVipOnlyFlaggedCfg `json:"only-flagged-cfg"`
    OnlyNotFlaggedCfg RouterBgpRedistributeVipOnlyNotFlaggedCfg `json:"only-not-flagged-cfg"`
}


type RouterBgpRedistributeVipOnlyFlaggedCfg struct {
    OnlyFlagged int `json:"only-flagged"`
    RouteMap string `json:"route-map"`
}


type RouterBgpRedistributeVipOnlyNotFlaggedCfg struct {
    OnlyNotFlagged int `json:"only-not-flagged"`
    RouteMap string `json:"route-map"`
}

func (p *RouterBgpRedistribute) GetId() string{
    return "1"
}

func (p *RouterBgpRedistribute) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/redistribute"
}

func (p *RouterBgpRedistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpRedistribute::Post")
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

func (p *RouterBgpRedistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpRedistribute::Get")
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
func (p *RouterBgpRedistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpRedistribute::Put")
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

func (p *RouterBgpRedistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpRedistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
