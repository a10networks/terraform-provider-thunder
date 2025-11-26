package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type VrrpAVridBladeParameters struct {
	Inst struct {
		FailOverPolicyTemplate string `json:"fail-over-policy-template"`

		Priority int `json:"priority" dval:"150"`

		TrackingOptions VrrpAVridBladeParametersTrackingOptions3730 `json:"tracking-options"`

		Uuid string `json:"uuid"`

		VridVal string
	} `json:"blade-parameters"`
}

type VrrpAVridBladeParametersTrackingOptions3730 struct {
	Interface []VrrpAVridBladeParametersTrackingOptionsInterface3731 `json:"interface"`
	Route     VrrpAVridBladeParametersTrackingOptionsRoute3732       `json:"route"`
	TrunkCfg  []VrrpAVridBladeParametersTrackingOptionsTrunkCfg3735  `json:"trunk-cfg"`
	Bgp       VrrpAVridBladeParametersTrackingOptionsBgp3736         `json:"bgp"`
	VlanCfg   []VrrpAVridBladeParametersTrackingOptionsVlanCfg3739   `json:"vlan-cfg"`
	Uuid      string                                                 `json:"uuid"`
	Gateway   VrrpAVridBladeParametersTrackingOptionsGateway3740     `json:"gateway"`
}

type VrrpAVridBladeParametersTrackingOptionsInterface3731 struct {
	Ethernet     int `json:"ethernet"`
	PriorityCost int `json:"priority-cost"`
}

type VrrpAVridBladeParametersTrackingOptionsRoute3732 struct {
	IpDestinationCfg   []VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733   `json:"ip-destination-cfg"`
	Ipv6DestinationCfg []VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734 `json:"ipv6-destination-cfg"`
}

type VrrpAVridBladeParametersTrackingOptionsRouteIpDestinationCfg3733 struct {
	IpDestination string `json:"ip-destination"`
	Mask          string `json:"mask"`
	PriorityCost  int    `json:"priority-cost"`
	Gateway       string `json:"gateway"`
	Distance      int    `json:"distance"`
	Protocol      string `json:"protocol" dval:"any"`
}

type VrrpAVridBladeParametersTrackingOptionsRouteIpv6DestinationCfg3734 struct {
	Ipv6Destination string `json:"ipv6-destination"`
	PriorityCost    int    `json:"priority-cost"`
	Gatewayv6       string `json:"gatewayv6"`
	Distance        int    `json:"distance"`
	Protocol        string `json:"protocol" dval:"any"`
}

type VrrpAVridBladeParametersTrackingOptionsTrunkCfg3735 struct {
	Trunk        int `json:"trunk"`
	PriorityCost int `json:"priority-cost"`
	PerPortPri   int `json:"per-port-pri"`
}

type VrrpAVridBladeParametersTrackingOptionsBgp3736 struct {
	BgpIpv4AddressCfg []VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737 `json:"bgp-ipv4-address-cfg"`
	BgpIpv6AddressCfg []VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738 `json:"bgp-ipv6-address-cfg"`
}

type VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg3737 struct {
	BgpIpv4Address string `json:"bgp-ipv4-address"`
	PriorityCost   int    `json:"priority-cost"`
}

type VrrpAVridBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg3738 struct {
	BgpIpv6Address string `json:"bgp-ipv6-address"`
	PriorityCost   int    `json:"priority-cost"`
}

type VrrpAVridBladeParametersTrackingOptionsVlanCfg3739 struct {
	Vlan         int `json:"vlan"`
	Timeout      int `json:"timeout"`
	PriorityCost int `json:"priority-cost"`
}

type VrrpAVridBladeParametersTrackingOptionsGateway3740 struct {
	Ipv4GatewayList []VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741 `json:"ipv4-gateway-list"`
	Ipv6GatewayList []VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742 `json:"ipv6-gateway-list"`
}

type VrrpAVridBladeParametersTrackingOptionsGatewayIpv4GatewayList3741 struct {
	IpAddress    string `json:"ip-address"`
	PriorityCost int    `json:"priority-cost"`
	Uuid         string `json:"uuid"`
}

type VrrpAVridBladeParametersTrackingOptionsGatewayIpv6GatewayList3742 struct {
	Ipv6Address  string `json:"ipv6-address"`
	PriorityCost int    `json:"priority-cost"`
	Uuid         string `json:"uuid"`
}

func (p *VrrpAVridBladeParameters) GetId() string {
	return "1"
}

func (p *VrrpAVridBladeParameters) getPath() string {
	return "vrrp-a/vrid/" + p.Inst.VridVal + "/blade-parameters"
}

func (p *VrrpAVridBladeParameters) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpAVridBladeParameters::Post")
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

func (p *VrrpAVridBladeParameters) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpAVridBladeParameters::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *VrrpAVridBladeParameters) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpAVridBladeParameters::Put")
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

func (p *VrrpAVridBladeParameters) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("VrrpAVridBladeParameters::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
