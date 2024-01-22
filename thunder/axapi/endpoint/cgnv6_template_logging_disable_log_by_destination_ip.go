

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateLoggingDisableLogByDestinationIp struct {
	Inst struct {

    Icmp int `json:"icmp"`
    Ipv4Addr string `json:"ipv4-addr"`
    Others int `json:"others"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationIpTcpList `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationIpUdpList `json:"udp-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"ip"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIpUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp) GetId() string{
    return url.QueryEscape(p.Inst.Ipv4Addr)
}

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp) getPath() string{
    return "cgnv6/template/logging/" +p.Inst.Name + "/disable-log-by-destination/ip"
}

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp::Post")
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

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp::Get")
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
func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp::Put")
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

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
