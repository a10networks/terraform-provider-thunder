

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Zone struct {
	Inst struct {

    Interface ZoneInterface3681 `json:"interface"`
    LocalZoneCfg ZoneLocalZoneCfg3687 `json:"local-zone-cfg"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vlan ZoneVlan3688 `json:"vlan"`

	} `json:"zone"`
}


type ZoneInterface3681 struct {
    EthernetList []ZoneInterfaceEthernetList3682 `json:"ethernet-list"`
    TrunkList []ZoneInterfaceTrunkList3683 `json:"trunk-list"`
    VeList []ZoneInterfaceVeList3684 `json:"ve-list"`
    LifList []ZoneInterfaceLifList3685 `json:"lif-list"`
    TunnelList []ZoneInterfaceTunnelList3686 `json:"tunnel-list"`
    Uuid string `json:"uuid"`
}


type ZoneInterfaceEthernetList3682 struct {
    InterfaceEthernetStart int `json:"interface-ethernet-start"`
    InterfaceEthernetEnd int `json:"interface-ethernet-end"`
}


type ZoneInterfaceTrunkList3683 struct {
    InterfaceTrunkStart int `json:"interface-trunk-start"`
    InterfaceTrunkEnd int `json:"interface-trunk-end"`
}


type ZoneInterfaceVeList3684 struct {
    InterfaceVeStart int `json:"interface-ve-start"`
    InterfaceVeEnd int `json:"interface-ve-end"`
}


type ZoneInterfaceLifList3685 struct {
    InterfaceLifStart int `json:"interface-lif-start"`
    InterfaceLifEnd int `json:"interface-lif-end"`
}


type ZoneInterfaceTunnelList3686 struct {
    InterfaceTunnelStart int `json:"interface-tunnel-start"`
    InterfaceTunnelEnd int `json:"interface-tunnel-end"`
}


type ZoneLocalZoneCfg3687 struct {
    LocalType int `json:"local-type"`
    Uuid string `json:"uuid"`
}


type ZoneVlan3688 struct {
    VlanList []ZoneVlanVlanList3689 `json:"vlan-list"`
    Uuid string `json:"uuid"`
}


type ZoneVlanVlanList3689 struct {
    VlanStart int `json:"vlan-start"`
    VlanEnd int `json:"vlan-end"`
}

func (p *Zone) GetId() string{
    return p.Inst.Name
}

func (p *Zone) getPath() string{
    return "zone"
}

func (p *Zone) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Zone::Post")
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

func (p *Zone) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Zone::Get")
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
func (p *Zone) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Zone::Put")
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

func (p *Zone) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Zone::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
