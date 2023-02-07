package go_thunder

import (
	"bytes"
	"encoding/json"
	json_marsh "github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"net/url"
	"util"
)

type VirtualServer struct {
	VirtualServerInstanceName VirtualServerInstance `json:"virtual-server,omitempty"`
}

type VirtualServerInstance struct {
	VirtualServerInstanceAclID                         int                             `json:"acl-id,omitempty"`
	VirtualServerInstanceAclIDShared                   int                             `json:"acl-id-shared,omitempty"`
	VirtualServerInstanceAclName                       string                          `json:"acl-name,omitempty"`
	VirtualServerInstanceAclNameShared                 string                          `json:"acl-name-shared,omitempty"`
	VirtualServerInstanceArpDisable                    int                             `json:"arp-disable,omitempty"`
	VirtualServerInstanceDescription                   string                          `json:"description,omitempty"`
	VirtualServerInstanceDisableVipAdv                 int                             `json:"disable-vip-adv,omitempty"`
	VirtualServerInstanceEnableDisableAction           string                          `json:"enable-disable-action,omitempty"`
	VirtualServerInstanceEthernet                      int                             `json:"ethernet,omitempty"`
	VirtualServerInstanceExtendedStats                 int                             `json:"extended-stats,omitempty"`
	VirtualServerInstanceHaDynamic                     int                             `json:"ha-dynamic,omitempty"`
	VirtualServerInstanceIPAddress                     string                          `json:"ip-address,omitempty"`
	VirtualServerInstanceIpv6Acl                       string                          `json:"ipv6-acl,omitempty"`
	VirtualServerInstanceIpv6AclShared                 string                          `json:"ipv6-acl-shared,omitempty"`
	VirtualServerInstanceIpv6Address                   string                          `json:"ipv6-address,omitempty"`
	VirtualServerInstanceMigrateVipTargetDataCPU       VirtualServerInstanceMigrateVip `json:"migrate-vip,omitempty"`
	VirtualServerInstanceName                          string                          `json:"name,omitempty"`
	VirtualServerInstanceNetmask                       string                          `json:"netmask,omitempty"`
	VirtualServerInstancePortListPortNumber            []VirtualServerInstancePortList `json:"port-list,omitempty"`
	VirtualServerInstanceRedistributeRouteMap          string                          `json:"redistribute-route-map,omitempty"`
	VirtualServerInstanceRedistributionFlagged         int                             `json:"redistribution-flagged,omitempty"`
	VirtualServerInstanceSharedPartitionPolicyTemplate int                             `json:"shared-partition-policy-template,omitempty"`
	VirtualServerInstanceSharedPartitionVsTemplate     int                             `json:"shared-partition-vs-template,omitempty"`
	VirtualServerInstanceStatsDataAction               string                          `json:"stats-data-action,omitempty"`
	VirtualServerInstanceSuppressInternalLoopback      int                             `json:"suppress-internal-loopback,omitempty"`
	VirtualServerInstanceTemplateLogging               string                          `json:"template-logging,omitempty"`
	VirtualServerInstanceTemplatePolicy                string                          `json:"template-policy,omitempty"`
	VirtualServerInstanceTemplatePolicyShared          string                          `json:"template-policy-shared,omitempty"`
	VirtualServerInstanceTemplateScaleout              string                          `json:"template-scaleout,omitempty"`
	VirtualServerInstanceTemplateVirtualServer         string                          `json:"template-virtual-server,omitempty"`
	VirtualServerInstanceTemplateVirtualServerShared   string                          `json:"template-virtual-server-shared,omitempty"`
	VirtualServerInstanceUUID                          string                          `json:"uuid,omitempty"`
	VirtualServerInstanceUseIfIP                       int                             `json:"use-if-ip,omitempty"`
	VirtualServerInstanceUserTag                       string                          `json:"user-tag,omitempty"`
	VirtualServerInstanceVportDisableAction            string                          `json:"vport-disable-action,omitempty"`
	VirtualServerInstanceVrid                          int                             `json:"vrid,omitempty"`
}

