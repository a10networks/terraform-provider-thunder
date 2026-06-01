package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type RouterBgpAddressFamilyIpv6 struct {
	Inst struct {
		AggregateAddressList []RouterBgpAddressFamilyIpv6AggregateAddressList `json:"aggregate-address-list"`

		AutoSummary int `json:"auto-summary"`

		Bgp RouterBgpAddressFamilyIpv6Bgp `json:"bgp"`

		Distance RouterBgpAddressFamilyIpv6Distance `json:"distance"`

		MaximumPathsValue int `json:"maximum-paths-value" dval:"1"`

		Neighbor RouterBgpAddressFamilyIpv6Neighbor1236 `json:"neighbor"`

		Network RouterBgpAddressFamilyIpv6Network1237 `json:"network"`

		Originate int `json:"originate"`

		PreferGlobal int `json:"prefer-global"`

		Redistribute RouterBgpAddressFamilyIpv6Redistribute1241 `json:"redistribute"`

		Synchronization int `json:"synchronization"`

		Uuid string `json:"uuid"`

		AsNumber string
	} `json:"ipv6"`
}

type RouterBgpAddressFamilyIpv6AggregateAddressList struct {
	AggregateAddress string `json:"aggregate-address"`
	AsSet            int    `json:"as-set"`
	SummaryOnly      int    `json:"summary-only"`
}

