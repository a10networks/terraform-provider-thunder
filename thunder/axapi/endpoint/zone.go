package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type Zone struct {
	Inst struct {
		Interface ZoneInterface3818 `json:"interface"`

		LocalZoneCfg ZoneLocalZoneCfg3824 `json:"local-zone-cfg"`

		Name string `json:"name"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Vlan ZoneVlan3825 `json:"vlan"`
	} `json:"zone"`
}

type ZoneInterface3818 struct {
	EthernetList []ZoneInterfaceEthernetList3819 `json:"ethernet-list"`
	TrunkList    []ZoneInterfaceTrunkList3820    `json:"trunk-list"`
	VeList       []ZoneInterfaceVeList3821       `json:"ve-list"`
	LifList      []ZoneInterfaceLifList3822      `json:"lif-list"`
	TunnelList   []ZoneInterfaceTunnelList3823   `json:"tunnel-list"`
	Uuid         string                          `json:"uuid"`
}

type ZoneInterfaceEthernetList3819 struct {
	InterfaceEthernetStart int `json:"interface-ethernet-start"`
	InterfaceEthernetEnd   int `json:"interface-ethernet-end"`
}

type ZoneInterfaceTrunkList3820 struct {
	InterfaceTrunkStart int `json:"interface-trunk-start"`
	InterfaceTrunkEnd   int `json:"interface-trunk-end"`
}

type ZoneInterfaceVeList3821 struct {
	InterfaceVeStart int `json:"interface-ve-start"`
	InterfaceVeEnd   int `json:"interface-ve-end"`
}

type ZoneInterfaceLifList3822 struct {
	InterfaceLifStart int `json:"interface-lif-start"`
	InterfaceLifEnd   int `json:"interface-lif-end"`
}

type ZoneInterfaceTunnelList3823 struct {
	InterfaceTunnelStart int `json:"interface-tunnel-start"`
	InterfaceTunnelEnd   int `json:"interface-tunnel-end"`
}

type ZoneLocalZoneCfg3824 struct {
	LocalType int    `json:"local-type"`
	Uuid      string `json:"uuid"`
}

type ZoneVlan3825 struct {
	VlanList []ZoneVlanVlanList3826 `json:"vlan-list"`
	Uuid     string                 `json:"uuid"`
}

type ZoneVlanVlanList3826 struct {
	VlanStart int `json:"vlan-start"`
	VlanEnd   int `json:"vlan-end"`
}

func (p *Zone) GetId() string {
	return p.Inst.Name
}

func (p *Zone) getPath() string {
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
		if len(axResp) > 0 {
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