type VirtualServerInstanceMigrateVip struct {
	VirtualServerInstanceMigrateVipCancelMigration    int    `json:"cancel-migration,omitempty"`
	VirtualServerInstanceMigrateVipFinishMigration    int    `json:"finish-migration,omitempty"`
	VirtualServerInstanceMigrateVipTargetDataCPU      int    `json:"target-data-cpu,omitempty"`
	VirtualServerInstanceMigrateVipTargetFloatingIpv4 string `json:"target-floating-ipv4,omitempty"`
	VirtualServerInstanceMigrateVipTargetFloatingIpv6 string `json:"target-floating-ipv6,omitempty"`
	VirtualServerInstanceMigrateVipUUID               string `json:"uuid,omitempty"`
}

type VirtualServerInstancePortList struct {
	VirtualServerInstancePortListAclListAclID                                []VirtualServerInstancePortListAclList        `json:"acl-list,omitempty"`
	VirtualServerInstancePortListAction                                      string                                        `json:"action,omitempty"`
	VirtualServerInstancePortListAflexScriptsAflex                           []VirtualServerInstancePortListAflexScripts   `json:"aflex-scripts,omitempty"`
	VirtualServerInstancePortListAflexTableEntrySynDisable                   int                                           `json:"aflex-table-entry-syn-disable,omitempty"`
	VirtualServerInstancePortListAflexTableEntrySynEnable                    int                                           `json:"aflex-table-entry-syn-enable,omitempty"`
	VirtualServerInstancePortListAltProtocol1                                string                                        `json:"alt-protocol1,omitempty"`
	VirtualServerInstancePortListAltProtocol2                                string                                        `json:"alt-protocol2,omitempty"`
	VirtualServerInstancePortListAlternatePort                               int                                           `json:"alternate-port,omitempty"`
	VirtualServerInstancePortListAlternatePortNumber                         int                                           `json:"alternate-port-number,omitempty"`
	VirtualServerInstancePortListAttackDetection                             int                                           `json:"attack-detection,omitempty"`
	VirtualServerInstancePortListAuthCfgAaaPolicy                            VirtualServerInstancePortListAuthCfg          `json:"auth-cfg,omitempty"`
	VirtualServerInstancePortListAuto                                        int                                           `json:"auto,omitempty"`
	VirtualServerInstancePortListCPUCompute                                  int                                           `json:"cpu-compute,omitempty"`
	VirtualServerInstancePortListClientipStickyNat                           int                                           `json:"clientip-sticky-nat,omitempty"`
	VirtualServerInstancePortListConnLimit                                   int                                           `json:"conn-limit,omitempty"`
	VirtualServerInstancePortListDefSelectionIfPrefFailed                    string                                        `json:"def-selection-if-pref-failed,omitempty"`
	VirtualServerInstancePortListEnablePlayeridCheck                         int                                           `json:"enable-playerid-check,omitempty"`
	VirtualServerInstancePortListEnableScaleout                              int                                           `json:"enable-scaleout,omitempty"`
	VirtualServerInstancePortListEthFwd                                      int                                           `json:"eth-fwd,omitempty"`
	VirtualServerInstancePortListEthRev                                      int                                           `json:"eth-rev,omitempty"`
	VirtualServerInstancePortListExpand                                      int                                           `json:"expand,omitempty"`
	VirtualServerInstancePortListExtendedStats                               int                                           `json:"extended-stats,omitempty"`
	VirtualServerInstancePortListForceRoutingMode                            int                                           `json:"force-routing-mode,omitempty"`
	VirtualServerInstancePortListGslbEnable                                  int                                           `json:"gslb-enable,omitempty"`
	VirtualServerInstancePortListGtpSessionLb                                int                                           `json:"gtp-session-lb,omitempty"`
	VirtualServerInstancePortListHaConnMirror                                int                                           `json:"ha-conn-mirror,omitempty"`
	VirtualServerInstancePortListIPMapList                                   string                                        `json:"ip-map-list,omitempty"`
	VirtualServerInstancePortListIPOnlyLb                                    int                                           `json:"ip-only-lb,omitempty"`
	VirtualServerInstancePortListIPSmartRr                                   int                                           `json:"ip-smart-rr,omitempty"`
	VirtualServerInstancePortListIgnoreGlobal                                int                                           `json:"ignore-global,omitempty"`
	VirtualServerInstancePortListIpinip                                      int                                           `json:"ipinip,omitempty"`
	VirtualServerInstancePortListL7HardwareAssist                            int                                           `json:"l7-hardware-assist,omitempty"`
	VirtualServerInstancePortListL7ServiceChain                              int                                           `json:"l7-service-chain,omitempty"`
	VirtualServerInstancePortListMemoryCompute                               int                                           `json:"memory-compute,omitempty"`
	VirtualServerInstancePortListMessageSwitching                            int                                           `json:"message-switching,omitempty"`
	VirtualServerInstancePortListName                                        string                                        `json:"name,omitempty"`
	VirtualServerInstancePortListNoAutoUpOnAflex                             int                                           `json:"no-auto-up-on-aflex,omitempty"`
	VirtualServerInstancePortListNoDestNat                                   int                                           `json:"no-dest-nat,omitempty"`
	VirtualServerInstancePortListNoLogging                                   int                                           `json:"no-logging,omitempty"`
	VirtualServerInstancePortListOnSyn                                       int                                           `json:"on-syn,omitempty"`
	VirtualServerInstancePortListOneServerConn                               int                                           `json:"one-server-conn,omitempty"`
	VirtualServerInstancePortListOptimizationLevel                           string                                        `json:"optimization-level,omitempty"`
	VirtualServerInstancePortListPTemplateSipShared                          int                                           `json:"p-template-sip-shared,omitempty"`
	VirtualServerInstancePortListPacketCaptureTemplate                       string                                        `json:"packet-capture-template,omitempty"`
	VirtualServerInstancePortListPersistType                                 string                                        `json:"persist-type,omitempty"`
	VirtualServerInstancePortListPool                                        string                                        `json:"pool,omitempty"`
	VirtualServerInstancePortListPoolShared                                  string                                        `json:"pool-shared,omitempty"`
	VirtualServerInstancePortListPortNumber                                  int                                           `json:"port-number,omitempty"`
	VirtualServerInstancePortListPortTranslation                             int                                           `json:"port-translation,omitempty"`
	VirtualServerInstancePortListPrecedence                                  int                                           `json:"precedence,omitempty"`
	VirtualServerInstancePortListProtocol                                    string                                        `json:"protocol,omitempty"`
	VirtualServerInstancePortListProxyLayer                                  string                                        `json:"proxy-layer,omitempty"`
	VirtualServerInstancePortListRange                                       int                                           `json:"range,omitempty"`
	VirtualServerInstancePortListRate                                        int                                           `json:"rate,omitempty"`
	VirtualServerInstancePortListRedirectToHTTPS                             int                                           `json:"redirect-to-https,omitempty"`
	VirtualServerInstancePortListReplyAcmeChallenge                          int                                           `json:"reply-acme-challenge,omitempty"`
	VirtualServerInstancePortListReqFail                                     int                                           `json:"req-fail,omitempty"`
	VirtualServerInstancePortListReselection                                 string                                        `json:"reselection,omitempty"`
	VirtualServerInstancePortListReset                                       int                                           `json:"reset,omitempty"`
	VirtualServerInstancePortListResetOnServerSelectionFail                  int                                           `json:"reset-on-server-selection-fail,omitempty"`
	VirtualServerInstancePortListResolveWebCatList                           string                                        `json:"resolve-web-cat-list,omitempty"`
	VirtualServerInstancePortListRtpSipCallIDMatch                           int                                           `json:"rtp-sip-call-id-match,omitempty"`
	VirtualServerInstancePortListSamplingEnableCounters1                     []VirtualServerInstancePortListSamplingEnable `json:"sampling-enable,omitempty"`
	VirtualServerInstancePortListScaleoutBucketCount                         int                                           `json:"scaleout-bucket-count,omitempty"`
	VirtualServerInstancePortListScaleoutDeviceGroup                         int                                           `json:"scaleout-device-group,omitempty"`
	VirtualServerInstancePortListSecs                                        int                                           `json:"secs,omitempty"`
	VirtualServerInstancePortListServSelFail                                 int                                           `json:"serv-sel-fail,omitempty"`
	VirtualServerInstancePortListServerGroup                                 string                                        `json:"server-group,omitempty"`
	VirtualServerInstancePortListServiceGroup                                string                                        `json:"service-group,omitempty"`
	VirtualServerInstancePortListSharedPartitionCacheTemplate                int                                           `json:"shared-partition-cache-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionClientSslTemplate            int                                           `json:"shared-partition-client-ssl-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionConnectionReuseTemplate      int                                           `json:"shared-partition-connection-reuse-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionDNSTemplate                  int                                           `json:"shared-partition-dns-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionDblbTemplate                 int                                           `json:"shared-partition-dblb-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionDiameterTemplate             int                                           `json:"shared-partition-diameter-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionDohTemplate                  int                                           `json:"shared-partition-doh-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionDynamicServiceTemplate       int                                           `json:"shared-partition-dynamic-service-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionExternalServiceTemplate      int                                           `json:"shared-partition-external-service-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionFixTemplate                  int                                           `json:"shared-partition-fix-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionHTTPPolicyTemplate           int                                           `json:"shared-partition-http-policy-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionHTTPTemplate                 int                                           `json:"shared-partition-http-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionImapPop3Template             int                                           `json:"shared-partition-imap-pop3-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionPersistCookieTemplate        int                                           `json:"shared-partition-persist-cookie-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionPersistDestinationIPTemplate int                                           `json:"shared-partition-persist-destination-ip-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionPersistSourceIPTemplate      int                                           `json:"shared-partition-persist-source-ip-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionPersistSslSidTemplate        int                                           `json:"shared-partition-persist-ssl-sid-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionPolicyTemplate               int                                           `json:"shared-partition-policy-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionPool                         int                                           `json:"shared-partition-pool,omitempty"`
	VirtualServerInstancePortListSharedPartitionSMTPTemplate                 int                                           `json:"shared-partition-smtp-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionServerSslTemplate            int                                           `json:"shared-partition-server-ssl-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionSmppTemplate                 int                                           `json:"shared-partition-smpp-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionTCP                          int                                           `json:"shared-partition-tcp,omitempty"`
	VirtualServerInstancePortListSharedPartitionTCPProxyTemplate             int                                           `json:"shared-partition-tcp-proxy-template,omitempty"`
	VirtualServerInstancePortListSharedPartitionUDP                          int                                           `json:"shared-partition-udp,omitempty"`
	VirtualServerInstancePortListSharedPartitionVirtualPortTemplate          int                                           `json:"shared-partition-virtual-port-template,omitempty"`
	VirtualServerInstancePortListSkipRevHash                                 int                                           `json:"skip-rev-hash,omitempty"`
	VirtualServerInstancePortListSnatOnVip                                   int                                           `json:"snat-on-vip,omitempty"`
	VirtualServerInstancePortListStatsDataAction                             string                                        `json:"stats-data-action,omitempty"`
	VirtualServerInstancePortListSubstituteSourceMac                         int                                           `json:"substitute-source-mac,omitempty"`
	VirtualServerInstancePortListSupportHTTP2                                int                                           `json:"support-http2,omitempty"`
	VirtualServerInstancePortListSynCookie                                   int                                           `json:"syn-cookie,omitempty"`
	VirtualServerInstancePortListTemplateCache                               string                                        `json:"template-cache,omitempty"`
	VirtualServerInstancePortListTemplateCacheShared                         string                                        `json:"template-cache-shared,omitempty"`
	VirtualServerInstancePortListTemplateClientSSH                           string                                        `json:"template-client-ssh,omitempty"`
	VirtualServerInstancePortListTemplateClientSsl                           string                                        `json:"template-client-ssl,omitempty"`
	VirtualServerInstancePortListTemplateClientSslShared                     string                                        `json:"template-client-ssl-shared,omitempty"`
	VirtualServerInstancePortListTemplateConnectionReuse                     string                                        `json:"template-connection-reuse,omitempty"`
	VirtualServerInstancePortListTemplateConnectionReuseShared               string                                        `json:"template-connection-reuse-shared,omitempty"`
	VirtualServerInstancePortListTemplateDNS                                 string                                        `json:"template-dns,omitempty"`
	VirtualServerInstancePortListTemplateDNSShared                           string                                        `json:"template-dns-shared,omitempty"`
	VirtualServerInstancePortListTemplateDblb                                string                                        `json:"template-dblb,omitempty"`
	VirtualServerInstancePortListTemplateDblbShared                          string                                        `json:"template-dblb-shared,omitempty"`
	VirtualServerInstancePortListTemplateDiameter                            string                                        `json:"template-diameter,omitempty"`
	VirtualServerInstancePortListTemplateDiameterShared                      string                                        `json:"template-diameter-shared,omitempty"`
	VirtualServerInstancePortListTemplateDoh                                 string                                        `json:"template-doh,omitempty"`
	VirtualServerInstancePortListTemplateDohShared                           string                                        `json:"template-doh-shared,omitempty"`
	VirtualServerInstancePortListTemplateDynamicService                      string                                        `json:"template-dynamic-service,omitempty"`
	VirtualServerInstancePortListTemplateDynamicServiceShared                string                                        `json:"template-dynamic-service-shared,omitempty"`
	VirtualServerInstancePortListTemplateExternalService                     string                                        `json:"template-external-service,omitempty"`
	VirtualServerInstancePortListTemplateExternalServiceShared               string                                        `json:"template-external-service-shared,omitempty"`
	VirtualServerInstancePortListTemplateFix                                 string                                        `json:"template-fix,omitempty"`
	VirtualServerInstancePortListTemplateFixShared                           string                                        `json:"template-fix-shared,omitempty"`
	VirtualServerInstancePortListTemplateFtp                                 string                                        `json:"template-ftp,omitempty"`
	VirtualServerInstancePortListTemplateHTTP                                string                                        `json:"template-http,omitempty"`
	VirtualServerInstancePortListTemplateHTTPPolicy                          string                                        `json:"template-http-policy,omitempty"`
	VirtualServerInstancePortListTemplateHTTPPolicyShared                    string                                        `json:"template-http-policy-shared,omitempty"`
	VirtualServerInstancePortListTemplateHTTPShared                          string                                        `json:"template-http-shared,omitempty"`
	VirtualServerInstancePortListTemplateImapPop3                            string                                        `json:"template-imap-pop3,omitempty"`
	VirtualServerInstancePortListTemplateImapPop3Shared                      string                                        `json:"template-imap-pop3-shared,omitempty"`
	VirtualServerInstancePortListTemplateMqtt                                string                                        `json:"template-mqtt,omitempty"`
	VirtualServerInstancePortListTemplatePersistCookie                       string                                        `json:"template-persist-cookie,omitempty"`
	VirtualServerInstancePortListTemplatePersistCookieShared                 string                                        `json:"template-persist-cookie-shared,omitempty"`
	VirtualServerInstancePortListTemplatePersistDestinationIP                string                                        `json:"template-persist-destination-ip,omitempty"`
	VirtualServerInstancePortListTemplatePersistDestinationIPShared          string                                        `json:"template-persist-destination-ip-shared,omitempty"`
	VirtualServerInstancePortListTemplatePersistSourceIP                     string                                        `json:"template-persist-source-ip,omitempty"`
	VirtualServerInstancePortListTemplatePersistSourceIPShared               string                                        `json:"template-persist-source-ip-shared,omitempty"`
	VirtualServerInstancePortListTemplatePersistSslSid                       string                                        `json:"template-persist-ssl-sid,omitempty"`
	VirtualServerInstancePortListTemplatePersistSslSidShared                 string                                        `json:"template-persist-ssl-sid-shared,omitempty"`
	VirtualServerInstancePortListTemplatePolicy                              string                                        `json:"template-policy,omitempty"`
	VirtualServerInstancePortListTemplatePolicyShared                        string                                        `json:"template-policy-shared,omitempty"`
	VirtualServerInstancePortListTemplateRAMCache                            string                                        `json:"template-ram-cache,omitempty"`
	VirtualServerInstancePortListTemplateReqmodIcap                          string                                        `json:"template-reqmod-icap,omitempty"`
	VirtualServerInstancePortListTemplateRespmodIcap                         string                                        `json:"template-respmod-icap,omitempty"`
	VirtualServerInstancePortListTemplateSMTP                                string                                        `json:"template-smtp,omitempty"`
	VirtualServerInstancePortListTemplateSMTPShared                          string                                        `json:"template-smtp-shared,omitempty"`
	VirtualServerInstancePortListTemplateScaleout                            string                                        `json:"template-scaleout,omitempty"`
	VirtualServerInstancePortListTemplateServerSSH                           string                                        `json:"template-server-ssh,omitempty"`
	VirtualServerInstancePortListTemplateServerSsl                           string                                        `json:"template-server-ssl,omitempty"`
	VirtualServerInstancePortListTemplateServerSslShared                     string                                        `json:"template-server-ssl-shared,omitempty"`
	VirtualServerInstancePortListTemplateSip                                 string                                        `json:"template-sip,omitempty"`
	VirtualServerInstancePortListTemplateSipShared                           string                                        `json:"template-sip-shared,omitempty"`
	VirtualServerInstancePortListTemplateSmpp                                string                                        `json:"template-smpp,omitempty"`
	VirtualServerInstancePortListTemplateSmppShared                          string                                        `json:"template-smpp-shared,omitempty"`
	VirtualServerInstancePortListTemplateSsli                                string                                        `json:"template-ssli,omitempty"`
	VirtualServerInstancePortListTemplateTCP                                 string                                        `json:"template-tcp,omitempty"`
	VirtualServerInstancePortListTemplateTCPProxy                            string                                        `json:"template-tcp-proxy,omitempty"`
	VirtualServerInstancePortListTemplateTCPProxyClient                      string                                        `json:"template-tcp-proxy-client,omitempty"`
	VirtualServerInstancePortListTemplateTCPProxyServer                      string                                        `json:"template-tcp-proxy-server,omitempty"`
	VirtualServerInstancePortListTemplateTCPProxyShared                      string                                        `json:"template-tcp-proxy-shared,omitempty"`
	VirtualServerInstancePortListTemplateTCPShared                           string                                        `json:"template-tcp-shared,omitempty"`
	VirtualServerInstancePortListTemplateUDP                                 string                                        `json:"template-udp,omitempty"`
	VirtualServerInstancePortListTemplateUDPShared                           string                                        `json:"template-udp-shared,omitempty"`
	VirtualServerInstancePortListTemplateVirtualPort                         string                                        `json:"template-virtual-port,omitempty"`
	VirtualServerInstancePortListTemplateVirtualPortShared                   string                                        `json:"template-virtual-port-shared,omitempty"`
	VirtualServerInstancePortListTrunkFwd                                    int                                           `json:"trunk-fwd,omitempty"`
	VirtualServerInstancePortListTrunkRev                                    int                                           `json:"trunk-rev,omitempty"`
	VirtualServerInstancePortListUUID                                        string                                        `json:"uuid,omitempty"`
	VirtualServerInstancePortListUseAlternatePort                            int                                           `json:"use-alternate-port,omitempty"`
	VirtualServerInstancePortListUseCgnv6                                    int                                           `json:"use-cgnv6,omitempty"`
	VirtualServerInstancePortListUseDefaultIfNoServer                        int                                           `json:"use-default-if-no-server,omitempty"`
	VirtualServerInstancePortListUseRcvHopForResp                            int                                           `json:"use-rcv-hop-for-resp,omitempty"`
	VirtualServerInstancePortListUseRcvHopGroup                              int                                           `json:"use-rcv-hop-group,omitempty"`
	VirtualServerInstancePortListUserTag                                     string                                        `json:"user-tag,omitempty"`
	VirtualServerInstancePortListView                                        int                                           `json:"view,omitempty"`
	VirtualServerInstancePortListWafTemplate                                 string                                        `json:"waf-template,omitempty"`
	VirtualServerInstancePortListWhenDown                                    int                                           `json:"when-down,omitempty"`
	VirtualServerInstancePortListWhenDownProtocol2                           int                                           `json:"when-down-protocol2,omitempty"`
}

