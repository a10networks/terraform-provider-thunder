package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgp struct {
	AsNumber RouterBgpInstance `json:"bgp,omitempty"`
}

type RouterBgpInstance struct {
	Ipv6                   RouterBgpAddressFamily          `json:"address-family,omitempty"`
	AggregateAddress       []RouterBgpAggregateAddressList `json:"aggregate-address-list,omitempty"`
	AsNumber               int                             `json:"as-number,omitempty"`
	AutoSummary            int                             `json:"auto-summary,omitempty"`
	AlwaysCompareMed       RouterBgpBgp                    `json:"bgp,omitempty"`
	AdminDistance          []RouterBgpDistanceList         `json:"distance-list,omitempty"`
	MaximumPathsValue      int                             `json:"maximum-paths-value,omitempty"`
	Ethernet               RouterBgpNeighbor               `json:"neighbor,omitempty"`
	NetworkSynchronization RouterBgpNetwork                `json:"network,omitempty"`
	Originate              int                             `json:"originate,omitempty"`
	ConnectedCfg           RouterBgpRedistributeBgp        `json:"redistribute,omitempty"`
	Synchronization        int                             `json:"synchronization,omitempty"`
	BgpKeepalive           RouterBgpTimers                 `json:"timers,omitempty"`
	UUID                   string                          `json:"uuid,omitempty"`
	UserTag                string                          `json:"user-tag,omitempty"`
}

type RouterBgpAddressFamily struct {
	Bgp RouterBgpIpv6 `json:"ipv6,omitempty"`
}

type RouterBgpAggregateAddressList struct {
	AggregateAddress string `json:"aggregate-address,omitempty"`
	AsSet            int    `json:"as-set,omitempty"`
	SummaryOnly      int    `json:"summary-only,omitempty"`
}

type RouterBgpBgp struct {
	AlwaysCompareMed     int                   `json:"always-compare-med,omitempty"`
	Ignore               RouterBgpBestpathCfg  `json:"bestpath-cfg,omitempty"`
	BgpRestartTime       int                   `json:"bgp-restart-time,omitempty"`
	BgpStalepathTime     int                   `json:"bgp-stalepath-time,omitempty"`
	Dampening            RouterBgpDampeningCfg `json:"dampening-cfg,omitempty"`
	DeterministicMed     int                   `json:"deterministic-med,omitempty"`
	EnforceFirstAs       int                   `json:"enforce-first-as,omitempty"`
	FastExternalFailover int                   `json:"fast-external-failover,omitempty"`
	GracefulRestart      int                   `json:"graceful-restart,omitempty"`
	LocalPreferenceValue int                   `json:"local-preference-value,omitempty"`
	LogNeighborChanges   int                   `json:"log-neighbor-changes,omitempty"`
	NexthopTriggerCount  int                   `json:"nexthop-trigger-count,omitempty"`
	OverrideValidation   int                   `json:"override-validation,omitempty"`
	RouterID             string                `json:"router-id,omitempty"`
	ScanTime             int                   `json:"scan-time,omitempty"`
}

type RouterBgpDistanceList struct {
	AclStr          string `json:"acl-str,omitempty"`
	AdminDistance   int    `json:"admin-distance,omitempty"`
	ExtRoutesDist   int    `json:"ext-routes-dist,omitempty"`
	IntRoutesDist   int    `json:"int-routes-dist,omitempty"`
	LocalRoutesDist int    `json:"local-routes-dist,omitempty"`
	SrcPrefix       string `json:"src-prefix,omitempty"`
}

type RouterBgpNeighbor struct {
	Ethernet     []RouterBgpEthernetNeighborList  `json:"ethernet-neighbor-list,omitempty"`
	NeighborIpv4 []RouterBgpIpv4NeighborList      `json:"ipv4-neighbor-list,omitempty"`
	NeighborIpv6 []RouterBgpIpv6NeighborList      `json:"ipv6-neighbor-list,omitempty"`
	PeerGroup    []RouterBgpPeerGroupNeighborList `json:"peer-group-neighbor-list,omitempty"`
	Trunk        []RouterBgpTrunkNeighborList     `json:"trunk-neighbor-list,omitempty"`
	Ve           []RouterBgpVeNeighborList        `json:"ve-neighbor-list,omitempty"`
	//TempVal string `json:"temp-val,omitempty"`
	DistanceExt RouterBgpDistance `json:"distance,omitempty"`
}

