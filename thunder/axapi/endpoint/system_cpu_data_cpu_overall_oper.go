

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuDataCpuOverallOper struct {
    
    Oper SystemCpuDataCpuOverallOperOper `json:"oper"`

}
type DataSystemCpuDataCpuOverallOper struct {
    DtSystemCpuDataCpuOverallOper SystemCpuDataCpuOverallOper `json:"data-cpu-overall"`
}


type SystemCpuDataCpuOverallOperOper struct {
    NumberOfCpu int `json:"number-of-cpu"`
    NumberOfDataCpu int `json:"number-of-data-cpu"`
    CpuUsage []SystemCpuDataCpuOverallOperOperCpuUsage `json:"cpu-usage"`
}


type SystemCpuDataCpuOverallOperOperCpuUsage struct {
    CpuId int `json:"cpu-id"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
    DcpuStr string `json:"dcpu-str"`
}

func (p *SystemCpuDataCpuOverallOper) GetId() string{
    return "1"
}

func (p *SystemCpuDataCpuOverallOper) getPath() string{
    return "system-cpu/data-cpu-overall/oper"
}

func (p *SystemCpuDataCpuOverallOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuDataCpuOverallOper,error) {
logger.Println("SystemCpuDataCpuOverallOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuDataCpuOverallOper
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
