

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbServiceGroup struct {
	Inst struct {

    BackupServerEventLog int `json:"backup-server-event-log"`
    ConnRate int `json:"conn-rate"`
    ConnRateDuration int `json:"conn-rate-duration"`
    ConnRateGracePeriod int `json:"conn-rate-grace-period"`
    ConnRateLog int `json:"conn-rate-log"`
    ConnRateRevertDuration int `json:"conn-rate-revert-duration"`
    ConnRevertRate int `json:"conn-revert-rate"`
    ExtendedStats int `json:"extended-stats"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    L4SessionRevertDuration int `json:"l4-session-revert-duration"`
    L4SessionUsage int `json:"l4-session-usage"`
    L4SessionUsageDuration int `json:"l4-session-usage-duration"`
    L4SessionUsageGracePeriod int `json:"l4-session-usage-grace-period"`
    L4SessionUsageLog int `json:"l4-session-usage-log"`
    L4SessionUsageRevertRate int `json:"l4-session-usage-revert-rate"`
    LbMethod string `json:"lb-method" dval:"round-robin"`
    LcMethod string `json:"lc-method"`
    LclbMethod string `json:"lclb-method"`
    LinkProbeTemplate string `json:"link-probe-template"`
    LlbMethod string `json:"llb-method"`
    LrprrMethod string `json:"lrprr-method"`
    MemberList []SlbServiceGroupMemberList `json:"member-list"`
    MinActiveMember int `json:"min-active-member"`
    MinActiveMemberAction string `json:"min-active-member-action"`
    Name string `json:"name"`
    PersistScoring string `json:"persist-scoring" dval:"global"`
    Priorities []SlbServiceGroupPriorities `json:"priorities"`
    PriorityAffinity int `json:"priority-affinity"`
    Protocol string `json:"protocol"`
    PseudoRoundRobin int `json:"pseudo-round-robin"`
    ReportDelay int `json:"report-delay"`
    Reset SlbServiceGroupReset1413 `json:"reset"`
    ResetOnServerSelectionFail int `json:"reset-on-server-selection-fail"`
    ResetPriorityAffinity int `json:"reset-priority-affinity"`
    RptExtServer int `json:"rpt-ext-server"`
    SampleRspTime int `json:"sample-rsp-time"`
    SamplingEnable []SlbServiceGroupSamplingEnable `json:"sampling-enable"`
    SharedPartitionPolicyTemplate int `json:"shared-partition-policy-template"`
    SharedPartitionSvcgrpHealthCheck int `json:"shared-partition-svcgrp-health-check"`
    StatelessAutoSwitch int `json:"stateless-auto-switch"`
    StatelessLbMethod string `json:"stateless-lb-method"`
    StatelessLbMethod2 string `json:"stateless-lb-method2"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    StrictSelect int `json:"strict-select"`
    SvcgrpHealthCheckShared string `json:"svcgrp-health-check-shared"`
    TemplatePolicy string `json:"template-policy"`
    TemplatePolicyShared string `json:"template-policy-shared"`
    TemplatePort string `json:"template-port"`
    TopFastest int `json:"top-fastest"`
    TopSlowest int `json:"top-slowest"`
    TrafficReplicationMirror int `json:"traffic-replication-mirror"`
    TrafficReplicationMirrorDaRepl int `json:"traffic-replication-mirror-da-repl"`
    TrafficReplicationMirrorIpRepl int `json:"traffic-replication-mirror-ip-repl"`
    TrafficReplicationMirrorSaDaRepl int `json:"traffic-replication-mirror-sa-da-repl"`
    TrafficReplicationMirrorSaRepl int `json:"traffic-replication-mirror-sa-repl"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service-group"`
}


type SlbServiceGroupMemberList struct {
    Name string `json:"name"`
    Port int `json:"port"`
    FqdnName string `json:"fqdn-name"`
    ResolveAs string `json:"resolve-as" dval:"resolve-to-ipv4"`
    Host string `json:"host"`
    ServerIpv6Addr string `json:"server-ipv6-addr"`
    MemberState string `json:"member-state" dval:"enable"`
    MemberStatsDataDisable int `json:"member-stats-data-disable"`
    MemberTemplate string `json:"member-template"`
    MemberPriority int `json:"member-priority" dval:"1"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbServiceGroupMemberListSamplingEnable `json:"sampling-enable"`
}


type SlbServiceGroupMemberListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbServiceGroupPriorities struct {
    Priority int `json:"priority"`
    PriorityAction string `json:"priority-action" dval:"proceed"`
}


type SlbServiceGroupReset1413 struct {
    AutoSwitch int `json:"auto-switch"`
}


type SlbServiceGroupSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbServiceGroup) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbServiceGroup) getPath() string{
    return "slb/service-group"
}

func (p *SlbServiceGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroup::Post")
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

func (p *SlbServiceGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroup::Get")
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
func (p *SlbServiceGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroup::Put")
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

func (p *SlbServiceGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServiceGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
