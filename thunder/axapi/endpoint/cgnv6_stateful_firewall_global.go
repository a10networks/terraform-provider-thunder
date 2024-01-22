

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallGlobal struct {
	Inst struct {

    RespondToUserMac int `json:"respond-to-user-mac"`
    SamplingEnable []Cgnv6StatefulFirewallGlobalSamplingEnable `json:"sampling-enable"`
    StatefulFirewallValue string `json:"stateful-firewall-value"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type Cgnv6StatefulFirewallGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6StatefulFirewallGlobal) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallGlobal) getPath() string{
    return "cgnv6/stateful-firewall/global"
}

func (p *Cgnv6StatefulFirewallGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallGlobal::Post")
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

func (p *Cgnv6StatefulFirewallGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallGlobal::Get")
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
func (p *Cgnv6StatefulFirewallGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallGlobal::Put")
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

func (p *Cgnv6StatefulFirewallGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
