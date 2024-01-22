

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyClassListLid struct {
	Inst struct {

    ActionValue string `json:"action-value"`
    BwPer int `json:"bw-per"`
    BwRateLimit int `json:"bw-rate-limit"`
    ConnLimit int `json:"conn-limit"`
    ConnPer int `json:"conn-per"`
    ConnRateLimit int `json:"conn-rate-limit"`
    DirectAction int `json:"direct-action"`
    DirectActionInterval int `json:"direct-action-interval" dval:"3"`
    DirectActionValue string `json:"direct-action-value"`
    DirectFail int `json:"direct-fail"`
    DirectLoggingDrpRst int `json:"direct-logging-drp-rst"`
    DirectPbslbInterval int `json:"direct-pbslb-interval" dval:"3"`
    DirectPbslbLogging int `json:"direct-pbslb-logging"`
    DirectServiceGroup string `json:"direct-service-group"`
    Dns64 SlbTemplatePolicyClassListLidDns64 `json:"dns64"`
    Interval int `json:"interval"`
    Lidnum int `json:"lidnum"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    OverLimitAction int `json:"over-limit-action"`
    RequestLimit int `json:"request-limit"`
    RequestPer int `json:"request-per"`
    RequestRateLimit int `json:"request-rate-limit"`
    ResponseCodeRateLimit []SlbTemplatePolicyClassListLidResponseCodeRateLimit `json:"response-code-rate-limit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"lid"`
}


type SlbTemplatePolicyClassListLidDns64 struct {
    Disable int `json:"disable"`
    ExclusiveAnswer int `json:"exclusive-answer"`
    Prefix string `json:"prefix"`
}


type SlbTemplatePolicyClassListLidResponseCodeRateLimit struct {
    CodeRangeStart int `json:"code-range-start"`
    CodeRangeEnd int `json:"code-range-end"`
    Threshold int `json:"threshold"`
    Period int `json:"period"`
}

func (p *SlbTemplatePolicyClassListLid) GetId() string{
    return strconv.Itoa(p.Inst.Lidnum)
}

func (p *SlbTemplatePolicyClassListLid) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/class-list/lid"
}

func (p *SlbTemplatePolicyClassListLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassListLid::Post")
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

func (p *SlbTemplatePolicyClassListLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassListLid::Get")
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
func (p *SlbTemplatePolicyClassListLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassListLid::Put")
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

func (p *SlbTemplatePolicyClassListLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassListLid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
