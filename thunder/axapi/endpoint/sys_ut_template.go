package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SysUtTemplate struct {
	Inst struct {
		IgnoreValidation SysUtTemplateIgnoreValidation1657 `json:"ignore-validation"`

		L1 SysUtTemplateL11658 `json:"l1"`

		L2 SysUtTemplateL21661 `json:"l2"`

		L3 SysUtTemplateL31663 `json:"l3"`

		Name string `json:"name"`

		Tcp SysUtTemplateTcp1665 `json:"tcp"`

		Udp SysUtTemplateUdp1669 `json:"udp"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"template"`
}

type SysUtTemplateIgnoreValidation1657 struct {
	L1   int    `json:"l1"`
	L2   int    `json:"l2"`
	L3   int    `json:"l3"`
	L4   int    `json:"l4"`
	All  int    `json:"all"`
	Uuid string `json:"uuid"`
}

type SysUtTemplateL11658 struct {
	EthList    []SysUtTemplateL1EthList1659    `json:"eth-list"`
	Trunk_list []SysUtTemplateL1Trunk_list1660 `json:"trunk_list"`
	Drop       int                             `json:"drop"`
	Length     int                             `json:"length"`
	Value      int                             `json:"value"`
	Auto       int                             `json:"auto"`
	Uuid       string                          `json:"uuid"`
}

type SysUtTemplateL1EthList1659 struct {
	EthernetStart int `json:"ethernet-start"`
	EthernetEnd   int `json:"ethernet-end"`
}

type SysUtTemplateL1Trunk_list1660 struct {
	TrunkStart int `json:"trunk-start"`
	TrunkEnd   int `json:"trunk-end"`
}

type SysUtTemplateL21661 struct {
	Ethertype int                          `json:"ethertype"`
	Protocol  string                       `json:"protocol" dval:"ipv4"`
	Value     int                          `json:"value"`
	Vlan      int                          `json:"vlan"`
	Uuid      string                       `json:"uuid"`
	MacList   []SysUtTemplateL2MacList1662 `json:"mac-list"`
}

type SysUtTemplateL2MacList1662 struct {
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

type SysUtTemplateL31663 struct {
	Protocol int                         `json:"protocol"`
	Type     string                      `json:"type"`
	Value    int                         `json:"value"`
	Checksum string                      `json:"checksum" dval:"valid"`
	Ttl      int                         `json:"ttl"`
	Uuid     string                      `json:"uuid"`
	IpList   []SysUtTemplateL3IpList1664 `json:"ip-list"`
}

type SysUtTemplateL3IpList1664 struct {
	SrcDst           string `json:"src-dst"`
	Ipv4StartAddress string `json:"ipv4-start-address"`
	Ipv4EndAddress   string `json:"ipv4-end-address"`
	Ipv6StartAddress string `json:"ipv6-start-address"`
	Ipv6EndAddress   string `json:"ipv6-end-address"`
	VirtualServer    string `json:"virtual-server"`
	NatPool          string `json:"nat-pool"`
	Ethernet         int    `json:"ethernet"`
	Ve               int    `json:"ve"`
	Trunk            int    `json:"trunk"`
	Uuid             string `json:"uuid"`
}

type SysUtTemplateTcp1665 struct {
	SrcPortRange  []SysUtTemplateTcpSrcPortRange1666 `json:"src-port-range"`
	DestPort      int                                `json:"dest-port"`
	DestPortValue int                                `json:"dest-port-value"`
	NatPool       string                             `json:"nat-pool"`
	SeqNumber     string                             `json:"seq-number" dval:"valid"`
	AckSeqNumber  string                             `json:"ack-seq-number" dval:"valid"`
	Checksum      string                             `json:"checksum" dval:"valid"`
	Urgent        string                             `json:"urgent" dval:"valid"`
	Window        string                             `json:"window" dval:"valid"`
	Uuid          string                             `json:"uuid"`
	Flags         SysUtTemplateTcpFlags1667          `json:"flags"`
	Options       SysUtTemplateTcpOptions1668        `json:"options"`
}

type SysUtTemplateTcpSrcPortRange1666 struct {
	SrcPortStart int `json:"src-port-start"`
	SrcPortEnd   int `json:"src-port-end"`
}

type SysUtTemplateTcpFlags1667 struct {
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

type SysUtTemplateTcpOptions1668 struct {
	Mss             int    `json:"mss"`
	Wscale          int    `json:"wscale"`
	SackType        string `json:"sack-type"`
	TimeStampEnable int    `json:"time-stamp-enable"`
	Nop             int    `json:"nop"`
	Uuid            string `json:"uuid"`
}

type SysUtTemplateUdp1669 struct {
	SrcPortRange  []SysUtTemplateUdpSrcPortRange1670 `json:"src-port-range"`
	DestPort      int                                `json:"dest-port"`
	DestPortValue int                                `json:"dest-port-value"`
	NatPool       string                             `json:"nat-pool"`
	Length        int                                `json:"length"`
	Checksum      string                             `json:"checksum" dval:"valid"`
	Uuid          string                             `json:"uuid"`
}

type SysUtTemplateUdpSrcPortRange1670 struct {
	SrcPortStart int `json:"src-port-start"`
	SrcPortEnd   int `json:"src-port-end"`
}

func (p *SysUtTemplate) GetId() string {
	return p.Inst.Name
}

func (p *SysUtTemplate) getPath() string {
	return "sys-ut/template"
}

func (p *SysUtTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtTemplate::Post")
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

func (p *SysUtTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtTemplate::Get")
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
func (p *SysUtTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtTemplate::Put")
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

func (p *SysUtTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SysUtTemplate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
