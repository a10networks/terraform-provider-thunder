

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CloudServicesKafkaServerProfileOper struct {
    
    Oper CloudServicesKafkaServerProfileOperOper `json:"oper"`

}
type DataCloudServicesKafkaServerProfileOper struct {
    DtCloudServicesKafkaServerProfileOper CloudServicesKafkaServerProfileOper `json:"kafka-server-profile"`
}


type CloudServicesKafkaServerProfileOperOper struct {
    ActiveNamespace string `json:"active-namespace"`
    KafkaBrokerState string `json:"kafka-broker-state"`
    RegistrationStatus string `json:"registration-status"`
}

func (p *CloudServicesKafkaServerProfileOper) GetId() string{
    return "1"
}

func (p *CloudServicesKafkaServerProfileOper) getPath() string{
    return "cloud-services/kafka-server-profile/oper"
}

func (p *CloudServicesKafkaServerProfileOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCloudServicesKafkaServerProfileOper,error) {
logger.Println("CloudServicesKafkaServerProfileOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCloudServicesKafkaServerProfileOper
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
