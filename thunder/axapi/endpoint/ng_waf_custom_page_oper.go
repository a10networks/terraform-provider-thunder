

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWafCustomPageOper struct {
    
    Oper NgWafCustomPageOperOper `json:"oper"`

}
type DataNgWafCustomPageOper struct {
    DtNgWafCustomPageOper NgWafCustomPageOper `json:"custom-page"`
}


type NgWafCustomPageOperOper struct {
    FileList []NgWafCustomPageOperOperFileList `json:"file-list"`
}


type NgWafCustomPageOperOperFileList struct {
    File string `json:"file"`
    Size int `json:"size"`
}

func (p *NgWafCustomPageOper) GetId() string{
    return "1"
}

func (p *NgWafCustomPageOper) getPath() string{
    return "ng-waf/custom-page/oper"
}

func (p *NgWafCustomPageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafCustomPageOper,error) {
logger.Println("NgWafCustomPageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafCustomPageOper
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
