package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type VrrpAVrid struct {
	VrrpAVridInstanceVridVal VrrpAVridInstance `json:"vrid,omitempty"`
}

type VrrpAVridInstance struct {
	VrrpAVridInstanceBladeParametersPriority VrrpAVridInstanceBladeParameters  `json:"blade-parameters,omitempty"`
	VrrpAVridInstanceFloatingIPIPAddressCfg  VrrpAVridInstanceFloatingIP       `json:"floating-ip,omitempty"`
	VrrpAVridInstanceFollowVridLead          VrrpAVridInstanceFollow           `json:"follow,omitempty"`
	VrrpAVridInstancePairFollowPairFollow    VrrpAVridInstancePairFollow       `json:"pair-follow,omitempty"`
	VrrpAVridInstancePreemptModeThreshold    VrrpAVridInstancePreemptMode      `json:"preempt-mode,omitempty"`
	VrrpAVridInstanceSamplingEnableCounters1 []VrrpAVridInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	VrrpAVridInstanceUUID                    string                            `json:"uuid,omitempty"`
	VrrpAVridInstanceUserTag                 string                            `json:"user-tag,omitempty"`
	VrrpAVridInstanceVridVal                 int                               `json:"vrid-val,omitempty"`
}

type VrrpAVridInstanceBladeParameters struct {
	VrrpAVridInstanceBladeParametersFailOverPolicyTemplate   string                                          `json:"fail-over-policy-template,omitempty"`
	VrrpAVridInstanceBladeParametersPriority                 int                                             `json:"priority,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsInterface VrrpAVridInstanceBladeParametersTrackingOptions `json:"tracking-options,omitempty"`
	VrrpAVridInstanceBladeParametersUUID                     string                                          `json:"uuid,omitempty"`
}

type VrrpAVridInstanceFloatingIP struct {
	VrrpAVridInstanceFloatingIPIPAddressCfgIPAddress                  []VrrpAVridInstanceFloatingIPIPAddressCfg       `json:"ip-address-cfg,omitempty"`
	VrrpAVridInstanceFloatingIPIPAddressPartCfgIPAddressPartition     []VrrpAVridInstanceFloatingIPIPAddressPartCfg   `json:"ip-address-part-cfg,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressCfgIpv6Address              []VrrpAVridInstanceFloatingIPIpv6AddressCfg     `json:"ipv6-address-cfg,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressPartCfgIpv6AddressPartition []VrrpAVridInstanceFloatingIPIpv6AddressPartCfg `json:"ipv6-address-part-cfg,omitempty"`
}

type VrrpAVridInstanceFollow struct {
	VrrpAVridInstanceFollowVridLead string `json:"vrid-lead,omitempty"`
}

type VrrpAVridInstancePairFollow struct {
	VrrpAVridInstancePairFollowPairFollow int    `json:"pair-follow,omitempty"`
	VrrpAVridInstancePairFollowVridLead   string `json:"vrid-lead,omitempty"`
}

type VrrpAVridInstancePreemptMode struct {
	VrrpAVridInstancePreemptModeDisable   int `json:"disable,omitempty"`
	VrrpAVridInstancePreemptModeThreshold int `json:"threshold,omitempty"`
}

type VrrpAVridInstanceSamplingEnable struct {
	VrrpAVridInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptions struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg   VrrpAVridInstanceBladeParametersTrackingOptionsBgp         `json:"bgp,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayList VrrpAVridInstanceBladeParametersTrackingOptionsGateway     `json:"gateway,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceEthernet      []VrrpAVridInstanceBladeParametersTrackingOptionsInterface `json:"interface,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfg  VrrpAVridInstanceBladeParametersTrackingOptionsRoute       `json:"route,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgTrunk          []VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfg  `json:"trunk-cfg,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsUUID                   string                                                     `json:"uuid,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgVlan            []VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfg   `json:"vlan-cfg,omitempty"`
}

type VrrpAVridInstanceFloatingIPIPAddressCfg struct {
	VrrpAVridInstanceFloatingIPIPAddressCfgIPAddress string `json:"ip-address,omitempty"`
}

