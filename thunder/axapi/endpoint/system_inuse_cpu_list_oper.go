

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemInuseCpuListOper struct {
    
    Oper SystemInuseCpuListOperOper `json:"oper"`

}
type DataSystemInuseCpuListOper struct {
    DtSystemInuseCpuListOper SystemInuseCpuListOper `json:"inuse-cpu-list"`
}


type SystemInuseCpuListOperOper struct {
    SystemInuseCpuList []SystemInuseCpuListOperOperSystemInuseCpuList `json:"system-inuse-cpu-list"`
}


type SystemInuseCpuListOperOperSystemInuseCpuList struct {
    CpuNum string `json:"CPU-Num"`
    NumaNode string `json:"NUMA-NODE"`
    CpuId string `json:"CPU-ID"`
    HtCpu string `json:"HT-CPU"`
    Status string `json:"status"`
}

func (p *SystemInuseCpuListOper) GetId() string{
    return "1"
}

func (p *SystemInuseCpuListOper) getPath() string{
    return "system/inuse-cpu-list/oper"
}

func (p *SystemInuseCpuListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemInuseCpuListOper,error) {
logger.Println("SystemInuseCpuListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemInuseCpuListOper
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
