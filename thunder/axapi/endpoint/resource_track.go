

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ResourceTrack struct {
	Inst struct {

    Bgp ResourceTrackBgp `json:"bgp"`
    Gateway ResourceTrackGateway `json:"gateway"`
    Interface []ResourceTrackInterface `json:"interface"`
    Name string `json:"name"`
    Route ResourceTrackRoute `json:"route"`
    ScaleoutCfg ResourceTrackScaleoutCfg `json:"scaleout-cfg"`
    TrunkCfg []ResourceTrackTrunkCfg `json:"trunk-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VlanCfg []ResourceTrackVlanCfg `json:"vlan-cfg"`

	} `json:"resource-track"`
}


type ResourceTrackBgp struct {
    BgpIpv4AddressCfg []ResourceTrackBgpBgpIpv4AddressCfg `json:"bgp-ipv4-address-cfg"`
    BgpIpv6AddressCfg []ResourceTrackBgpBgpIpv6AddressCfg `json:"bgp-ipv6-address-cfg"`
}


type ResourceTrackBgpBgpIpv4AddressCfg struct {
    BgpIpv4Address string `json:"bgp-ipv4-address"`
    Weight int `json:"weight"`
}


type ResourceTrackBgpBgpIpv6AddressCfg struct {
    BgpIpv6Address string `json:"bgp-ipv6-address"`
    Weight int `json:"weight"`
}


type ResourceTrackGateway struct {
    GwIpv4AddressCfg []ResourceTrackGatewayGwIpv4AddressCfg `json:"gw-ipv4-address-cfg"`
    GwIpv6AddressCfg []ResourceTrackGatewayGwIpv6AddressCfg `json:"gw-ipv6-address-cfg"`
}


type ResourceTrackGatewayGwIpv4AddressCfg struct {
    GwIpv4Address string `json:"gw-ipv4-address"`
    Weight int `json:"weight"`
}


type ResourceTrackGatewayGwIpv6AddressCfg struct {
    GwIpv6Address string `json:"gw-ipv6-address"`
    Weight int `json:"weight"`
}


type ResourceTrackInterface struct {
    Ethernet int `json:"ethernet"`
    Weight int `json:"weight"`
}


type ResourceTrackRoute struct {
    IpDestinationCfg []ResourceTrackRouteIpDestinationCfg `json:"ip-destination-cfg"`
    Ipv6DestinationCfg []ResourceTrackRouteIpv6DestinationCfg `json:"ipv6-destination-cfg"`
}


type ResourceTrackRouteIpDestinationCfg struct {
    IpDestination string `json:"ip-destination"`
    Mask string `json:"mask"`
    Weight int `json:"weight"`
    Gateway string `json:"gateway"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type ResourceTrackRouteIpv6DestinationCfg struct {
    Ipv6Destination string `json:"ipv6-destination"`
    Weight int `json:"weight"`
    Gatewayv6 string `json:"gatewayv6"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type ResourceTrackScaleoutCfg struct {
    Scaleout int `json:"scaleout"`
    Weight int `json:"weight"`
}


type ResourceTrackTrunkCfg struct {
    Trunk int `json:"trunk"`
    Weight int `json:"weight"`
    PerPortWeight int `json:"per-port-weight"`
}


type ResourceTrackVlanCfg struct {
    Vlan int `json:"vlan"`
    Timeout int `json:"timeout"`
    Weight int `json:"weight"`
}

func (p *ResourceTrack) GetId() string{
    return p.Inst.Name
}

func (p *ResourceTrack) getPath() string{
    return "resource-track"
}

func (p *ResourceTrack) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ResourceTrack::Post")
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

func (p *ResourceTrack) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ResourceTrack::Get")
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
func (p *ResourceTrack) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ResourceTrack::Put")
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

func (p *ResourceTrack) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ResourceTrack::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