type VrrpAVridInstanceFloatingIPIPAddressPartCfg struct {
	VrrpAVridInstanceFloatingIPIPAddressPartCfgIPAddressPartition string `json:"ip-address-partition,omitempty"`
}

type VrrpAVridInstanceFloatingIPIpv6AddressCfg struct {
	VrrpAVridInstanceFloatingIPIpv6AddressCfgEthernet    int    `json:"ethernet,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressCfgIpv6Address string `json:"ipv6-address,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressCfgTrunk       int    `json:"trunk,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressCfgVe          int    `json:"ve,omitempty"`
}

type VrrpAVridInstanceFloatingIPIpv6AddressPartCfg struct {
	VrrpAVridInstanceFloatingIPIpv6AddressPartCfgEthernet             int    `json:"ethernet,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressPartCfgIpv6AddressPartition string `json:"ipv6-address-partition,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressPartCfgTrunk                int    `json:"trunk,omitempty"`
	VrrpAVridInstanceFloatingIPIpv6AddressPartCfgVe                   int    `json:"ve,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsBgp struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgBgpIpv4Address []VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg `json:"bgp-ipv4-address-cfg,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgBgpIpv6Address []VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg `json:"bgp-ipv6-address-cfg,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsGateway struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListIPAddress   []VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayList `json:"ipv4-gateway-list,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListIpv6Address []VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayList `json:"ipv6-gateway-list,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsInterface struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsInterfaceEthernet     int `json:"ethernet,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsInterfacePriorityCost int `json:"priority-cost,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsRoute struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgIPDestination     []VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfg   `json:"ip-destination-cfg,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgIpv6Destination []VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfg `json:"ipv6-destination-cfg,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfg struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgPerPortPri   int `json:"per-port-pri,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgPriorityCost int `json:"priority-cost,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsTrunkCfgTrunk        int `json:"trunk,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfg struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgPriorityCost int `json:"priority-cost,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgTimeout      int `json:"timeout,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsVlanCfgVlan         int `json:"vlan,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfg struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgBgpIpv4Address string `json:"bgp-ipv4-address,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv4AddressCfgPriorityCost   int    `json:"priority-cost,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfg struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgBgpIpv6Address string `json:"bgp-ipv6-address,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsBgpBgpIpv6AddressCfgPriorityCost   int    `json:"priority-cost,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayList struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListIPAddress    string `json:"ip-address,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListPriorityCost int    `json:"priority-cost,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv4GatewayListUUID         string `json:"uuid,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayList struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListIpv6Address  string `json:"ipv6-address,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListPriorityCost int    `json:"priority-cost,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsGatewayIpv6GatewayListUUID         string `json:"uuid,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfg struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgDistance      int    `json:"distance,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgGateway       string `json:"gateway,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgIPDestination string `json:"ip-destination,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgMask          string `json:"mask,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgPriorityCost  int    `json:"priority-cost,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIPDestinationCfgProtocol      string `json:"protocol,omitempty"`
}

type VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfg struct {
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgDistance        int    `json:"distance,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgGatewayv6       string `json:"gatewayv6,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgIpv6Destination string `json:"ipv6-destination,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgPriorityCost    int    `json:"priority-cost,omitempty"`
	VrrpAVridInstanceBladeParametersTrackingOptionsRouteIpv6DestinationCfgProtocol        string `json:"protocol,omitempty"`
}

func PostVrrpAVrid(id string, inst VrrpAVrid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostVrrpAVrid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/vrrp-a/vrid", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VrrpAVrid
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostVrrpAVrid", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetVrrpAVrid(id string, name1 string, host string) (*VrrpAVrid, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetVrrpAVrid")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VrrpAVrid
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetVrrpAVrid", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutVrrpAVrid(id string, name1 string, inst VrrpAVrid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutVrrpAVrid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VrrpAVrid
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutVrrpAVrid", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteVrrpAVrid(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteVrrpAVrid")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VrrpAVrid
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteVrrpAVrid", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
