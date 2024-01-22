

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTsigOper struct {
    
    Oper SlbTsigOperOper `json:"oper"`

}
type DataSlbTsigOper struct {
    DtSlbTsigOper SlbTsigOper `json:"tsig"`
}


type SlbTsigOperOper struct {
    Filter_entry string `json:"filter_entry"`
    FileList []SlbTsigOperOperFileList `json:"file-list"`
    TsigCount int `json:"tsig-count"`
}


type SlbTsigOperOperFileList struct {
    File string `json:"file"`
}

func (p *SlbTsigOper) GetId() string{
    return "1"
}

func (p *SlbTsigOper) getPath() string{
    return "slb/tsig/oper"
}

func (p *SlbTsigOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTsigOper,error) {
logger.Println("SlbTsigOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTsigOper
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
