

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnable struct {
	Inst struct {

    Service int `json:"service"`
    Traps SnmpServerEnableTraps1504 `json:"traps"`
    Uuid string `json:"uuid"`

	} `json:"enable"`
}


type SnmpServerEnableTraps1504 struct {
    All int `json:"all"`
    Lldp int `json:"lldp"`
    Uuid string `json:"uuid"`
    Routing SnmpServerEnableTrapsRouting1505 `json:"routing"`
    Gslb SnmpServerEnableTrapsGslb1510 `json:"gslb"`
    Slb SnmpServerEnableTrapsSlb1511 `json:"slb"`
    Scaleout SnmpServerEnableTrapsScaleout1512 `json:"scaleout"`
    Snmp SnmpServerEnableTrapsSnmp1517 `json:"snmp"`
    VrrpA SnmpServerEnableTrapsVrrpA1518 `json:"vrrp-a"`
    Vcs SnmpServerEnableTrapsVcs1519 `json:"vcs"`
    System SnmpServerEnableTrapsSystem1520 `json:"system"`
    SlbChange SnmpServerEnableTrapsSlbChange1522 `json:"slb-change"`
    Lsn SnmpServerEnableTrapsLsn1523 `json:"lsn"`
    Network SnmpServerEnableTrapsNetwork1524 `json:"network"`
    Ssl SnmpServerEnableTrapsSsl1525 `json:"ssl"`
}


type SnmpServerEnableTrapsRouting1505 struct {
    Bgp SnmpServerEnableTrapsRoutingBgp1506 `json:"bgp"`
    Isis SnmpServerEnableTrapsRoutingIsis1508 `json:"isis"`
    Ospf SnmpServerEnableTrapsRoutingOspf1509 `json:"ospf"`
}


type SnmpServerEnableTrapsRoutingBgp1506 struct {
    Bgpestablishednotification int `json:"bgpEstablishedNotification"`
    Bgpbackwardtransnotification int `json:"bgpBackwardTransNotification"`
    Uuid string `json:"uuid"`
    Ax SnmpServerEnableTrapsRoutingBgpAx1507 `json:"ax"`
}


