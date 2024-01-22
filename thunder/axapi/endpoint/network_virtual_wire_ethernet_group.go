

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVirtualWireEthernetGroup struct {
	Inst struct {

    EthList []NetworkVirtualWireEthernetGroupEthList `json:"eth-list"`
    GroupId int `json:"group-id"`
    TrunkList []NetworkVirtualWireEthernetGroupTrunkList `json:"trunk-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"virtual-wire-ethernet-group"`
}


type NetworkVirtualWireEthernetGroupEthList struct {
    EthStart int `json:"eth-start"`
    EthEnd int `json:"eth-end"`
}


type NetworkVirtualWireEthernetGroupTrunkList struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}

func (p *NetworkVirtualWireEthernetGroup) GetId() string{
    return strconv.Itoa(p.Inst.GroupId)
}

func (p *NetworkVirtualWireEthernetGroup) getPath() string{
    return "network/virtual-wire-ethernet-group"
}

func (p *NetworkVirtualWireEthernetGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireEthernetGroup::Post")
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

func (p *NetworkVirtualWireEthernetGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireEthernetGroup::Get")
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
func (p *NetworkVirtualWireEthernetGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireEthernetGroup::Put")
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

func (p *NetworkVirtualWireEthernetGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVirtualWireEthernetGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
