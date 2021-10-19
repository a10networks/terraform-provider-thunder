package thunder

//Thunder resource ServiceGroup

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceGroupCreate,
		UpdateContext: resourceServiceGroupUpdate,
		ReadContext:   resourceServiceGroupRead,
		DeleteContext: resourceServiceGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"template_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_policy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_policy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"lb_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"lc_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"stateless_lb_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"llb_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"link_probe_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"lclb_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pseudo_round_robin": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_auto_switch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_lb_method2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_revert_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_revert_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_grace_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_revert_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_revert_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_grace_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_active_member": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_active_member_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"reset_on_server_selection_fail": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"priority_affinity": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"reset_priority_affinity": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"backup_server_event_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"strict_select": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror_da_repl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror_ip_repl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror_sa_da_repl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror_sa_repl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"shared_partition_svcgrp_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"svcgrp_health_check_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"priorities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"priority_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"sample_rsp_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rpt_ext_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"report_delay": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"top_slowest": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"top_fastest": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"persist_scoring": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"reset": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_switch": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"member_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"fqdn_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"resolve_as": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_ipv6_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"member_state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"member_stats_data_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"member_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"member_priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"sampling_enable": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"counters1": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceServiceGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating ServiceGroup (Inside resourceServiceGroupCreate) ")
		name1 := d.Get("name").(string)
		data := dataToServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to ServiceGroup --")
		d.SetId(name1)
		err := go_thunder.PostServiceGroup(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceServiceGroupRead(ctx, d, meta)

	}
	return diags
}

func resourceServiceGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading ServiceGroup (Inside resourceServiceGroupRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetServiceGroup(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceServiceGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying ServiceGroup   (Inside resourceServiceGroupUpdate) ")
		data := dataToServiceGroup(d)
		logger.Println("[INFO] received formatted data from method data to ServiceGroup ")
		err := go_thunder.PutServiceGroup(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceServiceGroupRead(ctx, d, meta)

	}
	return diags
}

func resourceServiceGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceServiceGroupDelete) " + name1)
		err := go_thunder.DeleteServiceGroup(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToServiceGroup(d *schema.ResourceData) go_thunder.ServiceGroup {
	var vc go_thunder.ServiceGroup
	var c go_thunder.ServiceGroupInstance
	c.ServiceGroupInstanceName = d.Get("name").(string)
	c.ServiceGroupInstanceProtocol = d.Get("protocol").(string)
	c.ServiceGroupInstanceTemplatePort = d.Get("template_port").(string)
	c.ServiceGroupInstanceTemplateServer = d.Get("template_server").(string)
	c.ServiceGroupInstanceTemplatePolicy = d.Get("template_policy").(string)
	c.ServiceGroupInstanceSharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	c.ServiceGroupInstanceTemplatePolicyShared = d.Get("template_policy_shared").(string)
	c.ServiceGroupInstanceLbMethod = d.Get("lb_method").(string)
	c.ServiceGroupInstanceLcMethod = d.Get("lc_method").(string)
	c.ServiceGroupInstanceStatelessLbMethod = d.Get("stateless_lb_method").(string)
	c.ServiceGroupInstanceLlbMethod = d.Get("llb_method").(string)
	c.ServiceGroupInstanceLinkProbeTemplate = d.Get("link_probe_template").(string)
	c.ServiceGroupInstanceLclbMethod = d.Get("lclb_method").(string)
	c.ServiceGroupInstancePseudoRoundRobin = d.Get("pseudo_round_robin").(int)
	c.ServiceGroupInstanceStatelessAutoSwitch = d.Get("stateless_auto_switch").(int)
	c.ServiceGroupInstanceStatelessLbMethod2 = d.Get("stateless_lb_method2").(string)
	c.ServiceGroupInstanceConnRate = d.Get("conn_rate").(int)
	c.ServiceGroupInstanceConnRateDuration = d.Get("conn_rate_duration").(int)
	c.ServiceGroupInstanceConnRevertRate = d.Get("conn_revert_rate").(int)
	c.ServiceGroupInstanceConnRateRevertDuration = d.Get("conn_rate_revert_duration").(int)
	c.ServiceGroupInstanceConnRateGracePeriod = d.Get("conn_rate_grace_period").(int)
	c.ServiceGroupInstanceConnRateLog = d.Get("conn_rate_log").(int)
	c.ServiceGroupInstanceL4SessionUsage = d.Get("l4_session_usage").(int)
	c.ServiceGroupInstanceL4SessionUsageDuration = d.Get("l4_session_usage_duration").(int)
	c.ServiceGroupInstanceL4SessionUsageRevertRate = d.Get("l4_session_usage_revert_rate").(int)
	c.ServiceGroupInstanceL4SessionRevertDuration = d.Get("l4_session_revert_duration").(int)
	c.ServiceGroupInstanceL4SessionUsageGracePeriod = d.Get("l4_session_usage_grace_period").(int)
	c.ServiceGroupInstanceL4SessionUsageLog = d.Get("l4_session_usage_log").(int)
	c.ServiceGroupInstanceMinActiveMember = d.Get("min_active_member").(int)
	c.ServiceGroupInstanceMinActiveMemberAction = d.Get("min_active_member_action").(string)
	c.ServiceGroupInstanceResetOnServerSelectionFail = d.Get("reset_on_server_selection_fail").(int)
	c.ServiceGroupInstancePriorityAffinity = d.Get("priority_affinity").(int)
	c.ServiceGroupInstanceResetPriorityAffinity = d.Get("reset_priority_affinity").(int)
	c.ServiceGroupInstanceBackupServerEventLog = d.Get("backup_server_event_log").(int)
	c.ServiceGroupInstanceStrictSelect = d.Get("strict_select").(int)
	c.ServiceGroupInstanceStatsDataAction = d.Get("stats_data_action").(string)
	c.ServiceGroupInstanceExtendedStats = d.Get("extended_stats").(int)
	c.ServiceGroupInstanceTrafficReplicationMirror = d.Get("traffic_replication_mirror").(int)
	c.ServiceGroupInstanceTrafficReplicationMirrorDaRepl = d.Get("traffic_replication_mirror_da_repl").(int)
	c.ServiceGroupInstanceTrafficReplicationMirrorIPRepl = d.Get("traffic_replication_mirror_ip_repl").(int)
	c.ServiceGroupInstanceTrafficReplicationMirrorSaDaRepl = d.Get("traffic_replication_mirror_sa_da_repl").(int)
	c.ServiceGroupInstanceTrafficReplicationMirrorSaRepl = d.Get("traffic_replication_mirror_sa_repl").(int)
	c.ServiceGroupInstanceHealthCheck = d.Get("health_check").(string)
	c.ServiceGroupInstanceSharedPartitionSvcgrpHealthCheck = d.Get("shared_partition_svcgrp_health_check").(int)
	c.ServiceGroupInstanceSvcgrpHealthCheckShared = d.Get("svcgrp_health_check_shared").(string)
	c.ServiceGroupInstanceHealthCheckDisable = d.Get("health_check_disable").(int)

	ServiceGroupInstancePrioritiesCount := d.Get("priorities.#").(int)
	c.ServiceGroupInstancePrioritiesPriority = make([]go_thunder.ServiceGroupInstancePriorities, 0, ServiceGroupInstancePrioritiesCount)

	for i := 0; i < ServiceGroupInstancePrioritiesCount; i++ {
		var obj1 go_thunder.ServiceGroupInstancePriorities
		prefix1 := fmt.Sprintf("priorities.%d.", i)
		obj1.ServiceGroupInstancePrioritiesPriority = d.Get(prefix1 + "priority").(int)
		obj1.ServiceGroupInstancePrioritiesPriorityAction = d.Get(prefix1 + "priority_action").(string)
		c.ServiceGroupInstancePrioritiesPriority = append(c.ServiceGroupInstancePrioritiesPriority, obj1)
	}

	c.ServiceGroupInstanceSampleRspTime = d.Get("sample_rsp_time").(int)
	c.ServiceGroupInstanceRptExtServer = d.Get("rpt_ext_server").(int)
	c.ServiceGroupInstanceReportDelay = d.Get("report_delay").(int)
	c.ServiceGroupInstanceTopSlowest = d.Get("top_slowest").(int)
	c.ServiceGroupInstanceTopFastest = d.Get("top_fastest").(int)
	c.ServiceGroupInstancePersistScoring = d.Get("persist_scoring").(string)
	c.ServiceGroupInstanceUserTag = d.Get("user_tag").(string)

	ServiceGroupInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.ServiceGroupInstanceSamplingEnableCounters1 = make([]go_thunder.ServiceGroupInstanceSamplingEnable, 0, ServiceGroupInstanceSamplingEnableCount)

	for i := 0; i < ServiceGroupInstanceSamplingEnableCount; i++ {
		var obj2 go_thunder.ServiceGroupInstanceSamplingEnable
		prefix2 := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.ServiceGroupInstanceSamplingEnableCounters1 = d.Get(prefix2 + "counters1").(string)
		c.ServiceGroupInstanceSamplingEnableCounters1 = append(c.ServiceGroupInstanceSamplingEnableCounters1, obj2)
	}

	var obj3 go_thunder.ServiceGroupInstanceReset
	prefix3 := "reset.0."
	obj3.ServiceGroupInstanceResetAutoSwitch = d.Get(prefix3 + "auto_switch").(int)

	c.ServiceGroupInstanceResetAutoSwitch = obj3

	ServiceGroupInstanceMemberListCount := d.Get("member_list.#").(int)
	c.ServiceGroupInstanceMemberListName = make([]go_thunder.ServiceGroupInstanceMemberList, 0, ServiceGroupInstanceMemberListCount)

	for i := 0; i < ServiceGroupInstanceMemberListCount; i++ {
		var obj4 go_thunder.ServiceGroupInstanceMemberList
		prefix4 := fmt.Sprintf("member_list.%d.", i)
		obj4.ServiceGroupInstanceMemberListName = d.Get(prefix4 + "name").(string)
		obj4.ServiceGroupInstanceMemberListPort = d.Get(prefix4 + "port").(int)
		obj4.ServiceGroupInstanceMemberListFqdnName = d.Get(prefix4 + "fqdn_name").(string)
		obj4.ServiceGroupInstanceMemberListResolveAs = d.Get(prefix4 + "resolve_as").(string)
		obj4.ServiceGroupInstanceMemberListHost = d.Get(prefix4 + "host").(string)
		obj4.ServiceGroupInstanceMemberListServerIpv6Addr = d.Get(prefix4 + "server_ipv6_addr").(string)
		obj4.ServiceGroupInstanceMemberListMemberState = d.Get(prefix4 + "member_state").(string)
		obj4.ServiceGroupInstanceMemberListMemberStatsDataDisable = d.Get(prefix4 + "member_stats_data_disable").(int)
		obj4.ServiceGroupInstanceMemberListMemberTemplate = d.Get(prefix4 + "member_template").(string)
		obj4.ServiceGroupInstanceMemberListMemberPriority = d.Get(prefix4 + "member_priority").(int)
		obj4.ServiceGroupInstanceMemberListUserTag = d.Get(prefix4 + "user_tag").(string)

		ServiceGroupInstanceMemberListSamplingEnableCount := d.Get(prefix4 + "sampling_enable.#").(int)
		obj4.ServiceGroupInstanceMemberListSamplingEnableCounters1 = make([]go_thunder.ServiceGroupInstanceMemberListSamplingEnable, 0, ServiceGroupInstanceMemberListSamplingEnableCount)

		for i := 0; i < ServiceGroupInstanceMemberListSamplingEnableCount; i++ {
			var obj4_1 go_thunder.ServiceGroupInstanceMemberListSamplingEnable
			prefix4_1 := prefix4 + fmt.Sprintf("sampling_enable.%d.", i)
			obj4_1.ServiceGroupInstanceMemberListSamplingEnableCounters1 = d.Get(prefix4_1 + "counters1").(string)
			obj4.ServiceGroupInstanceMemberListSamplingEnableCounters1 = append(obj4.ServiceGroupInstanceMemberListSamplingEnableCounters1, obj4_1)
		}

		c.ServiceGroupInstanceMemberListName = append(c.ServiceGroupInstanceMemberListName, obj4)
	}

	vc.ServiceGroupInstanceName = c
	return vc
}