type RouterBgpAddressFamilyIpv6Bgp struct {
	Dampening               int    `json:"dampening"`
	DampeningHalf           int    `json:"dampening-half"`
	DampeningStartReuse     int    `json:"dampening-start-reuse"`
	DampeningStartSupress   int    `json:"dampening-start-supress"`
	DampeningMaxSupress     int    `json:"dampening-max-supress"`
	DampeningUnreachability int    `json:"dampening-unreachability"`
	RouteMap                string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6Distance struct {
	DistanceExt   int `json:"distance-ext"`
	DistanceInt   int `json:"distance-int"`
	DistanceLocal int `json:"distance-local"`
}

type RouterBgpAddressFamilyIpv6Neighbor1236 struct {
	PeerGroupNeighborList    []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList    `json:"peer-group-neighbor-list"`
	Ipv4NeighborList         []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList         `json:"ipv4-neighbor-list"`
	Ipv6NeighborList         []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList         `json:"ipv6-neighbor-list"`
	EthernetNeighborIpv6List []RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List `json:"ethernet-neighbor-ipv6-list"`
	VeNeighborIpv6List       []RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List       `json:"ve-neighbor-ipv6-list"`
	TrunkNeighborIpv6List    []RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List    `json:"trunk-neighbor-ipv6-list"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList struct {
	PeerGroup             string                                                                         `json:"peer-group"`
	Activate              int                                                                            `json:"activate"`
	AllowasIn             int                                                                            `json:"allowas-in"`
	AllowasInCount        int                                                                            `json:"allowas-in-count" dval:"3"`
	MaximumPrefix         int                                                                            `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                            `json:"maximum-prefix-thres"`
	NextHopSelf           int                                                                            `json:"next-hop-self"`
	RemovePrivateAs       int                                                                            `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
	Inbound               int                                                                            `json:"inbound"`
	Weight                int                                                                            `json:"weight"`
	Uuid                  string                                                                         `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList struct {
	NeighborIpv4          string                                                                    `json:"neighbor-ipv4"`
	PeerGroupName         string                                                                    `json:"peer-group-name"`
	Activate              int                                                                       `json:"activate"`
	AllowasIn             int                                                                       `json:"allowas-in"`
	AllowasInCount        int                                                                       `json:"allowas-in-count" dval:"3"`
	PrefixListDirection   string                                                                    `json:"prefix-list-direction"`
	GracefulRestart       int                                                                       `json:"graceful-restart"`
	DefaultOriginate      int                                                                       `json:"default-originate"`
	RouteMap              string                                                                    `json:"route-map"`
	DistributeLists       []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists       `json:"distribute-lists"`
	NeighborFilterLists   []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists   `json:"neighbor-filter-lists"`
	MaximumPrefix         int                                                                       `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                       `json:"maximum-prefix-thres"`
	RestartMin            int                                                                       `json:"restart-min"`
	NextHopSelf           int                                                                       `json:"next-hop-self"`
	NeighborPrefixLists   []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists"`
	RemovePrivateAs       int                                                                       `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                    `json:"send-community-val" dval:"both"`
	Inbound               int                                                                       `json:"inbound"`
	UnsuppressMap         string                                                                    `json:"unsuppress-map"`
	Weight                int                                                                       `json:"weight"`
	Uuid                  string                                                                    `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList struct {
	NeighborIpv6          string                                                                    `json:"neighbor-ipv6"`
	PeerGroupName         string                                                                    `json:"peer-group-name"`
	Activate              int                                                                       `json:"activate"`
	AllowasIn             int                                                                       `json:"allowas-in"`
	AllowasInCount        int                                                                       `json:"allowas-in-count" dval:"3"`
	PrefixListDirection   string                                                                    `json:"prefix-list-direction"`
	GracefulRestart       int                                                                       `json:"graceful-restart"`
	DefaultOriginate      int                                                                       `json:"default-originate"`
	RouteMap              string                                                                    `json:"route-map"`
	DistributeLists       []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists       `json:"distribute-lists"`
	NeighborFilterLists   []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists   `json:"neighbor-filter-lists"`
	MaximumPrefix         int                                                                       `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                       `json:"maximum-prefix-thres"`
	RestartMin            int                                                                       `json:"restart-min"`
	NextHopSelf           int                                                                       `json:"next-hop-self"`
	NeighborPrefixLists   []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists"`
	RemovePrivateAs       int                                                                       `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                    `json:"send-community-val" dval:"both"`
	Inbound               int                                                                       `json:"inbound"`
	UnsuppressMap         string                                                                    `json:"unsuppress-map"`
	Weight                int                                                                       `json:"weight"`
	Uuid                  string                                                                    `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List struct {
	Ethernet      int    `json:"ethernet"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List struct {
	Ve            int    `json:"ve"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List struct {
	Trunk         int    `json:"trunk"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6Network1237 struct {
	Synchronization RouterBgpAddressFamilyIpv6NetworkSynchronization1238 `json:"synchronization"`
	Monitor         RouterBgpAddressFamilyIpv6NetworkMonitor1239         `json:"monitor"`
	Ipv6NetworkList []RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList   `json:"ipv6-network-list"`
}

type RouterBgpAddressFamilyIpv6NetworkSynchronization1238 struct {
	NetworkSynchronization int    `json:"network-synchronization"`
	Uuid                   string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NetworkMonitor1239 struct {
	Default RouterBgpAddressFamilyIpv6NetworkMonitorDefault1240 `json:"default"`
}

type RouterBgpAddressFamilyIpv6NetworkMonitorDefault1240 struct {
	NetworkMonitorDefault int    `json:"network-monitor-default"`
	Uuid                  string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList struct {
	NetworkIpv6 string `json:"network-ipv6"`
	RouteMap    string `json:"route-map"`
	Backdoor    int    `json:"backdoor"`
	Description string `json:"description"`
	CommValue   string `json:"comm-value"`
	LcommValue  string `json:"lcomm-value"`
	Uuid        string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6Redistribute1241 struct {
	ConnectedCfg  RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1242  `json:"connected-cfg"`
	FloatingIpCfg RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1243 `json:"floating-ip-cfg"`
	Nat64Cfg      RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1244      `json:"nat64-cfg"`
	NatMapCfg     RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1245     `json:"nat-map-cfg"`
	Lw4o6Cfg      RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1246      `json:"lw4o6-cfg"`
	StaticNatCfg  RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1247  `json:"static-nat-cfg"`
	IpNatCfg      RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1248      `json:"ip-nat-cfg"`
	IpNatListCfg  RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1249  `json:"ip-nat-list-cfg"`
	IsisCfg       RouterBgpAddressFamilyIpv6RedistributeIsisCfg1250       `json:"isis-cfg"`
	OspfCfg       RouterBgpAddressFamilyIpv6RedistributeOspfCfg1251       `json:"ospf-cfg"`
	RipCfg        RouterBgpAddressFamilyIpv6RedistributeRipCfg1252        `json:"rip-cfg"`
	StaticCfg     RouterBgpAddressFamilyIpv6RedistributeStaticCfg1253     `json:"static-cfg"`
	PublicIpCfg   RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1254   `json:"public-ip-cfg"`
	Vip           RouterBgpAddressFamilyIpv6RedistributeVip1255           `json:"vip"`
	Uuid          string                                                  `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1242 struct {
	Connected int    `json:"connected"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1243 struct {
	FloatingIp int    `json:"floating-ip"`
	RouteMap   string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1244 struct {
	Nat64    int    `json:"nat64"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1245 struct {
	NatMap   int    `json:"nat-map"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1246 struct {
	Lw4o6    int    `json:"lw4o6"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1247 struct {
	StaticNat int    `json:"static-nat"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1248 struct {
	IpNat    int    `json:"ip-nat"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1249 struct {
	IpNatList int    `json:"ip-nat-list"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIsisCfg1250 struct {
	Isis     int    `json:"isis"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeOspfCfg1251 struct {
	Ospf     int    `json:"ospf"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeRipCfg1252 struct {
	Rip      int    `json:"rip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticCfg1253 struct {
	Static   int    `json:"static"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1254 struct {
	PublicIp int    `json:"public-ip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeVip1255 struct {
	OnlyFlaggedCfg    RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1256    `json:"only-flagged-cfg"`
	OnlyNotFlaggedCfg RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1257 `json:"only-not-flagged-cfg"`
}

type RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1256 struct {
	OnlyFlagged int    `json:"only-flagged"`
	RouteMap    string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1257 struct {
	OnlyNotFlagged int    `json:"only-not-flagged"`
	RouteMap       string `json:"route-map"`
}

func (p *RouterBgpAddressFamilyIpv6) GetId() string {
	return "1"
}

func (p *RouterBgpAddressFamilyIpv6) getPath() string {
	return "router/bgp/" + p.Inst.AsNumber + "/address-family/ipv6"
}

func (p *RouterBgpAddressFamilyIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpAddressFamilyIpv6::Post")
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

func (p *RouterBgpAddressFamilyIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpAddressFamilyIpv6::Get")
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
func (p *RouterBgpAddressFamilyIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpAddressFamilyIpv6::Put")
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

func (p *RouterBgpAddressFamilyIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpAddressFamilyIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
