package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type Visibility struct {
	Inst struct {
		AnomalyDetection VisibilityAnomalyDetection3270 `json:"anomaly-detection"`

		DebugFiles VisibilityDebugFiles3271 `json:"debug-files"`

		File VisibilityFile3272 `json:"file"`

		FlowCollector VisibilityFlowCollector3274 `json:"flow-collector"`

		Granularity int `json:"granularity" dval:"5"`

		InitialLearningInterval int `json:"initial-learning-interval"`

		MonEntityTelemetryData VisibilityMonEntityTelemetryData3282 `json:"mon-entity-telemetry-data"`

		MonTopk VisibilityMonTopk3284 `json:"mon-topk"`

		Monitor VisibilityMonitor3286 `json:"monitor"`

		MonitoredEntity VisibilityMonitoredEntity3300 `json:"monitored-entity"`

		PacketCapture VisibilityPacketCapture3309 `json:"packet-capture"`

		PingSweepDetection VisibilityPingSweepDetection3719 `json:"ping-sweep-detection"`

		PortScanDetection VisibilityPortScanDetection3720 `json:"port-scan-detection"`

		Reporting VisibilityReporting3721 `json:"reporting"`

		ResourceUsage VisibilityResourceUsage3730 `json:"resource-usage"`

		SamplingEnable []VisibilitySamplingEnable `json:"sampling-enable"`

		SourceEntityTopk int `json:"source-entity-topk"`

		Topn VisibilityTopn3731 `json:"topn"`

		Uuid string `json:"uuid"`

		Zbar VisibilityZbar3745 `json:"zbar"`

		FileContent []byte `json:"-"`

		FileHandle string `json:"file-handle"`
	} `json:"visibility"`
}

type VisibilityAnomalyDetection3270 struct {
	Sensitivity              string `json:"sensitivity" dval:"low"`
	RestartLearningOnAnomaly int    `json:"restart-learning-on-anomaly"`
	FeatureStatus            string `json:"feature-status" dval:"disable"`
	Logging                  string `json:"logging" dval:"disable"`
	Uuid                     string `json:"uuid"`
}

type VisibilityDebugFiles3271 struct {
	Uuid string `json:"uuid"`
}

type VisibilityFile3272 struct {
	Metrics VisibilityFileMetrics3273 `json:"metrics"`
}

type VisibilityFileMetrics3273 struct {
	Action string `json:"action" dval:"disable"`
	Uuid   string `json:"uuid"`
}

type VisibilityFlowCollector3274 struct {
	Sflow   VisibilityFlowCollectorSflow3275   `json:"sflow"`
	Netflow VisibilityFlowCollectorNetflow3277 `json:"netflow"`
}

type VisibilityFlowCollectorSflow3275 struct {
	Uuid           string                                           `json:"uuid"`
	SamplingEnable []VisibilityFlowCollectorSflowSamplingEnable3276 `json:"sampling-enable"`
}

type VisibilityFlowCollectorSflowSamplingEnable3276 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityFlowCollectorNetflow3277 struct {
	Uuid           string                                             `json:"uuid"`
	SamplingEnable []VisibilityFlowCollectorNetflowSamplingEnable3278 `json:"sampling-enable"`
	Template       VisibilityFlowCollectorNetflowTemplate3279         `json:"template"`
}

type VisibilityFlowCollectorNetflowSamplingEnable3278 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityFlowCollectorNetflowTemplate3279 struct {
	Uuid           string                                                     `json:"uuid"`
	SamplingEnable []VisibilityFlowCollectorNetflowTemplateSamplingEnable3280 `json:"sampling-enable"`
	Detail         VisibilityFlowCollectorNetflowTemplateDetail3281           `json:"detail"`
}

type VisibilityFlowCollectorNetflowTemplateSamplingEnable3280 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityFlowCollectorNetflowTemplateDetail3281 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonEntityTelemetryData3282 struct {
	Uuid           string                                               `json:"uuid"`
	SamplingEnable []VisibilityMonEntityTelemetryDataSamplingEnable3283 `json:"sampling-enable"`
}

type VisibilityMonEntityTelemetryDataSamplingEnable3283 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityMonTopk3284 struct {
	Sources VisibilityMonTopkSources3285 `json:"sources"`
}

type VisibilityMonTopkSources3285 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitor3286 struct {
	PrimaryMonitor    string                                `json:"primary-monitor"`
	MonitorKey        string                                `json:"monitor-key"`
	MonEntityTopk     int                                   `json:"mon-entity-topk"`
	SourceEntityTopk  int                                   `json:"source-entity-topk"`
	IndexSessions     int                                   `json:"index-sessions"`
	IndexSessionsType string                                `json:"index-sessions-type"`
	Template          VisibilityMonitorTemplate3287         `json:"template"`
	Uuid              string                                `json:"uuid"`
	AgentList         []VisibilityMonitorAgentList3289      `json:"agent-list"`
	Sflow             VisibilityMonitorSflow3291            `json:"sflow"`
	Netflow           VisibilityMonitorNetflow3292          `json:"netflow"`
	DebugList         []VisibilityMonitorDebugList3293      `json:"debug-list"`
	ReplayDebugFile   VisibilityMonitorReplayDebugFile3294  `json:"replay-debug-file"`
	DeleteDebugFile   VisibilityMonitorDeleteDebugFile3295  `json:"delete-debug-file"`
	SecondaryMonitor  VisibilityMonitorSecondaryMonitor3296 `json:"secondary-monitor"`
}

type VisibilityMonitorTemplate3287 struct {
	Notification []VisibilityMonitorTemplateNotification3288 `json:"notification"`
}

type VisibilityMonitorTemplateNotification3288 struct {
	NotifTemplateName string `json:"notif-template-name"`
}

type VisibilityMonitorAgentList3289 struct {
	AgentName      string                                         `json:"agent-name"`
	AgentV4Addr    string                                         `json:"agent-v4-addr"`
	AgentV6Addr    string                                         `json:"agent-v6-addr"`
	Uuid           string                                         `json:"uuid"`
	UserTag        string                                         `json:"user-tag"`
	SamplingEnable []VisibilityMonitorAgentListSamplingEnable3290 `json:"sampling-enable"`
}

type VisibilityMonitorAgentListSamplingEnable3290 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityMonitorSflow3291 struct {
	ListeningPort int    `json:"listening-port" dval:"6343"`
	Uuid          string `json:"uuid"`
}

type VisibilityMonitorNetflow3292 struct {
	ListeningPort         int    `json:"listening-port" dval:"9996"`
	TemplateActiveTimeout int    `json:"template-active-timeout" dval:"30"`
	Uuid                  string `json:"uuid"`
}

type VisibilityMonitorDebugList3293 struct {
	DebugIpAddr   string `json:"debug-ip-addr"`
	DebugPort     int    `json:"debug-port"`
	DebugProtocol string `json:"debug-protocol"`
	Uuid          string `json:"uuid"`
}

type VisibilityMonitorReplayDebugFile3294 struct {
	DebugIpAddr   string `json:"debug-ip-addr"`
	DebugPort     int    `json:"debug-port"`
	DebugProtocol string `json:"debug-protocol"`
}

type VisibilityMonitorDeleteDebugFile3295 struct {
	DebugIpAddr   string `json:"debug-ip-addr"`
	DebugPort     int    `json:"debug-port"`
	DebugProtocol string `json:"debug-protocol"`
}

type VisibilityMonitorSecondaryMonitor3296 struct {
	SecondaryMonitoringKey string                                               `json:"secondary-monitoring-key"`
	MonEntityTopk          int                                                  `json:"mon-entity-topk"`
	SourceEntityTopk       int                                                  `json:"source-entity-topk"`
	Uuid                   string                                               `json:"uuid"`
	DebugList              []VisibilityMonitorSecondaryMonitorDebugList3297     `json:"debug-list"`
	DeleteDebugFile        VisibilityMonitorSecondaryMonitorDeleteDebugFile3298 `json:"delete-debug-file"`
	ReplayDebugFile        VisibilityMonitorSecondaryMonitorReplayDebugFile3299 `json:"replay-debug-file"`
}

type VisibilityMonitorSecondaryMonitorDebugList3297 struct {
	DebugIpAddr   string `json:"debug-ip-addr"`
	DebugPort     int    `json:"debug-port"`
	DebugProtocol string `json:"debug-protocol"`
	Uuid          string `json:"uuid"`
}

type VisibilityMonitorSecondaryMonitorDeleteDebugFile3298 struct {
	DebugIpAddr   string `json:"debug-ip-addr"`
	DebugPort     int    `json:"debug-port"`
	DebugProtocol string `json:"debug-protocol"`
}

type VisibilityMonitorSecondaryMonitorReplayDebugFile3299 struct {
	DebugIpAddr   string `json:"debug-ip-addr"`
	DebugPort     int    `json:"debug-port"`
	DebugProtocol string `json:"debug-protocol"`
}

type VisibilityMonitoredEntity3300 struct {
	Uuid      string                                 `json:"uuid"`
	Detail    VisibilityMonitoredEntityDetail3301    `json:"detail"`
	Sessions  VisibilityMonitoredEntitySessions3303  `json:"sessions"`
	MonTopk   VisibilityMonitoredEntityMonTopk3304   `json:"mon-topk"`
	Secondary VisibilityMonitoredEntitySecondary3306 `json:"secondary"`
}

type VisibilityMonitoredEntityDetail3301 struct {
	Uuid  string                                   `json:"uuid"`
	Debug VisibilityMonitoredEntityDetailDebug3302 `json:"debug"`
}

type VisibilityMonitoredEntityDetailDebug3302 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitoredEntitySessions3303 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitoredEntityMonTopk3304 struct {
	Uuid    string                                      `json:"uuid"`
	Sources VisibilityMonitoredEntityMonTopkSources3305 `json:"sources"`
}

type VisibilityMonitoredEntityMonTopkSources3305 struct {
	Uuid string `json:"uuid"`
}

type VisibilityMonitoredEntitySecondary3306 struct {
	MonTopk VisibilityMonitoredEntitySecondaryMonTopk3307 `json:"mon-topk"`
}

type VisibilityMonitoredEntitySecondaryMonTopk3307 struct {
	Uuid    string                                               `json:"uuid"`
	Sources VisibilityMonitoredEntitySecondaryMonTopkSources3308 `json:"sources"`
}

type VisibilityMonitoredEntitySecondaryMonTopkSources3308 struct {
	Uuid string `json:"uuid"`
}

type VisibilityPacketCapture3309 struct {
	Uuid                    string                                             `json:"uuid"`
	CaptureConfigList       []VisibilityPacketCaptureCaptureConfigList3310     `json:"capture-config-list"`
	DeletePacketCaptureFile VisibilityPacketCaptureDeletePacketCaptureFile3311 `json:"delete-packet-capture-file"`
	GlobalTemplates         VisibilityPacketCaptureGlobalTemplates3312         `json:"global-templates"`
	ObjectTemplates         VisibilityPacketCaptureObjectTemplates3569         `json:"object-templates"`
	AutomatedCaptures       VisibilityPacketCaptureAutomatedCaptures3718       `json:"automated-captures"`
}

type VisibilityPacketCaptureCaptureConfigList3310 struct {
	Name                          string `json:"name"`
	Disable                       int    `json:"disable"`
	ConcurrentCaptures            int    `json:"concurrent-captures"`
	ConcurrentConnPerCapture      int    `json:"concurrent-conn-per-capture" dval:"100"`
	ConcurrentCapturesAge         int    `json:"concurrent-captures-age" dval:"1"`
	ConcurrentConnTag             int    `json:"concurrent-conn-tag"`
	NumberOfPacketsPerConn        int    `json:"number-of-packets-per-conn"`
	PacketLength                  int    `json:"packet-length" dval:"128"`
	FileSize                      int    `json:"file-size" dval:"1"`
	FileCount                     int    `json:"file-count" dval:"10"`
	NumberOfPacketsPerCapture     int    `json:"number-of-packets-per-capture"`
	NumberOfPacketsTotal          int    `json:"number-of-packets-total"`
	EnableContinuousGlobalCapture int    `json:"enable-continuous-global-capture"`
	CreatePcapFilesNow            int    `json:"create-pcap-files-now"`
	DisableAutoMerge              int    `json:"disable-auto-merge"`
	KeepPcapFilesAfterMerge       int    `json:"keep-pcap-files-after-merge"`
	Uuid                          string `json:"uuid"`
	UserTag                       string `json:"user-tag"`
}

type VisibilityPacketCaptureDeletePacketCaptureFile3311 struct {
	FileName string `json:"file-name"`
	All      int    `json:"all"`
}

