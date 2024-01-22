

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosCloudIntegrationEcosystemK8s struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    ClusterConfigFile string `json:"cluster-config-file"`
    HealthCheckInterval string `json:"health-check-interval"`
    ServiceLabel []AcosCloudIntegrationEcosystemK8sServiceLabel `json:"service-label"`
    Uuid string `json:"uuid"`

	} `json:"k8s"`
}


type AcosCloudIntegrationEcosystemK8sServiceLabel struct {
    ServiceLabelName string `json:"service-label-name"`
}

func (p *AcosCloudIntegrationEcosystemK8s) GetId() string{
    return "1"
}

func (p *AcosCloudIntegrationEcosystemK8s) getPath() string{
    return "acos-cloud-integration/ecosystem/k8s"
}

func (p *AcosCloudIntegrationEcosystemK8s) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemK8s::Post")
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

func (p *AcosCloudIntegrationEcosystemK8s) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemK8s::Get")
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
func (p *AcosCloudIntegrationEcosystemK8s) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemK8s::Put")
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

func (p *AcosCloudIntegrationEcosystemK8s) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemK8s::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
