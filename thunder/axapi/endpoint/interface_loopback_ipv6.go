package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type InterfaceLoopbackIpv6 struct {
	Inst struct {
		AddressList []InterfaceLoopbackIpv6AddressList `json:"address-list"`

		Ipv6Enable int `json:"ipv6-enable"`

		Ospf InterfaceLoopbackIpv6Ospf759 `json:"ospf"`

		Rip InterfaceLoopbackIpv6Rip767 `json:"rip"`

		Router InterfaceLoopbackIpv6Router769 `json:"router"`

		Uuid string `json:"uuid"`

		Ifnum string
	} `json:"ipv6"`
}

type InterfaceLoopbackIpv6AddressList struct {
	Ipv6Addr  string `json:"ipv6-addr"`
	Anycast   int    `json:"anycast"`
	LinkLocal int    `json:"link-local"`
}

type InterfaceLoopbackIpv6Ospf759 struct {
	Bfd                   int                                                 `json:"bfd"`
	Disable               int                                                 `json:"disable"`
	CostCfg               []InterfaceLoopbackIpv6OspfCostCfg760               `json:"cost-cfg"`
	DeadIntervalCfg       []InterfaceLoopbackIpv6OspfDeadIntervalCfg761       `json:"dead-interval-cfg"`
	HelloIntervalCfg      []InterfaceLoopbackIpv6OspfHelloIntervalCfg762      `json:"hello-interval-cfg"`
	MtuIgnoreCfg          []InterfaceLoopbackIpv6OspfMtuIgnoreCfg763          `json:"mtu-ignore-cfg"`
	PriorityCfg           []InterfaceLoopbackIpv6OspfPriorityCfg764           `json:"priority-cfg"`
	RetransmitIntervalCfg []InterfaceLoopbackIpv6OspfRetransmitIntervalCfg765 `json:"retransmit-interval-cfg"`
	TransmitDelayCfg      []InterfaceLoopbackIpv6OspfTransmitDelayCfg766      `json:"transmit-delay-cfg"`
	Uuid                  string                                              `json:"uuid"`
}

type InterfaceLoopbackIpv6OspfCostCfg760 struct {
	Cost       int `json:"cost"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfDeadIntervalCfg761 struct {
	DeadInterval int `json:"dead-interval" dval:"40"`
	InstanceId   int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfHelloIntervalCfg762 struct {
	HelloInterval int `json:"hello-interval" dval:"10"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfMtuIgnoreCfg763 struct {
	MtuIgnore  int `json:"mtu-ignore"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfPriorityCfg764 struct {
	Priority   int `json:"priority" dval:"1"`
	InstanceId int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfRetransmitIntervalCfg765 struct {
	RetransmitInterval int `json:"retransmit-interval" dval:"5"`
	InstanceId         int `json:"instance-id"`
}

type InterfaceLoopbackIpv6OspfTransmitDelayCfg766 struct {
	TransmitDelay int `json:"transmit-delay" dval:"1"`
	InstanceId    int `json:"instance-id"`
}

type InterfaceLoopbackIpv6Rip767 struct {
	SplitHorizonCfg InterfaceLoopbackIpv6RipSplitHorizonCfg768 `json:"split-horizon-cfg"`
	Uuid            string                                     `json:"uuid"`
}

type InterfaceLoopbackIpv6RipSplitHorizonCfg768 struct {
	State string `json:"state" dval:"poisoned"`
}

type InterfaceLoopbackIpv6Router769 struct {
	Ripng InterfaceLoopbackIpv6RouterRipng770 `json:"ripng"`
	Ospf  InterfaceLoopbackIpv6RouterOspf771  `json:"ospf"`
	Isis  InterfaceLoopbackIpv6RouterIsis773  `json:"isis"`
}

type InterfaceLoopbackIpv6RouterRipng770 struct {
	Rip  int    `json:"rip"`
	Uuid string `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspf771 struct {
	AreaList []InterfaceLoopbackIpv6RouterOspfAreaList772 `json:"area-list"`
	Uuid     string                                       `json:"uuid"`
}

type InterfaceLoopbackIpv6RouterOspfAreaList772 struct {
	AreaIdNum  int    `json:"area-id-num"`
	AreaIdAddr string `json:"area-id-addr"`
	Tag        string `json:"tag"`
	InstanceId int    `json:"instance-id"`
}

type InterfaceLoopbackIpv6RouterIsis773 struct {
	Tag  string `json:"tag"`
	Uuid string `json:"uuid"`
}

func (p *InterfaceLoopbackIpv6) GetId() string {
	return "1"
}

func (p *InterfaceLoopbackIpv6) getPath() string {
	return "interface/loopback/" + p.Inst.Ifnum + "/ipv6"
}

func (p *InterfaceLoopbackIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopbackIpv6::Post")
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

func (p *InterfaceLoopbackIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopbackIpv6::Get")
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
func (p *InterfaceLoopbackIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopbackIpv6::Put")
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

func (p *InterfaceLoopbackIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("InterfaceLoopbackIpv6::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
