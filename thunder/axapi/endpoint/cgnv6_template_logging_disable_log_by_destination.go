

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateLoggingDisableLogByDestination struct {
	Inst struct {

    Icmp int `json:"icmp"`
    IpList []Cgnv6TemplateLoggingDisableLogByDestinationIpList `json:"ip-list"`
    Ip6List []Cgnv6TemplateLoggingDisableLogByDestinationIp6List `json:"ip6-list"`
    Others int `json:"others"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationTcpList `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationUdpList `json:"udp-list"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"disable-log-by-destination"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpList struct {
    Ipv4Addr string `json:"ipv4-addr"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6List struct {
    Ipv6Addr string `json:"ipv6-addr"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList `json:"udp-list"`
    Icmp int `json:"icmp"`
    Others int `json:"others"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}

func (p *Cgnv6TemplateLoggingDisableLogByDestination) GetId() string{
    return "1"
}

func (p *Cgnv6TemplateLoggingDisableLogByDestination) getPath() string{
    return "cgnv6/template/logging/" +p.Inst.Name + "/disable-log-by-destination"
}

func (p *Cgnv6TemplateLoggingDisableLogByDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestination::Post")
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

func (p *Cgnv6TemplateLoggingDisableLogByDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestination::Get")
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
func (p *Cgnv6TemplateLoggingDisableLogByDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestination::Put")
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

func (p *Cgnv6TemplateLoggingDisableLogByDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestination::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
