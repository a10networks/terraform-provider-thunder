

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDohForwarder struct {
	Inst struct {

    BypassDoh int `json:"bypass-doh"`
    ForwardingIpv4 string `json:"forwarding-ipv4"`
    ForwardingIpv6 string `json:"forwarding-ipv6"`
    TcpServiceGroup string `json:"tcp-service-group"`
    UdpServiceGroup string `json:"udp-service-group"`
    Uuid string `json:"uuid"`
    V4Internal int `json:"v4-internal"`
    V4L4Proto string `json:"v4-l4-proto" dval:"both"`
    V4Port int `json:"v4-port" dval:"53"`
    V6Internal int `json:"v6-internal"`
    V6L4Proto string `json:"v6-l4-proto" dval:"both"`
    V6Port int `json:"v6-port" dval:"53"`
    Name string 

	} `json:"forwarder"`
}

func (p *SlbTemplateDohForwarder) GetId() string{
    return "1"
}

func (p *SlbTemplateDohForwarder) getPath() string{
    return "slb/template/doh/" +p.Inst.Name + "/forwarder"
}

func (p *SlbTemplateDohForwarder) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDohForwarder::Post")
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

func (p *SlbTemplateDohForwarder) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDohForwarder::Get")
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
func (p *SlbTemplateDohForwarder) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDohForwarder::Put")
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

func (p *SlbTemplateDohForwarder) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDohForwarder::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
