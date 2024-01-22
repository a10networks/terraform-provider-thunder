

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCapture struct {
	Inst struct {

    AutomatedCaptures VisibilityPacketCaptureAutomatedCaptures2723 `json:"automated-captures"`
    CaptureConfigList []VisibilityPacketCaptureCaptureConfigList `json:"capture-config-list"`
    DeletePacketCaptureFile VisibilityPacketCaptureDeletePacketCaptureFile2724 `json:"delete-packet-capture-file"`
    GlobalTemplates VisibilityPacketCaptureGlobalTemplates2725 `json:"global-templates"`
    ObjectTemplates VisibilityPacketCaptureObjectTemplates2985 `json:"object-templates"`
    Uuid string `json:"uuid"`

	} `json:"packet-capture"`
}


type VisibilityPacketCaptureAutomatedCaptures2723 struct {
    Slb_port_tmpl_error_code_return_inc int `json:"slb_port_tmpl_error_code_return_inc"`
    Slb_port_tmpl_high_error_code_return int `json:"slb_port_tmpl_high_error_code_return"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureCaptureConfigList struct {
    Name string `json:"name"`
    Disable int `json:"disable"`
    ConcurrentCaptures int `json:"concurrent-captures"`
    ConcurrentConnPerCapture int `json:"concurrent-conn-per-capture" dval:"100"`
    ConcurrentCapturesAge int `json:"concurrent-captures-age" dval:"1"`
    ConcurrentConnTag int `json:"concurrent-conn-tag"`
    NumberOfPacketsPerConn int `json:"number-of-packets-per-conn"`
    PacketLength int `json:"packet-length" dval:"128"`
    FileSize int `json:"file-size" dval:"1"`
    FileCount int `json:"file-count" dval:"10"`
    NumberOfPacketsPerCapture int `json:"number-of-packets-per-capture"`
    NumberOfPacketsTotal int `json:"number-of-packets-total"`
    EnableContinuousGlobalCapture int `json:"enable-continuous-global-capture"`
    CreatePcapFilesNow int `json:"create-pcap-files-now"`
    DisableAutoMerge int `json:"disable-auto-merge"`
    KeepPcapFilesAfterMerge int `json:"keep-pcap-files-after-merge"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type VisibilityPacketCaptureDeletePacketCaptureFile2724 struct {
    FileName string `json:"file-name"`
    All int `json:"all"`
}


type VisibilityPacketCaptureGlobalTemplates2725 struct {
    Uuid string `json:"uuid"`
    TemplateList []VisibilityPacketCaptureGlobalTemplatesTemplateList2726 `json:"template-list"`
    Activate VisibilityPacketCaptureGlobalTemplatesActivate2984 `json:"activate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateList2726 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerSysObjStatsSeverity VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsSeverity2727 `json:"trigger-sys-obj-stats-severity"`
    TriggerSysObjStatsChange VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChange2728 `json:"trigger-sys-obj-stats-change"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsSeverity2727 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChange2728 struct {
    Uuid string `json:"uuid"`
    SystemCtrLibAcct VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcct2729 `json:"system-ctr-lib-acct"`
    SystemHardwareAccelerate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerate2732 `json:"system-hardware-accelerate"`
    SystemRadiusServer VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServer2735 `json:"system-radius-server"`
    SystemIpThreatList VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatList2738 `json:"system-ip-threat-list"`
    SystemFpgaDrop VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDrop2741 `json:"system-fpga-drop"`
    SystemDpdkStats VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStats2744 `json:"system-dpdk-stats"`
    IpAnomalyDrop VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDrop2747 `json:"ip-anomaly-drop"`
    AamAuthenticationGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobal2750 `json:"aam-authentication-global"`
    AamRdns VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdns2753 `json:"aam-rdns"`
    AamAuthServerLdap VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdap2756 `json:"aam-auth-server-ldap"`
    AamAuthServerOcsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcsp2759 `json:"aam-auth-server-ocsp"`
    AamAuthServerRadius VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadius2762 `json:"aam-auth-server-radius"`
    AamAuthServerWin VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWin2765 `json:"aam-auth-server-win"`
    AamAuthAccount VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccount2768 `json:"aam-auth-account"`
    AamAuthSamlGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobal2771 `json:"aam-auth-saml-global"`
    AamAuthRelayKerberos VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberos2774 `json:"aam-auth-relay-kerberos"`
    AamAuthCaptcha VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptcha2777 `json:"aam-auth-captcha"`
    SlbSslError VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslError2780 `json:"slb-ssl-error"`
    SlbSslCertRevoke VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevoke2783 `json:"slb-ssl-cert-revoke"`
    SlbSslForwardProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxy2786 `json:"slb-ssl-forward-proxy"`
    VpnError VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnError2789 `json:"vpn-error"`
    Cgnv6Global VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Global2792 `json:"cgnv6-global"`
    Cgnv6DdosProc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProc2795 `json:"cgnv6-ddos-proc"`
    Cgnv6Lsn VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Lsn2798 `json:"cgnv6-lsn"`
    Cgnv6LsnAlgEsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEsp2801 `json:"cgnv6-lsn-alg-esp"`
    Cgnv6LsnAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptp2804 `json:"cgnv6-lsn-alg-pptp"`
    Cgnv6LsnAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2807 `json:"cgnv6-lsn-alg-rtsp"`
    Cgnv6LsnAlgSip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSip2810 `json:"cgnv6-lsn-alg-sip"`
    Cgnv6LsnAlgMgcp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2813 `json:"cgnv6-lsn-alg-mgcp"`
    Cgnv6LsnAlgH323 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH3232816 `json:"cgnv6-lsn-alg-h323"`
    Cgnv6LsnRadius VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadius2819 `json:"cgnv6-lsn-radius"`
    Cgnv6Nat64Global VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64Global2822 `json:"cgnv6-nat64-global"`
    Cgnv6DsLiteGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobal2825 `json:"cgnv6-ds-lite-global"`
    Cgnv6FixedNatGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobal2828 `json:"cgnv6-fixed-nat-global"`
    Cgnv6FixedNatAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2831 `json:"cgnv6-fixed-nat-alg-pptp"`
    Cgnv6FixedNatAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2834 `json:"cgnv6-fixed-nat-alg-rtsp"`
    Cgnv6FixedNatAlgSip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2837 `json:"cgnv6-fixed-nat-alg-sip"`
    Cgnv6Pcp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Pcp2840 `json:"cgnv6-pcp"`
    Cgnv6Logging VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Logging2843 `json:"cgnv6-logging"`
    Cgnv6L4 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L42846 `json:"cgnv6-l4"`
    Cgnv6Icmp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Icmp2849 `json:"cgnv6-icmp"`
    Cgnv6HttpAlg VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlg2852 `json:"cgnv6-http-alg"`
    Cgnv6Dns64 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns642855 `json:"cgnv6-dns64"`
    Cgnv6Dhcpv6 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv62858 `json:"cgnv6-dhcpv6"`
    FwLogging VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLogging2861 `json:"fw-logging"`
    FwGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobal2864 `json:"fw-global"`
    FwAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtsp2867 `json:"fw-alg-rtsp"`
    FwAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptp2870 `json:"fw-alg-pptp"`
    FwRadServer VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServer2873 `json:"fw-rad-server"`
    FwTcpSynCookie VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookie2876 `json:"fw-tcp-syn-cookie"`
    FwDdosProtection VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtection2879 `json:"fw-ddos-protection"`
    FwGtp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtp2882 `json:"fw-gtp"`
    SystemTcp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcp2885 `json:"system-tcp"`
    SlbConnReuse VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuse2888 `json:"slb-conn-reuse"`
    SlbAflow VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflow2891 `json:"slb-aflow"`
    SlbFix VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFix2894 `json:"slb-fix"`
    SlbSpdyProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxy2897 `json:"slb-spdy-proxy"`
    SlbHttp2 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp22900 `json:"slb-http2"`
    SlbL7session VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7session2903 `json:"slb-l7session"`
    SlbSmpp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmpp2906 `json:"slb-smpp"`
    SlbSmtp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtp2909 `json:"slb-smtp"`
    SlbMqtt VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqtt2912 `json:"slb-mqtt"`
    SlbIcap VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcap2915 `json:"slb-icap"`
    SlbSip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSip2918 `json:"slb-sip"`
    SlbHwCompress VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompress2921 `json:"slb-hw-compress"`
    SlbMysql VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysql2924 `json:"slb-mysql"`
    SlbMssql VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssql2927 `json:"slb-mssql"`
    SlbCrlSrcip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcip2930 `json:"slb-crl-srcip"`
    SlbGeneric VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGeneric2933 `json:"slb-generic"`
    SlbPersist VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersist2936 `json:"slb-persist"`
    SlbHttpProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxy2939 `json:"slb-http-proxy"`
    SlbL4 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL42942 `json:"slb-l4"`
    SlbFastHttp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttp2945 `json:"slb-fast-http"`
    SlbFtpProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxy2948 `json:"slb-ftp-proxy"`
    SlbImapProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxy2951 `json:"slb-imap-proxy"`
    SlbPop3Proxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3Proxy2954 `json:"slb-pop3-proxy"`
    SlbSwitch VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitch2957 `json:"slb-switch"`
    SlbRcCache VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCache2960 `json:"slb-rc-cache"`
    SoCounters VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCounters2963 `json:"so-counters"`
    SlbPlyrIdGbl VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGbl2966 `json:"slb-plyr-id-gbl"`
    SlbSportRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRate2969 `json:"slb-sport-rate"`
    LoggingLocalLogGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobal2972 `json:"logging-local-log-global"`
    SlbMlb VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlb2975 `json:"slb-mlb"`
    SlbLinkProbe VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbe2978 `json:"slb-link-probe"`
    SlbRpz VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpz2981 `json:"slb-rpz"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcct2729 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2730 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2731 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc2730 struct {
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate2731 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerate2732 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2733 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2734 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc2733 struct {
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate2734 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServer2735 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2736 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2737 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc2736 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate2737 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatList2738 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2739 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2740 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc2739 struct {
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate2740 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDrop2741 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2742 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2743 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc2742 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate2743 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStats2744 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2745 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2746 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc2745 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate2746 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDrop2747 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2748 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2749 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc2748 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate2749 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobal2750 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2751 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2752 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc2751 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate2752 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdns2753 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2754 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2755 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsInc2754 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsRate2755 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdap2756 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2757 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2758 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc2757 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate2758 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcsp2759 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2760 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2761 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc2760 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate2761 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadius2762 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2763 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2764 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc2763 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate2764 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWin2765 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2766 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2767 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc2766 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate2767 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccount2768 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2769 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2770 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc2769 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate2770 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobal2771 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2772 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2773 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc2772 struct {
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate2773 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberos2774 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2775 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2776 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc2775 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate2776 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptcha2777 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2778 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2779 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc2778 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate2779 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslError2780 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2781 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2782 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc2781 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate2782 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevoke2783 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2784 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2785 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc2784 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate2785 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxy2786 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2787 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2788 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc2787 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate2788 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnError2789 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2790 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2791 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2790 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2791 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Global2792 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsInc2793 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsRate2794 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsInc2793 struct {
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsRate2794 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProc2795 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2796 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2797 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc2796 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate2797 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Lsn2798 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2799 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2800 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc2799 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate2800 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEsp2801 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2802 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2803 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc2802 struct {
    NatIpConflict int `json:"nat-ip-conflict"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate2803 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NatIpConflict int `json:"nat-ip-conflict"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptp2804 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2805 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2806 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc2805 struct {
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate2806 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtsp2807 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2808 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2809 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc2808 struct {
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate2809 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSip2810 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2811 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2812 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc2811 struct {
    MethodUnknown int `json:"method-unknown"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate2812 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MethodUnknown int `json:"method-unknown"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcp2813 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2814 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2815 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc2814 struct {
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate2815 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH3232816 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2817 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2818 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc2817 struct {
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate2818 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadius2819 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2820 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2821 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc2820 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate2821 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64Global2822 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2823 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2824 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc2823 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate2824 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobal2825 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2826 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2827 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc2826 struct {
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    Fullcone_failure int `json:"fullcone_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate2827 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    Fullcone_failure int `json:"fullcone_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobal2828 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2829 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2830 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc2829 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate2830 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp2831 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2832 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2833 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc2832 struct {
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate2833 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp2834 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2835 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2836 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc2835 struct {
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate2836 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSip2837 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2838 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2839 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc2838 struct {
    MethodUnknown int `json:"method-unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate2839 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MethodUnknown int `json:"method-unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Pcp2840 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2841 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2842 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc2841 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate2842 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Logging2843 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2844 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2845 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc2844 struct {
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate2845 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L42846 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2847 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2848 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc2847 struct {
    OutOfSessionMemory int `json:"out-of-session-memory"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate2848 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    OutOfSessionMemory int `json:"out-of-session-memory"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Icmp2849 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2850 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2851 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc2850 struct {
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate2851 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlg2852 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2853 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2854 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc2853 struct {
    RadiusRequstDropped int `json:"radius-requst-dropped"`
    RadiusResponseDropped int `json:"radius-response-dropped"`
    OutOfMemoryDropped int `json:"out-of-memory-dropped"`
    QueueLenExceedDropped int `json:"queue-len-exceed-dropped"`
    OutOfOrderDropped int `json:"out-of-order-dropped"`
    HeaderInsertionFailed int `json:"header-insertion-failed"`
    HeaderRemovalFailed int `json:"header-removal-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate2854 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns642855 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2856 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2857 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc2856 struct {
    QueryBadPkt int `json:"query-bad-pkt"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    Drop int `json:"drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate2857 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    QueryBadPkt int `json:"query-bad-pkt"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    Drop int `json:"drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv62858 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2859 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2860 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc2859 struct {
    PacketsDropped int `json:"packets-dropped"`
    PktsDroppedDuringClear int `json:"pkts-dropped-during-clear"`
    RcvNotSupportedMsg int `json:"rcv-not-supported-msg"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate2860 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    PacketsDropped int `json:"packets-dropped"`
    PktsDroppedDuringClear int `json:"pkts-dropped-during-clear"`
    RcvNotSupportedMsg int `json:"rcv-not-supported-msg"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLogging2861 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2862 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2863 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsInc2862 struct {
    LogDropped int `json:"log-dropped"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsRate2863 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobal2864 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2865 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2866 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsInc2865 struct {
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsRate2866 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtsp2867 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2868 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2869 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc2868 struct {
    TransportAllocFailure int `json:"transport-alloc-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate2869 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TransportAllocFailure int `json:"transport-alloc-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptp2870 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2871 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2872 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc2871 struct {
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate2872 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServer2873 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2874 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2875 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsInc2874 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsRate2875 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookie2876 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2877 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2878 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc2877 struct {
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate2878 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtection2879 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2880 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2881 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc2880 struct {
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate2881 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtp2882 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsInc2883 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsRate2884 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsInc2883 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsRate2884 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcp2885 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2886 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2887 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsInc2886 struct {
    Attemptfails int `json:"attemptfails"`
    Noroute int `json:"noroute"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsRate2887 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Attemptfails int `json:"attemptfails"`
    Noroute int `json:"noroute"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuse2888 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2889 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2890 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc2889 struct {
    Ntermi_err int `json:"ntermi_err"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate2890 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ntermi_err int `json:"ntermi_err"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflow2891 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2892 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2893 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsInc2892 struct {
    Pause_conn_fail int `json:"pause_conn_fail"`
    Error_resume_conn int `json:"error_resume_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsRate2893 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Error_resume_conn int `json:"error_resume_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFix2894 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsInc2895 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsRate2896 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsInc2895 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsRate2896 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxy2897 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2898 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2899 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc2898 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate2899 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp22900 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2901 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2902 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2901 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2902 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7session2903 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2904 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2905 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc2904 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate2905 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmpp2906 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2907 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2908 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsInc2907 struct {
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_fail int `json:"select_server_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsRate2908 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_fail int `json:"select_server_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtp2909 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2910 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2911 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc2910 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate2911 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqtt2912 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2913 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2914 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsInc2913 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsRate2914 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcap2915 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2916 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2917 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsInc2916 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsRate2917 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSip2918 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsInc2919 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsRate2920 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsInc2919 struct {
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsRate2920 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompress2921 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2922 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2923 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc2922 struct {
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate2923 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysql2924 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2925 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2926 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc2925 struct {
    Session_err int `json:"session_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate2926 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Session_err int `json:"session_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssql2927 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2928 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2929 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc2928 struct {
    Session_err int `json:"session_err"`
    Auth_failure int `json:"auth_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate2929 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Session_err int `json:"session_err"`
    Auth_failure int `json:"auth_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcip2930 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2931 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2932 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc2931 struct {
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Threshold_exceed int `json:"threshold_exceed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate2932 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Threshold_exceed int `json:"threshold_exceed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGeneric2933 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2934 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2935 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsInc2934 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsRate2935 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersist2936 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2937 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2938 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsInc2937 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsRate2938 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxy2939 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2940 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2941 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc2940 struct {
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Snat_fail int `json:"snat_fail"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate2941 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL42942 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsInc2943 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsRate2944 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsInc2943 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsRate2944 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttp2945 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2946 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2947 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc2946 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate2947 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxy2948 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2949 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2950 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc2949 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate2950 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxy2951 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2952 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2953 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc2952 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate2953 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3Proxy2954 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2955 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2956 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc2955 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate2956 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitch2957 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2958 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2959 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc2958 struct {
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate2959 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCache2960 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2961 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2962 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc2961 struct {
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate2962 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCounters2963 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsInc2964 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsRate2965 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsInc2964 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsRate2965 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGbl2966 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2967 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2968 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc2967 struct {
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate2968 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRate2969 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2970 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2971 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc2970 struct {
    Total_reset int `json:"total_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate2971 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_reset int `json:"total_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobal2972 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2973 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2974 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc2973 struct {
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate2974 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlb2975 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2976 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2977 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2976 struct {
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2977 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbe2978 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2979 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2980 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc2979 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate2980 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpz2981 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2982 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2983 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsInc2982 struct {
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsRate2983 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesActivate2984 struct {
    Template string `json:"template"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplates2985 struct {
    Uuid string `json:"uuid"`
    TemplGtpPlcyTmplList []VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList2986 `json:"templ-gtp-plcy-tmpl-list"`
    InterfaceEthernetTmplList []VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList2990 `json:"interface-ethernet-tmpl-list"`
    InterfaceTunnelTmplList []VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList2994 `json:"interface-tunnel-tmpl-list"`
    AamJwtAuthorizationTmplList []VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList2998 `json:"aam-jwt-authorization-tmpl-list"`
    AamAaaPolicyTmplList []VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList3002 `json:"aam-aaa-policy-tmpl-list"`
    AamAuthLogonHttpInsTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList3006 `json:"aam-auth-logon-http-ins-tmpl-list"`
    AamAuthServerLdapInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList3010 `json:"aam-auth-server-ldap-inst-tmpl-list"`
    AamAuthServerOcspInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList3014 `json:"aam-auth-server-ocsp-inst-tmpl-list"`
    AamAuthServerRadInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList3018 `json:"aam-auth-server-rad-inst-tmpl-list"`
    AamAuthServerWinInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList3022 `json:"aam-auth-server-win-inst-tmpl-list"`
    AamAuthSamlServiceProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList3026 `json:"aam-auth-saml-service-prov-tmpl-list"`
    AamAuthSamlIdProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList3030 `json:"aam-auth-saml-id-prov-tmpl-list"`
    AamAuthServiceGroupTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList3034 `json:"aam-auth-service-group-tmpl-list"`
    AamAuthServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList3038 `json:"aam-auth-service-group-mem-tmpl-list"`
    AamAuthRelayHbaseInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList3042 `json:"aam-auth-relay-hbase-inst-tmpl-list"`
    AamAuthRelayFormInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList3046 `json:"aam-auth-relay-form-inst-tmpl-list"`
    AamAuthRelayNtlmTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList3050 `json:"aam-auth-relay-ntlm-tmpl-list"`
    AamAuthRelayWsFedTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList3054 `json:"aam-auth-relay-ws-fed-tmpl-list"`
    AamAuthCaptchaInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList3058 `json:"aam-auth-captcha-inst-tmpl-list"`
    SlbTemplCacheTmplList []VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList3062 `json:"slb-templ-cache-tmpl-list"`
    SlbPortTmplList []VisibilityPacketCaptureObjectTemplatesSlbPortTmplList3066 `json:"slb-port-tmpl-list"`
    SlbVportTmplList []VisibilityPacketCaptureObjectTemplatesSlbVportTmplList3070 `json:"slb-vport-tmpl-list"`
    Cgnv6ServGroupTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList3074 `json:"cgnv6-serv-group-tmpl-list"`
    Cgnv6Dns64VsPortTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList3078 `json:"cgnv6-dns64-vs-port-tmpl-list"`
    Cgnv6MapTransDomainTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList3082 `json:"cgnv6-map-trans-domain-tmpl-list"`
    Cgnv6EncapDomainTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList3086 `json:"cgnv6-encap-domain-tmpl-list"`
    NetflowMonitorTmplList []VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList3090 `json:"netflow-monitor-tmpl-list"`
    RuleSetTmplList []VisibilityPacketCaptureObjectTemplatesRuleSetTmplList3094 `json:"rule-set-tmpl-list"`
    FwServerPortTmplList []VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList3098 `json:"fw-server-port-tmpl-list"`
    FwServiceGroupTmplList []VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList3102 `json:"fw-service-group-tmpl-list"`
    FwServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList3106 `json:"fw-service-group-mem-tmpl-list"`
    Dns_vportTmplList []VisibilityPacketCaptureObjectTemplatesDns_vportTmplList3110 `json:"dns_vport-tmpl-list"`
    SmtpVportTmplList []VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList3114 `json:"smtp-vport-tmpl-list"`
    Pop3VportTmplList []VisibilityPacketCaptureObjectTemplatesPop3VportTmplList3118 `json:"pop3-vport-tmpl-list"`
    ImapVportTmplList []VisibilityPacketCaptureObjectTemplatesImapVportTmplList3122 `json:"imap-vport-tmpl-list"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList2986 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity2987 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc2988 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate2989 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity2987 struct {
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


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc2988 struct {
    DropVldGtpIeRepeatCountExceed int `json:"drop-vld-gtp-ie-repeat-count-exceed"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldMandatoryIeInGroupedIe int `json:"drop-vld-mandatory-ie-in-grouped-ie"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    DropVldGtpBearerCountExceed int `json:"drop-vld-gtp-bearer-count-exceed"`
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    DropVldV0ReservedMessageDrop int `json:"drop-vld-v0-reserved-message-drop"`
    DropVldV1ReservedMessageDrop int `json:"drop-vld-v1-reserved-message-drop"`
    DropVldV2ReservedMessageDrop int `json:"drop-vld-v2-reserved-message-drop"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate2989 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    DropVldGtpIeRepeatCountExceed int `json:"drop-vld-gtp-ie-repeat-count-exceed"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldMandatoryIeInGroupedIe int `json:"drop-vld-mandatory-ie-in-grouped-ie"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    DropVldGtpBearerCountExceed int `json:"drop-vld-gtp-bearer-count-exceed"`
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    DropVldV0ReservedMessageDrop int `json:"drop-vld-v0-reserved-message-drop"`
    DropVldV1ReservedMessageDrop int `json:"drop-vld-v1-reserved-message-drop"`
    DropVldV2ReservedMessageDrop int `json:"drop-vld-v2-reserved-message-drop"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList2990 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity2991 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc2992 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate2993 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity2991 struct {
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


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc2992 struct {
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate2993 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList2994 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity2995 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc2996 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate2997 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity2995 struct {
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


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc2996 struct {
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate2997 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList2998 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity2999 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc3000 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate3001 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity2999 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc3000 struct {
    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtOtherError int `json:"jwt-other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate3001 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtOtherError int `json:"jwt-other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList3002 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity3003 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc3004 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate3005 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity3003 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc3004 struct {
    Error int `json:"error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate3005 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Error int `json:"error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList3006 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity3007 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc3008 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate3009 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity3007 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc3008 struct {
    Spn_krb_faiure int `json:"spn_krb_faiure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate3009 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Spn_krb_faiure int `json:"spn_krb_faiure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList3010 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity3011 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc3012 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate3013 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity3011 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc3012 struct {
    AdminBindFailure int `json:"admin-bind-failure"`
    BindFailure int `json:"bind-failure"`
    SearchFailure int `json:"search-failure"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    SslSessionFailure int `json:"ssl-session-failure"`
    Pw_change_failure int `json:"pw_change_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate3013 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AdminBindFailure int `json:"admin-bind-failure"`
    BindFailure int `json:"bind-failure"`
    SearchFailure int `json:"search-failure"`
    AuthorizeFailure int `json:"authorize-failure"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    SslSessionFailure int `json:"ssl-session-failure"`
    Pw_change_failure int `json:"pw_change_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList3014 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity3015 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc3016 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate3017 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity3015 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc3016 struct {
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate3017 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList3018 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity3019 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc3020 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate3021 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity3019 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc3020 struct {
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate3021 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList3022 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity3023 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc3024 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate3025 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity3023 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc3024 struct {
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate3025 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Krb_timeout_error int `json:"krb_timeout_error"`
    Krb_other_error int `json:"krb_other_error"`
    Krb_pw_expiry int `json:"krb_pw_expiry"`
    Krb_pw_change_failure int `json:"krb_pw_change_failure"`
    Ntlm_proto_nego_failure int `json:"ntlm_proto_nego_failure"`
    Ntlm_session_setup_failure int `json:"ntlm_session_setup_failure"`
    Ntlm_prepare_req_error int `json:"ntlm_prepare_req_error"`
    Ntlm_auth_failure int `json:"ntlm_auth_failure"`
    Ntlm_timeout_error int `json:"ntlm_timeout_error"`
    Ntlm_other_error int `json:"ntlm_other_error"`
    Krb_validate_kdc_failure int `json:"krb_validate_kdc_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList3026 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity3027 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc3028 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate3029 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity3027 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc3028 struct {
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate3029 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList3030 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity3031 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc3032 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate3033 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity3031 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc3032 struct {
    MdFail int `json:"md-fail"`
    AcsFail int `json:"acs-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate3033 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MdFail int `json:"md-fail"`
    AcsFail int `json:"acs-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList3034 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity3035 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc3036 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate3037 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity3035 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc3036 struct {
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate3037 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList3038 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity3039 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc3040 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate3041 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity3039 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc3040 struct {
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate3041 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList3042 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity3043 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc3044 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate3045 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity3043 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc3044 struct {
    NoCreds int `json:"no-creds"`
    BadReq int `json:"bad-req"`
    Unauth int `json:"unauth"`
    Forbidden int `json:"forbidden"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    Unavailable int `json:"unavailable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate3045 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NoCreds int `json:"no-creds"`
    BadReq int `json:"bad-req"`
    Unauth int `json:"unauth"`
    Forbidden int `json:"forbidden"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    Unavailable int `json:"unavailable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList3046 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity3047 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc3048 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate3049 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity3047 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc3048 struct {
    Invalid_srv_rsp int `json:"invalid_srv_rsp"`
    Post_fail int `json:"post_fail"`
    Invalid_cred int `json:"invalid_cred"`
    Bad_req int `json:"bad_req"`
    Not_fnd int `json:"not_fnd"`
    Error int `json:"error"`
    Other_error int `json:"other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate3049 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Invalid_srv_rsp int `json:"invalid_srv_rsp"`
    Post_fail int `json:"post_fail"`
    Invalid_cred int `json:"invalid_cred"`
    Bad_req int `json:"bad_req"`
    Not_fnd int `json:"not_fnd"`
    Error int `json:"error"`
    Other_error int `json:"other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList3050 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity3051 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc3052 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate3053 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity3051 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc3052 struct {
    Failure int `json:"failure"`
    BufferAllocFail int `json:"buffer-alloc-fail"`
    EncodingFail int `json:"encoding-fail"`
    InsertHeaderFail int `json:"insert-header-fail"`
    ParseHeaderFail int `json:"parse-header-fail"`
    InternalError int `json:"internal-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate3053 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure int `json:"failure"`
    BufferAllocFail int `json:"buffer-alloc-fail"`
    EncodingFail int `json:"encoding-fail"`
    InsertHeaderFail int `json:"insert-header-fail"`
    ParseHeaderFail int `json:"parse-header-fail"`
    InternalError int `json:"internal-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList3054 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity3055 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc3056 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate3057 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity3055 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc3056 struct {
    Failure int `json:"failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate3057 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure int `json:"failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList3058 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity3059 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc3060 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate3061 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity3059 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc3060 struct {
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate3061 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList3062 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity3063 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc3064 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate3065 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity3063 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc3064 struct {
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate3065 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplList3066 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity3067 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc3068 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate3069 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity3067 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc3068 struct {
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Resp3xx int `json:"resp-3xx"`
    Resp4xx int `json:"resp-4xx"`
    Resp5xx int `json:"resp-5xx"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate3069 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Resp3xx int `json:"resp-3xx"`
    Resp4xx int `json:"resp-4xx"`
    Resp5xx int `json:"resp-5xx"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplList3070 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity3071 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc3072 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate3073 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity3071 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc3072 struct {
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate3073 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList3074 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity3075 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc3076 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate3077 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity3075 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc3076 struct {
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate3077 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList3078 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity3079 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc3080 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate3081 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity3079 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc3080 struct {
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate3081 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList3082 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity3083 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc3084 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate3085 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity3083 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc3084 struct {
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate3085 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList3086 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity3087 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc3088 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate3089 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity3087 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc3088 struct {
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate3089 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList3090 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity3091 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc3092 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate3093 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity3091 struct {
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


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc3092 struct {
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate3093 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplList3094 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity3095 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc3096 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate3097 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity3095 struct {
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


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc3096 struct {
    UnmatchedDrops int `json:"unmatched-drops"`
    Deny int `json:"deny"`
    Reset int `json:"reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate3097 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    UnmatchedDrops int `json:"unmatched-drops"`
    Deny int `json:"deny"`
    Reset int `json:"reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList3098 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity3099 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc3100 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate3101 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity3099 struct {
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


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc3100 struct {
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate3101 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList3102 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity3103 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc3104 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate3105 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity3103 struct {
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


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc3104 struct {
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate3105 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList3106 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity3107 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc3108 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate3109 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity3107 struct {
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


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc3108 struct {
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate3109 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplList3110 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity3111 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc3112 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate3113 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity3111 struct {
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


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc3112 struct {
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Total_filter_drop int `json:"total_filter_drop"`
    Total_max_query_len_drop int `json:"total_max_query_len_drop"`
    Rcode_notimpl_receive int `json:"rcode_notimpl_receive"`
    Rcode_notimpl_response int `json:"rcode_notimpl_response"`
    Gslb_query_bad int `json:"gslb_query_bad"`
    Gslb_response_bad int `json:"gslb_response_bad"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate3113 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Total_filter_drop int `json:"total_filter_drop"`
    Total_max_query_len_drop int `json:"total_max_query_len_drop"`
    Rcode_notimpl_receive int `json:"rcode_notimpl_receive"`
    Rcode_notimpl_response int `json:"rcode_notimpl_response"`
    Gslb_query_bad int `json:"gslb_query_bad"`
    Gslb_response_bad int `json:"gslb_response_bad"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList3114 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity3115 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc3116 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate3117 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity3115 struct {
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


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc3116 struct {
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


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate3117 struct {
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


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplList3118 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity3119 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc3120 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate3121 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity3119 struct {
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


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc3120 struct {
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


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate3121 struct {
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


type VisibilityPacketCaptureObjectTemplatesImapVportTmplList3122 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity3123 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc3124 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate3125 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity3123 struct {
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


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc3124 struct {
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


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate3125 struct {
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

func (p *VisibilityPacketCapture) GetId() string{
    return "1"
}

func (p *VisibilityPacketCapture) getPath() string{
    return "visibility/packet-capture"
}

func (p *VisibilityPacketCapture) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCapture::Post")
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

func (p *VisibilityPacketCapture) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCapture::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *VisibilityPacketCapture) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCapture::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCapture) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCapture::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
