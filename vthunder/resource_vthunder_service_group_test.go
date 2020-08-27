package vthunder

import (
	"fmt"
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"reflect"
	"testing"
)

var NAME_SG = "sg9"
var SR_NAME = "rs9"

var TEST_SG_RESOURCE = `
resource "vthunder_service_group" "sg9" {
protocol="TCP"
name="` + NAME_SG + `"
member_list {
name="` + SR_NAME + `"
port=80
} 
}
`

//Acceptance test
func TestAccVthunderServiceGroup_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckServiceGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_SG_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckServiceGroupExists(NAME_SG, true),
					resource.TestCheckResourceAttr("vthunder_service_group.sg9", "name", NAME_SG),
					resource.TestCheckResourceAttr("vthunder_service_group.sg9", "protocol", "TCP"),
					resource.TestCheckResourceAttr("vthunder_service_group.sg9", "member_list.0.name", SR_NAME),
					resource.TestCheckResourceAttr("vthunder_service_group.sg9", "member_list.0.port", "80"),
				),
			},
		},
	})
}

func TestAccVthunderServiceGroup_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckServiceGroupDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_SG_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckServiceGroupExists(NAME_SG, true),
				),
				ResourceName:      NAME_SG,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}

//Unit test for utility method to create Service group structure
func TestDataToSg(t *testing.T) {

	ConnRate := 1
	ResetOnServerSelectionFail := 1
	HealthCheckDisable := 1
	Protocol := "HTTP"
	TrafficReplicationMirrorIPRepl := 1
	ResetPriorityAffinity := 1
	MinActiveMember := 1
	StatsDataAction := "enable"
	TrafficReplicationMirrorDaRepl := 1
	TemplatePolicyShared := "T"
	RptExtServer := 1
	TemplatePort := "80"
	ConnRateGracePeriod := 1
	L4SessionUsageDuration := 1
	UUID := "ID"
	BackupServerEventLog := 1
	LcMethod := "Method"
	PseudoRoundRobin := 1
	SharedPartitionPolicyTemplate := 1
	L4SessionUsageRevertRate := 1
	SharedPartitionSvcgrpHealthCheck := 1
	TemplateServer := "server"
	SvcgrpHealthCheckShared := "healthcheckshared"
	TrafficReplicationMirror := 1
	L4SessionRevertDuration := 1
	TrafficReplicationMirrorSaDaRepl := 1
	LbMethod := "method"
	StatelessAutoSwitch := 1
	MinActiveMemberAction := "enable"
	L4SessionUsage := 1
	ExtendedStats := 1
	ConnRateRevertDuration := 1
	StrictSelect := 1
	Name := "name"
	TrafficReplicationMirrorSaRepl := 1
	ReportDelay := 1
	ConnRateLog := 1
	L4SessionUsageLog := 1
	ConnRateDuration := 1
	StatelessLbMethod := "method"
	TemplatePolicy := "policy"
	StatelessLbMethod2 := "method2"
	UserTag := "tag"
	SampleRspTime := 1
	TopFastest := 1
	ConnRevertRate := 1
	L4SessionUsageGracePeriod := 1
	PriorityAffinity := 1
	TopSlowest := 1
	HealthCheck := "check"

	resourceSchema := map[string]*schema.Schema{
		"conn_rate": {
			Type: schema.TypeInt,
		},
		"reset_on_server_selection_fail": {
			Type: schema.TypeInt,
		},
		"health_check_disable": {
			Type: schema.TypeInt,
		},
		"protocol": {
			Type: schema.TypeString,
		},
		"traffic_replication_mirror_ip_repl": {
			Type: schema.TypeInt,
		},
		"reset_priority_affinity": {
			Type: schema.TypeInt,
		},
		"min_active_member": {
			Type: schema.TypeInt,
		},
		"stats_data_action": {
			Type: schema.TypeString,
		},
		"traffic_replication_mirror_da_repl": {
			Type: schema.TypeInt,
		},
		"template_policy_shared": {
			Type: schema.TypeString,
		},
		"rpt_ext_server": {
			Type: schema.TypeInt,
		},
		"template_port": {
			Type: schema.TypeString,
		},
		"conn_rate_grace_period": {
			Type: schema.TypeInt,
		},
		"L4SessionUsageDuration": {
			Type: schema.TypeInt,
		},
		"uuid": {
			Type: schema.TypeString,
		},
		"backup_server_event_log": {
			Type: schema.TypeInt,
		},
		"lc_method": {
			Type: schema.TypeString,
		},
		"pseudo_round_robin": {
			Type: schema.TypeInt,
		},
		"shared_partition_policy_template": {
			Type: schema.TypeInt,
		},
		"l4_session_usage_revert_rate": {
			Type: schema.TypeInt,
		},
		"shared_partition_svcgrp_health_check": {
			Type: schema.TypeInt,
		},
		"svcgrp_health_check_shared": {
			Type: schema.TypeString,
		},
		"template_server": {
			Type: schema.TypeString,
		},
		"traffic_replication_mirror": {
			Type: schema.TypeInt,
		},
		"l4_session_revert_duration": {
			Type: schema.TypeInt,
		},
		"traffic_replication_mirror_sa_da_repl": {
			Type: schema.TypeInt,
		},
		"lb_method": {
			Type: schema.TypeString,
		},
		"stateless_auto_switch": {
			Type: schema.TypeInt,
		},
		"min_active_member_action": {
			Type: schema.TypeString,
		},
		"l4_session_usage": {
			Type: schema.TypeInt,
		},
		"extended_stats": {
			Type: schema.TypeInt,
		},
		"conn_rate_revert_duration": {
			Type: schema.TypeInt,
		},
		"strict_select": {
			Type: schema.TypeInt,
		},
		"name": {
			Type: schema.TypeString,
		},
		"traffic_replication_mirror_sa_repl": {
			Type: schema.TypeInt,
		},
		"report_delay": {
			Type: schema.TypeInt,
		},
		"conn_rate_log": {
			Type: schema.TypeInt,
		},
		"l4_session_usage_log": {
			Type: schema.TypeInt,
		},
		"conn_rate_duration": {
			Type: schema.TypeInt,
		},
		"stateless_lb_method": {
			Type: schema.TypeString,
		},
		"template_policy": {
			Type: schema.TypeString,
		},
		"stateless_lb_method2": {
			Type: schema.TypeString,
		},
		"user_tag": {
			Type: schema.TypeString,
		},
		"sample_rsp_time": {
			Type: schema.TypeInt,
		},
		"top_fastest": {
			Type: schema.TypeInt,
		},
		"conn_revert_rate": {
			Type: schema.TypeInt,
		},
		"l4_session_usage_grace_period": {
			Type: schema.TypeInt,
		},
		"priority_affinity": {
			Type: schema.TypeInt,
		},
		"top_slowest": {
			Type: schema.TypeInt,
		},
		"health_check": {
			Type: schema.TypeString,
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
	}

	resourceDataMap := map[string]interface{}{
		"conn_rate":                             ConnRate,
		"reset_on_server_selection_fail":        ResetOnServerSelectionFail,
		"health_check_disable":                  HealthCheckDisable,
		"protocol":                              Protocol,
		"traffic_replication_mirror_ip_repl":    TrafficReplicationMirrorIPRepl,
		"reset_priority_affinity":               ResetPriorityAffinity,
		"min_active_member":                     MinActiveMember,
		"stats_data_action":                     StatsDataAction,
		"traffic_replication_mirror_da_repl":    TrafficReplicationMirrorDaRepl,
		"template_policy_shared":                TemplatePolicyShared,
		"rpt_ext_server":                        RptExtServer,
		"template_port":                         TemplatePort,
		"conn_rate_grace_period":                ConnRateGracePeriod,
		"L4SessionUsageDuration":                L4SessionUsageDuration,
		"uuid":                                  UUID,
		"backup_server_event_log":               BackupServerEventLog,
		"lc_method":                             LcMethod,
		"pseudo_round_robin":                    PseudoRoundRobin,
		"shared_partition_policy_template":      SharedPartitionPolicyTemplate,
		"l4_session_usage_revert_rate":          L4SessionUsageRevertRate,
		"shared_partition_svcgrp_health_check":  SharedPartitionSvcgrpHealthCheck,
		"svcgrp_health_check_shared":            SvcgrpHealthCheckShared,
		"template_server":                       TemplateServer,
		"traffic_replication_mirror":            TrafficReplicationMirror,
		"l4_session_revert_duration":            L4SessionRevertDuration,
		"traffic_replication_mirror_sa_da_repl": TrafficReplicationMirrorSaDaRepl,
		"lb_method":                             LbMethod,
		"stateless_auto_switch":                 StatelessAutoSwitch,
		"min_active_member_action":              MinActiveMemberAction,
		"l4_session_usage":                      L4SessionUsage,
		"extended_stats":                        ExtendedStats,
		"conn_rate_revert_duration":             ConnRateRevertDuration,
		"strict_select":                         StrictSelect,
		"name":                                  Name,
		"traffic_replication_mirror_sa_repl":    TrafficReplicationMirrorSaRepl,
		"report_delay":                          ReportDelay,
		"conn_rate_log":                         ConnRateLog,
		"l4_session_usage_log":                  L4SessionUsageLog,
		"conn_rate_duration":                    ConnRateDuration,
		"stateless_lb_method":                   StatelessLbMethod,
		"template_policy":                       TemplatePolicy,
		"stateless_lb_method2":                  StatelessLbMethod2,
		"user_tag":                              UserTag,
		"sample_rsp_time":                       SampleRspTime,
		"top_fastest":                           TopFastest,
		"conn_revert_rate":                      ConnRevertRate,
		"l4_session_usage_grace_period":         L4SessionUsageGracePeriod,
		"priority_affinity":                     PriorityAffinity,
		"top_slowest":                           TopSlowest,
		"health_check":                          HealthCheck,
		"sampling_enable":                       map[string]interface{}{},
		"priorities":                            map[string]interface{}{},
		"member_list":                           map[string]interface{}{},
	}

	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, resourceDataMap)

	var s go_vthunder.ServiceGroup

	var sInstance go_vthunder.ServiceGroupInstance

	sInstance.ConnRate = ConnRate
	sInstance.ResetOnServerSelectionFail = ResetOnServerSelectionFail
	sInstance.HealthCheckDisable = HealthCheckDisable
	sInstance.Protocol = Protocol
	sInstance.TrafficReplicationMirrorIPRepl = TrafficReplicationMirrorIPRepl
	sInstance.ResetPriorityAffinity = ResetPriorityAffinity
	sInstance.MinActiveMember = MinActiveMember
	sInstance.StatsDataAction = StatsDataAction
	sInstance.TrafficReplicationMirrorDaRepl = TrafficReplicationMirrorDaRepl
	sInstance.TemplatePolicyShared = TemplatePolicyShared
	sInstance.RptExtServer = RptExtServer
	sInstance.TemplatePort = TemplatePort
	sInstance.ConnRateGracePeriod = ConnRateGracePeriod
	sInstance.L4SessionUsageDuration = L4SessionUsageDuration
	sInstance.UUID = UUID
	sInstance.BackupServerEventLog = BackupServerEventLog
	sInstance.LcMethod = LcMethod
	sInstance.PseudoRoundRobin = PseudoRoundRobin
	sInstance.SharedPartitionPolicyTemplate = SharedPartitionPolicyTemplate
	sInstance.L4SessionUsageRevertRate = L4SessionUsageRevertRate
	sInstance.SharedPartitionSvcgrpHealthCheck = SharedPartitionSvcgrpHealthCheck
	sInstance.TemplateServer = TemplateServer
	sInstance.SvcgrpHealthCheckShared = SvcgrpHealthCheckShared
	sInstance.TrafficReplicationMirror = TrafficReplicationMirror
	sInstance.L4SessionRevertDuration = L4SessionRevertDuration
	sInstance.TrafficReplicationMirrorSaDaRepl = TrafficReplicationMirrorSaDaRepl
	sInstance.LbMethod = LbMethod
	sInstance.StatelessAutoSwitch = StatelessAutoSwitch
	sInstance.MinActiveMemberAction = MinActiveMemberAction
	sInstance.L4SessionUsage = L4SessionUsage
	sInstance.ExtendedStats = ExtendedStats
	sInstance.ConnRateRevertDuration = ConnRateRevertDuration
	sInstance.StrictSelect = StrictSelect
	sInstance.Name = Name
	sInstance.TrafficReplicationMirrorSaRepl = TrafficReplicationMirrorSaRepl
	sInstance.ReportDelay = ReportDelay
	sInstance.ConnRateLog = ConnRateLog
	sInstance.L4SessionUsageLog = L4SessionUsageLog
	sInstance.ConnRateDuration = ConnRateDuration
	sInstance.StatelessLbMethod = StatelessLbMethod
	sInstance.TemplatePolicy = TemplatePolicy
	sInstance.StatelessLbMethod2 = StatelessLbMethod2
	sInstance.UserTag = UserTag
	sInstance.SampleRspTime = SampleRspTime
	sInstance.TopFastest = TopFastest
	sInstance.ConnRevertRate = ConnRevertRate
	sInstance.L4SessionUsageGracePeriod = L4SessionUsageGracePeriod
	sInstance.PriorityAffinity = PriorityAffinity
	sInstance.TopSlowest = TopSlowest
	sInstance.HealthCheck = HealthCheck
	sInstance.Priority = make([]go_vthunder.Priorities, 0, 1)
	sInstance.Counters1 = make([]go_vthunder.SamplingEnable, 0, 1)
	sInstance.Host = make([]go_vthunder.MemberList, 0, 1)

	s.Name = sInstance

	cases := []struct {
		input  *schema.ResourceData
		output go_vthunder.ServiceGroup
	}{
		{
			input:  resourceLocalData,
			output: s,
		},
	}

	for _, tc := range cases {
		output := dataToSg("name", resourceLocalData)
		if !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("Got:\n\n%#v\n\nExpected:\n\n%#v", output, tc.output)
		}
	}

}

func testCheckServiceGroupExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		client := testAccProvider.Meta().(vThunder)
		sg, err := go_vthunder.GetSG(client.Token, name, client.Host)

		if err != nil {
			return err
		}
		if exists && sg == nil {
			return fmt.Errorf(" service group %s was not created.", name)
		}
		if !exists && sg != nil {
			return fmt.Errorf(" service group %s still exists.", name)
		}
		return nil
	}
}

func testCheckServiceGroupDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(vThunder)
	for _, rs := range s.RootModule().Resources {

		name := rs.Primary.ID
		sg, err := go_vthunder.GetSG(client.Token, name, client.Host)
		if err != nil {
			return err
		}
		if sg == nil {
			return fmt.Errorf(" service group %s not destroyed.", name)
		}
	}
	return nil
}
