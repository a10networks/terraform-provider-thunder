package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 5_2_1-P4_90
type System struct {
	Inst struct {
		AddCpuCore                  SystemAddCpuCore              `json:"add-cpu-core"`
		AddPort                     SystemAddPort                 `json:"add-port"`
		AllVlanLimit                SystemAllVlanLimit            `json:"all-vlan-limit"`
		AnomalyLog                  int                           `json:"anomaly-log"`
		AppPerformance              SystemAppPerformance          `json:"app-performance"`
		AppsGlobal                  SystemAppsGlobal              `json:"apps-global"`
		AsicDebugDump               SystemAsicDebugDump           `json:"asic-debug-dump"`
		AttackLog                   int                           `json:"attack-log"`
		Bandwidth                   SystemBandwidth               `json:"bandwidth"`
		Bfd                         SystemBfd                     `json:"bfd"`
		ClassListHitcountEnable     int                           `json:"class-list-hitcount-enable" dval:"1"`
		CmUpdateFileNameRef         SystemCmUpdateFileNameRef     `json:"cm-update-file-name-ref"`
		ControlCpu                  SystemControlCpu              `json:"control-cpu"`
		Core                        SystemCore                    `json:"core"`
		CosqShow                    SystemCosqShow                `json:"cosq-show"`
		CosqStats                   SystemCosqStats               `json:"cosq-stats"`
		CounterLibAccounting        SystemCounterLibAccounting    `json:"counter-lib-accounting"`
		CpuHyperThread              SystemCpuHyperThread          `json:"cpu-hyper-thread"`
		CpuList                     SystemCpuList                 `json:"cpu-list"`
		CpuLoadSharing              SystemCpuLoadSharing          `json:"cpu-load-sharing"`
		CpuMap                      SystemCpuMap                  `json:"cpu-map"`
		DataCpu                     SystemDataCpu                 `json:"data-cpu"`
		DdosAttack                  int                           `json:"ddos-attack"`
		DdosLog                     int                           `json:"ddos-log"`
		DeepHrxq                    SystemDeepHrxq                `json:"deep-hrxq"`
		DelPort                     SystemDelPort                 `json:"del-port"`
		DeleteCpuCore               SystemDeleteCpuCore           `json:"delete-cpu-core"`
		Dns                         SystemDns                     `json:"dns"`
		DnsCache                    SystemDnsCache                `json:"dns-cache"`
		DomainListHitcountEnable    int                           `json:"domain-list-hitcount-enable"`
		DomainListInfo              SystemDomainListInfo          `json:"domain-list-info"`
		DpdkStats                   SystemDpdkStats               `json:"dpdk-stats"`
		DynamicServiceDnsSocketPool int                           `json:"dynamic-service-dns-socket-pool"`
		Environment                 SystemEnvironment             `json:"environment"`
		FpgaCoreCrc                 SystemFpgaCoreCrc             `json:"fpga-core-crc"`
		FpgaDrop                    SystemFpgaDrop                `json:"fpga-drop"`
		Fw                          SystemFw                      `json:"fw"`
		GeoDbHitcountEnable         int                           `json:"geo-db-hitcount-enable"`
		GeoLocation                 SystemGeoLocation             `json:"geo-location"`
		Geoloc                      SystemGeoloc                  `json:"geoloc"`
		GeolocListList              []SystemGeolocListList        `json:"geoloc-list-list"`
		GeolocNameHelper            SystemGeolocNameHelper        `json:"geoloc-name-helper"`
		GeolocationFile             SystemGeolocationFile         `json:"geolocation-file"`
		Glid                        int                           `json:"glid"`
		GuestFile                   SystemGuestFile               `json:"guest-file"`
		GuiImageList                SystemGuiImageList            `json:"gui-image-list"`
		Hardware                    SystemHardware                `json:"hardware"`
		HardwareForward             SystemHardwareForward         `json:"hardware-forward"`
		HighMemoryL4Session         SystemHighMemoryL4Session     `json:"high-memory-l4-session"`
		HrxqStatus                  SystemHrxqStatus              `json:"hrxq-status"`
		Icmp                        SystemIcmp                    `json:"icmp"`
		IcmpRate                    SystemIcmpRate                `json:"icmp-rate"`
		Icmp6                       SystemIcmp6                   `json:"icmp6"`
		InuseCpuList                SystemInuseCpuList            `json:"inuse-cpu-list"`
		InusePortList               SystemInusePortList           `json:"inuse-port-list"`
		IoCpu                       SystemIoCpu                   `json:"io-cpu"`
		IpDnsCache                  SystemIpDnsCache              `json:"ip-dns-cache"`
		IpStats                     SystemIpStats                 `json:"ip-stats"`
		IpThreatList                SystemIpThreatList            `json:"ip-threat-list"`
		Ip6Stats                    SystemIp6Stats                `json:"ip6-stats"`
		Ipmi                        SystemIpmi                    `json:"ipmi"`
		IpmiService                 SystemIpmiService             `json:"ipmi-service"`
		Ipsec                       SystemIpsec                   `json:"ipsec"`
		Ipv6PrefixLength            int                           `json:"ipv6-prefix-length" dval:"128"`
		LinkCapability              SystemLinkCapability          `json:"link-capability"`
		LinkMonitor                 SystemLinkMonitor             `json:"link-monitor"`
		Lro                         SystemLro                     `json:"lro"`
		ManagementInterfaceMode     SystemManagementInterfaceMode `json:"management-interface-mode"`
		Memory                      SystemMemory                  `json:"memory"`
		MfaAuth                     SystemMfaAuth                 `json:"mfa-auth"`
		MfaCertStore                SystemMfaCertStore            `json:"mfa-cert-store"`
		MfaManagement               SystemMfaManagement           `json:"mfa-management"`
		MfaValidationType           SystemMfaValidationType       `json:"mfa-validation-type"`
		MgmtPort                    SystemMgmtPort                `json:"mgmt-port"`
		ModifyPort                  SystemModifyPort              `json:"modify-port"`
		ModuleCtrlCpu               string                        `json:"module-ctrl-cpu"`
		MonTemplate                 SystemMonTemplate             `json:"mon-template"`
		MultiQueueSupport           SystemMultiQueueSupport       `json:"multi-queue-support"`
		NdiscRa                     SystemNdiscRa                 `json:"ndisc-ra"`
		NsmA10lb                    SystemNsmA10lb                `json:"nsm-a10lb"`
		PasswordPolicy              SystemPasswordPolicy          `json:"password-policy"`
		PerVlanLimit                SystemPerVlanLimit            `json:"per-vlan-limit"`
		Platformtype                SystemPlatformtype            `json:"platformtype"`
		PortInfo                    SystemPortInfo                `json:"port-info"`
		PortList                    SystemPortList                `json:"port-list"`
		Ports                       SystemPorts                   `json:"ports"`
		ProbeNetworkDevices         SystemProbeNetworkDevices     `json:"probe-network-devices"`
		PromiscuousMode             int                           `json:"promiscuous-mode"`
		PsuInfo                     SystemPsuInfo                 `json:"psu-info"`
		QInQ                        SystemQInQ                    `json:"q-in-q"`
		QueuingBuffer               SystemQueuingBuffer           `json:"queuing-buffer"`
		Radius                      SystemRadius                  `json:"radius"`
		Reboot                      SystemReboot                  `json:"reboot"`
		ResourceAccounting          SystemResourceAccounting      `json:"resource-accounting"`
		ResourceUsage               SystemResourceUsage           `json:"resource-usage"`
		Session                     SystemSession                 `json:"session"`
		SessionReclaimLimit         SystemSessionReclaimLimit     `json:"session-reclaim-limit"`
		SetRxtxDescSize             SystemSetRxtxDescSize         `json:"set-rxtx-desc-size"`
		SetRxtxQueue                SystemSetRxtxQueue            `json:"set-rxtx-queue"`
		SetTcpSynPerSec             SystemSetTcpSynPerSec         `json:"set-tcp-syn-per-sec"`
		SharedPollMode              SystemSharedPollMode          `json:"shared-poll-mode"`
		ShellPrivileges             SystemShellPrivileges         `json:"shell-privileges"`
		ShmLogging                  SystemShmLogging              `json:"shm-logging"`
		Shutdown                    SystemShutdown                `json:"shutdown"`
		SockstressDisable           int                           `json:"sockstress-disable"`
		SpeProfile                  SystemSpeProfile              `json:"spe-profile"`
		SpeStatus                   SystemSpeStatus               `json:"spe-status"`
		SrcIpHashEnable             int                           `json:"src-ip-hash-enable"`
		SslReqQ                     SystemSslReqQ                 `json:"ssl-req-q"`
		SslScv                      SystemSslScv                  `json:"ssl-scv"`
		SslScvVerifyCrlSign         SystemSslScvVerifyCrlSign     `json:"ssl-scv-verify-crl-sign"`
		SslScvVerifyHost            SystemSslScvVerifyHost        `json:"ssl-scv-verify-host"`
		SslSetCompatibleCipher      SystemSslSetCompatibleCipher  `json:"ssl-set-compatible-cipher"`
		SslStatus                   SystemSslStatus               `json:"ssl-status"`
		SyslogTimeMsec              SystemSyslogTimeMsec          `json:"syslog-time-msec"`
		TableIntegrity              SystemTableIntegrity          `json:"table-integrity"`
		Tcp                         SystemTcp                     `json:"tcp"`
		TcpStats                    SystemTcpStats                `json:"tcp-stats"`
		TcpSynPerSec                SystemTcpSynPerSec            `json:"tcp-syn-per-sec"`
		TelemetryLog                SystemTelemetryLog            `json:"telemetry-log"`
		Template                    SystemTemplate                `json:"template"`
		TemplateBind                SystemTemplateBind            `json:"template-bind"`
		Throughput                  SystemThroughput              `json:"throughput"`
		TimeoutValue                SystemTimeoutValue            `json:"timeout-value"`
		Tls13Mgmt                   SystemTls13Mgmt               `json:"tls-1-3-mgmt"`
		Trunk                       SystemTrunk                   `json:"trunk"`
		TrunkHwHash                 SystemTrunkHwHash             `json:"trunk-hw-hash"`
		TrunkXauiHwHash             SystemTrunkXauiHwHash         `json:"trunk-xaui-hw-hash"`
		Tso                         SystemTso                     `json:"tso"`
		UpgradeStatus               SystemUpgradeStatus           `json:"upgrade-status"`
		Uuid                        string                        `json:"uuid"`
	} `json:"system"`
}