type RouterBgpNetwork struct {
	NetworkIpv4Cidr        []RouterBgpIPCidrList    `json:"ip-cidr-list,omitempty"`
	NetworkSynchronization RouterBgpSynchronization `json:"synchronization,omitempty"`
}

type RouterBgpRedistributeBgp struct {
	Connected      RouterBgpConnectedCfg  `json:"connected-cfg,omitempty"`
	FloatingIP     RouterBgpFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	IPNat          RouterBgpIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	IPNatList      RouterBgpIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	Isis           RouterBgpIsisCfg       `json:"isis-cfg,omitempty"`
	Lw4O6          RouterBgpLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	NatMap         RouterBgpNatMapCfg     `json:"nat-map-cfg,omitempty"`
	Ospf           RouterBgpOspfCfg       `json:"ospf-cfg,omitempty"`
	Rip            RouterBgpRipCfg        `json:"rip-cfg,omitempty"`
	Static         RouterBgpStaticCfg     `json:"static-cfg,omitempty"`
	StaticNat      RouterBgpStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	UUID           string                 `json:"uuid,omitempty"`
	OnlyFlaggedCfg RouterBgpVip           `json:"vip,omitempty"`
}

type RouterBgpTimers struct {
	BgpHoldtime  int `json:"bgp-holdtime,omitempty"`
	BgpKeepalive int `json:"bgp-keepalive,omitempty"`
}

type RouterBgpIpv6 struct {
	AggregateAddress       []RouterBgpAggregateAddressList `json:"aggregate-address-list,omitempty"`
	AutoSummary            int                             `json:"auto-summary,omitempty"`
	Dampening              RouterBgpBgp1                   `json:"bgp,omitempty"`
	DistanceExt            RouterBgpDistance               `json:"distance,omitempty"`
	MaximumPathsValue      int                             `json:"maximum-paths-value,omitempty"`
	Ethernet               RouterBgpNeighbor6              `json:"neighbor,omitempty"`
	NetworkSynchronization RouterBgpNetworkIpv6            `json:"network,omitempty"`
	Originate              int                             `json:"originate,omitempty"`
	Connected              RouterBgpRedistributeIpv6       `json:"redistribute,omitempty"`
	Synchronization        int                             `json:"synchronization,omitempty"`
	UUID                   string                          `json:"uuid,omitempty"`
}

type RouterBgpBestpathCfg struct {
	CompareRouterid int `json:"compare-routerid,omitempty"`
	Ignore          int `json:"ignore,omitempty"`
	MissingAsWorst  int `json:"missing-as-worst,omitempty"`
	RemoveRecvMed   int `json:"remove-recv-med,omitempty"`
	RemoveSendMed   int `json:"remove-send-med,omitempty"`
}

type RouterBgpDampeningCfg struct {
	Dampening           int    `json:"dampening,omitempty"`
	DampeningHalfTime   int    `json:"dampening-half-time,omitempty"`
	DampeningMaxSupress int    `json:"dampening-max-supress,omitempty"`
	DampeningPenalty    int    `json:"dampening-penalty,omitempty"`
	DampeningReuse      int    `json:"dampening-reuse,omitempty"`
	DampeningSupress    int    `json:"dampening-supress,omitempty"`
	RouteMap            string `json:"route-map,omitempty"`
}

type RouterBgpEthernetNeighborList struct {
	Ethernet      int    `json:"ethernet,omitempty"`
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Unnumbered    int    `json:"unnumbered,omitempty"`
}

