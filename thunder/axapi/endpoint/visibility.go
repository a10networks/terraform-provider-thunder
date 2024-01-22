

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Visibility struct {
	Inst struct {

    AnomalyDetection VisibilityAnomalyDetection3141 `json:"anomaly-detection"`
    DebugFiles VisibilityDebugFiles3142 `json:"debug-files"`
    File VisibilityFile3143 `json:"file"`
    FlowCollector VisibilityFlowCollector3145 `json:"flow-collector"`
    Granularity int `json:"granularity" dval:"5"`
    InitialLearningInterval int `json:"initial-learning-interval"`
    MonEntityTelemetryData VisibilityMonEntityTelemetryData3153 `json:"mon-entity-telemetry-data"`
    MonTopk VisibilityMonTopk3155 `json:"mon-topk"`
    Monitor VisibilityMonitor3157 `json:"monitor"`
    MonitoredEntity VisibilityMonitoredEntity3171 `json:"monitored-entity"`
    PacketCapture VisibilityPacketCapture3180 `json:"packet-capture"`
    PingSweepDetection VisibilityPingSweepDetection3585 `json:"ping-sweep-detection"`
    PortScanDetection VisibilityPortScanDetection3586 `json:"port-scan-detection"`
    Reporting VisibilityReporting3587 `json:"reporting"`
    ResourceUsage VisibilityResourceUsage3596 `json:"resource-usage"`
    SamplingEnable []VisibilitySamplingEnable `json:"sampling-enable"`
    SourceEntityTopk int `json:"source-entity-topk"`
    Topn VisibilityTopn3597 `json:"topn"`
    Uuid string `json:"uuid"`
    Zbar VisibilityZbar3608 `json:"zbar"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"visibility"`
}


type VisibilityAnomalyDetection3141 struct {
    Sensitivity string `json:"sensitivity" dval:"low"`
    RestartLearningOnAnomaly int `json:"restart-learning-on-anomaly"`
    FeatureStatus string `json:"feature-status" dval:"disable"`
    Logging string `json:"logging" dval:"disable"`
    Uuid string `json:"uuid"`
}


type VisibilityDebugFiles3142 struct {
    Uuid string `json:"uuid"`
}


type VisibilityFile3143 struct {
    Metrics VisibilityFileMetrics3144 `json:"metrics"`
}


type VisibilityFileMetrics3144 struct {
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
}


type VisibilityFlowCollector3145 struct {
    Sflow VisibilityFlowCollectorSflow3146 `json:"sflow"`
    Netflow VisibilityFlowCollectorNetflow3148 `json:"netflow"`
}


type VisibilityFlowCollectorSflow3146 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityFlowCollectorSflowSamplingEnable3147 `json:"sampling-enable"`
}


type VisibilityFlowCollectorSflowSamplingEnable3147 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityFlowCollectorNetflow3148 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityFlowCollectorNetflowSamplingEnable3149 `json:"sampling-enable"`
    Template VisibilityFlowCollectorNetflowTemplate3150 `json:"template"`
}


type VisibilityFlowCollectorNetflowSamplingEnable3149 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityFlowCollectorNetflowTemplate3150 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityFlowCollectorNetflowTemplateSamplingEnable3151 `json:"sampling-enable"`
    Detail VisibilityFlowCollectorNetflowTemplateDetail3152 `json:"detail"`
}


type VisibilityFlowCollectorNetflowTemplateSamplingEnable3151 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityFlowCollectorNetflowTemplateDetail3152 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonEntityTelemetryData3153 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityMonEntityTelemetryDataSamplingEnable3154 `json:"sampling-enable"`
}


type VisibilityMonEntityTelemetryDataSamplingEnable3154 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityMonTopk3155 struct {
    Sources VisibilityMonTopkSources3156 `json:"sources"`
}


type VisibilityMonTopkSources3156 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitor3157 struct {
    PrimaryMonitor string `json:"primary-monitor"`
    MonitorKey string `json:"monitor-key"`
    MonEntityTopk int `json:"mon-entity-topk"`
    SourceEntityTopk int `json:"source-entity-topk"`
    IndexSessions int `json:"index-sessions"`
    IndexSessionsType string `json:"index-sessions-type"`
    Template VisibilityMonitorTemplate3158 `json:"template"`
    Uuid string `json:"uuid"`
    AgentList []VisibilityMonitorAgentList3160 `json:"agent-list"`
    Sflow VisibilityMonitorSflow3162 `json:"sflow"`
    Netflow VisibilityMonitorNetflow3163 `json:"netflow"`
    DebugList []VisibilityMonitorDebugList3164 `json:"debug-list"`
    ReplayDebugFile VisibilityMonitorReplayDebugFile3165 `json:"replay-debug-file"`
    DeleteDebugFile VisibilityMonitorDeleteDebugFile3166 `json:"delete-debug-file"`
    SecondaryMonitor VisibilityMonitorSecondaryMonitor3167 `json:"secondary-monitor"`
}


type VisibilityMonitorTemplate3158 struct {
    Notification []VisibilityMonitorTemplateNotification3159 `json:"notification"`
}


type VisibilityMonitorTemplateNotification3159 struct {
    NotifTemplateName string `json:"notif-template-name"`
}


