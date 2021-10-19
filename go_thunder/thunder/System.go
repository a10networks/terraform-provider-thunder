package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type System struct {
	SystemInstanceAnomalyLog SystemInstance `json:"system,omitempty"`
}

type SystemInstance struct {
	SystemInstanceAddCPUCoreCoreIndex                  SystemInstanceAddCPUCore              `json:"add-cpu-core,omitempty"`
	SystemInstanceAddPortPortIndex                     SystemInstanceAddPort                 `json:"add-port,omitempty"`
	SystemInstanceAllVlanLimitBcast                    SystemInstanceAllVlanLimit            `json:"all-vlan-limit,omitempty"`
	SystemInstanceAnomalyLog                           int                                   `json:"anomaly-log,omitempty"`
	SystemInstanceAppPerformanceUUID                   SystemInstanceAppPerformance          `json:"app-performance,omitempty"`
	SystemInstanceAppsGlobalLogSessionOnEstablished    SystemInstanceAppsGlobal              `json:"apps-global,omitempty"`
	SystemInstanceAttackLog                            int                                   `json:"attack-log,omitempty"`
	SystemInstanceBandwidthUUID                        SystemInstanceBandwidth               `json:"bandwidth,omitempty"`
	SystemInstanceBfdUUID                              SystemInstanceBfd                     `json:"bfd,omitempty"`
	SystemInstanceCPUHyperThreadEnable                 SystemInstanceCPUHyperThread          `json:"cpu-hyper-thread,omitempty"`
	SystemInstanceCPUListUUID                          SystemInstanceCPUList                 `json:"cpu-list,omitempty"`
	SystemInstanceCPULoadSharingDisable                SystemInstanceCPULoadSharing          `json:"cpu-load-sharing,omitempty"`
	SystemInstanceCPUMapUUID                           SystemInstanceCPUMap                  `json:"cpu-map,omitempty"`
	SystemInstanceClassListHitcountEnable              int                                   `json:"class-list-hitcount-enable,omitempty"`
	SystemInstanceCmUpdateFileNameRefSourceName        SystemInstanceCmUpdateFileNameRef     `json:"cm-update-file-name-ref,omitempty"`
	SystemInstanceControlCPUUUID                       SystemInstanceControlCPU              `json:"control-cpu,omitempty"`
	SystemInstanceCoreUUID                             SystemInstanceCore                    `json:"core,omitempty"`
	SystemInstanceCosqShowUUID                         SystemInstanceCosqShow                `json:"cosq-show,omitempty"`
	SystemInstanceCosqStatsUUID                        SystemInstanceCosqStats               `json:"cosq-stats,omitempty"`
	SystemInstanceCounterLibAccountingUUID             SystemInstanceCounterLibAccounting    `json:"counter-lib-accounting,omitempty"`
	SystemInstanceDNSUUID                              SystemInstanceDNS                     `json:"dns,omitempty"`
	SystemInstanceDNSCacheUUID                         SystemInstanceDNSCache                `json:"dns-cache,omitempty"`
	SystemInstanceDataCPUUUID                          SystemInstanceDataCPU                 `json:"data-cpu,omitempty"`
	SystemInstanceDdosAttack                           int                                   `json:"ddos-attack,omitempty"`
	SystemInstanceDdosLog                              int                                   `json:"ddos-log,omitempty"`
	SystemInstanceDeepHrxqEnable                       SystemInstanceDeepHrxq                `json:"deep-hrxq,omitempty"`
	SystemInstanceDelPortPortIndex                     SystemInstanceDelPort                 `json:"del-port,omitempty"`
	SystemInstanceDeleteCPUCoreCoreIndex               SystemInstanceDeleteCPUCore           `json:"delete-cpu-core,omitempty"`
	SystemInstanceDomainListHitcountEnable             int                                   `json:"domain-list-hitcount-enable,omitempty"`
	SystemInstanceDomainListInfoUUID                   SystemInstanceDomainListInfo          `json:"domain-list-info,omitempty"`
	SystemInstanceDpdkStatsUUID                        SystemInstanceDpdkStats               `json:"dpdk-stats,omitempty"`
	SystemInstanceDynamicServiceDNSSocketPool          int                                   `json:"dynamic-service-dns-socket-pool,omitempty"`
	SystemInstanceEnvironmentUUID                      SystemInstanceEnvironment             `json:"environment,omitempty"`
	SystemInstanceFpgaCoreCrcMonitorDisable            SystemInstanceFpgaCoreCrc             `json:"fpga-core-crc,omitempty"`
	SystemInstanceFpgaDropUUID                         SystemInstanceFpgaDrop                `json:"fpga-drop,omitempty"`
	SystemInstanceFwApplicationMempool                 SystemInstanceFw                      `json:"fw,omitempty"`
	SystemInstanceGeoDbHitcountEnable                  int                                   `json:"geo-db-hitcount-enable,omitempty"`
	SystemInstanceGeoLocationGeoLocationIana           SystemInstanceGeoLocation             `json:"geo-location,omitempty"`
	SystemInstanceGeolocUUID                           SystemInstanceGeoloc                  `json:"geoloc,omitempty"`
	SystemInstanceGeolocListListName                   []SystemInstanceGeolocListList        `json:"geoloc-list-list,omitempty"`
	SystemInstanceGeolocNameHelperUUID                 SystemInstanceGeolocNameHelper        `json:"geoloc-name-helper,omitempty"`
	SystemInstanceGeolocationFileUUID                  SystemInstanceGeolocationFile         `json:"geolocation-file,omitempty"`
	SystemInstanceGlid                                 int                                   `json:"glid,omitempty"`
	SystemInstanceGuestFileUUID                        SystemInstanceGuestFile               `json:"guest-file,omitempty"`
	SystemInstanceGuiImageListUUID                     SystemInstanceGuiImageList            `json:"gui-image-list,omitempty"`
	SystemInstanceHardwareUUID                         SystemInstanceHardware                `json:"hardware,omitempty"`
	SystemInstanceHardwareForwardUUID                  SystemInstanceHardwareForward         `json:"hardware-forward,omitempty"`
	SystemInstanceHealthCheckListL2HmHcName            []SystemInstanceHealthCheckList       `json:"health-check-list,omitempty"`
	SystemInstanceHighMemoryL4SessionEnable            SystemInstanceHighMemoryL4Session     `json:"high-memory-l4-session,omitempty"`
	SystemInstanceHrxqStatusUUID                       SystemInstanceHrxqStatus              `json:"hrxq-status,omitempty"`
	SystemInstanceIP6StatsUUID                         SystemInstanceIP6Stats                `json:"ip6-stats,omitempty"`
	SystemInstanceIPDNSCacheUUID                       SystemInstanceIPDNSCache              `json:"ip-dns-cache,omitempty"`
	SystemInstanceIPStatsUUID                          SystemInstanceIPStats                 `json:"ip-stats,omitempty"`
	SystemInstanceIPThreatListUUID                     SystemInstanceIPThreatList            `json:"ip-threat-list,omitempty"`
	SystemInstanceIcmpUUID                             SystemInstanceIcmp                    `json:"icmp,omitempty"`
	SystemInstanceIcmp6UUID                            SystemInstanceIcmp6                   `json:"icmp6,omitempty"`
	SystemInstanceIcmpRateUUID                         SystemInstanceIcmpRate                `json:"icmp-rate,omitempty"`
	SystemInstanceInuseCPUListUUID                     SystemInstanceInuseCPUList            `json:"inuse-cpu-list,omitempty"`
	SystemInstanceInusePortListUUID                    SystemInstanceInusePortList           `json:"inuse-port-list,omitempty"`
	SystemInstanceIoCPUMaxCores                        SystemInstanceIoCPU                   `json:"io-cpu,omitempty"`
	SystemInstanceIpmiReset                            SystemInstanceIpmi                    `json:"ipmi,omitempty"`
	SystemInstanceIpmiServiceDisable                   SystemInstanceIpmiService             `json:"ipmi-service,omitempty"`
	SystemInstanceIpsecPacketRoundRobin                SystemInstanceIpsec                   `json:"ipsec,omitempty"`
	SystemInstanceJobOffloadUUID                       SystemInstanceJobOffload              `json:"job-offload,omitempty"`
	SystemInstanceLinkCapabilityEnable                 SystemInstanceLinkCapability          `json:"link-capability,omitempty"`
	SystemInstanceLinkMonitorEnable                    SystemInstanceLinkMonitor             `json:"link-monitor,omitempty"`
	SystemInstanceLroEnable                            SystemInstanceLro                     `json:"lro,omitempty"`
	SystemInstanceManagementInterfaceModeDedicated     SystemInstanceManagementInterfaceMode `json:"management-interface-mode,omitempty"`
	SystemInstanceMemoryUUID                           SystemInstanceMemory                  `json:"memory,omitempty"`
	SystemInstanceMfaAuthUsername                      SystemInstanceMfaAuth                 `json:"mfa-auth,omitempty"`
	SystemInstanceMfaCertStoreCertHost                 SystemInstanceMfaCertStore            `json:"mfa-cert-store,omitempty"`
	SystemInstanceMfaManagementEnable                  SystemInstanceMfaManagement           `json:"mfa-management,omitempty"`
	SystemInstanceMfaValidationTypeCaCert              SystemInstanceMfaValidationType       `json:"mfa-validation-type,omitempty"`
	SystemInstanceMgmtPortPortIndex                    SystemInstanceMgmtPort                `json:"mgmt-port,omitempty"`
	SystemInstanceModifyPortPortIndex                  SystemInstanceModifyPort              `json:"modify-port,omitempty"`
	SystemInstanceModuleCtrlCPU                        string                                `json:"module-ctrl-cpu,omitempty"`
	SystemInstanceMonTemplateMonitorList               SystemInstanceMonTemplate             `json:"mon-template,omitempty"`
	SystemInstanceMultiQueueSupportEnable              SystemInstanceMultiQueueSupport       `json:"multi-queue-support,omitempty"`
	SystemInstanceNdiscRaUUID                          SystemInstanceNdiscRa                 `json:"ndisc-ra,omitempty"`
	SystemInstancePasswordPolicyComplexity             SystemInstancePasswordPolicy          `json:"password-policy,omitempty"`
	SystemInstancePathListL2HmPathName                 []SystemInstancePathList              `json:"path-list,omitempty"`
	SystemInstancePerVlanLimitBcast                    SystemInstancePerVlanLimit            `json:"per-vlan-limit,omitempty"`
	SystemInstancePlatformtypeUUID                     SystemInstancePlatformtype            `json:"platformtype,omitempty"`
	SystemInstancePortCountPortCountKernel             SystemInstancePortCount               `json:"port-count,omitempty"`
	SystemInstancePortInfoUUID                         SystemInstancePortInfo                `json:"port-info,omitempty"`
	SystemInstancePortListUUID                         SystemInstancePortList                `json:"port-list,omitempty"`
	SystemInstancePortsLinkDetectionInterval           SystemInstancePorts                   `json:"ports,omitempty"`
	SystemInstancePromiscuousMode                      int                                   `json:"promiscuous-mode,omitempty"`
	SystemInstancePsuInfoUUID                          SystemInstancePsuInfo                 `json:"psu-info,omitempty"`
	SystemInstanceQInQInnerTpid                        SystemInstanceQInQ                    `json:"q-in-q,omitempty"`
	SystemInstanceQueuingBufferEnable                  SystemInstanceQueuingBuffer           `json:"queuing-buffer,omitempty"`
	SystemInstanceRadiusServer                         SystemInstanceRadius                  `json:"radius,omitempty"`
	SystemInstanceRebootUUID                           SystemInstanceReboot                  `json:"reboot,omitempty"`
	SystemInstanceResourceAccountingUUID               SystemInstanceResourceAccounting      `json:"resource-accounting,omitempty"`
	SystemInstanceResourceUsageSslContextMemory        SystemInstanceResourceUsage           `json:"resource-usage,omitempty"`
	SystemInstanceSessionUUID                          SystemInstanceSession                 `json:"session,omitempty"`
	SystemInstanceSessionReclaimLimitNscanLimit        SystemInstanceSessionReclaimLimit     `json:"session-reclaim-limit,omitempty"`
	SystemInstanceSetRxtxDescSizePortIndex             SystemInstanceSetRxtxDescSize         `json:"set-rxtx-desc-size,omitempty"`
	SystemInstanceSetRxtxQueuePortIndex                SystemInstanceSetRxtxQueue            `json:"set-rxtx-queue,omitempty"`
	SystemInstanceSetTCPSynPerSecTCPSynValue           SystemInstanceSetTCPSynPerSec         `json:"set-tcp-syn-per-sec,omitempty"`
	SystemInstanceSharedPollModeEnable                 SystemInstanceSharedPollMode          `json:"shared-poll-mode,omitempty"`
	SystemInstanceShellPrivilegesEnableShellPrivileges SystemInstanceShellPrivileges         `json:"shell-privileges,omitempty"`
	SystemInstanceShmLoggingEnable                     SystemInstanceShmLogging              `json:"shm-logging,omitempty"`
	SystemInstanceShutdownUUID                         SystemInstanceShutdown                `json:"shutdown,omitempty"`
	SystemInstanceSockstressDisable                    int                                   `json:"sockstress-disable,omitempty"`
	SystemInstanceSpeProfileAction                     SystemInstanceSpeProfile              `json:"spe-profile,omitempty"`
	SystemInstanceSpeStatusUUID                        SystemInstanceSpeStatus               `json:"spe-status,omitempty"`
	SystemInstanceSrcIPHashEnable                      int                                   `json:"src-ip-hash-enable,omitempty"`
	SystemInstanceSslReqQUUID                          SystemInstanceSslReqQ                 `json:"ssl-req-q,omitempty"`
	SystemInstanceSslScvEnable                         SystemInstanceSslScv                  `json:"ssl-scv,omitempty"`
	SystemInstanceSslScvVerifyCrlSignEnable            SystemInstanceSslScvVerifyCrlSign     `json:"ssl-scv-verify-crl-sign,omitempty"`
	SystemInstanceSslScvVerifyHostDisable              SystemInstanceSslScvVerifyHost        `json:"ssl-scv-verify-host,omitempty"`
	SystemInstanceSslSetCompatibleCipherDisable        SystemInstanceSslSetCompatibleCipher  `json:"ssl-set-compatible-cipher,omitempty"`
	SystemInstanceSslStatusUUID                        SystemInstanceSslStatus               `json:"ssl-status,omitempty"`
	SystemInstanceSyslogTimeMsecEnableFlag             SystemInstanceSyslogTimeMsec          `json:"syslog-time-msec,omitempty"`
	SystemInstanceSystemChassisPortSplitEnable         int                                   `json:"system-chassis-port-split-enable,omitempty"`
	SystemInstanceTCPUUID                              SystemInstanceTCP                     `json:"tcp,omitempty"`
	SystemInstanceTCPStatsUUID                         SystemInstanceTCPStats                `json:"tcp-stats,omitempty"`
	SystemInstanceTCPSynPerSecUUID                     SystemInstanceTCPSynPerSec            `json:"tcp-syn-per-sec,omitempty"`
	SystemInstanceTLS13MgmtEnable                      SystemInstanceTLS13Mgmt               `json:"tls-1-3-mgmt,omitempty"`
	SystemInstanceTableIntegrityTable                  SystemInstanceTableIntegrity          `json:"table-integrity,omitempty"`
	SystemInstanceTelemetryLogTopKSourceList           SystemInstanceTelemetryLog            `json:"telemetry-log,omitempty"`
	SystemInstanceTemplateTemplatePolicy               SystemInstanceTemplate                `json:"template,omitempty"`
	SystemInstanceTemplateBindMonitorList              SystemInstanceTemplateBind            `json:"template-bind,omitempty"`
	SystemInstanceThroughputUUID                       SystemInstanceThroughput              `json:"throughput,omitempty"`
	SystemInstanceTimeoutValueFtp                      SystemInstanceTimeoutValue            `json:"timeout-value,omitempty"`
	SystemInstanceTrunkLoadBalance                     SystemInstanceTrunk                   `json:"trunk,omitempty"`
	SystemInstanceTrunkHwHashMode                      SystemInstanceTrunkHwHash             `json:"trunk-hw-hash,omitempty"`
	SystemInstanceTrunkXauiHwHashMode                  SystemInstanceTrunkXauiHwHash         `json:"trunk-xaui-hw-hash,omitempty"`
	SystemInstanceTsoEnable                            SystemInstanceTso                     `json:"tso,omitempty"`
	SystemInstanceUUID                                 string                                `json:"uuid,omitempty"`
	SystemInstanceUpgradeStatusUUID                    SystemInstanceUpgradeStatus           `json:"upgrade-status,omitempty"`
	SystemInstanceVeMacSchemeVeMacSchemeVal            SystemInstanceVeMacScheme             `json:"ve-mac-scheme,omitempty"`
}

