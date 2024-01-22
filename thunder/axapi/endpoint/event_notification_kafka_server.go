

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EventNotificationKafkaServer struct {
	Inst struct {

    HostIpv4 string `json:"host-ipv4"`
    Port int `json:"port" dval:"9092"`
    SamplingEnable []EventNotificationKafkaServerSamplingEnable `json:"sampling-enable"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"server"`
}


type EventNotificationKafkaServerSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *EventNotificationKafkaServer) GetId() string{
    return "1"
}

func (p *EventNotificationKafkaServer) getPath() string{
    return "event-notification/kafka/server"
}

func (p *EventNotificationKafkaServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EventNotificationKafkaServer::Post")
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

func (p *EventNotificationKafkaServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EventNotificationKafkaServer::Get")
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
func (p *EventNotificationKafkaServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EventNotificationKafkaServer::Put")
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

func (p *EventNotificationKafkaServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EventNotificationKafkaServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