type VirtualServerInstancePortListAclList struct {
	VirtualServerInstancePortListAclListAclID                    int    `json:"acl-id,omitempty"`
	VirtualServerInstancePortListAclListAclIDSeqNum              int    `json:"acl-id-seq-num,omitempty"`
	VirtualServerInstancePortListAclListAclIDSeqNumShared        int    `json:"acl-id-seq-num-shared,omitempty"`
	VirtualServerInstancePortListAclListAclIDShared              int    `json:"acl-id-shared,omitempty"`
	VirtualServerInstancePortListAclListAclIDSrcNatPool          string `json:"acl-id-src-nat-pool,omitempty"`
	VirtualServerInstancePortListAclListAclIDSrcNatPoolShared    string `json:"acl-id-src-nat-pool-shared,omitempty"`
	VirtualServerInstancePortListAclListAclName                  string `json:"acl-name,omitempty"`
	VirtualServerInstancePortListAclListAclNameSeqNum            int    `json:"acl-name-seq-num,omitempty"`
	VirtualServerInstancePortListAclListAclNameSeqNumShared      int    `json:"acl-name-seq-num-shared,omitempty"`
	VirtualServerInstancePortListAclListAclNameShared            string `json:"acl-name-shared,omitempty"`
	VirtualServerInstancePortListAclListAclNameSrcNatPool        string `json:"acl-name-src-nat-pool,omitempty"`
	VirtualServerInstancePortListAclListAclNameSrcNatPoolShared  string `json:"acl-name-src-nat-pool-shared,omitempty"`
	VirtualServerInstancePortListAclListSharedPartitionPoolID    int    `json:"shared-partition-pool-id,omitempty"`
	VirtualServerInstancePortListAclListSharedPartitionPoolName  int    `json:"shared-partition-pool-name,omitempty"`
	VirtualServerInstancePortListAclListVAclIDSeqNum             int    `json:"v-acl-id-seq-num,omitempty"`
	VirtualServerInstancePortListAclListVAclIDSeqNumShared       int    `json:"v-acl-id-seq-num-shared,omitempty"`
	VirtualServerInstancePortListAclListVAclIDSrcNatPool         string `json:"v-acl-id-src-nat-pool,omitempty"`
	VirtualServerInstancePortListAclListVAclIDSrcNatPoolShared   string `json:"v-acl-id-src-nat-pool-shared,omitempty"`
	VirtualServerInstancePortListAclListVAclNameSeqNum           int    `json:"v-acl-name-seq-num,omitempty"`
	VirtualServerInstancePortListAclListVAclNameSeqNumShared     int    `json:"v-acl-name-seq-num-shared,omitempty"`
	VirtualServerInstancePortListAclListVAclNameSrcNatPool       string `json:"v-acl-name-src-nat-pool,omitempty"`
	VirtualServerInstancePortListAclListVAclNameSrcNatPoolShared string `json:"v-acl-name-src-nat-pool-shared,omitempty"`
	VirtualServerInstancePortListAclListVSharedPartitionPoolID   int    `json:"v-shared-partition-pool-id,omitempty"`
	VirtualServerInstancePortListAclListVSharedPartitionPoolName int    `json:"v-shared-partition-pool-name,omitempty"`
}

type VirtualServerInstancePortListAflexScripts struct {
	VirtualServerInstancePortListAflexScriptsAflex       string `json:"aflex,omitempty"`
	VirtualServerInstancePortListAflexScriptsAflexShared string `json:"aflex-shared,omitempty"`
}

type VirtualServerInstancePortListAuthCfg struct {
	VirtualServerInstancePortListAuthCfgAaaPolicy string `json:"aaa-policy,omitempty"`
}

type VirtualServerInstancePortListSamplingEnable struct {
	VirtualServerInstancePortListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

func PostVirtualServer(id string, inst VirtualServer, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostVirtualServer")
	payloadBytes, err := json_marsh.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/virtual-server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetVirtualServer(id string, name1 string, host string) (*VirtualServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetVirtualServer")
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/virtual-server/"+nameEncode, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetVirtualServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutVirtualServer(id string, name1 string, inst VirtualServer, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutVirtualServer")
	payloadBytes, err := json_marsh.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/virtual-server/"+nameEncode, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteVirtualServer(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteVirtualServer")
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/virtual-server/"+nameEncode, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