type RouterBgpIpv4NeighborList struct {
	Activate            int                              `json:"activate,omitempty"`
	AllowasIn           int                              `json:"allowas-in,omitempty"`
	AllowasInCount      int                              `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                              `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpDistributeLists       `json:"distribute-lists,omitempty"`
	GracefulRestart     int                              `json:"graceful-restart,omitempty"`
	Inbound             int                              `json:"inbound,omitempty"`
	MaximumPrefix       int                              `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                              `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv4        string                           `json:"neighbor-ipv4,omitempty"`
	NbrPrefixList       []RouterBgpNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                              `json:"next-hop-self,omitempty"`
	PeerGroupName       string                           `json:"peer-group-name,omitempty"`
	PrefixListDirection string                           `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                              `json:"remove-private-as,omitempty"`
	RouteMap            string                           `json:"route-map,omitempty"`
	SendCommunityVal    string                           `json:"send-community-val,omitempty"`
	UUID                string                           `json:"uuid,omitempty"`
	UnsuppressMap       string                           `json:"unsuppress-map,omitempty"`
	Weight              int                              `json:"weight,omitempty"`
}

type RouterBgpIpv6NeighborList struct {
	Activate            int                              `json:"activate,omitempty"`
	AllowasIn           int                              `json:"allowas-in,omitempty"`
	AllowasInCount      int                              `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                              `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpDistributeLists       `json:"distribute-lists,omitempty"`
	GracefulRestart     int                              `json:"graceful-restart,omitempty"`
	Inbound             int                              `json:"inbound,omitempty"`
	MaximumPrefix       int                              `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                              `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv6        string                           `json:"neighbor-ipv6,omitempty"`
	NbrPrefixList       []RouterBgpNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                              `json:"next-hop-self,omitempty"`
	PeerGroupName       string                           `json:"peer-group-name,omitempty"`
	PrefixListDirection string                           `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                              `json:"remove-private-as,omitempty"`
	RouteMap            string                           `json:"route-map,omitempty"`
	SendCommunityVal    string                           `json:"send-community-val,omitempty"`
	UUID                string                           `json:"uuid,omitempty"`
	UnsuppressMap       string                           `json:"unsuppress-map,omitempty"`
	Weight              int                              `json:"weight,omitempty"`
}

type RouterBgpPeerGroupNeighborList struct {
	Activate            int                              `json:"activate,omitempty"`
	AllowasIn           int                              `json:"allowas-in,omitempty"`
	AllowasInCount      int                              `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                              `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpDistributeLists       `json:"distribute-lists,omitempty"`
	Inbound             int                              `json:"inbound,omitempty"`
	MaximumPrefix       int                              `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                              `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NbrPrefixList       []RouterBgpNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                              `json:"next-hop-self,omitempty"`
	PeerGroup           string                           `json:"peer-group,omitempty"`
	PrefixListDirection string                           `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                              `json:"remove-private-as,omitempty"`
	RouteMap            string                           `json:"route-map,omitempty"`
	SendCommunityVal    string                           `json:"send-community-val,omitempty"`
	UUID                string                           `json:"uuid,omitempty"`
	UnsuppressMap       string                           `json:"unsuppress-map,omitempty"`
	Weight              int                              `json:"weight,omitempty"`
}

type RouterBgpTrunkNeighborList struct {
	PeerGroupName string `json:"peer-group-name,omitempty"`
	Trunk         int    `json:"trunk,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Unnumbered    int    `json:"unnumbered,omitempty"`
}

type RouterBgpVeNeighborList struct {
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Unnumbered    int    `json:"unnumbered,omitempty"`
	Ve            int    `json:"ve,omitempty"`
}

type RouterBgpIPCidrList struct {
	Backdoor        int    `json:"backdoor,omitempty"`
	CommValue       string `json:"comm-value,omitempty"`
	Description     string `json:"description,omitempty"`
	NetworkIpv4Cidr string `json:"network-ipv4-cidr,omitempty"`
	RouteMap        string `json:"route-map,omitempty"`
	UUID            string `json:"uuid,omitempty"`
}

type RouterBgpSynchronization struct {
	NetworkSynchronization int    `json:"network-synchronization,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
}

type RouterBgpConnectedCfg struct {
	Connected int    `json:"connected,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpFloatingIPCfg struct {
	FloatingIP int    `json:"floating-ip,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpIPNatCfg struct {
	IPNat    int    `json:"ip-nat,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpIPNatListCfg struct {
	IPNatList int    `json:"ip-nat-list,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpIsisCfg struct {
	Isis     int    `json:"isis,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpLw4O6Cfg struct {
	Lw4O6    int    `json:"lw4o6,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpNatMapCfg struct {
	NatMap   int    `json:"nat-map,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpOspfCfg struct {
	Ospf     int    `json:"ospf,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpRipCfg struct {
	Rip      int    `json:"rip,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpStaticCfg struct {
	RouteMap string `json:"route-map,omitempty"`
	Static   int    `json:"static,omitempty"`
} //Bgp

