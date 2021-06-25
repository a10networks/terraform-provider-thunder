package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type System struct {
	AnomalyLog SystemInstance `json:"system,omitempty"`
}

type SystemInstance struct {
	//CoreIndex  AddCPUCore   `json:"add-cpu-core,omitempty"`
	//PortIndex  AddPort      `json:"add-port,omitempty"`
	//Bcast      AllVlanLimit `json:"all-vlan-limit,omitempty"`
	AnomalyLog int `json:"anomaly-log,omitempty"`
	//UUID                        AppPerformance      `json:"app-performance,omitempty"`
	//LogSessionOnEstablished AppsGlobal `json:"apps-global,omitempty"`
	AttackLog               int        `json:"attack-log,omitempty"`
	//UUID                        AppPerformance      `json:"bandwidth,omitempty"`
	//UUID                        AppPerformance      `json:"bfd,omitempty"`
	//Enable CPUHyperThread `json:"cpu-hyper-thread,omitempty"`
	//UUID                        CPUList             `json:"cpu-list,omitempty"`
	//Disable                     CPULoadSharing      `json:"cpu-load-sharing,omitempty"`
	//UUID                        CPUList             `json:"cpu-map,omitempty"`
	ClassListHitcountEnable int `json:"class-list-hitcount-enable,omitempty"`
	//SourceName              CmUpdateFileNameRef `json:"cm-update-file-name-ref,omitempty"`
	//UUID                        CPUList             `json:"control-cpu,omitempty"`
	//UUID                        CPUList             `json:"core,omitempty"`
	//UUID                        CPUList             `json:"cosq-show,omitempty"`
	//UUID                        CPUList             `json:"cosq-stats,omitempty"`
	//UUID                        CPUList             `json:"counter-lib-accounting,omitempty"`
	//UUID                        AppPerformance      `json:"dns,omitempty"`
	//UUID                        AppPerformance      `json:"dns-cache,omitempty"`
	//UUID                        CPUList             `json:"data-cpu,omitempty"`
	DdosAttack int `json:"ddos-attack,omitempty"`
	DdosLog    int `json:"ddos-log,omitempty"`
	//Enable                      DeepHrxq            `json:"deep-hrxq,omitempty"`
	//PortIndex1               AddPort1    `json:"del-port,omitempty"`
	//CoreIndex1               AddCPUCore1 `json:"delete-cpu-core,omitempty"`
	DomainListHitcountEnable int `json:"domain-list-hitcount-enable,omitempty"`
	//UUID                        CPUList             `json:"domain-list-info,omitempty"`
	//UUID                        AppPerformance      `json:"dpdk-stats,omitempty"`
	DynamicServiceDNSSocketPool int `json:"dynamic-service-dns-socket-pool,omitempty"`
	//UUID                        CPUList             `json:"environment,omitempty"`
	//MonitorDisable FpgaCoreCrc `json:"fpga-core-crc,omitempty"`
	//UUID                        AppPerformance      `json:"fpga-drop,omitempty"`
	//ApplicationMempool  Fw          `json:"fw,omitempty"`
	GeoDbHitcountEnable int `json:"geo-db-hitcount-enable,omitempty"`
	//GeoLocationIana     GeoLocation `json:"geo-location,omitempty"`
	//UUID                        AppPerformance      `json:"geoloc,omitempty"`
	//Name                        []GeolocListList    `json:"geoloc-list-list,omitempty"`
	//UUID                        AppPerformance      `json:"geoloc-name-helper,omitempty"`
	//UUID                        GeolocationFile     `json:"geolocation-file,omitempty"`
	Glid int `json:"glid,omitempty"`
	//UUID                        CPUList             `json:"guest-file,omitempty"`
	//UUID                        CPUList             `json:"gui-image-list,omitempty"`
	//UUID                        CPUList             `json:"hardware,omitempty"`
	//UUID                        HardwareForward     `json:"hardware-forward,omitempty"`
	//UUID                        CPUList             `json:"hrxq-status,omitempty"`
	//UUID                        AppPerformance      `json:"ip6-stats,omitempty"`
	//UUID                        CPUList             `json:"ip-dns-cache,omitempty"`
	//UUID                        AppPerformance      `json:"ip-stats,omitempty"`
	//UUID                        IPThreatList        `json:"ip-threat-list,omitempty"`
	//UUID                        AppPerformance      `json:"icmp,omitempty"`
	//UUID                        AppPerformance      `json:"icmp6,omitempty"`
	//UUID                        AppPerformance      `json:"icmp-rate,omitempty"`
	//UUID                        CPUList             `json:"inuse-cpu-list,omitempty"`
	//UUID                        CPUList             `json:"inuse-port-list,omitempty"`
	//MaxCores IoCPU `json:"io-cpu,omitempty"`
	//Reset                       Ipmi                `json:"ipmi,omitempty"`
	//Disable                     IpmiService         `json:"ipmi-service,omitempty"`
	//PacketRoundRobin Ipsec           `json:"ipsec,omitempty"`
	//Enable1           LinkCapability  `json:"link-capability,omitempty"`
	//Disable1         CPUHyperThread1 `json:"link-monitor,omitempty"`
	//UUID                        AppPerformance      `json:"memory,omitempty"`
	//MacAddress     MgmtPort    `json:"mgmt-port,omitempty"`
	//PortNumber     ModifyPort  `json:"modify-port,omitempty"`
	ModuleCtrlCPU string `json:"module-ctrl-cpu,omitempty"`
	//MonitorList   MonTemplate `json:"mon-template,omitempty"`
	//DeepHrxq1Enable1       DeepHrxq1   `json:"multi-queue-support,omitempty"`
	//UUID                        AppPerformance      `json:"ndisc-ra,omitempty"`
	//Complexity PasswordPolicy `json:"password-policy,omitempty"`
	//Bcast      AllVlanLimit   `json:"per-vlan-limit,omitempty"`
	//UUID                        CPUList             `json:"platformtype,omitempty"`
	//UUID                        CPUList             `json:"port-info,omitempty"`
	//UUID                        CPUList             `json:"port-list,omitempty"`
	//LinkDetectionInterval Ports `json:"ports,omitempty"`
	PromiscuousMode int `json:"promiscuous-mode,omitempty"`
	//UUID                        CPUList             `json:"psu-info,omitempty"`
	//Enable                      LinkCapability      `json:"queuing-buffer,omitempty"`
	//Server                      Radius              `json:"radius,omitempty"`
	//UUID                        CPUList             `json:"reboot,omitempty"`
	//UUID                        ResourceAccounting  `json:"resource-accounting,omitempty"`
	//SslContextMemory SystemResourceUsage `json:"resource-usage,omitempty"`
	//UUID                        AppPerformance      `json:"session,omitempty"`
	//NscanLimit SessionReclaimLimit `json:"session-reclaim-limit,omitempty"`
	//TxdSize  SetRxtxDescSize     `json:"set-rxtx-desc-size,omitempty"`
	//RxqSize  SetRxtxQueue        `json:"set-rxtx-queue,omitempty"`
	//Enable2    CPUHyperThread2     `json:"shared-poll-mode,omitempty"`
	//UUID                        CPUList             `json:"shell-privileges,omitempty"`
	//UUID                        CPUList             `json:"shutdown,omitempty"`
	SockstressDisable int `json:"sockstress-disable,omitempty"`
	//Action            SpeProfile `json:"spe-profile,omitempty"`
	//UUID                        CPUList             `json:"spe-status,omitempty"`
	SrcIPHashEnable int `json:"src-ip-hash-enable,omitempty"`
	//UUID                        AppPerformance      `json:"ssl-req-q,omitempty"`
	//Enable                      LinkCapability      `json:"ssl-scv,omitempty"`
	//UUID                        CPUList             `json:"ssl-status,omitempty"`
	//EnableFlag SyslogTimeMsec `json:"syslog-time-msec,omitempty"`
	//UUID                        TCP                 `json:"tcp,omitempty"`
	//UUID                        AppPerformance      `json:"tcp-stats,omitempty"`
	//DeepHrxq2Enable2        DeepHrxq2      `json:"tls-1-3-mgmt,omitempty"`
	//TopKSourceList TelemetryLog   `json:"telemetry-log,omitempty"`
	//TemplatePolicy SystemTemplate `json:"template,omitempty"`
	//LinkdownEthernet1    TemplateBind   `json:"template-bind,omitempty"`
	//UUID                        AppPerformance      `json:"throughput,omitempty"`
	//Ftp         TimeoutValue `json:"timeout-value,omitempty"`
	//LoadBalance Trunk        `json:"trunk,omitempty"`
	//Mode                        TrunkHwHash         `json:"trunk-hw-hash,omitempty"`
	//Mode                        TrunkHwHash         `json:"trunk-xaui-hw-hash,omitempty"`
	UUID string `json:"uuid,omitempty"`
	//UUID                        CPUList             `json:"upgrade-status,omitempty"`
	//VeMacSchemeVal VeMacScheme `json:"ve-mac-scheme,omitempty"`
}

