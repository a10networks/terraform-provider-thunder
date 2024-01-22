

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CloudServicesKafkaServerProfile struct {
	Inst struct {

    Action string `json:"action"`
    Active_ns string `json:"active_ns"`
    Alias_ns string `json:"alias_ns"`
    Bootstrap_servers string `json:"bootstrap_servers"`
    Client_id string `json:"client_id"`
    Client_secret string `json:"client_secret"`
    Primary_ns string `json:"primary_ns"`
    Resource_group string `json:"resource_group"`
    Sasl_mechanisms string `json:"sasl_mechanisms"`
    Sasl_password string `json:"sasl_password"`
    Schema_group_name string `json:"schema_group_name"`
    Secondary_ns string `json:"secondary_ns"`
    Security_protocol string `json:"security_protocol"`
    Subscription_id string `json:"subscription_id"`
    Tenant_id string `json:"tenant_id"`
    Uuid string `json:"uuid"`

	} `json:"kafka-server-profile"`
}

func (p *CloudServicesKafkaServerProfile) GetId() string{
    return "1"
}

func (p *CloudServicesKafkaServerProfile) getPath() string{
    return "cloud-services/kafka-server-profile"
}

func (p *CloudServicesKafkaServerProfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesKafkaServerProfile::Post")
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

func (p *CloudServicesKafkaServerProfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesKafkaServerProfile::Get")
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
func (p *CloudServicesKafkaServerProfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesKafkaServerProfile::Put")
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

func (p *CloudServicesKafkaServerProfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("CloudServicesKafkaServerProfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
