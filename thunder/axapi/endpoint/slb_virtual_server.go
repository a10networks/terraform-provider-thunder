

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServer struct {
	Inst struct {

    AclId int `json:"acl-id"`
    AclIdShared int `json:"acl-id-shared"`
    AclName string `json:"acl-name"`
    AclNameShared string `json:"acl-name-shared"`
    ArpDisable int `json:"arp-disable"`
    Description string `json:"description"`
    DisableVipAdv int `json:"disable-vip-adv"`
    EnableDisableAction string `json:"enable-disable-action" dval:"enable"`
    Ethernet int `json:"ethernet"`
    ExtendedStats int `json:"extended-stats"`
    GamingProtocolCompliance int `json:"gaming-protocol-compliance"`
    HaDynamic int `json:"ha-dynamic"`
    IpAddress string `json:"ip-address"`
    Ipv6Acl string `json:"ipv6-acl"`
    Ipv6AclShared string `json:"ipv6-acl-shared"`
    Ipv6Address string `json:"ipv6-address"`
    MigrateVip SlbVirtualServerMigrateVip1476 `json:"migrate-vip"`
    Name string `json:"name"`
    Netmask string `json:"netmask"`
    PortList []SlbVirtualServerPortList `json:"port-list"`
    RedistributeRouteMap string `json:"redistribute-route-map"`
    RedistributionFlagged int `json:"redistribution-flagged"`
    SharedPartitionPolicyTemplate int `json:"shared-partition-policy-template"`
    SharedPartitionVsTemplate int `json:"shared-partition-vs-template"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    SuppressInternalLoopback int `json:"suppress-internal-loopback"`
    TemplateLogging string `json:"template-logging"`
    TemplatePolicy string `json:"template-policy"`
    TemplatePolicyShared string `json:"template-policy-shared"`
    TemplateScaleout string `json:"template-scaleout"`
    TemplateVirtualServer string `json:"template-virtual-server"`
    TemplateVirtualServerShared string `json:"template-virtual-server-shared"`
    UseIfIp int `json:"use-if-ip"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VportDisableAction string `json:"vport-disable-action"`
    Vrid int `json:"vrid"`

	} `json:"virtual-server"`
}


type SlbVirtualServerMigrateVip1476 struct {
    TargetDataCpu int `json:"target-data-cpu"`
    TargetFloatingIpv4 string `json:"target-floating-ipv4"`
    TargetFloatingIpv6 string `json:"target-floating-ipv6"`
    CancelMigration int `json:"cancel-migration"`
    FinishMigration int `json:"finish-migration"`
    Uuid string `json:"uuid"`
}


type SlbVirtualServerPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Range int `json:"range"`
    AlternatePort int `json:"alternate-port"`
    ProxyLayer string `json:"proxy-layer"`
    OptimizationLevel string `json:"optimization-level" dval:"0"`
    SupportHttp2 int `json:"support-http2"`
    IpOnlyLb int `json:"ip-only-lb"`
    Name string `json:"name"`
    ConnLimit int `json:"conn-limit" dval:"64000000"`
    Reset int `json:"reset"`
    NoLogging int `json:"no-logging"`
    UseAlternatePort int `json:"use-alternate-port"`
    AlternatePortNumber int `json:"alternate-port-number"`
    AltProtocol1 string `json:"alt-protocol1"`
    ServSelFail int `json:"serv-sel-fail"`
    WhenDown int `json:"when-down"`
    AltProtocol2 string `json:"alt-protocol2"`
    ReqFail int `json:"req-fail"`
    WhenDownProtocol2 int `json:"when-down-protocol2"`
    Action string `json:"action" dval:"enable"`
    L7ServiceChain int `json:"l7-service-chain"`
    DefSelectionIfPrefFailed string `json:"def-selection-if-pref-failed" dval:"def-selection-if-pref-failed"`
    HaConnMirror int `json:"ha-conn-mirror"`
    OnSyn int `json:"on-syn"`
    SkipRevHash int `json:"skip-rev-hash"`
    MessageSwitching int `json:"message-switching"`
    ForceRoutingMode int `json:"force-routing-mode"`
    OneServerConn int `json:"one-server-conn"`
    Rate int `json:"rate"`
    Secs int `json:"secs"`
    ResetOnServerSelectionFail int `json:"reset-on-server-selection-fail"`
    ClientipStickyNat int `json:"clientip-sticky-nat"`
    ExtendedStats int `json:"extended-stats"`
    GslbEnable int `json:"gslb-enable"`
    View int `json:"view"`
    SnatOnVip int `json:"snat-on-vip"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    SynCookie int `json:"syn-cookie"`
    ShowtechPrintExtendedStats int `json:"showtech-print-extended-stats"`
    Expand int `json:"expand"`
    AttackDetection int `json:"attack-detection"`
    AclList []SlbVirtualServerPortListAclList `json:"acl-list"`
    TemplatePolicy string `json:"template-policy"`
    SharedPartitionPolicyTemplate int `json:"shared-partition-policy-template"`
    TemplatePolicyShared string `json:"template-policy-shared"`
    AflexScripts []SlbVirtualServerPortListAflexScripts `json:"aflex-scripts"`
    NoAutoUpOnAflex int `json:"no-auto-up-on-aflex"`
    EnableScaleout int `json:"enable-scaleout"`
    Pool string `json:"pool"`
    SharedPartitionPool int `json:"shared-partition-pool"`
    PoolShared string `json:"pool-shared"`
    Auto int `json:"auto"`
    Precedence int `json:"precedence"`
    IpSmartRr int `json:"ip-smart-rr"`
    UseCgnv6 int `json:"use-cgnv6"`
    EnablePlayeridCheck int `json:"enable-playerid-check"`
    ServiceGroup string `json:"service-group"`
    Ipinip int `json:"ipinip"`
    IpMapList string `json:"ip-map-list"`
    RtpSipCallIdMatch int `json:"rtp-sip-call-id-match"`
    UseRcvHopForResp int `json:"use-rcv-hop-for-resp"`
    PersistType string `json:"persist-type"`
    UseRcvHopGroup int `json:"use-rcv-hop-group"`
    ServerGroup string `json:"server-group"`
    Reselection string `json:"reselection"`
    EthFwd int `json:"eth-fwd"`
    TrunkFwd int `json:"trunk-fwd"`
    EthRev int `json:"eth-rev"`
    TrunkRev int `json:"trunk-rev"`
    TemplateSip string `json:"template-sip"`
    PTemplateSipShared int `json:"p-template-sip-shared"`
    TemplateSipShared string `json:"template-sip-shared"`
    TemplateSmpp string `json:"template-smpp"`
    SharedPartitionSmppTemplate int `json:"shared-partition-smpp-template"`
    TemplateSmppShared string `json:"template-smpp-shared"`
    TemplateDblb string `json:"template-dblb"`
    SharedPartitionDblbTemplate int `json:"shared-partition-dblb-template"`
    TemplateDblbShared string `json:"template-dblb-shared"`
    TemplateConnectionReuse string `json:"template-connection-reuse"`
    SharedPartitionConnectionReuseTemplate int `json:"shared-partition-connection-reuse-template"`
    TemplateConnectionReuseShared string `json:"template-connection-reuse-shared"`
    TemplateDns string `json:"template-dns"`
    SharedPartitionDnsTemplate int `json:"shared-partition-dns-template"`
    TemplateDnsShared string `json:"template-dns-shared"`
    TemplateDynamicService string `json:"template-dynamic-service"`
    SharedPartitionDynamicServiceTemplate int `json:"shared-partition-dynamic-service-template"`
    TemplateDynamicServiceShared string `json:"template-dynamic-service-shared"`
    TemplatePersistSourceIp string `json:"template-persist-source-ip"`
    SharedPartitionPersistSourceIpTemplate int `json:"shared-partition-persist-source-ip-template"`
    TemplatePersistSourceIpShared string `json:"template-persist-source-ip-shared"`
    TemplatePersistDestinationIp string `json:"template-persist-destination-ip"`
    SharedPartitionPersistDestinationIpTemplate int `json:"shared-partition-persist-destination-ip-template"`
    TemplatePersistDestinationIpShared string `json:"template-persist-destination-ip-shared"`
    TemplatePersistSslSid string `json:"template-persist-ssl-sid"`
    SharedPartitionPersistSslSidTemplate int `json:"shared-partition-persist-ssl-sid-template"`
    TemplatePersistSslSidShared string `json:"template-persist-ssl-sid-shared"`
    TemplatePersistCookie string `json:"template-persist-cookie"`
    SharedPartitionPersistCookieTemplate int `json:"shared-partition-persist-cookie-template"`
    TemplatePersistCookieShared string `json:"template-persist-cookie-shared"`
    TemplateImapPop3 string `json:"template-imap-pop3"`
    SharedPartitionImapPop3Template int `json:"shared-partition-imap-pop3-template"`
    TemplateImapPop3Shared string `json:"template-imap-pop3-shared"`
    TemplateSmtp string `json:"template-smtp"`
    SharedPartitionSmtpTemplate int `json:"shared-partition-smtp-template"`
    TemplateSmtpShared string `json:"template-smtp-shared"`
    TemplateMqtt string `json:"template-mqtt"`
    TemplateHttp string `json:"template-http"`
    SharedPartitionHttpTemplate int `json:"shared-partition-http-template"`
    TemplateHttpShared string `json:"template-http-shared"`
    TemplateHttpPolicy string `json:"template-http-policy"`
    SharedPartitionHttpPolicyTemplate int `json:"shared-partition-http-policy-template"`
    TemplateHttpPolicyShared string `json:"template-http-policy-shared"`
    RedirectToHttps int `json:"redirect-to-https"`
    TemplateExternalService string `json:"template-external-service"`
    SharedPartitionExternalServiceTemplate int `json:"shared-partition-external-service-template"`
    TemplateExternalServiceShared string `json:"template-external-service-shared"`
    TemplateReqmodIcap string `json:"template-reqmod-icap"`
    TemplateRespmodIcap string `json:"template-respmod-icap"`
    TemplateDoh string `json:"template-doh"`
    SharedPartitionDohTemplate int `json:"shared-partition-doh-template"`
    TemplateDohShared string `json:"template-doh-shared"`
    TemplateServerSsl string `json:"template-server-ssl"`
    SharedPartitionServerSslTemplate int `json:"shared-partition-server-ssl-template"`
    TemplateServerSslShared string `json:"template-server-ssl-shared"`
    TemplateClientSsl string `json:"template-client-ssl"`
    SharedPartitionClientSslTemplate int `json:"shared-partition-client-ssl-template"`
    TemplateClientSslShared string `json:"template-client-ssl-shared"`
    TemplateServerSsh string `json:"template-server-ssh"`
    TemplateClientSsh string `json:"template-client-ssh"`
    TemplateUdp string `json:"template-udp" dval:"default"`
    SharedPartitionUdp int `json:"shared-partition-udp"`
    TemplateUdpShared string `json:"template-udp-shared" dval:"default"`
    TemplateTcp string `json:"template-tcp" dval:"default"`
    SharedPartitionTcp int `json:"shared-partition-tcp"`
    TemplateTcpShared string `json:"template-tcp-shared" dval:"default"`
    TemplateVirtualPort string `json:"template-virtual-port" dval:"default"`
    SharedPartitionVirtualPortTemplate int `json:"shared-partition-virtual-port-template"`
    TemplateVirtualPortShared string `json:"template-virtual-port-shared"`
    TemplateFtp string `json:"template-ftp"`
    TemplateDiameter string `json:"template-diameter"`
    SharedPartitionDiameterTemplate int `json:"shared-partition-diameter-template"`
    TemplateDiameterShared string `json:"template-diameter-shared"`
    TemplateCache string `json:"template-cache"`
    SharedPartitionCacheTemplate int `json:"shared-partition-cache-template"`
    TemplateCacheShared string `json:"template-cache-shared"`
    TemplateRamCache string `json:"template-ram-cache"`
    TemplateFix string `json:"template-fix"`
    SharedPartitionFixTemplate int `json:"shared-partition-fix-template"`
    TemplateFixShared string `json:"template-fix-shared"`
    TemplateSsli string `json:"template-ssli"`
    TemplateTcpProxyClient string `json:"template-tcp-proxy-client"`
    TemplateTcpProxyServer string `json:"template-tcp-proxy-server"`
    TemplateTcpProxy string `json:"template-tcp-proxy"`
    SharedPartitionTcpProxyTemplate int `json:"shared-partition-tcp-proxy-template"`
    TemplateTcpProxyShared string `json:"template-tcp-proxy-shared"`
    TemplateQuicClient string `json:"template-quic-client"`
    TemplateQuicServer string `json:"template-quic-server"`
    TemplateQuic string `json:"template-quic"`
    SharedPartitionQuicTemplate int `json:"shared-partition-quic-template"`
    TemplateQuicShared string `json:"template-quic-shared"`
    UseDefaultIfNoServer int `json:"use-default-if-no-server"`
    TemplateScaleout string `json:"template-scaleout"`
    NoDestNat int `json:"no-dest-nat"`
    PortTranslation int `json:"port-translation"`
    L7HardwareAssist int `json:"l7-hardware-assist"`
    AuthCfg SlbVirtualServerPortListAuthCfg `json:"auth-cfg"`
    CpuCompute int `json:"cpu-compute"`
    MemoryCompute int `json:"memory-compute"`
    SubstituteSourceMac int `json:"substitute-source-mac"`
    IgnoreGlobal int `json:"ignore-global"`
    AflexTableEntrySynDisable int `json:"aflex-table-entry-syn-disable"`
    AflexTableEntrySynEnable int `json:"aflex-table-entry-syn-enable"`
    GtpSessionLb int `json:"gtp-session-lb"`
    ReplyAcmeChallenge int `json:"reply-acme-challenge"`
    ResolveWebCatList string `json:"resolve-web-cat-list"`
    NgWaf int `json:"ng-waf"`
    FastPath string `json:"fast-path"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbVirtualServerPortListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type SlbVirtualServerPortListAclList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
    AclIdShared int `json:"acl-id-shared"`
    AclNameShared string `json:"acl-name-shared"`
    AclIdSrcNatPool string `json:"acl-id-src-nat-pool"`
    AclIdSeqNum int `json:"acl-id-seq-num"`
    SharedPartitionPoolId int `json:"shared-partition-pool-id"`
    AclIdSrcNatPoolShared string `json:"acl-id-src-nat-pool-shared"`
    AclIdSeqNumShared int `json:"acl-id-seq-num-shared"`
    VAclIdSrcNatPool string `json:"v-acl-id-src-nat-pool"`
    VAclIdSeqNum int `json:"v-acl-id-seq-num"`
    VSharedPartitionPoolId int `json:"v-shared-partition-pool-id"`
    VAclIdSrcNatPoolShared string `json:"v-acl-id-src-nat-pool-shared"`
    VAclIdSeqNumShared int `json:"v-acl-id-seq-num-shared"`
    AclNameSrcNatPool string `json:"acl-name-src-nat-pool"`
    AclNameSeqNum int `json:"acl-name-seq-num"`
    SharedPartitionPoolName int `json:"shared-partition-pool-name"`
    AclNameSrcNatPoolShared string `json:"acl-name-src-nat-pool-shared"`
    AclNameSeqNumShared int `json:"acl-name-seq-num-shared"`
    VAclNameSrcNatPool string `json:"v-acl-name-src-nat-pool"`
    VAclNameSeqNum int `json:"v-acl-name-seq-num"`
    VSharedPartitionPoolName int `json:"v-shared-partition-pool-name"`
    VAclNameSrcNatPoolShared string `json:"v-acl-name-src-nat-pool-shared"`
    VAclNameSeqNumShared int `json:"v-acl-name-seq-num-shared"`
}


type SlbVirtualServerPortListAflexScripts struct {
    Aflex string `json:"aflex"`
    AflexShared string `json:"aflex-shared"`
}


type SlbVirtualServerPortListAuthCfg struct {
    AaaPolicy string `json:"aaa-policy"`
}


type SlbVirtualServerPortListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbVirtualServer) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbVirtualServer) getPath() string{
    return "slb/virtual-server"
}

func (p *SlbVirtualServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServer::Post")
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

func (p *SlbVirtualServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServer::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SlbVirtualServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServer::Put")
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

func (p *SlbVirtualServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
