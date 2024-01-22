

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterOspfRedistribute struct {
	Inst struct {

    IpNat int `json:"ip-nat"`
    IpNatFloatingList []RouterOspfRedistributeIpNatFloatingList `json:"ip-nat-floating-list"`
    MetricIpNat int `json:"metric-ip-nat"`
    MetricTypeIpNat string `json:"metric-type-ip-nat" dval:"2"`
    OspfList []RouterOspfRedistributeOspfList `json:"ospf-list"`
    RedistList []RouterOspfRedistributeRedistList `json:"redist-list"`
    RouteMapIpNat string `json:"route-map-ip-nat"`
    TagIpNat int `json:"tag-ip-nat"`
    Uuid string `json:"uuid"`
    VipFloatingList []RouterOspfRedistributeVipFloatingList `json:"vip-floating-list"`
    VipList []RouterOspfRedistributeVipList `json:"vip-list"`
    ProcessId string 

	} `json:"redistribute"`
}


type RouterOspfRedistributeIpNatFloatingList struct {
    IpNatPrefix string `json:"ip-nat-prefix"`
    IpNatFloatingIpForward string `json:"ip-nat-floating-IP-forward"`
}


type RouterOspfRedistributeOspfList struct {
    Ospf int `json:"ospf"`
    ProcessId int `json:"process-id"`
    MetricOspf int `json:"metric-ospf"`
    MetricTypeOspf string `json:"metric-type-ospf" dval:"2"`
    RouteMapOspf string `json:"route-map-ospf"`
    TagOspf int `json:"tag-ospf"`
}


type RouterOspfRedistributeRedistList struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"2"`
    RouteMap string `json:"route-map"`
    Tag int `json:"tag"`
}


type RouterOspfRedistributeVipFloatingList struct {
    VipAddress string `json:"vip-address"`
    VipFloatingIpForward string `json:"vip-floating-IP-forward"`
}


type RouterOspfRedistributeVipList struct {
    TypeVip string `json:"type-vip"`
    MetricVip int `json:"metric-vip"`
    MetricTypeVip string `json:"metric-type-vip" dval:"2"`
    RouteMapVip string `json:"route-map-vip"`
    TagVip int `json:"tag-vip"`
}

func (p *RouterOspfRedistribute) GetId() string{
    return "1"
}

func (p *RouterOspfRedistribute) getPath() string{
    return "router/ospf/" +p.Inst.ProcessId + "/redistribute"
}

func (p *RouterOspfRedistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspfRedistribute::Post")
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

func (p *RouterOspfRedistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspfRedistribute::Get")
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
func (p *RouterOspfRedistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspfRedistribute::Put")
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

func (p *RouterOspfRedistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterOspfRedistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
