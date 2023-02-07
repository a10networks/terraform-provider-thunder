package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterOspf struct {
	ProcessID RouterOspfInstance `json:"ospf,omitempty"`
}

type RouterOspfInstance struct {
	AreaIpv4                   []RouterOspfAreaList               `json:"area-list,omitempty"`
	AutoCostReferenceBandwidth int                                `json:"auto-cost-reference-bandwidth,omitempty"`
	BfdAllInterfaces           int                                `json:"bfd-all-interfaces,omitempty"`
	Originate                  RouterOspfDefaultInformationOspf   `json:"default-information,omitempty"`
	DefaultMetric              int                                `json:"default-metric,omitempty"`
	DistanceValue              RouterOspfDistance                 `json:"distance,omitempty"`
	DiType                     []RouterOspfDistributeInternalList `json:"distribute-internal-list,omitempty"`
	Direction                  []RouterOspfDistributeLists        `json:"distribute-lists,omitempty"`
	ExtraCost                  []RouterOspfHaStandbyExtraCost     `json:"ha-standby-extra-cost,omitempty"`
	HostAddress                []RouterOspfHostList               `json:"host-list,omitempty"`
	State                      RouterOspfLogAdjacencyChangesCfg   `json:"log-adjacency-changes-cfg,omitempty"`
	MaxConcurrentDd            int                                `json:"max-concurrent-dd,omitempty"`
	MaximumArea                int                                `json:"maximum-area,omitempty"`
	Address                    []RouterOspfNeighborList           `json:"neighbor-list,omitempty"`
	NetworkIpv4                []RouterOspfNetworkList            `json:"network-list,omitempty"`
	AbrType                    RouterOspfOspf1                    `json:"ospf-1,omitempty"`
	Database                   RouterOspfOverflow                 `json:"overflow,omitempty"`
	LoopbackCfg                RouterOspfPassiveInterface         `json:"passive-interface,omitempty"`
	ProcessID                  int                                `json:"process-id,omitempty"`
	RedistList                 RouterOspfRedistributeOspf         `json:"redistribute,omitempty"`
	Rfc1583Compatible          int                                `json:"rfc1583-compatible,omitempty"`
	Value                      RouterOspfRouterID                 `json:"router-id,omitempty"`
	SummaryAddress             []RouterOspfSummaryAddressList     `json:"summary-address-list,omitempty"`
	Spf                        RouterOspfTimers                   `json:"timers,omitempty"`
	UUID                       string                             `json:"uuid,omitempty"`
	UserTag                    string                             `json:"user-tag,omitempty"`
}

type RouterOspfAreaList struct {
	AreaIpv4          string                      `json:"area-ipv4,omitempty"`
	AreaNum           int                         `json:"area-num,omitempty"`
	Authentication    RouterOspfAuthCfg           `json:"auth-cfg,omitempty"`
	DefaultCost       int                         `json:"default-cost,omitempty"`
	FilterList        []RouterOspfFilterLists     `json:"filter-lists,omitempty"`
	Nssa              RouterOspfNssaCfg           `json:"nssa-cfg,omitempty"`
	AreaRangePrefix   []RouterOspfRangeList       `json:"range-list,omitempty"`
	Shortcut          string                      `json:"shortcut,omitempty"`
	Stub              RouterOspfStubCfg           `json:"stub-cfg,omitempty"`
	UUID              string                      `json:"uuid,omitempty"`
	VirtualLinkIPAddr []RouterOspfVirtualLinkList `json:"virtual-link-list,omitempty"`
}

type RouterOspfDefaultInformationOspf struct {
	Always     int    `json:"always,omitempty"`
	Metric     int    `json:"metric,omitempty"`
	MetricType int    `json:"metric-type,omitempty"`
	Originate  int    `json:"originate,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}

type RouterOspfDistance struct {
	DistanceExternal RouterOspfDistanceOspf `json:"distance-ospf,omitempty"`
	DistanceValue    int                    `json:"distance-value,omitempty"`
}

type RouterOspfDistributeInternalList struct {
	DiAreaIpv4 string `json:"di-area-ipv4,omitempty"`
	DiAreaNum  int    `json:"di-area-num,omitempty"`
	DiCost     int    `json:"di-cost,omitempty"`
	DiType     string `json:"di-type,omitempty"`
}

type RouterOspfDistributeLists struct {
	Direction string `json:"direction,omitempty"`
	Option    string `json:"option,omitempty"`
	OspfID    int    `json:"ospf-id,omitempty"`
	Protocol  string `json:"protocol,omitempty"`
	Value     string `json:"value,omitempty"`
}

type RouterOspfHaStandbyExtraCost struct {
	ExtraCost int `json:"extra-cost,omitempty"`
	Group     int `json:"group,omitempty"`
}

type RouterOspfHostList struct {
	AreaIpv4    RouterOspfAreaCfg `json:"area-cfg,omitempty"`
	HostAddress string            `json:"host-address,omitempty"`
}

type RouterOspfLogAdjacencyChangesCfg struct {
	State string `json:"state,omitempty"`
}

type RouterOspfNeighborList struct {
	Address      string `json:"address,omitempty"`
	Cost         int    `json:"cost,omitempty"`
	PollInterval int    `json:"poll-interval,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type RouterOspfNetworkList struct {
	NetworkAreaIpv4 RouterOspfNetworkArea `json:"network-area,omitempty"`
	NetworkIpv4     string                `json:"network-ipv4,omitempty"`
	NetworkIpv4Cidr string                `json:"network-ipv4-cidr,omitempty"`
	NetworkIpv4Mask string                `json:"network-ipv4-mask,omitempty"`
}

