

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAPreferredSessionSyncPortTrunk struct {
	Inst struct {

    PreTrunk int `json:"pre-trunk"`
    PreVlan int `json:"pre-vlan"`
    Uuid string `json:"uuid"`

	} `json:"trunk"`
}

func (p *VrrpAPreferredSessionSyncPortTrunk) GetId() string{
    return strconv.Itoa(p.Inst.PreTrunk)
}

func (p *VrrpAPreferredSessionSyncPortTrunk) getPath() string{
    return "vrrp-a/preferred-session-sync-port/trunk"
}

func (p *VrrpAPreferredSessionSyncPortTrunk) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortTrunk::Post")
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

func (p *VrrpAPreferredSessionSyncPortTrunk) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortTrunk::Get")
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
func (p *VrrpAPreferredSessionSyncPortTrunk) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortTrunk::Put")
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

func (p *VrrpAPreferredSessionSyncPortTrunk) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortTrunk::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
