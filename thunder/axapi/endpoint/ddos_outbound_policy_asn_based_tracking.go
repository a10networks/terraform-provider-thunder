

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosOutboundPolicyAsnBasedTracking struct {
	Inst struct {

    Configuration string `json:"configuration"`
    PacketRateTriggered int `json:"packet-rate-triggered"`
    PerAsnGlid string `json:"per-asn-glid"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"asn-based-tracking"`
}

func (p *DdosOutboundPolicyAsnBasedTracking) GetId() string{
    return "1"
}

func (p *DdosOutboundPolicyAsnBasedTracking) getPath() string{
    return "ddos/outbound-policy/" +p.Inst.Name + "/asn-based-tracking"
}

func (p *DdosOutboundPolicyAsnBasedTracking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyAsnBasedTracking::Post")
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

func (p *DdosOutboundPolicyAsnBasedTracking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyAsnBasedTracking::Get")
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
func (p *DdosOutboundPolicyAsnBasedTracking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyAsnBasedTracking::Put")
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

func (p *DdosOutboundPolicyAsnBasedTracking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicyAsnBasedTracking::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
