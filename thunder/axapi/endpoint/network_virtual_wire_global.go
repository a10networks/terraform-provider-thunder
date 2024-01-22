

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireGlobal struct {
	Inst struct {

    SamplingEnable []NetworkVirtualWireGlobalSamplingEnable `json:"sampling-enable"`
    SrcMacLearning int `json:"src-mac-learning"`
    UpdateActiveVlan string `json:"update-active-vlan" dval:"l3-packet"`
    Uuid string `json:"uuid"`
    VlanUpdatePeriod int `json:"vlan-update-period" dval:"30"`

	} `json:"virtual-wire-global"`
}


type NetworkVirtualWireGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *NetworkVirtualWireGlobal) GetId() string{
    return "1"
}

func (p *NetworkVirtualWireGlobal) getPath() string{
    return "network/virtual-wire-global"
}

func (p *NetworkVirtualWireGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireGlobal::Post")
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

func (p *NetworkVirtualWireGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireGlobal::Get")
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
func (p *NetworkVirtualWireGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireGlobal::Put")
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

func (p *NetworkVirtualWireGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
