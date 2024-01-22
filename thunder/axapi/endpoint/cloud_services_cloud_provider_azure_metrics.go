

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CloudServicesCloudProviderAzureMetrics struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    ActivePartitions string `json:"active-partitions"`
    Cps string `json:"cps" dval:"disable"`
    Cpu string `json:"cpu" dval:"disable"`
    Disk string `json:"disk" dval:"disable"`
    Interfaces string `json:"interfaces" dval:"disable"`
    Memory string `json:"memory" dval:"disable"`
    PacketDrop string `json:"packet-drop" dval:"disable"`
    PacketRate string `json:"packet-rate" dval:"disable"`
    ResourceId string `json:"resource-id"`
    ServerDownCount string `json:"server-down-count" dval:"disable"`
    ServerDownPercentage string `json:"server-down-percentage" dval:"disable"`
    ServerError string `json:"server-error" dval:"disable"`
    Sessions string `json:"sessions" dval:"disable"`
    SslCert string `json:"ssl-cert" dval:"disable"`
    Throughput string `json:"throughput" dval:"disable"`
    Tps string `json:"tps" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"metrics"`
}

func (p *CloudServicesCloudProviderAzureMetrics) GetId() string{
    return "1"
}

func (p *CloudServicesCloudProviderAzureMetrics) getPath() string{
    return "cloud-services/cloud-provider/azure/metrics"
}

func (p *CloudServicesCloudProviderAzureMetrics) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAzureMetrics::Post")
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

func (p *CloudServicesCloudProviderAzureMetrics) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAzureMetrics::Get")
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
func (p *CloudServicesCloudProviderAzureMetrics) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAzureMetrics::Put")
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

func (p *CloudServicesCloudProviderAzureMetrics) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesCloudProviderAzureMetrics::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
