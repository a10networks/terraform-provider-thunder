

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type System struct {
	Inst struct {

    AddCpuCore SystemAddCpuCore1647 `json:"add-cpu-core"`
    AddPort SystemAddPort1648 `json:"add-port"`
    AllVlanLimit SystemAllVlanLimit1649 `json:"all-vlan-limit"`
    AnomalyLog int `json:"anomaly-log"`
    AnomalyLogRateLimit int `json:"anomaly-log-rate-limit"`
    AppPerformance SystemAppPerformance1650 `json:"app-performance"`
    AppsGlobal SystemAppsGlobal1652 `json:"apps-global"`
    AsicDebugDump SystemAsicDebugDump1653 `json:"asic-debug-dump"`
    AsicMmuFailSafe SystemAsicMmuFailSafe1654 `json:"asic-mmu-fail-safe"`
    AttackLog int `json:"attack-log"`
    Bandwidth SystemBandwidth1655 `json:"bandwidth"`
    Bfd SystemBfd1657 `json:"bfd"`
    ClassListHitcountEnable int `json:"class-list-hitcount-enable" dval:"1"`
    CliMonitorInterval SystemCliMonitorInterval1659 `json:"cli-monitor-interval"`
    CmUpdateFileNameRef SystemCmUpdateFileNameRef1660 `json:"cm-update-file-name-ref"`
    ControlCpu SystemControlCpu1661 `json:"control-cpu"`
    Core SystemCore1662 `json:"core"`
    CosqShow SystemCosqShow1663 `json:"cosq-show"`
    CosqStats SystemCosqStats1664 `json:"cosq-stats"`
    CounterLibAccounting SystemCounterLibAccounting1665 `json:"counter-lib-accounting"`
    CpuHyperThread SystemCpuHyperThread1666 `json:"cpu-hyper-thread"`
    CpuList SystemCpuList1667 `json:"cpu-list"`
    CpuLoadSharing SystemCpuLoadSharing1668 `json:"cpu-load-sharing"`
    CpuMap SystemCpuMap1671 `json:"cpu-map"`
    CpuPacketPrioSupport SystemCpuPacketPrioSupport1672 `json:"cpu-packet-prio-support"`
    DataCpu SystemDataCpu1673 `json:"data-cpu"`
    DdosAttack int `json:"ddos-attack"`
    DdosLog int `json:"ddos-log"`
    DefaultMtu int `json:"default-mtu"`
    DelPort SystemDelPort1674 `json:"del-port"`
    DeleteCpuCore SystemDeleteCpuCore1675 `json:"delete-cpu-core"`
    Dns SystemDns1676 `json:"dns"`
    DnsCache SystemDnsCache1680 `json:"dns-cache"`
    DomainListHitcountEnable int `json:"domain-list-hitcount-enable"`
    DomainListInfo SystemDomainListInfo1682 `json:"domain-list-info"`
    DpdkStats SystemDpdkStats1683 `json:"dpdk-stats"`
    DropLinuxClosedPortSyn string `json:"drop-linux-closed-port-syn" dval:"enable"`
    DynamicServiceDnsSocketPool int `json:"dynamic-service-dns-socket-pool"`
    EnablePassword SystemEnablePassword1685 `json:"enable-password"`
    Environment SystemEnvironment1686 `json:"environment"`
    ExtOnlyLogging SystemExtOnlyLogging1687 `json:"ext-only-logging"`
    FpgaCoreCrc SystemFpgaCoreCrc1688 `json:"fpga-core-crc"`
    FpgaDrop SystemFpgaDrop1689 `json:"fpga-drop"`
    Fw SystemFw1691 `json:"fw"`
    GeoDbHitcountEnable int `json:"geo-db-hitcount-enable"`
    GeoLocation SystemGeoLocation1692 `json:"geo-location"`
    Geoloc SystemGeoloc1696 `json:"geoloc"`
    GeolocListList []SystemGeolocListList `json:"geoloc-list-list"`
    GeolocNameHelper SystemGeolocNameHelper1698 `json:"geoloc-name-helper"`
    GeolocationFile SystemGeolocationFile1700 `json:"geolocation-file"`
    Glid SystemGlid1702 `json:"glid"`
    GuestFile SystemGuestFile1703 `json:"guest-file"`
    GuiImageList SystemGuiImageList1704 `json:"gui-image-list"`
    Hardware SystemHardware1705 `json:"hardware"`
    HardwareAccelerate SystemHardwareAccelerate1706 `json:"hardware-accelerate"`
    HealthCheckList []SystemHealthCheckList `json:"health-check-list"`
    HighMemoryL4Session SystemHighMemoryL4Session1710 `json:"high-memory-l4-session"`
    HrxqStatus SystemHrxqStatus1711 `json:"hrxq-status"`
    HwBlockingEnable int `json:"hw-blocking-enable"`
    Icmp SystemIcmp1712 `json:"icmp"`
    IcmpRate SystemIcmpRate1714 `json:"icmp-rate"`
    Icmp6 SystemIcmp61716 `json:"icmp6"`
    InuseCpuList SystemInuseCpuList1718 `json:"inuse-cpu-list"`
    InusePortList SystemInusePortList1719 `json:"inuse-port-list"`
    IoCpu SystemIoCpu1720 `json:"io-cpu"`
    IpDnsCache SystemIpDnsCache1721 `json:"ip-dns-cache"`
    IpStats SystemIpStats1722 `json:"ip-stats"`
    IpThreatList SystemIpThreatList1724 `json:"ip-threat-list"`
    Ip6Stats SystemIp6Stats1738 `json:"ip6-stats"`
    Ipmi SystemIpmi1740 `json:"ipmi"`
    IpmiService SystemIpmiService1745 `json:"ipmi-service"`
    Ipsec SystemIpsec1746 `json:"ipsec"`
    Ipv6PrefixLength int `json:"ipv6-prefix-length" dval:"128"`
    JobOffload SystemJobOffload1748 `json:"job-offload"`
    LinkCapability SystemLinkCapability1750 `json:"link-capability"`
    LinkMonitor SystemLinkMonitor1751 `json:"link-monitor"`
    Lro SystemLro1752 `json:"lro"`
    ManagementInterfaceMode SystemManagementInterfaceMode1753 `json:"management-interface-mode"`
    Memory SystemMemory1754 `json:"memory"`
    MemoryBlockDebug SystemMemoryBlockDebug1756 `json:"memory-block-debug"`
    MfaAuth SystemMfaAuth1757 `json:"mfa-auth"`
    MfaCertStore SystemMfaCertStore1758 `json:"mfa-cert-store"`
    MfaManagement SystemMfaManagement1759 `json:"mfa-management"`
    MfaValidationType SystemMfaValidationType1760 `json:"mfa-validation-type"`
    MgmtPort SystemMgmtPort1761 `json:"mgmt-port"`
    ModifyPort SystemModifyPort1762 `json:"modify-port"`
    ModuleCtrlCpu string `json:"module-ctrl-cpu"`
    MonTemplate SystemMonTemplate1763 `json:"mon-template"`
    MultiQueueSupport SystemMultiQueueSupport1766 `json:"multi-queue-support"`
    NdiscRa SystemNdiscRa1767 `json:"ndisc-ra"`
    NetvscMonitor SystemNetvscMonitor1769 `json:"netvsc-monitor"`
    NsmA10lb SystemNsmA10lb1770 `json:"nsm-a10lb"`
    PasswordPolicy SystemPasswordPolicy1771 `json:"password-policy"`
    PathList []SystemPathList `json:"path-list"`
    Pbslb SystemPbslb1772 `json:"pbslb"`
    PerVlanLimit SystemPerVlanLimit1774 `json:"per-vlan-limit"`
    Platformtype SystemPlatformtype1775 `json:"platformtype"`
    PortCount SystemPortCount1776 `json:"port-count"`
    PortInfo SystemPortInfo1777 `json:"port-info"`
    PortList SystemPortList1778 `json:"port-list"`
    Ports SystemPorts1779 `json:"ports"`
    PowerOnSelfTest SystemPowerOnSelfTest1780 `json:"power-on-self-test"`
    ProbeNetworkDevices SystemProbeNetworkDevices1781 `json:"probe-network-devices"`
    PromiscuousMode int `json:"promiscuous-mode"`
    PsuInfo SystemPsuInfo1782 `json:"psu-info"`
    QInQ SystemQInQ1783 `json:"q-in-q"`
    QueuingBuffer SystemQueuingBuffer1784 `json:"queuing-buffer"`
    Radius SystemRadius1785 `json:"radius"`
    Reboot SystemReboot1791 `json:"reboot"`
    ResourceAccounting SystemResourceAccounting1792 `json:"resource-accounting"`
    ResourceUsage SystemResourceUsage1846 `json:"resource-usage"`
    Session SystemSession1848 `json:"session"`
    SessionReclaimLimit SystemSessionReclaimLimit1850 `json:"session-reclaim-limit"`
    SetRxtxDescSize SystemSetRxtxDescSize1851 `json:"set-rxtx-desc-size"`
    SetRxtxQueue SystemSetRxtxQueue1852 `json:"set-rxtx-queue"`
    SetTcpSynPerSec SystemSetTcpSynPerSec1853 `json:"set-tcp-syn-per-sec"`
    SharedPollMode SystemSharedPollMode1854 `json:"shared-poll-mode"`
    ShellPrivileges SystemShellPrivileges1855 `json:"shell-privileges"`
    ShmLogging SystemShmLogging1856 `json:"shm-logging"`
    Shutdown SystemShutdown1857 `json:"shutdown"`
    SockstressDisable int `json:"sockstress-disable"`
    SpeProfile SystemSpeProfile1858 `json:"spe-profile"`
    SpeStatus SystemSpeStatus1859 `json:"spe-status"`
    SrcIpHashEnable int `json:"src-ip-hash-enable"`
    SslReqQ SystemSslReqQ1860 `json:"ssl-req-q"`
    SslScv SystemSslScv1862 `json:"ssl-scv"`
    SslScvVerifyCrlSign SystemSslScvVerifyCrlSign1863 `json:"ssl-scv-verify-crl-sign"`
    SslScvVerifyHost SystemSslScvVerifyHost1864 `json:"ssl-scv-verify-host"`
    SslSetCompatibleCipher SystemSslSetCompatibleCipher1865 `json:"ssl-set-compatible-cipher"`
    SslStatus SystemSslStatus1866 `json:"ssl-status"`
    SyslogTimeMsec SystemSyslogTimeMsec1867 `json:"syslog-time-msec"`
    SystemChassisPortSplitEnable int `json:"system-chassis-port-split-enable"`
    TableIntegrity SystemTableIntegrity1868 `json:"table-integrity"`
    Tcp SystemTcp1870 `json:"tcp"`
    TcpStats SystemTcpStats1873 `json:"tcp-stats"`
    TcpSynPerSec SystemTcpSynPerSec1875 `json:"tcp-syn-per-sec"`
    TelemetryLog SystemTelemetryLog1876 `json:"telemetry-log"`
    Template SystemTemplate1882 `json:"template"`
    TemplateBind SystemTemplateBind1883 `json:"template-bind"`
    Throughput SystemThroughput1884 `json:"throughput"`
    TimeoutValue SystemTimeoutValue1886 `json:"timeout-value"`
    Tls13Mgmt SystemTls13Mgmt1887 `json:"tls-1-3-mgmt"`
    Trunk SystemTrunk1888 `json:"trunk"`
    TrunkHwHash SystemTrunkHwHash1890 `json:"trunk-hw-hash"`
    TrunkXauiHwHash SystemTrunkXauiHwHash1891 `json:"trunk-xaui-hw-hash"`
    Tso SystemTso1892 `json:"tso"`
    UpgradeStatus SystemUpgradeStatus1893 `json:"upgrade-status"`
    Uuid string `json:"uuid"`
    VeMacScheme SystemVeMacScheme1894 `json:"ve-mac-scheme"`
    XauiDlbMode SystemXauiDlbMode1895 `json:"xaui-dlb-mode"`

	} `json:"system"`
}


