

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkTwampResponderIp struct {
	Inst struct {

    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`

	} `json:"ip"`
}

func (p *NetworkTwampResponderIp) GetId() string{
    return "1"
}

func (p *NetworkTwampResponderIp) getPath() string{
    return "network/twamp/responder/ip"
}

func (p *NetworkTwampResponderIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponderIp::Post")
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

func (p *NetworkTwampResponderIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponderIp::Get")
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
func (p *NetworkTwampResponderIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponderIp::Put")
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

func (p *NetworkTwampResponderIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponderIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
