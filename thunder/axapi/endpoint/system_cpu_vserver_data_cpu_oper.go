

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuVserverDataCpuOper struct {
    
    Oper SystemCpuVserverDataCpuOperOper `json:"oper"`

}
type DataSystemCpuVserverDataCpuOper struct {
    DtSystemCpuVserverDataCpuOper SystemCpuVserverDataCpuOper `json:"vserver-data-cpu"`
}


type SystemCpuVserverDataCpuOperOper struct {
    VserverCpuUsage []SystemCpuVserverDataCpuOperOperVserverCpuUsage `json:"vserver-cpu-usage"`
}


type SystemCpuVserverDataCpuOperOperVserverCpuUsage struct {
    VserverName string `json:"vserver-name"`
    CpuId int `json:"cpu-id"`
    DcpuStr string `json:"dcpu-str"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
}

func (p *SystemCpuVserverDataCpuOper) GetId() string{
    return "1"
}

func (p *SystemCpuVserverDataCpuOper) getPath() string{
    return "system-cpu/vserver-data-cpu/oper"
}

func (p *SystemCpuVserverDataCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuVserverDataCpuOper,error) {
logger.Println("SystemCpuVserverDataCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuVserverDataCpuOper
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
