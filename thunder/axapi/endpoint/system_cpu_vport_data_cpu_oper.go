

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuVportDataCpuOper struct {
    
    Oper SystemCpuVportDataCpuOperOper `json:"oper"`

}
type DataSystemCpuVportDataCpuOper struct {
    DtSystemCpuVportDataCpuOper SystemCpuVportDataCpuOper `json:"vport-data-cpu"`
}


type SystemCpuVportDataCpuOperOper struct {
    VportCpuUsage []SystemCpuVportDataCpuOperOperVportCpuUsage `json:"vport-cpu-usage"`
}


type SystemCpuVportDataCpuOperOperVportCpuUsage struct {
    VserverName string `json:"vserver-name"`
    Portnumber int `json:"portNumber"`
    Protocol string `json:"protocol"`
    CpuId int `json:"cpu-id"`
    VportType string `json:"vport-type"`
    DcpuStr string `json:"dcpu-str"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
}

func (p *SystemCpuVportDataCpuOper) GetId() string{
    return "1"
}

func (p *SystemCpuVportDataCpuOper) getPath() string{
    return "system-cpu/vport-data-cpu/oper"
}

func (p *SystemCpuVportDataCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuVportDataCpuOper,error) {
logger.Println("SystemCpuVportDataCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuVportDataCpuOper
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
