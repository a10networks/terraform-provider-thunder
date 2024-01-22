

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ZoneInterface struct {
	Inst struct {

    EthernetList []ZoneInterfaceEthernetList `json:"ethernet-list"`
    LifList []ZoneInterfaceLifList `json:"lif-list"`
    TrunkList []ZoneInterfaceTrunkList `json:"trunk-list"`
    TunnelList []ZoneInterfaceTunnelList `json:"tunnel-list"`
    Uuid string `json:"uuid"`
    VeList []ZoneInterfaceVeList `json:"ve-list"`
    Name string 

	} `json:"interface"`
}


type ZoneInterfaceEthernetList struct {
    InterfaceEthernetStart int `json:"interface-ethernet-start"`
    InterfaceEthernetEnd int `json:"interface-ethernet-end"`
}


type ZoneInterfaceLifList struct {
    InterfaceLifStart int `json:"interface-lif-start"`
    InterfaceLifEnd int `json:"interface-lif-end"`
}


type ZoneInterfaceTrunkList struct {
    InterfaceTrunkStart int `json:"interface-trunk-start"`
    InterfaceTrunkEnd int `json:"interface-trunk-end"`
}


type ZoneInterfaceTunnelList struct {
    InterfaceTunnelStart int `json:"interface-tunnel-start"`
    InterfaceTunnelEnd int `json:"interface-tunnel-end"`
}


type ZoneInterfaceVeList struct {
    InterfaceVeStart int `json:"interface-ve-start"`
    InterfaceVeEnd int `json:"interface-ve-end"`
}

func (p *ZoneInterface) GetId() string{
    return "1"
}

func (p *ZoneInterface) getPath() string{
    return "zone/" +p.Inst.Name + "/interface"
}

func (p *ZoneInterface) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneInterface::Post")
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

func (p *ZoneInterface) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneInterface::Get")
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
func (p *ZoneInterface) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneInterface::Put")
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

func (p *ZoneInterface) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ZoneInterface::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