type SystemAddCpuCore struct {
	CoreIndex int `json:"core-index"`
}

type SystemAddPort struct {
	PortIndex int `json:"port-index"`
}

type SystemAllVlanLimit struct {
	Bcast        int    `json:"bcast" dval:"5000"`
	Ipmcast      int    `json:"ipmcast" dval:"5000"`
	Mcast        int    `json:"mcast" dval:"5000"`
	UnknownUcast int    `json:"unknown-ucast" dval:"5000"`
	Uuid         string `json:"uuid"`
}

type SystemAppPerformance struct {
	Uuid           string                               `json:"uuid"`
	SamplingEnable []SystemAppPerformanceSamplingEnable `json:"sampling-enable"`
}

type SystemAppPerformanceSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemAppsGlobal struct {
	LogSessionOnEstablished int    `json:"log-session-on-established"`
	MslTime                 int    `json:"msl-time" dval:"2"`
	Uuid                    string `json:"uuid"`
}

type SystemAsicDebugDump struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemBandwidth struct {
	Uuid           string                          `json:"uuid"`
	SamplingEnable []SystemBandwidthSamplingEnable `json:"sampling-enable"`
}

type SystemBandwidthSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemBfd struct {
	Uuid           string                    `json:"uuid"`
	SamplingEnable []SystemBfdSamplingEnable `json:"sampling-enable"`
}

type SystemBfdSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemCmUpdateFileNameRef struct {
	Source_name string `json:"source_name"`
	Dest_name   string `json:"dest_name"`
	Id          int    `json:"id"`
}

type SystemControlCpu struct {
	Uuid string `json:"uuid"`
}

type SystemCore struct {
	Uuid string `json:"uuid"`
}

type SystemCosqShow struct {
	Uuid string `json:"uuid"`
}

type SystemCosqStats struct {
	Uuid string `json:"uuid"`
}

type SystemCounterLibAccounting struct {
	Uuid string `json:"uuid"`
}

type SystemCpuHyperThread struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemCpuList struct {
	Uuid string `json:"uuid"`
}

type SystemCpuLoadSharing struct {
	Disable          int                                  `json:"disable"`
	PacketsPerSecond SystemCpuLoadSharingPacketsPerSecond `json:"packets-per-second"`
	CpuUsage         SystemCpuLoadSharingCpuUsage         `json:"cpu-usage"`
	Tcp              int                                  `json:"tcp"`
	Udp              int                                  `json:"udp"`
	Uuid             string                               `json:"uuid"`
}

