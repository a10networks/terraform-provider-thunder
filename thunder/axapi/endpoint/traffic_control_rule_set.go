

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TrafficControlRuleSet struct {
	Inst struct {

    Name string `json:"name"`
    Remark string `json:"remark"`
    RuleList []TrafficControlRuleSetRuleList `json:"rule-list"`
    SamplingEnable []TrafficControlRuleSetSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"rule-set"`
}


type TrafficControlRuleSetRuleList struct {
    Name string `json:"name"`
    Remark string `json:"remark"`
    Status string `json:"status" dval:"enable"`
    IpVersion string `json:"ip-version" dval:"v4"`
    SrcGeolocName string `json:"src-geoloc-name"`
    SrcGeolocList string `json:"src-geoloc-list"`
    SrcGeolocListShared int `json:"src-geoloc-list-shared"`
    SrcIpv4Any string `json:"src-ipv4-any" dval:"any"`
    SrcIpv6Any string `json:"src-ipv6-any" dval:"any"`
    SrcClassList string `json:"src-class-list"`
    SourceList []TrafficControlRuleSetRuleListSourceList `json:"source-list"`
    SrcZone string `json:"src-zone"`
    SrcZoneAny string `json:"src-zone-any" dval:"any"`
    SrcThreatList string `json:"src-threat-list"`
    DstGeolocName string `json:"dst-geoloc-name"`
    DstGeolocList string `json:"dst-geoloc-list"`
    DstGeolocListShared int `json:"dst-geoloc-list-shared"`
    DstIpv4Any string `json:"dst-ipv4-any" dval:"any"`
    DstIpv6Any string `json:"dst-ipv6-any" dval:"any"`
    DstClassList string `json:"dst-class-list"`
    DestList []TrafficControlRuleSetRuleListDestList `json:"dest-list"`
    DstDomainList string `json:"dst-domain-list"`
    DstZone string `json:"dst-zone"`
    DstZoneAny string `json:"dst-zone-any" dval:"any"`
    DstThreatList string `json:"dst-threat-list"`
    ServiceAny string `json:"service-any" dval:"any"`
    ServiceList []TrafficControlRuleSetRuleListServiceList `json:"service-list"`
    ApplicationAny string `json:"application-any" dval:"any"`
    AppList []TrafficControlRuleSetRuleListAppList `json:"app-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []TrafficControlRuleSetRuleListSamplingEnable `json:"sampling-enable"`
    ActionGroup TrafficControlRuleSetRuleListActionGroup `json:"action-group"`
}


type TrafficControlRuleSetRuleListSourceList struct {
    SrcIpSubnet string `json:"src-ip-subnet"`
    SrcIpv6Subnet string `json:"src-ipv6-subnet"`
    SrcObjNetwork string `json:"src-obj-network"`
    SrcObjGrpNetwork string `json:"src-obj-grp-network"`
    SrcSlbServer string `json:"src-slb-server"`
}


type TrafficControlRuleSetRuleListDestList struct {
    DstIpSubnet string `json:"dst-ip-subnet"`
    DstIpv6Subnet string `json:"dst-ipv6-subnet"`
    DstObjNetwork string `json:"dst-obj-network"`
    DstObjGrpNetwork string `json:"dst-obj-grp-network"`
    DstSlbServer string `json:"dst-slb-server"`
    DstSlbVserver string `json:"dst-slb-vserver"`
}


type TrafficControlRuleSetRuleListServiceList struct {
    Protocols string `json:"protocols"`
    ProtoId int `json:"proto-id"`
    ObjGrpService string `json:"obj-grp-service"`
    Icmp int `json:"icmp"`
    Icmpv6 int `json:"icmpv6"`
    IcmpType int `json:"icmp-type"`
    SpecialType string `json:"special-type"`
    IcmpCode int `json:"icmp-code"`
    SpecialCode string `json:"special-code"`
    Icmpv6Type int `json:"icmpv6-type"`
    SpecialV6Type string `json:"special-v6-type"`
    Icmpv6Code int `json:"icmpv6-code"`
    SpecialV6Code string `json:"special-v6-code"`
    EqSrcPort int `json:"eq-src-port"`
    GtSrcPort int `json:"gt-src-port"`
    LtSrcPort int `json:"lt-src-port"`
    RangeSrcPort int `json:"range-src-port"`
    PortNumEndSrc int `json:"port-num-end-src"`
    EqDstPort int `json:"eq-dst-port"`
    GtDstPort int `json:"gt-dst-port"`
    LtDstPort int `json:"lt-dst-port"`
    RangeDstPort int `json:"range-dst-port"`
    PortNumEndDst int `json:"port-num-end-dst"`
    SctpTemplate string `json:"sctp-template"`
}


type TrafficControlRuleSetRuleListAppList struct {
    ObjGrpApplication string `json:"obj-grp-application"`
    Protocol string `json:"protocol"`
    ProtocolTag string `json:"protocol-tag"`
}


type TrafficControlRuleSetRuleListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type TrafficControlRuleSetRuleListActionGroup struct {
    LimitPolicy int `json:"limit-policy"`
    Uuid string `json:"uuid"`
}


type TrafficControlRuleSetSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *TrafficControlRuleSet) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TrafficControlRuleSet) getPath() string{
    return "traffic-control/rule-set"
}

func (p *TrafficControlRuleSet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TrafficControlRuleSet::Post")
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

func (p *TrafficControlRuleSet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TrafficControlRuleSet::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *TrafficControlRuleSet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TrafficControlRuleSet::Put")
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

func (p *TrafficControlRuleSet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TrafficControlRuleSet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
