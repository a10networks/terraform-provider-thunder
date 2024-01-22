

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type RuleSet struct {
	Inst struct {

    App RuleSetApp1318 `json:"app"`
    Application RuleSetApplication1319 `json:"application"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Remark string `json:"remark"`
    RuleList []RuleSetRuleList `json:"rule-list"`
    RulesByZone RuleSetRulesByZone1320 `json:"rules-by-zone"`
    SamplingEnable []RuleSetSamplingEnable `json:"sampling-enable"`
    SessionStatistic string `json:"session-statistic" dval:"enable"`
    Tag RuleSetTag1322 `json:"tag"`
    TrackAppRuleList RuleSetTrackAppRuleList1323 `json:"track-app-rule-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"rule-set"`
}


type RuleSetApp1318 struct {
    Uuid string `json:"uuid"`
}


type RuleSetApplication1319 struct {
    Uuid string `json:"uuid"`
}


type RuleSetRuleList struct {
    Name string `json:"name"`
    Remark string `json:"remark"`
    Status string `json:"status" dval:"enable"`
    IpVersion string `json:"ip-version" dval:"v4"`
    Action string `json:"action"`
    Log int `json:"log"`
    ResetLid int `json:"reset-lid"`
    ListenOnPort int `json:"listen-on-port"`
    Policy string `json:"policy"`
    VpnIpsecName string `json:"vpn-ipsec-name"`
    VpnIpsecGroupName string `json:"vpn-ipsec-group-name"`
    ForwardListenOnPort int `json:"forward-listen-on-port"`
    Lid int `json:"lid"`
    ListenOnPortLid int `json:"listen-on-port-lid"`
    FwLog int `json:"fw-log"`
    Fwlog int `json:"fwlog"`
    Cgnv6Log int `json:"cgnv6-log"`
    ForwardLog int `json:"forward-log"`
    Lidlog int `json:"lidlog"`
    ResetLidlog int `json:"reset-lidlog"`
    ListenOnPortLidlog int `json:"listen-on-port-lidlog"`
    Cgnv6Policy string `json:"cgnv6-policy"`
    Cgnv6FixedNatLog int `json:"cgnv6-fixed-nat-log"`
    Cgnv6LsnLid int `json:"cgnv6-lsn-lid"`
    Cgnv6DsLite string `json:"cgnv6-ds-lite"`
    Cgnv6DsLiteLsnLid int `json:"cgnv6-ds-lite-lsn-lid"`
    InspectPayload int `json:"inspect-payload"`
    Cgnv6DsLiteLog int `json:"cgnv6-ds-lite-log"`
    Cgnv6LsnLog int `json:"cgnv6-lsn-log"`
    GtpTemplate string `json:"gtp-template"`
    SrcGeolocName string `json:"src-geoloc-name"`
    SrcGeolocList string `json:"src-geoloc-list"`
    SrcGeolocListShared int `json:"src-geoloc-list-shared"`
    SrcIpv4Any string `json:"src-ipv4-any" dval:"any"`
    SrcIpv6Any string `json:"src-ipv6-any" dval:"any"`
    SrcClassList string `json:"src-class-list"`
    SourceList []RuleSetRuleListSourceList `json:"source-list"`
    SrcZone string `json:"src-zone"`
    SrcZoneAny string `json:"src-zone-any" dval:"any"`
    SrcThreatList string `json:"src-threat-list"`
    DstGeolocName string `json:"dst-geoloc-name"`
    DstGeolocList string `json:"dst-geoloc-list"`
    DstGeolocListShared int `json:"dst-geoloc-list-shared"`
    DstIpv4Any string `json:"dst-ipv4-any" dval:"any"`
    DstIpv6Any string `json:"dst-ipv6-any" dval:"any"`
    DstClassList string `json:"dst-class-list"`
    DestList []RuleSetRuleListDestList `json:"dest-list"`
    DstDomainList string `json:"dst-domain-list"`
    DstZone string `json:"dst-zone"`
    DstZoneAny string `json:"dst-zone-any" dval:"any"`
    DstThreatList string `json:"dst-threat-list"`
    ServiceAny string `json:"service-any" dval:"any"`
    ServiceList []RuleSetRuleListServiceList `json:"service-list"`
    IdleTimeout int `json:"idle-timeout"`
    DscpList []RuleSetRuleListDscpList `json:"dscp-list"`
    ApplicationAny string `json:"application-any" dval:"any"`
    AppList []RuleSetRuleListAppList `json:"app-list"`
    TrackApplication int `json:"track-application"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []RuleSetRuleListSamplingEnable `json:"sampling-enable"`
    ActionGroup RuleSetRuleListActionGroup `json:"action-group"`
    MoveRule RuleSetRuleListMoveRule `json:"move-rule"`
}


type RuleSetRuleListSourceList struct {
    SrcIpSubnet string `json:"src-ip-subnet"`
    SrcIpv6Subnet string `json:"src-ipv6-subnet"`
    SrcObjNetwork string `json:"src-obj-network"`
    SrcObjGrpNetwork string `json:"src-obj-grp-network"`
    SrcSlbServer string `json:"src-slb-server"`
}


type RuleSetRuleListDestList struct {
    DstIpSubnet string `json:"dst-ip-subnet"`
    DstIpv6Subnet string `json:"dst-ipv6-subnet"`
    DstObjNetwork string `json:"dst-obj-network"`
    DstObjGrpNetwork string `json:"dst-obj-grp-network"`
    DstSlbServer string `json:"dst-slb-server"`
    DstSlbVserver string `json:"dst-slb-vserver"`
}


type RuleSetRuleListServiceList struct {
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
    Alg string `json:"alg"`
}


type RuleSetRuleListDscpList struct {
    DscpValue string `json:"dscp-value"`
    DscpRangeStart int `json:"dscp-range-start"`
    DscpRangeEnd int `json:"dscp-range-end"`
}


type RuleSetRuleListAppList struct {
    ObjGrpApplication string `json:"obj-grp-application"`
    Protocol string `json:"protocol"`
    ProtocolTag string `json:"protocol-tag"`
}


type RuleSetRuleListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type RuleSetRuleListActionGroup struct {
    Type string `json:"type"`
    PermitLog int `json:"permit-log"`
    ResetLog int `json:"reset-log"`
    DenyLog int `json:"deny-log"`
    LoggingTemplateList []RuleSetRuleListActionGroupLoggingTemplateList `json:"logging-template-list"`
    ResetLogTemplateType string `json:"reset-log-template-type"`
    ResetFwLog string `json:"reset-fw-log"`
    DenyLogTemplateType string `json:"deny-log-template-type"`
    DenyFwLog string `json:"deny-fw-log"`
    ListenOnPort int `json:"listen-on-port"`
    Forward int `json:"forward"`
    Ipsec int `json:"ipsec"`
    IpsecGroup int `json:"ipsec-group"`
    VpnIpsecName string `json:"vpn-ipsec-name"`
    VpnIpsecGroupName string `json:"vpn-ipsec-group-name"`
    Cgnv6 int `json:"cgnv6"`
    Cgnv6Policy string `json:"cgnv6-policy"`
    Cgnv6LsnLid int `json:"cgnv6-lsn-lid"`
    Cgnv6DsLite string `json:"cgnv6-ds-lite"`
    Cgnv6DsLiteLsnLid int `json:"cgnv6-ds-lite-lsn-lid"`
    InspectPayload int `json:"inspect-payload"`
    PermitLimitPolicy int `json:"permit-limit-policy"`
    PermitRespondToUserMac int `json:"permit-respond-to-user-mac"`
    ResetRespondToUserMac int `json:"reset-respond-to-user-mac"`
    SetDscp int `json:"set-dscp"`
    DscpValue string `json:"dscp-value"`
    DscpNumber int `json:"dscp-number"`
    Uuid string `json:"uuid"`
}


type RuleSetRuleListActionGroupLoggingTemplateList struct {
    PermitLogTemplateType string `json:"permit-log-template-type"`
    PermitFwLog string `json:"permit-fw-log"`
    PermitCgnv6Log string `json:"permit-cgnv6-log"`
    PermitNetflowLog string `json:"permit-netflow-log"`
}


type RuleSetRuleListMoveRule struct {
    Location string `json:"location" dval:"bottom"`
    TargetRule string `json:"target-rule"`
}


type RuleSetRulesByZone1320 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []RuleSetRulesByZoneSamplingEnable1321 `json:"sampling-enable"`
}


type RuleSetRulesByZoneSamplingEnable1321 struct {
    Counters1 string `json:"counters1"`
}


type RuleSetSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type RuleSetTag1322 struct {
    Uuid string `json:"uuid"`
}


type RuleSetTrackAppRuleList1323 struct {
    Uuid string `json:"uuid"`
}

func (p *RuleSet) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *RuleSet) getPath() string{
    return "rule-set"
}

func (p *RuleSet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSet::Post")
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

func (p *RuleSet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSet::Get")
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
func (p *RuleSet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSet::Put")
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

func (p *RuleSet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
