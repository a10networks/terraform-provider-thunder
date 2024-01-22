

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsResponseRateLimiting struct {
	Inst struct {

    Action string `json:"action" dval:"rate-limit"`
    EnableLog int `json:"enable-log"`
    FilterResponseRate int `json:"filter-response-rate" dval:"10"`
    MatchSubnet string `json:"match-subnet" dval:"255.255.255.255"`
    MatchSubnetV6 int `json:"match-subnet-v6" dval:"128"`
    ResponseRate int `json:"response-rate" dval:"5"`
    RrlClassListList []SlbTemplateDnsResponseRateLimitingRrlClassListList `json:"rrl-class-list-list"`
    SlipRate int `json:"slip-rate"`
    SrcIpOnly int `json:"src-ip-only"`
    TcRate int `json:"TC-rate"`
    Uuid string `json:"uuid"`
    Window int `json:"window" dval:"1"`
    Name string 

	} `json:"response-rate-limiting"`
}


type SlbTemplateDnsResponseRateLimitingRrlClassListList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    LidList []SlbTemplateDnsResponseRateLimitingRrlClassListListLidList `json:"lid-list"`
}


type SlbTemplateDnsResponseRateLimitingRrlClassListListLidList struct {
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

func (p *SlbTemplateDnsResponseRateLimiting) GetId() string{
    return "1"
}

func (p *SlbTemplateDnsResponseRateLimiting) getPath() string{
    return "slb/template/dns/" +p.Inst.Name + "/response-rate-limiting"
}

func (p *SlbTemplateDnsResponseRateLimiting) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimiting::Post")
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

func (p *SlbTemplateDnsResponseRateLimiting) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimiting::Get")
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
func (p *SlbTemplateDnsResponseRateLimiting) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimiting::Put")
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

func (p *SlbTemplateDnsResponseRateLimiting) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsResponseRateLimiting::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
