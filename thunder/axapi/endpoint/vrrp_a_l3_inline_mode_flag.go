

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAL3InlineModeFlag struct {
	Inst struct {

    L3InlineMode int `json:"l3-inline-mode"`
    Uuid string `json:"uuid"`

	} `json:"l3-inline-mode-flag"`
}

func (p *VrrpAL3InlineModeFlag) GetId() string{
    return "1"
}

func (p *VrrpAL3InlineModeFlag) getPath() string{
    return "vrrp-a/l3-inline-mode-flag"
}

func (p *VrrpAL3InlineModeFlag) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL3InlineModeFlag::Post")
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

func (p *VrrpAL3InlineModeFlag) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL3InlineModeFlag::Get")
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
func (p *VrrpAL3InlineModeFlag) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL3InlineModeFlag::Put")
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

func (p *VrrpAL3InlineModeFlag) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL3InlineModeFlag::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
