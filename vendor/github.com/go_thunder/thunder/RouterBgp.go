package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgp struct {
	RouterBgpInstanceAsNumber RouterBgpInstance `json:"bgp,omitempty"`
}

type RouterBgpInstance struct {
	RouterBgpInstanceAddressFamilyIpv6                    RouterBgpInstanceAddressFamily          `json:"address-family,omitempty"`
	RouterBgpInstanceAggregateAddressListAggregateAddress []RouterBgpInstanceAggregateAddressList `json:"aggregate-address-list,omitempty"`
	RouterBgpInstanceAsNumber                             int                                     `json:"as-number,omitempty"`
	RouterBgpInstanceAutoSummary                          int                                     `json:"auto-summary,omitempty"`
	RouterBgpInstanceBgpAlwaysCompareMed                  RouterBgpInstanceBgp                    `json:"bgp,omitempty"`
	RouterBgpInstanceDistanceListAdminDistance            []RouterBgpInstanceDistanceList         `json:"distance-list,omitempty"`
	RouterBgpInstanceMaximumPathsValue                    int                                     `json:"maximum-paths-value,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborList        RouterBgpInstanceNeighbor               `json:"neighbor,omitempty"`
	RouterBgpInstanceNetworkSynchronization               RouterBgpInstanceNetwork                `json:"network,omitempty"`
	RouterBgpInstanceOriginate                            int                                     `json:"originate,omitempty"`
	RouterBgpInstanceRedistributeConnectedCfg             RouterBgpInstanceRedistribute           `json:"redistribute,omitempty"`
	RouterBgpInstanceSynchronization                      int                                     `json:"synchronization,omitempty"`
	RouterBgpInstanceTimersBgpKeepalive                   RouterBgpInstanceTimers                 `json:"timers,omitempty"`
	RouterBgpInstanceUUID                                 string                                  `json:"uuid,omitempty"`
	RouterBgpInstanceUserTag                              string                                  `json:"user-tag,omitempty"`
}

type RouterBgpInstanceAddressFamily struct {
	RouterBgpInstanceAddressFamilyIpv6Bgp RouterBgpInstanceAddressFamilyIpv6 `json:"ipv6,omitempty"`
}

type RouterBgpInstanceAggregateAddressList struct {
	RouterBgpInstanceAggregateAddressListAggregateAddress string `json:"aggregate-address,omitempty"`
	RouterBgpInstanceAggregateAddressListAsSet            int    `json:"as-set,omitempty"`
	RouterBgpInstanceAggregateAddressListSummaryOnly      int    `json:"summary-only,omitempty"`
}

type RouterBgpInstanceBgp struct {
	RouterBgpInstanceBgpAlwaysCompareMed      int                              `json:"always-compare-med,omitempty"`
	RouterBgpInstanceBgpBestpathCfgIgnore     RouterBgpInstanceBgpBestpathCfg  `json:"bestpath-cfg,omitempty"`
	RouterBgpInstanceBgpBgpRestartTime        int                              `json:"bgp-restart-time,omitempty"`
	RouterBgpInstanceBgpBgpStalepathTime      int                              `json:"bgp-stalepath-time,omitempty"`
	RouterBgpInstanceBgpDampeningCfgDampening RouterBgpInstanceBgpDampeningCfg `json:"dampening-cfg,omitempty"`
	RouterBgpInstanceBgpDeterministicMed      int                              `json:"deterministic-med,omitempty"`
	RouterBgpInstanceBgpEnforceFirstAs        int                              `json:"enforce-first-as,omitempty"`
	RouterBgpInstanceBgpFastExternalFailover  int                              `json:"fast-external-failover,omitempty"`
	RouterBgpInstanceBgpGracefulRestart       int                              `json:"graceful-restart,omitempty"`
	RouterBgpInstanceBgpLocalPreferenceValue  int                              `json:"local-preference-value,omitempty"`
	RouterBgpInstanceBgpLogNeighborChanges    int                              `json:"log-neighbor-changes,omitempty"`
	RouterBgpInstanceBgpNexthopTriggerCount   int                              `json:"nexthop-trigger-count,omitempty"`
	RouterBgpInstanceBgpOverrideValidation    int                              `json:"override-validation,omitempty"`
	RouterBgpInstanceBgpRouterID              string                           `json:"router-id,omitempty"`
	RouterBgpInstanceBgpScanTime              int                              `json:"scan-time,omitempty"`
}

type RouterBgpInstanceDistanceList struct {
	RouterBgpInstanceDistanceListAclStr          string `json:"acl-str,omitempty"`
	RouterBgpInstanceDistanceListAdminDistance   int    `json:"admin-distance,omitempty"`
	RouterBgpInstanceDistanceListExtRoutesDist   int    `json:"ext-routes-dist,omitempty"`
	RouterBgpInstanceDistanceListIntRoutesDist   int    `json:"int-routes-dist,omitempty"`
	RouterBgpInstanceDistanceListLocalRoutesDist int    `json:"local-routes-dist,omitempty"`
	RouterBgpInstanceDistanceListSrcPrefix       string `json:"src-prefix,omitempty"`
}

type RouterBgpInstanceNeighbor struct {
	RouterBgpInstanceNeighborEthernetNeighborListEthernet   []RouterBgpInstanceNeighborEthernetNeighborList  `json:"ethernet-neighbor-list,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborIpv4   []RouterBgpInstanceNeighborIpv4NeighborList      `json:"ipv4-neighbor-list,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborIpv6   []RouterBgpInstanceNeighborIpv6NeighborList      `json:"ipv6-neighbor-list,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListPeerGroup []RouterBgpInstanceNeighborPeerGroupNeighborList `json:"peer-group-neighbor-list,omitempty"`
	RouterBgpInstanceNeighborTrunkNeighborListTrunk         []RouterBgpInstanceNeighborTrunkNeighborList     `json:"trunk-neighbor-list,omitempty"`
	RouterBgpInstanceNeighborVeNeighborListVe               []RouterBgpInstanceNeighborVeNeighborList        `json:"ve-neighbor-list,omitempty"`
}

type RouterBgpInstanceNetwork struct {
	RouterBgpInstanceNetworkIPCidrListNetworkIpv4Cidr             []RouterBgpInstanceNetworkIPCidrList    `json:"ip-cidr-list,omitempty"`
	RouterBgpInstanceNetworkSynchronizationNetworkSynchronization RouterBgpInstanceNetworkSynchronization `json:"synchronization,omitempty"`
}

type RouterBgpInstanceRedistribute struct {
	RouterBgpInstanceRedistributeConnectedCfgConnected   RouterBgpInstanceRedistributeConnectedCfg  `json:"connected-cfg,omitempty"`
	RouterBgpInstanceRedistributeFloatingIPCfgFloatingIP RouterBgpInstanceRedistributeFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	RouterBgpInstanceRedistributeIPNatCfgIPNat           RouterBgpInstanceRedistributeIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	RouterBgpInstanceRedistributeIPNatListCfgIPNatList   RouterBgpInstanceRedistributeIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	RouterBgpInstanceRedistributeIsisCfgIsis             RouterBgpInstanceRedistributeIsisCfg       `json:"isis-cfg,omitempty"`
	RouterBgpInstanceRedistributeLw4O6CfgLw4O6           RouterBgpInstanceRedistributeLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	RouterBgpInstanceRedistributeNatMapCfgNatMap         RouterBgpInstanceRedistributeNatMapCfg     `json:"nat-map-cfg,omitempty"`
	RouterBgpInstanceRedistributeOspfCfgOspf             RouterBgpInstanceRedistributeOspfCfg       `json:"ospf-cfg,omitempty"`
	RouterBgpInstanceRedistributeRipCfgRip               RouterBgpInstanceRedistributeRipCfg        `json:"rip-cfg,omitempty"`
	RouterBgpInstanceRedistributeStaticCfgStatic         RouterBgpInstanceRedistributeStaticCfg     `json:"static-cfg,omitempty"`
	RouterBgpInstanceRedistributeStaticNatCfgStaticNat   RouterBgpInstanceRedistributeStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	RouterBgpInstanceRedistributeUUID                    string                                     `json:"uuid,omitempty"`
	RouterBgpInstanceRedistributeVipOnlyFlaggedCfg       RouterBgpInstanceRedistributeVip           `json:"vip,omitempty"`
}

type RouterBgpInstanceTimers struct {
	RouterBgpInstanceTimersBgpHoldtime  int `json:"bgp-holdtime,omitempty"`
	RouterBgpInstanceTimersBgpKeepalive int `json:"bgp-keepalive,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6 struct {
	RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAggregateAddress []RouterBgpInstanceAddressFamilyIpv6AggregateAddressList `json:"aggregate-address-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6AutoSummary                          int                                                      `json:"auto-summary,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpDampening                         RouterBgpInstanceAddressFamilyIpv6Bgp                    `json:"bgp,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6DistanceDistanceExt                  RouterBgpInstanceAddressFamilyIpv6Distance               `json:"distance,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6MaximumPathsValue                    int                                                      `json:"maximum-paths-value,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborList        RouterBgpInstanceAddressFamilyIpv6Neighbor               `json:"neighbor,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkSynchronization               RouterBgpInstanceAddressFamilyIpv6Network                `json:"network,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6Originate                            int                                                      `json:"originate,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfg             RouterBgpInstanceAddressFamilyIpv6Redistribute           `json:"redistribute,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6Synchronization                      int                                                      `json:"synchronization,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6UUID                                 string                                                   `json:"uuid,omitempty"`
}

type RouterBgpInstanceBgpBestpathCfg struct {
	RouterBgpInstanceBgpBestpathCfgCompareRouterid int `json:"compare-routerid,omitempty"`
	RouterBgpInstanceBgpBestpathCfgIgnore          int `json:"ignore,omitempty"`
	RouterBgpInstanceBgpBestpathCfgMissingAsWorst  int `json:"missing-as-worst,omitempty"`
	RouterBgpInstanceBgpBestpathCfgRemoveRecvMed   int `json:"remove-recv-med,omitempty"`
	RouterBgpInstanceBgpBestpathCfgRemoveSendMed   int `json:"remove-send-med,omitempty"`
}

type RouterBgpInstanceBgpDampeningCfg struct {
	RouterBgpInstanceBgpDampeningCfgDampening           int    `json:"dampening,omitempty"`
	RouterBgpInstanceBgpDampeningCfgDampeningHalfTime   int    `json:"dampening-half-time,omitempty"`
	RouterBgpInstanceBgpDampeningCfgDampeningMaxSupress int    `json:"dampening-max-supress,omitempty"`
	RouterBgpInstanceBgpDampeningCfgDampeningPenalty    int    `json:"dampening-penalty,omitempty"`
	RouterBgpInstanceBgpDampeningCfgDampeningReuse      int    `json:"dampening-reuse,omitempty"`
	RouterBgpInstanceBgpDampeningCfgDampeningSupress    int    `json:"dampening-supress,omitempty"`
	RouterBgpInstanceBgpDampeningCfgRouteMap            string `json:"route-map,omitempty"`
}

type RouterBgpInstanceNeighborEthernetNeighborList struct {
	RouterBgpInstanceNeighborEthernetNeighborListEthernet      int    `json:"ethernet,omitempty"`
	RouterBgpInstanceNeighborEthernetNeighborListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpInstanceNeighborEthernetNeighborListUUID          string `json:"uuid,omitempty"`
	RouterBgpInstanceNeighborEthernetNeighborListUnnumbered    int    `json:"unnumbered,omitempty"`
}

type RouterBgpInstanceNeighborIpv4NeighborList struct {
	RouterBgpInstanceNeighborIpv4NeighborListActivate                         int                                                              `json:"activate,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListAllowasIn                        int                                                              `json:"allowas-in,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListAllowasInCount                   int                                                              `json:"allowas-in-count,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListDefaultOriginate                 int                                                              `json:"default-originate,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeList    []RouterBgpInstanceNeighborIpv4NeighborListDistributeLists       `json:"distribute-lists,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListGracefulRestart                  int                                                              `json:"graceful-restart,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListInbound                          int                                                              `json:"inbound,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListMaximumPrefix                    int                                                              `json:"maximum-prefix,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListMaximumPrefixThres               int                                                              `json:"maximum-prefix-thres,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterList    []RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborIpv4                     string                                                           `json:"neighbor-ipv4,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList []RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNextHopSelf                      int                                                              `json:"next-hop-self,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListPeerGroupName                    string                                                           `json:"peer-group-name,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListPrefixListDirection              string                                                           `json:"prefix-list-direction,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListRemovePrivateAs                  int                                                              `json:"remove-private-as,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListRestartMin                       int                                                              `json:"restart-min,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListRouteMap                         string                                                           `json:"route-map,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListSendCommunityVal                 string                                                           `json:"send-community-val,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListUUID                             string                                                           `json:"uuid,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListUnsuppressMap                    string                                                           `json:"unsuppress-map,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListWeight                           int                                                              `json:"weight,omitempty"`
}

