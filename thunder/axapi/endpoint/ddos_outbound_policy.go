

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosOutboundPolicy struct {
	Inst struct {

    AsnBasedTracking DdosOutboundPolicyAsnBasedTracking287 `json:"asn-based-tracking"`
    CountryBasedTracking DdosOutboundPolicyCountryBasedTracking288 `json:"country-based-tracking"`
    Name string `json:"name"`
    PolicyClassListList []DdosOutboundPolicyPolicyClassListList `json:"policy-class-list-list"`
    PolicyDefaultClassList DdosOutboundPolicyPolicyDefaultClassList289 `json:"policy-default-class-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"outbound-policy"`
}


type DdosOutboundPolicyAsnBasedTracking287 struct {
    Configuration string `json:"configuration"`
    PerAsnGlid string `json:"per-asn-glid"`
    PacketRateTriggered int `json:"packet-rate-triggered"`
    Uuid string `json:"uuid"`
}


type DdosOutboundPolicyCountryBasedTracking288 struct {
    Configuration string `json:"configuration"`
    PerCountryGlid string `json:"per-country-glid"`
    PacketRateTriggered int `json:"packet-rate-triggered"`
    Uuid string `json:"uuid"`
}


type DdosOutboundPolicyPolicyClassListList struct {
    ClassListName string `json:"class-list-name"`
    ClassListGlid string `json:"class-list-glid"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosOutboundPolicyPolicyDefaultClassList289 struct {
    Configuration int `json:"configuration"`
    ClassListGlid string `json:"class-list-glid"`
    Uuid string `json:"uuid"`
}

func (p *DdosOutboundPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosOutboundPolicy) getPath() string{
    return "ddos/outbound-policy"
}

func (p *DdosOutboundPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicy::Post")
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

func (p *DdosOutboundPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicy::Get")
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
func (p *DdosOutboundPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicy::Put")
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

func (p *DdosOutboundPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosOutboundPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
