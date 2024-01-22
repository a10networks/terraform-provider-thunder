

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSessionAttrOper struct {
    
    Oper AamAuthenticationSessionAttrOperOper `json:"oper"`

}
type DataAamAuthenticationSessionAttrOper struct {
    DtAamAuthenticationSessionAttrOper AamAuthenticationSessionAttrOper `json:"session-attr"`
}


type AamAuthenticationSessionAttrOperOper struct {
    AttrList []AamAuthenticationSessionAttrOperOperAttrList `json:"attr-list"`
    Sid int `json:"sid"`
    Partition string `json:"partition"`
}


type AamAuthenticationSessionAttrOperOperAttrList struct {
    AttrName string `json:"attr-name"`
    AttrType string `json:"attr-type"`
    AttrValue string `json:"attr-value"`
}

func (p *AamAuthenticationSessionAttrOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationSessionAttrOper) getPath() string{
    return "aam/authentication/session-attr/oper"
}

func (p *AamAuthenticationSessionAttrOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationSessionAttrOper,error) {
logger.Println("AamAuthenticationSessionAttrOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationSessionAttrOper
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