type SystemCpuLoadSharingPacketsPerSecond struct {
	Min int `json:"min" dval:"100000"`
}

type SystemCpuLoadSharingCpuUsage struct {
	Low  int `json:"low" dval:"60"`
	High int `json:"high" dval:"75"`
}

type SystemCpuMap struct {
	Uuid string `json:"uuid"`
}

type SystemDataCpu struct {
	Uuid string `json:"uuid"`
}

type SystemDeepHrxq struct {
	Enable int `json:"enable"`
}

type SystemDelPort struct {
	PortIndex int `json:"port-index"`
}

type SystemDeleteCpuCore struct {
	CoreIndex int `json:"core-index"`
}

type SystemDns struct {
	Uuid                string                       `json:"uuid"`
	SamplingEnable      []SystemDnsSamplingEnable    `json:"sampling-enable"`
	RecursiveNameserver SystemDnsRecursiveNameserver `json:"recursive-nameserver"`
}

type SystemDnsSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemDnsRecursiveNameserver struct {
	FollowShared int                                      `json:"follow-shared"`
	ServerList   []SystemDnsRecursiveNameserverServerList `json:"server-list"`
	Uuid         string                                   `json:"uuid"`
}

type SystemDnsRecursiveNameserverServerList struct {
	Ipv4Addr string `json:"ipv4-addr"`
	V4Desc   string `json:"v4-desc"`
	Ipv6Addr string `json:"ipv6-addr"`
	V6Desc   string `json:"v6-desc"`
}

type SystemDnsCache struct {
	Uuid           string                         `json:"uuid"`
	SamplingEnable []SystemDnsCacheSamplingEnable `json:"sampling-enable"`
}

type SystemDnsCacheSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemDomainListInfo struct {
	Uuid string `json:"uuid"`
}

type SystemDpdkStats struct {
	Uuid           string                          `json:"uuid"`
	SamplingEnable []SystemDpdkStatsSamplingEnable `json:"sampling-enable"`
}

type SystemDpdkStatsSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemEnvironment struct {
	Uuid string `json:"uuid"`
}

type SystemFpgaCoreCrc struct {
	MonitorDisable int    `json:"monitor-disable"`
	RebootEnable   int    `json:"reboot-enable"`
	Uuid           string `json:"uuid"`
}

type SystemFpgaDrop struct {
	Uuid           string                         `json:"uuid"`
	SamplingEnable []SystemFpgaDropSamplingEnable `json:"sampling-enable"`
}

type SystemFpgaDropSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemFw struct {
	ApplicationMempool int    `json:"application-mempool"`
	ApplicationFlow    int    `json:"application-flow"`
	BasicDpiEnable     int    `json:"basic-dpi-enable"`
	Uuid               string `json:"uuid"`
}

type SystemGeoLocation struct {
	GeoLocationIana            int                                   `json:"geo-location-iana" dval:"1"`
	GeoLocationGeolite2City    int                                   `json:"geo-location-geolite2-city"`
	Geolite2CityIncludeIpv6    int                                   `json:"geolite2-city-include-ipv6"`
	GeoLocationGeolite2Country int                                   `json:"geo-location-geolite2-country"`
	Geolite2CountryIncludeIpv6 int                                   `json:"geolite2-country-include-ipv6"`
	GeolocLoadFileList         []SystemGeoLocationGeolocLoadFileList `json:"geoloc-load-file-list"`
	Uuid                       string                                `json:"uuid"`
	EntryList                  []SystemGeoLocationEntryList          `json:"entry-list"`
}

type SystemGeoLocationGeolocLoadFileList struct {
	GeoLocationLoadFilename string `json:"geo-location-load-filename"`
	TemplateName            string `json:"template-name"`
}

type SystemGeoLocationEntryList struct {
	GeoLocnObjName           string                                               `json:"geo-locn-obj-name"`
	GeoLocnMultipleAddresses []SystemGeoLocationEntryListGeoLocnMultipleAddresses `json:"geo-locn-multiple-addresses"`
	Uuid                     string                                               `json:"uuid"`
	UserTag                  string                                               `json:"user-tag"`
}

type SystemGeoLocationEntryListGeoLocnMultipleAddresses struct {
	FirstIpAddress   string `json:"first-ip-address"`
	GeolIpv4Mask     string `json:"geol-ipv4-mask"`
	IpAddr2          string `json:"ip-addr2"`
	FirstIpv6Address string `json:"first-ipv6-address"`
	GeolIpv6Mask     int    `json:"geol-ipv6-mask"`
	Ipv6Addr2        string `json:"ipv6-addr2"`
}

type SystemGeoloc struct {
	Uuid           string                       `json:"uuid"`
	SamplingEnable []SystemGeolocSamplingEnable `json:"sampling-enable"`
}

type SystemGeolocSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemGeolocListList struct {
	Name                  string                                      `json:"name"`
	Shared                int                                         `json:"shared"`
	IncludeGeolocNameList []SystemGeolocListListIncludeGeolocNameList `json:"include-geoloc-name-list"`
	ExcludeGeolocNameList []SystemGeolocListListExcludeGeolocNameList `json:"exclude-geoloc-name-list"`
	Uuid                  string                                      `json:"uuid"`
	UserTag               string                                      `json:"user-tag"`
	SamplingEnable        []SystemGeolocListListSamplingEnable        `json:"sampling-enable"`
}

type SystemGeolocListListIncludeGeolocNameList struct {
	IncludeGeolocNameVal string `json:"include-geoloc-name-val"`
}

type SystemGeolocListListExcludeGeolocNameList struct {
	ExcludeGeolocNameVal string `json:"exclude-geoloc-name-val"`
}

type SystemGeolocListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemGeolocNameHelper struct {
	Uuid           string                                 `json:"uuid"`
	SamplingEnable []SystemGeolocNameHelperSamplingEnable `json:"sampling-enable"`
}

type SystemGeolocNameHelperSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemGeolocationFile struct {
	Uuid      string                         `json:"uuid"`
	ErrorInfo SystemGeolocationFileErrorInfo `json:"error-info"`
}

type SystemGeolocationFileErrorInfo struct {
	Uuid string `json:"uuid"`
}

type SystemGuestFile struct {
	Uuid string `json:"uuid"`
}

