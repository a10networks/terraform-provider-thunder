package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceEthernet struct {
	InterfaceEthernetInstanceIfnum InterfaceEthernetInstance `json:"ethernet,omitempty"`
}

type InterfaceEthernetInstance struct {
	InterfaceEthernetInstanceAccessListAclID           InterfaceEthernetInstanceAccessList       `json:"access-list,omitempty"`
	InterfaceEthernetInstanceAction                    string                                    `json:"action,omitempty"`
	InterfaceEthernetInstanceAutoNegEnable             int                                       `json:"auto-neg-enable,omitempty"`
	InterfaceEthernetInstanceBfdAuthentication         InterfaceEthernetInstanceBfd              `json:"bfd,omitempty"`
	InterfaceEthernetInstanceCPUProcess                int                                       `json:"cpu-process,omitempty"`
	InterfaceEthernetInstanceCPUProcessDir             string                                    `json:"cpu-process-dir,omitempty"`
	InterfaceEthernetInstanceDdosOutside               InterfaceEthernetInstanceDdos             `json:"ddos,omitempty"`
	InterfaceEthernetInstanceDuplexity                 string                                    `json:"duplexity,omitempty"`
	InterfaceEthernetInstanceFecForcedOff              int                                       `json:"fec-forced-off,omitempty"`
	InterfaceEthernetInstanceFecForcedOn               int                                       `json:"fec-forced-on,omitempty"`
	InterfaceEthernetInstanceFlowControl               int                                       `json:"flow-control,omitempty"`
	InterfaceEthernetInstanceIPDhcp                    InterfaceEthernetInstanceIP               `json:"ip,omitempty"`
	InterfaceEthernetInstanceIcmpRateLimitNormal       InterfaceEthernetInstanceIcmpRateLimit    `json:"icmp-rate-limit,omitempty"`
	InterfaceEthernetInstanceIcmpv6RateLimitNormalV6   InterfaceEthernetInstanceIcmpv6RateLimit  `json:"icmpv6-rate-limit,omitempty"`
	InterfaceEthernetInstanceIfnum                     int                                       `json:"ifnum,omitempty"`
	InterfaceEthernetInstanceIpgBitTime                int                                       `json:"ipg-bit-time,omitempty"`
	InterfaceEthernetInstanceIpv6AddressList           InterfaceEthernetInstanceIpv6             `json:"ipv6,omitempty"`
	InterfaceEthernetInstanceIsisAuthentication        InterfaceEthernetInstanceIsis             `json:"isis,omitempty"`
	InterfaceEthernetInstanceL3VlanFwdDisable          int                                       `json:"l3-vlan-fwd-disable,omitempty"`
	InterfaceEthernetInstanceLldpEnableCfg             InterfaceEthernetInstanceLldp             `json:"lldp,omitempty"`
	InterfaceEthernetInstanceLoadInterval              int                                       `json:"load-interval,omitempty"`
	InterfaceEthernetInstanceLw4O6Outside              InterfaceEthernetInstanceLw4O6            `json:"lw-4o6,omitempty"`
	InterfaceEthernetInstanceMacLearning               string                                    `json:"mac-learning,omitempty"`
	InterfaceEthernetInstanceMapInside                 InterfaceEthernetInstanceMap              `json:"map,omitempty"`
	InterfaceEthernetInstanceMediaTypeCopper           int                                       `json:"media-type-copper,omitempty"`
	InterfaceEthernetInstanceMonitorListMonitor        []InterfaceEthernetInstanceMonitorList    `json:"monitor-list,omitempty"`
	InterfaceEthernetInstanceMtu                       int                                       `json:"mtu,omitempty"`
	InterfaceEthernetInstanceName                      string                                    `json:"name,omitempty"`
	InterfaceEthernetInstanceNptv6DomainList           InterfaceEthernetInstanceNptv6            `json:"nptv6,omitempty"`
	InterfaceEthernetInstancePacketCaptureTemplate     string                                    `json:"packet-capture-template,omitempty"`
	InterfaceEthernetInstancePingSweepDetection        string                                    `json:"ping-sweep-detection,omitempty"`
	InterfaceEthernetInstancePortBreakout              string                                    `json:"port-breakout,omitempty"`
	InterfaceEthernetInstancePortScanDetection         string                                    `json:"port-scan-detection,omitempty"`
	InterfaceEthernetInstanceRemoveVlanTag             int                                       `json:"remove-vlan-tag,omitempty"`
	InterfaceEthernetInstanceSamplingEnableCounters1   []InterfaceEthernetInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	InterfaceEthernetInstanceSpanningTreeAutoEdge      InterfaceEthernetInstanceSpanningTree     `json:"spanning-tree,omitempty"`
	InterfaceEthernetInstanceSpeed                     string                                    `json:"speed,omitempty"`
	InterfaceEthernetInstanceSpeedForced10G            int                                       `json:"speed-forced-10g,omitempty"`
	InterfaceEthernetInstanceSpeedForced1G             int                                       `json:"speed-forced-1g,omitempty"`
	InterfaceEthernetInstanceSpeedForced40G            int                                       `json:"speed-forced-40g,omitempty"`
	InterfaceEthernetInstanceTrafficDistributionMode   string                                    `json:"traffic-distribution-mode,omitempty"`
	InterfaceEthernetInstanceTrapSource                int                                       `json:"trap-source,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListTrunkNumber []InterfaceEthernetInstanceTrunkGroupList `json:"trunk-group-list,omitempty"`
	InterfaceEthernetInstanceUUID                      string                                    `json:"uuid,omitempty"`
	InterfaceEthernetInstanceUpdateL2Info              int                                       `json:"update-l2-info,omitempty"`
	InterfaceEthernetInstanceUserTag                   string                                    `json:"user-tag,omitempty"`
	InterfaceEthernetInstanceVirtualWire               int                                       `json:"virtual-wire,omitempty"`
	InterfaceEthernetInstanceVlanLearning              string                                    `json:"vlan-learning,omitempty"`
}

type InterfaceEthernetInstanceAccessList struct {
	InterfaceEthernetInstanceAccessListAclID   int    `json:"acl-id,omitempty"`
	InterfaceEthernetInstanceAccessListAclName string `json:"acl-name,omitempty"`
}

type InterfaceEthernetInstanceBfd struct {
	InterfaceEthernetInstanceBfdAuthenticationKeyID InterfaceEthernetInstanceBfdAuthentication `json:"authentication,omitempty"`
	InterfaceEthernetInstanceBfdDemand              int                                        `json:"demand,omitempty"`
	InterfaceEthernetInstanceBfdEcho                int                                        `json:"echo,omitempty"`
	InterfaceEthernetInstanceBfdIntervalCfgInterval InterfaceEthernetInstanceBfdIntervalCfg    `json:"interval-cfg,omitempty"`
	InterfaceEthernetInstanceBfdUUID                string                                     `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceDdos struct {
	InterfaceEthernetInstanceDdosInside  int    `json:"inside,omitempty"`
	InterfaceEthernetInstanceDdosOutside int    `json:"outside,omitempty"`
	InterfaceEthernetInstanceDdosUUID    string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIP struct {
	InterfaceEthernetInstanceIPAddressListIpv6Addr            []InterfaceEthernetInstanceIPAddressList       `json:"address-list,omitempty"`
	InterfaceEthernetInstanceIPAllowPromiscuousVip            int                                            `json:"allow-promiscuous-vip,omitempty"`
	InterfaceEthernetInstanceIPCacheSpoofingPort              int                                            `json:"cache-spoofing-port,omitempty"`
	InterfaceEthernetInstanceIPClient                         int                                            `json:"client,omitempty"`
	InterfaceEthernetInstanceIPDhcp                           int                                            `json:"dhcp,omitempty"`
	InterfaceEthernetInstanceIPGenerateMembershipQuery        int                                            `json:"generate-membership-query,omitempty"`
	InterfaceEthernetInstanceIPHelperAddressListHelperAddress []InterfaceEthernetInstanceIPHelperAddressList `json:"helper-address-list,omitempty"`
	InterfaceEthernetInstanceIPInside                         int                                            `json:"inside,omitempty"`
	InterfaceEthernetInstanceIPMaxRespTime                    int                                            `json:"max-resp-time,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobal                 InterfaceEthernetInstanceIPOspf                `json:"ospf,omitempty"`
	InterfaceEthernetInstanceIPOutside                        int                                            `json:"outside,omitempty"`
	InterfaceEthernetInstanceIPQueryInterval                  int                                            `json:"query-interval,omitempty"`
	InterfaceEthernetInstanceIPRipAuthentication              InterfaceEthernetInstanceIPRip                 `json:"rip,omitempty"`
	InterfaceEthernetInstanceIPRouterIsis                     InterfaceEthernetInstanceIPRouter              `json:"router,omitempty"`
	InterfaceEthernetInstanceIPServer                         int                                            `json:"server,omitempty"`
	InterfaceEthernetInstanceIPSlbPartitionRedirect           int                                            `json:"slb-partition-redirect,omitempty"`
	InterfaceEthernetInstanceIPStatefulFirewallInside         InterfaceEthernetInstanceIPStatefulFirewall    `json:"stateful-firewall,omitempty"`
	InterfaceEthernetInstanceIPSynCookie                      int                                            `json:"syn-cookie,omitempty"`
	InterfaceEthernetInstanceIPTTLIgnore                      int                                            `json:"ttl-ignore,omitempty"`
	InterfaceEthernetInstanceIPUUID                           string                                         `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIcmpRateLimit struct {
	InterfaceEthernetInstanceIcmpRateLimitLockup       int `json:"lockup,omitempty"`
	InterfaceEthernetInstanceIcmpRateLimitLockupPeriod int `json:"lockup-period,omitempty"`
	InterfaceEthernetInstanceIcmpRateLimitNormal       int `json:"normal,omitempty"`
}

type InterfaceEthernetInstanceIcmpv6RateLimit struct {
	InterfaceEthernetInstanceIcmpv6RateLimitLockupPeriodV6 int `json:"lockup-period-v6,omitempty"`
	InterfaceEthernetInstanceIcmpv6RateLimitLockupV6       int `json:"lockup-v6,omitempty"`
	InterfaceEthernetInstanceIcmpv6RateLimitNormalV6       int `json:"normal-v6,omitempty"`
}

type InterfaceEthernetInstanceIpv6 struct {
	InterfaceEthernetInstanceIpv6AccessListCfgV6AclName InterfaceEthernetInstanceIpv6AccessListCfg    `json:"access-list-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6AddressListIpv6Addr    []InterfaceEthernetInstanceIpv6AddressList    `json:"address-list,omitempty"`
	InterfaceEthernetInstanceIpv6Inside                 int                                           `json:"inside,omitempty"`
	InterfaceEthernetInstanceIpv6Ipv6Enable             int                                           `json:"ipv6-enable,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNetworkList        InterfaceEthernetInstanceIpv6Ospf             `json:"ospf,omitempty"`
	InterfaceEthernetInstanceIpv6Outside                int                                           `json:"outside,omitempty"`
	InterfaceEthernetInstanceIpv6RipSplitHorizonCfg     InterfaceEthernetInstanceIpv6Rip              `json:"rip,omitempty"`
	InterfaceEthernetInstanceIpv6RouterRipng            InterfaceEthernetInstanceIpv6Router           `json:"router,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverAction      InterfaceEthernetInstanceIpv6RouterAdver      `json:"router-adver,omitempty"`
	InterfaceEthernetInstanceIpv6StatefulFirewallInside InterfaceEthernetInstanceIpv6StatefulFirewall `json:"stateful-firewall,omitempty"`
	InterfaceEthernetInstanceIpv6TTLIgnore              int                                           `json:"ttl-ignore,omitempty"`
	InterfaceEthernetInstanceIpv6UUID                   string                                        `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIsis struct {
	InterfaceEthernetInstanceIsisAuthenticationSendOnlyList                   InterfaceEthernetInstanceIsisAuthentication             `json:"authentication,omitempty"`
	InterfaceEthernetInstanceIsisBfdCfgBfd                                    InterfaceEthernetInstanceIsisBfdCfg                     `json:"bfd-cfg,omitempty"`
	InterfaceEthernetInstanceIsisCircuitType                                  string                                                  `json:"circuit-type,omitempty"`
	InterfaceEthernetInstanceIsisCsnpIntervalListCsnpInterval                 []InterfaceEthernetInstanceIsisCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	InterfaceEthernetInstanceIsisHelloIntervalListHelloInterval               []InterfaceEthernetInstanceIsisHelloIntervalList        `json:"hello-interval-list,omitempty"`
	InterfaceEthernetInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal []InterfaceEthernetInstanceIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	InterfaceEthernetInstanceIsisHelloMultiplierListHelloMultiplier           []InterfaceEthernetInstanceIsisHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	InterfaceEthernetInstanceIsisLspInterval                                  int                                                     `json:"lsp-interval,omitempty"`
	InterfaceEthernetInstanceIsisMeshGroupValue                               InterfaceEthernetInstanceIsisMeshGroup                  `json:"mesh-group,omitempty"`
	InterfaceEthernetInstanceIsisMetricListMetric                             []InterfaceEthernetInstanceIsisMetricList               `json:"metric-list,omitempty"`
	InterfaceEthernetInstanceIsisNetwork                                      string                                                  `json:"network,omitempty"`
	InterfaceEthernetInstanceIsisPadding                                      int                                                     `json:"padding,omitempty"`
	InterfaceEthernetInstanceIsisPasswordListPassword                         []InterfaceEthernetInstanceIsisPasswordList             `json:"password-list,omitempty"`
	InterfaceEthernetInstanceIsisPriorityListPriority                         []InterfaceEthernetInstanceIsisPriorityList             `json:"priority-list,omitempty"`
	InterfaceEthernetInstanceIsisRetransmitInterval                           int                                                     `json:"retransmit-interval,omitempty"`
	InterfaceEthernetInstanceIsisUUID                                         string                                                  `json:"uuid,omitempty"`
	InterfaceEthernetInstanceIsisWideMetricListWideMetric                     []InterfaceEthernetInstanceIsisWideMetricList           `json:"wide-metric-list,omitempty"`
}

type InterfaceEthernetInstanceLldp struct {
	InterfaceEthernetInstanceLldpEnableCfgRtEnable           InterfaceEthernetInstanceLldpEnableCfg       `json:"enable-cfg,omitempty"`
	InterfaceEthernetInstanceLldpNotificationCfgNotification InterfaceEthernetInstanceLldpNotificationCfg `json:"notification-cfg,omitempty"`
	InterfaceEthernetInstanceLldpTxDot1CfgTxDot1Tlvs         InterfaceEthernetInstanceLldpTxDot1Cfg       `json:"tx-dot1-cfg,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgTxTlvs             InterfaceEthernetInstanceLldpTxTlvsCfg       `json:"tx-tlvs-cfg,omitempty"`
	InterfaceEthernetInstanceLldpUUID                        string                                       `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceLw4O6 struct {
	InterfaceEthernetInstanceLw4O6Inside  int    `json:"inside,omitempty"`
	InterfaceEthernetInstanceLw4O6Outside int    `json:"outside,omitempty"`
	InterfaceEthernetInstanceLw4O6UUID    string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceMap struct {
	InterfaceEthernetInstanceMapInside      int    `json:"inside,omitempty"`
	InterfaceEthernetInstanceMapMapTInside  int    `json:"map-t-inside,omitempty"`
	InterfaceEthernetInstanceMapMapTOutside int    `json:"map-t-outside,omitempty"`
	InterfaceEthernetInstanceMapOutside     int    `json:"outside,omitempty"`
	InterfaceEthernetInstanceMapUUID        string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceMonitorList struct {
	InterfaceEthernetInstanceMonitorListMirrorIndex int    `json:"mirror-index,omitempty"`
	InterfaceEthernetInstanceMonitorListMonitor     string `json:"monitor,omitempty"`
	InterfaceEthernetInstanceMonitorListMonitorVlan int    `json:"monitor-vlan,omitempty"`
}

type InterfaceEthernetInstanceNptv6 struct {
	InterfaceEthernetInstanceNptv6DomainListDomainName []InterfaceEthernetInstanceNptv6DomainList `json:"domain-list,omitempty"`
}

type InterfaceEthernetInstanceSamplingEnable struct {
	InterfaceEthernetInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type InterfaceEthernetInstanceSpanningTree struct {
	InterfaceEthernetInstanceSpanningTreeAdminEdge                 int                                                 `json:"admin-edge,omitempty"`
	InterfaceEthernetInstanceSpanningTreeAutoEdge                  int                                                 `json:"auto-edge,omitempty"`
	InterfaceEthernetInstanceSpanningTreeInstanceListInstanceStart []InterfaceEthernetInstanceSpanningTreeInstanceList `json:"instance-list,omitempty"`
	InterfaceEthernetInstanceSpanningTreePathCost                  int                                                 `json:"path-cost,omitempty"`
	InterfaceEthernetInstanceSpanningTreeUUID                      string                                              `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceTrunkGroupList struct {
	InterfaceEthernetInstanceTrunkGroupListAdminKey           int                                                   `json:"admin-key,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListMode               string                                                `json:"mode,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListPortPriority       int                                                   `json:"port-priority,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListTimeout            string                                                `json:"timeout,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListTrunkNumber        int                                                   `json:"trunk-number,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListType               string                                                `json:"type,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListUUID               string                                                `json:"uuid,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfgFast InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfg `json:"udld-timeout-cfg,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListUserTag            string                                                `json:"user-tag,omitempty"`
}

type InterfaceEthernetInstanceBfdAuthentication struct {
	InterfaceEthernetInstanceBfdAuthenticationKeyID    int    `json:"key-id,omitempty"`
	InterfaceEthernetInstanceBfdAuthenticationMethod   string `json:"method,omitempty"`
	InterfaceEthernetInstanceBfdAuthenticationPassword string `json:"password,omitempty"`
}

type InterfaceEthernetInstanceBfdIntervalCfg struct {
	InterfaceEthernetInstanceBfdIntervalCfgInterval   int `json:"interval,omitempty"`
	InterfaceEthernetInstanceBfdIntervalCfgMinRx      int `json:"min-rx,omitempty"`
	InterfaceEthernetInstanceBfdIntervalCfgMultiplier int `json:"multiplier,omitempty"`
}

type InterfaceEthernetInstanceIPAddressList struct {
	InterfaceEthernetInstanceIPAddressListAddressType string `json:"address-type,omitempty"`
	InterfaceEthernetInstanceIPAddressListIpv6Addr    string `json:"ipv6-addr,omitempty"`
}

type InterfaceEthernetInstanceIPHelperAddressList struct {
	InterfaceEthernetInstanceIPHelperAddressListHelperAddress string `json:"helper-address,omitempty"`
}

type InterfaceEthernetInstanceIPOspf struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfg InterfaceEthernetInstanceIPOspfOspfGlobal   `json:"ospf-global,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListIPAddr            []InterfaceEthernetInstanceIPOspfOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceEthernetInstanceIPRip struct {
	InterfaceEthernetInstanceIPRipAuthenticationStr    InterfaceEthernetInstanceIPRipAuthentication  `json:"authentication,omitempty"`
	InterfaceEthernetInstanceIPRipReceiveCfgReceive    InterfaceEthernetInstanceIPRipReceiveCfg      `json:"receive-cfg,omitempty"`
	InterfaceEthernetInstanceIPRipReceivePacket        int                                           `json:"receive-packet,omitempty"`
	InterfaceEthernetInstanceIPRipSendCfgSend          InterfaceEthernetInstanceIPRipSendCfg         `json:"send-cfg,omitempty"`
	InterfaceEthernetInstanceIPRipSendPacket           int                                           `json:"send-packet,omitempty"`
	InterfaceEthernetInstanceIPRipSplitHorizonCfgState InterfaceEthernetInstanceIPRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceEthernetInstanceIPRipUUID                 string                                        `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIPRouter struct {
	InterfaceEthernetInstanceIPRouterIsisTag InterfaceEthernetInstanceIPRouterIsis `json:"isis,omitempty"`
}

type InterfaceEthernetInstanceIPStatefulFirewall struct {
	InterfaceEthernetInstanceIPStatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceEthernetInstanceIPStatefulFirewallAclID      int    `json:"acl-id,omitempty"`
	InterfaceEthernetInstanceIPStatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceEthernetInstanceIPStatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceEthernetInstanceIPStatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceEthernetInstanceIPStatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6AccessListCfg struct {
	InterfaceEthernetInstanceIpv6AccessListCfgInbound   int    `json:"inbound,omitempty"`
	InterfaceEthernetInstanceIpv6AccessListCfgV6AclName string `json:"v6-acl-name,omitempty"`
}

type InterfaceEthernetInstanceIpv6AddressList struct {
	InterfaceEthernetInstanceIpv6AddressListAddressType string `json:"address-type,omitempty"`
	InterfaceEthernetInstanceIpv6AddressListIpv6Addr    string `json:"ipv6-addr,omitempty"`
}

type InterfaceEthernetInstanceIpv6Ospf struct {
	InterfaceEthernetInstanceIpv6OspfBfd                                     int                                                      `json:"bfd,omitempty"`
	InterfaceEthernetInstanceIpv6OspfCostCfgCost                             []InterfaceEthernetInstanceIpv6OspfCostCfg               `json:"cost-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgDeadInterval             []InterfaceEthernetInstanceIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfDisable                                 int                                                      `json:"disable,omitempty"`
	InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgHelloInterval           []InterfaceEthernetInstanceIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgMtuIgnore                   []InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighbor                     []InterfaceEthernetInstanceIpv6OspfNeighborCfg           `json:"neighbor-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNetworkListBroadcastType                []InterfaceEthernetInstanceIpv6OspfNetworkList           `json:"network-list,omitempty"`
	InterfaceEthernetInstanceIpv6OspfPriorityCfgPriority                     []InterfaceEthernetInstanceIpv6OspfPriorityCfg           `json:"priority-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval []InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgTransmitDelay           []InterfaceEthernetInstanceIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6OspfUUID                                    string                                                   `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6Rip struct {
	InterfaceEthernetInstanceIpv6RipSplitHorizonCfgState InterfaceEthernetInstanceIpv6RipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceEthernetInstanceIpv6RipUUID                 string                                          `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6Router struct {
	InterfaceEthernetInstanceIpv6RouterIsisTag      InterfaceEthernetInstanceIpv6RouterIsis  `json:"isis,omitempty"`
	InterfaceEthernetInstanceIpv6RouterOspfAreaList InterfaceEthernetInstanceIpv6RouterOspf  `json:"ospf,omitempty"`
	InterfaceEthernetInstanceIpv6RouterRipngRip     InterfaceEthernetInstanceIpv6RouterRipng `json:"ripng,omitempty"`
}

type InterfaceEthernetInstanceIpv6RouterAdver struct {
	InterfaceEthernetInstanceIpv6RouterAdverAction                   string                                               `json:"action,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverAdverMtu                 int                                                  `json:"adver-mtu,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverAdverMtuDisable          int                                                  `json:"adver-mtu-disable,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverAdverVrid                int                                                  `json:"adver-vrid,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverAdverVridDefault         int                                                  `json:"adver-vrid-default,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverDefaultLifetime          int                                                  `json:"default-lifetime,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverFloatingIP               string                                               `json:"floating-ip,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverFloatingIPDefaultVrid    string                                               `json:"floating-ip-default-vrid,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverHopLimit                 int                                                  `json:"hop-limit,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverManagedConfigAction      string                                               `json:"managed-config-action,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverMaxInterval              int                                                  `json:"max-interval,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverMinInterval              int                                                  `json:"min-interval,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverOtherConfigAction        string                                               `json:"other-config-action,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverPrefixListPrefix         []InterfaceEthernetInstanceIpv6RouterAdverPrefixList `json:"prefix-list,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverRateLimit                int                                                  `json:"rate-limit,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverReachableTime            int                                                  `json:"reachable-time,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverRetransmitTimer          int                                                  `json:"retransmit-timer,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverUseFloatingIP            int                                                  `json:"use-floating-ip,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverUseFloatingIPDefaultVrid int                                                  `json:"use-floating-ip-default-vrid,omitempty"`
}

type InterfaceEthernetInstanceIpv6StatefulFirewall struct {
	InterfaceEthernetInstanceIpv6StatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceEthernetInstanceIpv6StatefulFirewallAclName    string `json:"acl-name,omitempty"`
	InterfaceEthernetInstanceIpv6StatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceEthernetInstanceIpv6StatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceEthernetInstanceIpv6StatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceEthernetInstanceIpv6StatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIsisAuthentication struct {
	InterfaceEthernetInstanceIsisAuthenticationKeyChainListKeyChain []InterfaceEthernetInstanceIsisAuthenticationKeyChainList `json:"key-chain-list,omitempty"`
	InterfaceEthernetInstanceIsisAuthenticationModeListMode         []InterfaceEthernetInstanceIsisAuthenticationModeList     `json:"mode-list,omitempty"`
	InterfaceEthernetInstanceIsisAuthenticationSendOnlyListSendOnly []InterfaceEthernetInstanceIsisAuthenticationSendOnlyList `json:"send-only-list,omitempty"`
}

type InterfaceEthernetInstanceIsisBfdCfg struct {
	InterfaceEthernetInstanceIsisBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceEthernetInstanceIsisBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceEthernetInstanceIsisCsnpIntervalList struct {
	InterfaceEthernetInstanceIsisCsnpIntervalListCsnpInterval int    `json:"csnp-interval,omitempty"`
	InterfaceEthernetInstanceIsisCsnpIntervalListLevel        string `json:"level,omitempty"`
}

type InterfaceEthernetInstanceIsisHelloIntervalList struct {
	InterfaceEthernetInstanceIsisHelloIntervalListHelloInterval int    `json:"hello-interval,omitempty"`
	InterfaceEthernetInstanceIsisHelloIntervalListLevel         string `json:"level,omitempty"`
}

type InterfaceEthernetInstanceIsisHelloIntervalMinimalList struct {
	InterfaceEthernetInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	InterfaceEthernetInstanceIsisHelloIntervalMinimalListLevel                string `json:"level,omitempty"`
}

type InterfaceEthernetInstanceIsisHelloMultiplierList struct {
	InterfaceEthernetInstanceIsisHelloMultiplierListHelloMultiplier int    `json:"hello-multiplier,omitempty"`
	InterfaceEthernetInstanceIsisHelloMultiplierListLevel           string `json:"level,omitempty"`
}

type InterfaceEthernetInstanceIsisMeshGroup struct {
	InterfaceEthernetInstanceIsisMeshGroupBlocked int `json:"blocked,omitempty"`
	InterfaceEthernetInstanceIsisMeshGroupValue   int `json:"value,omitempty"`
}

type InterfaceEthernetInstanceIsisMetricList struct {
	InterfaceEthernetInstanceIsisMetricListLevel  string `json:"level,omitempty"`
	InterfaceEthernetInstanceIsisMetricListMetric int    `json:"metric,omitempty"`
}

type InterfaceEthernetInstanceIsisPasswordList struct {
	InterfaceEthernetInstanceIsisPasswordListLevel    string `json:"level,omitempty"`
	InterfaceEthernetInstanceIsisPasswordListPassword string `json:"password,omitempty"`
}

type InterfaceEthernetInstanceIsisPriorityList struct {
	InterfaceEthernetInstanceIsisPriorityListLevel    string `json:"level,omitempty"`
	InterfaceEthernetInstanceIsisPriorityListPriority int    `json:"priority,omitempty"`
}

type InterfaceEthernetInstanceIsisWideMetricList struct {
	InterfaceEthernetInstanceIsisWideMetricListLevel      string `json:"level,omitempty"`
	InterfaceEthernetInstanceIsisWideMetricListWideMetric int    `json:"wide-metric,omitempty"`
}

type InterfaceEthernetInstanceLldpEnableCfg struct {
	InterfaceEthernetInstanceLldpEnableCfgRtEnable int `json:"rt-enable,omitempty"`
	InterfaceEthernetInstanceLldpEnableCfgRx       int `json:"rx,omitempty"`
	InterfaceEthernetInstanceLldpEnableCfgTx       int `json:"tx,omitempty"`
}

type InterfaceEthernetInstanceLldpNotificationCfg struct {
	InterfaceEthernetInstanceLldpNotificationCfgNotifEnable  int `json:"notif-enable,omitempty"`
	InterfaceEthernetInstanceLldpNotificationCfgNotification int `json:"notification,omitempty"`
}

type InterfaceEthernetInstanceLldpTxDot1Cfg struct {
	InterfaceEthernetInstanceLldpTxDot1CfgLinkAggregation int `json:"link-aggregation,omitempty"`
	InterfaceEthernetInstanceLldpTxDot1CfgTxDot1Tlvs      int `json:"tx-dot1-tlvs,omitempty"`
	InterfaceEthernetInstanceLldpTxDot1CfgVlan            int `json:"vlan,omitempty"`
}

type InterfaceEthernetInstanceLldpTxTlvsCfg struct {
	InterfaceEthernetInstanceLldpTxTlvsCfgExclude            int `json:"exclude,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgManagementAddress  int `json:"management-address,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgPortDescription    int `json:"port-description,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgSystemCapabilities int `json:"system-capabilities,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgSystemDescription  int `json:"system-description,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgSystemName         int `json:"system-name,omitempty"`
	InterfaceEthernetInstanceLldpTxTlvsCfgTxTlvs             int `json:"tx-tlvs,omitempty"`
}

type InterfaceEthernetInstanceNptv6DomainList struct {
	InterfaceEthernetInstanceNptv6DomainListBindType   string `json:"bind-type,omitempty"`
	InterfaceEthernetInstanceNptv6DomainListDomainName string `json:"domain-name,omitempty"`
	InterfaceEthernetInstanceNptv6DomainListUUID       string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceSpanningTreeInstanceList struct {
	InterfaceEthernetInstanceSpanningTreeInstanceListInstanceStart int `json:"instance-start,omitempty"`
	InterfaceEthernetInstanceSpanningTreeInstanceListMstpPathCost  int `json:"mstp-path-cost,omitempty"`
}

type InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfg struct {
	InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfgFast int `json:"fast,omitempty"`
	InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfgSlow int `json:"slow,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfGlobal struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication  InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationKey                string                                                      `json:"authentication-key,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfgBfd                        InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfg             `json:"bfd-cfg,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalCost                             int                                                         `json:"cost,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter  InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalDeadInterval                     int                                                         `json:"dead-interval,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalDisable                          string                                                      `json:"disable,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalHelloInterval                    int                                                         `json:"hello-interval,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey []InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalMtu                              int                                                         `json:"mtu,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalMtuIgnore                        int                                                         `json:"mtu-ignore,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalNetworkBroadcast                 InterfaceEthernetInstanceIPOspfOspfGlobalNetwork            `json:"network,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalPriority                         int                                                         `json:"priority,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalRetransmitInterval               int                                                         `json:"retransmit-interval,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalTransmitDelay                    int                                                         `json:"transmit-delay,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalUUID                             string                                                      `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfIPList struct {
	InterfaceEthernetInstanceIPOspfOspfIPListAuthentication                   int                                                         `json:"authentication,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListAuthenticationKey                string                                                      `json:"authentication-key,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListCost                             int                                                         `json:"cost,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListDatabaseFilter                   string                                                      `json:"database-filter,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListDeadInterval                     int                                                         `json:"dead-interval,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListHelloInterval                    int                                                         `json:"hello-interval,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListIPAddr                           string                                                      `json:"ip-addr,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey []InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListMtuIgnore                        int                                                         `json:"mtu-ignore,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListOut                              int                                                         `json:"out,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListPriority                         int                                                         `json:"priority,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListRetransmitInterval               int                                                         `json:"retransmit-interval,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListTransmitDelay                    int                                                         `json:"transmit-delay,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListUUID                             string                                                      `json:"uuid,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListValue                            string                                                      `json:"value,omitempty"`
}

type InterfaceEthernetInstanceIPRipAuthentication struct {
	InterfaceEthernetInstanceIPRipAuthenticationKeyChainKeyChain InterfaceEthernetInstanceIPRipAuthenticationKeyChain `json:"key-chain,omitempty"`
	InterfaceEthernetInstanceIPRipAuthenticationModeMode         InterfaceEthernetInstanceIPRipAuthenticationMode     `json:"mode,omitempty"`
	InterfaceEthernetInstanceIPRipAuthenticationStrString        InterfaceEthernetInstanceIPRipAuthenticationStr      `json:"str,omitempty"`
}

type InterfaceEthernetInstanceIPRipReceiveCfg struct {
	InterfaceEthernetInstanceIPRipReceiveCfgReceive int    `json:"receive,omitempty"`
	InterfaceEthernetInstanceIPRipReceiveCfgVersion string `json:"version,omitempty"`
}

type InterfaceEthernetInstanceIPRipSendCfg struct {
	InterfaceEthernetInstanceIPRipSendCfgSend    int    `json:"send,omitempty"`
	InterfaceEthernetInstanceIPRipSendCfgVersion string `json:"version,omitempty"`
}

type InterfaceEthernetInstanceIPRipSplitHorizonCfg struct {
	InterfaceEthernetInstanceIPRipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceEthernetInstanceIPRouterIsis struct {
	InterfaceEthernetInstanceIPRouterIsisTag  string `json:"tag,omitempty"`
	InterfaceEthernetInstanceIPRouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfCostCfg struct {
	InterfaceEthernetInstanceIpv6OspfCostCfgCost       int `json:"cost,omitempty"`
	InterfaceEthernetInstanceIpv6OspfCostCfgInstanceID int `json:"instance-id,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfDeadIntervalCfg struct {
	InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgDeadInterval int `json:"dead-interval,omitempty"`
	InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgInstanceID   int `json:"instance-id,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfHelloIntervalCfg struct {
	InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgHelloInterval int `json:"hello-interval,omitempty"`
	InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgInstanceID    int `json:"instance-id,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfg struct {
	InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgMtuIgnore  int `json:"mtu-ignore,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfNeighborCfg struct {
	InterfaceEthernetInstanceIpv6OspfNeighborCfgNeigInst             int    `json:"neig-inst,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighbor             string `json:"neighbor,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighborCost         int    `json:"neighbor-cost,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighborPriority     int    `json:"neighbor-priority,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfNetworkList struct {
	InterfaceEthernetInstanceIpv6OspfNetworkListBroadcastType     string `json:"broadcast-type,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNetworkListNetworkInstanceID int    `json:"network-instance-id,omitempty"`
	InterfaceEthernetInstanceIpv6OspfNetworkListP2MpNbma          int    `json:"p2mp-nbma,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfPriorityCfg struct {
	InterfaceEthernetInstanceIpv6OspfPriorityCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceEthernetInstanceIpv6OspfPriorityCfgPriority   int `json:"priority,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfg struct {
	InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgInstanceID         int `json:"instance-id,omitempty"`
	InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval int `json:"retransmit-interval,omitempty"`
}

type InterfaceEthernetInstanceIpv6OspfTransmitDelayCfg struct {
	InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgInstanceID    int `json:"instance-id,omitempty"`
	InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgTransmitDelay int `json:"transmit-delay,omitempty"`
}

type InterfaceEthernetInstanceIpv6RipSplitHorizonCfg struct {
	InterfaceEthernetInstanceIpv6RipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceEthernetInstanceIpv6RouterIsis struct {
	InterfaceEthernetInstanceIpv6RouterIsisTag  string `json:"tag,omitempty"`
	InterfaceEthernetInstanceIpv6RouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6RouterOspf struct {
	InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDNum []InterfaceEthernetInstanceIpv6RouterOspfAreaList `json:"area-list,omitempty"`
	InterfaceEthernetInstanceIpv6RouterOspfUUID              string                                            `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6RouterRipng struct {
	InterfaceEthernetInstanceIpv6RouterRipngRip  int    `json:"rip,omitempty"`
	InterfaceEthernetInstanceIpv6RouterRipngUUID string `json:"uuid,omitempty"`
}

type InterfaceEthernetInstanceIpv6RouterAdverPrefixList struct {
	InterfaceEthernetInstanceIpv6RouterAdverPrefixListNotAutonomous     int    `json:"not-autonomous,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverPrefixListNotOnLink         int    `json:"not-on-link,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverPrefixListPreferredLifetime int    `json:"preferred-lifetime,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverPrefixListPrefix            string `json:"prefix,omitempty"`
	InterfaceEthernetInstanceIpv6RouterAdverPrefixListValidLifetime     int    `json:"valid-lifetime,omitempty"`
}

type InterfaceEthernetInstanceIsisAuthenticationKeyChainList struct {
	InterfaceEthernetInstanceIsisAuthenticationKeyChainListKeyChain string `json:"key-chain,omitempty"`
	InterfaceEthernetInstanceIsisAuthenticationKeyChainListLevel    string `json:"level,omitempty"`
}

type InterfaceEthernetInstanceIsisAuthenticationModeList struct {
	InterfaceEthernetInstanceIsisAuthenticationModeListLevel string `json:"level,omitempty"`
	InterfaceEthernetInstanceIsisAuthenticationModeListMode  string `json:"mode,omitempty"`
}

type InterfaceEthernetInstanceIsisAuthenticationSendOnlyList struct {
	InterfaceEthernetInstanceIsisAuthenticationSendOnlyListLevel    string `json:"level,omitempty"`
	InterfaceEthernetInstanceIsisAuthenticationSendOnlyListSendOnly int    `json:"send-only,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfg struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication int    `json:"authentication,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfgValue          string `json:"value,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfg struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfg struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter string `json:"database-filter,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfgOut            int    `json:"out,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfg struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfGlobalNetwork struct {
	InterfaceEthernetInstanceIPOspfOspfGlobalNetworkBroadcast         int `json:"broadcast,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalNetworkNonBroadcast      int `json:"non-broadcast,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalNetworkP2MpNbma          int `json:"p2mp-nbma,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalNetworkPointToMultipoint int `json:"point-to-multipoint,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfGlobalNetworkPointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfg struct {
	InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceEthernetInstanceIPRipAuthenticationKeyChain struct {
	InterfaceEthernetInstanceIPRipAuthenticationKeyChainKeyChain string `json:"key-chain,omitempty"`
}

type InterfaceEthernetInstanceIPRipAuthenticationMode struct {
	InterfaceEthernetInstanceIPRipAuthenticationModeMode string `json:"mode,omitempty"`
}

type InterfaceEthernetInstanceIPRipAuthenticationStr struct {
	InterfaceEthernetInstanceIPRipAuthenticationStrString string `json:"string,omitempty"`
}

type InterfaceEthernetInstanceIpv6RouterOspfAreaList struct {
	InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDAddr string `json:"area-id-addr,omitempty"`
	InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDNum  int    `json:"area-id-num,omitempty"`
	InterfaceEthernetInstanceIpv6RouterOspfAreaListInstanceID int    `json:"instance-id,omitempty"`
	InterfaceEthernetInstanceIpv6RouterOspfAreaListTag        string `json:"tag,omitempty"`
}

func PostInterfaceEthernetObject(id string, inst InterfaceEthernet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceEthernet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func PostInterfaceEthernetInstance(id string, ifnum string, inst InterfaceEthernet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+ifnum, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceEthernet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceEthernet(id string, name1 string, host string) (*InterfaceEthernet, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernet")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceEthernet", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutInterfaceEthernet(id string, name1 string, inst InterfaceEthernet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceEthernet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/ethernet/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutInterfaceEthernet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteInterfaceEthernet(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceEthernet")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/ethernet/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceEthernet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteInterfaceEthernet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
