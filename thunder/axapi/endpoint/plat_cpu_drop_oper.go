

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PlatCpuDropOper struct {
    
    Oper PlatCpuDropOperOper `json:"oper"`

}
type DataPlatCpuDropOper struct {
    DtPlatCpuDropOper PlatCpuDropOper `json:"plat-cpu-drop"`
}


type PlatCpuDropOperOper struct {
    FpgaSeg []PlatCpuDropOperOperFpgaSeg `json:"fpga-seg"`
    DropSeg []PlatCpuDropOperOperDropSeg `json:"drop-seg"`
    RateLimit int `json:"rate-limit"`
    RateLimitDrp []PlatCpuDropOperOperRateLimitDrp `json:"rate-limit-drp"`
}


type PlatCpuDropOperOperFpgaSeg struct {
    FpgaSegName string `json:"fpga-seg-name"`
}


type PlatCpuDropOperOperDropSeg struct {
    DropName string `json:"drop-name"`
    DropCnt []PlatCpuDropOperOperDropSegDropCnt `json:"drop-cnt"`
}


type PlatCpuDropOperOperDropSegDropCnt struct {
    DropCount string `json:"drop-count"`
}


type PlatCpuDropOperOperRateLimitDrp struct {
    RateLimitDrop string `json:"rate-limit-drop"`
}

func (p *PlatCpuDropOper) GetId() string{
    return "1"
}

func (p *PlatCpuDropOper) getPath() string{
    return "plat-cpu-drop/oper"
}

func (p *PlatCpuDropOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPlatCpuDropOper,error) {
logger.Println("PlatCpuDropOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataPlatCpuDropOper
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
