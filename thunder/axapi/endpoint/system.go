package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type System struct {
	Inst struct {
		AddCpuCore SystemAddCpuCore1755 `json:"add-cpu-core"`

		AddPort SystemAddPort1756 `json:"add-port"`

		AllVlanLimit SystemAllVlanLimit1757 `json:"all-vlan-limit"`

		AnomalyLog int `json:"anomaly-log"`

		AnomalyLogRateLimit int `json:"anomaly-log-rate-limit"`

		AppPerformance SystemAppPerformance1758 `json:"app-performance"`

		AppsGlobal SystemAppsGlobal1760 `json:"apps-global"`

		AsicDebugDump SystemAsicDebugDump1761 `json:"asic-debug-dump"`

		AsicMmuFailSafe SystemAsicMmuFailSafe1762 `json:"asic-mmu-fail-safe"`

		AttackLog int `json:"attack-log"`

		Bandwidth SystemBandwidth1763 `json:"bandwidth"`

		Bfd SystemBfd1765 `json:"bfd"`

		ClThreatCategory SystemClThreatCategory1767 `json:"cl-threat-category"`

		ClassListHitcountEnable int `json:"class-list-hitcount-enable"`

		CliMonitorInterval SystemCliMonitorInterval1768 `json:"cli-monitor-interval"`

		CmUpdateFileNameRef SystemCmUpdateFileNameRef1769 `json:"cm-update-file-name-ref"`

		ConfigMgmt SystemConfigMgmt1770 `json:"config-mgmt"`

		ControlCpu SystemControlCpu1774 `json:"control-cpu"`

		Core SystemCore1775 `json:"core"`

		CosqShow SystemCosqShow1776 `json:"cosq-show"`

		CosqStats SystemCosqStats1777 `json:"cosq-stats"`

		CounterLibAccounting SystemCounterLibAccounting1778 `json:"counter-lib-accounting"`

		CpuHyperThread SystemCpuHyperThread1779 `json:"cpu-hyper-thread"`

		CpuList SystemCpuList1780 `json:"cpu-list"`

		CpuLoadSharing SystemCpuLoadSharing1781 `json:"cpu-load-sharing"`

		CpuMap SystemCpuMap1784 `json:"cpu-map"`

		CpuPacketPrioSupport SystemCpuPacketPrioSupport1785 `json:"cpu-packet-prio-support"`

		DataCpu SystemDataCpu1786 `json:"data-cpu"`

		DdosAttack int `json:"ddos-attack"`

		DdosLog int `json:"ddos-log"`

		DefaultMtu int `json:"default-mtu"`

		DelPort SystemDelPort1787 `json:"del-port"`

		DeleteCpuCore SystemDeleteCpuCore1788 `json:"delete-cpu-core"`

		DisableSshAgentForwarding int `json:"disable-ssh-agent-forwarding"`

		Dns SystemDns1789 `json:"dns"`

		DnsCache SystemDnsCache1793 `json:"dns-cache"`

		DomainListHitcountEnable int `json:"domain-list-hitcount-enable"`

		DomainListInfo SystemDomainListInfo1795 `json:"domain-list-info"`

		DomainListSettings SystemDomainListSettings1796 `json:"domain-list-settings"`

		DpdkStats SystemDpdkStats1797 `json:"dpdk-stats"`

		DropLinuxClosedPortSyn string `json:"drop-linux-closed-port-syn" dval:"enable"`

		DynamicServiceDnsSocketPool int `json:"dynamic-service-dns-socket-pool"`

		EnableDiskEncryption SystemEnableDiskEncryption1799 `json:"enable-disk-encryption"`

		EnablePassword SystemEnablePassword1800 `json:"enable-password"`

		Environment SystemEnvironment1801 `json:"environment"`

		EvenPortHashEnable int `json:"even-port-hash-enable"`

		ExtOnlyLogging SystemExtOnlyLogging1802 `json:"ext-only-logging"`

		ForcedGroupSpeedList []SystemForcedGroupSpeedList `json:"forced-group-speed-list"`

		FpgaCoreCrc SystemFpgaCoreCrc1803 `json:"fpga-core-crc"`

		FpgaDrop SystemFpgaDrop1804 `json:"fpga-drop"`

		Fw SystemFw1806 `json:"fw"`

		GeoDbHitcountEnable int `json:"geo-db-hitcount-enable"`

		GeoLocation SystemGeoLocation1807 `json:"geo-location"`

		Geoloc SystemGeoloc1811 `json:"geoloc"`

		GeolocListList []SystemGeolocListList `json:"geoloc-list-list"`

		GeolocNameHelper SystemGeolocNameHelper1813 `json:"geoloc-name-helper"`

		GeolocationFile SystemGeolocationFile1815 `json:"geolocation-file"`

		Glid SystemGlid1817 `json:"glid"`

		GuestFile SystemGuestFile1818 `json:"guest-file"`

		GuiImageList SystemGuiImageList1819 `json:"gui-image-list"`

		Hardware SystemHardware1820 `json:"hardware"`

		HardwareAccelerate SystemHardwareAccelerate1821 `json:"hardware-accelerate"`

		HealthCheckList []SystemHealthCheckList `json:"health-check-list"`

		HighMemoryL4Session SystemHighMemoryL4Session1825 `json:"high-memory-l4-session"`

		HrxqStatus SystemHrxqStatus1826 `json:"hrxq-status"`

		HwBlockingEnable int `json:"hw-blocking-enable"`

		Icmp SystemIcmp1827 `json:"icmp"`

		IcmpRate SystemIcmpRate1829 `json:"icmp-rate"`

		Icmp6 SystemIcmp61831 `json:"icmp6"`

		InuseCpuList SystemInuseCpuList1833 `json:"inuse-cpu-list"`

		InusePortList SystemInusePortList1834 `json:"inuse-port-list"`

		IoCpu SystemIoCpu1835 `json:"io-cpu"`

		Ip SystemIp1836 `json:"ip"`

		IpDnsCache SystemIpDnsCache1837 `json:"ip-dns-cache"`

		IpStats SystemIpStats1838 `json:"ip-stats"`

		IpThreatList SystemIpThreatList1840 `json:"ip-threat-list"`

		Ip6Stats SystemIp6Stats1854 `json:"ip6-stats"`

		Ipmi SystemIpmi1856 `json:"ipmi"`

		IpmiService SystemIpmiService1861 `json:"ipmi-service"`

		Ipsec SystemIpsec1862 `json:"ipsec"`

		Ipv6 SystemIpv61864 `json:"ipv6"`

		Ipv6PrefixLength int `json:"ipv6-prefix-length" dval:"128"`

		JobOffload SystemJobOffload1865 `json:"job-offload"`

		LinkCapability SystemLinkCapability1867 `json:"link-capability"`

		LinkMonitor SystemLinkMonitor1868 `json:"link-monitor"`

		Lro SystemLro1869 `json:"lro"`

		ManagementInterfaceMode SystemManagementInterfaceMode1870 `json:"management-interface-mode"`

		Memory SystemMemory1871 `json:"memory"`

		MemoryBlockDebug SystemMemoryBlockDebug1873 `json:"memory-block-debug"`

		MfaAuth SystemMfaAuth1874 `json:"mfa-auth"`

		MfaCertStore SystemMfaCertStore1875 `json:"mfa-cert-store"`

		MfaManagement SystemMfaManagement1876 `json:"mfa-management"`

		MfaValidationType SystemMfaValidationType1877 `json:"mfa-validation-type"`

		MgmtPort SystemMgmtPort1878 `json:"mgmt-port"`

		ModifyPort SystemModifyPort1879 `json:"modify-port"`

		ModuleCtrlCpu string `json:"module-ctrl-cpu"`

		MonTemplate SystemMonTemplate1880 `json:"mon-template"`

		MultiQueueSupport SystemMultiQueueSupport1884 `json:"multi-queue-support"`

		NdiscRa SystemNdiscRa1885 `json:"ndisc-ra"`

		NetvscMonitor SystemNetvscMonitor1887 `json:"netvsc-monitor"`

		NsmA10lb SystemNsmA10lb1888 `json:"nsm-a10lb"`

		PasswordPolicy SystemPasswordPolicy1889 `json:"password-policy"`

		PathList []SystemPathList `json:"path-list"`

		Pbslb SystemPbslb1890 `json:"pbslb"`

		PerVlanLimit SystemPerVlanLimit1892 `json:"per-vlan-limit"`

		Platformtype SystemPlatformtype1893 `json:"platformtype"`

		PortCount SystemPortCount1894 `json:"port-count"`

		PortInfo SystemPortInfo1895 `json:"port-info"`

		PortList SystemPortList1896 `json:"port-list"`

		Ports SystemPorts1897 `json:"ports"`

		PowerOnSelfTest SystemPowerOnSelfTest1898 `json:"power-on-self-test"`

		ProbeNetworkDevices SystemProbeNetworkDevices1899 `json:"probe-network-devices"`

		PromiscuousMode int `json:"promiscuous-mode"`

		PsuInfo SystemPsuInfo1900 `json:"psu-info"`

		QInQ SystemQInQ1901 `json:"q-in-q"`

		QueuingBuffer SystemQueuingBuffer1902 `json:"queuing-buffer"`

		Radius SystemRadius1903 `json:"radius"`

		Reboot SystemReboot1912 `json:"reboot"`

		ResourceAccounting SystemResourceAccounting1913 `json:"resource-accounting"`

		ResourceUsage SystemResourceUsage1967 `json:"resource-usage"`

		RfcIpfixIeSpec string `json:"rfc-ipfix-ie-spec" dval:"disable"`

		Session SystemSession1969 `json:"session"`

		SessionReclaimLimit SystemSessionReclaimLimit1971 `json:"session-reclaim-limit"`

		SetRxtxDescSize SystemSetRxtxDescSize1972 `json:"set-rxtx-desc-size"`

		SetRxtxQueue SystemSetRxtxQueue1973 `json:"set-rxtx-queue"`

		SetTcpSynPerSec SystemSetTcpSynPerSec1974 `json:"set-tcp-syn-per-sec"`

		SharedPollMode SystemSharedPollMode1975 `json:"shared-poll-mode"`

		ShellPrivileges SystemShellPrivileges1976 `json:"shell-privileges"`

		ShmLogging SystemShmLogging1977 `json:"shm-logging"`

		Shutdown SystemShutdown1978 `json:"shutdown"`

		SoftwareTcam SystemSoftwareTcam1979 `json:"software-tcam"`

		SpeProfile SystemSpeProfile1980 `json:"spe-profile"`

		SpeStatus SystemSpeStatus1981 `json:"spe-status"`

		SrcIpHashEnable int `json:"src-ip-hash-enable"`

		SslHwMemory SystemSslHwMemory1982 `json:"ssl-hw-memory"`

		SslReqQ SystemSslReqQ1984 `json:"ssl-req-q"`

		SslScv SystemSslScv1986 `json:"ssl-scv"`

		SslScvVerifyCrlSign SystemSslScvVerifyCrlSign1987 `json:"ssl-scv-verify-crl-sign"`

		SslScvVerifyHost SystemSslScvVerifyHost1988 `json:"ssl-scv-verify-host"`

		SslSetCompatibleCipher SystemSslSetCompatibleCipher1989 `json:"ssl-set-compatible-cipher"`

		SslStatus SystemSslStatus1990 `json:"ssl-status"`

		SyslogTimeMsec SystemSyslogTimeMsec1991 `json:"syslog-time-msec"`

		SystemChassisPortSplitEnable int `json:"system-chassis-port-split-enable"`

		TableIntegrity SystemTableIntegrity1992 `json:"table-integrity"`

		Tcp SystemTcp1994 `json:"tcp"`

		TcpStats SystemTcpStats1997 `json:"tcp-stats"`

		TcpSynPerSec SystemTcpSynPerSec1999 `json:"tcp-syn-per-sec"`

		TelemetryLog SystemTelemetryLog2000 `json:"telemetry-log"`

		Template SystemTemplate2006 `json:"template"`

		TemplateBind SystemTemplateBind2007 `json:"template-bind"`

		Throughput SystemThroughput2008 `json:"throughput"`

		TimeoutValue SystemTimeoutValue2010 `json:"timeout-value"`

		Trunk SystemTrunk2011 `json:"trunk"`

		TrunkHwHash SystemTrunkHwHash2013 `json:"trunk-hw-hash"`

		TrunkXauiHwHash SystemTrunkXauiHwHash2014 `json:"trunk-xaui-hw-hash"`

		Tso SystemTso2015 `json:"tso"`

		Udp SystemUdp2016 `json:"udp"`

		UpgradeStatus SystemUpgradeStatus2017 `json:"upgrade-status"`

		Uuid string `json:"uuid"`

		VeMacScheme SystemVeMacScheme2018 `json:"ve-mac-scheme"`

		XauiDlbMode SystemXauiDlbMode2019 `json:"xaui-dlb-mode"`
	} `json:"system"`
}

