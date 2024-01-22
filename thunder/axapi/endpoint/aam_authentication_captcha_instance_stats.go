

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationCaptchaInstanceStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationCaptchaInstanceStatsStats `json:"stats"`

}
type DataAamAuthenticationCaptchaInstanceStats struct {
    DtAamAuthenticationCaptchaInstanceStats AamAuthenticationCaptchaInstanceStats `json:"instance"`
}


type AamAuthenticationCaptchaInstanceStatsStats struct {
    Request int `json:"request"`
    VerifySucc int `json:"verify-succ"`
    ParseFail int `json:"parse-fail"`
    JsonFail int `json:"json-fail"`
    AttrFail int `json:"attr-fail"`
    TimeoutError int `json:"timeout-error"`
    OtherError int `json:"other-error"`
}

func (p *AamAuthenticationCaptchaInstanceStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationCaptchaInstanceStats) getPath() string{
    
    return "aam/authentication/captcha/instance/"+p.Name+"/stats"
}

func (p *AamAuthenticationCaptchaInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationCaptchaInstanceStats,error) {
logger.Println("AamAuthenticationCaptchaInstanceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationCaptchaInstanceStats
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
