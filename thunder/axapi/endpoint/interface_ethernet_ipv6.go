package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type InterfaceEthernetIpv6 struct {
	Inst struct {
		AccessListCfg InterfaceEthernetIpv6AccessListCfg `json:"access-list-cfg"`

		AddressList []InterfaceEthernetIpv6AddressList `json:"address-list"`

		Inside int `json:"inside"`

		Ipv6Enable int `json:"ipv6-enable"`

		Ospf InterfaceEthernetIpv6Ospf559 `json:"ospf"`

		Outside int `json:"outside"`

		Rip InterfaceEthernetIpv6Rip569 `json:"rip"`

		Router InterfaceEthernetIpv6Router571 `json:"router"`

		RouterAdver InterfaceEthernetIpv6RouterAdver `json:"router-adver"`

		TtlIgnore int `json:"ttl-ignore"`

		Uuid string `json:"uuid"`

		Ifnum string
	} `json:"ipv6"`
}

type InterfaceEthernetIpv6AccessListCfg struct {
	V6AclName string `json:"v6-acl-name"`
	Inbound   int    `json:"inbound"`
}

type InterfaceEthernetIpv6AddressList struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceEthernetIpv6Ospf559 struct {
	NetworkList           []InterfaceEthernetIpv6OspfNetworkList560           `json:"network-list"`
	Bfd                   int                                                 `json:"bfd"`
	Disable               int                                                 `json:"disable"`
	CostCfg               []InterfaceEthernetIpv6OspfCostCfg561               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceEthernetIpv6OspfDeadIntervalCfg562       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceEthernetIpv6OspfHelloIntervalCfg563      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceEthernetIpv6OspfMtuIgnoreCfg564          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceEthernetIpv6OspfNeighborCfg565           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceEthernetIpv6OspfPriorityCfg566           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceEthernetIpv6OspfRetransmitIntervalCfg567 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceEthernetIpv6OspfTransmitDelayCfg568      `json:"transmit-delay-cfg"`
	Uuid                  string                                              `json:"uuid"`
}

type InterfaceEthernetIpv6OspfNetworkList560 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceEthernetIpv6OspfCostCfg561 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfDeadIntervalCfg562 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfHelloIntervalCfg563 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfMtuIgnoreCfg564 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfNeighborCfg565 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceEthernetIpv6OspfPriorityCfg566 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfRetransmitIntervalCfg567 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceEthernetIpv6OspfTransmitDelayCfg568 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceEthernetIpv6Rip569 struct {
	SplitHorizonCfg InterfaceEthernetIpv6RipSplitHorizonCfg570 `json:"split-horizon-cfg"`
	Uuid            string                                     `json:"uuid"`
}

type InterfaceEthernetIpv6RipSplitHorizonCfg570 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceEthernetIpv6Router571 struct {
	Ripng InterfaceEthernetIpv6RouterRipng572 `json:"ripng"`
	Ospf  InterfaceEthernetIpv6RouterOspf573  `json:"ospf"`
	Isis  InterfaceEthernetIpv6RouterIsis575  `json:"isis"`
}

type InterfaceEthernetIpv6RouterRipng572 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceEthernetIpv6RouterOspf573 struct {
	AreaList []InterfaceEthernetIpv6RouterOspfAreaList574 `json:"area-list"`
	Uuid     string                                       `json:"uuid"`
}

type InterfaceEthernetIpv6RouterOspfAreaList574 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceEthernetIpv6RouterIsis575 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceEthernetIpv6RouterAdver struct {
	Action                   string                                       `json:"action" dval:"disable"`
	HopLimit                 int                                          `json:"hop-limit" dval:"255"`
	MaxInterval              int                                          `json:"max-interval" dval:"600"`
	MinInterval              int                                          `json:"min-interval" dval:"200"`
	DefaultLifetime          int                                          `json:"default-lifetime" dval:"1800"`
	RateLimit                int                                          `json:"rate-limit" dval:"100000"`
	ReachableTime            int                                          `json:"reachable-time"`
	RetransmitTimer          int                                          `json:"retransmit-timer"`
	AdverMtuDisable          int                                          `json:"adver-mtu-disable" dval:"1"`
	AdverMtu                 int                                          `json:"adver-mtu"`
	PrefixList               []InterfaceEthernetIpv6RouterAdverPrefixList `json:"prefix-list"`
	ManagedConfigAction      string                                       `json:"managed-config-action" dval:"disable"`
	OtherConfigAction        string                                       `json:"other-config-action" dval:"disable"`
	AdverVrid                int                                          `json:"adver-vrid"`
	UseFloatingIp            int                                          `json:"use-floating-ip"`
	FloatingIp               string                                       `json:"floating-ip"`
	AdverVridDefault         int                                          `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int                                          `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string                                       `json:"floating-ip-default-vrid"`
}

type InterfaceEthernetIpv6RouterAdverPrefixList struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

func (p *InterfaceEthernetIpv6) GetId() string {
	return "1"
}

func (p *InterfaceEthernetIpv6) getPath() string {
	return "interface/ethernet/" + p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceEthernetIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernetIpv6::Post")
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

func (p *InterfaceEthernetIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernetIpv6::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *InterfaceEthernetIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernetIpv6::Put")
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

func (p *InterfaceEthernetIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceEthernetIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