type SystemAddCpuCore1647 struct {
    CoreIndex int `json:"core-index"`
}


type SystemAddPort1648 struct {
    PortIndex int `json:"port-index"`
}


type SystemAllVlanLimit1649 struct {
    Bcast int `json:"bcast" dval:"5000"`
    Ipmcast int `json:"ipmcast" dval:"5000"`
    Mcast int `json:"mcast" dval:"5000"`
    UnknownUcast int `json:"unknown-ucast" dval:"5000"`
    Uuid string `json:"uuid"`
}


type SystemAppPerformance1650 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemAppPerformanceSamplingEnable1651 `json:"sampling-enable"`
}


type SystemAppPerformanceSamplingEnable1651 struct {
    Counters1 string `json:"counters1"`
}


type SystemAppsGlobal1652 struct {
    LogSessionOnEstablished int `json:"log-session-on-established"`
    MslTime int `json:"msl-time" dval:"2"`
    TimerWheelWalkLimit int `json:"timer-wheel-walk-limit" dval:"100"`
    SessionsThreshold int `json:"sessions-threshold"`
    CpsThreshold int `json:"cps-threshold"`
    Uuid string `json:"uuid"`
}


type SystemAsicDebugDump1653 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemAsicMmuFailSafe1654 struct {
    RecoveryThreshold int `json:"recovery-threshold" dval:"2"`
    MonitorInterval int `json:"monitor-interval" dval:"60"`
    MonitorDisable int `json:"monitor-disable"`
    RebootDisable int `json:"reboot-disable"`
    InjectError int `json:"inject-error"`
    TestPatternType string `json:"test-pattern-type" dval:"lcb"`
    Uuid string `json:"uuid"`
}


type SystemBandwidth1655 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemBandwidthSamplingEnable1656 `json:"sampling-enable"`
}


type SystemBandwidthSamplingEnable1656 struct {
    Counters1 string `json:"counters1"`
}


type SystemBfd1657 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemBfdSamplingEnable1658 `json:"sampling-enable"`
}


type SystemBfdSamplingEnable1658 struct {
    Counters1 string `json:"counters1"`
}


type SystemCliMonitorInterval1659 struct {
    Interval int `json:"interval"`
    Uuid string `json:"uuid"`
}


type SystemCmUpdateFileNameRef1660 struct {
    Source_name string `json:"source_name"`
    Dest_name string `json:"dest_name"`
    Id1 int `json:"id1"`
}


type SystemControlCpu1661 struct {
    Uuid string `json:"uuid"`
}


type SystemCore1662 struct {
    Uuid string `json:"uuid"`
}


type SystemCosqShow1663 struct {
    Uuid string `json:"uuid"`
}


type SystemCosqStats1664 struct {
    Uuid string `json:"uuid"`
}


type SystemCounterLibAccounting1665 struct {
    Uuid string `json:"uuid"`
}


type SystemCpuHyperThread1666 struct {
    Enable int `json:"enable"`
    Disable int `json:"disable"`
}


type SystemCpuList1667 struct {
    Uuid string `json:"uuid"`
}


