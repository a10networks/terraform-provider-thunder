

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplate struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerSysObjStatsChange VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChange2360 `json:"trigger-sys-obj-stats-change"`
    TriggerSysObjStatsSeverity VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsSeverity2616 `json:"trigger-sys-obj-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"template"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChange2360 struct {
    Uuid string `json:"uuid"`
    SystemCtrLibAcct VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcct2361 `json:"system-ctr-lib-acct"`
    SystemHardwareAccelerate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate2364 `json:"system-hardware-accelerate"`
    SystemRadiusServer VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServer2367 `json:"system-radius-server"`
    SystemIpThreatList VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatList2370 `json:"system-ip-threat-list"`
    SystemFpgaDrop VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop2373 `json:"system-fpga-drop"`
    SystemDpdkStats VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats2376 `json:"system-dpdk-stats"`
    IpAnomalyDrop VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop2379 `json:"ip-anomaly-drop"`
    AamAuthenticationGlobal VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobal2382 `json:"aam-authentication-global"`
    AamRdns VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdns2385 `json:"aam-rdns"`
    AamAuthServerLdap VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdap2388 `json:"aam-auth-server-ldap"`
    AamAuthServerOcsp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcsp2391 `json:"aam-auth-server-ocsp"`
    AamAuthServerRadius VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadius2394 `json:"aam-auth-server-radius"`
    AamAuthServerWin VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin2397 `json:"aam-auth-server-win"`
    AamAuthAccount VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccount2400 `json:"aam-auth-account"`
    AamAuthSamlGlobal VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobal2403 `json:"aam-auth-saml-global"`
    AamAuthRelayKerberos VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberos2406 `json:"aam-auth-relay-kerberos"`
    AamAuthCaptcha VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptcha2409 `json:"aam-auth-captcha"`
    SlbSslError VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslError2412 `json:"slb-ssl-error"`
    SlbSslCertRevoke VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke2415 `json:"slb-ssl-cert-revoke"`
    SlbSslForwardProxy VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy2418 `json:"slb-ssl-forward-proxy"`
    VpnError VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError2421 `json:"vpn-error"`
    Cgnv6Global VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Global2424 `json:"cgnv6-global"`
    Cgnv6DdosProc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc2427 `json:"cgnv6-ddos-proc"`
    Cgnv6Lsn VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn2430 `json:"cgnv6-lsn"`
    Cgnv6LsnAlgEsp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEsp2433 `json:"cgnv6-lsn-alg-esp"`
    Cgnv6LsnAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptp2436 `json:"cgnv6-lsn-alg-pptp"`
    Cgnv6LsnAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2439 `json:"cgnv6-lsn-alg-rtsp"`
    Cgnv6LsnAlgSip VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip2442 `json:"cgnv6-lsn-alg-sip"`
    Cgnv6LsnAlgMgcp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2445 `json:"cgnv6-lsn-alg-mgcp"`
    Cgnv6LsnAlgH323 VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH3232448 `json:"cgnv6-lsn-alg-h323"`
    Cgnv6LsnRadius VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadius2451 `json:"cgnv6-lsn-radius"`
    Cgnv6Nat64Global VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global2454 `json:"cgnv6-nat64-global"`
    Cgnv6DsLiteGlobal VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobal2457 `json:"cgnv6-ds-lite-global"`
    Cgnv6FixedNatGlobal VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal2460 `json:"cgnv6-fixed-nat-global"`
    Cgnv6FixedNatAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2463 `json:"cgnv6-fixed-nat-alg-pptp"`
    Cgnv6FixedNatAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2466 `json:"cgnv6-fixed-nat-alg-rtsp"`
    Cgnv6FixedNatAlgSip VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2469 `json:"cgnv6-fixed-nat-alg-sip"`
    Cgnv6Pcp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp2472 `json:"cgnv6-pcp"`
    Cgnv6Logging VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging2475 `json:"cgnv6-logging"`
    Cgnv6L4 VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L42478 `json:"cgnv6-l4"`
    Cgnv6Icmp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Icmp2481 `json:"cgnv6-icmp"`
    Cgnv6HttpAlg VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg2484 `json:"cgnv6-http-alg"`
    Cgnv6Dns64 VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns642487 `json:"cgnv6-dns64"`
    Cgnv6Dhcpv6 VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv62490 `json:"cgnv6-dhcpv6"`
    FwLogging VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLogging2493 `json:"fw-logging"`
    FwGlobal VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobal2496 `json:"fw-global"`
    FwAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtsp2499 `json:"fw-alg-rtsp"`
    FwAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp2502 `json:"fw-alg-pptp"`
    FwRadServer VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServer2505 `json:"fw-rad-server"`
    FwTcpSynCookie VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie2508 `json:"fw-tcp-syn-cookie"`
    FwDdosProtection VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtection2511 `json:"fw-ddos-protection"`
    FwGtp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp2514 `json:"fw-gtp"`
    SystemTcp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcp2517 `json:"system-tcp"`
    SlbConnReuse VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuse2520 `json:"slb-conn-reuse"`
    SlbAflow VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflow2523 `json:"slb-aflow"`
    SlbFix VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFix2526 `json:"slb-fix"`
    SlbSpdyProxy VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy2529 `json:"slb-spdy-proxy"`
    SlbHttp2 VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp22532 `json:"slb-http2"`
    SlbL7session VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7session2535 `json:"slb-l7session"`
    SlbSmpp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmpp2538 `json:"slb-smpp"`
    SlbSmtp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtp2541 `json:"slb-smtp"`
    SlbMqtt VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqtt2544 `json:"slb-mqtt"`
    SlbIcap VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap2547 `json:"slb-icap"`
    SlbSip VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSip2550 `json:"slb-sip"`
    SlbHwCompress VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompress2553 `json:"slb-hw-compress"`
    SlbMysql VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysql2556 `json:"slb-mysql"`
    SlbMssql VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssql2559 `json:"slb-mssql"`
    SlbCrlSrcip VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcip2562 `json:"slb-crl-srcip"`
    SlbGeneric VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric2565 `json:"slb-generic"`
    SlbPersist VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist2568 `json:"slb-persist"`
    SlbHttpProxy VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxy2571 `json:"slb-http-proxy"`
    SlbL4 VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL42574 `json:"slb-l4"`
    SlbFastHttp VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttp2577 `json:"slb-fast-http"`
    SlbFtpProxy VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxy2580 `json:"slb-ftp-proxy"`
    SlbImapProxy VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy2583 `json:"slb-imap-proxy"`
    SlbPop3Proxy VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3Proxy2586 `json:"slb-pop3-proxy"`
    SlbSwitch VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitch2589 `json:"slb-switch"`
    SlbRcCache VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCache2592 `json:"slb-rc-cache"`
    SoCounters VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters2595 `json:"so-counters"`
    SlbPlyrIdGbl VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGbl2598 `json:"slb-plyr-id-gbl"`
    SlbSportRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRate2601 `json:"slb-sport-rate"`
    LoggingLocalLogGlobal VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobal2604 `json:"logging-local-log-global"`
    SlbMlb VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb2607 `json:"slb-mlb"`
    SlbLinkProbe VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbe2610 `json:"slb-link-probe"`
    SlbRpz VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpz2613 `json:"slb-rpz"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcct2361 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2362 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2363 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2362 struct {
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2363 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerate2364 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2365 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2366 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2365 struct {
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2366 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServer2367 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2368 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2369 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2368 struct {
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RadiusTableFull int `json:"radius-table-full"`
    SecretNotConfiguredDropped int `json:"secret-not-configured-dropped"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    Ipv6PrefixLengthMismatch int `json:"ipv6-prefix-length-mismatch"`
    InvalidKey int `json:"invalid-key"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2369 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RadiusTableFull int `json:"radius-table-full"`
    SecretNotConfiguredDropped int `json:"secret-not-configured-dropped"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    Ipv6PrefixLengthMismatch int `json:"ipv6-prefix-length-mismatch"`
    InvalidKey int `json:"invalid-key"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatList2370 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2371 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2372 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2371 struct {
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2372 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDrop2373 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2374 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2375 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2374 struct {
    MrxDrop int `json:"mrx-drop"`
    HrxDrop int `json:"hrx-drop"`
    SizDrop int `json:"siz-drop"`
    FcsDrop int `json:"fcs-drop"`
    LandDrop int `json:"land-drop"`
    EmptyFragDrop int `json:"empty-frag-drop"`
    MicFragDrop int `json:"mic-frag-drop"`
    Ipv4OptDrop int `json:"ipv4-opt-drop"`
    Ipv4Frag int `json:"ipv4-frag"`
    BadIpHdrLen int `json:"bad-ip-hdr-len"`
    BadIpFlagsDrop int `json:"bad-ip-flags-drop"`
    BadIpTtlDrop int `json:"bad-ip-ttl-drop"`
    NoIpPayloadDrop int `json:"no-ip-payload-drop"`
    OversizeIpPayload int `json:"oversize-ip-payload"`
    BadIpPayloadLen int `json:"bad-ip-payload-len"`
    BadIpFragOffset int `json:"bad-ip-frag-offset"`
    BadIpChksumDrop int `json:"bad-ip-chksum-drop"`
    IcmpPodDrop int `json:"icmp-pod-drop"`
    TcpBadUrgOffet int `json:"tcp-bad-urg-offet"`
    TcpShortHdr int `json:"tcp-short-hdr"`
    TcpBadIpLen int `json:"tcp-bad-ip-len"`
    TcpNullFlags int `json:"tcp-null-flags"`
    TcpNullScan int `json:"tcp-null-scan"`
    TcpFinSin int `json:"tcp-fin-sin"`
    TcpXmasFlags int `json:"tcp-xmas-flags"`
    TcpXmasScan int `json:"tcp-xmas-scan"`
    TcpSynFrag int `json:"tcp-syn-frag"`
    TcpFragHdr int `json:"tcp-frag-hdr"`
    TcpBadChksum int `json:"tcp-bad-chksum"`
    UdpShortHdr int `json:"udp-short-hdr"`
    UdpBadIpLen int `json:"udp-bad-ip-len"`
    UdpKbFrags int `json:"udp-kb-frags"`
    UdpPortLb int `json:"udp-port-lb"`
    UdpBadChksum int `json:"udp-bad-chksum"`
    RuntIpHdr int `json:"runt-ip-hdr"`
    RuntTcpudpHdr int `json:"runt-tcpudp-hdr"`
    TunMismatch int `json:"tun-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2375 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MrxDrop int `json:"mrx-drop"`
    HrxDrop int `json:"hrx-drop"`
    SizDrop int `json:"siz-drop"`
    FcsDrop int `json:"fcs-drop"`
    LandDrop int `json:"land-drop"`
    EmptyFragDrop int `json:"empty-frag-drop"`
    MicFragDrop int `json:"mic-frag-drop"`
    Ipv4OptDrop int `json:"ipv4-opt-drop"`
    Ipv4Frag int `json:"ipv4-frag"`
    BadIpHdrLen int `json:"bad-ip-hdr-len"`
    BadIpFlagsDrop int `json:"bad-ip-flags-drop"`
    BadIpTtlDrop int `json:"bad-ip-ttl-drop"`
    NoIpPayloadDrop int `json:"no-ip-payload-drop"`
    OversizeIpPayload int `json:"oversize-ip-payload"`
    BadIpPayloadLen int `json:"bad-ip-payload-len"`
    BadIpFragOffset int `json:"bad-ip-frag-offset"`
    BadIpChksumDrop int `json:"bad-ip-chksum-drop"`
    IcmpPodDrop int `json:"icmp-pod-drop"`
    TcpBadUrgOffet int `json:"tcp-bad-urg-offet"`
    TcpShortHdr int `json:"tcp-short-hdr"`
    TcpBadIpLen int `json:"tcp-bad-ip-len"`
    TcpNullFlags int `json:"tcp-null-flags"`
    TcpNullScan int `json:"tcp-null-scan"`
    TcpFinSin int `json:"tcp-fin-sin"`
    TcpXmasFlags int `json:"tcp-xmas-flags"`
    TcpXmasScan int `json:"tcp-xmas-scan"`
    TcpSynFrag int `json:"tcp-syn-frag"`
    TcpFragHdr int `json:"tcp-frag-hdr"`
    TcpBadChksum int `json:"tcp-bad-chksum"`
    UdpShortHdr int `json:"udp-short-hdr"`
    UdpBadIpLen int `json:"udp-bad-ip-len"`
    UdpKbFrags int `json:"udp-kb-frags"`
    UdpPortLb int `json:"udp-port-lb"`
    UdpBadChksum int `json:"udp-bad-chksum"`
    RuntIpHdr int `json:"runt-ip-hdr"`
    RuntTcpudpHdr int `json:"runt-tcpudp-hdr"`
    TunMismatch int `json:"tun-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStats2376 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2377 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2378 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2377 struct {
    PktDrop int `json:"pkt-drop"`
    PktLnkDownDrop int `json:"pkt-lnk-down-drop"`
    ErrPktDrop int `json:"err-pkt-drop"`
    RxErr int `json:"rx-err"`
    TxErr int `json:"tx-err"`
    TxDrop int `json:"tx-drop"`
    RxLenErr int `json:"rx-len-err"`
    RxOverErr int `json:"rx-over-err"`
    RxCrcErr int `json:"rx-crc-err"`
    RxFrameErr int `json:"rx-frame-err"`
    RxNoBuffErr int `json:"rx-no-buff-err"`
    RxMissErr int `json:"rx-miss-err"`
    TxAbortErr int `json:"tx-abort-err"`
    TxCarrierErr int `json:"tx-carrier-err"`
    TxFifoErr int `json:"tx-fifo-err"`
    TxHbeatErr int `json:"tx-hbeat-err"`
    TxWindowsErr int `json:"tx-windows-err"`
    RxLongLenErr int `json:"rx-long-len-err"`
    RxShortLenErr int `json:"rx-short-len-err"`
    RxAlignErr int `json:"rx-align-err"`
    RxCsumOffloadErr int `json:"rx-csum-offload-err"`
    IoRxQueDrop int `json:"io-rx-que-drop"`
    IoTxQueDrop int `json:"io-tx-que-drop"`
    IoRingDrop int `json:"io-ring-drop"`
    WTxQueDrop int `json:"w-tx-que-drop"`
    WLinkDownDrop int `json:"w-link-down-drop"`
    WRingDrop int `json:"w-ring-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2378 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    PktDrop int `json:"pkt-drop"`
    PktLnkDownDrop int `json:"pkt-lnk-down-drop"`
    ErrPktDrop int `json:"err-pkt-drop"`
    RxErr int `json:"rx-err"`
    TxErr int `json:"tx-err"`
    TxDrop int `json:"tx-drop"`
    RxLenErr int `json:"rx-len-err"`
    RxOverErr int `json:"rx-over-err"`
    RxCrcErr int `json:"rx-crc-err"`
    RxFrameErr int `json:"rx-frame-err"`
    RxNoBuffErr int `json:"rx-no-buff-err"`
    RxMissErr int `json:"rx-miss-err"`
    TxAbortErr int `json:"tx-abort-err"`
    TxCarrierErr int `json:"tx-carrier-err"`
    TxFifoErr int `json:"tx-fifo-err"`
    TxHbeatErr int `json:"tx-hbeat-err"`
    TxWindowsErr int `json:"tx-windows-err"`
    RxLongLenErr int `json:"rx-long-len-err"`
    RxShortLenErr int `json:"rx-short-len-err"`
    RxAlignErr int `json:"rx-align-err"`
    RxCsumOffloadErr int `json:"rx-csum-offload-err"`
    IoRxQueDrop int `json:"io-rx-que-drop"`
    IoTxQueDrop int `json:"io-tx-que-drop"`
    IoRingDrop int `json:"io-ring-drop"`
    WTxQueDrop int `json:"w-tx-que-drop"`
    WLinkDownDrop int `json:"w-link-down-drop"`
    WRingDrop int `json:"w-ring-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDrop2379 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2380 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2381 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2380 struct {
    Land int `json:"land"`
    Emp_frg int `json:"emp_frg"`
    Emp_mic_frg int `json:"emp_mic_frg"`
    Opt int `json:"opt"`
    Frg int `json:"frg"`
    Bad_ip_hdrlen int `json:"bad_ip_hdrlen"`
    Bad_ip_flg int `json:"bad_ip_flg"`
    Bad_ip_ttl int `json:"bad_ip_ttl"`
    No_ip_payload int `json:"no_ip_payload"`
    Over_ip_payload int `json:"over_ip_payload"`
    Bad_ip_payload_len int `json:"bad_ip_payload_len"`
    Bad_ip_frg_offset int `json:"bad_ip_frg_offset"`
    Csum int `json:"csum"`
    Pod int `json:"pod"`
    Bad_tcp_urg_offset int `json:"bad_tcp_urg_offset"`
    Tcp_sht_hdr int `json:"tcp_sht_hdr"`
    Tcp_bad_iplen int `json:"tcp_bad_iplen"`
    Tcp_null_frg int `json:"tcp_null_frg"`
    Tcp_null_scan int `json:"tcp_null_scan"`
    Tcp_syn_fin int `json:"tcp_syn_fin"`
    Tcp_xmas int `json:"tcp_xmas"`
    Tcp_xmas_scan int `json:"tcp_xmas_scan"`
    Tcp_syn_frg int `json:"tcp_syn_frg"`
    Tcp_frg_hdr int `json:"tcp_frg_hdr"`
    Tcp_bad_csum int `json:"tcp_bad_csum"`
    Udp_srt_hdr int `json:"udp_srt_hdr"`
    Udp_bad_len int `json:"udp_bad_len"`
    Udp_kerb_frg int `json:"udp_kerb_frg"`
    Udp_port_lb int `json:"udp_port_lb"`
    Udp_bad_csum int `json:"udp_bad_csum"`
    Runt_ip_hdr int `json:"runt_ip_hdr"`
    Runt_tcp_udp_hdr int `json:"runt_tcp_udp_hdr"`
    Ipip_tnl_msmtch int `json:"ipip_tnl_msmtch"`
    Tcp_opt_err int `json:"tcp_opt_err"`
    Ipip_tnl_err int `json:"ipip_tnl_err"`
    Vxlan_err int `json:"vxlan_err"`
    Nvgre_err int `json:"nvgre_err"`
    Gre_pptp_err int `json:"gre_pptp_err"`
    Ipv6_eh_hbh int `json:"ipv6_eh_hbh"`
    Ipv6_eh_dest int `json:"ipv6_eh_dest"`
    Ipv6_eh_routing int `json:"ipv6_eh_routing"`
    Ipv6_eh_frag int `json:"ipv6_eh_frag"`
    Ipv6_eh_ah int `json:"ipv6_eh_ah"`
    Ipv6_eh_esp int `json:"ipv6_eh_esp"`
    Ipv6_eh_mobility int `json:"ipv6_eh_mobility"`
    Ipv6_eh_none int `json:"ipv6_eh_none"`
    Ipv6_eh_other int `json:"ipv6_eh_other"`
    Ipv6_eh_malformed int `json:"ipv6_eh_malformed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2381 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Land int `json:"land"`
    Emp_frg int `json:"emp_frg"`
    Emp_mic_frg int `json:"emp_mic_frg"`
    Opt int `json:"opt"`
    Frg int `json:"frg"`
    Bad_ip_hdrlen int `json:"bad_ip_hdrlen"`
    Bad_ip_flg int `json:"bad_ip_flg"`
    Bad_ip_ttl int `json:"bad_ip_ttl"`
    No_ip_payload int `json:"no_ip_payload"`
    Over_ip_payload int `json:"over_ip_payload"`
    Bad_ip_payload_len int `json:"bad_ip_payload_len"`
    Bad_ip_frg_offset int `json:"bad_ip_frg_offset"`
    Csum int `json:"csum"`
    Pod int `json:"pod"`
    Bad_tcp_urg_offset int `json:"bad_tcp_urg_offset"`
    Tcp_sht_hdr int `json:"tcp_sht_hdr"`
    Tcp_bad_iplen int `json:"tcp_bad_iplen"`
    Tcp_null_frg int `json:"tcp_null_frg"`
    Tcp_null_scan int `json:"tcp_null_scan"`
    Tcp_syn_fin int `json:"tcp_syn_fin"`
    Tcp_xmas int `json:"tcp_xmas"`
    Tcp_xmas_scan int `json:"tcp_xmas_scan"`
    Tcp_syn_frg int `json:"tcp_syn_frg"`
    Tcp_frg_hdr int `json:"tcp_frg_hdr"`
    Tcp_bad_csum int `json:"tcp_bad_csum"`
    Udp_srt_hdr int `json:"udp_srt_hdr"`
    Udp_bad_len int `json:"udp_bad_len"`
    Udp_kerb_frg int `json:"udp_kerb_frg"`
    Udp_port_lb int `json:"udp_port_lb"`
    Udp_bad_csum int `json:"udp_bad_csum"`
    Runt_ip_hdr int `json:"runt_ip_hdr"`
    Runt_tcp_udp_hdr int `json:"runt_tcp_udp_hdr"`
    Ipip_tnl_msmtch int `json:"ipip_tnl_msmtch"`
    Tcp_opt_err int `json:"tcp_opt_err"`
    Ipip_tnl_err int `json:"ipip_tnl_err"`
    Vxlan_err int `json:"vxlan_err"`
    Nvgre_err int `json:"nvgre_err"`
    Gre_pptp_err int `json:"gre_pptp_err"`
    Ipv6_eh_hbh int `json:"ipv6_eh_hbh"`
    Ipv6_eh_dest int `json:"ipv6_eh_dest"`
    Ipv6_eh_routing int `json:"ipv6_eh_routing"`
    Ipv6_eh_frag int `json:"ipv6_eh_frag"`
    Ipv6_eh_ah int `json:"ipv6_eh_ah"`
    Ipv6_eh_esp int `json:"ipv6_eh_esp"`
    Ipv6_eh_mobility int `json:"ipv6_eh_mobility"`
    Ipv6_eh_none int `json:"ipv6_eh_none"`
    Ipv6_eh_other int `json:"ipv6_eh_other"`
    Ipv6_eh_malformed int `json:"ipv6_eh_malformed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobal2382 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2383 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2384 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2383 struct {
    Misses int `json:"misses"`
    OpenSocketFailed int `json:"open-socket-failed"`
    ConnectFailed int `json:"connect-failed"`
    CreateTimerFailed int `json:"create-timer-failed"`
    GetSocketOptionFailed int `json:"get-socket-option-failed"`
    AflexAuthzFail int `json:"aflex-authz-fail"`
    AuthnFailure int `json:"authn-failure"`
    AuthzFailure int `json:"authz-failure"`
    DnsResolveFailed int `json:"dns-resolve-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2384 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Misses int `json:"misses"`
    OpenSocketFailed int `json:"open-socket-failed"`
    ConnectFailed int `json:"connect-failed"`
    CreateTimerFailed int `json:"create-timer-failed"`
    GetSocketOptionFailed int `json:"get-socket-option-failed"`
    AflexAuthzFail int `json:"aflex-authz-fail"`
    AuthnFailure int `json:"authn-failure"`
    AuthzFailure int `json:"authz-failure"`
    DnsResolveFailed int `json:"dns-resolve-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdns2385 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2386 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2387 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2386 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2387 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdap2388 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2389 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2390 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2389 struct {
    AdminBindFailure int `json:"admin-bind-failure"`
    BindFailure int `json:"bind-failure"`
    SearchFailure int `json:"search-failure"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    SslSessionFailure int `json:"ssl-session-failure"`
    PwChangeFailure int `json:"pw-change-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2390 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AdminBindFailure int `json:"admin-bind-failure"`
    BindFailure int `json:"bind-failure"`
    SearchFailure int `json:"search-failure"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    SslSessionFailure int `json:"ssl-session-failure"`
    PwChangeFailure int `json:"pw-change-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcsp2391 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2392 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2393 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2392 struct {
    StaplingRequestDropped int `json:"stapling-request-dropped"`
    StaplingResponseFailure int `json:"stapling-response-failure"`
    StaplingResponseError int `json:"stapling-response-error"`
    StaplingResponseTimeout int `json:"stapling-response-timeout"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2393 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StaplingRequestDropped int `json:"stapling-request-dropped"`
    StaplingResponseFailure int `json:"stapling-response-failure"`
    StaplingResponseError int `json:"stapling-response-error"`
    StaplingResponseTimeout int `json:"stapling-response-timeout"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadius2394 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2395 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2396 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2395 struct {
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2396 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWin2397 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2398 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2399 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2398 struct {
    KerberosTimeoutError int `json:"kerberos-timeout-error"`
    KerberosOtherError int `json:"kerberos-other-error"`
    NtlmAuthenticationFailure int `json:"ntlm-authentication-failure"`
    NtlmProtoNegotiationFailure int `json:"ntlm-proto-negotiation-failure"`
    NtlmSessionSetupFailed int `json:"ntlm-session-setup-failed"`
    KerberosRequestDropped int `json:"kerberos-request-dropped"`
    KerberosResponseFailure int `json:"kerberos-response-failure"`
    KerberosResponseError int `json:"kerberos-response-error"`
    KerberosResponseTimeout int `json:"kerberos-response-timeout"`
    KerberosJobStartError int `json:"kerberos-job-start-error"`
    KerberosPollingControlError int `json:"kerberos-polling-control-error"`
    NtlmPrepareReqFailed int `json:"ntlm-prepare-req-failed"`
    NtlmTimeoutError int `json:"ntlm-timeout-error"`
    NtlmOtherError int `json:"ntlm-other-error"`
    NtlmRequestDropped int `json:"ntlm-request-dropped"`
    NtlmResponseFailure int `json:"ntlm-response-failure"`
    NtlmResponseError int `json:"ntlm-response-error"`
    NtlmResponseTimeout int `json:"ntlm-response-timeout"`
    NtlmJobStartError int `json:"ntlm-job-start-error"`
    NtlmPollingControlError int `json:"ntlm-polling-control-error"`
    KerberosPwExpiry int `json:"kerberos-pw-expiry"`
    KerberosPwChangeFailure int `json:"kerberos-pw-change-failure"`
    KerberosValidateKdcFailure int `json:"kerberos-validate-kdc-failure"`
    KerberosGenerateKdcKeytabFailure int `json:"kerberos-generate-kdc-keytab-failure"`
    KerberosDeleteKdcKeytabFailure int `json:"kerberos-delete-kdc-keytab-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2399 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    KerberosTimeoutError int `json:"kerberos-timeout-error"`
    KerberosOtherError int `json:"kerberos-other-error"`
    NtlmAuthenticationFailure int `json:"ntlm-authentication-failure"`
    NtlmProtoNegotiationFailure int `json:"ntlm-proto-negotiation-failure"`
    NtlmSessionSetupFailed int `json:"ntlm-session-setup-failed"`
    KerberosRequestDropped int `json:"kerberos-request-dropped"`
    KerberosResponseFailure int `json:"kerberos-response-failure"`
    KerberosResponseError int `json:"kerberos-response-error"`
    KerberosResponseTimeout int `json:"kerberos-response-timeout"`
    KerberosJobStartError int `json:"kerberos-job-start-error"`
    KerberosPollingControlError int `json:"kerberos-polling-control-error"`
    NtlmPrepareReqFailed int `json:"ntlm-prepare-req-failed"`
    NtlmTimeoutError int `json:"ntlm-timeout-error"`
    NtlmOtherError int `json:"ntlm-other-error"`
    NtlmRequestDropped int `json:"ntlm-request-dropped"`
    NtlmResponseFailure int `json:"ntlm-response-failure"`
    NtlmResponseError int `json:"ntlm-response-error"`
    NtlmResponseTimeout int `json:"ntlm-response-timeout"`
    NtlmJobStartError int `json:"ntlm-job-start-error"`
    NtlmPollingControlError int `json:"ntlm-polling-control-error"`
    KerberosPwExpiry int `json:"kerberos-pw-expiry"`
    KerberosPwChangeFailure int `json:"kerberos-pw-change-failure"`
    KerberosValidateKdcFailure int `json:"kerberos-validate-kdc-failure"`
    KerberosGenerateKdcKeytabFailure int `json:"kerberos-generate-kdc-keytab-failure"`
    KerberosDeleteKdcKeytabFailure int `json:"kerberos-delete-kdc-keytab-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccount2400 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2401 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2402 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2401 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2402 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobal2403 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2404 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2405 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2404 struct {
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2405 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberos2406 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2407 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2408 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2407 struct {
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2408 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptcha2409 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2410 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2411 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2410 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2411 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslError2412 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2413 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2414 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2413 struct {
    AppDataInHandshake int `json:"app-data-in-handshake"`
    AttemptToReuseSessInDiffContext int `json:"attempt-to-reuse-sess-in-diff-context"`
    BadAlertRecord int `json:"bad-alert-record"`
    BadAuthenticationType int `json:"bad-authentication-type"`
    BadChangeCipherSpec int `json:"bad-change-cipher-spec"`
    BadChecksum int `json:"bad-checksum"`
    BadDataReturnedByCallback int `json:"bad-data-returned-by-callback"`
    BadDecompression int `json:"bad-decompression"`
    BadDhGLength int `json:"bad-dh-g-length"`
    BadDhPubKeyLength int `json:"bad-dh-pub-key-length"`
    BadDhPLength int `json:"bad-dh-p-length"`
    BadDigestLength int `json:"bad-digest-length"`
    BadDsaSignature int `json:"bad-dsa-signature"`
    BadHelloRequest int `json:"bad-hello-request"`
    BadLength int `json:"bad-length"`
    BadMacDecode int `json:"bad-mac-decode"`
    BadMessageType int `json:"bad-message-type"`
    BadPacketLength int `json:"bad-packet-length"`
    BadProtocolVersionCounter int `json:"bad-protocol-version-counter"`
    BadResponseArgument int `json:"bad-response-argument"`
    BadRsaDecrypt int `json:"bad-rsa-decrypt"`
    BadRsaEncrypt int `json:"bad-rsa-encrypt"`
    BadRsaELength int `json:"bad-rsa-e-length"`
    BadRsaModulusLength int `json:"bad-rsa-modulus-length"`
    BadRsaSignature int `json:"bad-rsa-signature"`
    BadSignature int `json:"bad-signature"`
    BadSslFiletype int `json:"bad-ssl-filetype"`
    BadSslSessionIdLength int `json:"bad-ssl-session-id-length"`
    BadState int `json:"bad-state"`
    BadWriteRetry int `json:"bad-write-retry"`
    BioNotSet int `json:"bio-not-set"`
    BlockCipherPadIsWrong int `json:"block-cipher-pad-is-wrong"`
    BnLib int `json:"bn-lib"`
    CaDnLengthMismatch int `json:"ca-dn-length-mismatch"`
    CaDnTooLong int `json:"ca-dn-too-long"`
    CcsReceivedEarly int `json:"ccs-received-early"`
    CertificateVerifyFailed int `json:"certificate-verify-failed"`
    CertLengthMismatch int `json:"cert-length-mismatch"`
    ChallengeIsDifferent int `json:"challenge-is-different"`
    CipherCodeWrongLength int `json:"cipher-code-wrong-length"`
    CipherOrHashUnavailable int `json:"cipher-or-hash-unavailable"`
    CipherTableSrcError int `json:"cipher-table-src-error"`
    CompressedLengthTooLong int `json:"compressed-length-too-long"`
    CompressionFailure int `json:"compression-failure"`
    CompressionLibraryError int `json:"compression-library-error"`
    ConnectionIdIsDifferent int `json:"connection-id-is-different"`
    ConnectionTypeNotSet int `json:"connection-type-not-set"`
    DataBetweenCcsAndFinished int `json:"data-between-ccs-and-finished"`
    DataLengthTooLong int `json:"data-length-too-long"`
    DecryptionFailed int `json:"decryption-failed"`
    DecryptionFailedOrBadRecordMac int `json:"decryption-failed-or-bad-record-mac"`
    DhPublicValueLengthIsWrong int `json:"dh-public-value-length-is-wrong"`
    DigestCheckFailed int `json:"digest-check-failed"`
    EncryptedLengthTooLong int `json:"encrypted-length-too-long"`
    ErrorGeneratingTmpRsaKey int `json:"error-generating-tmp-rsa-key"`
    ErrorInReceivedCipherList int `json:"error-in-received-cipher-list"`
    ExcessiveMessageSize int `json:"excessive-message-size"`
    ExtraDataInMessage int `json:"extra-data-in-message"`
    GotAFinBeforeACcs int `json:"got-a-fin-before-a-ccs"`
    HttpsProxyRequest int `json:"https-proxy-request"`
    HttpRequest int `json:"http-request"`
    IllegalPadding int `json:"illegal-padding"`
    InappropriateFallback int `json:"inappropriate-fallback"`
    InvalidChallengeLength int `json:"invalid-challenge-length"`
    InvalidCommand int `json:"invalid-command"`
    InvalidPurpose int `json:"invalid-purpose"`
    InvalidStatusResponse int `json:"invalid-status-response"`
    InvalidTrust int `json:"invalid-trust"`
    KeyArgTooLong int `json:"key-arg-too-long"`
    Krb5 int `json:"krb5"`
    Krb5ClientCcPrincipal int `json:"krb5-client-cc-principal"`
    Krb5ClientGetCred int `json:"krb5-client-get-cred"`
    Krb5ClientInit int `json:"krb5-client-init"`
    Krb5ClientMkReq int `json:"krb5-client-mk-req"`
    Krb5ServerBadTicket int `json:"krb5-server-bad-ticket"`
    Krb5ServerInit int `json:"krb5-server-init"`
    Krb5ServerRdReq int `json:"krb5-server-rd-req"`
    Krb5ServerTktExpired int `json:"krb5-server-tkt-expired"`
    Krb5ServerTktNotYetValid int `json:"krb5-server-tkt-not-yet-valid"`
    Krb5ServerTktSkew int `json:"krb5-server-tkt-skew"`
    LengthMismatch int `json:"length-mismatch"`
    LengthTooShort int `json:"length-too-short"`
    LibraryBug int `json:"library-bug"`
    LibraryHasNoCiphers int `json:"library-has-no-ciphers"`
    MastKeyTooLong int `json:"mast-key-too-long"`
    MessageTooLong int `json:"message-too-long"`
    MissingDhDsaCert int `json:"missing-dh-dsa-cert"`
    MissingDhKey int `json:"missing-dh-key"`
    MissingDhRsaCert int `json:"missing-dh-rsa-cert"`
    MissingDsaSigningCert int `json:"missing-dsa-signing-cert"`
    MissingExportTmpDhKey int `json:"missing-export-tmp-dh-key"`
    MissingExportTmpRsaKey int `json:"missing-export-tmp-rsa-key"`
    MissingRsaCertificate int `json:"missing-rsa-certificate"`
    MissingRsaEncryptingCert int `json:"missing-rsa-encrypting-cert"`
    MissingRsaSigningCert int `json:"missing-rsa-signing-cert"`
    MissingTmpDhKey int `json:"missing-tmp-dh-key"`
    MissingTmpRsaKey int `json:"missing-tmp-rsa-key"`
    MissingTmpRsaPkey int `json:"missing-tmp-rsa-pkey"`
    MissingVerifyMessage int `json:"missing-verify-message"`
    NonSslv2InitialPacket int `json:"non-sslv2-initial-packet"`
    NoCertificatesReturned int `json:"no-certificates-returned"`
    NoCertificateAssigned int `json:"no-certificate-assigned"`
    NoCertificateReturned int `json:"no-certificate-returned"`
    NoCertificateSet int `json:"no-certificate-set"`
    NoCertificateSpecified int `json:"no-certificate-specified"`
    NoCiphersAvailable int `json:"no-ciphers-available"`
    NoCiphersPassed int `json:"no-ciphers-passed"`
    NoCiphersSpecified int `json:"no-ciphers-specified"`
    NoCipherList int `json:"no-cipher-list"`
    NoCipherMatch int `json:"no-cipher-match"`
    NoClientCertReceived int `json:"no-client-cert-received"`
    NoCompressionSpecified int `json:"no-compression-specified"`
    NoMethodSpecified int `json:"no-method-specified"`
    NoPrivatekey int `json:"no-privatekey"`
    NoPrivateKeyAssigned int `json:"no-private-key-assigned"`
    NoProtocolsAvailable int `json:"no-protocols-available"`
    NoPublickey int `json:"no-publickey"`
    NoSharedCipher int `json:"no-shared-cipher"`
    NoVerifyCallback int `json:"no-verify-callback"`
    NullSslCtx int `json:"null-ssl-ctx"`
    NullSslMethodPassed int `json:"null-ssl-method-passed"`
    OldSessionCipherNotReturned int `json:"old-session-cipher-not-returned"`
    PacketLengthTooLong int `json:"packet-length-too-long"`
    PathTooLong int `json:"path-too-long"`
    PeerDidNotReturnACertificate int `json:"peer-did-not-return-a-certificate"`
    PeerError int `json:"peer-error"`
    PeerErrorCertificate int `json:"peer-error-certificate"`
    PeerErrorNoCertificate int `json:"peer-error-no-certificate"`
    PeerErrorNoCipher int `json:"peer-error-no-cipher"`
    PeerErrorUnsupportedCertificateType int `json:"peer-error-unsupported-certificate-type"`
    PreMacLengthTooLong int `json:"pre-mac-length-too-long"`
    ProblemsMappingCipherFunctions int `json:"problems-mapping-cipher-functions"`
    ProtocolIsShutdown int `json:"protocol-is-shutdown"`
    PublicKeyEncryptError int `json:"public-key-encrypt-error"`
    PublicKeyIsNotRsa int `json:"public-key-is-not-rsa"`
    PublicKeyNotRsa int `json:"public-key-not-rsa"`
    ReadBioNotSet int `json:"read-bio-not-set"`
    ReadWrongPacketType int `json:"read-wrong-packet-type"`
    RecordLengthMismatch int `json:"record-length-mismatch"`
    RecordTooLarge int `json:"record-too-large"`
    RecordTooSmall int `json:"record-too-small"`
    RequiredCipherMissing int `json:"required-cipher-missing"`
    ReuseCertLengthNotZero int `json:"reuse-cert-length-not-zero"`
    ReuseCertTypeNotZero int `json:"reuse-cert-type-not-zero"`
    ReuseCipherListNotZero int `json:"reuse-cipher-list-not-zero"`
    ScsvReceivedWhenRenegotiating int `json:"scsv-received-when-renegotiating"`
    SessionIdContextUninitialized int `json:"session-id-context-uninitialized"`
    ShortRead int `json:"short-read"`
    SignatureForNonSigningCertificate int `json:"signature-for-non-signing-certificate"`
    Ssl23DoingSessionIdReuse int `json:"ssl23-doing-session-id-reuse"`
    Ssl2ConnectionIdTooLong int `json:"ssl2-connection-id-too-long"`
    Ssl3SessionIdTooLong int `json:"ssl3-session-id-too-long"`
    Ssl3SessionIdTooShort int `json:"ssl3-session-id-too-short"`
    Sslv3AlertBadCertificate int `json:"sslv3-alert-bad-certificate"`
    Sslv3AlertBadRecordMac int `json:"sslv3-alert-bad-record-mac"`
    Sslv3AlertCertificateExpired int `json:"sslv3-alert-certificate-expired"`
    Sslv3AlertCertificateRevoked int `json:"sslv3-alert-certificate-revoked"`
    Sslv3AlertCertificateUnknown int `json:"sslv3-alert-certificate-unknown"`
    Sslv3AlertDecompressionFailure int `json:"sslv3-alert-decompression-failure"`
    Sslv3AlertHandshakeFailure int `json:"sslv3-alert-handshake-failure"`
    Sslv3AlertIllegalParameter int `json:"sslv3-alert-illegal-parameter"`
    Sslv3AlertNoCertificate int `json:"sslv3-alert-no-certificate"`
    Sslv3AlertPeerErrorCert int `json:"sslv3-alert-peer-error-cert"`
    Sslv3AlertPeerErrorNoCert int `json:"sslv3-alert-peer-error-no-cert"`
    Sslv3AlertPeerErrorNoCipher int `json:"sslv3-alert-peer-error-no-cipher"`
    Sslv3AlertPeerErrorUnsuppCertType int `json:"sslv3-alert-peer-error-unsupp-cert-type"`
    Sslv3AlertUnexpectedMsg int `json:"sslv3-alert-unexpected-msg"`
    Sslv3AlertUnknownRemoteErrType int `json:"sslv3-alert-unknown-remote-err-type"`
    Sslv3AlertUnspportedCert int `json:"sslv3-alert-unspported-cert"`
    SslCtxHasNoDefaultSslVersion int `json:"ssl-ctx-has-no-default-ssl-version"`
    SslHandshakeFailure int `json:"ssl-handshake-failure"`
    SslLibraryHasNoCiphers int `json:"ssl-library-has-no-ciphers"`
    SslSessionIdCallbackFailed int `json:"ssl-session-id-callback-failed"`
    SslSessionIdConflict int `json:"ssl-session-id-conflict"`
    SslSessionIdContextTooLong int `json:"ssl-session-id-context-too-long"`
    SslSessionIdHasBadLength int `json:"ssl-session-id-has-bad-length"`
    SslSessionIdIsDifferent int `json:"ssl-session-id-is-different"`
    Tlsv1AlertAccessDenied int `json:"tlsv1-alert-access-denied"`
    Tlsv1AlertDecodeError int `json:"tlsv1-alert-decode-error"`
    Tlsv1AlertDecryptionFailed int `json:"tlsv1-alert-decryption-failed"`
    Tlsv1AlertDecryptError int `json:"tlsv1-alert-decrypt-error"`
    Tlsv1AlertExportRestriction int `json:"tlsv1-alert-export-restriction"`
    Tlsv1AlertInsufficientSecurity int `json:"tlsv1-alert-insufficient-security"`
    Tlsv1AlertInternalError int `json:"tlsv1-alert-internal-error"`
    Tlsv1AlertNoRenegotiation int `json:"tlsv1-alert-no-renegotiation"`
    Tlsv1AlertProtocolVersion int `json:"tlsv1-alert-protocol-version"`
    Tlsv1AlertRecordOverflow int `json:"tlsv1-alert-record-overflow"`
    Tlsv1AlertUnknownCa int `json:"tlsv1-alert-unknown-ca"`
    Tlsv1AlertUserCancelled int `json:"tlsv1-alert-user-cancelled"`
    TlsClientCertReqWithAnonCipher int `json:"tls-client-cert-req-with-anon-cipher"`
    TlsPeerDidNotRespondWithCertList int `json:"tls-peer-did-not-respond-with-cert-list"`
    TlsRsaEncryptedValueLengthIsWrong int `json:"tls-rsa-encrypted-value-length-is-wrong"`
    TriedToUseUnsupportedCipher int `json:"tried-to-use-unsupported-cipher"`
    UnableToDecodeDhCerts int `json:"unable-to-decode-dh-certs"`
    UnableToExtractPublicKey int `json:"unable-to-extract-public-key"`
    UnableToFindDhParameters int `json:"unable-to-find-dh-parameters"`
    UnableToFindPublicKeyParameters int `json:"unable-to-find-public-key-parameters"`
    UnableToFindSslMethod int `json:"unable-to-find-ssl-method"`
    UnableToLoadSsl2Md5Routines int `json:"unable-to-load-ssl2-md5-routines"`
    UnableToLoadSsl3Md5Routines int `json:"unable-to-load-ssl3-md5-routines"`
    UnableToLoadSsl3Sha1Routines int `json:"unable-to-load-ssl3-sha1-routines"`
    UnexpectedMessage int `json:"unexpected-message"`
    UnexpectedRecord int `json:"unexpected-record"`
    Uninitialized int `json:"uninitialized"`
    UnknownAlertType int `json:"unknown-alert-type"`
    UnknownCertificateType int `json:"unknown-certificate-type"`
    UnknownCipherReturned int `json:"unknown-cipher-returned"`
    UnknownCipherType int `json:"unknown-cipher-type"`
    UnknownKeyExchangeType int `json:"unknown-key-exchange-type"`
    UnknownPkeyType int `json:"unknown-pkey-type"`
    UnknownProtocol int `json:"unknown-protocol"`
    UnknownRemoteErrorType int `json:"unknown-remote-error-type"`
    UnknownSslVersion int `json:"unknown-ssl-version"`
    UnknownState int `json:"unknown-state"`
    UnsupportedCipher int `json:"unsupported-cipher"`
    UnsupportedCompressionAlgorithm int `json:"unsupported-compression-algorithm"`
    UnsupportedOption int `json:"unsupported-option"`
    UnsupportedProtocol int `json:"unsupported-protocol"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    UnsupportedStatusType int `json:"unsupported-status-type"`
    WriteBioNotSet int `json:"write-bio-not-set"`
    WrongCipherReturned int `json:"wrong-cipher-returned"`
    WrongMessageType int `json:"wrong-message-type"`
    WrongCounterOfKeyBits int `json:"wrong-counter-of-key-bits"`
    WrongSignatureLength int `json:"wrong-signature-length"`
    WrongSignatureSize int `json:"wrong-signature-size"`
    WrongSslVersion int `json:"wrong-ssl-version"`
    WrongVersionCounter int `json:"wrong-version-counter"`
    X509Lib int `json:"x509-lib"`
    X509VerificationSetupProblems int `json:"x509-verification-setup-problems"`
    ClienthelloTlsext int `json:"clienthello-tlsext"`
    ParseTlsext int `json:"parse-tlsext"`
    ServerhelloTlsext int `json:"serverhello-tlsext"`
    Ssl3ExtInvalidServername int `json:"ssl3-ext-invalid-servername"`
    Ssl3ExtInvalidServernameType int `json:"ssl3-ext-invalid-servername-type"`
    MultipleSgcRestarts int `json:"multiple-sgc-restarts"`
    TlsInvalidEcpointformatList int `json:"tls-invalid-ecpointformat-list"`
    BadEccCert int `json:"bad-ecc-cert"`
    BadEcdsaSig int `json:"bad-ecdsa-sig"`
    BadEcpoint int `json:"bad-ecpoint"`
    CookieMismatch int `json:"cookie-mismatch"`
    UnsupportedEllipticCurve int `json:"unsupported-elliptic-curve"`
    NoRequiredDigest int `json:"no-required-digest"`
    UnsupportedDigestType int `json:"unsupported-digest-type"`
    BadHandshakeLength int `json:"bad-handshake-length"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2414 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AppDataInHandshake int `json:"app-data-in-handshake"`
    AttemptToReuseSessInDiffContext int `json:"attempt-to-reuse-sess-in-diff-context"`
    BadAlertRecord int `json:"bad-alert-record"`
    BadAuthenticationType int `json:"bad-authentication-type"`
    BadChangeCipherSpec int `json:"bad-change-cipher-spec"`
    BadChecksum int `json:"bad-checksum"`
    BadDataReturnedByCallback int `json:"bad-data-returned-by-callback"`
    BadDecompression int `json:"bad-decompression"`
    BadDhGLength int `json:"bad-dh-g-length"`
    BadDhPubKeyLength int `json:"bad-dh-pub-key-length"`
    BadDhPLength int `json:"bad-dh-p-length"`
    BadDigestLength int `json:"bad-digest-length"`
    BadDsaSignature int `json:"bad-dsa-signature"`
    BadHelloRequest int `json:"bad-hello-request"`
    BadLength int `json:"bad-length"`
    BadMacDecode int `json:"bad-mac-decode"`
    BadMessageType int `json:"bad-message-type"`
    BadPacketLength int `json:"bad-packet-length"`
    BadProtocolVersionCounter int `json:"bad-protocol-version-counter"`
    BadResponseArgument int `json:"bad-response-argument"`
    BadRsaDecrypt int `json:"bad-rsa-decrypt"`
    BadRsaEncrypt int `json:"bad-rsa-encrypt"`
    BadRsaELength int `json:"bad-rsa-e-length"`
    BadRsaModulusLength int `json:"bad-rsa-modulus-length"`
    BadRsaSignature int `json:"bad-rsa-signature"`
    BadSignature int `json:"bad-signature"`
    BadSslFiletype int `json:"bad-ssl-filetype"`
    BadSslSessionIdLength int `json:"bad-ssl-session-id-length"`
    BadState int `json:"bad-state"`
    BadWriteRetry int `json:"bad-write-retry"`
    BioNotSet int `json:"bio-not-set"`
    BlockCipherPadIsWrong int `json:"block-cipher-pad-is-wrong"`
    BnLib int `json:"bn-lib"`
    CaDnLengthMismatch int `json:"ca-dn-length-mismatch"`
    CaDnTooLong int `json:"ca-dn-too-long"`
    CcsReceivedEarly int `json:"ccs-received-early"`
    CertificateVerifyFailed int `json:"certificate-verify-failed"`
    CertLengthMismatch int `json:"cert-length-mismatch"`
    ChallengeIsDifferent int `json:"challenge-is-different"`
    CipherCodeWrongLength int `json:"cipher-code-wrong-length"`
    CipherOrHashUnavailable int `json:"cipher-or-hash-unavailable"`
    CipherTableSrcError int `json:"cipher-table-src-error"`
    CompressedLengthTooLong int `json:"compressed-length-too-long"`
    CompressionFailure int `json:"compression-failure"`
    CompressionLibraryError int `json:"compression-library-error"`
    ConnectionIdIsDifferent int `json:"connection-id-is-different"`
    ConnectionTypeNotSet int `json:"connection-type-not-set"`
    DataBetweenCcsAndFinished int `json:"data-between-ccs-and-finished"`
    DataLengthTooLong int `json:"data-length-too-long"`
    DecryptionFailed int `json:"decryption-failed"`
    DecryptionFailedOrBadRecordMac int `json:"decryption-failed-or-bad-record-mac"`
    DhPublicValueLengthIsWrong int `json:"dh-public-value-length-is-wrong"`
    DigestCheckFailed int `json:"digest-check-failed"`
    EncryptedLengthTooLong int `json:"encrypted-length-too-long"`
    ErrorGeneratingTmpRsaKey int `json:"error-generating-tmp-rsa-key"`
    ErrorInReceivedCipherList int `json:"error-in-received-cipher-list"`
    ExcessiveMessageSize int `json:"excessive-message-size"`
    ExtraDataInMessage int `json:"extra-data-in-message"`
    GotAFinBeforeACcs int `json:"got-a-fin-before-a-ccs"`
    HttpsProxyRequest int `json:"https-proxy-request"`
    HttpRequest int `json:"http-request"`
    IllegalPadding int `json:"illegal-padding"`
    InappropriateFallback int `json:"inappropriate-fallback"`
    InvalidChallengeLength int `json:"invalid-challenge-length"`
    InvalidCommand int `json:"invalid-command"`
    InvalidPurpose int `json:"invalid-purpose"`
    InvalidStatusResponse int `json:"invalid-status-response"`
    InvalidTrust int `json:"invalid-trust"`
    KeyArgTooLong int `json:"key-arg-too-long"`
    Krb5 int `json:"krb5"`
    Krb5ClientCcPrincipal int `json:"krb5-client-cc-principal"`
    Krb5ClientGetCred int `json:"krb5-client-get-cred"`
    Krb5ClientInit int `json:"krb5-client-init"`
    Krb5ClientMkReq int `json:"krb5-client-mk-req"`
    Krb5ServerBadTicket int `json:"krb5-server-bad-ticket"`
    Krb5ServerInit int `json:"krb5-server-init"`
    Krb5ServerRdReq int `json:"krb5-server-rd-req"`
    Krb5ServerTktExpired int `json:"krb5-server-tkt-expired"`
    Krb5ServerTktNotYetValid int `json:"krb5-server-tkt-not-yet-valid"`
    Krb5ServerTktSkew int `json:"krb5-server-tkt-skew"`
    LengthMismatch int `json:"length-mismatch"`
    LengthTooShort int `json:"length-too-short"`
    LibraryBug int `json:"library-bug"`
    LibraryHasNoCiphers int `json:"library-has-no-ciphers"`
    MastKeyTooLong int `json:"mast-key-too-long"`
    MessageTooLong int `json:"message-too-long"`
    MissingDhDsaCert int `json:"missing-dh-dsa-cert"`
    MissingDhKey int `json:"missing-dh-key"`
    MissingDhRsaCert int `json:"missing-dh-rsa-cert"`
    MissingDsaSigningCert int `json:"missing-dsa-signing-cert"`
    MissingExportTmpDhKey int `json:"missing-export-tmp-dh-key"`
    MissingExportTmpRsaKey int `json:"missing-export-tmp-rsa-key"`
    MissingRsaCertificate int `json:"missing-rsa-certificate"`
    MissingRsaEncryptingCert int `json:"missing-rsa-encrypting-cert"`
    MissingRsaSigningCert int `json:"missing-rsa-signing-cert"`
    MissingTmpDhKey int `json:"missing-tmp-dh-key"`
    MissingTmpRsaKey int `json:"missing-tmp-rsa-key"`
    MissingTmpRsaPkey int `json:"missing-tmp-rsa-pkey"`
    MissingVerifyMessage int `json:"missing-verify-message"`
    NonSslv2InitialPacket int `json:"non-sslv2-initial-packet"`
    NoCertificatesReturned int `json:"no-certificates-returned"`
    NoCertificateAssigned int `json:"no-certificate-assigned"`
    NoCertificateReturned int `json:"no-certificate-returned"`
    NoCertificateSet int `json:"no-certificate-set"`
    NoCertificateSpecified int `json:"no-certificate-specified"`
    NoCiphersAvailable int `json:"no-ciphers-available"`
    NoCiphersPassed int `json:"no-ciphers-passed"`
    NoCiphersSpecified int `json:"no-ciphers-specified"`
    NoCipherList int `json:"no-cipher-list"`
    NoCipherMatch int `json:"no-cipher-match"`
    NoClientCertReceived int `json:"no-client-cert-received"`
    NoCompressionSpecified int `json:"no-compression-specified"`
    NoMethodSpecified int `json:"no-method-specified"`
    NoPrivatekey int `json:"no-privatekey"`
    NoPrivateKeyAssigned int `json:"no-private-key-assigned"`
    NoProtocolsAvailable int `json:"no-protocols-available"`
    NoPublickey int `json:"no-publickey"`
    NoSharedCipher int `json:"no-shared-cipher"`
    NoVerifyCallback int `json:"no-verify-callback"`
    NullSslCtx int `json:"null-ssl-ctx"`
    NullSslMethodPassed int `json:"null-ssl-method-passed"`
    OldSessionCipherNotReturned int `json:"old-session-cipher-not-returned"`
    PacketLengthTooLong int `json:"packet-length-too-long"`
    PathTooLong int `json:"path-too-long"`
    PeerDidNotReturnACertificate int `json:"peer-did-not-return-a-certificate"`
    PeerError int `json:"peer-error"`
    PeerErrorCertificate int `json:"peer-error-certificate"`
    PeerErrorNoCertificate int `json:"peer-error-no-certificate"`
    PeerErrorNoCipher int `json:"peer-error-no-cipher"`
    PeerErrorUnsupportedCertificateType int `json:"peer-error-unsupported-certificate-type"`
    PreMacLengthTooLong int `json:"pre-mac-length-too-long"`
    ProblemsMappingCipherFunctions int `json:"problems-mapping-cipher-functions"`
    ProtocolIsShutdown int `json:"protocol-is-shutdown"`
    PublicKeyEncryptError int `json:"public-key-encrypt-error"`
    PublicKeyIsNotRsa int `json:"public-key-is-not-rsa"`
    PublicKeyNotRsa int `json:"public-key-not-rsa"`
    ReadBioNotSet int `json:"read-bio-not-set"`
    ReadWrongPacketType int `json:"read-wrong-packet-type"`
    RecordLengthMismatch int `json:"record-length-mismatch"`
    RecordTooLarge int `json:"record-too-large"`
    RecordTooSmall int `json:"record-too-small"`
    RequiredCipherMissing int `json:"required-cipher-missing"`
    ReuseCertLengthNotZero int `json:"reuse-cert-length-not-zero"`
    ReuseCertTypeNotZero int `json:"reuse-cert-type-not-zero"`
    ReuseCipherListNotZero int `json:"reuse-cipher-list-not-zero"`
    ScsvReceivedWhenRenegotiating int `json:"scsv-received-when-renegotiating"`
    SessionIdContextUninitialized int `json:"session-id-context-uninitialized"`
    ShortRead int `json:"short-read"`
    SignatureForNonSigningCertificate int `json:"signature-for-non-signing-certificate"`
    Ssl23DoingSessionIdReuse int `json:"ssl23-doing-session-id-reuse"`
    Ssl2ConnectionIdTooLong int `json:"ssl2-connection-id-too-long"`
    Ssl3SessionIdTooLong int `json:"ssl3-session-id-too-long"`
    Ssl3SessionIdTooShort int `json:"ssl3-session-id-too-short"`
    Sslv3AlertBadCertificate int `json:"sslv3-alert-bad-certificate"`
    Sslv3AlertBadRecordMac int `json:"sslv3-alert-bad-record-mac"`
    Sslv3AlertCertificateExpired int `json:"sslv3-alert-certificate-expired"`
    Sslv3AlertCertificateRevoked int `json:"sslv3-alert-certificate-revoked"`
    Sslv3AlertCertificateUnknown int `json:"sslv3-alert-certificate-unknown"`
    Sslv3AlertDecompressionFailure int `json:"sslv3-alert-decompression-failure"`
    Sslv3AlertHandshakeFailure int `json:"sslv3-alert-handshake-failure"`
    Sslv3AlertIllegalParameter int `json:"sslv3-alert-illegal-parameter"`
    Sslv3AlertNoCertificate int `json:"sslv3-alert-no-certificate"`
    Sslv3AlertPeerErrorCert int `json:"sslv3-alert-peer-error-cert"`
    Sslv3AlertPeerErrorNoCert int `json:"sslv3-alert-peer-error-no-cert"`
    Sslv3AlertPeerErrorNoCipher int `json:"sslv3-alert-peer-error-no-cipher"`
    Sslv3AlertPeerErrorUnsuppCertType int `json:"sslv3-alert-peer-error-unsupp-cert-type"`
    Sslv3AlertUnexpectedMsg int `json:"sslv3-alert-unexpected-msg"`
    Sslv3AlertUnknownRemoteErrType int `json:"sslv3-alert-unknown-remote-err-type"`
    Sslv3AlertUnspportedCert int `json:"sslv3-alert-unspported-cert"`
    SslCtxHasNoDefaultSslVersion int `json:"ssl-ctx-has-no-default-ssl-version"`
    SslHandshakeFailure int `json:"ssl-handshake-failure"`
    SslLibraryHasNoCiphers int `json:"ssl-library-has-no-ciphers"`
    SslSessionIdCallbackFailed int `json:"ssl-session-id-callback-failed"`
    SslSessionIdConflict int `json:"ssl-session-id-conflict"`
    SslSessionIdContextTooLong int `json:"ssl-session-id-context-too-long"`
    SslSessionIdHasBadLength int `json:"ssl-session-id-has-bad-length"`
    SslSessionIdIsDifferent int `json:"ssl-session-id-is-different"`
    Tlsv1AlertAccessDenied int `json:"tlsv1-alert-access-denied"`
    Tlsv1AlertDecodeError int `json:"tlsv1-alert-decode-error"`
    Tlsv1AlertDecryptionFailed int `json:"tlsv1-alert-decryption-failed"`
    Tlsv1AlertDecryptError int `json:"tlsv1-alert-decrypt-error"`
    Tlsv1AlertExportRestriction int `json:"tlsv1-alert-export-restriction"`
    Tlsv1AlertInsufficientSecurity int `json:"tlsv1-alert-insufficient-security"`
    Tlsv1AlertInternalError int `json:"tlsv1-alert-internal-error"`
    Tlsv1AlertNoRenegotiation int `json:"tlsv1-alert-no-renegotiation"`
    Tlsv1AlertProtocolVersion int `json:"tlsv1-alert-protocol-version"`
    Tlsv1AlertRecordOverflow int `json:"tlsv1-alert-record-overflow"`
    Tlsv1AlertUnknownCa int `json:"tlsv1-alert-unknown-ca"`
    Tlsv1AlertUserCancelled int `json:"tlsv1-alert-user-cancelled"`
    TlsClientCertReqWithAnonCipher int `json:"tls-client-cert-req-with-anon-cipher"`
    TlsPeerDidNotRespondWithCertList int `json:"tls-peer-did-not-respond-with-cert-list"`
    TlsRsaEncryptedValueLengthIsWrong int `json:"tls-rsa-encrypted-value-length-is-wrong"`
    TriedToUseUnsupportedCipher int `json:"tried-to-use-unsupported-cipher"`
    UnableToDecodeDhCerts int `json:"unable-to-decode-dh-certs"`
    UnableToExtractPublicKey int `json:"unable-to-extract-public-key"`
    UnableToFindDhParameters int `json:"unable-to-find-dh-parameters"`
    UnableToFindPublicKeyParameters int `json:"unable-to-find-public-key-parameters"`
    UnableToFindSslMethod int `json:"unable-to-find-ssl-method"`
    UnableToLoadSsl2Md5Routines int `json:"unable-to-load-ssl2-md5-routines"`
    UnableToLoadSsl3Md5Routines int `json:"unable-to-load-ssl3-md5-routines"`
    UnableToLoadSsl3Sha1Routines int `json:"unable-to-load-ssl3-sha1-routines"`
    UnexpectedMessage int `json:"unexpected-message"`
    UnexpectedRecord int `json:"unexpected-record"`
    Uninitialized int `json:"uninitialized"`
    UnknownAlertType int `json:"unknown-alert-type"`
    UnknownCertificateType int `json:"unknown-certificate-type"`
    UnknownCipherReturned int `json:"unknown-cipher-returned"`
    UnknownCipherType int `json:"unknown-cipher-type"`
    UnknownKeyExchangeType int `json:"unknown-key-exchange-type"`
    UnknownPkeyType int `json:"unknown-pkey-type"`
    UnknownProtocol int `json:"unknown-protocol"`
    UnknownRemoteErrorType int `json:"unknown-remote-error-type"`
    UnknownSslVersion int `json:"unknown-ssl-version"`
    UnknownState int `json:"unknown-state"`
    UnsupportedCipher int `json:"unsupported-cipher"`
    UnsupportedCompressionAlgorithm int `json:"unsupported-compression-algorithm"`
    UnsupportedOption int `json:"unsupported-option"`
    UnsupportedProtocol int `json:"unsupported-protocol"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    UnsupportedStatusType int `json:"unsupported-status-type"`
    WriteBioNotSet int `json:"write-bio-not-set"`
    WrongCipherReturned int `json:"wrong-cipher-returned"`
    WrongMessageType int `json:"wrong-message-type"`
    WrongCounterOfKeyBits int `json:"wrong-counter-of-key-bits"`
    WrongSignatureLength int `json:"wrong-signature-length"`
    WrongSignatureSize int `json:"wrong-signature-size"`
    WrongSslVersion int `json:"wrong-ssl-version"`
    WrongVersionCounter int `json:"wrong-version-counter"`
    X509Lib int `json:"x509-lib"`
    X509VerificationSetupProblems int `json:"x509-verification-setup-problems"`
    ClienthelloTlsext int `json:"clienthello-tlsext"`
    ParseTlsext int `json:"parse-tlsext"`
    ServerhelloTlsext int `json:"serverhello-tlsext"`
    Ssl3ExtInvalidServername int `json:"ssl3-ext-invalid-servername"`
    Ssl3ExtInvalidServernameType int `json:"ssl3-ext-invalid-servername-type"`
    MultipleSgcRestarts int `json:"multiple-sgc-restarts"`
    TlsInvalidEcpointformatList int `json:"tls-invalid-ecpointformat-list"`
    BadEccCert int `json:"bad-ecc-cert"`
    BadEcdsaSig int `json:"bad-ecdsa-sig"`
    BadEcpoint int `json:"bad-ecpoint"`
    CookieMismatch int `json:"cookie-mismatch"`
    UnsupportedEllipticCurve int `json:"unsupported-elliptic-curve"`
    NoRequiredDigest int `json:"no-required-digest"`
    UnsupportedDigestType int `json:"unsupported-digest-type"`
    BadHandshakeLength int `json:"bad-handshake-length"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevoke2415 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2416 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2417 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2416 struct {
    Ocsp_chain_status_revoked int `json:"ocsp_chain_status_revoked"`
    Ocsp_chain_status_unknown int `json:"ocsp_chain_status_unknown"`
    Ocsp_connection_error int `json:"ocsp_connection_error"`
    Ocsp_uri_not_found int `json:"ocsp_uri_not_found"`
    Ocsp_uri_https int `json:"ocsp_uri_https"`
    Ocsp_uri_unsupported int `json:"ocsp_uri_unsupported"`
    Ocsp_response_status_revoked int `json:"ocsp_response_status_revoked"`
    Ocsp_response_status_unknown int `json:"ocsp_response_status_unknown"`
    Ocsp_cache_status_revoked int `json:"ocsp_cache_status_revoked"`
    Ocsp_cache_miss int `json:"ocsp_cache_miss"`
    Ocsp_other_error int `json:"ocsp_other_error"`
    Ocsp_response_no_nonce int `json:"ocsp_response_no_nonce"`
    Ocsp_response_nonce_error int `json:"ocsp_response_nonce_error"`
    Crl_connection_error int `json:"crl_connection_error"`
    Crl_uri_not_found int `json:"crl_uri_not_found"`
    Crl_uri_https int `json:"crl_uri_https"`
    Crl_uri_unsupported int `json:"crl_uri_unsupported"`
    Crl_response_status_revoked int `json:"crl_response_status_revoked"`
    Crl_response_status_unknown int `json:"crl_response_status_unknown"`
    Crl_cache_status_revoked int `json:"crl_cache_status_revoked"`
    Crl_other_error int `json:"crl_other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2417 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ocsp_chain_status_revoked int `json:"ocsp_chain_status_revoked"`
    Ocsp_chain_status_unknown int `json:"ocsp_chain_status_unknown"`
    Ocsp_connection_error int `json:"ocsp_connection_error"`
    Ocsp_uri_not_found int `json:"ocsp_uri_not_found"`
    Ocsp_uri_https int `json:"ocsp_uri_https"`
    Ocsp_uri_unsupported int `json:"ocsp_uri_unsupported"`
    Ocsp_response_status_revoked int `json:"ocsp_response_status_revoked"`
    Ocsp_response_status_unknown int `json:"ocsp_response_status_unknown"`
    Ocsp_cache_status_revoked int `json:"ocsp_cache_status_revoked"`
    Ocsp_cache_miss int `json:"ocsp_cache_miss"`
    Ocsp_other_error int `json:"ocsp_other_error"`
    Ocsp_response_no_nonce int `json:"ocsp_response_no_nonce"`
    Ocsp_response_nonce_error int `json:"ocsp_response_nonce_error"`
    Crl_connection_error int `json:"crl_connection_error"`
    Crl_uri_not_found int `json:"crl_uri_not_found"`
    Crl_uri_https int `json:"crl_uri_https"`
    Crl_uri_unsupported int `json:"crl_uri_unsupported"`
    Crl_response_status_revoked int `json:"crl_response_status_revoked"`
    Crl_response_status_unknown int `json:"crl_response_status_unknown"`
    Crl_cache_status_revoked int `json:"crl_cache_status_revoked"`
    Crl_other_error int `json:"crl_other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxy2418 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2419 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2420 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2419 struct {
    FailedInSslHandshakes int `json:"failed-in-ssl-handshakes"`
    FailedInCryptoOperations int `json:"failed-in-crypto-operations"`
    FailedInTcp int `json:"failed-in-tcp"`
    FailedInCertificateVerification int `json:"failed-in-certificate-verification"`
    FailedInCertificateSigning int `json:"failed-in-certificate-signing"`
    InvalidOcspStaplingResponse int `json:"invalid-ocsp-stapling-response"`
    RevokedOcspResponse int `json:"revoked-ocsp-response"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    ConnectionsFailed int `json:"connections-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2420 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    FailedInSslHandshakes int `json:"failed-in-ssl-handshakes"`
    FailedInCryptoOperations int `json:"failed-in-crypto-operations"`
    FailedInTcp int `json:"failed-in-tcp"`
    FailedInCertificateVerification int `json:"failed-in-certificate-verification"`
    FailedInCertificateSigning int `json:"failed-in-certificate-signing"`
    InvalidOcspStaplingResponse int `json:"invalid-ocsp-stapling-response"`
    RevokedOcspResponse int `json:"revoked-ocsp-response"`
    UnsupportedSslVersion int `json:"unsupported-ssl-version"`
    ConnectionsFailed int `json:"connections-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError2421 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2422 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2423 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2422 struct {
    Bad_opcode int `json:"bad_opcode"`
    Bad_sg_write_len int `json:"bad_sg_write_len"`
    Bad_len int `json:"bad_len"`
    Bad_ipsec_protocol int `json:"bad_ipsec_protocol"`
    Bad_ipsec_auth int `json:"bad_ipsec_auth"`
    Bad_ipsec_padding int `json:"bad_ipsec_padding"`
    Bad_ip_version int `json:"bad_ip_version"`
    Bad_auth_type int `json:"bad_auth_type"`
    Bad_encrypt_type int `json:"bad_encrypt_type"`
    Bad_ipsec_spi int `json:"bad_ipsec_spi"`
    Bad_checksum int `json:"bad_checksum"`
    Bad_ipsec_context int `json:"bad_ipsec_context"`
    Bad_ipsec_context_direction int `json:"bad_ipsec_context_direction"`
    Bad_ipsec_context_flag_mismatch int `json:"bad_ipsec_context_flag_mismatch"`
    Ipcomp_payload int `json:"ipcomp_payload"`
    Bad_selector_match int `json:"bad_selector_match"`
    Bad_fragment_size int `json:"bad_fragment_size"`
    Bad_inline_data int `json:"bad_inline_data"`
    Bad_frag_size_configuration int `json:"bad_frag_size_configuration"`
    Dummy_payload int `json:"dummy_payload"`
    Bad_ip_payload_type int `json:"bad_ip_payload_type"`
    Bad_min_frag_size_auth_sha384_512 int `json:"bad_min_frag_size_auth_sha384_512"`
    Bad_esp_next_header int `json:"bad_esp_next_header"`
    Bad_gre_header int `json:"bad_gre_header"`
    Bad_gre_protocol int `json:"bad_gre_protocol"`
    Ipv6_extension_headers_too_big int `json:"ipv6_extension_headers_too_big"`
    Ipv6_hop_by_hop_error int `json:"ipv6_hop_by_hop_error"`
    Error_ipv6_decrypt_rh_segs_left_error int `json:"error_ipv6_decrypt_rh_segs_left_error"`
    Ipv6_rh_length_error int `json:"ipv6_rh_length_error"`
    Ipv6_outbound_rh_copy_addr_error int `json:"ipv6_outbound_rh_copy_addr_error"`
    Error_ipv6_extension_header_bad int `json:"error_IPv6_extension_header_bad"`
    Bad_encrypt_type_ctr_gcm int `json:"bad_encrypt_type_ctr_gcm"`
    Ah_not_supported_with_gcm_gmac_sha2 int `json:"ah_not_supported_with_gcm_gmac_sha2"`
    Tfc_padding_with_prefrag_not_supported int `json:"tfc_padding_with_prefrag_not_supported"`
    Bad_srtp_auth_tag int `json:"bad_srtp_auth_tag"`
    Bad_ipcomp_configuration int `json:"bad_ipcomp_configuration"`
    Dsiv_incorrect_param int `json:"dsiv_incorrect_param"`
    Bad_ipsec_unknown int `json:"bad_ipsec_unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2423 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Bad_opcode int `json:"bad_opcode"`
    Bad_sg_write_len int `json:"bad_sg_write_len"`
    Bad_len int `json:"bad_len"`
    Bad_ipsec_protocol int `json:"bad_ipsec_protocol"`
    Bad_ipsec_auth int `json:"bad_ipsec_auth"`
    Bad_ipsec_padding int `json:"bad_ipsec_padding"`
    Bad_ip_version int `json:"bad_ip_version"`
    Bad_auth_type int `json:"bad_auth_type"`
    Bad_encrypt_type int `json:"bad_encrypt_type"`
    Bad_ipsec_spi int `json:"bad_ipsec_spi"`
    Bad_checksum int `json:"bad_checksum"`
    Bad_ipsec_context int `json:"bad_ipsec_context"`
    Bad_ipsec_context_direction int `json:"bad_ipsec_context_direction"`
    Bad_ipsec_context_flag_mismatch int `json:"bad_ipsec_context_flag_mismatch"`
    Ipcomp_payload int `json:"ipcomp_payload"`
    Bad_selector_match int `json:"bad_selector_match"`
    Bad_fragment_size int `json:"bad_fragment_size"`
    Bad_inline_data int `json:"bad_inline_data"`
    Bad_frag_size_configuration int `json:"bad_frag_size_configuration"`
    Dummy_payload int `json:"dummy_payload"`
    Bad_ip_payload_type int `json:"bad_ip_payload_type"`
    Bad_min_frag_size_auth_sha384_512 int `json:"bad_min_frag_size_auth_sha384_512"`
    Bad_esp_next_header int `json:"bad_esp_next_header"`
    Bad_gre_header int `json:"bad_gre_header"`
    Bad_gre_protocol int `json:"bad_gre_protocol"`
    Ipv6_extension_headers_too_big int `json:"ipv6_extension_headers_too_big"`
    Ipv6_hop_by_hop_error int `json:"ipv6_hop_by_hop_error"`
    Error_ipv6_decrypt_rh_segs_left_error int `json:"error_ipv6_decrypt_rh_segs_left_error"`
    Ipv6_rh_length_error int `json:"ipv6_rh_length_error"`
    Ipv6_outbound_rh_copy_addr_error int `json:"ipv6_outbound_rh_copy_addr_error"`
    Error_ipv6_extension_header_bad int `json:"error_IPv6_extension_header_bad"`
    Bad_encrypt_type_ctr_gcm int `json:"bad_encrypt_type_ctr_gcm"`
    Ah_not_supported_with_gcm_gmac_sha2 int `json:"ah_not_supported_with_gcm_gmac_sha2"`
    Tfc_padding_with_prefrag_not_supported int `json:"tfc_padding_with_prefrag_not_supported"`
    Bad_srtp_auth_tag int `json:"bad_srtp_auth_tag"`
    Bad_ipcomp_configuration int `json:"bad_ipcomp_configuration"`
    Dsiv_incorrect_param int `json:"dsiv_incorrect_param"`
    Bad_ipsec_unknown int `json:"bad_ipsec_unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Global2424 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsInc2425 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsRate2426 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsInc2425 struct {
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsRate2426 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProc2427 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2428 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2429 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2428 struct {
    L3_entry_match_drop int `json:"l3_entry_match_drop"`
    L3_entry_match_drop_hw int `json:"l3_entry_match_drop_hw"`
    L3_entry_drop_max_hw_exceeded int `json:"l3_entry_drop_max_hw_exceeded"`
    L4_entry_match_drop int `json:"l4_entry_match_drop"`
    L4_entry_match_drop_hw int `json:"l4_entry_match_drop_hw"`
    L4_entry_drop_max_hw_exceeded int `json:"l4_entry_drop_max_hw_exceeded"`
    L4_entry_list_alloc_failure int `json:"l4_entry_list_alloc_failure"`
    Ip_node_alloc_failure int `json:"ip_node_alloc_failure"`
    Ip_port_block_alloc_failure int `json:"ip_port_block_alloc_failure"`
    Ip_other_block_alloc_failure int `json:"ip_other_block_alloc_failure"`
    L3_entry_add_to_bgp_failure int `json:"l3_entry_add_to_bgp_failure"`
    L3_entry_remove_from_bgp_failure int `json:"l3_entry_remove_from_bgp_failure"`
    L3_entry_add_to_hw_failure int `json:"l3_entry_add_to_hw_failure"`
    Syn_cookie_verification_failed int `json:"syn_cookie_verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2429 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    L3_entry_match_drop int `json:"l3_entry_match_drop"`
    L3_entry_match_drop_hw int `json:"l3_entry_match_drop_hw"`
    L3_entry_drop_max_hw_exceeded int `json:"l3_entry_drop_max_hw_exceeded"`
    L4_entry_match_drop int `json:"l4_entry_match_drop"`
    L4_entry_match_drop_hw int `json:"l4_entry_match_drop_hw"`
    L4_entry_drop_max_hw_exceeded int `json:"l4_entry_drop_max_hw_exceeded"`
    L4_entry_list_alloc_failure int `json:"l4_entry_list_alloc_failure"`
    Ip_node_alloc_failure int `json:"ip_node_alloc_failure"`
    Ip_port_block_alloc_failure int `json:"ip_port_block_alloc_failure"`
    Ip_other_block_alloc_failure int `json:"ip_other_block_alloc_failure"`
    L3_entry_add_to_bgp_failure int `json:"l3_entry_add_to_bgp_failure"`
    L3_entry_remove_from_bgp_failure int `json:"l3_entry_remove_from_bgp_failure"`
    L3_entry_add_to_hw_failure int `json:"l3_entry_add_to_hw_failure"`
    Syn_cookie_verification_failed int `json:"syn_cookie_verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Lsn2430 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2431 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2432 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2431 struct {
    User_quota_failure int `json:"user_quota_failure"`
    Data_sesn_user_quota_exceeded int `json:"data_sesn_user_quota_exceeded"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    Sip_alg_quota_inc_failure int `json:"sip_alg_quota_inc_failure"`
    Sip_alg_alloc_rtp_rtcp_port_failure int `json:"sip_alg_alloc_rtp_rtcp_port_failure"`
    Sip_alg_alloc_single_port_failure int `json:"sip_alg_alloc_single_port_failure"`
    Sip_alg_create_single_fullcone_failure int `json:"sip_alg_create_single_fullcone_failure"`
    Sip_alg_create_rtp_fullcone_failure int `json:"sip_alg_create_rtp_fullcone_failure"`
    Sip_alg_create_rtcp_fullcone_failure int `json:"sip_alg_create_rtcp_fullcone_failure"`
    H323_alg_alloc_single_port_failure int `json:"h323_alg_alloc_single_port_failure"`
    H323_alg_create_single_fullcone_failure int `json:"h323_alg_create_single_fullcone_failure"`
    H323_alg_create_rtp_fullcone_failure int `json:"h323_alg_create_rtp_fullcone_failure"`
    H323_alg_create_rtcp_fullcone_failure int `json:"h323_alg_create_rtcp_fullcone_failure"`
    Port_overloading_out_of_memory int `json:"port_overloading_out_of_memory"`
    Port_overloading_inc_overflow int `json:"port_overloading_inc_overflow"`
    Fullcone_ext_mem_alloc_failure int `json:"fullcone_ext_mem_alloc_failure"`
    Fullcone_ext_mem_alloc_init_faulure int `json:"fullcone_ext_mem_alloc_init_faulure"`
    Mgcp_alg_create_rtp_fullcone_failure int `json:"mgcp_alg_create_rtp_fullcone_failure"`
    Mgcp_alg_create_rtcp_fullcone_failure int `json:"mgcp_alg_create_rtcp_fullcone_failure"`
    Mgcp_alg_port_pair_alloc_from_quota_par int `json:"mgcp_alg_port_pair_alloc_from_quota_par"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
    Adc_port_allocation_failed int `json:"adc_port_allocation_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2432 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    User_quota_failure int `json:"user_quota_failure"`
    Data_sesn_user_quota_exceeded int `json:"data_sesn_user_quota_exceeded"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    Sip_alg_quota_inc_failure int `json:"sip_alg_quota_inc_failure"`
    Sip_alg_alloc_rtp_rtcp_port_failure int `json:"sip_alg_alloc_rtp_rtcp_port_failure"`
    Sip_alg_alloc_single_port_failure int `json:"sip_alg_alloc_single_port_failure"`
    Sip_alg_create_single_fullcone_failure int `json:"sip_alg_create_single_fullcone_failure"`
    Sip_alg_create_rtp_fullcone_failure int `json:"sip_alg_create_rtp_fullcone_failure"`
    Sip_alg_create_rtcp_fullcone_failure int `json:"sip_alg_create_rtcp_fullcone_failure"`
    H323_alg_alloc_single_port_failure int `json:"h323_alg_alloc_single_port_failure"`
    H323_alg_create_single_fullcone_failure int `json:"h323_alg_create_single_fullcone_failure"`
    H323_alg_create_rtp_fullcone_failure int `json:"h323_alg_create_rtp_fullcone_failure"`
    H323_alg_create_rtcp_fullcone_failure int `json:"h323_alg_create_rtcp_fullcone_failure"`
    Port_overloading_out_of_memory int `json:"port_overloading_out_of_memory"`
    Port_overloading_inc_overflow int `json:"port_overloading_inc_overflow"`
    Fullcone_ext_mem_alloc_failure int `json:"fullcone_ext_mem_alloc_failure"`
    Fullcone_ext_mem_alloc_init_faulure int `json:"fullcone_ext_mem_alloc_init_faulure"`
    Mgcp_alg_create_rtp_fullcone_failure int `json:"mgcp_alg_create_rtp_fullcone_failure"`
    Mgcp_alg_create_rtcp_fullcone_failure int `json:"mgcp_alg_create_rtcp_fullcone_failure"`
    Mgcp_alg_port_pair_alloc_from_quota_par int `json:"mgcp_alg_port_pair_alloc_from_quota_par"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
    Adc_port_allocation_failed int `json:"adc_port_allocation_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEsp2433 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2434 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2435 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2434 struct {
    NatIpConflict int `json:"nat-ip-conflict"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2435 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NatIpConflict int `json:"nat-ip-conflict"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptp2436 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2437 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2438 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2437 struct {
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2438 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2439 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2440 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2441 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2440 struct {
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2441 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSip2442 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2443 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2444 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2443 struct {
    MethodUnknown int `json:"method-unknown"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2444 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MethodUnknown int `json:"method-unknown"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2445 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2446 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2447 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2446 struct {
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2447 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH3232448 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2449 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2450 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2449 struct {
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2450 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadius2451 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2452 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2453 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2452 struct {
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RequestIgnored int `json:"request-ignored"`
    RadiusTableFull int `json:"radius-table-full"`
    SecretNotConfiguredDropped int `json:"secret-not-configured-dropped"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    InvalidKey int `json:"invalid-key"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2453 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RequestIgnored int `json:"request-ignored"`
    RadiusTableFull int `json:"radius-table-full"`
    SecretNotConfiguredDropped int `json:"secret-not-configured-dropped"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    InvalidKey int `json:"invalid-key"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64Global2454 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2455 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2456 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2455 struct {
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    New_user_resource_unavailable int `json:"new_user_resource_unavailable"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Eif_limit_exceeded int `json:"eif_limit_exceeded"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    No_radius_profile_match int `json:"no_radius_profile_match"`
    No_class_list_match int `json:"no_class_list_match"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2456 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    New_user_resource_unavailable int `json:"new_user_resource_unavailable"`
    Fullcone_failure int `json:"fullcone_failure"`
    Fullcone_self_hairpinning_drop int `json:"fullcone_self_hairpinning_drop"`
    Eif_limit_exceeded int `json:"eif_limit_exceeded"`
    Nat_pool_unusable int `json:"nat_pool_unusable"`
    Ha_nat_pool_unusable int `json:"ha_nat_pool_unusable"`
    Ha_nat_pool_batch_type_mismatch int `json:"ha_nat_pool_batch_type_mismatch"`
    No_radius_profile_match int `json:"no_radius_profile_match"`
    No_class_list_match int `json:"no_class_list_match"`
    User_quota_unusable_drop int `json:"user_quota_unusable_drop"`
    User_quota_unusable int `json:"user_quota_unusable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobal2457 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2458 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2459 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2458 struct {
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    Fullcone_failure int `json:"fullcone_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2459 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    Fullcone_failure int `json:"fullcone_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobal2460 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2461 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2462 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2461 struct {
    NatPortUnavailableTcp int `json:"nat-port-unavailable-tcp"`
    NatPortUnavailableUdp int `json:"nat-port-unavailable-udp"`
    NatPortUnavailableIcmp int `json:"nat-port-unavailable-icmp"`
    SessionUserQuotaExceeded int `json:"session-user-quota-exceeded"`
    FullconeFailure int `json:"fullcone-failure"`
    Nat44InboundFiltered int `json:"nat44-inbound-filtered"`
    Nat64InboundFiltered int `json:"nat64-inbound-filtered"`
    DsliteInboundFiltered int `json:"dslite-inbound-filtered"`
    Nat44EifLimitExceeded int `json:"nat44-eif-limit-exceeded"`
    Nat64EifLimitExceeded int `json:"nat64-eif-limit-exceeded"`
    DsliteEifLimitExceeded int `json:"dslite-eif-limit-exceeded"`
    StandbyDrop int `json:"standby-drop"`
    FixedNatFullconeSelfHairpinningDro int `json:"fixed-nat-fullcone-self-hairpinning-dro"`
    SixrdDrop int `json:"sixrd-drop"`
    DestRlistDrop int `json:"dest-rlist-drop"`
    DestRlistPassThrough int `json:"dest-rlist-pass-through"`
    DestRlistSnatDrop int `json:"dest-rlist-snat-drop"`
    ConfigNotFound int `json:"config-not-found"`
    PortOverloadFailed int `json:"port-overload-failed"`
    HaSessionUserQuotaExceeded int `json:"ha-session-user-quota-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2462 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NatPortUnavailableTcp int `json:"nat-port-unavailable-tcp"`
    NatPortUnavailableUdp int `json:"nat-port-unavailable-udp"`
    NatPortUnavailableIcmp int `json:"nat-port-unavailable-icmp"`
    SessionUserQuotaExceeded int `json:"session-user-quota-exceeded"`
    FullconeFailure int `json:"fullcone-failure"`
    Nat44InboundFiltered int `json:"nat44-inbound-filtered"`
    Nat64InboundFiltered int `json:"nat64-inbound-filtered"`
    DsliteInboundFiltered int `json:"dslite-inbound-filtered"`
    Nat44EifLimitExceeded int `json:"nat44-eif-limit-exceeded"`
    Nat64EifLimitExceeded int `json:"nat64-eif-limit-exceeded"`
    DsliteEifLimitExceeded int `json:"dslite-eif-limit-exceeded"`
    StandbyDrop int `json:"standby-drop"`
    FixedNatFullconeSelfHairpinningDro int `json:"fixed-nat-fullcone-self-hairpinning-dro"`
    SixrdDrop int `json:"sixrd-drop"`
    DestRlistDrop int `json:"dest-rlist-drop"`
    DestRlistPassThrough int `json:"dest-rlist-pass-through"`
    DestRlistSnatDrop int `json:"dest-rlist-snat-drop"`
    ConfigNotFound int `json:"config-not-found"`
    PortOverloadFailed int `json:"port-overload-failed"`
    HaSessionUserQuotaExceeded int `json:"ha-session-user-quota-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2463 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2464 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2465 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2464 struct {
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2465 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2466 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2467 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2468 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2467 struct {
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2468 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2469 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2470 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2471 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2470 struct {
    MethodUnknown int `json:"method-unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2471 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MethodUnknown int `json:"method-unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Pcp2472 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2473 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2474 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2473 struct {
    PktNotRequestDrop int `json:"pkt-not-request-drop"`
    PktTooShortDrop int `json:"pkt-too-short-drop"`
    NorouteDrop int `json:"noroute-drop"`
    UnsupportedVersion int `json:"unsupported-version"`
    NotAuthorized int `json:"not-authorized"`
    MalformRequest int `json:"malform-request"`
    UnsuppOpcode int `json:"unsupp-opcode"`
    UnsuppOption int `json:"unsupp-option"`
    MalformOption int `json:"malform-option"`
    NoResources int `json:"no-resources"`
    UnsuppProtocol int `json:"unsupp-protocol"`
    CannotProvideSuggest int `json:"cannot-provide-suggest"`
    AddressMismatch int `json:"address-mismatch"`
    ExcessiveRemotePeers int `json:"excessive-remote-peers"`
    PktNotFromNatInside int `json:"pkt-not-from-nat-inside"`
    L4ProcessError int `json:"l4-process-error"`
    InternalErrorDrop int `json:"internal-error-drop"`
    Unsol_ance_sent_fail int `json:"unsol_ance_sent_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2474 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    PktNotRequestDrop int `json:"pkt-not-request-drop"`
    PktTooShortDrop int `json:"pkt-too-short-drop"`
    NorouteDrop int `json:"noroute-drop"`
    UnsupportedVersion int `json:"unsupported-version"`
    NotAuthorized int `json:"not-authorized"`
    MalformRequest int `json:"malform-request"`
    UnsuppOpcode int `json:"unsupp-opcode"`
    UnsuppOption int `json:"unsupp-option"`
    MalformOption int `json:"malform-option"`
    NoResources int `json:"no-resources"`
    UnsuppProtocol int `json:"unsupp-protocol"`
    CannotProvideSuggest int `json:"cannot-provide-suggest"`
    AddressMismatch int `json:"address-mismatch"`
    ExcessiveRemotePeers int `json:"excessive-remote-peers"`
    PktNotFromNatInside int `json:"pkt-not-from-nat-inside"`
    L4ProcessError int `json:"l4-process-error"`
    InternalErrorDrop int `json:"internal-error-drop"`
    Unsol_ance_sent_fail int `json:"unsol_ance_sent_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Logging2475 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2476 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2477 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2476 struct {
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2477 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L42478 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2479 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2480 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2479 struct {
    OutOfSessionMemory int `json:"out-of-session-memory"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2480 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    OutOfSessionMemory int `json:"out-of-session-memory"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Icmp2481 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2482 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2483 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2482 struct {
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2483 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlg2484 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2485 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2486 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2485 struct {
    RadiusRequstDropped int `json:"radius-requst-dropped"`
    RadiusResponseDropped int `json:"radius-response-dropped"`
    OutOfMemoryDropped int `json:"out-of-memory-dropped"`
    QueueLenExceedDropped int `json:"queue-len-exceed-dropped"`
    OutOfOrderDropped int `json:"out-of-order-dropped"`
    HeaderInsertionFailed int `json:"header-insertion-failed"`
    HeaderRemovalFailed int `json:"header-removal-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2486 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RadiusRequstDropped int `json:"radius-requst-dropped"`
    RadiusResponseDropped int `json:"radius-response-dropped"`
    OutOfMemoryDropped int `json:"out-of-memory-dropped"`
    QueueLenExceedDropped int `json:"queue-len-exceed-dropped"`
    OutOfOrderDropped int `json:"out-of-order-dropped"`
    HeaderInsertionFailed int `json:"header-insertion-failed"`
    HeaderRemovalFailed int `json:"header-removal-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns642487 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2488 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2489 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2488 struct {
    QueryBadPkt int `json:"query-bad-pkt"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    Drop int `json:"drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2489 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    QueryBadPkt int `json:"query-bad-pkt"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    Drop int `json:"drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv62490 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2491 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2492 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2491 struct {
    PacketsDropped int `json:"packets-dropped"`
    PktsDroppedDuringClear int `json:"pkts-dropped-during-clear"`
    RcvNotSupportedMsg int `json:"rcv-not-supported-msg"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2492 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    PacketsDropped int `json:"packets-dropped"`
    PktsDroppedDuringClear int `json:"pkts-dropped-during-clear"`
    RcvNotSupportedMsg int `json:"rcv-not-supported-msg"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLogging2493 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2494 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2495 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2494 struct {
    LogDropped int `json:"log-dropped"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2495 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobal2496 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2497 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2498 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2497 struct {
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2498 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtsp2499 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2500 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2501 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2500 struct {
    TransportAllocFailure int `json:"transport-alloc-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2501 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TransportAllocFailure int `json:"transport-alloc-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptp2502 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2503 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2504 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2503 struct {
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2504 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServer2505 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2506 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2507 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2506 struct {
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RequestIgnored int `json:"request-ignored"`
    RadiusTableFull int `json:"radius-table-full"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    Ipv6PrefixLengthMismatch int `json:"ipv6-prefix-length-mismatch"`
    InvalidKey int `json:"invalid-key"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2507 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RequestIgnored int `json:"request-ignored"`
    RadiusTableFull int `json:"radius-table-full"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    Ipv6PrefixLengthMismatch int `json:"ipv6-prefix-length-mismatch"`
    InvalidKey int `json:"invalid-key"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookie2508 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2509 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2510 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2509 struct {
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2510 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtection2511 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2512 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2513 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2512 struct {
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2513 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtp2514 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2515 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2516 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsInc2515 struct {
    OutOfSessionMemory int `json:"out-of-session-memory"`
    GtpSmpPathCheckFailed int `json:"gtp-smp-path-check-failed"`
    GtpSmpCheckFailed int `json:"gtp-smp-check-failed"`
    GtpSmpSessionCountCheckFailed int `json:"gtp-smp-session-count-check-failed"`
    GtpCRefCountSmpExceeded int `json:"gtp-c-ref-count-smp-exceeded"`
    GtpUSmpInRmlWithSess int `json:"gtp-u-smp-in-rml-with-sess"`
    GtpTunnelRateLimitEntryCreateFail int `json:"gtp-tunnel-rate-limit-entry-create-fail"`
    GtpRateLimitSmpCreateFailure int `json:"gtp-rate-limit-smp-create-failure"`
    GtpRateLimitT3CtrCreateFailure int `json:"gtp-rate-limit-t3-ctr-create-failure"`
    GtpRateLimitEntryCreateFailure int `json:"gtp-rate-limit-entry-create-failure"`
    GtpSmpDecSessCountCheckFailed int `json:"gtp-smp-dec-sess-count-check-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeFwGtpTriggerStatsRate2516 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    OutOfSessionMemory int `json:"out-of-session-memory"`
    GtpSmpPathCheckFailed int `json:"gtp-smp-path-check-failed"`
    GtpSmpCheckFailed int `json:"gtp-smp-check-failed"`
    GtpSmpSessionCountCheckFailed int `json:"gtp-smp-session-count-check-failed"`
    GtpCRefCountSmpExceeded int `json:"gtp-c-ref-count-smp-exceeded"`
    GtpUSmpInRmlWithSess int `json:"gtp-u-smp-in-rml-with-sess"`
    GtpTunnelRateLimitEntryCreateFail int `json:"gtp-tunnel-rate-limit-entry-create-fail"`
    GtpRateLimitSmpCreateFailure int `json:"gtp-rate-limit-smp-create-failure"`
    GtpRateLimitT3CtrCreateFailure int `json:"gtp-rate-limit-t3-ctr-create-failure"`
    GtpRateLimitEntryCreateFailure int `json:"gtp-rate-limit-entry-create-failure"`
    GtpSmpDecSessCountCheckFailed int `json:"gtp-smp-dec-sess-count-check-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcp2517 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2518 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2519 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2518 struct {
    Attemptfails int `json:"attemptfails"`
    Noroute int `json:"noroute"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2519 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Attemptfails int `json:"attemptfails"`
    Noroute int `json:"noroute"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuse2520 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2521 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2522 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2521 struct {
    Ntermi_err int `json:"ntermi_err"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2522 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ntermi_err int `json:"ntermi_err"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflow2523 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2524 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2525 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2524 struct {
    Pause_conn_fail int `json:"pause_conn_fail"`
    Error_resume_conn int `json:"error_resume_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2525 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Error_resume_conn int `json:"error_resume_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFix2526 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsInc2527 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsRate2528 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsInc2527 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFixTriggerStatsRate2528 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxy2529 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2530 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2531 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2530 struct {
    Tcp_err int `json:"tcp_err"`
    Stream_not_found int `json:"stream_not_found"`
    Stream_err int `json:"stream_err"`
    Session_err int `json:"session_err"`
    Data_no_stream int `json:"data_no_stream"`
    Data_no_stream_no_goaway int `json:"data_no_stream_no_goaway"`
    Data_no_stream_goaway_close int `json:"data_no_stream_goaway_close"`
    Est_cb_no_tuple int `json:"est_cb_no_tuple"`
    Data_cb_no_tuple int `json:"data_cb_no_tuple"`
    Ctx_alloc_fail int `json:"ctx_alloc_fail"`
    Stream_alloc_fail int `json:"stream_alloc_fail"`
    Http_conn_alloc_fail int `json:"http_conn_alloc_fail"`
    Request_header_alloc_fail int `json:"request_header_alloc_fail"`
    Decompress_fail int `json:"decompress_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_version int `json:"invalid_version"`
    Compress_ctx_alloc_fail int `json:"compress_ctx_alloc_fail"`
    Header_compress_fail int `json:"header_compress_fail"`
    Http_err_stream_closed int `json:"http_err_stream_closed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2531 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Tcp_err int `json:"tcp_err"`
    Stream_not_found int `json:"stream_not_found"`
    Stream_err int `json:"stream_err"`
    Session_err int `json:"session_err"`
    Data_no_stream int `json:"data_no_stream"`
    Data_no_stream_no_goaway int `json:"data_no_stream_no_goaway"`
    Data_no_stream_goaway_close int `json:"data_no_stream_goaway_close"`
    Est_cb_no_tuple int `json:"est_cb_no_tuple"`
    Data_cb_no_tuple int `json:"data_cb_no_tuple"`
    Ctx_alloc_fail int `json:"ctx_alloc_fail"`
    Stream_alloc_fail int `json:"stream_alloc_fail"`
    Http_conn_alloc_fail int `json:"http_conn_alloc_fail"`
    Request_header_alloc_fail int `json:"request_header_alloc_fail"`
    Decompress_fail int `json:"decompress_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_version int `json:"invalid_version"`
    Compress_ctx_alloc_fail int `json:"compress_ctx_alloc_fail"`
    Header_compress_fail int `json:"header_compress_fail"`
    Http_err_stream_closed int `json:"http_err_stream_closed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp22532 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2533 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2534 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2533 struct {
    Protocol_error int `json:"protocol_error"`
    Internal_error int `json:"internal_error"`
    Proxy_alloc_error int `json:"proxy_alloc_error"`
    Split_buff_fail int `json:"split_buff_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Error_max_invalid_stream int `json:"error_max_invalid_stream"`
    Data_no_stream int `json:"data_no_stream"`
    Flow_control_error int `json:"flow_control_error"`
    Settings_timeout int `json:"settings_timeout"`
    Frame_size_error int `json:"frame_size_error"`
    Refused_stream int `json:"refused_stream"`
    Cancel int `json:"cancel"`
    Compression_error int `json:"compression_error"`
    Connect_error int `json:"connect_error"`
    Enhance_your_calm int `json:"enhance_your_calm"`
    Inadequate_security int `json:"inadequate_security"`
    Http_1_1_required int `json:"http_1_1_required"`
    Deflate_alloc_fail int `json:"deflate_alloc_fail"`
    Inflate_alloc_fail int `json:"inflate_alloc_fail"`
    Inflate_header_fail int `json:"inflate_header_fail"`
    Bad_connection_preface int `json:"bad_connection_preface"`
    Cant_allocate_control_frame int `json:"cant_allocate_control_frame"`
    Cant_allocate_settings_frame int `json:"cant_allocate_settings_frame"`
    Bad_frame_type_for_stream_state int `json:"bad_frame_type_for_stream_state"`
    Wrong_stream_state int `json:"wrong_stream_state"`
    Data_queue_alloc_error int `json:"data_queue_alloc_error"`
    Buff_alloc_error int `json:"buff_alloc_error"`
    Cant_allocate_rst_frame int `json:"cant_allocate_rst_frame"`
    Cant_allocate_goaway_frame int `json:"cant_allocate_goaway_frame"`
    Cant_allocate_ping_frame int `json:"cant_allocate_ping_frame"`
    Cant_allocate_stream int `json:"cant_allocate_stream"`
    Cant_allocate_window_frame int `json:"cant_allocate_window_frame"`
    Header_no_stream int `json:"header_no_stream"`
    Header_padlen_gt_frame_payload int `json:"header_padlen_gt_frame_payload"`
    Streams_gt_max_concur_streams int `json:"streams_gt_max_concur_streams"`
    Idle_state_unexpected_frame int `json:"idle_state_unexpected_frame"`
    Reserved_local_state_unexpected_frame int `json:"reserved_local_state_unexpected_frame"`
    Reserved_remote_state_unexpected_frame int `json:"reserved_remote_state_unexpected_frame"`
    Half_closed_remote_state_unexpected_fra int `json:"half_closed_remote_state_unexpected_fra"`
    Closed_state_unexpected_frame int `json:"closed_state_unexpected_frame"`
    Zero_window_size_on_stream int `json:"zero_window_size_on_stream"`
    Exceeds_max_window_size_stream int `json:"exceeds_max_window_size_stream"`
    Continuation_before_headers int `json:"continuation_before_headers"`
    Invalid_frame_during_headers int `json:"invalid_frame_during_headers"`
    Headers_after_continuation int `json:"headers_after_continuation"`
    Invalid_push_promise int `json:"invalid_push_promise"`
    Invalid_stream_id int `json:"invalid_stream_id"`
    Headers_interleaved int `json:"headers_interleaved"`
    Trailers_no_end_stream int `json:"trailers_no_end_stream"`
    Invalid_setting_value int `json:"invalid_setting_value"`
    Invalid_window_update int `json:"invalid_window_update"`
    Alloc_fail_total int `json:"alloc_fail_total"`
    Err_rcvd_total int `json:"err_rcvd_total"`
    Err_sent_total int `json:"err_sent_total"`
    Err_sent_proto_err int `json:"err_sent_proto_err"`
    Err_sent_internal_err int `json:"err_sent_internal_err"`
    Err_sent_flow_control int `json:"err_sent_flow_control"`
    Err_sent_setting_timeout int `json:"err_sent_setting_timeout"`
    Err_sent_stream_closed int `json:"err_sent_stream_closed"`
    Err_sent_frame_size_err int `json:"err_sent_frame_size_err"`
    Err_sent_refused_stream int `json:"err_sent_refused_stream"`
    Err_sent_cancel int `json:"err_sent_cancel"`
    Err_sent_compression_err int `json:"err_sent_compression_err"`
    Err_sent_connect_err int `json:"err_sent_connect_err"`
    Err_sent_your_calm int `json:"err_sent_your_calm"`
    Err_sent_inadequate_security int `json:"err_sent_inadequate_security"`
    Err_sent_http11_required int `json:"err_sent_http11_required"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2534 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Protocol_error int `json:"protocol_error"`
    Internal_error int `json:"internal_error"`
    Proxy_alloc_error int `json:"proxy_alloc_error"`
    Split_buff_fail int `json:"split_buff_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Error_max_invalid_stream int `json:"error_max_invalid_stream"`
    Data_no_stream int `json:"data_no_stream"`
    Flow_control_error int `json:"flow_control_error"`
    Settings_timeout int `json:"settings_timeout"`
    Frame_size_error int `json:"frame_size_error"`
    Refused_stream int `json:"refused_stream"`
    Cancel int `json:"cancel"`
    Compression_error int `json:"compression_error"`
    Connect_error int `json:"connect_error"`
    Enhance_your_calm int `json:"enhance_your_calm"`
    Inadequate_security int `json:"inadequate_security"`
    Http_1_1_required int `json:"http_1_1_required"`
    Deflate_alloc_fail int `json:"deflate_alloc_fail"`
    Inflate_alloc_fail int `json:"inflate_alloc_fail"`
    Inflate_header_fail int `json:"inflate_header_fail"`
    Bad_connection_preface int `json:"bad_connection_preface"`
    Cant_allocate_control_frame int `json:"cant_allocate_control_frame"`
    Cant_allocate_settings_frame int `json:"cant_allocate_settings_frame"`
    Bad_frame_type_for_stream_state int `json:"bad_frame_type_for_stream_state"`
    Wrong_stream_state int `json:"wrong_stream_state"`
    Data_queue_alloc_error int `json:"data_queue_alloc_error"`
    Buff_alloc_error int `json:"buff_alloc_error"`
    Cant_allocate_rst_frame int `json:"cant_allocate_rst_frame"`
    Cant_allocate_goaway_frame int `json:"cant_allocate_goaway_frame"`
    Cant_allocate_ping_frame int `json:"cant_allocate_ping_frame"`
    Cant_allocate_stream int `json:"cant_allocate_stream"`
    Cant_allocate_window_frame int `json:"cant_allocate_window_frame"`
    Header_no_stream int `json:"header_no_stream"`
    Header_padlen_gt_frame_payload int `json:"header_padlen_gt_frame_payload"`
    Streams_gt_max_concur_streams int `json:"streams_gt_max_concur_streams"`
    Idle_state_unexpected_frame int `json:"idle_state_unexpected_frame"`
    Reserved_local_state_unexpected_frame int `json:"reserved_local_state_unexpected_frame"`
    Reserved_remote_state_unexpected_frame int `json:"reserved_remote_state_unexpected_frame"`
    Half_closed_remote_state_unexpected_fra int `json:"half_closed_remote_state_unexpected_fra"`
    Closed_state_unexpected_frame int `json:"closed_state_unexpected_frame"`
    Zero_window_size_on_stream int `json:"zero_window_size_on_stream"`
    Exceeds_max_window_size_stream int `json:"exceeds_max_window_size_stream"`
    Continuation_before_headers int `json:"continuation_before_headers"`
    Invalid_frame_during_headers int `json:"invalid_frame_during_headers"`
    Headers_after_continuation int `json:"headers_after_continuation"`
    Invalid_push_promise int `json:"invalid_push_promise"`
    Invalid_stream_id int `json:"invalid_stream_id"`
    Headers_interleaved int `json:"headers_interleaved"`
    Trailers_no_end_stream int `json:"trailers_no_end_stream"`
    Invalid_setting_value int `json:"invalid_setting_value"`
    Invalid_window_update int `json:"invalid_window_update"`
    Alloc_fail_total int `json:"alloc_fail_total"`
    Err_rcvd_total int `json:"err_rcvd_total"`
    Err_sent_total int `json:"err_sent_total"`
    Err_sent_proto_err int `json:"err_sent_proto_err"`
    Err_sent_internal_err int `json:"err_sent_internal_err"`
    Err_sent_flow_control int `json:"err_sent_flow_control"`
    Err_sent_setting_timeout int `json:"err_sent_setting_timeout"`
    Err_sent_stream_closed int `json:"err_sent_stream_closed"`
    Err_sent_frame_size_err int `json:"err_sent_frame_size_err"`
    Err_sent_refused_stream int `json:"err_sent_refused_stream"`
    Err_sent_cancel int `json:"err_sent_cancel"`
    Err_sent_compression_err int `json:"err_sent_compression_err"`
    Err_sent_connect_err int `json:"err_sent_connect_err"`
    Err_sent_your_calm int `json:"err_sent_your_calm"`
    Err_sent_inadequate_security int `json:"err_sent_inadequate_security"`
    Err_sent_http11_required int `json:"err_sent_http11_required"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7session2535 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2536 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2537 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2536 struct {
    Conn_not_exist int `json:"conn_not_exist"`
    Wbuf_cb_failed int `json:"wbuf_cb_failed"`
    Err_event int `json:"err_event"`
    Err_cb_failed int `json:"err_cb_failed"`
    Server_conn_failed int `json:"server_conn_failed"`
    Server_select_fail int `json:"server_select_fail"`
    Data_cb_failed int `json:"data_cb_failed"`
    Hps_fwdreq_fail int `json:"hps_fwdreq_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2537 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Conn_not_exist int `json:"conn_not_exist"`
    Wbuf_cb_failed int `json:"wbuf_cb_failed"`
    Err_event int `json:"err_event"`
    Err_cb_failed int `json:"err_cb_failed"`
    Server_conn_failed int `json:"server_conn_failed"`
    Server_select_fail int `json:"server_select_fail"`
    Data_cb_failed int `json:"data_cb_failed"`
    Hps_fwdreq_fail int `json:"hps_fwdreq_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmpp2538 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2539 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2540 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2539 struct {
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_fail int `json:"select_server_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2540 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_fail int `json:"select_server_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtp2541 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2542 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2543 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2542 struct {
    No_proxy int `json:"no_proxy"`
    Parse_req_fail int `json:"parse_req_fail"`
    Server_select_fail int `json:"server_select_fail"`
    Forward_req_fail int `json:"forward_req_fail"`
    Forward_req_data_fail int `json:"forward_req_data_fail"`
    Snat_fail int `json:"snat_fail"`
    Send_client_service_not_ready int `json:"send_client_service_not_ready"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
    Read_request_line_fail int `json:"read_request_line_fail"`
    Get_all_headers_fail int `json:"get_all_headers_fail"`
    Too_many_headers int `json:"too_many_headers"`
    Line_too_long int `json:"line_too_long"`
    Line_extend_fail int `json:"line_extend_fail"`
    Line_table_extend_fail int `json:"line_table_extend_fail"`
    Parse_request_line_fail int `json:"parse_request_line_fail"`
    Insert_resonse_line_fail int `json:"insert_resonse_line_fail"`
    Remove_resonse_line_fail int `json:"remove_resonse_line_fail"`
    Parse_resonse_line_fail int `json:"parse_resonse_line_fail"`
    Server_starttls_fail int `json:"server_STARTTLS_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2543 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    No_proxy int `json:"no_proxy"`
    Parse_req_fail int `json:"parse_req_fail"`
    Server_select_fail int `json:"server_select_fail"`
    Forward_req_fail int `json:"forward_req_fail"`
    Forward_req_data_fail int `json:"forward_req_data_fail"`
    Snat_fail int `json:"snat_fail"`
    Send_client_service_not_ready int `json:"send_client_service_not_ready"`
    Recv_server_unknow_reply_code int `json:"recv_server_unknow_reply_code"`
    Read_request_line_fail int `json:"read_request_line_fail"`
    Get_all_headers_fail int `json:"get_all_headers_fail"`
    Too_many_headers int `json:"too_many_headers"`
    Line_too_long int `json:"line_too_long"`
    Line_extend_fail int `json:"line_extend_fail"`
    Line_table_extend_fail int `json:"line_table_extend_fail"`
    Parse_request_line_fail int `json:"parse_request_line_fail"`
    Insert_resonse_line_fail int `json:"insert_resonse_line_fail"`
    Remove_resonse_line_fail int `json:"remove_resonse_line_fail"`
    Parse_resonse_line_fail int `json:"parse_resonse_line_fail"`
    Server_starttls_fail int `json:"server_STARTTLS_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqtt2544 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2545 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2546 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2545 struct {
    Parse_connect_fail int `json:"parse_connect_fail"`
    Parse_publish_fail int `json:"parse_publish_fail"`
    Parse_subscribe_fail int `json:"parse_subscribe_fail"`
    Parse_unsubscribe_fail int `json:"parse_unsubscribe_fail"`
    Tuple_not_linked int `json:"tuple_not_linked"`
    Tuple_already_linked int `json:"tuple_already_linked"`
    Conn_null int `json:"conn_null"`
    Client_id_null int `json:"client_id_null"`
    Session_exist int `json:"session_exist"`
    Insertion_failed int `json:"insertion_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2546 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Parse_connect_fail int `json:"parse_connect_fail"`
    Parse_publish_fail int `json:"parse_publish_fail"`
    Parse_subscribe_fail int `json:"parse_subscribe_fail"`
    Parse_unsubscribe_fail int `json:"parse_unsubscribe_fail"`
    Tuple_not_linked int `json:"tuple_not_linked"`
    Tuple_already_linked int `json:"tuple_already_linked"`
    Conn_null int `json:"conn_null"`
    Client_id_null int `json:"client_id_null"`
    Session_exist int `json:"session_exist"`
    Insertion_failed int `json:"insertion_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcap2547 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2548 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2549 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2548 struct {
    App_serv_conn_no_pcb_err int `json:"app_serv_conn_no_pcb_err"`
    App_serv_conn_err int `json:"app_serv_conn_err"`
    Chunk1_hdr_err int `json:"chunk1_hdr_err"`
    Chunk2_hdr_err int `json:"chunk2_hdr_err"`
    Chunk_bad_trail_err int `json:"chunk_bad_trail_err"`
    No_payload_next_buff_err int `json:"no_payload_next_buff_err"`
    No_payload_buff_err int `json:"no_payload_buff_err"`
    Resp_hdr_incomplete_err int `json:"resp_hdr_incomplete_err"`
    Serv_sel_fail_err int `json:"serv_sel_fail_err"`
    Start_icap_conn_fail_err int `json:"start_icap_conn_fail_err"`
    Prep_req_fail_err int `json:"prep_req_fail_err"`
    Icap_ver_err int `json:"icap_ver_err"`
    Icap_line_err int `json:"icap_line_err"`
    Encap_hdr_incomplete_err int `json:"encap_hdr_incomplete_err"`
    No_icap_resp_err int `json:"no_icap_resp_err"`
    Resp_line_read_err int `json:"resp_line_read_err"`
    Resp_line_parse_err int `json:"resp_line_parse_err"`
    Resp_hdr_err int `json:"resp_hdr_err"`
    Req_hdr_incomplete_err int `json:"req_hdr_incomplete_err"`
    No_status_code_err int `json:"no_status_code_err"`
    Http_resp_line_read_err int `json:"http_resp_line_read_err"`
    Http_resp_line_parse_err int `json:"http_resp_line_parse_err"`
    Http_resp_hdr_err int `json:"http_resp_hdr_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2549 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    App_serv_conn_no_pcb_err int `json:"app_serv_conn_no_pcb_err"`
    App_serv_conn_err int `json:"app_serv_conn_err"`
    Chunk1_hdr_err int `json:"chunk1_hdr_err"`
    Chunk2_hdr_err int `json:"chunk2_hdr_err"`
    Chunk_bad_trail_err int `json:"chunk_bad_trail_err"`
    No_payload_next_buff_err int `json:"no_payload_next_buff_err"`
    No_payload_buff_err int `json:"no_payload_buff_err"`
    Resp_hdr_incomplete_err int `json:"resp_hdr_incomplete_err"`
    Serv_sel_fail_err int `json:"serv_sel_fail_err"`
    Start_icap_conn_fail_err int `json:"start_icap_conn_fail_err"`
    Prep_req_fail_err int `json:"prep_req_fail_err"`
    Icap_ver_err int `json:"icap_ver_err"`
    Icap_line_err int `json:"icap_line_err"`
    Encap_hdr_incomplete_err int `json:"encap_hdr_incomplete_err"`
    No_icap_resp_err int `json:"no_icap_resp_err"`
    Resp_line_read_err int `json:"resp_line_read_err"`
    Resp_line_parse_err int `json:"resp_line_parse_err"`
    Resp_hdr_err int `json:"resp_hdr_err"`
    Req_hdr_incomplete_err int `json:"req_hdr_incomplete_err"`
    No_status_code_err int `json:"no_status_code_err"`
    Http_resp_line_read_err int `json:"http_resp_line_read_err"`
    Http_resp_line_parse_err int `json:"http_resp_line_parse_err"`
    Http_resp_hdr_err int `json:"http_resp_hdr_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSip2550 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsInc2551 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsRate2552 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsInc2551 struct {
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSipTriggerStatsRate2552 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompress2553 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2554 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2555 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2554 struct {
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2555 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysql2556 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2557 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2558 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2557 struct {
    Session_err int `json:"session_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2558 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Session_err int `json:"session_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssql2559 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2560 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2561 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2560 struct {
    Session_err int `json:"session_err"`
    Auth_failure int `json:"auth_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2561 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Session_err int `json:"session_err"`
    Auth_failure int `json:"auth_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcip2562 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2563 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2564 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2563 struct {
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Threshold_exceed int `json:"threshold_exceed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2564 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Threshold_exceed int `json:"threshold_exceed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGeneric2565 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2566 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2567 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2566 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Client_fail int `json:"client_fail"`
    Server_fail int `json:"server_fail"`
    Mismatch_fwd_id int `json:"mismatch_fwd_id"`
    Mismatch_rev_id int `json:"mismatch_rev_id"`
    Unkwn_cmd_code int `json:"unkwn_cmd_code"`
    No_session_id int `json:"no_session_id"`
    No_fwd_tuple int `json:"no_fwd_tuple"`
    No_rev_tuple int `json:"no_rev_tuple"`
    Dcmsg_error int `json:"dcmsg_error"`
    Retry_client_request_fail int `json:"retry_client_request_fail"`
    Reply_unknown_session_id int `json:"reply_unknown_session_id"`
    Client_select_fail int `json:"client_select_fail"`
    Invalid_avp int `json:"invalid_avp"`
    Reply_error_info_fail int `json:"reply_error_info_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2567 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Client_fail int `json:"client_fail"`
    Server_fail int `json:"server_fail"`
    Mismatch_fwd_id int `json:"mismatch_fwd_id"`
    Mismatch_rev_id int `json:"mismatch_rev_id"`
    Unkwn_cmd_code int `json:"unkwn_cmd_code"`
    No_session_id int `json:"no_session_id"`
    No_fwd_tuple int `json:"no_fwd_tuple"`
    No_rev_tuple int `json:"no_rev_tuple"`
    Dcmsg_error int `json:"dcmsg_error"`
    Retry_client_request_fail int `json:"retry_client_request_fail"`
    Reply_unknown_session_id int `json:"reply_unknown_session_id"`
    Client_select_fail int `json:"client_select_fail"`
    Invalid_avp int `json:"invalid_avp"`
    Reply_error_info_fail int `json:"reply_error_info_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersist2568 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2569 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2570 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2569 struct {
    Hash_tbl_trylock_fail int `json:"hash_tbl_trylock_fail"`
    Hash_tbl_create_fail int `json:"hash_tbl_create_fail"`
    Hash_tbl_rst_updown int `json:"hash_tbl_rst_updown"`
    Hash_tbl_rst_adddel int `json:"hash_tbl_rst_adddel"`
    Url_hash_fail int `json:"url_hash_fail"`
    Header_hash_fail int `json:"header_hash_fail"`
    Src_ip_fail int `json:"src_ip_fail"`
    Src_ip_new_sess_cache_fail int `json:"src_ip_new_sess_cache_fail"`
    Src_ip_new_sess_sel_fail int `json:"src_ip_new_sess_sel_fail"`
    Src_ip_hash_fail int `json:"src_ip_hash_fail"`
    Dst_ip_fail int `json:"dst_ip_fail"`
    Dst_ip_new_sess_cache_fail int `json:"dst_ip_new_sess_cache_fail"`
    Dst_ip_new_sess_sel_fail int `json:"dst_ip_new_sess_sel_fail"`
    Dst_ip_hash_fail int `json:"dst_ip_hash_fail"`
    Cssl_sid_not_found int `json:"cssl_sid_not_found"`
    Cssl_sid_not_match int `json:"cssl_sid_not_match"`
    Sssl_sid_not_found int `json:"sssl_sid_not_found"`
    Sssl_sid_not_match int `json:"sssl_sid_not_match"`
    Ssl_sid_persist_fail int `json:"ssl_sid_persist_fail"`
    Ssl_sid_session_fail int `json:"ssl_sid_session_fail"`
    Cookie_persist_fail int `json:"cookie_persist_fail"`
    Cookie_not_found int `json:"cookie_not_found"`
    Cookie_invalid int `json:"cookie_invalid"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2570 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Hash_tbl_trylock_fail int `json:"hash_tbl_trylock_fail"`
    Hash_tbl_create_fail int `json:"hash_tbl_create_fail"`
    Hash_tbl_rst_updown int `json:"hash_tbl_rst_updown"`
    Hash_tbl_rst_adddel int `json:"hash_tbl_rst_adddel"`
    Url_hash_fail int `json:"url_hash_fail"`
    Header_hash_fail int `json:"header_hash_fail"`
    Src_ip_fail int `json:"src_ip_fail"`
    Src_ip_new_sess_cache_fail int `json:"src_ip_new_sess_cache_fail"`
    Src_ip_new_sess_sel_fail int `json:"src_ip_new_sess_sel_fail"`
    Src_ip_hash_fail int `json:"src_ip_hash_fail"`
    Dst_ip_fail int `json:"dst_ip_fail"`
    Dst_ip_new_sess_cache_fail int `json:"dst_ip_new_sess_cache_fail"`
    Dst_ip_new_sess_sel_fail int `json:"dst_ip_new_sess_sel_fail"`
    Dst_ip_hash_fail int `json:"dst_ip_hash_fail"`
    Cssl_sid_not_found int `json:"cssl_sid_not_found"`
    Cssl_sid_not_match int `json:"cssl_sid_not_match"`
    Sssl_sid_not_found int `json:"sssl_sid_not_found"`
    Sssl_sid_not_match int `json:"sssl_sid_not_match"`
    Ssl_sid_persist_fail int `json:"ssl_sid_persist_fail"`
    Ssl_sid_session_fail int `json:"ssl_sid_session_fail"`
    Cookie_persist_fail int `json:"cookie_persist_fail"`
    Cookie_not_found int `json:"cookie_not_found"`
    Cookie_invalid int `json:"cookie_invalid"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxy2571 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2572 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2573 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2572 struct {
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Snat_fail int `json:"snat_fail"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2573 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Snat_fail int `json:"snat_fail"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL42574 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2575 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2576 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsInc2575 struct {
    Syncookiessentfailed int `json:"syncookiessentfailed"`
    Svrselfail int `json:"svrselfail"`
    Snat_fail int `json:"snat_fail"`
    Snat_no_fwd_route int `json:"snat_no_fwd_route"`
    Snat_no_rev_route int `json:"snat_no_rev_route"`
    Snat_icmp_error_process int `json:"snat_icmp_error_process"`
    Snat_icmp_no_match int `json:"snat_icmp_no_match"`
    Smart_nat_id_mismatch int `json:"smart_nat_id_mismatch"`
    Syncookiescheckfailed int `json:"syncookiescheckfailed"`
    Connlimit_drop int `json:"connlimit_drop"`
    Conn_rate_limit_drop int `json:"conn_rate_limit_drop"`
    Conn_rate_limit_reset int `json:"conn_rate_limit_reset"`
    Dns_policy_drop int `json:"dns_policy_drop"`
    No_resourse_drop int `json:"no_resourse_drop"`
    Bw_rate_limit_exceed int `json:"bw_rate_limit_exceed"`
    L4_cps_exceed int `json:"l4_cps_exceed"`
    Nat_cps_exceed int `json:"nat_cps_exceed"`
    L7_cps_exceed int `json:"l7_cps_exceed"`
    Ssl_cps_exceed int `json:"ssl_cps_exceed"`
    Ssl_tpt_exceed int `json:"ssl_tpt_exceed"`
    Concurrent_conn_exceed int `json:"concurrent_conn_exceed"`
    Svr_syn_handshake_fail int `json:"svr_syn_handshake_fail"`
    Synattack int `json:"synattack"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL4TriggerStatsRate2576 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Syncookiessentfailed int `json:"syncookiessentfailed"`
    Svrselfail int `json:"svrselfail"`
    Snat_fail int `json:"snat_fail"`
    Snat_no_fwd_route int `json:"snat_no_fwd_route"`
    Snat_no_rev_route int `json:"snat_no_rev_route"`
    Snat_icmp_error_process int `json:"snat_icmp_error_process"`
    Snat_icmp_no_match int `json:"snat_icmp_no_match"`
    Smart_nat_id_mismatch int `json:"smart_nat_id_mismatch"`
    Syncookiescheckfailed int `json:"syncookiescheckfailed"`
    Connlimit_drop int `json:"connlimit_drop"`
    Conn_rate_limit_drop int `json:"conn_rate_limit_drop"`
    Conn_rate_limit_reset int `json:"conn_rate_limit_reset"`
    Dns_policy_drop int `json:"dns_policy_drop"`
    No_resourse_drop int `json:"no_resourse_drop"`
    Bw_rate_limit_exceed int `json:"bw_rate_limit_exceed"`
    L4_cps_exceed int `json:"l4_cps_exceed"`
    Nat_cps_exceed int `json:"nat_cps_exceed"`
    L7_cps_exceed int `json:"l7_cps_exceed"`
    Ssl_cps_exceed int `json:"ssl_cps_exceed"`
    Ssl_tpt_exceed int `json:"ssl_tpt_exceed"`
    Concurrent_conn_exceed int `json:"concurrent_conn_exceed"`
    Svr_syn_handshake_fail int `json:"svr_syn_handshake_fail"`
    Synattack int `json:"synattack"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttp2577 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2578 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2579 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2578 struct {
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Snat_fail int `json:"snat_fail"`
    Full_proxy_fpga_err int `json:"full_proxy_fpga_err"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2579 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Snat_fail int `json:"snat_fail"`
    Full_proxy_fpga_err int `json:"full_proxy_fpga_err"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxy2580 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2581 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2582 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2581 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Data_conn_start_err int `json:"data_conn_start_err"`
    Data_serv_connecting_err int `json:"data_serv_connecting_err"`
    Data_serv_connected_err int `json:"data_serv_connected_err"`
    Auth_fail int `json:"auth_fail"`
    Ds_fail int `json:"ds_fail"`
    Cant_find_port int `json:"cant_find_port"`
    Cant_find_eprt int `json:"cant_find_eprt"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2582 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Data_conn_start_err int `json:"data_conn_start_err"`
    Data_serv_connecting_err int `json:"data_serv_connecting_err"`
    Data_serv_connected_err int `json:"data_serv_connected_err"`
    Auth_fail int `json:"auth_fail"`
    Ds_fail int `json:"ds_fail"`
    Cant_find_port int `json:"cant_find_port"`
    Cant_find_eprt int `json:"cant_find_eprt"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxy2583 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2584 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2585 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2584 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Auth_unsupported int `json:"auth_unsupported"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2585 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Cant_find_pasv int `json:"cant_find_pasv"`
    Smp_create_fail int `json:"smp_create_fail"`
    Data_server_conn_fail int `json:"data_server_conn_fail"`
    Data_send_fail int `json:"data_send_fail"`
    Cant_find_epsv int `json:"cant_find_epsv"`
    Auth_unsupported int `json:"auth_unsupported"`
    Unsupported_pbsz_value int `json:"unsupported_pbsz_value"`
    Unsupported_prot_value int `json:"unsupported_prot_value"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Cl_request_err int `json:"cl_request_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3Proxy2586 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2587 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2588 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2587 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2588 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Line_too_long int `json:"line_too_long"`
    Invalid_start_line int `json:"invalid_start_line"`
    Unsupported_command int `json:"unsupported_command"`
    Bad_sequence int `json:"bad_sequence"`
    Rsv_persist_conn_fail int `json:"rsv_persist_conn_fail"`
    Smp_v6_fail int `json:"smp_v6_fail"`
    Smp_v4_fail int `json:"smp_v4_fail"`
    Insert_tuple_fail int `json:"insert_tuple_fail"`
    Cl_est_err int `json:"cl_est_err"`
    Ser_connecting_err int `json:"ser_connecting_err"`
    Server_response_err int `json:"server_response_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitch2589 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2590 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2591 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2590 struct {
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2591 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCache2592 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2593 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2594 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2593 struct {
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2594 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCounters2595 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2596 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2597 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsInc2596 struct {
    So_pkts_slb_nat_reserve_fail int `json:"so_pkts_slb_nat_reserve_fail"`
    So_pkts_slb_nat_release_fail int `json:"so_pkts_slb_nat_release_fail"`
    So_pkts_l2redirect_dest_mac_zero_drop int `json:"so_pkts_l2redirect_dest_mac_zero_drop"`
    So_pkts_l2redirect_interface_not_up int `json:"so_pkts_l2redirect_interface_not_up"`
    So_pkts_l2redirect_invalid_redirect_inf int `json:"so_pkts_l2redirect_invalid_redirect_inf"`
    So_pkts_l3_redirect_encap_error_drop int `json:"so_pkts_l3_redirect_encap_error_drop"`
    So_pkts_l3_redirect_inner_mac_zero_drop int `json:"so_pkts_l3_redirect_inner_mac_zero_drop"`
    So_pkts_l3_redirect_table_error int `json:"so_pkts_l3_redirect_table_error"`
    So_pkts_l3_redirect_fragmentation_error int `json:"so_pkts_l3_redirect_fragmentation_error"`
    So_pkts_l3_redirect_table_no_entry_foun int `json:"so_pkts_l3_redirect_table_no_entry_foun"`
    So_pkts_l3_redirect_invalid_dev_dir int `json:"so_pkts_l3_redirect_invalid_dev_dir"`
    So_pkts_l3_redirect_chassis_dest_mac_er int `json:"so_pkts_l3_redirect_chassis_dest_mac_er"`
    So_pkts_l2redirect_vlan_retrieval_error int `json:"so_pkts_l2redirect_vlan_retrieval_error"`
    So_pkts_l2redirect_port_retrieval_error int `json:"so_pkts_l2redirect_port_retrieval_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSoCountersTriggerStatsRate2597 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    So_pkts_slb_nat_reserve_fail int `json:"so_pkts_slb_nat_reserve_fail"`
    So_pkts_slb_nat_release_fail int `json:"so_pkts_slb_nat_release_fail"`
    So_pkts_l2redirect_dest_mac_zero_drop int `json:"so_pkts_l2redirect_dest_mac_zero_drop"`
    So_pkts_l2redirect_interface_not_up int `json:"so_pkts_l2redirect_interface_not_up"`
    So_pkts_l2redirect_invalid_redirect_inf int `json:"so_pkts_l2redirect_invalid_redirect_inf"`
    So_pkts_l3_redirect_encap_error_drop int `json:"so_pkts_l3_redirect_encap_error_drop"`
    So_pkts_l3_redirect_inner_mac_zero_drop int `json:"so_pkts_l3_redirect_inner_mac_zero_drop"`
    So_pkts_l3_redirect_table_error int `json:"so_pkts_l3_redirect_table_error"`
    So_pkts_l3_redirect_fragmentation_error int `json:"so_pkts_l3_redirect_fragmentation_error"`
    So_pkts_l3_redirect_table_no_entry_foun int `json:"so_pkts_l3_redirect_table_no_entry_foun"`
    So_pkts_l3_redirect_invalid_dev_dir int `json:"so_pkts_l3_redirect_invalid_dev_dir"`
    So_pkts_l3_redirect_chassis_dest_mac_er int `json:"so_pkts_l3_redirect_chassis_dest_mac_er"`
    So_pkts_l2redirect_vlan_retrieval_error int `json:"so_pkts_l2redirect_vlan_retrieval_error"`
    So_pkts_l2redirect_port_retrieval_error int `json:"so_pkts_l2redirect_port_retrieval_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGbl2598 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2599 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2600 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2599 struct {
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2600 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRate2601 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2602 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2603 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2602 struct {
    Total_reset int `json:"total_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2603 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_reset int `json:"total_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobal2604 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2605 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2606 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2605 struct {
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2606 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb2607 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2608 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2609 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2608 struct {
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2609 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbe2610 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2611 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2612 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2611 struct {
    Err_entry_create_failed int `json:"err_entry_create_failed"`
    Err_entry_create_oom int `json:"err_entry_create_oom"`
    Err_entry_insert_failed int `json:"err_entry_insert_failed"`
    Err_tmpl_probe_create_failed int `json:"err_tmpl_probe_create_failed"`
    Err_tmpl_probe_create_oom int `json:"err_tmpl_probe_create_oom"`
    Total_http_response_bad int `json:"total_http_response_bad"`
    Total_tcp_err int `json:"total_tcp_err"`
    Err_smart_nat_alloc int `json:"err_smart_nat_alloc"`
    Err_smart_nat_port_alloc int `json:"err_smart_nat_port_alloc"`
    Err_l4_sess_alloc int `json:"err_l4_sess_alloc"`
    Err_probe_tcp_conn_send int `json:"err_probe_tcp_conn_send"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2612 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Err_entry_create_failed int `json:"err_entry_create_failed"`
    Err_entry_create_oom int `json:"err_entry_create_oom"`
    Err_entry_insert_failed int `json:"err_entry_insert_failed"`
    Err_tmpl_probe_create_failed int `json:"err_tmpl_probe_create_failed"`
    Err_tmpl_probe_create_oom int `json:"err_tmpl_probe_create_oom"`
    Total_http_response_bad int `json:"total_http_response_bad"`
    Total_tcp_err int `json:"total_tcp_err"`
    Err_smart_nat_alloc int `json:"err_smart_nat_alloc"`
    Err_smart_nat_port_alloc int `json:"err_smart_nat_port_alloc"`
    Err_l4_sess_alloc int `json:"err_l4_sess_alloc"`
    Err_probe_tcp_conn_send int `json:"err_probe_tcp_conn_send"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpz2613 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2614 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2615 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2614 struct {
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2615 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsSeverity2616 struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplate) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplate) getPath() string{
    return "visibility/packet-capture/global-templates/template"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplate::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplate::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplate::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
