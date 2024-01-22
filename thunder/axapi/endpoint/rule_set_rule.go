

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type RuleSetRule struct {
	Inst struct {

    Action string `json:"action"`
    ActionGroup RuleSetRuleActionGroup1315 `json:"action-group"`
    AppList []RuleSetRuleAppList `json:"app-list"`
    ApplicationAny string `json:"application-any" dval:"any"`
    Cgnv6DsLite string `json:"cgnv6-ds-lite"`
    Cgnv6DsLiteLog int `json:"cgnv6-ds-lite-log"`
    Cgnv6DsLiteLsnLid int `json:"cgnv6-ds-lite-lsn-lid"`
    Cgnv6FixedNatLog int `json:"cgnv6-fixed-nat-log"`
    Cgnv6Log int `json:"cgnv6-log"`
    Cgnv6LsnLid int `json:"cgnv6-lsn-lid"`
    Cgnv6LsnLog int `json:"cgnv6-lsn-log"`
    Cgnv6Policy string `json:"cgnv6-policy"`
    DestList []RuleSetRuleDestList `json:"dest-list"`
    DscpList []RuleSetRuleDscpList `json:"dscp-list"`
    DstClassList string `json:"dst-class-list"`
    DstDomainList string `json:"dst-domain-list"`
    DstGeolocList string `json:"dst-geoloc-list"`
    DstGeolocListShared int `json:"dst-geoloc-list-shared"`
    DstGeolocName string `json:"dst-geoloc-name"`
    DstIpv4Any string `json:"dst-ipv4-any" dval:"any"`
    DstIpv6Any string `json:"dst-ipv6-any" dval:"any"`
    DstThreatList string `json:"dst-threat-list"`
    DstZone string `json:"dst-zone"`
    DstZoneAny string `json:"dst-zone-any" dval:"any"`
    ForwardListenOnPort int `json:"forward-listen-on-port"`
    ForwardLog int `json:"forward-log"`
    FwLog int `json:"fw-log"`
    Fwlog int `json:"fwlog"`
    GtpTemplate string `json:"gtp-template"`
    IdleTimeout int `json:"idle-timeout"`
    InspectPayload int `json:"inspect-payload"`
    IpVersion string `json:"ip-version" dval:"v4"`
    Lid int `json:"lid"`
    Lidlog int `json:"lidlog"`
    ListenOnPort int `json:"listen-on-port"`
    ListenOnPortLid int `json:"listen-on-port-lid"`
    ListenOnPortLidlog int `json:"listen-on-port-lidlog"`
    Log int `json:"log"`
    MoveRule RuleSetRuleMoveRule1317 `json:"move-rule"`
    Name string `json:"name"`
    Policy string `json:"policy"`
    Remark string `json:"remark"`
    ResetLid int `json:"reset-lid"`
    ResetLidlog int `json:"reset-lidlog"`
    SamplingEnable []RuleSetRuleSamplingEnable `json:"sampling-enable"`
    ServiceAny string `json:"service-any" dval:"any"`
    ServiceList []RuleSetRuleServiceList `json:"service-list"`
    SourceList []RuleSetRuleSourceList `json:"source-list"`
    SrcClassList string `json:"src-class-list"`
    SrcGeolocList string `json:"src-geoloc-list"`
    SrcGeolocListShared int `json:"src-geoloc-list-shared"`
    SrcGeolocName string `json:"src-geoloc-name"`
    SrcIpv4Any string `json:"src-ipv4-any" dval:"any"`
    SrcIpv6Any string `json:"src-ipv6-any" dval:"any"`
    SrcThreatList string `json:"src-threat-list"`
    SrcZone string `json:"src-zone"`
    SrcZoneAny string `json:"src-zone-any" dval:"any"`
    Status string `json:"status" dval:"enable"`
    TrackApplication int `json:"track-application"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VpnIpsecGroupName string `json:"vpn-ipsec-group-name"`
    VpnIpsecName string `json:"vpn-ipsec-name"`

	} `json:"rule"`
}


type RuleSetRuleActionGroup1315 struct {
    Type string `json:"type"`
    PermitLog int `json:"permit-log"`
    ResetLog int `json:"reset-log"`
    DenyLog int `json:"deny-log"`
    LoggingTemplateList []RuleSetRuleActionGroupLoggingTemplateList1316 `json:"logging-template-list"`
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


type RuleSetRuleActionGroupLoggingTemplateList1316 struct {
    PermitLogTemplateType string `json:"permit-log-template-type"`
    PermitFwLog string `json:"permit-fw-log"`
    PermitCgnv6Log string `json:"permit-cgnv6-log"`
    PermitNetflowLog string `json:"permit-netflow-log"`
}


type RuleSetRuleAppList struct {
    ObjGrpApplication string `json:"obj-grp-application"`
    Protocol string `json:"protocol"`
    ProtocolTag string `json:"protocol-tag"`
}


type RuleSetRuleDestList struct {
    DstIpSubnet string `json:"dst-ip-subnet"`
    DstIpv6Subnet string `json:"dst-ipv6-subnet"`
    DstObjNetwork string `json:"dst-obj-network"`
    DstObjGrpNetwork string `json:"dst-obj-grp-network"`
    DstSlbServer string `json:"dst-slb-server"`
    DstSlbVserver string `json:"dst-slb-vserver"`
}


type RuleSetRuleDscpList struct {
    DscpValue string `json:"dscp-value"`
    DscpRangeStart int `json:"dscp-range-start"`
    DscpRangeEnd int `json:"dscp-range-end"`
}


type RuleSetRuleMoveRule1317 struct {
    Location string `json:"location" dval:"bottom"`
    TargetRule string `json:"target-rule"`
}


type RuleSetRuleSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type RuleSetRuleServiceList struct {
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


type RuleSetRuleSourceList struct {
    SrcIpSubnet string `json:"src-ip-subnet"`
    SrcIpv6Subnet string `json:"src-ipv6-subnet"`
    SrcObjNetwork string `json:"src-obj-network"`
    SrcObjGrpNetwork string `json:"src-obj-grp-network"`
    SrcSlbServer string `json:"src-slb-server"`
}

func (p *RuleSetRule) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *RuleSetRule) getPath() string{
    return "rule-set/rule"
}

func (p *RuleSetRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRule::Post")
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

func (p *RuleSetRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRule::Get")
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
func (p *RuleSetRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRule::Put")
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

func (p *RuleSetRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
