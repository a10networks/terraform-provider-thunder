

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosCloudIntegration struct {
	Inst struct {

    Dummy int `json:"dummy"`
    Ecosystem AcosCloudIntegrationEcosystem47 `json:"ecosystem"`
    Uuid string `json:"uuid"`

	} `json:"acos-cloud-integration"`
}


type AcosCloudIntegrationEcosystem47 struct {
    Dummy int `json:"dummy"`
    Uuid string `json:"uuid"`
    Consul AcosCloudIntegrationEcosystemConsul48 `json:"consul"`
    Oracle AcosCloudIntegrationEcosystemOracle50 `json:"oracle"`
    K8s AcosCloudIntegrationEcosystemK8s52 `json:"k8s"`
}


type AcosCloudIntegrationEcosystemConsul48 struct {
    ServiceLabel []AcosCloudIntegrationEcosystemConsulServiceLabel49 `json:"service-label"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    HostName string `json:"host-name"`
    Port int `json:"port"`
    HealthCheckInterval string `json:"health-check-interval"`
    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`
}


type AcosCloudIntegrationEcosystemConsulServiceLabel49 struct {
    ServiceLabelName string `json:"service-label-name"`
}


type AcosCloudIntegrationEcosystemOracle50 struct {
    ServiceLabel []AcosCloudIntegrationEcosystemOracleServiceLabel51 `json:"service-label"`
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


type AcosCloudIntegrationEcosystemOracleServiceLabel51 struct {
    ServiceLabelName string `json:"service-label-name"`
}


type AcosCloudIntegrationEcosystemK8s52 struct {
    Action string `json:"action" dval:"disable"`
    HealthCheckInterval string `json:"health-check-interval"`
    ClusterConfigFile string `json:"cluster-config-file"`
    ServiceLabel []AcosCloudIntegrationEcosystemK8sServiceLabel53 `json:"service-label"`
    Uuid string `json:"uuid"`
}


type AcosCloudIntegrationEcosystemK8sServiceLabel53 struct {
    ServiceLabelName string `json:"service-label-name"`
}

func (p *AcosCloudIntegration) GetId() string{
    return "1"
}

func (p *AcosCloudIntegration) getPath() string{
    return "acos-cloud-integration"
}

func (p *AcosCloudIntegration) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegration::Post")
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

func (p *AcosCloudIntegration) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegration::Get")
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
func (p *AcosCloudIntegration) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegration::Put")
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

func (p *AcosCloudIntegration) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegration::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
