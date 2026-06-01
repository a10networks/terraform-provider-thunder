package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceTrunkIpv6 struct {
	Inst struct {
		AccessListCfg InterfaceTrunkIpv6AccessListCfg `json:"access-list-cfg"`

		AddressList []InterfaceTrunkIpv6AddressList `json:"address-list"`

		Ipv6Enable int `json:"ipv6-enable"`

		Nat InterfaceTrunkIpv6Nat `json:"nat"`

		Ospf InterfaceTrunkIpv6Ospf843 `json:"ospf"`

		Rip InterfaceTrunkIpv6Rip853 `json:"rip"`

		Router InterfaceTrunkIpv6Router855 `json:"router"`

		RouterAdver InterfaceTrunkIpv6RouterAdver `json:"router-adver"`

		StatefulFirewall InterfaceTrunkIpv6StatefulFirewall860 `json:"stateful-firewall"`

		TtlIgnore int `json:"ttl-ignore"`

		Uuid string `json:"uuid"`

		Ifnum string
	} `json:"ipv6"`
}

type InterfaceTrunkIpv6AccessListCfg struct {
	V6AclName string `json:"v6-acl-name"`
	Inbound   int    `json:"inbound"`
}

type InterfaceTrunkIpv6AddressList struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceTrunkIpv6Nat struct {
	Inside  int `json:"inside"`
	Outside int `json:"outside"`
}

type InterfaceTrunkIpv6Ospf843 struct {
	NetworkList           []InterfaceTrunkIpv6OspfNetworkList844           `json:"network-list"`
	Bfd                   int                                              `json:"bfd"`
	Disable               int                                              `json:"disable"`
	CostCfg               []InterfaceTrunkIpv6OspfCostCfg845               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceTrunkIpv6OspfDeadIntervalCfg846       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceTrunkIpv6OspfHelloIntervalCfg847      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceTrunkIpv6OspfMtuIgnoreCfg848          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceTrunkIpv6OspfNeighborCfg849           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceTrunkIpv6OspfPriorityCfg850           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceTrunkIpv6OspfRetransmitIntervalCfg851 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceTrunkIpv6OspfTransmitDelayCfg852      `json:"transmit-delay-cfg"`
	Uuid                  string                                           `json:"uuid"`
}

type InterfaceTrunkIpv6OspfNetworkList844 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceTrunkIpv6OspfCostCfg845 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfDeadIntervalCfg846 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfHelloIntervalCfg847 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfMtuIgnoreCfg848 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfNeighborCfg849 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceTrunkIpv6OspfPriorityCfg850 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfRetransmitIntervalCfg851 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceTrunkIpv6OspfTransmitDelayCfg852 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceTrunkIpv6Rip853 struct {
	SplitHorizonCfg InterfaceTrunkIpv6RipSplitHorizonCfg854 `json:"split-horizon-cfg"`
	Uuid            string                                  `json:"uuid"`
}

type InterfaceTrunkIpv6RipSplitHorizonCfg854 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceTrunkIpv6Router855 struct {
	Ripng InterfaceTrunkIpv6RouterRipng856 `json:"ripng"`
	Ospf  InterfaceTrunkIpv6RouterOspf857  `json:"ospf"`
	Isis  InterfaceTrunkIpv6RouterIsis859  `json:"isis"`
}

type InterfaceTrunkIpv6RouterRipng856 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpv6RouterOspf857 struct {
	AreaList []InterfaceTrunkIpv6RouterOspfAreaList858 `json:"area-list"`
	Uuid     string                                    `json:"uuid"`
}

type InterfaceTrunkIpv6RouterOspfAreaList858 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceTrunkIpv6RouterIsis859 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceTrunkIpv6RouterAdver struct {
	Action              string                                    `json:"action" dval:"disable"`
	DefaultLifetime     int                                       `json:"default-lifetime" dval:"1800"`
	HopLimit            int                                       `json:"hop-limit" dval:"255"`
	MaxInterval         int                                       `json:"max-interval" dval:"600"`
	MinInterval         int                                       `json:"min-interval" dval:"200"`
	RateLimit           int                                       `json:"rate-limit" dval:"100000"`
	ReachableTime       int                                       `json:"reachable-time"`
	RetransmitTimer     int                                       `json:"retransmit-timer"`
	Mtu                 InterfaceTrunkIpv6RouterAdverMtu          `json:"mtu"`
	PrefixList          []InterfaceTrunkIpv6RouterAdverPrefixList `json:"prefix-list"`
	ManagedConfigAction string                                    `json:"managed-config-action" dval:"disable"`
	OtherConfigAction   string                                    `json:"other-config-action" dval:"disable"`
	Vrid                InterfaceTrunkIpv6RouterAdverVrid         `json:"vrid"`
}

type InterfaceTrunkIpv6RouterAdverMtu struct {
	AdverMtuDisable int `json:"adver-mtu-disable" dval:"1"`
	AdverMtu        int `json:"adver-mtu"`
}

type InterfaceTrunkIpv6RouterAdverPrefixList struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceTrunkIpv6RouterAdverVrid struct {
	AdverVrid                int    `json:"adver-vrid"`
	UseFloatingIp            int    `json:"use-floating-ip"`
	FloatingIp               string `json:"floating-ip"`
	AdverVridDefault         int    `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int    `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string `json:"floating-ip-default-vrid"`
}

type InterfaceTrunkIpv6StatefulFirewall860 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclName    string `json:"acl-name"`
	Uuid       string `json:"uuid"`
}

func (p *InterfaceTrunkIpv6) GetId() string {
	return "1"
}

func (p *InterfaceTrunkIpv6) getPath() string {
	return "interface/trunk/" + p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceTrunkIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIpv6::Post")
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

func (p *InterfaceTrunkIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIpv6::Get")
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
func (p *InterfaceTrunkIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIpv6::Put")
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

func (p *InterfaceTrunkIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceTrunkIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