type VisibilityPacketCaptureGlobalTemplates3312 struct {
	Uuid         string                                                   `json:"uuid"`
	TemplateList []VisibilityPacketCaptureGlobalTemplatesTemplateList3313 `json:"template-list"`
	Activate     VisibilityPacketCaptureGlobalTemplatesActivate3568       `json:"activate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateList3313 struct {
	Name                       string                                                                           `json:"name"`
	CaptureConfig              string                                                                           `json:"capture-config"`
	Uuid                       string                                                                           `json:"uuid"`
	UserTag                    string                                                                           `json:"user-tag"`
	TriggerSysObjStatsSeverity VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsSeverity3314 `json:"trigger-sys-obj-stats-severity"`
	TriggerSysObjStatsChange   VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChange3315   `json:"trigger-sys-obj-stats-change"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsSeverity3314 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChange3315 struct {
	Uuid                     string                                                                                                 `json:"uuid"`
	SystemCtrLibAcct         VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcct3316         `json:"system-ctr-lib-acct"`
	SystemHardwareAccelerate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerate3319 `json:"system-hardware-accelerate"`
	SystemRadiusServer       VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServer3322       `json:"system-radius-server"`
	SystemIpThreatList       VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatList3325       `json:"system-ip-threat-list"`
	SystemFpgaDrop           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDrop3328           `json:"system-fpga-drop"`
	SystemDpdkStats          VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStats3331          `json:"system-dpdk-stats"`
	IpAnomalyDrop            VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDrop3334            `json:"ip-anomaly-drop"`
	AamAuthenticationGlobal  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobal3337  `json:"aam-authentication-global"`
	AamRdns                  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdns3340                  `json:"aam-rdns"`
	AamAuthServerLdap        VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdap3343        `json:"aam-auth-server-ldap"`
	AamAuthServerOcsp        VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcsp3346        `json:"aam-auth-server-ocsp"`
	AamAuthServerRadius      VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadius3349      `json:"aam-auth-server-radius"`
	AamAuthServerWin         VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWin3352         `json:"aam-auth-server-win"`
	AamAuthAccount           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccount3355           `json:"aam-auth-account"`
	AamAuthSamlGlobal        VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobal3358        `json:"aam-auth-saml-global"`
	AamAuthRelayKerberos     VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberos3361     `json:"aam-auth-relay-kerberos"`
	AamAuthCaptcha           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptcha3364           `json:"aam-auth-captcha"`
	SlbSslError              VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslError3367              `json:"slb-ssl-error"`
	SlbSslCertRevoke         VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevoke3370         `json:"slb-ssl-cert-revoke"`
	SlbSslForwardProxy       VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxy3373       `json:"slb-ssl-forward-proxy"`
	VpnError                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnError3376                 `json:"vpn-error"`
	Cgnv6DdosProc            VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProc3379            `json:"cgnv6-ddos-proc"`
	Cgnv6Lsn                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Lsn3382                 `json:"cgnv6-lsn"`
	Cgnv6LsnAlgEsp           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEsp3385           `json:"cgnv6-lsn-alg-esp"`
	Cgnv6LsnAlgPptp          VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptp3388          `json:"cgnv6-lsn-alg-pptp"`
	Cgnv6LsnAlgRtsp          VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtsp3391          `json:"cgnv6-lsn-alg-rtsp"`
	Cgnv6LsnAlgSip           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSip3394           `json:"cgnv6-lsn-alg-sip"`
	Cgnv6LsnAlgMgcp          VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcp3397          `json:"cgnv6-lsn-alg-mgcp"`
	Cgnv6LsnAlgH323          VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH3233400          `json:"cgnv6-lsn-alg-h323"`
	Cgnv6LsnRadius           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadius3403           `json:"cgnv6-lsn-radius"`
	Cgnv6Nat64Global         VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64Global3406         `json:"cgnv6-nat64-global"`
	Cgnv6DsLiteGlobal        VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobal3409        `json:"cgnv6-ds-lite-global"`
	Cgnv6FixedNatGlobal      VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobal3412      `json:"cgnv6-fixed-nat-global"`
	Cgnv6FixedNatAlgPptp     VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp3415     `json:"cgnv6-fixed-nat-alg-pptp"`
	Cgnv6FixedNatAlgRtsp     VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp3418     `json:"cgnv6-fixed-nat-alg-rtsp"`
	Cgnv6FixedNatAlgSip      VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSip3421      `json:"cgnv6-fixed-nat-alg-sip"`
	Cgnv6Pcp                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Pcp3424                 `json:"cgnv6-pcp"`
	Cgnv6Logging             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Logging3427             `json:"cgnv6-logging"`
	Cgnv6L4                  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L43430                  `json:"cgnv6-l4"`
	Cgnv6Icmp                VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Icmp3433                `json:"cgnv6-icmp"`
	Cgnv6HttpAlg             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlg3436             `json:"cgnv6-http-alg"`
	Cgnv6Dns64               VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns643439               `json:"cgnv6-dns64"`
	Cgnv6Dhcpv6              VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv63442              `json:"cgnv6-dhcpv6"`
	FwLogging                VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLogging3445                `json:"fw-logging"`
	FwGlobal                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobal3448                 `json:"fw-global"`
	FwAlgRtsp                VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtsp3451                `json:"fw-alg-rtsp"`
	FwAlgPptp                VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptp3454                `json:"fw-alg-pptp"`
	FwRadServer              VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServer3457              `json:"fw-rad-server"`
	FwTcpSynCookie           VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookie3460           `json:"fw-tcp-syn-cookie"`
	FwDdosProtection         VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtection3463         `json:"fw-ddos-protection"`
	FwGtp                    VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtp3466                    `json:"fw-gtp"`
	SystemTcp                VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcp3469                `json:"system-tcp"`
	SlbConnReuse             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuse3472             `json:"slb-conn-reuse"`
	SlbAflow                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflow3475                 `json:"slb-aflow"`
	SlbFix                   VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFix3478                   `json:"slb-fix"`
	SlbSpdyProxy             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxy3481             `json:"slb-spdy-proxy"`
	SlbHttp2                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp23484                 `json:"slb-http2"`
	SlbL7session             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7session3487             `json:"slb-l7session"`
	SlbSmpp                  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmpp3490                  `json:"slb-smpp"`
	SlbSmtp                  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtp3493                  `json:"slb-smtp"`
	SlbMqtt                  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqtt3496                  `json:"slb-mqtt"`
	SlbIcap                  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcap3499                  `json:"slb-icap"`
	SlbSip                   VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSip3502                   `json:"slb-sip"`
	SlbHwCompress            VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompress3505            `json:"slb-hw-compress"`
	SlbMysql                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysql3508                 `json:"slb-mysql"`
	SlbMssql                 VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssql3511                 `json:"slb-mssql"`
	SlbCrlSrcip              VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcip3514              `json:"slb-crl-srcip"`
	SlbGeneric               VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGeneric3517               `json:"slb-generic"`
	SlbPersist               VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersist3520               `json:"slb-persist"`
	SlbHttpProxy             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxy3523             `json:"slb-http-proxy"`
	SlbL4                    VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL43526                    `json:"slb-l4"`
	SlbFastHttp              VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttp3529              `json:"slb-fast-http"`
	SlbFtpProxy              VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxy3532              `json:"slb-ftp-proxy"`
	SlbImapProxy             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxy3535             `json:"slb-imap-proxy"`
	SlbPop3Proxy             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3Proxy3538             `json:"slb-pop3-proxy"`
	SlbSwitch                VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitch3541                `json:"slb-switch"`
	SlbRcCache               VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCache3544               `json:"slb-rc-cache"`
	SoCounters               VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCounters3547               `json:"so-counters"`
	SlbPlyrIdGbl             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGbl3550             `json:"slb-plyr-id-gbl"`
	SlbSportRate             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRate3553             `json:"slb-sport-rate"`
	LoggingLocalLogGlobal    VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobal3556    `json:"logging-local-log-global"`
	SlbMlb                   VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlb3559                   `json:"slb-mlb"`
	SlbLinkProbe             VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbe3562             `json:"slb-link-probe"`
	SlbRpz                   VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpz3565                   `json:"slb-rpz"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcct3316 struct {
	Uuid             string                                                                                                         `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc3317  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate3318 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsInc3317 struct {
	TotalNodesFreeFailed   int    `json:"total-nodes-free-failed"`
	TotalNodesUnlinkFailed int    `json:"total-nodes-unlink-failed"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemCtrLibAcctTriggerStatsRate3318 struct {
	ThresholdExceededBy    int    `json:"threshold-exceeded-by" dval:"5"`
	Duration               int    `json:"duration" dval:"60"`
	TotalNodesFreeFailed   int    `json:"total-nodes-free-failed"`
	TotalNodesUnlinkFailed int    `json:"total-nodes-unlink-failed"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerate3319 struct {
	Uuid             string                                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc3320  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate3321 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsInc3320 struct {
	HwFwdProgErrors          int    `json:"hw-fwd-prog-errors"`
	HwFwdFlowSinglebitErrors int    `json:"hw-fwd-flow-singlebit-errors"`
	HwFwdFlowTagMismatch     int    `json:"hw-fwd-flow-tag-mismatch"`
	HwFwdFlowSeqMismatch     int    `json:"hw-fwd-flow-seq-mismatch"`
	HwFwdFlowErrorCount      int    `json:"hw-fwd-flow-error-count"`
	HwFwdFlowUnalignCount    int    `json:"hw-fwd-flow-unalign-count"`
	HwFwdFlowUnderflowCount  int    `json:"hw-fwd-flow-underflow-count"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemHardwareAccelerateTriggerStatsRate3321 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	HwFwdProgErrors          int    `json:"hw-fwd-prog-errors"`
	HwFwdFlowSinglebitErrors int    `json:"hw-fwd-flow-singlebit-errors"`
	HwFwdFlowTagMismatch     int    `json:"hw-fwd-flow-tag-mismatch"`
	HwFwdFlowSeqMismatch     int    `json:"hw-fwd-flow-seq-mismatch"`
	HwFwdFlowErrorCount      int    `json:"hw-fwd-flow-error-count"`
	HwFwdFlowUnalignCount    int    `json:"hw-fwd-flow-unalign-count"`
	HwFwdFlowUnderflowCount  int    `json:"hw-fwd-flow-underflow-count"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServer3322 struct {
	Uuid             string                                                                                                           `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc3323  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate3324 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsInc3323 struct {
	RadiusRequestDropped       int    `json:"radius-request-dropped"`
	RequestBadSecretDropped    int    `json:"request-bad-secret-dropped"`
	RequestNoKeyVapDropped     int    `json:"request-no-key-vap-dropped"`
	RequestMalformedDropped    int    `json:"request-malformed-dropped"`
	RadiusTableFull            int    `json:"radius-table-full"`
	SecretNotConfiguredDropped int    `json:"secret-not-configured-dropped"`
	HaStandbyDropped           int    `json:"ha-standby-dropped"`
	Ipv6PrefixLengthMismatch   int    `json:"ipv6-prefix-length-mismatch"`
	InvalidKey                 int    `json:"invalid-key"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemRadiusServerTriggerStatsRate3324 struct {
	ThresholdExceededBy        int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                   int    `json:"duration" dval:"60"`
	RadiusRequestDropped       int    `json:"radius-request-dropped"`
	RequestBadSecretDropped    int    `json:"request-bad-secret-dropped"`
	RequestNoKeyVapDropped     int    `json:"request-no-key-vap-dropped"`
	RequestMalformedDropped    int    `json:"request-malformed-dropped"`
	RadiusTableFull            int    `json:"radius-table-full"`
	SecretNotConfiguredDropped int    `json:"secret-not-configured-dropped"`
	HaStandbyDropped           int    `json:"ha-standby-dropped"`
	Ipv6PrefixLengthMismatch   int    `json:"ipv6-prefix-length-mismatch"`
	InvalidKey                 int    `json:"invalid-key"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatList3325 struct {
	Uuid             string                                                                                                           `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc3326  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate3327 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsInc3326 struct {
	Error_out_of_memory      int    `json:"error_out_of_memory"`
	Error_out_of_spe_entries int    `json:"error_out_of_spe_entries"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemIpThreatListTriggerStatsRate3327 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	Error_out_of_memory      int    `json:"error_out_of_memory"`
	Error_out_of_spe_entries int    `json:"error_out_of_spe_entries"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDrop3328 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc3329  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate3330 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsInc3329 struct {
	MrxDrop           int    `json:"mrx-drop"`
	HrxDrop           int    `json:"hrx-drop"`
	SizDrop           int    `json:"siz-drop"`
	FcsDrop           int    `json:"fcs-drop"`
	LandDrop          int    `json:"land-drop"`
	EmptyFragDrop     int    `json:"empty-frag-drop"`
	MicFragDrop       int    `json:"mic-frag-drop"`
	Ipv4OptDrop       int    `json:"ipv4-opt-drop"`
	Ipv4Frag          int    `json:"ipv4-frag"`
	BadIpHdrLen       int    `json:"bad-ip-hdr-len"`
	BadIpFlagsDrop    int    `json:"bad-ip-flags-drop"`
	BadIpTtlDrop      int    `json:"bad-ip-ttl-drop"`
	NoIpPayloadDrop   int    `json:"no-ip-payload-drop"`
	OversizeIpPayload int    `json:"oversize-ip-payload"`
	BadIpPayloadLen   int    `json:"bad-ip-payload-len"`
	BadIpFragOffset   int    `json:"bad-ip-frag-offset"`
	BadIpChksumDrop   int    `json:"bad-ip-chksum-drop"`
	IcmpPodDrop       int    `json:"icmp-pod-drop"`
	TcpBadUrgOffet    int    `json:"tcp-bad-urg-offet"`
	TcpShortHdr       int    `json:"tcp-short-hdr"`
	TcpBadIpLen       int    `json:"tcp-bad-ip-len"`
	TcpNullFlags      int    `json:"tcp-null-flags"`
	TcpNullScan       int    `json:"tcp-null-scan"`
	TcpFinSin         int    `json:"tcp-fin-sin"`
	TcpXmasFlags      int    `json:"tcp-xmas-flags"`
	TcpXmasScan       int    `json:"tcp-xmas-scan"`
	TcpSynFrag        int    `json:"tcp-syn-frag"`
	TcpFragHdr        int    `json:"tcp-frag-hdr"`
	TcpBadChksum      int    `json:"tcp-bad-chksum"`
	UdpShortHdr       int    `json:"udp-short-hdr"`
	UdpBadIpLen       int    `json:"udp-bad-ip-len"`
	UdpKbFrags        int    `json:"udp-kb-frags"`
	UdpPortLb         int    `json:"udp-port-lb"`
	UdpBadChksum      int    `json:"udp-bad-chksum"`
	RuntIpHdr         int    `json:"runt-ip-hdr"`
	RuntTcpudpHdr     int    `json:"runt-tcpudp-hdr"`
	TunMismatch       int    `json:"tun-mismatch"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemFpgaDropTriggerStatsRate3330 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	MrxDrop             int    `json:"mrx-drop"`
	HrxDrop             int    `json:"hrx-drop"`
	SizDrop             int    `json:"siz-drop"`
	FcsDrop             int    `json:"fcs-drop"`
	LandDrop            int    `json:"land-drop"`
	EmptyFragDrop       int    `json:"empty-frag-drop"`
	MicFragDrop         int    `json:"mic-frag-drop"`
	Ipv4OptDrop         int    `json:"ipv4-opt-drop"`
	Ipv4Frag            int    `json:"ipv4-frag"`
	BadIpHdrLen         int    `json:"bad-ip-hdr-len"`
	BadIpFlagsDrop      int    `json:"bad-ip-flags-drop"`
	BadIpTtlDrop        int    `json:"bad-ip-ttl-drop"`
	NoIpPayloadDrop     int    `json:"no-ip-payload-drop"`
	OversizeIpPayload   int    `json:"oversize-ip-payload"`
	BadIpPayloadLen     int    `json:"bad-ip-payload-len"`
	BadIpFragOffset     int    `json:"bad-ip-frag-offset"`
	BadIpChksumDrop     int    `json:"bad-ip-chksum-drop"`
	IcmpPodDrop         int    `json:"icmp-pod-drop"`
	TcpBadUrgOffet      int    `json:"tcp-bad-urg-offet"`
	TcpShortHdr         int    `json:"tcp-short-hdr"`
	TcpBadIpLen         int    `json:"tcp-bad-ip-len"`
	TcpNullFlags        int    `json:"tcp-null-flags"`
	TcpNullScan         int    `json:"tcp-null-scan"`
	TcpFinSin           int    `json:"tcp-fin-sin"`
	TcpXmasFlags        int    `json:"tcp-xmas-flags"`
	TcpXmasScan         int    `json:"tcp-xmas-scan"`
	TcpSynFrag          int    `json:"tcp-syn-frag"`
	TcpFragHdr          int    `json:"tcp-frag-hdr"`
	TcpBadChksum        int    `json:"tcp-bad-chksum"`
	UdpShortHdr         int    `json:"udp-short-hdr"`
	UdpBadIpLen         int    `json:"udp-bad-ip-len"`
	UdpKbFrags          int    `json:"udp-kb-frags"`
	UdpPortLb           int    `json:"udp-port-lb"`
	UdpBadChksum        int    `json:"udp-bad-chksum"`
	RuntIpHdr           int    `json:"runt-ip-hdr"`
	RuntTcpudpHdr       int    `json:"runt-tcpudp-hdr"`
	TunMismatch         int    `json:"tun-mismatch"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStats3331 struct {
	Uuid             string                                                                                                        `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc3332  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate3333 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsInc3332 struct {
	PktDrop          int    `json:"pkt-drop"`
	PktLnkDownDrop   int    `json:"pkt-lnk-down-drop"`
	ErrPktDrop       int    `json:"err-pkt-drop"`
	RxErr            int    `json:"rx-err"`
	TxErr            int    `json:"tx-err"`
	TxDrop           int    `json:"tx-drop"`
	RxLenErr         int    `json:"rx-len-err"`
	RxOverErr        int    `json:"rx-over-err"`
	RxCrcErr         int    `json:"rx-crc-err"`
	RxFrameErr       int    `json:"rx-frame-err"`
	RxNoBuffErr      int    `json:"rx-no-buff-err"`
	RxMissErr        int    `json:"rx-miss-err"`
	TxAbortErr       int    `json:"tx-abort-err"`
	TxCarrierErr     int    `json:"tx-carrier-err"`
	TxFifoErr        int    `json:"tx-fifo-err"`
	TxHbeatErr       int    `json:"tx-hbeat-err"`
	TxWindowsErr     int    `json:"tx-windows-err"`
	RxLongLenErr     int    `json:"rx-long-len-err"`
	RxShortLenErr    int    `json:"rx-short-len-err"`
	RxAlignErr       int    `json:"rx-align-err"`
	RxCsumOffloadErr int    `json:"rx-csum-offload-err"`
	IoRxQueDrop      int    `json:"io-rx-que-drop"`
	IoTxQueDrop      int    `json:"io-tx-que-drop"`
	IoRingDrop       int    `json:"io-ring-drop"`
	WTxQueDrop       int    `json:"w-tx-que-drop"`
	WLinkDownDrop    int    `json:"w-link-down-drop"`
	WRingDrop        int    `json:"w-ring-drop"`
	Uuid             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemDpdkStatsTriggerStatsRate3333 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	PktDrop             int    `json:"pkt-drop"`
	PktLnkDownDrop      int    `json:"pkt-lnk-down-drop"`
	ErrPktDrop          int    `json:"err-pkt-drop"`
	RxErr               int    `json:"rx-err"`
	TxErr               int    `json:"tx-err"`
	TxDrop              int    `json:"tx-drop"`
	RxLenErr            int    `json:"rx-len-err"`
	RxOverErr           int    `json:"rx-over-err"`
	RxCrcErr            int    `json:"rx-crc-err"`
	RxFrameErr          int    `json:"rx-frame-err"`
	RxNoBuffErr         int    `json:"rx-no-buff-err"`
	RxMissErr           int    `json:"rx-miss-err"`
	TxAbortErr          int    `json:"tx-abort-err"`
	TxCarrierErr        int    `json:"tx-carrier-err"`
	TxFifoErr           int    `json:"tx-fifo-err"`
	TxHbeatErr          int    `json:"tx-hbeat-err"`
	TxWindowsErr        int    `json:"tx-windows-err"`
	RxLongLenErr        int    `json:"rx-long-len-err"`
	RxShortLenErr       int    `json:"rx-short-len-err"`
	RxAlignErr          int    `json:"rx-align-err"`
	RxCsumOffloadErr    int    `json:"rx-csum-offload-err"`
	IoRxQueDrop         int    `json:"io-rx-que-drop"`
	IoTxQueDrop         int    `json:"io-tx-que-drop"`
	IoRingDrop          int    `json:"io-ring-drop"`
	WTxQueDrop          int    `json:"w-tx-que-drop"`
	WLinkDownDrop       int    `json:"w-link-down-drop"`
	WRingDrop           int    `json:"w-ring-drop"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDrop3334 struct {
	Uuid             string                                                                                                      `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc3335  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate3336 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc3335 struct {
	Land               int    `json:"land"`
	Emp_frg            int    `json:"emp_frg"`
	Emp_mic_frg        int    `json:"emp_mic_frg"`
	Opt                int    `json:"opt"`
	Frg                int    `json:"frg"`
	Bad_ip_hdrlen      int    `json:"bad_ip_hdrlen"`
	Bad_ip_flg         int    `json:"bad_ip_flg"`
	Bad_ip_ttl         int    `json:"bad_ip_ttl"`
	No_ip_payload      int    `json:"no_ip_payload"`
	Over_ip_payload    int    `json:"over_ip_payload"`
	Bad_ip_payload_len int    `json:"bad_ip_payload_len"`
	Bad_ip_frg_offset  int    `json:"bad_ip_frg_offset"`
	Csum               int    `json:"csum"`
	Pod                int    `json:"pod"`
	Bad_tcp_urg_offset int    `json:"bad_tcp_urg_offset"`
	Tcp_sht_hdr        int    `json:"tcp_sht_hdr"`
	Tcp_bad_iplen      int    `json:"tcp_bad_iplen"`
	Tcp_null_frg       int    `json:"tcp_null_frg"`
	Tcp_null_scan      int    `json:"tcp_null_scan"`
	Tcp_syn_fin        int    `json:"tcp_syn_fin"`
	Tcp_xmas           int    `json:"tcp_xmas"`
	Tcp_xmas_scan      int    `json:"tcp_xmas_scan"`
	Tcp_syn_frg        int    `json:"tcp_syn_frg"`
	Tcp_frg_hdr        int    `json:"tcp_frg_hdr"`
	Tcp_bad_csum       int    `json:"tcp_bad_csum"`
	Udp_srt_hdr        int    `json:"udp_srt_hdr"`
	Udp_bad_len        int    `json:"udp_bad_len"`
	Udp_kerb_frg       int    `json:"udp_kerb_frg"`
	Udp_port_lb        int    `json:"udp_port_lb"`
	Udp_bad_csum       int    `json:"udp_bad_csum"`
	Runt_ip_hdr        int    `json:"runt_ip_hdr"`
	Runt_tcp_udp_hdr   int    `json:"runt_tcp_udp_hdr"`
	Ipip_tnl_msmtch    int    `json:"ipip_tnl_msmtch"`
	Tcp_opt_err        int    `json:"tcp_opt_err"`
	Ipip_tnl_err       int    `json:"ipip_tnl_err"`
	Vxlan_err          int    `json:"vxlan_err"`
	Nvgre_err          int    `json:"nvgre_err"`
	Gre_pptp_err       int    `json:"gre_pptp_err"`
	Ipv6_eh_hbh        int    `json:"ipv6_eh_hbh"`
	Ipv6_eh_dest       int    `json:"ipv6_eh_dest"`
	Ipv6_eh_routing    int    `json:"ipv6_eh_routing"`
	Ipv6_eh_frag       int    `json:"ipv6_eh_frag"`
	Ipv6_eh_ah         int    `json:"ipv6_eh_ah"`
	Ipv6_eh_esp        int    `json:"ipv6_eh_esp"`
	Ipv6_eh_mobility   int    `json:"ipv6_eh_mobility"`
	Ipv6_eh_none       int    `json:"ipv6_eh_none"`
	Ipv6_eh_other      int    `json:"ipv6_eh_other"`
	Ipv6_eh_malformed  int    `json:"ipv6_eh_malformed"`
	Tcp_udp_zero_port  int    `json:"tcp_udp_zero_port"`
	Uuid               string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsRate3336 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Land                int    `json:"land"`
	Emp_frg             int    `json:"emp_frg"`
	Emp_mic_frg         int    `json:"emp_mic_frg"`
	Opt                 int    `json:"opt"`
	Frg                 int    `json:"frg"`
	Bad_ip_hdrlen       int    `json:"bad_ip_hdrlen"`
	Bad_ip_flg          int    `json:"bad_ip_flg"`
	Bad_ip_ttl          int    `json:"bad_ip_ttl"`
	No_ip_payload       int    `json:"no_ip_payload"`
	Over_ip_payload     int    `json:"over_ip_payload"`
	Bad_ip_payload_len  int    `json:"bad_ip_payload_len"`
	Bad_ip_frg_offset   int    `json:"bad_ip_frg_offset"`
	Csum                int    `json:"csum"`
	Pod                 int    `json:"pod"`
	Bad_tcp_urg_offset  int    `json:"bad_tcp_urg_offset"`
	Tcp_sht_hdr         int    `json:"tcp_sht_hdr"`
	Tcp_bad_iplen       int    `json:"tcp_bad_iplen"`
	Tcp_null_frg        int    `json:"tcp_null_frg"`
	Tcp_null_scan       int    `json:"tcp_null_scan"`
	Tcp_syn_fin         int    `json:"tcp_syn_fin"`
	Tcp_xmas            int    `json:"tcp_xmas"`
	Tcp_xmas_scan       int    `json:"tcp_xmas_scan"`
	Tcp_syn_frg         int    `json:"tcp_syn_frg"`
	Tcp_frg_hdr         int    `json:"tcp_frg_hdr"`
	Tcp_bad_csum        int    `json:"tcp_bad_csum"`
	Udp_srt_hdr         int    `json:"udp_srt_hdr"`
	Udp_bad_len         int    `json:"udp_bad_len"`
	Udp_kerb_frg        int    `json:"udp_kerb_frg"`
	Udp_port_lb         int    `json:"udp_port_lb"`
	Udp_bad_csum        int    `json:"udp_bad_csum"`
	Runt_ip_hdr         int    `json:"runt_ip_hdr"`
	Runt_tcp_udp_hdr    int    `json:"runt_tcp_udp_hdr"`
	Ipip_tnl_msmtch     int    `json:"ipip_tnl_msmtch"`
	Tcp_opt_err         int    `json:"tcp_opt_err"`
	Ipip_tnl_err        int    `json:"ipip_tnl_err"`
	Vxlan_err           int    `json:"vxlan_err"`
	Nvgre_err           int    `json:"nvgre_err"`
	Gre_pptp_err        int    `json:"gre_pptp_err"`
	Ipv6_eh_hbh         int    `json:"ipv6_eh_hbh"`
	Ipv6_eh_dest        int    `json:"ipv6_eh_dest"`
	Ipv6_eh_routing     int    `json:"ipv6_eh_routing"`
	Ipv6_eh_frag        int    `json:"ipv6_eh_frag"`
	Ipv6_eh_ah          int    `json:"ipv6_eh_ah"`
	Ipv6_eh_esp         int    `json:"ipv6_eh_esp"`
	Ipv6_eh_mobility    int    `json:"ipv6_eh_mobility"`
	Ipv6_eh_none        int    `json:"ipv6_eh_none"`
	Ipv6_eh_other       int    `json:"ipv6_eh_other"`
	Ipv6_eh_malformed   int    `json:"ipv6_eh_malformed"`
	Tcp_udp_zero_port   int    `json:"tcp_udp_zero_port"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobal3337 struct {
	Uuid             string                                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc3338  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate3339 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsInc3338 struct {
	Misses                int    `json:"misses"`
	OpenSocketFailed      int    `json:"open-socket-failed"`
	ConnectFailed         int    `json:"connect-failed"`
	CreateTimerFailed     int    `json:"create-timer-failed"`
	GetSocketOptionFailed int    `json:"get-socket-option-failed"`
	AflexAuthzFail        int    `json:"aflex-authz-fail"`
	AuthnFailure          int    `json:"authn-failure"`
	AuthzFailure          int    `json:"authz-failure"`
	DnsResolveFailed      int    `json:"dns-resolve-failed"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthenticationGlobalTriggerStatsRate3339 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	Misses                int    `json:"misses"`
	OpenSocketFailed      int    `json:"open-socket-failed"`
	ConnectFailed         int    `json:"connect-failed"`
	CreateTimerFailed     int    `json:"create-timer-failed"`
	GetSocketOptionFailed int    `json:"get-socket-option-failed"`
	AflexAuthzFail        int    `json:"aflex-authz-fail"`
	AuthnFailure          int    `json:"authn-failure"`
	AuthzFailure          int    `json:"authz-failure"`
	DnsResolveFailed      int    `json:"dns-resolve-failed"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdns3340 struct {
	Uuid             string                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsInc3341  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsRate3342 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsInc3341 struct {
	RequestDropped  int    `json:"request-dropped"`
	ResponseFailure int    `json:"response-failure"`
	ResponseError   int    `json:"response-error"`
	ResponseTimeout int    `json:"response-timeout"`
	Uuid            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamRdnsTriggerStatsRate3342 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdap3343 struct {
	Uuid             string                                                                                                          `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc3344  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate3345 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsInc3344 struct {
	AdminBindFailure    int    `json:"admin-bind-failure"`
	BindFailure         int    `json:"bind-failure"`
	SearchFailure       int    `json:"search-failure"`
	AuthorizeFailure    int    `json:"authorize-failure"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	SslSessionFailure   int    `json:"ssl-session-failure"`
	PwChangeFailure     int    `json:"pw-change-failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerLdapTriggerStatsRate3345 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	AdminBindFailure    int    `json:"admin-bind-failure"`
	BindFailure         int    `json:"bind-failure"`
	SearchFailure       int    `json:"search-failure"`
	AuthorizeFailure    int    `json:"authorize-failure"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	SslSessionFailure   int    `json:"ssl-session-failure"`
	PwChangeFailure     int    `json:"pw-change-failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcsp3346 struct {
	Uuid             string                                                                                                          `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc3347  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate3348 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsInc3347 struct {
	StaplingRequestDropped  int    `json:"stapling-request-dropped"`
	StaplingResponseFailure int    `json:"stapling-response-failure"`
	StaplingResponseError   int    `json:"stapling-response-error"`
	StaplingResponseTimeout int    `json:"stapling-response-timeout"`
	RequestDropped          int    `json:"request-dropped"`
	ResponseFailure         int    `json:"response-failure"`
	ResponseError           int    `json:"response-error"`
	ResponseTimeout         int    `json:"response-timeout"`
	JobStartError           int    `json:"job-start-error"`
	PollingControlError     int    `json:"polling-control-error"`
	Uuid                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerOcspTriggerStatsRate3348 struct {
	ThresholdExceededBy     int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                int    `json:"duration" dval:"60"`
	StaplingRequestDropped  int    `json:"stapling-request-dropped"`
	StaplingResponseFailure int    `json:"stapling-response-failure"`
	StaplingResponseError   int    `json:"stapling-response-error"`
	StaplingResponseTimeout int    `json:"stapling-response-timeout"`
	RequestDropped          int    `json:"request-dropped"`
	ResponseFailure         int    `json:"response-failure"`
	ResponseError           int    `json:"response-error"`
	ResponseTimeout         int    `json:"response-timeout"`
	JobStartError           int    `json:"job-start-error"`
	PollingControlError     int    `json:"polling-control-error"`
	Uuid                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadius3349 struct {
	Uuid             string                                                                                                            `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc3350  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate3351 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsInc3350 struct {
	Authen_failure      int    `json:"authen_failure"`
	Authorize_failure   int    `json:"authorize_failure"`
	Timeout_error       int    `json:"timeout_error"`
	Other_error         int    `json:"other_error"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	AccountingFailure   int    `json:"accounting-failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerRadiusTriggerStatsRate3351 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Authen_failure      int    `json:"authen_failure"`
	Authorize_failure   int    `json:"authorize_failure"`
	Timeout_error       int    `json:"timeout_error"`
	Other_error         int    `json:"other_error"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	AccountingFailure   int    `json:"accounting-failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWin3352 struct {
	Uuid             string                                                                                                         `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc3353  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate3354 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsInc3353 struct {
	KerberosTimeoutError             int    `json:"kerberos-timeout-error"`
	KerberosOtherError               int    `json:"kerberos-other-error"`
	KerberosRequestDropped           int    `json:"kerberos-request-dropped"`
	KerberosResponseFailure          int    `json:"kerberos-response-failure"`
	KerberosResponseError            int    `json:"kerberos-response-error"`
	KerberosResponseTimeout          int    `json:"kerberos-response-timeout"`
	KerberosJobStartError            int    `json:"kerberos-job-start-error"`
	KerberosPollingControlError      int    `json:"kerberos-polling-control-error"`
	KerberosPwExpiry                 int    `json:"kerberos-pw-expiry"`
	KerberosPwChangeFailure          int    `json:"kerberos-pw-change-failure"`
	KerberosValidateKdcFailure       int    `json:"kerberos-validate-kdc-failure"`
	KerberosGenerateKdcKeytabFailure int    `json:"kerberos-generate-kdc-keytab-failure"`
	KerberosDeleteKdcKeytabFailure   int    `json:"kerberos-delete-kdc-keytab-failure"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthServerWinTriggerStatsRate3354 struct {
	ThresholdExceededBy              int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                         int    `json:"duration" dval:"60"`
	KerberosTimeoutError             int    `json:"kerberos-timeout-error"`
	KerberosOtherError               int    `json:"kerberos-other-error"`
	KerberosRequestDropped           int    `json:"kerberos-request-dropped"`
	KerberosResponseFailure          int    `json:"kerberos-response-failure"`
	KerberosResponseError            int    `json:"kerberos-response-error"`
	KerberosResponseTimeout          int    `json:"kerberos-response-timeout"`
	KerberosJobStartError            int    `json:"kerberos-job-start-error"`
	KerberosPollingControlError      int    `json:"kerberos-polling-control-error"`
	KerberosPwExpiry                 int    `json:"kerberos-pw-expiry"`
	KerberosPwChangeFailure          int    `json:"kerberos-pw-change-failure"`
	KerberosValidateKdcFailure       int    `json:"kerberos-validate-kdc-failure"`
	KerberosGenerateKdcKeytabFailure int    `json:"kerberos-generate-kdc-keytab-failure"`
	KerberosDeleteKdcKeytabFailure   int    `json:"kerberos-delete-kdc-keytab-failure"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccount3355 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc3356  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate3357 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsInc3356 struct {
	RequestDropped  int    `json:"request-dropped"`
	ResponseFailure int    `json:"response-failure"`
	ResponseError   int    `json:"response-error"`
	ResponseTimeout int    `json:"response-timeout"`
	ResponseOther   int    `json:"response-other"`
	Uuid            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthAccountTriggerStatsRate3357 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	ResponseOther       int    `json:"response-other"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobal3358 struct {
	Uuid             string                                                                                                          `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc3359  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate3360 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsInc3359 struct {
	AcsAuthzFail int    `json:"acs-authz-fail"`
	AcsError     int    `json:"acs-error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthSamlGlobalTriggerStatsRate3360 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	AcsAuthzFail        int    `json:"acs-authz-fail"`
	AcsError            int    `json:"acs-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberos3361 struct {
	Uuid             string                                                                                                             `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc3362  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate3363 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsInc3362 struct {
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthRelayKerberosTriggerStatsRate3363 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptcha3364 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc3365  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate3366 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsInc3365 struct {
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JsonFail            int    `json:"json-fail"`
	AttrFail            int    `json:"attr-fail"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeAamAuthCaptchaTriggerStatsRate3366 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	RequestDropped      int    `json:"request-dropped"`
	ResponseFailure     int    `json:"response-failure"`
	ResponseError       int    `json:"response-error"`
	ResponseTimeout     int    `json:"response-timeout"`
	JsonFail            int    `json:"json-fail"`
	AttrFail            int    `json:"attr-fail"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	JobStartError       int    `json:"job-start-error"`
	PollingControlError int    `json:"polling-control-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslError3367 struct {
	Uuid             string                                                                                                    `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc3368  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate3369 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsInc3368 struct {
	AppDataInHandshake                  int    `json:"app-data-in-handshake"`
	AttemptToReuseSessInDiffContext     int    `json:"attempt-to-reuse-sess-in-diff-context"`
	BadAlertRecord                      int    `json:"bad-alert-record"`
	BadAuthenticationType               int    `json:"bad-authentication-type"`
	BadChangeCipherSpec                 int    `json:"bad-change-cipher-spec"`
	BadChecksum                         int    `json:"bad-checksum"`
	BadDataReturnedByCallback           int    `json:"bad-data-returned-by-callback"`
	BadDecompression                    int    `json:"bad-decompression"`
	BadDhGLength                        int    `json:"bad-dh-g-length"`
	BadDhPubKeyLength                   int    `json:"bad-dh-pub-key-length"`
	BadDhPLength                        int    `json:"bad-dh-p-length"`
	BadDigestLength                     int    `json:"bad-digest-length"`
	BadDsaSignature                     int    `json:"bad-dsa-signature"`
	BadHelloRequest                     int    `json:"bad-hello-request"`
	BadLength                           int    `json:"bad-length"`
	BadMacDecode                        int    `json:"bad-mac-decode"`
	BadMessageType                      int    `json:"bad-message-type"`
	BadPacketLength                     int    `json:"bad-packet-length"`
	BadProtocolVersionCounter           int    `json:"bad-protocol-version-counter"`
	BadResponseArgument                 int    `json:"bad-response-argument"`
	BadRsaDecrypt                       int    `json:"bad-rsa-decrypt"`
	BadRsaEncrypt                       int    `json:"bad-rsa-encrypt"`
	BadRsaELength                       int    `json:"bad-rsa-e-length"`
	BadRsaModulusLength                 int    `json:"bad-rsa-modulus-length"`
	BadRsaSignature                     int    `json:"bad-rsa-signature"`
	BadSignature                        int    `json:"bad-signature"`
	BadSslFiletype                      int    `json:"bad-ssl-filetype"`
	BadSslSessionIdLength               int    `json:"bad-ssl-session-id-length"`
	BadState                            int    `json:"bad-state"`
	BadWriteRetry                       int    `json:"bad-write-retry"`
	BioNotSet                           int    `json:"bio-not-set"`
	BlockCipherPadIsWrong               int    `json:"block-cipher-pad-is-wrong"`
	BnLib                               int    `json:"bn-lib"`
	CaDnLengthMismatch                  int    `json:"ca-dn-length-mismatch"`
	CaDnTooLong                         int    `json:"ca-dn-too-long"`
	CcsReceivedEarly                    int    `json:"ccs-received-early"`
	CertificateVerifyFailed             int    `json:"certificate-verify-failed"`
	CertLengthMismatch                  int    `json:"cert-length-mismatch"`
	ChallengeIsDifferent                int    `json:"challenge-is-different"`
	CipherCodeWrongLength               int    `json:"cipher-code-wrong-length"`
	CipherOrHashUnavailable             int    `json:"cipher-or-hash-unavailable"`
	CipherTableSrcError                 int    `json:"cipher-table-src-error"`
	CompressedLengthTooLong             int    `json:"compressed-length-too-long"`
	CompressionFailure                  int    `json:"compression-failure"`
	CompressionLibraryError             int    `json:"compression-library-error"`
	ConnectionIdIsDifferent             int    `json:"connection-id-is-different"`
	ConnectionTypeNotSet                int    `json:"connection-type-not-set"`
	DataBetweenCcsAndFinished           int    `json:"data-between-ccs-and-finished"`
	DataLengthTooLong                   int    `json:"data-length-too-long"`
	DecryptionFailed                    int    `json:"decryption-failed"`
	DecryptionFailedOrBadRecordMac      int    `json:"decryption-failed-or-bad-record-mac"`
	DhPublicValueLengthIsWrong          int    `json:"dh-public-value-length-is-wrong"`
	DigestCheckFailed                   int    `json:"digest-check-failed"`
	EncryptedLengthTooLong              int    `json:"encrypted-length-too-long"`
	ErrorGeneratingTmpRsaKey            int    `json:"error-generating-tmp-rsa-key"`
	ErrorInReceivedCipherList           int    `json:"error-in-received-cipher-list"`
	ExcessiveMessageSize                int    `json:"excessive-message-size"`
	ExtraDataInMessage                  int    `json:"extra-data-in-message"`
	GotAFinBeforeACcs                   int    `json:"got-a-fin-before-a-ccs"`
	HttpsProxyRequest                   int    `json:"https-proxy-request"`
	HttpRequest                         int    `json:"http-request"`
	IllegalPadding                      int    `json:"illegal-padding"`
	InappropriateFallback               int    `json:"inappropriate-fallback"`
	InvalidChallengeLength              int    `json:"invalid-challenge-length"`
	InvalidCommand                      int    `json:"invalid-command"`
	InvalidPurpose                      int    `json:"invalid-purpose"`
	InvalidStatusResponse               int    `json:"invalid-status-response"`
	InvalidTrust                        int    `json:"invalid-trust"`
	KeyArgTooLong                       int    `json:"key-arg-too-long"`
	Krb5                                int    `json:"krb5"`
	Krb5ClientCcPrincipal               int    `json:"krb5-client-cc-principal"`
	Krb5ClientGetCred                   int    `json:"krb5-client-get-cred"`
	Krb5ClientInit                      int    `json:"krb5-client-init"`
	Krb5ClientMkReq                     int    `json:"krb5-client-mk-req"`
	Krb5ServerBadTicket                 int    `json:"krb5-server-bad-ticket"`
	Krb5ServerInit                      int    `json:"krb5-server-init"`
	Krb5ServerRdReq                     int    `json:"krb5-server-rd-req"`
	Krb5ServerTktExpired                int    `json:"krb5-server-tkt-expired"`
	Krb5ServerTktNotYetValid            int    `json:"krb5-server-tkt-not-yet-valid"`
	Krb5ServerTktSkew                   int    `json:"krb5-server-tkt-skew"`
	LengthMismatch                      int    `json:"length-mismatch"`
	LengthTooShort                      int    `json:"length-too-short"`
	LibraryBug                          int    `json:"library-bug"`
	LibraryHasNoCiphers                 int    `json:"library-has-no-ciphers"`
	MastKeyTooLong                      int    `json:"mast-key-too-long"`
	MessageTooLong                      int    `json:"message-too-long"`
	MissingDhDsaCert                    int    `json:"missing-dh-dsa-cert"`
	MissingDhKey                        int    `json:"missing-dh-key"`
	MissingDhRsaCert                    int    `json:"missing-dh-rsa-cert"`
	MissingDsaSigningCert               int    `json:"missing-dsa-signing-cert"`
	MissingExportTmpDhKey               int    `json:"missing-export-tmp-dh-key"`
	MissingExportTmpRsaKey              int    `json:"missing-export-tmp-rsa-key"`
	MissingRsaCertificate               int    `json:"missing-rsa-certificate"`
	MissingRsaEncryptingCert            int    `json:"missing-rsa-encrypting-cert"`
	MissingRsaSigningCert               int    `json:"missing-rsa-signing-cert"`
	MissingTmpDhKey                     int    `json:"missing-tmp-dh-key"`
	MissingTmpRsaKey                    int    `json:"missing-tmp-rsa-key"`
	MissingTmpRsaPkey                   int    `json:"missing-tmp-rsa-pkey"`
	MissingVerifyMessage                int    `json:"missing-verify-message"`
	NonSslv2InitialPacket               int    `json:"non-sslv2-initial-packet"`
	NoCertificatesReturned              int    `json:"no-certificates-returned"`
	NoCertificateAssigned               int    `json:"no-certificate-assigned"`
	NoCertificateReturned               int    `json:"no-certificate-returned"`
	NoCertificateSet                    int    `json:"no-certificate-set"`
	NoCertificateSpecified              int    `json:"no-certificate-specified"`
	NoCiphersAvailable                  int    `json:"no-ciphers-available"`
	NoCiphersPassed                     int    `json:"no-ciphers-passed"`
	NoCiphersSpecified                  int    `json:"no-ciphers-specified"`
	NoCipherList                        int    `json:"no-cipher-list"`
	NoCipherMatch                       int    `json:"no-cipher-match"`
	NoClientCertReceived                int    `json:"no-client-cert-received"`
	NoCompressionSpecified              int    `json:"no-compression-specified"`
	NoMethodSpecified                   int    `json:"no-method-specified"`
	NoPrivatekey                        int    `json:"no-privatekey"`
	NoPrivateKeyAssigned                int    `json:"no-private-key-assigned"`
	NoProtocolsAvailable                int    `json:"no-protocols-available"`
	NoPublickey                         int    `json:"no-publickey"`
	NoSharedCipher                      int    `json:"no-shared-cipher"`
	NoVerifyCallback                    int    `json:"no-verify-callback"`
	NullSslCtx                          int    `json:"null-ssl-ctx"`
	NullSslMethodPassed                 int    `json:"null-ssl-method-passed"`
	OldSessionCipherNotReturned         int    `json:"old-session-cipher-not-returned"`
	PacketLengthTooLong                 int    `json:"packet-length-too-long"`
	PathTooLong                         int    `json:"path-too-long"`
	PeerDidNotReturnACertificate        int    `json:"peer-did-not-return-a-certificate"`
	PeerError                           int    `json:"peer-error"`
	PeerErrorCertificate                int    `json:"peer-error-certificate"`
	PeerErrorNoCertificate              int    `json:"peer-error-no-certificate"`
	PeerErrorNoCipher                   int    `json:"peer-error-no-cipher"`
	PeerErrorUnsupportedCertificateType int    `json:"peer-error-unsupported-certificate-type"`
	PreMacLengthTooLong                 int    `json:"pre-mac-length-too-long"`
	ProblemsMappingCipherFunctions      int    `json:"problems-mapping-cipher-functions"`
	ProtocolIsShutdown                  int    `json:"protocol-is-shutdown"`
	PublicKeyEncryptError               int    `json:"public-key-encrypt-error"`
	PublicKeyIsNotRsa                   int    `json:"public-key-is-not-rsa"`
	PublicKeyNotRsa                     int    `json:"public-key-not-rsa"`
	ReadBioNotSet                       int    `json:"read-bio-not-set"`
	ReadWrongPacketType                 int    `json:"read-wrong-packet-type"`
	RecordLengthMismatch                int    `json:"record-length-mismatch"`
	RecordTooLarge                      int    `json:"record-too-large"`
	RecordTooSmall                      int    `json:"record-too-small"`
	RequiredCipherMissing               int    `json:"required-cipher-missing"`
	ReuseCertLengthNotZero              int    `json:"reuse-cert-length-not-zero"`
	ReuseCertTypeNotZero                int    `json:"reuse-cert-type-not-zero"`
	ReuseCipherListNotZero              int    `json:"reuse-cipher-list-not-zero"`
	ScsvReceivedWhenRenegotiating       int    `json:"scsv-received-when-renegotiating"`
	SessionIdContextUninitialized       int    `json:"session-id-context-uninitialized"`
	ShortRead                           int    `json:"short-read"`
	SignatureForNonSigningCertificate   int    `json:"signature-for-non-signing-certificate"`
	Ssl23DoingSessionIdReuse            int    `json:"ssl23-doing-session-id-reuse"`
	Ssl2ConnectionIdTooLong             int    `json:"ssl2-connection-id-too-long"`
	Ssl3SessionIdTooLong                int    `json:"ssl3-session-id-too-long"`
	Ssl3SessionIdTooShort               int    `json:"ssl3-session-id-too-short"`
	Sslv3AlertBadCertificate            int    `json:"sslv3-alert-bad-certificate"`
	Sslv3AlertBadRecordMac              int    `json:"sslv3-alert-bad-record-mac"`
	Sslv3AlertCertificateExpired        int    `json:"sslv3-alert-certificate-expired"`
	Sslv3AlertCertificateRevoked        int    `json:"sslv3-alert-certificate-revoked"`
	Sslv3AlertCertificateUnknown        int    `json:"sslv3-alert-certificate-unknown"`
	Sslv3AlertDecompressionFailure      int    `json:"sslv3-alert-decompression-failure"`
	Sslv3AlertHandshakeFailure          int    `json:"sslv3-alert-handshake-failure"`
	Sslv3AlertIllegalParameter          int    `json:"sslv3-alert-illegal-parameter"`
	Sslv3AlertNoCertificate             int    `json:"sslv3-alert-no-certificate"`
	Sslv3AlertPeerErrorCert             int    `json:"sslv3-alert-peer-error-cert"`
	Sslv3AlertPeerErrorNoCert           int    `json:"sslv3-alert-peer-error-no-cert"`
	Sslv3AlertPeerErrorNoCipher         int    `json:"sslv3-alert-peer-error-no-cipher"`
	Sslv3AlertPeerErrorUnsuppCertType   int    `json:"sslv3-alert-peer-error-unsupp-cert-type"`
	Sslv3AlertUnexpectedMsg             int    `json:"sslv3-alert-unexpected-msg"`
	Sslv3AlertUnknownRemoteErrType      int    `json:"sslv3-alert-unknown-remote-err-type"`
	Sslv3AlertUnspportedCert            int    `json:"sslv3-alert-unspported-cert"`
	SslCtxHasNoDefaultSslVersion        int    `json:"ssl-ctx-has-no-default-ssl-version"`
	SslHandshakeFailure                 int    `json:"ssl-handshake-failure"`
	SslLibraryHasNoCiphers              int    `json:"ssl-library-has-no-ciphers"`
	SslSessionIdCallbackFailed          int    `json:"ssl-session-id-callback-failed"`
	SslSessionIdConflict                int    `json:"ssl-session-id-conflict"`
	SslSessionIdContextTooLong          int    `json:"ssl-session-id-context-too-long"`
	SslSessionIdHasBadLength            int    `json:"ssl-session-id-has-bad-length"`
	SslSessionIdIsDifferent             int    `json:"ssl-session-id-is-different"`
	Tlsv1AlertAccessDenied              int    `json:"tlsv1-alert-access-denied"`
	Tlsv1AlertDecodeError               int    `json:"tlsv1-alert-decode-error"`
	Tlsv1AlertDecryptionFailed          int    `json:"tlsv1-alert-decryption-failed"`
	Tlsv1AlertDecryptError              int    `json:"tlsv1-alert-decrypt-error"`
	Tlsv1AlertExportRestriction         int    `json:"tlsv1-alert-export-restriction"`
	Tlsv1AlertInsufficientSecurity      int    `json:"tlsv1-alert-insufficient-security"`
	Tlsv1AlertInternalError             int    `json:"tlsv1-alert-internal-error"`
	Tlsv1AlertNoRenegotiation           int    `json:"tlsv1-alert-no-renegotiation"`
	Tlsv1AlertProtocolVersion           int    `json:"tlsv1-alert-protocol-version"`
	Tlsv1AlertRecordOverflow            int    `json:"tlsv1-alert-record-overflow"`
	Tlsv1AlertUnknownCa                 int    `json:"tlsv1-alert-unknown-ca"`
	Tlsv1AlertUserCancelled             int    `json:"tlsv1-alert-user-cancelled"`
	TlsClientCertReqWithAnonCipher      int    `json:"tls-client-cert-req-with-anon-cipher"`
	TlsPeerDidNotRespondWithCertList    int    `json:"tls-peer-did-not-respond-with-cert-list"`
	TlsRsaEncryptedValueLengthIsWrong   int    `json:"tls-rsa-encrypted-value-length-is-wrong"`
	TriedToUseUnsupportedCipher         int    `json:"tried-to-use-unsupported-cipher"`
	UnableToDecodeDhCerts               int    `json:"unable-to-decode-dh-certs"`
	UnableToExtractPublicKey            int    `json:"unable-to-extract-public-key"`
	UnableToFindDhParameters            int    `json:"unable-to-find-dh-parameters"`
	UnableToFindPublicKeyParameters     int    `json:"unable-to-find-public-key-parameters"`
	UnableToFindSslMethod               int    `json:"unable-to-find-ssl-method"`
	UnableToLoadSsl2Md5Routines         int    `json:"unable-to-load-ssl2-md5-routines"`
	UnableToLoadSsl3Md5Routines         int    `json:"unable-to-load-ssl3-md5-routines"`
	UnableToLoadSsl3Sha1Routines        int    `json:"unable-to-load-ssl3-sha1-routines"`
	UnexpectedMessage                   int    `json:"unexpected-message"`
	UnexpectedRecord                    int    `json:"unexpected-record"`
	Uninitialized                       int    `json:"uninitialized"`
	UnknownAlertType                    int    `json:"unknown-alert-type"`
	UnknownCertificateType              int    `json:"unknown-certificate-type"`
	UnknownCipherReturned               int    `json:"unknown-cipher-returned"`
	UnknownCipherType                   int    `json:"unknown-cipher-type"`
	UnknownKeyExchangeType              int    `json:"unknown-key-exchange-type"`
	UnknownPkeyType                     int    `json:"unknown-pkey-type"`
	UnknownProtocol                     int    `json:"unknown-protocol"`
	UnknownRemoteErrorType              int    `json:"unknown-remote-error-type"`
	UnknownSslVersion                   int    `json:"unknown-ssl-version"`
	UnknownState                        int    `json:"unknown-state"`
	UnsupportedCipher                   int    `json:"unsupported-cipher"`
	UnsupportedCompressionAlgorithm     int    `json:"unsupported-compression-algorithm"`
	UnsupportedOption                   int    `json:"unsupported-option"`
	UnsupportedProtocol                 int    `json:"unsupported-protocol"`
	UnsupportedSslVersion               int    `json:"unsupported-ssl-version"`
	UnsupportedStatusType               int    `json:"unsupported-status-type"`
	WriteBioNotSet                      int    `json:"write-bio-not-set"`
	WrongCipherReturned                 int    `json:"wrong-cipher-returned"`
	WrongMessageType                    int    `json:"wrong-message-type"`
	WrongCounterOfKeyBits               int    `json:"wrong-counter-of-key-bits"`
	WrongSignatureLength                int    `json:"wrong-signature-length"`
	WrongSignatureSize                  int    `json:"wrong-signature-size"`
	WrongSslVersion                     int    `json:"wrong-ssl-version"`
	WrongVersionCounter                 int    `json:"wrong-version-counter"`
	X509Lib                             int    `json:"x509-lib"`
	X509VerificationSetupProblems       int    `json:"x509-verification-setup-problems"`
	ClienthelloTlsext                   int    `json:"clienthello-tlsext"`
	ParseTlsext                         int    `json:"parse-tlsext"`
	ServerhelloTlsext                   int    `json:"serverhello-tlsext"`
	Ssl3ExtInvalidServername            int    `json:"ssl3-ext-invalid-servername"`
	Ssl3ExtInvalidServernameType        int    `json:"ssl3-ext-invalid-servername-type"`
	MultipleSgcRestarts                 int    `json:"multiple-sgc-restarts"`
	TlsInvalidEcpointformatList         int    `json:"tls-invalid-ecpointformat-list"`
	BadEccCert                          int    `json:"bad-ecc-cert"`
	BadEcdsaSig                         int    `json:"bad-ecdsa-sig"`
	BadEcpoint                          int    `json:"bad-ecpoint"`
	CookieMismatch                      int    `json:"cookie-mismatch"`
	UnsupportedEllipticCurve            int    `json:"unsupported-elliptic-curve"`
	NoRequiredDigest                    int    `json:"no-required-digest"`
	UnsupportedDigestType               int    `json:"unsupported-digest-type"`
	BadHandshakeLength                  int    `json:"bad-handshake-length"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslErrorTriggerStatsRate3369 struct {
	ThresholdExceededBy                 int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                            int    `json:"duration" dval:"60"`
	AppDataInHandshake                  int    `json:"app-data-in-handshake"`
	AttemptToReuseSessInDiffContext     int    `json:"attempt-to-reuse-sess-in-diff-context"`
	BadAlertRecord                      int    `json:"bad-alert-record"`
	BadAuthenticationType               int    `json:"bad-authentication-type"`
	BadChangeCipherSpec                 int    `json:"bad-change-cipher-spec"`
	BadChecksum                         int    `json:"bad-checksum"`
	BadDataReturnedByCallback           int    `json:"bad-data-returned-by-callback"`
	BadDecompression                    int    `json:"bad-decompression"`
	BadDhGLength                        int    `json:"bad-dh-g-length"`
	BadDhPubKeyLength                   int    `json:"bad-dh-pub-key-length"`
	BadDhPLength                        int    `json:"bad-dh-p-length"`
	BadDigestLength                     int    `json:"bad-digest-length"`
	BadDsaSignature                     int    `json:"bad-dsa-signature"`
	BadHelloRequest                     int    `json:"bad-hello-request"`
	BadLength                           int    `json:"bad-length"`
	BadMacDecode                        int    `json:"bad-mac-decode"`
	BadMessageType                      int    `json:"bad-message-type"`
	BadPacketLength                     int    `json:"bad-packet-length"`
	BadProtocolVersionCounter           int    `json:"bad-protocol-version-counter"`
	BadResponseArgument                 int    `json:"bad-response-argument"`
	BadRsaDecrypt                       int    `json:"bad-rsa-decrypt"`
	BadRsaEncrypt                       int    `json:"bad-rsa-encrypt"`
	BadRsaELength                       int    `json:"bad-rsa-e-length"`
	BadRsaModulusLength                 int    `json:"bad-rsa-modulus-length"`
	BadRsaSignature                     int    `json:"bad-rsa-signature"`
	BadSignature                        int    `json:"bad-signature"`
	BadSslFiletype                      int    `json:"bad-ssl-filetype"`
	BadSslSessionIdLength               int    `json:"bad-ssl-session-id-length"`
	BadState                            int    `json:"bad-state"`
	BadWriteRetry                       int    `json:"bad-write-retry"`
	BioNotSet                           int    `json:"bio-not-set"`
	BlockCipherPadIsWrong               int    `json:"block-cipher-pad-is-wrong"`
	BnLib                               int    `json:"bn-lib"`
	CaDnLengthMismatch                  int    `json:"ca-dn-length-mismatch"`
	CaDnTooLong                         int    `json:"ca-dn-too-long"`
	CcsReceivedEarly                    int    `json:"ccs-received-early"`
	CertificateVerifyFailed             int    `json:"certificate-verify-failed"`
	CertLengthMismatch                  int    `json:"cert-length-mismatch"`
	ChallengeIsDifferent                int    `json:"challenge-is-different"`
	CipherCodeWrongLength               int    `json:"cipher-code-wrong-length"`
	CipherOrHashUnavailable             int    `json:"cipher-or-hash-unavailable"`
	CipherTableSrcError                 int    `json:"cipher-table-src-error"`
	CompressedLengthTooLong             int    `json:"compressed-length-too-long"`
	CompressionFailure                  int    `json:"compression-failure"`
	CompressionLibraryError             int    `json:"compression-library-error"`
	ConnectionIdIsDifferent             int    `json:"connection-id-is-different"`
	ConnectionTypeNotSet                int    `json:"connection-type-not-set"`
	DataBetweenCcsAndFinished           int    `json:"data-between-ccs-and-finished"`
	DataLengthTooLong                   int    `json:"data-length-too-long"`
	DecryptionFailed                    int    `json:"decryption-failed"`
	DecryptionFailedOrBadRecordMac      int    `json:"decryption-failed-or-bad-record-mac"`
	DhPublicValueLengthIsWrong          int    `json:"dh-public-value-length-is-wrong"`
	DigestCheckFailed                   int    `json:"digest-check-failed"`
	EncryptedLengthTooLong              int    `json:"encrypted-length-too-long"`
	ErrorGeneratingTmpRsaKey            int    `json:"error-generating-tmp-rsa-key"`
	ErrorInReceivedCipherList           int    `json:"error-in-received-cipher-list"`
	ExcessiveMessageSize                int    `json:"excessive-message-size"`
	ExtraDataInMessage                  int    `json:"extra-data-in-message"`
	GotAFinBeforeACcs                   int    `json:"got-a-fin-before-a-ccs"`
	HttpsProxyRequest                   int    `json:"https-proxy-request"`
	HttpRequest                         int    `json:"http-request"`
	IllegalPadding                      int    `json:"illegal-padding"`
	InappropriateFallback               int    `json:"inappropriate-fallback"`
	InvalidChallengeLength              int    `json:"invalid-challenge-length"`
	InvalidCommand                      int    `json:"invalid-command"`
	InvalidPurpose                      int    `json:"invalid-purpose"`
	InvalidStatusResponse               int    `json:"invalid-status-response"`
	InvalidTrust                        int    `json:"invalid-trust"`
	KeyArgTooLong                       int    `json:"key-arg-too-long"`
	Krb5                                int    `json:"krb5"`
	Krb5ClientCcPrincipal               int    `json:"krb5-client-cc-principal"`
	Krb5ClientGetCred                   int    `json:"krb5-client-get-cred"`
	Krb5ClientInit                      int    `json:"krb5-client-init"`
	Krb5ClientMkReq                     int    `json:"krb5-client-mk-req"`
	Krb5ServerBadTicket                 int    `json:"krb5-server-bad-ticket"`
	Krb5ServerInit                      int    `json:"krb5-server-init"`
	Krb5ServerRdReq                     int    `json:"krb5-server-rd-req"`
	Krb5ServerTktExpired                int    `json:"krb5-server-tkt-expired"`
	Krb5ServerTktNotYetValid            int    `json:"krb5-server-tkt-not-yet-valid"`
	Krb5ServerTktSkew                   int    `json:"krb5-server-tkt-skew"`
	LengthMismatch                      int    `json:"length-mismatch"`
	LengthTooShort                      int    `json:"length-too-short"`
	LibraryBug                          int    `json:"library-bug"`
	LibraryHasNoCiphers                 int    `json:"library-has-no-ciphers"`
	MastKeyTooLong                      int    `json:"mast-key-too-long"`
	MessageTooLong                      int    `json:"message-too-long"`
	MissingDhDsaCert                    int    `json:"missing-dh-dsa-cert"`
	MissingDhKey                        int    `json:"missing-dh-key"`
	MissingDhRsaCert                    int    `json:"missing-dh-rsa-cert"`
	MissingDsaSigningCert               int    `json:"missing-dsa-signing-cert"`
	MissingExportTmpDhKey               int    `json:"missing-export-tmp-dh-key"`
	MissingExportTmpRsaKey              int    `json:"missing-export-tmp-rsa-key"`
	MissingRsaCertificate               int    `json:"missing-rsa-certificate"`
	MissingRsaEncryptingCert            int    `json:"missing-rsa-encrypting-cert"`
	MissingRsaSigningCert               int    `json:"missing-rsa-signing-cert"`
	MissingTmpDhKey                     int    `json:"missing-tmp-dh-key"`
	MissingTmpRsaKey                    int    `json:"missing-tmp-rsa-key"`
	MissingTmpRsaPkey                   int    `json:"missing-tmp-rsa-pkey"`
	MissingVerifyMessage                int    `json:"missing-verify-message"`
	NonSslv2InitialPacket               int    `json:"non-sslv2-initial-packet"`
	NoCertificatesReturned              int    `json:"no-certificates-returned"`
	NoCertificateAssigned               int    `json:"no-certificate-assigned"`
	NoCertificateReturned               int    `json:"no-certificate-returned"`
	NoCertificateSet                    int    `json:"no-certificate-set"`
	NoCertificateSpecified              int    `json:"no-certificate-specified"`
	NoCiphersAvailable                  int    `json:"no-ciphers-available"`
	NoCiphersPassed                     int    `json:"no-ciphers-passed"`
	NoCiphersSpecified                  int    `json:"no-ciphers-specified"`
	NoCipherList                        int    `json:"no-cipher-list"`
	NoCipherMatch                       int    `json:"no-cipher-match"`
	NoClientCertReceived                int    `json:"no-client-cert-received"`
	NoCompressionSpecified              int    `json:"no-compression-specified"`
	NoMethodSpecified                   int    `json:"no-method-specified"`
	NoPrivatekey                        int    `json:"no-privatekey"`
	NoPrivateKeyAssigned                int    `json:"no-private-key-assigned"`
	NoProtocolsAvailable                int    `json:"no-protocols-available"`
	NoPublickey                         int    `json:"no-publickey"`
	NoSharedCipher                      int    `json:"no-shared-cipher"`
	NoVerifyCallback                    int    `json:"no-verify-callback"`
	NullSslCtx                          int    `json:"null-ssl-ctx"`
	NullSslMethodPassed                 int    `json:"null-ssl-method-passed"`
	OldSessionCipherNotReturned         int    `json:"old-session-cipher-not-returned"`
	PacketLengthTooLong                 int    `json:"packet-length-too-long"`
	PathTooLong                         int    `json:"path-too-long"`
	PeerDidNotReturnACertificate        int    `json:"peer-did-not-return-a-certificate"`
	PeerError                           int    `json:"peer-error"`
	PeerErrorCertificate                int    `json:"peer-error-certificate"`
	PeerErrorNoCertificate              int    `json:"peer-error-no-certificate"`
	PeerErrorNoCipher                   int    `json:"peer-error-no-cipher"`
	PeerErrorUnsupportedCertificateType int    `json:"peer-error-unsupported-certificate-type"`
	PreMacLengthTooLong                 int    `json:"pre-mac-length-too-long"`
	ProblemsMappingCipherFunctions      int    `json:"problems-mapping-cipher-functions"`
	ProtocolIsShutdown                  int    `json:"protocol-is-shutdown"`
	PublicKeyEncryptError               int    `json:"public-key-encrypt-error"`
	PublicKeyIsNotRsa                   int    `json:"public-key-is-not-rsa"`
	PublicKeyNotRsa                     int    `json:"public-key-not-rsa"`
	ReadBioNotSet                       int    `json:"read-bio-not-set"`
	ReadWrongPacketType                 int    `json:"read-wrong-packet-type"`
	RecordLengthMismatch                int    `json:"record-length-mismatch"`
	RecordTooLarge                      int    `json:"record-too-large"`
	RecordTooSmall                      int    `json:"record-too-small"`
	RequiredCipherMissing               int    `json:"required-cipher-missing"`
	ReuseCertLengthNotZero              int    `json:"reuse-cert-length-not-zero"`
	ReuseCertTypeNotZero                int    `json:"reuse-cert-type-not-zero"`
	ReuseCipherListNotZero              int    `json:"reuse-cipher-list-not-zero"`
	ScsvReceivedWhenRenegotiating       int    `json:"scsv-received-when-renegotiating"`
	SessionIdContextUninitialized       int    `json:"session-id-context-uninitialized"`
	ShortRead                           int    `json:"short-read"`
	SignatureForNonSigningCertificate   int    `json:"signature-for-non-signing-certificate"`
	Ssl23DoingSessionIdReuse            int    `json:"ssl23-doing-session-id-reuse"`
	Ssl2ConnectionIdTooLong             int    `json:"ssl2-connection-id-too-long"`
	Ssl3SessionIdTooLong                int    `json:"ssl3-session-id-too-long"`
	Ssl3SessionIdTooShort               int    `json:"ssl3-session-id-too-short"`
	Sslv3AlertBadCertificate            int    `json:"sslv3-alert-bad-certificate"`
	Sslv3AlertBadRecordMac              int    `json:"sslv3-alert-bad-record-mac"`
	Sslv3AlertCertificateExpired        int    `json:"sslv3-alert-certificate-expired"`
	Sslv3AlertCertificateRevoked        int    `json:"sslv3-alert-certificate-revoked"`
	Sslv3AlertCertificateUnknown        int    `json:"sslv3-alert-certificate-unknown"`
	Sslv3AlertDecompressionFailure      int    `json:"sslv3-alert-decompression-failure"`
	Sslv3AlertHandshakeFailure          int    `json:"sslv3-alert-handshake-failure"`
	Sslv3AlertIllegalParameter          int    `json:"sslv3-alert-illegal-parameter"`
	Sslv3AlertNoCertificate             int    `json:"sslv3-alert-no-certificate"`
	Sslv3AlertPeerErrorCert             int    `json:"sslv3-alert-peer-error-cert"`
	Sslv3AlertPeerErrorNoCert           int    `json:"sslv3-alert-peer-error-no-cert"`
	Sslv3AlertPeerErrorNoCipher         int    `json:"sslv3-alert-peer-error-no-cipher"`
	Sslv3AlertPeerErrorUnsuppCertType   int    `json:"sslv3-alert-peer-error-unsupp-cert-type"`
	Sslv3AlertUnexpectedMsg             int    `json:"sslv3-alert-unexpected-msg"`
	Sslv3AlertUnknownRemoteErrType      int    `json:"sslv3-alert-unknown-remote-err-type"`
	Sslv3AlertUnspportedCert            int    `json:"sslv3-alert-unspported-cert"`
	SslCtxHasNoDefaultSslVersion        int    `json:"ssl-ctx-has-no-default-ssl-version"`
	SslHandshakeFailure                 int    `json:"ssl-handshake-failure"`
	SslLibraryHasNoCiphers              int    `json:"ssl-library-has-no-ciphers"`
	SslSessionIdCallbackFailed          int    `json:"ssl-session-id-callback-failed"`
	SslSessionIdConflict                int    `json:"ssl-session-id-conflict"`
	SslSessionIdContextTooLong          int    `json:"ssl-session-id-context-too-long"`
	SslSessionIdHasBadLength            int    `json:"ssl-session-id-has-bad-length"`
	SslSessionIdIsDifferent             int    `json:"ssl-session-id-is-different"`
	Tlsv1AlertAccessDenied              int    `json:"tlsv1-alert-access-denied"`
	Tlsv1AlertDecodeError               int    `json:"tlsv1-alert-decode-error"`
	Tlsv1AlertDecryptionFailed          int    `json:"tlsv1-alert-decryption-failed"`
	Tlsv1AlertDecryptError              int    `json:"tlsv1-alert-decrypt-error"`
	Tlsv1AlertExportRestriction         int    `json:"tlsv1-alert-export-restriction"`
	Tlsv1AlertInsufficientSecurity      int    `json:"tlsv1-alert-insufficient-security"`
	Tlsv1AlertInternalError             int    `json:"tlsv1-alert-internal-error"`
	Tlsv1AlertNoRenegotiation           int    `json:"tlsv1-alert-no-renegotiation"`
	Tlsv1AlertProtocolVersion           int    `json:"tlsv1-alert-protocol-version"`
	Tlsv1AlertRecordOverflow            int    `json:"tlsv1-alert-record-overflow"`
	Tlsv1AlertUnknownCa                 int    `json:"tlsv1-alert-unknown-ca"`
	Tlsv1AlertUserCancelled             int    `json:"tlsv1-alert-user-cancelled"`
	TlsClientCertReqWithAnonCipher      int    `json:"tls-client-cert-req-with-anon-cipher"`
	TlsPeerDidNotRespondWithCertList    int    `json:"tls-peer-did-not-respond-with-cert-list"`
	TlsRsaEncryptedValueLengthIsWrong   int    `json:"tls-rsa-encrypted-value-length-is-wrong"`
	TriedToUseUnsupportedCipher         int    `json:"tried-to-use-unsupported-cipher"`
	UnableToDecodeDhCerts               int    `json:"unable-to-decode-dh-certs"`
	UnableToExtractPublicKey            int    `json:"unable-to-extract-public-key"`
	UnableToFindDhParameters            int    `json:"unable-to-find-dh-parameters"`
	UnableToFindPublicKeyParameters     int    `json:"unable-to-find-public-key-parameters"`
	UnableToFindSslMethod               int    `json:"unable-to-find-ssl-method"`
	UnableToLoadSsl2Md5Routines         int    `json:"unable-to-load-ssl2-md5-routines"`
	UnableToLoadSsl3Md5Routines         int    `json:"unable-to-load-ssl3-md5-routines"`
	UnableToLoadSsl3Sha1Routines        int    `json:"unable-to-load-ssl3-sha1-routines"`
	UnexpectedMessage                   int    `json:"unexpected-message"`
	UnexpectedRecord                    int    `json:"unexpected-record"`
	Uninitialized                       int    `json:"uninitialized"`
	UnknownAlertType                    int    `json:"unknown-alert-type"`
	UnknownCertificateType              int    `json:"unknown-certificate-type"`
	UnknownCipherReturned               int    `json:"unknown-cipher-returned"`
	UnknownCipherType                   int    `json:"unknown-cipher-type"`
	UnknownKeyExchangeType              int    `json:"unknown-key-exchange-type"`
	UnknownPkeyType                     int    `json:"unknown-pkey-type"`
	UnknownProtocol                     int    `json:"unknown-protocol"`
	UnknownRemoteErrorType              int    `json:"unknown-remote-error-type"`
	UnknownSslVersion                   int    `json:"unknown-ssl-version"`
	UnknownState                        int    `json:"unknown-state"`
	UnsupportedCipher                   int    `json:"unsupported-cipher"`
	UnsupportedCompressionAlgorithm     int    `json:"unsupported-compression-algorithm"`
	UnsupportedOption                   int    `json:"unsupported-option"`
	UnsupportedProtocol                 int    `json:"unsupported-protocol"`
	UnsupportedSslVersion               int    `json:"unsupported-ssl-version"`
	UnsupportedStatusType               int    `json:"unsupported-status-type"`
	WriteBioNotSet                      int    `json:"write-bio-not-set"`
	WrongCipherReturned                 int    `json:"wrong-cipher-returned"`
	WrongMessageType                    int    `json:"wrong-message-type"`
	WrongCounterOfKeyBits               int    `json:"wrong-counter-of-key-bits"`
	WrongSignatureLength                int    `json:"wrong-signature-length"`
	WrongSignatureSize                  int    `json:"wrong-signature-size"`
	WrongSslVersion                     int    `json:"wrong-ssl-version"`
	WrongVersionCounter                 int    `json:"wrong-version-counter"`
	X509Lib                             int    `json:"x509-lib"`
	X509VerificationSetupProblems       int    `json:"x509-verification-setup-problems"`
	ClienthelloTlsext                   int    `json:"clienthello-tlsext"`
	ParseTlsext                         int    `json:"parse-tlsext"`
	ServerhelloTlsext                   int    `json:"serverhello-tlsext"`
	Ssl3ExtInvalidServername            int    `json:"ssl3-ext-invalid-servername"`
	Ssl3ExtInvalidServernameType        int    `json:"ssl3-ext-invalid-servername-type"`
	MultipleSgcRestarts                 int    `json:"multiple-sgc-restarts"`
	TlsInvalidEcpointformatList         int    `json:"tls-invalid-ecpointformat-list"`
	BadEccCert                          int    `json:"bad-ecc-cert"`
	BadEcdsaSig                         int    `json:"bad-ecdsa-sig"`
	BadEcpoint                          int    `json:"bad-ecpoint"`
	CookieMismatch                      int    `json:"cookie-mismatch"`
	UnsupportedEllipticCurve            int    `json:"unsupported-elliptic-curve"`
	NoRequiredDigest                    int    `json:"no-required-digest"`
	UnsupportedDigestType               int    `json:"unsupported-digest-type"`
	BadHandshakeLength                  int    `json:"bad-handshake-length"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevoke3370 struct {
	Uuid             string                                                                                                         `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc3371  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate3372 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsInc3371 struct {
	Ocsp_chain_status_revoked    int    `json:"ocsp_chain_status_revoked"`
	Ocsp_chain_status_unknown    int    `json:"ocsp_chain_status_unknown"`
	Ocsp_connection_error        int    `json:"ocsp_connection_error"`
	Ocsp_uri_not_found           int    `json:"ocsp_uri_not_found"`
	Ocsp_uri_https               int    `json:"ocsp_uri_https"`
	Ocsp_uri_unsupported         int    `json:"ocsp_uri_unsupported"`
	Ocsp_response_status_revoked int    `json:"ocsp_response_status_revoked"`
	Ocsp_response_status_unknown int    `json:"ocsp_response_status_unknown"`
	Ocsp_cache_status_revoked    int    `json:"ocsp_cache_status_revoked"`
	Ocsp_cache_miss              int    `json:"ocsp_cache_miss"`
	Ocsp_other_error             int    `json:"ocsp_other_error"`
	Ocsp_response_no_nonce       int    `json:"ocsp_response_no_nonce"`
	Ocsp_response_nonce_error    int    `json:"ocsp_response_nonce_error"`
	Crl_connection_error         int    `json:"crl_connection_error"`
	Crl_uri_not_found            int    `json:"crl_uri_not_found"`
	Crl_uri_https                int    `json:"crl_uri_https"`
	Crl_uri_unsupported          int    `json:"crl_uri_unsupported"`
	Crl_response_status_revoked  int    `json:"crl_response_status_revoked"`
	Crl_response_status_unknown  int    `json:"crl_response_status_unknown"`
	Crl_cache_status_revoked     int    `json:"crl_cache_status_revoked"`
	Crl_other_error              int    `json:"crl_other_error"`
	Uuid                         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslCertRevokeTriggerStatsRate3372 struct {
	ThresholdExceededBy          int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                     int    `json:"duration" dval:"60"`
	Ocsp_chain_status_revoked    int    `json:"ocsp_chain_status_revoked"`
	Ocsp_chain_status_unknown    int    `json:"ocsp_chain_status_unknown"`
	Ocsp_connection_error        int    `json:"ocsp_connection_error"`
	Ocsp_uri_not_found           int    `json:"ocsp_uri_not_found"`
	Ocsp_uri_https               int    `json:"ocsp_uri_https"`
	Ocsp_uri_unsupported         int    `json:"ocsp_uri_unsupported"`
	Ocsp_response_status_revoked int    `json:"ocsp_response_status_revoked"`
	Ocsp_response_status_unknown int    `json:"ocsp_response_status_unknown"`
	Ocsp_cache_status_revoked    int    `json:"ocsp_cache_status_revoked"`
	Ocsp_cache_miss              int    `json:"ocsp_cache_miss"`
	Ocsp_other_error             int    `json:"ocsp_other_error"`
	Ocsp_response_no_nonce       int    `json:"ocsp_response_no_nonce"`
	Ocsp_response_nonce_error    int    `json:"ocsp_response_nonce_error"`
	Crl_connection_error         int    `json:"crl_connection_error"`
	Crl_uri_not_found            int    `json:"crl_uri_not_found"`
	Crl_uri_https                int    `json:"crl_uri_https"`
	Crl_uri_unsupported          int    `json:"crl_uri_unsupported"`
	Crl_response_status_revoked  int    `json:"crl_response_status_revoked"`
	Crl_response_status_unknown  int    `json:"crl_response_status_unknown"`
	Crl_cache_status_revoked     int    `json:"crl_cache_status_revoked"`
	Crl_other_error              int    `json:"crl_other_error"`
	Uuid                         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxy3373 struct {
	Uuid             string                                                                                                           `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc3374  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate3375 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsInc3374 struct {
	FailedInSslHandshakes           int    `json:"failed-in-ssl-handshakes"`
	FailedInCryptoOperations        int    `json:"failed-in-crypto-operations"`
	FailedInTcp                     int    `json:"failed-in-tcp"`
	FailedInCertificateVerification int    `json:"failed-in-certificate-verification"`
	FailedInCertificateSigning      int    `json:"failed-in-certificate-signing"`
	InvalidOcspStaplingResponse     int    `json:"invalid-ocsp-stapling-response"`
	RevokedOcspResponse             int    `json:"revoked-ocsp-response"`
	UnsupportedSslVersion           int    `json:"unsupported-ssl-version"`
	ConnectionsFailed               int    `json:"connections-failed"`
	Uuid                            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSslForwardProxyTriggerStatsRate3375 struct {
	ThresholdExceededBy             int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                        int    `json:"duration" dval:"60"`
	FailedInSslHandshakes           int    `json:"failed-in-ssl-handshakes"`
	FailedInCryptoOperations        int    `json:"failed-in-crypto-operations"`
	FailedInTcp                     int    `json:"failed-in-tcp"`
	FailedInCertificateVerification int    `json:"failed-in-certificate-verification"`
	FailedInCertificateSigning      int    `json:"failed-in-certificate-signing"`
	InvalidOcspStaplingResponse     int    `json:"invalid-ocsp-stapling-response"`
	RevokedOcspResponse             int    `json:"revoked-ocsp-response"`
	UnsupportedSslVersion           int    `json:"unsupported-ssl-version"`
	ConnectionsFailed               int    `json:"connections-failed"`
	Uuid                            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnError3376 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsInc3377  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsRate3378 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsInc3377 struct {
	Bad_opcode                             int    `json:"bad_opcode"`
	Bad_sg_write_len                       int    `json:"bad_sg_write_len"`
	Bad_len                                int    `json:"bad_len"`
	Bad_ipsec_protocol                     int    `json:"bad_ipsec_protocol"`
	Bad_ipsec_auth                         int    `json:"bad_ipsec_auth"`
	Bad_ipsec_padding                      int    `json:"bad_ipsec_padding"`
	Bad_ip_version                         int    `json:"bad_ip_version"`
	Bad_auth_type                          int    `json:"bad_auth_type"`
	Bad_encrypt_type                       int    `json:"bad_encrypt_type"`
	Bad_ipsec_spi                          int    `json:"bad_ipsec_spi"`
	Bad_checksum                           int    `json:"bad_checksum"`
	Bad_ipsec_context                      int    `json:"bad_ipsec_context"`
	Bad_ipsec_context_direction            int    `json:"bad_ipsec_context_direction"`
	Bad_ipsec_context_flag_mismatch        int    `json:"bad_ipsec_context_flag_mismatch"`
	Ipcomp_payload                         int    `json:"ipcomp_payload"`
	Bad_selector_match                     int    `json:"bad_selector_match"`
	Bad_fragment_size                      int    `json:"bad_fragment_size"`
	Bad_inline_data                        int    `json:"bad_inline_data"`
	Bad_frag_size_configuration            int    `json:"bad_frag_size_configuration"`
	Dummy_payload                          int    `json:"dummy_payload"`
	Bad_ip_payload_type                    int    `json:"bad_ip_payload_type"`
	Bad_min_frag_size_auth_sha384_512      int    `json:"bad_min_frag_size_auth_sha384_512"`
	Bad_esp_next_header                    int    `json:"bad_esp_next_header"`
	Bad_gre_header                         int    `json:"bad_gre_header"`
	Bad_gre_protocol                       int    `json:"bad_gre_protocol"`
	Ipv6_extension_headers_too_big         int    `json:"ipv6_extension_headers_too_big"`
	Ipv6_hop_by_hop_error                  int    `json:"ipv6_hop_by_hop_error"`
	Error_ipv6_decrypt_rh_segs_left_error  int    `json:"error_ipv6_decrypt_rh_segs_left_error"`
	Ipv6_rh_length_error                   int    `json:"ipv6_rh_length_error"`
	Ipv6_outbound_rh_copy_addr_error       int    `json:"ipv6_outbound_rh_copy_addr_error"`
	Error_ipv6_extension_header_bad        int    `json:"error_IPv6_extension_header_bad"`
	Bad_encrypt_type_ctr_gcm               int    `json:"bad_encrypt_type_ctr_gcm"`
	Ah_not_supported_with_gcm_gmac_sha2    int    `json:"ah_not_supported_with_gcm_gmac_sha2"`
	Tfc_padding_with_prefrag_not_supported int    `json:"tfc_padding_with_prefrag_not_supported"`
	Bad_srtp_auth_tag                      int    `json:"bad_srtp_auth_tag"`
	Bad_ipcomp_configuration               int    `json:"bad_ipcomp_configuration"`
	Dsiv_incorrect_param                   int    `json:"dsiv_incorrect_param"`
	Bad_ipsec_unknown                      int    `json:"bad_ipsec_unknown"`
	Uuid                                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeVpnErrorTriggerStatsRate3378 struct {
	ThresholdExceededBy                    int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                               int    `json:"duration" dval:"60"`
	Bad_opcode                             int    `json:"bad_opcode"`
	Bad_sg_write_len                       int    `json:"bad_sg_write_len"`
	Bad_len                                int    `json:"bad_len"`
	Bad_ipsec_protocol                     int    `json:"bad_ipsec_protocol"`
	Bad_ipsec_auth                         int    `json:"bad_ipsec_auth"`
	Bad_ipsec_padding                      int    `json:"bad_ipsec_padding"`
	Bad_ip_version                         int    `json:"bad_ip_version"`
	Bad_auth_type                          int    `json:"bad_auth_type"`
	Bad_encrypt_type                       int    `json:"bad_encrypt_type"`
	Bad_ipsec_spi                          int    `json:"bad_ipsec_spi"`
	Bad_checksum                           int    `json:"bad_checksum"`
	Bad_ipsec_context                      int    `json:"bad_ipsec_context"`
	Bad_ipsec_context_direction            int    `json:"bad_ipsec_context_direction"`
	Bad_ipsec_context_flag_mismatch        int    `json:"bad_ipsec_context_flag_mismatch"`
	Ipcomp_payload                         int    `json:"ipcomp_payload"`
	Bad_selector_match                     int    `json:"bad_selector_match"`
	Bad_fragment_size                      int    `json:"bad_fragment_size"`
	Bad_inline_data                        int    `json:"bad_inline_data"`
	Bad_frag_size_configuration            int    `json:"bad_frag_size_configuration"`
	Dummy_payload                          int    `json:"dummy_payload"`
	Bad_ip_payload_type                    int    `json:"bad_ip_payload_type"`
	Bad_min_frag_size_auth_sha384_512      int    `json:"bad_min_frag_size_auth_sha384_512"`
	Bad_esp_next_header                    int    `json:"bad_esp_next_header"`
	Bad_gre_header                         int    `json:"bad_gre_header"`
	Bad_gre_protocol                       int    `json:"bad_gre_protocol"`
	Ipv6_extension_headers_too_big         int    `json:"ipv6_extension_headers_too_big"`
	Ipv6_hop_by_hop_error                  int    `json:"ipv6_hop_by_hop_error"`
	Error_ipv6_decrypt_rh_segs_left_error  int    `json:"error_ipv6_decrypt_rh_segs_left_error"`
	Ipv6_rh_length_error                   int    `json:"ipv6_rh_length_error"`
	Ipv6_outbound_rh_copy_addr_error       int    `json:"ipv6_outbound_rh_copy_addr_error"`
	Error_ipv6_extension_header_bad        int    `json:"error_IPv6_extension_header_bad"`
	Bad_encrypt_type_ctr_gcm               int    `json:"bad_encrypt_type_ctr_gcm"`
	Ah_not_supported_with_gcm_gmac_sha2    int    `json:"ah_not_supported_with_gcm_gmac_sha2"`
	Tfc_padding_with_prefrag_not_supported int    `json:"tfc_padding_with_prefrag_not_supported"`
	Bad_srtp_auth_tag                      int    `json:"bad_srtp_auth_tag"`
	Bad_ipcomp_configuration               int    `json:"bad_ipcomp_configuration"`
	Dsiv_incorrect_param                   int    `json:"dsiv_incorrect_param"`
	Bad_ipsec_unknown                      int    `json:"bad_ipsec_unknown"`
	Uuid                                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProc3379 struct {
	Uuid             string                                                                                                      `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc3380  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate3381 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsInc3380 struct {
	L3_entry_match_drop              int    `json:"l3_entry_match_drop"`
	L3_entry_match_drop_hw           int    `json:"l3_entry_match_drop_hw"`
	L3_entry_drop_max_hw_exceeded    int    `json:"l3_entry_drop_max_hw_exceeded"`
	L4_entry_match_drop              int    `json:"l4_entry_match_drop"`
	L4_entry_match_drop_hw           int    `json:"l4_entry_match_drop_hw"`
	L4_entry_drop_max_hw_exceeded    int    `json:"l4_entry_drop_max_hw_exceeded"`
	L4_entry_list_alloc_failure      int    `json:"l4_entry_list_alloc_failure"`
	Ip_node_alloc_failure            int    `json:"ip_node_alloc_failure"`
	Ip_port_block_alloc_failure      int    `json:"ip_port_block_alloc_failure"`
	Ip_other_block_alloc_failure     int    `json:"ip_other_block_alloc_failure"`
	L3_entry_add_to_bgp_failure      int    `json:"l3_entry_add_to_bgp_failure"`
	L3_entry_remove_from_bgp_failure int    `json:"l3_entry_remove_from_bgp_failure"`
	L3_entry_add_to_hw_failure       int    `json:"l3_entry_add_to_hw_failure"`
	Syn_cookie_verification_failed   int    `json:"syn_cookie_verification_failed"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DdosProcTriggerStatsRate3381 struct {
	ThresholdExceededBy              int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                         int    `json:"duration" dval:"60"`
	L3_entry_match_drop              int    `json:"l3_entry_match_drop"`
	L3_entry_match_drop_hw           int    `json:"l3_entry_match_drop_hw"`
	L3_entry_drop_max_hw_exceeded    int    `json:"l3_entry_drop_max_hw_exceeded"`
	L4_entry_match_drop              int    `json:"l4_entry_match_drop"`
	L4_entry_match_drop_hw           int    `json:"l4_entry_match_drop_hw"`
	L4_entry_drop_max_hw_exceeded    int    `json:"l4_entry_drop_max_hw_exceeded"`
	L4_entry_list_alloc_failure      int    `json:"l4_entry_list_alloc_failure"`
	Ip_node_alloc_failure            int    `json:"ip_node_alloc_failure"`
	Ip_port_block_alloc_failure      int    `json:"ip_port_block_alloc_failure"`
	Ip_other_block_alloc_failure     int    `json:"ip_other_block_alloc_failure"`
	L3_entry_add_to_bgp_failure      int    `json:"l3_entry_add_to_bgp_failure"`
	L3_entry_remove_from_bgp_failure int    `json:"l3_entry_remove_from_bgp_failure"`
	L3_entry_add_to_hw_failure       int    `json:"l3_entry_add_to_hw_failure"`
	Syn_cookie_verification_failed   int    `json:"syn_cookie_verification_failed"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Lsn3382 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc3383  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate3384 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsInc3383 struct {
	User_quota_failure                      int    `json:"user_quota_failure"`
	Data_sesn_user_quota_exceeded           int    `json:"data_sesn_user_quota_exceeded"`
	Fullcone_failure                        int    `json:"fullcone_failure"`
	Fullcone_self_hairpinning_drop          int    `json:"fullcone_self_hairpinning_drop"`
	Nat_pool_unusable                       int    `json:"nat_pool_unusable"`
	Ha_nat_pool_unusable                    int    `json:"ha_nat_pool_unusable"`
	Ha_nat_pool_batch_type_mismatch         int    `json:"ha_nat_pool_batch_type_mismatch"`
	Sip_alg_quota_inc_failure               int    `json:"sip_alg_quota_inc_failure"`
	Sip_alg_alloc_rtp_rtcp_port_failure     int    `json:"sip_alg_alloc_rtp_rtcp_port_failure"`
	Sip_alg_alloc_single_port_failure       int    `json:"sip_alg_alloc_single_port_failure"`
	Sip_alg_create_single_fullcone_failure  int    `json:"sip_alg_create_single_fullcone_failure"`
	Sip_alg_create_rtp_fullcone_failure     int    `json:"sip_alg_create_rtp_fullcone_failure"`
	Sip_alg_create_rtcp_fullcone_failure    int    `json:"sip_alg_create_rtcp_fullcone_failure"`
	H323_alg_alloc_single_port_failure      int    `json:"h323_alg_alloc_single_port_failure"`
	H323_alg_create_single_fullcone_failure int    `json:"h323_alg_create_single_fullcone_failure"`
	H323_alg_create_rtp_fullcone_failure    int    `json:"h323_alg_create_rtp_fullcone_failure"`
	H323_alg_create_rtcp_fullcone_failure   int    `json:"h323_alg_create_rtcp_fullcone_failure"`
	Port_overloading_out_of_memory          int    `json:"port_overloading_out_of_memory"`
	Port_overloading_inc_overflow           int    `json:"port_overloading_inc_overflow"`
	Fullcone_ext_mem_alloc_failure          int    `json:"fullcone_ext_mem_alloc_failure"`
	Fullcone_ext_mem_alloc_init_faulure     int    `json:"fullcone_ext_mem_alloc_init_faulure"`
	Mgcp_alg_create_rtp_fullcone_failure    int    `json:"mgcp_alg_create_rtp_fullcone_failure"`
	Mgcp_alg_create_rtcp_fullcone_failure   int    `json:"mgcp_alg_create_rtcp_fullcone_failure"`
	Mgcp_alg_port_pair_alloc_from_quota_par int    `json:"mgcp_alg_port_pair_alloc_from_quota_par"`
	User_quota_unusable_drop                int    `json:"user_quota_unusable_drop"`
	User_quota_unusable                     int    `json:"user_quota_unusable"`
	Adc_port_allocation_failed              int    `json:"adc_port_allocation_failed"`
	Uuid                                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnTriggerStatsRate3384 struct {
	ThresholdExceededBy                     int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                                int    `json:"duration" dval:"60"`
	User_quota_failure                      int    `json:"user_quota_failure"`
	Data_sesn_user_quota_exceeded           int    `json:"data_sesn_user_quota_exceeded"`
	Fullcone_failure                        int    `json:"fullcone_failure"`
	Fullcone_self_hairpinning_drop          int    `json:"fullcone_self_hairpinning_drop"`
	Nat_pool_unusable                       int    `json:"nat_pool_unusable"`
	Ha_nat_pool_unusable                    int    `json:"ha_nat_pool_unusable"`
	Ha_nat_pool_batch_type_mismatch         int    `json:"ha_nat_pool_batch_type_mismatch"`
	Sip_alg_quota_inc_failure               int    `json:"sip_alg_quota_inc_failure"`
	Sip_alg_alloc_rtp_rtcp_port_failure     int    `json:"sip_alg_alloc_rtp_rtcp_port_failure"`
	Sip_alg_alloc_single_port_failure       int    `json:"sip_alg_alloc_single_port_failure"`
	Sip_alg_create_single_fullcone_failure  int    `json:"sip_alg_create_single_fullcone_failure"`
	Sip_alg_create_rtp_fullcone_failure     int    `json:"sip_alg_create_rtp_fullcone_failure"`
	Sip_alg_create_rtcp_fullcone_failure    int    `json:"sip_alg_create_rtcp_fullcone_failure"`
	H323_alg_alloc_single_port_failure      int    `json:"h323_alg_alloc_single_port_failure"`
	H323_alg_create_single_fullcone_failure int    `json:"h323_alg_create_single_fullcone_failure"`
	H323_alg_create_rtp_fullcone_failure    int    `json:"h323_alg_create_rtp_fullcone_failure"`
	H323_alg_create_rtcp_fullcone_failure   int    `json:"h323_alg_create_rtcp_fullcone_failure"`
	Port_overloading_out_of_memory          int    `json:"port_overloading_out_of_memory"`
	Port_overloading_inc_overflow           int    `json:"port_overloading_inc_overflow"`
	Fullcone_ext_mem_alloc_failure          int    `json:"fullcone_ext_mem_alloc_failure"`
	Fullcone_ext_mem_alloc_init_faulure     int    `json:"fullcone_ext_mem_alloc_init_faulure"`
	Mgcp_alg_create_rtp_fullcone_failure    int    `json:"mgcp_alg_create_rtp_fullcone_failure"`
	Mgcp_alg_create_rtcp_fullcone_failure   int    `json:"mgcp_alg_create_rtcp_fullcone_failure"`
	Mgcp_alg_port_pair_alloc_from_quota_par int    `json:"mgcp_alg_port_pair_alloc_from_quota_par"`
	User_quota_unusable_drop                int    `json:"user_quota_unusable_drop"`
	User_quota_unusable                     int    `json:"user_quota_unusable"`
	Adc_port_allocation_failed              int    `json:"adc_port_allocation_failed"`
	Uuid                                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEsp3385 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc3386  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate3387 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsInc3386 struct {
	NatIpConflict int    `json:"nat-ip-conflict"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgEspTriggerStatsRate3387 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	NatIpConflict       int    `json:"nat-ip-conflict"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptp3388 struct {
	Uuid             string                                                                                                        `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc3389  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate3390 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsInc3389 struct {
	NoGreSessionMatch          int    `json:"no-gre-session-match"`
	CallReqPnsCallIdMismatch   int    `json:"call-req-pns-call-id-mismatch"`
	CallReplyPnsCallIdMismatch int    `json:"call-reply-pns-call-id-mismatch"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgPptpTriggerStatsRate3390 struct {
	ThresholdExceededBy        int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                   int    `json:"duration" dval:"60"`
	NoGreSessionMatch          int    `json:"no-gre-session-match"`
	CallReqPnsCallIdMismatch   int    `json:"call-req-pns-call-id-mismatch"`
	CallReplyPnsCallIdMismatch int    `json:"call-reply-pns-call-id-mismatch"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtsp3391 struct {
	Uuid             string                                                                                                        `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc3392  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate3393 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsInc3392 struct {
	StreamCreationFailure       int    `json:"stream-creation-failure"`
	PortAllocationFailure       int    `json:"port-allocation-failure"`
	UnknownClientPortFromServer int    `json:"unknown-client-port-from-server"`
	NoSessionMem                int    `json:"no-session-mem"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgRtspTriggerStatsRate3393 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	StreamCreationFailure       int    `json:"stream-creation-failure"`
	PortAllocationFailure       int    `json:"port-allocation-failure"`
	UnknownClientPortFromServer int    `json:"unknown-client-port-from-server"`
	NoSessionMem                int    `json:"no-session-mem"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSip3394 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc3395  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate3396 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsInc3395 struct {
	MethodUnknown     int    `json:"method-unknown"`
	ParseError        int    `json:"parse-error"`
	TcpOutOfOrderDrop int    `json:"tcp-out-of-order-drop"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgSipTriggerStatsRate3396 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	MethodUnknown       int    `json:"method-unknown"`
	ParseError          int    `json:"parse-error"`
	TcpOutOfOrderDrop   int    `json:"tcp-out-of-order-drop"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcp3397 struct {
	Uuid             string                                                                                                        `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc3398  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate3399 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsInc3398 struct {
	ParseError        int    `json:"parse-error"`
	TcpOutOfOrderDrop int    `json:"tcp-out-of-order-drop"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgMgcpTriggerStatsRate3399 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	ParseError          int    `json:"parse-error"`
	TcpOutOfOrderDrop   int    `json:"tcp-out-of-order-drop"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH3233400 struct {
	Uuid             string                                                                                                        `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc3401  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate3402 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsInc3401 struct {
	ParseError        int    `json:"parse-error"`
	TcpOutOfOrderDrop int    `json:"tcp-out-of-order-drop"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnAlgH323TriggerStatsRate3402 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	ParseError          int    `json:"parse-error"`
	TcpOutOfOrderDrop   int    `json:"tcp-out-of-order-drop"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadius3403 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc3404  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate3405 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsInc3404 struct {
	RadiusRequestDropped       int    `json:"radius-request-dropped"`
	RequestBadSecretDropped    int    `json:"request-bad-secret-dropped"`
	RequestNoKeyVapDropped     int    `json:"request-no-key-vap-dropped"`
	RequestMalformedDropped    int    `json:"request-malformed-dropped"`
	RequestIgnored             int    `json:"request-ignored"`
	RadiusTableFull            int    `json:"radius-table-full"`
	SecretNotConfiguredDropped int    `json:"secret-not-configured-dropped"`
	HaStandbyDropped           int    `json:"ha-standby-dropped"`
	InvalidKey                 int    `json:"invalid-key"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LsnRadiusTriggerStatsRate3405 struct {
	ThresholdExceededBy        int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                   int    `json:"duration" dval:"60"`
	RadiusRequestDropped       int    `json:"radius-request-dropped"`
	RequestBadSecretDropped    int    `json:"request-bad-secret-dropped"`
	RequestNoKeyVapDropped     int    `json:"request-no-key-vap-dropped"`
	RequestMalformedDropped    int    `json:"request-malformed-dropped"`
	RequestIgnored             int    `json:"request-ignored"`
	RadiusTableFull            int    `json:"radius-table-full"`
	SecretNotConfiguredDropped int    `json:"secret-not-configured-dropped"`
	HaStandbyDropped           int    `json:"ha-standby-dropped"`
	InvalidKey                 int    `json:"invalid-key"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64Global3406 struct {
	Uuid             string                                                                                                         `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc3407  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate3408 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsInc3407 struct {
	User_quota_failure              int    `json:"user_quota_failure"`
	Nat_port_unavailable_tcp        int    `json:"nat_port_unavailable_tcp"`
	Nat_port_unavailable_udp        int    `json:"nat_port_unavailable_udp"`
	Nat_port_unavailable_icmp       int    `json:"nat_port_unavailable_icmp"`
	New_user_resource_unavailable   int    `json:"new_user_resource_unavailable"`
	Fullcone_failure                int    `json:"fullcone_failure"`
	Fullcone_self_hairpinning_drop  int    `json:"fullcone_self_hairpinning_drop"`
	Eif_limit_exceeded              int    `json:"eif_limit_exceeded"`
	Nat_pool_unusable               int    `json:"nat_pool_unusable"`
	Ha_nat_pool_unusable            int    `json:"ha_nat_pool_unusable"`
	Ha_nat_pool_batch_type_mismatch int    `json:"ha_nat_pool_batch_type_mismatch"`
	No_radius_profile_match         int    `json:"no_radius_profile_match"`
	No_class_list_match             int    `json:"no_class_list_match"`
	User_quota_unusable_drop        int    `json:"user_quota_unusable_drop"`
	User_quota_unusable             int    `json:"user_quota_unusable"`
	Uuid                            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Nat64GlobalTriggerStatsRate3408 struct {
	ThresholdExceededBy             int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                        int    `json:"duration" dval:"60"`
	User_quota_failure              int    `json:"user_quota_failure"`
	Nat_port_unavailable_tcp        int    `json:"nat_port_unavailable_tcp"`
	Nat_port_unavailable_udp        int    `json:"nat_port_unavailable_udp"`
	Nat_port_unavailable_icmp       int    `json:"nat_port_unavailable_icmp"`
	New_user_resource_unavailable   int    `json:"new_user_resource_unavailable"`
	Fullcone_failure                int    `json:"fullcone_failure"`
	Fullcone_self_hairpinning_drop  int    `json:"fullcone_self_hairpinning_drop"`
	Eif_limit_exceeded              int    `json:"eif_limit_exceeded"`
	Nat_pool_unusable               int    `json:"nat_pool_unusable"`
	Ha_nat_pool_unusable            int    `json:"ha_nat_pool_unusable"`
	Ha_nat_pool_batch_type_mismatch int    `json:"ha_nat_pool_batch_type_mismatch"`
	No_radius_profile_match         int    `json:"no_radius_profile_match"`
	No_class_list_match             int    `json:"no_class_list_match"`
	User_quota_unusable_drop        int    `json:"user_quota_unusable_drop"`
	User_quota_unusable             int    `json:"user_quota_unusable"`
	Uuid                            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobal3409 struct {
	Uuid             string                                                                                                          `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc3410  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate3411 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsInc3410 struct {
	User_quota_failure        int    `json:"user_quota_failure"`
	Nat_port_unavailable_tcp  int    `json:"nat_port_unavailable_tcp"`
	Nat_port_unavailable_udp  int    `json:"nat_port_unavailable_udp"`
	Nat_port_unavailable_icmp int    `json:"nat_port_unavailable_icmp"`
	Fullcone_failure          int    `json:"fullcone_failure"`
	Uuid                      string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6DsLiteGlobalTriggerStatsRate3411 struct {
	ThresholdExceededBy       int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                  int    `json:"duration" dval:"60"`
	User_quota_failure        int    `json:"user_quota_failure"`
	Nat_port_unavailable_tcp  int    `json:"nat_port_unavailable_tcp"`
	Nat_port_unavailable_udp  int    `json:"nat_port_unavailable_udp"`
	Nat_port_unavailable_icmp int    `json:"nat_port_unavailable_icmp"`
	Fullcone_failure          int    `json:"fullcone_failure"`
	Uuid                      string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobal3412 struct {
	Uuid             string                                                                                                            `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc3413  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate3414 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsInc3413 struct {
	NatPortUnavailableTcp              int    `json:"nat-port-unavailable-tcp"`
	NatPortUnavailableUdp              int    `json:"nat-port-unavailable-udp"`
	NatPortUnavailableIcmp             int    `json:"nat-port-unavailable-icmp"`
	SessionUserQuotaExceeded           int    `json:"session-user-quota-exceeded"`
	FullconeFailure                    int    `json:"fullcone-failure"`
	Nat44InboundFiltered               int    `json:"nat44-inbound-filtered"`
	Nat64InboundFiltered               int    `json:"nat64-inbound-filtered"`
	DsliteInboundFiltered              int    `json:"dslite-inbound-filtered"`
	Nat44EifLimitExceeded              int    `json:"nat44-eif-limit-exceeded"`
	Nat64EifLimitExceeded              int    `json:"nat64-eif-limit-exceeded"`
	DsliteEifLimitExceeded             int    `json:"dslite-eif-limit-exceeded"`
	StandbyDrop                        int    `json:"standby-drop"`
	FixedNatFullconeSelfHairpinningDro int    `json:"fixed-nat-fullcone-self-hairpinning-dro"`
	SixrdDrop                          int    `json:"sixrd-drop"`
	DestRlistDrop                      int    `json:"dest-rlist-drop"`
	DestRlistPassThrough               int    `json:"dest-rlist-pass-through"`
	DestRlistSnatDrop                  int    `json:"dest-rlist-snat-drop"`
	ConfigNotFound                     int    `json:"config-not-found"`
	PortOverloadFailed                 int    `json:"port-overload-failed"`
	HaSessionUserQuotaExceeded         int    `json:"ha-session-user-quota-exceeded"`
	UnusedIpDrop                       int    `json:"unused-ip-drop"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatGlobalTriggerStatsRate3414 struct {
	ThresholdExceededBy                int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                           int    `json:"duration" dval:"60"`
	NatPortUnavailableTcp              int    `json:"nat-port-unavailable-tcp"`
	NatPortUnavailableUdp              int    `json:"nat-port-unavailable-udp"`
	NatPortUnavailableIcmp             int    `json:"nat-port-unavailable-icmp"`
	SessionUserQuotaExceeded           int    `json:"session-user-quota-exceeded"`
	FullconeFailure                    int    `json:"fullcone-failure"`
	Nat44InboundFiltered               int    `json:"nat44-inbound-filtered"`
	Nat64InboundFiltered               int    `json:"nat64-inbound-filtered"`
	DsliteInboundFiltered              int    `json:"dslite-inbound-filtered"`
	Nat44EifLimitExceeded              int    `json:"nat44-eif-limit-exceeded"`
	Nat64EifLimitExceeded              int    `json:"nat64-eif-limit-exceeded"`
	DsliteEifLimitExceeded             int    `json:"dslite-eif-limit-exceeded"`
	StandbyDrop                        int    `json:"standby-drop"`
	FixedNatFullconeSelfHairpinningDro int    `json:"fixed-nat-fullcone-self-hairpinning-dro"`
	SixrdDrop                          int    `json:"sixrd-drop"`
	DestRlistDrop                      int    `json:"dest-rlist-drop"`
	DestRlistPassThrough               int    `json:"dest-rlist-pass-through"`
	DestRlistSnatDrop                  int    `json:"dest-rlist-snat-drop"`
	ConfigNotFound                     int    `json:"config-not-found"`
	PortOverloadFailed                 int    `json:"port-overload-failed"`
	HaSessionUserQuotaExceeded         int    `json:"ha-session-user-quota-exceeded"`
	UnusedIpDrop                       int    `json:"unused-ip-drop"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptp3415 struct {
	Uuid             string                                                                                                             `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc3416  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate3417 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsInc3416 struct {
	CallReqPnsCallIdMismatch   int    `json:"call-req-pns-call-id-mismatch"`
	CallReplyPnsCallIdMismatch int    `json:"call-reply-pns-call-id-mismatch"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgPptpTriggerStatsRate3417 struct {
	ThresholdExceededBy        int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                   int    `json:"duration" dval:"60"`
	CallReqPnsCallIdMismatch   int    `json:"call-req-pns-call-id-mismatch"`
	CallReplyPnsCallIdMismatch int    `json:"call-reply-pns-call-id-mismatch"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtsp3418 struct {
	Uuid             string                                                                                                             `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc3419  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate3420 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsInc3419 struct {
	StreamCreationFailure int    `json:"stream-creation-failure"`
	PortAllocationFailure int    `json:"port-allocation-failure"`
	NoSessionMem          int    `json:"no-session-mem"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgRtspTriggerStatsRate3420 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	StreamCreationFailure int    `json:"stream-creation-failure"`
	PortAllocationFailure int    `json:"port-allocation-failure"`
	NoSessionMem          int    `json:"no-session-mem"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSip3421 struct {
	Uuid             string                                                                                                            `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc3422  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate3423 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsInc3422 struct {
	MethodUnknown int    `json:"method-unknown"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6FixedNatAlgSipTriggerStatsRate3423 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	MethodUnknown       int    `json:"method-unknown"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Pcp3424 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc3425  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate3426 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsInc3425 struct {
	PktNotRequestDrop    int    `json:"pkt-not-request-drop"`
	PktTooShortDrop      int    `json:"pkt-too-short-drop"`
	NorouteDrop          int    `json:"noroute-drop"`
	UnsupportedVersion   int    `json:"unsupported-version"`
	NotAuthorized        int    `json:"not-authorized"`
	MalformRequest       int    `json:"malform-request"`
	UnsuppOpcode         int    `json:"unsupp-opcode"`
	UnsuppOption         int    `json:"unsupp-option"`
	MalformOption        int    `json:"malform-option"`
	NoResources          int    `json:"no-resources"`
	UnsuppProtocol       int    `json:"unsupp-protocol"`
	CannotProvideSuggest int    `json:"cannot-provide-suggest"`
	AddressMismatch      int    `json:"address-mismatch"`
	ExcessiveRemotePeers int    `json:"excessive-remote-peers"`
	PktNotFromNatInside  int    `json:"pkt-not-from-nat-inside"`
	L4ProcessError       int    `json:"l4-process-error"`
	InternalErrorDrop    int    `json:"internal-error-drop"`
	Unsol_ance_sent_fail int    `json:"unsol_ance_sent_fail"`
	Uuid                 string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6PcpTriggerStatsRate3426 struct {
	ThresholdExceededBy  int    `json:"threshold-exceeded-by" dval:"5"`
	Duration             int    `json:"duration" dval:"60"`
	PktNotRequestDrop    int    `json:"pkt-not-request-drop"`
	PktTooShortDrop      int    `json:"pkt-too-short-drop"`
	NorouteDrop          int    `json:"noroute-drop"`
	UnsupportedVersion   int    `json:"unsupported-version"`
	NotAuthorized        int    `json:"not-authorized"`
	MalformRequest       int    `json:"malform-request"`
	UnsuppOpcode         int    `json:"unsupp-opcode"`
	UnsuppOption         int    `json:"unsupp-option"`
	MalformOption        int    `json:"malform-option"`
	NoResources          int    `json:"no-resources"`
	UnsuppProtocol       int    `json:"unsupp-protocol"`
	CannotProvideSuggest int    `json:"cannot-provide-suggest"`
	AddressMismatch      int    `json:"address-mismatch"`
	ExcessiveRemotePeers int    `json:"excessive-remote-peers"`
	PktNotFromNatInside  int    `json:"pkt-not-from-nat-inside"`
	L4ProcessError       int    `json:"l4-process-error"`
	InternalErrorDrop    int    `json:"internal-error-drop"`
	Unsol_ance_sent_fail int    `json:"unsol_ance_sent_fail"`
	Uuid                 string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Logging3427 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc3428  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate3429 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsInc3428 struct {
	LogDropped     int    `json:"log-dropped"`
	ConnTcpDropped int    `json:"conn-tcp-dropped"`
	Uuid           string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6LoggingTriggerStatsRate3429 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	LogDropped          int    `json:"log-dropped"`
	ConnTcpDropped      int    `json:"conn-tcp-dropped"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L43430 struct {
	Uuid             string                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc3431  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate3432 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsInc3431 struct {
	OutOfSessionMemory      int    `json:"out-of-session-memory"`
	IcmpHostUnreachableSent int    `json:"icmp-host-unreachable-sent"`
	Uuid                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6L4TriggerStatsRate3432 struct {
	ThresholdExceededBy     int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                int    `json:"duration" dval:"60"`
	OutOfSessionMemory      int    `json:"out-of-session-memory"`
	IcmpHostUnreachableSent int    `json:"icmp-host-unreachable-sent"`
	Uuid                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Icmp3433 struct {
	Uuid             string                                                                                                  `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc3434  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate3435 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsInc3434 struct {
	IcmpToIcmpErr     int    `json:"icmp-to-icmp-err"`
	IcmpToIcmpv6Err   int    `json:"icmp-to-icmpv6-err"`
	Icmpv6ToIcmpErr   int    `json:"icmpv6-to-icmp-err"`
	Icmpv6ToIcmpv6Err int    `json:"icmpv6-to-icmpv6-err"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6IcmpTriggerStatsRate3435 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	IcmpToIcmpErr       int    `json:"icmp-to-icmp-err"`
	IcmpToIcmpv6Err     int    `json:"icmp-to-icmpv6-err"`
	Icmpv6ToIcmpErr     int    `json:"icmpv6-to-icmp-err"`
	Icmpv6ToIcmpv6Err   int    `json:"icmpv6-to-icmpv6-err"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlg3436 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc3437  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate3438 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsInc3437 struct {
	RadiusRequstDropped   int    `json:"radius-requst-dropped"`
	RadiusResponseDropped int    `json:"radius-response-dropped"`
	OutOfMemoryDropped    int    `json:"out-of-memory-dropped"`
	QueueLenExceedDropped int    `json:"queue-len-exceed-dropped"`
	OutOfOrderDropped     int    `json:"out-of-order-dropped"`
	HeaderInsertionFailed int    `json:"header-insertion-failed"`
	HeaderRemovalFailed   int    `json:"header-removal-failed"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6HttpAlgTriggerStatsRate3438 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	RadiusRequstDropped   int    `json:"radius-requst-dropped"`
	RadiusResponseDropped int    `json:"radius-response-dropped"`
	OutOfMemoryDropped    int    `json:"out-of-memory-dropped"`
	QueueLenExceedDropped int    `json:"queue-len-exceed-dropped"`
	OutOfOrderDropped     int    `json:"out-of-order-dropped"`
	HeaderInsertionFailed int    `json:"header-insertion-failed"`
	HeaderRemovalFailed   int    `json:"header-removal-failed"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns643439 struct {
	Uuid             string                                                                                                   `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc3440  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate3441 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsInc3440 struct {
	QueryBadPkt int    `json:"query-bad-pkt"`
	RespBadPkt  int    `json:"resp-bad-pkt"`
	RespBadQr   int    `json:"resp-bad-qr"`
	Drop        int    `json:"drop"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dns64TriggerStatsRate3441 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	QueryBadPkt         int    `json:"query-bad-pkt"`
	RespBadPkt          int    `json:"resp-bad-pkt"`
	RespBadQr           int    `json:"resp-bad-qr"`
	Drop                int    `json:"drop"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv63442 struct {
	Uuid             string                                                                                                    `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc3443  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate3444 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsInc3443 struct {
	PacketsDropped         int    `json:"packets-dropped"`
	PktsDroppedDuringClear int    `json:"pkts-dropped-during-clear"`
	RcvNotSupportedMsg     int    `json:"rcv-not-supported-msg"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeCgnv6Dhcpv6TriggerStatsRate3444 struct {
	ThresholdExceededBy    int    `json:"threshold-exceeded-by" dval:"5"`
	Duration               int    `json:"duration" dval:"60"`
	PacketsDropped         int    `json:"packets-dropped"`
	PktsDroppedDuringClear int    `json:"pkts-dropped-during-clear"`
	RcvNotSupportedMsg     int    `json:"rcv-not-supported-msg"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLogging3445 struct {
	Uuid             string                                                                                                  `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsInc3446  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsRate3447 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsInc3446 struct {
	LogDropped               int    `json:"log-dropped"`
	HttpLoggingInvalidFormat int    `json:"http-logging-invalid-format"`
	SessionLimitExceeded     int    `json:"session-limit-exceeded"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwLoggingTriggerStatsRate3447 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	LogDropped               int    `json:"log-dropped"`
	HttpLoggingInvalidFormat int    `json:"http-logging-invalid-format"`
	SessionLimitExceeded     int    `json:"session-limit-exceeded"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobal3448 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsInc3449  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsRate3450 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsInc3449 struct {
	Fullcone_creation_failure int    `json:"fullcone_creation_failure"`
	Uuid                      string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGlobalTriggerStatsRate3450 struct {
	ThresholdExceededBy       int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                  int    `json:"duration" dval:"60"`
	Fullcone_creation_failure int    `json:"fullcone_creation_failure"`
	Uuid                      string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtsp3451 struct {
	Uuid             string                                                                                                  `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc3452  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate3453 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsInc3452 struct {
	TransportAllocFailure int    `json:"transport-alloc-failure"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgRtspTriggerStatsRate3453 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	TransportAllocFailure int    `json:"transport-alloc-failure"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptp3454 struct {
	Uuid             string                                                                                                  `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc3455  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate3456 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsInc3455 struct {
	CallReqPnsCallIdMismatch   int    `json:"call-req-pns-call-id-mismatch"`
	CallReplyPnsCallIdMismatch int    `json:"call-reply-pns-call-id-mismatch"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwAlgPptpTriggerStatsRate3456 struct {
	ThresholdExceededBy        int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                   int    `json:"duration" dval:"60"`
	CallReqPnsCallIdMismatch   int    `json:"call-req-pns-call-id-mismatch"`
	CallReplyPnsCallIdMismatch int    `json:"call-reply-pns-call-id-mismatch"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServer3457 struct {
	Uuid             string                                                                                                    `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsInc3458  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsRate3459 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsInc3458 struct {
	RadiusRequestDropped     int    `json:"radius-request-dropped"`
	RequestBadSecretDropped  int    `json:"request-bad-secret-dropped"`
	RequestNoKeyVapDropped   int    `json:"request-no-key-vap-dropped"`
	RequestMalformedDropped  int    `json:"request-malformed-dropped"`
	RequestIgnored           int    `json:"request-ignored"`
	RadiusTableFull          int    `json:"radius-table-full"`
	HaStandbyDropped         int    `json:"ha-standby-dropped"`
	Ipv6PrefixLengthMismatch int    `json:"ipv6-prefix-length-mismatch"`
	InvalidKey               int    `json:"invalid-key"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwRadServerTriggerStatsRate3459 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	RadiusRequestDropped     int    `json:"radius-request-dropped"`
	RequestBadSecretDropped  int    `json:"request-bad-secret-dropped"`
	RequestNoKeyVapDropped   int    `json:"request-no-key-vap-dropped"`
	RequestMalformedDropped  int    `json:"request-malformed-dropped"`
	RequestIgnored           int    `json:"request-ignored"`
	RadiusTableFull          int    `json:"radius-table-full"`
	HaStandbyDropped         int    `json:"ha-standby-dropped"`
	Ipv6PrefixLengthMismatch int    `json:"ipv6-prefix-length-mismatch"`
	InvalidKey               int    `json:"invalid-key"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookie3460 struct {
	Uuid             string                                                                                                       `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc3461  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate3462 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsInc3461 struct {
	Verification_failed int    `json:"verification_failed"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwTcpSynCookieTriggerStatsRate3462 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Verification_failed int    `json:"verification_failed"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtection3463 struct {
	Uuid             string                                                                                                         `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc3464  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate3465 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsInc3464 struct {
	Ddos_entries_too_many              int    `json:"ddos_entries_too_many"`
	Ddos_entry_add_to_bgp_failure      int    `json:"ddos_entry_add_to_bgp_failure"`
	Ddos_entry_remove_from_bgp_failure int    `json:"ddos_entry_remove_from_bgp_failure"`
	Ddos_packet_dropped                int    `json:"ddos_packet_dropped"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwDdosProtectionTriggerStatsRate3465 struct {
	ThresholdExceededBy                int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                           int    `json:"duration" dval:"60"`
	Ddos_entries_too_many              int    `json:"ddos_entries_too_many"`
	Ddos_entry_add_to_bgp_failure      int    `json:"ddos_entry_add_to_bgp_failure"`
	Ddos_entry_remove_from_bgp_failure int    `json:"ddos_entry_remove_from_bgp_failure"`
	Ddos_packet_dropped                int    `json:"ddos_packet_dropped"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtp3466 struct {
	Uuid             string                                                                                              `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsInc3467  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsRate3468 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsInc3467 struct {
	OutOfSessionMemory                int    `json:"out-of-session-memory"`
	BladeOutOfSessionMemory           int    `json:"blade-out-of-session-memory"`
	GtpSmpPathCheckFailed             int    `json:"gtp-smp-path-check-failed"`
	GtpSmpCCheckFailed                int    `json:"gtp-smp-c-check-failed"`
	BladeGtpSmpPathCheckFailed        int    `json:"blade-gtp-smp-path-check-failed"`
	BladeGtpSmpCCheckFailed           int    `json:"blade-gtp-smp-c-check-failed"`
	GtpTunnelRateLimitEntryCreateFail int    `json:"gtp-tunnel-rate-limit-entry-create-fail"`
	GtpUTunnelRateLimitEntryCreateFa  int    `json:"gtp-u-tunnel-rate-limit-entry-create-fa"`
	GtpRateLimitSmpCreateFailure      int    `json:"gtp-rate-limit-smp-create-failure"`
	GtpRateLimitT3CtrCreateFailure    int    `json:"gtp-rate-limit-t3-ctr-create-failure"`
	GtpRateLimitEntryCreateFailure    int    `json:"gtp-rate-limit-entry-create-failure"`
	BladeGtpRateLimitSmpCreateFailure int    `json:"blade-gtp-rate-limit-smp-create-failure"`
	BladeGtpRateLimitT3CtrCreateFail  int    `json:"blade-gtp-rate-limit-t3-ctr-create-fail"`
	BladeGtpRateLimitEntryCreateFailu int    `json:"blade-gtp-rate-limit-entry-create-failu"`
	GtpSmpDecSessCountCheckFailed     int    `json:"gtp-smp-dec-sess-count-check-failed"`
	GtpUSmpCheckFailed                int    `json:"gtp-u-smp-check-failed"`
	GtpInfoExtNotFound                int    `json:"gtp-info-ext-not-found"`
	BladeGtpSmpDecSessCountCheckFail  int    `json:"blade-gtp-smp-dec-sess-count-check-fail"`
	BladeGtpUSmpCheckFailed           int    `json:"blade-gtp-u-smp-check-failed"`
	BladeGtpInfoExtNotFound           int    `json:"blade-gtp-info-ext-not-found"`
	BladeGtpSmpSessionCountCheckFaile int    `json:"blade-gtp-smp-session-count-check-faile"`
	GtpCSmpSigCheckFailed             int    `json:"gtp-c-smp-sig-check-failed"`
	BladeGtpCSmpSigCheckFailed        int    `json:"blade-gtp-c-smp-sig-check-failed"`
	GtpUSmpSigCheckFailed             int    `json:"gtp-u-smp-sig-check-failed"`
	BladeGtpUSmpSigCheckFailed        int    `json:"blade-gtp-u-smp-sig-check-failed"`
	GtpSmpSigCheckFailed              int    `json:"gtp-smp-sig-check-failed"`
	BladeGtpSmpSigCheckFailed         int    `json:"blade-gtp-smp-sig-check-failed"`
	GtpCFailConnCreateSlow            int    `json:"gtp-c-fail-conn-create-slow"`
	GtpUFailConnCreateSlow            int    `json:"gtp-u-fail-conn-create-slow"`
	GtpPathmFailConnCreateSlow        int    `json:"gtp-pathm-fail-conn-create-slow"`
	Uuid                              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeFwGtpTriggerStatsRate3468 struct {
	ThresholdExceededBy               int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                          int    `json:"duration" dval:"60"`
	OutOfSessionMemory                int    `json:"out-of-session-memory"`
	BladeOutOfSessionMemory           int    `json:"blade-out-of-session-memory"`
	GtpSmpPathCheckFailed             int    `json:"gtp-smp-path-check-failed"`
	GtpSmpCCheckFailed                int    `json:"gtp-smp-c-check-failed"`
	BladeGtpSmpPathCheckFailed        int    `json:"blade-gtp-smp-path-check-failed"`
	BladeGtpSmpCCheckFailed           int    `json:"blade-gtp-smp-c-check-failed"`
	GtpTunnelRateLimitEntryCreateFail int    `json:"gtp-tunnel-rate-limit-entry-create-fail"`
	GtpUTunnelRateLimitEntryCreateFa  int    `json:"gtp-u-tunnel-rate-limit-entry-create-fa"`
	GtpRateLimitSmpCreateFailure      int    `json:"gtp-rate-limit-smp-create-failure"`
	GtpRateLimitT3CtrCreateFailure    int    `json:"gtp-rate-limit-t3-ctr-create-failure"`
	GtpRateLimitEntryCreateFailure    int    `json:"gtp-rate-limit-entry-create-failure"`
	BladeGtpRateLimitSmpCreateFailure int    `json:"blade-gtp-rate-limit-smp-create-failure"`
	BladeGtpRateLimitT3CtrCreateFail  int    `json:"blade-gtp-rate-limit-t3-ctr-create-fail"`
	BladeGtpRateLimitEntryCreateFailu int    `json:"blade-gtp-rate-limit-entry-create-failu"`
	GtpSmpDecSessCountCheckFailed     int    `json:"gtp-smp-dec-sess-count-check-failed"`
	GtpUSmpCheckFailed                int    `json:"gtp-u-smp-check-failed"`
	GtpInfoExtNotFound                int    `json:"gtp-info-ext-not-found"`
	BladeGtpSmpDecSessCountCheckFail  int    `json:"blade-gtp-smp-dec-sess-count-check-fail"`
	BladeGtpUSmpCheckFailed           int    `json:"blade-gtp-u-smp-check-failed"`
	BladeGtpInfoExtNotFound           int    `json:"blade-gtp-info-ext-not-found"`
	BladeGtpSmpSessionCountCheckFaile int    `json:"blade-gtp-smp-session-count-check-faile"`
	GtpCSmpSigCheckFailed             int    `json:"gtp-c-smp-sig-check-failed"`
	BladeGtpCSmpSigCheckFailed        int    `json:"blade-gtp-c-smp-sig-check-failed"`
	GtpUSmpSigCheckFailed             int    `json:"gtp-u-smp-sig-check-failed"`
	BladeGtpUSmpSigCheckFailed        int    `json:"blade-gtp-u-smp-sig-check-failed"`
	GtpSmpSigCheckFailed              int    `json:"gtp-smp-sig-check-failed"`
	BladeGtpSmpSigCheckFailed         int    `json:"blade-gtp-smp-sig-check-failed"`
	GtpCFailConnCreateSlow            int    `json:"gtp-c-fail-conn-create-slow"`
	GtpUFailConnCreateSlow            int    `json:"gtp-u-fail-conn-create-slow"`
	GtpPathmFailConnCreateSlow        int    `json:"gtp-pathm-fail-conn-create-slow"`
	Uuid                              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcp3469 struct {
	Uuid             string                                                                                                  `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsInc3470  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsRate3471 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsInc3470 struct {
	Attemptfails int    `json:"attemptfails"`
	Noroute      int    `json:"noroute"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSystemTcpTriggerStatsRate3471 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Attemptfails        int    `json:"attemptfails"`
	Noroute             int    `json:"noroute"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuse3472 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc3473  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate3474 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsInc3473 struct {
	Ntermi_err      int    `json:"ntermi_err"`
	Pause_conn_fail int    `json:"pause_conn_fail"`
	Uuid            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbConnReuseTriggerStatsRate3474 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Ntermi_err          int    `json:"ntermi_err"`
	Pause_conn_fail     int    `json:"pause_conn_fail"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflow3475 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsInc3476  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsRate3477 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsInc3476 struct {
	Pause_conn_fail   int    `json:"pause_conn_fail"`
	Error_resume_conn int    `json:"error_resume_conn"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbAflowTriggerStatsRate3477 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Pause_conn_fail     int    `json:"pause_conn_fail"`
	Error_resume_conn   int    `json:"error_resume_conn"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFix3478 struct {
	Uuid             string                                                                                               `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsInc3479  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsRate3480 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsInc3479 struct {
	Svrsel_fail int    `json:"svrsel_fail"`
	Noroute     int    `json:"noroute"`
	Snat_fail   int    `json:"snat_fail"`
	Client_err  int    `json:"client_err"`
	Server_err  int    `json:"server_err"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFixTriggerStatsRate3480 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Svrsel_fail         int    `json:"svrsel_fail"`
	Noroute             int    `json:"noroute"`
	Snat_fail           int    `json:"snat_fail"`
	Client_err          int    `json:"client_err"`
	Server_err          int    `json:"server_err"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxy3481 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc3482  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate3483 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc3482 struct {
	Tcp_err                     int    `json:"tcp_err"`
	Stream_not_found            int    `json:"stream_not_found"`
	Stream_err                  int    `json:"stream_err"`
	Session_err                 int    `json:"session_err"`
	Data_no_stream              int    `json:"data_no_stream"`
	Data_no_stream_no_goaway    int    `json:"data_no_stream_no_goaway"`
	Data_no_stream_goaway_close int    `json:"data_no_stream_goaway_close"`
	Est_cb_no_tuple             int    `json:"est_cb_no_tuple"`
	Data_cb_no_tuple            int    `json:"data_cb_no_tuple"`
	Ctx_alloc_fail              int    `json:"ctx_alloc_fail"`
	Stream_alloc_fail           int    `json:"stream_alloc_fail"`
	Http_conn_alloc_fail        int    `json:"http_conn_alloc_fail"`
	Request_header_alloc_fail   int    `json:"request_header_alloc_fail"`
	Decompress_fail             int    `json:"decompress_fail"`
	Invalid_frame_size          int    `json:"invalid_frame_size"`
	Invalid_version             int    `json:"invalid_version"`
	Compress_ctx_alloc_fail     int    `json:"compress_ctx_alloc_fail"`
	Header_compress_fail        int    `json:"header_compress_fail"`
	Http_err_stream_closed      int    `json:"http_err_stream_closed"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsRate3483 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	Tcp_err                     int    `json:"tcp_err"`
	Stream_not_found            int    `json:"stream_not_found"`
	Stream_err                  int    `json:"stream_err"`
	Session_err                 int    `json:"session_err"`
	Data_no_stream              int    `json:"data_no_stream"`
	Data_no_stream_no_goaway    int    `json:"data_no_stream_no_goaway"`
	Data_no_stream_goaway_close int    `json:"data_no_stream_goaway_close"`
	Est_cb_no_tuple             int    `json:"est_cb_no_tuple"`
	Data_cb_no_tuple            int    `json:"data_cb_no_tuple"`
	Ctx_alloc_fail              int    `json:"ctx_alloc_fail"`
	Stream_alloc_fail           int    `json:"stream_alloc_fail"`
	Http_conn_alloc_fail        int    `json:"http_conn_alloc_fail"`
	Request_header_alloc_fail   int    `json:"request_header_alloc_fail"`
	Decompress_fail             int    `json:"decompress_fail"`
	Invalid_frame_size          int    `json:"invalid_frame_size"`
	Invalid_version             int    `json:"invalid_version"`
	Compress_ctx_alloc_fail     int    `json:"compress_ctx_alloc_fail"`
	Header_compress_fail        int    `json:"header_compress_fail"`
	Http_err_stream_closed      int    `json:"http_err_stream_closed"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp23484 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc3485  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate3486 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc3485 struct {
	Protocol_error                          int    `json:"protocol_error"`
	Internal_error                          int    `json:"internal_error"`
	Proxy_alloc_error                       int    `json:"proxy_alloc_error"`
	Split_buff_fail                         int    `json:"split_buff_fail"`
	Invalid_frame_size                      int    `json:"invalid_frame_size"`
	Error_max_invalid_stream                int    `json:"error_max_invalid_stream"`
	Data_no_stream                          int    `json:"data_no_stream"`
	Flow_control_error                      int    `json:"flow_control_error"`
	Settings_timeout                        int    `json:"settings_timeout"`
	Frame_size_error                        int    `json:"frame_size_error"`
	Refused_stream                          int    `json:"refused_stream"`
	Cancel                                  int    `json:"cancel"`
	Compression_error                       int    `json:"compression_error"`
	Connect_error                           int    `json:"connect_error"`
	Enhance_your_calm                       int    `json:"enhance_your_calm"`
	Inadequate_security                     int    `json:"inadequate_security"`
	Http_1_1_required                       int    `json:"http_1_1_required"`
	Deflate_alloc_fail                      int    `json:"deflate_alloc_fail"`
	Inflate_alloc_fail                      int    `json:"inflate_alloc_fail"`
	Inflate_header_fail                     int    `json:"inflate_header_fail"`
	Bad_connection_preface                  int    `json:"bad_connection_preface"`
	Cant_allocate_control_frame             int    `json:"cant_allocate_control_frame"`
	Cant_allocate_settings_frame            int    `json:"cant_allocate_settings_frame"`
	Bad_frame_type_for_stream_state         int    `json:"bad_frame_type_for_stream_state"`
	Wrong_stream_state                      int    `json:"wrong_stream_state"`
	Data_queue_alloc_error                  int    `json:"data_queue_alloc_error"`
	Buff_alloc_error                        int    `json:"buff_alloc_error"`
	Cant_allocate_rst_frame                 int    `json:"cant_allocate_rst_frame"`
	Cant_allocate_goaway_frame              int    `json:"cant_allocate_goaway_frame"`
	Cant_allocate_ping_frame                int    `json:"cant_allocate_ping_frame"`
	Cant_allocate_stream                    int    `json:"cant_allocate_stream"`
	Cant_allocate_window_frame              int    `json:"cant_allocate_window_frame"`
	Header_no_stream                        int    `json:"header_no_stream"`
	Header_padlen_gt_frame_payload          int    `json:"header_padlen_gt_frame_payload"`
	Streams_gt_max_concur_streams           int    `json:"streams_gt_max_concur_streams"`
	Idle_state_unexpected_frame             int    `json:"idle_state_unexpected_frame"`
	Reserved_local_state_unexpected_frame   int    `json:"reserved_local_state_unexpected_frame"`
	Reserved_remote_state_unexpected_frame  int    `json:"reserved_remote_state_unexpected_frame"`
	Half_closed_remote_state_unexpected_fra int    `json:"half_closed_remote_state_unexpected_fra"`
	Closed_state_unexpected_frame           int    `json:"closed_state_unexpected_frame"`
	Zero_window_size_on_stream              int    `json:"zero_window_size_on_stream"`
	Exceeds_max_window_size_stream          int    `json:"exceeds_max_window_size_stream"`
	Continuation_before_headers             int    `json:"continuation_before_headers"`
	Invalid_frame_during_headers            int    `json:"invalid_frame_during_headers"`
	Headers_after_continuation              int    `json:"headers_after_continuation"`
	Invalid_push_promise                    int    `json:"invalid_push_promise"`
	Invalid_stream_id                       int    `json:"invalid_stream_id"`
	Headers_interleaved                     int    `json:"headers_interleaved"`
	Trailers_no_end_stream                  int    `json:"trailers_no_end_stream"`
	Invalid_setting_value                   int    `json:"invalid_setting_value"`
	Invalid_window_update                   int    `json:"invalid_window_update"`
	Alloc_fail_total                        int    `json:"alloc_fail_total"`
	Err_rcvd_total                          int    `json:"err_rcvd_total"`
	Err_sent_total                          int    `json:"err_sent_total"`
	Err_sent_proto_err                      int    `json:"err_sent_proto_err"`
	Err_sent_internal_err                   int    `json:"err_sent_internal_err"`
	Err_sent_flow_control                   int    `json:"err_sent_flow_control"`
	Err_sent_setting_timeout                int    `json:"err_sent_setting_timeout"`
	Err_sent_stream_closed                  int    `json:"err_sent_stream_closed"`
	Err_sent_frame_size_err                 int    `json:"err_sent_frame_size_err"`
	Err_sent_refused_stream                 int    `json:"err_sent_refused_stream"`
	Err_sent_cancel                         int    `json:"err_sent_cancel"`
	Err_sent_compression_err                int    `json:"err_sent_compression_err"`
	Err_sent_connect_err                    int    `json:"err_sent_connect_err"`
	Err_sent_your_calm                      int    `json:"err_sent_your_calm"`
	Err_sent_inadequate_security            int    `json:"err_sent_inadequate_security"`
	Err_sent_http11_required                int    `json:"err_sent_http11_required"`
	Uuid                                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate3486 struct {
	ThresholdExceededBy                     int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                                int    `json:"duration" dval:"60"`
	Protocol_error                          int    `json:"protocol_error"`
	Internal_error                          int    `json:"internal_error"`
	Proxy_alloc_error                       int    `json:"proxy_alloc_error"`
	Split_buff_fail                         int    `json:"split_buff_fail"`
	Invalid_frame_size                      int    `json:"invalid_frame_size"`
	Error_max_invalid_stream                int    `json:"error_max_invalid_stream"`
	Data_no_stream                          int    `json:"data_no_stream"`
	Flow_control_error                      int    `json:"flow_control_error"`
	Settings_timeout                        int    `json:"settings_timeout"`
	Frame_size_error                        int    `json:"frame_size_error"`
	Refused_stream                          int    `json:"refused_stream"`
	Cancel                                  int    `json:"cancel"`
	Compression_error                       int    `json:"compression_error"`
	Connect_error                           int    `json:"connect_error"`
	Enhance_your_calm                       int    `json:"enhance_your_calm"`
	Inadequate_security                     int    `json:"inadequate_security"`
	Http_1_1_required                       int    `json:"http_1_1_required"`
	Deflate_alloc_fail                      int    `json:"deflate_alloc_fail"`
	Inflate_alloc_fail                      int    `json:"inflate_alloc_fail"`
	Inflate_header_fail                     int    `json:"inflate_header_fail"`
	Bad_connection_preface                  int    `json:"bad_connection_preface"`
	Cant_allocate_control_frame             int    `json:"cant_allocate_control_frame"`
	Cant_allocate_settings_frame            int    `json:"cant_allocate_settings_frame"`
	Bad_frame_type_for_stream_state         int    `json:"bad_frame_type_for_stream_state"`
	Wrong_stream_state                      int    `json:"wrong_stream_state"`
	Data_queue_alloc_error                  int    `json:"data_queue_alloc_error"`
	Buff_alloc_error                        int    `json:"buff_alloc_error"`
	Cant_allocate_rst_frame                 int    `json:"cant_allocate_rst_frame"`
	Cant_allocate_goaway_frame              int    `json:"cant_allocate_goaway_frame"`
	Cant_allocate_ping_frame                int    `json:"cant_allocate_ping_frame"`
	Cant_allocate_stream                    int    `json:"cant_allocate_stream"`
	Cant_allocate_window_frame              int    `json:"cant_allocate_window_frame"`
	Header_no_stream                        int    `json:"header_no_stream"`
	Header_padlen_gt_frame_payload          int    `json:"header_padlen_gt_frame_payload"`
	Streams_gt_max_concur_streams           int    `json:"streams_gt_max_concur_streams"`
	Idle_state_unexpected_frame             int    `json:"idle_state_unexpected_frame"`
	Reserved_local_state_unexpected_frame   int    `json:"reserved_local_state_unexpected_frame"`
	Reserved_remote_state_unexpected_frame  int    `json:"reserved_remote_state_unexpected_frame"`
	Half_closed_remote_state_unexpected_fra int    `json:"half_closed_remote_state_unexpected_fra"`
	Closed_state_unexpected_frame           int    `json:"closed_state_unexpected_frame"`
	Zero_window_size_on_stream              int    `json:"zero_window_size_on_stream"`
	Exceeds_max_window_size_stream          int    `json:"exceeds_max_window_size_stream"`
	Continuation_before_headers             int    `json:"continuation_before_headers"`
	Invalid_frame_during_headers            int    `json:"invalid_frame_during_headers"`
	Headers_after_continuation              int    `json:"headers_after_continuation"`
	Invalid_push_promise                    int    `json:"invalid_push_promise"`
	Invalid_stream_id                       int    `json:"invalid_stream_id"`
	Headers_interleaved                     int    `json:"headers_interleaved"`
	Trailers_no_end_stream                  int    `json:"trailers_no_end_stream"`
	Invalid_setting_value                   int    `json:"invalid_setting_value"`
	Invalid_window_update                   int    `json:"invalid_window_update"`
	Alloc_fail_total                        int    `json:"alloc_fail_total"`
	Err_rcvd_total                          int    `json:"err_rcvd_total"`
	Err_sent_total                          int    `json:"err_sent_total"`
	Err_sent_proto_err                      int    `json:"err_sent_proto_err"`
	Err_sent_internal_err                   int    `json:"err_sent_internal_err"`
	Err_sent_flow_control                   int    `json:"err_sent_flow_control"`
	Err_sent_setting_timeout                int    `json:"err_sent_setting_timeout"`
	Err_sent_stream_closed                  int    `json:"err_sent_stream_closed"`
	Err_sent_frame_size_err                 int    `json:"err_sent_frame_size_err"`
	Err_sent_refused_stream                 int    `json:"err_sent_refused_stream"`
	Err_sent_cancel                         int    `json:"err_sent_cancel"`
	Err_sent_compression_err                int    `json:"err_sent_compression_err"`
	Err_sent_connect_err                    int    `json:"err_sent_connect_err"`
	Err_sent_your_calm                      int    `json:"err_sent_your_calm"`
	Err_sent_inadequate_security            int    `json:"err_sent_inadequate_security"`
	Err_sent_http11_required                int    `json:"err_sent_http11_required"`
	Uuid                                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7session3487 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc3488  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate3489 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsInc3488 struct {
	Conn_not_exist     int    `json:"conn_not_exist"`
	Wbuf_cb_failed     int    `json:"wbuf_cb_failed"`
	Err_event          int    `json:"err_event"`
	Err_cb_failed      int    `json:"err_cb_failed"`
	Server_conn_failed int    `json:"server_conn_failed"`
	Server_select_fail int    `json:"server_select_fail"`
	Data_cb_failed     int    `json:"data_cb_failed"`
	Hps_fwdreq_fail    int    `json:"hps_fwdreq_fail"`
	Uuid               string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate3489 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Conn_not_exist      int    `json:"conn_not_exist"`
	Wbuf_cb_failed      int    `json:"wbuf_cb_failed"`
	Err_event           int    `json:"err_event"`
	Err_cb_failed       int    `json:"err_cb_failed"`
	Server_conn_failed  int    `json:"server_conn_failed"`
	Server_select_fail  int    `json:"server_select_fail"`
	Data_cb_failed      int    `json:"data_cb_failed"`
	Hps_fwdreq_fail     int    `json:"hps_fwdreq_fail"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmpp3490 struct {
	Uuid             string                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsInc3491  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsRate3492 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsInc3491 struct {
	Msg_proxy_client_fail            int    `json:"msg_proxy_client_fail"`
	Msg_proxy_server_fail            int    `json:"msg_proxy_server_fail"`
	Msg_proxy_fail_start_server_conn int    `json:"msg_proxy_fail_start_server_conn"`
	Select_client_fail               int    `json:"select_client_fail"`
	Select_server_fail               int    `json:"select_server_fail"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmppTriggerStatsRate3492 struct {
	ThresholdExceededBy              int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                         int    `json:"duration" dval:"60"`
	Msg_proxy_client_fail            int    `json:"msg_proxy_client_fail"`
	Msg_proxy_server_fail            int    `json:"msg_proxy_server_fail"`
	Msg_proxy_fail_start_server_conn int    `json:"msg_proxy_fail_start_server_conn"`
	Select_client_fail               int    `json:"select_client_fail"`
	Select_server_fail               int    `json:"select_server_fail"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtp3493 struct {
	Uuid             string                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc3494  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate3495 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsInc3494 struct {
	No_proxy                      int    `json:"no_proxy"`
	Parse_req_fail                int    `json:"parse_req_fail"`
	Server_select_fail            int    `json:"server_select_fail"`
	Forward_req_fail              int    `json:"forward_req_fail"`
	Forward_req_data_fail         int    `json:"forward_req_data_fail"`
	Snat_fail                     int    `json:"snat_fail"`
	Send_client_service_not_ready int    `json:"send_client_service_not_ready"`
	Recv_server_unknow_reply_code int    `json:"recv_server_unknow_reply_code"`
	Read_request_line_fail        int    `json:"read_request_line_fail"`
	Get_all_headers_fail          int    `json:"get_all_headers_fail"`
	Too_many_headers              int    `json:"too_many_headers"`
	Line_too_long                 int    `json:"line_too_long"`
	Line_extend_fail              int    `json:"line_extend_fail"`
	Line_table_extend_fail        int    `json:"line_table_extend_fail"`
	Parse_request_line_fail       int    `json:"parse_request_line_fail"`
	Insert_resonse_line_fail      int    `json:"insert_resonse_line_fail"`
	Remove_resonse_line_fail      int    `json:"remove_resonse_line_fail"`
	Parse_resonse_line_fail       int    `json:"parse_resonse_line_fail"`
	Server_starttls_fail          int    `json:"server_STARTTLS_fail"`
	Uuid                          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSmtpTriggerStatsRate3495 struct {
	ThresholdExceededBy           int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                      int    `json:"duration" dval:"60"`
	No_proxy                      int    `json:"no_proxy"`
	Parse_req_fail                int    `json:"parse_req_fail"`
	Server_select_fail            int    `json:"server_select_fail"`
	Forward_req_fail              int    `json:"forward_req_fail"`
	Forward_req_data_fail         int    `json:"forward_req_data_fail"`
	Snat_fail                     int    `json:"snat_fail"`
	Send_client_service_not_ready int    `json:"send_client_service_not_ready"`
	Recv_server_unknow_reply_code int    `json:"recv_server_unknow_reply_code"`
	Read_request_line_fail        int    `json:"read_request_line_fail"`
	Get_all_headers_fail          int    `json:"get_all_headers_fail"`
	Too_many_headers              int    `json:"too_many_headers"`
	Line_too_long                 int    `json:"line_too_long"`
	Line_extend_fail              int    `json:"line_extend_fail"`
	Line_table_extend_fail        int    `json:"line_table_extend_fail"`
	Parse_request_line_fail       int    `json:"parse_request_line_fail"`
	Insert_resonse_line_fail      int    `json:"insert_resonse_line_fail"`
	Remove_resonse_line_fail      int    `json:"remove_resonse_line_fail"`
	Parse_resonse_line_fail       int    `json:"parse_resonse_line_fail"`
	Server_starttls_fail          int    `json:"server_STARTTLS_fail"`
	Uuid                          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqtt3496 struct {
	Uuid             string                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsInc3497  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsRate3498 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsInc3497 struct {
	Parse_connect_fail     int    `json:"parse_connect_fail"`
	Parse_publish_fail     int    `json:"parse_publish_fail"`
	Parse_subscribe_fail   int    `json:"parse_subscribe_fail"`
	Parse_unsubscribe_fail int    `json:"parse_unsubscribe_fail"`
	Tuple_not_linked       int    `json:"tuple_not_linked"`
	Tuple_already_linked   int    `json:"tuple_already_linked"`
	Conn_null              int    `json:"conn_null"`
	Client_id_null         int    `json:"client_id_null"`
	Session_exist          int    `json:"session_exist"`
	Insertion_failed       int    `json:"insertion_failed"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMqttTriggerStatsRate3498 struct {
	ThresholdExceededBy    int    `json:"threshold-exceeded-by" dval:"5"`
	Duration               int    `json:"duration" dval:"60"`
	Parse_connect_fail     int    `json:"parse_connect_fail"`
	Parse_publish_fail     int    `json:"parse_publish_fail"`
	Parse_subscribe_fail   int    `json:"parse_subscribe_fail"`
	Parse_unsubscribe_fail int    `json:"parse_unsubscribe_fail"`
	Tuple_not_linked       int    `json:"tuple_not_linked"`
	Tuple_already_linked   int    `json:"tuple_already_linked"`
	Conn_null              int    `json:"conn_null"`
	Client_id_null         int    `json:"client_id_null"`
	Session_exist          int    `json:"session_exist"`
	Insertion_failed       int    `json:"insertion_failed"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcap3499 struct {
	Uuid             string                                                                                                `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsInc3500  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsRate3501 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsInc3500 struct {
	App_serv_conn_no_pcb_err int    `json:"app_serv_conn_no_pcb_err"`
	App_serv_conn_err        int    `json:"app_serv_conn_err"`
	Chunk1_hdr_err           int    `json:"chunk1_hdr_err"`
	Chunk2_hdr_err           int    `json:"chunk2_hdr_err"`
	Chunk_bad_trail_err      int    `json:"chunk_bad_trail_err"`
	No_payload_next_buff_err int    `json:"no_payload_next_buff_err"`
	No_payload_buff_err      int    `json:"no_payload_buff_err"`
	Resp_hdr_incomplete_err  int    `json:"resp_hdr_incomplete_err"`
	Serv_sel_fail_err        int    `json:"serv_sel_fail_err"`
	Start_icap_conn_fail_err int    `json:"start_icap_conn_fail_err"`
	Prep_req_fail_err        int    `json:"prep_req_fail_err"`
	Icap_ver_err             int    `json:"icap_ver_err"`
	Icap_line_err            int    `json:"icap_line_err"`
	Encap_hdr_incomplete_err int    `json:"encap_hdr_incomplete_err"`
	No_icap_resp_err         int    `json:"no_icap_resp_err"`
	Resp_line_read_err       int    `json:"resp_line_read_err"`
	Resp_line_parse_err      int    `json:"resp_line_parse_err"`
	Resp_hdr_err             int    `json:"resp_hdr_err"`
	Req_hdr_incomplete_err   int    `json:"req_hdr_incomplete_err"`
	No_status_code_err       int    `json:"no_status_code_err"`
	Http_resp_line_read_err  int    `json:"http_resp_line_read_err"`
	Http_resp_line_parse_err int    `json:"http_resp_line_parse_err"`
	Http_resp_hdr_err        int    `json:"http_resp_hdr_err"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbIcapTriggerStatsRate3501 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	App_serv_conn_no_pcb_err int    `json:"app_serv_conn_no_pcb_err"`
	App_serv_conn_err        int    `json:"app_serv_conn_err"`
	Chunk1_hdr_err           int    `json:"chunk1_hdr_err"`
	Chunk2_hdr_err           int    `json:"chunk2_hdr_err"`
	Chunk_bad_trail_err      int    `json:"chunk_bad_trail_err"`
	No_payload_next_buff_err int    `json:"no_payload_next_buff_err"`
	No_payload_buff_err      int    `json:"no_payload_buff_err"`
	Resp_hdr_incomplete_err  int    `json:"resp_hdr_incomplete_err"`
	Serv_sel_fail_err        int    `json:"serv_sel_fail_err"`
	Start_icap_conn_fail_err int    `json:"start_icap_conn_fail_err"`
	Prep_req_fail_err        int    `json:"prep_req_fail_err"`
	Icap_ver_err             int    `json:"icap_ver_err"`
	Icap_line_err            int    `json:"icap_line_err"`
	Encap_hdr_incomplete_err int    `json:"encap_hdr_incomplete_err"`
	No_icap_resp_err         int    `json:"no_icap_resp_err"`
	Resp_line_read_err       int    `json:"resp_line_read_err"`
	Resp_line_parse_err      int    `json:"resp_line_parse_err"`
	Resp_hdr_err             int    `json:"resp_hdr_err"`
	Req_hdr_incomplete_err   int    `json:"req_hdr_incomplete_err"`
	No_status_code_err       int    `json:"no_status_code_err"`
	Http_resp_line_read_err  int    `json:"http_resp_line_read_err"`
	Http_resp_line_parse_err int    `json:"http_resp_line_parse_err"`
	Http_resp_hdr_err        int    `json:"http_resp_hdr_err"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSip3502 struct {
	Uuid             string                                                                                               `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsInc3503  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsRate3504 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsInc3503 struct {
	Msg_proxy_client_fail            int    `json:"msg_proxy_client_fail"`
	Msg_proxy_server_fail            int    `json:"msg_proxy_server_fail"`
	Msg_proxy_fail_start_server_conn int    `json:"msg_proxy_fail_start_server_conn"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSipTriggerStatsRate3504 struct {
	ThresholdExceededBy              int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                         int    `json:"duration" dval:"60"`
	Msg_proxy_client_fail            int    `json:"msg_proxy_client_fail"`
	Msg_proxy_server_fail            int    `json:"msg_proxy_server_fail"`
	Msg_proxy_fail_start_server_conn int    `json:"msg_proxy_fail_start_server_conn"`
	Uuid                             string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompress3505 struct {
	Uuid             string                                                                                                      `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc3506  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate3507 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsInc3506 struct {
	Failure_count                 int    `json:"failure_count"`
	Failure_code                  int    `json:"failure_code"`
	Ring_full_count               int    `json:"ring_full_count"`
	Max_outstanding_request_count int    `json:"max_outstanding_request_count"`
	Max_outstanding_submit_count  int    `json:"max_outstanding_submit_count"`
	Uuid                          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHwCompressTriggerStatsRate3507 struct {
	ThresholdExceededBy           int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                      int    `json:"duration" dval:"60"`
	Failure_count                 int    `json:"failure_count"`
	Failure_code                  int    `json:"failure_code"`
	Ring_full_count               int    `json:"ring_full_count"`
	Max_outstanding_request_count int    `json:"max_outstanding_request_count"`
	Max_outstanding_submit_count  int    `json:"max_outstanding_submit_count"`
	Uuid                          string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysql3508 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc3509  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate3510 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsInc3509 struct {
	Session_err int    `json:"session_err"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMysqlTriggerStatsRate3510 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Session_err         int    `json:"session_err"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssql3511 struct {
	Uuid             string                                                                                                 `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc3512  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate3513 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsInc3512 struct {
	Session_err  int    `json:"session_err"`
	Auth_failure int    `json:"auth_failure"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMssqlTriggerStatsRate3513 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Session_err         int    `json:"session_err"`
	Auth_failure        int    `json:"auth_failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcip3514 struct {
	Uuid             string                                                                                                    `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc3515  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate3516 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsInc3515 struct {
	Out_of_sessions   int    `json:"out_of_sessions"`
	Too_many_sessions int    `json:"too_many_sessions"`
	Threshold_exceed  int    `json:"threshold_exceed"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbCrlSrcipTriggerStatsRate3516 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Out_of_sessions     int    `json:"out_of_sessions"`
	Too_many_sessions   int    `json:"too_many_sessions"`
	Threshold_exceed    int    `json:"threshold_exceed"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGeneric3517 struct {
	Uuid             string                                                                                                   `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsInc3518  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsRate3519 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsInc3518 struct {
	Svrsel_fail               int    `json:"svrsel_fail"`
	No_route                  int    `json:"no_route"`
	Snat_fail                 int    `json:"snat_fail"`
	Client_fail               int    `json:"client_fail"`
	Server_fail               int    `json:"server_fail"`
	Mismatch_fwd_id           int    `json:"mismatch_fwd_id"`
	Mismatch_rev_id           int    `json:"mismatch_rev_id"`
	Unkwn_cmd_code            int    `json:"unkwn_cmd_code"`
	No_session_id             int    `json:"no_session_id"`
	No_fwd_tuple              int    `json:"no_fwd_tuple"`
	No_rev_tuple              int    `json:"no_rev_tuple"`
	Dcmsg_error               int    `json:"dcmsg_error"`
	Retry_client_request_fail int    `json:"retry_client_request_fail"`
	Reply_unknown_session_id  int    `json:"reply_unknown_session_id"`
	Client_select_fail        int    `json:"client_select_fail"`
	Invalid_avp               int    `json:"invalid_avp"`
	Reply_error_info_fail     int    `json:"reply_error_info_fail"`
	Uuid                      string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbGenericTriggerStatsRate3519 struct {
	ThresholdExceededBy       int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                  int    `json:"duration" dval:"60"`
	Svrsel_fail               int    `json:"svrsel_fail"`
	No_route                  int    `json:"no_route"`
	Snat_fail                 int    `json:"snat_fail"`
	Client_fail               int    `json:"client_fail"`
	Server_fail               int    `json:"server_fail"`
	Mismatch_fwd_id           int    `json:"mismatch_fwd_id"`
	Mismatch_rev_id           int    `json:"mismatch_rev_id"`
	Unkwn_cmd_code            int    `json:"unkwn_cmd_code"`
	No_session_id             int    `json:"no_session_id"`
	No_fwd_tuple              int    `json:"no_fwd_tuple"`
	No_rev_tuple              int    `json:"no_rev_tuple"`
	Dcmsg_error               int    `json:"dcmsg_error"`
	Retry_client_request_fail int    `json:"retry_client_request_fail"`
	Reply_unknown_session_id  int    `json:"reply_unknown_session_id"`
	Client_select_fail        int    `json:"client_select_fail"`
	Invalid_avp               int    `json:"invalid_avp"`
	Reply_error_info_fail     int    `json:"reply_error_info_fail"`
	Uuid                      string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersist3520 struct {
	Uuid             string                                                                                                   `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsInc3521  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsRate3522 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsInc3521 struct {
	Hash_tbl_trylock_fail      int    `json:"hash_tbl_trylock_fail"`
	Hash_tbl_create_fail       int    `json:"hash_tbl_create_fail"`
	Hash_tbl_rst_updown        int    `json:"hash_tbl_rst_updown"`
	Hash_tbl_rst_adddel        int    `json:"hash_tbl_rst_adddel"`
	Url_hash_fail              int    `json:"url_hash_fail"`
	Header_hash_fail           int    `json:"header_hash_fail"`
	Src_ip_fail                int    `json:"src_ip_fail"`
	Src_ip_new_sess_cache_fail int    `json:"src_ip_new_sess_cache_fail"`
	Src_ip_new_sess_sel_fail   int    `json:"src_ip_new_sess_sel_fail"`
	Src_ip_hash_fail           int    `json:"src_ip_hash_fail"`
	Dst_ip_fail                int    `json:"dst_ip_fail"`
	Dst_ip_new_sess_cache_fail int    `json:"dst_ip_new_sess_cache_fail"`
	Dst_ip_new_sess_sel_fail   int    `json:"dst_ip_new_sess_sel_fail"`
	Dst_ip_hash_fail           int    `json:"dst_ip_hash_fail"`
	Cssl_sid_not_found         int    `json:"cssl_sid_not_found"`
	Cssl_sid_not_match         int    `json:"cssl_sid_not_match"`
	Sssl_sid_not_found         int    `json:"sssl_sid_not_found"`
	Sssl_sid_not_match         int    `json:"sssl_sid_not_match"`
	Ssl_sid_persist_fail       int    `json:"ssl_sid_persist_fail"`
	Ssl_sid_session_fail       int    `json:"ssl_sid_session_fail"`
	Cookie_persist_fail        int    `json:"cookie_persist_fail"`
	Cookie_not_found           int    `json:"cookie_not_found"`
	Cookie_invalid             int    `json:"cookie_invalid"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPersistTriggerStatsRate3522 struct {
	ThresholdExceededBy        int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                   int    `json:"duration" dval:"60"`
	Hash_tbl_trylock_fail      int    `json:"hash_tbl_trylock_fail"`
	Hash_tbl_create_fail       int    `json:"hash_tbl_create_fail"`
	Hash_tbl_rst_updown        int    `json:"hash_tbl_rst_updown"`
	Hash_tbl_rst_adddel        int    `json:"hash_tbl_rst_adddel"`
	Url_hash_fail              int    `json:"url_hash_fail"`
	Header_hash_fail           int    `json:"header_hash_fail"`
	Src_ip_fail                int    `json:"src_ip_fail"`
	Src_ip_new_sess_cache_fail int    `json:"src_ip_new_sess_cache_fail"`
	Src_ip_new_sess_sel_fail   int    `json:"src_ip_new_sess_sel_fail"`
	Src_ip_hash_fail           int    `json:"src_ip_hash_fail"`
	Dst_ip_fail                int    `json:"dst_ip_fail"`
	Dst_ip_new_sess_cache_fail int    `json:"dst_ip_new_sess_cache_fail"`
	Dst_ip_new_sess_sel_fail   int    `json:"dst_ip_new_sess_sel_fail"`
	Dst_ip_hash_fail           int    `json:"dst_ip_hash_fail"`
	Cssl_sid_not_found         int    `json:"cssl_sid_not_found"`
	Cssl_sid_not_match         int    `json:"cssl_sid_not_match"`
	Sssl_sid_not_found         int    `json:"sssl_sid_not_found"`
	Sssl_sid_not_match         int    `json:"sssl_sid_not_match"`
	Ssl_sid_persist_fail       int    `json:"ssl_sid_persist_fail"`
	Ssl_sid_session_fail       int    `json:"ssl_sid_session_fail"`
	Cookie_persist_fail        int    `json:"cookie_persist_fail"`
	Cookie_not_found           int    `json:"cookie_not_found"`
	Cookie_invalid             int    `json:"cookie_invalid"`
	Uuid                       string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxy3523 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc3524  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate3525 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsInc3524 struct {
	Parsereq_fail       int    `json:"parsereq_fail"`
	Svrsel_fail         int    `json:"svrsel_fail"`
	Fwdreq_fail         int    `json:"fwdreq_fail"`
	Fwdreqdata_fail     int    `json:"fwdreqdata_fail"`
	Snat_fail           int    `json:"snat_fail"`
	Req_over_limit      int    `json:"req_over_limit"`
	Req_rate_over_limit int    `json:"req_rate_over_limit"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbHttpProxyTriggerStatsRate3525 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Parsereq_fail       int    `json:"parsereq_fail"`
	Svrsel_fail         int    `json:"svrsel_fail"`
	Fwdreq_fail         int    `json:"fwdreq_fail"`
	Fwdreqdata_fail     int    `json:"fwdreqdata_fail"`
	Snat_fail           int    `json:"snat_fail"`
	Req_over_limit      int    `json:"req_over_limit"`
	Req_rate_over_limit int    `json:"req_rate_over_limit"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL43526 struct {
	Uuid             string                                                                                              `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsInc3527  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsRate3528 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsInc3527 struct {
	Syncookiessentfailed    int    `json:"syncookiessentfailed"`
	Svrselfail              int    `json:"svrselfail"`
	Snat_fail               int    `json:"snat_fail"`
	Snat_no_fwd_route       int    `json:"snat_no_fwd_route"`
	Snat_no_rev_route       int    `json:"snat_no_rev_route"`
	Snat_icmp_error_process int    `json:"snat_icmp_error_process"`
	Snat_icmp_no_match      int    `json:"snat_icmp_no_match"`
	Smart_nat_id_mismatch   int    `json:"smart_nat_id_mismatch"`
	Syncookiescheckfailed   int    `json:"syncookiescheckfailed"`
	Connlimit_drop          int    `json:"connlimit_drop"`
	Conn_rate_limit_drop    int    `json:"conn_rate_limit_drop"`
	Conn_rate_limit_reset   int    `json:"conn_rate_limit_reset"`
	Dns_policy_drop         int    `json:"dns_policy_drop"`
	No_resourse_drop        int    `json:"no_resourse_drop"`
	Bw_rate_limit_exceed    int    `json:"bw_rate_limit_exceed"`
	L4_cps_exceed           int    `json:"l4_cps_exceed"`
	Nat_cps_exceed          int    `json:"nat_cps_exceed"`
	L7_cps_exceed           int    `json:"l7_cps_exceed"`
	Ssl_cps_exceed          int    `json:"ssl_cps_exceed"`
	Ssl_tpt_exceed          int    `json:"ssl_tpt_exceed"`
	Concurrent_conn_exceed  int    `json:"concurrent_conn_exceed"`
	Svr_syn_handshake_fail  int    `json:"svr_syn_handshake_fail"`
	Synattack               int    `json:"synattack"`
	Uuid                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbL4TriggerStatsRate3528 struct {
	ThresholdExceededBy     int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                int    `json:"duration" dval:"60"`
	Syncookiessentfailed    int    `json:"syncookiessentfailed"`
	Svrselfail              int    `json:"svrselfail"`
	Snat_fail               int    `json:"snat_fail"`
	Snat_no_fwd_route       int    `json:"snat_no_fwd_route"`
	Snat_no_rev_route       int    `json:"snat_no_rev_route"`
	Snat_icmp_error_process int    `json:"snat_icmp_error_process"`
	Snat_icmp_no_match      int    `json:"snat_icmp_no_match"`
	Smart_nat_id_mismatch   int    `json:"smart_nat_id_mismatch"`
	Syncookiescheckfailed   int    `json:"syncookiescheckfailed"`
	Connlimit_drop          int    `json:"connlimit_drop"`
	Conn_rate_limit_drop    int    `json:"conn_rate_limit_drop"`
	Conn_rate_limit_reset   int    `json:"conn_rate_limit_reset"`
	Dns_policy_drop         int    `json:"dns_policy_drop"`
	No_resourse_drop        int    `json:"no_resourse_drop"`
	Bw_rate_limit_exceed    int    `json:"bw_rate_limit_exceed"`
	L4_cps_exceed           int    `json:"l4_cps_exceed"`
	Nat_cps_exceed          int    `json:"nat_cps_exceed"`
	L7_cps_exceed           int    `json:"l7_cps_exceed"`
	Ssl_cps_exceed          int    `json:"ssl_cps_exceed"`
	Ssl_tpt_exceed          int    `json:"ssl_tpt_exceed"`
	Concurrent_conn_exceed  int    `json:"concurrent_conn_exceed"`
	Svr_syn_handshake_fail  int    `json:"svr_syn_handshake_fail"`
	Synattack               int    `json:"synattack"`
	Uuid                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttp3529 struct {
	Uuid             string                                                                                                    `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc3530  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate3531 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsInc3530 struct {
	Parsereq_fail       int    `json:"parsereq_fail"`
	Svrsel_fail         int    `json:"svrsel_fail"`
	Fwdreq_fail         int    `json:"fwdreq_fail"`
	Fwdreqdata_fail     int    `json:"fwdreqdata_fail"`
	Snat_fail           int    `json:"snat_fail"`
	Full_proxy_fpga_err int    `json:"full_proxy_fpga_err"`
	Req_over_limit      int    `json:"req_over_limit"`
	Req_rate_over_limit int    `json:"req_rate_over_limit"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFastHttpTriggerStatsRate3531 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Parsereq_fail       int    `json:"parsereq_fail"`
	Svrsel_fail         int    `json:"svrsel_fail"`
	Fwdreq_fail         int    `json:"fwdreq_fail"`
	Fwdreqdata_fail     int    `json:"fwdreqdata_fail"`
	Snat_fail           int    `json:"snat_fail"`
	Full_proxy_fpga_err int    `json:"full_proxy_fpga_err"`
	Req_over_limit      int    `json:"req_over_limit"`
	Req_rate_over_limit int    `json:"req_rate_over_limit"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxy3532 struct {
	Uuid             string                                                                                                    `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc3533  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate3534 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsInc3533 struct {
	Svrsel_fail              int    `json:"svrsel_fail"`
	No_route                 int    `json:"no_route"`
	Snat_fail                int    `json:"snat_fail"`
	Line_too_long            int    `json:"line_too_long"`
	Invalid_start_line       int    `json:"invalid_start_line"`
	Smp_create_fail          int    `json:"smp_create_fail"`
	Data_server_conn_fail    int    `json:"data_server_conn_fail"`
	Data_send_fail           int    `json:"data_send_fail"`
	Unsupported_pbsz_value   int    `json:"unsupported_pbsz_value"`
	Unsupported_prot_value   int    `json:"unsupported_prot_value"`
	Unsupported_command      int    `json:"unsupported_command"`
	Bad_sequence             int    `json:"bad_sequence"`
	Rsv_persist_conn_fail    int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail              int    `json:"smp_v6_fail"`
	Smp_v4_fail              int    `json:"smp_v4_fail"`
	Insert_tuple_fail        int    `json:"insert_tuple_fail"`
	Cl_est_err               int    `json:"cl_est_err"`
	Ser_connecting_err       int    `json:"ser_connecting_err"`
	Server_response_err      int    `json:"server_response_err"`
	Cl_request_err           int    `json:"cl_request_err"`
	Data_conn_start_err      int    `json:"data_conn_start_err"`
	Data_serv_connecting_err int    `json:"data_serv_connecting_err"`
	Data_serv_connected_err  int    `json:"data_serv_connected_err"`
	Auth_fail                int    `json:"auth_fail"`
	Ds_fail                  int    `json:"ds_fail"`
	Cant_find_port           int    `json:"cant_find_port"`
	Cant_find_eprt           int    `json:"cant_find_eprt"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbFtpProxyTriggerStatsRate3534 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	Svrsel_fail              int    `json:"svrsel_fail"`
	No_route                 int    `json:"no_route"`
	Snat_fail                int    `json:"snat_fail"`
	Line_too_long            int    `json:"line_too_long"`
	Invalid_start_line       int    `json:"invalid_start_line"`
	Smp_create_fail          int    `json:"smp_create_fail"`
	Data_server_conn_fail    int    `json:"data_server_conn_fail"`
	Data_send_fail           int    `json:"data_send_fail"`
	Unsupported_pbsz_value   int    `json:"unsupported_pbsz_value"`
	Unsupported_prot_value   int    `json:"unsupported_prot_value"`
	Unsupported_command      int    `json:"unsupported_command"`
	Bad_sequence             int    `json:"bad_sequence"`
	Rsv_persist_conn_fail    int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail              int    `json:"smp_v6_fail"`
	Smp_v4_fail              int    `json:"smp_v4_fail"`
	Insert_tuple_fail        int    `json:"insert_tuple_fail"`
	Cl_est_err               int    `json:"cl_est_err"`
	Ser_connecting_err       int    `json:"ser_connecting_err"`
	Server_response_err      int    `json:"server_response_err"`
	Cl_request_err           int    `json:"cl_request_err"`
	Data_conn_start_err      int    `json:"data_conn_start_err"`
	Data_serv_connecting_err int    `json:"data_serv_connecting_err"`
	Data_serv_connected_err  int    `json:"data_serv_connected_err"`
	Auth_fail                int    `json:"auth_fail"`
	Ds_fail                  int    `json:"ds_fail"`
	Cant_find_port           int    `json:"cant_find_port"`
	Cant_find_eprt           int    `json:"cant_find_eprt"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxy3535 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc3536  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate3537 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsInc3536 struct {
	Svrsel_fail            int    `json:"svrsel_fail"`
	No_route               int    `json:"no_route"`
	Snat_fail              int    `json:"snat_fail"`
	Line_too_long          int    `json:"line_too_long"`
	Invalid_start_line     int    `json:"invalid_start_line"`
	Cant_find_pasv         int    `json:"cant_find_pasv"`
	Smp_create_fail        int    `json:"smp_create_fail"`
	Data_server_conn_fail  int    `json:"data_server_conn_fail"`
	Data_send_fail         int    `json:"data_send_fail"`
	Cant_find_epsv         int    `json:"cant_find_epsv"`
	Auth_unsupported       int    `json:"auth_unsupported"`
	Unsupported_pbsz_value int    `json:"unsupported_pbsz_value"`
	Unsupported_prot_value int    `json:"unsupported_prot_value"`
	Bad_sequence           int    `json:"bad_sequence"`
	Rsv_persist_conn_fail  int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail            int    `json:"smp_v6_fail"`
	Smp_v4_fail            int    `json:"smp_v4_fail"`
	Insert_tuple_fail      int    `json:"insert_tuple_fail"`
	Cl_est_err             int    `json:"cl_est_err"`
	Ser_connecting_err     int    `json:"ser_connecting_err"`
	Server_response_err    int    `json:"server_response_err"`
	Cl_request_err         int    `json:"cl_request_err"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbImapProxyTriggerStatsRate3537 struct {
	ThresholdExceededBy    int    `json:"threshold-exceeded-by" dval:"5"`
	Duration               int    `json:"duration" dval:"60"`
	Svrsel_fail            int    `json:"svrsel_fail"`
	No_route               int    `json:"no_route"`
	Snat_fail              int    `json:"snat_fail"`
	Line_too_long          int    `json:"line_too_long"`
	Invalid_start_line     int    `json:"invalid_start_line"`
	Cant_find_pasv         int    `json:"cant_find_pasv"`
	Smp_create_fail        int    `json:"smp_create_fail"`
	Data_server_conn_fail  int    `json:"data_server_conn_fail"`
	Data_send_fail         int    `json:"data_send_fail"`
	Cant_find_epsv         int    `json:"cant_find_epsv"`
	Auth_unsupported       int    `json:"auth_unsupported"`
	Unsupported_pbsz_value int    `json:"unsupported_pbsz_value"`
	Unsupported_prot_value int    `json:"unsupported_prot_value"`
	Bad_sequence           int    `json:"bad_sequence"`
	Rsv_persist_conn_fail  int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail            int    `json:"smp_v6_fail"`
	Smp_v4_fail            int    `json:"smp_v4_fail"`
	Insert_tuple_fail      int    `json:"insert_tuple_fail"`
	Cl_est_err             int    `json:"cl_est_err"`
	Ser_connecting_err     int    `json:"ser_connecting_err"`
	Server_response_err    int    `json:"server_response_err"`
	Cl_request_err         int    `json:"cl_request_err"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3Proxy3538 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc3539  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate3540 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsInc3539 struct {
	Svrsel_fail           int    `json:"svrsel_fail"`
	No_route              int    `json:"no_route"`
	Snat_fail             int    `json:"snat_fail"`
	Line_too_long         int    `json:"line_too_long"`
	Invalid_start_line    int    `json:"invalid_start_line"`
	Unsupported_command   int    `json:"unsupported_command"`
	Bad_sequence          int    `json:"bad_sequence"`
	Rsv_persist_conn_fail int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail           int    `json:"smp_v6_fail"`
	Smp_v4_fail           int    `json:"smp_v4_fail"`
	Insert_tuple_fail     int    `json:"insert_tuple_fail"`
	Cl_est_err            int    `json:"cl_est_err"`
	Ser_connecting_err    int    `json:"ser_connecting_err"`
	Server_response_err   int    `json:"server_response_err"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPop3ProxyTriggerStatsRate3540 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	Svrsel_fail           int    `json:"svrsel_fail"`
	No_route              int    `json:"no_route"`
	Snat_fail             int    `json:"snat_fail"`
	Line_too_long         int    `json:"line_too_long"`
	Invalid_start_line    int    `json:"invalid_start_line"`
	Unsupported_command   int    `json:"unsupported_command"`
	Bad_sequence          int    `json:"bad_sequence"`
	Rsv_persist_conn_fail int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail           int    `json:"smp_v6_fail"`
	Smp_v4_fail           int    `json:"smp_v4_fail"`
	Insert_tuple_fail     int    `json:"insert_tuple_fail"`
	Cl_est_err            int    `json:"cl_est_err"`
	Ser_connecting_err    int    `json:"ser_connecting_err"`
	Server_response_err   int    `json:"server_response_err"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitch3541 struct {
	Uuid             string                                                                                                  `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc3542  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate3543 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsInc3542 struct {
	Lacp_tx_intf_err_drop       int    `json:"lacp_tx_intf_err_drop"`
	Unnumbered_nat_error        int    `json:"unnumbered_nat_error"`
	Unnumbered_unsupported_drop int    `json:"unnumbered_unsupported_drop"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSwitchTriggerStatsRate3543 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	Lacp_tx_intf_err_drop       int    `json:"lacp_tx_intf_err_drop"`
	Unnumbered_nat_error        int    `json:"unnumbered_nat_error"`
	Unnumbered_unsupported_drop int    `json:"unnumbered_unsupported_drop"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCache3544 struct {
	Uuid             string                                                                                                   `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc3545  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate3546 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsInc3545 struct {
	Rv_failure            int    `json:"rv_failure"`
	Content_toobig        int    `json:"content_toobig"`
	Content_toosmall      int    `json:"content_toosmall"`
	Entry_create_failures int    `json:"entry_create_failures"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRcCacheTriggerStatsRate3546 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	Rv_failure            int    `json:"rv_failure"`
	Content_toobig        int    `json:"content_toobig"`
	Content_toosmall      int    `json:"content_toosmall"`
	Entry_create_failures int    `json:"entry_create_failures"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCounters3547 struct {
	Uuid             string                                                                                                   `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsInc3548  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsRate3549 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsInc3548 struct {
	So_pkts_slb_nat_reserve_fail            int    `json:"so_pkts_slb_nat_reserve_fail"`
	So_pkts_slb_nat_release_fail            int    `json:"so_pkts_slb_nat_release_fail"`
	So_pkts_l2redirect_dest_mac_zero_drop   int    `json:"so_pkts_l2redirect_dest_mac_zero_drop"`
	So_pkts_l2redirect_interface_not_up     int    `json:"so_pkts_l2redirect_interface_not_up"`
	So_pkts_l2redirect_invalid_redirect_inf int    `json:"so_pkts_l2redirect_invalid_redirect_inf"`
	So_pkts_l3_redirect_encap_error_drop    int    `json:"so_pkts_l3_redirect_encap_error_drop"`
	So_pkts_l3_redirect_inner_mac_zero_drop int    `json:"so_pkts_l3_redirect_inner_mac_zero_drop"`
	So_pkts_l3_redirect_table_error         int    `json:"so_pkts_l3_redirect_table_error"`
	So_pkts_l3_redirect_fragmentation_error int    `json:"so_pkts_l3_redirect_fragmentation_error"`
	So_pkts_l3_redirect_table_no_entry_foun int    `json:"so_pkts_l3_redirect_table_no_entry_foun"`
	So_pkts_l3_redirect_invalid_dev_dir     int    `json:"so_pkts_l3_redirect_invalid_dev_dir"`
	So_pkts_l3_redirect_chassis_dest_mac_er int    `json:"so_pkts_l3_redirect_chassis_dest_mac_er"`
	So_pkts_l2redirect_vlan_retrieval_error int    `json:"so_pkts_l2redirect_vlan_retrieval_error"`
	So_pkts_l2redirect_port_retrieval_error int    `json:"so_pkts_l2redirect_port_retrieval_error"`
	So_pkts_l2redirect_loop_detect_drop     int    `json:"so_pkts_l2redirect_loop_detect_drop"`
	So_pkts_l2redirect_frag_vlan_retrieval_ int    `json:"so_pkts_l2redirect_frag_vlan_retrieval_"`
	So_pkts_l2redirect_tx_frag_vlan_add_fai int    `json:"so_pkts_l2redirect_tx_frag_vlan_add_fai"`
	Uuid                                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSoCountersTriggerStatsRate3549 struct {
	ThresholdExceededBy                     int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                                int    `json:"duration" dval:"60"`
	So_pkts_slb_nat_reserve_fail            int    `json:"so_pkts_slb_nat_reserve_fail"`
	So_pkts_slb_nat_release_fail            int    `json:"so_pkts_slb_nat_release_fail"`
	So_pkts_l2redirect_dest_mac_zero_drop   int    `json:"so_pkts_l2redirect_dest_mac_zero_drop"`
	So_pkts_l2redirect_interface_not_up     int    `json:"so_pkts_l2redirect_interface_not_up"`
	So_pkts_l2redirect_invalid_redirect_inf int    `json:"so_pkts_l2redirect_invalid_redirect_inf"`
	So_pkts_l3_redirect_encap_error_drop    int    `json:"so_pkts_l3_redirect_encap_error_drop"`
	So_pkts_l3_redirect_inner_mac_zero_drop int    `json:"so_pkts_l3_redirect_inner_mac_zero_drop"`
	So_pkts_l3_redirect_table_error         int    `json:"so_pkts_l3_redirect_table_error"`
	So_pkts_l3_redirect_fragmentation_error int    `json:"so_pkts_l3_redirect_fragmentation_error"`
	So_pkts_l3_redirect_table_no_entry_foun int    `json:"so_pkts_l3_redirect_table_no_entry_foun"`
	So_pkts_l3_redirect_invalid_dev_dir     int    `json:"so_pkts_l3_redirect_invalid_dev_dir"`
	So_pkts_l3_redirect_chassis_dest_mac_er int    `json:"so_pkts_l3_redirect_chassis_dest_mac_er"`
	So_pkts_l2redirect_vlan_retrieval_error int    `json:"so_pkts_l2redirect_vlan_retrieval_error"`
	So_pkts_l2redirect_port_retrieval_error int    `json:"so_pkts_l2redirect_port_retrieval_error"`
	So_pkts_l2redirect_loop_detect_drop     int    `json:"so_pkts_l2redirect_loop_detect_drop"`
	So_pkts_l2redirect_frag_vlan_retrieval_ int    `json:"so_pkts_l2redirect_frag_vlan_retrieval_"`
	So_pkts_l2redirect_tx_frag_vlan_add_fai int    `json:"so_pkts_l2redirect_tx_frag_vlan_add_fai"`
	Uuid                                    string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGbl3550 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc3551  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate3552 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsInc3551 struct {
	Total_invalid_playerid_pkts int    `json:"total_invalid_playerid_pkts"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbPlyrIdGblTriggerStatsRate3552 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	Total_invalid_playerid_pkts int    `json:"total_invalid_playerid_pkts"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRate3553 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc3554  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate3555 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsInc3554 struct {
	Total_reset int    `json:"total_reset"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbSportRateTriggerStatsRate3555 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Total_reset         int    `json:"total_reset"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobal3556 struct {
	Uuid             string                                                                                                              `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc3557  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate3558 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsInc3557 struct {
	EnqueueFull  int    `json:"enqueue-full"`
	EnqueueError int    `json:"enqueue-error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeLoggingLocalLogGlobalTriggerStatsRate3558 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	EnqueueFull         int    `json:"enqueue-full"`
	EnqueueError        int    `json:"enqueue-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlb3559 struct {
	Uuid             string                                                                                               `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsInc3560  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsRate3561 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsInc3560 struct {
	Mlb_dcmsg_error int    `json:"mlb_dcmsg_error"`
	Uuid            string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbMlbTriggerStatsRate3561 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Mlb_dcmsg_error     int    `json:"mlb_dcmsg_error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbe3562 struct {
	Uuid             string                                                                                                     `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc3563  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate3564 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsInc3563 struct {
	Err_entry_create_failed      int    `json:"err_entry_create_failed"`
	Err_entry_create_oom         int    `json:"err_entry_create_oom"`
	Err_entry_insert_failed      int    `json:"err_entry_insert_failed"`
	Err_tmpl_probe_create_failed int    `json:"err_tmpl_probe_create_failed"`
	Err_tmpl_probe_create_oom    int    `json:"err_tmpl_probe_create_oom"`
	Total_http_response_bad      int    `json:"total_http_response_bad"`
	Total_tcp_err                int    `json:"total_tcp_err"`
	Err_smart_nat_alloc          int    `json:"err_smart_nat_alloc"`
	Err_smart_nat_port_alloc     int    `json:"err_smart_nat_port_alloc"`
	Err_l4_sess_alloc            int    `json:"err_l4_sess_alloc"`
	Err_probe_tcp_conn_send      int    `json:"err_probe_tcp_conn_send"`
	Uuid                         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbLinkProbeTriggerStatsRate3564 struct {
	ThresholdExceededBy          int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                     int    `json:"duration" dval:"60"`
	Err_entry_create_failed      int    `json:"err_entry_create_failed"`
	Err_entry_create_oom         int    `json:"err_entry_create_oom"`
	Err_entry_insert_failed      int    `json:"err_entry_insert_failed"`
	Err_tmpl_probe_create_failed int    `json:"err_tmpl_probe_create_failed"`
	Err_tmpl_probe_create_oom    int    `json:"err_tmpl_probe_create_oom"`
	Total_http_response_bad      int    `json:"total_http_response_bad"`
	Total_tcp_err                int    `json:"total_tcp_err"`
	Err_smart_nat_alloc          int    `json:"err_smart_nat_alloc"`
	Err_smart_nat_port_alloc     int    `json:"err_smart_nat_port_alloc"`
	Err_l4_sess_alloc            int    `json:"err_l4_sess_alloc"`
	Err_probe_tcp_conn_send      int    `json:"err_probe_tcp_conn_send"`
	Uuid                         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpz3565 struct {
	Uuid             string                                                                                               `json:"uuid"`
	TriggerStatsInc  VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsInc3566  `json:"trigger-stats-inc"`
	TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsRate3567 `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsInc3566 struct {
	Set_bw_error int    `json:"set_bw_error"`
	Parse_error  int    `json:"parse_error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesTemplateListTriggerSysObjStatsChangeSlbRpzTriggerStatsRate3567 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Set_bw_error        int    `json:"set_bw_error"`
	Parse_error         int    `json:"parse_error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureGlobalTemplatesActivate3568 struct {
	Template string `json:"template"`
	Uuid     string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplates3569 struct {
	Uuid                           string                                                                     `json:"uuid"`
	TemplGtpPlcyTmplList           []VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList3570           `json:"templ-gtp-plcy-tmpl-list"`
	InterfaceEthernetTmplList      []VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList3574      `json:"interface-ethernet-tmpl-list"`
	InterfaceTunnelTmplList        []VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList3578        `json:"interface-tunnel-tmpl-list"`
	AamJwtAuthorizationTmplList    []VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList3582    `json:"aam-jwt-authorization-tmpl-list"`
	AamAaaPolicyTmplList           []VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList3586           `json:"aam-aaa-policy-tmpl-list"`
	AamAuthLogonHttpInsTmplList    []VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList3590    `json:"aam-auth-logon-http-ins-tmpl-list"`
	AamAuthServerLdapInstTmplList  []VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList3594  `json:"aam-auth-server-ldap-inst-tmpl-list"`
	AamAuthServerOcspInstTmplList  []VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList3598  `json:"aam-auth-server-ocsp-inst-tmpl-list"`
	AamAuthServerRadInstTmplList   []VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList3602   `json:"aam-auth-server-rad-inst-tmpl-list"`
	AamAuthServerWinInstTmplList   []VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList3606   `json:"aam-auth-server-win-inst-tmpl-list"`
	AamAuthSamlServiceProvTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList3610 `json:"aam-auth-saml-service-prov-tmpl-list"`
	AamAuthSamlIdProvTmplList      []VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList3614      `json:"aam-auth-saml-id-prov-tmpl-list"`
	AamAuthServiceGroupTmplList    []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList3618    `json:"aam-auth-service-group-tmpl-list"`
	AamAuthServiceGroupMemTmplList []VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList3622 `json:"aam-auth-service-group-mem-tmpl-list"`
	AamAuthRelayHbaseInstTmplList  []VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList3626  `json:"aam-auth-relay-hbase-inst-tmpl-list"`
	AamAuthRelayFormInstTmplList   []VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList3630   `json:"aam-auth-relay-form-inst-tmpl-list"`
	AamAuthRelayWsFedTmplList      []VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList3634      `json:"aam-auth-relay-ws-fed-tmpl-list"`
	AamAuthCaptchaInstTmplList     []VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList3638     `json:"aam-auth-captcha-inst-tmpl-list"`
	SlbTemplCacheTmplList          []VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList3642          `json:"slb-templ-cache-tmpl-list"`
	CaptchaTemplateInstTmplList    []VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplList3646    `json:"captcha-template-inst-tmpl-list"`
	SlbPortTmplList                []VisibilityPacketCaptureObjectTemplatesSlbPortTmplList3650                `json:"slb-port-tmpl-list"`
	SlbServiceTmplList             []VisibilityPacketCaptureObjectTemplatesSlbServiceTmplList3654             `json:"slb-service-tmpl-list"`
	SlbVportTmplList               []VisibilityPacketCaptureObjectTemplatesSlbVportTmplList3658               `json:"slb-vport-tmpl-list"`
	Cgnv6ServGroupTmplList         []VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList3662         `json:"cgnv6-serv-group-tmpl-list"`
	Cgnv6Dns64VsPortTmplList       []VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList3666       `json:"cgnv6-dns64-vs-port-tmpl-list"`
	Cgnv6MapTransDomainTmplList    []VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList3670    `json:"cgnv6-map-trans-domain-tmpl-list"`
	Cgnv6EncapDomainTmplList       []VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList3674       `json:"cgnv6-encap-domain-tmpl-list"`
	NetflowMonitorTmplList         []VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList3678         `json:"netflow-monitor-tmpl-list"`
	RuleSetTmplList                []VisibilityPacketCaptureObjectTemplatesRuleSetTmplList3682                `json:"rule-set-tmpl-list"`
	RuleSetRuleTmplList            []VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplList3686            `json:"rule-set-rule-tmpl-list"`
	FwServerPortTmplList           []VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList3690           `json:"fw-server-port-tmpl-list"`
	FwServiceGroupTmplList         []VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList3694         `json:"fw-service-group-tmpl-list"`
	FwServiceGroupMemTmplList      []VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList3698      `json:"fw-service-group-mem-tmpl-list"`
	Dns_vportTmplList              []VisibilityPacketCaptureObjectTemplatesDns_vportTmplList3702              `json:"dns_vport-tmpl-list"`
	SmtpVportTmplList              []VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList3706              `json:"smtp-vport-tmpl-list"`
	Pop3VportTmplList              []VisibilityPacketCaptureObjectTemplatesPop3VportTmplList3710              `json:"pop3-vport-tmpl-list"`
	ImapVportTmplList              []VisibilityPacketCaptureObjectTemplatesImapVportTmplList3714              `json:"imap-vport-tmpl-list"`
}

type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplList3570 struct {
	Name                 string                                                                             `json:"name"`
	CaptureConfig        string                                                                             `json:"capture-config"`
	Uuid                 string                                                                             `json:"uuid"`
	UserTag              string                                                                             `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity3571 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc3572      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate3573     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsSeverity3571 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsInc3572 struct {
	DropVldGtpIeRepeatCountExceed      int    `json:"drop-vld-gtp-ie-repeat-count-exceed"`
	DropVldReservedFieldSet            int    `json:"drop-vld-reserved-field-set"`
	DropVldTunnelIdFlag                int    `json:"drop-vld-tunnel-id-flag"`
	DropVldInvalidFlowLabelV0          int    `json:"drop-vld-invalid-flow-label-v0"`
	DropVldInvalidTeid                 int    `json:"drop-vld-invalid-teid"`
	DropVldOutOfState                  int    `json:"drop-vld-out-of-state"`
	DropVldMandatoryInformationElement int    `json:"drop-vld-mandatory-information-element"`
	DropVldMandatoryIeInGroupedIe      int    `json:"drop-vld-mandatory-ie-in-grouped-ie"`
	DropVldOutOfOrderIe                int    `json:"drop-vld-out-of-order-ie"`
	DropVldOutOfStateIe                int    `json:"drop-vld-out-of-state-ie"`
	DropVldReservedInformationElement  int    `json:"drop-vld-reserved-information-element"`
	DropVldVersionNotSupported         int    `json:"drop-vld-version-not-supported"`
	DropVldMessageLength               int    `json:"drop-vld-message-length"`
	DropVldCrossLayerCorrelation       int    `json:"drop-vld-cross-layer-correlation"`
	DropVldCountryCodeMismatch         int    `json:"drop-vld-country-code-mismatch"`
	DropVldGtpUSpoofedSourceAddress    int    `json:"drop-vld-gtp-u-spoofed-source-address"`
	DropVldGtpBearerCountExceed        int    `json:"drop-vld-gtp-bearer-count-exceed"`
	DropVldGtpV2WrongLbiCreateBearer   int    `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
	DropVldV0ReservedMessageDrop       int    `json:"drop-vld-v0-reserved-message-drop"`
	DropVldV1ReservedMessageDrop       int    `json:"drop-vld-v1-reserved-message-drop"`
	DropVldV2ReservedMessageDrop       int    `json:"drop-vld-v2-reserved-message-drop"`
	DropVldInvalidPktLenPiggyback      int    `json:"drop-vld-invalid-pkt-len-piggyback"`
	DropVldSanityFailedPiggyback       int    `json:"drop-vld-sanity-failed-piggyback"`
	DropVldSequenceNumCorrelation      int    `json:"drop-vld-sequence-num-correlation"`
	DropVldGtpv0SeqnumBufferFull       int    `json:"drop-vld-gtpv0-seqnum-buffer-full"`
	DropVldGtpv1SeqnumBufferFull       int    `json:"drop-vld-gtpv1-seqnum-buffer-full"`
	DropVldGtpv2SeqnumBufferFull       int    `json:"drop-vld-gtpv2-seqnum-buffer-full"`
	DropVldGtpInvalidImsiLenDrop       int    `json:"drop-vld-gtp-invalid-imsi-len-drop"`
	DropVldGtpInvalidApnLenDrop        int    `json:"drop-vld-gtp-invalid-apn-len-drop"`
	DropVldProtocolFlagUnset           int    `json:"drop-vld-protocol-flag-unset"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplListTriggerStatsRate3573 struct {
	ThresholdExceededBy                int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                           int    `json:"duration" dval:"60"`
	DropVldGtpIeRepeatCountExceed      int    `json:"drop-vld-gtp-ie-repeat-count-exceed"`
	DropVldReservedFieldSet            int    `json:"drop-vld-reserved-field-set"`
	DropVldTunnelIdFlag                int    `json:"drop-vld-tunnel-id-flag"`
	DropVldInvalidFlowLabelV0          int    `json:"drop-vld-invalid-flow-label-v0"`
	DropVldInvalidTeid                 int    `json:"drop-vld-invalid-teid"`
	DropVldOutOfState                  int    `json:"drop-vld-out-of-state"`
	DropVldMandatoryInformationElement int    `json:"drop-vld-mandatory-information-element"`
	DropVldMandatoryIeInGroupedIe      int    `json:"drop-vld-mandatory-ie-in-grouped-ie"`
	DropVldOutOfOrderIe                int    `json:"drop-vld-out-of-order-ie"`
	DropVldOutOfStateIe                int    `json:"drop-vld-out-of-state-ie"`
	DropVldReservedInformationElement  int    `json:"drop-vld-reserved-information-element"`
	DropVldVersionNotSupported         int    `json:"drop-vld-version-not-supported"`
	DropVldMessageLength               int    `json:"drop-vld-message-length"`
	DropVldCrossLayerCorrelation       int    `json:"drop-vld-cross-layer-correlation"`
	DropVldCountryCodeMismatch         int    `json:"drop-vld-country-code-mismatch"`
	DropVldGtpUSpoofedSourceAddress    int    `json:"drop-vld-gtp-u-spoofed-source-address"`
	DropVldGtpBearerCountExceed        int    `json:"drop-vld-gtp-bearer-count-exceed"`
	DropVldGtpV2WrongLbiCreateBearer   int    `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
	DropVldV0ReservedMessageDrop       int    `json:"drop-vld-v0-reserved-message-drop"`
	DropVldV1ReservedMessageDrop       int    `json:"drop-vld-v1-reserved-message-drop"`
	DropVldV2ReservedMessageDrop       int    `json:"drop-vld-v2-reserved-message-drop"`
	DropVldInvalidPktLenPiggyback      int    `json:"drop-vld-invalid-pkt-len-piggyback"`
	DropVldSanityFailedPiggyback       int    `json:"drop-vld-sanity-failed-piggyback"`
	DropVldSequenceNumCorrelation      int    `json:"drop-vld-sequence-num-correlation"`
	DropVldGtpv0SeqnumBufferFull       int    `json:"drop-vld-gtpv0-seqnum-buffer-full"`
	DropVldGtpv1SeqnumBufferFull       int    `json:"drop-vld-gtpv1-seqnum-buffer-full"`
	DropVldGtpv2SeqnumBufferFull       int    `json:"drop-vld-gtpv2-seqnum-buffer-full"`
	DropVldGtpInvalidImsiLenDrop       int    `json:"drop-vld-gtp-invalid-imsi-len-drop"`
	DropVldGtpInvalidApnLenDrop        int    `json:"drop-vld-gtp-invalid-apn-len-drop"`
	DropVldProtocolFlagUnset           int    `json:"drop-vld-protocol-flag-unset"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplList3574 struct {
	Name                 string                                                                                  `json:"name"`
	CaptureConfig        string                                                                                  `json:"capture-config"`
	Uuid                 string                                                                                  `json:"uuid"`
	UserTag              string                                                                                  `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity3575 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc3576      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate3577     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsSeverity3575 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsInc3576 struct {
	Input_errors  int    `json:"input_errors"`
	Crc           int    `json:"crc"`
	Runts         int    `json:"runts"`
	Giants        int    `json:"giants"`
	Output_errors int    `json:"output_errors"`
	Collisions    int    `json:"collisions"`
	Giants_output int    `json:"giants_output"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplListTriggerStatsRate3577 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Input_errors        int    `json:"input_errors"`
	Crc                 int    `json:"crc"`
	Runts               int    `json:"runts"`
	Giants              int    `json:"giants"`
	Output_errors       int    `json:"output_errors"`
	Collisions          int    `json:"collisions"`
	Giants_output       int    `json:"giants_output"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplList3578 struct {
	Name                 string                                                                                `json:"name"`
	CaptureConfig        string                                                                                `json:"capture-config"`
	Uuid                 string                                                                                `json:"uuid"`
	UserTag              string                                                                                `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity3579 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc3580      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate3581     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsSeverity3579 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsInc3580 struct {
	NumRxErrPkts int    `json:"num-rx-err-pkts"`
	NumTxErrPkts int    `json:"num-tx-err-pkts"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplListTriggerStatsRate3581 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	NumRxErrPkts        int    `json:"num-rx-err-pkts"`
	NumTxErrPkts        int    `json:"num-tx-err-pkts"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplList3582 struct {
	Name                 string                                                                                    `json:"name"`
	CaptureConfig        string                                                                                    `json:"capture-config"`
	Uuid                 string                                                                                    `json:"uuid"`
	UserTag              string                                                                                    `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity3583 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc3584      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate3585     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsSeverity3583 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsInc3584 struct {
	JwtAuthorizeFailure int    `json:"jwt-authorize-failure"`
	JwtMissingToken     int    `json:"jwt-missing-token"`
	JwtMissingClaim     int    `json:"jwt-missing-claim"`
	JwtTokenExpired     int    `json:"jwt-token-expired"`
	JwtSignatureFailure int    `json:"jwt-signature-failure"`
	JwtOtherError       int    `json:"jwt-other-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplListTriggerStatsRate3585 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	JwtAuthorizeFailure int    `json:"jwt-authorize-failure"`
	JwtMissingToken     int    `json:"jwt-missing-token"`
	JwtMissingClaim     int    `json:"jwt-missing-claim"`
	JwtTokenExpired     int    `json:"jwt-token-expired"`
	JwtSignatureFailure int    `json:"jwt-signature-failure"`
	JwtOtherError       int    `json:"jwt-other-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplList3586 struct {
	Name                 string                                                                             `json:"name"`
	CaptureConfig        string                                                                             `json:"capture-config"`
	Uuid                 string                                                                             `json:"uuid"`
	UserTag              string                                                                             `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity3587 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc3588      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate3589     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsSeverity3587 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsInc3588 struct {
	Error int    `json:"error"`
	Uuid  string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplListTriggerStatsRate3589 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Error               int    `json:"error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplList3590 struct {
	Name                 string                                                                                    `json:"name"`
	CaptureConfig        string                                                                                    `json:"capture-config"`
	Uuid                 string                                                                                    `json:"uuid"`
	UserTag              string                                                                                    `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity3591 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc3592      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate3593     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsSeverity3591 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsInc3592 struct {
	Spn_krb_faiure int    `json:"spn_krb_faiure"`
	Uuid           string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplListTriggerStatsRate3593 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Spn_krb_faiure      int    `json:"spn_krb_faiure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplList3594 struct {
	Name                 string                                                                                      `json:"name"`
	CaptureConfig        string                                                                                      `json:"capture-config"`
	Uuid                 string                                                                                      `json:"uuid"`
	UserTag              string                                                                                      `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity3595 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc3596      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate3597     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsSeverity3595 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsInc3596 struct {
	AdminBindFailure  int    `json:"admin-bind-failure"`
	BindFailure       int    `json:"bind-failure"`
	SearchFailure     int    `json:"search-failure"`
	AuthorizeFailure  int    `json:"authorize-failure"`
	TimeoutError      int    `json:"timeout-error"`
	OtherError        int    `json:"other-error"`
	SslSessionFailure int    `json:"ssl-session-failure"`
	Pw_change_failure int    `json:"pw_change_failure"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerLdapInstTmplListTriggerStatsRate3597 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	AdminBindFailure    int    `json:"admin-bind-failure"`
	BindFailure         int    `json:"bind-failure"`
	SearchFailure       int    `json:"search-failure"`
	AuthorizeFailure    int    `json:"authorize-failure"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	SslSessionFailure   int    `json:"ssl-session-failure"`
	Pw_change_failure   int    `json:"pw_change_failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplList3598 struct {
	Name                 string                                                                                      `json:"name"`
	CaptureConfig        string                                                                                      `json:"capture-config"`
	Uuid                 string                                                                                      `json:"uuid"`
	UserTag              string                                                                                      `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity3599 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc3600      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate3601     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsSeverity3599 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsInc3600 struct {
	Timeout         int    `json:"timeout"`
	Fail            int    `json:"fail"`
	StaplingTimeout int    `json:"stapling-timeout"`
	StaplingFail    int    `json:"stapling-fail"`
	Uuid            string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplListTriggerStatsRate3601 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Timeout             int    `json:"timeout"`
	Fail                int    `json:"fail"`
	StaplingTimeout     int    `json:"stapling-timeout"`
	StaplingFail        int    `json:"stapling-fail"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplList3602 struct {
	Name                 string                                                                                     `json:"name"`
	CaptureConfig        string                                                                                     `json:"capture-config"`
	Uuid                 string                                                                                     `json:"uuid"`
	UserTag              string                                                                                     `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity3603 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc3604      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate3605     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsSeverity3603 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsInc3604 struct {
	Authen_failure    int    `json:"authen_failure"`
	Authorize_failure int    `json:"authorize_failure"`
	Timeout_error     int    `json:"timeout_error"`
	Other_error       int    `json:"other_error"`
	AccountingFailure int    `json:"accounting-failure"`
	Uuid              string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerRadInstTmplListTriggerStatsRate3605 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Authen_failure      int    `json:"authen_failure"`
	Authorize_failure   int    `json:"authorize_failure"`
	Timeout_error       int    `json:"timeout_error"`
	Other_error         int    `json:"other_error"`
	AccountingFailure   int    `json:"accounting-failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplList3606 struct {
	Name                 string                                                                                     `json:"name"`
	CaptureConfig        string                                                                                     `json:"capture-config"`
	Uuid                 string                                                                                     `json:"uuid"`
	UserTag              string                                                                                     `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity3607 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc3608      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate3609     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsSeverity3607 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsInc3608 struct {
	Krb_timeout_error        int    `json:"krb_timeout_error"`
	Krb_other_error          int    `json:"krb_other_error"`
	Krb_pw_expiry            int    `json:"krb_pw_expiry"`
	Krb_pw_change_failure    int    `json:"krb_pw_change_failure"`
	Krb_validate_kdc_failure int    `json:"krb_validate_kdc_failure"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServerWinInstTmplListTriggerStatsRate3609 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	Krb_timeout_error        int    `json:"krb_timeout_error"`
	Krb_other_error          int    `json:"krb_other_error"`
	Krb_pw_expiry            int    `json:"krb_pw_expiry"`
	Krb_pw_change_failure    int    `json:"krb_pw_change_failure"`
	Krb_validate_kdc_failure int    `json:"krb_validate_kdc_failure"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplList3610 struct {
	Name                 string                                                                                       `json:"name"`
	CaptureConfig        string                                                                                       `json:"capture-config"`
	Uuid                 string                                                                                       `json:"uuid"`
	UserTag              string                                                                                       `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity3611 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc3612      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate3613     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsSeverity3611 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsInc3612 struct {
	AcsAuthzFail int    `json:"acs-authz-fail"`
	AcsError     int    `json:"acs-error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlServiceProvTmplListTriggerStatsRate3613 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	AcsAuthzFail        int    `json:"acs-authz-fail"`
	AcsError            int    `json:"acs-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplList3614 struct {
	Name                 string                                                                                  `json:"name"`
	CaptureConfig        string                                                                                  `json:"capture-config"`
	Uuid                 string                                                                                  `json:"uuid"`
	UserTag              string                                                                                  `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity3615 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc3616      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate3617     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsSeverity3615 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsInc3616 struct {
	MdFail  int    `json:"md-fail"`
	AcsFail int    `json:"acs-fail"`
	Uuid    string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthSamlIdProvTmplListTriggerStatsRate3617 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	MdFail              int    `json:"md-fail"`
	AcsFail             int    `json:"acs-fail"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplList3618 struct {
	Name                 string                                                                                    `json:"name"`
	CaptureConfig        string                                                                                    `json:"capture-config"`
	Uuid                 string                                                                                    `json:"uuid"`
	UserTag              string                                                                                    `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity3619 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc3620      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate3621     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsSeverity3619 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsInc3620 struct {
	Server_selection_fail_reset int    `json:"server_selection_fail_reset"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplListTriggerStatsRate3621 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	Server_selection_fail_reset int    `json:"server_selection_fail_reset"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplList3622 struct {
	Name                 string                                                                                       `json:"name"`
	CaptureConfig        string                                                                                       `json:"capture-config"`
	Uuid                 string                                                                                       `json:"uuid"`
	UserTag              string                                                                                       `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity3623 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc3624      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate3625     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsSeverity3623 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsInc3624 struct {
	Curr_conn_overflow int    `json:"curr_conn_overflow"`
	Uuid               string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupMemTmplListTriggerStatsRate3625 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Curr_conn_overflow  int    `json:"curr_conn_overflow"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplList3626 struct {
	Name                 string                                                                                      `json:"name"`
	CaptureConfig        string                                                                                      `json:"capture-config"`
	Uuid                 string                                                                                      `json:"uuid"`
	UserTag              string                                                                                      `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity3627 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc3628      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate3629     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsSeverity3627 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsInc3628 struct {
	NoCreds     int    `json:"no-creds"`
	BadReq      int    `json:"bad-req"`
	Unauth      int    `json:"unauth"`
	Forbidden   int    `json:"forbidden"`
	NotFound    int    `json:"not-found"`
	ServerError int    `json:"server-error"`
	Unavailable int    `json:"unavailable"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayHbaseInstTmplListTriggerStatsRate3629 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	NoCreds             int    `json:"no-creds"`
	BadReq              int    `json:"bad-req"`
	Unauth              int    `json:"unauth"`
	Forbidden           int    `json:"forbidden"`
	NotFound            int    `json:"not-found"`
	ServerError         int    `json:"server-error"`
	Unavailable         int    `json:"unavailable"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplList3630 struct {
	Name                 string                                                                                     `json:"name"`
	CaptureConfig        string                                                                                     `json:"capture-config"`
	Uuid                 string                                                                                     `json:"uuid"`
	UserTag              string                                                                                     `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity3631 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc3632      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate3633     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsSeverity3631 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsInc3632 struct {
	Invalid_srv_rsp int    `json:"invalid_srv_rsp"`
	Post_fail       int    `json:"post_fail"`
	Invalid_cred    int    `json:"invalid_cred"`
	Bad_req         int    `json:"bad_req"`
	Not_fnd         int    `json:"not_fnd"`
	Error           int    `json:"error"`
	Other_error     int    `json:"other_error"`
	Uuid            string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayFormInstTmplListTriggerStatsRate3633 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Invalid_srv_rsp     int    `json:"invalid_srv_rsp"`
	Post_fail           int    `json:"post_fail"`
	Invalid_cred        int    `json:"invalid_cred"`
	Bad_req             int    `json:"bad_req"`
	Not_fnd             int    `json:"not_fnd"`
	Error               int    `json:"error"`
	Other_error         int    `json:"other_error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplList3634 struct {
	Name                 string                                                                                  `json:"name"`
	CaptureConfig        string                                                                                  `json:"capture-config"`
	Uuid                 string                                                                                  `json:"uuid"`
	UserTag              string                                                                                  `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity3635 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc3636      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate3637     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsSeverity3635 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsInc3636 struct {
	Failure int    `json:"failure"`
	Uuid    string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthRelayWsFedTmplListTriggerStatsRate3637 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Failure             int    `json:"failure"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplList3638 struct {
	Name                 string                                                                                   `json:"name"`
	CaptureConfig        string                                                                                   `json:"capture-config"`
	Uuid                 string                                                                                   `json:"uuid"`
	UserTag              string                                                                                   `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity3639 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc3640      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate3641     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsSeverity3639 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsInc3640 struct {
	ParseFail    int    `json:"parse-fail"`
	JsonFail     int    `json:"json-fail"`
	AttrFail     int    `json:"attr-fail"`
	TimeoutError int    `json:"timeout-error"`
	OtherError   int    `json:"other-error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplListTriggerStatsRate3641 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	ParseFail           int    `json:"parse-fail"`
	JsonFail            int    `json:"json-fail"`
	AttrFail            int    `json:"attr-fail"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplList3642 struct {
	Name                 string                                                                              `json:"name"`
	CaptureConfig        string                                                                              `json:"capture-config"`
	Uuid                 string                                                                              `json:"uuid"`
	UserTag              string                                                                              `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity3643 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc3644      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate3645     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsSeverity3643 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsInc3644 struct {
	Nc_req_header         int    `json:"nc_req_header"`
	Nc_res_header         int    `json:"nc_res_header"`
	Rv_failure            int    `json:"rv_failure"`
	Content_toobig        int    `json:"content_toobig"`
	Content_toosmall      int    `json:"content_toosmall"`
	Entry_create_failures int    `json:"entry_create_failures"`
	Header_save_error     int    `json:"header_save_error"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplListTriggerStatsRate3645 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	Nc_req_header         int    `json:"nc_req_header"`
	Nc_res_header         int    `json:"nc_res_header"`
	Rv_failure            int    `json:"rv_failure"`
	Content_toobig        int    `json:"content_toobig"`
	Content_toosmall      int    `json:"content_toosmall"`
	Entry_create_failures int    `json:"entry_create_failures"`
	Header_save_error     int    `json:"header_save_error"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplList3646 struct {
	Name                 string                                                                                    `json:"name"`
	CaptureConfig        string                                                                                    `json:"capture-config"`
	Uuid                 string                                                                                    `json:"uuid"`
	UserTag              string                                                                                    `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplListTriggerStatsSeverity3647 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplListTriggerStatsInc3648      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplListTriggerStatsRate3649     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplListTriggerStatsSeverity3647 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplListTriggerStatsInc3648 struct {
	ParseFail    int    `json:"parse-fail"`
	JsonFail     int    `json:"json-fail"`
	TimeoutError int    `json:"timeout-error"`
	OtherError   int    `json:"other-error"`
	Uuid         string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplListTriggerStatsRate3649 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	ParseFail           int    `json:"parse-fail"`
	JsonFail            int    `json:"json-fail"`
	TimeoutError        int    `json:"timeout-error"`
	OtherError          int    `json:"other-error"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbPortTmplList3650 struct {
	Name                 string                                                                        `json:"name"`
	CaptureConfig        string                                                                        `json:"capture-config"`
	Uuid                 string                                                                        `json:"uuid"`
	UserTag              string                                                                        `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity3651 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc3652      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate3653     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsSeverity3651 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsInc3652 struct {
	Es_resp_300 int    `json:"es_resp_300"`
	Es_resp_400 int    `json:"es_resp_400"`
	Es_resp_500 int    `json:"es_resp_500"`
	Resp3xx     int    `json:"resp-3xx"`
	Resp4xx     int    `json:"resp-4xx"`
	Resp5xx     int    `json:"resp-5xx"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbPortTmplListTriggerStatsRate3653 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Es_resp_300         int    `json:"es_resp_300"`
	Es_resp_400         int    `json:"es_resp_400"`
	Es_resp_500         int    `json:"es_resp_500"`
	Resp3xx             int    `json:"resp-3xx"`
	Resp4xx             int    `json:"resp-4xx"`
	Resp5xx             int    `json:"resp-5xx"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplList3654 struct {
	Name                 string                                                                           `json:"name"`
	CaptureConfig        string                                                                           `json:"capture-config"`
	Uuid                 string                                                                           `json:"uuid"`
	UserTag              string                                                                           `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbServiceTmplListTriggerStatsSeverity3655 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesSlbServiceTmplListTriggerStatsInc3656      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesSlbServiceTmplListTriggerStatsRate3657     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplListTriggerStatsSeverity3655 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplListTriggerStatsInc3656 struct {
	Es_resp_300 int    `json:"es_resp_300"`
	Es_resp_400 int    `json:"es_resp_400"`
	Es_resp_500 int    `json:"es_resp_500"`
	Resp3xx     int    `json:"resp-3xx"`
	Resp4xx     int    `json:"resp-4xx"`
	Resp5xx     int    `json:"resp-5xx"`
	Uuid        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbServiceTmplListTriggerStatsRate3657 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Es_resp_300         int    `json:"es_resp_300"`
	Es_resp_400         int    `json:"es_resp_400"`
	Es_resp_500         int    `json:"es_resp_500"`
	Resp3xx             int    `json:"resp-3xx"`
	Resp4xx             int    `json:"resp-4xx"`
	Resp5xx             int    `json:"resp-5xx"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbVportTmplList3658 struct {
	Name                 string                                                                         `json:"name"`
	CaptureConfig        string                                                                         `json:"capture-config"`
	Uuid                 string                                                                         `json:"uuid"`
	UserTag              string                                                                         `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity3659 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc3660      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate3661     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsSeverity3659 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsInc3660 struct {
	Total_mf_dns_pkts                   int    `json:"total_mf_dns_pkts"`
	Es_total_failure_actions            int    `json:"es_total_failure_actions"`
	Compression_miss_no_client          int    `json:"compression_miss_no_client"`
	Compression_miss_template_exclusion int    `json:"compression_miss_template_exclusion"`
	Loc_deny                            int    `json:"loc_deny"`
	Dnsrrl_total_dropped                int    `json:"dnsrrl_total_dropped"`
	Dnsrrl_bad_fqdn                     int    `json:"dnsrrl_bad_fqdn"`
	Dnsrrl_nx_exceed                    int    `json:"dnsrrl_nx_exceed"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSlbVportTmplListTriggerStatsRate3661 struct {
	ThresholdExceededBy                 int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                            int    `json:"duration" dval:"60"`
	Total_mf_dns_pkts                   int    `json:"total_mf_dns_pkts"`
	Es_total_failure_actions            int    `json:"es_total_failure_actions"`
	Compression_miss_no_client          int    `json:"compression_miss_no_client"`
	Compression_miss_template_exclusion int    `json:"compression_miss_template_exclusion"`
	Loc_deny                            int    `json:"loc_deny"`
	Dnsrrl_total_dropped                int    `json:"dnsrrl_total_dropped"`
	Dnsrrl_bad_fqdn                     int    `json:"dnsrrl_bad_fqdn"`
	Dnsrrl_nx_exceed                    int    `json:"dnsrrl_nx_exceed"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplList3662 struct {
	Name                 string                                                                               `json:"name"`
	CaptureConfig        string                                                                               `json:"capture-config"`
	Uuid                 string                                                                               `json:"uuid"`
	UserTag              string                                                                               `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity3663 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc3664      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate3665     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsSeverity3663 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsInc3664 struct {
	Server_selection_fail_drop  int    `json:"server_selection_fail_drop"`
	Server_selection_fail_reset int    `json:"server_selection_fail_reset"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6ServGroupTmplListTriggerStatsRate3665 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	Server_selection_fail_drop  int    `json:"server_selection_fail_drop"`
	Server_selection_fail_reset int    `json:"server_selection_fail_reset"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplList3666 struct {
	Name                 string                                                                                 `json:"name"`
	CaptureConfig        string                                                                                 `json:"capture-config"`
	Uuid                 string                                                                                 `json:"uuid"`
	UserTag              string                                                                                 `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity3667 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc3668      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate3669     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsSeverity3667 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsInc3668 struct {
	Es_total_failure_actions int    `json:"es_total_failure_actions"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplListTriggerStatsRate3669 struct {
	ThresholdExceededBy      int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                 int    `json:"duration" dval:"60"`
	Es_total_failure_actions int    `json:"es_total_failure_actions"`
	Uuid                     string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplList3670 struct {
	Name                 string                                                                                    `json:"name"`
	CaptureConfig        string                                                                                    `json:"capture-config"`
	Uuid                 string                                                                                    `json:"uuid"`
	UserTag              string                                                                                    `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity3671 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc3672      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate3673     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsSeverity3671 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsInc3672 struct {
	Inbound_addr_port_validation_failed int    `json:"inbound_addr_port_validation_failed"`
	Inbound_rev_lookup_failed           int    `json:"inbound_rev_lookup_failed"`
	Inbound_dest_unreachable            int    `json:"inbound_dest_unreachable"`
	Outbound_addr_validation_failed     int    `json:"outbound_addr_validation_failed"`
	Outbound_rev_lookup_failed          int    `json:"outbound_rev_lookup_failed"`
	Outbound_dest_unreachable           int    `json:"outbound_dest_unreachable"`
	Packet_mtu_exceeded                 int    `json:"packet_mtu_exceeded"`
	Interface_not_configured            int    `json:"interface_not_configured"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6MapTransDomainTmplListTriggerStatsRate3673 struct {
	ThresholdExceededBy                 int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                            int    `json:"duration" dval:"60"`
	Inbound_addr_port_validation_failed int    `json:"inbound_addr_port_validation_failed"`
	Inbound_rev_lookup_failed           int    `json:"inbound_rev_lookup_failed"`
	Inbound_dest_unreachable            int    `json:"inbound_dest_unreachable"`
	Outbound_addr_validation_failed     int    `json:"outbound_addr_validation_failed"`
	Outbound_rev_lookup_failed          int    `json:"outbound_rev_lookup_failed"`
	Outbound_dest_unreachable           int    `json:"outbound_dest_unreachable"`
	Packet_mtu_exceeded                 int    `json:"packet_mtu_exceeded"`
	Interface_not_configured            int    `json:"interface_not_configured"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplList3674 struct {
	Name                 string                                                                                 `json:"name"`
	CaptureConfig        string                                                                                 `json:"capture-config"`
	Uuid                 string                                                                                 `json:"uuid"`
	UserTag              string                                                                                 `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity3675 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc3676      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate3677     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsSeverity3675 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsInc3676 struct {
	Inbound_addr_port_validation_failed int    `json:"inbound_addr_port_validation_failed"`
	Inbound_rev_lookup_failed           int    `json:"inbound_rev_lookup_failed"`
	Inbound_dest_unreachable            int    `json:"inbound_dest_unreachable"`
	Outbound_addr_validation_failed     int    `json:"outbound_addr_validation_failed"`
	Outbound_rev_lookup_failed          int    `json:"outbound_rev_lookup_failed"`
	Outbound_dest_unreachable           int    `json:"outbound_dest_unreachable"`
	Packet_mtu_exceeded                 int    `json:"packet_mtu_exceeded"`
	Interface_not_configured            int    `json:"interface_not_configured"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplListTriggerStatsRate3677 struct {
	ThresholdExceededBy                 int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                            int    `json:"duration" dval:"60"`
	Inbound_addr_port_validation_failed int    `json:"inbound_addr_port_validation_failed"`
	Inbound_rev_lookup_failed           int    `json:"inbound_rev_lookup_failed"`
	Inbound_dest_unreachable            int    `json:"inbound_dest_unreachable"`
	Outbound_addr_validation_failed     int    `json:"outbound_addr_validation_failed"`
	Outbound_rev_lookup_failed          int    `json:"outbound_rev_lookup_failed"`
	Outbound_dest_unreachable           int    `json:"outbound_dest_unreachable"`
	Packet_mtu_exceeded                 int    `json:"packet_mtu_exceeded"`
	Interface_not_configured            int    `json:"interface_not_configured"`
	Uuid                                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplList3678 struct {
	Name                 string                                                                               `json:"name"`
	CaptureConfig        string                                                                               `json:"capture-config"`
	Uuid                 string                                                                               `json:"uuid"`
	UserTag              string                                                                               `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity3679 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc3680      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate3681     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsSeverity3679 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsInc3680 struct {
	Nat44RecordsSentFailure            int    `json:"nat44-records-sent-failure"`
	Nat64RecordsSentFailure            int    `json:"nat64-records-sent-failure"`
	DsliteRecordsSentFailure           int    `json:"dslite-records-sent-failure"`
	SessionEventNat44RecordsSentFailur int    `json:"session-event-nat44-records-sent-failur"`
	SessionEventNat64RecordsSentFailur int    `json:"session-event-nat64-records-sent-failur"`
	SessionEventDsliteRecordsSentFailu int    `json:"session-event-dslite-records-sent-failu"`
	SessionEventFw4RecordsSentFailure  int    `json:"session-event-fw4-records-sent-failure"`
	SessionEventFw6RecordsSentFailure  int    `json:"session-event-fw6-records-sent-failure"`
	PortMappingNat44RecordsSentFailure int    `json:"port-mapping-nat44-records-sent-failure"`
	PortMappingNat64RecordsSentFailure int    `json:"port-mapping-nat64-records-sent-failure"`
	PortMappingDsliteRecordsSentFailur int    `json:"port-mapping-dslite-records-sent-failur"`
	NetflowV5RecordsSentFailure        int    `json:"netflow-v5-records-sent-failure"`
	NetflowV5ExtRecordsSentFailure     int    `json:"netflow-v5-ext-records-sent-failure"`
	PortBatchingNat44RecordsSentFailur int    `json:"port-batching-nat44-records-sent-failur"`
	PortBatchingNat64RecordsSentFailur int    `json:"port-batching-nat64-records-sent-failur"`
	PortBatchingDsliteRecordsSentFailu int    `json:"port-batching-dslite-records-sent-failu"`
	PortBatchingV2Nat44RecordsSentFai  int    `json:"port-batching-v2-nat44-records-sent-fai"`
	PortBatchingV2Nat64RecordsSentFai  int    `json:"port-batching-v2-nat64-records-sent-fai"`
	PortBatchingV2DsliteRecordsSentFa  int    `json:"port-batching-v2-dslite-records-sent-fa"`
	CustomSessionEventNat44CreationRec int    `json:"custom-session-event-nat44-creation-rec"`
	CustomSessionEventNat64CreationRec int    `json:"custom-session-event-nat64-creation-rec"`
	CustomSessionEventDsliteCreationRe int    `json:"custom-session-event-dslite-creation-re"`
	CustomSessionEventNat44DeletionRec int    `json:"custom-session-event-nat44-deletion-rec"`
	CustomSessionEventNat64DeletionRec int    `json:"custom-session-event-nat64-deletion-rec"`
	CustomSessionEventDsliteDeletionRe int    `json:"custom-session-event-dslite-deletion-re"`
	CustomSessionEventFw4CreationRecor int    `json:"custom-session-event-fw4-creation-recor"`
	CustomSessionEventFw6CreationRecor int    `json:"custom-session-event-fw6-creation-recor"`
	CustomSessionEventFw4DeletionRecor int    `json:"custom-session-event-fw4-deletion-recor"`
	CustomSessionEventFw6DeletionRecor int    `json:"custom-session-event-fw6-deletion-recor"`
	CustomDenyResetEventFw4RecordsSen  int    `json:"custom-deny-reset-event-fw4-records-sen"`
	CustomDenyResetEventFw6RecordsSen  int    `json:"custom-deny-reset-event-fw6-records-sen"`
	CustomPortMappingNat44CreationReco int    `json:"custom-port-mapping-nat44-creation-reco"`
	CustomPortMappingNat64CreationReco int    `json:"custom-port-mapping-nat64-creation-reco"`
	CustomPortMappingDsliteCreationRec int    `json:"custom-port-mapping-dslite-creation-rec"`
	CustomPortMappingNat44DeletionReco int    `json:"custom-port-mapping-nat44-deletion-reco"`
	CustomPortMappingNat64DeletionReco int    `json:"custom-port-mapping-nat64-deletion-reco"`
	CustomPortMappingDsliteDeletionRec int    `json:"custom-port-mapping-dslite-deletion-rec"`
	CustomPortBatchingNat44CreationRec int    `json:"custom-port-batching-nat44-creation-rec"`
	CustomPortBatchingNat64CreationRec int    `json:"custom-port-batching-nat64-creation-rec"`
	CustomPortBatchingDsliteCreationRe int    `json:"custom-port-batching-dslite-creation-re"`
	CustomPortBatchingNat44DeletionRec int    `json:"custom-port-batching-nat44-deletion-rec"`
	CustomPortBatchingNat64DeletionRec int    `json:"custom-port-batching-nat64-deletion-rec"`
	CustomPortBatchingDsliteDeletionRe int    `json:"custom-port-batching-dslite-deletion-re"`
	CustomPortBatchingV2Nat44Creation  int    `json:"custom-port-batching-v2-nat44-creation-"`
	CustomPortBatchingV2Nat64Creation  int    `json:"custom-port-batching-v2-nat64-creation-"`
	CustomPortBatchingV2DsliteCreation int    `json:"custom-port-batching-v2-dslite-creation"`
	CustomPortBatchingV2Nat44Deletion  int    `json:"custom-port-batching-v2-nat44-deletion-"`
	CustomPortBatchingV2Nat64Deletion  int    `json:"custom-port-batching-v2-nat64-deletion-"`
	CustomPortBatchingV2DsliteDeletion int    `json:"custom-port-batching-v2-dslite-deletion"`
	CustomGtpCTunnelEventRecordsSent   int    `json:"custom-gtp-c-tunnel-event-records-sent-"`
	CustomGtpUTunnelEventRecordsSent   int    `json:"custom-gtp-u-tunnel-event-records-sent-"`
	CustomGtpDenyEventRecordsSentFail  int    `json:"custom-gtp-deny-event-records-sent-fail"`
	CustomGtpInfoEventRecordsSentFail  int    `json:"custom-gtp-info-event-records-sent-fail"`
	CustomFwIddosEntryCreatedRecordsS  int    `json:"custom-fw-iddos-entry-created-records-s"`
	CustomFwIddosEntryDeletedRecordsS  int    `json:"custom-fw-iddos-entry-deleted-records-s"`
	CustomFwSesnLimitExceededRecordsS  int    `json:"custom-fw-sesn-limit-exceeded-records-s"`
	CustomNatIddosL3EntryCreatedRecor  int    `json:"custom-nat-iddos-l3-entry-created-recor"`
	CustomNatIddosL3EntryDeletedRecor  int    `json:"custom-nat-iddos-l3-entry-deleted-recor"`
	CustomNatIddosL4EntryCreatedRecor  int    `json:"custom-nat-iddos-l4-entry-created-recor"`
	CustomNatIddosL4EntryDeletedRecor  int    `json:"custom-nat-iddos-l4-entry-deleted-recor"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplListTriggerStatsRate3681 struct {
	ThresholdExceededBy                int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                           int    `json:"duration" dval:"60"`
	Nat44RecordsSentFailure            int    `json:"nat44-records-sent-failure"`
	Nat64RecordsSentFailure            int    `json:"nat64-records-sent-failure"`
	DsliteRecordsSentFailure           int    `json:"dslite-records-sent-failure"`
	SessionEventNat44RecordsSentFailur int    `json:"session-event-nat44-records-sent-failur"`
	SessionEventNat64RecordsSentFailur int    `json:"session-event-nat64-records-sent-failur"`
	SessionEventDsliteRecordsSentFailu int    `json:"session-event-dslite-records-sent-failu"`
	SessionEventFw4RecordsSentFailure  int    `json:"session-event-fw4-records-sent-failure"`
	SessionEventFw6RecordsSentFailure  int    `json:"session-event-fw6-records-sent-failure"`
	PortMappingNat44RecordsSentFailure int    `json:"port-mapping-nat44-records-sent-failure"`
	PortMappingNat64RecordsSentFailure int    `json:"port-mapping-nat64-records-sent-failure"`
	PortMappingDsliteRecordsSentFailur int    `json:"port-mapping-dslite-records-sent-failur"`
	NetflowV5RecordsSentFailure        int    `json:"netflow-v5-records-sent-failure"`
	NetflowV5ExtRecordsSentFailure     int    `json:"netflow-v5-ext-records-sent-failure"`
	PortBatchingNat44RecordsSentFailur int    `json:"port-batching-nat44-records-sent-failur"`
	PortBatchingNat64RecordsSentFailur int    `json:"port-batching-nat64-records-sent-failur"`
	PortBatchingDsliteRecordsSentFailu int    `json:"port-batching-dslite-records-sent-failu"`
	PortBatchingV2Nat44RecordsSentFai  int    `json:"port-batching-v2-nat44-records-sent-fai"`
	PortBatchingV2Nat64RecordsSentFai  int    `json:"port-batching-v2-nat64-records-sent-fai"`
	PortBatchingV2DsliteRecordsSentFa  int    `json:"port-batching-v2-dslite-records-sent-fa"`
	CustomSessionEventNat44CreationRec int    `json:"custom-session-event-nat44-creation-rec"`
	CustomSessionEventNat64CreationRec int    `json:"custom-session-event-nat64-creation-rec"`
	CustomSessionEventDsliteCreationRe int    `json:"custom-session-event-dslite-creation-re"`
	CustomSessionEventNat44DeletionRec int    `json:"custom-session-event-nat44-deletion-rec"`
	CustomSessionEventNat64DeletionRec int    `json:"custom-session-event-nat64-deletion-rec"`
	CustomSessionEventDsliteDeletionRe int    `json:"custom-session-event-dslite-deletion-re"`
	CustomSessionEventFw4CreationRecor int    `json:"custom-session-event-fw4-creation-recor"`
	CustomSessionEventFw6CreationRecor int    `json:"custom-session-event-fw6-creation-recor"`
	CustomSessionEventFw4DeletionRecor int    `json:"custom-session-event-fw4-deletion-recor"`
	CustomSessionEventFw6DeletionRecor int    `json:"custom-session-event-fw6-deletion-recor"`
	CustomDenyResetEventFw4RecordsSen  int    `json:"custom-deny-reset-event-fw4-records-sen"`
	CustomDenyResetEventFw6RecordsSen  int    `json:"custom-deny-reset-event-fw6-records-sen"`
	CustomPortMappingNat44CreationReco int    `json:"custom-port-mapping-nat44-creation-reco"`
	CustomPortMappingNat64CreationReco int    `json:"custom-port-mapping-nat64-creation-reco"`
	CustomPortMappingDsliteCreationRec int    `json:"custom-port-mapping-dslite-creation-rec"`
	CustomPortMappingNat44DeletionReco int    `json:"custom-port-mapping-nat44-deletion-reco"`
	CustomPortMappingNat64DeletionReco int    `json:"custom-port-mapping-nat64-deletion-reco"`
	CustomPortMappingDsliteDeletionRec int    `json:"custom-port-mapping-dslite-deletion-rec"`
	CustomPortBatchingNat44CreationRec int    `json:"custom-port-batching-nat44-creation-rec"`
	CustomPortBatchingNat64CreationRec int    `json:"custom-port-batching-nat64-creation-rec"`
	CustomPortBatchingDsliteCreationRe int    `json:"custom-port-batching-dslite-creation-re"`
	CustomPortBatchingNat44DeletionRec int    `json:"custom-port-batching-nat44-deletion-rec"`
	CustomPortBatchingNat64DeletionRec int    `json:"custom-port-batching-nat64-deletion-rec"`
	CustomPortBatchingDsliteDeletionRe int    `json:"custom-port-batching-dslite-deletion-re"`
	CustomPortBatchingV2Nat44Creation  int    `json:"custom-port-batching-v2-nat44-creation-"`
	CustomPortBatchingV2Nat64Creation  int    `json:"custom-port-batching-v2-nat64-creation-"`
	CustomPortBatchingV2DsliteCreation int    `json:"custom-port-batching-v2-dslite-creation"`
	CustomPortBatchingV2Nat44Deletion  int    `json:"custom-port-batching-v2-nat44-deletion-"`
	CustomPortBatchingV2Nat64Deletion  int    `json:"custom-port-batching-v2-nat64-deletion-"`
	CustomPortBatchingV2DsliteDeletion int    `json:"custom-port-batching-v2-dslite-deletion"`
	CustomGtpCTunnelEventRecordsSent   int    `json:"custom-gtp-c-tunnel-event-records-sent-"`
	CustomGtpUTunnelEventRecordsSent   int    `json:"custom-gtp-u-tunnel-event-records-sent-"`
	CustomGtpDenyEventRecordsSentFail  int    `json:"custom-gtp-deny-event-records-sent-fail"`
	CustomGtpInfoEventRecordsSentFail  int    `json:"custom-gtp-info-event-records-sent-fail"`
	CustomFwIddosEntryCreatedRecordsS  int    `json:"custom-fw-iddos-entry-created-records-s"`
	CustomFwIddosEntryDeletedRecordsS  int    `json:"custom-fw-iddos-entry-deleted-records-s"`
	CustomFwSesnLimitExceededRecordsS  int    `json:"custom-fw-sesn-limit-exceeded-records-s"`
	CustomNatIddosL3EntryCreatedRecor  int    `json:"custom-nat-iddos-l3-entry-created-recor"`
	CustomNatIddosL3EntryDeletedRecor  int    `json:"custom-nat-iddos-l3-entry-deleted-recor"`
	CustomNatIddosL4EntryCreatedRecor  int    `json:"custom-nat-iddos-l4-entry-created-recor"`
	CustomNatIddosL4EntryDeletedRecor  int    `json:"custom-nat-iddos-l4-entry-deleted-recor"`
	Uuid                               string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetTmplList3682 struct {
	Name                 string                                                                        `json:"name"`
	CaptureConfig        string                                                                        `json:"capture-config"`
	Uuid                 string                                                                        `json:"uuid"`
	UserTag              string                                                                        `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity3683 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc3684      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate3685     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsSeverity3683 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsInc3684 struct {
	UnmatchedDrops int    `json:"unmatched-drops"`
	Deny           int    `json:"deny"`
	Reset          int    `json:"reset"`
	Uuid           string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetTmplListTriggerStatsRate3685 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	UnmatchedDrops      int    `json:"unmatched-drops"`
	Deny                int    `json:"deny"`
	Reset               int    `json:"reset"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplList3686 struct {
	Name                 string                                                                            `json:"name"`
	CaptureConfig        string                                                                            `json:"capture-config"`
	Uuid                 string                                                                            `json:"uuid"`
	UserTag              string                                                                            `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplListTriggerStatsSeverity3687 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplListTriggerStatsInc3688      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplListTriggerStatsRate3689     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplListTriggerStatsSeverity3687 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplListTriggerStatsInc3688 struct {
	SynCookieVerificationFailed int    `json:"syn-cookie-verification-failed"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesRuleSetRuleTmplListTriggerStatsRate3689 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	SynCookieVerificationFailed int    `json:"syn-cookie-verification-failed"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplList3690 struct {
	Name                 string                                                                             `json:"name"`
	CaptureConfig        string                                                                             `json:"capture-config"`
	Uuid                 string                                                                             `json:"uuid"`
	UserTag              string                                                                             `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity3691 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc3692      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate3693     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsSeverity3691 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsInc3692 struct {
	Es_resp_400          int    `json:"es_resp_400"`
	Es_resp_500          int    `json:"es_resp_500"`
	Es_resp_invalid_http int    `json:"es_resp_invalid_http"`
	Uuid                 string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplListTriggerStatsRate3693 struct {
	ThresholdExceededBy  int    `json:"threshold-exceeded-by" dval:"5"`
	Duration             int    `json:"duration" dval:"60"`
	Es_resp_400          int    `json:"es_resp_400"`
	Es_resp_500          int    `json:"es_resp_500"`
	Es_resp_invalid_http int    `json:"es_resp_invalid_http"`
	Uuid                 string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplList3694 struct {
	Name                 string                                                                               `json:"name"`
	CaptureConfig        string                                                                               `json:"capture-config"`
	Uuid                 string                                                                               `json:"uuid"`
	UserTag              string                                                                               `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity3695 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc3696      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate3697     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsSeverity3695 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsInc3696 struct {
	Server_selection_fail_reset int    `json:"server_selection_fail_reset"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupTmplListTriggerStatsRate3697 struct {
	ThresholdExceededBy         int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                    int    `json:"duration" dval:"60"`
	Server_selection_fail_reset int    `json:"server_selection_fail_reset"`
	Uuid                        string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplList3698 struct {
	Name                 string                                                                                  `json:"name"`
	CaptureConfig        string                                                                                  `json:"capture-config"`
	Uuid                 string                                                                                  `json:"uuid"`
	UserTag              string                                                                                  `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity3699 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc3700      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate3701     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsSeverity3699 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsInc3700 struct {
	Curr_conn_overflow int    `json:"curr_conn_overflow"`
	Uuid               string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesFwServiceGroupMemTmplListTriggerStatsRate3701 struct {
	ThresholdExceededBy int    `json:"threshold-exceeded-by" dval:"5"`
	Duration            int    `json:"duration" dval:"60"`
	Curr_conn_overflow  int    `json:"curr_conn_overflow"`
	Uuid                string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesDns_vportTmplList3702 struct {
	Name                 string                                                                          `json:"name"`
	CaptureConfig        string                                                                          `json:"capture-config"`
	Uuid                 string                                                                          `json:"uuid"`
	UserTag              string                                                                          `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity3703 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc3704      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate3705     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsSeverity3703 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsInc3704 struct {
	Dnsrrl_total_dropped           int    `json:"dnsrrl_total_dropped"`
	Total_filter_drop              int    `json:"total_filter_drop"`
	Total_max_query_len_drop       int    `json:"total_max_query_len_drop"`
	Rcode_notimpl_receive          int    `json:"rcode_notimpl_receive"`
	Rcode_notimpl_response         int    `json:"rcode_notimpl_response"`
	Gslb_query_bad                 int    `json:"gslb_query_bad"`
	Gslb_response_bad              int    `json:"gslb_response_bad"`
	Total_dns_filter_type_drop     int    `json:"total_dns_filter_type_drop"`
	Total_dns_filter_class_drop    int    `json:"total_dns_filter_class_drop"`
	Dns_filter_type_a_drop         int    `json:"dns_filter_type_a_drop"`
	Dns_filter_type_aaaa_drop      int    `json:"dns_filter_type_aaaa_drop"`
	Dns_filter_type_cname_drop     int    `json:"dns_filter_type_cname_drop"`
	Dns_filter_type_mx_drop        int    `json:"dns_filter_type_mx_drop"`
	Dns_filter_type_ns_drop        int    `json:"dns_filter_type_ns_drop"`
	Dns_filter_type_srv_drop       int    `json:"dns_filter_type_srv_drop"`
	Dns_filter_type_ptr_drop       int    `json:"dns_filter_type_ptr_drop"`
	Dns_filter_type_soa_drop       int    `json:"dns_filter_type_soa_drop"`
	Dns_filter_type_txt_drop       int    `json:"dns_filter_type_txt_drop"`
	Dns_filter_type_any_drop       int    `json:"dns_filter_type_any_drop"`
	Dns_filter_type_others_drop    int    `json:"dns_filter_type_others_drop"`
	Dns_filter_class_internet_drop int    `json:"dns_filter_class_internet_drop"`
	Dns_filter_class_chaos_drop    int    `json:"dns_filter_class_chaos_drop"`
	Dns_filter_class_hesiod_drop   int    `json:"dns_filter_class_hesiod_drop"`
	Dns_filter_class_none_drop     int    `json:"dns_filter_class_none_drop"`
	Dns_filter_class_any_drop      int    `json:"dns_filter_class_any_drop"`
	Dns_filter_class_others_drop   int    `json:"dns_filter_class_others_drop"`
	Dns_rpz_action_drop            int    `json:"dns_rpz_action_drop"`
	Dnsrrl_bad_fqdn                int    `json:"dnsrrl_bad_fqdn"`
	Dns_filter_tld_drop            int    `json:"dns_filter_tld_drop"`
	Uuid                           string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesDns_vportTmplListTriggerStatsRate3705 struct {
	ThresholdExceededBy            int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                       int    `json:"duration" dval:"60"`
	Dnsrrl_total_dropped           int    `json:"dnsrrl_total_dropped"`
	Total_filter_drop              int    `json:"total_filter_drop"`
	Total_max_query_len_drop       int    `json:"total_max_query_len_drop"`
	Rcode_notimpl_receive          int    `json:"rcode_notimpl_receive"`
	Rcode_notimpl_response         int    `json:"rcode_notimpl_response"`
	Gslb_query_bad                 int    `json:"gslb_query_bad"`
	Gslb_response_bad              int    `json:"gslb_response_bad"`
	Total_dns_filter_type_drop     int    `json:"total_dns_filter_type_drop"`
	Total_dns_filter_class_drop    int    `json:"total_dns_filter_class_drop"`
	Dns_filter_type_a_drop         int    `json:"dns_filter_type_a_drop"`
	Dns_filter_type_aaaa_drop      int    `json:"dns_filter_type_aaaa_drop"`
	Dns_filter_type_cname_drop     int    `json:"dns_filter_type_cname_drop"`
	Dns_filter_type_mx_drop        int    `json:"dns_filter_type_mx_drop"`
	Dns_filter_type_ns_drop        int    `json:"dns_filter_type_ns_drop"`
	Dns_filter_type_srv_drop       int    `json:"dns_filter_type_srv_drop"`
	Dns_filter_type_ptr_drop       int    `json:"dns_filter_type_ptr_drop"`
	Dns_filter_type_soa_drop       int    `json:"dns_filter_type_soa_drop"`
	Dns_filter_type_txt_drop       int    `json:"dns_filter_type_txt_drop"`
	Dns_filter_type_any_drop       int    `json:"dns_filter_type_any_drop"`
	Dns_filter_type_others_drop    int    `json:"dns_filter_type_others_drop"`
	Dns_filter_class_internet_drop int    `json:"dns_filter_class_internet_drop"`
	Dns_filter_class_chaos_drop    int    `json:"dns_filter_class_chaos_drop"`
	Dns_filter_class_hesiod_drop   int    `json:"dns_filter_class_hesiod_drop"`
	Dns_filter_class_none_drop     int    `json:"dns_filter_class_none_drop"`
	Dns_filter_class_any_drop      int    `json:"dns_filter_class_any_drop"`
	Dns_filter_class_others_drop   int    `json:"dns_filter_class_others_drop"`
	Dns_rpz_action_drop            int    `json:"dns_rpz_action_drop"`
	Dnsrrl_bad_fqdn                int    `json:"dnsrrl_bad_fqdn"`
	Dns_filter_tld_drop            int    `json:"dns_filter_tld_drop"`
	Uuid                           string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplList3706 struct {
	Name                 string                                                                          `json:"name"`
	CaptureConfig        string                                                                          `json:"capture-config"`
	Uuid                 string                                                                          `json:"uuid"`
	UserTag              string                                                                          `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity3707 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc3708      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate3709     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsSeverity3707 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsInc3708 struct {
	No_proxy                      int    `json:"no_proxy"`
	Parse_req_fail                int    `json:"parse_req_fail"`
	Server_select_fail            int    `json:"server_select_fail"`
	Forward_req_fail              int    `json:"forward_req_fail"`
	Forward_req_data_fail         int    `json:"forward_req_data_fail"`
	Snat_fail                     int    `json:"snat_fail"`
	Send_client_service_not_ready int    `json:"send_client_service_not_ready"`
	Recv_server_unknow_reply_code int    `json:"recv_server_unknow_reply_code"`
	Read_request_line_fail        int    `json:"read_request_line_fail"`
	Get_all_headers_fail          int    `json:"get_all_headers_fail"`
	Too_many_headers              int    `json:"too_many_headers"`
	Line_too_long                 int    `json:"line_too_long"`
	Line_extend_fail              int    `json:"line_extend_fail"`
	Line_table_extend_fail        int    `json:"line_table_extend_fail"`
	Parse_request_line_fail       int    `json:"parse_request_line_fail"`
	Insert_resonse_line_fail      int    `json:"insert_resonse_line_fail"`
	Remove_resonse_line_fail      int    `json:"remove_resonse_line_fail"`
	Parse_resonse_line_fail       int    `json:"parse_resonse_line_fail"`
	Server_starttls_fail          int    `json:"server_STARTTLS_fail"`
	Uuid                          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesSmtpVportTmplListTriggerStatsRate3709 struct {
	ThresholdExceededBy           int    `json:"threshold-exceeded-by" dval:"5"`
	Duration                      int    `json:"duration" dval:"60"`
	No_proxy                      int    `json:"no_proxy"`
	Parse_req_fail                int    `json:"parse_req_fail"`
	Server_select_fail            int    `json:"server_select_fail"`
	Forward_req_fail              int    `json:"forward_req_fail"`
	Forward_req_data_fail         int    `json:"forward_req_data_fail"`
	Snat_fail                     int    `json:"snat_fail"`
	Send_client_service_not_ready int    `json:"send_client_service_not_ready"`
	Recv_server_unknow_reply_code int    `json:"recv_server_unknow_reply_code"`
	Read_request_line_fail        int    `json:"read_request_line_fail"`
	Get_all_headers_fail          int    `json:"get_all_headers_fail"`
	Too_many_headers              int    `json:"too_many_headers"`
	Line_too_long                 int    `json:"line_too_long"`
	Line_extend_fail              int    `json:"line_extend_fail"`
	Line_table_extend_fail        int    `json:"line_table_extend_fail"`
	Parse_request_line_fail       int    `json:"parse_request_line_fail"`
	Insert_resonse_line_fail      int    `json:"insert_resonse_line_fail"`
	Remove_resonse_line_fail      int    `json:"remove_resonse_line_fail"`
	Parse_resonse_line_fail       int    `json:"parse_resonse_line_fail"`
	Server_starttls_fail          int    `json:"server_STARTTLS_fail"`
	Uuid                          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesPop3VportTmplList3710 struct {
	Name                 string                                                                          `json:"name"`
	CaptureConfig        string                                                                          `json:"capture-config"`
	Uuid                 string                                                                          `json:"uuid"`
	UserTag              string                                                                          `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity3711 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc3712      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate3713     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsSeverity3711 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsInc3712 struct {
	Svrsel_fail           int    `json:"svrsel_fail"`
	No_route              int    `json:"no_route"`
	Snat_fail             int    `json:"snat_fail"`
	Line_too_long         int    `json:"line_too_long"`
	Invalid_start_line    int    `json:"invalid_start_line"`
	Unsupported_command   int    `json:"unsupported_command"`
	Bad_sequence          int    `json:"bad_sequence"`
	Rsv_persist_conn_fail int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail           int    `json:"smp_v6_fail"`
	Smp_v4_fail           int    `json:"smp_v4_fail"`
	Insert_tuple_fail     int    `json:"insert_tuple_fail"`
	Cl_est_err            int    `json:"cl_est_err"`
	Ser_connecting_err    int    `json:"ser_connecting_err"`
	Server_response_err   int    `json:"server_response_err"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesPop3VportTmplListTriggerStatsRate3713 struct {
	ThresholdExceededBy   int    `json:"threshold-exceeded-by" dval:"5"`
	Duration              int    `json:"duration" dval:"60"`
	Svrsel_fail           int    `json:"svrsel_fail"`
	No_route              int    `json:"no_route"`
	Snat_fail             int    `json:"snat_fail"`
	Line_too_long         int    `json:"line_too_long"`
	Invalid_start_line    int    `json:"invalid_start_line"`
	Unsupported_command   int    `json:"unsupported_command"`
	Bad_sequence          int    `json:"bad_sequence"`
	Rsv_persist_conn_fail int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail           int    `json:"smp_v6_fail"`
	Smp_v4_fail           int    `json:"smp_v4_fail"`
	Insert_tuple_fail     int    `json:"insert_tuple_fail"`
	Cl_est_err            int    `json:"cl_est_err"`
	Ser_connecting_err    int    `json:"ser_connecting_err"`
	Server_response_err   int    `json:"server_response_err"`
	Uuid                  string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesImapVportTmplList3714 struct {
	Name                 string                                                                          `json:"name"`
	CaptureConfig        string                                                                          `json:"capture-config"`
	Uuid                 string                                                                          `json:"uuid"`
	UserTag              string                                                                          `json:"user-tag"`
	TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity3715 `json:"trigger-stats-severity"`
	TriggerStatsInc      VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc3716      `json:"trigger-stats-inc"`
	TriggerStatsRate     VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate3717     `json:"trigger-stats-rate"`
}

type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsSeverity3715 struct {
	Error         int    `json:"error"`
	ErrorAlert    int    `json:"error-alert"`
	ErrorWarning  int    `json:"error-warning"`
	ErrorCritical int    `json:"error-critical"`
	Drop          int    `json:"drop"`
	DropAlert     int    `json:"drop-alert"`
	DropWarning   int    `json:"drop-warning"`
	DropCritical  int    `json:"drop-critical"`
	Uuid          string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsInc3716 struct {
	Svrsel_fail            int    `json:"svrsel_fail"`
	No_route               int    `json:"no_route"`
	Snat_fail              int    `json:"snat_fail"`
	Line_too_long          int    `json:"line_too_long"`
	Invalid_start_line     int    `json:"invalid_start_line"`
	Cant_find_pasv         int    `json:"cant_find_pasv"`
	Smp_create_fail        int    `json:"smp_create_fail"`
	Data_server_conn_fail  int    `json:"data_server_conn_fail"`
	Data_send_fail         int    `json:"data_send_fail"`
	Cant_find_epsv         int    `json:"cant_find_epsv"`
	Auth_unsupported       int    `json:"auth_unsupported"`
	Unsupported_pbsz_value int    `json:"unsupported_pbsz_value"`
	Unsupported_prot_value int    `json:"unsupported_prot_value"`
	Bad_sequence           int    `json:"bad_sequence"`
	Rsv_persist_conn_fail  int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail            int    `json:"smp_v6_fail"`
	Smp_v4_fail            int    `json:"smp_v4_fail"`
	Insert_tuple_fail      int    `json:"insert_tuple_fail"`
	Cl_est_err             int    `json:"cl_est_err"`
	Ser_connecting_err     int    `json:"ser_connecting_err"`
	Server_response_err    int    `json:"server_response_err"`
	Cl_request_err         int    `json:"cl_request_err"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureObjectTemplatesImapVportTmplListTriggerStatsRate3717 struct {
	ThresholdExceededBy    int    `json:"threshold-exceeded-by" dval:"5"`
	Duration               int    `json:"duration" dval:"60"`
	Svrsel_fail            int    `json:"svrsel_fail"`
	No_route               int    `json:"no_route"`
	Snat_fail              int    `json:"snat_fail"`
	Line_too_long          int    `json:"line_too_long"`
	Invalid_start_line     int    `json:"invalid_start_line"`
	Cant_find_pasv         int    `json:"cant_find_pasv"`
	Smp_create_fail        int    `json:"smp_create_fail"`
	Data_server_conn_fail  int    `json:"data_server_conn_fail"`
	Data_send_fail         int    `json:"data_send_fail"`
	Cant_find_epsv         int    `json:"cant_find_epsv"`
	Auth_unsupported       int    `json:"auth_unsupported"`
	Unsupported_pbsz_value int    `json:"unsupported_pbsz_value"`
	Unsupported_prot_value int    `json:"unsupported_prot_value"`
	Bad_sequence           int    `json:"bad_sequence"`
	Rsv_persist_conn_fail  int    `json:"rsv_persist_conn_fail"`
	Smp_v6_fail            int    `json:"smp_v6_fail"`
	Smp_v4_fail            int    `json:"smp_v4_fail"`
	Insert_tuple_fail      int    `json:"insert_tuple_fail"`
	Cl_est_err             int    `json:"cl_est_err"`
	Ser_connecting_err     int    `json:"ser_connecting_err"`
	Server_response_err    int    `json:"server_response_err"`
	Cl_request_err         int    `json:"cl_request_err"`
	Uuid                   string `json:"uuid"`
}

type VisibilityPacketCaptureAutomatedCaptures3718 struct {
	Slb_port_tmpl_error_code_return_inc  int    `json:"slb_port_tmpl_error_code_return_inc"`
	Slb_port_tmpl_high_error_code_return int    `json:"slb_port_tmpl_high_error_code_return"`
	Uuid                                 string `json:"uuid"`
}

type VisibilityPingSweepDetection3719 struct {
	Events   int    `json:"events" dval:"10"`
	Interval int    `json:"interval" dval:"60"`
	V4List   string `json:"v4-list"`
	V6List   string `json:"v6-list"`
	Uuid     string `json:"uuid"`
}

type VisibilityPortScanDetection3720 struct {
	Events   int    `json:"events" dval:"10"`
	Interval int    `json:"interval" dval:"60"`
	V4List   string `json:"v4-list"`
	V6List   string `json:"v6-list"`
	Uuid     string `json:"uuid"`
}

type VisibilityReporting3721 struct {
	SessionLogging          string                                         `json:"session-logging" dval:"disable"`
	Uuid                    string                                         `json:"uuid"`
	SamplingEnable          []VisibilityReportingSamplingEnable3722        `json:"sampling-enable"`
	TelemetryExportInterval VisibilityReportingTelemetryExportInterval3723 `json:"telemetry-export-interval"`
	Template                VisibilityReportingTemplate3724                `json:"template"`
}

type VisibilityReportingSamplingEnable3722 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityReportingTelemetryExportInterval3723 struct {
	Value int    `json:"value" dval:"5"`
	Uuid  string `json:"uuid"`
}

type VisibilityReportingTemplate3724 struct {
	Notification VisibilityReportingTemplateNotification3725 `json:"notification"`
}

type VisibilityReportingTemplateNotification3725 struct {
	TemplateNameList []VisibilityReportingTemplateNotificationTemplateNameList3726 `json:"template-name-list"`
	Debug            VisibilityReportingTemplateNotificationDebug3729              `json:"debug"`
}

type VisibilityReportingTemplateNotificationTemplateNameList3726 struct {
	Name             string                                                                      `json:"name"`
	Ipv4Address      string                                                                      `json:"ipv4-address"`
	Ipv6Address      string                                                                      `json:"ipv6-address"`
	HostName         string                                                                      `json:"host-name"`
	UseMgmtPort      int                                                                         `json:"use-mgmt-port"`
	Protocol         string                                                                      `json:"protocol" dval:"https"`
	HttpPort         int                                                                         `json:"http-port" dval:"80"`
	HttpsPort        int                                                                         `json:"https-port" dval:"443"`
	RelativeUri      string                                                                      `json:"relative-uri" dval:"/"`
	Action           string                                                                      `json:"action" dval:"enable"`
	DebugMode        int                                                                         `json:"debug-mode"`
	TestConnectivity int                                                                         `json:"test-connectivity"`
	Uuid             string                                                                      `json:"uuid"`
	SamplingEnable   []VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable3727 `json:"sampling-enable"`
	Authentication   VisibilityReportingTemplateNotificationTemplateNameListAuthentication3728   `json:"authentication"`
}

type VisibilityReportingTemplateNotificationTemplateNameListSamplingEnable3727 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityReportingTemplateNotificationTemplateNameListAuthentication3728 struct {
	RelativeLoginUri   string `json:"relative-login-uri"`
	RelativeLogoffUri  string `json:"relative-logoff-uri"`
	AuthUsername       string `json:"auth-username"`
	AuthPassword       int    `json:"auth-password"`
	AuthPasswordString string `json:"auth-password-string"`
	Encrypted          string `json:"encrypted"`
	ApiKey             int    `json:"api-key"`
	ApiKeyString       string `json:"api-key-string"`
	ApiKeyEncrypted    string `json:"api-key-encrypted"`
	Uuid               string `json:"uuid"`
}

type VisibilityReportingTemplateNotificationDebug3729 struct {
	Uuid string `json:"uuid"`
}

type VisibilityResourceUsage3730 struct {
	Uuid string `json:"uuid"`
}

type VisibilitySamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type VisibilityTopn3731 struct {
	Uuid                          string                                            `json:"uuid"`
	SamplingEnable                []VisibilityTopnSamplingEnable3732                `json:"sampling-enable"`
	TemplGtpPlcyTopnTmplList      []VisibilityTopnTemplGtpPlcyTopnTmplList3733      `json:"templ-gtp-plcy-topn-tmpl-list"`
	TemplGtpPlcyTopnNode          VisibilityTopnTemplGtpPlcyTopnNode3735            `json:"templ-gtp-plcy-topn-node"`
	Cgnv6NatPoolTopnTmplList      []VisibilityTopnCgnv6NatPoolTopnTmplList3736      `json:"cgnv6-nat-pool-topn-tmpl-list"`
	Cgnv6NatPoolTopnNode          VisibilityTopnCgnv6NatPoolTopnNode3738            `json:"cgnv6-nat-pool-topn-node"`
	GtpApnPrefixTopnTmplList      []VisibilityTopnGtpApnPrefixTopnTmplList3739      `json:"gtp-apn-prefix-topn-tmpl-list"`
	GtpApnPrefixTopnNode          VisibilityTopnGtpApnPrefixTopnNode3741            `json:"gtp-apn-prefix-topn-node"`
	GtpNetworkElementTopnTmplList []VisibilityTopnGtpNetworkElementTopnTmplList3742 `json:"gtp-network-element-topn-tmpl-list"`
	GtpNetworkElementTopnNode     VisibilityTopnGtpNetworkElementTopnNode3744       `json:"gtp-network-element-topn-node"`
}

type VisibilityTopnSamplingEnable3732 struct {
	Counters1 string `json:"counters1"`
}

type VisibilityTopnTemplGtpPlcyTopnTmplList3733 struct {
	Name     string                                            `json:"name"`
	TopnSize int                                               `json:"topn-size"`
	Interval string                                            `json:"interval"`
	Uuid     string                                            `json:"uuid"`
	UserTag  string                                            `json:"user-tag"`
	Metrics  VisibilityTopnTemplGtpPlcyTopnTmplListMetrics3734 `json:"metrics"`
}

type VisibilityTopnTemplGtpPlcyTopnTmplListMetrics3734 struct {
	RlMessageMonitor int    `json:"rl-message-monitor"`
	Uuid             string `json:"uuid"`
}

type VisibilityTopnTemplGtpPlcyTopnNode3735 struct {
	Activate string `json:"activate"`
	Uuid     string `json:"uuid"`
}

type VisibilityTopnCgnv6NatPoolTopnTmplList3736 struct {
	Name     string                                            `json:"name"`
	TopnSize int                                               `json:"topn-size"`
	Interval string                                            `json:"interval"`
	Uuid     string                                            `json:"uuid"`
	UserTag  string                                            `json:"user-tag"`
	Metrics  VisibilityTopnCgnv6NatPoolTopnTmplListMetrics3737 `json:"metrics"`
}

type VisibilityTopnCgnv6NatPoolTopnTmplListMetrics3737 struct {
	UdpTotal int    `json:"udp-total"`
	TcpTotal int    `json:"tcp-total"`
	Uuid     string `json:"uuid"`
}

type VisibilityTopnCgnv6NatPoolTopnNode3738 struct {
	Activate string `json:"activate"`
	Uuid     string `json:"uuid"`
}

type VisibilityTopnGtpApnPrefixTopnTmplList3739 struct {
	Name     string                                            `json:"name"`
	TopnSize int                                               `json:"topn-size"`
	Interval string                                            `json:"interval"`
	Uuid     string                                            `json:"uuid"`
	UserTag  string                                            `json:"user-tag"`
	Metrics  VisibilityTopnGtpApnPrefixTopnTmplListMetrics3740 `json:"metrics"`
}

type VisibilityTopnGtpApnPrefixTopnTmplListMetrics3740 struct {
	UplinkBytes                        int    `json:"uplink-bytes"`
	DownlinkBytes                      int    `json:"downlink-bytes"`
	UplinkPkts                         int    `json:"uplink-pkts"`
	DownlinkPkts                       int    `json:"downlink-pkts"`
	UUplinkBytes                       int    `json:"u-uplink-bytes"`
	UDownlinkBytes                     int    `json:"u-downlink-bytes"`
	UUplinkPkts                        int    `json:"u-uplink-pkts"`
	UDownlinkPkts                      int    `json:"u-downlink-pkts"`
	GtpV0CTunnelCreated                int    `json:"gtp-v0-c-tunnel-created"`
	GtpV0CTunnelHalfOpen               int    `json:"gtp-v0-c-tunnel-half-open"`
	GtpV0CTunnelHalfClosed             int    `json:"gtp-v0-c-tunnel-half-closed"`
	GtpV0CTunnelClosed                 int    `json:"gtp-v0-c-tunnel-closed"`
	GtpV0CTunnelDeleted                int    `json:"gtp-v0-c-tunnel-deleted"`
	GtpV0CHalfOpenTunnelClosed         int    `json:"gtp-v0-c-half-open-tunnel-closed"`
	GtpV1CTunnelCreated                int    `json:"gtp-v1-c-tunnel-created"`
	GtpV1CTunnelHalfOpen               int    `json:"gtp-v1-c-tunnel-half-open"`
	GtpV1CTunnelHalfClosed             int    `json:"gtp-v1-c-tunnel-half-closed"`
	GtpV1CTunnelClosed                 int    `json:"gtp-v1-c-tunnel-closed"`
	GtpV1CTunnelDeleted                int    `json:"gtp-v1-c-tunnel-deleted"`
	GtpV1CHalfOpenTunnelClosed         int    `json:"gtp-v1-c-half-open-tunnel-closed"`
	GtpV2CTunnelCreated                int    `json:"gtp-v2-c-tunnel-created"`
	GtpV2CTunnelHalfOpen               int    `json:"gtp-v2-c-tunnel-half-open"`
	GtpV2CTunnelHalfClosed             int    `json:"gtp-v2-c-tunnel-half-closed"`
	GtpV2CTunnelClosed                 int    `json:"gtp-v2-c-tunnel-closed"`
	GtpV2CTunnelDeleted                int    `json:"gtp-v2-c-tunnel-deleted"`
	GtpV2CHalfOpenTunnelClosed         int    `json:"gtp-v2-c-half-open-tunnel-closed"`
	GtpUTunnelCreated                  int    `json:"gtp-u-tunnel-created"`
	GtpUTunnelDeleted                  int    `json:"gtp-u-tunnel-deleted"`
	GtpV0CUpdatePdpRespUnsuccess       int    `json:"gtp-v0-c-update-pdp-resp-unsuccess"`
	GtpV1CUpdatePdpRespUnsuccess       int    `json:"gtp-v1-c-update-pdp-resp-unsuccess"`
	GtpV2CMod_bearerRespUnsuccess      int    `json:"gtp-v2-c-mod_bearer-resp-unsuccess"`
	GtpV0CCreatePdpRespUnsuccess       int    `json:"gtp-v0-c-create-pdp-resp-unsuccess"`
	GtpV1CCreatePdpRespUnsuccess       int    `json:"gtp-v1-c-create-pdp-resp-unsuccess"`
	GtpV2CCreateSessRespUnsuccess      int    `json:"gtp-v2-c-create-sess-resp-unsuccess"`
	GtpV2CPiggybackMessage             int    `json:"gtp-v2-c-piggyback-message"`
	GtpPathManagementMessage           int    `json:"gtp-path-management-message"`
	GtpV0CTunnelDeletedRestart         int    `json:"gtp-v0-c-tunnel-deleted-restart"`
	GtpV1CTunnelDeletedRestart         int    `json:"gtp-v1-c-tunnel-deleted-restart"`
	GtpV2CTunnelDeletedRestart         int    `json:"gtp-v2-c-tunnel-deleted-restart"`
	DropVldReservedFieldSet            int    `json:"drop-vld-reserved-field-set"`
	DropVldTunnelIdFlag                int    `json:"drop-vld-tunnel-id-flag"`
	DropVldInvalidFlowLabelV0          int    `json:"drop-vld-invalid-flow-label-v0"`
	DropVldInvalidTeid                 int    `json:"drop-vld-invalid-teid"`
	DropVldUnsupportedMessageType      int    `json:"drop-vld-unsupported-message-type"`
	DropVldOutOfState                  int    `json:"drop-vld-out-of-state"`
	DropVldMandatoryInformationElement int    `json:"drop-vld-mandatory-information-element"`
	DropVldOutOfOrderIe                int    `json:"drop-vld-out-of-order-ie"`
	DropVldOutOfStateIe                int    `json:"drop-vld-out-of-state-ie"`
	DropVldReservedInformationElement  int    `json:"drop-vld-reserved-information-element"`
	DropVldVersionNotSupported         int    `json:"drop-vld-version-not-supported"`
	DropVldMessageLength               int    `json:"drop-vld-message-length"`
	DropVldCrossLayerCorrelation       int    `json:"drop-vld-cross-layer-correlation"`
	DropVldCountryCodeMismatch         int    `json:"drop-vld-country-code-mismatch"`
	DropVldGtpUSpoofedSourceAddress    int    `json:"drop-vld-gtp-u-spoofed-source-address"`
	DropVldGtpBearerCountExceed        int    `json:"drop-vld-gtp-bearer-count-exceed"`
	DropVldGtpV2WrongLbiCreateBearer   int    `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
	GtpCHandoverInProgressWithConn     int    `json:"gtp-c-handover-in-progress-with-conn"`
	DropVldInvalidPktLenPiggyback      int    `json:"drop-vld-invalid-pkt-len-piggyback"`
	DropVldSanityFailedPiggyback       int    `json:"drop-vld-sanity-failed-piggyback"`
	DropVldSequenceNumCorrelation      int    `json:"drop-vld-sequence-num-correlation"`
	DropVldGtpv0SeqnumBufferFull       int    `json:"drop-vld-gtpv0-seqnum-buffer-full"`
	DropVldGtpv1SeqnumBufferFull       int    `json:"drop-vld-gtpv1-seqnum-buffer-full"`
	DropVldGtpv2SeqnumBufferFull       int    `json:"drop-vld-gtpv2-seqnum-buffer-full"`
	DropVldGtpInvalidImsiLenDrop       int    `json:"drop-vld-gtp-invalid-imsi-len-drop"`
	DropVldGtpInvalidApnLenDrop        int    `json:"drop-vld-gtp-invalid-apn-len-drop"`
	DropVldProtocolFlagUnset           int    `json:"drop-vld-protocol-flag-unset"`
	DropFltMessageFiltering            int    `json:"drop-flt-message-filtering"`
	DropFltApnFiltering                int    `json:"drop-flt-apn-filtering"`
	DropFltMsisdnFiltering             int    `json:"drop-flt-msisdn-filtering"`
	DropFltRatTypeFiltering            int    `json:"drop-flt-rat-type-filtering"`
	DropFltGtpInGtp                    int    `json:"drop-flt-gtp-in-gtp"`
	DropRlGtpV0CAgg                    int    `json:"drop-rl-gtp-v0-c-agg"`
	DropRlGtpV1CAgg                    int    `json:"drop-rl-gtp-v1-c-agg"`
	DropRlGtpV2CAgg                    int    `json:"drop-rl-gtp-v2-c-agg"`
	DropRlGtpV1CCreatePdpRequest       int    `json:"drop-rl-gtp-v1-c-create-pdp-request"`
	DropRlGtpV2CCreateSessionRequest   int    `json:"drop-rl-gtp-v2-c-create-session-request"`
	DropRlGtpV1CUpdatePdpRequest       int    `json:"drop-rl-gtp-v1-c-update-pdp-request"`
	DropRlGtpV2CModifyBearerRequest    int    `json:"drop-rl-gtp-v2-c-modify-bearer-request"`
	DropRlGtpUTunnelCreate             int    `json:"drop-rl-gtp-u-tunnel-create"`
	DropRlGtpUUplinkByte               int    `json:"drop-rl-gtp-u-uplink-byte"`
	DropRlGtpUUplinkPacket             int    `json:"drop-rl-gtp-u-uplink-packet"`
	DropRlGtpUDownlinkByte             int    `json:"drop-rl-gtp-u-downlink-byte"`
	DropRlGtpUDownlinkPacket           int    `json:"drop-rl-gtp-u-downlink-packet"`
	DropRlGtpUTotalByte                int    `json:"drop-rl-gtp-u-total-byte"`
	DropRlGtpUTotalPacket              int    `json:"drop-rl-gtp-u-total-packet"`
	DropRlGtpUMaxConcurrentTunnels     int    `json:"drop-rl-gtp-u-max-concurrent-tunnels"`
	RlMessageMonitor                   int    `json:"rl-message-monitor"`
	Uuid                               string `json:"uuid"`
}

type VisibilityTopnGtpApnPrefixTopnNode3741 struct {
	Activate string `json:"activate"`
	Uuid     string `json:"uuid"`
}

type VisibilityTopnGtpNetworkElementTopnTmplList3742 struct {
	Name     string                                                 `json:"name"`
	TopnSize int                                                    `json:"topn-size"`
	Interval string                                                 `json:"interval"`
	Uuid     string                                                 `json:"uuid"`
	UserTag  string                                                 `json:"user-tag"`
	Metrics  VisibilityTopnGtpNetworkElementTopnTmplListMetrics3743 `json:"metrics"`
}

type VisibilityTopnGtpNetworkElementTopnTmplListMetrics3743 struct {
	UplinkBytes                        int    `json:"uplink-bytes"`
	DownlinkBytes                      int    `json:"downlink-bytes"`
	UplinkPkts                         int    `json:"uplink-pkts"`
	DownlinkPkts                       int    `json:"downlink-pkts"`
	UUplinkBytes                       int    `json:"u-uplink-bytes"`
	UDownlinkBytes                     int    `json:"u-downlink-bytes"`
	UUplinkPkts                        int    `json:"u-uplink-pkts"`
	UDownlinkPkts                      int    `json:"u-downlink-pkts"`
	GtpV0CTunnelCreated                int    `json:"gtp-v0-c-tunnel-created"`
	GtpV0CTunnelHalfOpen               int    `json:"gtp-v0-c-tunnel-half-open"`
	GtpV0CTunnelHalfClosed             int    `json:"gtp-v0-c-tunnel-half-closed"`
	GtpV0CTunnelClosed                 int    `json:"gtp-v0-c-tunnel-closed"`
	GtpV0CTunnelDeleted                int    `json:"gtp-v0-c-tunnel-deleted"`
	GtpV0CHalfOpenTunnelClosed         int    `json:"gtp-v0-c-half-open-tunnel-closed"`
	GtpV1CTunnelCreated                int    `json:"gtp-v1-c-tunnel-created"`
	GtpV1CTunnelHalfOpen               int    `json:"gtp-v1-c-tunnel-half-open"`
	GtpV1CTunnelHalfClosed             int    `json:"gtp-v1-c-tunnel-half-closed"`
	GtpV1CTunnelClosed                 int    `json:"gtp-v1-c-tunnel-closed"`
	GtpV1CTunnelDeleted                int    `json:"gtp-v1-c-tunnel-deleted"`
	GtpV1CHalfOpenTunnelClosed         int    `json:"gtp-v1-c-half-open-tunnel-closed"`
	GtpV2CTunnelCreated                int    `json:"gtp-v2-c-tunnel-created"`
	GtpV2CTunnelHalfOpen               int    `json:"gtp-v2-c-tunnel-half-open"`
	GtpV2CTunnelHalfClosed             int    `json:"gtp-v2-c-tunnel-half-closed"`
	GtpV2CTunnelClosed                 int    `json:"gtp-v2-c-tunnel-closed"`
	GtpV2CTunnelDeleted                int    `json:"gtp-v2-c-tunnel-deleted"`
	GtpV2CHalfOpenTunnelClosed         int    `json:"gtp-v2-c-half-open-tunnel-closed"`
	GtpUTunnelCreated                  int    `json:"gtp-u-tunnel-created"`
	GtpUTunnelDeleted                  int    `json:"gtp-u-tunnel-deleted"`
	GtpV0CUpdatePdpRespUnsuccess       int    `json:"gtp-v0-c-update-pdp-resp-unsuccess"`
	GtpV1CUpdatePdpRespUnsuccess       int    `json:"gtp-v1-c-update-pdp-resp-unsuccess"`
	GtpV2CMod_bearerRespUnsuccess      int    `json:"gtp-v2-c-mod_bearer-resp-unsuccess"`
	GtpV0CCreatePdpRespUnsuccess       int    `json:"gtp-v0-c-create-pdp-resp-unsuccess"`
	GtpV1CCreatePdpRespUnsuccess       int    `json:"gtp-v1-c-create-pdp-resp-unsuccess"`
	GtpV2CCreateSessRespUnsuccess      int    `json:"gtp-v2-c-create-sess-resp-unsuccess"`
	GtpV2CPiggybackMessage             int    `json:"gtp-v2-c-piggyback-message"`
	GtpPathManagementMessage           int    `json:"gtp-path-management-message"`
	GtpV0CTunnelDeletedRestart         int    `json:"gtp-v0-c-tunnel-deleted-restart"`
	GtpV1CTunnelDeletedRestart         int    `json:"gtp-v1-c-tunnel-deleted-restart"`
	GtpV2CTunnelDeletedRestart         int    `json:"gtp-v2-c-tunnel-deleted-restart"`
	GtpV0CReservedMessageAllow         int    `json:"gtp-v0-c-reserved-message-allow"`
	GtpV1CReservedMessageAllow         int    `json:"gtp-v1-c-reserved-message-allow"`
	GtpV2CReservedMessageAllow         int    `json:"gtp-v2-c-reserved-message-allow"`
	DropVldReservedFieldSet            int    `json:"drop-vld-reserved-field-set"`
	DropVldTunnelIdFlag                int    `json:"drop-vld-tunnel-id-flag"`
	DropVldInvalidFlowLabelV0          int    `json:"drop-vld-invalid-flow-label-v0"`
	DropVldInvalidTeid                 int    `json:"drop-vld-invalid-teid"`
	DropVldUnsupportedMessageType      int    `json:"drop-vld-unsupported-message-type"`
	DropVldOutOfState                  int    `json:"drop-vld-out-of-state"`
	DropVldMandatoryInformationElement int    `json:"drop-vld-mandatory-information-element"`
	DropVldOutOfOrderIe                int    `json:"drop-vld-out-of-order-ie"`
	DropVldOutOfStateIe                int    `json:"drop-vld-out-of-state-ie"`
	DropVldReservedInformationElement  int    `json:"drop-vld-reserved-information-element"`
	DropVldVersionNotSupported         int    `json:"drop-vld-version-not-supported"`
	DropVldMessageLength               int    `json:"drop-vld-message-length"`
	DropVldCrossLayerCorrelation       int    `json:"drop-vld-cross-layer-correlation"`
	DropVldCountryCodeMismatch         int    `json:"drop-vld-country-code-mismatch"`
	DropVldGtpUSpoofedSourceAddress    int    `json:"drop-vld-gtp-u-spoofed-source-address"`
	Drop_vldGtpBearerCountExceed       int    `json:"drop_vld-gtp-bearer-count-exceed"`
	DropVldGtpV2WrongLbiCreateBearer   int    `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
	GtpCHandoverInProgressWithConn     int    `json:"gtp-c-handover-in-progress-with-conn"`
	DropVldInvalidPktLenPiggyback      int    `json:"drop-vld-invalid-pkt-len-piggyback"`
	DropVldSanityFailedPiggyback       int    `json:"drop-vld-sanity-failed-piggyback"`
	DropVldSequenceNumCorrelation      int    `json:"drop-vld-sequence-num-correlation"`
	DropVldGtpv0SeqnumBufferFull       int    `json:"drop-vld-gtpv0-seqnum-buffer-full"`
	DropVldGtpv1SeqnumBufferFull       int    `json:"drop-vld-gtpv1-seqnum-buffer-full"`
	DropVldGtpv2SeqnumBufferFull       int    `json:"drop-vld-gtpv2-seqnum-buffer-full"`
	DropVldGtpInvalidImsiLenDrop       int    `json:"drop-vld-gtp-invalid-imsi-len-drop"`
	DropVldGtpInvalidApnLenDrop        int    `json:"drop-vld-gtp-invalid-apn-len-drop"`
	DropVldProtocolFlagUnset           int    `json:"drop-vld-protocol-flag-unset"`
	DropFltMessageFiltering            int    `json:"drop-flt-message-filtering"`
	DropFltApnFiltering                int    `json:"drop-flt-apn-filtering"`
	DropFltMsisdnFiltering             int    `json:"drop-flt-msisdn-filtering"`
	DropFltRatTypeFiltering            int    `json:"drop-flt-rat-type-filtering"`
	DropFltGtpInGtp                    int    `json:"drop-flt-gtp-in-gtp"`
	DropRlGtpV0CAgg                    int    `json:"drop-rl-gtp-v0-c-agg"`
	DropRlGtpV1CAgg                    int    `json:"drop-rl-gtp-v1-c-agg"`
	DropRlGtpV2CAgg                    int    `json:"drop-rl-gtp-v2-c-agg"`
	DropRlGtpV1CCreatePdpRequest       int    `json:"drop-rl-gtp-v1-c-create-pdp-request"`
	DropRlGtpV2CCreateSessionRequest   int    `json:"drop-rl-gtp-v2-c-create-session-request"`
	DropRlGtpV1CUpdatePdpRequest       int    `json:"drop-rl-gtp-v1-c-update-pdp-request"`
	DropRlGtpV2CModifyBearerRequest    int    `json:"drop-rl-gtp-v2-c-modify-bearer-request"`
	DropRlGtpUTunnelCreate             int    `json:"drop-rl-gtp-u-tunnel-create"`
	DropRlGtpUUplinkByte               int    `json:"drop-rl-gtp-u-uplink-byte"`
	DropRlGtpUUplinkPacket             int    `json:"drop-rl-gtp-u-uplink-packet"`
	DropRlGtpUDownlinkByte             int    `json:"drop-rl-gtp-u-downlink-byte"`
	DropRlGtpUDownlinkPacket           int    `json:"drop-rl-gtp-u-downlink-packet"`
	DropRlGtpUTotalByte                int    `json:"drop-rl-gtp-u-total-byte"`
	DropRlGtpUTotalPacket              int    `json:"drop-rl-gtp-u-total-packet"`
	DropRlGtpUMaxConcurrentTunnels     int    `json:"drop-rl-gtp-u-max-concurrent-tunnels"`
	RlMessageMonitor                   int    `json:"rl-message-monitor"`
	Uuid                               string `json:"uuid"`
}

type VisibilityTopnGtpNetworkElementTopnNode3744 struct {
	Activate string `json:"activate"`
	Uuid     string `json:"uuid"`
}

type VisibilityZbar3745 struct {
	Action  string                    `json:"action" dval:"disable"`
	Uuid    string                    `json:"uuid"`
	Dest    VisibilityZbarDest3746    `json:"dest"`
	Truples VisibilityZbarTruples3748 `json:"truples"`
}

type VisibilityZbarDest3746 struct {
	Uuid       string                           `json:"uuid"`
	BadSources VisibilityZbarDestBadSources3747 `json:"bad-sources"`
}

type VisibilityZbarDestBadSources3747 struct {
	Uuid string `json:"uuid"`
}

type VisibilityZbarTruples3748 struct {
	Uuid string `json:"uuid"`
}

func (p *Visibility) GetId() string {
	return "1"
}

func (p *Visibility) getPath() string {
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
		if len(axResp) > 0 {
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
