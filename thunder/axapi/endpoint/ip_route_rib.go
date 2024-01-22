

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type IpRouteRib struct {
	Inst struct {

    IpDestAddr string `json:"ip-dest-addr"`
    IpMask string `json:"ip-mask"`
    IpNexthopIpv4 []IpRouteRibIpNexthopIpv4 `json:"ip-nexthop-ipv4"`
    IpNexthopLif []IpRouteRibIpNexthopLif `json:"ip-nexthop-lif"`
    IpNexthopPartition []IpRouteRibIpNexthopPartition `json:"ip-nexthop-partition"`
    IpNexthopTunnel []IpRouteRibIpNexthopTunnel `json:"ip-nexthop-tunnel"`
    Uuid string `json:"uuid"`

	} `json:"rib"`
}


type IpRouteRibIpNexthopIpv4 struct {
    IpNextHop string `json:"ip-next-hop"`
    DistanceNexthopIp int `json:"distance-nexthop-ip" dval:"1"`
    DescriptionNexthopIp string `json:"description-nexthop-ip"`
}


type IpRouteRibIpNexthopLif struct {
    Lif string `json:"lif"`
    DescriptionNexthopLif string `json:"description-nexthop-lif"`
}


type IpRouteRibIpNexthopPartition struct {
    PartitionName string `json:"partition-name"`
    VridNumInPartition int `json:"vrid-num-in-partition"`
    DescriptionNexthopPartition string `json:"description-nexthop-partition"`
    DescriptionPartitionVrid string `json:"description-partition-vrid"`
}


type IpRouteRibIpNexthopTunnel struct {
    Tunnel int `json:"tunnel"`
    IpNextHopTunnel string `json:"ip-next-hop-tunnel"`
    DistanceNexthopTunnel int `json:"distance-nexthop-tunnel" dval:"1"`
    DescriptionNexthopTunnel string `json:"description-nexthop-tunnel"`
}

func (p *IpRouteRib) GetId() string{
    return p.Inst.IpDestAddr+"+"+url.QueryEscape(p.Inst.IpMask)
}

func (p *IpRouteRib) getPath() string{
    return "ip/route/rib"
}

func (p *IpRouteRib) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteRib::Post")
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

func (p *IpRouteRib) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteRib::Get")
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
func (p *IpRouteRib) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteRib::Put")
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

func (p *IpRouteRib) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteRib::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
