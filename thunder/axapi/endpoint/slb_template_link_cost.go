

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateLinkCost struct {
	Inst struct {

    BandwidthInterval int `json:"bandwidth-interval" dval:"5"`
    Name string `json:"name"`
    OverageBandwidthCost int `json:"overage-bandwidth-cost"`
    PrepaidBandwidth int `json:"prepaid-bandwidth"`
    SamplingEnable []SlbTemplateLinkCostSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"link-cost"`
}


type SlbTemplateLinkCostSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplateLinkCost) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateLinkCost) getPath() string{
    return "slb/template/link-cost"
}

func (p *SlbTemplateLinkCost) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkCost::Post")
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

func (p *SlbTemplateLinkCost) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkCost::Get")
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
func (p *SlbTemplateLinkCost) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkCost::Put")
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

func (p *SlbTemplateLinkCost) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkCost::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
