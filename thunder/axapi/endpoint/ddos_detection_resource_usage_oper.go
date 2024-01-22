

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionResourceUsageOper struct {
    
    Oper DdosDetectionResourceUsageOperOper `json:"oper"`

}
type DataDdosDetectionResourceUsageOper struct {
    DtDdosDetectionResourceUsageOper DdosDetectionResourceUsageOper `json:"resource-usage"`
}


type DdosDetectionResourceUsageOperOper struct {
    StaticResources []DdosDetectionResourceUsageOperOperStaticResources `json:"static-resources"`
    DynamicResources []DdosDetectionResourceUsageOperOperDynamicResources `json:"dynamic-resources"`
}


type DdosDetectionResourceUsageOperOperStaticResources struct {
    ResName string `json:"res-name"`
    ResAlloc int `json:"res-alloc"`
    ResLimit int `json:"res-limit"`
}


type DdosDetectionResourceUsageOperOperDynamicResources struct {
    ResName string `json:"res-name"`
    ResAlloc int `json:"res-alloc"`
    ResLimit int `json:"res-limit"`
}

func (p *DdosDetectionResourceUsageOper) GetId() string{
    return "1"
}

func (p *DdosDetectionResourceUsageOper) getPath() string{
    return "ddos/detection/resource-usage/oper"
}

func (p *DdosDetectionResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDetectionResourceUsageOper,error) {
logger.Println("DdosDetectionResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDetectionResourceUsageOper
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