type SnmpServerEnableTrapsRoutingBgpAx1507 struct {
    Bgpestablishednotification int `json:"bgpEstablishedNotification"`
    Bgpbackwardtransnotification int `json:"bgpBackwardTransNotification"`
    Bgpprefixthresholdexceedednotification int `json:"bgpPrefixThresholdExceededNotification"`
    Bgpprefixthresholdclearnotification int `json:"bgpPrefixThresholdClearNotification"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsRoutingIsis1508 struct {
    Isisadjacencychange int `json:"isisAdjacencyChange"`
    Isisareamismatch int `json:"isisAreaMismatch"`
    Isisattempttoexceedmaxsequence int `json:"isisAttemptToExceedMaxSequence"`
    Isisauthenticationfailure int `json:"isisAuthenticationFailure"`
    Isisauthenticationtypefailure int `json:"isisAuthenticationTypeFailure"`
    Isiscorruptedlspdetected int `json:"isisCorruptedLSPDetected"`
    Isisdatabaseoverload int `json:"isisDatabaseOverload"`
    Isisidlenmismatch int `json:"isisIDLenMismatch"`
    Isislsptoolargetopropagate int `json:"isisLSPTooLargeToPropagate"`
    Isismanualaddressdrops int `json:"isisManualAddressDrops"`
    Isismaxareaaddressesmismatch int `json:"isisMaxAreaAddressesMismatch"`
    Isisoriginatinglspbuffersizemismatch int `json:"isisOriginatingLSPBufferSizeMismatch"`
    Isisownlsppurge int `json:"isisOwnLSPPurge"`
    Isisprotocolssupportedmismatch int `json:"isisProtocolsSupportedMismatch"`
    Isisrejectedadjacency int `json:"isisRejectedAdjacency"`
    Isissequencenumberskip int `json:"isisSequenceNumberSkip"`
    Isisversionskew int `json:"isisVersionSkew"`
    Isislsperrordetected int `json:"isisLSPErrorDetected"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsRoutingOspf1509 struct {
    Ospfifauthfailure int `json:"ospfIfAuthFailure"`
    Ospfifconfigerror int `json:"ospfIfConfigError"`
    Ospfifrxbadpacket int `json:"ospfIfRxBadPacket"`
    Ospfifstatechange int `json:"ospfIfStateChange"`
    Ospflsdbapproachingoverflow int `json:"ospfLsdbApproachingOverflow"`
    Ospflsdboverflow int `json:"ospfLsdbOverflow"`
    Ospfmaxagelsa int `json:"ospfMaxAgeLsa"`
    Ospfnbrstatechange int `json:"ospfNbrStateChange"`
    Ospforiginatelsa int `json:"ospfOriginateLsa"`
    Ospftxretransmit int `json:"ospfTxRetransmit"`
    Ospfvirtifauthfailure int `json:"ospfVirtIfAuthFailure"`
    Ospfvirtifconfigerror int `json:"ospfVirtIfConfigError"`
    Ospfvirtifrxbadpacket int `json:"ospfVirtIfRxBadPacket"`
    Ospfvirtifstatechange int `json:"ospfVirtIfStateChange"`
    Ospfvirtiftxretransmit int `json:"ospfVirtIfTxRetransmit"`
    Ospfvirtnbrstatechange int `json:"ospfVirtNbrStateChange"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsGslb1510 struct {
    All int `json:"all"`
    Zone int `json:"zone"`
    Site int `json:"site"`
    Group int `json:"group"`
    ServiceIp int `json:"service-ip"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsSlb1511 struct {
    All int `json:"all"`
    ApplicationBufferLimit int `json:"application-buffer-limit"`
    GatewayUp int `json:"gateway-up"`
    GatewayDown int `json:"gateway-down"`
    ServerConnLimit int `json:"server-conn-limit"`
    ServerConnResume int `json:"server-conn-resume"`
    ServerUp int `json:"server-up"`
    ServerDown int `json:"server-down"`
    ServerDisabled int `json:"server-disabled"`
    ServerSelectionFailure int `json:"server-selection-failure"`
    ServiceConnLimit int `json:"service-conn-limit"`
    ServiceConnResume int `json:"service-conn-resume"`
    ServiceDown int `json:"service-down"`
    ServiceUp int `json:"service-up"`
    ServiceGroupUp int `json:"service-group-up"`
    ServiceGroupDown int `json:"service-group-down"`
    ServiceGroupMemberUp int `json:"service-group-member-up"`
    ServiceGroupMemberDown int `json:"service-group-member-down"`
    VipConnlimit int `json:"vip-connlimit"`
    VipConnratelimit int `json:"vip-connratelimit"`
    VipDown int `json:"vip-down"`
    VipPortConnlimit int `json:"vip-port-connlimit"`
    VipPortConnratelimit int `json:"vip-port-connratelimit"`
    VipPortDown int `json:"vip-port-down"`
    VipPortUp int `json:"vip-port-up"`
    VipUp int `json:"vip-up"`
    BwRateLimitExceed int `json:"bw-rate-limit-exceed"`
    BwRateLimitResume int `json:"bw-rate-limit-resume"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsScaleout1512 struct {
    Infrastructure SnmpServerEnableTrapsScaleoutInfrastructure1513 `json:"infrastructure"`
}


type SnmpServerEnableTrapsScaleoutInfrastructure1513 struct {
    All int `json:"all"`
    TestSendAllTraps int `json:"test-send-all-traps"`
    Uuid string `json:"uuid"`
    Cluster SnmpServerEnableTrapsScaleoutInfrastructureCluster1514 `json:"cluster"`
    ServiceNode SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1515 `json:"service-node"`
    MasterNode SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1516 `json:"master-node"`
}


type SnmpServerEnableTrapsScaleoutInfrastructureCluster1514 struct {
    Election int `json:"election"`
    MasterCallingReElection int `json:"master-calling-re-election"`
    NodeStatus int `json:"node-status"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsScaleoutInfrastructureServiceNode1515 struct {
    LocalDeviceDisabled int `json:"local-device-disabled"`
    ServiceMaster int `json:"service-master"`
    TrafficMapUpdate int `json:"traffic-map-update"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsScaleoutInfrastructureMasterNode1516 struct {
    TrafficMapDistribution int `json:"traffic-map-distribution"`
    VserverTrafficMapUpdate int `json:"vserver-traffic-map-update"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsSnmp1517 struct {
    All int `json:"all"`
    Linkdown int `json:"linkdown"`
    Linkup int `json:"linkup"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsVrrpA1518 struct {
    All int `json:"all"`
    Active int `json:"active"`
    Standby int `json:"standby"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsVcs1519 struct {
    StateChange int `json:"state-change"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsSystem1520 struct {
    All int `json:"all"`
    ControlCpuHigh int `json:"control-cpu-high"`
    DataCpuHigh int `json:"data-cpu-high"`
    Fan int `json:"fan"`
    FileSysReadOnly int `json:"file-sys-read-only"`
    HighDiskUse int `json:"high-disk-use"`
    HighMemoryUse int `json:"high-memory-use"`
    HighTemp int `json:"high-temp"`
    LowTemp int `json:"low-temp"`
    LicenseManagement int `json:"license-management"`
    PacketDrop int `json:"packet-drop"`
    Power int `json:"power"`
    PriDisk int `json:"pri-disk"`
    Restart int `json:"restart"`
    SecDisk int `json:"sec-disk"`
    Shutdown int `json:"shutdown"`
    SmpResourceEvent int `json:"smp-resource-event"`
    SyslogSeverityOne int `json:"syslog-severity-one"`
    TacacsServerUpDown int `json:"tacacs-server-up-down"`
    Start int `json:"start"`
    Uuid string `json:"uuid"`
    AppsGlobal SnmpServerEnableTrapsSystemAppsGlobal1521 `json:"apps-global"`
}


type SnmpServerEnableTrapsSystemAppsGlobal1521 struct {
    SessionsThreshold int `json:"sessions-threshold"`
    CpsThreshold int `json:"cps-threshold"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsSlbChange1522 struct {
    All int `json:"all"`
    ResourceUsageWarning int `json:"resource-usage-warning"`
    ConnectionResourceEvent int `json:"connection-resource-event"`
    Server int `json:"server"`
    ServerPort int `json:"server-port"`
    SslCertChange int `json:"ssl-cert-change"`
    SslCertExpire int `json:"ssl-cert-expire"`
    Vip int `json:"vip"`
    VipPort int `json:"vip-port"`
    SystemThreshold int `json:"system-threshold"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsLsn1523 struct {
    All int `json:"all"`
    TotalPortUsageThreshold int `json:"total-port-usage-threshold"`
    PerIpPortUsageThreshold int `json:"per-ip-port-usage-threshold"`
    MaxPortThreshold int `json:"max-port-threshold" dval:"655350000"`
    MaxIpportThreshold int `json:"max-ipport-threshold" dval:"64512"`
    FixedNatPortMappingFileChange int `json:"fixed-nat-port-mapping-file-change"`
    TrafficExceeded int `json:"traffic-exceeded"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsNetwork1524 struct {
    TrunkPortThreshold int `json:"trunk-port-threshold"`
    Uuid string `json:"uuid"`
}


type SnmpServerEnableTrapsSsl1525 struct {
    ServerCertificateError int `json:"server-certificate-error"`
    Uuid string `json:"uuid"`
}

func (p *SnmpServerEnable) GetId() string{
    return "1"
}

func (p *SnmpServerEnable) getPath() string{
    return "snmp-server/enable"
}

func (p *SnmpServerEnable) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnable::Post")
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

func (p *SnmpServerEnable) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnable::Get")
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
func (p *SnmpServerEnable) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnable::Put")
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

func (p *SnmpServerEnable) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnable::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