type SystemInstanceAddCPUCore struct {
	SystemInstanceAddCPUCoreCoreIndex int `json:"core-index,omitempty"`
}

type SystemInstanceAddPort struct {
	SystemInstanceAddPortPortIndex int `json:"port-index,omitempty"`
}

type SystemInstanceAllVlanLimit struct {
	SystemInstanceAllVlanLimitBcast        int    `json:"bcast,omitempty"`
	SystemInstanceAllVlanLimitIpmcast      int    `json:"ipmcast,omitempty"`
	SystemInstanceAllVlanLimitMcast        int    `json:"mcast,omitempty"`
	SystemInstanceAllVlanLimitUUID         string `json:"uuid,omitempty"`
	SystemInstanceAllVlanLimitUnknownUcast int    `json:"unknown-ucast,omitempty"`
}

type SystemInstanceAppPerformance struct {
	SystemInstanceAppPerformanceSamplingEnableCounters1 []SystemInstanceAppPerformanceSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceAppPerformanceUUID                    string                                       `json:"uuid,omitempty"`
}

type SystemInstanceAppsGlobal struct {
	SystemInstanceAppsGlobalLogSessionOnEstablished int    `json:"log-session-on-established,omitempty"`
	SystemInstanceAppsGlobalMslTime                 int    `json:"msl-time,omitempty"`
	SystemInstanceAppsGlobalUUID                    string `json:"uuid,omitempty"`
}

type SystemInstanceBandwidth struct {
	SystemInstanceBandwidthSamplingEnableCounters1 []SystemInstanceBandwidthSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceBandwidthUUID                    string                                  `json:"uuid,omitempty"`
}

type SystemInstanceBfd struct {
	SystemInstanceBfdSamplingEnableCounters1 []SystemInstanceBfdSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceBfdUUID                    string                            `json:"uuid,omitempty"`
}

type SystemInstanceCPUHyperThread struct {
	SystemInstanceCPUHyperThreadDisable int `json:"disable,omitempty"`
	SystemInstanceCPUHyperThreadEnable  int `json:"enable,omitempty"`
}

type SystemInstanceCPUList struct {
	SystemInstanceCPUListUUID string `json:"uuid,omitempty"`
}

type SystemInstanceCPULoadSharing struct {
	SystemInstanceCPULoadSharingCPUUsageLow         SystemInstanceCPULoadSharingCPUUsage         `json:"cpu-usage,omitempty"`
	SystemInstanceCPULoadSharingDisable             int                                          `json:"disable,omitempty"`
	SystemInstanceCPULoadSharingPacketsPerSecondMin SystemInstanceCPULoadSharingPacketsPerSecond `json:"packets-per-second,omitempty"`
	SystemInstanceCPULoadSharingTCP                 int                                          `json:"tcp,omitempty"`
	SystemInstanceCPULoadSharingUDP                 int                                          `json:"udp,omitempty"`
	SystemInstanceCPULoadSharingUUID                string                                       `json:"uuid,omitempty"`
}

type SystemInstanceCPUMap struct {
	SystemInstanceCPUMapUUID string `json:"uuid,omitempty"`
}

type SystemInstanceCmUpdateFileNameRef struct {
	SystemInstanceCmUpdateFileNameRefDestName   string `json:"dest_name,omitempty"`
	SystemInstanceCmUpdateFileNameRefID         int    `json:"id,omitempty"`
	SystemInstanceCmUpdateFileNameRefSourceName string `json:"source_name,omitempty"`
}

type SystemInstanceControlCPU struct {
	SystemInstanceControlCPUUUID string `json:"uuid,omitempty"`
}

type SystemInstanceCore struct {
	SystemInstanceCoreUUID string `json:"uuid,omitempty"`
}

type SystemInstanceCosqShow struct {
	SystemInstanceCosqShowUUID string `json:"uuid,omitempty"`
}

type SystemInstanceCosqStats struct {
	SystemInstanceCosqStatsUUID string `json:"uuid,omitempty"`
}