type SystemCpuLoadSharing1668 struct {
    Disable int `json:"disable"`
    PacketsPerSecond SystemCpuLoadSharingPacketsPerSecond1669 `json:"packets-per-second"`
    CpuUsage SystemCpuLoadSharingCpuUsage1670 `json:"cpu-usage"`
    Tcp int `json:"tcp"`
    Udp int `json:"udp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
}


type SystemCpuLoadSharingPacketsPerSecond1669 struct {
    Min int `json:"min" dval:"100000"`
}


type SystemCpuLoadSharingCpuUsage1670 struct {
    Low int `json:"low" dval:"60"`
    High int `json:"high" dval:"75"`
}


type SystemCpuMap1671 struct {
    Uuid string `json:"uuid"`
}


type SystemCpuPacketPrioSupport1672 struct {
    Enable int `json:"enable"`
    Disable int `json:"disable"`
}


type SystemDataCpu1673 struct {
    Uuid string `json:"uuid"`
}


type SystemDelPort1674 struct {
    PortIndex int `json:"port-index"`
}


type SystemDeleteCpuCore1675 struct {
    CoreIndex int `json:"core-index"`
}


type SystemDns1676 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemDnsSamplingEnable1677 `json:"sampling-enable"`
    RecursiveNameserver SystemDnsRecursiveNameserver1678 `json:"recursive-nameserver"`
}


type SystemDnsSamplingEnable1677 struct {
    Counters1 string `json:"counters1"`
}


type SystemDnsRecursiveNameserver1678 struct {
    FollowShared int `json:"follow-shared"`
    ServerList []SystemDnsRecursiveNameserverServerList1679 `json:"server-list"`
    Uuid string `json:"uuid"`
}


type SystemDnsRecursiveNameserverServerList1679 struct {
    Ipv4Addr string `json:"ipv4-addr"`
    V4Desc string `json:"v4-desc"`
    Ipv6Addr string `json:"ipv6-addr"`
    V6Desc string `json:"v6-desc"`
}


type SystemDnsCache1680 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemDnsCacheSamplingEnable1681 `json:"sampling-enable"`
}


type SystemDnsCacheSamplingEnable1681 struct {
    Counters1 string `json:"counters1"`
}


type SystemDomainListInfo1682 struct {
    Uuid string `json:"uuid"`
}


type SystemDpdkStats1683 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemDpdkStatsSamplingEnable1684 `json:"sampling-enable"`
}


type SystemDpdkStatsSamplingEnable1684 struct {
    Counters1 string `json:"counters1"`
}


type SystemEnablePassword1685 struct {
    FollowPasswordPolicy int `json:"follow-password-policy"`
    Uuid string `json:"uuid"`
}


type SystemEnvironment1686 struct {
    Uuid string `json:"uuid"`
}


type SystemExtOnlyLogging1687 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemFpgaCoreCrc1688 struct {
    MonitorDisable int `json:"monitor-disable"`
    RebootEnable int `json:"reboot-enable"`
    Uuid string `json:"uuid"`
}


type SystemFpgaDrop1689 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemFpgaDropSamplingEnable1690 `json:"sampling-enable"`
}


type SystemFpgaDropSamplingEnable1690 struct {
    Counters1 string `json:"counters1"`
}


type SystemFw1691 struct {
    ApplicationMempool int `json:"application-mempool"`
    ApplicationFlow int `json:"application-flow"`
    BasicDpiEnable int `json:"basic-dpi-enable"`
    Uuid string `json:"uuid"`
}


type SystemGeoLocation1692 struct {
    GeoLocationIana int `json:"geo-location-iana" dval:"1"`
    GeoLocationIanaSystem int `json:"geo-location-iana-system"`
    GeoLocationGeolite2Asn int `json:"geo-location-geolite2-asn"`
    Geolite2AsnIncludeIpv6 int `json:"geolite2-asn-include-ipv6"`
    GeoLocationGeolite2City int `json:"geo-location-geolite2-city"`
    Geolite2CityIncludeIpv6 int `json:"geolite2-city-include-ipv6"`
    GeoLocationGeolite2Country int `json:"geo-location-geolite2-country"`
    Geolite2CountryIncludeIpv6 int `json:"geolite2-country-include-ipv6"`
    GeolocLoadFileList []SystemGeoLocationGeolocLoadFileList1693 `json:"geoloc-load-file-list"`
    Uuid string `json:"uuid"`
    EntryList []SystemGeoLocationEntryList1694 `json:"entry-list"`
}


type SystemGeoLocationGeolocLoadFileList1693 struct {
    GeoLocationLoadFilename string `json:"geo-location-load-filename"`
    GeoLocationLoadFileIncludeIpv6 int `json:"geo-location-load-file-include-ipv6"`
    TemplateName string `json:"template-name"`
}


type SystemGeoLocationEntryList1694 struct {
    GeoLocnObjName string `json:"geo-locn-obj-name"`
    GeoLocnMultipleAddresses []SystemGeoLocationEntryListGeoLocnMultipleAddresses1695 `json:"geo-locn-multiple-addresses"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SystemGeoLocationEntryListGeoLocnMultipleAddresses1695 struct {
    FirstIpAddress string `json:"first-ip-address"`
    GeolIpv4Mask string `json:"geol-ipv4-mask"`
    IpAddr2 string `json:"ip-addr2"`
    FirstIpv6Address string `json:"first-ipv6-address"`
    GeolIpv6Mask int `json:"geol-ipv6-mask"`
    Ipv6Addr2 string `json:"ipv6-addr2"`
}


type SystemGeoloc1696 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemGeolocSamplingEnable1697 `json:"sampling-enable"`
}


type SystemGeolocSamplingEnable1697 struct {
    Counters1 string `json:"counters1"`
}


type SystemGeolocListList struct {
    Name string `json:"name"`
    Shared int `json:"shared"`
    IncludeGeolocNameList []SystemGeolocListListIncludeGeolocNameList `json:"include-geoloc-name-list"`
    ExcludeGeolocNameList []SystemGeolocListListExcludeGeolocNameList `json:"exclude-geoloc-name-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SystemGeolocListListSamplingEnable `json:"sampling-enable"`
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


type SystemGeolocNameHelper1698 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemGeolocNameHelperSamplingEnable1699 `json:"sampling-enable"`
}


type SystemGeolocNameHelperSamplingEnable1699 struct {
    Counters1 string `json:"counters1"`
}


type SystemGeolocationFile1700 struct {
    Uuid string `json:"uuid"`
    ErrorInfo SystemGeolocationFileErrorInfo1701 `json:"error-info"`
}


type SystemGeolocationFileErrorInfo1701 struct {
    Uuid string `json:"uuid"`
}


type SystemGlid1702 struct {
    GlidId string `json:"glid-id"`
    NonShared int `json:"non-shared"`
    Uuid string `json:"uuid"`
}


type SystemGuestFile1703 struct {
    Uuid string `json:"uuid"`
}


type SystemGuiImageList1704 struct {
    Uuid string `json:"uuid"`
}


type SystemHardware1705 struct {
    Uuid string `json:"uuid"`
}


type SystemHardwareAccelerate1706 struct {
    SessionForwarding int `json:"session-forwarding"`
    Uuid string `json:"uuid"`
    SamplingEnable []SystemHardwareAccelerateSamplingEnable1707 `json:"sampling-enable"`
    Slb SystemHardwareAccelerateSlb1708 `json:"slb"`
}


type SystemHardwareAccelerateSamplingEnable1707 struct {
    Counters1 string `json:"counters1"`
}


type SystemHardwareAccelerateSlb1708 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemHardwareAccelerateSlbSamplingEnable1709 `json:"sampling-enable"`
}


type SystemHardwareAccelerateSlbSamplingEnable1709 struct {
    Counters1 string `json:"counters1"`
}


type SystemHealthCheckList struct {
    L2hmHcName string `json:"l2hm-hc-name"`
    MethodL2bfd int `json:"method-l2bfd"`
    L2bfdTxInterval int `json:"l2bfd-tx-interval"`
    L2bfdRxInterval int `json:"l2bfd-rx-interval"`
    L2bfdMultiplier int `json:"l2bfd-multiplier"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SystemHighMemoryL4Session1710 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemHrxqStatus1711 struct {
    Uuid string `json:"uuid"`
}


type SystemIcmp1712 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemIcmpSamplingEnable1713 `json:"sampling-enable"`
}


