

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerWindowsOper struct {
    
    Oper AamAuthenticationServerWindowsOperOper `json:"oper"`

}
type DataAamAuthenticationServerWindowsOper struct {
    DtAamAuthenticationServerWindowsOper AamAuthenticationServerWindowsOper `json:"windows"`
}


type AamAuthenticationServerWindowsOperOper struct {
    StatsClearType string `json:"stats-clear-type"`
    Name string `json:"name"`
}

func (p *AamAuthenticationServerWindowsOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerWindowsOper) getPath() string{
    return "aam/authentication/server/windows/oper"
}

func (p *AamAuthenticationServerWindowsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerWindowsOper,error) {
logger.Println("AamAuthenticationServerWindowsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerWindowsOper
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
