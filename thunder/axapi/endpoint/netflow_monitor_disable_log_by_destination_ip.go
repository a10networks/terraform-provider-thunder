

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorDisableLogByDestinationIp struct {
	Inst struct {

    Icmp int `json:"icmp"`
    Ipv4Addr string `json:"ipv4-addr"`
    Others int `json:"others"`
    TcpList []NetflowMonitorDisableLogByDestinationIpTcpList `json:"tcp-list"`
    UdpList []NetflowMonitorDisableLogByDestinationIpUdpList `json:"udp-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"ip"`
}


type NetflowMonitorDisableLogByDestinationIpTcpList struct {
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
}


type NetflowMonitorDisableLogByDestinationIpUdpList struct {
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
}

func (p *NetflowMonitorDisableLogByDestinationIp) GetId() string{
    return url.QueryEscape(p.Inst.Ipv4Addr)
}

func (p *NetflowMonitorDisableLogByDestinationIp) getPath() string{
    return "netflow/monitor/" +p.Inst.Name + "/disable-log-by-destination/ip"
}

func (p *NetflowMonitorDisableLogByDestinationIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestinationIp::Post")
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

func (p *NetflowMonitorDisableLogByDestinationIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestinationIp::Get")
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
func (p *NetflowMonitorDisableLogByDestinationIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestinationIp::Put")
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

func (p *NetflowMonitorDisableLogByDestinationIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDisableLogByDestinationIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
