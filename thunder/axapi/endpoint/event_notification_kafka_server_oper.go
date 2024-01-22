

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EventNotificationKafkaServerOper struct {
    
    Oper EventNotificationKafkaServerOperOper `json:"oper"`

}
type DataEventNotificationKafkaServerOper struct {
    DtEventNotificationKafkaServerOper EventNotificationKafkaServerOper `json:"server"`
}


type EventNotificationKafkaServerOperOper struct {
    KafkaBrokerState string `json:"kafka-broker-state"`
}

func (p *EventNotificationKafkaServerOper) GetId() string{
    return "1"
}

func (p *EventNotificationKafkaServerOper) getPath() string{
    return "event-notification/kafka/server/oper"
}

func (p *EventNotificationKafkaServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEventNotificationKafkaServerOper,error) {
logger.Println("EventNotificationKafkaServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEventNotificationKafkaServerOper
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
