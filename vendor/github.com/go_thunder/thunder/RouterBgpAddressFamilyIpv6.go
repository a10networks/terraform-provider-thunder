package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6 struct {
	RouterBgpAddressFamilyIpv6InstanceBgp RouterBgpAddressFamilyIpv6Instance `json:"ipv6,omitempty"`
}

type RouterBgpAddressFamilyIpv6Instance struct {
	RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAggregateAddress []RouterBgpAddressFamilyIpv6InstanceAggregateAddressList `json:"aggregate-address-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceAutoSummary                          int                                                      `json:"auto-summary,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpDampening                         RouterBgpAddressFamilyIpv6InstanceBgp                    `json:"bgp,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceDistanceDistanceExt                  RouterBgpAddressFamilyIpv6InstanceDistance               `json:"distance,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceMaximumPathsValue                    int                                                      `json:"maximum-paths-value,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborList        RouterBgpAddressFamilyIpv6InstanceNeighbor               `json:"neighbor,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkSynchronization               RouterBgpAddressFamilyIpv6InstanceNetwork                `json:"network,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceOriginate                            int                                                      `json:"originate,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfg             RouterBgpAddressFamilyIpv6InstanceRedistribute           `json:"redistribute,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceSynchronization                      int                                                      `json:"synchronization,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceUUID                                 string                                                   `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceAggregateAddressList struct {
	RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAggregateAddress string `json:"aggregate-address,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceAggregateAddressListAsSet            int    `json:"as-set,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceAggregateAddressListSummaryOnly      int    `json:"summary-only,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceBgp struct {
	RouterBgpAddressFamilyIpv6InstanceBgpDampening               int    `json:"dampening,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpDampeningHalf           int    `json:"dampening-half,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpDampeningMaxSupress     int    `json:"dampening-max-supress,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpDampeningStartReuse     int    `json:"dampening-start-reuse,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpDampeningStartSupress   int    `json:"dampening-start-supress,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpDampeningUnreachability int    `json:"dampening-unreachability,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceBgpRouteMap                string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceDistance struct {
	RouterBgpAddressFamilyIpv6InstanceDistanceDistanceExt   int `json:"distance-ext,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceDistanceDistanceInt   int `json:"distance-int,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceDistanceDistanceLocal int `json:"distance-local,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighbor struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListEthernet []RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6List `json:"ethernet-neighbor-ipv6-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborIpv4     []RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborList         `json:"ipv4-neighbor-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborIpv6     []RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborList         `json:"ipv6-neighbor-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListPeerGroup   []RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborList    `json:"peer-group-neighbor-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListTrunk       []RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6List    `json:"trunk-neighbor-ipv6-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListVe             []RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6List       `json:"ve-neighbor-ipv6-list,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNetwork struct {
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListNetworkIpv6            []RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkList `json:"ipv6-network-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkSynchronizationNetworkSynchronization RouterBgpAddressFamilyIpv6InstanceNetworkSynchronization   `json:"synchronization,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistribute struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfgConnected   RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfg  `json:"connected-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfgFloatingIP RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfg `json:"floating-ip-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfgIPNat           RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfg      `json:"ip-nat-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfgIPNatList   RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfg  `json:"ip-nat-list-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfgIsis             RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfg       `json:"isis-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6CfgLw4O6           RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6Cfg      `json:"lw4o6-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeNat64CfgNat64           RouterBgpAddressFamilyIpv6InstanceRedistributeNat64Cfg      `json:"nat64-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfgNatMap         RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfg     `json:"nat-map-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfgOspf             RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfg       `json:"ospf-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfgRip               RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfg        `json:"rip-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfgStatic         RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfg     `json:"static-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfgStaticNat   RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfg  `json:"static-nat-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeUUID                    string                                                      `json:"uuid,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfg       RouterBgpAddressFamilyIpv6InstanceRedistributeVip           `json:"vip,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6List struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListEthernet      int    `json:"ethernet,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborEthernetNeighborIpv6ListUUID          string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborList struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListActivate                         int                                                                               `json:"activate,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListAllowasIn                        int                                                                               `json:"allowas-in,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListAllowasInCount                   int                                                                               `json:"allowas-in-count,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDefaultOriginate                 int                                                                               `json:"default-originate,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeList    []RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeLists       `json:"distribute-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListGracefulRestart                  int                                                                               `json:"graceful-restart,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListInbound                          int                                                                               `json:"inbound,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListMaximumPrefix                    int                                                                               `json:"maximum-prefix,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListMaximumPrefixThres               int                                                                               `json:"maximum-prefix-thres,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterList    []RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborIpv4                     string                                                                            `json:"neighbor-ipv4,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList []RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNextHopSelf                      int                                                                               `json:"next-hop-self,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListPeerGroupName                    string                                                                            `json:"peer-group-name,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListPrefixListDirection              string                                                                            `json:"prefix-list-direction,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListRemovePrivateAs                  int                                                                               `json:"remove-private-as,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListRestartMin                       int                                                                               `json:"restart-min,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListRouteMap                         string                                                                            `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListSendCommunityVal                 string                                                                            `json:"send-community-val,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListUUID                             string                                                                            `json:"uuid,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListUnsuppressMap                    string                                                                            `json:"unsuppress-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListWeight                           int                                                                               `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborList struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListActivate                         int                                                                               `json:"activate,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListAllowasIn                        int                                                                               `json:"allowas-in,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListAllowasInCount                   int                                                                               `json:"allowas-in-count,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDefaultOriginate                 int                                                                               `json:"default-originate,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeList    []RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeLists       `json:"distribute-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListGracefulRestart                  int                                                                               `json:"graceful-restart,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListInbound                          int                                                                               `json:"inbound,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListMaximumPrefix                    int                                                                               `json:"maximum-prefix,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListMaximumPrefixThres               int                                                                               `json:"maximum-prefix-thres,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterList    []RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborIpv6                     string                                                                            `json:"neighbor-ipv6,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList []RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNextHopSelf                      int                                                                               `json:"next-hop-self,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListPeerGroupName                    string                                                                            `json:"peer-group-name,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListPrefixListDirection              string                                                                            `json:"prefix-list-direction,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListRemovePrivateAs                  int                                                                               `json:"remove-private-as,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListRestartMin                       int                                                                               `json:"restart-min,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListRouteMap                         string                                                                            `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListSendCommunityVal                 string                                                                            `json:"send-community-val,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListUUID                             string                                                                            `json:"uuid,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListUnsuppressMap                    string                                                                            `json:"unsuppress-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListWeight                           int                                                                               `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborList struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListActivate                         int                                                                                    `json:"activate,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListAllowasIn                        int                                                                                    `json:"allowas-in,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListAllowasInCount                   int                                                                                    `json:"allowas-in-count,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListInbound                          int                                                                                    `json:"inbound,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListMaximumPrefix                    int                                                                                    `json:"maximum-prefix,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListMaximumPrefixThres               int                                                                                    `json:"maximum-prefix-thres,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap []RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNextHopSelf                      int                                                                                    `json:"next-hop-self,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListPeerGroup                        string                                                                                 `json:"peer-group,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListRemovePrivateAs                  int                                                                                    `json:"remove-private-as,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListUUID                             string                                                                                 `json:"uuid,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListWeight                           int                                                                                    `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6List struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListTrunk         int    `json:"trunk,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborTrunkNeighborIpv6ListUUID          string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6List struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListPeerGroupName string `json:"peer-group-name,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListUUID          string `json:"uuid,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborVeNeighborIpv6ListVe            int    `json:"ve,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkList struct {
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListBackdoor    int    `json:"backdoor,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListCommValue   string `json:"comm-value,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListDescription string `json:"description,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListNetworkIpv6 string `json:"network-ipv6,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListRouteMap    string `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkIpv6NetworkListUUID        string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNetworkSynchronization struct {
	RouterBgpAddressFamilyIpv6InstanceNetworkSynchronizationNetworkSynchronization int    `json:"network-synchronization,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNetworkSynchronizationUUID                   string `json:"uuid,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfgConnected int    `json:"connected,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeConnectedCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfgFloatingIP int    `json:"floating-ip,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeFloatingIPCfgRouteMap   string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfgIPNat    int    `json:"ip-nat,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfgIPNatList int    `json:"ip-nat-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeIPNatListCfgRouteMap  string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfgIsis     int    `json:"isis,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeIsisCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6Cfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6CfgLw4O6    int    `json:"lw4o6,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeLw4O6CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeNat64Cfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeNat64CfgNat64    int    `json:"nat64,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeNat64CfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfgNatMap   int    `json:"nat-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeNatMapCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfgOspf     int    `json:"ospf,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeOspfCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfgRip      int    `json:"rip,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeRipCfgRouteMap string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfgRouteMap string `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeStaticCfgStatic   int    `json:"static,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfgRouteMap  string `json:"route-map,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeStaticNatCfgStaticNat int    `json:"static-nat,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeVip struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged       RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfg    `json:"only-flagged-cfg,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfg `json:"only-not-flagged-cfg,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeList          string `json:"distribute-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListDistributeListsDistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterList          string `json:"filter-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborFilterListsFilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborPrefixListsNbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv4NeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeList          string `json:"distribute-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListDistributeListsDistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterList          string `json:"filter-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborFilterListsFilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborPrefixListsNbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborIpv6NeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapLists struct {
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceNeighborPeerGroupNeighborListNeighborRouteMapListsNbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfgOnlyFlagged int    `json:"only-flagged,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyFlaggedCfgRouteMap    string `json:"route-map,omitempty"`
}

type RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfg struct {
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfgOnlyNotFlagged int    `json:"only-not-flagged,omitempty"`
	RouterBgpAddressFamilyIpv6InstanceRedistributeVipOnlyNotFlaggedCfgRouteMap       string `json:"route-map,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6(id string, name1 string, inst RouterBgpAddressFamilyIpv6, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpAddressFamilyIpv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
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
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpAddressFamilyIpv6", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
