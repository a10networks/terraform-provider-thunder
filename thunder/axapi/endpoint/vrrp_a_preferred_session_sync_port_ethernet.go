

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAPreferredSessionSyncPortEthernet struct {
	Inst struct {

    PreEth int `json:"pre-eth"`
    PreVlan int `json:"pre-vlan"`
    Uuid string `json:"uuid"`

	} `json:"ethernet"`
}

func (p *VrrpAPreferredSessionSyncPortEthernet) GetId() string{
    return strconv.Itoa(p.Inst.PreEth)
}

func (p *VrrpAPreferredSessionSyncPortEthernet) getPath() string{
    return "vrrp-a/preferred-session-sync-port/ethernet"
}

func (p *VrrpAPreferredSessionSyncPortEthernet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortEthernet::Post")
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

func (p *VrrpAPreferredSessionSyncPortEthernet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortEthernet::Get")
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
func (p *VrrpAPreferredSessionSyncPortEthernet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortEthernet::Put")
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

func (p *VrrpAPreferredSessionSyncPortEthernet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAPreferredSessionSyncPortEthernet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
