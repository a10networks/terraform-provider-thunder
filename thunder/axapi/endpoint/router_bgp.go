package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type RouterBgp struct {
	Inst struct {
		AddressFamily RouterBgpAddressFamily1259 `json:"address-family"`

		AggregateAddressList []RouterBgpAggregateAddressList `json:"aggregate-address-list"`

		AsNumber string `json:"as-number"`

		AutoSummary int `json:"auto-summary"`

		Bgp RouterBgpBgp `json:"bgp"`

		DistanceList []RouterBgpDistanceList `json:"distance-list"`

		MaximumPathsValue int `json:"maximum-paths-value" dval:"1"`

		Neighbor RouterBgpNeighbor1314 `json:"neighbor"`

		Network RouterBgpNetwork1315 `json:"network"`

		Originate int `json:"originate"`

		Redistribute RouterBgpRedistribute1319 `json:"redistribute"`

		Synchronization int `json:"synchronization"`

		Timers RouterBgpTimers `json:"timers"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"bgp"`
}

type RouterBgpAddressFamily1259 struct {
	Ipv6         RouterBgpAddressFamilyIpv61260         `json:"ipv6"`
	Ipv4Flowspec RouterBgpAddressFamilyIpv4Flowspec1302 `json:"ipv4-flowspec"`
	Ipv6Flowspec RouterBgpAddressFamilyIpv6Flowspec1308 `json:"ipv6-flowspec"`
}

type RouterBgpAddressFamilyIpv61260 struct {
	Bgp                  RouterBgpAddressFamilyIpv6Bgp1261                    `json:"bgp"`
	Distance             RouterBgpAddressFamilyIpv6Distance1262               `json:"distance"`
	MaximumPathsValue    int                                                  `json:"maximum-paths-value" dval:"1"`
	Originate            int                                                  `json:"originate"`
	AggregateAddressList []RouterBgpAddressFamilyIpv6AggregateAddressList1263 `json:"aggregate-address-list"`
	AutoSummary          int                                                  `json:"auto-summary"`
	Synchronization      int                                                  `json:"synchronization"`
	Uuid                 string                                               `json:"uuid"`
	Network              RouterBgpAddressFamilyIpv6Network1264                `json:"network"`
	Neighbor             RouterBgpAddressFamilyIpv6Neighbor1269               `json:"neighbor"`
	Redistribute         RouterBgpAddressFamilyIpv6Redistribute1285           `json:"redistribute"`
}

type RouterBgpAddressFamilyIpv6Bgp1261 struct {
	Dampening               int    `json:"dampening"`
	DampeningHalf           int    `json:"dampening-half"`
	DampeningStartReuse     int    `json:"dampening-start-reuse"`
	DampeningStartSupress   int    `json:"dampening-start-supress"`
	DampeningMaxSupress     int    `json:"dampening-max-supress"`
	DampeningUnreachability int    `json:"dampening-unreachability"`
	RouteMap                string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6Distance1262 struct {
	DistanceExt   int `json:"distance-ext"`
	DistanceInt   int `json:"distance-int"`
	DistanceLocal int `json:"distance-local"`
}

type RouterBgpAddressFamilyIpv6AggregateAddressList1263 struct {
	AggregateAddress string `json:"aggregate-address"`
	AsSet            int    `json:"as-set"`
	SummaryOnly      int    `json:"summary-only"`
}

type RouterBgpAddressFamilyIpv6Network1264 struct {
	Synchronization RouterBgpAddressFamilyIpv6NetworkSynchronization1265   `json:"synchronization"`
	Monitor         RouterBgpAddressFamilyIpv6NetworkMonitor1266           `json:"monitor"`
	Ipv6NetworkList []RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1268 `json:"ipv6-network-list"`
}

type RouterBgpAddressFamilyIpv6NetworkSynchronization1265 struct {
	NetworkSynchronization int    `json:"network-synchronization"`
	Uuid                   string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NetworkMonitor1266 struct {
	Default RouterBgpAddressFamilyIpv6NetworkMonitorDefault1267 `json:"default"`
}

type RouterBgpAddressFamilyIpv6NetworkMonitorDefault1267 struct {
	NetworkMonitorDefault int    `json:"network-monitor-default"`
	Uuid                  string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1268 struct {
	NetworkIpv6 string `json:"network-ipv6"`
	RouteMap    string `json:"route-map"`
	Backdoor    int    `json:"backdoor"`
	Description string `json:"description"`
	CommValue   string `json:"comm-value"`
	LcommValue  string `json:"lcomm-value"`
	Uuid        string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6Neighbor1269 struct {
	PeerGroupNeighborList    []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1270    `json:"peer-group-neighbor-list"`
	Ipv4NeighborList         []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1272         `json:"ipv4-neighbor-list"`
	Ipv6NeighborList         []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1277         `json:"ipv6-neighbor-list"`
	EthernetNeighborIpv6List []RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1282 `json:"ethernet-neighbor-ipv6-list"`
	VeNeighborIpv6List       []RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1283       `json:"ve-neighbor-ipv6-list"`
	TrunkNeighborIpv6List    []RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1284    `json:"trunk-neighbor-ipv6-list"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1270 struct {
	PeerGroup             string                                                                             `json:"peer-group"`
	Activate              int                                                                                `json:"activate"`
	AllowasIn             int                                                                                `json:"allowas-in"`
	AllowasInCount        int                                                                                `json:"allowas-in-count" dval:"3"`
	MaximumPrefix         int                                                                                `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                                `json:"maximum-prefix-thres"`
	NextHopSelf           int                                                                                `json:"next-hop-self"`
	RemovePrivateAs       int                                                                                `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1271 `json:"neighbor-route-map-lists"`
	Inbound               int                                                                                `json:"inbound"`
	Weight                int                                                                                `json:"weight"`
	Uuid                  string                                                                             `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1271 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1272 struct {
	NeighborIpv4          string                                                                        `json:"neighbor-ipv4"`
	PeerGroupName         string                                                                        `json:"peer-group-name"`
	Activate              int                                                                           `json:"activate"`
	AllowasIn             int                                                                           `json:"allowas-in"`
	AllowasInCount        int                                                                           `json:"allowas-in-count" dval:"3"`
	PrefixListDirection   string                                                                        `json:"prefix-list-direction"`
	GracefulRestart       int                                                                           `json:"graceful-restart"`
	DefaultOriginate      int                                                                           `json:"default-originate"`
	RouteMap              string                                                                        `json:"route-map"`
	DistributeLists       []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1273       `json:"distribute-lists"`
	NeighborFilterLists   []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1274   `json:"neighbor-filter-lists"`
	MaximumPrefix         int                                                                           `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                           `json:"maximum-prefix-thres"`
	RestartMin            int                                                                           `json:"restart-min"`
	NextHopSelf           int                                                                           `json:"next-hop-self"`
	NeighborPrefixLists   []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1275   `json:"neighbor-prefix-lists"`
	RemovePrivateAs       int                                                                           `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1276 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                        `json:"send-community-val" dval:"both"`
	Inbound               int                                                                           `json:"inbound"`
	UnsuppressMap         string                                                                        `json:"unsuppress-map"`
	Weight                int                                                                           `json:"weight"`
	Uuid                  string                                                                        `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1273 struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1274 struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1275 struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1276 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1277 struct {
	NeighborIpv6          string                                                                        `json:"neighbor-ipv6"`
	PeerGroupName         string                                                                        `json:"peer-group-name"`
	Activate              int                                                                           `json:"activate"`
	AllowasIn             int                                                                           `json:"allowas-in"`
	AllowasInCount        int                                                                           `json:"allowas-in-count" dval:"3"`
	PrefixListDirection   string                                                                        `json:"prefix-list-direction"`
	GracefulRestart       int                                                                           `json:"graceful-restart"`
	DefaultOriginate      int                                                                           `json:"default-originate"`
	RouteMap              string                                                                        `json:"route-map"`
	DistributeLists       []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1278       `json:"distribute-lists"`
	NeighborFilterLists   []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1279   `json:"neighbor-filter-lists"`
	MaximumPrefix         int                                                                           `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                           `json:"maximum-prefix-thres"`
	RestartMin            int                                                                           `json:"restart-min"`
	NextHopSelf           int                                                                           `json:"next-hop-self"`
	NeighborPrefixLists   []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1280   `json:"neighbor-prefix-lists"`
	RemovePrivateAs       int                                                                           `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1281 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                        `json:"send-community-val" dval:"both"`
	Inbound               int                                                                           `json:"inbound"`
	UnsuppressMap         string                                                                        `json:"unsuppress-map"`
	Weight                int                                                                           `json:"weight"`
	Uuid                  string                                                                        `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1278 struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1279 struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1280 struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1281 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1282 struct {
	Ethernet      int    `json:"ethernet"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1283 struct {
	Ve            int    `json:"ve"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1284 struct {
	Trunk         int    `json:"trunk"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6Redistribute1285 struct {
	ConnectedCfg  RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1286  `json:"connected-cfg"`
	FloatingIpCfg RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1287 `json:"floating-ip-cfg"`
	Nat64Cfg      RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1288      `json:"nat64-cfg"`
	NatMapCfg     RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1289     `json:"nat-map-cfg"`
	Lw4o6Cfg      RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1290      `json:"lw4o6-cfg"`
	StaticNatCfg  RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1291  `json:"static-nat-cfg"`
	IpNatCfg      RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1292      `json:"ip-nat-cfg"`
	IpNatListCfg  RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1293  `json:"ip-nat-list-cfg"`
	IsisCfg       RouterBgpAddressFamilyIpv6RedistributeIsisCfg1294       `json:"isis-cfg"`
	OspfCfg       RouterBgpAddressFamilyIpv6RedistributeOspfCfg1295       `json:"ospf-cfg"`
	RipCfg        RouterBgpAddressFamilyIpv6RedistributeRipCfg1296        `json:"rip-cfg"`
	StaticCfg     RouterBgpAddressFamilyIpv6RedistributeStaticCfg1297     `json:"static-cfg"`
	PublicIpCfg   RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1298   `json:"public-ip-cfg"`
	Vip           RouterBgpAddressFamilyIpv6RedistributeVip1299           `json:"vip"`
	Uuid          string                                                  `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1286 struct {
	Connected int    `json:"connected"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1287 struct {
	FloatingIp int    `json:"floating-ip"`
	RouteMap   string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1288 struct {
	Nat64    int    `json:"nat64"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1289 struct {
	NatMap   int    `json:"nat-map"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1290 struct {
	Lw4o6    int    `json:"lw4o6"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1291 struct {
	StaticNat int    `json:"static-nat"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1292 struct {
	IpNat    int    `json:"ip-nat"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1293 struct {
	IpNatList int    `json:"ip-nat-list"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIsisCfg1294 struct {
	Isis     int    `json:"isis"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeOspfCfg1295 struct {
	Ospf     int    `json:"ospf"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeRipCfg1296 struct {
	Rip      int    `json:"rip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticCfg1297 struct {
	Static   int    `json:"static"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1298 struct {
	PublicIp int    `json:"public-ip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeVip1299 struct {
	OnlyFlaggedCfg    RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1300    `json:"only-flagged-cfg"`
	OnlyNotFlaggedCfg RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1301 `json:"only-not-flagged-cfg"`
}

type RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1300 struct {
	OnlyFlagged int    `json:"only-flagged"`
	RouteMap    string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1301 struct {
	OnlyNotFlagged int    `json:"only-not-flagged"`
	RouteMap       string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv4Flowspec1302 struct {
	Uuid     string                                         `json:"uuid"`
	Neighbor RouterBgpAddressFamilyIpv4FlowspecNeighbor1303 `json:"neighbor"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighbor1303 struct {
	Ipv4NeighborList []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1304 `json:"ipv4-neighbor-list"`
	Ipv6NeighborList []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1306 `json:"ipv6-neighbor-list"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1304 struct {
	NeighborIpv4          string                                                                                `json:"neighbor-ipv4"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1305 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1305 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1306 struct {
	NeighborIpv6          string                                                                                `json:"neighbor-ipv6"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1307 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1307 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6Flowspec1308 struct {
	Uuid     string                                         `json:"uuid"`
	Neighbor RouterBgpAddressFamilyIpv6FlowspecNeighbor1309 `json:"neighbor"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighbor1309 struct {
	Ipv4NeighborList []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1310 `json:"ipv4-neighbor-list"`
	Ipv6NeighborList []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1312 `json:"ipv6-neighbor-list"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1310 struct {
	NeighborIpv4          string                                                                                `json:"neighbor-ipv4"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1311 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1311 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1312 struct {
	NeighborIpv6          string                                                                                `json:"neighbor-ipv6"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1313 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1313 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAggregateAddressList struct {
	AggregateAddress string `json:"aggregate-address"`
	AsSet            int    `json:"as-set"`
	SummaryOnly      int    `json:"summary-only"`
}

type RouterBgpBgp struct {
	AlwaysCompareMed     int                      `json:"always-compare-med"`
	BestpathCfg          RouterBgpBgpBestpathCfg  `json:"bestpath-cfg"`
	DampeningCfg         RouterBgpBgpDampeningCfg `json:"dampening-cfg"`
	LocalPreferenceValue int                      `json:"local-preference-value" dval:"100"`
	DeterministicMed     int                      `json:"deterministic-med"`
	EnforceFirstAs       int                      `json:"enforce-first-as"`
	FastExternalFailover int                      `json:"fast-external-failover" dval:"1"`
	LogNeighborChanges   int                      `json:"log-neighbor-changes"`
	NexthopTriggerCount  int                      `json:"nexthop-trigger-count"`
	RouterId             string                   `json:"router-id"`
	ScanTime             int                      `json:"scan-time" dval:"60"`
	GracefulRestart      int                      `json:"graceful-restart"`
	BgpRestartTime       int                      `json:"bgp-restart-time" dval:"90"`
	BgpStalepathTime     int                      `json:"bgp-stalepath-time" dval:"360"`
}

type RouterBgpBgpBestpathCfg struct {
	Ignore          int `json:"ignore"`
	CompareRouterid int `json:"compare-routerid"`
	RemoveRecvMed   int `json:"remove-recv-med"`
	RemoveSendMed   int `json:"remove-send-med"`
	MissingAsWorst  int `json:"missing-as-worst"`
}

type RouterBgpBgpDampeningCfg struct {
	Dampening           int    `json:"dampening"`
	DampeningHalfTime   int    `json:"dampening-half-time"`
	DampeningReuse      int    `json:"dampening-reuse"`
	DampeningSupress    int    `json:"dampening-supress"`
	DampeningMaxSupress int    `json:"dampening-max-supress"`
	DampeningPenalty    int    `json:"dampening-penalty"`
	RouteMap            string `json:"route-map"`
}

type RouterBgpDistanceList struct {
	AdminDistance   int    `json:"admin-distance"`
	SrcPrefix       string `json:"src-prefix"`
	AclStr          string `json:"acl-str"`
	ExtRoutesDist   int    `json:"ext-routes-dist"`
	IntRoutesDist   int    `json:"int-routes-dist"`
	LocalRoutesDist int    `json:"local-routes-dist"`
}

type RouterBgpNeighbor1314 struct {
	PeerGroupNeighborList []RouterBgpNeighborPeerGroupNeighborList `json:"peer-group-neighbor-list"`
	Ipv4NeighborList      []RouterBgpNeighborIpv4NeighborList      `json:"ipv4-neighbor-list"`
	Ipv6NeighborList      []RouterBgpNeighborIpv6NeighborList      `json:"ipv6-neighbor-list"`
	EthernetNeighborList  []RouterBgpNeighborEthernetNeighborList  `json:"ethernet-neighbor-list"`
	VeNeighborList        []RouterBgpNeighborVeNeighborList        `json:"ve-neighbor-list"`
	TrunkNeighborList     []RouterBgpNeighborTrunkNeighborList     `json:"trunk-neighbor-list"`
}

type RouterBgpNeighborPeerGroupNeighborList struct {
	PeerGroup               string                                                        `json:"peer-group"`
	PeerGroupKey            int                                                           `json:"peer-group-key"`
	PeerGroupRemoteAs       string                                                        `json:"peer-group-remote-as"`
	Activate                int                                                           `json:"activate" dval:"1"`
	AdvertisementInterval   int                                                           `json:"advertisement-interval"`
	AllowasIn               int                                                           `json:"allowas-in"`
	AllowasInCount          int                                                           `json:"allowas-in-count" dval:"3"`
	AsOriginationInterval   int                                                           `json:"as-origination-interval"`
	Dynamic                 int                                                           `json:"dynamic"`
	RouteRefresh            int                                                           `json:"route-refresh" dval:"1"`
	ExtendedNexthop         int                                                           `json:"extended-nexthop"`
	CollideEstablished      int                                                           `json:"collide-established"`
	DefaultOriginate        int                                                           `json:"default-originate"`
	RouteMap                string                                                        `json:"route-map"`
	Description             string                                                        `json:"description"`
	DontCapabilityNegotiate int                                                           `json:"dont-capability-negotiate"`
	EbgpMultihop            int                                                           `json:"ebgp-multihop"`
	EbgpMultihopHopCount    int                                                           `json:"ebgp-multihop-hop-count"`
	EnforceMultihop         int                                                           `json:"enforce-multihop"`
	Bfd                     int                                                           `json:"bfd"`
	Multihop                int                                                           `json:"multihop"`
	MaximumPrefix           int                                                           `json:"maximum-prefix"`
	MaximumPrefixThres      int                                                           `json:"maximum-prefix-thres"`
	OverrideCapability      int                                                           `json:"override-capability"`
	PassValue               string                                                        `json:"pass-value"`
	PassEncrypted           string                                                        `json:"pass-encrypted"`
	Passive                 int                                                           `json:"passive"`
	RemovePrivateAs         int                                                           `json:"remove-private-as"`
	NeighborRouteMapLists   []RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
	Inbound                 int                                                           `json:"inbound"`
	Shutdown                int                                                           `json:"shutdown"`
	StrictCapabilityMatch   int                                                           `json:"strict-capability-match"`
	TimersKeepalive         int                                                           `json:"timers-keepalive" dval:"30"`
	TimersHoldtime          int                                                           `json:"timers-holdtime" dval:"90"`
	Connect                 int                                                           `json:"connect"`
	UpdateSourceIp          string                                                        `json:"update-source-ip"`
	UpdateSourceIpv6        string                                                        `json:"update-source-ipv6"`
	Ethernet                int                                                           `json:"ethernet"`
	Loopback                int                                                           `json:"loopback"`
	Ve                      int                                                           `json:"ve"`
	Trunk                   int                                                           `json:"trunk"`
	Lif                     string                                                        `json:"lif"`
	Tunnel                  int                                                           `json:"tunnel"`
	Weight                  int                                                           `json:"weight"`
	Uuid                    string                                                        `json:"uuid"`
}

type RouterBgpNeighborPeerGroupNeighborListNeighborRouteMapLists struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpNeighborIpv4NeighborList struct {
	NeighborIpv4             string                                                   `json:"neighbor-ipv4"`
	NbrRemoteAs              string                                                   `json:"nbr-remote-as"`
	PeerGroupName            string                                                   `json:"peer-group-name"`
	Activate                 int                                                      `json:"activate" dval:"1"`
	AdvertisementInterval    int                                                      `json:"advertisement-interval"`
	AllowasIn                int                                                      `json:"allowas-in"`
	AllowasInCount           int                                                      `json:"allowas-in-count" dval:"3"`
	AsOriginationInterval    int                                                      `json:"as-origination-interval"`
	Dynamic                  int                                                      `json:"dynamic"`
	PrefixListDirection      string                                                   `json:"prefix-list-direction"`
	RouteRefresh             int                                                      `json:"route-refresh" dval:"1"`
	GracefulRestart          int                                                      `json:"graceful-restart"`
	CollideEstablished       int                                                      `json:"collide-established"`
	DefaultOriginate         int                                                      `json:"default-originate"`
	RouteMap                 string                                                   `json:"route-map"`
	Description              string                                                   `json:"description"`
	DisallowInfiniteHoldtime int                                                      `json:"disallow-infinite-holdtime"`
	DistributeLists          []RouterBgpNeighborIpv4NeighborListDistributeLists       `json:"distribute-lists"`
	AcosApplicationOnly      int                                                      `json:"acos-application-only"`
	Telemetry                int                                                      `json:"telemetry"`
	DontCapabilityNegotiate  int                                                      `json:"dont-capability-negotiate"`
	EbgpMultihop             int                                                      `json:"ebgp-multihop"`
	EbgpMultihopHopCount     int                                                      `json:"ebgp-multihop-hop-count"`
	EnforceMultihop          int                                                      `json:"enforce-multihop"`
	Bfd                      int                                                      `json:"bfd"`
	Multihop                 int                                                      `json:"multihop"`
	KeyId                    int                                                      `json:"key-id"`
	KeyType                  string                                                   `json:"key-type"`
	BfdValue                 string                                                   `json:"bfd-value"`
	BfdEncrypted             string                                                   `json:"bfd-encrypted"`
	NeighborFilterLists      []RouterBgpNeighborIpv4NeighborListNeighborFilterLists   `json:"neighbor-filter-lists"`
	MaximumPrefix            int                                                      `json:"maximum-prefix"`
	MaximumPrefixThres       int                                                      `json:"maximum-prefix-thres"`
	RestartMin               int                                                      `json:"restart-min"`
	NextHopSelf              int                                                      `json:"next-hop-self"`
	OverrideCapability       int                                                      `json:"override-capability"`
	PassValue                string                                                   `json:"pass-value"`
	PassEncrypted            string                                                   `json:"pass-encrypted"`
	Passive                  int                                                      `json:"passive"`
	NeighborPrefixLists      []RouterBgpNeighborIpv4NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists"`
	RemovePrivateAs          int                                                      `json:"remove-private-as"`
	NeighborRouteMapLists    []RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
	SendCommunityVal         string                                                   `json:"send-community-val" dval:"both"`
	Inbound                  int                                                      `json:"inbound"`
	Shutdown                 int                                                      `json:"shutdown"`
	StrictCapabilityMatch    int                                                      `json:"strict-capability-match"`
	TimersKeepalive          int                                                      `json:"timers-keepalive" dval:"30"`
	TimersHoldtime           int                                                      `json:"timers-holdtime" dval:"90"`
	Connect                  int                                                      `json:"connect"`
	UnsuppressMap            string                                                   `json:"unsuppress-map"`
	UpdateSourceIp           string                                                   `json:"update-source-ip"`
	UpdateSourceIpv6         string                                                   `json:"update-source-ipv6"`
	Ethernet                 int                                                      `json:"ethernet"`
	Loopback                 int                                                      `json:"loopback"`
	Ve                       int                                                      `json:"ve"`
	Trunk                    int                                                      `json:"trunk"`
	Lif                      string                                                   `json:"lif"`
	Tunnel                   int                                                      `json:"tunnel"`
	Weight                   int                                                      `json:"weight"`
	Uuid                     string                                                   `json:"uuid"`
}

type RouterBgpNeighborIpv4NeighborListDistributeLists struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpNeighborIpv4NeighborListNeighborFilterLists struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpNeighborIpv4NeighborListNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpNeighborIpv4NeighborListNeighborRouteMapLists struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpNeighborIpv6NeighborList struct {
	NeighborIpv6             string                                                   `json:"neighbor-ipv6"`
	NbrRemoteAs              string                                                   `json:"nbr-remote-as"`
	PeerGroupName            string                                                   `json:"peer-group-name"`
	Activate                 int                                                      `json:"activate" dval:"1"`
	AdvertisementInterval    int                                                      `json:"advertisement-interval"`
	AllowasIn                int                                                      `json:"allowas-in"`
	AllowasInCount           int                                                      `json:"allowas-in-count" dval:"3"`
	AsOriginationInterval    int                                                      `json:"as-origination-interval"`
	Dynamic                  int                                                      `json:"dynamic"`
	PrefixListDirection      string                                                   `json:"prefix-list-direction"`
	RouteRefresh             int                                                      `json:"route-refresh" dval:"1"`
	GracefulRestart          int                                                      `json:"graceful-restart"`
	ExtendedNexthop          int                                                      `json:"extended-nexthop"`
	CollideEstablished       int                                                      `json:"collide-established"`
	DefaultOriginate         int                                                      `json:"default-originate"`
	RouteMap                 string                                                   `json:"route-map"`
	Description              string                                                   `json:"description"`
	DisallowInfiniteHoldtime int                                                      `json:"disallow-infinite-holdtime"`
	DistributeLists          []RouterBgpNeighborIpv6NeighborListDistributeLists       `json:"distribute-lists"`
	AcosApplicationOnly      int                                                      `json:"acos-application-only"`
	Telemetry                int                                                      `json:"telemetry"`
	DontCapabilityNegotiate  int                                                      `json:"dont-capability-negotiate"`
	EbgpMultihop             int                                                      `json:"ebgp-multihop"`
	EbgpMultihopHopCount     int                                                      `json:"ebgp-multihop-hop-count"`
	EnforceMultihop          int                                                      `json:"enforce-multihop"`
	Bfd                      int                                                      `json:"bfd"`
	Multihop                 int                                                      `json:"multihop"`
	KeyId                    int                                                      `json:"key-id"`
	KeyType                  string                                                   `json:"key-type"`
	BfdValue                 string                                                   `json:"bfd-value"`
	BfdEncrypted             string                                                   `json:"bfd-encrypted"`
	NeighborFilterLists      []RouterBgpNeighborIpv6NeighborListNeighborFilterLists   `json:"neighbor-filter-lists"`
	MaximumPrefix            int                                                      `json:"maximum-prefix"`
	MaximumPrefixThres       int                                                      `json:"maximum-prefix-thres"`
	RestartMin               int                                                      `json:"restart-min"`
	NextHopSelf              int                                                      `json:"next-hop-self"`
	OverrideCapability       int                                                      `json:"override-capability"`
	PassValue                string                                                   `json:"pass-value"`
	PassEncrypted            string                                                   `json:"pass-encrypted"`
	Passive                  int                                                      `json:"passive"`
	NeighborPrefixLists      []RouterBgpNeighborIpv6NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists"`
	RemovePrivateAs          int                                                      `json:"remove-private-as"`
	NeighborRouteMapLists    []RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
	SendCommunityVal         string                                                   `json:"send-community-val" dval:"both"`
	Inbound                  int                                                      `json:"inbound"`
	Shutdown                 int                                                      `json:"shutdown"`
	StrictCapabilityMatch    int                                                      `json:"strict-capability-match"`
	TimersKeepalive          int                                                      `json:"timers-keepalive" dval:"30"`
	TimersHoldtime           int                                                      `json:"timers-holdtime" dval:"90"`
	Connect                  int                                                      `json:"connect"`
	UnsuppressMap            string                                                   `json:"unsuppress-map"`
	UpdateSourceIp           string                                                   `json:"update-source-ip"`
	UpdateSourceIpv6         string                                                   `json:"update-source-ipv6"`
	Ethernet                 int                                                      `json:"ethernet"`
	Loopback                 int                                                      `json:"loopback"`
	Ve                       int                                                      `json:"ve"`
	Trunk                    int                                                      `json:"trunk"`
	Lif                      string                                                   `json:"lif"`
	Tunnel                   int                                                      `json:"tunnel"`
	Weight                   int                                                      `json:"weight"`
	Uuid                     string                                                   `json:"uuid"`
}

type RouterBgpNeighborIpv6NeighborListDistributeLists struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpNeighborIpv6NeighborListNeighborFilterLists struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpNeighborIpv6NeighborListNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpNeighborIpv6NeighborListNeighborRouteMapLists struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpNeighborEthernetNeighborList struct {
	Ethernet      int    `json:"ethernet"`
	Unnumbered    int    `json:"unnumbered"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpNeighborVeNeighborList struct {
	Ve            int    `json:"ve"`
	Unnumbered    int    `json:"unnumbered"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpNeighborTrunkNeighborList struct {
	Trunk         int    `json:"trunk"`
	Unnumbered    int    `json:"unnumbered"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpNetwork1315 struct {
	Synchronization RouterBgpNetworkSynchronization1316 `json:"synchronization"`
	Monitor         RouterBgpNetworkMonitor1317         `json:"monitor"`
	IpCidrList      []RouterBgpNetworkIpCidrList        `json:"ip-cidr-list"`
}

type RouterBgpNetworkSynchronization1316 struct {
	NetworkSynchronization int    `json:"network-synchronization"`
	Uuid                   string `json:"uuid"`
}

type RouterBgpNetworkMonitor1317 struct {
	Default RouterBgpNetworkMonitorDefault1318 `json:"default"`
}

type RouterBgpNetworkMonitorDefault1318 struct {
	NetworkMonitorDefault int    `json:"network-monitor-default"`
	Uuid                  string `json:"uuid"`
}

type RouterBgpNetworkIpCidrList struct {
	NetworkIpv4Cidr string `json:"network-ipv4-cidr"`
	RouteMap        string `json:"route-map"`
	Backdoor        int    `json:"backdoor"`
	Description     string `json:"description"`
	CommValue       string `json:"comm-value"`
	LcommValue      string `json:"lcomm-value"`
	Uuid            string `json:"uuid"`
}

type RouterBgpRedistribute1319 struct {
	ConnectedCfg  RouterBgpRedistributeConnectedCfg1320  `json:"connected-cfg"`
	FloatingIpCfg RouterBgpRedistributeFloatingIpCfg1321 `json:"floating-ip-cfg"`
	Lw4o6Cfg      RouterBgpRedistributeLw4o6Cfg1322      `json:"lw4o6-cfg"`
	StaticNatCfg  RouterBgpRedistributeStaticNatCfg1323  `json:"static-nat-cfg"`
	IpNatCfg      RouterBgpRedistributeIpNatCfg1324      `json:"ip-nat-cfg"`
	IpNatListCfg  RouterBgpRedistributeIpNatListCfg1325  `json:"ip-nat-list-cfg"`
	IsisCfg       RouterBgpRedistributeIsisCfg1326       `json:"isis-cfg"`
	OspfCfg       RouterBgpRedistributeOspfCfg1327       `json:"ospf-cfg"`
	RipCfg        RouterBgpRedistributeRipCfg1328        `json:"rip-cfg"`
	StaticCfg     RouterBgpRedistributeStaticCfg1329     `json:"static-cfg"`
	NatMapCfg     RouterBgpRedistributeNatMapCfg1330     `json:"nat-map-cfg"`
	PublicIpCfg   RouterBgpRedistributePublicIpCfg1331   `json:"public-ip-cfg"`
	Vip           RouterBgpRedistributeVip1332           `json:"vip"`
	Uuid          string                                 `json:"uuid"`
}

type RouterBgpRedistributeConnectedCfg1320 struct {
	Connected int    `json:"connected"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpRedistributeFloatingIpCfg1321 struct {
	FloatingIp int    `json:"floating-ip"`
	RouteMap   string `json:"route-map"`
}

type RouterBgpRedistributeLw4o6Cfg1322 struct {
	Lw4o6    int    `json:"lw4o6"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeStaticNatCfg1323 struct {
	StaticNat int    `json:"static-nat"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpRedistributeIpNatCfg1324 struct {
	IpNat    int    `json:"ip-nat"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeIpNatListCfg1325 struct {
	IpNatList int    `json:"ip-nat-list"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpRedistributeIsisCfg1326 struct {
	Isis     int    `json:"isis"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeOspfCfg1327 struct {
	Ospf     int    `json:"ospf"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeRipCfg1328 struct {
	Rip      int    `json:"rip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeStaticCfg1329 struct {
	Static   int    `json:"static"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeNatMapCfg1330 struct {
	NatMap   int    `json:"nat-map"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributePublicIpCfg1331 struct {
	PublicIp int    `json:"public-ip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeVip1332 struct {
	OnlyFlaggedCfg    RouterBgpRedistributeVipOnlyFlaggedCfg1333    `json:"only-flagged-cfg"`
	OnlyNotFlaggedCfg RouterBgpRedistributeVipOnlyNotFlaggedCfg1334 `json:"only-not-flagged-cfg"`
}

type RouterBgpRedistributeVipOnlyFlaggedCfg1333 struct {
	OnlyFlagged int    `json:"only-flagged"`
	RouteMap    string `json:"route-map"`
}

type RouterBgpRedistributeVipOnlyNotFlaggedCfg1334 struct {
	OnlyNotFlagged int    `json:"only-not-flagged"`
	RouteMap       string `json:"route-map"`
}

type RouterBgpTimers struct {
	BgpKeepalive int `json:"bgp-keepalive" dval:"30"`
	BgpHoldtime  int `json:"bgp-holdtime" dval:"90"`
}

func (p *RouterBgp) GetId() string {
	return p.Inst.AsNumber
}

func (p *RouterBgp) getPath() string {
	return "router/bgp"
}

func (p *RouterBgp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgp::Post")
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

func (p *RouterBgp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgp::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *RouterBgp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgp::Put")
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

func (p *RouterBgp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgp::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