type RouterOspfOspf1 struct {
	Option RouterOspfAbrType `json:"abr-type,omitempty"`
}

type RouterOspfOverflow struct {
	Count RouterOspfDatabase `json:"database,omitempty"`
}

type RouterOspfPassiveInterface struct {
	Ethernet []RouterOspfEthCfg      `json:"eth-cfg,omitempty"`
	Lif      []RouterOspfLifCfg      `json:"lif-cfg,omitempty"`
	Loopback []RouterOspfLoopbackCfg `json:"loopback-cfg,omitempty"`
	Trunk    []RouterOspfTrunkCfg    `json:"trunk-cfg,omitempty"`
	Tunnel   []RouterOspfTunnelCfg   `json:"tunnel-cfg,omitempty"`
	Ve       []RouterOspfVeCfg       `json:"ve-cfg,omitempty"`
}

type RouterOspfRedistributeOspf struct {
	IPNat           int                           `json:"ip-nat,omitempty"`
	IPNatPrefix     []RouterOspfIPNatFloatingList `json:"ip-nat-floating-list,omitempty"`
	MetricIPNat     int                           `json:"metric-ip-nat,omitempty"`
	MetricTypeIPNat string                        `json:"metric-type-ip-nat,omitempty"`
	Ospf            []RouterOspfOspfList          `json:"ospf-list,omitempty"`
	Type            []RouterOspfRedistList        `json:"redist-list,omitempty"`
	RouteMapIPNat   string                        `json:"route-map-ip-nat,omitempty"`
	TagIPNat        int                           `json:"tag-ip-nat,omitempty"`
	UUID            string                        `json:"uuid,omitempty"`
	VipAddress      []RouterOspfVipFloatingList   `json:"vip-floating-list,omitempty"`
	TypeVip         []RouterOspfVipList           `json:"vip-list,omitempty"`
}

type RouterOspfRouterID struct {
	Value string `json:"value,omitempty"`
}

type RouterOspfSummaryAddressList struct {
	NotAdvertise   int    `json:"not-advertise,omitempty"`
	SummaryAddress string `json:"summary-address,omitempty"`
	Tag            int    `json:"tag,omitempty"`
}

type RouterOspfTimers struct {
	Exp RouterOspfSpf `json:"spf,omitempty"`
}

type RouterOspfAuthCfg struct {
	Authentication int `json:"authentication,omitempty"`
	MessageDigest  int `json:"message-digest,omitempty"`
}

type RouterOspfFilterLists struct {
	AclDirection   string `json:"acl-direction,omitempty"`
	AclName        string `json:"acl-name,omitempty"`
	FilterList     int    `json:"filter-list,omitempty"`
	PlistDirection string `json:"plist-direction,omitempty"`
	PlistName      string `json:"plist-name,omitempty"`
}

type RouterOspfNssaCfg struct {
	DefaultInformationOriginate int    `json:"default-information-originate,omitempty"`
	Metric                      int    `json:"metric,omitempty"`
	MetricType                  int    `json:"metric-type,omitempty"`
	NoRedistribution            int    `json:"no-redistribution,omitempty"`
	NoSummary                   int    `json:"no-summary,omitempty"`
	Nssa                        int    `json:"nssa,omitempty"`
	TranslatorRole              string `json:"translator-role,omitempty"`
}

type RouterOspfRangeList struct {
	AreaRangePrefix string `json:"area-range-prefix,omitempty"`
	Option          string `json:"option,omitempty"`
}

type RouterOspfStubCfg struct {
	NoSummary int `json:"no-summary,omitempty"`
	Stub      int `json:"stub,omitempty"`
}

