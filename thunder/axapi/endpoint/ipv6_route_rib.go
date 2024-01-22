

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RouteRib struct {
	Inst struct {

    Ipv6Address string `json:"ipv6-address"`
    Ipv6NexthopIpv6 []Ipv6RouteRibIpv6NexthopIpv6 `json:"ipv6-nexthop-ipv6"`
    Ipv6NexthopTunnel []Ipv6RouteRibIpv6NexthopTunnel `json:"ipv6-nexthop-tunnel"`
    Uuid string `json:"uuid"`

	} `json:"rib"`
}


type Ipv6RouteRibIpv6NexthopIpv6 struct {
    Ipv6Nexthop string `json:"ipv6-nexthop"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Distance int `json:"distance" dval:"1"`
    Description string `json:"description"`
}


type Ipv6RouteRibIpv6NexthopTunnel struct {
    Tunnel int `json:"tunnel"`
    Ipv6NexthopTunnelAddr string `json:"ipv6-nexthop-tunnel-addr"`
    DistanceNexthopTunnel int `json:"distance-nexthop-tunnel" dval:"1"`
    Description string `json:"description"`
}

func (p *Ipv6RouteRib) GetId() string{
    return url.QueryEscape(p.Inst.Ipv6Address)
}

func (p *Ipv6RouteRib) getPath() string{
    return "ipv6/route/rib"
}

func (p *Ipv6RouteRib) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteRib::Post")
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

func (p *Ipv6RouteRib) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteRib::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *Ipv6RouteRib) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteRib::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *Ipv6RouteRib) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteRib::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
