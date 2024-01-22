

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HarmonyControllerProfileThunderMgmtIp struct {
	Inst struct {

    IpAddress string `json:"ip-address"`
    Ipv6Addr string `json:"ipv6-addr"`
    Uuid string `json:"uuid"`

	} `json:"thunder-mgmt-ip"`
}

func (p *HarmonyControllerProfileThunderMgmtIp) GetId() string{
    return "1"
}

func (p *HarmonyControllerProfileThunderMgmtIp) getPath() string{
    return "harmony-controller/profile/thunder-mgmt-ip"
}

func (p *HarmonyControllerProfileThunderMgmtIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileThunderMgmtIp::Post")
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

func (p *HarmonyControllerProfileThunderMgmtIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileThunderMgmtIp::Get")
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
func (p *HarmonyControllerProfileThunderMgmtIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileThunderMgmtIp::Put")
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

func (p *HarmonyControllerProfileThunderMgmtIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileThunderMgmtIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
