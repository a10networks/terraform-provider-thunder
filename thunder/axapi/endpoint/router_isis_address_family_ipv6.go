

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIsisAddressFamilyIpv6 struct {
	Inst struct {

    AdjacencyCheck int `json:"adjacency-check" dval:"1"`
    DefaultInformation string `json:"default-information"`
    Distance int `json:"distance" dval:"115"`
    MultiTopologyCfg RouterIsisAddressFamilyIpv6MultiTopologyCfg `json:"multi-topology-cfg"`
    Redistribute RouterIsisAddressFamilyIpv6Redistribute1268 `json:"redistribute"`
    SummaryPrefixList []RouterIsisAddressFamilyIpv6SummaryPrefixList `json:"summary-prefix-list"`
    Uuid string `json:"uuid"`
    Tag string 

	} `json:"ipv6"`
}


type RouterIsisAddressFamilyIpv6MultiTopologyCfg struct {
    MultiTopology int `json:"multi-topology"`
    Level string `json:"level"`
    Transition int `json:"transition"`
    LevelTransition int `json:"level-transition"`
}


type RouterIsisAddressFamilyIpv6Redistribute1268 struct {
    RedistList []RouterIsisAddressFamilyIpv6RedistributeRedistList1269 `json:"redist-list"`
    VipList []RouterIsisAddressFamilyIpv6RedistributeVipList1270 `json:"vip-list"`
    Isis RouterIsisAddressFamilyIpv6RedistributeIsis1271 `json:"isis"`
    Uuid string `json:"uuid"`
}


type RouterIsisAddressFamilyIpv6RedistributeRedistList1269 struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"internal"`
    RouteMap string `json:"route-map"`
    Level string `json:"level" dval:"level-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeVipList1270 struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
    VipMetricType string `json:"vip-metric-type" dval:"internal"`
    VipLevel string `json:"vip-level" dval:"level-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsis1271 struct {
    Level1From RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1272 `json:"level-1-from"`
    Level2From RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1274 `json:"level-2-from"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel1From1272 struct {
    Into1 RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11273 `json:"into-1"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel1FromInto11273 struct {
    Level2 int `json:"level-2"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel2From1274 struct {
    Into2 RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21275 `json:"into-2"`
}


type RouterIsisAddressFamilyIpv6RedistributeIsisLevel2FromInto21275 struct {
    Level1 int `json:"level-1"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisAddressFamilyIpv6SummaryPrefixList struct {
    Prefix string `json:"prefix"`
    Level string `json:"level" dval:"level-2"`
}

func (p *RouterIsisAddressFamilyIpv6) GetId() string{
    return "1"
}

func (p *RouterIsisAddressFamilyIpv6) getPath() string{
    return "router/isis/" +p.Inst.Tag + "/address-family/ipv6"
}

func (p *RouterIsisAddressFamilyIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6::Post")
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

func (p *RouterIsisAddressFamilyIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6::Get")
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
func (p *RouterIsisAddressFamilyIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6::Put")
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

func (p *RouterIsisAddressFamilyIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisAddressFamilyIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
