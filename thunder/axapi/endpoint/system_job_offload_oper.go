

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemJobOffloadOper struct {
    
    Oper SystemJobOffloadOperOper `json:"oper"`

}
type DataSystemJobOffloadOper struct {
    DtSystemJobOffloadOper SystemJobOffloadOper `json:"job-offload"`
}


type SystemJobOffloadOperOper struct {
    JobOffloadCpuList []SystemJobOffloadOperOperJobOffloadCpuList `json:"job-offload-cpu-list"`
    CpuCount int `json:"cpu-count"`
    OffloadCpus int `json:"offload-cpus"`
}


type SystemJobOffloadOperOperJobOffloadCpuList struct {
    Jobs int `json:"jobs"`
    Submit int `json:"submit"`
    Receive int `json:"receive"`
    Execute int `json:"execute"`
    Snt_home int `json:"snt_home"`
    Rcv_home int `json:"rcv_home"`
    Complete int `json:"complete"`
    Fail_submit int `json:"fail_submit"`
    Q_no_space int `json:"q_no_space"`
    Fail_execute int `json:"fail_execute"`
    Fail_complete int `json:"fail_complete"`
}

func (p *SystemJobOffloadOper) GetId() string{
    return "1"
}

func (p *SystemJobOffloadOper) getPath() string{
    return "system/job-offload/oper"
}

func (p *SystemJobOffloadOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemJobOffloadOper,error) {
logger.Println("SystemJobOffloadOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemJobOffloadOper
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