type SystemAddCpuCore1755 struct {
	CoreIndex int `json:"core-index"`
}

type SystemAddPort1756 struct {
	PortIndex int `json:"port-index"`
}

type SystemAllVlanLimit1757 struct {
	Bcast        int    `json:"bcast" dval:"5000"`
	Ipmcast      int    `json:"ipmcast" dval:"5000"`
	Mcast        int    `json:"mcast" dval:"5000"`
	UnknownUcast int    `json:"unknown-ucast" dval:"5000"`
	Uuid         string `json:"uuid"`
}

type SystemAppPerformance1758 struct {
	Uuid           string                                   `json:"uuid"`
	SamplingEnable []SystemAppPerformanceSamplingEnable1759 `json:"sampling-enable"`
}

type SystemAppPerformanceSamplingEnable1759 struct {
	Counters1 string `json:"counters1"`
}

type SystemAppsGlobal1760 struct {
	LogSessionOnEstablished int    `json:"log-session-on-established"`
	MslTime                 int    `json:"msl-time" dval:"2"`
	TimerWheelWalkLimit     int    `json:"timer-wheel-walk-limit" dval:"100"`
	SessionsThreshold       int    `json:"sessions-threshold"`
	CpsThreshold            int    `json:"cps-threshold"`
	Uuid                    string `json:"uuid"`
}

type SystemAsicDebugDump1761 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemAsicMmuFailSafe1762 struct {
	RecoveryThreshold int    `json:"recovery-threshold" dval:"2"`
	MonitorInterval   int    `json:"monitor-interval" dval:"60"`
	MonitorDisable    int    `json:"monitor-disable"`
	RebootDisable     int    `json:"reboot-disable"`
	InjectError       int    `json:"inject-error"`
	TestPatternType   string `json:"test-pattern-type" dval:"lcb"`
	Uuid              string `json:"uuid"`
}

type SystemBandwidth1763 struct {
	WarningThreshold  int                                 `json:"warning-threshold" dval:"75"`
	CriticalThreshold int                                 `json:"critical-threshold" dval:"95"`
	Uuid              string                              `json:"uuid"`
	SamplingEnable    []SystemBandwidthSamplingEnable1764 `json:"sampling-enable"`
}

type SystemBandwidthSamplingEnable1764 struct {
	Counters1 string `json:"counters1"`
}

type SystemBfd1765 struct {
	Uuid           string                        `json:"uuid"`
	SamplingEnable []SystemBfdSamplingEnable1766 `json:"sampling-enable"`
}

type SystemBfdSamplingEnable1766 struct {
	Counters1 string `json:"counters1"`
}

type SystemClThreatCategory1767 struct {
	Uuid string `json:"uuid"`
}

type SystemCliMonitorInterval1768 struct {
	Interval int    `json:"interval"`
	Uuid     string `json:"uuid"`
}

type SystemCmUpdateFileNameRef1769 struct {
	Source_name string `json:"source_name"`
	Dest_name   string `json:"dest_name"`
	Id1         int    `json:"id1"`
}