type SystemInstanceCounterLibAccounting struct {
	SystemInstanceCounterLibAccountingUUID string `json:"uuid,omitempty"`
}

type SystemInstanceDNS struct {
	SystemInstanceDNSSamplingEnableCounters1 []SystemInstanceDNSSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceDNSUUID                    string                            `json:"uuid,omitempty"`
}

type SystemInstanceDNSCache struct {
	SystemInstanceDNSCacheSamplingEnableCounters1 []SystemInstanceDNSCacheSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceDNSCacheUUID                    string                                 `json:"uuid,omitempty"`
}

type SystemInstanceDataCPU struct {
	SystemInstanceDataCPUUUID string `json:"uuid,omitempty"`
}

type SystemInstanceDeepHrxq struct {
	SystemInstanceDeepHrxqEnable int `json:"enable,omitempty"`
}

type SystemInstanceDelPort struct {
	SystemInstanceDelPortPortIndex int `json:"port-index,omitempty"`
}

type SystemInstanceDeleteCPUCore struct {
	SystemInstanceDeleteCPUCoreCoreIndex int `json:"core-index,omitempty"`
}

type SystemInstanceDomainListInfo struct {
	SystemInstanceDomainListInfoUUID string `json:"uuid,omitempty"`
}

type SystemInstanceDpdkStats struct {
	SystemInstanceDpdkStatsSamplingEnableCounters1 []SystemInstanceDpdkStatsSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceDpdkStatsUUID                    string                                  `json:"uuid,omitempty"`
}

type SystemInstanceEnvironment struct {
	SystemInstanceEnvironmentUUID string `json:"uuid,omitempty"`
}

type SystemInstanceFpgaCoreCrc struct {
	SystemInstanceFpgaCoreCrcMonitorDisable int    `json:"monitor-disable,omitempty"`
	SystemInstanceFpgaCoreCrcRebootEnable   int    `json:"reboot-enable,omitempty"`
	SystemInstanceFpgaCoreCrcUUID           string `json:"uuid,omitempty"`
}

type SystemInstanceFpgaDrop struct {
	SystemInstanceFpgaDropSamplingEnableCounters1 []SystemInstanceFpgaDropSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceFpgaDropUUID                    string                                 `json:"uuid,omitempty"`
}

type SystemInstanceFw struct {
	SystemInstanceFwApplicationFlow    int    `json:"application-flow,omitempty"`
	SystemInstanceFwApplicationMempool int    `json:"application-mempool,omitempty"`
	SystemInstanceFwBasicDpiEnable     int    `json:"basic-dpi-enable,omitempty"`
	SystemInstanceFwUUID               string `json:"uuid,omitempty"`
}

type SystemInstanceGeoLocation struct {
	SystemInstanceGeoLocationEntryListGeoLocnObjName                   []SystemInstanceGeoLocationEntryList          `json:"entry-list,omitempty"`
	SystemInstanceGeoLocationGeoLocationGeolite2City                   int                                           `json:"geo-location-geolite2-city,omitempty"`
	SystemInstanceGeoLocationGeoLocationGeolite2Country                int                                           `json:"geo-location-geolite2-country,omitempty"`
	SystemInstanceGeoLocationGeoLocationIana                           int                                           `json:"geo-location-iana,omitempty"`
	SystemInstanceGeoLocationGeolite2CityIncludeIpv6                   int                                           `json:"geolite2-city-include-ipv6,omitempty"`
	SystemInstanceGeoLocationGeolite2CountryIncludeIpv6                int                                           `json:"geolite2-country-include-ipv6,omitempty"`
	SystemInstanceGeoLocationGeolocLoadFileListGeoLocationLoadFilename []SystemInstanceGeoLocationGeolocLoadFileList `json:"geoloc-load-file-list,omitempty"`
	SystemInstanceGeoLocationUUID                                      string                                        `json:"uuid,omitempty"`
}

type SystemInstanceGeoloc struct {
	SystemInstanceGeolocSamplingEnableCounters1 []SystemInstanceGeolocSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceGeolocUUID                    string                               `json:"uuid,omitempty"`
}

type SystemInstanceGeolocListList struct {
	SystemInstanceGeolocListListExcludeGeolocNameListExcludeGeolocNameVal []SystemInstanceGeolocListListExcludeGeolocNameList `json:"exclude-geoloc-name-list,omitempty"`
	SystemInstanceGeolocListListIncludeGeolocNameListIncludeGeolocNameVal []SystemInstanceGeolocListListIncludeGeolocNameList `json:"include-geoloc-name-list,omitempty"`
	SystemInstanceGeolocListListName                                      string                                              `json:"name,omitempty"`
	SystemInstanceGeolocListListSamplingEnableCounters1                   []SystemInstanceGeolocListListSamplingEnable        `json:"sampling-enable,omitempty"`
	SystemInstanceGeolocListListShared                                    int                                                 `json:"shared,omitempty"`
	SystemInstanceGeolocListListUUID                                      string                                              `json:"uuid,omitempty"`
	SystemInstanceGeolocListListUserTag                                   string                                              `json:"user-tag,omitempty"`
}

type SystemInstanceGeolocNameHelper struct {
	SystemInstanceGeolocNameHelperSamplingEnableCounters1 []SystemInstanceGeolocNameHelperSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceGeolocNameHelperUUID                    string                                         `json:"uuid,omitempty"`
}

type SystemInstanceGeolocationFile struct {
	SystemInstanceGeolocationFileErrorInfoUUID SystemInstanceGeolocationFileErrorInfo `json:"error-info,omitempty"`
	SystemInstanceGeolocationFileUUID          string                                 `json:"uuid,omitempty"`
}

type SystemInstanceGuestFile struct {
	SystemInstanceGuestFileUUID string `json:"uuid,omitempty"`
}

type SystemInstanceGuiImageList struct {
	SystemInstanceGuiImageListUUID string `json:"uuid,omitempty"`
}

type SystemInstanceHardware struct {
	SystemInstanceHardwareUUID string `json:"uuid,omitempty"`
}

type SystemInstanceHardwareForward struct {
	SystemInstanceHardwareForwardSamplingEnableCounters1 []SystemInstanceHardwareForwardSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceHardwareForwardSlbUUID                 SystemInstanceHardwareForwardSlb              `json:"slb,omitempty"`
	SystemInstanceHardwareForwardUUID                    string                                        `json:"uuid,omitempty"`
}

type SystemInstanceHealthCheckList struct {
	SystemInstanceHealthCheckListL2BfdMultiplier int    `json:"l2bfd-multiplier,omitempty"`
	SystemInstanceHealthCheckListL2BfdRxInterval int    `json:"l2bfd-rx-interval,omitempty"`
	SystemInstanceHealthCheckListL2BfdTxInterval int    `json:"l2bfd-tx-interval,omitempty"`
	SystemInstanceHealthCheckListL2HmHcName      string `json:"l2hm-hc-name,omitempty"`
	SystemInstanceHealthCheckListMethodL2Bfd     int    `json:"method-l2bfd,omitempty"`
	SystemInstanceHealthCheckListUUID            string `json:"uuid,omitempty"`
	SystemInstanceHealthCheckListUserTag         string `json:"user-tag,omitempty"`
}

