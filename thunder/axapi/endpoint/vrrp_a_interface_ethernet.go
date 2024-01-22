

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAInterfaceEthernet struct {
	Inst struct {

    Both int `json:"both"`
    EthernetVal int `json:"ethernet-val"`
    NoHeartbeat int `json:"no-heartbeat"`
    RouterInterface int `json:"router-interface"`
    ServerInterface int `json:"server-interface"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VlanCfg []VrrpAInterfaceEthernetVlanCfg `json:"vlan-cfg"`

	} `json:"ethernet"`
}


type VrrpAInterfaceEthernetVlanCfg struct {
    Vlan int `json:"vlan"`
}

func (p *VrrpAInterfaceEthernet) GetId() string{
    return strconv.Itoa(p.Inst.EthernetVal)
}

func (p *VrrpAInterfaceEthernet) getPath() string{
    return "vrrp-a/interface/ethernet"
}

func (p *VrrpAInterfaceEthernet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceEthernet::Post")
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

func (p *VrrpAInterfaceEthernet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceEthernet::Get")
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
func (p *VrrpAInterfaceEthernet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceEthernet::Put")
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

func (p *VrrpAInterfaceEthernet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpAInterfaceEthernet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