type RouterBgpInstanceNeighborIpv6NeighborList struct {
	RouterBgpInstanceNeighborIpv6NeighborListActivate                         int                                                              `json:"activate,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListAllowasIn                        int                                                              `json:"allowas-in,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListAllowasInCount                   int                                                              `json:"allowas-in-count,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListDefaultOriginate                 int                                                              `json:"default-originate,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeList    []RouterBgpInstanceNeighborIpv6NeighborListDistributeLists       `json:"distribute-lists,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListGracefulRestart                  int                                                              `json:"graceful-restart,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListInbound                          int                                                              `json:"inbound,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListMaximumPrefix                    int                                                              `json:"maximum-prefix,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListMaximumPrefixThres               int                                                              `json:"maximum-prefix-thres,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterList    []RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborIpv6                     string                                                           `json:"neighbor-ipv6,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList []RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNextHopSelf                      int                                                              `json:"next-hop-self,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListPeerGroupName                    string                                                           `json:"peer-group-name,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListPrefixListDirection              string                                                           `json:"prefix-list-direction,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListRemovePrivateAs                  int                                                              `json:"remove-private-as,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListRestartMin                       int                                                              `json:"restart-min,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListRouteMap                         string                                                           `json:"route-map,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListSendCommunityVal                 string                                                           `json:"send-community-val,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListUUID                             string                                                           `json:"uuid,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListUnsuppressMap                    string                                                           `json:"unsuppress-map,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListWeight                           int                                                              `json:"weight,omitempty"`
}

type RouterBgpInstanceNeighborPeerGroupNeighborList struct {
	RouterBgpInstanceNeighborPeerGroupNeighborListActivate                         int                                                                   `json:"activate,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListAllowasIn                        int                                                                   `json:"allowas-in,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListAllowasInCount                   int                                                                   `json:"allowas-in-count,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListInbound                          int                                                                   `json:"inbound,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListMaximumPrefix                    int                                                                   `json:"maximum-prefix,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListMaximumPrefixThres               int                                                                   `json:"maximum-prefix-thres,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListNextHopSelf                      int                                                                   `json:"next-hop-self,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListPeerGroup                        string                                                                `json:"peer-group,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListRemovePrivateAs                  int                                                                   `json:"remove-private-as,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListUUID                             string                                                                `json:"uuid,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListWeight                           int                                                                   `json:"weight,omitempty"`
}

