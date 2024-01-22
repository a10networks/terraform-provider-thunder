

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FlowspecFilteringAction struct {
	Inst struct {

    CopyIpHost int `json:"copy-ip-host"`
    CopyIpHostNlri int `json:"copy-ip-host-nlri"`
    CopyIpv6Host int `json:"copy-ipv6-host"`
    CopyIpv6HostNlri int `json:"copy-ipv6-host-nlri"`
    DscpVal int `json:"dscp-val"`
    EcommCustomHex string `json:"ecomm-custom-hex"`
    IpHost string `json:"ip-host"`
    IpHostNlri string `json:"ip-host-nlri"`
    IpHostRt string `json:"ip-host-rt"`
    Ipv6Host string `json:"ipv6-host"`
    Ipv6HostNlri string `json:"ipv6-host-nlri"`
    NextHopNlriType string `json:"next-hop-nlri-type"`
    NextHopType string `json:"next-hop-type"`
    Redirect string `json:"redirect"`
    SampleLog int `json:"sample-log"`
    TerminalAction int `json:"terminal-action"`
    TrafficClass int `json:"traffic-class"`
    TrafficMarking string `json:"traffic-marking"`
    TrafficRate int `json:"traffic-rate"`
    Uuid string `json:"uuid"`
    ValueIpHost int `json:"value-ip-host"`
    VrfTargetIp string `json:"vrf-target-ip"`
    VrfTargetString string `json:"vrf-target-string"`
    Name string 

	} `json:"filtering-action"`
}

func (p *FlowspecFilteringAction) GetId() string{
    return "1"
}

func (p *FlowspecFilteringAction) getPath() string{
    return "flowspec/" +p.Inst.Name + "/filtering-action"
}

func (p *FlowspecFilteringAction) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecFilteringAction::Post")
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

func (p *FlowspecFilteringAction) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecFilteringAction::Get")
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
func (p *FlowspecFilteringAction) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecFilteringAction::Put")
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

func (p *FlowspecFilteringAction) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FlowspecFilteringAction::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
