

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorDisableLogByDestination struct {
	Inst struct {

    Icmp int `json:"icmp"`
    IpList []NetflowMonitorDisableLogByDestinationIpList `json:"ip-list"`
    Ip6List []NetflowMonitorDisableLogByDestinationIp6List `json:"ip6-list"`
    Others int `json:"others"`
    TcpList []NetflowMonitorDisableLogByDestinationTcpList `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationUdpList `json:"udp-list"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"disable-log-by-destination"`
}


type NetflowMonitorDisableLogByDestinationIpList struct {
    Ipv4Addr string `json:"ipv4-addr"`
    TcpList []NetflowMonitorDisableLogByDestinationIpListTcpList `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationIpListUdpList `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type NetflowMonitorDisableLogByDestinationIpListTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIpListUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIp6List struct {
    Ipv6Addr string `json:"ipv6-addr"`
    TcpList []NetflowMonitorDisableLogByDestinationIp6ListTcpList `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationIp6ListUdpList `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type NetflowMonitorDisableLogByDestinationIp6ListTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIp6ListUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type NetflowMonitorDisableLogByDestinationTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}

func (p *NetflowMonitorDisableLogByDestination) GetId() string{
    return "1"
}

func (p *NetflowMonitorDisableLogByDestination) getPath() string{
    return "netflow/monitor/" +p.Inst.Name + "/disable-log-by-destination"
}

func (p *NetflowMonitorDisableLogByDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestination::Post")
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

func (p *NetflowMonitorDisableLogByDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestination::Get")
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
func (p *NetflowMonitorDisableLogByDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestination::Put")
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

func (p *NetflowMonitorDisableLogByDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestination::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
