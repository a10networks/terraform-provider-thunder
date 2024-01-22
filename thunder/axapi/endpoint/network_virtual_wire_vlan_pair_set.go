

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireVlanPairSet struct {
	Inst struct {

    SetCfg []NetworkVirtualWireVlanPairSetSetCfg `json:"set-cfg"`
    SetId int `json:"set-id"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"virtual-wire-vlan-pair-set"`
}


type NetworkVirtualWireVlanPairSetSetCfg struct {
    VirtualWireVlanPair int `json:"virtual-wire-vlan-pair"`
}

func (p *NetworkVirtualWireVlanPairSet) GetId() string{
    return strconv.Itoa(p.Inst.SetId)
}

func (p *NetworkVirtualWireVlanPairSet) getPath() string{
    return "network/virtual-wire-vlan-pair-set"
}

func (p *NetworkVirtualWireVlanPairSet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPairSet::Post")
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

func (p *NetworkVirtualWireVlanPairSet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPairSet::Get")
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
func (p *NetworkVirtualWireVlanPairSet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPairSet::Put")
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

func (p *NetworkVirtualWireVlanPairSet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPairSet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
