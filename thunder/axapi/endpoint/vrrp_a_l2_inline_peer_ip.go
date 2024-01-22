

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAL2InlinePeerIp struct {
	Inst struct {

    IpAddress string `json:"ip-address"`
    Uuid string `json:"uuid"`

	} `json:"l2-inline-peer-ip"`
}

func (p *VrrpAL2InlinePeerIp) GetId() string{
    return p.Inst.IpAddress
}

func (p *VrrpAL2InlinePeerIp) getPath() string{
    return "vrrp-a/l2-inline-peer-ip"
}

func (p *VrrpAL2InlinePeerIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL2InlinePeerIp::Post")
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

func (p *VrrpAL2InlinePeerIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL2InlinePeerIp::Get")
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
func (p *VrrpAL2InlinePeerIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL2InlinePeerIp::Put")
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

func (p *VrrpAL2InlinePeerIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAL2InlinePeerIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