type RouterBgpInstanceNeighborTrunkNeighborList struct {
	RouterBgpInstanceNeighborTrunkNeighborListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpInstanceNeighborTrunkNeighborListTrunk         int    `json:"trunk,omitempty"`
	RouterBgpInstanceNeighborTrunkNeighborListUUID          string `json:"uuid,omitempty"`
	RouterBgpInstanceNeighborTrunkNeighborListUnnumbered    int    `json:"unnumbered,omitempty"`
}

type RouterBgpInstanceNeighborVeNeighborList struct {
	RouterBgpInstanceNeighborVeNeighborListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpInstanceNeighborVeNeighborListUUID          string `json:"uuid,omitempty"`
	RouterBgpInstanceNeighborVeNeighborListUnnumbered    int    `json:"unnumbered,omitempty"`
	RouterBgpInstanceNeighborVeNeighborListVe            int    `json:"ve,omitempty"`
}

type RouterBgpInstanceNetworkIPCidrList struct {
	RouterBgpInstanceNetworkIPCidrListBackdoor        int    `json:"backdoor,omitempty"`
	RouterBgpInstanceNetworkIPCidrListCommValue       string `json:"comm-value,omitempty"`
	RouterBgpInstanceNetworkIPCidrListDescription     string `json:"description,omitempty"`
	RouterBgpInstanceNetworkIPCidrListNetworkIpv4Cidr string `json:"network-ipv4-cidr,omitempty"`
	RouterBgpInstanceNetworkIPCidrListRouteMap        string `json:"route-map,omitempty"`
	RouterBgpInstanceNetworkIPCidrListUUID            string `json:"uuid,omitempty"`
}

