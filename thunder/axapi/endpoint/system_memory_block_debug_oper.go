

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemMemoryBlockDebugOper struct {
    
    Oper SystemMemoryBlockDebugOperOper `json:"oper"`

}
type DataSystemMemoryBlockDebugOper struct {
    DtSystemMemoryBlockDebugOper SystemMemoryBlockDebugOper `json:"memory-block-debug"`
}


type SystemMemoryBlockDebugOperOper struct {
    A10MemIdMax int `json:"a10-mem-id-max"`
    BlockDebugArrZero string `json:"block-debug-arr-zero"`
    BlockDebugArrOne string `json:"block-debug-arr-one"`
    BlockDebugArrTwo string `json:"block-debug-arr-two"`
    BlockDebugArrThree string `json:"block-debug-arr-three"`
    AllEverything []SystemMemoryBlockDebugOperOperAllEverything `json:"all-everything"`
    ThreadId int `json:"thread-id"`
    A10Mem []SystemMemoryBlockDebugOperOperA10Mem `json:"a10-mem"`
}


type SystemMemoryBlockDebugOperOperAllEverything struct {
    Id1 int `json:"id1"`
    ModuleName string `json:"module-name"`
    AllocZero int `json:"alloc-zero"`
    AllocOne int `json:"alloc-one"`
    AllocTwo int `json:"alloc-two"`
    AllocThree int `json:"alloc-three"`
    FreeFour int `json:"free-four"`
    FreeFive int `json:"free-five"`
    FreeSix int `json:"free-six"`
    FreeSeven int `json:"free-seven"`
    BlockDebugAllocOver int `json:"block-debug-alloc-over"`
    BlockDebugFreeOver int `json:"block-debug-free-over"`
    BlockDebugAllocMaxOver int `json:"block-debug-alloc-max-over"`
    DebugDecisionOverSize int `json:"debug-decision-over-size"`
    DebugDecisionAlloc int `json:"debug-decision-alloc"`
    DebugDecisionFree int `json:"debug-decision-free"`
}


type SystemMemoryBlockDebugOperOperA10Mem struct {
    TId int `json:"t-id"`
    TModuleName string `json:"t-module-name"`
    TAllocZero int `json:"t-alloc-zero"`
    TAllocOne int `json:"t-alloc-one"`
    TAllocTwo int `json:"t-alloc-two"`
    TAllocThree int `json:"t-alloc-three"`
    TFreeFour int `json:"t-free-four"`
    TFreeFive int `json:"t-free-five"`
    TFreeSix int `json:"t-free-six"`
    TFreeSeven int `json:"t-free-seven"`
    TBlockDebugAllocOver int `json:"t-block-debug-alloc-over"`
    TBlockDebugFreeOver int `json:"t-block-debug-free-over"`
    TBlockDebugAllocMaxOver int `json:"t-block-debug-alloc-max-over"`
    TDebugDecisionOverSize int `json:"t-debug-decision-over-size"`
    TDebugDecisionAlloc int `json:"t-debug-decision-alloc"`
    TDebugDecisionFree int `json:"t-debug-decision-free"`
}

func (p *SystemMemoryBlockDebugOper) GetId() string{
    return "1"
}

func (p *SystemMemoryBlockDebugOper) getPath() string{
    return "system/memory-block-debug/oper"
}

func (p *SystemMemoryBlockDebugOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemMemoryBlockDebugOper,error) {
logger.Println("SystemMemoryBlockDebugOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemMemoryBlockDebugOper
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
