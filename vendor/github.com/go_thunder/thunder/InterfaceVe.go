package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceVe struct {
	InterfaceVeInstanceIfnum InterfaceVeInstance `json:"ve,omitempty"`
}

type InterfaceVeInstance struct {
	InterfaceVeInstanceAccessListAclID         InterfaceVeInstanceAccessList       `json:"access-list,omitempty"`
	InterfaceVeInstanceAction                  string                              `json:"action,omitempty"`
	InterfaceVeInstanceBfdAuthentication       InterfaceVeInstanceBfd              `json:"bfd,omitempty"`
	InterfaceVeInstanceDdosOutside             InterfaceVeInstanceDdos             `json:"ddos,omitempty"`
	InterfaceVeInstanceIPDhcp                  InterfaceVeInstanceIP               `json:"ip,omitempty"`
	InterfaceVeInstanceIcmpRateLimitNormal     InterfaceVeInstanceIcmpRateLimit    `json:"icmp-rate-limit,omitempty"`
	InterfaceVeInstanceIcmpv6RateLimitNormalV6 InterfaceVeInstanceIcmpv6RateLimit  `json:"icmpv6-rate-limit,omitempty"`
	InterfaceVeInstanceIfnum                   int                                 `json:"ifnum,omitempty"`
	InterfaceVeInstanceIpv6AddressList         InterfaceVeInstanceIpv6             `json:"ipv6,omitempty"`
	InterfaceVeInstanceIsisAuthentication      InterfaceVeInstanceIsis             `json:"isis,omitempty"`
	InterfaceVeInstanceL3VlanFwdDisable        int                                 `json:"l3-vlan-fwd-disable,omitempty"`
	InterfaceVeInstanceLw4O6Outside            InterfaceVeInstanceLw4O6            `json:"lw-4o6,omitempty"`
	InterfaceVeInstanceMapInside               InterfaceVeInstanceMap              `json:"map,omitempty"`
	InterfaceVeInstanceMtu                     int                                 `json:"mtu,omitempty"`
	InterfaceVeInstanceName                    string                              `json:"name,omitempty"`
	InterfaceVeInstanceNptv6DomainList         InterfaceVeInstanceNptv6            `json:"nptv6,omitempty"`
	InterfaceVeInstancePingSweepDetection      string                              `json:"ping-sweep-detection,omitempty"`
	InterfaceVeInstancePortScanDetection       string                              `json:"port-scan-detection,omitempty"`
	InterfaceVeInstanceSamplingEnableCounters1 []InterfaceVeInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	InterfaceVeInstanceTrapSource              int                                 `json:"trap-source,omitempty"`
	InterfaceVeInstanceUUID                    string                              `json:"uuid,omitempty"`
	InterfaceVeInstanceUserTag                 string                              `json:"user-tag,omitempty"`
}

type InterfaceVeInstanceAccessList struct {
	InterfaceVeInstanceAccessListAclID   int    `json:"acl-id,omitempty"`
	InterfaceVeInstanceAccessListAclName string `json:"acl-name,omitempty"`
}

type InterfaceVeInstanceBfd struct {
	InterfaceVeInstanceBfdAuthenticationKeyID InterfaceVeInstanceBfdAuthentication `json:"authentication,omitempty"`
	InterfaceVeInstanceBfdDemand              int                                  `json:"demand,omitempty"`
	InterfaceVeInstanceBfdEcho                int                                  `json:"echo,omitempty"`
	InterfaceVeInstanceBfdIntervalCfgInterval InterfaceVeInstanceBfdIntervalCfg    `json:"interval-cfg,omitempty"`
	InterfaceVeInstanceBfdUUID                string                               `json:"uuid,omitempty"`
}

