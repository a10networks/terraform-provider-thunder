

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationCaptchaStats struct {
    
    InstanceList []AamAuthenticationCaptchaStatsInstanceList `json:"instance-list"`
    Stats AamAuthenticationCaptchaStatsStats `json:"stats"`

}
type DataAamAuthenticationCaptchaStats struct {
    DtAamAuthenticationCaptchaStats AamAuthenticationCaptchaStats `json:"captcha"`
}


type AamAuthenticationCaptchaStatsInstanceList struct {
    Name string `json:"name"`
    Stats AamAuthenticationCaptchaStatsInstanceListStats `json:"stats"`
}


type AamAuthenticationCaptchaStatsInstanceListStats struct {
    Request int `json:"request"`
    VerifySucc int `json:"verify-succ"`
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
}


type AamAuthenticationCaptchaStatsStats struct {
    RequestNormal int `json:"request-normal"`
    RequestDropped int `json:"request-dropped"`
    ResponseSuccess int `json:"response-success"`
    ResponseFailure int `json:"response-failure"`
    ResponseError int `json:"response-error"`
    ResponseTimeout int `json:"response-timeout"`
    ResponseOther int `json:"response-other"`
    Request int `json:"request"`
    VerifySucc int `json:"verify-succ"`
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
    JobStartError int `json:"job-start-error"`
    PollingControlError int `json:"polling-control-error"`
}

func (p *AamAuthenticationCaptchaStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationCaptchaStats) getPath() string{
    return "aam/authentication/captcha/stats"
}

func (p *AamAuthenticationCaptchaStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationCaptchaStats,error) {
logger.Println("AamAuthenticationCaptchaStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationCaptchaStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
