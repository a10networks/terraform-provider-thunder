

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6RipRedistribute struct {
	Inst struct {

    RedistList []RouterIpv6RipRedistributeRedistList `json:"redist-list"`
    Uuid string `json:"uuid"`
    VipList []RouterIpv6RipRedistributeVipList `json:"vip-list"`

	} `json:"redistribute"`
}


type RouterIpv6RipRedistributeRedistList struct {
    Type string `json:"type"`
    Metric int `json:"metric"`
    RouteMap string `json:"route-map"`
}


type RouterIpv6RipRedistributeVipList struct {
    VipType string `json:"vip-type"`
    VipMetric int `json:"vip-metric"`
    VipRouteMap string `json:"vip-route-map"`
}

func (p *RouterIpv6RipRedistribute) GetId() string{
    return "1"
}

func (p *RouterIpv6RipRedistribute) getPath() string{
    return "router/ipv6/rip/redistribute"
}

func (p *RouterIpv6RipRedistribute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRedistribute::Post")
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

func (p *RouterIpv6RipRedistribute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRedistribute::Get")
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
func (p *RouterIpv6RipRedistribute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRedistribute::Put")
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

func (p *RouterIpv6RipRedistribute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRedistribute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
