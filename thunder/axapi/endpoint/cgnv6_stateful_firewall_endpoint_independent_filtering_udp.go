

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallEndpointIndependentFilteringUdp struct {
	Inst struct {

    PortList []Cgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList `json:"port-list"`
    Uuid string `json:"uuid"`

	} `json:"udp"`
}


type Cgnv6StatefulFirewallEndpointIndependentFilteringUdpPortList struct {
    Port int `json:"port"`
    PortEnd int `json:"port-end"`
}

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringUdp) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringUdp) getPath() string{
    return "cgnv6/stateful-firewall/endpoint-independent-filtering/udp"
}

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringUdp::Post")
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

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringUdp::Get")
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
func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringUdp::Put")
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

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