type RouterOspfVirtualLinkList struct {
	AuthenticationKey         string `json:"authentication-key,omitempty"`
	Bfd                       int    `json:"bfd,omitempty"`
	DeadInterval              int    `json:"dead-interval,omitempty"`
	HelloInterval             int    `json:"hello-interval,omitempty"`
	Md5                       string `json:"md5,omitempty"`
	MessageDigestKey          int    `json:"message-digest-key,omitempty"`
	RetransmitInterval        int    `json:"retransmit-interval,omitempty"`
	TransmitDelay             int    `json:"transmit-delay,omitempty"`
	VirtualLinkAuthType       string `json:"virtual-link-auth-type,omitempty"`
	VirtualLinkAuthentication int    `json:"virtual-link-authentication,omitempty"`
	VirtualLinkIPAddr         string `json:"virtual-link-ip-addr,omitempty"`
}

type RouterOspfDistanceOspf struct {
	DistanceExternal  int `json:"distance-external,omitempty"`
	DistanceInterArea int `json:"distance-inter-area,omitempty"`
	DistanceIntraArea int `json:"distance-intra-area,omitempty"`
}

type RouterOspfAreaCfg struct {
	AreaIpv4 string `json:"area-ipv4,omitempty"`
	AreaNum  int    `json:"area-num,omitempty"`
	Cost     int    `json:"cost,omitempty"`
}

type RouterOspfNetworkArea struct {
	InstanceValue   int    `json:"instance-value,omitempty"`
	NetworkAreaIpv4 string `json:"network-area-ipv4,omitempty"`
	NetworkAreaNum  int    `json:"network-area-num,omitempty"`
}

type RouterOspfAbrType struct {
	Option string `json:"option,omitempty"`
}

type RouterOspfDatabase struct {
	Count        int    `json:"count,omitempty"`
	DbExternal   int    `json:"db-external,omitempty"`
	Limit        string `json:"limit,omitempty"`
	RecoveryTime int    `json:"recovery-time,omitempty"`
}

type RouterOspfEthCfg struct {
	EthAddress string `json:"eth-address,omitempty"`
	Ethernet   int    `json:"ethernet,omitempty"`
}

type RouterOspfLifCfg struct {
	Lif        int    `json:"lif,omitempty"`
	LifAddress string `json:"lif-address,omitempty"`
}

type RouterOspfLoopbackCfg struct {
	Loopback        int    `json:"loopback,omitempty"`
	LoopbackAddress string `json:"loopback-address,omitempty"`
}

type RouterOspfTrunkCfg struct {
	Trunk        int    `json:"trunk,omitempty"`
	TrunkAddress string `json:"trunk-address,omitempty"`
}

type RouterOspfTunnelCfg struct {
	Tunnel        int    `json:"tunnel,omitempty"`
	TunnelAddress string `json:"tunnel-address,omitempty"`
}

type RouterOspfVeCfg struct {
	Ve        int    `json:"ve,omitempty"`
	VeAddress string `json:"ve-address,omitempty"`
}

type RouterOspfIPNatFloatingList struct {
	IPNatFloatingIPForward string `json:"ip-nat-floating-IP-forward,omitempty"`
	IPNatPrefix            string `json:"ip-nat-prefix,omitempty"`
}

type RouterOspfOspfList struct {
	MetricOspf     int    `json:"metric-ospf,omitempty"`
	MetricTypeOspf string `json:"metric-type-ospf,omitempty"`
	Ospf           int    `json:"ospf,omitempty"`
	ProcessID      int    `json:"process-id,omitempty"`
	RouteMapOspf   string `json:"route-map-ospf,omitempty"`
	TagOspf        int    `json:"tag-ospf,omitempty"`
}

type RouterOspfRedistList struct {
	Metric     int    `json:"metric,omitempty"`
	MetricType string `json:"metric-type,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
	Tag        int    `json:"tag,omitempty"`
	Type       string `json:"type,omitempty"`
}

type RouterOspfVipFloatingList struct {
	VipAddress           string `json:"vip-address,omitempty"`
	VipFloatingIPForward string `json:"vip-floating-IP-forward,omitempty"`
}

type RouterOspfVipList struct {
	MetricTypeVip string `json:"metric-type-vip,omitempty"`
	MetricVip     int    `json:"metric-vip,omitempty"`
	RouteMapVip   string `json:"route-map-vip,omitempty"`
	TagVip        int    `json:"tag-vip,omitempty"`
	TypeVip       string `json:"type-vip,omitempty"`
}

type RouterOspfSpf struct {
	MinDelay RouterOspfExp `json:"exp,omitempty"`
}

type RouterOspfExp struct {
	MaxDelay int `json:"max-delay,omitempty"`
	MinDelay int `json:"min-delay,omitempty"`
}

func PostRouterOspf(id string, inst RouterOspf, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterOspf")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/ospf", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterOspf", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterOspf(id string, name1 string, host string) (*RouterOspf, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterOspf")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/ospf/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterOspf", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterOspf(id string, name1 string, inst RouterOspf, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterOspf")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/ospf/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterOspf", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRouterOspf(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterOspf")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/ospf/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouterOspf", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
