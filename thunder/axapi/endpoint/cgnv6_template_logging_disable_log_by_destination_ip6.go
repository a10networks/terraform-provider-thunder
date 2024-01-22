

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateLoggingDisableLogByDestinationIp6 struct {
	Inst struct {

    Icmp int `json:"icmp"`
    Ipv6Addr string `json:"ipv6-addr"`
    Others int `json:"others"`
    TcpList []Cgnv6TemplateLoggingDisableLogByDestinationIp6TcpList `json:"tcp-list"`
    UdpList []Cgnv6TemplateLoggingDisableLogByDestinationIp6UdpList `json:"udp-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"ip6"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6TcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type Cgnv6TemplateLoggingDisableLogByDestinationIp6UdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp6) GetId() string{
    return url.QueryEscape(p.Inst.Ipv6Addr)
}

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp6) getPath() string{
    return "cgnv6/template/logging/" +p.Inst.Name + "/disable-log-by-destination/ip6"
}

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp6::Post")
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

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp6::Get")
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
func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp6::Put")
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

func (p *Cgnv6TemplateLoggingDisableLogByDestinationIp6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateLoggingDisableLogByDestinationIp6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
