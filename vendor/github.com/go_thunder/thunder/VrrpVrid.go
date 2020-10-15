package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type VridInstance struct {
	UUID Vrid `json:"vrid,omitempty"`
}
type VridPairFollow struct {
	PairFollow int    `json:"pair-follow,omitempty"`
	VridLead   string `json:"vrid-lead,omitempty"`
}
type VridVlanCfg struct {
	Vlan         int `json:"vlan,omitempty"`
	Timeout      int `json:"timeout,omitempty"`
	PriorityCost int `json:"priority-cost,omitempty"`
}
type VridIpv6DestinationCfg struct {
	Distance        int    `json:"distance,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	PriorityCost    int    `json:"priority-cost,omitempty"`
	Ipv6Destination string `json:"ipv6-destination,omitempty"`
	Gatewayv6       string `json:"gatewayv6,omitempty"`
}
type VridIPDestinationCfg struct {
	Distance      int    `json:"distance,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Mask          string `json:"mask,omitempty"`
	PriorityCost  int    `json:"priority-cost,omitempty"`
	IPDestination string `json:"ip-destination,omitempty"`
	Gateway       string `json:"gateway,omitempty"`
}
type VridRoute struct {
	Distance []VridIpv6DestinationCfg `json:"ipv6-destination-cfg,omitempty"`
	Protocol []VridIPDestinationCfg   `json:"ip-destination-cfg,omitempty"`
}
type VridBgpIpv4AddressCfg struct {
	BgpIpv4Address string `json:"bgp-ipv4-address,omitempty"`
	PriorityCost   int    `json:"priority-cost,omitempty"`
}
type VridBgpIpv6AddressCfg struct {
	BgpIpv6Address string `json:"bgp-ipv6-address,omitempty"`
	PriorityCost   int    `json:"priority-cost,omitempty"`
}
type VridBgp struct {
	BgpIpv4Address []VridBgpIpv4AddressCfg `json:"bgp-ipv4-address-cfg,omitempty"`
	BgpIpv6Address []VridBgpIpv6AddressCfg `json:"bgp-ipv6-address-cfg,omitempty"`
}
type VridInterface struct {
	Ethernet     int `json:"ethernet,omitempty"`
	PriorityCost int `json:"priority-cost,omitempty"`
}
type VridIpv4GatewayList struct {
	UUID         string `json:"uuid,omitempty"`
	IPAddress    string `json:"ip-address,omitempty"`
	PriorityCost int    `json:"priority-cost,omitempty"`
}
type VridIpv6GatewayList struct {
	Ipv6Address  string `json:"ipv6-address,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	PriorityCost int    `json:"priority-cost,omitempty"`
}
type VridGateway struct {
	IPAddress   []VridIpv4GatewayList `json:"ipv4-gateway-list,omitempty"`
	Ipv6Address []VridIpv6GatewayList `json:"ipv6-gateway-list,omitempty"`
}
type VridTrunkCfg struct {
	PriorityCost int `json:"priority-cost,omitempty"`
	Trunk        int `json:"trunk,omitempty"`
	PerPortPri   int `json:"per-port-pri,omitempty"`
}
type VridTrackingOptions struct {
	Vlan           []VridVlanCfg   `json:"vlan-cfg,omitempty"`
	UUID           string          `json:"uuid,omitempty"`
	Distance       VridRoute       `json:"route,omitempty"`
	BgpIpv4Address VridBgp         `json:"bgp,omitempty"`
	Ethernet       []VridInterface `json:"interface,omitempty"`
	IPAddress      VridGateway     `json:"gateway,omitempty"`
	PriorityCost   []VridTrunkCfg  `json:"trunk-cfg,omitempty"`
}
type VridBladeParameters struct {
	Priority               int                 `json:"priority,omitempty"`
	FailOverPolicyTemplate string              `json:"fail-over-policy-template,omitempty"`
	UUID                   string              `json:"uuid,omitempty"`
	Vlan                   VridTrackingOptions `json:"tracking-options,omitempty"`
}
type VridPreemptMode struct {
	Threshold int `json:"threshold,omitempty"`
	Disable   int `json:"disable,omitempty"`
}
type VridSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}
type VridIpv6AddressPartCfg struct {
	Ethernet             int    `json:"ethernet,omitempty"`
	Ipv6AddressPartition string `json:"ipv6-address-partition,omitempty"`
	Ve                   int    `json:"ve,omitempty"`
	Trunk                int    `json:"trunk,omitempty"`
}
type VridIPAddressCfg struct {
	IPAddress string `json:"ip-address,omitempty"`
}
type VridIPAddressPartCfg struct {
	IPAddressPartition string `json:"ip-address-partition,omitempty"`
}
type VridIpv6AddressCfg struct {
	Ipv6Address string `json:"ipv6-address,omitempty"`
	Ethernet    int    `json:"ethernet,omitempty"`
	Ve          int    `json:"ve,omitempty"`
	Trunk       int    `json:"trunk,omitempty"`
}
type VridFloatingIP struct {
	Ethernet           []VridIpv6AddressPartCfg `json:"ipv6-address-part-cfg,omitempty"`
	IPAddress          []VridIPAddressCfg       `json:"ip-address-cfg,omitempty"`
	IPAddressPartition []VridIPAddressPartCfg   `json:"ip-address-part-cfg,omitempty"`
	Ipv6Address        []VridIpv6AddressCfg     `json:"ipv6-address-cfg,omitempty"`
}
type VridFollow struct {
	VridLead string `json:"vrid-lead,omitempty"`
}
type Vrid struct {
	PairFollow VridPairFollow       `json:"pair-follow,omitempty"`
	Priority   VridBladeParameters  `json:"blade-parameters,omitempty"`
	UUID       string               `json:"uuid,omitempty"`
	VridVal    int                  `json:"vrid-val,omitempty"`
	UserTag    string               `json:"user-tag,omitempty"`
	Threshold  VridPreemptMode      `json:"preempt-mode,omitempty"`
	Counters1  []VridSamplingEnable `json:"sampling-enable,omitempty"`
	Ethernet   VridFloatingIP       `json:"floating-ip,omitempty"`
	VridLead   VridFollow           `json:"follow,omitempty"`
}

func GetVrrpVrid(id string, name string, host string) (*VridInstance, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VridInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}
}

func PostVrrpVrid(id string, vc VridInstance, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/vrrp-a/vrid", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v VridInstance
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func PutVrrpVrid(id string, name string, vc VridInstance, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] Inside putVrrpVrid")

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VridInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func DeleteVrrpVrid(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/vrrp-a/vrid/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VridInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
