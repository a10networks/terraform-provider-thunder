

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpsSignatureOper struct {
    
    Oper IpsSignatureOperOper `json:"oper"`

}
type DataIpsSignatureOper struct {
    DtIpsSignatureOper IpsSignatureOper `json:"signature"`
}


type IpsSignatureOperOper struct {
    SignatureList []IpsSignatureOperOperSignatureList `json:"signature-list"`
    Sid int `json:"sid"`
    Category string `json:"category"`
    Type string `json:"type"`
}


type IpsSignatureOperOperSignatureList struct {
    Sid int `json:"sid"`
    DefaultAction string `json:"default-action"`
    Category string `json:"category"`
    Revision int `json:"revision"`
    IsEnabled int `json:"is-enabled"`
    IsCustom int `json:"is-custom"`
    Direction string `json:"direction"`
    Attack_target string `json:"attack_target"`
    Severity int `json:"severity"`
    Type string `json:"type"`
    CveNumber string `json:"cve-number"`
    Message string `json:"message"`
    PatternList []IpsSignatureOperOperSignatureListPatternList `json:"pattern-list"`
    ParameterList IpsSignatureOperOperSignatureListParameterList `json:"parameter-list"`
    Hash string `json:"hash"`
}


type IpsSignatureOperOperSignatureListPatternList struct {
    IsInvert int `json:"is-invert"`
    IgnoreCase int `json:"ignore-case"`
    HasSpaceBar int `json:"has-space-bar"`
    IsMultiLine int `json:"is-multi-line"`
    PatternType string `json:"pattern-type"`
    PatternString string `json:"pattern-string"`
}


type IpsSignatureOperOperSignatureListParameterList struct {
    FastPatternIndex int `json:"fast-pattern-index"`
    MaxPayloadLen int `json:"max-payload-len"`
    MinPayloadLen int `json:"min-payload-len"`
    MaxUriLen int `json:"max-uri-len"`
    MinUriLen int `json:"min-uri-len"`
    HttpStatusCode int `json:"http-status-code"`
    HttpMethod string `json:"http-method"`
    DnsQuery int `json:"dns-query"`
    Threshold IpsSignatureOperOperSignatureListParameterListThreshold `json:"threshold"`
}


type IpsSignatureOperOperSignatureListParameterListThreshold struct {
    Count1 int `json:"count1"`
    Interval int `json:"interval"`
    TrackDirection string `json:"track-direction"`
    DetectionFilter int `json:"detection-filter"`
}

func (p *IpsSignatureOper) GetId() string{
    return "1"
}

func (p *IpsSignatureOper) getPath() string{
    return "ips/signature/oper"
}

func (p *IpsSignatureOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpsSignatureOper,error) {
logger.Println("IpsSignatureOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpsSignatureOper
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
