

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosCloudIntegrationEcosystemConsul struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    HealthCheckInterval string `json:"health-check-interval"`
    HostName string `json:"host-name"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    Port int `json:"port"`
    ServiceLabel []AcosCloudIntegrationEcosystemConsulServiceLabel `json:"service-label"`
    Uuid string `json:"uuid"`

	} `json:"consul"`
}


type AcosCloudIntegrationEcosystemConsulServiceLabel struct {
    ServiceLabelName string `json:"service-label-name"`
}

func (p *AcosCloudIntegrationEcosystemConsul) GetId() string{
    return "1"
}

func (p *AcosCloudIntegrationEcosystemConsul) getPath() string{
    return "acos-cloud-integration/ecosystem/consul"
}

func (p *AcosCloudIntegrationEcosystemConsul) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemConsul::Post")
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

func (p *AcosCloudIntegrationEcosystemConsul) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemConsul::Get")
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
func (p *AcosCloudIntegrationEcosystemConsul) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemConsul::Put")
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

func (p *AcosCloudIntegrationEcosystemConsul) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemConsul::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
