

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAForceSelfStandby struct {
	Inst struct {

    Action string `json:"action"`
    AllPartitions int `json:"all-partitions"`
    SkipCheck int `json:"skip-check"`
    Vrid int `json:"vrid"`

	} `json:"force-self-standby"`
}

func (p *VrrpAForceSelfStandby) GetId() string{
    return "1"
}

func (p *VrrpAForceSelfStandby) getPath() string{
    return "vrrp-a/force-self-standby"
}

func (p *VrrpAForceSelfStandby) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandby::Post")
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

func (p *VrrpAForceSelfStandby) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandby::Get")
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
func (p *VrrpAForceSelfStandby) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandby::Put")
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

func (p *VrrpAForceSelfStandby) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAForceSelfStandby::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
