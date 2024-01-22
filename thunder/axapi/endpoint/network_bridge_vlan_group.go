

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkBridgeVlanGroup struct {
	Inst struct {

    BridgeVlanGroupNumber int `json:"bridge-vlan-group-number"`
    ForwardTraffic string `json:"forward-traffic" dval:"forward-ip-traffic"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Ve int `json:"ve"`
    VlanList []NetworkBridgeVlanGroupVlanList `json:"vlan-list"`

	} `json:"bridge-vlan-group"`
}


type NetworkBridgeVlanGroupVlanList struct {
    VlanStart int `json:"vlan-start"`
    VlanEnd int `json:"vlan-end"`
}

func (p *NetworkBridgeVlanGroup) GetId() string{
    return strconv.Itoa(p.Inst.BridgeVlanGroupNumber)
}

func (p *NetworkBridgeVlanGroup) getPath() string{
    return "network/bridge-vlan-group"
}

func (p *NetworkBridgeVlanGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBridgeVlanGroup::Post")
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

func (p *NetworkBridgeVlanGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBridgeVlanGroup::Get")
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
func (p *NetworkBridgeVlanGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBridgeVlanGroup::Put")
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

func (p *NetworkBridgeVlanGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkBridgeVlanGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
