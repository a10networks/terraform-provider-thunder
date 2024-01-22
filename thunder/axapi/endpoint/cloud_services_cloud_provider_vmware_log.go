

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CloudServicesCloudProviderVmwareLog struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    ActivePartitions string `json:"active-partitions"`
    Uuid string `json:"uuid"`
    VrliHost string `json:"vrli-host"`

	} `json:"log"`
}

func (p *CloudServicesCloudProviderVmwareLog) GetId() string{
    return "1"
}

func (p *CloudServicesCloudProviderVmwareLog) getPath() string{
    return "cloud-services/cloud-provider/vmware/log"
}

func (p *CloudServicesCloudProviderVmwareLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderVmwareLog::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *CloudServicesCloudProviderVmwareLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderVmwareLog::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *CloudServicesCloudProviderVmwareLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderVmwareLog::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *CloudServicesCloudProviderVmwareLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderVmwareLog::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