type SystemIcmpSamplingEnable1713 struct {
    Counters1 string `json:"counters1"`
}


type SystemIcmpRate1714 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemIcmpRateSamplingEnable1715 `json:"sampling-enable"`
}


type SystemIcmpRateSamplingEnable1715 struct {
    Counters1 string `json:"counters1"`
}


type SystemIcmp61716 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemIcmp6SamplingEnable1717 `json:"sampling-enable"`
}


type SystemIcmp6SamplingEnable1717 struct {
    Counters1 string `json:"counters1"`
}


type SystemInuseCpuList1718 struct {
    Uuid string `json:"uuid"`
}


type SystemInusePortList1719 struct {
    Uuid string `json:"uuid"`
}


type SystemIoCpu1720 struct {
    MaxCores int `json:"max-cores"`
}


type SystemIpDnsCache1721 struct {
    Uuid string `json:"uuid"`
}


type SystemIpStats1722 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemIpStatsSamplingEnable1723 `json:"sampling-enable"`
}


type SystemIpStatsSamplingEnable1723 struct {
    Counters1 string `json:"counters1"`
}


type SystemIpThreatList1724 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemIpThreatListSamplingEnable1725 `json:"sampling-enable"`
    Ipv4SourceList SystemIpThreatListIpv4SourceList1726 `json:"ipv4-source-list"`
    Ipv4DestList SystemIpThreatListIpv4DestList1728 `json:"ipv4-dest-list"`
    Ipv6SourceList SystemIpThreatListIpv6SourceList1730 `json:"ipv6-source-list"`
    Ipv6DestList SystemIpThreatListIpv6DestList1732 `json:"ipv6-dest-list"`
    Ipv4InternetHostList SystemIpThreatListIpv4InternetHostList1734 `json:"ipv4-internet-host-list"`
    Ipv6InternetHostList SystemIpThreatListIpv6InternetHostList1736 `json:"ipv6-internet-host-list"`
}


type SystemIpThreatListSamplingEnable1725 struct {
    Counters1 string `json:"counters1"`
}


type SystemIpThreatListIpv4SourceList1726 struct {
    ClassListCfg []SystemIpThreatListIpv4SourceListClassListCfg1727 `json:"class-list-cfg"`
    Uuid string `json:"uuid"`
}


type SystemIpThreatListIpv4SourceListClassListCfg1727 struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}


type SystemIpThreatListIpv4DestList1728 struct {
    ClassListCfg []SystemIpThreatListIpv4DestListClassListCfg1729 `json:"class-list-cfg"`
    Uuid string `json:"uuid"`
}


type SystemIpThreatListIpv4DestListClassListCfg1729 struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}


type SystemIpThreatListIpv6SourceList1730 struct {
    ClassListCfg []SystemIpThreatListIpv6SourceListClassListCfg1731 `json:"class-list-cfg"`
    Uuid string `json:"uuid"`
}


type SystemIpThreatListIpv6SourceListClassListCfg1731 struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}


type SystemIpThreatListIpv6DestList1732 struct {
    ClassListCfg []SystemIpThreatListIpv6DestListClassListCfg1733 `json:"class-list-cfg"`
    Uuid string `json:"uuid"`
}


type SystemIpThreatListIpv6DestListClassListCfg1733 struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}


type SystemIpThreatListIpv4InternetHostList1734 struct {
    ClassListCfg []SystemIpThreatListIpv4InternetHostListClassListCfg1735 `json:"class-list-cfg"`
    Uuid string `json:"uuid"`
}


type SystemIpThreatListIpv4InternetHostListClassListCfg1735 struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}


type SystemIpThreatListIpv6InternetHostList1736 struct {
    ClassListCfg []SystemIpThreatListIpv6InternetHostListClassListCfg1737 `json:"class-list-cfg"`
    Uuid string `json:"uuid"`
}


type SystemIpThreatListIpv6InternetHostListClassListCfg1737 struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}


type SystemIp6Stats1738 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemIp6StatsSamplingEnable1739 `json:"sampling-enable"`
}


type SystemIp6StatsSamplingEnable1739 struct {
    Counters1 string `json:"counters1"`
}


type SystemIpmi1740 struct {
    Reset int `json:"reset"`
    Ip SystemIpmiIp1741 `json:"ip"`
    Ipsrc SystemIpmiIpsrc1742 `json:"ipsrc"`
    User SystemIpmiUser1743 `json:"user"`
    Tool SystemIpmiTool1744 `json:"tool"`
}


type SystemIpmiIp1741 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    DefaultGateway string `json:"default-gateway"`
}


type SystemIpmiIpsrc1742 struct {
    Dhcp int `json:"dhcp"`
    Static int `json:"static"`
}


type SystemIpmiUser1743 struct {
    Add string `json:"add"`
    Password string `json:"password"`
    Administrator int `json:"administrator"`
    Callback int `json:"callback"`
    Operator int `json:"operator"`
    User int `json:"user"`
    Disable string `json:"disable"`
    Privilege string `json:"privilege"`
    Setname string `json:"setname"`
    Newname string `json:"newname"`
    Setpass string `json:"setpass"`
    Newpass string `json:"newpass"`
}


type SystemIpmiTool1744 struct {
    Cmd string `json:"cmd"`
}


type SystemIpmiService1745 struct {
    Disable int `json:"disable"`
    Uuid string `json:"uuid"`
}


type SystemIpsec1746 struct {
    PacketRoundRobin int `json:"packet-round-robin"`
    CryptoCore int `json:"crypto-core"`
    CryptoMem int `json:"crypto-mem"`
    Qat int `json:"QAT"`
    Uuid string `json:"uuid"`
    FpgaDecrypt SystemIpsecFpgaDecrypt1747 `json:"fpga-decrypt"`
}


type SystemIpsecFpgaDecrypt1747 struct {
    Action string `json:"action" dval:"disable"`
}


type SystemJobOffload1748 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemJobOffloadSamplingEnable1749 `json:"sampling-enable"`
}


type SystemJobOffloadSamplingEnable1749 struct {
    Counters1 string `json:"counters1"`
}


type SystemLinkCapability1750 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemLinkMonitor1751 struct {
    Enable int `json:"enable"`
    Disable int `json:"disable"`
}


type SystemLro1752 struct {
    Enable int `json:"enable"`
    Disable int `json:"disable"`
}


type SystemManagementInterfaceMode1753 struct {
    Dedicated int `json:"dedicated"`
    NonDedicated int `json:"non-dedicated"`
}


type SystemMemory1754 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemMemorySamplingEnable1755 `json:"sampling-enable"`
}


type SystemMemorySamplingEnable1755 struct {
    Counters1 string `json:"counters1"`
}


