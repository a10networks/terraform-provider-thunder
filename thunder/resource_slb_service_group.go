package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbServiceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_service_group`: Service Group\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbServiceGroupCreate,
		UpdateContext: resourceSlbServiceGroupUpdate,
		ReadContext:   resourceSlbServiceGroupRead,
		DeleteContext: resourceSlbServiceGroupDelete,

		Schema: map[string]*schema.Schema{
			"backup_server_event_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send log info on back up server events",
			},
			"conn_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Dynamically enable stateless method by conn-rate (Rate to trigger stateless method(conn/sec))",
			},
			"conn_rate_duration": {
				Type: schema.TypeInt, Optional: true, Description: "Period that trigger condition consistently happens(seconds)",
			},
			"conn_rate_grace_period": {
				Type: schema.TypeInt, Optional: true, Description: "Define the grace period during transition (Define the grace period during transition(seconds))",
			},
			"conn_rate_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send log if transition happens",
			},
			"conn_rate_revert_duration": {
				Type: schema.TypeInt, Optional: true, Description: "Period that revert condition consistently happens(seconds)",
			},
			"conn_revert_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Rate to revert to statelful method (conn/sec)",
			},
			"extended_stats": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable extended statistics on service group",
			},
			"health_check": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
			},
			"health_check_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable health check",
			},
			"l4_session_revert_duration": {
				Type: schema.TypeInt, Optional: true, Description: "Period that revert condition consistently happens(seconds)",
			},
			"l4_session_usage": {
				Type: schema.TypeInt, Optional: true, Description: "Dynamically enable stateless method by session usage (Usage to trigger stateless method)",
			},
			"l4_session_usage_duration": {
				Type: schema.TypeInt, Optional: true, Description: "Period that trigger condition consistently happens(seconds)",
			},
			"l4_session_usage_grace_period": {
				Type: schema.TypeInt, Optional: true, Description: "Define the grace period during transition (Define the grace period during transition(seconds))",
			},
			"l4_session_usage_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send log if transition happens",
			},
			"l4_session_usage_revert_rate": {
				Type: schema.TypeInt, Optional: true, Description: "Usage to revert to statelful method",
			},
			"lb_method": {
				Type: schema.TypeString, Optional: true, Default: "round-robin", Description: "'dst-ip-hash': Load-balancing based on only Dst IP and Port hash; 'dst-ip-only-hash': Load-balancing based on only Dst IP hash; 'fastest-response': Fastest response time on service port level; 'least-request': Least request on service port level; 'src-ip-hash': Load-balancing based on only Src IP and Port hash; 'src-ip-only-hash': Load-balancing based on only Src IP hash; 'weighted-rr': Weighted round robin on server level; 'service-weighted-rr': Weighted round robin on service port level; 'round-robin': Round robin on server level; 'round-robin-strict': Strict mode round robin on server level; 'odd-even-hash': odd/even hash based of client src-ip;",
			},
			"lc_method": {
				Type: schema.TypeString, Optional: true, Description: "'least-connection': Least connection on server level; 'service-least-connection': Least connection on service port level; 'weighted-least-connection': Weighted least connection on server level; 'service-weighted-least-connection': Weighted least connection on service port level;",
			},
			"lclb_method": {
				Type: schema.TypeString, Optional: true, Description: "'link-cost-load-balance': Link cost load balance;",
			},
			"link_probe_template": {
				Type: schema.TypeString, Optional: true, Description: "Link Probe template (Link Probe template name)",
			},
			"llb_method": {
				Type: schema.TypeString, Optional: true, Description: "'next-hop-link': Server selection w/ link probe template on service port level;",
			},
			"lrprr_method": {
				Type: schema.TypeString, Optional: true, Description: "'service-least-request-pseudo-round-robin': Least request on service port level and select the oldest node for sub-select;",
			},
			"member_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Member name",
						},
						"port": {
							Type: schema.TypeInt, Required: true, Description: "Port number",
						},
						"fqdn_name": {
							Type: schema.TypeString, Optional: true, Description: "Server hostname - Not applicable if real server is already defined",
						},
						"resolve_as": {
							Type: schema.TypeString, Optional: true, Default: "resolve-to-ipv4", Description: "'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;",
						},
						"host": {
							Type: schema.TypeString, Optional: true, Description: "IP Address - Not applicable if real server is already defined",
						},
						"server_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 Address - Not applicable if real server is already defined",
						},
						"member_state": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable member service port; 'disable': Disable member service port; 'disable-with-health-check': disable member service port, but health check work;",
						},
						"member_stats_data_disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable statistical data collection",
						},
						"member_template": {
							Type: schema.TypeString, Optional: true, Description: "Real server port template (Real server port template name)",
						},
						"member_priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Priority of Port in the Group (Priority of Port in the Group, default is 1)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"sampling_enable": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total established connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests successful; 'peak_conn': Peak connections; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count; 'state_flaps': State flaps count;",
									},
								},
							},
						},
					},
				},
			},
			"min_active_member": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum Active Member Per Priority (Minimum Active Member before Action)",
			},
			"min_active_member_action": {
				Type: schema.TypeString, Optional: true, Description: "'dynamic-priority': dynamic change member priority to met the min-active-member requirement; 'skip-pri-set': Skip Current Priority Set If Min not met;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "SLB Service Name",
			},
			"persist_scoring": {
				Type: schema.TypeString, Optional: true, Default: "global", Description: "'global': Use Global Configuration; 'enable': Enable persist-scoring; 'disable': Disable persist-scoring;",
			},
			"priorities": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "Priority option. Define different action for each priority node. (Priority in the Group)",
						},
						"priority_action": {
							Type: schema.TypeString, Optional: true, Default: "proceed", Description: "'drop': Drop request when all priority nodes fail; 'drop-if-exceed-limit': Drop request when connection over limit; 'proceed': Proceed to next priority when all priority nodes fail(default); 'reset': Send client reset when all priority nodes fail; 'reset-if-exceed-limit': Send client reset when connection over limit;",
						},
					},
				},
			},
			"priority_affinity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Priority affinity. Persist to the same priority if possible.",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Description: "'tcp': TCP LB service; 'udp': UDP LB service;",
			},
			"pseudo_round_robin": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "PRR, select the oldest node for sub-select",
			},
			"report_delay": {
				Type: schema.TypeInt, Optional: true, Description: "Reporting frequency (in minutes)",
			},
			"reset": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset auto stateless state",
						},
					},
				},
			},
			"reset_on_server_selection_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send reset to client if server selection fails",
			},
			"reset_priority_affinity": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reset",
			},
			"rpt_ext_server": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Report top 10 fastest/slowest servers",
			},
			"sample_rsp_time": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "sample server response time",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'server_selection_fail_drop': Drops due to Service selection failure; 'server_selection_fail_reset': Resets sent out for Service selection failure; 'service_peak_conn': Peak connection count for the Service Group; 'service_healthy_host': Service Group healthy host count; 'service_unhealthy_host': Service Group unhealthy host count; 'service_req_count': Service Group request count; 'service_resp_count': Service Group response count; 'service_resp_2xx': Service Group response 2xx count; 'service_resp_3xx': Service Group response 3xx count; 'service_resp_4xx': Service Group response 4xx count; 'service_resp_5xx': Service Group response 5xx count; 'service_curr_conn_overflow': Current connection counter overflow count;",
						},
					},
				},
			},
			"shared_partition_policy_template": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a policy template from shared partition",
			},
			"shared_partition_svcgrp_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reference a health-check from shared partition",
			},
			"stateless_auto_switch": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable auto stateless method",
			},
			"stateless_lb_method": {
				Type: schema.TypeString, Optional: true, Description: "'stateless-dst-ip-hash': Stateless load-balancing based on Dst IP and Dst port hash; 'stateless-per-pkt-round-robin': Stateless load-balancing using per-packet round-robin; 'stateless-src-dst-ip-hash': Stateless load-balancing based on IP and port hash for both Src and Dst; 'stateless-src-dst-ip-only-hash': Stateless load-balancing based on only IP hash for both Src and Dst; 'stateless-src-ip-hash': Stateless load-balancing based on Src IP and Src port hash; 'stateless-src-ip-only-hash': Stateless load-balancing based on only Src IP hash; 'stateless-per-pkt-weighted-rr': Stateless load-balancing using per-packet weighted round robin on server level; 'stateless-per-pkt-service-weighted-rr': Stateless load-balancing using per-packet weighted round robin on service port level;",
			},
			"stateless_lb_method2": {
				Type: schema.TypeString, Optional: true, Description: "'stateless-dst-ip-hash': Stateless load-balancing based on Dst IP and Dst port hash; 'stateless-per-pkt-round-robin': Stateless load-balancing using per-packet round-robin; 'stateless-src-dst-ip-hash': Stateless load-balancing based on IP and port hash for both Src and Dst; 'stateless-src-dst-ip-only-hash': Stateless load-balancing based on only IP hash for both Src and Dst; 'stateless-src-ip-hash': Stateless load-balancing based on Src IP and Src port hash; 'stateless-src-ip-only-hash': Stateless load-balancing based on only Src IP hash; 'stateless-per-pkt-weighted-rr': Stateless load-balancing using per-packet weighted round robin on server level; 'stateless-per-pkt-service-weighted-rr': Stateless load-balancing using per-packet weighted round robin on service port level;",
			},
			"stats_data_action": {
				Type: schema.TypeString, Optional: true, Default: "stats-data-enable", Description: "'stats-data-enable': Enable statistical data collection for service group; 'stats-data-disable': Disable statistical data collection for service group;",
			},
			"strict_select": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "strict selection",
			},
			"svcgrp_health_check_shared": {
				Type: schema.TypeString, Optional: true, Description: "Health Check (Monitor Name)",
			},
			"template_policy": {
				Type: schema.TypeString, Optional: true, Description: "Policy template (Policy template name)",
			},
			"template_policy_shared": {
				Type: schema.TypeString, Optional: true, Description: "Policy template",
			},
			"template_port": {
				Type: schema.TypeString, Optional: true, Description: "Port template (Port template name)",
			},
			"top_fastest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Report top 10 fastest servers",
			},
			"top_slowest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Report top 10 slowest servers",
			},
			"traffic_replication_mirror": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mirror Bi-directional Packet",
			},
			"traffic_replication_mirror_da_repl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace Destination MAC",
			},
			"traffic_replication_mirror_ip_repl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replaces IP with server-IP",
			},
			"traffic_replication_mirror_sa_da_repl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace Source MAC and Destination MAC",
			},
			"traffic_replication_mirror_sa_repl": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Replace Source MAC",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbServiceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServiceGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbServiceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbServiceGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbServiceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbServiceGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbServiceGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbServiceGroupMemberList(d []interface{}) []edpt.SlbServiceGroupMemberList {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupMemberList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupMemberList
		oi.Name = in["name"].(string)
		oi.Port = in["port"].(int)
		oi.FqdnName = in["fqdn_name"].(string)
		oi.ResolveAs = in["resolve_as"].(string)
		oi.Host = in["host"].(string)
		oi.ServerIpv6Addr = in["server_ipv6_addr"].(string)
		oi.MemberState = in["member_state"].(string)
		oi.MemberStatsDataDisable = in["member_stats_data_disable"].(int)
		oi.MemberTemplate = in["member_template"].(string)
		oi.MemberPriority = in["member_priority"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceSlbServiceGroupMemberListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServiceGroupMemberListSamplingEnable(d []interface{}) []edpt.SlbServiceGroupMemberListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupMemberListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupMemberListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbServiceGroupPriorities(d []interface{}) []edpt.SlbServiceGroupPriorities {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupPriorities, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupPriorities
		oi.Priority = in["priority"].(int)
		oi.PriorityAction = in["priority_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbServiceGroupReset1413(d []interface{}) edpt.SlbServiceGroupReset1413 {

	count1 := len(d)
	var ret edpt.SlbServiceGroupReset1413
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AutoSwitch = in["auto_switch"].(int)
	}
	return ret
}

func getSliceSlbServiceGroupSamplingEnable(d []interface{}) []edpt.SlbServiceGroupSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbServiceGroupSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbServiceGroupSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbServiceGroup(d *schema.ResourceData) edpt.SlbServiceGroup {
	var ret edpt.SlbServiceGroup
	ret.Inst.BackupServerEventLog = d.Get("backup_server_event_log").(int)
	ret.Inst.ConnRate = d.Get("conn_rate").(int)
	ret.Inst.ConnRateDuration = d.Get("conn_rate_duration").(int)
	ret.Inst.ConnRateGracePeriod = d.Get("conn_rate_grace_period").(int)
	ret.Inst.ConnRateLog = d.Get("conn_rate_log").(int)
	ret.Inst.ConnRateRevertDuration = d.Get("conn_rate_revert_duration").(int)
	ret.Inst.ConnRevertRate = d.Get("conn_revert_rate").(int)
	ret.Inst.ExtendedStats = d.Get("extended_stats").(int)
	ret.Inst.HealthCheck = d.Get("health_check").(string)
	ret.Inst.HealthCheckDisable = d.Get("health_check_disable").(int)
	ret.Inst.L4SessionRevertDuration = d.Get("l4_session_revert_duration").(int)
	ret.Inst.L4SessionUsage = d.Get("l4_session_usage").(int)
	ret.Inst.L4SessionUsageDuration = d.Get("l4_session_usage_duration").(int)
	ret.Inst.L4SessionUsageGracePeriod = d.Get("l4_session_usage_grace_period").(int)
	ret.Inst.L4SessionUsageLog = d.Get("l4_session_usage_log").(int)
	ret.Inst.L4SessionUsageRevertRate = d.Get("l4_session_usage_revert_rate").(int)
	ret.Inst.LbMethod = d.Get("lb_method").(string)
	ret.Inst.LcMethod = d.Get("lc_method").(string)
	ret.Inst.LclbMethod = d.Get("lclb_method").(string)
	ret.Inst.LinkProbeTemplate = d.Get("link_probe_template").(string)
	ret.Inst.LlbMethod = d.Get("llb_method").(string)
	ret.Inst.LrprrMethod = d.Get("lrprr_method").(string)
	ret.Inst.MemberList = getSliceSlbServiceGroupMemberList(d.Get("member_list").([]interface{}))
	ret.Inst.MinActiveMember = d.Get("min_active_member").(int)
	ret.Inst.MinActiveMemberAction = d.Get("min_active_member_action").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PersistScoring = d.Get("persist_scoring").(string)
	ret.Inst.Priorities = getSliceSlbServiceGroupPriorities(d.Get("priorities").([]interface{}))
	ret.Inst.PriorityAffinity = d.Get("priority_affinity").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.PseudoRoundRobin = d.Get("pseudo_round_robin").(int)
	ret.Inst.ReportDelay = d.Get("report_delay").(int)
	ret.Inst.Reset = getObjectSlbServiceGroupReset1413(d.Get("reset").([]interface{}))
	ret.Inst.ResetOnServerSelectionFail = d.Get("reset_on_server_selection_fail").(int)
	ret.Inst.ResetPriorityAffinity = d.Get("reset_priority_affinity").(int)
	ret.Inst.RptExtServer = d.Get("rpt_ext_server").(int)
	ret.Inst.SampleRspTime = d.Get("sample_rsp_time").(int)
	ret.Inst.SamplingEnable = getSliceSlbServiceGroupSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	ret.Inst.SharedPartitionSvcgrpHealthCheck = d.Get("shared_partition_svcgrp_health_check").(int)
	ret.Inst.StatelessAutoSwitch = d.Get("stateless_auto_switch").(int)
	ret.Inst.StatelessLbMethod = d.Get("stateless_lb_method").(string)
	ret.Inst.StatelessLbMethod2 = d.Get("stateless_lb_method2").(string)
	ret.Inst.StatsDataAction = d.Get("stats_data_action").(string)
	ret.Inst.StrictSelect = d.Get("strict_select").(int)
	ret.Inst.SvcgrpHealthCheckShared = d.Get("svcgrp_health_check_shared").(string)
	ret.Inst.TemplatePolicy = d.Get("template_policy").(string)
	ret.Inst.TemplatePolicyShared = d.Get("template_policy_shared").(string)
	ret.Inst.TemplatePort = d.Get("template_port").(string)
	ret.Inst.TopFastest = d.Get("top_fastest").(int)
	ret.Inst.TopSlowest = d.Get("top_slowest").(int)
	ret.Inst.TrafficReplicationMirror = d.Get("traffic_replication_mirror").(int)
	ret.Inst.TrafficReplicationMirrorDaRepl = d.Get("traffic_replication_mirror_da_repl").(int)
	ret.Inst.TrafficReplicationMirrorIpRepl = d.Get("traffic_replication_mirror_ip_repl").(int)
	ret.Inst.TrafficReplicationMirrorSaDaRepl = d.Get("traffic_replication_mirror_sa_da_repl").(int)
	ret.Inst.TrafficReplicationMirrorSaRepl = d.Get("traffic_replication_mirror_sa_repl").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