type InterfaceVeInstanceDdos struct {
	InterfaceVeInstanceDdosInside  int    `json:"inside,omitempty"`
	InterfaceVeInstanceDdosOutside int    `json:"outside,omitempty"`
	InterfaceVeInstanceDdosUUID    string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIP struct {
	InterfaceVeInstanceIPAddressListIpv6Addr            []InterfaceVeInstanceIPAddressList       `json:"address-list,omitempty"`
	InterfaceVeInstanceIPAllowPromiscuousVip            int                                      `json:"allow-promiscuous-vip,omitempty"`
	InterfaceVeInstanceIPClient                         int                                      `json:"client,omitempty"`
	InterfaceVeInstanceIPDhcp                           int                                      `json:"dhcp,omitempty"`
	InterfaceVeInstanceIPGenerateMembershipQuery        int                                      `json:"generate-membership-query,omitempty"`
	InterfaceVeInstanceIPHelperAddressListHelperAddress []InterfaceVeInstanceIPHelperAddressList `json:"helper-address-list,omitempty"`
	InterfaceVeInstanceIPInside                         int                                      `json:"inside,omitempty"`
	InterfaceVeInstanceIPMaxRespTime                    int                                      `json:"max-resp-time,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobal                 InterfaceVeInstanceIPOspf                `json:"ospf,omitempty"`
	InterfaceVeInstanceIPOutside                        int                                      `json:"outside,omitempty"`
	InterfaceVeInstanceIPQueryInterval                  int                                      `json:"query-interval,omitempty"`
	InterfaceVeInstanceIPRipAuthentication              InterfaceVeInstanceIPRip                 `json:"rip,omitempty"`
	InterfaceVeInstanceIPRouterIsis                     InterfaceVeInstanceIPRouter              `json:"router,omitempty"`
	InterfaceVeInstanceIPServer                         int                                      `json:"server,omitempty"`
	InterfaceVeInstanceIPSlbPartitionRedirect           int                                      `json:"slb-partition-redirect,omitempty"`
	InterfaceVeInstanceIPStatefulFirewallInside         InterfaceVeInstanceIPStatefulFirewall    `json:"stateful-firewall,omitempty"`
	InterfaceVeInstanceIPSynCookie                      int                                      `json:"syn-cookie,omitempty"`
	InterfaceVeInstanceIPTTLIgnore                      int                                      `json:"ttl-ignore,omitempty"`
	InterfaceVeInstanceIPUUID                           string                                   `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIcmpRateLimit struct {
	InterfaceVeInstanceIcmpRateLimitLockup       int `json:"lockup,omitempty"`
	InterfaceVeInstanceIcmpRateLimitLockupPeriod int `json:"lockup-period,omitempty"`
	InterfaceVeInstanceIcmpRateLimitNormal       int `json:"normal,omitempty"`
}

type InterfaceVeInstanceIcmpv6RateLimit struct {
	InterfaceVeInstanceIcmpv6RateLimitLockupPeriodV6 int `json:"lockup-period-v6,omitempty"`
	InterfaceVeInstanceIcmpv6RateLimitLockupV6       int `json:"lockup-v6,omitempty"`
	InterfaceVeInstanceIcmpv6RateLimitNormalV6       int `json:"normal-v6,omitempty"`
}

type InterfaceVeInstanceIpv6 struct {
	InterfaceVeInstanceIpv6AddressListIpv6Addr    []InterfaceVeInstanceIpv6AddressList    `json:"address-list,omitempty"`
	InterfaceVeInstanceIpv6Inbound                int                                     `json:"inbound,omitempty"`
	InterfaceVeInstanceIpv6Inside                 int                                     `json:"inside,omitempty"`
	InterfaceVeInstanceIpv6Ipv6Enable             int                                     `json:"ipv6-enable,omitempty"`
	InterfaceVeInstanceIpv6OspfNetworkList        InterfaceVeInstanceIpv6Ospf             `json:"ospf,omitempty"`
	InterfaceVeInstanceIpv6Outside                int                                     `json:"outside,omitempty"`
	InterfaceVeInstanceIpv6RipSplitHorizonCfg     InterfaceVeInstanceIpv6Rip              `json:"rip,omitempty"`
	InterfaceVeInstanceIpv6RouterRipng            InterfaceVeInstanceIpv6Router           `json:"router,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverAction      InterfaceVeInstanceIpv6RouterAdver      `json:"router-adver,omitempty"`
	InterfaceVeInstanceIpv6StatefulFirewallInside InterfaceVeInstanceIpv6StatefulFirewall `json:"stateful-firewall,omitempty"`
	InterfaceVeInstanceIpv6TTLIgnore              int                                     `json:"ttl-ignore,omitempty"`
	InterfaceVeInstanceIpv6UUID                   string                                  `json:"uuid,omitempty"`
	InterfaceVeInstanceIpv6V6AclName              string                                  `json:"v6-acl-name,omitempty"`
}

type InterfaceVeInstanceIsis struct {
	InterfaceVeInstanceIsisAuthenticationSendOnlyList                   InterfaceVeInstanceIsisAuthentication             `json:"authentication,omitempty"`
	InterfaceVeInstanceIsisBfdCfgBfd                                    InterfaceVeInstanceIsisBfdCfg                     `json:"bfd-cfg,omitempty"`
	InterfaceVeInstanceIsisCircuitType                                  string                                            `json:"circuit-type,omitempty"`
	InterfaceVeInstanceIsisCsnpIntervalListCsnpInterval                 []InterfaceVeInstanceIsisCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	InterfaceVeInstanceIsisHelloIntervalListHelloInterval               []InterfaceVeInstanceIsisHelloIntervalList        `json:"hello-interval-list,omitempty"`
	InterfaceVeInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal []InterfaceVeInstanceIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	InterfaceVeInstanceIsisHelloMultiplierListHelloMultiplier           []InterfaceVeInstanceIsisHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	InterfaceVeInstanceIsisLspInterval                                  int                                               `json:"lsp-interval,omitempty"`
	InterfaceVeInstanceIsisMeshGroupValue                               InterfaceVeInstanceIsisMeshGroup                  `json:"mesh-group,omitempty"`
	InterfaceVeInstanceIsisMetricListMetric                             []InterfaceVeInstanceIsisMetricList               `json:"metric-list,omitempty"`
	InterfaceVeInstanceIsisNetwork                                      string                                            `json:"network,omitempty"`
	InterfaceVeInstanceIsisPadding                                      int                                               `json:"padding,omitempty"`
	InterfaceVeInstanceIsisPasswordListPassword                         []InterfaceVeInstanceIsisPasswordList             `json:"password-list,omitempty"`
	InterfaceVeInstanceIsisPriorityListPriority                         []InterfaceVeInstanceIsisPriorityList             `json:"priority-list,omitempty"`
	InterfaceVeInstanceIsisRetransmitInterval                           int                                               `json:"retransmit-interval,omitempty"`
	InterfaceVeInstanceIsisUUID                                         string                                            `json:"uuid,omitempty"`
	InterfaceVeInstanceIsisWideMetricListWideMetric                     []InterfaceVeInstanceIsisWideMetricList           `json:"wide-metric-list,omitempty"`
}

type InterfaceVeInstanceLw4O6 struct {
	InterfaceVeInstanceLw4O6Inside  int    `json:"inside,omitempty"`
	InterfaceVeInstanceLw4O6Outside int    `json:"outside,omitempty"`
	InterfaceVeInstanceLw4O6UUID    string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceMap struct {
	InterfaceVeInstanceMapInside      int    `json:"inside,omitempty"`
	InterfaceVeInstanceMapMapTInside  int    `json:"map-t-inside,omitempty"`
	InterfaceVeInstanceMapMapTOutside int    `json:"map-t-outside,omitempty"`
	InterfaceVeInstanceMapOutside     int    `json:"outside,omitempty"`
	InterfaceVeInstanceMapUUID        string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceNptv6 struct {
	InterfaceVeInstanceNptv6DomainListDomainName []InterfaceVeInstanceNptv6DomainList `json:"domain-list,omitempty"`
}

type InterfaceVeInstanceSamplingEnable struct {
	InterfaceVeInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type InterfaceVeInstanceBfdAuthentication struct {
	InterfaceVeInstanceBfdAuthenticationKeyID    int    `json:"key-id,omitempty"`
	InterfaceVeInstanceBfdAuthenticationMethod   string `json:"method,omitempty"`
	InterfaceVeInstanceBfdAuthenticationPassword string `json:"password,omitempty"`
}

type InterfaceVeInstanceBfdIntervalCfg struct {
	InterfaceVeInstanceBfdIntervalCfgInterval   int `json:"interval,omitempty"`
	InterfaceVeInstanceBfdIntervalCfgMinRx      int `json:"min-rx,omitempty"`
	InterfaceVeInstanceBfdIntervalCfgMultiplier int `json:"multiplier,omitempty"`
}

type InterfaceVeInstanceIPAddressList struct {
	InterfaceVeInstanceIPAddressListAddressType string `json:"address-type,omitempty"`
	InterfaceVeInstanceIPAddressListIpv6Addr    string `json:"ipv6-addr,omitempty"`
}

type InterfaceVeInstanceIPHelperAddressList struct {
	InterfaceVeInstanceIPHelperAddressListHelperAddress string `json:"helper-address,omitempty"`
}

type InterfaceVeInstanceIPOspf struct {
	InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfg InterfaceVeInstanceIPOspfOspfGlobal   `json:"ospf-global,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListIPAddr            []InterfaceVeInstanceIPOspfOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceVeInstanceIPRip struct {
	InterfaceVeInstanceIPRipAuthenticationStr    InterfaceVeInstanceIPRipAuthentication  `json:"authentication,omitempty"`
	InterfaceVeInstanceIPRipReceiveCfgReceive    InterfaceVeInstanceIPRipReceiveCfg      `json:"receive-cfg,omitempty"`
	InterfaceVeInstanceIPRipReceivePacket        int                                     `json:"receive-packet,omitempty"`
	InterfaceVeInstanceIPRipSendCfgSend          InterfaceVeInstanceIPRipSendCfg         `json:"send-cfg,omitempty"`
	InterfaceVeInstanceIPRipSendPacket           int                                     `json:"send-packet,omitempty"`
	InterfaceVeInstanceIPRipSplitHorizonCfgState InterfaceVeInstanceIPRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceVeInstanceIPRipUUID                 string                                  `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIPRouter struct {
	InterfaceVeInstanceIPRouterIsisTag InterfaceVeInstanceIPRouterIsis `json:"isis,omitempty"`
}

type InterfaceVeInstanceIPStatefulFirewall struct {
	InterfaceVeInstanceIPStatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceVeInstanceIPStatefulFirewallAclID      int    `json:"acl-id,omitempty"`
	InterfaceVeInstanceIPStatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceVeInstanceIPStatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceVeInstanceIPStatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceVeInstanceIPStatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6AddressList struct {
	InterfaceVeInstanceIpv6AddressListAddressType string `json:"address-type,omitempty"`
	InterfaceVeInstanceIpv6AddressListIpv6Addr    string `json:"ipv6-addr,omitempty"`
}

type InterfaceVeInstanceIpv6Ospf struct {
	InterfaceVeInstanceIpv6OspfBfd                                     int                                                `json:"bfd,omitempty"`
	InterfaceVeInstanceIpv6OspfCostCfgCost                             []InterfaceVeInstanceIpv6OspfCostCfg               `json:"cost-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfDeadIntervalCfgDeadInterval             []InterfaceVeInstanceIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfDisable                                 int                                                `json:"disable,omitempty"`
	InterfaceVeInstanceIpv6OspfHelloIntervalCfgHelloInterval           []InterfaceVeInstanceIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfMtuIgnoreCfgMtuIgnore                   []InterfaceVeInstanceIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfNeighborCfgNeighbor                     []InterfaceVeInstanceIpv6OspfNeighborCfg           `json:"neighbor-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfNetworkListBroadcastType                []InterfaceVeInstanceIpv6OspfNetworkList           `json:"network-list,omitempty"`
	InterfaceVeInstanceIpv6OspfPriorityCfgPriority                     []InterfaceVeInstanceIpv6OspfPriorityCfg           `json:"priority-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval []InterfaceVeInstanceIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfTransmitDelayCfgTransmitDelay           []InterfaceVeInstanceIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	InterfaceVeInstanceIpv6OspfUUID                                    string                                             `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6Rip struct {
	InterfaceVeInstanceIpv6RipSplitHorizonCfgState InterfaceVeInstanceIpv6RipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceVeInstanceIpv6RipUUID                 string                                    `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6Router struct {
	InterfaceVeInstanceIpv6RouterIsisTag      InterfaceVeInstanceIpv6RouterIsis  `json:"isis,omitempty"`
	InterfaceVeInstanceIpv6RouterOspfAreaList InterfaceVeInstanceIpv6RouterOspf  `json:"ospf,omitempty"`
	InterfaceVeInstanceIpv6RouterRipngRip     InterfaceVeInstanceIpv6RouterRipng `json:"ripng,omitempty"`
}

type InterfaceVeInstanceIpv6RouterAdver struct {
	InterfaceVeInstanceIpv6RouterAdverAction                   string                                         `json:"action,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverAdverMtu                 int                                            `json:"adver-mtu,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverAdverMtuDisable          int                                            `json:"adver-mtu-disable,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverAdverVrid                int                                            `json:"adver-vrid,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverAdverVridDefault         int                                            `json:"adver-vrid-default,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverDefaultLifetime          int                                            `json:"default-lifetime,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverFloatingIP               string                                         `json:"floating-ip,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverFloatingIPDefaultVrid    string                                         `json:"floating-ip-default-vrid,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverHopLimit                 int                                            `json:"hop-limit,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverManagedConfigAction      string                                         `json:"managed-config-action,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverMaxInterval              int                                            `json:"max-interval,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverMinInterval              int                                            `json:"min-interval,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverOtherConfigAction        string                                         `json:"other-config-action,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverPrefixListPrefix         []InterfaceVeInstanceIpv6RouterAdverPrefixList `json:"prefix-list,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverRateLimit                int                                            `json:"rate-limit,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverReachableTime            int                                            `json:"reachable-time,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverRetransmitTimer          int                                            `json:"retransmit-timer,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverUseFloatingIP            int                                            `json:"use-floating-ip,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverUseFloatingIPDefaultVrid int                                            `json:"use-floating-ip-default-vrid,omitempty"`
}

type InterfaceVeInstanceIpv6StatefulFirewall struct {
	InterfaceVeInstanceIpv6StatefulFirewallAccessList int    `json:"access-list,omitempty"`
	InterfaceVeInstanceIpv6StatefulFirewallAclName    string `json:"acl-name,omitempty"`
	InterfaceVeInstanceIpv6StatefulFirewallClassList  string `json:"class-list,omitempty"`
	InterfaceVeInstanceIpv6StatefulFirewallInside     int    `json:"inside,omitempty"`
	InterfaceVeInstanceIpv6StatefulFirewallOutside    int    `json:"outside,omitempty"`
	InterfaceVeInstanceIpv6StatefulFirewallUUID       string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIsisAuthentication struct {
	InterfaceVeInstanceIsisAuthenticationKeyChainListKeyChain []InterfaceVeInstanceIsisAuthenticationKeyChainList `json:"key-chain-list,omitempty"`
	InterfaceVeInstanceIsisAuthenticationModeListMode         []InterfaceVeInstanceIsisAuthenticationModeList     `json:"mode-list,omitempty"`
	InterfaceVeInstanceIsisAuthenticationSendOnlyListSendOnly []InterfaceVeInstanceIsisAuthenticationSendOnlyList `json:"send-only-list,omitempty"`
}

type InterfaceVeInstanceIsisBfdCfg struct {
	InterfaceVeInstanceIsisBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceVeInstanceIsisBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceVeInstanceIsisCsnpIntervalList struct {
	InterfaceVeInstanceIsisCsnpIntervalListCsnpInterval int    `json:"csnp-interval,omitempty"`
	InterfaceVeInstanceIsisCsnpIntervalListLevel        string `json:"level,omitempty"`
}

type InterfaceVeInstanceIsisHelloIntervalList struct {
	InterfaceVeInstanceIsisHelloIntervalListHelloInterval int    `json:"hello-interval,omitempty"`
	InterfaceVeInstanceIsisHelloIntervalListLevel         string `json:"level,omitempty"`
}

type InterfaceVeInstanceIsisHelloIntervalMinimalList struct {
	InterfaceVeInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	InterfaceVeInstanceIsisHelloIntervalMinimalListLevel                string `json:"level,omitempty"`
}

type InterfaceVeInstanceIsisHelloMultiplierList struct {
	InterfaceVeInstanceIsisHelloMultiplierListHelloMultiplier int    `json:"hello-multiplier,omitempty"`
	InterfaceVeInstanceIsisHelloMultiplierListLevel           string `json:"level,omitempty"`
}

type InterfaceVeInstanceIsisMeshGroup struct {
	InterfaceVeInstanceIsisMeshGroupBlocked int `json:"blocked,omitempty"`
	InterfaceVeInstanceIsisMeshGroupValue   int `json:"value,omitempty"`
}

type InterfaceVeInstanceIsisMetricList struct {
	InterfaceVeInstanceIsisMetricListLevel  string `json:"level,omitempty"`
	InterfaceVeInstanceIsisMetricListMetric int    `json:"metric,omitempty"`
}

type InterfaceVeInstanceIsisPasswordList struct {
	InterfaceVeInstanceIsisPasswordListLevel    string `json:"level,omitempty"`
	InterfaceVeInstanceIsisPasswordListPassword string `json:"password,omitempty"`
}

type InterfaceVeInstanceIsisPriorityList struct {
	InterfaceVeInstanceIsisPriorityListLevel    string `json:"level,omitempty"`
	InterfaceVeInstanceIsisPriorityListPriority int    `json:"priority,omitempty"`
}

type InterfaceVeInstanceIsisWideMetricList struct {
	InterfaceVeInstanceIsisWideMetricListLevel      string `json:"level,omitempty"`
	InterfaceVeInstanceIsisWideMetricListWideMetric int    `json:"wide-metric,omitempty"`
}

type InterfaceVeInstanceNptv6DomainList struct {
	InterfaceVeInstanceNptv6DomainListBindType   string `json:"bind-type,omitempty"`
	InterfaceVeInstanceNptv6DomainListDomainName string `json:"domain-name,omitempty"`
	InterfaceVeInstanceNptv6DomainListUUID       string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfGlobal struct {
	InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication  InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalAuthenticationKey                string                                                `json:"authentication-key,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalBfdCfgBfd                        InterfaceVeInstanceIPOspfOspfGlobalBfdCfg             `json:"bfd-cfg,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalCost                             int                                                   `json:"cost,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter  InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalDeadInterval                     int                                                   `json:"dead-interval,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalDisable                          string                                                `json:"disable,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalHelloInterval                    int                                                   `json:"hello-interval,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey []InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalMtu                              int                                                   `json:"mtu,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalMtuIgnore                        int                                                   `json:"mtu-ignore,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalNetworkBroadcast                 InterfaceVeInstanceIPOspfOspfGlobalNetwork            `json:"network,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalPriority                         int                                                   `json:"priority,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalRetransmitInterval               int                                                   `json:"retransmit-interval,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalTransmitDelay                    int                                                   `json:"transmit-delay,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalUUID                             string                                                `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfIPList struct {
	InterfaceVeInstanceIPOspfOspfIPListAuthentication                   int                                                   `json:"authentication,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListAuthenticationKey                string                                                `json:"authentication-key,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListCost                             int                                                   `json:"cost,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListDatabaseFilter                   string                                                `json:"database-filter,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListDeadInterval                     int                                                   `json:"dead-interval,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListHelloInterval                    int                                                   `json:"hello-interval,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListIPAddr                           string                                                `json:"ip-addr,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey []InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListMtuIgnore                        int                                                   `json:"mtu-ignore,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListOut                              int                                                   `json:"out,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListPriority                         int                                                   `json:"priority,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListRetransmitInterval               int                                                   `json:"retransmit-interval,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListTransmitDelay                    int                                                   `json:"transmit-delay,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListUUID                             string                                                `json:"uuid,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListValue                            string                                                `json:"value,omitempty"`
}

type InterfaceVeInstanceIPRipAuthentication struct {
	InterfaceVeInstanceIPRipAuthenticationKeyChainKeyChain InterfaceVeInstanceIPRipAuthenticationKeyChain `json:"key-chain,omitempty"`
	InterfaceVeInstanceIPRipAuthenticationModeMode         InterfaceVeInstanceIPRipAuthenticationMode     `json:"mode,omitempty"`
	InterfaceVeInstanceIPRipAuthenticationStrString        InterfaceVeInstanceIPRipAuthenticationStr      `json:"str,omitempty"`
}

type InterfaceVeInstanceIPRipReceiveCfg struct {
	InterfaceVeInstanceIPRipReceiveCfgReceive int    `json:"receive,omitempty"`
	InterfaceVeInstanceIPRipReceiveCfgVersion string `json:"version,omitempty"`
}

type InterfaceVeInstanceIPRipSendCfg struct {
	InterfaceVeInstanceIPRipSendCfgSend    int    `json:"send,omitempty"`
	InterfaceVeInstanceIPRipSendCfgVersion string `json:"version,omitempty"`
}

type InterfaceVeInstanceIPRipSplitHorizonCfg struct {
	InterfaceVeInstanceIPRipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceVeInstanceIPRouterIsis struct {
	InterfaceVeInstanceIPRouterIsisTag  string `json:"tag,omitempty"`
	InterfaceVeInstanceIPRouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6OspfCostCfg struct {
	InterfaceVeInstanceIpv6OspfCostCfgCost       int `json:"cost,omitempty"`
	InterfaceVeInstanceIpv6OspfCostCfgInstanceID int `json:"instance-id,omitempty"`
}

type InterfaceVeInstanceIpv6OspfDeadIntervalCfg struct {
	InterfaceVeInstanceIpv6OspfDeadIntervalCfgDeadInterval int `json:"dead-interval,omitempty"`
	InterfaceVeInstanceIpv6OspfDeadIntervalCfgInstanceID   int `json:"instance-id,omitempty"`
}

type InterfaceVeInstanceIpv6OspfHelloIntervalCfg struct {
	InterfaceVeInstanceIpv6OspfHelloIntervalCfgHelloInterval int `json:"hello-interval,omitempty"`
	InterfaceVeInstanceIpv6OspfHelloIntervalCfgInstanceID    int `json:"instance-id,omitempty"`
}

type InterfaceVeInstanceIpv6OspfMtuIgnoreCfg struct {
	InterfaceVeInstanceIpv6OspfMtuIgnoreCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceVeInstanceIpv6OspfMtuIgnoreCfgMtuIgnore  int `json:"mtu-ignore,omitempty"`
}

type InterfaceVeInstanceIpv6OspfNeighborCfg struct {
	InterfaceVeInstanceIpv6OspfNeighborCfgNeigInst             int    `json:"neig-inst,omitempty"`
	InterfaceVeInstanceIpv6OspfNeighborCfgNeighbor             string `json:"neighbor,omitempty"`
	InterfaceVeInstanceIpv6OspfNeighborCfgNeighborCost         int    `json:"neighbor-cost,omitempty"`
	InterfaceVeInstanceIpv6OspfNeighborCfgNeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	InterfaceVeInstanceIpv6OspfNeighborCfgNeighborPriority     int    `json:"neighbor-priority,omitempty"`
}

type InterfaceVeInstanceIpv6OspfNetworkList struct {
	InterfaceVeInstanceIpv6OspfNetworkListBroadcastType     string `json:"broadcast-type,omitempty"`
	InterfaceVeInstanceIpv6OspfNetworkListNetworkInstanceID int    `json:"network-instance-id,omitempty"`
	InterfaceVeInstanceIpv6OspfNetworkListP2MpNbma          int    `json:"p2mp-nbma,omitempty"`
}

type InterfaceVeInstanceIpv6OspfPriorityCfg struct {
	InterfaceVeInstanceIpv6OspfPriorityCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceVeInstanceIpv6OspfPriorityCfgPriority   int `json:"priority,omitempty"`
}

type InterfaceVeInstanceIpv6OspfRetransmitIntervalCfg struct {
	InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgInstanceID         int `json:"instance-id,omitempty"`
	InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval int `json:"retransmit-interval,omitempty"`
}

type InterfaceVeInstanceIpv6OspfTransmitDelayCfg struct {
	InterfaceVeInstanceIpv6OspfTransmitDelayCfgInstanceID    int `json:"instance-id,omitempty"`
	InterfaceVeInstanceIpv6OspfTransmitDelayCfgTransmitDelay int `json:"transmit-delay,omitempty"`
}

type InterfaceVeInstanceIpv6RipSplitHorizonCfg struct {
	InterfaceVeInstanceIpv6RipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceVeInstanceIpv6RouterIsis struct {
	InterfaceVeInstanceIpv6RouterIsisTag  string `json:"tag,omitempty"`
	InterfaceVeInstanceIpv6RouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6RouterOspf struct {
	InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDNum []InterfaceVeInstanceIpv6RouterOspfAreaList `json:"area-list,omitempty"`
	InterfaceVeInstanceIpv6RouterOspfUUID              string                                      `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6RouterRipng struct {
	InterfaceVeInstanceIpv6RouterRipngRip  int    `json:"rip,omitempty"`
	InterfaceVeInstanceIpv6RouterRipngUUID string `json:"uuid,omitempty"`
}

type InterfaceVeInstanceIpv6RouterAdverPrefixList struct {
	InterfaceVeInstanceIpv6RouterAdverPrefixListNotAutonomous     int    `json:"not-autonomous,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverPrefixListNotOnLink         int    `json:"not-on-link,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverPrefixListPreferredLifetime int    `json:"preferred-lifetime,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverPrefixListPrefix            string `json:"prefix,omitempty"`
	InterfaceVeInstanceIpv6RouterAdverPrefixListValidLifetime     int    `json:"valid-lifetime,omitempty"`
}

type InterfaceVeInstanceIsisAuthenticationKeyChainList struct {
	InterfaceVeInstanceIsisAuthenticationKeyChainListKeyChain string `json:"key-chain,omitempty"`
	InterfaceVeInstanceIsisAuthenticationKeyChainListLevel    string `json:"level,omitempty"`
}

type InterfaceVeInstanceIsisAuthenticationModeList struct {
	InterfaceVeInstanceIsisAuthenticationModeListLevel string `json:"level,omitempty"`
	InterfaceVeInstanceIsisAuthenticationModeListMode  string `json:"mode,omitempty"`
}

type InterfaceVeInstanceIsisAuthenticationSendOnlyList struct {
	InterfaceVeInstanceIsisAuthenticationSendOnlyListLevel    string `json:"level,omitempty"`
	InterfaceVeInstanceIsisAuthenticationSendOnlyListSendOnly int    `json:"send-only,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfg struct {
	InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication int    `json:"authentication,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfgValue          string `json:"value,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfGlobalBfdCfg struct {
	InterfaceVeInstanceIPOspfOspfGlobalBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfg struct {
	InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter string `json:"database-filter,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfgOut            int    `json:"out,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfg struct {
	InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfGlobalNetwork struct {
	InterfaceVeInstanceIPOspfOspfGlobalNetworkBroadcast         int `json:"broadcast,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalNetworkNonBroadcast      int `json:"non-broadcast,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalNetworkP2MpNbma          int `json:"p2mp-nbma,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalNetworkPointToMultipoint int `json:"point-to-multipoint,omitempty"`
	InterfaceVeInstanceIPOspfOspfGlobalNetworkPointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfg struct {
	InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceVeInstanceIPRipAuthenticationKeyChain struct {
	InterfaceVeInstanceIPRipAuthenticationKeyChainKeyChain string `json:"key-chain,omitempty"`
}

type InterfaceVeInstanceIPRipAuthenticationMode struct {
	InterfaceVeInstanceIPRipAuthenticationModeMode string `json:"mode,omitempty"`
}

type InterfaceVeInstanceIPRipAuthenticationStr struct {
	InterfaceVeInstanceIPRipAuthenticationStrString string `json:"string,omitempty"`
}

type InterfaceVeInstanceIpv6RouterOspfAreaList struct {
	InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDAddr string `json:"area-id-addr,omitempty"`
	InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDNum  int    `json:"area-id-num,omitempty"`
	InterfaceVeInstanceIpv6RouterOspfAreaListInstanceID int    `json:"instance-id,omitempty"`
	InterfaceVeInstanceIpv6RouterOspfAreaListTag        string `json:"tag,omitempty"`
}

func PostInterfaceVe(id string, inst InterfaceVe, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceVe")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ve", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVe
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceVe", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceVe(id string, name1 string, host string) (*InterfaceVe, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceVe")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ve/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVe
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceVe", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutInterfaceVe(id string, name1 string, inst InterfaceVe, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceVe")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/ve/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVe
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutInterfaceVe", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteInterfaceVe(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceVe")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/ve/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceVe
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteInterfaceVe", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
