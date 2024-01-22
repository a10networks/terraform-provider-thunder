

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionDdosScriptOper struct {
    
    Oper DdosDetectionDdosScriptOperOper `json:"oper"`

}
type DataDdosDetectionDdosScriptOper struct {
    DtDdosDetectionDdosScriptOper DdosDetectionDdosScriptOper `json:"ddos-script"`
}


type DdosDetectionDdosScriptOperOper struct {
    FileList []DdosDetectionDdosScriptOperOperFileList `json:"file-list"`
    TotalRecords int `json:"total-records"`
}


type DdosDetectionDdosScriptOperOperFileList struct {
    File string `json:"file"`
    ReferenceCount int `json:"reference-count"`
    FileSize int `json:"file-size"`
}

func (p *DdosDetectionDdosScriptOper) GetId() string{
    return "1"
}

func (p *DdosDetectionDdosScriptOper) getPath() string{
    return "ddos/detection/ddos-script/oper"
}

func (p *DdosDetectionDdosScriptOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDetectionDdosScriptOper,error) {
logger.Println("DdosDetectionDdosScriptOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDetectionDdosScriptOper
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