type SystemGuiImageList struct {
	Uuid string `json:"uuid"`
}

type SystemHardware struct {
	Uuid string `json:"uuid"`
}

type SystemHardwareForward struct {
	Uuid           string                                `json:"uuid"`
	SamplingEnable []SystemHardwareForwardSamplingEnable `json:"sampling-enable"`
	Slb            SystemHardwareForwardSlb              `json:"slb"`
}

type SystemHardwareForwardSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemHardwareForwardSlb struct {
	Uuid           string                                   `json:"uuid"`
	SamplingEnable []SystemHardwareForwardSlbSamplingEnable `json:"sampling-enable"`
}

type SystemHardwareForwardSlbSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemHighMemoryL4Session struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemHrxqStatus struct {
	Uuid string `json:"uuid"`
}

type SystemIcmp struct {
	Uuid           string                     `json:"uuid"`
	SamplingEnable []SystemIcmpSamplingEnable `json:"sampling-enable"`
}

type SystemIcmpSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemIcmpRate struct {
	Uuid           string                         `json:"uuid"`
	SamplingEnable []SystemIcmpRateSamplingEnable `json:"sampling-enable"`
}

type SystemIcmpRateSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemIcmp6 struct {
	Uuid           string                      `json:"uuid"`
	SamplingEnable []SystemIcmp6SamplingEnable `json:"sampling-enable"`
}

type SystemIcmp6SamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemInuseCpuList struct {
	Uuid string `json:"uuid"`
}

type SystemInusePortList struct {
	Uuid string `json:"uuid"`
}

type SystemIoCpu struct {
	MaxCores int `json:"max-cores"`
}

type SystemIpDnsCache struct {
	Uuid string `json:"uuid"`
}

type SystemIpStats struct {
	Uuid           string                        `json:"uuid"`
	SamplingEnable []SystemIpStatsSamplingEnable `json:"sampling-enable"`
}

type SystemIpStatsSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemIpThreatList struct {
	Uuid                 string                                 `json:"uuid"`
	SamplingEnable       []SystemIpThreatListSamplingEnable     `json:"sampling-enable"`
	Ipv4SourceList       SystemIpThreatListIpv4SourceList       `json:"ipv4-source-list"`
	Ipv4DestList         SystemIpThreatListIpv4DestList         `json:"ipv4-dest-list"`
	Ipv6SourceList       SystemIpThreatListIpv6SourceList       `json:"ipv6-source-list"`
	Ipv6DestList         SystemIpThreatListIpv6DestList         `json:"ipv6-dest-list"`
	Ipv4InternetHostList SystemIpThreatListIpv4InternetHostList `json:"ipv4-internet-host-list"`
	Ipv6InternetHostList SystemIpThreatListIpv6InternetHostList `json:"ipv6-internet-host-list"`
}

type SystemIpThreatListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemIpThreatListIpv4SourceList struct {
	ClassListCfg []SystemIpThreatListIpv4SourceListClassListCfg `json:"class-list-cfg"`
	Uuid         string                                         `json:"uuid"`
}

type SystemIpThreatListIpv4SourceListClassListCfg struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv4DestList struct {
	ClassListCfg []SystemIpThreatListIpv4DestListClassListCfg `json:"class-list-cfg"`
	Uuid         string                                       `json:"uuid"`
}

type SystemIpThreatListIpv4DestListClassListCfg struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6SourceList struct {
	ClassListCfg []SystemIpThreatListIpv6SourceListClassListCfg `json:"class-list-cfg"`
	Uuid         string                                         `json:"uuid"`
}

type SystemIpThreatListIpv6SourceListClassListCfg struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6DestList struct {
	ClassListCfg []SystemIpThreatListIpv6DestListClassListCfg `json:"class-list-cfg"`
	Uuid         string                                       `json:"uuid"`
}

type SystemIpThreatListIpv6DestListClassListCfg struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv4InternetHostList struct {
	ClassListCfg []SystemIpThreatListIpv4InternetHostListClassListCfg `json:"class-list-cfg"`
	Uuid         string                                               `json:"uuid"`
}

type SystemIpThreatListIpv4InternetHostListClassListCfg struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6InternetHostList struct {
	ClassListCfg []SystemIpThreatListIpv6InternetHostListClassListCfg `json:"class-list-cfg"`
	Uuid         string                                               `json:"uuid"`
}

type SystemIpThreatListIpv6InternetHostListClassListCfg struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIp6Stats struct {
	Uuid           string                         `json:"uuid"`
	SamplingEnable []SystemIp6StatsSamplingEnable `json:"sampling-enable"`
}

type SystemIp6StatsSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemIpmi struct {
	Reset int             `json:"reset"`
	Ip    SystemIpmiIp    `json:"ip"`
	Ipsrc SystemIpmiIpsrc `json:"ipsrc"`
	User  SystemIpmiUser  `json:"user"`
	Tool  SystemIpmiTool  `json:"tool"`
}

type SystemIpmiIp struct {
	Ipv4Address    string `json:"ipv4-address"`
	Ipv4Netmask    string `json:"ipv4-netmask"`
	DefaultGateway string `json:"default-gateway"`
}

type SystemIpmiIpsrc struct {
	Dhcp   int `json:"dhcp"`
	Static int `json:"static"`
}

type SystemIpmiUser struct {
	Add           string `json:"add"`
	Password      string `json:"password"`
	Administrator int    `json:"administrator"`
	Callback      int    `json:"callback"`
	Operator      int    `json:"operator"`
	User          int    `json:"user"`
	Disable       string `json:"disable"`
	Privilege     string `json:"privilege"`
	Setname       string `json:"setname"`
	Newname       string `json:"newname"`
	Setpass       string `json:"setpass"`
	Newpass       string `json:"newpass"`
}

type SystemIpmiTool struct {
	Cmd string `json:"cmd"`
}

type SystemIpmiService struct {
	Disable int    `json:"disable"`
	Uuid    string `json:"uuid"`
}

type SystemIpsec struct {
	PacketRoundRobin int                    `json:"packet-round-robin"`
	CryptoCore       int                    `json:"crypto-core"`
	CryptoMem        int                    `json:"crypto-mem"`
	Uuid             string                 `json:"uuid"`
	FpgaDecrypt      SystemIpsecFpgaDecrypt `json:"fpga-decrypt"`
}

type SystemIpsecFpgaDecrypt struct {
	Action string `json:"action" dval:"disable"`
}

