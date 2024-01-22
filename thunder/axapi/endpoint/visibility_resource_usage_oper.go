

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityResourceUsageOper struct {
    
    Oper VisibilityResourceUsageOperOper `json:"oper"`

}
type DataVisibilityResourceUsageOper struct {
    DtVisibilityResourceUsageOper VisibilityResourceUsageOper `json:"resource-usage"`
}


type VisibilityResourceUsageOperOper struct {
    ResourceList []VisibilityResourceUsageOperOperResourceList `json:"resource-list"`
}


type VisibilityResourceUsageOperOperResourceList struct {
    ResourceName string `json:"resource-name"`
    ResourceCurrent int `json:"resource-current"`
    ResourceLimit int `json:"resource-limit"`
}

func (p *VisibilityResourceUsageOper) GetId() string{
    return "1"
}

func (p *VisibilityResourceUsageOper) getPath() string{
    return "visibility/resource-usage/oper"
}

func (p *VisibilityResourceUsageOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityResourceUsageOper,error) {
logger.Println("VisibilityResourceUsageOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityResourceUsageOper
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
