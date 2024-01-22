

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallUdpStunTimeout struct {
	Inst struct {

    Port int `json:"port"`
    PortEnd int `json:"port-end"`
    StunTimeoutValPortRange int `json:"stun-timeout-val-port-range" dval:"2"`
    Uuid string `json:"uuid"`

	} `json:"stun-timeout"`
}

func (p *Cgnv6StatefulFirewallUdpStunTimeout) GetId() string{
    return strconv.Itoa(p.Inst.Port)+"+"+strconv.Itoa(p.Inst.PortEnd)
}

func (p *Cgnv6StatefulFirewallUdpStunTimeout) getPath() string{
    return "cgnv6/stateful-firewall/udp/stun-timeout"
}

func (p *Cgnv6StatefulFirewallUdpStunTimeout) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallUdpStunTimeout::Post")
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

func (p *Cgnv6StatefulFirewallUdpStunTimeout) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallUdpStunTimeout::Get")
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
func (p *Cgnv6StatefulFirewallUdpStunTimeout) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallUdpStunTimeout::Put")
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

func (p *Cgnv6StatefulFirewallUdpStunTimeout) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallUdpStunTimeout::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