type RouterBgpStaticNatCfg struct {
	RouteMap  string `json:"route-map,omitempty"`
	StaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpVip struct {
	OnlyFlagged    RouterBgpOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	OnlyNotFlagged RouterBgpOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpBgp1 struct {
	Dampening               int    `json:"dampening,omitempty"`
	DampeningHalf           int    `json:"dampening-half,omitempty"`
	DampeningMaxSupress     int    `json:"dampening-max-supress,omitempty"`
	DampeningStartReuse     int    `json:"dampening-start-reuse,omitempty"`
	DampeningStartSupress   int    `json:"dampening-start-supress,omitempty"`
	DampeningUnreachability int    `json:"dampening-unreachability,omitempty"`
	RouteMap                string `json:"route-map,omitempty"`
}

type RouterBgpDistance struct {
	DistanceExt   int `json:"distance-ext,omitempty"`
	DistanceInt   int `json:"distance-int,omitempty"`
	DistanceLocal int `json:"distance-local,omitempty"`
}

type RouterBgpNeighbor6 struct {
	Ethernet     []RouterBgpEthernetNeighborIpv6List `json:"ethernet-neighbor-ipv6-list,omitempty"`
	NeighborIpv4 []RouterBgpIpv4NeighborList         `json:"ipv4-neighbor-list,omitempty"`
	NeighborIpv6 []RouterBgpIpv6NeighborList         `json:"ipv6-neighbor-list,omitempty"`
	PeerGroup    []RouterBgpPeerGroupNeighborList    `json:"peer-group-neighbor-list,omitempty"`
	Trunk        []RouterBgpTrunkNeighborIpv6List    `json:"trunk-neighbor-ipv6-list,omitempty"`
	Ve           []RouterBgpVeNeighborIpv6List       `json:"ve-neighbor-ipv6-list,omitempty"`
	//TempVal string `json:"temp-val,omitempty"`
	DistanceExt RouterBgpDistance `json:"distance,omitempty"`
}

type RouterBgpNetworkIpv6 struct {
	NetworkIpv6            []RouterBgpIpv6NetworkList `json:"ipv6-network-list,omitempty"`
	NetworkSynchronization RouterBgpSynchronization   `json:"synchronization,omitempty"`
}

type RouterBgpRedistributeIpv6 struct {
	Connected      RouterBgpConnectedCfg  `json:"connected-cfg,omitempty"`
	FloatingIP     RouterBgpFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	IPNat          RouterBgpIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	IPNatList      RouterBgpIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	Isis           RouterBgpIsisCfg       `json:"isis-cfg,omitempty"`
	Lw4O6          RouterBgpLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	Nat64          RouterBgpNat64Cfg      `json:"nat64-cfg,omitempty"`
	NatMap         RouterBgpNatMapCfg     `json:"nat-map-cfg,omitempty"`
	Ospf           RouterBgpOspfCfg       `json:"ospf-cfg,omitempty"`
	Rip            RouterBgpRipCfg        `json:"rip-cfg,omitempty"`
	Static         RouterBgpStaticCfg     `json:"static-cfg,omitempty"`
	StaticNat      RouterBgpStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	UUID           string                 `json:"uuid,omitempty"`
	OnlyFlaggedCfg RouterBgpVip           `json:"vip,omitempty"`
}

type RouterBgpDistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpNeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpNeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpOnlyFlaggedCfg struct {
	OnlyFlagged int    `json:"only-flagged,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpOnlyNotFlaggedCfg struct {
	OnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouteMap       string `json:"route-map,omitempty"`
}

type RouterBgpEthernetNeighborIpv6List struct {
	Ethernet      int    `json:"ethernet,omitempty"`
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

type RouterBgpTrunkNeighborIpv6List struct {
	PeerGroupName string `json:"peer-group-name,omitempty"`
	Trunk         int    `json:"trunk,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

type RouterBgpVeNeighborIpv6List struct {
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Ve            int    `json:"ve,omitempty"`
}

type RouterBgpIpv6NetworkList struct {
	Backdoor    int    `json:"backdoor,omitempty"`
	CommValue   string `json:"comm-value,omitempty"`
	Description string `json:"description,omitempty"`
	NetworkIpv6 string `json:"network-ipv6,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

type RouterBgpNat64Cfg struct {
	Nat64    int    `json:"nat64,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

func PostRouterBgp(id string, inst RouterBgp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgp(id string, name1 string, host string) (*RouterBgp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgp(id string, name1 string, inst RouterBgp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRouterBgp(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterBgp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/bgp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouterBgp", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