type SystemMemoryBlockDebug1756 struct {
    AssertBlock int `json:"assert-block" dval:"65536"`
    PktdumpBlock int `json:"pktdump-block"`
    FirstBlk int `json:"first-blk" dval:"8192"`
    SecondBlk int `json:"second-blk" dval:"16384"`
    ThirdBlk int `json:"third-blk" dval:"32768"`
    FourthBlk int `json:"fourth-blk" dval:"65536"`
    Uuid string `json:"uuid"`
}


type SystemMfaAuth1757 struct {
    Username string `json:"username"`
    SecondFactor string `json:"second-factor"`
}


type SystemMfaCertStore1758 struct {
    CertHost string `json:"cert-host"`
    Protocol string `json:"protocol"`
    CertStorePath string `json:"cert-store-path"`
    Username string `json:"username"`
    PasswdString string `json:"passwd-string"`
    Encrypted string `json:"encrypted"`
    Uuid string `json:"uuid"`
}


type SystemMfaManagement1759 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemMfaValidationType1760 struct {
    CaCert string `json:"ca-cert"`
    Uuid string `json:"uuid"`
}


type SystemMgmtPort1761 struct {
    PortIndex int `json:"port-index"`
    MacAddress string `json:"mac-address"`
    PciAddress string `json:"pci-address"`
}


type SystemModifyPort1762 struct {
    PortIndex int `json:"port-index"`
    PortNumber int `json:"port-number"`
}


type SystemMonTemplate1763 struct {
    MonitorList []SystemMonTemplateMonitorList `json:"monitor-list"`
    LinkBlockAsDown SystemMonTemplateLinkBlockAsDown1764 `json:"link-block-as-down"`
    LinkDownOnRestart SystemMonTemplateLinkDownOnRestart1765 `json:"link-down-on-restart"`
}


type SystemMonTemplateMonitorList struct {
    Id1 int `json:"id1"`
    ClearCfg []SystemMonTemplateMonitorListClearCfg `json:"clear-cfg"`
    LinkDisableCfg []SystemMonTemplateMonitorListLinkDisableCfg `json:"link-disable-cfg"`
    LinkEnableCfg []SystemMonTemplateMonitorListLinkEnableCfg `json:"link-enable-cfg"`
    MonitorRelation string `json:"monitor-relation" dval:"monitor-and"`
    LinkUpCfg []SystemMonTemplateMonitorListLinkUpCfg `json:"link-up-cfg"`
    LinkDownCfg []SystemMonTemplateMonitorListLinkDownCfg `json:"link-down-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SystemMonTemplateMonitorListClearCfg struct {
    Sessions string `json:"sessions"`
    ClearAllSequence int `json:"clear-all-sequence"`
    ClearSequence int `json:"clear-sequence"`
}


type SystemMonTemplateMonitorListLinkDisableCfg struct {
    Diseth int `json:"diseth"`
    DisSequence int `json:"dis-sequence"`
}


type SystemMonTemplateMonitorListLinkEnableCfg struct {
    Enaeth int `json:"enaeth"`
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


type SystemMonTemplateLinkBlockAsDown1764 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemMonTemplateLinkDownOnRestart1765 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemMultiQueueSupport1766 struct {
    Enable int `json:"enable"`
}


type SystemNdiscRa1767 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemNdiscRaSamplingEnable1768 `json:"sampling-enable"`
}


type SystemNdiscRaSamplingEnable1768 struct {
    Counters1 string `json:"counters1"`
}


type SystemNetvscMonitor1769 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemNsmA10lb1770 struct {
    Kill int `json:"kill"`
    Uuid string `json:"uuid"`
}


type SystemPasswordPolicy1771 struct {
    Complexity string `json:"complexity"`
    Aging string `json:"aging"`
    History string `json:"history"`
    MinPswdLen int `json:"min-pswd-len"`
    UsernameCheck string `json:"username-check" dval:"disable"`
    RepeatCharacterCheck string `json:"repeat-character-check" dval:"disable"`
    ForbidConsecutiveCharacter string `json:"forbid-consecutive-character" dval:"0"`
    Uuid string `json:"uuid"`
}


type SystemPathList struct {
    L2hmPathName string `json:"l2hm-path-name"`
    L2hmVlan int `json:"l2hm-vlan"`
    L2hmSetupTestApi int `json:"l2hm-setup-test-api"`
    IfpairEthStart int `json:"ifpair-eth-start"`
    IfpairEthEnd int `json:"ifpair-eth-end"`
    IfpairTrunkStart int `json:"ifpair-trunk-start"`
    IfpairTrunkEnd int `json:"ifpair-trunk-end"`
    L2hmAttach string `json:"l2hm-attach"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SystemPbslb1772 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemPbslbSamplingEnable1773 `json:"sampling-enable"`
}


type SystemPbslbSamplingEnable1773 struct {
    Counters1 string `json:"counters1"`
}


type SystemPerVlanLimit1774 struct {
    Bcast int `json:"bcast" dval:"1000"`
    Ipmcast int `json:"ipmcast" dval:"1000"`
    Mcast int `json:"mcast" dval:"1000"`
    UnknownUcast int `json:"unknown-ucast" dval:"1000"`
    Uuid string `json:"uuid"`
}


type SystemPlatformtype1775 struct {
    Uuid string `json:"uuid"`
}


type SystemPortCount1776 struct {
    PortCountKernel int `json:"port-count-kernel" dval:"18000"`
    PortCountHm int `json:"port-count-hm" dval:"1024"`
    PortCountLogging int `json:"port-count-logging" dval:"4096"`
    PortCountAlg int `json:"port-count-alg" dval:"6000"`
    Uuid string `json:"uuid"`
}


type SystemPortInfo1777 struct {
    Uuid string `json:"uuid"`
}


type SystemPortList1778 struct {
    Uuid string `json:"uuid"`
}


type SystemPorts1779 struct {
    LinkDetectionInterval int `json:"link-detection-interval" dval:"1000"`
    Uuid string `json:"uuid"`
}


type SystemPowerOnSelfTest1780 struct {
    Uuid string `json:"uuid"`
}


type SystemProbeNetworkDevices1781 struct {
}


type SystemPsuInfo1782 struct {
    Uuid string `json:"uuid"`
}


type SystemQInQ1783 struct {
    EnableAllPorts int `json:"enable-all-ports"`
    InnerTpid string `json:"inner-tpid"`
    OuterTpid string `json:"outer-tpid"`
    Uuid string `json:"uuid"`
}


type SystemQueuingBuffer1784 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemRadius1785 struct {
    Server SystemRadiusServer1786 `json:"server"`
}


type SystemRadiusServer1786 struct {
    ListenPort int `json:"listen-port" dval:"1813"`
    Remote SystemRadiusServerRemote1787 `json:"remote"`
    Secret int `json:"secret"`
    SecretString string `json:"secret-string"`
    Encrypted string `json:"encrypted"`
    Vrid int `json:"vrid"`
    Attribute []SystemRadiusServerAttribute1789 `json:"attribute"`
    DisableReply int `json:"disable-reply"`
    AccountingStart string `json:"accounting-start" dval:"append-entry"`
    AccountingStop string `json:"accounting-stop" dval:"delete-entry"`
    AccountingInterimUpdate string `json:"accounting-interim-update" dval:"ignore"`
    AccountingOn string `json:"accounting-on" dval:"ignore"`
    AttributeName string `json:"attribute-name"`
    CustomAttributeName string `json:"custom-attribute-name"`
    Uuid string `json:"uuid"`
    SamplingEnable []SystemRadiusServerSamplingEnable1790 `json:"sampling-enable"`
}


type SystemRadiusServerRemote1787 struct {
    IpList []SystemRadiusServerRemoteIpList1788 `json:"ip-list"`
}


