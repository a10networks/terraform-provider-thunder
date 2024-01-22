

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NgWafCpuOper struct {
    
    Oper NgWafCpuOperOper `json:"oper"`

}
type DataNgWafCpuOper struct {
    DtNgWafCpuOper NgWafCpuOper `json:"cpu"`
}


type NgWafCpuOperOper struct {
    NumberOfCpus int `json:"number-of-cpus"`
    CpuInfo []NgWafCpuOperOperCpuInfo `json:"cpu-info"`
}


type NgWafCpuOperOperCpuInfo struct {
    CpuId int `json:"cpu-id"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
}

func (p *NgWafCpuOper) GetId() string{
    return "1"
}

func (p *NgWafCpuOper) getPath() string{
    return "ng-waf/cpu/oper"
}

func (p *NgWafCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNgWafCpuOper,error) {
logger.Println("NgWafCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNgWafCpuOper
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
