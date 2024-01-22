

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosNetworkObjectNotification struct {
	Inst struct {

    Configuration string `json:"configuration"`
    Notification []DdosNetworkObjectNotificationNotification `json:"notification"`
    Uuid string `json:"uuid"`
    ObjectName string 

	} `json:"notification"`
}


type DdosNetworkObjectNotificationNotification struct {
    NotificationTemplateName string `json:"notification-template-name"`
}

func (p *DdosNetworkObjectNotification) GetId() string{
    return "1"
}

func (p *DdosNetworkObjectNotification) getPath() string{
    return "ddos/network-object/" +p.Inst.ObjectName + "/notification"
}

func (p *DdosNetworkObjectNotification) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectNotification::Post")
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

func (p *DdosNetworkObjectNotification) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectNotification::Get")
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
func (p *DdosNetworkObjectNotification) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectNotification::Put")
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

func (p *DdosNetworkObjectNotification) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosNetworkObjectNotification::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