type SystemInstanceHighMemoryL4Session struct {
	SystemInstanceHighMemoryL4SessionEnable int    `json:"enable,omitempty"`
	SystemInstanceHighMemoryL4SessionUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceHrxqStatus struct {
	SystemInstanceHrxqStatusUUID string `json:"uuid,omitempty"`
}

type SystemInstanceIP6Stats struct {
	SystemInstanceIP6StatsSamplingEnableCounters1 []SystemInstanceIP6StatsSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceIP6StatsUUID                    string                                 `json:"uuid,omitempty"`
}

type SystemInstanceIPDNSCache struct {
	SystemInstanceIPDNSCacheUUID string `json:"uuid,omitempty"`
}

type SystemInstanceIPStats struct {
	SystemInstanceIPStatsSamplingEnableCounters1 []SystemInstanceIPStatsSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceIPStatsUUID                    string                                `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatList struct {
	SystemInstanceIPThreatListIpv4DestListClassListCfg         SystemInstanceIPThreatListIpv4DestList         `json:"ipv4-dest-list,omitempty"`
	SystemInstanceIPThreatListIpv4InternetHostListClassListCfg SystemInstanceIPThreatListIpv4InternetHostList `json:"ipv4-internet-host-list,omitempty"`
	SystemInstanceIPThreatListIpv4SourceListClassListCfg       SystemInstanceIPThreatListIpv4SourceList       `json:"ipv4-source-list,omitempty"`
	SystemInstanceIPThreatListIpv6DestListClassListCfg         SystemInstanceIPThreatListIpv6DestList         `json:"ipv6-dest-list,omitempty"`
	SystemInstanceIPThreatListIpv6InternetHostListClassListCfg SystemInstanceIPThreatListIpv6InternetHostList `json:"ipv6-internet-host-list,omitempty"`
	SystemInstanceIPThreatListIpv6SourceListClassListCfg       SystemInstanceIPThreatListIpv6SourceList       `json:"ipv6-source-list,omitempty"`
	SystemInstanceIPThreatListSamplingEnableCounters1          []SystemInstanceIPThreatListSamplingEnable     `json:"sampling-enable,omitempty"`
	SystemInstanceIPThreatListUUID                             string                                         `json:"uuid,omitempty"`
}

type SystemInstanceIcmp struct {
	SystemInstanceIcmpSamplingEnableCounters1 []SystemInstanceIcmpSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceIcmpUUID                    string                             `json:"uuid,omitempty"`
}

type SystemInstanceIcmp6 struct {
	SystemInstanceIcmp6SamplingEnableCounters1 []SystemInstanceIcmp6SamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceIcmp6UUID                    string                              `json:"uuid,omitempty"`
}

type SystemInstanceIcmpRate struct {
	SystemInstanceIcmpRateSamplingEnableCounters1 []SystemInstanceIcmpRateSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceIcmpRateUUID                    string                                 `json:"uuid,omitempty"`
}

type SystemInstanceInuseCPUList struct {
	SystemInstanceInuseCPUListUUID string `json:"uuid,omitempty"`
}

type SystemInstanceInusePortList struct {
	SystemInstanceInusePortListUUID string `json:"uuid,omitempty"`
}

type SystemInstanceIoCPU struct {
	SystemInstanceIoCPUMaxCores int `json:"max-cores,omitempty"`
}

type SystemInstanceIpmi struct {
	SystemInstanceIpmiIPIpv4Address SystemInstanceIpmiIP    `json:"ip,omitempty"`
	SystemInstanceIpmiIpsrcDhcp     SystemInstanceIpmiIpsrc `json:"ipsrc,omitempty"`
	SystemInstanceIpmiReset         int                     `json:"reset,omitempty"`
	SystemInstanceIpmiToolCmd       SystemInstanceIpmiTool  `json:"tool,omitempty"`
	SystemInstanceIpmiUserAdd       SystemInstanceIpmiUser  `json:"user,omitempty"`
}

type SystemInstanceIpmiService struct {
	SystemInstanceIpmiServiceDisable int    `json:"disable,omitempty"`
	SystemInstanceIpmiServiceUUID    string `json:"uuid,omitempty"`
}

type SystemInstanceIpsec struct {
	SystemInstanceIpsecCryptoCore        int                            `json:"crypto-core,omitempty"`
	SystemInstanceIpsecCryptoMem         int                            `json:"crypto-mem,omitempty"`
	SystemInstanceIpsecFpgaDecryptAction SystemInstanceIpsecFpgaDecrypt `json:"fpga-decrypt,omitempty"`
	SystemInstanceIpsecPacketRoundRobin  int                            `json:"packet-round-robin,omitempty"`
	SystemInstanceIpsecQAT               int                            `json:"QAT,omitempty"`
	SystemInstanceIpsecUUID              string                         `json:"uuid,omitempty"`
}

type SystemInstanceJobOffload struct {
	SystemInstanceJobOffloadSamplingEnableCounters1 []SystemInstanceJobOffloadSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceJobOffloadUUID                    string                                   `json:"uuid,omitempty"`
}

type SystemInstanceLinkCapability struct {
	SystemInstanceLinkCapabilityEnable int    `json:"enable,omitempty"`
	SystemInstanceLinkCapabilityUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceLinkMonitor struct {
	SystemInstanceLinkMonitorDisable int `json:"disable,omitempty"`
	SystemInstanceLinkMonitorEnable  int `json:"enable,omitempty"`
}

type SystemInstanceLro struct {
	SystemInstanceLroDisable int `json:"disable,omitempty"`
	SystemInstanceLroEnable  int `json:"enable,omitempty"`
}

type SystemInstanceManagementInterfaceMode struct {
	SystemInstanceManagementInterfaceModeDedicated    int `json:"dedicated,omitempty"`
	SystemInstanceManagementInterfaceModeNonDedicated int `json:"non-dedicated,omitempty"`
}

type SystemInstanceMemory struct {
	SystemInstanceMemorySamplingEnableCounters1 []SystemInstanceMemorySamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceMemoryUUID                    string                               `json:"uuid,omitempty"`
}

type SystemInstanceMfaAuth struct {
	SystemInstanceMfaAuthSecondFactor string `json:"second-factor,omitempty"`
	SystemInstanceMfaAuthUsername     string `json:"username,omitempty"`
}

type SystemInstanceMfaCertStore struct {
	SystemInstanceMfaCertStoreCertHost      string `json:"cert-host,omitempty"`
	SystemInstanceMfaCertStoreCertStorePath string `json:"cert-store-path,omitempty"`
	SystemInstanceMfaCertStorePasswdString  string `json:"passwd-string,omitempty"`
	SystemInstanceMfaCertStoreProtocol      string `json:"protocol,omitempty"`
	SystemInstanceMfaCertStoreUUID          string `json:"uuid,omitempty"`
	SystemInstanceMfaCertStoreUsername      string `json:"username,omitempty"`
}

type SystemInstanceMfaManagement struct {
	SystemInstanceMfaManagementEnable int    `json:"enable,omitempty"`
	SystemInstanceMfaManagementUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceMfaValidationType struct {
	SystemInstanceMfaValidationTypeCaCert string `json:"ca-cert,omitempty"`
	SystemInstanceMfaValidationTypeUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceMgmtPort struct {
	SystemInstanceMgmtPortMacAddress string `json:"mac-address,omitempty"`
	SystemInstanceMgmtPortPciAddress string `json:"pci-address,omitempty"`
	SystemInstanceMgmtPortPortIndex  int    `json:"port-index,omitempty"`
}

type SystemInstanceModifyPort struct {
	SystemInstanceModifyPortPortIndex  int `json:"port-index,omitempty"`
	SystemInstanceModifyPortPortNumber int `json:"port-number,omitempty"`
}

type SystemInstanceMonTemplate struct {
	SystemInstanceMonTemplateLinkBlockAsDownEnable   SystemInstanceMonTemplateLinkBlockAsDown   `json:"link-block-as-down,omitempty"`
	SystemInstanceMonTemplateLinkDownOnRestartEnable SystemInstanceMonTemplateLinkDownOnRestart `json:"link-down-on-restart,omitempty"`
	SystemInstanceMonTemplateMonitorListID           []SystemInstanceMonTemplateMonitorList     `json:"monitor-list,omitempty"`
}

type SystemInstanceMultiQueueSupport struct {
	SystemInstanceMultiQueueSupportEnable int `json:"enable,omitempty"`
}

type SystemInstanceNdiscRa struct {
	SystemInstanceNdiscRaSamplingEnableCounters1 []SystemInstanceNdiscRaSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceNdiscRaUUID                    string                                `json:"uuid,omitempty"`
}

type SystemInstancePasswordPolicy struct {
	SystemInstancePasswordPolicyAging      string `json:"aging,omitempty"`
	SystemInstancePasswordPolicyComplexity string `json:"complexity,omitempty"`
	SystemInstancePasswordPolicyHistory    string `json:"history,omitempty"`
	SystemInstancePasswordPolicyMinPswdLen int    `json:"min-pswd-len,omitempty"`
	SystemInstancePasswordPolicyUUID       string `json:"uuid,omitempty"`
}

type SystemInstancePathList struct {
	SystemInstancePathListIfpairEthEnd     int    `json:"ifpair-eth-end,omitempty"`
	SystemInstancePathListIfpairEthStart   int    `json:"ifpair-eth-start,omitempty"`
	SystemInstancePathListIfpairTrunkEnd   int    `json:"ifpair-trunk-end,omitempty"`
	SystemInstancePathListIfpairTrunkStart int    `json:"ifpair-trunk-start,omitempty"`
	SystemInstancePathListL2HmAttach       string `json:"l2hm-attach,omitempty"`
	SystemInstancePathListL2HmPathName     string `json:"l2hm-path-name,omitempty"`
	SystemInstancePathListL2HmSetupTestAPI int    `json:"l2hm-setup-test-api,omitempty"`
	SystemInstancePathListL2HmVlan         int    `json:"l2hm-vlan,omitempty"`
	SystemInstancePathListUUID             string `json:"uuid,omitempty"`
	SystemInstancePathListUserTag          string `json:"user-tag,omitempty"`
}

type SystemInstancePerVlanLimit struct {
	SystemInstancePerVlanLimitBcast        int    `json:"bcast,omitempty"`
	SystemInstancePerVlanLimitIpmcast      int    `json:"ipmcast,omitempty"`
	SystemInstancePerVlanLimitMcast        int    `json:"mcast,omitempty"`
	SystemInstancePerVlanLimitUUID         string `json:"uuid,omitempty"`
	SystemInstancePerVlanLimitUnknownUcast int    `json:"unknown-ucast,omitempty"`
}

type SystemInstancePlatformtype struct {
	SystemInstancePlatformtypeUUID string `json:"uuid,omitempty"`
}

type SystemInstancePortCount struct {
	SystemInstancePortCountPortCountAlg     int    `json:"port-count-alg,omitempty"`
	SystemInstancePortCountPortCountHm      int    `json:"port-count-hm,omitempty"`
	SystemInstancePortCountPortCountKernel  int    `json:"port-count-kernel,omitempty"`
	SystemInstancePortCountPortCountLogging int    `json:"port-count-logging,omitempty"`
	SystemInstancePortCountUUID             string `json:"uuid,omitempty"`
}

type SystemInstancePortInfo struct {
	SystemInstancePortInfoUUID string `json:"uuid,omitempty"`
}

type SystemInstancePortList struct {
	SystemInstancePortListUUID string `json:"uuid,omitempty"`
}

type SystemInstancePorts struct {
	SystemInstancePortsLinkDetectionInterval int    `json:"link-detection-interval,omitempty"`
	SystemInstancePortsUUID                  string `json:"uuid,omitempty"`
}

type SystemInstanceProbeNetworkDevices struct {
}

type SystemInstancePsuInfo struct {
	SystemInstancePsuInfoUUID string `json:"uuid,omitempty"`
}

type SystemInstanceQInQ struct {
	SystemInstanceQInQEnableAllPorts int    `json:"enable-all-ports,omitempty"`
	SystemInstanceQInQInnerTpid      string `json:"inner-tpid,omitempty"`
	SystemInstanceQInQOuterTpid      string `json:"outer-tpid,omitempty"`
	SystemInstanceQInQUUID           string `json:"uuid,omitempty"`
}

type SystemInstanceQueuingBuffer struct {
	SystemInstanceQueuingBufferEnable int    `json:"enable,omitempty"`
	SystemInstanceQueuingBufferUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceRadius struct {
	SystemInstanceRadiusServerListenPort SystemInstanceRadiusServer `json:"server,omitempty"`
}

type SystemInstanceReboot struct {
	SystemInstanceRebootUUID string `json:"uuid,omitempty"`
}

type SystemInstanceResourceAccounting struct {
	SystemInstanceResourceAccountingTemplateListName []SystemInstanceResourceAccountingTemplateList `json:"template-list,omitempty"`
	SystemInstanceResourceAccountingUUID             string                                         `json:"uuid,omitempty"`
}

type SystemInstanceResourceUsage struct {
	SystemInstanceResourceUsageAflexTableEntryCount           int                                   `json:"aflex-table-entry-count,omitempty"`
	SystemInstanceResourceUsageAuthPortalHTMLFileSize         int                                   `json:"auth-portal-html-file-size,omitempty"`
	SystemInstanceResourceUsageAuthPortalImageFileSize        int                                   `json:"auth-portal-image-file-size,omitempty"`
	SystemInstanceResourceUsageAuthSessionCount               int                                   `json:"auth-session-count,omitempty"`
	SystemInstanceResourceUsageAuthzPolicyNumber              int                                   `json:"authz-policy-number,omitempty"`
	SystemInstanceResourceUsageClassListAcEntryCount          int                                   `json:"class-list-ac-entry-count,omitempty"`
	SystemInstanceResourceUsageClassListEntryCount            int                                   `json:"class-list-entry-count,omitempty"`
	SystemInstanceResourceUsageClassListIpv6AddrCount         int                                   `json:"class-list-ipv6-addr-count,omitempty"`
	SystemInstanceResourceUsageIpsecSaNumber                  int                                   `json:"ipsec-sa-number,omitempty"`
	SystemInstanceResourceUsageL4SessionCount                 int                                   `json:"l4-session-count,omitempty"`
	SystemInstanceResourceUsageMaxAflexAuthzCollectionNumber  int                                   `json:"max-aflex-authz-collection-number,omitempty"`
	SystemInstanceResourceUsageMaxAflexFileSize               int                                   `json:"max-aflex-file-size,omitempty"`
	SystemInstanceResourceUsageNatPoolAddrCount               int                                   `json:"nat-pool-addr-count,omitempty"`
	SystemInstanceResourceUsageRAMCacheMemoryLimit            int                                   `json:"ram-cache-memory-limit,omitempty"`
	SystemInstanceResourceUsageRadiusTableSize                int                                   `json:"radius-table-size,omitempty"`
	SystemInstanceResourceUsageSslContextMemory               int                                   `json:"ssl-context-memory,omitempty"`
	SystemInstanceResourceUsageSslDmaMemory                   int                                   `json:"ssl-dma-memory,omitempty"`
	SystemInstanceResourceUsageUUID                           string                                `json:"uuid,omitempty"`
	SystemInstanceResourceUsageVisibilityMonitoredEntityCount SystemInstanceResourceUsageVisibility `json:"visibility,omitempty"`
	SystemInstanceResourceUsageWafTemplateCount               int                                   `json:"waf-template-count,omitempty"`
}

type SystemInstanceSession struct {
	SystemInstanceSessionSamplingEnableCounters1 []SystemInstanceSessionSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceSessionUUID                    string                                `json:"uuid,omitempty"`
}

type SystemInstanceSessionReclaimLimit struct {
	SystemInstanceSessionReclaimLimitNscanLimit int    `json:"nscan-limit,omitempty"`
	SystemInstanceSessionReclaimLimitScanFreq   int    `json:"scan-freq,omitempty"`
	SystemInstanceSessionReclaimLimitUUID       string `json:"uuid,omitempty"`
}

type SystemInstanceSetRxtxDescSize struct {
	SystemInstanceSetRxtxDescSizePortIndex int `json:"port-index,omitempty"`
	SystemInstanceSetRxtxDescSizeRxdSize   int `json:"rxd-size,omitempty"`
	SystemInstanceSetRxtxDescSizeTxdSize   int `json:"txd-size,omitempty"`
}

type SystemInstanceSetRxtxQueue struct {
	SystemInstanceSetRxtxQueuePortIndex int `json:"port-index,omitempty"`
	SystemInstanceSetRxtxQueueRxqSize   int `json:"rxq-size,omitempty"`
	SystemInstanceSetRxtxQueueTxqSize   int `json:"txq-size,omitempty"`
}

type SystemInstanceSetTCPSynPerSec struct {
	SystemInstanceSetTCPSynPerSecTCPSynValue int    `json:"tcp-syn-value,omitempty"`
	SystemInstanceSetTCPSynPerSecUUID        string `json:"uuid,omitempty"`
}

type SystemInstanceSharedPollMode struct {
	SystemInstanceSharedPollModeDisable int `json:"disable,omitempty"`
	SystemInstanceSharedPollModeEnable  int `json:"enable,omitempty"`
}

type SystemInstanceShellPrivileges struct {
	SystemInstanceShellPrivilegesEnableShellPrivileges int    `json:"enable-shell-privileges,omitempty"`
	SystemInstanceShellPrivilegesUUID                  string `json:"uuid,omitempty"`
}

type SystemInstanceShmLogging struct {
	SystemInstanceShmLoggingEnable int    `json:"enable,omitempty"`
	SystemInstanceShmLoggingUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceShutdown struct {
	SystemInstanceShutdownUUID string `json:"uuid,omitempty"`
}

type SystemInstanceSpeProfile struct {
	SystemInstanceSpeProfileAction string `json:"action,omitempty"`
}

type SystemInstanceSpeStatus struct {
	SystemInstanceSpeStatusUUID string `json:"uuid,omitempty"`
}

type SystemInstanceSslReqQ struct {
	SystemInstanceSslReqQSamplingEnableCounters1 []SystemInstanceSslReqQSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceSslReqQUUID                    string                                `json:"uuid,omitempty"`
}

type SystemInstanceSslScv struct {
	SystemInstanceSslScvEnable int    `json:"enable,omitempty"`
	SystemInstanceSslScvUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceSslScvVerifyCrlSign struct {
	SystemInstanceSslScvVerifyCrlSignEnable int    `json:"enable,omitempty"`
	SystemInstanceSslScvVerifyCrlSignUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceSslScvVerifyHost struct {
	SystemInstanceSslScvVerifyHostDisable int    `json:"disable,omitempty"`
	SystemInstanceSslScvVerifyHostUUID    string `json:"uuid,omitempty"`
}

type SystemInstanceSslSetCompatibleCipher struct {
	SystemInstanceSslSetCompatibleCipherDisable int    `json:"disable,omitempty"`
	SystemInstanceSslSetCompatibleCipherUUID    string `json:"uuid,omitempty"`
}

type SystemInstanceSslStatus struct {
	SystemInstanceSslStatusUUID string `json:"uuid,omitempty"`
}

type SystemInstanceSyslogTimeMsec struct {
	SystemInstanceSyslogTimeMsecEnableFlag int `json:"enable-flag,omitempty"`
}

type SystemInstanceTCP struct {
	SystemInstanceTCPRateLimitResetUnknownConnPktRateForResetUnknownConn SystemInstanceTCPRateLimitResetUnknownConn `json:"rate-limit-reset-unknown-conn,omitempty"`
	SystemInstanceTCPSamplingEnableCounters1                             []SystemInstanceTCPSamplingEnable          `json:"sampling-enable,omitempty"`
	SystemInstanceTCPUUID                                                string                                     `json:"uuid,omitempty"`
}

type SystemInstanceTCPStats struct {
	SystemInstanceTCPStatsSamplingEnableCounters1 []SystemInstanceTCPStatsSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceTCPStatsUUID                    string                                 `json:"uuid,omitempty"`
}

type SystemInstanceTCPSynPerSec struct {
	SystemInstanceTCPSynPerSecUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTLS13Mgmt struct {
	SystemInstanceTLS13MgmtEnable int    `json:"enable,omitempty"`
	SystemInstanceTLS13MgmtUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceTableIntegrity struct {
	SystemInstanceTableIntegrityAuditAction             string                                       `json:"audit-action,omitempty"`
	SystemInstanceTableIntegrityAutoSyncAction          string                                       `json:"auto-sync-action,omitempty"`
	SystemInstanceTableIntegritySamplingEnableCounters1 []SystemInstanceTableIntegritySamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceTableIntegrityTable                   string                                       `json:"table,omitempty"`
	SystemInstanceTableIntegrityUUID                    string                                       `json:"uuid,omitempty"`
}

type SystemInstanceTelemetryLog struct {
	SystemInstanceTelemetryLogDeviceStatusUUID     SystemInstanceTelemetryLogDeviceStatus     `json:"device-status,omitempty"`
	SystemInstanceTelemetryLogEnvironmentUUID      SystemInstanceTelemetryLogEnvironment      `json:"environment,omitempty"`
	SystemInstanceTelemetryLogPartitionMetricsUUID SystemInstanceTelemetryLogPartitionMetrics `json:"partition-metrics,omitempty"`
	SystemInstanceTelemetryLogTopKAppSvcListUUID   SystemInstanceTelemetryLogTopKAppSvcList   `json:"top-k-app-svc-list,omitempty"`
	SystemInstanceTelemetryLogTopKSourceListUUID   SystemInstanceTelemetryLogTopKSourceList   `json:"top-k-source-list,omitempty"`
}

type SystemInstanceTemplate struct {
	SystemInstanceTemplateTemplatePolicy string `json:"template-policy,omitempty"`
	SystemInstanceTemplateUUID           string `json:"uuid,omitempty"`
}

type SystemInstanceTemplateBind struct {
	SystemInstanceTemplateBindMonitorListID []SystemInstanceTemplateBindMonitorList `json:"monitor-list,omitempty"`
}

type SystemInstanceThroughput struct {
	SystemInstanceThroughputSamplingEnableCounters1 []SystemInstanceThroughputSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceThroughputUUID                    string                                   `json:"uuid,omitempty"`
}

type SystemInstanceTimeoutValue struct {
	SystemInstanceTimeoutValueFtp   int    `json:"ftp,omitempty"`
	SystemInstanceTimeoutValueHTTP  int    `json:"http,omitempty"`
	SystemInstanceTimeoutValueHTTPS int    `json:"https,omitempty"`
	SystemInstanceTimeoutValueScp   int    `json:"scp,omitempty"`
	SystemInstanceTimeoutValueSftp  int    `json:"sftp,omitempty"`
	SystemInstanceTimeoutValueTftp  int    `json:"tftp,omitempty"`
	SystemInstanceTimeoutValueUUID  string `json:"uuid,omitempty"`
}

type SystemInstanceTrunk struct {
	SystemInstanceTrunkLoadBalanceUseL3 SystemInstanceTrunkLoadBalance `json:"load-balance,omitempty"`
}

type SystemInstanceTrunkHwHash struct {
	SystemInstanceTrunkHwHashMode int    `json:"mode,omitempty"`
	SystemInstanceTrunkHwHashUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTrunkXauiHwHash struct {
	SystemInstanceTrunkXauiHwHashMode int    `json:"mode,omitempty"`
	SystemInstanceTrunkXauiHwHashUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTso struct {
	SystemInstanceTsoDisable int `json:"disable,omitempty"`
	SystemInstanceTsoEnable  int `json:"enable,omitempty"`
}

type SystemInstanceUpgradeStatus struct {
	SystemInstanceUpgradeStatusUUID string `json:"uuid,omitempty"`
}

type SystemInstanceVeMacScheme struct {
	SystemInstanceVeMacSchemeUUID           string `json:"uuid,omitempty"`
	SystemInstanceVeMacSchemeVeMacSchemeVal string `json:"ve-mac-scheme-val,omitempty"`
}

type SystemInstanceAppPerformanceSamplingEnable struct {
	SystemInstanceAppPerformanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceBandwidthSamplingEnable struct {
	SystemInstanceBandwidthSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceBfdSamplingEnable struct {
	SystemInstanceBfdSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceCPULoadSharingCPUUsage struct {
	SystemInstanceCPULoadSharingCPUUsageHigh int `json:"high,omitempty"`
	SystemInstanceCPULoadSharingCPUUsageLow  int `json:"low,omitempty"`
}

type SystemInstanceCPULoadSharingPacketsPerSecond struct {
	SystemInstanceCPULoadSharingPacketsPerSecondMin int `json:"min,omitempty"`
}

type SystemInstanceDNSSamplingEnable struct {
	SystemInstanceDNSSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceDNSCacheSamplingEnable struct {
	SystemInstanceDNSCacheSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceDpdkStatsSamplingEnable struct {
	SystemInstanceDpdkStatsSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceFpgaDropSamplingEnable struct {
	SystemInstanceFpgaDropSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceGeoLocationEntryList struct {
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIPAddress []SystemInstanceGeoLocationEntryListGeoLocnMultipleAddresses `json:"geo-locn-multiple-addresses,omitempty"`
	SystemInstanceGeoLocationEntryListGeoLocnObjName                         string                                                       `json:"geo-locn-obj-name,omitempty"`
	SystemInstanceGeoLocationEntryListUUID                                   string                                                       `json:"uuid,omitempty"`
	SystemInstanceGeoLocationEntryListUserTag                                string                                                       `json:"user-tag,omitempty"`
}

type SystemInstanceGeoLocationGeolocLoadFileList struct {
	SystemInstanceGeoLocationGeolocLoadFileListGeoLocationLoadFilename string `json:"geo-location-load-filename,omitempty"`
	SystemInstanceGeoLocationGeolocLoadFileListTemplateName            string `json:"template-name,omitempty"`
}

type SystemInstanceGeolocSamplingEnable struct {
	SystemInstanceGeolocSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceGeolocListListExcludeGeolocNameList struct {
	SystemInstanceGeolocListListExcludeGeolocNameListExcludeGeolocNameVal string `json:"exclude-geoloc-name-val,omitempty"`
}

type SystemInstanceGeolocListListIncludeGeolocNameList struct {
	SystemInstanceGeolocListListIncludeGeolocNameListIncludeGeolocNameVal string `json:"include-geoloc-name-val,omitempty"`
}

type SystemInstanceGeolocListListSamplingEnable struct {
	SystemInstanceGeolocListListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceGeolocNameHelperSamplingEnable struct {
	SystemInstanceGeolocNameHelperSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceGeolocationFileErrorInfo struct {
	SystemInstanceGeolocationFileErrorInfoUUID string `json:"uuid,omitempty"`
}

type SystemInstanceHardwareForwardSamplingEnable struct {
	SystemInstanceHardwareForwardSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceHardwareForwardSlb struct {
	SystemInstanceHardwareForwardSlbSamplingEnableCounters1 []SystemInstanceHardwareForwardSlbSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceHardwareForwardSlbUUID                    string                                           `json:"uuid,omitempty"`
}

type SystemInstanceIP6StatsSamplingEnable struct {
	SystemInstanceIP6StatsSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIPStatsSamplingEnable struct {
	SystemInstanceIPStatsSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIPThreatListIpv4DestList struct {
	SystemInstanceIPThreatListIpv4DestListClassListCfgClassList []SystemInstanceIPThreatListIpv4DestListClassListCfg `json:"class-list-cfg,omitempty"`
	SystemInstanceIPThreatListIpv4DestListUUID                  string                                               `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatListIpv4InternetHostList struct {
	SystemInstanceIPThreatListIpv4InternetHostListClassListCfgClassList []SystemInstanceIPThreatListIpv4InternetHostListClassListCfg `json:"class-list-cfg,omitempty"`
	SystemInstanceIPThreatListIpv4InternetHostListUUID                  string                                                       `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatListIpv4SourceList struct {
	SystemInstanceIPThreatListIpv4SourceListClassListCfgClassList []SystemInstanceIPThreatListIpv4SourceListClassListCfg `json:"class-list-cfg,omitempty"`
	SystemInstanceIPThreatListIpv4SourceListUUID                  string                                                 `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatListIpv6DestList struct {
	SystemInstanceIPThreatListIpv6DestListClassListCfgClassList []SystemInstanceIPThreatListIpv6DestListClassListCfg `json:"class-list-cfg,omitempty"`
	SystemInstanceIPThreatListIpv6DestListUUID                  string                                               `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatListIpv6InternetHostList struct {
	SystemInstanceIPThreatListIpv6InternetHostListClassListCfgClassList []SystemInstanceIPThreatListIpv6InternetHostListClassListCfg `json:"class-list-cfg,omitempty"`
	SystemInstanceIPThreatListIpv6InternetHostListUUID                  string                                                       `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatListIpv6SourceList struct {
	SystemInstanceIPThreatListIpv6SourceListClassListCfgClassList []SystemInstanceIPThreatListIpv6SourceListClassListCfg `json:"class-list-cfg,omitempty"`
	SystemInstanceIPThreatListIpv6SourceListUUID                  string                                                 `json:"uuid,omitempty"`
}

type SystemInstanceIPThreatListSamplingEnable struct {
	SystemInstanceIPThreatListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIcmpSamplingEnable struct {
	SystemInstanceIcmpSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIcmp6SamplingEnable struct {
	SystemInstanceIcmp6SamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIcmpRateSamplingEnable struct {
	SystemInstanceIcmpRateSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIpmiIP struct {
	SystemInstanceIpmiIPDefaultGateway string `json:"default-gateway,omitempty"`
	SystemInstanceIpmiIPIpv4Address    string `json:"ipv4-address,omitempty"`
	SystemInstanceIpmiIPIpv4Netmask    string `json:"ipv4-netmask,omitempty"`
}

type SystemInstanceIpmiIpsrc struct {
	SystemInstanceIpmiIpsrcDhcp   int `json:"dhcp,omitempty"`
	SystemInstanceIpmiIpsrcStatic int `json:"static,omitempty"`
}

type SystemInstanceIpmiTool struct {
	SystemInstanceIpmiToolCmd string `json:"cmd,omitempty"`
}

type SystemInstanceIpmiUser struct {
	SystemInstanceIpmiUserAdd           string `json:"add,omitempty"`
	SystemInstanceIpmiUserAdministrator int    `json:"administrator,omitempty"`
	SystemInstanceIpmiUserCallback      int    `json:"callback,omitempty"`
	SystemInstanceIpmiUserDisable       string `json:"disable,omitempty"`
	SystemInstanceIpmiUserNewname       string `json:"newname,omitempty"`
	SystemInstanceIpmiUserNewpass       string `json:"newpass,omitempty"`
	SystemInstanceIpmiUserOperator      int    `json:"operator,omitempty"`
	SystemInstanceIpmiUserPassword      string `json:"password,omitempty"`
	SystemInstanceIpmiUserPrivilege     string `json:"privilege,omitempty"`
	SystemInstanceIpmiUserSetname       string `json:"setname,omitempty"`
	SystemInstanceIpmiUserSetpass       string `json:"setpass,omitempty"`
	SystemInstanceIpmiUserUser          int    `json:"user,omitempty"`
}

type SystemInstanceIpsecFpgaDecrypt struct {
	SystemInstanceIpsecFpgaDecryptAction string `json:"action,omitempty"`
}

type SystemInstanceJobOffloadSamplingEnable struct {
	SystemInstanceJobOffloadSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceMemorySamplingEnable struct {
	SystemInstanceMemorySamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceMonTemplateLinkBlockAsDown struct {
	SystemInstanceMonTemplateLinkBlockAsDownEnable int    `json:"enable,omitempty"`
	SystemInstanceMonTemplateLinkBlockAsDownUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceMonTemplateLinkDownOnRestart struct {
	SystemInstanceMonTemplateLinkDownOnRestartEnable int    `json:"enable,omitempty"`
	SystemInstanceMonTemplateLinkDownOnRestartUUID   string `json:"uuid,omitempty"`
}

type SystemInstanceMonTemplateMonitorList struct {
	SystemInstanceMonTemplateMonitorListClearCfgSessions             []SystemInstanceMonTemplateMonitorListClearCfg       `json:"clear-cfg,omitempty"`
	SystemInstanceMonTemplateMonitorListID                           int                                                  `json:"id,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDisableCfgDiseth         []SystemInstanceMonTemplateMonitorListLinkDisableCfg `json:"link-disable-cfg,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet1 []SystemInstanceMonTemplateMonitorListLinkDownCfg    `json:"link-down-cfg,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaeth          []SystemInstanceMonTemplateMonitorListLinkEnableCfg  `json:"link-enable-cfg,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet1     []SystemInstanceMonTemplateMonitorListLinkUpCfg      `json:"link-up-cfg,omitempty"`
	SystemInstanceMonTemplateMonitorListMonitorRelation              string                                               `json:"monitor-relation,omitempty"`
	SystemInstanceMonTemplateMonitorListUUID                         string                                               `json:"uuid,omitempty"`
	SystemInstanceMonTemplateMonitorListUserTag                      string                                               `json:"user-tag,omitempty"`
}

type SystemInstanceNdiscRaSamplingEnable struct {
	SystemInstanceNdiscRaSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceRadiusServer struct {
	SystemInstanceRadiusServerAccountingInterimUpdate string                                     `json:"accounting-interim-update,omitempty"`
	SystemInstanceRadiusServerAccountingOn            string                                     `json:"accounting-on,omitempty"`
	SystemInstanceRadiusServerAccountingStart         string                                     `json:"accounting-start,omitempty"`
	SystemInstanceRadiusServerAccountingStop          string                                     `json:"accounting-stop,omitempty"`
	SystemInstanceRadiusServerAttributeAttributeValue []SystemInstanceRadiusServerAttribute      `json:"attribute,omitempty"`
	SystemInstanceRadiusServerAttributeName           string                                     `json:"attribute-name,omitempty"`
	SystemInstanceRadiusServerCustomAttributeName     string                                     `json:"custom-attribute-name,omitempty"`
	SystemInstanceRadiusServerDisableReply            int                                        `json:"disable-reply,omitempty"`
	SystemInstanceRadiusServerListenPort              int                                        `json:"listen-port,omitempty"`
	SystemInstanceRadiusServerRemoteIPList            SystemInstanceRadiusServerRemote           `json:"remote,omitempty"`
	SystemInstanceRadiusServerSamplingEnableCounters1 []SystemInstanceRadiusServerSamplingEnable `json:"sampling-enable,omitempty"`
	SystemInstanceRadiusServerSecret                  int                                        `json:"secret,omitempty"`
	SystemInstanceRadiusServerSecretString            string                                     `json:"secret-string,omitempty"`
	SystemInstanceRadiusServerUUID                    string                                     `json:"uuid,omitempty"`
	SystemInstanceRadiusServerVrid                    int                                        `json:"vrid,omitempty"`
}

type SystemInstanceResourceAccountingTemplateList struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfg          SystemInstanceResourceAccountingTemplateListAppResources     `json:"app-resources,omitempty"`
	SystemInstanceResourceAccountingTemplateListName                               string                                                       `json:"name,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg SystemInstanceResourceAccountingTemplateListNetworkResources `json:"network-resources,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfg          SystemInstanceResourceAccountingTemplateListSystemResources  `json:"system-resources,omitempty"`
	SystemInstanceResourceAccountingTemplateListUUID                               string                                                       `json:"uuid,omitempty"`
	SystemInstanceResourceAccountingTemplateListUserTag                            string                                                       `json:"user-tag,omitempty"`
}

type SystemInstanceResourceUsageVisibility struct {
	SystemInstanceResourceUsageVisibilityMonitoredEntityCount int    `json:"monitored-entity-count,omitempty"`
	SystemInstanceResourceUsageVisibilityUUID                 string `json:"uuid,omitempty"`
}

type SystemInstanceSessionSamplingEnable struct {
	SystemInstanceSessionSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceSslReqQSamplingEnable struct {
	SystemInstanceSslReqQSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceTCPRateLimitResetUnknownConn struct {
	SystemInstanceTCPRateLimitResetUnknownConnLogForResetUnknownConn     int    `json:"log-for-reset-unknown-conn,omitempty"`
	SystemInstanceTCPRateLimitResetUnknownConnPktRateForResetUnknownConn int    `json:"pkt-rate-for-reset-unknown-conn,omitempty"`
	SystemInstanceTCPRateLimitResetUnknownConnUUID                       string `json:"uuid,omitempty"`
}

type SystemInstanceTCPSamplingEnable struct {
	SystemInstanceTCPSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceTCPStatsSamplingEnable struct {
	SystemInstanceTCPStatsSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceTableIntegritySamplingEnable struct {
	SystemInstanceTableIntegritySamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceTelemetryLogDeviceStatus struct {
	SystemInstanceTelemetryLogDeviceStatusUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTelemetryLogEnvironment struct {
	SystemInstanceTelemetryLogEnvironmentUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTelemetryLogPartitionMetrics struct {
	SystemInstanceTelemetryLogPartitionMetricsUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTelemetryLogTopKAppSvcList struct {
	SystemInstanceTelemetryLogTopKAppSvcListUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTelemetryLogTopKSourceList struct {
	SystemInstanceTelemetryLogTopKSourceListUUID string `json:"uuid,omitempty"`
}

type SystemInstanceTemplateBindMonitorList struct {
	SystemInstanceTemplateBindMonitorListClearCfgSessions             []SystemInstanceTemplateBindMonitorListClearCfg       `json:"clear-cfg,omitempty"`
	SystemInstanceTemplateBindMonitorListID                           int                                                   `json:"id,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDisableCfgDiseth         []SystemInstanceTemplateBindMonitorListLinkDisableCfg `json:"link-disable-cfg,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet1 []SystemInstanceTemplateBindMonitorListLinkDownCfg    `json:"link-down-cfg,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaeth          []SystemInstanceTemplateBindMonitorListLinkEnableCfg  `json:"link-enable-cfg,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet1     []SystemInstanceTemplateBindMonitorListLinkUpCfg      `json:"link-up-cfg,omitempty"`
	SystemInstanceTemplateBindMonitorListMonitorRelation              string                                                `json:"monitor-relation,omitempty"`
	SystemInstanceTemplateBindMonitorListUUID                         string                                                `json:"uuid,omitempty"`
	SystemInstanceTemplateBindMonitorListUserTag                      string                                                `json:"user-tag,omitempty"`
}

type SystemInstanceThroughputSamplingEnable struct {
	SystemInstanceThroughputSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceTrunkLoadBalance struct {
	SystemInstanceTrunkLoadBalanceUUID  string `json:"uuid,omitempty"`
	SystemInstanceTrunkLoadBalanceUseL3 int    `json:"use-l3,omitempty"`
	SystemInstanceTrunkLoadBalanceUseL4 int    `json:"use-l4,omitempty"`
}

type SystemInstanceGeoLocationEntryListGeoLocnMultipleAddresses struct {
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIPAddress   string `json:"first-ip-address,omitempty"`
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIpv6Address string `json:"first-ipv6-address,omitempty"`
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesGeolIpv4Mask     string `json:"geol-ipv4-mask,omitempty"`
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesGeolIpv6Mask     int    `json:"geol-ipv6-mask,omitempty"`
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesIPAddr2          string `json:"ip-addr2,omitempty"`
	SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesIpv6Addr2        string `json:"ipv6-addr2,omitempty"`
}

type SystemInstanceHardwareForwardSlbSamplingEnable struct {
	SystemInstanceHardwareForwardSlbSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceIPThreatListIpv4DestListClassListCfg struct {
	SystemInstanceIPThreatListIpv4DestListClassListCfgClassList          string `json:"class-list,omitempty"`
	SystemInstanceIPThreatListIpv4DestListClassListCfgIPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemInstanceIPThreatListIpv4InternetHostListClassListCfg struct {
	SystemInstanceIPThreatListIpv4InternetHostListClassListCfgClassList          string `json:"class-list,omitempty"`
	SystemInstanceIPThreatListIpv4InternetHostListClassListCfgIPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemInstanceIPThreatListIpv4SourceListClassListCfg struct {
	SystemInstanceIPThreatListIpv4SourceListClassListCfgClassList          string `json:"class-list,omitempty"`
	SystemInstanceIPThreatListIpv4SourceListClassListCfgIPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemInstanceIPThreatListIpv6DestListClassListCfg struct {
	SystemInstanceIPThreatListIpv6DestListClassListCfgClassList          string `json:"class-list,omitempty"`
	SystemInstanceIPThreatListIpv6DestListClassListCfgIPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemInstanceIPThreatListIpv6InternetHostListClassListCfg struct {
	SystemInstanceIPThreatListIpv6InternetHostListClassListCfgClassList          string `json:"class-list,omitempty"`
	SystemInstanceIPThreatListIpv6InternetHostListClassListCfgIPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemInstanceIPThreatListIpv6SourceListClassListCfg struct {
	SystemInstanceIPThreatListIpv6SourceListClassListCfgClassList          string `json:"class-list,omitempty"`
	SystemInstanceIPThreatListIpv6SourceListClassListCfgIPThreatActionTmpl int    `json:"ip-threat-action-tmpl,omitempty"`
}

type SystemInstanceMonTemplateMonitorListClearCfg struct {
	SystemInstanceMonTemplateMonitorListClearCfgClearAllSequence int    `json:"clear-all-sequence,omitempty"`
	SystemInstanceMonTemplateMonitorListClearCfgClearSequence    int    `json:"clear-sequence,omitempty"`
	SystemInstanceMonTemplateMonitorListClearCfgSessions         string `json:"sessions,omitempty"`
}

type SystemInstanceMonTemplateMonitorListLinkDisableCfg struct {
	SystemInstanceMonTemplateMonitorListLinkDisableCfgDisSequence int `json:"dis-sequence,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDisableCfgDiseth      int `json:"diseth,omitempty"`
}

type SystemInstanceMonTemplateMonitorListLinkDownCfg struct {
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkDownSequence1 int `json:"link-down-sequence1,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkDownSequence2 int `json:"link-down-sequence2,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkDownSequence3 int `json:"link-down-sequence3,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet1 int `json:"linkdown-ethernet1,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet2 int `json:"linkdown-ethernet2,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet3 int `json:"linkdown-ethernet3,omitempty"`
}

type SystemInstanceMonTemplateMonitorListLinkEnableCfg struct {
	SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaSequence int `json:"ena-sequence,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaeth      int `json:"enaeth,omitempty"`
}

type SystemInstanceMonTemplateMonitorListLinkUpCfg struct {
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkUpSequence1 int `json:"link-up-sequence1,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkUpSequence2 int `json:"link-up-sequence2,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkUpSequence3 int `json:"link-up-sequence3,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet1 int `json:"linkup-ethernet1,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet2 int `json:"linkup-ethernet2,omitempty"`
	SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet3 int `json:"linkup-ethernet3,omitempty"`
}

type SystemInstanceRadiusServerAttribute struct {
	SystemInstanceRadiusServerAttributeAttributeValue string `json:"attribute-value,omitempty"`
	SystemInstanceRadiusServerAttributeCustomNumber   int    `json:"custom-number,omitempty"`
	SystemInstanceRadiusServerAttributeCustomVendor   int    `json:"custom-vendor,omitempty"`
	SystemInstanceRadiusServerAttributeName           string `json:"name,omitempty"`
	SystemInstanceRadiusServerAttributeNumber         int    `json:"number,omitempty"`
	SystemInstanceRadiusServerAttributePrefixLength   string `json:"prefix-length,omitempty"`
	SystemInstanceRadiusServerAttributePrefixNumber   int    `json:"prefix-number,omitempty"`
	SystemInstanceRadiusServerAttributePrefixVendor   int    `json:"prefix-vendor,omitempty"`
	SystemInstanceRadiusServerAttributeValue          string `json:"value,omitempty"`
	SystemInstanceRadiusServerAttributeVendor         int    `json:"vendor,omitempty"`
}

type SystemInstanceRadiusServerRemote struct {
	SystemInstanceRadiusServerRemoteIPListIPListName []SystemInstanceRadiusServerRemoteIPList `json:"ip-list,omitempty"`
}

type SystemInstanceRadiusServerSamplingEnable struct {
	SystemInstanceRadiusServerSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResources struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfgCacheTemplateMax                 SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfg         `json:"cache-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfgClientSslTemplateMax         SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfg     `json:"client-ssl-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfgConnReuseTemplateMax         SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg     `json:"conn-reuse-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfgFastTCPTemplateMax             SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfg       `json:"fast-tcp-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfgFastUDPTemplateMax             SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfg       `json:"fast-udp-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfgFixTemplateMax                     SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfg           `json:"fix-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfgGslbDeviceMax                       SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfg            `json:"gslb-device-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfgGslbGeoLocationMax             SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg       `json:"gslb-geo-location-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfgGslbIPListMax                       SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfg            `json:"gslb-ip-list-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfgGslbPolicyMax                       SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfg            `json:"gslb-policy-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfgGslbServiceMax                     SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfg           `json:"gslb-service-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfgGslbServiceIPMax                 SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfg         `json:"gslb-service-ip-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfgGslbServicePortMax             SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfg       `json:"gslb-service-port-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfgGslbSiteMax                           SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfg              `json:"gslb-site-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfgGslbSvcGroupMax                   SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg          `json:"gslb-svc-group-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfgGslbTemplateMax                   SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfg          `json:"gslb-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfgGslbZoneMax                           SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfg              `json:"gslb-zone-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfgHTTPTemplateMax                   SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfg          `json:"http-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfgHealthMonitorMax                 SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfg         `json:"health-monitor-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfgLinkCostTemplateMax           SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg      `json:"link-cost-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfgPersistCookieTemplateMax SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg `json:"persist-cookie-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfgPersistSrcipTemplateMax   SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg  `json:"persist-srcip-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfgProxyTemplateMax                 SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfg         `json:"proxy-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfgRealPortMax                           SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfg              `json:"real-port-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfgRealServerMax                       SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfg            `json:"real-server-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfgServerSslTemplateMax         SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfg     `json:"server-ssl-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfgServiceGroupMax                   SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfg          `json:"service-group-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfgStreamTemplateMax               SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfg        `json:"stream-template-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesThreshold                                        int                                                                              `json:"threshold,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesUUID                                             string                                                                           `json:"uuid,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfgVirtualPortMax                     SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfg           `json:"virtual-port-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfgVirtualServerMax                 SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfg         `json:"virtual-server-cfg,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResources struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfgIpv4AclLineMax             SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg       `json:"ipv4-acl-line-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfgIpv6AclLineMax             SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg       `json:"ipv6-acl-line-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfgObjectGroupMax             SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfg       `json:"object-group-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfgObjectGroupClauseMax SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg `json:"object-group-clause-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfgStaticArpMax                 SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfg         `json:"static-arp-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfgStaticIpv4RouteMax     SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg   `json:"static-ipv4-route-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfgStaticIpv6RouteMax     SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg   `json:"static-ipv6-route-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfgStaticMacMax                 SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfg         `json:"static-mac-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfgStaticNeighborMax       SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg    `json:"static-neighbor-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesThreshold                                int                                                                              `json:"threshold,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesUUID                                     string                                                                           `json:"uuid,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResources struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfgBwLimitMax                               SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfg                `json:"bw-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfgConcurrentSessionLimitMax SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg `json:"concurrent-session-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfgFwcpsLimitMax                         SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg             `json:"fwcps-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfgL4CpsLimitMax                         SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfg             `json:"l4cps-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfgL4SessionLimitMax                 SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg         `json:"l4-session-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfgL7CpsLimitMax                         SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfg             `json:"l7cps-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfgNatcpsLimitMax                       SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg            `json:"natcps-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfgSslThroughputLimitMax         SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg     `json:"ssl-throughput-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfgSslcpsLimitMax                       SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg            `json:"sslcps-limit-cfg,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesThreshold                                          int                                                                                  `json:"threshold,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesUUID                                               string                                                                               `json:"uuid,omitempty"`
}

type SystemInstanceTemplateBindMonitorListClearCfg struct {
	SystemInstanceTemplateBindMonitorListClearCfgClearAllSequence int    `json:"clear-all-sequence,omitempty"`
	SystemInstanceTemplateBindMonitorListClearCfgClearSequence    int    `json:"clear-sequence,omitempty"`
	SystemInstanceTemplateBindMonitorListClearCfgSessions         string `json:"sessions,omitempty"`
}

type SystemInstanceTemplateBindMonitorListLinkDisableCfg struct {
	SystemInstanceTemplateBindMonitorListLinkDisableCfgDisSequence int `json:"dis-sequence,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDisableCfgDiseth      int `json:"diseth,omitempty"`
}

type SystemInstanceTemplateBindMonitorListLinkDownCfg struct {
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkDownSequence1 int `json:"link-down-sequence1,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkDownSequence2 int `json:"link-down-sequence2,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkDownSequence3 int `json:"link-down-sequence3,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet1 int `json:"linkdown-ethernet1,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet2 int `json:"linkdown-ethernet2,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet3 int `json:"linkdown-ethernet3,omitempty"`
}

type SystemInstanceTemplateBindMonitorListLinkEnableCfg struct {
	SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaSequence int `json:"ena-sequence,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaeth      int `json:"enaeth,omitempty"`
}

type SystemInstanceTemplateBindMonitorListLinkUpCfg struct {
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkUpSequence1 int `json:"link-up-sequence1,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkUpSequence2 int `json:"link-up-sequence2,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkUpSequence3 int `json:"link-up-sequence3,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet1 int `json:"linkup-ethernet1,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet2 int `json:"linkup-ethernet2,omitempty"`
	SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet3 int `json:"linkup-ethernet3,omitempty"`
}

type SystemInstanceRadiusServerRemoteIPList struct {
	SystemInstanceRadiusServerRemoteIPListIPListName         string `json:"ip-list-name,omitempty"`
	SystemInstanceRadiusServerRemoteIPListIPListSecret       int    `json:"ip-list-secret,omitempty"`
	SystemInstanceRadiusServerRemoteIPListIPListSecretString string `json:"ip-list-secret-string,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfgCacheTemplateMax          int `json:"cache-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfgCacheTemplateMinGuarantee int `json:"cache-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfgClientSslTemplateMax          int `json:"client-ssl-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfgClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfgConnReuseTemplateMax          int `json:"conn-reuse-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfgConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfgFastTCPTemplateMax          int `json:"fast-tcp-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfgFastTCPTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfgFastUDPTemplateMax          int `json:"fast-udp-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfgFastUDPTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfgFixTemplateMax          int `json:"fix-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfgFixTemplateMinGuarantee int `json:"fix-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfgGslbDeviceMax          int `json:"gslb-device-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfgGslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfgGslbGeoLocationMax          int `json:"gslb-geo-location-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfgGslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfgGslbIPListMax          int `json:"gslb-ip-list-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfgGslbIPListMinGuarantee int `json:"gslb-ip-list-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfgGslbPolicyMax          int `json:"gslb-policy-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfgGslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfgGslbServiceMax          int `json:"gslb-service-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfgGslbServiceMinGuarantee int `json:"gslb-service-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfgGslbServiceIPMax          int `json:"gslb-service-ip-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfgGslbServiceIPMinGuarantee int `json:"gslb-service-ip-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfgGslbServicePortMax          int `json:"gslb-service-port-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfgGslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfgGslbSiteMax          int `json:"gslb-site-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfgGslbSiteMinGuarantee int `json:"gslb-site-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfgGslbSvcGroupMax          int `json:"gslb-svc-group-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfgGslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfgGslbTemplateMax          int `json:"gslb-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfgGslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfgGslbZoneMax          int `json:"gslb-zone-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfgGslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfgHTTPTemplateMax          int `json:"http-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfgHTTPTemplateMinGuarantee int `json:"http-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfgHealthMonitorMax          int `json:"health-monitor-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfgHealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfgLinkCostTemplateMax          int `json:"link-cost-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfgLinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfgPersistCookieTemplateMax          int `json:"persist-cookie-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfgPersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfgPersistSrcipTemplateMax          int `json:"persist-srcip-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfgPersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfgProxyTemplateMax          int `json:"proxy-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfgProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfgRealPortMax          int `json:"real-port-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfgRealPortMinGuarantee int `json:"real-port-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfgRealServerMax          int `json:"real-server-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfgRealServerMinGuarantee int `json:"real-server-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfgServerSslTemplateMax          int `json:"server-ssl-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfgServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfgServiceGroupMax          int `json:"service-group-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfgServiceGroupMinGuarantee int `json:"service-group-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfgStreamTemplateMax          int `json:"stream-template-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfgStreamTemplateMinGuarantee int `json:"stream-template-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfgVirtualPortMax          int `json:"virtual-port-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfgVirtualPortMinGuarantee int `json:"virtual-port-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfg struct {
	SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfgVirtualServerMax          int `json:"virtual-server-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfgVirtualServerMinGuarantee int `json:"virtual-server-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfgIpv4AclLineMax          int `json:"ipv4-acl-line-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfgIpv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfgIpv6AclLineMax          int `json:"ipv6-acl-line-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfgIpv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfgObjectGroupMax          int `json:"object-group-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfgObjectGroupMinGuarantee int `json:"object-group-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfgObjectGroupClauseMax          int `json:"object-group-clause-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfgObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfgStaticArpMax          int `json:"static-arp-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfgStaticArpMinGuarantee int `json:"static-arp-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfgStaticIpv4RouteMax          int `json:"static-ipv4-route-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfgStaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfgStaticIpv6RouteMax          int `json:"static-ipv6-route-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfgStaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfgStaticMacMax          int `json:"static-mac-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfgStaticMacMinGuarantee int `json:"static-mac-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg struct {
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfgStaticNeighborMax          int `json:"static-neighbor-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfgStaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfgBwLimitMax              int `json:"bw-limit-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfgBwLimitWatermarkDisable int `json:"bw-limit-watermark-disable,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfgConcurrentSessionLimitMax int `json:"concurrent-session-limit-max,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfgFwcpsLimitMax int `json:"fwcps-limit-max,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfgL4CpsLimitMax int `json:"l4cps-limit-max,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfgL4SessionLimitMax          string `json:"l4-session-limit-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfgL4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfgL7CpsLimitMax int `json:"l7cps-limit-max,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfgNatcpsLimitMax int `json:"natcps-limit-max,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfgSslThroughputLimitMax              int `json:"ssl-throughput-limit-max,omitempty"`
	SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfgSslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable,omitempty"`
}

type SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg struct {
	SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfgSslcpsLimitMax int `json:"sslcps-limit-max,omitempty"`
}

func PostSystem(id string, inst System, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/system", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m System
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSystem", data)
			if err != nil {
				return err
			}

		}
	}
	return err
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
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSystem", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
