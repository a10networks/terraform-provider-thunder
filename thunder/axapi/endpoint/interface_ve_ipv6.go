package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceVeIpv6 struct {
	Inst struct {
		AddressList []InterfaceVeIpv6AddressList `json:"address-list"`

		Inbound int `json:"inbound"`

		Inside int `json:"inside"`

		Ipv6Enable int `json:"ipv6-enable"`

		Ospf InterfaceVeIpv6Ospf1025 `json:"ospf"`

		Outside int `json:"outside"`

		Rip InterfaceVeIpv6Rip1035 `json:"rip"`

		Router InterfaceVeIpv6Router1037 `json:"router"`

		RouterAdver InterfaceVeIpv6RouterAdver `json:"router-adver"`

		StatefulFirewall InterfaceVeIpv6StatefulFirewall1042 `json:"stateful-firewall"`

		TtlIgnore int `json:"ttl-ignore"`

		Uuid string `json:"uuid"`

		V6AclName string `json:"v6-acl-name"`

		Ifnum string
	} `json:"ipv6"`
}

type InterfaceVeIpv6AddressList struct {
	Ipv6Addr    string `json:"ipv6-addr"`
	AddressType string `json:"address-type"`
}

type InterfaceVeIpv6Ospf1025 struct {
	NetworkList           []InterfaceVeIpv6OspfNetworkList1026           `json:"network-list"`
	Bfd                   int                                            `json:"bfd"`
	Disable               int                                            `json:"disable"`
	CostCfg               []InterfaceVeIpv6OspfCostCfg1027               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceVeIpv6OspfDeadIntervalCfg1028       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceVeIpv6OspfHelloIntervalCfg1029      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceVeIpv6OspfMtuIgnoreCfg1030          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceVeIpv6OspfNeighborCfg1031           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceVeIpv6OspfPriorityCfg1032           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceVeIpv6OspfRetransmitIntervalCfg1033 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceVeIpv6OspfTransmitDelayCfg1034      `json:"transmit-delay-cfg"`
	Uuid                  string                                         `json:"uuid"`
}

type InterfaceVeIpv6OspfNetworkList1026 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceVeIpv6OspfCostCfg1027 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfDeadIntervalCfg1028 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceVeIpv6OspfHelloIntervalCfg1029 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceVeIpv6OspfMtuIgnoreCfg1030 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfNeighborCfg1031 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceVeIpv6OspfPriorityCfg1032 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceVeIpv6OspfRetransmitIntervalCfg1033 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceVeIpv6OspfTransmitDelayCfg1034 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceVeIpv6Rip1035 struct {
	SplitHorizonCfg InterfaceVeIpv6RipSplitHorizonCfg1036 `json:"split-horizon-cfg"`
	Uuid            string                                `json:"uuid"`
}

type InterfaceVeIpv6RipSplitHorizonCfg1036 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceVeIpv6Router1037 struct {
	Ripng InterfaceVeIpv6RouterRipng1038 `json:"ripng"`
	Ospf  InterfaceVeIpv6RouterOspf1039  `json:"ospf"`
	Isis  InterfaceVeIpv6RouterIsis1041  `json:"isis"`
}

type InterfaceVeIpv6RouterRipng1038 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpv6RouterOspf1039 struct {
	AreaList []InterfaceVeIpv6RouterOspfAreaList1040 `json:"area-list"`
	Uuid     string                                  `json:"uuid"`
}

type InterfaceVeIpv6RouterOspfAreaList1040 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceVeIpv6RouterIsis1041 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

type InterfaceVeIpv6RouterAdver struct {
	Action                   string                                 `json:"action" dval:"disable"`
	DefaultLifetime          int                                    `json:"default-lifetime" dval:"1800"`
	HopLimit                 int                                    `json:"hop-limit" dval:"255"`
	MaxInterval              int                                    `json:"max-interval" dval:"600"`
	MinInterval              int                                    `json:"min-interval" dval:"200"`
	RateLimit                int                                    `json:"rate-limit" dval:"100000"`
	ReachableTime            int                                    `json:"reachable-time"`
	RetransmitTimer          int                                    `json:"retransmit-timer"`
	AdverMtuDisable          int                                    `json:"adver-mtu-disable" dval:"1"`
	AdverMtu                 int                                    `json:"adver-mtu"`
	PrefixList               []InterfaceVeIpv6RouterAdverPrefixList `json:"prefix-list"`
	ManagedConfigAction      string                                 `json:"managed-config-action" dval:"disable"`
	OtherConfigAction        string                                 `json:"other-config-action" dval:"disable"`
	AdverVrid                int                                    `json:"adver-vrid"`
	UseFloatingIp            int                                    `json:"use-floating-ip"`
	FloatingIp               string                                 `json:"floating-ip"`
	AdverVridDefault         int                                    `json:"adver-vrid-default"`
	UseFloatingIpDefaultVrid int                                    `json:"use-floating-ip-default-vrid"`
	FloatingIpDefaultVrid    string                                 `json:"floating-ip-default-vrid"`
}

type InterfaceVeIpv6RouterAdverPrefixList struct {
	Prefix            string `json:"prefix"`
	NotAutonomous     int    `json:"not-autonomous"`
	NotOnLink         int    `json:"not-on-link"`
	PreferredLifetime int    `json:"preferred-lifetime" dval:"604800"`
	ValidLifetime     int    `json:"valid-lifetime" dval:"2592000"`
}

type InterfaceVeIpv6StatefulFirewall1042 struct {
	Inside     int    `json:"inside"`
	ClassList  string `json:"class-list"`
	Outside    int    `json:"outside"`
	AccessList int    `json:"access-list"`
	AclName    string `json:"acl-name"`
	Uuid       string `json:"uuid"`
}

func (p *InterfaceVeIpv6) GetId() string {
	return "1"
}

func (p *InterfaceVeIpv6) getPath() string {
	return "interface/ve/" + p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceVeIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIpv6::Post")
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

func (p *InterfaceVeIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIpv6::Get")
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
func (p *InterfaceVeIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIpv6::Put")
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

func (p *InterfaceVeIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceVeIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
