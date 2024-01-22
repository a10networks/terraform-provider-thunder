

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CloudServicesCloudProviderAwsLog struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    ActivePartitions string `json:"active-partitions"`
    LogGroupName string `json:"log-group-name"`
    Uuid string `json:"uuid"`

	} `json:"log"`
}

func (p *CloudServicesCloudProviderAwsLog) GetId() string{
    return "1"
}

func (p *CloudServicesCloudProviderAwsLog) getPath() string{
    return "cloud-services/cloud-provider/aws/log"
}

func (p *CloudServicesCloudProviderAwsLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAwsLog::Post")
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

func (p *CloudServicesCloudProviderAwsLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAwsLog::Get")
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
func (p *CloudServicesCloudProviderAwsLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAwsLog::Put")
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

func (p *CloudServicesCloudProviderAwsLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAwsLog::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
