

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationFilePortalOper struct {
    
    Oper AamAuthenticationFilePortalOperOper `json:"oper"`

}
type DataAamAuthenticationFilePortalOper struct {
    DtAamAuthenticationFilePortalOper AamAuthenticationFilePortalOper `json:"portal"`
}


type AamAuthenticationFilePortalOperOper struct {
    FileList []AamAuthenticationFilePortalOperOperFileList `json:"file-list"`
    PortalName string `json:"portal-name"`
}


type AamAuthenticationFilePortalOperOperFileList struct {
    File string `json:"file"`
    Size int `json:"size"`
}

func (p *AamAuthenticationFilePortalOper) GetId() string{
    return "1"
}

func (p *AamAuthenticationFilePortalOper) getPath() string{
    return "aam/authentication/file/portal/oper"
}

func (p *AamAuthenticationFilePortalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthenticationFilePortalOper,error) {
logger.Println("AamAuthenticationFilePortalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthenticationFilePortalOper
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
