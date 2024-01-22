

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HarmonyControllerProfileTunnel struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"tunnel"`
}

func (p *HarmonyControllerProfileTunnel) GetId() string{
    return "1"
}

func (p *HarmonyControllerProfileTunnel) getPath() string{
    return "harmony-controller/profile/tunnel"
}

func (p *HarmonyControllerProfileTunnel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileTunnel::Post")
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

func (p *HarmonyControllerProfileTunnel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileTunnel::Get")
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
func (p *HarmonyControllerProfileTunnel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileTunnel::Put")
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

func (p *HarmonyControllerProfileTunnel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HarmonyControllerProfileTunnel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
