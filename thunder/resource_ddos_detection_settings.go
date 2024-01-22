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
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure periodical auto-saving interval in minutes(default: 0) and 0 to disable.",
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
			"network_object_window_size": {
				Type: schema.TypeString, Optional: true, Default: "30", Description: "'5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '30': 30 seconds;  (DDoS detection window size in seconds(default: 30))",
			},
			"notification_debug_log": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable detection notification debug log (default: disabled);",
			},
			"pkt_sampling": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"override_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Sample 1 in X packets (default: X=1)",
						},
						"assign_index": {
							Type: schema.TypeInt, Optional: true, Description: "Lower index is more aggressive sampling",
						},
						"assign_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Assign rate to given index",
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
						"de_escalation_quiet_time": {
							Type: schema.TypeInt, Optional: true, Description: "Configure de-escalation needed time in minutes from level 1 to 0.(legacy)",
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

func getObjectDdosDetectionSettingsEntrySaving132(d []interface{}) edpt.DdosDetectionSettingsEntrySaving132 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsEntrySaving132
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interval = in["interval"].(int)
		ret.ManualSave = in["manual_save"].(int)
		ret.ManualRestore = in["manual_restore"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsPktSampling(d []interface{}) edpt.DdosDetectionSettingsPktSampling {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsPktSampling
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OverrideRate = in["override_rate"].(int)
		ret.AssignIndex = in["assign_index"].(int)
		ret.AssignRate = in["assign_rate"].(int)
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettings133(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettings133 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettings133
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.DeEscalationQuietTime = in["de_escalation_quiet_time"].(int)
		//omit uuid
		ret.Sflow = getObjectDdosDetectionSettingsStandaloneSettingsSflow134(in["sflow"].([]interface{}))
		ret.Netflow = getObjectDdosDetectionSettingsStandaloneSettingsNetflow135(in["netflow"].([]interface{}))
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsSflow134(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsSflow134 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsSflow134
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDetectionSettingsStandaloneSettingsNetflow135(d []interface{}) edpt.DdosDetectionSettingsStandaloneSettingsNetflow135 {

	count1 := len(d)
	var ret edpt.DdosDetectionSettingsStandaloneSettingsNetflow135
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		ret.TemplateActiveTimeout = in["template_active_timeout"].(int)
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
	ret.Inst.EntrySaving = getObjectDdosDetectionSettingsEntrySaving132(d.Get("entry_saving").([]interface{}))
	ret.Inst.ExportInterval = d.Get("export_interval").(int)
	ret.Inst.FullCoreEnable = d.Get("full_core_enable").(int)
	ret.Inst.HistogramDeEscalatePercentage = d.Get("histogram_de_escalate_percentage").(int)
	ret.Inst.HistogramEscalatePercentage = d.Get("histogram_escalate_percentage").(int)
	ret.Inst.InitialLearningInterval = d.Get("initial_learning_interval").(int)
	ret.Inst.NetworkObjectFloodingMultiple = d.Get("network_object_flooding_multiple").(int)
	ret.Inst.NetworkObjectWindowSize = d.Get("network_object_window_size").(string)
	ret.Inst.NotificationDebugLog = d.Get("notification_debug_log").(string)
	ret.Inst.PktSampling = getObjectDdosDetectionSettingsPktSampling(d.Get("pkt_sampling").([]interface{}))
	ret.Inst.StandaloneSettings = getObjectDdosDetectionSettingsStandaloneSettings133(d.Get("standalone_settings").([]interface{}))
	ret.Inst.TopKResetInterval = d.Get("top_k_reset_interval").(int)
	//omit uuid
	return ret
}
