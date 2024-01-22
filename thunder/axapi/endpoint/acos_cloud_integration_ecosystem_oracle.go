

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosCloudIntegrationEcosystemOracle struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    CompartmentId string `json:"compartment-id"`
    Fingerprint string `json:"fingerprint"`
    HealthCheckInterval string `json:"health-check-interval"`
    HostName string `json:"host-name"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv6Address string `json:"ipv6-address"`
    Port int `json:"port"`
    PrivateKeyPath string `json:"private-key-path"`
    ServiceLabel []AcosCloudIntegrationEcosystemOracleServiceLabel `json:"service-label"`
    TenancyId string `json:"tenancy-id"`
    UserId string `json:"user-id"`
    Uuid string `json:"uuid"`

	} `json:"oracle"`
}


type AcosCloudIntegrationEcosystemOracleServiceLabel struct {
    ServiceLabelName string `json:"service-label-name"`
}

func (p *AcosCloudIntegrationEcosystemOracle) GetId() string{
    return "1"
}

func (p *AcosCloudIntegrationEcosystemOracle) getPath() string{
    return "acos-cloud-integration/ecosystem/oracle"
}

func (p *AcosCloudIntegrationEcosystemOracle) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemOracle::Post")
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

func (p *AcosCloudIntegrationEcosystemOracle) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemOracle::Get")
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
func (p *AcosCloudIntegrationEcosystemOracle) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemOracle::Put")
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

func (p *AcosCloudIntegrationEcosystemOracle) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosCloudIntegrationEcosystemOracle::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
