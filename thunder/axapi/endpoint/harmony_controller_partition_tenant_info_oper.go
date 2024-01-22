

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HarmonyControllerPartitionTenantInfoOper struct {
    
    Oper HarmonyControllerPartitionTenantInfoOperOper `json:"oper"`

}
type DataHarmonyControllerPartitionTenantInfoOper struct {
    DtHarmonyControllerPartitionTenantInfoOper HarmonyControllerPartitionTenantInfoOper `json:"partition-tenant-info"`
}


type HarmonyControllerPartitionTenantInfoOperOper struct {
    PartitionName string `json:"partition-name"`
    TenantName string `json:"tenant-name"`
    TenantId string `json:"tenant-id"`
    ClusterName string `json:"cluster-name"`
    ClusterId string `json:"cluster-id"`
    LogRatePerSec int `json:"log-rate-per-sec"`
}

func (p *HarmonyControllerPartitionTenantInfoOper) GetId() string{
    return "1"
}

func (p *HarmonyControllerPartitionTenantInfoOper) getPath() string{
    return "harmony-controller/partition-tenant-info/oper"
}

func (p *HarmonyControllerPartitionTenantInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataHarmonyControllerPartitionTenantInfoOper,error) {
logger.Println("HarmonyControllerPartitionTenantInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataHarmonyControllerPartitionTenantInfoOper
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
