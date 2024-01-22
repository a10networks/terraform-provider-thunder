

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIsisAddressFamilyIpv6Redistribute struct {
	Inst struct {

    Isis RouterIsisAddressFamilyIpv6RedistributeIsis `json:"isis"`
    RedistList []RouterIsisAddressFamilyIpv6RedistributeRedistList `json:"redist-list"`
    Uuid string `json:"uuid"`
    VipList []RouterIsisAddressFamilyIpv6RedistributeVipList `json:"vip-list"`
    Tag string 

	} `json:"redistribute"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsis struct {
    Level1From RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From `json:"level-1-from"`
    Level2From RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From `json:"level-2-from"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From struct {
    Into1 RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1 `json:"into-1"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto1 struct {
    Level2 int `json:"level-2"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From struct {
    Into2 RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2 `json:"into-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto2 struct {
    Level1 int `json:"level-1"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisAddressFamilyIpv6RedistributeRedistList struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"internal"`
    RouteMap string `json:"route-map"`
    Level string `json:"level" dval:"level-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeVipList struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
    VipMetricType string `json:"vip-metric-type" dval:"internal"`
    VipLevel string `json:"vip-level" dval:"level-2"`
}

func (p *RouterIsisAddressFamilyIpv6Redistribute) GetId() string{
    return "1"
}

func (p *RouterIsisAddressFamilyIpv6Redistribute) getPath() string{
    return "router/isis/" +p.Inst.Tag + "/address-family/ipv6/redistribute"
}

func (p *RouterIsisAddressFamilyIpv6Redistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6Redistribute::Post")
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

func (p *RouterIsisAddressFamilyIpv6Redistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6Redistribute::Get")
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
func (p *RouterIsisAddressFamilyIpv6Redistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6Redistribute::Put")
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

func (p *RouterIsisAddressFamilyIpv6Redistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6Redistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
