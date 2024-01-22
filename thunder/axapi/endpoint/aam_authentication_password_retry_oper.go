

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationPasswordRetryOper struct {
    
    Oper AamAuthenticationPasswordRetryOperOper `json:"oper"`

}
type DataAamAuthenticationPasswordRetryOper struct {
    DtAamAuthenticationPasswordRetryOper AamAuthenticationPasswordRetryOper `json:"password-retry"`
}


type AamAuthenticationPasswordRetryOperOper struct {
    EntryList []AamAuthenticationPasswordRetryOperOperEntryList `json:"entry-list"`
    LogonName string `json:"logon-name"`
    Name string `json:"name"`
}


type AamAuthenticationPasswordRetryOperOperEntryList struct {
    Account string `json:"account"`
    Logon string `json:"logon"`
    PwFailure int `json:"pw-failure"`
    Ttl int `json:"ttl"`
    LockedOut string `json:"locked-out"`
}

func (p *AamAuthenticationPasswordRetryOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationPasswordRetryOper) getPath() string{
    return "aam/authentication/password-retry/oper"
}

func (p *AamAuthenticationPasswordRetryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationPasswordRetryOper,error) {
logger.Println("AamAuthenticationPasswordRetryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationPasswordRetryOper
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
