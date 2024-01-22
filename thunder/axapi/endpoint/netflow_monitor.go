

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitor struct {
	Inst struct {

    CounterPollingInterval int `json:"counter-polling-interval" dval:"60"`
    CustomRecord NetflowMonitorCustomRecord1055 `json:"custom-record"`
    Destination NetflowMonitorDestination1057 `json:"destination"`
    Disable int `json:"disable"`
    DisableLogByDestination NetflowMonitorDisableLogByDestination1060 `json:"disable-log-by-destination"`
    FlowTimeout int `json:"flow-timeout" dval:"10"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Protocol string `json:"protocol" dval:"v9"`
    Record NetflowMonitorRecord1069 `json:"record"`
    ResendTemplate NetflowMonitorResendTemplate1070 `json:"resend-template"`
    Sample NetflowMonitorSample1071 `json:"sample"`
    SamplingEnable []NetflowMonitorSamplingEnable `json:"sampling-enable"`
    Scope string `json:"scope" dval:"global"`
    SourceAddress NetflowMonitorSourceAddress1072 `json:"source-address"`
    SourceIpUseMgmt int `json:"source-ip-use-mgmt"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"monitor"`
}


type NetflowMonitorCustomRecord1055 struct {
    CustomCfg []NetflowMonitorCustomRecordCustomCfg1056 `json:"custom-cfg"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorCustomRecordCustomCfg1056 struct {
    Event string `json:"event"`
    IpfixTemplate string `json:"ipfix-template"`
}


type NetflowMonitorDestination1057 struct {
    ServiceGroup string `json:"service-group"`
    IpCfg NetflowMonitorDestinationIpCfg1058 `json:"ip-cfg"`
    Ipv6Cfg NetflowMonitorDestinationIpv6Cfg1059 `json:"ipv6-cfg"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorDestinationIpCfg1058 struct {
    Ip string `json:"ip"`
    Port4 int `json:"port4" dval:"9996"`
}


type NetflowMonitorDestinationIpv6Cfg1059 struct {
    Ipv6 string `json:"ipv6"`
    Port6 int `json:"port6" dval:"9996"`
}


type NetflowMonitorDisableLogByDestination1060 struct {
    TcpList []NetflowMonitorDisableLogByDestinationTcpList1061 `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationUdpList1062 `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    IpList []NetflowMonitorDisableLogByDestinationIpList1063 `json:"ip-list"`
    Ip6List []NetflowMonitorDisableLogByDestinationIp6List1066 `json:"ip6-list"`
}


type NetflowMonitorDisableLogByDestinationTcpList1061 struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationUdpList1062 struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIpList1063 struct {
    Ipv4Addr string `json:"ipv4-addr"`
    TcpList []NetflowMonitorDisableLogByDestinationIpListTcpList1064 `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationIpListUdpList1065 `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type NetflowMonitorDisableLogByDestinationIpListTcpList1064 struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIpListUdpList1065 struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIp6List1066 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    TcpList []NetflowMonitorDisableLogByDestinationIp6ListTcpList1067 `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationIp6ListUdpList1068 `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type NetflowMonitorDisableLogByDestinationIp6ListTcpList1067 struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIp6ListUdpList1068 struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type NetflowMonitorRecord1069 struct {
    NetflowV5 int `json:"netflow-v5"`
    NetflowV5Ext int `json:"netflow-v5-ext"`
    Nat44 int `json:"nat44"`
    Nat64 int `json:"nat64"`
    Dslite int `json:"dslite"`
    SesnEventNat44 string `json:"sesn-event-nat44"`
    SesnEventNat64 string `json:"sesn-event-nat64"`
    SesnEventDslite string `json:"sesn-event-dslite"`
    SesnEventFw4 string `json:"sesn-event-fw4"`
    SesnEventFw6 string `json:"sesn-event-fw6"`
    PortMappingNat44 string `json:"port-mapping-nat44"`
    PortMappingNat64 string `json:"port-mapping-nat64"`
    PortMappingDslite string `json:"port-mapping-dslite"`
    PortBatchNat44 string `json:"port-batch-nat44"`
    PortBatchNat64 string `json:"port-batch-nat64"`
    PortBatchDslite string `json:"port-batch-dslite"`
    PortBatchV2Nat44 string `json:"port-batch-v2-nat44"`
    PortBatchV2Nat64 string `json:"port-batch-v2-nat64"`
    PortBatchV2Dslite string `json:"port-batch-v2-dslite"`
    DdosGeneralStat int `json:"ddos-general-stat"`
    DdosHttpStat int `json:"ddos-http-stat"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorResendTemplate1070 struct {
    Timeout int `json:"timeout" dval:"1800"`
    Records int `json:"records" dval:"1000"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorSample1071 struct {
    EthernetList []NetflowMonitorSampleEthernetList `json:"ethernet-list"`
    VeList []NetflowMonitorSampleVeList `json:"ve-list"`
    NatPoolList []NetflowMonitorSampleNatPoolList `json:"nat-pool-list"`
}


type NetflowMonitorSampleEthernetList struct {
    Ifindex int `json:"ifindex"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorSampleVeList struct {
    VeNum int `json:"ve-num"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorSampleNatPoolList struct {
    PoolName string `json:"pool-name"`
    Uuid string `json:"uuid"`
}


type NetflowMonitorSamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
}


type NetflowMonitorSourceAddress1072 struct {
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`
}

func (p *NetflowMonitor) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *NetflowMonitor) getPath() string{
    return "netflow/monitor"
}

func (p *NetflowMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitor::Post")
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

func (p *NetflowMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitor::Get")
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
func (p *NetflowMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitor::Put")
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

func (p *NetflowMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
