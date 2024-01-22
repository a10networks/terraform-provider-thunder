

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAFailOverPolicyTemplate struct {
	Inst struct {

    Bgp VrrpAFailOverPolicyTemplateBgp `json:"bgp"`
    Gateway VrrpAFailOverPolicyTemplateGateway `json:"gateway"`
    Interface []VrrpAFailOverPolicyTemplateInterface `json:"interface"`
    Name string `json:"name"`
    Route VrrpAFailOverPolicyTemplateRoute `json:"route"`
    TrunkCfg []VrrpAFailOverPolicyTemplateTrunkCfg `json:"trunk-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VlanCfg []VrrpAFailOverPolicyTemplateVlanCfg `json:"vlan-cfg"`

	} `json:"fail-over-policy-template"`
}


type VrrpAFailOverPolicyTemplateBgp struct {
    BgpIpv4AddressCfg []VrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg `json:"bgp-ipv4-address-cfg"`
    BgpIpv6AddressCfg []VrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg `json:"bgp-ipv6-address-cfg"`
}


type VrrpAFailOverPolicyTemplateBgpBgpIpv4AddressCfg struct {
    BgpIpv4Address string `json:"bgp-ipv4-address"`
    Weight int `json:"weight"`
}


type VrrpAFailOverPolicyTemplateBgpBgpIpv6AddressCfg struct {
    BgpIpv6Address string `json:"bgp-ipv6-address"`
    Weight int `json:"weight"`
}


type VrrpAFailOverPolicyTemplateGateway struct {
    GwIpv4AddressCfg []VrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg `json:"gw-ipv4-address-cfg"`
    GwIpv6AddressCfg []VrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg `json:"gw-ipv6-address-cfg"`
}


type VrrpAFailOverPolicyTemplateGatewayGwIpv4AddressCfg struct {
    GwIpv4Address string `json:"gw-ipv4-address"`
    Weight int `json:"weight"`
}


type VrrpAFailOverPolicyTemplateGatewayGwIpv6AddressCfg struct {
    GwIpv6Address string `json:"gw-ipv6-address"`
    Weight int `json:"weight"`
}


type VrrpAFailOverPolicyTemplateInterface struct {
    Ethernet int `json:"ethernet"`
    Weight int `json:"weight"`
}


type VrrpAFailOverPolicyTemplateRoute struct {
    IpDestinationCfg []VrrpAFailOverPolicyTemplateRouteIpDestinationCfg `json:"ip-destination-cfg"`
    Ipv6DestinationCfg []VrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg `json:"ipv6-destination-cfg"`
}


type VrrpAFailOverPolicyTemplateRouteIpDestinationCfg struct {
    IpDestination string `json:"ip-destination"`
    Mask string `json:"mask"`
    Weight int `json:"weight"`
    Gateway string `json:"gateway"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type VrrpAFailOverPolicyTemplateRouteIpv6DestinationCfg struct {
    Ipv6Destination string `json:"ipv6-destination"`
    Weight int `json:"weight"`
    Gatewayv6 string `json:"gatewayv6"`
    Distance int `json:"distance"`
    Protocol string `json:"protocol" dval:"any"`
}


type VrrpAFailOverPolicyTemplateTrunkCfg struct {
    Trunk int `json:"trunk"`
    Weight int `json:"weight"`
    PerPortWeight int `json:"per-port-weight"`
}


type VrrpAFailOverPolicyTemplateVlanCfg struct {
    Vlan int `json:"vlan"`
    Timeout int `json:"timeout"`
    Weight int `json:"weight"`
}

func (p *VrrpAFailOverPolicyTemplate) GetId() string{
    return p.Inst.Name
}

func (p *VrrpAFailOverPolicyTemplate) getPath() string{
    return "vrrp-a/fail-over-policy-template"
}

func (p *VrrpAFailOverPolicyTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAFailOverPolicyTemplate::Post")
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

func (p *VrrpAFailOverPolicyTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAFailOverPolicyTemplate::Get")
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
func (p *VrrpAFailOverPolicyTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAFailOverPolicyTemplate::Put")
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

func (p *VrrpAFailOverPolicyTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAFailOverPolicyTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