type SystemRadiusServerRemoteIpList1788 struct {
    IpListName string `json:"ip-list-name"`
    IpListSecret int `json:"ip-list-secret"`
    IpListSecretString string `json:"ip-list-secret-string"`
    IpListEncrypted string `json:"ip-list-encrypted"`
}


type SystemRadiusServerAttribute1789 struct {
    AttributeValue string `json:"attribute-value"`
    PrefixLength string `json:"prefix-length"`
    PrefixVendor int `json:"prefix-vendor"`
    PrefixNumber int `json:"prefix-number"`
    Name string `json:"name"`
    Value string `json:"value"`
    CustomVendor int `json:"custom-vendor"`
    CustomNumber int `json:"custom-number"`
    Vendor int `json:"vendor"`
    Number int `json:"number"`
}


type SystemRadiusServerSamplingEnable1790 struct {
    Counters1 string `json:"counters1"`
}


type SystemReboot1791 struct {
    Uuid string `json:"uuid"`
}


type SystemResourceAccounting1792 struct {
    Uuid string `json:"uuid"`
    TemplateList []SystemResourceAccountingTemplateList1793 `json:"template-list"`
}


type SystemResourceAccountingTemplateList1793 struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    AppResources SystemResourceAccountingTemplateListAppResources1794 `json:"app-resources"`
    NetworkResources SystemResourceAccountingTemplateListNetworkResources1826 `json:"network-resources"`
    SystemResources SystemResourceAccountingTemplateListSystemResources1836 `json:"system-resources"`
}


type SystemResourceAccountingTemplateListAppResources1794 struct {
    GslbDeviceCfg SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1795 `json:"gslb-device-cfg"`
    GslbGeoLocationCfg SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1796 `json:"gslb-geo-location-cfg"`
    GslbIpListCfg SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1797 `json:"gslb-ip-list-cfg"`
    GslbPolicyCfg SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1798 `json:"gslb-policy-cfg"`
    GslbServiceCfg SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1799 `json:"gslb-service-cfg"`
    GslbServiceIpCfg SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1800 `json:"gslb-service-ip-cfg"`
    GslbServicePortCfg SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1801 `json:"gslb-service-port-cfg"`
    GslbSiteCfg SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1802 `json:"gslb-site-cfg"`
    GslbSvcGroupCfg SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1803 `json:"gslb-svc-group-cfg"`
    GslbTemplateCfg SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1804 `json:"gslb-template-cfg"`
    GslbZoneCfg SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1805 `json:"gslb-zone-cfg"`
    HealthMonitorCfg SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1806 `json:"health-monitor-cfg"`
    RealPortCfg SystemResourceAccountingTemplateListAppResourcesRealPortCfg1807 `json:"real-port-cfg"`
    RealServerCfg SystemResourceAccountingTemplateListAppResourcesRealServerCfg1808 `json:"real-server-cfg"`
    ServiceGroupCfg SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1809 `json:"service-group-cfg"`
    VirtualServerCfg SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1810 `json:"virtual-server-cfg"`
    VirtualPortCfg SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1811 `json:"virtual-port-cfg"`
    CacheTemplateCfg SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1812 `json:"cache-template-cfg"`
    ClientSslTemplateCfg SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1813 `json:"client-ssl-template-cfg"`
    ConnReuseTemplateCfg SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1814 `json:"conn-reuse-template-cfg"`
    FastTcpTemplateCfg SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1815 `json:"fast-tcp-template-cfg"`
    FastUdpTemplateCfg SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1816 `json:"fast-udp-template-cfg"`
    FixTemplateCfg SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1817 `json:"fix-template-cfg"`
    HttpTemplateCfg SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1818 `json:"http-template-cfg"`
    LinkCostTemplateCfg SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1819 `json:"link-cost-template-cfg"`
    PbslbEntryCfg SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1820 `json:"pbslb-entry-cfg"`
    PersistCookieTemplateCfg SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1821 `json:"persist-cookie-template-cfg"`
    PersistSrcipTemplateCfg SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1822 `json:"persist-srcip-template-cfg"`
    ServerSslTemplateCfg SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1823 `json:"server-ssl-template-cfg"`
    ProxyTemplateCfg SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1824 `json:"proxy-template-cfg"`
    StreamTemplateCfg SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1825 `json:"stream-template-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbDeviceCfg1795 struct {
    GslbDeviceMax int `json:"gslb-device-max"`
    GslbDeviceMinGuarantee int `json:"gslb-device-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg1796 struct {
    GslbGeoLocationMax int `json:"gslb-geo-location-max"`
    GslbGeoLocationMinGuarantee int `json:"gslb-geo-location-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbIpListCfg1797 struct {
    GslbIpListMax int `json:"gslb-ip-list-max"`
    GslbIpListMinGuarantee int `json:"gslb-ip-list-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbPolicyCfg1798 struct {
    GslbPolicyMax int `json:"gslb-policy-max"`
    GslbPolicyMinGuarantee int `json:"gslb-policy-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbServiceCfg1799 struct {
    GslbServiceMax int `json:"gslb-service-max"`
    GslbServiceMinGuarantee int `json:"gslb-service-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbServiceIpCfg1800 struct {
    GslbServiceIpMax int `json:"gslb-service-ip-max"`
    GslbServiceIpMinGuarantee int `json:"gslb-service-ip-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbServicePortCfg1801 struct {
    GslbServicePortMax int `json:"gslb-service-port-max"`
    GslbServicePortMinGuarantee int `json:"gslb-service-port-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbSiteCfg1802 struct {
    GslbSiteMax int `json:"gslb-site-max"`
    GslbSiteMinGuarantee int `json:"gslb-site-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg1803 struct {
    GslbSvcGroupMax int `json:"gslb-svc-group-max"`
    GslbSvcGroupMinGuarantee int `json:"gslb-svc-group-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbTemplateCfg1804 struct {
    GslbTemplateMax int `json:"gslb-template-max"`
    GslbTemplateMinGuarantee int `json:"gslb-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesGslbZoneCfg1805 struct {
    GslbZoneMax int `json:"gslb-zone-max"`
    GslbZoneMinGuarantee int `json:"gslb-zone-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesHealthMonitorCfg1806 struct {
    HealthMonitorMax int `json:"health-monitor-max"`
    HealthMonitorMinGuarantee int `json:"health-monitor-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesRealPortCfg1807 struct {
    RealPortMax int `json:"real-port-max"`
    RealPortMinGuarantee int `json:"real-port-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesRealServerCfg1808 struct {
    RealServerMax int `json:"real-server-max"`
    RealServerMinGuarantee int `json:"real-server-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesServiceGroupCfg1809 struct {
    ServiceGroupMax int `json:"service-group-max"`
    ServiceGroupMinGuarantee int `json:"service-group-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesVirtualServerCfg1810 struct {
    VirtualServerMax int `json:"virtual-server-max"`
    VirtualServerMinGuarantee int `json:"virtual-server-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesVirtualPortCfg1811 struct {
    VirtualPortMax int `json:"virtual-port-max"`
    VirtualPortMinGuarantee int `json:"virtual-port-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesCacheTemplateCfg1812 struct {
    CacheTemplateMax int `json:"cache-template-max"`
    CacheTemplateMinGuarantee int `json:"cache-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesClientSslTemplateCfg1813 struct {
    ClientSslTemplateMax int `json:"client-ssl-template-max"`
    ClientSslTemplateMinGuarantee int `json:"client-ssl-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg1814 struct {
    ConnReuseTemplateMax int `json:"conn-reuse-template-max"`
    ConnReuseTemplateMinGuarantee int `json:"conn-reuse-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesFastTcpTemplateCfg1815 struct {
    FastTcpTemplateMax int `json:"fast-tcp-template-max"`
    FastTcpTemplateMinGuarantee int `json:"fast-tcp-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesFastUdpTemplateCfg1816 struct {
    FastUdpTemplateMax int `json:"fast-udp-template-max"`
    FastUdpTemplateMinGuarantee int `json:"fast-udp-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesFixTemplateCfg1817 struct {
    FixTemplateMax int `json:"fix-template-max"`
    FixTemplateMinGuarantee int `json:"fix-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesHttpTemplateCfg1818 struct {
    HttpTemplateMax int `json:"http-template-max"`
    HttpTemplateMinGuarantee int `json:"http-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg1819 struct {
    LinkCostTemplateMax int `json:"link-cost-template-max"`
    LinkCostTemplateMinGuarantee int `json:"link-cost-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesPbslbEntryCfg1820 struct {
    PbslbEntryMax int `json:"pbslb-entry-max"`
    PbslbEntryMinGuarantee int `json:"pbslb-entry-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg1821 struct {
    PersistCookieTemplateMax int `json:"persist-cookie-template-max"`
    PersistCookieTemplateMinGuarantee int `json:"persist-cookie-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg1822 struct {
    PersistSrcipTemplateMax int `json:"persist-srcip-template-max"`
    PersistSrcipTemplateMinGuarantee int `json:"persist-srcip-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesServerSslTemplateCfg1823 struct {
    ServerSslTemplateMax int `json:"server-ssl-template-max"`
    ServerSslTemplateMinGuarantee int `json:"server-ssl-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesProxyTemplateCfg1824 struct {
    ProxyTemplateMax int `json:"proxy-template-max"`
    ProxyTemplateMinGuarantee int `json:"proxy-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListAppResourcesStreamTemplateCfg1825 struct {
    StreamTemplateMax int `json:"stream-template-max"`
    StreamTemplateMinGuarantee int `json:"stream-template-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResources1826 struct {
    StaticIpv4RouteCfg SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1827 `json:"static-ipv4-route-cfg"`
    StaticIpv6RouteCfg SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1828 `json:"static-ipv6-route-cfg"`
    Ipv4AclLineCfg SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1829 `json:"ipv4-acl-line-cfg"`
    Ipv6AclLineCfg SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1830 `json:"ipv6-acl-line-cfg"`
    StaticArpCfg SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1831 `json:"static-arp-cfg"`
    StaticNeighborCfg SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1832 `json:"static-neighbor-cfg"`
    StaticMacCfg SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1833 `json:"static-mac-cfg"`
    ObjectGroupCfg SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1834 `json:"object-group-cfg"`
    ObjectGroupClauseCfg SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1835 `json:"object-group-clause-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg1827 struct {
    StaticIpv4RouteMax int `json:"static-ipv4-route-max"`
    StaticIpv4RouteMinGuarantee int `json:"static-ipv4-route-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg1828 struct {
    StaticIpv6RouteMax int `json:"static-ipv6-route-max"`
    StaticIpv6RouteMinGuarantee int `json:"static-ipv6-route-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg1829 struct {
    Ipv4AclLineMax int `json:"ipv4-acl-line-max"`
    Ipv4AclLineMinGuarantee int `json:"ipv4-acl-line-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg1830 struct {
    Ipv6AclLineMax int `json:"ipv6-acl-line-max"`
    Ipv6AclLineMinGuarantee int `json:"ipv6-acl-line-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticArpCfg1831 struct {
    StaticArpMax int `json:"static-arp-max"`
    StaticArpMinGuarantee int `json:"static-arp-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg1832 struct {
    StaticNeighborMax int `json:"static-neighbor-max"`
    StaticNeighborMinGuarantee int `json:"static-neighbor-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesStaticMacCfg1833 struct {
    StaticMacMax int `json:"static-mac-max"`
    StaticMacMinGuarantee int `json:"static-mac-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupCfg1834 struct {
    ObjectGroupMax int `json:"object-group-max"`
    ObjectGroupMinGuarantee int `json:"object-group-min-guarantee"`
}


type SystemResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg1835 struct {
    ObjectGroupClauseMax int `json:"object-group-clause-max"`
    ObjectGroupClauseMinGuarantee int `json:"object-group-clause-min-guarantee"`
}


type SystemResourceAccountingTemplateListSystemResources1836 struct {
    BwLimitCfg SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1837 `json:"bw-limit-cfg"`
    ConcurrentSessionLimitCfg SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1838 `json:"concurrent-session-limit-cfg"`
    L4SessionLimitCfg SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1839 `json:"l4-session-limit-cfg"`
    L4cpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1840 `json:"l4cps-limit-cfg"`
    L7cpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1841 `json:"l7cps-limit-cfg"`
    NatcpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1842 `json:"natcps-limit-cfg"`
    FwcpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1843 `json:"fwcps-limit-cfg"`
    SslThroughputLimitCfg SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1844 `json:"ssl-throughput-limit-cfg"`
    SslcpsLimitCfg SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1845 `json:"sslcps-limit-cfg"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`
}


type SystemResourceAccountingTemplateListSystemResourcesBwLimitCfg1837 struct {
    BwLimitMax int `json:"bw-limit-max"`
    BwLimitWatermarkDisable int `json:"bw-limit-watermark-disable"`
}


type SystemResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg1838 struct {
    ConcurrentSessionLimitMax int `json:"concurrent-session-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg1839 struct {
    L4SessionLimitMax string `json:"l4-session-limit-max"`
    L4SessionLimitMinGuarantee string `json:"l4-session-limit-min-guarantee" dval:"0"`
}


type SystemResourceAccountingTemplateListSystemResourcesL4cpsLimitCfg1840 struct {
    L4cpsLimitMax int `json:"l4cps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesL7cpsLimitCfg1841 struct {
    L7cpsLimitMax int `json:"l7cps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg1842 struct {
    NatcpsLimitMax int `json:"natcps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg1843 struct {
    FwcpsLimitMax int `json:"fwcps-limit-max"`
}


type SystemResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg1844 struct {
    SslThroughputLimitMax int `json:"ssl-throughput-limit-max"`
    SslThroughputLimitWatermarkDisable int `json:"ssl-throughput-limit-watermark-disable"`
}


type SystemResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg1845 struct {
    SslcpsLimitMax int `json:"sslcps-limit-max"`
}


type SystemResourceUsage1846 struct {
    SslContextMemory int `json:"ssl-context-memory" dval:"2048"`
    SslDmaMemory int `json:"ssl-dma-memory" dval:"256"`
    NatPoolAddrCount int `json:"nat-pool-addr-count"`
    L4SessionCount int `json:"l4-session-count"`
    AuthPortalHtmlFileSize int `json:"auth-portal-html-file-size" dval:"20"`
    AuthPortalImageFileSize int `json:"auth-portal-image-file-size" dval:"6"`
    MaxAflexFileSize int `json:"max-aflex-file-size" dval:"32"`
    AflexTableEntryCount int `json:"aflex-table-entry-count"`
    ClassListIpv6AddrCount int `json:"class-list-ipv6-addr-count"`
    ClassListAcEntryCount int `json:"class-list-ac-entry-count"`
    ClassListEntryCount int `json:"class-list-entry-count"`
    MaxAflexAuthzCollectionNumber int `json:"max-aflex-authz-collection-number" dval:"512"`
    RadiusTableSize int `json:"radius-table-size"`
    AuthzPolicyNumber int `json:"authz-policy-number"`
    IpsecSaNumber int `json:"ipsec-sa-number"`
    RamCacheMemoryLimit int `json:"ram-cache-memory-limit"`
    AuthSessionCount int `json:"auth-session-count"`
    Uuid string `json:"uuid"`
    Visibility SystemResourceUsageVisibility1847 `json:"visibility"`
}