type VisibilityMonitorAgentList3160 struct {
    AgentName string `json:"agent-name"`
    AgentV4Addr string `json:"agent-v4-addr"`
    AgentV6Addr string `json:"agent-v6-addr"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []VisibilityMonitorAgentListSamplingEnable3161 `json:"sampling-enable"`
}


type VisibilityMonitorAgentListSamplingEnable3161 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityMonitorSflow3162 struct {
    ListeningPort int `json:"listening-port" dval:"6343"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorNetflow3163 struct {
    ListeningPort int `json:"listening-port" dval:"9996"`
    TemplateActiveTimeout int `json:"template-active-timeout" dval:"30"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorDebugList3164 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorReplayDebugFile3165 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorDeleteDebugFile3166 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorSecondaryMonitor3167 struct {
    SecondaryMonitoringKey string `json:"secondary-monitoring-key"`
    MonEntityTopk int `json:"mon-entity-topk"`
    SourceEntityTopk int `json:"source-entity-topk"`
    Uuid string `json:"uuid"`
    DebugList []VisibilityMonitorSecondaryMonitorDebugList3168 `json:"debug-list"`
    DeleteDebugFile VisibilityMonitorSecondaryMonitorDeleteDebugFile3169 `json:"delete-debug-file"`
    ReplayDebugFile VisibilityMonitorSecondaryMonitorReplayDebugFile3170 `json:"replay-debug-file"`
}


type VisibilityMonitorSecondaryMonitorDebugList3168 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorSecondaryMonitorDeleteDebugFile3169 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorSecondaryMonitorReplayDebugFile3170 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitoredEntity3171 struct {
    Uuid string `json:"uuid"`
    Detail VisibilityMonitoredEntityDetail3172 `json:"detail"`
    Sessions VisibilityMonitoredEntitySessions3174 `json:"sessions"`
    MonTopk VisibilityMonitoredEntityMonTopk3175 `json:"mon-topk"`
    Secondary VisibilityMonitoredEntitySecondary3177 `json:"secondary"`
}


type VisibilityMonitoredEntityDetail3172 struct {
    Uuid string `json:"uuid"`
    Debug VisibilityMonitoredEntityDetailDebug3173 `json:"debug"`
}


type VisibilityMonitoredEntityDetailDebug3173 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitoredEntitySessions3174 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitoredEntityMonTopk3175 struct {
    Uuid string `json:"uuid"`
    Sources VisibilityMonitoredEntityMonTopkSources3176 `json:"sources"`
}


type VisibilityMonitoredEntityMonTopkSources3176 struct {
    Uuid string `json:"uuid"`
}


type VisibilityMonitoredEntitySecondary3177 struct {
    MonTopk VisibilityMonitoredEntitySecondaryMonTopk3178 `json:"mon-topk"`
}


type VisibilityMonitoredEntitySecondaryMonTopk3178 struct {
    Uuid string `json:"uuid"`
    Sources VisibilityMonitoredEntitySecondaryMonTopkSources3179 `json:"sources"`
}


type VisibilityMonitoredEntitySecondaryMonTopkSources3179 struct {
    Uuid string `json:"uuid"`
}


type VisibilityPacketCapture3180 struct {
    Uuid string `json:"uuid"`
    CaptureConfigList []VisibilityPacketCaptureCaptureConfigList3181 `json:"capture-config-list"`
    DeletePacketCaptureFile VisibilityPacketCaptureDeletePacketCaptureFile3182 `json:"delete-packet-capture-file"`
    GlobalTemplates VisibilityPacketCaptureGlobalTemplates3183 `json:"global-templates"`
    ObjectTemplates VisibilityPacketCaptureObjectTemplates3443 `json:"object-templates"`
    AutomatedCaptures VisibilityPacketCaptureAutomatedCaptures3584 `json:"automated-captures"`
}


type VisibilityPacketCaptureCaptureConfigList3181 struct {
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


type VisibilityPacketCaptureDeletePacketCaptureFile3182 struct {
    FileName string `json:"file-name"`
    All int `json:"all"`
}


type VisibilityPacketCaptureGlobalTemplates3183 struct {
    Uuid string `json:"uuid"`
    TemplateList []VisibilityPacketCaptureGlobalTemplatesTemplateList3184 `json:"template-list"`
    Activate VisibilityPacketCaptureGlobalTemplatesActivate3442 `json:"activate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateList3184 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerSysObjStatsSeverity VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsSeverity3185 `json:"trigger-sys-obj-stats-severity"`
    TriggerSysObjStatsChange VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChange3186 `json:"trigger-sys-obj-stats-change"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsSeverity3185 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChange3186 struct {
    Uuid string `json:"uuid"`
    SystemCtrLibAcct VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcct3187 `json:"system-ctr-lib-acct"`
    SystemHardwareAccelerate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerate3190 `json:"system-hardware-accelerate"`
    SystemRadiusServer VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServer3193 `json:"system-radius-server"`
    SystemIpThreatList VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatList3196 `json:"system-ip-threat-list"`
    SystemFpgaDrop VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDrop3199 `json:"system-fpga-drop"`
    SystemDpdkStats VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStats3202 `json:"system-dpdk-stats"`
    IpAnomalyDrop VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDrop3205 `json:"ip-anomaly-drop"`
    AamAuthenticationGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobal3208 `json:"aam-authentication-global"`
    AamRdns VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdns3211 `json:"aam-rdns"`
    AamAuthServerLdap VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdap3214 `json:"aam-auth-server-ldap"`
    AamAuthServerOcsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcsp3217 `json:"aam-auth-server-ocsp"`
    AamAuthServerRadius VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadius3220 `json:"aam-auth-server-radius"`
    AamAuthServerWin VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWin3223 `json:"aam-auth-server-win"`
    AamAuthAccount VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccount3226 `json:"aam-auth-account"`
    AamAuthSamlGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobal3229 `json:"aam-auth-saml-global"`
    AamAuthRelayKerberos VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberos3232 `json:"aam-auth-relay-kerberos"`
    AamAuthCaptcha VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptcha3235 `json:"aam-auth-captcha"`
    SlbSslError VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslError3238 `json:"slb-ssl-error"`
    SlbSslCertRevoke VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevoke3241 `json:"slb-ssl-cert-revoke"`
    SlbSslForwardProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxy3244 `json:"slb-ssl-forward-proxy"`
    VpnError VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnError3247 `json:"vpn-error"`
    Cgnv6Global VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Global3250 `json:"cgnv6-global"`
    Cgnv6DdosProc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProc3253 `json:"cgnv6-ddos-proc"`
    Cgnv6Lsn VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Lsn3256 `json:"cgnv6-lsn"`
    Cgnv6LsnAlgEsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEsp3259 `json:"cgnv6-lsn-alg-esp"`
    Cgnv6LsnAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptp3262 `json:"cgnv6-lsn-alg-pptp"`
    Cgnv6LsnAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtsp3265 `json:"cgnv6-lsn-alg-rtsp"`
    Cgnv6LsnAlgSip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSip3268 `json:"cgnv6-lsn-alg-sip"`
    Cgnv6LsnAlgMgcp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcp3271 `json:"cgnv6-lsn-alg-mgcp"`
    Cgnv6LsnAlgH323 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH3233274 `json:"cgnv6-lsn-alg-h323"`
    Cgnv6LsnRadius VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadius3277 `json:"cgnv6-lsn-radius"`
    Cgnv6Nat64Global VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64Global3280 `json:"cgnv6-nat64-global"`
    Cgnv6DsLiteGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobal3283 `json:"cgnv6-ds-lite-global"`
    Cgnv6FixedNatGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobal3286 `json:"cgnv6-fixed-nat-global"`
    Cgnv6FixedNatAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp3289 `json:"cgnv6-fixed-nat-alg-pptp"`
    Cgnv6FixedNatAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp3292 `json:"cgnv6-fixed-nat-alg-rtsp"`
    Cgnv6FixedNatAlgSip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSip3295 `json:"cgnv6-fixed-nat-alg-sip"`
    Cgnv6Pcp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Pcp3298 `json:"cgnv6-pcp"`
    Cgnv6Logging VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Logging3301 `json:"cgnv6-logging"`
    Cgnv6L4 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L43304 `json:"cgnv6-l4"`
    Cgnv6Icmp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Icmp3307 `json:"cgnv6-icmp"`
    Cgnv6HttpAlg VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlg3310 `json:"cgnv6-http-alg"`
    Cgnv6Dns64 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns643313 `json:"cgnv6-dns64"`
    Cgnv6Dhcpv6 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv63316 `json:"cgnv6-dhcpv6"`
    FwLogging VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLogging3319 `json:"fw-logging"`
    FwGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobal3322 `json:"fw-global"`
    FwAlgRtsp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtsp3325 `json:"fw-alg-rtsp"`
    FwAlgPptp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptp3328 `json:"fw-alg-pptp"`
    FwRadServer VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServer3331 `json:"fw-rad-server"`
    FwTcpSynCookie VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookie3334 `json:"fw-tcp-syn-cookie"`
    FwDdosProtection VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtection3337 `json:"fw-ddos-protection"`
    FwGtp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtp3340 `json:"fw-gtp"`
    SystemTcp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcp3343 `json:"system-tcp"`
    SlbConnReuse VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuse3346 `json:"slb-conn-reuse"`
    SlbAflow VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflow3349 `json:"slb-aflow"`
    SlbFix VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFix3352 `json:"slb-fix"`
    SlbSpdyProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxy3355 `json:"slb-spdy-proxy"`
    SlbHttp2 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp23358 `json:"slb-http2"`
    SlbL7session VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7session3361 `json:"slb-l7session"`
    SlbSmpp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmpp3364 `json:"slb-smpp"`
    SlbSmtp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtp3367 `json:"slb-smtp"`
    SlbMqtt VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqtt3370 `json:"slb-mqtt"`
    SlbIcap VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcap3373 `json:"slb-icap"`
    SlbSip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSip3376 `json:"slb-sip"`
    SlbHwCompress VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompress3379 `json:"slb-hw-compress"`
    SlbMysql VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysql3382 `json:"slb-mysql"`
    SlbMssql VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssql3385 `json:"slb-mssql"`
    SlbCrlSrcip VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcip3388 `json:"slb-crl-srcip"`
    SlbGeneric VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGeneric3391 `json:"slb-generic"`
    SlbPersist VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersist3394 `json:"slb-persist"`
    SlbHttpProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxy3397 `json:"slb-http-proxy"`
    SlbL4 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL43400 `json:"slb-l4"`
    SlbFastHttp VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttp3403 `json:"slb-fast-http"`
    SlbFtpProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxy3406 `json:"slb-ftp-proxy"`
    SlbImapProxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxy3409 `json:"slb-imap-proxy"`
    SlbPop3Proxy VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3Proxy3412 `json:"slb-pop3-proxy"`
    SlbSwitch VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitch3415 `json:"slb-switch"`
    SlbRcCache VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCache3418 `json:"slb-rc-cache"`
    SoCounters VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCounters3421 `json:"so-counters"`
    SlbPlyrIdGbl VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGbl3424 `json:"slb-plyr-id-gbl"`
    SlbSportRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRate3427 `json:"slb-sport-rate"`
    LoggingLocalLogGlobal VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobal3430 `json:"logging-local-log-global"`
    SlbMlb VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlb3433 `json:"slb-mlb"`
    SlbLinkProbe VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbe3436 `json:"slb-link-probe"`
    SlbRpz VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpz3439 `json:"slb-rpz"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcct3187 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc3188 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate3189 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc3188 struct {
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate3189 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerate3190 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc3191 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate3192 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc3191 struct {
    HwFwdProgErrors int `json:"hw-fwd-prog-errors"`
    HwFwdFlowSinglebitErrors int `json:"hw-fwd-flow-singlebit-errors"`
    HwFwdFlowTagMismatch int `json:"hw-fwd-flow-tag-mismatch"`
    HwFwdFlowSeqMismatch int `json:"hw-fwd-flow-seq-mismatch"`
    HwFwdFlowErrorCount int `json:"hw-fwd-flow-error-count"`
    HwFwdFlowUnalignCount int `json:"hw-fwd-flow-unalign-count"`
    HwFwdFlowUnderflowCount int `json:"hw-fwd-flow-underflow-count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate3192 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServer3193 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc3194 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate3195 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc3194 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate3195 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatList3196 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc3197 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate3198 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc3197 struct {
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate3198 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Error_out_of_memory int `json:"error_out_of_memory"`
    Error_out_of_spe_entries int `json:"error_out_of_spe_entries"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDrop3199 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc3200 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate3201 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc3200 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate3201 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStats3202 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc3203 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate3204 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc3203 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate3204 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDrop3205 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc3206 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate3207 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc3206 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate3207 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobal3208 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc3209 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate3210 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc3209 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate3210 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdns3211 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsInc3212 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsRate3213 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsInc3212 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsRate3213 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdap3214 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc3215 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate3216 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc3215 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate3216 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcsp3217 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc3218 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate3219 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc3218 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate3219 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadius3220 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc3221 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate3222 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc3221 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate3222 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWin3223 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc3224 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate3225 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc3224 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate3225 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccount3226 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc3227 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate3228 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc3227 struct {
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate3228 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    RequestDropped int `json:"request-dropped"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobal3229 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc3230 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate3231 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc3230 struct {
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate3231 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberos3232 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc3233 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate3234 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc3233 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate3234 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptcha3235 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc3236 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate3237 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc3236 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate3237 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslError3238 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc3239 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate3240 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc3239 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate3240 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevoke3241 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc3242 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate3243 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc3242 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate3243 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxy3244 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc3245 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate3246 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc3245 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate3246 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnError3247 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsInc3248 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsRate3249 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsInc3248 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsRate3249 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Global3250 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsInc3251 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsRate3252 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsInc3251 struct {
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6GlobalTriggerStatsRate3252 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    UdpTotalPortsAllocated int `json:"udp-total-ports-allocated"`
    IcmpTotalPortsAllocated int `json:"icmp-total-ports-allocated"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProc3253 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc3254 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate3255 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc3254 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate3255 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Lsn3256 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc3257 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate3258 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc3257 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate3258 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEsp3259 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc3260 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate3261 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc3260 struct {
    NatIpConflict int `json:"nat-ip-conflict"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate3261 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NatIpConflict int `json:"nat-ip-conflict"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptp3262 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc3263 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate3264 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc3263 struct {
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate3264 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NoGreSessionMatch int `json:"no-gre-session-match"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtsp3265 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc3266 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate3267 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc3266 struct {
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate3267 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    UnknownClientPortFromServer int `json:"unknown-client-port-from-server"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSip3268 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc3269 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate3270 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc3269 struct {
    MethodUnknown int `json:"method-unknown"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate3270 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MethodUnknown int `json:"method-unknown"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcp3271 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc3272 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate3273 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc3272 struct {
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate3273 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH3233274 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc3275 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate3276 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc3275 struct {
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate3276 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseError int `json:"parse-error"`
    TcpOutOfOrderDrop int `json:"tcp-out-of-order-drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadius3277 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc3278 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate3279 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc3278 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate3279 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64Global3280 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc3281 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate3282 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc3281 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate3282 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobal3283 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc3284 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate3285 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc3284 struct {
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    Fullcone_failure int `json:"fullcone_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate3285 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    User_quota_failure int `json:"user_quota_failure"`
    Nat_port_unavailable_tcp int `json:"nat_port_unavailable_tcp"`
    Nat_port_unavailable_udp int `json:"nat_port_unavailable_udp"`
    Nat_port_unavailable_icmp int `json:"nat_port_unavailable_icmp"`
    Fullcone_failure int `json:"fullcone_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobal3286 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc3287 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate3288 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc3287 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate3288 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp3289 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc3290 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate3291 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc3290 struct {
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate3291 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp3292 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc3293 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate3294 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc3293 struct {
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate3294 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    StreamCreationFailure int `json:"stream-creation-failure"`
    PortAllocationFailure int `json:"port-allocation-failure"`
    NoSessionMem int `json:"no-session-mem"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSip3295 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc3296 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate3297 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc3296 struct {
    MethodUnknown int `json:"method-unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate3297 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MethodUnknown int `json:"method-unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Pcp3298 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc3299 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate3300 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc3299 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate3300 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Logging3301 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc3302 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate3303 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc3302 struct {
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate3303 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    ConnTcpDropped int `json:"conn-tcp-dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L43304 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc3305 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate3306 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc3305 struct {
    OutOfSessionMemory int `json:"out-of-session-memory"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate3306 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    OutOfSessionMemory int `json:"out-of-session-memory"`
    IcmpHostUnreachableSent int `json:"icmp-host-unreachable-sent"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Icmp3307 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc3308 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate3309 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc3308 struct {
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate3309 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    IcmpToIcmpErr int `json:"icmp-to-icmp-err"`
    IcmpToIcmpv6Err int `json:"icmp-to-icmpv6-err"`
    Icmpv6ToIcmpErr int `json:"icmpv6-to-icmp-err"`
    Icmpv6ToIcmpv6Err int `json:"icmpv6-to-icmpv6-err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlg3310 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc3311 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate3312 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc3311 struct {
    RadiusRequstDropped int `json:"radius-requst-dropped"`
    RadiusResponseDropped int `json:"radius-response-dropped"`
    OutOfMemoryDropped int `json:"out-of-memory-dropped"`
    QueueLenExceedDropped int `json:"queue-len-exceed-dropped"`
    OutOfOrderDropped int `json:"out-of-order-dropped"`
    HeaderInsertionFailed int `json:"header-insertion-failed"`
    HeaderRemovalFailed int `json:"header-removal-failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate3312 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns643313 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc3314 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate3315 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc3314 struct {
    QueryBadPkt int `json:"query-bad-pkt"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    Drop int `json:"drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate3315 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    QueryBadPkt int `json:"query-bad-pkt"`
    RespBadPkt int `json:"resp-bad-pkt"`
    RespBadQr int `json:"resp-bad-qr"`
    Drop int `json:"drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv63316 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc3317 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate3318 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc3317 struct {
    PacketsDropped int `json:"packets-dropped"`
    PktsDroppedDuringClear int `json:"pkts-dropped-during-clear"`
    RcvNotSupportedMsg int `json:"rcv-not-supported-msg"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate3318 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    PacketsDropped int `json:"packets-dropped"`
    PktsDroppedDuringClear int `json:"pkts-dropped-during-clear"`
    RcvNotSupportedMsg int `json:"rcv-not-supported-msg"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLogging3319 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsInc3320 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsRate3321 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsInc3320 struct {
    LogDropped int `json:"log-dropped"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsRate3321 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    LogDropped int `json:"log-dropped"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobal3322 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsInc3323 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsRate3324 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsInc3323 struct {
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsRate3324 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Fullcone_creation_failure int `json:"fullcone_creation_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtsp3325 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc3326 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate3327 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc3326 struct {
    TransportAllocFailure int `json:"transport-alloc-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate3327 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    TransportAllocFailure int `json:"transport-alloc-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptp3328 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc3329 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate3330 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc3329 struct {
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate3330 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    CallReqPnsCallIdMismatch int `json:"call-req-pns-call-id-mismatch"`
    CallReplyPnsCallIdMismatch int `json:"call-reply-pns-call-id-mismatch"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServer3331 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsInc3332 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsRate3333 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsInc3332 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsRate3333 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookie3334 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc3335 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate3336 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc3335 struct {
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate3336 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Verification_failed int `json:"verification_failed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtection3337 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc3338 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate3339 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc3338 struct {
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate3339 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ddos_entries_too_many int `json:"ddos_entries_too_many"`
    Ddos_entry_add_to_bgp_failure int `json:"ddos_entry_add_to_bgp_failure"`
    Ddos_entry_remove_from_bgp_failure int `json:"ddos_entry_remove_from_bgp_failure"`
    Ddos_packet_dropped int `json:"ddos_packet_dropped"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtp3340 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsInc3341 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsRate3342 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsInc3341 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsRate3342 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcp3343 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsInc3344 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsRate3345 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsInc3344 struct {
    Attemptfails int `json:"attemptfails"`
    Noroute int `json:"noroute"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsRate3345 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Attemptfails int `json:"attemptfails"`
    Noroute int `json:"noroute"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuse3346 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc3347 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate3348 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc3347 struct {
    Ntermi_err int `json:"ntermi_err"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate3348 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Ntermi_err int `json:"ntermi_err"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflow3349 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsInc3350 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsRate3351 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsInc3350 struct {
    Pause_conn_fail int `json:"pause_conn_fail"`
    Error_resume_conn int `json:"error_resume_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsRate3351 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Pause_conn_fail int `json:"pause_conn_fail"`
    Error_resume_conn int `json:"error_resume_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFix3352 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsInc3353 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsRate3354 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsInc3353 struct {
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsRate3354 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Svrsel_fail int `json:"svrsel_fail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Client_err int `json:"client_err"`
    Server_err int `json:"server_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxy3355 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc3356 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate3357 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc3356 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate3357 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp23358 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc3359 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate3360 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc3359 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate3360 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7session3361 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc3362 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate3363 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc3362 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate3363 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmpp3364 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsInc3365 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsRate3366 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsInc3365 struct {
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_fail int `json:"select_server_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsRate3366 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Select_client_fail int `json:"select_client_fail"`
    Select_server_fail int `json:"select_server_fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtp3367 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc3368 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate3369 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc3368 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate3369 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqtt3370 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsInc3371 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsRate3372 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsInc3371 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsRate3372 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcap3373 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsInc3374 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsRate3375 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsInc3374 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsRate3375 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSip3376 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsInc3377 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsRate3378 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsInc3377 struct {
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsRate3378 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Msg_proxy_client_fail int `json:"msg_proxy_client_fail"`
    Msg_proxy_server_fail int `json:"msg_proxy_server_fail"`
    Msg_proxy_fail_start_server_conn int `json:"msg_proxy_fail_start_server_conn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompress3379 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc3380 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate3381 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc3380 struct {
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate3381 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure_count int `json:"failure_count"`
    Failure_code int `json:"failure_code"`
    Ring_full_count int `json:"ring_full_count"`
    Max_outstanding_request_count int `json:"max_outstanding_request_count"`
    Max_outstanding_submit_count int `json:"max_outstanding_submit_count"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysql3382 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc3383 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate3384 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc3383 struct {
    Session_err int `json:"session_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate3384 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Session_err int `json:"session_err"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssql3385 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc3386 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate3387 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc3386 struct {
    Session_err int `json:"session_err"`
    Auth_failure int `json:"auth_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate3387 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Session_err int `json:"session_err"`
    Auth_failure int `json:"auth_failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcip3388 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc3389 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate3390 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc3389 struct {
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Threshold_exceed int `json:"threshold_exceed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate3390 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Out_of_sessions int `json:"out_of_sessions"`
    Too_many_sessions int `json:"too_many_sessions"`
    Threshold_exceed int `json:"threshold_exceed"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGeneric3391 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsInc3392 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsRate3393 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsInc3392 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsRate3393 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersist3394 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsInc3395 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsRate3396 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsInc3395 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsRate3396 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxy3397 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc3398 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate3399 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc3398 struct {
    Parsereq_fail int `json:"parsereq_fail"`
    Svrsel_fail int `json:"svrsel_fail"`
    Fwdreq_fail int `json:"fwdreq_fail"`
    Fwdreqdata_fail int `json:"fwdreqdata_fail"`
    Snat_fail int `json:"snat_fail"`
    Req_over_limit int `json:"req_over_limit"`
    Req_rate_over_limit int `json:"req_rate_over_limit"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate3399 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL43400 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsInc3401 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsRate3402 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsInc3401 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsRate3402 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttp3403 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc3404 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate3405 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc3404 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate3405 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxy3406 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc3407 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate3408 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc3407 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate3408 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxy3409 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc3410 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate3411 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc3410 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate3411 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3Proxy3412 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc3413 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate3414 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc3413 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate3414 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitch3415 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc3416 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate3417 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc3416 struct {
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate3417 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCache3418 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc3419 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate3420 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc3419 struct {
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate3420 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCounters3421 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsInc3422 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsRate3423 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsInc3422 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsRate3423 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGbl3424 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc3425 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate3426 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc3425 struct {
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate3426 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_invalid_playerid_pkts int `json:"total_invalid_playerid_pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRate3427 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc3428 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate3429 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc3428 struct {
    Total_reset int `json:"total_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate3429 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_reset int `json:"total_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobal3430 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc3431 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate3432 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc3431 struct {
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate3432 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    EnqueueFull int `json:"enqueue-full"`
    EnqueueError int `json:"enqueue-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlb3433 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsInc3434 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsRate3435 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsInc3434 struct {
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsRate3435 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbe3436 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc3437 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate3438 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc3437 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate3438 struct {
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


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpz3439 struct {
    Uuid string `json:"uuid"`
    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsInc3440 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsRate3441 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsInc3440 struct {
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsRate3441 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Set_bw_error int `json:"set_bw_error"`
    Parse_error int `json:"parse_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesActivate3442 struct {
    Template string `json:"template"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplates3443 struct {
    Uuid string `json:"uuid"`
    TemplGtpPlcyTmplList []VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList3444 `json:"templ-gtp-plcy-tmpl-list"`
    InterfaceEthernetTmplList []VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList3448 `json:"interface-ethernet-tmpl-list"`
    InterfaceTunnelTmplList []VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList3452 `json:"interface-tunnel-tmpl-list"`
    AamJwtAuthorizationTmplList []VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList3456 `json:"aam-jwt-authorization-tmpl-list"`
    AamAaaPolicyTmplList []VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList3460 `json:"aam-aaa-policy-tmpl-list"`
    AamAuthLogonHttpInsTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList3464 `json:"aam-auth-logon-http-ins-tmpl-list"`
    AamAuthServerLdapInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList3468 `json:"aam-auth-server-ldap-inst-tmpl-list"`
    AamAuthServerOcspInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList3472 `json:"aam-auth-server-ocsp-inst-tmpl-list"`
    AamAuthServerRadInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList3476 `json:"aam-auth-server-rad-inst-tmpl-list"`
    AamAuthServerWinInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList3480 `json:"aam-auth-server-win-inst-tmpl-list"`
    AamAuthSamlServiceProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList3484 `json:"aam-auth-saml-service-prov-tmpl-list"`
    AamAuthSamlIdProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList3488 `json:"aam-auth-saml-id-prov-tmpl-list"`
    AamAuthServiceGroupTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList3492 `json:"aam-auth-service-group-tmpl-list"`
    AamAuthServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList3496 `json:"aam-auth-service-group-mem-tmpl-list"`
    AamAuthRelayHbaseInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList3500 `json:"aam-auth-relay-hbase-inst-tmpl-list"`
    AamAuthRelayFormInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList3504 `json:"aam-auth-relay-form-inst-tmpl-list"`
    AamAuthRelayNtlmTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList3508 `json:"aam-auth-relay-ntlm-tmpl-list"`
    AamAuthRelayWsFedTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList3512 `json:"aam-auth-relay-ws-fed-tmpl-list"`
    AamAuthCaptchaInstTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList3516 `json:"aam-auth-captcha-inst-tmpl-list"`
    SlbTemplCacheTmplList []VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList3520 `json:"slb-templ-cache-tmpl-list"`
    SlbPortTmplList []VisibilityPacketCaptureObjectTemplatesSlbPortTmplList3524 `json:"slb-port-tmpl-list"`
    SlbVportTmplList []VisibilityPacketCaptureObjectTemplatesSlbVportTmplList3528 `json:"slb-vport-tmpl-list"`
    Cgnv6ServGroupTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList3532 `json:"cgnv6-serv-group-tmpl-list"`
    Cgnv6Dns64VsPortTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList3536 `json:"cgnv6-dns64-vs-port-tmpl-list"`
    Cgnv6MapTransDomainTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList3540 `json:"cgnv6-map-trans-domain-tmpl-list"`
    Cgnv6EncapDomainTmplList []VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList3544 `json:"cgnv6-encap-domain-tmpl-list"`
    NetflowMonitorTmplList []VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList3548 `json:"netflow-monitor-tmpl-list"`
    RuleSetTmplList []VisibilityPacketCaptureObjectTemplatesRuleSetTmplList3552 `json:"rule-set-tmpl-list"`
    FwServerPortTmplList []VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList3556 `json:"fw-server-port-tmpl-list"`
    FwServiceGroupTmplList []VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList3560 `json:"fw-service-group-tmpl-list"`
    FwServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList3564 `json:"fw-service-group-mem-tmpl-list"`
    Dns_vportTmplList []VisibilityPacketCaptureObjectTemplatesDns_vportTmplList3568 `json:"dns_vport-tmpl-list"`
    SmtpVportTmplList []VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList3572 `json:"smtp-vport-tmpl-list"`
    Pop3VportTmplList []VisibilityPacketCaptureObjectTemplatesPop3VportTmplList3576 `json:"pop3-vport-tmpl-list"`
    ImapVportTmplList []VisibilityPacketCaptureObjectTemplatesImapVportTmplList3580 `json:"imap-vport-tmpl-list"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList3444 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity3445 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc3446 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate3447 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity3445 struct {
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


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc3446 struct {
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


type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate3447 struct {
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


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList3448 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity3449 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc3450 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate3451 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity3449 struct {
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


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc3450 struct {
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate3451 struct {
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


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList3452 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity3453 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc3454 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate3455 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity3453 struct {
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


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc3454 struct {
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate3455 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList3456 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity3457 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc3458 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate3459 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity3457 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc3458 struct {
    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtOtherError int `json:"jwt-other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate3459 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList3460 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity3461 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc3462 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate3463 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity3461 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc3462 struct {
    Error int `json:"error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate3463 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Error int `json:"error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList3464 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity3465 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc3466 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate3467 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity3465 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc3466 struct {
    Spn_krb_faiure int `json:"spn_krb_faiure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate3467 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Spn_krb_faiure int `json:"spn_krb_faiure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList3468 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity3469 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc3470 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate3471 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity3469 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc3470 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate3471 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList3472 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity3473 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc3474 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate3475 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity3473 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc3474 struct {
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate3475 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList3476 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity3477 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc3478 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate3479 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity3477 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc3478 struct {
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate3479 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Authen_failure int `json:"authen_failure"`
    Authorize_failure int `json:"authorize_failure"`
    Timeout_error int `json:"timeout_error"`
    Other_error int `json:"other_error"`
    AccountingFailure int `json:"accounting-failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList3480 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity3481 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc3482 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate3483 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity3481 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc3482 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate3483 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList3484 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity3485 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc3486 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate3487 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity3485 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc3486 struct {
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate3487 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    AcsAuthzFail int `json:"acs-authz-fail"`
    AcsError int `json:"acs-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList3488 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity3489 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc3490 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate3491 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity3489 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc3490 struct {
    MdFail int `json:"md-fail"`
    AcsFail int `json:"acs-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate3491 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    MdFail int `json:"md-fail"`
    AcsFail int `json:"acs-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList3492 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity3493 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc3494 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate3495 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity3493 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc3494 struct {
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate3495 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList3496 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity3497 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc3498 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate3499 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity3497 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc3498 struct {
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate3499 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList3500 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity3501 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc3502 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate3503 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity3501 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc3502 struct {
    NoCreds int `json:"no-creds"`
    BadReq int `json:"bad-req"`
    Unauth int `json:"unauth"`
    Forbidden int `json:"forbidden"`
    NotFound int `json:"not-found"`
    ServerError int `json:"server-error"`
    Unavailable int `json:"unavailable"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate3503 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList3504 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity3505 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc3506 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate3507 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity3505 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc3506 struct {
    Invalid_srv_rsp int `json:"invalid_srv_rsp"`
    Post_fail int `json:"post_fail"`
    Invalid_cred int `json:"invalid_cred"`
    Bad_req int `json:"bad_req"`
    Not_fnd int `json:"not_fnd"`
    Error int `json:"error"`
    Other_error int `json:"other_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate3507 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplList3508 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity3509 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc3510 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate3511 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsSeverity3509 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsInc3510 struct {
    Failure int `json:"failure"`
    BufferAllocFail int `json:"buffer-alloc-fail"`
    EncodingFail int `json:"encoding-fail"`
    InsertHeaderFail int `json:"insert-header-fail"`
    ParseHeaderFail int `json:"parse-header-fail"`
    InternalError int `json:"internal-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayNtlmTmplListTriggerStatsRate3511 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList3512 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity3513 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc3514 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate3515 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity3513 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc3514 struct {
    Failure int `json:"failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate3515 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Failure int `json:"failure"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList3516 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity3517 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc3518 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate3519 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity3517 struct {
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


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc3518 struct {
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate3519 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList3520 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity3521 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc3522 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate3523 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity3521 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc3522 struct {
    Nc_req_header int `json:"nc_req_header"`
    Nc_res_header int `json:"nc_res_header"`
    Rv_failure int `json:"rv_failure"`
    Content_toobig int `json:"content_toobig"`
    Content_toosmall int `json:"content_toosmall"`
    Entry_create_failures int `json:"entry_create_failures"`
    Header_save_error int `json:"header_save_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate3523 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplList3524 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity3525 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc3526 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate3527 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity3525 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc3526 struct {
    Es_resp_300 int `json:"es_resp_300"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Resp3xx int `json:"resp-3xx"`
    Resp4xx int `json:"resp-4xx"`
    Resp5xx int `json:"resp-5xx"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate3527 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplList3528 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity3529 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc3530 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate3531 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity3529 struct {
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


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc3530 struct {
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate3531 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList3532 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity3533 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc3534 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate3535 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity3533 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc3534 struct {
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate3535 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_drop int `json:"server_selection_fail_drop"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList3536 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity3537 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc3538 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate3539 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity3537 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc3538 struct {
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate3539 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList3540 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity3541 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc3542 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate3543 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity3541 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc3542 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate3543 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList3544 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity3545 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc3546 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate3547 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity3545 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc3546 struct {
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


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate3547 struct {
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


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList3548 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity3549 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc3550 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate3551 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity3549 struct {
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


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc3550 struct {
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


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate3551 struct {
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


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplList3552 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity3553 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc3554 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate3555 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity3553 struct {
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


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc3554 struct {
    UnmatchedDrops int `json:"unmatched-drops"`
    Deny int `json:"deny"`
    Reset int `json:"reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate3555 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    UnmatchedDrops int `json:"unmatched-drops"`
    Deny int `json:"deny"`
    Reset int `json:"reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList3556 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity3557 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc3558 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate3559 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity3557 struct {
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


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc3558 struct {
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate3559 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList3560 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity3561 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc3562 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate3563 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity3561 struct {
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


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc3562 struct {
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate3563 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Server_selection_fail_reset int `json:"server_selection_fail_reset"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList3564 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity3565 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc3566 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate3567 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity3565 struct {
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


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc3566 struct {
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate3567 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Curr_conn_overflow int `json:"curr_conn_overflow"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplList3568 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity3569 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc3570 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate3571 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity3569 struct {
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


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc3570 struct {
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


type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate3571 struct {
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


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList3572 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity3573 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc3574 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate3575 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity3573 struct {
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


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc3574 struct {
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


type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate3575 struct {
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


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplList3576 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity3577 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc3578 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate3579 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity3577 struct {
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


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc3578 struct {
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


type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate3579 struct {
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


type VisibilityPacketCaptureObjectTemplatesImapVportTmplList3580 struct {
    Name string `json:"name"`
    CaptureConfig string `json:"capture-config"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity3581 `json:"trigger-stats-severity"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc3582 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate3583 `json:"trigger-stats-rate"`
}


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity3581 struct {
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


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc3582 struct {
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


type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate3583 struct {
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


type VisibilityPacketCaptureAutomatedCaptures3584 struct {
    Slb_port_tmpl_error_code_return_inc int `json:"slb_port_tmpl_error_code_return_inc"`
    Slb_port_tmpl_high_error_code_return int `json:"slb_port_tmpl_high_error_code_return"`
    Uuid string `json:"uuid"`
}


type VisibilityPingSweepDetection3585 struct {
    Events int `json:"events" dval:"10"`
    Interval int `json:"interval" dval:"60"`
    V4List string `json:"v4-list"`
    V6List string `json:"v6-list"`
    Uuid string `json:"uuid"`
}


type VisibilityPortScanDetection3586 struct {
    Events int `json:"events" dval:"10"`
    Interval int `json:"interval" dval:"60"`
    V4List string `json:"v4-list"`
    V6List string `json:"v6-list"`
    Uuid string `json:"uuid"`
}


type VisibilityReporting3587 struct {
    SessionLogging string `json:"session-logging" dval:"disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityReportingSamplingEnable3588 `json:"sampling-enable"`
    TelemetryExportInterval VisibilityReportingTelemetryExportInterval3589 `json:"telemetry-export-interval"`
    Template VisibilityReportingTemplate3590 `json:"template"`
}


type VisibilityReportingSamplingEnable3588 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityReportingTelemetryExportInterval3589 struct {
    Value int `json:"value" dval:"5"`
    Uuid string `json:"uuid"`
}


type VisibilityReportingTemplate3590 struct {
    Notification VisibilityReportingTemplateNotification3591 `json:"notification"`
}


type VisibilityReportingTemplateNotification3591 struct {
    TemplateNameList []VisibilityReportingTemplateNotificationTemplateNameList3592 `json:"template-name-list"`
    Debug VisibilityReportingTemplateNotificationDebug3595 `json:"debug"`
}


type VisibilityReportingTemplateNotificationTemplateNameList3592 struct {
    Name string `json:"name"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    HostName string `json:"host-name"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Protocol string `json:"protocol" dval:"https"`
    HttpPort int `json:"http-port" dval:"80"`
    HttpsPort int `json:"https-port" dval:"443"`
    RelativeUri string `json:"relative-uri" dval:"/"`
    Action string `json:"action" dval:"enable"`
    DebugMode int `json:"debug-mode"`
    TestConnectivity int `json:"test-connectivity"`
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable3593 `json:"sampling-enable"`
    Authentication VisibilityReportingTemplateNotificationTemplateNameListAuthentication3594 `json:"authentication"`
}


type VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable3593 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityReportingTemplateNotificationTemplateNameListAuthentication3594 struct {
    RelativeLoginUri string `json:"relative-login-uri"`
    RelativeLogoffUri string `json:"relative-logoff-uri"`
    AuthUsername string `json:"auth-username"`
    AuthPassword int `json:"auth-password"`
    AuthPasswordString string `json:"auth-password-string"`
    Encrypted string `json:"encrypted"`
    ApiKey int `json:"api-key"`
    ApiKeyString string `json:"api-key-string"`
    ApiKeyEncrypted string `json:"api-key-encrypted"`
    Uuid string `json:"uuid"`
}


type VisibilityReportingTemplateNotificationDebug3595 struct {
    Uuid string `json:"uuid"`
}


type VisibilityResourceUsage3596 struct {
    Uuid string `json:"uuid"`
}


type VisibilitySamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VisibilityTopn3597 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VisibilityTopnSamplingEnable3598 `json:"sampling-enable"`
    Cgnv6NatPoolTopnTmplList []VisibilityTopnCgnv6NatPoolTopnTmplList3599 `json:"cgnv6-nat-pool-topn-tmpl-list"`
    Cgnv6NatPoolTopnNode VisibilityTopnCgnv6NatPoolTopnNode3601 `json:"cgnv6-nat-pool-topn-node"`
    GtpApnPrefixTopnTmplList []VisibilityTopnGtpApnPrefixTopnTmplList3602 `json:"gtp-apn-prefix-topn-tmpl-list"`
    GtpApnPrefixTopnNode VisibilityTopnGtpApnPrefixTopnNode3604 `json:"gtp-apn-prefix-topn-node"`
    GtpNetworkElementTopnTmplList []VisibilityTopnGtpNetworkElementTopnTmplList3605 `json:"gtp-network-element-topn-tmpl-list"`
    GtpNetworkElementTopnNode VisibilityTopnGtpNetworkElementTopnNode3607 `json:"gtp-network-element-topn-node"`
}


type VisibilityTopnSamplingEnable3598 struct {
    Counters1 string `json:"counters1"`
}


type VisibilityTopnCgnv6NatPoolTopnTmplList3599 struct {
    Name string `json:"name"`
    TopnSize int `json:"topn-size"`
    Interval string `json:"interval"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    Metrics VisibilityTopnCgnv6NatPoolTopnTmplListMetrics3600 `json:"metrics"`
}


type VisibilityTopnCgnv6NatPoolTopnTmplListMetrics3600 struct {
    UdpTotal int `json:"udp-total"`
    TcpTotal int `json:"tcp-total"`
    Uuid string `json:"uuid"`
}


type VisibilityTopnCgnv6NatPoolTopnNode3601 struct {
    Activate string `json:"activate"`
    Uuid string `json:"uuid"`
}


type VisibilityTopnGtpApnPrefixTopnTmplList3602 struct {
    Name string `json:"name"`
    TopnSize int `json:"topn-size"`
    Interval string `json:"interval"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    Metrics VisibilityTopnGtpApnPrefixTopnTmplListMetrics3603 `json:"metrics"`
}


type VisibilityTopnGtpApnPrefixTopnTmplListMetrics3603 struct {
    UplinkBytes int `json:"uplink-bytes"`
    DownlinkBytes int `json:"downlink-bytes"`
    UplinkPkts int `json:"uplink-pkts"`
    DownlinkPkts int `json:"downlink-pkts"`
    GtpV0CTunnelCreated int `json:"gtp-v0-c-tunnel-created"`
    GtpV0CTunnelHalfOpen int `json:"gtp-v0-c-tunnel-half-open"`
    GtpV0CTunnelHalfClosed int `json:"gtp-v0-c-tunnel-half-closed"`
    GtpV0CTunnelClosed int `json:"gtp-v0-c-tunnel-closed"`
    GtpV0CTunnelDeleted int `json:"gtp-v0-c-tunnel-deleted"`
    GtpV0CHalfOpenTunnelClosed int `json:"gtp-v0-c-half-open-tunnel-closed"`
    GtpV1CTunnelCreated int `json:"gtp-v1-c-tunnel-created"`
    GtpV1CTunnelHalfOpen int `json:"gtp-v1-c-tunnel-half-open"`
    GtpV1CTunnelHalfClosed int `json:"gtp-v1-c-tunnel-half-closed"`
    GtpV1CTunnelClosed int `json:"gtp-v1-c-tunnel-closed"`
    GtpV1CTunnelDeleted int `json:"gtp-v1-c-tunnel-deleted"`
    GtpV1CHalfOpenTunnelClosed int `json:"gtp-v1-c-half-open-tunnel-closed"`
    GtpV2CTunnelCreated int `json:"gtp-v2-c-tunnel-created"`
    GtpV2CTunnelHalfOpen int `json:"gtp-v2-c-tunnel-half-open"`
    GtpV2CTunnelHalfClosed int `json:"gtp-v2-c-tunnel-half-closed"`
    GtpV2CTunnelClosed int `json:"gtp-v2-c-tunnel-closed"`
    GtpV2CTunnelDeleted int `json:"gtp-v2-c-tunnel-deleted"`
    GtpV2CHalfOpenTunnelClosed int `json:"gtp-v2-c-half-open-tunnel-closed"`
    GtpUTunnelCreated int `json:"gtp-u-tunnel-created"`
    GtpUTunnelDeleted int `json:"gtp-u-tunnel-deleted"`
    GtpV0CUpdatePdpRespUnsuccess int `json:"gtp-v0-c-update-pdp-resp-unsuccess"`
    GtpV1CUpdatePdpRespUnsuccess int `json:"gtp-v1-c-update-pdp-resp-unsuccess"`
    GtpV2CMod_bearerRespUnsuccess int `json:"gtp-v2-c-mod_bearer-resp-unsuccess"`
    GtpV0CCreatePdpRespUnsuccess int `json:"gtp-v0-c-create-pdp-resp-unsuccess"`
    GtpV1CCreatePdpRespUnsuccess int `json:"gtp-v1-c-create-pdp-resp-unsuccess"`
    GtpV2CCreateSessRespUnsuccess int `json:"gtp-v2-c-create-sess-resp-unsuccess"`
    GtpV2CPiggybackMessage int `json:"gtp-v2-c-piggyback-message"`
    GtpPathManagementMessage int `json:"gtp-path-management-message"`
    GtpV0CTunnelDeletedRestart int `json:"gtp-v0-c-tunnel-deleted-restart"`
    GtpV1CTunnelDeletedRestart int `json:"gtp-v1-c-tunnel-deleted-restart"`
    GtpV2CTunnelDeletedRestart int `json:"gtp-v2-c-tunnel-deleted-restart"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldUnsupportedMessageType int `json:"drop-vld-unsupported-message-type"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
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
    GtpCHandoverInProgressWithConn int `json:"gtp-c-handover-in-progress-with-conn"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    DropFltMessageFiltering int `json:"drop-flt-message-filtering"`
    DropFltApnFiltering int `json:"drop-flt-apn-filtering"`
    DropFltMsisdnFiltering int `json:"drop-flt-msisdn-filtering"`
    DropFltRatTypeFiltering int `json:"drop-flt-rat-type-filtering"`
    DropFltGtpInGtp int `json:"drop-flt-gtp-in-gtp"`
    DropRlGtpV0CAgg int `json:"drop-rl-gtp-v0-c-agg"`
    DropRlGtpV1CAgg int `json:"drop-rl-gtp-v1-c-agg"`
    DropRlGtpV2CAgg int `json:"drop-rl-gtp-v2-c-agg"`
    DropRlGtpV1CCreatePdpRequest int `json:"drop-rl-gtp-v1-c-create-pdp-request"`
    DropRlGtpV2CCreateSessionRequest int `json:"drop-rl-gtp-v2-c-create-session-request"`
    DropRlGtpV1CUpdatePdpRequest int `json:"drop-rl-gtp-v1-c-update-pdp-request"`
    DropRlGtpV2CModifyBearerRequest int `json:"drop-rl-gtp-v2-c-modify-bearer-request"`
    DropRlGtpUTunnelCreate int `json:"drop-rl-gtp-u-tunnel-create"`
    DropRlGtpUUplinkByte int `json:"drop-rl-gtp-u-uplink-byte"`
    DropRlGtpUUplinkPacket int `json:"drop-rl-gtp-u-uplink-packet"`
    DropRlGtpUDownlinkByte int `json:"drop-rl-gtp-u-downlink-byte"`
    DropRlGtpUDownlinkPacket int `json:"drop-rl-gtp-u-downlink-packet"`
    DropRlGtpUTotalByte int `json:"drop-rl-gtp-u-total-byte"`
    DropRlGtpUTotalPacket int `json:"drop-rl-gtp-u-total-packet"`
    DropRlGtpUMaxConcurrentTunnels int `json:"drop-rl-gtp-u-max-concurrent-tunnels"`
    Uuid string `json:"uuid"`
}


type VisibilityTopnGtpApnPrefixTopnNode3604 struct {
    Activate string `json:"activate"`
    Uuid string `json:"uuid"`
}


type VisibilityTopnGtpNetworkElementTopnTmplList3605 struct {
    Name string `json:"name"`
    TopnSize int `json:"topn-size"`
    Interval string `json:"interval"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    Metrics VisibilityTopnGtpNetworkElementTopnTmplListMetrics3606 `json:"metrics"`
}


type VisibilityTopnGtpNetworkElementTopnTmplListMetrics3606 struct {
    UplinkBytes int `json:"uplink-bytes"`
    DownlinkBytes int `json:"downlink-bytes"`
    UplinkPkts int `json:"uplink-pkts"`
    DownlinkPkts int `json:"downlink-pkts"`
    GtpV0CTunnelCreated int `json:"gtp-v0-c-tunnel-created"`
    GtpV0CTunnelHalfOpen int `json:"gtp-v0-c-tunnel-half-open"`
    GtpV0CTunnelHalfClosed int `json:"gtp-v0-c-tunnel-half-closed"`
    GtpV0CTunnelClosed int `json:"gtp-v0-c-tunnel-closed"`
    GtpV0CTunnelDeleted int `json:"gtp-v0-c-tunnel-deleted"`
    GtpV0CHalfOpenTunnelClosed int `json:"gtp-v0-c-half-open-tunnel-closed"`
    GtpV1CTunnelCreated int `json:"gtp-v1-c-tunnel-created"`
    GtpV1CTunnelHalfOpen int `json:"gtp-v1-c-tunnel-half-open"`
    GtpV1CTunnelHalfClosed int `json:"gtp-v1-c-tunnel-half-closed"`
    GtpV1CTunnelClosed int `json:"gtp-v1-c-tunnel-closed"`
    GtpV1CTunnelDeleted int `json:"gtp-v1-c-tunnel-deleted"`
    GtpV1CHalfOpenTunnelClosed int `json:"gtp-v1-c-half-open-tunnel-closed"`
    GtpV2CTunnelCreated int `json:"gtp-v2-c-tunnel-created"`
    GtpV2CTunnelHalfOpen int `json:"gtp-v2-c-tunnel-half-open"`
    GtpV2CTunnelHalfClosed int `json:"gtp-v2-c-tunnel-half-closed"`
    GtpV2CTunnelClosed int `json:"gtp-v2-c-tunnel-closed"`
    GtpV2CTunnelDeleted int `json:"gtp-v2-c-tunnel-deleted"`
    GtpV2CHalfOpenTunnelClosed int `json:"gtp-v2-c-half-open-tunnel-closed"`
    GtpUTunnelCreated int `json:"gtp-u-tunnel-created"`
    GtpUTunnelDeleted int `json:"gtp-u-tunnel-deleted"`
    GtpV0CUpdatePdpRespUnsuccess int `json:"gtp-v0-c-update-pdp-resp-unsuccess"`
    GtpV1CUpdatePdpRespUnsuccess int `json:"gtp-v1-c-update-pdp-resp-unsuccess"`
    GtpV2CMod_bearerRespUnsuccess int `json:"gtp-v2-c-mod_bearer-resp-unsuccess"`
    GtpV0CCreatePdpRespUnsuccess int `json:"gtp-v0-c-create-pdp-resp-unsuccess"`
    GtpV1CCreatePdpRespUnsuccess int `json:"gtp-v1-c-create-pdp-resp-unsuccess"`
    GtpV2CCreateSessRespUnsuccess int `json:"gtp-v2-c-create-sess-resp-unsuccess"`
    GtpV2CPiggybackMessage int `json:"gtp-v2-c-piggyback-message"`
    GtpPathManagementMessage int `json:"gtp-path-management-message"`
    GtpV0CTunnelDeletedRestart int `json:"gtp-v0-c-tunnel-deleted-restart"`
    GtpV1CTunnelDeletedRestart int `json:"gtp-v1-c-tunnel-deleted-restart"`
    GtpV2CTunnelDeletedRestart int `json:"gtp-v2-c-tunnel-deleted-restart"`
    GtpV0CReservedMessageAllow int `json:"gtp-v0-c-reserved-message-allow"`
    GtpV1CReservedMessageAllow int `json:"gtp-v1-c-reserved-message-allow"`
    GtpV2CReservedMessageAllow int `json:"gtp-v2-c-reserved-message-allow"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldUnsupportedMessageType int `json:"drop-vld-unsupported-message-type"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    Drop_vldGtpBearerCountExceed int `json:"drop_vld-gtp-bearer-count-exceed"`
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    GtpCHandoverInProgressWithConn int `json:"gtp-c-handover-in-progress-with-conn"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    DropFltMessageFiltering int `json:"drop-flt-message-filtering"`
    DropFltApnFiltering int `json:"drop-flt-apn-filtering"`
    DropFltMsisdnFiltering int `json:"drop-flt-msisdn-filtering"`
    DropFltRatTypeFiltering int `json:"drop-flt-rat-type-filtering"`
    DropFltGtpInGtp int `json:"drop-flt-gtp-in-gtp"`
    DropRlGtpV0CAgg int `json:"drop-rl-gtp-v0-c-agg"`
    DropRlGtpV1CAgg int `json:"drop-rl-gtp-v1-c-agg"`
    DropRlGtpV2CAgg int `json:"drop-rl-gtp-v2-c-agg"`
    DropRlGtpV1CCreatePdpRequest int `json:"drop-rl-gtp-v1-c-create-pdp-request"`
    DropRlGtpV2CCreateSessionRequest int `json:"drop-rl-gtp-v2-c-create-session-request"`
    DropRlGtpV1CUpdatePdpRequest int `json:"drop-rl-gtp-v1-c-update-pdp-request"`
    DropRlGtpV2CModifyBearerRequest int `json:"drop-rl-gtp-v2-c-modify-bearer-request"`
    DropRlGtpUTunnelCreate int `json:"drop-rl-gtp-u-tunnel-create"`
    DropRlGtpUUplinkByte int `json:"drop-rl-gtp-u-uplink-byte"`
    DropRlGtpUUplinkPacket int `json:"drop-rl-gtp-u-uplink-packet"`
    DropRlGtpUDownlinkByte int `json:"drop-rl-gtp-u-downlink-byte"`
    DropRlGtpUDownlinkPacket int `json:"drop-rl-gtp-u-downlink-packet"`
    DropRlGtpUTotalByte int `json:"drop-rl-gtp-u-total-byte"`
    DropRlGtpUTotalPacket int `json:"drop-rl-gtp-u-total-packet"`
    DropRlGtpUMaxConcurrentTunnels int `json:"drop-rl-gtp-u-max-concurrent-tunnels"`
    Uuid string `json:"uuid"`
}


type VisibilityTopnGtpNetworkElementTopnNode3607 struct {
    Activate string `json:"activate"`
    Uuid string `json:"uuid"`
}


type VisibilityZbar3608 struct {
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
    Dest VisibilityZbarDest3609 `json:"dest"`
    Truples VisibilityZbarTruples3611 `json:"truples"`
}


type VisibilityZbarDest3609 struct {
    Uuid string `json:"uuid"`
    BadSources VisibilityZbarDestBadSources3610 `json:"bad-sources"`
}


type VisibilityZbarDestBadSources3610 struct {
    Uuid string `json:"uuid"`
}


type VisibilityZbarTruples3611 struct {
    Uuid string `json:"uuid"`
}

func (p *Visibility) GetId() string{
    return "1"
}

func (p *Visibility) getPath() string{
    return "visibility"
}

func (p *Visibility) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Visibility::Post")
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

func (p *Visibility) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Visibility::Get")
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
func (p *Visibility) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Visibility::Put")
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

func (p *Visibility) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Visibility::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
