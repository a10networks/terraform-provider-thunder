package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type InterfaceLifIpv6 struct {
	Ipv6 InterfaceLifIpv6Inst `json:"ipv6"`
}

type InterfaceLifIpv6Inst struct {
	AddressList []InterfaceLifIpv6AddressList `json:"address-list"`
	Inside      int                           `json:"inside"`
	Ipv6Enable  int                           `json:"ipv6-enable"`
	Ospf        InterfaceLifIpv6Ospf659       `json:"ospf"`
	Outside     int                           `json:"outside"`
	Router      InterfaceLifIpv6Router669     `json:"router"`
	Uuid        string                        `json:"uuid"`
	Ifname      string                        `json:"-"`
}


type InterfaceLifIpv6AddressList struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLifIpv6Ospf659 struct {
	NetworkList           []InterfaceLifIpv6OspfNetworkList660           `json:"network-list"`
	Bfd                   int                                            `json:"bfd"`
	Disable               int                                            `json:"disable"`
	CostCfg               []InterfaceLifIpv6OspfCostCfg661               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLifIpv6OspfDeadIntervalCfg662       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLifIpv6OspfHelloIntervalCfg663      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLifIpv6OspfMtuIgnoreCfg664          `json:"mtu-ignore-cfg"`
	NeighborCfg           []InterfaceLifIpv6OspfNeighborCfg665           `json:"neighbor-cfg"`
	PriorityCfg           []InterfaceLifIpv6OspfPriorityCfg666           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLifIpv6OspfRetransmitIntervalCfg667 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLifIpv6OspfTransmitDelayCfg668      `json:"transmit-delay-cfg"`
	Uuid                  string                                         `json:"uuid"`
}

type InterfaceLifIpv6OspfNetworkList660 struct {
	BroadcastType     string `json:"broadcast-type"`
	P2mpNbma          int    `json:"p2mp-nbma"`
	NetworkInstanceId int    `json:"network-instance-id"`
}

type InterfaceLifIpv6OspfCostCfg661 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfDeadIntervalCfg662 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLifIpv6OspfHelloIntervalCfg663 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIpv6OspfMtuIgnoreCfg664 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfNeighborCfg665 struct {
	Neighbor             string `json:"neighbor" dval:"::"`
	NeigInst             int    `json:"neig-inst"`
	NeighborCost         int    `json:"neighbor-cost"`
	NeighborPollInterval int    `json:"neighbor-poll-interval"`
	NeighborPriority     int    `json:"neighbor-priority"`
}

type InterfaceLifIpv6OspfPriorityCfg666 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLifIpv6OspfRetransmitIntervalCfg667 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLifIpv6OspfTransmitDelayCfg668 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLifIpv6Router669 struct {
	Ripng InterfaceLifIpv6RouterRipng670 `json:"ripng"`
	Ospf  InterfaceLifIpv6RouterOspf671  `json:"ospf"`
	Isis  InterfaceLifIpv6RouterIsis673  `json:"isis"`
}

type InterfaceLifIpv6RouterRipng670 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLifIpv6RouterOspf671 struct {
	AreaList []InterfaceLifIpv6RouterOspfAreaList672 `json:"area-list"`
	Uuid     string                                  `json:"uuid"`
}

type InterfaceLifIpv6RouterOspfAreaList672 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLifIpv6RouterIsis673 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

func (p *InterfaceLifIpv6) GetId() string {
	return "1"
}

func (p *InterfaceLifIpv6) getPath() string {
	return "interface/lif/" + p.Ipv6.Ifname + "/ipv6"
}


func (p *InterfaceLifIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIpv6::Post")
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

func (p *InterfaceLifIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIpv6::Get")
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
func (p *InterfaceLifIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIpv6::Put")
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

func (p *InterfaceLifIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLifIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
