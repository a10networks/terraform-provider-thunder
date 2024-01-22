

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterIpv6RipRouteMap struct {
	Inst struct {

    MapCfg []RouterIpv6RipRouteMapMapCfg `json:"map-cfg"`
    Uuid string `json:"uuid"`

	} `json:"route-map"`
}


type RouterIpv6RipRouteMapMapCfg struct {
    Map string `json:"map"`
    RouteMapDirection string `json:"route-map-direction"`
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Ve int `json:"ve"`
}

func (p *RouterIpv6RipRouteMap) GetId() string{
    return "1"
}

func (p *RouterIpv6RipRouteMap) getPath() string{
    return "router/ipv6/rip/route-map"
}

func (p *RouterIpv6RipRouteMap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRouteMap::Post")
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

func (p *RouterIpv6RipRouteMap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRouteMap::Get")
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
func (p *RouterIpv6RipRouteMap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRouteMap::Put")
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

func (p *RouterIpv6RipRouteMap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterIpv6RipRouteMap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
