package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection`: DDoS Detection Commands\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionCreate,
		UpdateContext: resourceDdosDetectionUpdate,
		ReadContext:   resourceDdosDetectionRead,
		DeleteContext: resourceDdosDetectionDelete,

		Schema: map[string]*schema.Schema{
			"agent_group_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_group_name": {
							Type: schema.TypeString, Required: true, Description: "Specify name for the agent-group",
						},
						"agent": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agent_name": {
										Type: schema.TypeString, Optional: true, Description: "detection agent name",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"agent_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"agent_name": {
							Type: schema.TypeString, Required: true, Description: "Specify name for the agent",
						},
						"agent_v4_addr": {
							Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv4 address",
						},
						"agent_v6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Configure agent's IPv6 address",
						},
						"agent_type": {
							Type: schema.TypeString, Optional: true, Description: "'Cisco': Cisco; 'Juniper': Juniper;",
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'sflow-packets-received': sFlow Packets Received; 'sflow-samples-received': sFlow Samples Received; 'sflow-samples-bad-len': sFlow Samples Bad Length; 'sflow-samples-non-std': sFlow Samples Non-standard; 'sflow-samples-skipped': sFlow Samples Skipped; 'sflow-sample-record-bad-len': sFlow Sample Records Bad Length; 'sflow-samples-sent-for-detection': sFlow Samples Processed For Detection; 'sflow-sample-record-invalid-layer2': sFlow Sample Records Unknown Layer-2; 'sflow-sample-ipv6-hdr-parse-fail': sFlow Sample IPv6 Record Header Parse Failures; 'sflow-disabled': sFlow Packet Samples Processing Disabled; 'netflow-disabled': Netflow Flow Samples Processing Disabled; 'netflow-v5-packets-received': Netflow v5 Packets Received; 'netflow-v5-samples-received': Netflow v5 Samples Received; 'netflow-v5-samples-sent-for-detection': Netflow v5 Samples Processed For Detection; 'netflow-v5-sample-records-bad-len': Netflow v5 Sample Records Bad Length; 'netflow-v5-max-records-exceed': Netflow v5 Sample Max Records Error; 'netflow-v9-packets-received': Netflow v9 Packets Received; 'netflow-v9-samples-received': Netflow v9 Samples Received; 'netflow-v9-samples-sent-for-detection': Netflow v9 Samples Processed For Detection; 'netflow-v9-sample-records-bad-len': Netflow v9 Sample Records Bad Length; 'netflow-v9-sample-flowset-bad-padding': Netflow v9 Sample Flowset Bad Padding; 'netflow-v9-max-records-exceed': Netflow v9 Sample Max Records Error; 'netflow-v9-template-not-found': Netflow v9 Template Not Found; 'netflow-v10-packets-received': Netflow v10 Packets Received; 'netflow-v10-samples-received': Netflow v10 Samples Received; 'netflow-v10-samples-sent-for-detection': Netflow v10 Samples Procssed For Detection; 'netflow-v10-sample-records-bad-len': Netflow v10 Sample Records Bad Length; 'netflow-v10-max-records-exceed': Netflow v10 Sample Max records Error; 'netflow-tcp-sample-received': Netflow TCP Samples Received; 'netflow-udp-sample-received': Netflow UDP Samples received; 'netflow-icmp-sample-received': Netflow ICMP Samples Received; 'netflow-other-sample-received': Netflow OTHER Samples Received; 'netflow-record-copy-oom-error': Netflow Data Record Copy Fail, Local MEM size error; 'netflow-record-rse-invalid': Netflow Data Record Reduced Size Invalid; 'netflow-sample-flow-dur-error': Netflow Sample Flow Duration Error; 'flow-dst-entry-miss': DDoS Destination Entry Lookup Failures; 'flow-ip-proto-or-port-miss': DDoS Destination Service Lookup Failures; 'flow-detection-msgq-full': Detection Message Enqueue Failures; 'flow-network-entry-miss': DDoS Destination Network-object Entry Lookup Failures; 'xflow-extend-pkt-rcv': XFlow Sample Extend Packets Received; 'xflow-extend-byte-rcv': XFlow Sample Extend Bytes Received; 'xflow-dst-entry-miss-extend-pkt-rcv': Extend Packets Received of DDoS Destination Entry Miss; 'xflow-dst-entry-miss-extend-byte-rcv': Extend Bytes Received of DDoS Destination Entry Miss; 'xflow-dst-svc-miss-extend-pkt-rcv': Extend Packets Received of DDoS Destination Service Miss; 'xflow-dst-svc-miss-extend-byte-rcv': Extend Bytes Received of DDoS Destination Service Miss;",
									},
								},
							},
						},
						"sflow": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sflow_pkt_samples_collection": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable sflow packet samples collection(default); 'disable': Disable sflow packet samples collection;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"netflow": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"netflow_samples_collection": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable Netflow flow samples collection(default); 'disable': Disable Netflow flow samples collection;",
									},
									"netflow_sampling_rate": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure agent's netflow sampling rate",
									},
									"active_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Configure agent's flow active timeout (seconds)",
									},
									"inactive_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Configure agent's flow inactive timeout (seconds)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"ddos_script": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file": {
							Type: schema.TypeString, Optional: true, Description: "startup-config local file name",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'delete': delete;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DDoS detection (default: enabled)",
			},
			"entry_saving": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"clear_saved_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear all saved network-object-based detection entries and learned indicators",
						},
						"manual_save": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually save network-object-based detection entries and learned indicators",
						},
						"manual_restore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually restore network-object-based detection entries and learned indicators",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"resource_usage": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"settings": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"detector_mode": {
							Type: schema.TypeString, Optional: true, Description: "'standalone': Standalone detector; 'on-box': Mitigator and Detector on the same box; 'auto-svc-discovery': Auto Service discovery using Visibility module (Deprecatd);",
						},
						"dedicated_cpus": {
							Type: schema.TypeInt, Optional: true, Description: "Configure the number of dedicated cores for detection",
						},
						"ctrl_cpu_usage": {
							Type: schema.TypeInt, Optional: true, Description: "Control cpu usage threshold for DDoS detection",
						},
						"full_core_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable full core",
						},
						"top_k_reset_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Configure top-k reset interval",
						},
						"pkt_sampling": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Sample 1 in X packets (default: X=1)",
									},
									"start_level": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure the start level for dynamic sampling adjustment (Sample 1 in N packets, the larger level the larger value of N (default: 1))",
									},
								},
							},
						},
						"histogram_escalate_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "histogram escalate sensitivity for DDoS detection",
						},
						"histogram_de_escalate_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "histogram de-escalate sensitivity for DDoS detection",
						},
						"detection_window_size": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure detection window size in seconds (DDoS detection window size in seconds(default: 1))",
						},
						"initial_learning_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Initial learning interval (in hours) before processing",
						},
						"export_interval": {
							Type: schema.TypeInt, Optional: true, Default: 20, Description: "Configure Baselining and export interval in seconds (DDoS Baselining and export interval in seconds(default: 20))",
						},
						"notification_debug_log": {
							Type: schema.TypeString, Optional: true, Description: "'enable': Enable detection notification debug log (default: disabled);",
						},
						"network_object_window_size": {
							Type: schema.TypeString, Optional: true, Default: "30", Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '30': 30 seconds;  (DDoS detection window size in seconds(default: 30))",
						},
						"network_object_flooding_multiple": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "multiplier for flooding detection threshold in network objects (default 2x threshold)",
						},
						"de_escalation_quiet_time": {
							Type: schema.TypeInt, Optional: true, Description: "Configure de-escalation needed time in minutes from level 1 to 0.(default 1 minutes)",
						},
						"network_object_subnet_notify_percent": {
							Type: schema.TypeInt, Optional: true, Description: "Send subnet notification when anomaly children subnet entries over configured percentage.(default 50%)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"entry_saving": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disable_bootup_restore": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable auto-restoring when system boots up",
									},
									"interval": {
										Type: schema.TypeInt, Optional: true, Description: "Configure periodical auto-saving interval in minutes",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"standalone_settings": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable standalone detector; 'disable': Disable standalone detector (default);",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"sflow": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"listening_port": {
													Type: schema.TypeInt, Optional: true, Default: 6343, Description: "sFlow port to receive packets (sFlow port number(default 6343))",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"netflow": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"listening_port": {
													Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Netflow port to receive packets (Netflow port number(default 9996))",
												},
												"template_active_timeout": {
													Type: schema.TypeInt, Optional: true, Default: 30, Description: "Configure active timeout of the netflow templates received in mins (Template active timeout(mins)(default 30mins))",
												},
												"distribute_by_duration": {
													Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable data distribution by flow duration(default); 'disable': Disable data distribution by flow duration;",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
								},
							},
						},
						"zone_notifications": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_entry": {
										Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable source entry detection notification; 'disable': Disable source entry detection notification(default);",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"statistics": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trustlist": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"v4_class_list": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Class-list name",
						},
						"v6_class_list": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Class-list name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDetectionAgentGroupList(d []interface{}) []edpt.DdosDetectionAgentGroupList {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentGroupList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentGroupList
		oi.AgentGroupName = in["agent_group_name"].(string)
		oi.Agent = getSliceDdosDetectionAgentGroupListAgent(in["agent"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDetectionAgentGroupListAgent(d []interface{}) []edpt.DdosDetectionAgentGroupListAgent {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentGroupListAgent, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentGroupListAgent
		oi.AgentName = in["agent_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDetectionAgentList(d []interface{}) []edpt.DdosDetectionAgentList {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentList
		oi.AgentName = in["agent_name"].(string)
		oi.AgentV4Addr = in["agent_v4_addr"].(string)
		oi.AgentV6Addr = in["agent_v6_addr"].(string)
		oi.AgentType = in["agent_type"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDetectionAgentListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.Sflow = getObjectDdosDetectionAgentListSflow(in["sflow"].([]interface{}))
		oi.Netflow = getObjectDdosDetectionAgentListNetflow(in["netflow"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDetectionAgentListSamplingEnable(d []interface{}) []edpt.DdosDetectionAgentListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionAgentListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionAgentListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDetectionAgentListSflow(d []interface{}) edpt.DdosDetectionAgentListSflow {

	count1 := len(d)
	var ret edpt.DdosDetectionAgentListSflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPktSamplesCollection = in["sflow_pkt_samples_collection"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionAgentListNetflow(d []interface{}) edpt.DdosDetectionAgentListNetflow {

	count1 := len(d)
	var ret edpt.DdosDetectionAgentListNetflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetflowSamplesCollection = in["netflow_samples_collection"].(string)
		ret.NetflowSamplingRate = in["netflow_sampling_rate"].(int)
		ret.ActiveTimeout = in["active_timeout"].(int)
		ret.InactiveTimeout = in["inactive_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionDdosScript154(d []interface{}) edpt.DdosDetectionDdosScript154 {

	count1 := len(d)
	var ret edpt.DdosDetectionDdosScript154
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.File = in["file"].(string)
		ret.Action = in["action"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionEntrySaving155(d []interface{}) edpt.DdosDetectionEntrySaving155 {

	count1 := len(d)
	var ret edpt.DdosDetectionEntrySaving155
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ClearSavedData = in["clear_saved_data"].(int)
		ret.ManualSave = in["manual_save"].(int)
		ret.ManualRestore = in["manual_restore"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionResourceUsage156(d []interface{}) edpt.DdosDetectionResourceUsage156 {

	var ret edpt.DdosDetectionResourceUsage156
	return ret
}

func getObjectDdosDetectionSettings157(d []interface{}) edpt.DdosDetectionSettings157 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettings157
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DetectorMode = in["detector_mode"].(string)
		ret.DedicatedCpus = in["dedicated_cpus"].(int)
		ret.CtrlCpuUsage = in["ctrl_cpu_usage"].(int)
		ret.FullCoreEnable = in["full_core_enable"].(int)
		ret.TopKResetInterval = in["top_k_reset_interval"].(int)
		ret.PktSampling = getSliceDdosDetectionSettingsPktSampling158(in["pkt_sampling"].([]interface{}))
		ret.HistogramEscalatePercentage = in["histogram_escalate_percentage"].(int)
		ret.HistogramDeEscalatePercentage = in["histogram_de_escalate_percentage"].(int)
		ret.DetectionWindowSize = in["detection_window_size"].(int)
		ret.InitialLearningInterval = in["initial_learning_interval"].(int)
		ret.ExportInterval = in["export_interval"].(int)
		ret.NotificationDebugLog = in["notification_debug_log"].(string)
		ret.NetworkObjectWindowSize = in["network_object_window_size"].(string)
		ret.NetworkObjectFloodingMultiple = in["network_object_flooding_multiple"].(int)
		ret.DeEscalationQuietTime = in["de_escalation_quiet_time"].(int)
		ret.NetworkObjectSubnetNotifyPercent = in["network_object_subnet_notify_percent"].(int)
		//omit uuid
		ret.EntrySaving = getObjectDdosDetectionSettingsEntrySaving159(in["entry_saving"].([]interface{}))
		ret.StandaloneSettings = getObjectDdosDetectionSettingsStandaloneSettings160(in["standalone_settings"].([]interface{}))
		ret.ZoneNotifications = getObjectDdosDetectionSettingsZoneNotifications163(in["zone_notifications"].([]interface{}))
	}
	return ret
}

func getSliceDdosDetectionSettingsPktSampling158(d []interface{}) []edpt.DdosDetectionSettingsPktSampling158 {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionSettingsPktSampling158, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionSettingsPktSampling158
		oi.OverrideRate = in["override_rate"].(int)
		oi.StartLevel = in["start_level"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDetectionSettingsEntrySaving159(d []interface{}) edpt.DdosDetectionSettingsEntrySaving159 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsEntrySaving159
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DisableBootupRestore = in["disable_bootup_restore"].(int)
		ret.Interval = in["interval"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettings160(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettings160 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettings160
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
		ret.Sflow = getObjectDdosDetectionSettingsStandaloneSettingsSflow161(in["sflow"].([]interface{}))
		ret.Netflow = getObjectDdosDetectionSettingsStandaloneSettingsNetflow162(in["netflow"].([]interface{}))
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsSflow161(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsSflow161 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsSflow161
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsNetflow162(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsNetflow162 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsNetflow162
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		ret.TemplateActiveTimeout = in["template_active_timeout"].(int)
		ret.DistributeByDuration = in["distribute_by_duration"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsZoneNotifications163(d []interface{}) edpt.DdosDetectionSettingsZoneNotifications163 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsZoneNotifications163
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SourceEntry = in["source_entry"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionStatistics164(d []interface{}) edpt.DdosDetectionStatistics164 {

	var ret edpt.DdosDetectionStatistics164
	return ret
}

func getObjectDdosDetectionTrustlist165(d []interface{}) edpt.DdosDetectionTrustlist165 {

	count1 := len(d)
	var ret edpt.DdosDetectionTrustlist165
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.V4ClassList = in["v4_class_list"].(string)
		ret.V6ClassList = in["v6_class_list"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosDetection(d *schema.ResourceData) edpt.DdosDetection {
	var ret edpt.DdosDetection
	ret.Inst.AgentGroupList = getSliceDdosDetectionAgentGroupList(d.Get("agent_group_list").([]interface{}))
	ret.Inst.AgentList = getSliceDdosDetectionAgentList(d.Get("agent_list").([]interface{}))
	ret.Inst.DdosScript = getObjectDdosDetectionDdosScript154(d.Get("ddos_script").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.EntrySaving = getObjectDdosDetectionEntrySaving155(d.Get("entry_saving").([]interface{}))
	ret.Inst.ResourceUsage = getObjectDdosDetectionResourceUsage156(d.Get("resource_usage").([]interface{}))
	ret.Inst.Settings = getObjectDdosDetectionSettings157(d.Get("settings").([]interface{}))
	ret.Inst.Statistics = getObjectDdosDetectionStatistics164(d.Get("statistics").([]interface{}))
	ret.Inst.Trustlist = getObjectDdosDetectionTrustlist165(d.Get("trustlist").([]interface{}))
	//omit uuid
	return ret
}
