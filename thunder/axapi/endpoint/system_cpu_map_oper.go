

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuMapOper struct {
    
    Oper SystemCpuMapOperOper `json:"oper"`

}
type DataSystemCpuMapOper struct {
    DtSystemCpuMapOper SystemCpuMapOper `json:"cpu-map"`
}


type SystemCpuMapOperOper struct {
    SystemCpuMap []SystemCpuMapOperOperSystemCpuMap `json:"system-cpu-map"`
}


type SystemCpuMapOperOperSystemCpuMap struct {
    CpuIdx string `json:"CPU-IDX"`
    CpuType string `json:"cpu-type"`
}

func (p *SystemCpuMapOper) GetId() string{
    return "1"
}

func (p *SystemCpuMapOper) getPath() string{
    return "system/cpu-map/oper"
}

func (p *SystemCpuMapOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuMapOper,error) {
logger.Println("SystemCpuMapOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuMapOper
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
