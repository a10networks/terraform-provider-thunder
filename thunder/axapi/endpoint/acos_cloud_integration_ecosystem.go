

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosCloudIntegrationEcosystem struct {
	Inst struct {

    Consul AcosCloudIntegrationEcosystemConsul41 `json:"consul"`
    Dummy int `json:"dummy"`
    K8s AcosCloudIntegrationEcosystemK8s43 `json:"k8s"`
    Oracle AcosCloudIntegrationEcosystemOracle45 `json:"oracle"`
    Uuid string `json:"uuid"`

	} `json:"ecosystem"`
}


type AcosCloudIntegrationEcosystemConsul41 struct {
    ServiceLabel []AcosCloudIntegrationEcosystemConsulServiceLabel42 `json:"service-label"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    HostName string `json:"host-name"`
    Port int `json:"port"`
    HealthCheckInterval string `json:"health-check-interval"`
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
}


type AcosCloudIntegrationEcosystemConsulServiceLabel42 struct {
    ServiceLabelName string `json:"service-label-name"`
}


type AcosCloudIntegrationEcosystemK8s43 struct {
    Action string `json:"action" dval:"disable"`
    HealthCheckInterval string `json:"health-check-interval"`
    ClusterConfigFile string `json:"cluster-config-file"`
    ServiceLabel []AcosCloudIntegrationEcosystemK8sServiceLabel44 `json:"service-label"`
    Uuid string `json:"uuid"`
}


type AcosCloudIntegrationEcosystemK8sServiceLabel44 struct {
    ServiceLabelName string `json:"service-label-name"`
}


type AcosCloudIntegrationEcosystemOracle45 struct {
    ServiceLabel []AcosCloudIntegrationEcosystemOracleServiceLabel46 `json:"service-label"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    HostName string `json:"host-name"`
    Port int `json:"port"`
    HealthCheckInterval string `json:"health-check-interval"`
    CompartmentId string `json:"compartment-id"`
    TenancyId string `json:"tenancy-id"`
    UserId string `json:"user-id"`
    Fingerprint string `json:"fingerprint"`
    PrivateKeyPath string `json:"private-key-path"`
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
}


type AcosCloudIntegrationEcosystemOracleServiceLabel46 struct {
    ServiceLabelName string `json:"service-label-name"`
}

func (p *AcosCloudIntegrationEcosystem) GetId() string{
    return "1"
}

func (p *AcosCloudIntegrationEcosystem) getPath() string{
    return "acos-cloud-integration/ecosystem"
}

func (p *AcosCloudIntegrationEcosystem) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystem::Post")
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

func (p *AcosCloudIntegrationEcosystem) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystem::Get")
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
func (p *AcosCloudIntegrationEcosystem) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystem::Put")
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

func (p *AcosCloudIntegrationEcosystem) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystem::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
