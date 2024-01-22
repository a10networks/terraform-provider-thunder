

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuCtrlCpuOper struct {
    
    Oper SystemCpuCtrlCpuOperOper `json:"oper"`

}
type DataSystemCpuCtrlCpuOper struct {
    DtSystemCpuCtrlCpuOper SystemCpuCtrlCpuOper `json:"ctrl-cpu"`
}


type SystemCpuCtrlCpuOperOper struct {
    CurrentTime string `json:"current-time"`
    NumberOfCpu int `json:"number-of-cpu"`
    CpuUsage []SystemCpuCtrlCpuOperOperCpuUsage `json:"cpu-usage"`
}


type SystemCpuCtrlCpuOperOperCpuUsage struct {
    CpuId int `json:"cpu-id"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
}

func (p *SystemCpuCtrlCpuOper) GetId() string{
    return "1"
}

func (p *SystemCpuCtrlCpuOper) getPath() string{
    return "system-cpu/ctrl-cpu/oper"
}

func (p *SystemCpuCtrlCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuCtrlCpuOper,error) {
logger.Println("SystemCpuCtrlCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuCtrlCpuOper
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