type SystemLinkCapability struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemLinkMonitor struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemLro struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemManagementInterfaceMode struct {
	Dedicated    int `json:"dedicated"`
	NonDedicated int `json:"non-dedicated"`
}

type SystemMemory struct {
	Uuid           string                       `json:"uuid"`
	SamplingEnable []SystemMemorySamplingEnable `json:"sampling-enable"`
}

type SystemMemorySamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemMfaAuth struct {
	Username     string `json:"username"`
	SecondFactor string `json:"second-factor"`
}

type SystemMfaCertStore struct {
	CertHost      string `json:"cert-host"`
	Protocol      string `json:"protocol"`
	CertStorePath string `json:"cert-store-path"`
	Username      string `json:"username"`
	PasswdString  string `json:"passwd-string"`
	Encrypted     string `json:"encrypted"`
	Uuid          string `json:"uuid"`
}

type SystemMfaManagement struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemMfaValidationType struct {
	CaCert string `json:"ca-cert"`
	Uuid   string `json:"uuid"`
}

type SystemMgmtPort struct {
	PortIndex  int    `json:"port-index"`
	MacAddress string `json:"mac-address"`
	PciAddress string `json:"pci-address"`
}

type SystemModifyPort struct {
	PortIndex  int `json:"port-index"`
	PortNumber int `json:"port-number"`
}

type SystemMonTemplate struct {
	MonitorList       []SystemMonTemplateMonitorList     `json:"monitor-list"`
	LinkBlockAsDown   SystemMonTemplateLinkBlockAsDown   `json:"link-block-as-down"`
	LinkDownOnRestart SystemMonTemplateLinkDownOnRestart `json:"link-down-on-restart"`
}

type SystemMonTemplateMonitorList struct {
	Id              int                                          `json:"id"`
	ClearCfg        []SystemMonTemplateMonitorListClearCfg       `json:"clear-cfg"`
	LinkDisableCfg  []SystemMonTemplateMonitorListLinkDisableCfg `json:"link-disable-cfg"`
	LinkEnableCfg   []SystemMonTemplateMonitorListLinkEnableCfg  `json:"link-enable-cfg"`
	MonitorRelation string                                       `json:"monitor-relation" dval:"monitor-and"`
	LinkUpCfg       []SystemMonTemplateMonitorListLinkUpCfg      `json:"link-up-cfg"`
	LinkDownCfg     []SystemMonTemplateMonitorListLinkDownCfg    `json:"link-down-cfg"`
	Uuid            string                                       `json:"uuid"`
	UserTag         string                                       `json:"user-tag"`
}

type SystemMonTemplateMonitorListClearCfg struct {
	Sessions         string `json:"sessions"`
	ClearAllSequence int    `json:"clear-all-sequence"`
	ClearSequence    int    `json:"clear-sequence"`
}

type SystemMonTemplateMonitorListLinkDisableCfg struct {
	Diseth      int `json:"diseth"`
	DisSequence int `json:"dis-sequence"`
}

type SystemMonTemplateMonitorListLinkEnableCfg struct {
	Enaeth      int `json:"enaeth"`
	EnaSequence int `json:"ena-sequence"`
}

type SystemMonTemplateMonitorListLinkUpCfg struct {
	LinkupEthernet1 int `json:"linkup-ethernet1"`
	LinkUpSequence1 int `json:"link-up-sequence1"`
	LinkupEthernet2 int `json:"linkup-ethernet2"`
	LinkUpSequence2 int `json:"link-up-sequence2"`
	LinkupEthernet3 int `json:"linkup-ethernet3"`
	LinkUpSequence3 int `json:"link-up-sequence3"`
}

type SystemMonTemplateMonitorListLinkDownCfg struct {
	LinkdownEthernet1 int `json:"linkdown-ethernet1"`
	LinkDownSequence1 int `json:"link-down-sequence1"`
	LinkdownEthernet2 int `json:"linkdown-ethernet2"`
	LinkDownSequence2 int `json:"link-down-sequence2"`
	LinkdownEthernet3 int `json:"linkdown-ethernet3"`
	LinkDownSequence3 int `json:"link-down-sequence3"`
}

type SystemMonTemplateLinkBlockAsDown struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemMonTemplateLinkDownOnRestart struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemMultiQueueSupport struct {
	Enable int `json:"enable"`
}

type SystemNdiscRa struct {
	Uuid           string                        `json:"uuid"`
	SamplingEnable []SystemNdiscRaSamplingEnable `json:"sampling-enable"`
}

type SystemNdiscRaSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemNsmA10lb struct {
	Kill int    `json:"kill"`
	Uuid string `json:"uuid"`
}

type SystemPasswordPolicy struct {
	Complexity string `json:"complexity"`
	Aging      string `json:"aging"`
	History    string `json:"history"`
	MinPswdLen int    `json:"min-pswd-len"`
	Uuid       string `json:"uuid"`
}

type SystemPerVlanLimit struct {
	Bcast        int    `json:"bcast" dval:"1000"`
	Ipmcast      int    `json:"ipmcast" dval:"1000"`
	Mcast        int    `json:"mcast" dval:"1000"`
	UnknownUcast int    `json:"unknown-ucast" dval:"1000"`
	Uuid         string `json:"uuid"`
}

type SystemPlatformtype struct {
	Uuid string `json:"uuid"`
}

type SystemPortInfo struct {
	Uuid string `json:"uuid"`
}

type SystemPortList struct {
	Uuid string `json:"uuid"`
}

type SystemPorts struct {
	LinkDetectionInterval int    `json:"link-detection-interval" dval:"1000"`
	Uuid                  string `json:"uuid"`
}

type SystemProbeNetworkDevices struct {
}

type SystemPsuInfo struct {
	Uuid string `json:"uuid"`
}

type SystemQInQ struct {
	InnerTpid      string `json:"inner-tpid"`
	OuterTpid      string `json:"outer-tpid"`
	EnableAllPorts int    `json:"enable-all-ports"`
	Uuid           string `json:"uuid"`
}

type SystemQueuingBuffer struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemRadius struct {
	Server SystemRadiusServer `json:"server"`
}

