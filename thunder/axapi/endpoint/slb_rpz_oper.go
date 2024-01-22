

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbRpzOper struct {
    
    Oper SlbRpzOperOper `json:"oper"`

}
type DataSlbRpzOper struct {
    DtSlbRpzOper SlbRpzOper `json:"rpz"`
}


type SlbRpzOperOper struct {
    Filter_entry string `json:"filter_entry"`
    FileList []SlbRpzOperOperFileList `json:"file-list"`
    RpzFileSizeMax int `json:"rpz-file-size-max"`
    RpzCount int `json:"rpz-count"`
}


type SlbRpzOperOperFileList struct {
    File string `json:"file"`
    Dns_template_bound string `json:"dns_template_bound"`
}

func (p *SlbRpzOper) GetId() string{
    return "1"
}

func (p *SlbRpzOper) getPath() string{
    return "slb/rpz/oper"
}

func (p *SlbRpzOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbRpzOper,error) {
logger.Println("SlbRpzOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbRpzOper
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
