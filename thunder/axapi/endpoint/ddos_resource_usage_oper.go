

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosResourceUsageOper struct {
    
    Oper DdosResourceUsageOperOper `json:"oper"`

}
type DataDdosResourceUsageOper struct {
    DtDdosResourceUsageOper DdosResourceUsageOper `json:"resource-usage"`
}


type DdosResourceUsageOperOper struct {
    StaticResources []DdosResourceUsageOperOperStaticResources `json:"static-resources"`
    DynamicResources []DdosResourceUsageOperOperDynamicResources `json:"dynamic-resources"`
    PerResourceLimit []DdosResourceUsageOperOperPerResourceLimit `json:"per-resource-limit"`
    HwResourceLimit []DdosResourceUsageOperOperHwResourceLimit `json:"hw-resource-limit"`
    Enabled int `json:"enabled"`
    SystemCapacityRemaining int `json:"system-capacity-remaining"`
    SystemCapacityTotal int `json:"system-capacity-total"`
    SystemCapacityPercentage int `json:"system-capacity-percentage"`
    SystemCapacity int `json:"system-capacity"`
}


type DdosResourceUsageOperOperStaticResources struct {
    ResName string `json:"res-name"`
    ResAlloc int `json:"res-alloc"`
    ResLimit int `json:"res-limit"`
}


type DdosResourceUsageOperOperDynamicResources struct {
    ResName string `json:"res-name"`
    ResAlloc int `json:"res-alloc"`
    ResLimit int `json:"res-limit"`
}


type DdosResourceUsageOperOperPerResourceLimit struct {
    ResName string `json:"res-name"`
    ResAlloc string `json:"res-alloc"`
    ResLimit int `json:"res-limit"`
}


type DdosResourceUsageOperOperHwResourceLimit struct {
    ResName string `json:"res-name"`
    ResAlloc int `json:"res-alloc"`
    ResLimit int `json:"res-limit"`
}

func (p *DdosResourceUsageOper) GetId() string{
    return "1"
}

func (p *DdosResourceUsageOper) getPath() string{
    return "ddos/resource-usage/oper"
}

func (p *DdosResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosResourceUsageOper,error) {
logger.Println("DdosResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosResourceUsageOper
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
