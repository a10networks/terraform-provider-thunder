

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuListOper struct {
    
    Oper SystemCpuListOperOper `json:"oper"`

}
type DataSystemCpuListOper struct {
    DtSystemCpuListOper SystemCpuListOper `json:"cpu-list"`
}


type SystemCpuListOperOper struct {
    SystemCpuList []SystemCpuListOperOperSystemCpuList `json:"system-cpu-list"`
}


type SystemCpuListOperOperSystemCpuList struct {
    CpuIdx string `json:"CPU-IDX"`
    CpuNum string `json:"CPU-Num"`
    NumaNode string `json:"NUMA-NODE"`
    HtCpu string `json:"HT-CPU"`
    Status string `json:"status"`
}

func (p *SystemCpuListOper) GetId() string{
    return "1"
}

func (p *SystemCpuListOper) getPath() string{
    return "system/cpu-list/oper"
}

func (p *SystemCpuListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuListOper,error) {
logger.Println("SystemCpuListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuListOper
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
