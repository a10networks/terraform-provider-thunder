

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewShowProcessOper struct {
    
    Oper SystemViewShowProcessOperOper `json:"oper"`

}
type DataSystemViewShowProcessOper struct {
    DtSystemViewShowProcessOper SystemViewShowProcessOper `json:"show-process"`
}


type SystemViewShowProcessOperOper struct {
    ProcInfo []SystemViewShowProcessOperOperProcInfo `json:"proc-info"`
}


type SystemViewShowProcessOperOperProcInfo struct {
    ProcData string `json:"proc-data"`
}

func (p *SystemViewShowProcessOper) GetId() string{
    return "1"
}

func (p *SystemViewShowProcessOper) getPath() string{
    return "system-view/show-process/oper"
}

func (p *SystemViewShowProcessOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewShowProcessOper,error) {
logger.Println("SystemViewShowProcessOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewShowProcessOper
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
