package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

//based on ACOS 5_2_1-P4_81
type RouterOspfArea struct {
	Inst struct {
		ProcessId       int //outside key
		UseAreaIpv4     bool
		AreaIpv4        string                          `json:"area-ipv4"`
		AreaNum         int                             `json:"area-num" dval:"-1"`
		AuthCfg         RouterOspfAreaAuthCfg           `json:"auth-cfg"`
		DefaultCost     int                             `json:"default-cost" dval:"1"`
		FilterLists     []RouterOspfAreaFilterLists     `json:"filter-lists"`
		NssaCfg         RouterOspfAreaNssaCfg           `json:"nssa-cfg"`
		RangeList       []RouterOspfAreaRangeList       `json:"range-list"`
		Shortcut        string                          `json:"shortcut" dval:"default"`
		StubCfg         RouterOspfAreaStubCfg           `json:"stub-cfg"`
		Uuid            string                          `json:"uuid"`
		VirtualLinkList []RouterOspfAreaVirtualLinkList `json:"virtual-link-list"`
	} `json:"area"`
}

type RouterOspfAreaAuthCfg struct {
	Authentication int `json:"authentication"`
	MessageDigest  int `json:"message-digest"`
}

type RouterOspfAreaFilterLists struct {
	FilterList     int    `json:"filter-list"`
	AclName        string `json:"acl-name"`
	AclDirection   string `json:"acl-direction"`
	PlistName      string `json:"plist-name"`
	PlistDirection string `json:"plist-direction"`
}

type RouterOspfAreaNssaCfg struct {
	Nssa                        int    `json:"nssa"`
	NoRedistribution            int    `json:"no-redistribution"`
	NoSummary                   int    `json:"no-summary"`
	TranslatorRole              string `json:"translator-role" dval:"candidate"`
	DefaultInformationOriginate int    `json:"default-information-originate"`
	Metric                      int    `json:"metric" dval:"1"`
	MetricType                  int    `json:"metric-type" dval:"2"`
}

type RouterOspfAreaRangeList struct {
	AreaRangePrefix string `json:"area-range-prefix"`
	Option          string `json:"option" dval:"advertise"`
}

type RouterOspfAreaStubCfg struct {
	Stub      int `json:"stub"`
	NoSummary int `json:"no-summary"`
}

type RouterOspfAreaVirtualLinkList struct {
	VirtualLinkIpAddr         string `json:"virtual-link-ip-addr"`
	Bfd                       int    `json:"bfd"`
	HelloInterval             int    `json:"hello-interval"`
	DeadInterval              int    `json:"dead-interval"`
	RetransmitInterval        int    `json:"retransmit-interval"`
	TransmitDelay             int    `json:"transmit-delay" dval:"1"`
	VirtualLinkAuthentication int    `json:"virtual-link-authentication"`
	VirtualLinkAuthType       string `json:"virtual-link-auth-type"`
	AuthenticationKey         string `json:"authentication-key"`
	MessageDigestKey          int    `json:"message-digest-key"`
	Md5                       string `json:"md5"`
}

func (p *RouterOspfArea) GetId() string {
	if p.Inst.UseAreaIpv4 {
		return p.Inst.AreaIpv4
	}
	return strconv.Itoa(p.Inst.AreaNum)
}

func (p *RouterOspfArea) getPath() string {
	return "router/ospf/" + strconv.Itoa(p.Inst.ProcessId) + "/area"
}

func (p *RouterOspfArea) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterOspfArea::Post")
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

func (p *RouterOspfArea) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterOspfArea::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *RouterOspfArea) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterOspfArea::Put")
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

func (p *RouterOspfArea) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterOspfArea::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
