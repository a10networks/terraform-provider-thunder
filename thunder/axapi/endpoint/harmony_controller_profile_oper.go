

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HarmonyControllerProfileOper struct {
    
    Oper HarmonyControllerProfileOperOper `json:"oper"`

}
type DataHarmonyControllerProfileOper struct {
    DtHarmonyControllerProfileOper HarmonyControllerProfileOper `json:"profile"`
}


type HarmonyControllerProfileOperOper struct {
    OverallStatus string `json:"overall-status"`
    HeartbeatStatus string `json:"heartbeat-status"`
    HeartbeatErrorMessage string `json:"heartbeat-error-message"`
    ServiceRegistry string `json:"service-registry"`
    ServiceRegistryErrorMessage string `json:"service-registry-error-message"`
    RegistrationStatus string `json:"registration-status"`
    RegistrationStatusCode int `json:"registration-status-code"`
    RegistrationErrorMessage string `json:"registration-error-message"`
    DeregistrationStatus string `json:"deregistration-status"`
    DeregistrationStatusCode int `json:"deregistration-status-code"`
    DeregistrationErrorMessage string `json:"deregistration-error-message"`
    SchemaRegistryStatus string `json:"schema-registry-status"`
    Broker_info string `json:"broker_info"`
    KafkaBrokerState string `json:"kafka-broker-state"`
    NumberOfTenantMappedPartitions int `json:"Number-of-tenant-mapped-partitions"`
    NumberOfTenantUnmappedPartitions int `json:"Number-of-tenant-unmapped-partitions"`
    TunnelStatus string `json:"tunnel-status"`
    TunnelErrorMessage string `json:"tunnel-error-message"`
    PeerDeviceInfo string `json:"peer-device-info"`
}

func (p *HarmonyControllerProfileOper) GetId() string{
    return "1"
}

func (p *HarmonyControllerProfileOper) getPath() string{
    return "harmony-controller/profile/oper"
}

func (p *HarmonyControllerProfileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataHarmonyControllerProfileOper,error) {
logger.Println("HarmonyControllerProfileOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataHarmonyControllerProfileOper
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
