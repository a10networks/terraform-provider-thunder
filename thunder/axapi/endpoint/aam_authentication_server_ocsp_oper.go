

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationServerOcspOper struct {
    
    Oper AamAuthenticationServerOcspOperOper `json:"oper"`

}
type DataAamAuthenticationServerOcspOper struct {
    DtAamAuthenticationServerOcspOper AamAuthenticationServerOcspOper `json:"ocsp"`
}


type AamAuthenticationServerOcspOperOper struct {
    StatsClearType string `json:"stats-clear-type"`
    Name string `json:"name"`
}

func (p *AamAuthenticationServerOcspOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationServerOcspOper) getPath() string{
    return "aam/authentication/server/ocsp/oper"
}

func (p *AamAuthenticationServerOcspOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationServerOcspOper,error) {
logger.Println("AamAuthenticationServerOcspOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationServerOcspOper
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
