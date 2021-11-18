package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type InterfaceLif struct {
	InterfaceLifInstanceIfname InterfaceLifInstance `json:"lif,omitempty"`
}

type InterfaceLifInstance struct {
	InterfaceLifInstanceAccessListAclID         InterfaceLifInstanceAccessList       `json:"access-list,omitempty"`
	InterfaceLifInstanceAction                  string                               `json:"action,omitempty"`
	InterfaceLifInstanceBfdAuthentication       InterfaceLifInstanceBfd              `json:"bfd,omitempty"`
	InterfaceLifInstanceIPDhcp                  InterfaceLifInstanceIP               `json:"ip,omitempty"`
	InterfaceLifInstanceIfname                  string                               `json:"ifname,omitempty"`
	InterfaceLifInstanceIpv6AddressList         InterfaceLifInstanceIpv6             `json:"ipv6,omitempty"`
	InterfaceLifInstanceIsisAuthentication      InterfaceLifInstanceIsis             `json:"isis,omitempty"`
	InterfaceLifInstanceMtu                     int                                  `json:"mtu,omitempty"`
	InterfaceLifInstanceName                    string                               `json:"name,omitempty"`
	InterfaceLifInstanceSamplingEnableCounters1 []InterfaceLifInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	InterfaceLifInstanceUUID                    string                               `json:"uuid,omitempty"`
	InterfaceLifInstanceUserTag                 string                               `json:"user-tag,omitempty"`
}

type InterfaceLifInstanceAccessList struct {
	InterfaceLifInstanceAccessListAclID   int    `json:"acl-id,omitempty"`
	InterfaceLifInstanceAccessListAclName string `json:"acl-name,omitempty"`
}

