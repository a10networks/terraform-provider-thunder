

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCoreOper struct {
    
    Oper SystemCoreOperOper `json:"oper"`

}
type DataSystemCoreOper struct {
    DtSystemCoreOper SystemCoreOper `json:"core"`
}


type SystemCoreOperOper struct {
    FileList []SystemCoreOperOperFileList `json:"file-list"`
}


type SystemCoreOperOperFileList struct {
    File string `json:"file"`
    UpdateTime string `json:"update-time"`
    Size int `json:"size"`
}

func (p *SystemCoreOper) GetId() string{
    return "1"
}

func (p *SystemCoreOper) getPath() string{
    return "system/core/oper"
}

func (p *SystemCoreOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCoreOper,error) {
logger.Println("SystemCoreOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCoreOper
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