type SystemResourceUsageVisibility1847 struct {
    MonitoredEntityCount int `json:"monitored-entity-count"`
    Uuid string `json:"uuid"`
}


type SystemSession1848 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemSessionSamplingEnable1849 `json:"sampling-enable"`
}


type SystemSessionSamplingEnable1849 struct {
    Counters1 string `json:"counters1"`
}


type SystemSessionReclaimLimit1850 struct {
    NscanLimit int `json:"nscan-limit" dval:"4096"`
    ScanFreq int `json:"scan-freq" dval:"5"`
    Uuid string `json:"uuid"`
}


type SystemSetRxtxDescSize1851 struct {
    PortIndex int `json:"port-index"`
    RxdSize int `json:"rxd-size"`
    TxdSize int `json:"txd-size"`
}


type SystemSetRxtxQueue1852 struct {
    PortIndex int `json:"port-index"`
    RxqSize int `json:"rxq-size"`
    TxqSize int `json:"txq-size"`
}


type SystemSetTcpSynPerSec1853 struct {
    TcpSynValue int `json:"tcp-syn-value" dval:"70"`
    Uuid string `json:"uuid"`
}


type SystemSharedPollMode1854 struct {
    Enable int `json:"enable"`
    Disable int `json:"disable"`
}


type SystemShellPrivileges1855 struct {
    EnableShellPrivileges int `json:"enable-shell-privileges"`
    Uuid string `json:"uuid"`
}


type SystemShmLogging1856 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemShutdown1857 struct {
    Uuid string `json:"uuid"`
}


type SystemSpeProfile1858 struct {
    Action string `json:"action" dval:"ipv4-ipv6"`
}


type SystemSpeStatus1859 struct {
    Uuid string `json:"uuid"`
}


type SystemSslReqQ1860 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemSslReqQSamplingEnable1861 `json:"sampling-enable"`
}


type SystemSslReqQSamplingEnable1861 struct {
    Counters1 string `json:"counters1"`
}


type SystemSslScv1862 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemSslScvVerifyCrlSign1863 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemSslScvVerifyHost1864 struct {
    Disable int `json:"disable"`
    Uuid string `json:"uuid"`
}


type SystemSslSetCompatibleCipher1865 struct {
    Disable int `json:"disable"`
    Uuid string `json:"uuid"`
}


type SystemSslStatus1866 struct {
    Uuid string `json:"uuid"`
}


type SystemSyslogTimeMsec1867 struct {
    EnableFlag int `json:"enable-flag"`
}


type SystemTableIntegrity1868 struct {
    Table string `json:"table" dval:"all"`
    AuditAction string `json:"audit-action" dval:"enable"`
    AutoSyncAction string `json:"auto-sync-action" dval:"enable"`
    Uuid string `json:"uuid"`
    SamplingEnable []SystemTableIntegritySamplingEnable1869 `json:"sampling-enable"`
}


type SystemTableIntegritySamplingEnable1869 struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
    Counters3 string `json:"counters3"`
}


type SystemTcp1870 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemTcpSamplingEnable1871 `json:"sampling-enable"`
    RateLimitResetUnknownConn SystemTcpRateLimitResetUnknownConn1872 `json:"rate-limit-reset-unknown-conn"`
}


type SystemTcpSamplingEnable1871 struct {
    Counters1 string `json:"counters1"`
}


type SystemTcpRateLimitResetUnknownConn1872 struct {
    PktRateForResetUnknownConn int `json:"pkt-rate-for-reset-unknown-conn"`
    LogForResetUnknownConn int `json:"log-for-reset-unknown-conn"`
    Uuid string `json:"uuid"`
}


type SystemTcpStats1873 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemTcpStatsSamplingEnable1874 `json:"sampling-enable"`
}


type SystemTcpStatsSamplingEnable1874 struct {
    Counters1 string `json:"counters1"`
}


type SystemTcpSynPerSec1875 struct {
    Uuid string `json:"uuid"`
}


type SystemTelemetryLog1876 struct {
    TopKSourceList SystemTelemetryLogTopKSourceList1877 `json:"top-k-source-list"`
    TopKAppSvcList SystemTelemetryLogTopKAppSvcList1878 `json:"top-k-app-svc-list"`
    DeviceStatus SystemTelemetryLogDeviceStatus1879 `json:"device-status"`
    Environment SystemTelemetryLogEnvironment1880 `json:"environment"`
    PartitionMetrics SystemTelemetryLogPartitionMetrics1881 `json:"partition-metrics"`
}


type SystemTelemetryLogTopKSourceList1877 struct {
    Uuid string `json:"uuid"`
}


type SystemTelemetryLogTopKAppSvcList1878 struct {
    Uuid string `json:"uuid"`
}


type SystemTelemetryLogDeviceStatus1879 struct {
    Uuid string `json:"uuid"`
}


type SystemTelemetryLogEnvironment1880 struct {
    Uuid string `json:"uuid"`
}


type SystemTelemetryLogPartitionMetrics1881 struct {
    Uuid string `json:"uuid"`
}


type SystemTemplate1882 struct {
    TemplatePolicy string `json:"template-policy"`
    Uuid string `json:"uuid"`
}


type SystemTemplateBind1883 struct {
    MonitorList []SystemTemplateBindMonitorList `json:"monitor-list"`
}


type SystemTemplateBindMonitorList struct {
    TemplateMonitor int `json:"template-monitor"`
    Uuid string `json:"uuid"`
}


type SystemThroughput1884 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []SystemThroughputSamplingEnable1885 `json:"sampling-enable"`
}


type SystemThroughputSamplingEnable1885 struct {
    Counters1 string `json:"counters1"`
}


type SystemTimeoutValue1886 struct {
    Ftp int `json:"ftp" dval:"120"`
    Scp int `json:"scp" dval:"300"`
    Sftp int `json:"sftp"`
    Tftp int `json:"tftp" dval:"300"`
    Http int `json:"http"`
    Https int `json:"https"`
    Uuid string `json:"uuid"`
}


type SystemTls13Mgmt1887 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}


type SystemTrunk1888 struct {
    LoadBalance SystemTrunkLoadBalance1889 `json:"load-balance"`
}


type SystemTrunkLoadBalance1889 struct {
    UseL3 int `json:"use-l3"`
    UseL4 int `json:"use-l4"`
    Uuid string `json:"uuid"`
}


type SystemTrunkHwHash1890 struct {
    Mode int `json:"mode" dval:"6"`
    Uuid string `json:"uuid"`
}


type SystemTrunkXauiHwHash1891 struct {
    Mode int `json:"mode" dval:"6"`
    Uuid string `json:"uuid"`
}


type SystemTso1892 struct {
    Enable int `json:"enable"`
    Disable int `json:"disable"`
}


type SystemUpgradeStatus1893 struct {
    Uuid string `json:"uuid"`
}


type SystemVeMacScheme1894 struct {
    VeMacSchemeVal string `json:"ve-mac-scheme-val" dval:"hash-based"`
    Uuid string `json:"uuid"`
}


type SystemXauiDlbMode1895 struct {
    Enable int `json:"enable"`
    Uuid string `json:"uuid"`
}

func (p *System) GetId() string{
    return "1"
}

func (p *System) getPath() string{
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
        if len(axResp) > 0{
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