type SystemRadiusServer struct {
	ListenPort              int                                `json:"listen-port" dval:"1813"`
	Remote                  SystemRadiusServerRemote           `json:"remote"`
	Secret                  int                                `json:"secret"`
	SecretString            string                             `json:"secret-string"`
	Encrypted               string                             `json:"encrypted"`
	Vrid                    int                                `json:"vrid"`
	Attribute               []SystemRadiusServerAttribute      `json:"attribute"`
	DisableReply            int                                `json:"disable-reply"`
	AccountingStart         string                             `json:"accounting-start" dval:"append-entry"`
	AccountingStop          string                             `json:"accounting-stop" dval:"delete-entry"`
	AccountingInterimUpdate string                             `json:"accounting-interim-update" dval:"ignore"`
	AccountingOn            string                             `json:"accounting-on" dval:"ignore"`
	AttributeName           string                             `json:"attribute-name"`
	CustomAttributeName     string                             `json:"custom-attribute-name"`
	Uuid                    string                             `json:"uuid"`
	SamplingEnable          []SystemRadiusServerSamplingEnable `json:"sampling-enable"`
}

type SystemRadiusServerRemote struct {
	IpList []SystemRadiusServerRemoteIpList `json:"ip-list"`
}

type SystemRadiusServerRemoteIpList struct {
	IpListName         string `json:"ip-list-name"`
	IpListSecret       int    `json:"ip-list-secret"`
	IpListSecretString string `json:"ip-list-secret-string"`
	IpListEncrypted    string `json:"ip-list-encrypted"`
}

type SystemRadiusServerAttribute struct {
	AttributeValue string `json:"attribute-value"`
	PrefixLength   string `json:"prefix-length"`
	PrefixVendor   int    `json:"prefix-vendor"`
	PrefixNumber   int    `json:"prefix-number"`
	Name           string `json:"name"`
	Value          string `json:"value"`
	CustomVendor   int    `json:"custom-vendor"`
	CustomNumber   int    `json:"custom-number"`
	Vendor         int    `json:"vendor"`
	Number         int    `json:"number"`
}

type SystemRadiusServerSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemReboot struct {
	Uuid string `json:"uuid"`
}

type SystemResourceAccounting struct {
	Uuid         string                                 `json:"uuid"`
	TemplateList []SystemResourceAccountingTemplateList `json:"template-list"`
}

type SystemResourceAccountingTemplateList struct {
	Name             string                                               `json:"name"`
	Uuid             string                                               `json:"uuid"`
	UserTag          string                                               `json:"user-tag"`
	AppResources     SystemResourceAccountingTemplateListAppResources     `json:"app-resources"`
	NetworkResources SystemResourceAccountingTemplateListNetworkResources `json:"network-resources"`
	SystemResources  SystemResourceAccountingTemplateListSystemResources  `json:"system-resources"`
}