type InterfaceLifInstanceBfd struct {
	InterfaceLifInstanceBfdAuthenticationKeyID InterfaceLifInstanceBfdAuthentication `json:"authentication,omitempty"`
	InterfaceLifInstanceBfdDemand              int                                   `json:"demand,omitempty"`
	InterfaceLifInstanceBfdEcho                int                                   `json:"echo,omitempty"`
	InterfaceLifInstanceBfdIntervalCfgInterval InterfaceLifInstanceBfdIntervalCfg    `json:"interval-cfg,omitempty"`
	InterfaceLifInstanceBfdUUID                string                                `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIP struct {
	InterfaceLifInstanceIPAddressListIpv6Addr     []InterfaceLifInstanceIPAddressList `json:"address-list,omitempty"`
	InterfaceLifInstanceIPAllowPromiscuousVip     int                                 `json:"allow-promiscuous-vip,omitempty"`
	InterfaceLifInstanceIPCacheSpoofingPort       int                                 `json:"cache-spoofing-port,omitempty"`
	InterfaceLifInstanceIPDhcp                    int                                 `json:"dhcp,omitempty"`
	InterfaceLifInstanceIPGenerateMembershipQuery int                                 `json:"generate-membership-query,omitempty"`
	InterfaceLifInstanceIPInside                  int                                 `json:"inside,omitempty"`
	InterfaceLifInstanceIPMaxRespTime             int                                 `json:"max-resp-time,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobal          InterfaceLifInstanceIPOspf          `json:"ospf,omitempty"`
	InterfaceLifInstanceIPOutside                 int                                 `json:"outside,omitempty"`
	InterfaceLifInstanceIPQueryInterval           int                                 `json:"query-interval,omitempty"`
	InterfaceLifInstanceIPRipAuthentication       InterfaceLifInstanceIPRip           `json:"rip,omitempty"`
	InterfaceLifInstanceIPRouterIsis              InterfaceLifInstanceIPRouter        `json:"router,omitempty"`
	InterfaceLifInstanceIPUUID                    string                              `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIpv6 struct {
	InterfaceLifInstanceIpv6AddressListIpv6Addr []InterfaceLifInstanceIpv6AddressList `json:"address-list,omitempty"`
	InterfaceLifInstanceIpv6Inside              int                                   `json:"inside,omitempty"`
	InterfaceLifInstanceIpv6Ipv6Enable          int                                   `json:"ipv6-enable,omitempty"`
	InterfaceLifInstanceIpv6OspfNetworkList     InterfaceLifInstanceIpv6Ospf          `json:"ospf,omitempty"`
	InterfaceLifInstanceIpv6Outside             int                                   `json:"outside,omitempty"`
	InterfaceLifInstanceIpv6RouterRipng         InterfaceLifInstanceIpv6Router        `json:"router,omitempty"`
	InterfaceLifInstanceIpv6UUID                string                                `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIsis struct {
	InterfaceLifInstanceIsisAuthenticationSendOnlyList                   InterfaceLifInstanceIsisAuthentication             `json:"authentication,omitempty"`
	InterfaceLifInstanceIsisBfdCfgBfd                                    InterfaceLifInstanceIsisBfdCfg                     `json:"bfd-cfg,omitempty"`
	InterfaceLifInstanceIsisCircuitType                                  string                                             `json:"circuit-type,omitempty"`
	InterfaceLifInstanceIsisCsnpIntervalListCsnpInterval                 []InterfaceLifInstanceIsisCsnpIntervalList         `json:"csnp-interval-list,omitempty"`
	InterfaceLifInstanceIsisHelloIntervalListHelloInterval               []InterfaceLifInstanceIsisHelloIntervalList        `json:"hello-interval-list,omitempty"`
	InterfaceLifInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal []InterfaceLifInstanceIsisHelloIntervalMinimalList `json:"hello-interval-minimal-list,omitempty"`
	InterfaceLifInstanceIsisHelloMultiplierListHelloMultiplier           []InterfaceLifInstanceIsisHelloMultiplierList      `json:"hello-multiplier-list,omitempty"`
	InterfaceLifInstanceIsisLspInterval                                  int                                                `json:"lsp-interval,omitempty"`
	InterfaceLifInstanceIsisMeshGroupValue                               InterfaceLifInstanceIsisMeshGroup                  `json:"mesh-group,omitempty"`
	InterfaceLifInstanceIsisMetricListMetric                             []InterfaceLifInstanceIsisMetricList               `json:"metric-list,omitempty"`
	InterfaceLifInstanceIsisNetwork                                      string                                             `json:"network,omitempty"`
	InterfaceLifInstanceIsisPadding                                      int                                                `json:"padding,omitempty"`
	InterfaceLifInstanceIsisPasswordListPassword                         []InterfaceLifInstanceIsisPasswordList             `json:"password-list,omitempty"`
	InterfaceLifInstanceIsisPriorityListPriority                         []InterfaceLifInstanceIsisPriorityList             `json:"priority-list,omitempty"`
	InterfaceLifInstanceIsisRetransmitInterval                           int                                                `json:"retransmit-interval,omitempty"`
	InterfaceLifInstanceIsisUUID                                         string                                             `json:"uuid,omitempty"`
	InterfaceLifInstanceIsisWideMetricListWideMetric                     []InterfaceLifInstanceIsisWideMetricList           `json:"wide-metric-list,omitempty"`
}

type InterfaceLifInstanceSamplingEnable struct {
	InterfaceLifInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type InterfaceLifInstanceBfdAuthentication struct {
	InterfaceLifInstanceBfdAuthenticationKeyID    int    `json:"key-id,omitempty"`
	InterfaceLifInstanceBfdAuthenticationMethod   string `json:"method,omitempty"`
	InterfaceLifInstanceBfdAuthenticationPassword string `json:"password,omitempty"`
}

type InterfaceLifInstanceBfdIntervalCfg struct {
	InterfaceLifInstanceBfdIntervalCfgInterval   int `json:"interval,omitempty"`
	InterfaceLifInstanceBfdIntervalCfgMinRx      int `json:"min-rx,omitempty"`
	InterfaceLifInstanceBfdIntervalCfgMultiplier int `json:"multiplier,omitempty"`
}

type InterfaceLifInstanceIPAddressList struct {
	InterfaceLifInstanceIPAddressListAnycast   int    `json:"anycast,omitempty"`
	InterfaceLifInstanceIPAddressListIpv6Addr  string `json:"ipv6-addr,omitempty"`
	InterfaceLifInstanceIPAddressListLinkLocal int    `json:"link-local,omitempty"`
}

type InterfaceLifInstanceIPOspf struct {
	InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfg InterfaceLifInstanceIPOspfOspfGlobal   `json:"ospf-global,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListIPAddr            []InterfaceLifInstanceIPOspfOspfIPList `json:"ospf-ip-list,omitempty"`
}

type InterfaceLifInstanceIPRip struct {
	InterfaceLifInstanceIPRipAuthenticationStr    InterfaceLifInstanceIPRipAuthentication  `json:"authentication,omitempty"`
	InterfaceLifInstanceIPRipReceiveCfgReceive    InterfaceLifInstanceIPRipReceiveCfg      `json:"receive-cfg,omitempty"`
	InterfaceLifInstanceIPRipReceivePacket        int                                      `json:"receive-packet,omitempty"`
	InterfaceLifInstanceIPRipSendCfgSend          InterfaceLifInstanceIPRipSendCfg         `json:"send-cfg,omitempty"`
	InterfaceLifInstanceIPRipSendPacket           int                                      `json:"send-packet,omitempty"`
	InterfaceLifInstanceIPRipSplitHorizonCfgState InterfaceLifInstanceIPRipSplitHorizonCfg `json:"split-horizon-cfg,omitempty"`
	InterfaceLifInstanceIPRipUUID                 string                                   `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIPRouter struct {
	InterfaceLifInstanceIPRouterIsisTag InterfaceLifInstanceIPRouterIsis `json:"isis,omitempty"`
}

type InterfaceLifInstanceIpv6AddressList struct {
	InterfaceLifInstanceIpv6AddressListAnycast   int    `json:"anycast,omitempty"`
	InterfaceLifInstanceIpv6AddressListIpv6Addr  string `json:"ipv6-addr,omitempty"`
	InterfaceLifInstanceIpv6AddressListLinkLocal int    `json:"link-local,omitempty"`
}

type InterfaceLifInstanceIpv6Ospf struct {
	InterfaceLifInstanceIpv6OspfBfd                                     int                                                 `json:"bfd,omitempty"`
	InterfaceLifInstanceIpv6OspfCostCfgCost                             []InterfaceLifInstanceIpv6OspfCostCfg               `json:"cost-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfDeadIntervalCfgDeadInterval             []InterfaceLifInstanceIpv6OspfDeadIntervalCfg       `json:"dead-interval-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfDisable                                 int                                                 `json:"disable,omitempty"`
	InterfaceLifInstanceIpv6OspfHelloIntervalCfgHelloInterval           []InterfaceLifInstanceIpv6OspfHelloIntervalCfg      `json:"hello-interval-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfMtuIgnoreCfgMtuIgnore                   []InterfaceLifInstanceIpv6OspfMtuIgnoreCfg          `json:"mtu-ignore-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfNeighborCfgNeighbor                     []InterfaceLifInstanceIpv6OspfNeighborCfg           `json:"neighbor-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfNetworkListBroadcastType                []InterfaceLifInstanceIpv6OspfNetworkList           `json:"network-list,omitempty"`
	InterfaceLifInstanceIpv6OspfPriorityCfgPriority                     []InterfaceLifInstanceIpv6OspfPriorityCfg           `json:"priority-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval []InterfaceLifInstanceIpv6OspfRetransmitIntervalCfg `json:"retransmit-interval-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfTransmitDelayCfgTransmitDelay           []InterfaceLifInstanceIpv6OspfTransmitDelayCfg      `json:"transmit-delay-cfg,omitempty"`
	InterfaceLifInstanceIpv6OspfUUID                                    string                                              `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIpv6Router struct {
	InterfaceLifInstanceIpv6RouterIsisTag      InterfaceLifInstanceIpv6RouterIsis  `json:"isis,omitempty"`
	InterfaceLifInstanceIpv6RouterOspfAreaList InterfaceLifInstanceIpv6RouterOspf  `json:"ospf,omitempty"`
	InterfaceLifInstanceIpv6RouterRipngRip     InterfaceLifInstanceIpv6RouterRipng `json:"ripng,omitempty"`
}

type InterfaceLifInstanceIsisAuthentication struct {
	InterfaceLifInstanceIsisAuthenticationKeyChainListKeyChain []InterfaceLifInstanceIsisAuthenticationKeyChainList `json:"key-chain-list,omitempty"`
	InterfaceLifInstanceIsisAuthenticationModeListMode         []InterfaceLifInstanceIsisAuthenticationModeList     `json:"mode-list,omitempty"`
	InterfaceLifInstanceIsisAuthenticationSendOnlyListSendOnly []InterfaceLifInstanceIsisAuthenticationSendOnlyList `json:"send-only-list,omitempty"`
}

type InterfaceLifInstanceIsisBfdCfg struct {
	InterfaceLifInstanceIsisBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceLifInstanceIsisBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceLifInstanceIsisCsnpIntervalList struct {
	InterfaceLifInstanceIsisCsnpIntervalListCsnpInterval int    `json:"csnp-interval,omitempty"`
	InterfaceLifInstanceIsisCsnpIntervalListLevel        string `json:"level,omitempty"`
}

type InterfaceLifInstanceIsisHelloIntervalList struct {
	InterfaceLifInstanceIsisHelloIntervalListHelloInterval int    `json:"hello-interval,omitempty"`
	InterfaceLifInstanceIsisHelloIntervalListLevel         string `json:"level,omitempty"`
}

type InterfaceLifInstanceIsisHelloIntervalMinimalList struct {
	InterfaceLifInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal int    `json:"hello-interval-minimal,omitempty"`
	InterfaceLifInstanceIsisHelloIntervalMinimalListLevel                string `json:"level,omitempty"`
}

type InterfaceLifInstanceIsisHelloMultiplierList struct {
	InterfaceLifInstanceIsisHelloMultiplierListHelloMultiplier int    `json:"hello-multiplier,omitempty"`
	InterfaceLifInstanceIsisHelloMultiplierListLevel           string `json:"level,omitempty"`
}

type InterfaceLifInstanceIsisMeshGroup struct {
	InterfaceLifInstanceIsisMeshGroupBlocked int `json:"blocked,omitempty"`
	InterfaceLifInstanceIsisMeshGroupValue   int `json:"value,omitempty"`
}

type InterfaceLifInstanceIsisMetricList struct {
	InterfaceLifInstanceIsisMetricListLevel  string `json:"level,omitempty"`
	InterfaceLifInstanceIsisMetricListMetric int    `json:"metric,omitempty"`
}

type InterfaceLifInstanceIsisPasswordList struct {
	InterfaceLifInstanceIsisPasswordListLevel    string `json:"level,omitempty"`
	InterfaceLifInstanceIsisPasswordListPassword string `json:"password,omitempty"`
}

type InterfaceLifInstanceIsisPriorityList struct {
	InterfaceLifInstanceIsisPriorityListLevel    string `json:"level,omitempty"`
	InterfaceLifInstanceIsisPriorityListPriority int    `json:"priority,omitempty"`
}

type InterfaceLifInstanceIsisWideMetricList struct {
	InterfaceLifInstanceIsisWideMetricListLevel      string `json:"level,omitempty"`
	InterfaceLifInstanceIsisWideMetricListWideMetric int    `json:"wide-metric,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfGlobal struct {
	InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication  InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfg  `json:"authentication-cfg,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalAuthenticationKey                string                                                 `json:"authentication-key,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalBfdCfgBfd                        InterfaceLifInstanceIPOspfOspfGlobalBfdCfg             `json:"bfd-cfg,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalCost                             int                                                    `json:"cost,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter  InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfg  `json:"database-filter-cfg,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalDeadInterval                     int                                                    `json:"dead-interval,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalDisable                          string                                                 `json:"disable,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalHelloInterval                    int                                                    `json:"hello-interval,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey []InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalMtu                              int                                                    `json:"mtu,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalMtuIgnore                        int                                                    `json:"mtu-ignore,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalNetworkBroadcast                 InterfaceLifInstanceIPOspfOspfGlobalNetwork            `json:"network,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalPriority                         int                                                    `json:"priority,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalRetransmitInterval               int                                                    `json:"retransmit-interval,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalTransmitDelay                    int                                                    `json:"transmit-delay,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalUUID                             string                                                 `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfIPList struct {
	InterfaceLifInstanceIPOspfOspfIPListAuthentication                   int                                                    `json:"authentication,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListAuthenticationKey                string                                                 `json:"authentication-key,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListCost                             int                                                    `json:"cost,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListDatabaseFilter                   string                                                 `json:"database-filter,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListDeadInterval                     int                                                    `json:"dead-interval,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListHelloInterval                    int                                                    `json:"hello-interval,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListIPAddr                           string                                                 `json:"ip-addr,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey []InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfg `json:"message-digest-cfg,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListMtuIgnore                        int                                                    `json:"mtu-ignore,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListOut                              int                                                    `json:"out,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListPriority                         int                                                    `json:"priority,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListRetransmitInterval               int                                                    `json:"retransmit-interval,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListTransmitDelay                    int                                                    `json:"transmit-delay,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListUUID                             string                                                 `json:"uuid,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListValue                            string                                                 `json:"value,omitempty"`
}

type InterfaceLifInstanceIPRipAuthentication struct {
	InterfaceLifInstanceIPRipAuthenticationKeyChainKeyChain InterfaceLifInstanceIPRipAuthenticationKeyChain `json:"key-chain,omitempty"`
	InterfaceLifInstanceIPRipAuthenticationModeMode         InterfaceLifInstanceIPRipAuthenticationMode     `json:"mode,omitempty"`
	InterfaceLifInstanceIPRipAuthenticationStrString        InterfaceLifInstanceIPRipAuthenticationStr      `json:"str,omitempty"`
}

type InterfaceLifInstanceIPRipReceiveCfg struct {
	InterfaceLifInstanceIPRipReceiveCfgReceive int    `json:"receive,omitempty"`
	InterfaceLifInstanceIPRipReceiveCfgVersion string `json:"version,omitempty"`
}

type InterfaceLifInstanceIPRipSendCfg struct {
	InterfaceLifInstanceIPRipSendCfgSend    int    `json:"send,omitempty"`
	InterfaceLifInstanceIPRipSendCfgVersion string `json:"version,omitempty"`
}

type InterfaceLifInstanceIPRipSplitHorizonCfg struct {
	InterfaceLifInstanceIPRipSplitHorizonCfgState string `json:"state,omitempty"`
}

type InterfaceLifInstanceIPRouterIsis struct {
	InterfaceLifInstanceIPRouterIsisTag  string `json:"tag,omitempty"`
	InterfaceLifInstanceIPRouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIpv6OspfCostCfg struct {
	InterfaceLifInstanceIpv6OspfCostCfgCost       int `json:"cost,omitempty"`
	InterfaceLifInstanceIpv6OspfCostCfgInstanceID int `json:"instance-id,omitempty"`
}

type InterfaceLifInstanceIpv6OspfDeadIntervalCfg struct {
	InterfaceLifInstanceIpv6OspfDeadIntervalCfgDeadInterval int `json:"dead-interval,omitempty"`
	InterfaceLifInstanceIpv6OspfDeadIntervalCfgInstanceID   int `json:"instance-id,omitempty"`
}

type InterfaceLifInstanceIpv6OspfHelloIntervalCfg struct {
	InterfaceLifInstanceIpv6OspfHelloIntervalCfgHelloInterval int `json:"hello-interval,omitempty"`
	InterfaceLifInstanceIpv6OspfHelloIntervalCfgInstanceID    int `json:"instance-id,omitempty"`
}

type InterfaceLifInstanceIpv6OspfMtuIgnoreCfg struct {
	InterfaceLifInstanceIpv6OspfMtuIgnoreCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceLifInstanceIpv6OspfMtuIgnoreCfgMtuIgnore  int `json:"mtu-ignore,omitempty"`
}

type InterfaceLifInstanceIpv6OspfNeighborCfg struct {
	InterfaceLifInstanceIpv6OspfNeighborCfgNeigInst             int    `json:"neig-inst,omitempty"`
	InterfaceLifInstanceIpv6OspfNeighborCfgNeighbor             string `json:"neighbor,omitempty"`
	InterfaceLifInstanceIpv6OspfNeighborCfgNeighborCost         int    `json:"neighbor-cost,omitempty"`
	InterfaceLifInstanceIpv6OspfNeighborCfgNeighborPollInterval int    `json:"neighbor-poll-interval,omitempty"`
	InterfaceLifInstanceIpv6OspfNeighborCfgNeighborPriority     int    `json:"neighbor-priority,omitempty"`
}

type InterfaceLifInstanceIpv6OspfNetworkList struct {
	InterfaceLifInstanceIpv6OspfNetworkListBroadcastType     string `json:"broadcast-type,omitempty"`
	InterfaceLifInstanceIpv6OspfNetworkListNetworkInstanceID int    `json:"network-instance-id,omitempty"`
	InterfaceLifInstanceIpv6OspfNetworkListP2MpNbma          int    `json:"p2mp-nbma,omitempty"`
}

type InterfaceLifInstanceIpv6OspfPriorityCfg struct {
	InterfaceLifInstanceIpv6OspfPriorityCfgInstanceID int `json:"instance-id,omitempty"`
	InterfaceLifInstanceIpv6OspfPriorityCfgPriority   int `json:"priority,omitempty"`
}

type InterfaceLifInstanceIpv6OspfRetransmitIntervalCfg struct {
	InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgInstanceID         int `json:"instance-id,omitempty"`
	InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval int `json:"retransmit-interval,omitempty"`
}

type InterfaceLifInstanceIpv6OspfTransmitDelayCfg struct {
	InterfaceLifInstanceIpv6OspfTransmitDelayCfgInstanceID    int `json:"instance-id,omitempty"`
	InterfaceLifInstanceIpv6OspfTransmitDelayCfgTransmitDelay int `json:"transmit-delay,omitempty"`
}

type InterfaceLifInstanceIpv6RouterIsis struct {
	InterfaceLifInstanceIpv6RouterIsisTag  string `json:"tag,omitempty"`
	InterfaceLifInstanceIpv6RouterIsisUUID string `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIpv6RouterOspf struct {
	InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDNum []InterfaceLifInstanceIpv6RouterOspfAreaList `json:"area-list,omitempty"`
	InterfaceLifInstanceIpv6RouterOspfUUID              string                                       `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIpv6RouterRipng struct {
	InterfaceLifInstanceIpv6RouterRipngRip  int    `json:"rip,omitempty"`
	InterfaceLifInstanceIpv6RouterRipngUUID string `json:"uuid,omitempty"`
}

type InterfaceLifInstanceIsisAuthenticationKeyChainList struct {
	InterfaceLifInstanceIsisAuthenticationKeyChainListKeyChain string `json:"key-chain,omitempty"`
	InterfaceLifInstanceIsisAuthenticationKeyChainListLevel    string `json:"level,omitempty"`
}

type InterfaceLifInstanceIsisAuthenticationModeList struct {
	InterfaceLifInstanceIsisAuthenticationModeListLevel string `json:"level,omitempty"`
	InterfaceLifInstanceIsisAuthenticationModeListMode  string `json:"mode,omitempty"`
}

type InterfaceLifInstanceIsisAuthenticationSendOnlyList struct {
	InterfaceLifInstanceIsisAuthenticationSendOnlyListLevel    string `json:"level,omitempty"`
	InterfaceLifInstanceIsisAuthenticationSendOnlyListSendOnly int    `json:"send-only,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfg struct {
	InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication int    `json:"authentication,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfgValue          string `json:"value,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfGlobalBfdCfg struct {
	InterfaceLifInstanceIPOspfOspfGlobalBfdCfgBfd     int `json:"bfd,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalBfdCfgDisable int `json:"disable,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfg struct {
	InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter string `json:"database-filter,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfgOut            int    `json:"out,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfg struct {
	InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfGlobalNetwork struct {
	InterfaceLifInstanceIPOspfOspfGlobalNetworkBroadcast         int `json:"broadcast,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalNetworkNonBroadcast      int `json:"non-broadcast,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalNetworkP2MpNbma          int `json:"p2mp-nbma,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalNetworkPointToMultipoint int `json:"point-to-multipoint,omitempty"`
	InterfaceLifInstanceIPOspfOspfGlobalNetworkPointToPoint      int `json:"point-to-point,omitempty"`
}

type InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfg struct {
	InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMd5Value         string `json:"md5-value,omitempty"`
	InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey int    `json:"message-digest-key,omitempty"`
}

type InterfaceLifInstanceIPRipAuthenticationKeyChain struct {
	InterfaceLifInstanceIPRipAuthenticationKeyChainKeyChain string `json:"key-chain,omitempty"`
}

type InterfaceLifInstanceIPRipAuthenticationMode struct {
	InterfaceLifInstanceIPRipAuthenticationModeMode string `json:"mode,omitempty"`
}

type InterfaceLifInstanceIPRipAuthenticationStr struct {
	InterfaceLifInstanceIPRipAuthenticationStrString string `json:"string,omitempty"`
}

type InterfaceLifInstanceIpv6RouterOspfAreaList struct {
	InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDAddr string `json:"area-id-addr,omitempty"`
	InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDNum  int    `json:"area-id-num,omitempty"`
	InterfaceLifInstanceIpv6RouterOspfAreaListInstanceID int    `json:"instance-id,omitempty"`
	InterfaceLifInstanceIpv6RouterOspfAreaListTag        string `json:"tag,omitempty"`
}

func PostInterfaceLif(id string, inst InterfaceLif, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceLif")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/lif", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostInterfaceLif", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceLif(id string, name1 string, host string) (*InterfaceLif, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceLif")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/lif/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetInterfaceLif", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutInterfaceLif(id string, name1 string, inst InterfaceLif, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutInterfaceLif")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/lif/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutInterfaceLif", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteInterfaceLif(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteInterfaceLif")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/interface/lif/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m InterfaceLif
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteInterfaceLif", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
