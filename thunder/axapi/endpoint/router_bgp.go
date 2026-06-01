package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type RouterBgp struct {
	Inst struct {
		AddressFamily RouterBgpAddressFamily1258 `json:"address-family"`

		AggregateAddressList []RouterBgpAggregateAddressList `json:"aggregate-address-list"`

		AsNumber string `json:"as-number"`

		AutoSummary int `json:"auto-summary"`

		Bgp RouterBgpBgp `json:"bgp"`

		DistanceList []RouterBgpDistanceList `json:"distance-list"`

		MaximumPathsValue int `json:"maximum-paths-value" dval:"1"`

		Neighbor RouterBgpNeighbor1313 `json:"neighbor"`

		Network RouterBgpNetwork1314 `json:"network"`

		Originate int `json:"originate"`

		Redistribute RouterBgpRedistribute1318 `json:"redistribute"`

		Synchronization int `json:"synchronization"`

		Timers RouterBgpTimers `json:"timers"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"bgp"`
}

type RouterBgpAddressFamily1258 struct {
	Ipv6         RouterBgpAddressFamilyIpv61259         `json:"ipv6"`
	Ipv4Flowspec RouterBgpAddressFamilyIpv4Flowspec1301 `json:"ipv4-flowspec"`
	Ipv6Flowspec RouterBgpAddressFamilyIpv6Flowspec1307 `json:"ipv6-flowspec"`
}

type RouterBgpAddressFamilyIpv61259 struct {
	Bgp                  RouterBgpAddressFamilyIpv6Bgp1260                    `json:"bgp"`
	Distance             RouterBgpAddressFamilyIpv6Distance1261               `json:"distance"`
	MaximumPathsValue    int                                                  `json:"maximum-paths-value" dval:"1"`
	Originate            int                                                  `json:"originate"`
	PreferGlobal         int                                                  `json:"prefer-global"`
	AggregateAddressList []RouterBgpAddressFamilyIpv6AggregateAddressList1262 `json:"aggregate-address-list"`
	AutoSummary          int                                                  `json:"auto-summary"`
	Synchronization      int                                                  `json:"synchronization"`
	Uuid                 string                                               `json:"uuid"`
	Network              RouterBgpAddressFamilyIpv6Network1263                `json:"network"`
	Neighbor             RouterBgpAddressFamilyIpv6Neighbor1268               `json:"neighbor"`
	Redistribute         RouterBgpAddressFamilyIpv6Redistribute1284           `json:"redistribute"`
}

type RouterBgpAddressFamilyIpv6Bgp1260 struct {
	Dampening               int    `json:"dampening"`
	DampeningHalf           int    `json:"dampening-half"`
	DampeningStartReuse     int    `json:"dampening-start-reuse"`
	DampeningStartSupress   int    `json:"dampening-start-supress"`
	DampeningMaxSupress     int    `json:"dampening-max-supress"`
	DampeningUnreachability int    `json:"dampening-unreachability"`
	RouteMap                string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6Distance1261 struct {
	DistanceExt   int `json:"distance-ext"`
	DistanceInt   int `json:"distance-int"`
	DistanceLocal int `json:"distance-local"`
}

type RouterBgpAddressFamilyIpv6AggregateAddressList1262 struct {
	AggregateAddress string `json:"aggregate-address"`
	AsSet            int    `json:"as-set"`
	SummaryOnly      int    `json:"summary-only"`
}

type RouterBgpAddressFamilyIpv6Network1263 struct {
	Synchronization RouterBgpAddressFamilyIpv6NetworkSynchronization1264   `json:"synchronization"`
	Monitor         RouterBgpAddressFamilyIpv6NetworkMonitor1265           `json:"monitor"`
	Ipv6NetworkList []RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267 `json:"ipv6-network-list"`
}

type RouterBgpAddressFamilyIpv6NetworkSynchronization1264 struct {
	NetworkSynchronization int    `json:"network-synchronization"`
	Uuid                   string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NetworkMonitor1265 struct {
	Default RouterBgpAddressFamilyIpv6NetworkMonitorDefault1266 `json:"default"`
}

type RouterBgpAddressFamilyIpv6NetworkMonitorDefault1266 struct {
	NetworkMonitorDefault int    `json:"network-monitor-default"`
	Uuid                  string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NetworkIpv6NetworkList1267 struct {
	NetworkIpv6 string `json:"network-ipv6"`
	RouteMap    string `json:"route-map"`
	Backdoor    int    `json:"backdoor"`
	Description string `json:"description"`
	CommValue   string `json:"comm-value"`
	LcommValue  string `json:"lcomm-value"`
	Uuid        string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6Neighbor1268 struct {
	PeerGroupNeighborList    []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269    `json:"peer-group-neighbor-list"`
	Ipv4NeighborList         []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271         `json:"ipv4-neighbor-list"`
	Ipv6NeighborList         []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276         `json:"ipv6-neighbor-list"`
	EthernetNeighborIpv6List []RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281 `json:"ethernet-neighbor-ipv6-list"`
	VeNeighborIpv6List       []RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282       `json:"ve-neighbor-ipv6-list"`
	TrunkNeighborIpv6List    []RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283    `json:"trunk-neighbor-ipv6-list"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborList1269 struct {
	PeerGroup             string                                                                             `json:"peer-group"`
	Activate              int                                                                                `json:"activate"`
	AllowasIn             int                                                                                `json:"allowas-in"`
	AllowasInCount        int                                                                                `json:"allowas-in-count" dval:"3"`
	MaximumPrefix         int                                                                                `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                                `json:"maximum-prefix-thres"`
	NextHopSelf           int                                                                                `json:"next-hop-self"`
	RemovePrivateAs       int                                                                                `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270 `json:"neighbor-route-map-lists"`
	Inbound               int                                                                                `json:"inbound"`
	Weight                int                                                                                `json:"weight"`
	Uuid                  string                                                                             `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists1270 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborList1271 struct {
	NeighborIpv4          string                                                                        `json:"neighbor-ipv4"`
	PeerGroupName         string                                                                        `json:"peer-group-name"`
	Activate              int                                                                           `json:"activate"`
	AllowasIn             int                                                                           `json:"allowas-in"`
	AllowasInCount        int                                                                           `json:"allowas-in-count" dval:"3"`
	PrefixListDirection   string                                                                        `json:"prefix-list-direction"`
	GracefulRestart       int                                                                           `json:"graceful-restart"`
	DefaultOriginate      int                                                                           `json:"default-originate"`
	RouteMap              string                                                                        `json:"route-map"`
	DistributeLists       []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272       `json:"distribute-lists"`
	NeighborFilterLists   []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273   `json:"neighbor-filter-lists"`
	MaximumPrefix         int                                                                           `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                           `json:"maximum-prefix-thres"`
	RestartMin            int                                                                           `json:"restart-min"`
	NextHopSelf           int                                                                           `json:"next-hop-self"`
	NeighborPrefixLists   []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274   `json:"neighbor-prefix-lists"`
	RemovePrivateAs       int                                                                           `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                        `json:"send-community-val" dval:"both"`
	Inbound               int                                                                           `json:"inbound"`
	UnsuppressMap         string                                                                        `json:"unsuppress-map"`
	Weight                int                                                                           `json:"weight"`
	Uuid                  string                                                                        `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists1272 struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists1273 struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists1274 struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists1275 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborList1276 struct {
	NeighborIpv6          string                                                                        `json:"neighbor-ipv6"`
	PeerGroupName         string                                                                        `json:"peer-group-name"`
	Activate              int                                                                           `json:"activate"`
	AllowasIn             int                                                                           `json:"allowas-in"`
	AllowasInCount        int                                                                           `json:"allowas-in-count" dval:"3"`
	PrefixListDirection   string                                                                        `json:"prefix-list-direction"`
	GracefulRestart       int                                                                           `json:"graceful-restart"`
	DefaultOriginate      int                                                                           `json:"default-originate"`
	RouteMap              string                                                                        `json:"route-map"`
	DistributeLists       []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277       `json:"distribute-lists"`
	NeighborFilterLists   []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278   `json:"neighbor-filter-lists"`
	MaximumPrefix         int                                                                           `json:"maximum-prefix"`
	MaximumPrefixThres    int                                                                           `json:"maximum-prefix-thres"`
	RestartMin            int                                                                           `json:"restart-min"`
	NextHopSelf           int                                                                           `json:"next-hop-self"`
	NeighborPrefixLists   []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279   `json:"neighbor-prefix-lists"`
	RemovePrivateAs       int                                                                           `json:"remove-private-as"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                        `json:"send-community-val" dval:"both"`
	Inbound               int                                                                           `json:"inbound"`
	UnsuppressMap         string                                                                        `json:"unsuppress-map"`
	Weight                int                                                                           `json:"weight"`
	Uuid                  string                                                                        `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists1277 struct {
	DistributeList          string `json:"distribute-list"`
	DistributeListDirection string `json:"distribute-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists1278 struct {
	FilterList          string `json:"filter-list"`
	FilterListDirection string `json:"filter-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists1279 struct {
	NbrPrefixList          string `json:"nbr-prefix-list"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists1280 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6List1281 struct {
	Ethernet      int    `json:"ethernet"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborVeNeighborIpv6List1282 struct {
	Ve            int    `json:"ve"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6NeighborTrunkNeighborIpv6List1283 struct {
	Trunk         int    `json:"trunk"`
	PeerGroupName string `json:"peer-group-name"`
	Uuid          string `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6Redistribute1284 struct {
	ConnectedCfg  RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1285  `json:"connected-cfg"`
	FloatingIpCfg RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1286 `json:"floating-ip-cfg"`
	Nat64Cfg      RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1287      `json:"nat64-cfg"`
	NatMapCfg     RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1288     `json:"nat-map-cfg"`
	Lw4o6Cfg      RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1289      `json:"lw4o6-cfg"`
	StaticNatCfg  RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1290  `json:"static-nat-cfg"`
	IpNatCfg      RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1291      `json:"ip-nat-cfg"`
	IpNatListCfg  RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1292  `json:"ip-nat-list-cfg"`
	IsisCfg       RouterBgpAddressFamilyIpv6RedistributeIsisCfg1293       `json:"isis-cfg"`
	OspfCfg       RouterBgpAddressFamilyIpv6RedistributeOspfCfg1294       `json:"ospf-cfg"`
	RipCfg        RouterBgpAddressFamilyIpv6RedistributeRipCfg1295        `json:"rip-cfg"`
	StaticCfg     RouterBgpAddressFamilyIpv6RedistributeStaticCfg1296     `json:"static-cfg"`
	PublicIpCfg   RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1297   `json:"public-ip-cfg"`
	Vip           RouterBgpAddressFamilyIpv6RedistributeVip1298           `json:"vip"`
	Uuid          string                                                  `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6RedistributeConnectedCfg1285 struct {
	Connected int    `json:"connected"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeFloatingIpCfg1286 struct {
	FloatingIp int    `json:"floating-ip"`
	RouteMap   string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeNat64Cfg1287 struct {
	Nat64    int    `json:"nat64"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeNatMapCfg1288 struct {
	NatMap   int    `json:"nat-map"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeLw4o6Cfg1289 struct {
	Lw4o6    int    `json:"lw4o6"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticNatCfg1290 struct {
	StaticNat int    `json:"static-nat"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpNatCfg1291 struct {
	IpNat    int    `json:"ip-nat"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIpNatListCfg1292 struct {
	IpNatList int    `json:"ip-nat-list"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeIsisCfg1293 struct {
	Isis     int    `json:"isis"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeOspfCfg1294 struct {
	Ospf     int    `json:"ospf"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeRipCfg1295 struct {
	Rip      int    `json:"rip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeStaticCfg1296 struct {
	Static   int    `json:"static"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributePublicIpCfg1297 struct {
	PublicIp int    `json:"public-ip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeVip1298 struct {
	OnlyFlaggedCfg    RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1299    `json:"only-flagged-cfg"`
	OnlyNotFlaggedCfg RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1300 `json:"only-not-flagged-cfg"`
}

type RouterBgpAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg1299 struct {
	OnlyFlagged int    `json:"only-flagged"`
	RouteMap    string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg1300 struct {
	OnlyNotFlagged int    `json:"only-not-flagged"`
	RouteMap       string `json:"route-map"`
}

type RouterBgpAddressFamilyIpv4Flowspec1301 struct {
	Uuid     string                                         `json:"uuid"`
	Neighbor RouterBgpAddressFamilyIpv4FlowspecNeighbor1302 `json:"neighbor"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighbor1302 struct {
	Ipv4NeighborList []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303 `json:"ipv4-neighbor-list"`
	Ipv6NeighborList []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305 `json:"ipv6-neighbor-list"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList1303 struct {
	NeighborIpv4          string                                                                                `json:"neighbor-ipv4"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1304 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList1305 struct {
	NeighborIpv6          string                                                                                `json:"neighbor-ipv6"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1306 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6Flowspec1307 struct {
	Uuid     string                                         `json:"uuid"`
	Neighbor RouterBgpAddressFamilyIpv6FlowspecNeighbor1308 `json:"neighbor"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighbor1308 struct {
	Ipv4NeighborList []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309 `json:"ipv4-neighbor-list"`
	Ipv6NeighborList []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311 `json:"ipv6-neighbor-list"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborList1309 struct {
	NeighborIpv4          string                                                                                `json:"neighbor-ipv4"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborListNeighborRouteMapLists1310 struct {
	NbrRouteMap      string `json:"nbr-route-map"`
	NbrRmapDirection string `json:"nbr-rmap-direction"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborList1311 struct {
	NeighborIpv6          string                                                                                `json:"neighbor-ipv6"`
	Activate              int                                                                                   `json:"activate"`
	NeighborRouteMapLists []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312 `json:"neighbor-route-map-lists"`
	SendCommunityVal      string                                                                                `json:"send-community-val" dval:"both"`
	Uuid                  string                                                                                `json:"uuid"`
}

type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv6NeighborListNeighborRouteMapLists1312 struct {
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

type RouterBgpNeighbor1313 struct {
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

type RouterBgpNetwork1314 struct {
	Synchronization RouterBgpNetworkSynchronization1315 `json:"synchronization"`
	Monitor         RouterBgpNetworkMonitor1316         `json:"monitor"`
	IpCidrList      []RouterBgpNetworkIpCidrList        `json:"ip-cidr-list"`
}

type RouterBgpNetworkSynchronization1315 struct {
	NetworkSynchronization int    `json:"network-synchronization"`
	Uuid                   string `json:"uuid"`
}

type RouterBgpNetworkMonitor1316 struct {
	Default RouterBgpNetworkMonitorDefault1317 `json:"default"`
}

type RouterBgpNetworkMonitorDefault1317 struct {
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

type RouterBgpRedistribute1318 struct {
	ConnectedCfg  RouterBgpRedistributeConnectedCfg1319  `json:"connected-cfg"`
	FloatingIpCfg RouterBgpRedistributeFloatingIpCfg1320 `json:"floating-ip-cfg"`
	Lw4o6Cfg      RouterBgpRedistributeLw4o6Cfg1321      `json:"lw4o6-cfg"`
	StaticNatCfg  RouterBgpRedistributeStaticNatCfg1322  `json:"static-nat-cfg"`
	IpNatCfg      RouterBgpRedistributeIpNatCfg1323      `json:"ip-nat-cfg"`
	IpNatListCfg  RouterBgpRedistributeIpNatListCfg1324  `json:"ip-nat-list-cfg"`
	IsisCfg       RouterBgpRedistributeIsisCfg1325       `json:"isis-cfg"`
	OspfCfg       RouterBgpRedistributeOspfCfg1326       `json:"ospf-cfg"`
	RipCfg        RouterBgpRedistributeRipCfg1327        `json:"rip-cfg"`
	StaticCfg     RouterBgpRedistributeStaticCfg1328     `json:"static-cfg"`
	NatMapCfg     RouterBgpRedistributeNatMapCfg1329     `json:"nat-map-cfg"`
	PublicIpCfg   RouterBgpRedistributePublicIpCfg1330   `json:"public-ip-cfg"`
	Vip           RouterBgpRedistributeVip1331           `json:"vip"`
	Uuid          string                                 `json:"uuid"`
}

type RouterBgpRedistributeConnectedCfg1319 struct {
	Connected int    `json:"connected"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpRedistributeFloatingIpCfg1320 struct {
	FloatingIp int    `json:"floating-ip"`
	RouteMap   string `json:"route-map"`
}

type RouterBgpRedistributeLw4o6Cfg1321 struct {
	Lw4o6    int    `json:"lw4o6"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeStaticNatCfg1322 struct {
	StaticNat int    `json:"static-nat"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpRedistributeIpNatCfg1323 struct {
	IpNat    int    `json:"ip-nat"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeIpNatListCfg1324 struct {
	IpNatList int    `json:"ip-nat-list"`
	RouteMap  string `json:"route-map"`
}

type RouterBgpRedistributeIsisCfg1325 struct {
	Isis     int    `json:"isis"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeOspfCfg1326 struct {
	Ospf     int    `json:"ospf"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeRipCfg1327 struct {
	Rip      int    `json:"rip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeStaticCfg1328 struct {
	Static   int    `json:"static"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeNatMapCfg1329 struct {
	NatMap   int    `json:"nat-map"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributePublicIpCfg1330 struct {
	PublicIp int    `json:"public-ip"`
	RouteMap string `json:"route-map"`
}

type RouterBgpRedistributeVip1331 struct {
	OnlyFlaggedCfg    RouterBgpRedistributeVipOnlyFlaggedCfg1332    `json:"only-flagged-cfg"`
	OnlyNotFlaggedCfg RouterBgpRedistributeVipOnlyNotFlaggedCfg1333 `json:"only-not-flagged-cfg"`
}

type RouterBgpRedistributeVipOnlyFlaggedCfg1332 struct {
	OnlyFlagged int    `json:"only-flagged"`
	RouteMap    string `json:"route-map"`
}

type RouterBgpRedistributeVipOnlyNotFlaggedCfg1333 struct {
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
