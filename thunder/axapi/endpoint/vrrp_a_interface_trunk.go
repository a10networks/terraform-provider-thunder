

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAInterfaceTrunk struct {
	Inst struct {

    Both int `json:"both"`
    NoHeartbeat int `json:"no-heartbeat"`
    RouterInterface int `json:"router-interface"`
    ServerInterface int `json:"server-interface"`
    TrunkVal int `json:"trunk-val"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VlanCfg []VrrpAInterfaceTrunkVlanCfg `json:"vlan-cfg"`

	} `json:"trunk"`
}


type VrrpAInterfaceTrunkVlanCfg struct {
    Vlan int `json:"vlan"`
}

func (p *VrrpAInterfaceTrunk) GetId() string{
    return strconv.Itoa(p.Inst.TrunkVal)
}

func (p *VrrpAInterfaceTrunk) getPath() string{
    return "vrrp-a/interface/trunk"
}

func (p *VrrpAInterfaceTrunk) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceTrunk::Post")
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

func (p *VrrpAInterfaceTrunk) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceTrunk::Get")
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
func (p *VrrpAInterfaceTrunk) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceTrunk::Put")
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

func (p *VrrpAInterfaceTrunk) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceTrunk::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
