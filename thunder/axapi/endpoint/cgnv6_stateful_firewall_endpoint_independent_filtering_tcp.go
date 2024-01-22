

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6StatefulFirewallEndpointIndependentFilteringTcp struct {
	Inst struct {

    PortList []Cgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList `json:"port-list"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type Cgnv6StatefulFirewallEndpointIndependentFilteringTcpPortList struct {
    Port int `json:"port"`
    PortEnd int `json:"port-end"`
}

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringTcp) GetId() string{
    return "1"
}

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringTcp) getPath() string{
    return "cgnv6/stateful-firewall/endpoint-independent-filtering/tcp"
}

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringTcp::Post")
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

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringTcp::Get")
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
func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringTcp::Put")
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

func (p *Cgnv6StatefulFirewallEndpointIndependentFilteringTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6StatefulFirewallEndpointIndependentFilteringTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
