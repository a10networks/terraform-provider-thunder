

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6OspfRedistribute struct {
	Inst struct {

    IpNat int `json:"ip-nat"`
    IpNatFloatingList []RouterIpv6OspfRedistributeIpNatFloatingList `json:"ip-nat-floating-list"`
    MetricIpNat int `json:"metric-ip-nat"`
    MetricTypeIpNat string `json:"metric-type-ip-nat" dval:"2"`
    OspfList []RouterIpv6OspfRedistributeOspfList `json:"ospf-list"`
    RedistList []RouterIpv6OspfRedistributeRedistList `json:"redist-list"`
    RouteMapIpNat string `json:"route-map-ip-nat"`
    Uuid string `json:"uuid"`
    VipFloatingList []RouterIpv6OspfRedistributeVipFloatingList `json:"vip-floating-list"`
    VipList []RouterIpv6OspfRedistributeVipList `json:"vip-list"`
    ProcessId string 

	} `json:"redistribute"`
}


type RouterIpv6OspfRedistributeIpNatFloatingList struct {
    IpNatPrefix string `json:"ip-nat-prefix"`
    IpNatFloatingIpForward string `json:"ip-nat-floating-IP-forward"`
}


type RouterIpv6OspfRedistributeOspfList struct {
    Ospf int `json:"ospf"`
    ProcessId string `json:"process-id"`
    MetricOspf int `json:"metric-ospf"`
    MetricTypeOspf string `json:"metric-type-ospf" dval:"2"`
    RouteMapOspf string `json:"route-map-ospf"`
}


type RouterIpv6OspfRedistributeRedistList struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"2"`
    RouteMap string `json:"route-map"`
}


type RouterIpv6OspfRedistributeVipFloatingList struct {
    VipAddress string `json:"vip-address"`
    VipFloatingIpForward string `json:"vip-floating-IP-forward"`
}


type RouterIpv6OspfRedistributeVipList struct {
    TypeVip string `json:"type-vip"`
    MetricVip int `json:"metric-vip"`
    MetricTypeVip string `json:"metric-type-vip" dval:"2"`
    RouteMapVip string `json:"route-map-vip"`
}

func (p *RouterIpv6OspfRedistribute) GetId() string{
    return "1"
}

func (p *RouterIpv6OspfRedistribute) getPath() string{
    return "router/ipv6/ospf/" +p.Inst.ProcessId + "/redistribute"
}

func (p *RouterIpv6OspfRedistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfRedistribute::Post")
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

func (p *RouterIpv6OspfRedistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfRedistribute::Get")
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
func (p *RouterIpv6OspfRedistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfRedistribute::Put")
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

func (p *RouterIpv6OspfRedistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6OspfRedistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
