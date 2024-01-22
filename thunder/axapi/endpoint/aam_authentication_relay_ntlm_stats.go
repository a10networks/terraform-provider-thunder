

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayNtlmStats struct {
    
    Name string `json:"name"`
    Stats AamAuthenticationRelayNtlmStatsStats `json:"stats"`

}
type DataAamAuthenticationRelayNtlmStats struct {
    DtAamAuthenticationRelayNtlmStats AamAuthenticationRelayNtlmStats `json:"ntlm"`
}


type AamAuthenticationRelayNtlmStatsStats struct {
    Success int `json:"success"`
    Failure int `json:"failure"`
    Request int `json:"request"`
    Response int `json:"response"`
    HttpCode200 int `json:"http-code-200"`
    HttpCode400 int `json:"http-code-400"`
    HttpCode401 int `json:"http-code-401"`
    HttpCode403 int `json:"http-code-403"`
    HttpCode404 int `json:"http-code-404"`
    HttpCode500 int `json:"http-code-500"`
    HttpCode503 int `json:"http-code-503"`
    HttpCodeOther int `json:"http-code-other"`
    BufferAllocFail int `json:"buffer-alloc-fail"`
    EncodingFail int `json:"encoding-fail"`
    InsertHeaderFail int `json:"insert-header-fail"`
    ParseHeaderFail int `json:"parse-header-fail"`
    InternalError int `json:"internal-error"`
    NtlmAuthSkipped int `json:"ntlm-auth-skipped"`
    LargeRequestProcessing int `json:"large-request-processing"`
    LargeRequestFlushed int `json:"large-request-flushed"`
    HeadNegotiateRequestSent int `json:"head-negotiate-request-sent"`
    HeadAuthRequestSent int `json:"head-auth-request-sent"`
}

func (p *AamAuthenticationRelayNtlmStats) GetId() string{
    return "1"
}

func (p *AamAuthenticationRelayNtlmStats) getPath() string{
    
    return "aam/authentication/relay/ntlm/"+p.Name+"/stats"
}

func (p *AamAuthenticationRelayNtlmStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationRelayNtlmStats,error) {
logger.Println("AamAuthenticationRelayNtlmStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationRelayNtlmStats
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
