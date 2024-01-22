

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkVlanGroup struct {
	Inst struct {

    EthList []NetworkVlanGroupEthList `json:"eth-list"`
    TrunkList []NetworkVlanGroupTrunkList `json:"trunk-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VlanNumEnd int `json:"vlan-num-end"`
    VlanNumStart int `json:"vlan-num-start"`

	} `json:"vlan-group"`
}


type NetworkVlanGroupEthList struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type NetworkVlanGroupTrunkList struct {
    TrunkStart int `json:"trunk-start"`
    TrunkEnd int `json:"trunk-end"`
}

func (p *NetworkVlanGroup) GetId() string{
    return strconv.Itoa(p.Inst.VlanNumStart)+"+"+strconv.Itoa(p.Inst.VlanNumEnd)
}

func (p *NetworkVlanGroup) getPath() string{
    return "network/vlan-group"
}

func (p *NetworkVlanGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGroup::Post")
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

func (p *NetworkVlanGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGroup::Get")
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
func (p *NetworkVlanGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGroup::Put")
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

func (p *NetworkVlanGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkVlanGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
