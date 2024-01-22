

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlanGlobal struct {
	Inst struct {

    EnableDefVlanL2Forwarding int `json:"enable-def-vlan-l2-forwarding"`
    L3VlanFwdDisable int `json:"l3-vlan-fwd-disable"`
    SamplingEnable []NetworkVlanGlobalSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"vlan-global"`
}


type NetworkVlanGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *NetworkVlanGlobal) GetId() string{
    return "1"
}

func (p *NetworkVlanGlobal) getPath() string{
    return "network/vlan-global"
}

func (p *NetworkVlanGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGlobal::Post")
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

func (p *NetworkVlanGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGlobal::Get")
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
func (p *NetworkVlanGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGlobal::Put")
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

func (p *NetworkVlanGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