type SystemConfigMgmt1770 struct {
	DeleteReferencedTaggedObjects string                              `json:"delete-referenced-tagged-objects" dval:"enable"`
	Uuid                          string                              `json:"uuid"`
	PuSyncDetection               SystemConfigMgmtPuSyncDetection1771 `json:"pu-sync-detection"`
	Mpm                           SystemConfigMgmtMpm1772             `json:"mpm"`
	Notification                  SystemConfigMgmtNotification1773    `json:"notification"`
}

type SystemConfigMgmtPuSyncDetection1771 struct {
	Interval int    `json:"interval" dval:"30"`
	Action   string `json:"action" dval:"disable"`
	Uuid     string `json:"uuid"`
}

type SystemConfigMgmtMpm1772 struct {
	MaxWorkers     int    `json:"max-workers" dval:"1"`
	MinIdleWorkers int    `json:"min-idle-workers" dval:"1"`
	StartWorkers   int    `json:"start-workers" dval:"1"`
	Uuid           string `json:"uuid"`
}

type SystemConfigMgmtNotification1773 struct {
	Period int    `json:"period" dval:"15"`
	Uuid   string `json:"uuid"`
}

type SystemControlCpu1774 struct {
	Uuid string `json:"uuid"`
}

type SystemCore1775 struct {
	Uuid string `json:"uuid"`
}

type SystemCosqShow1776 struct {
	Uuid string `json:"uuid"`
}

type SystemCosqStats1777 struct {
	Uuid string `json:"uuid"`
}

type SystemCounterLibAccounting1778 struct {
	Uuid string `json:"uuid"`
}

type SystemCpuHyperThread1779 struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemCpuList1780 struct {
	Uuid string `json:"uuid"`
}

type SystemCpuLoadSharing1781 struct {
	Disable          int                                      `json:"disable"`
	PacketsPerSecond SystemCpuLoadSharingPacketsPerSecond1782 `json:"packets-per-second"`
	CpuUsage         SystemCpuLoadSharingCpuUsage1783         `json:"cpu-usage"`
	AllowL7Sessions  int                                      `json:"allow-l7-sessions"`
	Tcp              int                                      `json:"tcp"`
	Udp              int                                      `json:"udp"`
	Others           int                                      `json:"others"`
	Uuid             string                                   `json:"uuid"`
}

type SystemCpuLoadSharingPacketsPerSecond1782 struct {
	Min int `json:"min" dval:"100000"`
}

type SystemCpuLoadSharingCpuUsage1783 struct {
	Low  int `json:"low" dval:"60"`
	High int `json:"high" dval:"75"`
}

type SystemCpuMap1784 struct {
	Uuid string `json:"uuid"`
}

type SystemCpuPacketPrioSupport1785 struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemDataCpu1786 struct {
	Uuid string `json:"uuid"`
}

type SystemDelPort1787 struct {
	PortIndex int `json:"port-index"`
}

type SystemDeleteCpuCore1788 struct {
	CoreIndex int `json:"core-index"`
}

type SystemDns1789 struct {
	Uuid                string                           `json:"uuid"`
	SamplingEnable      []SystemDnsSamplingEnable1790    `json:"sampling-enable"`
	RecursiveNameserver SystemDnsRecursiveNameserver1791 `json:"recursive-nameserver"`
}

type SystemDnsSamplingEnable1790 struct {
	Counters1 string `json:"counters1"`
}

type SystemDnsRecursiveNameserver1791 struct {
	FollowShared int                                          `json:"follow-shared"`
	ServerList   []SystemDnsRecursiveNameserverServerList1792 `json:"server-list"`
	Uuid         string                                       `json:"uuid"`
}

type SystemDnsRecursiveNameserverServerList1792 struct {
	Ipv4Addr string `json:"ipv4-addr"`
	V4Desc   string `json:"v4-desc"`
	Ipv6Addr string `json:"ipv6-addr"`
	V6Desc   string `json:"v6-desc"`
}

type SystemDnsCache1793 struct {
	Uuid           string                             `json:"uuid"`
	SamplingEnable []SystemDnsCacheSamplingEnable1794 `json:"sampling-enable"`
}

type SystemDnsCacheSamplingEnable1794 struct {
	Counters1 string `json:"counters1"`
}

type SystemDomainListInfo1795 struct {
	Uuid string `json:"uuid"`
}

type SystemDomainListSettings1796 struct {
	PollingInterval    string `json:"polling-interval" dval:"10-second"`
	ConcurrentTask     int    `json:"concurrent-task" dval:"6"`
	DomainListPerGroup string `json:"domain-list-per-group" dval:"16"`
	Uuid               string `json:"uuid"`
}

type SystemDpdkStats1797 struct {
	Uuid           string                              `json:"uuid"`
	SamplingEnable []SystemDpdkStatsSamplingEnable1798 `json:"sampling-enable"`
}

type SystemDpdkStatsSamplingEnable1798 struct {
	Counters1 string `json:"counters1"`
}

type SystemEnableDiskEncryption1799 struct {
	Cipher           string `json:"cipher" dval:"aes"`
	Passphrase       string `json:"passphrase"`
	PassphraseBase64 string `json:"passphrase-base64"`
}

type SystemEnablePassword1800 struct {
	FollowPasswordPolicy int    `json:"follow-password-policy"`
	Uuid                 string `json:"uuid"`
}

type SystemEnvironment1801 struct {
	Uuid string `json:"uuid"`
}

type SystemExtOnlyLogging1802 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemForcedGroupSpeedList struct {
	Eth01_to_04 int    `json:"eth01_to_04"`
	Eth05_to_08 int    `json:"eth05_to_08"`
	Eth09_to_12 int    `json:"eth09_to_12"`
	Eth13_to_16 int    `json:"eth13_to_16"`
	Eth17_to_20 int    `json:"eth17_to_20"`
	Eth21_to_24 int    `json:"eth21_to_24"`
	Speed       string `json:"speed" dval:"10G"`
	Uuid        string `json:"uuid"`
}

type SystemFpgaCoreCrc1803 struct {
	MonitorDisable int    `json:"monitor-disable"`
	RebootEnable   int    `json:"reboot-enable"`
	Uuid           string `json:"uuid"`
}

type SystemFpgaDrop1804 struct {
	Uuid           string                             `json:"uuid"`
	SamplingEnable []SystemFpgaDropSamplingEnable1805 `json:"sampling-enable"`
}

type SystemFpgaDropSamplingEnable1805 struct {
	Counters1 string `json:"counters1"`
}

type SystemFw1806 struct {
	ApplicationMempool int    `json:"application-mempool"`
	ApplicationFlow    int    `json:"application-flow"`
	BasicDpiEnable     int    `json:"basic-dpi-enable"`
	Uuid               string `json:"uuid"`
}

type SystemGeoLocation1807 struct {
	GeoLocationIana            int                                       `json:"geo-location-iana" dval:"1"`
	GeoLocationIanaSystem      int                                       `json:"geo-location-iana-system"`
	GeoLocationGeolite2Asn     int                                       `json:"geo-location-geolite2-asn"`
	Geolite2AsnIncludeIpv6     int                                       `json:"geolite2-asn-include-ipv6"`
	GeoLocationGeolite2City    int                                       `json:"geo-location-geolite2-city"`
	Geolite2CityIncludeIpv6    int                                       `json:"geolite2-city-include-ipv6"`
	GeoLocationGeolite2Country int                                       `json:"geo-location-geolite2-country"`
	Geolite2CountryIncludeIpv6 int                                       `json:"geolite2-country-include-ipv6"`
	GeolocLoadFileList         []SystemGeoLocationGeolocLoadFileList1808 `json:"geoloc-load-file-list"`
	Uuid                       string                                    `json:"uuid"`
	EntryList                  []SystemGeoLocationEntryList1809          `json:"entry-list"`
}

type SystemGeoLocationGeolocLoadFileList1808 struct {
	GeoLocationLoadFilename        string `json:"geo-location-load-filename"`
	GeoLocationLoadFileIncludeIpv6 int    `json:"geo-location-load-file-include-ipv6"`
	TemplateName                   string `json:"template-name"`
	GeoLocationLoadTempIncludeIpv6 int    `json:"geo-location-load-temp-include-ipv6"`
}

