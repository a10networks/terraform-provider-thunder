package vthunder

//vThunder resource - Service Group

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/go_vthunder/vthunder"
	"log"
	"util"
)

func resourceServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceGroupCreate,
		Update: resourceServiceGroupUpdate,
		Read:   resourceServiceGroupRead,
		Delete: resourceServiceGroupDelete,

		Schema: map[string]*schema.Schema{
			"extended_stats": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_lb_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"svcgrp_health_check_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pseudo_round_robin": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"priorities": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"conn_rate_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"health_check_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"report_delay": {
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
			"template_policy_shared": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"min_active_member": {
				Type:        schema.TypeInt,
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
			"l4_session_usage": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"min_active_member_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_revert_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_auto_switch": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_revert_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			//			"reset": {
			//				Type:     schema.TypeList,
			//				Optional: true,
			//				MaxItems: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"auto_switch": {
			//							Type:        schema.TypeInt,
			//							Optional:    true,
			//							Description: "",
			//						},
			//					},
			//				},
			//			},
			"traffic_replication_mirror_sa_da_repl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"health_check": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"stateless_lb_method2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_grace_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"lc_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"lb_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"strict_select": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"member_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"member_template": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"member_stats_data_disable": {
							Type:        schema.TypeInt,
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"resolve_as": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"server_ipv6_addr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"fqdn_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"member_state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"stats_data_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"backup_server_event_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sample_rsp_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_revert_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"l4_session_usage_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_policy_template": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"top_slowest": {
				Type:        schema.TypeInt,
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
			"reset_priority_affinity": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"traffic_replication_mirror_sa_repl": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"shared_partition_svcgrp_health_check": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_revert_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"template_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"conn_rate_grace_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"top_fastest": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"rpt_ext_server": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"conn_rate_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}

}

func resourceServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Creating service group   (Inside resourceServiceGroupCreate    " + name)
		v := dataToSg(name, d)
		logger.Println("[INFO] received V from method data to sg --" + v.Name.Name + ",--" + v.Name.UUID)
		d.SetId(name)
		go_vthunder.PostSG(client.Token, v, client.Host)

		return resourceServiceGroupRead(d, meta)
	}
	return nil
}

func resourceServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading service group (Inside resourceServiceGroupRead)")

	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()

		logger.Println("[INFO] Fetching service group Read" + name)
		
		sg,err := go_vthunder.GetSG(client.Token, name, client.Host)
		
		if(sg == nil){
			logger.Println("[INFO] No service group found "+ name)
			d.SetId("")
			return nil
		}

		return err
	}
	return nil
}

func resourceServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Get("name").(string)
		logger.Println("[INFO] Modifying service group   (Inside resourceServiceGroupUpdate    " + name)
		v := dataToSg(name, d)
		logger.Println("[INFO] received V from method data to sg --" + v.Name.Name)
		d.SetId(name)
		go_vthunder.PutSG(client.Token, name, v, client.Host)

		return resourceServiceGroupRead(d, meta)
	}
	return nil
}

func resourceServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(vThunder)
	logger := util.GetLoggerInstance()

	if client.Host != "" {

		name := d.Id()
		logger.Println("[INFO] Deleting service group (Inside resourceServiceGroupDelete) " + name)

		err := go_vthunder.DeleteSG(client.Token, name, client.Host)
		if err != nil {
			log.Printf("[ERROR] Unable to Delete service group  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

//Utility method to instantiate ServiceGroup structure
func dataToSg(name string, d *schema.ResourceData) go_vthunder.ServiceGroup {
	//	logger := util.GetLoggerInstance()
	var s go_vthunder.ServiceGroup

	var sInstance go_vthunder.ServiceGroupInstance

	sInstance.ConnRate = d.Get("conn_rate").(int)
	sInstance.ResetOnServerSelectionFail = d.Get("reset_on_server_selection_fail").(int)
	sInstance.HealthCheckDisable = d.Get("health_check_disable").(int)
	sInstance.Protocol = d.Get("protocol").(string)
	sInstance.TrafficReplicationMirrorIPRepl = d.Get("traffic_replication_mirror_ip_repl").(int)
	sInstance.ResetPriorityAffinity = d.Get("reset_priority_affinity").(int)
	sInstance.MinActiveMember = d.Get("min_active_member").(int)
	sInstance.StatsDataAction = d.Get("stats_data_action").(string)
	sInstance.TrafficReplicationMirrorDaRepl = d.Get("traffic_replication_mirror_da_repl").(int)
	sInstance.TemplatePolicyShared = d.Get("template_policy_shared").(string)
	sInstance.RptExtServer = d.Get("rpt_ext_server").(int)
	sInstance.TemplatePort = d.Get("template_port").(string)
	sInstance.ConnRateGracePeriod = d.Get("conn_rate_grace_period").(int)
	sInstance.L4SessionUsageDuration = d.Get("l4_session_usage").(int)
	sInstance.UUID = d.Get("uuid").(string)
	sInstance.BackupServerEventLog = d.Get("backup_server_event_log").(int)
	sInstance.LcMethod = d.Get("lc_method").(string)
	sInstance.PseudoRoundRobin = d.Get("pseudo_round_robin").(int)
	sInstance.SharedPartitionPolicyTemplate = d.Get("shared_partition_policy_template").(int)
	sInstance.L4SessionUsageRevertRate = d.Get("l4_session_usage_revert_rate").(int)
	sInstance.SharedPartitionSvcgrpHealthCheck = d.Get("shared_partition_svcgrp_health_check").(int)
	sInstance.TemplateServer = d.Get("template_server").(string)
	sInstance.SvcgrpHealthCheckShared = d.Get("svcgrp_health_check_shared").(string)
	sInstance.TrafficReplicationMirror = d.Get("traffic_replication_mirror").(int)
	sInstance.L4SessionRevertDuration = d.Get("l4_session_revert_duration").(int)
	sInstance.TrafficReplicationMirrorSaDaRepl = d.Get("traffic_replication_mirror_sa_da_repl").(int)
	sInstance.LbMethod = d.Get("lb_method").(string)
	sInstance.StatelessAutoSwitch = d.Get("stateless_auto_switch").(int)
	sInstance.MinActiveMemberAction = d.Get("min_active_member_action").(string)
	sInstance.L4SessionUsage = d.Get("l4_session_usage").(int)
	sInstance.ExtendedStats = d.Get("extended_stats").(int)
	sInstance.ConnRateRevertDuration = d.Get("conn_rate_revert_duration").(int)
	sInstance.StrictSelect = d.Get("strict_select").(int)
	sInstance.Name = d.Get("name").(string)
	sInstance.TrafficReplicationMirrorSaRepl = d.Get("traffic_replication_mirror_sa_repl").(int)
	sInstance.ReportDelay = d.Get("report_delay").(int)
	sInstance.ConnRateLog = d.Get("conn_rate_log").(int)
	sInstance.L4SessionUsageLog = d.Get("l4_session_usage_log").(int)
	sInstance.ConnRateDuration = d.Get("conn_rate_duration").(int)
	sInstance.StatelessLbMethod = d.Get("stateless_lb_method").(string)
	sInstance.TemplatePolicy = d.Get("template_policy").(string)
	sInstance.StatelessLbMethod2 = d.Get("stateless_lb_method2").(string)
	sInstance.UserTag = d.Get("user_tag").(string)
	sInstance.SampleRspTime = d.Get("sample_rsp_time").(int)
	sInstance.TopFastest = d.Get("top_fastest").(int)
	sInstance.ConnRevertRate = d.Get("conn_revert_rate").(int)
	sInstance.L4SessionUsageGracePeriod = d.Get("l4_session_usage_grace_period").(int)
	sInstance.PriorityAffinity = d.Get("priority_affinity").(int)
	sInstance.TopSlowest = d.Get("top_slowest").(int)
	sInstance.HealthCheck = d.Get("health_check").(string)

	priorityCount := d.Get("priorities.#").(int)
	sInstance.Priority = make([]go_vthunder.Priorities, 0, priorityCount)
	for i := 0; i < priorityCount; i++ {
		var pr go_vthunder.Priorities
		prefix := fmt.Sprintf("priorities.%d", i)
		pr.Priority = d.Get(prefix + ".priority").(int)
		pr.PriorityAction = d.Get(prefix + ".priority_action").(string)

		sInstance.Priority = append(sInstance.Priority, pr)
	}

	samplingCount := d.Get("sampling_enable.#").(int)
	sInstance.Counters1 = make([]go_vthunder.SamplingEnable, 0, samplingCount)
	for i := 0; i < samplingCount; i++ {
		var sm go_vthunder.SamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d", i)
		sm.Counters1 = d.Get(prefix + ".counters1").(string)

		sInstance.Counters1 = append(sInstance.Counters1, sm)
	}

	//NEED TO FIGURE OUT IF VALUE IS PROVIDED IN TF FILE OR DEFAULT IS BEING USED
	//	var as Reset
	//	as.AutoSwitch = d.Get("reset.0.auto_switch").(int)
	//	logger.Println("[INFO] Auto switch is- ", d.Get("reset.0.auto_switch").(int))
	//	sInstance.AutoSwitch = as

	memberCount := d.Get("member_list.#").(int)
	sInstance.Host = make([]go_vthunder.MemberList, 0, memberCount)
	for i := 0; i < memberCount; i++ {
		var ml go_vthunder.MemberList
		prefix := fmt.Sprintf("member_list.%d", i)
		ml.FqdnName = d.Get(prefix + ".fqdn_name").(string)
		ml.Host = d.Get(prefix + ".host").(string)
		ml.MemberPriority = d.Get(prefix + ".member_priority").(int)
		ml.MemberState = d.Get(prefix + ".member_state").(string)
		ml.MemberStatsDataDisable = d.Get(prefix + ".member_stats_data_disable").(int)
		ml.MemberTemplate = d.Get(prefix + ".member_template").(string)
		ml.Name = d.Get(prefix + ".name").(string)
		ml.Port = d.Get(prefix + ".port").(int)
		ml.ResolveAs = d.Get(prefix + ".resolve_as").(string)
		ml.ServerIpv6Addr = d.Get(prefix + ".server_ipv6_addr").(string)
		ml.UUID = d.Get(prefix + ".uuid").(string)
		ml.UserTag = d.Get(prefix + ".user_tag").(string)

		sampleCount := d.Get(prefix + ".sampling_enable.#").(int)
		ml.Counters1 = make([]go_vthunder.SamplingEnable, sampleCount, sampleCount)

		for x := 0; x < sampleCount; x++ {
			var s go_vthunder.SamplingEnable
			mapEntity(d.Get(fmt.Sprintf("%s.sampling_enable.%d", prefix, x)).(map[string]interface{}), &s)
			ml.Counters1[x] = s
		}

		sInstance.Host = append(sInstance.Host, ml)
	}

	s.Name = sInstance

	return s
}