type SystemResourceAccountingTemplateListAppResources struct {
	GslbDeviceCfg            SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg            `json:"gslb-device-cfg"`
	GslbGeoLocationCfg       SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg       `json:"gslb-geo-location-cfg"`
	GslbIpListCfg            SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg            `json:"gslb-ip-list-cfg"`
	GslbPolicyCfg            SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg            `json:"gslb-policy-cfg"`
	GslbServiceCfg           SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg           `json:"gslb-service-cfg"`
	GslbServiceIpCfg         SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg         `json:"gslb-service-ip-cfg"`
	GslbServicePortCfg       SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg       `json:"gslb-service-port-cfg"`
	GslbSiteCfg              SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg              `json:"gslb-site-cfg"`
	GslbSvcGroupCfg          SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg          `json:"gslb-svc-group-cfg"`
	GslbTemplateCfg          SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg          `json:"gslb-template-cfg"`
	GslbZoneCfg              SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg              `json:"gslb-zone-cfg"`
	HealthMonitorCfg         SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg         `json:"health-monitor-cfg"`
	RealPortCfg              SystemResourceAccountingTemplateListAppResourcesRealPortCfg              `json:"real-port-cfg"`
	RealServerCfg            SystemResourceAccountingTemplateListAppResourcesRealServerCfg            `json:"real-server-cfg"`
	ServiceGroupCfg          SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg          `json:"service-group-cfg"`
	VirtualServerCfg         SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg         `json:"virtual-server-cfg"`
	VirtualPortCfg           SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg           `json:"virtual-port-cfg"`
	CacheTemplateCfg         SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg         `json:"cache-template-cfg"`
	ClientSslTemplateCfg     SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg     `json:"client-ssl-template-cfg"`
	ConnReuseTemplateCfg     SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg     `json:"conn-reuse-template-cfg"`
	FastTcpTemplateCfg       SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg       `json:"fast-tcp-template-cfg"`
	FastUdpTemplateCfg       SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg       `json:"fast-udp-template-cfg"`
	FixTemplateCfg           SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg           `json:"fix-template-cfg"`
	HttpTemplateCfg          SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg          `json:"http-template-cfg"`
	LinkCostTemplateCfg      SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg      `json:"link-cost-template-cfg"`
	PersistCookieTemplateCfg SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg `json:"persist-cookie-template-cfg"`
	PersistSrcipTemplateCfg  SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg  `json:"persist-srcip-template-cfg"`
	ServerSslTemplateCfg     SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg     `json:"server-ssl-template-cfg"`
	ProxyTemplateCfg         SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg         `json:"proxy-template-cfg"`
	StreamTemplateCfg        SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg        `json:"stream-template-cfg"`
	Threshold                int                                                                      `json:"threshold"`
	Uuid                     string                                                                   `json:"uuid"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg struct {
	GslbDeviceMax          int `json:"gslb-device-max"`
	GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg struct {
	GslbGeoLocationMax          int `json:"gslb-geo-location-max"`
	GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg struct {
	GslbIpListMax          int `json:"gslb-ip-list-max"`
	GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg struct {
	GslbPolicyMax          int `json:"gslb-policy-max"`
	GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg struct {
	GslbServiceMax          int `json:"gslb-service-max"`
	GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg struct {
	GslbServiceIpMax          int `json:"gslb-service-ip-max"`
	GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg struct {
	GslbServicePortMax          int `json:"gslb-service-port-max"`
	GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg struct {
	GslbSiteMax          int `json:"gslb-site-max"`
	GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg struct {
	GslbSvcGroupMax          int `json:"gslb-svc-group-max"`
	GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg struct {
	GslbTemplateMax          int `json:"gslb-template-max"`
	GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg struct {
	GslbZoneMax          int `json:"gslb-zone-max"`
	GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg struct {
	HealthMonitorMax          int `json:"health-monitor-max"`
	HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesRealPortCfg struct {
	RealPortMax          int `json:"real-port-max"`
	RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesRealServerCfg struct {
	RealServerMax          int `json:"real-server-max"`
	RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg struct {
	ServiceGroupMax          int `json:"service-group-max"`
	ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg struct {
	VirtualServerMax          int `json:"virtual-server-max"`
	VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg struct {
	VirtualPortMax          int `json:"virtual-port-max"`
	VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg struct {
	CacheTemplateMax          int `json:"cache-template-max"`
	CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg struct {
	ClientSslTemplateMax          int `json:"client-ssl-template-max"`
	ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg struct {
	ConnReuseTemplateMax          int `json:"conn-reuse-template-max"`
	ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg struct {
	FastTcpTemplateMax          int `json:"fast-tcp-template-max"`
	FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg struct {
	FastUdpTemplateMax          int `json:"fast-udp-template-max"`
	FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg struct {
	FixTemplateMax          int `json:"fix-template-max"`
	FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg struct {
	HttpTemplateMax          int `json:"http-template-max"`
	HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg struct {
	LinkCostTemplateMax          int `json:"link-cost-template-max"`
	LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg struct {
	PersistCookieTemplateMax          int `json:"persist-cookie-template-max"`
	PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg struct {
	PersistSrcipTemplateMax          int `json:"persist-srcip-template-max"`
	PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg struct {
	ServerSslTemplateMax          int `json:"server-ssl-template-max"`
	ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg struct {
	ProxyTemplateMax          int `json:"proxy-template-max"`
	ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg struct {
	StreamTemplateMax          int `json:"stream-template-max"`
	StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResources struct {
	StaticIpv4RouteCfg   SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg   `json:"static-ipv4-route-cfg"`
	StaticIpv6RouteCfg   SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg   `json:"static-ipv6-route-cfg"`
	Ipv4AclLineCfg       SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg       `json:"ipv4-acl-line-cfg"`
	Ipv6AclLineCfg       SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg       `json:"ipv6-acl-line-cfg"`
	StaticArpCfg         SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg         `json:"static-arp-cfg"`
	StaticNeighborCfg    SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg    `json:"static-neighbor-cfg"`
	StaticMacCfg         SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg         `json:"static-mac-cfg"`
	ObjectGroupCfg       SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg       `json:"object-group-cfg"`
	ObjectGroupClauseCfg SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg `json:"object-group-clause-cfg"`
	Threshold            int                                                                      `json:"threshold"`
	Uuid                 string                                                                   `json:"uuid"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg struct {
	StaticIpv4RouteMax          int `json:"static-ipv4-route-max"`
	StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg struct {
	StaticIpv6RouteMax          int `json:"static-ipv6-route-max"`
	StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg struct {
	Ipv4AclLineMax          int `json:"ipv4-acl-line-max"`
	Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg struct {
	Ipv6AclLineMax          int `json:"ipv6-acl-line-max"`
	Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg struct {
	StaticArpMax          int `json:"static-arp-max"`
	StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg struct {
	StaticNeighborMax          int `json:"static-neighbor-max"`
	StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg struct {
	StaticMacMax          int `json:"static-mac-max"`
	StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg struct {
	ObjectGroupMax          int `json:"object-group-max"`
	ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg struct {
	ObjectGroupClauseMax          int `json:"object-group-clause-max"`
	ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}

type SystemResourceAccountingTemplateListSystemResources struct {
	BwLimitCfg                SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg                `json:"bw-limit-cfg"`
	ConcurrentSessionLimitCfg SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg `json:"concurrent-session-limit-cfg"`
	L4SessionLimitCfg         SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg         `json:"l4-session-limit-cfg"`
	L4cpsLimitCfg             SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg             `json:"l4cps-limit-cfg"`
	L7cpsLimitCfg             SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg             `json:"l7cps-limit-cfg"`
	NatcpsLimitCfg            SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg            `json:"natcps-limit-cfg"`
	FwcpsLimitCfg             SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg             `json:"fwcps-limit-cfg"`
	SslThroughputLimitCfg     SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg     `json:"ssl-throughput-limit-cfg"`
	SslcpsLimitCfg            SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg            `json:"sslcps-limit-cfg"`
	Threshold                 int                                                                          `json:"threshold"`
	Uuid                      string                                                                       `json:"uuid"`
}

type SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg struct {
	BwLimitMax              int `json:"bw-limit-max"`
	BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg struct {
	ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg struct {
	L4SessionLimitMax          string `json:"l4-session-limit-max"`
	L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}

type SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg struct {
	L4cpsLimitMax int `json:"l4cps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg struct {
	L7cpsLimitMax int `json:"l7cps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg struct {
	NatcpsLimitMax int `json:"natcps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg struct {
	FwcpsLimitMax int `json:"fwcps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg struct {
	SslThroughputLimitMax              int `json:"ssl-throughput-limit-max"`
	SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg struct {
	SslcpsLimitMax int `json:"sslcps-limit-max"`
}

type SystemResourceUsage struct {
	SslContextMemory              int                           `json:"ssl-context-memory" dval:"2048"`
	SslDmaMemory                  int                           `json:"ssl-dma-memory" dval:"256"`
	NatPoolAddrCount              int                           `json:"nat-pool-addr-count"`
	L4SessionCount                int                           `json:"l4-session-count"`
	AuthPortalHtmlFileSize        int                           `json:"auth-portal-html-file-size" dval:"20"`
	AuthPortalImageFileSize       int                           `json:"auth-portal-image-file-size" dval:"6"`
	MaxAflexFileSize              int                           `json:"max-aflex-file-size" dval:"32"`
	AflexTableEntryCount          int                           `json:"aflex-table-entry-count"`
	ClassListIpv6AddrCount        int                           `json:"class-list-ipv6-addr-count"`
	ClassListAcEntryCount         int                           `json:"class-list-ac-entry-count"`
	ClassListEntryCount           int                           `json:"class-list-entry-count"`
	MaxAflexAuthzCollectionNumber int                           `json:"max-aflex-authz-collection-number" dval:"512"`
	RadiusTableSize               int                           `json:"radius-table-size"`
	AuthzPolicyNumber             int                           `json:"authz-policy-number"`
	IpsecSaNumber                 int                           `json:"ipsec-sa-number"`
	RamCacheMemoryLimit           int                           `json:"ram-cache-memory-limit"`
	WafTemplateCount              int                           `json:"waf-template-count"`
	AuthSessionCount              int                           `json:"auth-session-count"`
	Uuid                          string                        `json:"uuid"`
	Visibility                    SystemResourceUsageVisibility `json:"visibility"`
}

type SystemResourceUsageVisibility struct {
	MonitoredEntityCount int    `json:"monitored-entity-count"`
	Uuid                 string `json:"uuid"`
}

type SystemSession struct {
	Uuid           string                        `json:"uuid"`
	SamplingEnable []SystemSessionSamplingEnable `json:"sampling-enable"`
}

type SystemSessionSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemSessionReclaimLimit struct {
	NscanLimit int    `json:"nscan-limit" dval:"4096"`
	ScanFreq   int    `json:"scan-freq" dval:"5"`
	Uuid       string `json:"uuid"`
}

type SystemSetRxtxDescSize struct {
	PortIndex int `json:"port-index"`
	RxdSize   int `json:"rxd-size"`
	TxdSize   int `json:"txd-size"`
}

type SystemSetRxtxQueue struct {
	PortIndex int `json:"port-index"`
	RxqSize   int `json:"rxq-size"`
	TxqSize   int `json:"txq-size"`
}

type SystemSetTcpSynPerSec struct {
	TcpSynValue int    `json:"tcp-syn-value" dval:"70"`
	Uuid        string `json:"uuid"`
}

type SystemSharedPollMode struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemShellPrivileges struct {
	EnableShellPrivileges int    `json:"enable-shell-privileges"`
	Uuid                  string `json:"uuid"`
}

type SystemShmLogging struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemShutdown struct {
	Uuid string `json:"uuid"`
}

type SystemSpeProfile struct {
	Action string `json:"action" dval:"ipv4-ipv6"`
}

type SystemSpeStatus struct {
	Uuid string `json:"uuid"`
}

type SystemSslReqQ struct {
	Uuid           string                        `json:"uuid"`
	SamplingEnable []SystemSslReqQSamplingEnable `json:"sampling-enable"`
}

type SystemSslReqQSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemSslScv struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemSslScvVerifyCrlSign struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemSslScvVerifyHost struct {
	Disable int    `json:"disable"`
	Uuid    string `json:"uuid"`
}

type SystemSslSetCompatibleCipher struct {
	Disable int    `json:"disable"`
	Uuid    string `json:"uuid"`
}

type SystemSslStatus struct {
	Uuid string `json:"uuid"`
}

type SystemSyslogTimeMsec struct {
	EnableFlag int `json:"enable-flag"`
}

type SystemTableIntegrity struct {
	Table          string                               `json:"table" dval:"all"`
	AuditAction    string                               `json:"audit-action" dval:"enable"`
	AutoSyncAction string                               `json:"auto-sync-action" dval:"enable"`
	Uuid           string                               `json:"uuid"`
	SamplingEnable []SystemTableIntegritySamplingEnable `json:"sampling-enable"`
}

type SystemTableIntegritySamplingEnable struct {
	Counters1 string `json:"counters1"`
	Counters2 string `json:"counters2"`
	Counters3 string `json:"counters3"`
}

type SystemTcp struct {
	Uuid                      string                             `json:"uuid"`
	SamplingEnable            []SystemTcpSamplingEnable          `json:"sampling-enable"`
	RateLimitResetUnknownConn SystemTcpRateLimitResetUnknownConn `json:"rate-limit-reset-unknown-conn"`
}

type SystemTcpSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemTcpRateLimitResetUnknownConn struct {
	PktRateForResetUnknownConn int    `json:"pkt-rate-for-reset-unknown-conn"`
	LogForResetUnknownConn     int    `json:"log-for-reset-unknown-conn"`
	Uuid                       string `json:"uuid"`
}

type SystemTcpStats struct {
	Uuid           string                         `json:"uuid"`
	SamplingEnable []SystemTcpStatsSamplingEnable `json:"sampling-enable"`
}

type SystemTcpStatsSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemTcpSynPerSec struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLog struct {
	TopKSourceList   SystemTelemetryLogTopKSourceList   `json:"top-k-source-list"`
	TopKAppSvcList   SystemTelemetryLogTopKAppSvcList   `json:"top-k-app-svc-list"`
	DeviceStatus     SystemTelemetryLogDeviceStatus     `json:"device-status"`
	Environment      SystemTelemetryLogEnvironment      `json:"environment"`
	PartitionMetrics SystemTelemetryLogPartitionMetrics `json:"partition-metrics"`
}

type SystemTelemetryLogTopKSourceList struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogTopKAppSvcList struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogDeviceStatus struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogEnvironment struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogPartitionMetrics struct {
	Uuid string `json:"uuid"`
}

type SystemTemplate struct {
	TemplatePolicy string `json:"template-policy"`
	Uuid           string `json:"uuid"`
}

type SystemTemplateBind struct {
	MonitorList []SystemTemplateBindMonitorList `json:"monitor-list"`
}

type SystemTemplateBindMonitorList struct {
	TemplateMonitor int    `json:"template-monitor"`
	Uuid            string `json:"uuid"`
}

type SystemThroughput struct {
	Uuid           string                           `json:"uuid"`
	SamplingEnable []SystemThroughputSamplingEnable `json:"sampling-enable"`
}

type SystemThroughputSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type SystemTimeoutValue struct {
	Ftp   int    `json:"ftp" dval:"120"`
	Scp   int    `json:"scp" dval:"300"`
	Sftp  int    `json:"sftp"`
	Tftp  int    `json:"tftp"`
	Http  int    `json:"http"`
	Https int    `json:"https"`
	Uuid  string `json:"uuid"`
}

type SystemTls13Mgmt struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemTrunk struct {
	LoadBalance SystemTrunkLoadBalance `json:"load-balance"`
}

type SystemTrunkLoadBalance struct {
	UseL3 int    `json:"use-l3"`
	UseL4 int    `json:"use-l4"`
	Uuid  string `json:"uuid"`
}

type SystemTrunkHwHash struct {
	Mode int    `json:"mode" dval:"6"`
	Uuid string `json:"uuid"`
}

type SystemTrunkXauiHwHash struct {
	Mode int    `json:"mode" dval:"6"`
	Uuid string `json:"uuid"`
}

type SystemTso struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemUpgradeStatus struct {
	Uuid string `json:"uuid"`
}

func (p *System) GetId() string {
	return "1"
}

func (p *System) getPath() string {
	return "system"
}

func (p *System) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("System::Post")
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

func (p *System) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("System::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *System) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("System::Put")
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

func (p *System) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("System::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