type SystemGeoLocationEntryList1809 struct {
	GeoLocnObjName           string                                                   `json:"geo-locn-obj-name"`
	GeoLocnMultipleAddresses []SystemGeoLocationEntryListGeoLocnMultipleAddresses1810 `json:"geo-locn-multiple-addresses"`
	Uuid                     string                                                   `json:"uuid"`
	UserTag                  string                                                   `json:"user-tag"`
}

type SystemGeoLocationEntryListGeoLocnMultipleAddresses1810 struct {
	FirstIpAddress   string `json:"first-ip-address"`
	GeolIpv4Mask     string `json:"geol-ipv4-mask"`
	IpAddr2          string `json:"ip-addr2"`
	FirstIpv6Address string `json:"first-ipv6-address"`
	GeolIpv6Mask     int    `json:"geol-ipv6-mask"`
	Ipv6Addr2        string `json:"ipv6-addr2"`
}

type SystemGeoloc1811 struct {
	Uuid           string                           `json:"uuid"`
	SamplingEnable []SystemGeolocSamplingEnable1812 `json:"sampling-enable"`
}

type SystemGeolocSamplingEnable1812 struct {
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

type SystemGeolocNameHelper1813 struct {
	Uuid           string                                     `json:"uuid"`
	SamplingEnable []SystemGeolocNameHelperSamplingEnable1814 `json:"sampling-enable"`
}

type SystemGeolocNameHelperSamplingEnable1814 struct {
	Counters1 string `json:"counters1"`
}

type SystemGeolocationFile1815 struct {
	Uuid      string                             `json:"uuid"`
	ErrorInfo SystemGeolocationFileErrorInfo1816 `json:"error-info"`
}

type SystemGeolocationFileErrorInfo1816 struct {
	Uuid string `json:"uuid"`
}

type SystemGlid1817 struct {
	GlidId    string `json:"glid-id"`
	NonShared int    `json:"non-shared"`
	Uuid      string `json:"uuid"`
}

type SystemGuestFile1818 struct {
	Uuid string `json:"uuid"`
}

type SystemGuiImageList1819 struct {
	Uuid string `json:"uuid"`
}

type SystemHardware1820 struct {
	Uuid string `json:"uuid"`
}

type SystemHardwareAccelerate1821 struct {
	SessionForwarding int                                          `json:"session-forwarding"`
	Uuid              string                                       `json:"uuid"`
	SamplingEnable    []SystemHardwareAccelerateSamplingEnable1822 `json:"sampling-enable"`
	Slb               SystemHardwareAccelerateSlb1823              `json:"slb"`
}

type SystemHardwareAccelerateSamplingEnable1822 struct {
	Counters1 string `json:"counters1"`
}

type SystemHardwareAccelerateSlb1823 struct {
	Uuid           string                                          `json:"uuid"`
	SamplingEnable []SystemHardwareAccelerateSlbSamplingEnable1824 `json:"sampling-enable"`
}

type SystemHardwareAccelerateSlbSamplingEnable1824 struct {
	Counters1 string `json:"counters1"`
}

type SystemHealthCheckList struct {
	L2hmHcName      string `json:"l2hm-hc-name"`
	MethodL2bfd     int    `json:"method-l2bfd"`
	L2bfdTxInterval int    `json:"l2bfd-tx-interval"`
	L2bfdRxInterval int    `json:"l2bfd-rx-interval"`
	L2bfdMultiplier int    `json:"l2bfd-multiplier"`
	Uuid            string `json:"uuid"`
	UserTag         string `json:"user-tag"`
}

type SystemHighMemoryL4Session1825 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemHrxqStatus1826 struct {
	Uuid string `json:"uuid"`
}

type SystemIcmp1827 struct {
	Uuid           string                         `json:"uuid"`
	SamplingEnable []SystemIcmpSamplingEnable1828 `json:"sampling-enable"`
}

type SystemIcmpSamplingEnable1828 struct {
	Counters1 string `json:"counters1"`
}

type SystemIcmpRate1829 struct {
	Uuid           string                             `json:"uuid"`
	SamplingEnable []SystemIcmpRateSamplingEnable1830 `json:"sampling-enable"`
}

type SystemIcmpRateSamplingEnable1830 struct {
	Counters1 string `json:"counters1"`
}

type SystemIcmp61831 struct {
	Uuid           string                          `json:"uuid"`
	SamplingEnable []SystemIcmp6SamplingEnable1832 `json:"sampling-enable"`
}

type SystemIcmp6SamplingEnable1832 struct {
	Counters1 string `json:"counters1"`
}

type SystemInuseCpuList1833 struct {
	Uuid string `json:"uuid"`
}

type SystemInusePortList1834 struct {
	Uuid string `json:"uuid"`
}

type SystemIoCpu1835 struct {
	MaxCores int `json:"max-cores"`
}

type SystemIp1836 struct {
	IcmpUnreachableDisable   int    `json:"icmp-unreachable-disable"`
	IcmpRedirectDisable      int    `json:"icmp-redirect-disable"`
	RpfCheckEnable           int    `json:"rpf-check-enable"`
	SourceRoutePktDropEnable int    `json:"source-route-pkt-drop-enable"`
	Uuid                     string `json:"uuid"`
}

type SystemIpDnsCache1837 struct {
	Uuid string `json:"uuid"`
}

type SystemIpStats1838 struct {
	Uuid           string                            `json:"uuid"`
	SamplingEnable []SystemIpStatsSamplingEnable1839 `json:"sampling-enable"`
}

type SystemIpStatsSamplingEnable1839 struct {
	Counters1 string `json:"counters1"`
}

type SystemIpThreatList1840 struct {
	Uuid                 string                                     `json:"uuid"`
	SamplingEnable       []SystemIpThreatListSamplingEnable1841     `json:"sampling-enable"`
	Ipv4SourceList       SystemIpThreatListIpv4SourceList1842       `json:"ipv4-source-list"`
	Ipv4DestList         SystemIpThreatListIpv4DestList1844         `json:"ipv4-dest-list"`
	Ipv6SourceList       SystemIpThreatListIpv6SourceList1846       `json:"ipv6-source-list"`
	Ipv6DestList         SystemIpThreatListIpv6DestList1848         `json:"ipv6-dest-list"`
	Ipv4InternetHostList SystemIpThreatListIpv4InternetHostList1850 `json:"ipv4-internet-host-list"`
	Ipv6InternetHostList SystemIpThreatListIpv6InternetHostList1852 `json:"ipv6-internet-host-list"`
}

type SystemIpThreatListSamplingEnable1841 struct {
	Counters1 string `json:"counters1"`
}

type SystemIpThreatListIpv4SourceList1842 struct {
	ClassListCfg []SystemIpThreatListIpv4SourceListClassListCfg1843 `json:"class-list-cfg"`
	Uuid         string                                             `json:"uuid"`
}

type SystemIpThreatListIpv4SourceListClassListCfg1843 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv4DestList1844 struct {
	ClassListCfg []SystemIpThreatListIpv4DestListClassListCfg1845 `json:"class-list-cfg"`
	Uuid         string                                           `json:"uuid"`
}

type SystemIpThreatListIpv4DestListClassListCfg1845 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6SourceList1846 struct {
	ClassListCfg []SystemIpThreatListIpv6SourceListClassListCfg1847 `json:"class-list-cfg"`
	Uuid         string                                             `json:"uuid"`
}

type SystemIpThreatListIpv6SourceListClassListCfg1847 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6DestList1848 struct {
	ClassListCfg []SystemIpThreatListIpv6DestListClassListCfg1849 `json:"class-list-cfg"`
	Uuid         string                                           `json:"uuid"`
}

type SystemIpThreatListIpv6DestListClassListCfg1849 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv4InternetHostList1850 struct {
	WhiteList    string                                                   `json:"white-list"`
	ClassListCfg []SystemIpThreatListIpv4InternetHostListClassListCfg1851 `json:"class-list-cfg"`
	Uuid         string                                                   `json:"uuid"`
}

type SystemIpThreatListIpv4InternetHostListClassListCfg1851 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6InternetHostList1852 struct {
	WhiteList    string                                                   `json:"white-list"`
	ClassListCfg []SystemIpThreatListIpv6InternetHostListClassListCfg1853 `json:"class-list-cfg"`
	Uuid         string                                                   `json:"uuid"`
}

type SystemIpThreatListIpv6InternetHostListClassListCfg1853 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIp6Stats1854 struct {
	Uuid           string                             `json:"uuid"`
	SamplingEnable []SystemIp6StatsSamplingEnable1855 `json:"sampling-enable"`
}

type SystemIp6StatsSamplingEnable1855 struct {
	Counters1 string `json:"counters1"`
}

type SystemIpmi1856 struct {
	Reset int                 `json:"reset"`
	Ip    SystemIpmiIp1857    `json:"ip"`
	Ipsrc SystemIpmiIpsrc1858 `json:"ipsrc"`
	User  SystemIpmiUser1859  `json:"user"`
	Tool  SystemIpmiTool1860  `json:"tool"`
}

type SystemIpmiIp1857 struct {
	Ipv4Address    string `json:"ipv4-address"`
	Ipv4Netmask    string `json:"ipv4-netmask"`
	DefaultGateway string `json:"default-gateway"`
}

type SystemIpmiIpsrc1858 struct {
	Dhcp   int `json:"dhcp"`
	Static int `json:"static"`
}

type SystemIpmiUser1859 struct {
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

type SystemIpmiTool1860 struct {
	Cmd string `json:"cmd"`
}

type SystemIpmiService1861 struct {
	Disable int    `json:"disable"`
	Uuid    string `json:"uuid"`
}

type SystemIpsec1862 struct {
	PacketRoundRobin int                        `json:"packet-round-robin"`
	CryptoCore       int                        `json:"crypto-core"`
	CryptoMem        int                        `json:"crypto-mem"`
	Qat              int                        `json:"QAT"`
	Uuid             string                     `json:"uuid"`
	FpgaDecrypt      SystemIpsecFpgaDecrypt1863 `json:"fpga-decrypt"`
}

type SystemIpsecFpgaDecrypt1863 struct {
	Action string `json:"action" dval:"disable"`
}

type SystemIpv61864 struct {
	Icmpv6UnreachableDisable int    `json:"icmpv6-unreachable-disable"`
	Icmpv6RedirectDisable    int    `json:"icmpv6-redirect-disable"`
	RpfCheckEnable           int    `json:"rpf-check-enable"`
	SourceRoutePktDropEnable int    `json:"source-route-pkt-drop-enable"`
	Uuid                     string `json:"uuid"`
}

type SystemJobOffload1865 struct {
	Uuid           string                               `json:"uuid"`
	SamplingEnable []SystemJobOffloadSamplingEnable1866 `json:"sampling-enable"`
}

type SystemJobOffloadSamplingEnable1866 struct {
	Counters1 string `json:"counters1"`
}

type SystemLinkCapability1867 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemLinkMonitor1868 struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemLro1869 struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemManagementInterfaceMode1870 struct {
	Dedicated    int `json:"dedicated"`
	NonDedicated int `json:"non-dedicated"`
}

type SystemMemory1871 struct {
	Uuid           string                           `json:"uuid"`
	SamplingEnable []SystemMemorySamplingEnable1872 `json:"sampling-enable"`
}

type SystemMemorySamplingEnable1872 struct {
	Counters1 string `json:"counters1"`
}

type SystemMemoryBlockDebug1873 struct {
	AssertBlock  int    `json:"assert-block" dval:"65536"`
	PktdumpBlock int    `json:"pktdump-block"`
	FirstBlk     int    `json:"first-blk" dval:"8192"`
	SecondBlk    int    `json:"second-blk" dval:"16384"`
	ThirdBlk     int    `json:"third-blk" dval:"32768"`
	FourthBlk    int    `json:"fourth-blk" dval:"65536"`
	Uuid         string `json:"uuid"`
}

type SystemMfaAuth1874 struct {
	Username     string `json:"username"`
	SecondFactor string `json:"second-factor"`
}

type SystemMfaCertStore1875 struct {
	CertHost      string `json:"cert-host"`
	Protocol      string `json:"protocol"`
	CertStorePath string `json:"cert-store-path"`
	Username      string `json:"username"`
	PasswdString  string `json:"passwd-string"`
	Encrypted     string `json:"encrypted"`
	Uuid          string `json:"uuid"`
}

type SystemMfaManagement1876 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemMfaValidationType1877 struct {
	CaCert string `json:"ca-cert"`
	Uuid   string `json:"uuid"`
}

type SystemMgmtPort1878 struct {
	PortIndex  int    `json:"port-index"`
	MacAddress string `json:"mac-address"`
	PciAddress string `json:"pci-address"`
}

type SystemModifyPort1879 struct {
	PortIndex  int `json:"port-index"`
	PortNumber int `json:"port-number"`
}

type SystemMonTemplate1880 struct {
	MonitorList       []SystemMonTemplateMonitorList         `json:"monitor-list"`
	LinkBlockAsDown   SystemMonTemplateLinkBlockAsDown1881   `json:"link-block-as-down"`
	LinkDownOnRestart SystemMonTemplateLinkDownOnRestart1882 `json:"link-down-on-restart"`
	MonitoringMode    SystemMonTemplateMonitoringMode1883    `json:"monitoring-mode"`
}

type SystemMonTemplateMonitorList struct {
	Id1             int                                          `json:"id1"`
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
	Sessions             string `json:"sessions"`
	ClearAllSequence     int    `json:"clear-all-sequence"`
	ClearAllPartition    string `json:"clear-all-partition"`
	ClearAllPartitionAll int    `json:"clear-all-partition-all"`
	ClearSequence        int    `json:"clear-sequence"`
	ClearPartition       string `json:"clear-partition"`
	ClearPartitionAll    int    `json:"clear-partition-all"`
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

type SystemMonTemplateLinkBlockAsDown1881 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemMonTemplateLinkDownOnRestart1882 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemMonTemplateMonitoringMode1883 struct {
	Mmode string `json:"mmode" dval:"and"`
	Uuid  string `json:"uuid"`
}

type SystemMultiQueueSupport1884 struct {
	Enable int `json:"enable"`
}

type SystemNdiscRa1885 struct {
	Uuid           string                            `json:"uuid"`
	SamplingEnable []SystemNdiscRaSamplingEnable1886 `json:"sampling-enable"`
}

type SystemNdiscRaSamplingEnable1886 struct {
	Counters1 string `json:"counters1"`
}

type SystemNetvscMonitor1887 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemNsmA10lb1888 struct {
	Kill int    `json:"kill"`
	Uuid string `json:"uuid"`
}

type SystemPasswordPolicy1889 struct {
	Complexity                 string `json:"complexity"`
	Aging                      string `json:"aging"`
	History                    string `json:"history"`
	MinPswdLen                 int    `json:"min-pswd-len"`
	UsernameCheck              string `json:"username-check" dval:"disable"`
	RepeatCharacterCheck       string `json:"repeat-character-check" dval:"disable"`
	ForbidConsecutiveCharacter string `json:"forbid-consecutive-character" dval:"0"`
	Uuid                       string `json:"uuid"`
}

type SystemPathList struct {
	L2hmPathName     string `json:"l2hm-path-name"`
	L2hmVlan         int    `json:"l2hm-vlan"`
	L2hmSetupTestApi int    `json:"l2hm-setup-test-api"`
	IfpairEthStart   int    `json:"ifpair-eth-start"`
	IfpairEthEnd     int    `json:"ifpair-eth-end"`
	IfpairTrunkStart int    `json:"ifpair-trunk-start"`
	IfpairTrunkEnd   int    `json:"ifpair-trunk-end"`
	L2hmAttach       string `json:"l2hm-attach"`
	Uuid             string `json:"uuid"`
	UserTag          string `json:"user-tag"`
}

type SystemPbslb1890 struct {
	SockstressDisable int                             `json:"sockstress-disable"`
	Uuid              string                          `json:"uuid"`
	SamplingEnable    []SystemPbslbSamplingEnable1891 `json:"sampling-enable"`
}

type SystemPbslbSamplingEnable1891 struct {
	Counters1 string `json:"counters1"`
}

type SystemPerVlanLimit1892 struct {
	Bcast        int    `json:"bcast" dval:"1000"`
	Ipmcast      int    `json:"ipmcast" dval:"1000"`
	Mcast        int    `json:"mcast" dval:"1000"`
	UnknownUcast int    `json:"unknown-ucast" dval:"1000"`
	Uuid         string `json:"uuid"`
}

type SystemPlatformtype1893 struct {
	Uuid string `json:"uuid"`
}

type SystemPortCount1894 struct {
	PortCountKernel  int    `json:"port-count-kernel" dval:"18000"`
	PortCountHm      int    `json:"port-count-hm" dval:"1024"`
	PortCountLogging int    `json:"port-count-logging" dval:"4096"`
	PortCountAlg     int    `json:"port-count-alg" dval:"6000"`
	Uuid             string `json:"uuid"`
}

type SystemPortInfo1895 struct {
	Uuid string `json:"uuid"`
}

type SystemPortList1896 struct {
	Uuid string `json:"uuid"`
}

type SystemPorts1897 struct {
	LinkDetectionInterval int    `json:"link-detection-interval" dval:"1000"`
	Uuid                  string `json:"uuid"`
}

type SystemPowerOnSelfTest1898 struct {
	Uuid string `json:"uuid"`
}

type SystemProbeNetworkDevices1899 struct {
}

type SystemPsuInfo1900 struct {
	Uuid string `json:"uuid"`
}

type SystemQInQ1901 struct {
	EnableAllPorts int    `json:"enable-all-ports"`
	InnerTpid      string `json:"inner-tpid"`
	OuterTpid      string `json:"outer-tpid"`
	Uuid           string `json:"uuid"`
}

type SystemQueuingBuffer1902 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemRadius1903 struct {
	Server SystemRadiusServer1904 `json:"server"`
}

type SystemRadiusServer1904 struct {
	ListenPort              int                                    `json:"listen-port" dval:"1813"`
	Remote                  SystemRadiusServerRemote1905           `json:"remote"`
	Secret                  int                                    `json:"secret"`
	SecretString            string                                 `json:"secret-string"`
	Encrypted               string                                 `json:"encrypted"`
	Vrid                    int                                    `json:"vrid"`
	Attribute               []SystemRadiusServerAttribute1907      `json:"attribute"`
	DisableReply            int                                    `json:"disable-reply"`
	AccountingStart         string                                 `json:"accounting-start" dval:"append-entry"`
	AccountingStop          string                                 `json:"accounting-stop" dval:"delete-entry"`
	AccountingInterimUpdate string                                 `json:"accounting-interim-update" dval:"ignore"`
	AccountingOn            string                                 `json:"accounting-on" dval:"ignore"`
	AttributeName           string                                 `json:"attribute-name"`
	CustomAttributeName     string                                 `json:"custom-attribute-name"`
	Uuid                    string                                 `json:"uuid"`
	SamplingEnable          []SystemRadiusServerSamplingEnable1908 `json:"sampling-enable"`
	DerivedAttribute        SystemRadiusServerDerivedAttribute1909 `json:"derived-attribute"`
}

type SystemRadiusServerRemote1905 struct {
	IpList []SystemRadiusServerRemoteIpList1906 `json:"ip-list"`
}

type SystemRadiusServerRemoteIpList1906 struct {
	IpListName         string `json:"ip-list-name"`
	IpListSecret       int    `json:"ip-list-secret"`
	IpListSecretString string `json:"ip-list-secret-string"`
	IpListEncrypted    string `json:"ip-list-encrypted"`
}

type SystemRadiusServerAttribute1907 struct {
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

type SystemRadiusServerSamplingEnable1908 struct {
	Counters1 string `json:"counters1"`
}

type SystemRadiusServerDerivedAttribute1909 struct {
	Usergroup SystemRadiusServerDerivedAttributeUsergroup1910 `json:"usergroup"`
	Userid    SystemRadiusServerDerivedAttributeUserid1911    `json:"userid"`
}

type SystemRadiusServerDerivedAttributeUsergroup1910 struct {
	Attribute string `json:"attribute"`
	Regex     string `json:"regex"`
	Uuid      string `json:"uuid"`
}

type SystemRadiusServerDerivedAttributeUserid1911 struct {
	Attribute string `json:"attribute"`
	Regex     string `json:"regex"`
	Uuid      string `json:"uuid"`
}

type SystemReboot1912 struct {
	Uuid string `json:"uuid"`
}

type SystemResourceAccounting1913 struct {
	Uuid         string                                     `json:"uuid"`
	TemplateList []SystemResourceAccountingTemplateList1914 `json:"template-list"`
}

type SystemResourceAccountingTemplateList1914 struct {
	Name             string                                                   `json:"name"`
	Uuid             string                                                   `json:"uuid"`
	UserTag          string                                                   `json:"user-tag"`
	AppResources     SystemResourceAccountingTemplateListAppResources1915     `json:"app-resources"`
	NetworkResources SystemResourceAccountingTemplateListNetworkResources1947 `json:"network-resources"`
	SystemResources  SystemResourceAccountingTemplateListSystemResources1957  `json:"system-resources"`
}

type SystemResourceAccountingTemplateListAppResources1915 struct {
	GslbDeviceCfg            SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1916            `json:"gslb-device-cfg"`
	GslbGeoLocationCfg       SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1917       `json:"gslb-geo-location-cfg"`
	GslbIpListCfg            SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1918            `json:"gslb-ip-list-cfg"`
	GslbPolicyCfg            SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1919            `json:"gslb-policy-cfg"`
	GslbServiceCfg           SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1920           `json:"gslb-service-cfg"`
	GslbServiceIpCfg         SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1921         `json:"gslb-service-ip-cfg"`
	GslbServicePortCfg       SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1922       `json:"gslb-service-port-cfg"`
	GslbSiteCfg              SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1923              `json:"gslb-site-cfg"`
	GslbSvcGroupCfg          SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1924          `json:"gslb-svc-group-cfg"`
	GslbTemplateCfg          SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1925          `json:"gslb-template-cfg"`
	GslbZoneCfg              SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1926              `json:"gslb-zone-cfg"`
	HealthMonitorCfg         SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1927         `json:"health-monitor-cfg"`
	RealPortCfg              SystemResourceAccountingTemplateListAppResourcesRealPortCfg1928              `json:"real-port-cfg"`
	RealServerCfg            SystemResourceAccountingTemplateListAppResourcesRealServerCfg1929            `json:"real-server-cfg"`
	ServiceGroupCfg          SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1930          `json:"service-group-cfg"`
	VirtualServerCfg         SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1931         `json:"virtual-server-cfg"`
	VirtualPortCfg           SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1932           `json:"virtual-port-cfg"`
	CacheTemplateCfg         SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1933         `json:"cache-template-cfg"`
	ClientSslTemplateCfg     SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1934     `json:"client-ssl-template-cfg"`
	ConnReuseTemplateCfg     SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1935     `json:"conn-reuse-template-cfg"`
	FastTcpTemplateCfg       SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1936       `json:"fast-tcp-template-cfg"`
	FastUdpTemplateCfg       SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1937       `json:"fast-udp-template-cfg"`
	FixTemplateCfg           SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1938           `json:"fix-template-cfg"`
	HttpTemplateCfg          SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1939          `json:"http-template-cfg"`
	LinkCostTemplateCfg      SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1940      `json:"link-cost-template-cfg"`
	PbslbEntryCfg            SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1941            `json:"pbslb-entry-cfg"`
	PersistCookieTemplateCfg SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1942 `json:"persist-cookie-template-cfg"`
	PersistSrcipTemplateCfg  SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1943  `json:"persist-srcip-template-cfg"`
	ServerSslTemplateCfg     SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1944     `json:"server-ssl-template-cfg"`
	ProxyTemplateCfg         SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1945         `json:"proxy-template-cfg"`
	StreamTemplateCfg        SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1946        `json:"stream-template-cfg"`
	Threshold                int                                                                          `json:"threshold"`
	Uuid                     string                                                                       `json:"uuid"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1916 struct {
	GslbDeviceMax          int `json:"gslb-device-max"`
	GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1917 struct {
	GslbGeoLocationMax          int `json:"gslb-geo-location-max"`
	GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1918 struct {
	GslbIpListMax          int `json:"gslb-ip-list-max"`
	GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1919 struct {
	GslbPolicyMax          int `json:"gslb-policy-max"`
	GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1920 struct {
	GslbServiceMax          int `json:"gslb-service-max"`
	GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1921 struct {
	GslbServiceIpMax          int `json:"gslb-service-ip-max"`
	GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1922 struct {
	GslbServicePortMax          int `json:"gslb-service-port-max"`
	GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1923 struct {
	GslbSiteMax          int `json:"gslb-site-max"`
	GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1924 struct {
	GslbSvcGroupMax          int `json:"gslb-svc-group-max"`
	GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1925 struct {
	GslbTemplateMax          int `json:"gslb-template-max"`
	GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1926 struct {
	GslbZoneMax          int `json:"gslb-zone-max"`
	GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1927 struct {
	HealthMonitorMax          int `json:"health-monitor-max"`
	HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesRealPortCfg1928 struct {
	RealPortMax          int `json:"real-port-max"`
	RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesRealServerCfg1929 struct {
	RealServerMax          int `json:"real-server-max"`
	RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1930 struct {
	ServiceGroupMax          int `json:"service-group-max"`
	ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1931 struct {
	VirtualServerMax          int `json:"virtual-server-max"`
	VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1932 struct {
	VirtualPortMax          int `json:"virtual-port-max"`
	VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1933 struct {
	CacheTemplateMax          int `json:"cache-template-max"`
	CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1934 struct {
	ClientSslTemplateMax          int `json:"client-ssl-template-max"`
	ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1935 struct {
	ConnReuseTemplateMax          int `json:"conn-reuse-template-max"`
	ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1936 struct {
	FastTcpTemplateMax          int `json:"fast-tcp-template-max"`
	FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1937 struct {
	FastUdpTemplateMax          int `json:"fast-udp-template-max"`
	FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1938 struct {
	FixTemplateMax          int `json:"fix-template-max"`
	FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1939 struct {
	HttpTemplateMax          int `json:"http-template-max"`
	HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1940 struct {
	LinkCostTemplateMax          int `json:"link-cost-template-max"`
	LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1941 struct {
	PbslbEntryMax          int `json:"pbslb-entry-max"`
	PbslbEntryMinGuarantee int `json:"pbslb-entry-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1942 struct {
	PersistCookieTemplateMax          int `json:"persist-cookie-template-max"`
	PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1943 struct {
	PersistSrcipTemplateMax          int `json:"persist-srcip-template-max"`
	PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1944 struct {
	ServerSslTemplateMax          int `json:"server-ssl-template-max"`
	ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1945 struct {
	ProxyTemplateMax          int `json:"proxy-template-max"`
	ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1946 struct {
	StreamTemplateMax          int `json:"stream-template-max"`
	StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResources1947 struct {
	StaticIpv4RouteCfg   SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1948   `json:"static-ipv4-route-cfg"`
	StaticIpv6RouteCfg   SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1949   `json:"static-ipv6-route-cfg"`
	Ipv4AclLineCfg       SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1950       `json:"ipv4-acl-line-cfg"`
	Ipv6AclLineCfg       SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1951       `json:"ipv6-acl-line-cfg"`
	StaticArpCfg         SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1952         `json:"static-arp-cfg"`
	StaticNeighborCfg    SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1953    `json:"static-neighbor-cfg"`
	StaticMacCfg         SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1954         `json:"static-mac-cfg"`
	ObjectGroupCfg       SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1955       `json:"object-group-cfg"`
	ObjectGroupClauseCfg SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1956 `json:"object-group-clause-cfg"`
	Threshold            int                                                                          `json:"threshold"`
	Uuid                 string                                                                       `json:"uuid"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1948 struct {
	StaticIpv4RouteMax          int `json:"static-ipv4-route-max"`
	StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1949 struct {
	StaticIpv6RouteMax          int `json:"static-ipv6-route-max"`
	StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1950 struct {
	Ipv4AclLineMax          int `json:"ipv4-acl-line-max"`
	Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1951 struct {
	Ipv6AclLineMax          int `json:"ipv6-acl-line-max"`
	Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1952 struct {
	StaticArpMax          int `json:"static-arp-max"`
	StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1953 struct {
	StaticNeighborMax          int `json:"static-neighbor-max"`
	StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1954 struct {
	StaticMacMax          int `json:"static-mac-max"`
	StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1955 struct {
	ObjectGroupMax          int `json:"object-group-max"`
	ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}

type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1956 struct {
	ObjectGroupClauseMax          int `json:"object-group-clause-max"`
	ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}

type SystemResourceAccountingTemplateListSystemResources1957 struct {
	BwLimitCfg                SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1958                `json:"bw-limit-cfg"`
	ConcurrentSessionLimitCfg SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1959 `json:"concurrent-session-limit-cfg"`
	L4SessionLimitCfg         SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1960         `json:"l4-session-limit-cfg"`
	L4cpsLimitCfg             SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1961             `json:"l4cps-limit-cfg"`
	L7cpsLimitCfg             SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1962             `json:"l7cps-limit-cfg"`
	NatcpsLimitCfg            SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1963            `json:"natcps-limit-cfg"`
	FwcpsLimitCfg             SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1964             `json:"fwcps-limit-cfg"`
	SslThroughputLimitCfg     SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1965     `json:"ssl-throughput-limit-cfg"`
	SslcpsLimitCfg            SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1966            `json:"sslcps-limit-cfg"`
	Threshold                 int                                                                              `json:"threshold"`
	Uuid                      string                                                                           `json:"uuid"`
}

type SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1958 struct {
	BwLimitMax              int `json:"bw-limit-max"`
	BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1959 struct {
	ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1960 struct {
	L4SessionLimitMax          string `json:"l4-session-limit-max"`
	L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}

type SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1961 struct {
	L4cpsLimitMax int `json:"l4cps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1962 struct {
	L7cpsLimitMax int `json:"l7cps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1963 struct {
	NatcpsLimitMax int `json:"natcps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1964 struct {
	FwcpsLimitMax int `json:"fwcps-limit-max"`
}

type SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1965 struct {
	SslThroughputLimitMax              int `json:"ssl-throughput-limit-max"`
	SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}

type SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1966 struct {
	SslcpsLimitMax int `json:"sslcps-limit-max"`
}

type SystemResourceUsage1967 struct {
	NatPoolAddrCount              int                               `json:"nat-pool-addr-count"`
	L4SessionCount                int                               `json:"l4-session-count"`
	AuthPortalHtmlFileSize        int                               `json:"auth-portal-html-file-size" dval:"20"`
	AuthPortalImageFileSize       int                               `json:"auth-portal-image-file-size" dval:"6"`
	MaxAflexFileSize              int                               `json:"max-aflex-file-size" dval:"32"`
	AflexTableEntryCount          int                               `json:"aflex-table-entry-count"`
	ClassListIpv6AddrCount        int                               `json:"class-list-ipv6-addr-count"`
	ClassListAcEntryCount         int                               `json:"class-list-ac-entry-count"`
	ClassListEntryCount           int                               `json:"class-list-entry-count"`
	MaxAflexAuthzCollectionNumber int                               `json:"max-aflex-authz-collection-number" dval:"512"`
	RadiusTableSize               int                               `json:"radius-table-size"`
	AuthzPolicyNumber             int                               `json:"authz-policy-number"`
	IpsecSaNumber                 int                               `json:"ipsec-sa-number"`
	RamCacheMemoryLimit           int                               `json:"ram-cache-memory-limit"`
	AuthSessionCount              int                               `json:"auth-session-count"`
	NgwafCacheEntry               int                               `json:"ngwaf-cache-entry"`
	Uuid                          string                            `json:"uuid"`
	Visibility                    SystemResourceUsageVisibility1968 `json:"visibility"`
}

type SystemResourceUsageVisibility1968 struct {
	MonitoredEntityCount int    `json:"monitored-entity-count"`
	Uuid                 string `json:"uuid"`
}

type SystemSession1969 struct {
	Uuid           string                            `json:"uuid"`
	SamplingEnable []SystemSessionSamplingEnable1970 `json:"sampling-enable"`
}

type SystemSessionSamplingEnable1970 struct {
	Counters1 string `json:"counters1"`
}

type SystemSessionReclaimLimit1971 struct {
	NscanLimit int    `json:"nscan-limit" dval:"4096"`
	ScanFreq   int    `json:"scan-freq" dval:"5"`
	Uuid       string `json:"uuid"`
}

type SystemSetRxtxDescSize1972 struct {
	PortIndex int `json:"port-index"`
	RxdSize   int `json:"rxd-size"`
	TxdSize   int `json:"txd-size"`
}

type SystemSetRxtxQueue1973 struct {
	PortIndex int `json:"port-index"`
	RxqSize   int `json:"rxq-size"`
	TxqSize   int `json:"txq-size"`
}

type SystemSetTcpSynPerSec1974 struct {
	TcpSynValue int    `json:"tcp-syn-value" dval:"70"`
	Uuid        string `json:"uuid"`
}

type SystemSharedPollMode1975 struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemShellPrivileges1976 struct {
	EnableShellPrivileges int    `json:"enable-shell-privileges"`
	Uuid                  string `json:"uuid"`
}

type SystemShmLogging1977 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemShutdown1978 struct {
	Uuid string `json:"uuid"`
}

type SystemSoftwareTcam1979 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemSpeProfile1980 struct {
	Action string `json:"action" dval:"ipv4-ipv6"`
}

type SystemSpeStatus1981 struct {
	Uuid string `json:"uuid"`
}

type SystemSslHwMemory1982 struct {
	MemBlockCfg []SystemSslHwMemoryMemBlockCfg1983 `json:"mem-block-cfg"`
	Uuid        string                             `json:"uuid"`
}

type SystemSslHwMemoryMemBlockCfg1983 struct {
	MemBlock string `json:"mem-block"`
	Size     int    `json:"size"`
}

type SystemSslReqQ1984 struct {
	Uuid           string                            `json:"uuid"`
	SamplingEnable []SystemSslReqQSamplingEnable1985 `json:"sampling-enable"`
}

type SystemSslReqQSamplingEnable1985 struct {
	Counters1 string `json:"counters1"`
}

type SystemSslScv1986 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemSslScvVerifyCrlSign1987 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
}

type SystemSslScvVerifyHost1988 struct {
	Disable int    `json:"disable"`
	Uuid    string `json:"uuid"`
}

type SystemSslSetCompatibleCipher1989 struct {
	Disable int    `json:"disable"`
	Uuid    string `json:"uuid"`
}

type SystemSslStatus1990 struct {
	Uuid string `json:"uuid"`
}

type SystemSyslogTimeMsec1991 struct {
	EnableFlag int `json:"enable-flag"`
}

type SystemTableIntegrity1992 struct {
	Table          string                                   `json:"table" dval:"all"`
	AuditAction    string                                   `json:"audit-action" dval:"enable"`
	AutoSyncAction string                                   `json:"auto-sync-action" dval:"enable"`
	Uuid           string                                   `json:"uuid"`
	SamplingEnable []SystemTableIntegritySamplingEnable1993 `json:"sampling-enable"`
}

type SystemTableIntegritySamplingEnable1993 struct {
	Counters1 string `json:"counters1"`
	Counters2 string `json:"counters2"`
	Counters3 string `json:"counters3"`
}

type SystemTcp1994 struct {
	Uuid                      string                                 `json:"uuid"`
	SamplingEnable            []SystemTcpSamplingEnable1995          `json:"sampling-enable"`
	RateLimitResetUnknownConn SystemTcpRateLimitResetUnknownConn1996 `json:"rate-limit-reset-unknown-conn"`
}

type SystemTcpSamplingEnable1995 struct {
	Counters1 string `json:"counters1"`
}

type SystemTcpRateLimitResetUnknownConn1996 struct {
	PktRateForResetUnknownConn int    `json:"pkt-rate-for-reset-unknown-conn"`
	LogForResetUnknownConn     int    `json:"log-for-reset-unknown-conn"`
	Uuid                       string `json:"uuid"`
}

type SystemTcpStats1997 struct {
	Uuid           string                             `json:"uuid"`
	SamplingEnable []SystemTcpStatsSamplingEnable1998 `json:"sampling-enable"`
}

type SystemTcpStatsSamplingEnable1998 struct {
	Counters1 string `json:"counters1"`
}

type SystemTcpSynPerSec1999 struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLog2000 struct {
	TopKSourceList   SystemTelemetryLogTopKSourceList2001   `json:"top-k-source-list"`
	TopKAppSvcList   SystemTelemetryLogTopKAppSvcList2002   `json:"top-k-app-svc-list"`
	DeviceStatus     SystemTelemetryLogDeviceStatus2003     `json:"device-status"`
	Environment      SystemTelemetryLogEnvironment2004      `json:"environment"`
	PartitionMetrics SystemTelemetryLogPartitionMetrics2005 `json:"partition-metrics"`
}

type SystemTelemetryLogTopKSourceList2001 struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogTopKAppSvcList2002 struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogDeviceStatus2003 struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogEnvironment2004 struct {
	Uuid string `json:"uuid"`
}

type SystemTelemetryLogPartitionMetrics2005 struct {
	Uuid string `json:"uuid"`
}

type SystemTemplate2006 struct {
	TemplatePolicy string `json:"template-policy"`
	Uuid           string `json:"uuid"`
}

type SystemTemplateBind2007 struct {
	MonitorList []SystemTemplateBindMonitorList `json:"monitor-list"`
}

type SystemTemplateBindMonitorList struct {
	TemplateMonitor int    `json:"template-monitor"`
	Uuid            string `json:"uuid"`
}

type SystemThroughput2008 struct {
	Uuid           string                               `json:"uuid"`
	SamplingEnable []SystemThroughputSamplingEnable2009 `json:"sampling-enable"`
}

type SystemThroughputSamplingEnable2009 struct {
	Counters1 string `json:"counters1"`
}

type SystemTimeoutValue2010 struct {
	Ftp   int    `json:"ftp" dval:"120"`
	Scp   int    `json:"scp" dval:"300"`
	Sftp  int    `json:"sftp"`
	Tftp  int    `json:"tftp" dval:"300"`
	Http  int    `json:"http" dval:"120"`
	Https int    `json:"https" dval:"120"`
	Uuid  string `json:"uuid"`
}

type SystemTrunk2011 struct {
	LoadBalance SystemTrunkLoadBalance2012 `json:"load-balance"`
}

type SystemTrunkLoadBalance2012 struct {
	UseL3 int    `json:"use-l3"`
	UseL4 int    `json:"use-l4"`
	Uuid  string `json:"uuid"`
}

type SystemTrunkHwHash2013 struct {
	Mode int    `json:"mode" dval:"6"`
	Uuid string `json:"uuid"`
}

type SystemTrunkXauiHwHash2014 struct {
	Mode int    `json:"mode" dval:"6"`
	Uuid string `json:"uuid"`
}

type SystemTso2015 struct {
	Enable  int `json:"enable"`
	Disable int `json:"disable"`
}

type SystemUdp2016 struct {
	SkipChecksumWhenZero int    `json:"skip-checksum-when-zero"`
	Uuid                 string `json:"uuid"`
}

type SystemUpgradeStatus2017 struct {
	Uuid string `json:"uuid"`
}

type SystemVeMacScheme2018 struct {
	VeMacSchemeVal string `json:"ve-mac-scheme-val" dval:"hash-based"`
	Uuid           string `json:"uuid"`
}

type SystemXauiDlbMode2019 struct {
	Enable int    `json:"enable"`
	Uuid   string `json:"uuid"`
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
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
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