type RouterBgpInstanceNetworkSynchronization struct {
	RouterBgpInstanceNetworkSynchronizationNetworkSynchronization int    `json:"network-synchronization,omitempty"`
	RouterBgpInstanceNetworkSynchronizationUUID                   string `json:"uuid,omitempty"`
}

type RouterBgpInstanceRedistributeConnectedCfg struct {
	RouterBgpInstanceRedistributeConnectedCfgConnected int    `json:"connected,omitempty"`
	RouterBgpInstanceRedistributeConnectedCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeFloatingIPCfg struct {
	RouterBgpInstanceRedistributeFloatingIPCfgFloatingIP int    `json:"floating-ip,omitempty"`
	RouterBgpInstanceRedistributeFloatingIPCfgRouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeIPNatCfg struct {
	RouterBgpInstanceRedistributeIPNatCfgIPNat    int    `json:"ip-nat,omitempty"`
	RouterBgpInstanceRedistributeIPNatCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeIPNatListCfg struct {
	RouterBgpInstanceRedistributeIPNatListCfgIPNatList int    `json:"ip-nat-list,omitempty"`
	RouterBgpInstanceRedistributeIPNatListCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeIsisCfg struct {
	RouterBgpInstanceRedistributeIsisCfgIsis     int    `json:"isis,omitempty"`
	RouterBgpInstanceRedistributeIsisCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeLw4O6Cfg struct {
	RouterBgpInstanceRedistributeLw4O6CfgLw4O6    int    `json:"lw4o6,omitempty"`
	RouterBgpInstanceRedistributeLw4O6CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeNatMapCfg struct {
	RouterBgpInstanceRedistributeNatMapCfgNatMap   int    `json:"nat-map,omitempty"`
	RouterBgpInstanceRedistributeNatMapCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeOspfCfg struct {
	RouterBgpInstanceRedistributeOspfCfgOspf     int    `json:"ospf,omitempty"`
	RouterBgpInstanceRedistributeOspfCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeRipCfg struct {
	RouterBgpInstanceRedistributeRipCfgRip      int    `json:"rip,omitempty"`
	RouterBgpInstanceRedistributeRipCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeStaticCfg struct {
	RouterBgpInstanceRedistributeStaticCfgRouteMap string `json:"route-map,omitempty"`
	RouterBgpInstanceRedistributeStaticCfgStatic   int    `json:"static,omitempty"`
}

type RouterBgpInstanceRedistributeStaticNatCfg struct {
	RouterBgpInstanceRedistributeStaticNatCfgRouteMap  string `json:"route-map,omitempty"`
	RouterBgpInstanceRedistributeStaticNatCfgStaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpInstanceRedistributeVip struct {
	RouterBgpInstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged       RouterBgpInstanceRedistributeVipOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6AggregateAddressList struct {
	RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAggregateAddress string `json:"aggregate-address,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6AggregateAddressListAsSet            int    `json:"as-set,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6AggregateAddressListSummaryOnly      int    `json:"summary-only,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6Bgp struct {
	RouterBgpInstanceAddressFamilyIpv6BgpDampening               int    `json:"dampening,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpDampeningHalf           int    `json:"dampening-half,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpDampeningMaxSupress     int    `json:"dampening-max-supress,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpDampeningStartReuse     int    `json:"dampening-start-reuse,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpDampeningStartSupress   int    `json:"dampening-start-supress,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpDampeningUnreachability int    `json:"dampening-unreachability,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6BgpRouteMap                string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6Distance struct {
	RouterBgpInstanceAddressFamilyIpv6DistanceDistanceExt   int `json:"distance-ext,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6DistanceDistanceInt   int `json:"distance-int,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6DistanceDistanceLocal int `json:"distance-local,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6Neighbor struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListEthernet []RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6List `json:"ethernet-neighbor-ipv6-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborIpv4     []RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborList         `json:"ipv4-neighbor-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborIpv6     []RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborList         `json:"ipv6-neighbor-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListPeerGroup   []RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborList    `json:"peer-group-neighbor-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListTrunk       []RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6List    `json:"trunk-neighbor-ipv6-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListVe             []RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6List       `json:"ve-neighbor-ipv6-list,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6Network struct {
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListNetworkIpv6            []RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkList `json:"ipv6-network-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkSynchronizationNetworkSynchronization RouterBgpInstanceAddressFamilyIpv6NetworkSynchronization   `json:"synchronization,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6Redistribute struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfgConnected   RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfg  `json:"connected-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfgFloatingIP RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfgIPNat           RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfgIPNatList   RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfgIsis             RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfg       `json:"isis-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6CfgLw4O6           RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeNat64CfgNat64           RouterBgpInstanceAddressFamilyIpv6RedistributeNat64Cfg      `json:"nat64-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfgNatMap         RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfg     `json:"nat-map-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfgOspf             RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfg       `json:"ospf-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfgRip               RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfg        `json:"rip-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfgStatic         RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfg     `json:"static-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfgStaticNat   RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeUUID                    string                                                      `json:"uuid,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg       RouterBgpInstanceAddressFamilyIpv6RedistributeVip           `json:"vip,omitempty"`
}

type RouterBgpInstanceNeighborIpv4NeighborListDistributeLists struct {
	RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeList          string `json:"distribute-list,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListDistributeListsDistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterLists struct {
	RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterList          string `json:"filter-list,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborFilterListsFilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixLists struct {
	RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapLists struct {
	RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpInstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpInstanceNeighborIpv6NeighborListDistributeLists struct {
	RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeList          string `json:"distribute-list,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListDistributeListsDistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterLists struct {
	RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterList          string `json:"filter-list,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborFilterListsFilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixLists struct {
	RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapLists struct {
	RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpInstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapLists struct {
	RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpInstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpInstanceRedistributeVipOnlyFlaggedCfg struct {
	RouterBgpInstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged int    `json:"only-flagged,omitempty"`
	RouterBgpInstanceRedistributeVipOnlyFlaggedCfgRouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfg struct {
	RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouterBgpInstanceRedistributeVipOnlyNotFlaggedCfgRouteMap       string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6List struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListEthernet      int    `json:"ethernet,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborEthernetNeighborIpv6ListUUID          string `json:"uuid,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborList struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListActivate                         int                                                                               `json:"activate,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListAllowasIn                        int                                                                               `json:"allowas-in,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListAllowasInCount                   int                                                                               `json:"allowas-in-count,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDefaultOriginate                 int                                                                               `json:"default-originate,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeList    []RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists       `json:"distribute-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListGracefulRestart                  int                                                                               `json:"graceful-restart,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListInbound                          int                                                                               `json:"inbound,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListMaximumPrefix                    int                                                                               `json:"maximum-prefix,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListMaximumPrefixThres               int                                                                               `json:"maximum-prefix-thres,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterList    []RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborIpv4                     string                                                                            `json:"neighbor-ipv4,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList []RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNextHopSelf                      int                                                                               `json:"next-hop-self,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListPeerGroupName                    string                                                                            `json:"peer-group-name,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListPrefixListDirection              string                                                                            `json:"prefix-list-direction,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListRemovePrivateAs                  int                                                                               `json:"remove-private-as,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListRestartMin                       int                                                                               `json:"restart-min,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListRouteMap                         string                                                                            `json:"route-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListSendCommunityVal                 string                                                                            `json:"send-community-val,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListUUID                             string                                                                            `json:"uuid,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListUnsuppressMap                    string                                                                            `json:"unsuppress-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListWeight                           int                                                                               `json:"weight,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborList struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListActivate                         int                                                                               `json:"activate,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListAllowasIn                        int                                                                               `json:"allowas-in,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListAllowasInCount                   int                                                                               `json:"allowas-in-count,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDefaultOriginate                 int                                                                               `json:"default-originate,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeList    []RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists       `json:"distribute-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListGracefulRestart                  int                                                                               `json:"graceful-restart,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListInbound                          int                                                                               `json:"inbound,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListMaximumPrefix                    int                                                                               `json:"maximum-prefix,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListMaximumPrefixThres               int                                                                               `json:"maximum-prefix-thres,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterList    []RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborIpv6                     string                                                                            `json:"neighbor-ipv6,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList []RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNextHopSelf                      int                                                                               `json:"next-hop-self,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListPeerGroupName                    string                                                                            `json:"peer-group-name,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListPrefixListDirection              string                                                                            `json:"prefix-list-direction,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListRemovePrivateAs                  int                                                                               `json:"remove-private-as,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListRestartMin                       int                                                                               `json:"restart-min,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListRouteMap                         string                                                                            `json:"route-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListSendCommunityVal                 string                                                                            `json:"send-community-val,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListUUID                             string                                                                            `json:"uuid,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListUnsuppressMap                    string                                                                            `json:"unsuppress-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListWeight                           int                                                                               `json:"weight,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborList struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListActivate                         int                                                                                    `json:"activate,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListAllowasIn                        int                                                                                    `json:"allowas-in,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListAllowasInCount                   int                                                                                    `json:"allowas-in-count,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListInbound                          int                                                                                    `json:"inbound,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListMaximumPrefix                    int                                                                                    `json:"maximum-prefix,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListMaximumPrefixThres               int                                                                                    `json:"maximum-prefix-thres,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNextHopSelf                      int                                                                                    `json:"next-hop-self,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListPeerGroup                        string                                                                                 `json:"peer-group,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListRemovePrivateAs                  int                                                                                    `json:"remove-private-as,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListUUID                             string                                                                                 `json:"uuid,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListWeight                           int                                                                                    `json:"weight,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6List struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListTrunk         int    `json:"trunk,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborTrunkNeighborIpv6ListUUID          string `json:"uuid,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6List struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListUUID          string `json:"uuid,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborVeNeighborIpv6ListVe            int    `json:"ve,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkList struct {
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListBackdoor    int    `json:"backdoor,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListCommValue   string `json:"comm-value,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListDescription string `json:"description,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListNetworkIpv6 string `json:"network-ipv6,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListRouteMap    string `json:"route-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkIpv6NetworkListUUID        string `json:"uuid,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NetworkSynchronization struct {
	RouterBgpInstanceAddressFamilyIpv6NetworkSynchronizationNetworkSynchronization int    `json:"network-synchronization,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NetworkSynchronizationUUID                   string `json:"uuid,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfgConnected int    `json:"connected,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeConnectedCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfgFloatingIP int    `json:"floating-ip,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeFloatingIPCfgRouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfgIPNat    int    `json:"ip-nat,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfgIPNatList int    `json:"ip-nat-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeIPNatListCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfgIsis     int    `json:"isis,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeIsisCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6Cfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6CfgLw4O6    int    `json:"lw4o6,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeLw4O6CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeNat64Cfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeNat64CfgNat64    int    `json:"nat64,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeNat64CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfgNatMap   int    `json:"nat-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeNatMapCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfgOspf     int    `json:"ospf,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeOspfCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfgRip      int    `json:"rip,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeRipCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfgRouteMap string `json:"route-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeStaticCfgStatic   int    `json:"static,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfgRouteMap  string `json:"route-map,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeStaticNatCfgStaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeVip struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfgOnlyFlagged       RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeList          string `json:"distribute-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListDistributeListsDistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterList          string `json:"filter-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborFilterListsFilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborPrefixListsNbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeList          string `json:"distribute-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListDistributeListsDistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterList          string `json:"filter-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborFilterListsFilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborPrefixListsNbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapLists struct {
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6NeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfgOnlyFlagged int    `json:"only-flagged,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyFlaggedCfgRouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfg struct {
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouterBgpInstanceAddressFamilyIpv6RedistributeVipOnlyNotFlaggedCfgRouteMap       string `json:"route-map,omitempty"`
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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
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
		err := json.Unmarshal(data, &m)
		if err != nil {
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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
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
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgp
		err := json.Unmarshal(data, &m)
		if err != nil {
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
	return err
}
