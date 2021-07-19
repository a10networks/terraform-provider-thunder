package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6 struct {
	Bgp RouterBgpAddressFamilyIpv6Instance `json:"ipv6,omitempty"`
}

type RouterBgpAddressFamilyIpv6Instance struct {
	AggregateAddress       []RouterBgpAddressFamilyIpv6AggregateAddressList `json:"aggregate-address-list,omitempty"`
	AutoSummary            int                                              `json:"auto-summary,omitempty"`
	Dampening              RouterBgpAddressFamilyIpv6Bgp                    `json:"bgp,omitempty"`
	DistanceExt            RouterBgpAddressFamilyIpv6Ipv6Distance           `json:"distance,omitempty"`
	MaximumPathsValue      int                                              `json:"maximum-paths-value,omitempty"`
	PeerGroupNeighborList  RouterBgpAddressFamilyIpv6Ipv6Neighbor           `json:"neighbor,omitempty"`
	NetworkSynchronization RouterBgpAddressFamilyIpv6Network                `json:"network,omitempty"`
	Originate              int                                              `json:"originate,omitempty"`
	ConnectedCfg           RouterBgpAddressFamilyIpv6RedistributeIpv6       `json:"redistribute,omitempty"`
	Synchronization        int                                              `json:"synchronization,omitempty"`
	UUID                   string                                           `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6AggregateAddressList struct {
	AggregateAddress string `json:"aggregate-address,omitempty"`
	AsSet            int    `json:"as-set,omitempty"`
	SummaryOnly      int    `json:"summary-only,omitempty"`
}

type RouterBgpAddressFamilyIpv6Bgp struct {
	Dampening               int    `json:"dampening,omitempty"`
	DampeningHalf           int    `json:"dampening-half,omitempty"`
	DampeningMaxSupress     int    `json:"dampening-max-supress,omitempty"`
	DampeningStartReuse     int    `json:"dampening-start-reuse,omitempty"`
	DampeningStartSupress   int    `json:"dampening-start-supress,omitempty"`
	DampeningUnreachability int    `json:"dampening-unreachability,omitempty"`
	RouteMap                string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6Ipv6Distance struct {
	DistanceExt   int `json:"distance-ext,omitempty"`
	DistanceInt   int `json:"distance-int,omitempty"`
	DistanceLocal int `json:"distance-local,omitempty"`
}

type RouterBgpAddressFamilyIpv6Ipv6Neighbor struct {
	Ethernet               []RouterBgpAddressFamilyIpv6EthernetNeighborIpv6List `json:"ethernet-neighbor-ipv6-list,omitempty"`
	NeighborIpv4           []RouterBgpAddressFamilyIpv6Ipv4NeighborList         `json:"ipv4-neighbor-list,omitempty"`
	NeighborIpv6           []RouterBgpAddressFamilyIpv6Ipv6NeighborList         `json:"ipv6-neighbor-list,omitempty"`
	PeerGroup              []RouterBgpAddressFamilyIpv6PeerGroupNeighborList    `json:"peer-group-neighbor-list,omitempty"`
	Trunk                  []RouterBgpAddressFamilyIpv6TrunkNeighborIpv6List    `json:"trunk-neighbor-ipv6-list,omitempty"`
	Ve                     []RouterBgpAddressFamilyIpv6VeNeighborIpv6List       `json:"ve-neighbor-ipv6-list,omitempty"`
	NetworkSynchronization RouterBgpAddressFamilyIpv6Synchronization            `json:"synchronization,omitempty"`
}

type RouterBgpAddressFamilyIpv6Network struct {
	NetworkIpv6            []RouterBgpAddressFamilyIpv6NetworkList   `json:"ipv6-network-list,omitempty"`
	NetworkSynchronization RouterBgpAddressFamilyIpv6Synchronization `json:"synchronization,omitempty"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpv6 struct {
	Connected      RouterBgpAddressFamilyIpv6ConnectedCfg  `json:"connected-cfg,omitempty"`
	FloatingIP     RouterBgpAddressFamilyIpv6FloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	IPNat          RouterBgpAddressFamilyIpv6IPNatCfg      `json:"ip-nat-cfg,omitempty"`
	IPNatList      RouterBgpAddressFamilyIpv6IPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	Isis           RouterBgpAddressFamilyIpv6IsisCfg       `json:"isis-cfg,omitempty"`
	Lw4O6          RouterBgpAddressFamilyIpv6Lw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	Nat64          RouterBgpAddressFamilyIpv6Nat64Cfg      `json:"nat64-cfg,omitempty"`
	NatMap         RouterBgpAddressFamilyIpv6NatMapCfg     `json:"nat-map-cfg,omitempty"`
	Ospf           RouterBgpAddressFamilyIpv6OspfCfg       `json:"ospf-cfg,omitempty"`
	Rip            RouterBgpAddressFamilyIpv6RipCfg        `json:"rip-cfg,omitempty"`
	Static         RouterBgpAddressFamilyIpv6StaticCfg     `json:"static-cfg,omitempty"`
	StaticNat      RouterBgpAddressFamilyIpv6StaticNatCfg  `json:"static-nat-cfg,omitempty"`
	UUID           string                                  `json:"uuid,omitempty"`
	OnlyFlaggedCfg RouterBgpAddressFamilyIpv6Vip           `json:"vip,omitempty"`
}

type RouterBgpAddressFamilyIpv6EthernetNeighborIpv6List struct {
	Ethernet      int    `json:"ethernet,omitempty"`
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6Ipv4NeighborList struct {
	Activate            int                                               `json:"activate,omitempty"`
	AllowasIn           int                                               `json:"allowas-in,omitempty"`
	AllowasInCount      int                                               `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                                               `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpAddressFamilyIpv6DistributeLists       `json:"distribute-lists,omitempty"`
	GracefulRestart     int                                               `json:"graceful-restart,omitempty"`
	Inbound             int                                               `json:"inbound,omitempty"`
	MaximumPrefix       int                                               `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                                               `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpAddressFamilyIpv6NeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv4        string                                            `json:"neighbor-ipv4,omitempty"`
	NbrPrefixList       []RouterBgpAddressFamilyIpv6NeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpAddressFamilyIpv6NeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                                               `json:"next-hop-self,omitempty"`
	PeerGroupName       string                                            `json:"peer-group-name,omitempty"`
	PrefixListDirection string                                            `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                                               `json:"remove-private-as,omitempty"`
	RouteMap            string                                            `json:"route-map,omitempty"`
	SendCommunityVal    string                                            `json:"send-community-val,omitempty"`
	UUID                string                                            `json:"uuid,omitempty"`
	UnsuppressMap       string                                            `json:"unsuppress-map,omitempty"`
	Weight              int                                               `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6Ipv6NeighborList struct {
	Activate            int                                               `json:"activate,omitempty"`
	AllowasIn           int                                               `json:"allowas-in,omitempty"`
	AllowasInCount      int                                               `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                                               `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpAddressFamilyIpv6DistributeLists       `json:"distribute-lists,omitempty"`
	GracefulRestart     int                                               `json:"graceful-restart,omitempty"`
	Inbound             int                                               `json:"inbound,omitempty"`
	MaximumPrefix       int                                               `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                                               `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpAddressFamilyIpv6NeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv6        string                                            `json:"neighbor-ipv6,omitempty"`
	NbrPrefixList       []RouterBgpAddressFamilyIpv6NeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpAddressFamilyIpv6NeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                                               `json:"next-hop-self,omitempty"`
	PeerGroupName       string                                            `json:"peer-group-name,omitempty"`
	PrefixListDirection string                                            `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                                               `json:"remove-private-as,omitempty"`
	RouteMap            string                                            `json:"route-map,omitempty"`
	SendCommunityVal    string                                            `json:"send-community-val,omitempty"`
	UUID                string                                            `json:"uuid,omitempty"`
	UnsuppressMap       string                                            `json:"unsuppress-map,omitempty"`
	Weight              int                                               `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6PeerGroupNeighborList struct {
	Activate            int                                               `json:"activate,omitempty"`
	AllowasIn           int                                               `json:"allowas-in,omitempty"`
	AllowasInCount      int                                               `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                                               `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpAddressFamilyIpv6DistributeLists       `json:"distribute-lists,omitempty"`
	Inbound             int                                               `json:"inbound,omitempty"`
	MaximumPrefix       int                                               `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                                               `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpAddressFamilyIpv6NeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NbrPrefixList       []RouterBgpAddressFamilyIpv6NeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpAddressFamilyIpv6NeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                                               `json:"next-hop-self,omitempty"`
	PeerGroup           string                                            `json:"peer-group,omitempty"`
	PrefixListDirection string                                            `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                                               `json:"remove-private-as,omitempty"`
	RouteMap            string                                            `json:"route-map,omitempty"`
	SendCommunityVal    string                                            `json:"send-community-val,omitempty"`
	UUID                string                                            `json:"uuid,omitempty"`
	UnsuppressMap       string                                            `json:"unsuppress-map,omitempty"`
	Weight              int                                               `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6TrunkNeighborIpv6List struct {
	PeerGroupName string `json:"peer-group-name,omitempty"`
	Trunk         int    `json:"trunk,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6VeNeighborIpv6List struct {
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	Ve            int    `json:"ve,omitempty"`
}

type RouterBgpAddressFamilyIpv6NetworkList struct {
	Backdoor    int    `json:"backdoor,omitempty"`
	CommValue   string `json:"comm-value,omitempty"`
	Description string `json:"description,omitempty"`
	NetworkIpv6 string `json:"network-ipv6,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6Synchronization struct {
	NetworkSynchronization int    `json:"network-synchronization,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6ConnectedCfg struct {
	Connected int    `json:"connected,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6FloatingIPCfg struct {
	FloatingIP int    `json:"floating-ip,omitempty"`
	RouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6IPNatCfg struct {
	IPNat    int    `json:"ip-nat,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6IPNatListCfg struct {
	IPNatList int    `json:"ip-nat-list,omitempty"`
	RouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6IsisCfg struct {
	Isis     int    `json:"isis,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6Lw4O6Cfg struct {
	Lw4O6    int    `json:"lw4o6,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6Nat64Cfg struct {
	Nat64    int    `json:"nat64,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6NatMapCfg struct {
	NatMap   int    `json:"nat-map,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6OspfCfg struct {
	Ospf     int    `json:"ospf,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6RipCfg struct {
	Rip      int    `json:"rip,omitempty"`
	RouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6StaticCfg struct {
	RouteMap string `json:"route-map,omitempty"`
	Static   int    `json:"static,omitempty"`
}

type RouterBgpAddressFamilyIpv6StaticNatCfg struct {
	RouteMap  string `json:"route-map,omitempty"`
	StaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpAddressFamilyIpv6Vip struct {
	OnlyFlagged    RouterBgpAddressFamilyIpv6OnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	OnlyNotFlagged RouterBgpAddressFamilyIpv6OnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpAddressFamilyIpv6DistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6OnlyFlaggedCfg struct {
	OnlyFlagged int    `json:"only-flagged,omitempty"`
	RouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6OnlyNotFlaggedCfg struct {
	OnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouteMap       string `json:"route-map,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6(id string, name1 string, inst RouterBgpAddressFamilyIpv6, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostRouterBgpAddressFamilyIpv6", data)

		}
	}

}

func GetRouterBgpAddressFamilyIpv6(id string, name1 string, host string) (*RouterBgpAddressFamilyIpv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetRouterBgpAddressFamilyIpv6", data)
			return &m, nil
		}
	}

}
