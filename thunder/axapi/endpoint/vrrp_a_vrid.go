

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVrid struct {
	Inst struct {

    BladeParameters VrrpAVridBladeParameters3644 `json:"blade-parameters"`
    FloatingIp VrrpAVridFloatingIp `json:"floating-ip"`
    Follow VrrpAVridFollow `json:"follow"`
    PairFollow VrrpAVridPairFollow `json:"pair-follow"`
    PreemptMode VrrpAVridPreemptMode `json:"preempt-mode"`
    SamplingEnable []VrrpAVridSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VridVal int `json:"vrid-val"`

	} `json:"vrid"`
}


type VrrpAVridBladeParameters3644 struct {
    Priority int `json:"priority" dval:"150"`
    FailOverPolicyTemplate string `json:"fail-over-policy-template"`
    Uuid string `json:"uuid"`
    TrackingOptions VrrpAVridBladeParametersTrackingOptions3645 `json:"tracking-options"`
}


type VrrpAVridBladeParametersTrackingOptions3645 struct {
    Interface []VrrpAVridBladeParametersTrackingOptionsInterface3646 `json:"interface"`
    Route VrrpAVridBladeParametersTrackingOptionsRoute3647 `json:"route"`
    TrunkCfg []VrrpAVridBladeParametersTrackingOptionsTrunkCfg3650 `json:"trunk-cfg"`
    Bgp VrrpAVridBladeParametersTrackingOptionsBgp3651 `json:"bgp"`
    VlanCfg []VrrpAVridBladeParametersTrackingOptionsVlanCfg3654 `json:"vlan-cfg"`
    Uuid string `json:"uuid"`
    Gateway VrrpAVridBladeParametersTrackingOptionsGateway3655 `json:"gateway"`
}


type VrrpAVridBladeParametersTrackingOptionsInterface3646 struct {
    Ethernet int `json:"ethernet"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsRoute3647 struct {
    IpDestinationCfg []VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648 `json:"ip-destination-cfg"`
    Ipv6DestinationCfg []VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649 `json:"ipv6-destination-cfg"`
}


type VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3648 struct {
    IpDestination string `json:"ip-destination"`
    Mask string `json:"mask"`
    PriorityCost int `json:"priority-cost"`
    Gateway string `json:"gateway"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3649 struct {
    Ipv6Destination string `json:"ipv6-destination"`
    PriorityCost int `json:"priority-cost"`
    Gatewayv6 string `json:"gatewayv6"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type VrrpAVridBladeParametersTrackingOptionsTrunkCfg3650 struct {
    Trunk int `json:"trunk"`
    PriorityCost int `json:"priority-cost"`
    PerPortPri int `json:"per-port-pri"`
}


type VrrpAVridBladeParametersTrackingOptionsBgp3651 struct {
    BgpIpv4AddressCfg []VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652 `json:"bgp-ipv4-address-cfg"`
    BgpIpv6AddressCfg []VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653 `json:"bgp-ipv6-address-cfg"`
}


type VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3652 struct {
    BgpIpv4Address string `json:"bgp-ipv4-address"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3653 struct {
    BgpIpv6Address string `json:"bgp-ipv6-address"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsVlanCfg3654 struct {
    Vlan int `json:"vlan"`
    Timeout int `json:"timeout"`
    PriorityCost int `json:"priority-cost"`
}


type VrrpAVridBladeParametersTrackingOptionsGateway3655 struct {
    Ipv4GatewayList []VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656 `json:"ipv4-gateway-list"`
    Ipv6GatewayList []VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657 `json:"ipv6-gateway-list"`
}


type VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3656 struct {
    IpAddress string `json:"ip-address"`
    PriorityCost int `json:"priority-cost"`
    Uuid string `json:"uuid"`
}


type VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3657 struct {
    Ipv6Address string `json:"ipv6-address"`
    PriorityCost int `json:"priority-cost"`
    Uuid string `json:"uuid"`
}


type VrrpAVridFloatingIp struct {
    IpAddressCfg []VrrpAVridFloatingIpIpAddressCfg `json:"ip-address-cfg"`
    IpAddressPartCfg []VrrpAVridFloatingIpIpAddressPartCfg `json:"ip-address-part-cfg"`
    Ipv6AddressCfg []VrrpAVridFloatingIpIpv6AddressCfg `json:"ipv6-address-cfg"`
    Ipv6AddressPartCfg []VrrpAVridFloatingIpIpv6AddressPartCfg `json:"ipv6-address-part-cfg"`
}


type VrrpAVridFloatingIpIpAddressCfg struct {
    IpAddress string `json:"ip-address"`
}


type VrrpAVridFloatingIpIpAddressPartCfg struct {
    IpAddressPartition string `json:"ip-address-partition"`
}


type VrrpAVridFloatingIpIpv6AddressCfg struct {
    Ipv6Address string `json:"ipv6-address"`
    Ethernet int `json:"ethernet"`
    Trunk int `json:"trunk"`
    Ve int `json:"ve"`
}


type VrrpAVridFloatingIpIpv6AddressPartCfg struct {
    Ipv6AddressPartition string `json:"ipv6-address-partition"`
    Ethernet int `json:"ethernet"`
    Trunk int `json:"trunk"`
    Ve int `json:"ve"`
}


type VrrpAVridFollow struct {
    VridLead string `json:"vrid-lead"`
}


type VrrpAVridPairFollow struct {
    PairFollow int `json:"pair-follow"`
    VridLead string `json:"vrid-lead"`
}


type VrrpAVridPreemptMode struct {
    Threshold int `json:"threshold"`
    Disable int `json:"disable"`
}


type VrrpAVridSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *VrrpAVrid) GetId() string{
    return strconv.Itoa(p.Inst.VridVal)
}

func (p *VrrpAVrid) getPath() string{
    return "vrrp-a/vrid"
}

func (p *VrrpAVrid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVrid::Post")
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

func (p *VrrpAVrid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVrid::Get")
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
func (p *VrrpAVrid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVrid::Put")
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

func (p *VrrpAVrid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAVrid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