type AddCPUCore struct {
	CoreIndex int `json:"core-index,omitempty"`
}

type AddCPUCore1 struct {
	CoreIndex1 int `json:"core-index,omitempty"`
}

type AddPort struct {
	PortIndex int `json:"port-index,omitempty"`
}

type AddPort1 struct {
	PortIndex1 int `json:"port-index,omitempty"`
}

type AllVlanLimit struct {
	Bcast        int    `json:"bcast,omitempty"`
	Ipmcast      int    `json:"ipmcast,omitempty"`
	Mcast        int    `json:"mcast,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	UnknownUcast int    `json:"unknown-ucast,omitempty"`
}

// type AppPerformance struct {
// 	Counters1 []SamplingEnable `json:"sampling-enable,omitempty"`
// 	UUID      string           `json:"uuid,omitempty"`
// }

type AppsGlobal struct {
	LogSessionOnEstablished int    `json:"log-session-on-established,omitempty"`
	MslTime                 int    `json:"msl-time,omitempty"`
	UUID                    string `json:"uuid,omitempty"`
}

type CPUHyperThread struct {
	Disable int `json:"disable,omitempty"`
	Enable  int `json:"enable,omitempty"`
}

type CPUHyperThread1 struct {
	Disable1 int `json:"disable,omitempty"`
	Enable1  int `json:"enable,omitempty"`
}

type CPUHyperThread2 struct {
	Disable2 int `json:"disable,omitempty"`
	Enable2  int `json:"enable,omitempty"`
}

type CPUList struct {
	UUID string `json:"uuid,omitempty"`
}

// type CPULoadSharing struct {
// 	Low     CPUUsage         `json:"cpu-usage,omitempty"`
// 	Disable int              `json:"disable,omitempty"`
// //	Min     PacketsPerSecond `json:"packets-per-second,omitempty"`
// 	UUID    string           `json:"uuid,omitempty"`
// }

type CmUpdateFileNameRef struct {
	DestName   string `json:"dest_name,omitempty"`
	ID         int    `json:"id,omitempty"`
	SourceName string `json:"source_name,omitempty"`
}

type DeepHrxq struct {
	Enable int `json:"enable,omitempty"`
}

type DeepHrxq1 struct {
	DeepHrxq1Enable1 int `json:"enable,omitempty"`
}

type DeepHrxq2 struct {
	DeepHrxq2Enable2 int `json:"enable,omitempty"`
}

type FpgaCoreCrc struct {
	MonitorDisable int    `json:"monitor-disable,omitempty"`
	RebootEnable   int    `json:"reboot-enable,omitempty"`
	UUID           string `json:"uuid,omitempty"`
}

type Fw struct {
	ApplicationFlow    int    `json:"application-flow,omitempty"`
	ApplicationMempool int    `json:"application-mempool,omitempty"`
	BasicDpiEnable     int    `json:"basic-dpi-enable,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}

type GeoLocation struct {
	GeoLocnObjName             []EntryList          `json:"entry-list,omitempty"`
	GeoLocationGeolite2City    int                  `json:"geo-location-geolite2-city,omitempty"`
	GeoLocationGeolite2Country int                  `json:"geo-location-geolite2-country,omitempty"`
	GeoLocationIana            int                  `json:"geo-location-iana,omitempty"`
	Geolite2CityIncludeIpv6    int                  `json:"geolite2-city-include-ipv6,omitempty"`
	Geolite2CountryIncludeIpv6 int                  `json:"geolite2-country-include-ipv6,omitempty"`
	GeoLocationLoadFilename    []GeolocLoadFileList `json:"geoloc-load-file-list,omitempty"`
	UUID                       string               `json:"uuid,omitempty"`
}

type GeolocListList struct {
	ExcludeGeolocNameVal []ExcludeGeolocNameList `json:"exclude-geoloc-name-list,omitempty"`
	IncludeGeolocNameVal []IncludeGeolocNameList `json:"include-geoloc-name-list,omitempty"`
	Name                 string                  `json:"name,omitempty"`
	Counters1            []SamplingEnable        `json:"sampling-enable,omitempty"`
	Shared               int                     `json:"shared,omitempty"`
	UUID                 string                  `json:"uuid,omitempty"`
	UserTag              string                  `json:"user-tag,omitempty"`
}

// type GeolocationFile struct {
// 	UUID CPUList `json:"error-info,omitempty"`
// 	UUID string  `json:"uuid,omitempty"`
// }

type HardwareForward struct {
	Counters1 []SamplingEnable `json:"sampling-enable,omitempty"`
	//UUID      AppPerformance   `json:"slb,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// type IPThreatList struct {
// 	ClassListCfg Ipv4DestList     `json:"ipv4-dest-list,omitempty"`
// 	ClassListCfg Ipv4DestList     `json:"ipv4-source-list,omitempty"`
// 	ClassListCfg Ipv4DestList     `json:"ipv6-dest-list,omitempty"`
// 	ClassListCfg Ipv4DestList     `json:"ipv6-source-list,omitempty"`
// 	Counters1    []SamplingEnable `json:"sampling-enable,omitempty"`
// 	UUID         string           `json:"uuid,omitempty"`
// }

type IoCPU struct {
	MaxCores int `json:"max-cores,omitempty"`
}

// type Ipmi struct {
// 	Ipv4Address IP    `json:"ip,omitempty"`
// 	Dhcp        Ipsrc `json:"ipsrc,omitempty"`
// 	Reset       int   `json:"reset,omitempty"`
// 	Cmd         Tool  `json:"tool,omitempty"`
// 	Add         User  `json:"user,omitempty"`
// }

type IpmiService struct {
	Disable int    `json:"disable,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type Ipsec struct {
	CryptoCore       int        `json:"crypto-core,omitempty"`
	CryptoMem        int        `json:"crypto-mem,omitempty"`
	Action           SpeProfile `json:"fpga-decrypt,omitempty"`
	PacketRoundRobin int        `json:"packet-round-robin,omitempty"`
	UUID             string     `json:"uuid,omitempty"`
}

type LinkCapability struct {
	Enable1 int    `json:"enable,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type MgmtPort struct {
	MacAddress string `json:"mac-address,omitempty"`
	PciAddress string `json:"pci-address,omitempty"`
	PortIndex  int    `json:"port-index,omitempty"`
}

type ModifyPort struct {
	PortIndex  int `json:"port-index,omitempty"`
	PortNumber int `json:"port-number,omitempty"`
}

type MonTemplate struct {
	Enable1 LinkCapability `json:"link-block-as-down,omitempty"`
	ID      []MonitorList  `json:"monitor-list,omitempty"`
}

type PasswordPolicy struct {
	Aging      string `json:"aging,omitempty"`
	Complexity string `json:"complexity,omitempty"`
	History    string `json:"history,omitempty"`
	MinPswdLen int    `json:"min-pswd-len,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}

type Ports struct {
	LinkDetectionInterval int    `json:"link-detection-interval,omitempty"`
	UUID                  string `json:"uuid,omitempty"`
}

// type Radius struct {
// 	ListenPort Server `json:"server,omitempty"`
// }

type ResourceAccounting struct {
	Name []TemplateList `json:"template-list,omitempty"`
	UUID string         `json:"uuid,omitempty"`
}

type SystemResourceUsage struct {
	AflexTableEntryCount          int        `json:"aflex-table-entry-count,omitempty"`
	AuthPortalHTMLFileSize        int        `json:"auth-portal-html-file-size,omitempty"`
	AuthPortalImageFileSize       int        `json:"auth-portal-image-file-size,omitempty"`
	AuthzPolicyNumber             int        `json:"authz-policy-number,omitempty"`
	ClassListAcEntryCount         int        `json:"class-list-ac-entry-count,omitempty"`
	ClassListIpv6AddrCount        int        `json:"class-list-ipv6-addr-count,omitempty"`
	IpsecSaNumber                 int        `json:"ipsec-sa-number,omitempty"`
	L4SessionCount                int        `json:"l4-session-count,omitempty"`
	MaxAflexAuthzCollectionNumber int        `json:"max-aflex-authz-collection-number,omitempty"`
	MaxAflexFileSize              int        `json:"max-aflex-file-size,omitempty"`
	NatPoolAddrCount              int        `json:"nat-pool-addr-count,omitempty"`
	RAMCacheMemoryLimit           int        `json:"ram-cache-memory-limit,omitempty"`
	RadiusTableSize               int        `json:"radius-table-size,omitempty"`
	SslContextMemory              int        `json:"ssl-context-memory,omitempty"`
	SslDmaMemory                  int        `json:"ssl-dma-memory,omitempty"`
	UUID                          string     `json:"uuid,omitempty"`
	MonitoredEntityCount          Visibility `json:"visibility,omitempty"`
}

type SessionReclaimLimit struct {
	NscanLimit int    `json:"nscan-limit,omitempty"`
	ScanFreq   int    `json:"scan-freq,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}

type SetRxtxDescSize struct {
	PortIndex int `json:"port-index,omitempty"`
	RxdSize   int `json:"rxd-size,omitempty"`
	TxdSize   int `json:"txd-size,omitempty"`
}

type SetRxtxQueue struct {
	PortIndex int `json:"port-index,omitempty"`
	RxqSize   int `json:"rxq-size,omitempty"`
	TxqSize   int `json:"txq-size,omitempty"`
}

type SpeProfile struct {
	Action string `json:"action,omitempty"`
}

type SyslogTimeMsec struct {
	EnableFlag int `json:"enable-flag,omitempty"`
}

// type TCP struct {
// 	PktRateForResetUnknownConn RateLimitResetUnknownConn `json:"rate-limit-reset-unknown-conn,omitempty"`
// 	Counters1                  []SamplingEnable          `json:"sampling-enable,omitempty"`
// 	UUID                       string                    `json:"uuid,omitempty"`
// }

// type TelemetryLog struct {
// 	UUID CPUList `json:"device-status,omitempty"`
// 	UUID CPUList `json:"environment,omitempty"`
// 	UUID CPUList `json:"partition-metrics,omitempty"`
// 	UUID CPUList `json:"top-k-app-svc-list,omitempty"`
// 	UUID CPUList `json:"top-k-source-list,omitempty"`
// }

type SystemTemplate struct {
	TemplatePolicy string `json:"template-policy,omitempty"`
	UUID           string `json:"uuid,omitempty"`
}

type TemplateBind struct {
	LinkdownEthernet1 []MonitorList `json:"monitor-list,omitempty"`
}

type TimeoutValue struct {
	Ftp   int    `json:"ftp,omitempty"`
	HTTP  int    `json:"http,omitempty"`
	HTTPS int    `json:"https,omitempty"`
	Scp   int    `json:"scp,omitempty"`
	Sftp  int    `json:"sftp,omitempty"`
	Tftp  int    `json:"tftp,omitempty"`
	UUID  string `json:"uuid,omitempty"`
}

type Trunk struct {
	UseL3 LoadBalance `json:"load-balance,omitempty"`
}

type TrunkHwHash struct {
	Mode int    `json:"mode,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type VeMacScheme struct {
	UUID           string `json:"uuid,omitempty"`
	VeMacSchemeVal string `json:"ve-mac-scheme-val,omitempty"`
}

type SystemSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type CPUUsage struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}

// type PacketsPerSecond struct {
// 	Min int `json:"min,omitempty"`
// }

type EntryList struct {
	FirstIPAddress []GeoLocnMultipleAddresses `json:"geo-locn-multiple-addresses,omitempty"`
	GeoLocnObjName string                     `json:"geo-locn-obj-name,omitempty"`
	UUID           string                     `json:"uuid,omitempty"`
	UserTag        string                     `json:"user-tag,omitempty"`
}

type GeolocLoadFileList struct {
	GeoLocationLoadFilename string `json:"geo-location-load-filename,omitempty"`
	TemplateName            string `json:"template-name,omitempty"`
}

type ExcludeGeolocNameList struct {
	ExcludeGeolocNameVal string `json:"exclude-geoloc-name-val,omitempty"`
}

type IncludeGeolocNameList struct {
	IncludeGeolocNameVal string `json:"include-geoloc-name-val,omitempty"`
}

type Ipv4DestList struct {
	ClassList []ClassListCfg `json:"class-list-cfg,omitempty"`
	UUID      string         `json:"uuid,omitempty"`
}

// type IP struct {
// 	DefaultGateway string `json:"default-gateway,omitempty"`
// 	Ipv4Address    string `json:"ipv4-address,omitempty"`
// 	Ipv4Netmask    string `json:"ipv4-netmask,omitempty"`
// }

type Ipsrc struct {
	Dhcp   int `json:"dhcp,omitempty"`
	Static int `json:"static,omitempty"`
}

type Tool struct {
	Cmd string `json:"cmd,omitempty"`
}

type User struct {
	Add           string `json:"add,omitempty"`
	Administrator int    `json:"administrator,omitempty"`
	Callback      int    `json:"callback,omitempty"`
	Disable       string `json:"disable,omitempty"`
	Newname       string `json:"newname,omitempty"`
	Newpass       string `json:"newpass,omitempty"`
	Operator      int    `json:"operator,omitempty"`
	Password      string `json:"password,omitempty"`
	Privilege     string `json:"privilege,omitempty"`
	Setname       string `json:"setname,omitempty"`
	Setpass       string `json:"setpass,omitempty"`
	User          int    `json:"user,omitempty"`
}

type MonitorList struct {
	Sessions          []SystemClearCfg       `json:"clear-cfg,omitempty"`
	ID                int                    `json:"id,omitempty"`
	Diseth            []SystemLinkDisableCfg `json:"link-disable-cfg,omitempty"`
	LinkdownEthernet1 []SystemLinkDownCfg    `json:"link-down-cfg,omitempty"`
	Enaeth            []SystemLinkEnableCfg  `json:"link-enable-cfg,omitempty"`
	LinkupEthernet1   []SystemLinkUpCfg      `json:"link-up-cfg,omitempty"`
	MonitorRelation   string                 `json:"monitor-relation,omitempty"`
	UUID              string                 `json:"uuid,omitempty"`
	UserTag           string                 `json:"user-tag,omitempty"`
}

// type Server struct {
// 	AccountingInterimUpdate string           `json:"accounting-interim-update,omitempty"`
// 	AccountingOn            string           `json:"accounting-on,omitempty"`
// 	AccountingStart         string           `json:"accounting-start,omitempty"`
// 	AccountingStop          string           `json:"accounting-stop,omitempty"`
// 	AttributeValue          []Attribute      `json:"attribute,omitempty"`
// 	AttributeName           string           `json:"attribute-name,omitempty"`
// 	DisableReply            int              `json:"disable-reply,omitempty"`
// 	ListenPort              int              `json:"listen-port,omitempty"`
// 	IPList                  Remote           `json:"remote,omitempty"`
// 	Counters1               []SamplingEnable `json:"sampling-enable,omitempty"`
// 	Secret                  int              `json:"secret,omitempty"`
// 	SecretString            string           `json:"secret-string,omitempty"`
// 	UUID                    string           `json:"uuid,omitempty"`
// 	Vrid                    int              `json:"vrid,omitempty"`
// }

type TemplateList struct {
	GslbDeviceCfg      AppResources     `json:"app-resources,omitempty"`
	Name               string           `json:"name,omitempty"`
	StaticIpv4RouteCfg NetworkResources `json:"network-resources,omitempty"`
	BwLimitCfg         SystemResources  `json:"system-resources,omitempty"`
	UUID               string           `json:"uuid,omitempty"`
	UserTag            string           `json:"user-tag,omitempty"`
}

type Visibility struct {
	MonitoredEntityCount int    `json:"monitored-entity-count,omitempty"`
	UUID                 string `json:"uuid,omitempty"`
}

type RateLimitResetUnknownConn struct {
	LogForResetUnknownConn     int    `json:"log-for-reset-unknown-conn,omitempty"`
	PktRateForResetUnknownConn int    `json:"pkt-rate-for-reset-unknown-conn,omitempty"`
	UUID                       string `json:"uuid,omitempty"`
}

type LoadBalance struct {
	UUID  string `json:"uuid,omitempty"`
	UseL3 int    `json:"use-l3,omitempty"`
	UseL4 int    `json:"use-l4,omitempty"`
}

type GeoLocnMultipleAddresses struct {
	FirstIPAddress   string `json:"first-ip-address,omitempty"`
	FirstIpv6Address string `json:"first-ipv6-address,omitempty"`
	GeolIpv4Mask     string `json:"geol-ipv4-mask,omitempty"`
	GeolIpv6Mask     int    `json:"geol-ipv6-mask,omitempty"`
	IPAddr2          string `json:"ip-addr2,omitempty"`
	Ipv6Addr2        string `json:"ipv6-addr2,omitempty"`
}

type ClassListCfg struct {
	ClassList          string `json:"class-list,omitempty"`
	IPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemClearCfg struct {
	ClearAllSequence int    `json:"clear-all-sequence,omitempty"`
	ClearSequence    int    `json:"clear-sequence,omitempty"`
	Sessions         string `json:"sessions,omitempty"`
}

type SystemLinkDisableCfg struct {
	DisSequence int `json:"dis-sequence,omitempty"`
	Diseth      int `json:"diseth,omitempty"`
}

type SystemLinkDownCfg struct {
	LinkDownSequence1 int `json:"link-down-sequence1,omitempty"`
	LinkDownSequence2 int `json:"link-down-sequence2,omitempty"`
	LinkDownSequence3 int `json:"link-down-sequence3,omitempty"`
	LinkdownEthernet1 int `json:"linkdown-ethernet1,omitempty"`
	LinkdownEthernet2 int `json:"linkdown-ethernet2,omitempty"`
	LinkdownEthernet3 int `json:"linkdown-ethernet3,omitempty"`
}

type SystemLinkEnableCfg struct {
	EnaSequence int `json:"ena-sequence,omitempty"`
	Enaeth      int `json:"enaeth,omitempty"`
}

type SystemLinkUpCfg struct {
	LinkUpSequence1 int `json:"link-up-sequence1,omitempty"`
	LinkUpSequence2 int `json:"link-up-sequence2,omitempty"`
	LinkUpSequence3 int `json:"link-up-sequence3,omitempty"`
	LinkupEthernet1 int `json:"linkup-ethernet1,omitempty"`
	LinkupEthernet2 int `json:"linkup-ethernet2,omitempty"`
	LinkupEthernet3 int `json:"linkup-ethernet3,omitempty"`
}

type Attribute struct {
	AttributeValue string `json:"attribute-value,omitempty"`
	CustomNumber   int    `json:"custom-number,omitempty"`
	CustomVendor   int    `json:"custom-vendor,omitempty"`
	Name           string `json:"name,omitempty"`
	Number         int    `json:"number,omitempty"`
	PrefixLength   string `json:"prefix-length,omitempty"`
	PrefixNumber   int    `json:"prefix-number,omitempty"`
	PrefixVendor   int    `json:"prefix-vendor,omitempty"`
	Value          string `json:"value,omitempty"`
	Vendor         int    `json:"vendor,omitempty"`
}

type Remote struct {
	IPListName []IPList `json:"ip-list,omitempty"`
}

type AppResources struct {
	CacheTemplateMax         CacheTemplateCfg         `json:"cache-template-cfg,omitempty"`
	ClientSslTemplateMax     ClientSslTemplateCfg     `json:"client-ssl-template-cfg,omitempty"`
	ConnReuseTemplateMax     ConnReuseTemplateCfg     `json:"conn-reuse-template-cfg,omitempty"`
	FastTCPTemplateMax       FastTCPTemplateCfg       `json:"fast-tcp-template-cfg,omitempty"`
	FastUDPTemplateMax       FastUDPTemplateCfg       `json:"fast-udp-template-cfg,omitempty"`
	FixTemplateMax           FixTemplateCfg           `json:"fix-template-cfg,omitempty"`
	GslbDeviceMax            GslbDeviceCfg            `json:"gslb-device-cfg,omitempty"`
	GslbGeoLocationMax       GslbGeoLocationCfg       `json:"gslb-geo-location-cfg,omitempty"`
	GslbIPListMax            GslbIPListCfg            `json:"gslb-ip-list-cfg,omitempty"`
	GslbPolicyMax            GslbPolicyCfg            `json:"gslb-policy-cfg,omitempty"`
	GslbServiceMax           GslbServiceCfg           `json:"gslb-service-cfg,omitempty"`
	GslbServiceIPMax         GslbServiceIPCfg         `json:"gslb-service-ip-cfg,omitempty"`
	GslbServicePortMax       GslbServicePortCfg       `json:"gslb-service-port-cfg,omitempty"`
	GslbSiteMax              GslbSiteCfg              `json:"gslb-site-cfg,omitempty"`
	GslbSvcGroupMax          GslbSvcGroupCfg          `json:"gslb-svc-group-cfg,omitempty"`
	GslbTemplateMax          GslbTemplateCfg          `json:"gslb-template-cfg,omitempty"`
	GslbZoneMax              GslbZoneCfg              `json:"gslb-zone-cfg,omitempty"`
	HTTPTemplateMax          HTTPTemplateCfg          `json:"http-template-cfg,omitempty"`
	HealthMonitorMax         HealthMonitorCfg         `json:"health-monitor-cfg,omitempty"`
	LinkCostTemplateMax      LinkCostTemplateCfg      `json:"link-cost-template-cfg,omitempty"`
	PersistCookieTemplateMax PersistCookieTemplateCfg `json:"persist-cookie-template-cfg,omitempty"`
	PersistSrcipTemplateMax  PersistSrcipTemplateCfg  `json:"persist-srcip-template-cfg,omitempty"`
	ProxyTemplateMax         ProxyTemplateCfg         `json:"proxy-template-cfg,omitempty"`
	RealPortMax              RealPortCfg              `json:"real-port-cfg,omitempty"`
	RealServerMax            RealServerCfg            `json:"real-server-cfg,omitempty"`
	ServerSslTemplateMax     ServerSslTemplateCfg     `json:"server-ssl-template-cfg,omitempty"`
	ServiceGroupMax          ServiceGroupCfg          `json:"service-group-cfg,omitempty"`
	StreamTemplateMax        StreamTemplateCfg        `json:"stream-template-cfg,omitempty"`
	Threshold                int                      `json:"threshold,omitempty"`
	UUID                     string                   `json:"uuid,omitempty"`
	VirtualPortMax           VirtualPortCfg           `json:"virtual-port-cfg,omitempty"`
	VirtualServerMax         VirtualServerCfg         `json:"virtual-server-cfg,omitempty"`
}

type NetworkResources struct {
	Ipv4AclLineMax       Ipv4AclLineCfg       `json:"ipv4-acl-line-cfg,omitempty"`
	Ipv6AclLineMax       Ipv6AclLineCfg       `json:"ipv6-acl-line-cfg,omitempty"`
	ObjectGroupMax       ObjectGroupCfg       `json:"object-group-cfg,omitempty"`
	ObjectGroupClauseMax ObjectGroupClauseCfg `json:"object-group-clause-cfg,omitempty"`
	StaticArpMax         StaticArpCfg         `json:"static-arp-cfg,omitempty"`
	StaticIpv4RouteMax   StaticIpv4RouteCfg   `json:"static-ipv4-route-cfg,omitempty"`
	StaticIpv6RouteMax   StaticIpv6RouteCfg   `json:"static-ipv6-route-cfg,omitempty"`
	StaticMacMax         StaticMacCfg         `json:"static-mac-cfg,omitempty"`
	StaticNeighborMax    StaticNeighborCfg    `json:"static-neighbor-cfg,omitempty"`
	Threshold            int                  `json:"threshold,omitempty"`
	UUID                 string               `json:"uuid,omitempty"`
}

type SystemResources struct {
	BwLimitMax                BwLimitCfg                `json:"bw-limit-cfg,omitempty"`
	ConcurrentSessionLimitMax ConcurrentSessionLimitCfg `json:"concurrent-session-limit-cfg,omitempty"`
	FwcpsLimitMax             FwcpsLimitCfg             `json:"fwcps-limit-cfg,omitempty"`
	L4CpsLimitMax             L4CpsLimitCfg             `json:"l4cps-limit-cfg,omitempty"`
	L4SessionLimitMax         L4SessionLimitCfg         `json:"l4-session-limit-cfg,omitempty"`
	L7CpsLimitMax             L7CpsLimitCfg             `json:"l7cps-limit-cfg,omitempty"`
	NatcpsLimitMax            NatcpsLimitCfg            `json:"natcps-limit-cfg,omitempty"`
	SslThroughputLimitMax     SslThroughputLimitCfg     `json:"ssl-throughput-limit-cfg,omitempty"`
	SslcpsLimitMax            SslcpsLimitCfg            `json:"sslcps-limit-cfg,omitempty"`
	Threshold                 int                       `json:"threshold,omitempty"`
	UUID                      string                    `json:"uuid,omitempty"`
}

type IPList struct {
	IPListName         string `json:"ip-list-name,omitempty"`
	IPListSecret       int    `json:"ip-list-secret,omitempty"`
	IPListSecretString string `json:"ip-list-secret-string,omitempty"`
}

type CacheTemplateCfg struct {
	CacheTemplateMax          int `json:"cache-template-max,omitempty"`
	CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee,omitempty"`
}

type ClientSslTemplateCfg struct {
	ClientSslTemplateMax          int `json:"client-ssl-template-max,omitempty"`
	ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee,omitempty"`
}

type ConnReuseTemplateCfg struct {
	ConnReuseTemplateMax          int `json:"conn-reuse-template-max,omitempty"`
	ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee,omitempty"`
}

type FastTCPTemplateCfg struct {
	FastTCPTemplateMax          int `json:"fast-tcp-template-max,omitempty"`
	FastTCPTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee,omitempty"`
}

type FastUDPTemplateCfg struct {
	FastUDPTemplateMax          int `json:"fast-udp-template-max,omitempty"`
	FastUDPTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee,omitempty"`
}

type FixTemplateCfg struct {
	FixTemplateMax          int `json:"fix-template-max,omitempty"`
	FixTemplateMinGuarantee int `json:"fix-template-min-guarantee,omitempty"`
}

type GslbDeviceCfg struct {
	GslbDeviceMax          int `json:"gslb-device-max,omitempty"`
	GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee,omitempty"`
}

type GslbGeoLocationCfg struct {
	GslbGeoLocationMax          int `json:"gslb-geo-location-max,omitempty"`
	GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee,omitempty"`
}

type GslbIPListCfg struct {
	GslbIPListMax          int `json:"gslb-ip-list-max,omitempty"`
	GslbIPListMinGuarantee int `json:"gslb-ip-list-min-guarantee,omitempty"`
}

type GslbPolicyCfg struct {
	GslbPolicyMax          int `json:"gslb-policy-max,omitempty"`
	GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee,omitempty"`
}

type GslbServiceCfg struct {
	GslbServiceMax          int `json:"gslb-service-max,omitempty"`
	GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee,omitempty"`
}

type GslbServiceIPCfg struct {
	GslbServiceIPMax          int `json:"gslb-service-ip-max,omitempty"`
	GslbServiceIPMinGuarantee int `json:"gslb-service-ip-min-guarantee,omitempty"`
}

type GslbServicePortCfg struct {
	GslbServicePortMax          int `json:"gslb-service-port-max,omitempty"`
	GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee,omitempty"`
}

type GslbSiteCfg struct {
	GslbSiteMax          int `json:"gslb-site-max,omitempty"`
	GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee,omitempty"`
}

type GslbSvcGroupCfg struct {
	GslbSvcGroupMax          int `json:"gslb-svc-group-max,omitempty"`
	GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee,omitempty"`
}

type GslbTemplateCfg struct {
	GslbTemplateMax          int `json:"gslb-template-max,omitempty"`
	GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee,omitempty"`
}

type GslbZoneCfg struct {
	GslbZoneMax          int `json:"gslb-zone-max,omitempty"`
	GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee,omitempty"`
}

type HTTPTemplateCfg struct {
	HTTPTemplateMax          int `json:"http-template-max,omitempty"`
	HTTPTemplateMinGuarantee int `json:"http-template-min-guarantee,omitempty"`
}

type HealthMonitorCfg struct {
	HealthMonitorMax          int `json:"health-monitor-max,omitempty"`
	HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee,omitempty"`
}

type LinkCostTemplateCfg struct {
	LinkCostTemplateMax          int `json:"link-cost-template-max,omitempty"`
	LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee,omitempty"`
}

type PersistCookieTemplateCfg struct {
	PersistCookieTemplateMax          int `json:"persist-cookie-template-max,omitempty"`
	PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee,omitempty"`
}

type PersistSrcipTemplateCfg struct {
	PersistSrcipTemplateMax          int `json:"persist-srcip-template-max,omitempty"`
	PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee,omitempty"`
}

type ProxyTemplateCfg struct {
	ProxyTemplateMax          int `json:"proxy-template-max,omitempty"`
	ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee,omitempty"`
}

type RealPortCfg struct {
	RealPortMax          int `json:"real-port-max,omitempty"`
	RealPortMinGuarantee int `json:"real-port-min-guarantee,omitempty"`
}

type RealServerCfg struct {
	RealServerMax          int `json:"real-server-max,omitempty"`
	RealServerMinGuarantee int `json:"real-server-min-guarantee,omitempty"`
}

type ServerSslTemplateCfg struct {
	ServerSslTemplateMax          int `json:"server-ssl-template-max,omitempty"`
	ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee,omitempty"`
}

type ServiceGroupCfg struct {
	ServiceGroupMax          int `json:"service-group-max,omitempty"`
	ServiceGroupMinGuarantee int `json:"service-group-min-guarantee,omitempty"`
}

type StreamTemplateCfg struct {
	StreamTemplateMax          int `json:"stream-template-max,omitempty"`
	StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee,omitempty"`
}

type VirtualPortCfg struct {
	VirtualPortMax          int `json:"virtual-port-max,omitempty"`
	VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee,omitempty"`
}

type VirtualServerCfg struct {
	VirtualServerMax          int `json:"virtual-server-max,omitempty"`
	VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee,omitempty"`
}

type Ipv4AclLineCfg struct {
	Ipv4AclLineMax          int `json:"ipv4-acl-line-max,omitempty"`
	Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee,omitempty"`
}

type Ipv6AclLineCfg struct {
	Ipv6AclLineMax          int `json:"ipv6-acl-line-max,omitempty"`
	Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee,omitempty"`
}

type ObjectGroupCfg struct {
	ObjectGroupMax          int `json:"object-group-max,omitempty"`
	ObjectGroupMinGuarantee int `json:"object-group-min-guarantee,omitempty"`
}

type ObjectGroupClauseCfg struct {
	ObjectGroupClauseMax          int `json:"object-group-clause-max,omitempty"`
	ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee,omitempty"`
}

type StaticArpCfg struct {
	StaticArpMax          int `json:"static-arp-max,omitempty"`
	StaticArpMinGuarantee int `json:"static-arp-min-guarantee,omitempty"`
}

type StaticIpv4RouteCfg struct {
	StaticIpv4RouteMax          int `json:"static-ipv4-route-max,omitempty"`
	StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee,omitempty"`
}

type StaticIpv6RouteCfg struct {
	StaticIpv6RouteMax          int `json:"static-ipv6-route-max,omitempty"`
	StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee,omitempty"`
}

type StaticMacCfg struct {
	StaticMacMax          int `json:"static-mac-max,omitempty"`
	StaticMacMinGuarantee int `json:"static-mac-min-guarantee,omitempty"`
}

type StaticNeighborCfg struct {
	StaticNeighborMax          int `json:"static-neighbor-max,omitempty"`
	StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee,omitempty"`
}

type BwLimitCfg struct {
	BwLimitMax              int `json:"bw-limit-max,omitempty"`
	BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable,omitempty"`
}

type ConcurrentSessionLimitCfg struct {
	ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max,omitempty"`
}

type FwcpsLimitCfg struct {
	FwcpsLimitMax int `json:"fwcps-limit-max,omitempty"`
}

type L4CpsLimitCfg struct {
	L4CpsLimitMax int `json:"l4cps-limit-max,omitempty"`
}

type L4SessionLimitCfg struct {
	L4SessionLimitMax          string `json:"l4-session-limit-max,omitempty"`
	L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee,omitempty"`
}

type L7CpsLimitCfg struct {
	L7CpsLimitMax int `json:"l7cps-limit-max,omitempty"`
}

type NatcpsLimitCfg struct {
	NatcpsLimitMax int `json:"natcps-limit-max,omitempty"`
}

type SslThroughputLimitCfg struct {
	SslThroughputLimitMax              int `json:"ssl-throughput-limit-max,omitempty"`
	SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable,omitempty"`
}

type SslcpsLimitCfg struct {
	SslcpsLimitMax int `json:"sslcps-limit-max,omitempty"`
}

func PostSystem(id string, inst System, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSystem")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/system", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m System
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSystem", data)

		}
	}

}

func GetSystem(id string, host string) (*System, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSystem")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/system", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m System
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSystem", data)
			return &m, nil
		}
	}

}
