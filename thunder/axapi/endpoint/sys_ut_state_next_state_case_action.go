package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SysUtStateNextStateCaseAction struct {
	Inst struct {
		Delay int `json:"delay"`

		Direction string `json:"direction"`

		Drop int `json:"drop"`

		L1 SysUtStateNextStateCaseActionL11644 `json:"l1"`

		L2 SysUtStateNextStateCaseActionL21647 `json:"l2"`

		L3 SysUtStateNextStateCaseActionL31649 `json:"l3"`

		Tcp SysUtStateNextStateCaseActionTcp1651 `json:"tcp"`

		Template string `json:"template"`

		Udp SysUtStateNextStateCaseActionUdp1654 `json:"udp"`

		Uuid string `json:"uuid"`

		State_name string

		CaseNumber string

		Next_state_name string
	} `json:"action"`
}

type SysUtStateNextStateCaseActionL11644 struct {
	EthList    []SysUtStateNextStateCaseActionL1EthList1645    `json:"eth-list"`
	Trunk_list []SysUtStateNextStateCaseActionL1Trunk_list1646 `json:"trunk_list"`
	Length     int                                             `json:"length"`
	Value      int                                             `json:"value"`
	Auto       int                                             `json:"auto"`
	Uuid       string                                          `json:"uuid"`
}

type SysUtStateNextStateCaseActionL1EthList1645 struct {
	EthernetStart int `json:"ethernet-start"`
	EthernetEnd   int `json:"ethernet-end"`
}

type SysUtStateNextStateCaseActionL1Trunk_list1646 struct {
	TrunkStart int `json:"trunk-start"`
	TrunkEnd   int `json:"trunk-end"`
}

type SysUtStateNextStateCaseActionL21647 struct {
	Ethertype int                                          `json:"ethertype"`
	Protocol  string                                       `json:"protocol" dval:"ipv4"`
	Value     int                                          `json:"value"`
	Vlan      int                                          `json:"vlan"`
	Uuid      string                                       `json:"uuid"`
	MacList   []SysUtStateNextStateCaseActionL2MacList1648 `json:"mac-list"`
}

type SysUtStateNextStateCaseActionL2MacList1648 struct {
	SrcDst        string `json:"src-dst"`
	AddressType   string `json:"address-type"`
	VirtualServer string `json:"virtual-server"`
	NatPool       string `json:"nat-pool"`
	Ethernet      int    `json:"ethernet"`
	Ve            int    `json:"ve"`
	Trunk         int    `json:"trunk"`
	Value         string `json:"value"`
	Uuid          string `json:"uuid"`
}

type SysUtStateNextStateCaseActionL31649 struct {
	Protocol int                                         `json:"protocol"`
	Type     string                                      `json:"type"`
	Value    int                                         `json:"value"`
	Checksum string                                      `json:"checksum" dval:"valid"`
	Ttl      int                                         `json:"ttl"`
	Uuid     string                                      `json:"uuid"`
	IpList   []SysUtStateNextStateCaseActionL3IpList1650 `json:"ip-list"`
}

type SysUtStateNextStateCaseActionL3IpList1650 struct {
	SrcDst        string `json:"src-dst"`
	Ipv4Address   string `json:"ipv4-address"`
	Ipv6Address   string `json:"ipv6-address"`
	VirtualServer string `json:"virtual-server"`
	NatPool       string `json:"nat-pool"`
	Ethernet      int    `json:"ethernet"`
	Ve            int    `json:"ve"`
	Trunk         int    `json:"trunk"`
	Uuid          string `json:"uuid"`
}

type SysUtStateNextStateCaseActionTcp1651 struct {
	SrcPort       int                                         `json:"src-port"`
	DestPort      int                                         `json:"dest-port"`
	DestPortValue int                                         `json:"dest-port-value"`
	NatPool       string                                      `json:"nat-pool"`
	SeqNumber     string                                      `json:"seq-number" dval:"valid"`
	AckSeqNumber  string                                      `json:"ack-seq-number" dval:"valid"`
	Checksum      string                                      `json:"checksum" dval:"valid"`
	Urgent        string                                      `json:"urgent" dval:"valid"`
	Window        string                                      `json:"window" dval:"valid"`
	Uuid          string                                      `json:"uuid"`
	Flags         SysUtStateNextStateCaseActionTcpFlags1652   `json:"flags"`
	Options       SysUtStateNextStateCaseActionTcpOptions1653 `json:"options"`
}

type SysUtStateNextStateCaseActionTcpFlags1652 struct {
	Syn  int    `json:"syn"`
	Ack  int    `json:"ack"`
	Fin  int    `json:"fin"`
	Rst  int    `json:"rst"`
	Psh  int    `json:"psh"`
	Ece  int    `json:"ece"`
	Urg  int    `json:"urg"`
	Cwr  int    `json:"cwr"`
	Uuid string `json:"uuid"`
}

type SysUtStateNextStateCaseActionTcpOptions1653 struct {
	Mss             int    `json:"mss"`
	Wscale          int    `json:"wscale"`
	SackType        string `json:"sack-type"`
	TimeStampEnable int    `json:"time-stamp-enable"`
	Nop             int    `json:"nop"`
	Uuid            string `json:"uuid"`
}

type SysUtStateNextStateCaseActionUdp1654 struct {
	SrcPort       int    `json:"src-port"`
	DestPort      int    `json:"dest-port"`
	DestPortValue int    `json:"dest-port-value"`
	NatPool       string `json:"nat-pool"`
	Length        int    `json:"length"`
	Checksum      string `json:"checksum" dval:"valid"`
	Uuid          string `json:"uuid"`
}

func (p *SysUtStateNextStateCaseAction) GetId() string {
	return p.Inst.Direction
}

func (p *SysUtStateNextStateCaseAction) getPath() string {
	return "sys-ut/state/" + p.Inst.State_name + "/next-state/" + p.Inst.Next_state_name + "/case/" + p.Inst.CaseNumber + "/action"
}

func (p *SysUtStateNextStateCaseAction) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseAction::Post")
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

func (p *SysUtStateNextStateCaseAction) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseAction::Get")
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
func (p *SysUtStateNextStateCaseAction) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseAction::Put")
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

func (p *SysUtStateNextStateCaseAction) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtStateNextStateCaseAction::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
