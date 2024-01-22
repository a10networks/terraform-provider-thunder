

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIsisRedistribute struct {
	Inst struct {

    Isis RouterIsisRedistributeIsis `json:"isis"`
    RedistList []RouterIsisRedistributeRedistList `json:"redist-list"`
    Uuid string `json:"uuid"`
    VipList []RouterIsisRedistributeVipList `json:"vip-list"`
    Tag string 

	} `json:"redistribute"`
}


type RouterIsisRedistributeIsis struct {
    Level1From RouterIsisRedistributeIsisLevel1From `json:"level-1-from"`
    Level2From RouterIsisRedistributeIsisLevel2From `json:"level-2-from"`
}


type RouterIsisRedistributeIsisLevel1From struct {
    Into1 RouterIsisRedistributeIsisLevel1FromInto1 `json:"into-1"`
}


type RouterIsisRedistributeIsisLevel1FromInto1 struct {
    Level2 int `json:"level-2"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisRedistributeIsisLevel2From struct {
    Into2 RouterIsisRedistributeIsisLevel2FromInto2 `json:"into-2"`
}


type RouterIsisRedistributeIsisLevel2FromInto2 struct {
    Level1 int `json:"level-1"`
    DistributeList string `json:"distribute-list"`
}


type RouterIsisRedistributeRedistList struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    MetricType string `json:"metric-type" dval:"internal"`
    RouteMap string `json:"route-map"`
    Level string `json:"level" dval:"level-2"`
}


type RouterIsisRedistributeVipList struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
    VipMetricType string `json:"vip-metric-type" dval:"internal"`
    VipLevel string `json:"vip-level" dval:"level-2"`
}

func (p *RouterIsisRedistribute) GetId() string{
    return "1"
}

func (p *RouterIsisRedistribute) getPath() string{
    return "router/isis/" +p.Inst.Tag + "/redistribute"
}

func (p *RouterIsisRedistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisRedistribute::Post")
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

func (p *RouterIsisRedistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisRedistribute::Get")
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
func (p *RouterIsisRedistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisRedistribute::Put")
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

func (p *RouterIsisRedistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIsisRedistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
