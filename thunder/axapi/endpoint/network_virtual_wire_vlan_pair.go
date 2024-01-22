

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireVlanPair struct {
	Inst struct {

    PairId int `json:"pair-id"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vlan1 int `json:"vlan1"`
    Vlan2 int `json:"vlan2"`

	} `json:"virtual-wire-vlan-pair"`
}

func (p *NetworkVirtualWireVlanPair) GetId() string{
    return strconv.Itoa(p.Inst.PairId)
}

func (p *NetworkVirtualWireVlanPair) getPath() string{
    return "network/virtual-wire-vlan-pair"
}

func (p *NetworkVirtualWireVlanPair) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPair::Post")
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

func (p *NetworkVirtualWireVlanPair) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPair::Get")
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
func (p *NetworkVirtualWireVlanPair) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPair::Put")
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

func (p *NetworkVirtualWireVlanPair) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireVlanPair::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
