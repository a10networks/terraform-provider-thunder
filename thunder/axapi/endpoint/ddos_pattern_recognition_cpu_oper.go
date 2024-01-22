

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosPatternRecognitionCpuOper struct {
    
    Oper DdosPatternRecognitionCpuOperOper `json:"oper"`

}
type DataDdosPatternRecognitionCpuOper struct {
    DtDdosPatternRecognitionCpuOper DdosPatternRecognitionCpuOper `json:"cpu"`
}


type DdosPatternRecognitionCpuOperOper struct {
    NumberOfCpus int `json:"number-of-cpus"`
    CpuInfo []DdosPatternRecognitionCpuOperOperCpuInfo `json:"cpu-info"`
}


type DdosPatternRecognitionCpuOperOperCpuInfo struct {
    CpuId int `json:"cpu-id"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
}

func (p *DdosPatternRecognitionCpuOper) GetId() string{
    return "1"
}

func (p *DdosPatternRecognitionCpuOper) getPath() string{
    return "ddos/pattern-recognition/cpu/oper"
}

func (p *DdosPatternRecognitionCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosPatternRecognitionCpuOper,error) {
logger.Println("DdosPatternRecognitionCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosPatternRecognitionCpuOper
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
