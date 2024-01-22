

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkTwampResponder struct {
	Inst struct {

    EnableBothIpIpv6 int `json:"enable-both-ip-ipv6"`
    EnableIp int `json:"enable-ip"`
    EnableIpv6 int `json:"enable-ipv6"`
    Ip NetworkTwampResponderIp1074 `json:"ip"`
    Ipv6 NetworkTwampResponderIpv61075 `json:"ipv6"`
    Port int `json:"port"`
    Uuid string `json:"uuid"`

	} `json:"responder"`
}


type NetworkTwampResponderIp1074 struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
    Uuid string `json:"uuid"`
}


type NetworkTwampResponderIpv61075 struct {
    V6AclName string `json:"v6-acl-name"`
    Uuid string `json:"uuid"`
}

func (p *NetworkTwampResponder) GetId() string{
    return "1"
}

func (p *NetworkTwampResponder) getPath() string{
    return "network/twamp/responder"
}

func (p *NetworkTwampResponder) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponder::Post")
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

func (p *NetworkTwampResponder) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponder::Get")
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
func (p *NetworkTwampResponder) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponder::Put")
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

func (p *NetworkTwampResponder) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkTwampResponder::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
