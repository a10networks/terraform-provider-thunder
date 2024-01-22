

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallTcpStunTimeout struct {
	Inst struct {

    Port int `json:"port"`
    PortEnd int `json:"port-end"`
    StunTimeoutValPortRange int `json:"stun-timeout-val-port-range" dval:"2"`
    Uuid string `json:"uuid"`

	} `json:"stun-timeout"`
}

func (p *Cgnv6StatefulFirewallTcpStunTimeout) GetId() string{
    return strconv.Itoa(p.Inst.Port)+"+"+strconv.Itoa(p.Inst.PortEnd)
}

func (p *Cgnv6StatefulFirewallTcpStunTimeout) getPath() string{
    return "cgnv6/stateful-firewall/tcp/stun-timeout"
}

func (p *Cgnv6StatefulFirewallTcpStunTimeout) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallTcpStunTimeout::Post")
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

func (p *Cgnv6StatefulFirewallTcpStunTimeout) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallTcpStunTimeout::Get")
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
func (p *Cgnv6StatefulFirewallTcpStunTimeout) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallTcpStunTimeout::Put")
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

func (p *Cgnv6StatefulFirewallTcpStunTimeout) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallTcpStunTimeout::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
