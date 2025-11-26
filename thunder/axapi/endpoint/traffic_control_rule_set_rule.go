package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type TrafficControlRuleSetRule struct {
	Inst struct {
		ActionGroup TrafficControlRuleSetRuleActionGroup2030 `json:"action-group"`

		AppList []TrafficControlRuleSetRuleAppList `json:"app-list"`

		ApplicationAny string `json:"application-any" dval:"any"`

		DerivedAttribute string `json:"derived-attribute"`

		DestList []TrafficControlRuleSetRuleDestList `json:"dest-list"`

		DstClassList string `json:"dst-class-list"`

		DstDomainList string `json:"dst-domain-list"`

		DstGeolocList string `json:"dst-geoloc-list"`

		DstGeolocListShared int `json:"dst-geoloc-list-shared"`

		DstGeolocName string `json:"dst-geoloc-name"`

		DstIpv4Any string `json:"dst-ipv4-any" dval:"any"`

		DstIpv6Any string `json:"dst-ipv6-any" dval:"any"`

		DstZone string `json:"dst-zone"`

		DstZoneAny string `json:"dst-zone-any" dval:"any"`

		IpVersion string `json:"ip-version" dval:"v4"`

		MoveRule TrafficControlRuleSetRuleMoveRule2031 `json:"move-rule"`

		Name string `json:"name"`

		Remark string `json:"remark"`

		SamplingEnable []TrafficControlRuleSetRuleSamplingEnable `json:"sampling-enable"`

		ServiceAny string `json:"service-any" dval:"any"`

		ServiceList []TrafficControlRuleSetRuleServiceList `json:"service-list"`

		SourceList []TrafficControlRuleSetRuleSourceList `json:"source-list"`

		SrcClassList string `json:"src-class-list"`

		SrcClassListType string `json:"src-class-list-type"`

		SrcGeolocList string `json:"src-geoloc-list"`

		SrcGeolocListShared int `json:"src-geoloc-list-shared"`

		SrcGeolocName string `json:"src-geoloc-name"`

		SrcIpv4Any string `json:"src-ipv4-any" dval:"any"`

		SrcIpv6Any string `json:"src-ipv6-any" dval:"any"`

		SrcZone string `json:"src-zone"`

		SrcZoneAny string `json:"src-zone-any" dval:"any"`

		Status string `json:"status" dval:"enable"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Rule_set_name string
	} `json:"rule"`
}

type TrafficControlRuleSetRuleActionGroup2030 struct {
	LimitPolicy int    `json:"limit-policy"`
	Uuid        string `json:"uuid"`
}

type TrafficControlRuleSetRuleAppList struct {
	ObjGrpApplication string `json:"obj-grp-application"`
	Protocol          string `json:"protocol"`
	ProtocolTag       string `json:"protocol-tag"`
}

type TrafficControlRuleSetRuleDestList struct {
	DstIpSubnet      string `json:"dst-ip-subnet"`
	DstIpv6Subnet    string `json:"dst-ipv6-subnet"`
	DstObjNetwork    string `json:"dst-obj-network"`
	DstObjGrpNetwork string `json:"dst-obj-grp-network"`
	DstSlbVserver    string `json:"dst-slb-vserver"`
}

type TrafficControlRuleSetRuleMoveRule2031 struct {
	Location   string `json:"location" dval:"bottom"`
	TargetRule string `json:"target-rule"`
}

type TrafficControlRuleSetRuleSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type TrafficControlRuleSetRuleServiceList struct {
	Protocols     string `json:"protocols"`
	ProtoId       int    `json:"proto-id"`
	ObjGrpService string `json:"obj-grp-service"`
	Icmp          int    `json:"icmp"`
	Icmpv6        int    `json:"icmpv6"`
	IcmpType      int    `json:"icmp-type"`
	SpecialType   string `json:"special-type"`
	IcmpCode      int    `json:"icmp-code"`
	SpecialCode   string `json:"special-code"`
	Icmpv6Type    int    `json:"icmpv6-type"`
	SpecialV6Type string `json:"special-v6-type"`
	Icmpv6Code    int    `json:"icmpv6-code"`
	SpecialV6Code string `json:"special-v6-code"`
	EqSrcPort     int    `json:"eq-src-port"`
	GtSrcPort     int    `json:"gt-src-port"`
	LtSrcPort     int    `json:"lt-src-port"`
	RangeSrcPort  int    `json:"range-src-port"`
	PortNumEndSrc int    `json:"port-num-end-src"`
	EqDstPort     int    `json:"eq-dst-port"`
	GtDstPort     int    `json:"gt-dst-port"`
	LtDstPort     int    `json:"lt-dst-port"`
	RangeDstPort  int    `json:"range-dst-port"`
	PortNumEndDst int    `json:"port-num-end-dst"`
	SctpTemplate  string `json:"sctp-template"`
}

type TrafficControlRuleSetRuleSourceList struct {
	SrcIpSubnet      string `json:"src-ip-subnet"`
	SrcIpv6Subnet    string `json:"src-ipv6-subnet"`
	SrcObjNetwork    string `json:"src-obj-network"`
	SrcObjGrpNetwork string `json:"src-obj-grp-network"`
}

func (p *TrafficControlRuleSetRule) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *TrafficControlRuleSetRule) getPath() string {
	return "traffic-control/rule-set/" + p.Inst.Rule_set_name + "/rule"
}

func (p *TrafficControlRuleSetRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("TrafficControlRuleSetRule::Post")
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

func (p *TrafficControlRuleSetRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("TrafficControlRuleSetRule::Get")
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
func (p *TrafficControlRuleSetRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("TrafficControlRuleSetRule::Put")
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

func (p *TrafficControlRuleSetRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("TrafficControlRuleSetRule::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
