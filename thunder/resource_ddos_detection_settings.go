package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_settings`: Configure ddos detection settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionSettingsCreate,
		UpdateContext: resourceDdosDetectionSettingsUpdate,
		ReadContext:   resourceDdosDetectionSettingsRead,
		DeleteContext: resourceDdosDetectionSettingsDelete,

		Schema: map[string]*schema.Schema{
			"ctrl_cpu_usage": {
				Type: schema.TypeInt, Optional: true, Description: "Control cpu usage threshold for DDoS detection",
			},
			"de_escalation_quiet_time": {
				Type: schema.TypeInt, Optional: true, Description: "Configure de-escalation needed time in minutes from level 1 to 0.(default 1 minutes)",
			},
			"dedicated_cpus": {
				Type: schema.TypeInt, Optional: true, Description: "Configure the number of dedicated cores for detection",
			},
			"detection_window_size": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Configure detection window size in seconds (DDoS detection window size in seconds(default: 1))",
			},
			"detector_mode": {
				Type: schema.TypeString, Optional: true, Description: "'standalone': Standalone detector; 'on-box': Mitigator and Detector on the same box; 'auto-svc-discovery': Auto Service discovery using Visibility module (Deprecatd);",
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
			"export_interval": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Configure Baselining and export interval in seconds (DDoS Baselining and export interval in seconds(default: 20))",
			},
			"full_core_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable full core",
			},
			"histogram_de_escalate_percentage": {
				Type: schema.TypeInt, Optional: true, Description: "histogram de-escalate sensitivity for DDoS detection",
			},
			"histogram_escalate_percentage": {
				Type: schema.TypeInt, Optional: true, Description: "histogram escalate sensitivity for DDoS detection",
			},
			"initial_learning_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Initial learning interval (in hours) before processing",
			},
			"network_object_flooding_multiple": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "multiplier for flooding detection threshold in network objects (default 2x threshold)",
			},
			"network_object_subnet_notify_percent": {
				Type: schema.TypeInt, Optional: true, Description: "Send subnet notification when anomaly children subnet entries over configured percentage.(default 50%)",
			},
			"network_object_window_size": {
				Type: schema.TypeString, Optional: true, Default: "30", Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '30': 30 seconds;  (DDoS detection window size in seconds(default: 30))",
			},
			"notification_debug_log": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable detection notification debug log (default: disabled);",
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
			"top_k_reset_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Configure top-k reset interval",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
	}
}
func resourceDdosDetectionSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettings(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettings(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionSettingsRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettings(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionSettingsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionSettings(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDetectionSettingsEntrySaving149(d []interface{}) edpt.DdosDetectionSettingsEntrySaving149 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsEntrySaving149
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DisableBootupRestore = in["disable_bootup_restore"].(int)
		ret.Interval = in["interval"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosDetectionSettingsPktSampling(d []interface{}) []edpt.DdosDetectionSettingsPktSampling {

	count1 := len(d)
	ret := make([]edpt.DdosDetectionSettingsPktSampling, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDetectionSettingsPktSampling
		oi.OverrideRate = in["override_rate"].(int)
		oi.StartLevel = in["start_level"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettings150(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettings150 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettings150
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		//omit uuid
		ret.Sflow = getObjectDdosDetectionSettingsStandaloneSettingsSflow151(in["sflow"].([]interface{}))
		ret.Netflow = getObjectDdosDetectionSettingsStandaloneSettingsNetflow152(in["netflow"].([]interface{}))
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsSflow151(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsSflow151 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsSflow151
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsNetflow152(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsNetflow152 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsNetflow152
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		ret.TemplateActiveTimeout = in["template_active_timeout"].(int)
		ret.DistributeByDuration = in["distribute_by_duration"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsZoneNotifications153(d []interface{}) edpt.DdosDetectionSettingsZoneNotifications153 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsZoneNotifications153
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SourceEntry = in["source_entry"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosDetectionSettings(d *schema.ResourceData) edpt.DdosDetectionSettings {
	var ret edpt.DdosDetectionSettings
	ret.Inst.CtrlCpuUsage = d.Get("ctrl_cpu_usage").(int)
	ret.Inst.DeEscalationQuietTime = d.Get("de_escalation_quiet_time").(int)
	ret.Inst.DedicatedCpus = d.Get("dedicated_cpus").(int)
	ret.Inst.DetectionWindowSize = d.Get("detection_window_size").(int)
	ret.Inst.DetectorMode = d.Get("detector_mode").(string)
	ret.Inst.EntrySaving = getObjectDdosDetectionSettingsEntrySaving149(d.Get("entry_saving").([]interface{}))
	ret.Inst.ExportInterval = d.Get("export_interval").(int)
	ret.Inst.FullCoreEnable = d.Get("full_core_enable").(int)
	ret.Inst.HistogramDeEscalatePercentage = d.Get("histogram_de_escalate_percentage").(int)
	ret.Inst.HistogramEscalatePercentage = d.Get("histogram_escalate_percentage").(int)
	ret.Inst.InitialLearningInterval = d.Get("initial_learning_interval").(int)
	ret.Inst.NetworkObjectFloodingMultiple = d.Get("network_object_flooding_multiple").(int)
	ret.Inst.NetworkObjectSubnetNotifyPercent = d.Get("network_object_subnet_notify_percent").(int)
	ret.Inst.NetworkObjectWindowSize = d.Get("network_object_window_size").(string)
	ret.Inst.NotificationDebugLog = d.Get("notification_debug_log").(string)
	ret.Inst.PktSampling = getSliceDdosDetectionSettingsPktSampling(d.Get("pkt_sampling").([]interface{}))
	ret.Inst.StandaloneSettings = getObjectDdosDetectionSettingsStandaloneSettings150(d.Get("standalone_settings").([]interface{}))
	ret.Inst.TopKResetInterval = d.Get("top_k_reset_interval").(int)
	//omit uuid
	ret.Inst.ZoneNotifications = getObjectDdosDetectionSettingsZoneNotifications153(d.Get("zone_notifications").([]interface{}))
	return ret
}
