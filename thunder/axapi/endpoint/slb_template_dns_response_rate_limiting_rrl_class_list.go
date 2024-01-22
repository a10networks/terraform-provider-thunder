

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsResponseRateLimitingRrlClassList struct {
	Inst struct {

    LidList []SlbTemplateDnsResponseRateLimitingRrlClassListLidList `json:"lid-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"rrl-class-list"`
}


type SlbTemplateDnsResponseRateLimitingRrlClassListLidList struct {
    Lidnum int `json:"lidnum"`
    LidResponseRate int `json:"lid-response-rate" dval:"5"`
    LidSlipRate int `json:"lid-slip-rate"`
    LidTcRate int `json:"lid-tc-rate"`
    LidMatchSubnet string `json:"lid-match-subnet" dval:"255.255.255.255"`
    LidMatchSubnetV6 int `json:"lid-match-subnet-v6" dval:"128"`
    LidWindow int `json:"lid-window" dval:"1"`
    LidSrcIpOnly int `json:"lid-src-ip-only"`
    LidEnableLog int `json:"lid-enable-log"`
    LidAction string `json:"lid-action" dval:"rate-limit"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *SlbTemplateDnsResponseRateLimitingRrlClassList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDnsResponseRateLimitingRrlClassList) getPath() string{
    return "slb/template/dns/response-rate-limiting/rrl-class-list"
}

func (p *SlbTemplateDnsResponseRateLimitingRrlClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassList::Post")
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

func (p *SlbTemplateDnsResponseRateLimitingRrlClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassList::Get")
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
func (p *SlbTemplateDnsResponseRateLimitingRrlClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassList::Put")
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

func (p *SlbTemplateDnsResponseRateLimitingRrlClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimitingRrlClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
