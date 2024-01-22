

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyClassList struct {
	Inst struct {

    ClientIpL3Dest int `json:"client-ip-l3-dest"`
    ClientIpL7Header int `json:"client-ip-l7-header"`
    HeaderName string `json:"header-name"`
    LidList []SlbTemplatePolicyClassListLidList `json:"lid-list"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`

	} `json:"class-list"`
}


type SlbTemplatePolicyClassListLidList struct {
    Lidnum int `json:"lidnum"`
    ConnLimit int `json:"conn-limit"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnPer int `json:"conn-per"`
    RequestLimit int `json:"request-limit"`
    RequestRateLimit int `json:"request-rate-limit"`
    RequestPer int `json:"request-per"`
    BwRateLimit int `json:"bw-rate-limit"`
    BwPer int `json:"bw-per"`
    OverLimitAction int `json:"over-limit-action"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    Interval int `json:"interval"`
    DirectAction int `json:"direct-action"`
    DirectServiceGroup string `json:"direct-service-group"`
    DirectPbslbLogging int `json:"direct-pbslb-logging"`
    DirectPbslbInterval int `json:"direct-pbslb-interval" dval:"3"`
    DirectFail int `json:"direct-fail"`
    DirectActionValue string `json:"direct-action-value"`
    DirectLoggingDrpRst int `json:"direct-logging-drp-rst"`
    DirectActionInterval int `json:"direct-action-interval" dval:"3"`
    ResponseCodeRateLimit []SlbTemplatePolicyClassListLidListResponseCodeRateLimit `json:"response-code-rate-limit"`
    Dns64 SlbTemplatePolicyClassListLidListDns64 `json:"dns64"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SlbTemplatePolicyClassListLidListResponseCodeRateLimit struct {
    CodeRangeStart int `json:"code-range-start"`
    CodeRangeEnd int `json:"code-range-end"`
    Threshold int `json:"threshold"`
    Period int `json:"period"`
}


type SlbTemplatePolicyClassListLidListDns64 struct {
    Disable int `json:"disable"`
    ExclusiveAnswer int `json:"exclusive-answer"`
    Prefix string `json:"prefix"`
}

func (p *SlbTemplatePolicyClassList) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyClassList) getPath() string{
    return "slb/template/policy/"+p.Inst.Name+"/class-list"
}

func (p *SlbTemplatePolicyClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassList::Post")
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

func (p *SlbTemplatePolicyClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassList::Get")
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
func (p *SlbTemplatePolicyClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassList::Put")
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

func (p *SlbTemplatePolicyClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
