

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVridBladeParametersTrackingOptions struct {
	Inst struct {

    Bgp VrrpAVridBladeParametersTrackingOptionsBgp `json:"bgp"`
    Gateway VrrpAVridBladeParametersTrackingOptionsGateway3630 `json:"gateway"`
    Interface []VrrpAVridBladeParametersTrackingOptionsInterface `json:"interface"`
    Route VrrpAVridBladeParametersTrackingOptionsRoute `json:"route"`
    TrunkCfg []VrrpAVridBladeParametersTrackingOptionsTrunkCfg `json:"trunk-cfg"`
    Uuid string `json:"uuid"`
    VlanCfg []VrrpAVridBladeParametersTrackingOptionsVlanCfg `json:"vlan-cfg"`
    VridVal string 

	} `json:"tracking-options"`
}


type VrrpAVridBladeParametersTrackingOptionsBgp struct {
    BgpIpv4AddressCfg []VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg `json:"bgp-ipv4-address-cfg"`
    BgpIpv6AddressCfg []VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg `json:"bgp-ipv6-address-cfg"`
}


type VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg struct {
    BgpIpv4Address string `json:"bgp-ipv4-address"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg struct {
    BgpIpv6Address string `json:"bgp-ipv6-address"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsGateway3630 struct {
    Ipv4GatewayList []VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList `json:"ipv4-gateway-list"`
    Ipv6GatewayList []VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList `json:"ipv6-gateway-list"`
}


type VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList struct {
    IpAddress string `json:"ip-address"`
    PriorityCost int `json:"priority-cost"`
    Uuid string `json:"uuid"`
}


type VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList struct {
    Ipv6Address string `json:"ipv6-address"`
    PriorityCost int `json:"priority-cost"`
    Uuid string `json:"uuid"`
}


type VrrpAVridBladeParametersTrackingOptionsInterface struct {
    Ethernet int `json:"ethernet"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsRoute struct {
    IpDestinationCfg []VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg `json:"ip-destination-cfg"`
    Ipv6DestinationCfg []VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg `json:"ipv6-destination-cfg"`
}


type VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg struct {
    IpDestination string `json:"ip-destination"`
    Mask string `json:"mask"`
    PriorityCost int `json:"priority-cost"`
    Gateway string `json:"gateway"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg struct {
    Ipv6Destination string `json:"ipv6-destination"`
    PriorityCost int `json:"priority-cost"`
    Gatewayv6 string `json:"gatewayv6"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type VrrpAVridBladeParametersTrackingOptionsTrunkCfg struct {
    Trunk int `json:"trunk"`
    PriorityCost int `json:"priority-cost"`
    PerPortPri int `json:"per-port-pri"`
}


type VrrpAVridBladeParametersTrackingOptionsVlanCfg struct {
    Vlan int `json:"vlan"`
    Timeout int `json:"timeout"`
    PriorityCost int `json:"priority-cost"`
}

func (p *VrrpAVridBladeParametersTrackingOptions) GetId() string{
    return "1"
}

func (p *VrrpAVridBladeParametersTrackingOptions) getPath() string{
    return "vrrp-a/vrid/" +p.Inst.VridVal + "/blade-parameters/tracking-options"
}

func (p *VrrpAVridBladeParametersTrackingOptions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptions::Post")
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

func (p *VrrpAVridBladeParametersTrackingOptions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptions::Get")
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
func (p *VrrpAVridBladeParametersTrackingOptions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptions::Put")
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

func (p *VrrpAVridBladeParametersTrackingOptions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVridBladeParametersTrackingOptions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
