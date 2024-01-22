

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemResourceAccountingOper struct {
    
    Oper SystemResourceAccountingOperOper `json:"oper"`

}
type DataSystemResourceAccountingOper struct {
    DtSystemResourceAccountingOper SystemResourceAccountingOper `json:"resource-accounting"`
}


type SystemResourceAccountingOperOper struct {
    Scope string `json:"scope"`
    PartitionResource []SystemResourceAccountingOperOperPartitionResource `json:"partition-resource"`
}


type SystemResourceAccountingOperOperPartitionResource struct {
    PartitionName string `json:"partition-name"`
    ResType []SystemResourceAccountingOperOperPartitionResourceResType `json:"res-type"`
}


type SystemResourceAccountingOperOperPartitionResourceResType struct {
    ResourceType string `json:"resource-type"`
    Resources []SystemResourceAccountingOperOperPartitionResourceResTypeResources `json:"resources"`
}


type SystemResourceAccountingOperOperPartitionResourceResTypeResources struct {
    ResourceName string `json:"resource-name"`
    Current string `json:"current"`
    Minimum string `json:"minimum"`
    Maximum string `json:"maximum"`
    Utilization string `json:"utilization"`
    MaxExceed string `json:"max-exceed"`
    ThresholdExceed string `json:"threshold-exceed"`
    Average string `json:"average"`
    Peak string `json:"peak"`
    Cap string `json:"cap"`
    CapUtilization string `json:"cap-utilization"`
}

func (p *SystemResourceAccountingOper) GetId() string{
    return "1"
}

func (p *SystemResourceAccountingOper) getPath() string{
    return "system/resource-accounting/oper"
}

func (p *SystemResourceAccountingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemResourceAccountingOper,error) {
logger.Println("SystemResourceAccountingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemResourceAccountingOper
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
