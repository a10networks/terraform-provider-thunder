

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpsProfileOper struct {
    
    Oper IpsProfileOperOper `json:"oper"`

}
type DataIpsProfileOper struct {
    DtIpsProfileOper IpsProfileOper `json:"profile"`
}


type IpsProfileOperOper struct {
    ProfileList []IpsProfileOperOperProfileList `json:"profile-list"`
    Name string `json:"name"`
    Sid int `json:"sid"`
    ProfileDetail int `json:"profile-detail"`
}


type IpsProfileOperOperProfileList struct {
    ProfileName string `json:"profile-name"`
    IsPredefined int `json:"is-predefined"`
    ReferenceCount int `json:"reference-count"`
    SignatureCount int `json:"signature-count"`
    SignatureList []IpsProfileOperOperProfileListSignatureList `json:"signature-list"`
}


type IpsProfileOperOperProfileListSignatureList struct {
    Sid int `json:"sid"`
    Enable int `json:"enable"`
    Action string `json:"action"`
    MissingSignature int `json:"missing-signature"`
    Message string `json:"message"`
}

func (p *IpsProfileOper) GetId() string{
    return "1"
}

func (p *IpsProfileOper) getPath() string{
    return "ips/profile/oper"
}

func (p *IpsProfileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpsProfileOper,error) {
logger.Println("IpsProfileOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpsProfileOper
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
