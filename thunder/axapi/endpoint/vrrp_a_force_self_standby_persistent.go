

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAForceSelfStandbyPersistent struct {
	Inst struct {

    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"force-self-standby-persistent"`
}

func (p *VrrpAForceSelfStandbyPersistent) GetId() string{
    return strconv.Itoa(p.Inst.Vrid)
}

func (p *VrrpAForceSelfStandbyPersistent) getPath() string{
    return "vrrp-a/force-self-standby-persistent"
}

func (p *VrrpAForceSelfStandbyPersistent) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandbyPersistent::Post")
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

func (p *VrrpAForceSelfStandbyPersistent) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandbyPersistent::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *VrrpAForceSelfStandbyPersistent) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandbyPersistent::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *VrrpAForceSelfStandbyPersistent) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandbyPersistent::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
