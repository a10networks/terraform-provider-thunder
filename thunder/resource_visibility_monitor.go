package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_monitor`: Configure monitoring keys\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityMonitorCreate,
		UpdateContext: resourceVisibilityMonitorUpdate,
		ReadContext:   resourceVisibilityMonitorRead,
		DeleteContext: resourceVisibilityMonitorDelete,

		Schema: map[string]*schema.Schema{
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
										Type: schema.TypeString, Optional: true, Description: "'all': all; 'sflow-packets-received': sFlow Packets Received; 'sflow-samples-received': sFlow Samples Received; 'sflow-samples-bad-len': sFlow Samples Bad Length; 'sflow-samples-non-std': sFlow Samples Non-standard; 'sflow-samples-skipped': sFlow Samples Skipped; 'sflow-sample-record-bad-len': sFlow Sample Records Bad Length; 'sflow-samples-sent-for-detection': sFlow Samples Processed For Detection; 'sflow-sample-record-invalid-layer2': sFlow Sample Records Unknown Layer-2; 'sflow-sample-ipv6-hdr-parse-fail': sFlow Sample IPv6 Record Header Parse Failures; 'sflow-disabled': sFlow Packet Samples Processing Disabled; 'netflow-disabled': Netflow Flow Samples Processing Disabled; 'netflow-v5-packets-received': Netflow v5 Packets Received; 'netflow-v5-samples-received': Netflow v5 Samples Received; 'netflow-v5-samples-sent-for-detection': Netflow v5 Samples Processed For Detection; 'netflow-v5-sample-records-bad-len': Netflow v5 Sample Records Bad Length; 'netflow-v5-max-records-exceed': Netflow v5 Sample Max Records Error; 'netflow-v9-packets-received': Netflow v9 Packets Received; 'netflow-v9-samples-received': Netflow v9 Samples Received; 'netflow-v9-samples-sent-for-detection': Netflow v9 Samples Processed For Detection; 'netflow-v9-sample-records-bad-len': Netflow v9 Sample Records Bad Length; 'netflow-v9-max-records-exceed': Netflow v9 Sample Max Records Error; 'netflow-v10-packets-received': Netflow v10 Packets Received; 'netflow-v10-samples-received': Netflow v10 Samples Received; 'netflow-v10-samples-sent-for-detection': Netflow v10 Samples Procssed For Detection; 'netflow-v10-sample-records-bad-len': Netflow v10 Sample Records Bad Length; 'netflow-v10-max-records-exceed': Netflow v10 Sample Max records Error; 'netflow-tcp-sample-received': Netflow TCP Samples Received; 'netflow-udp-sample-received': Netflow UDP Samples received; 'netflow-icmp-sample-received': Netflow ICMP Samples Received; 'netflow-other-sample-received': Netflow OTHER Samples Received; 'netflow-record-copy-oom-error': Netflow Data Record Copy Fail OOM; 'netflow-record-rse-invalid': Netflow Data Record Reduced Size Invalid; 'netflow-sample-flow-dur-error': Netflow Sample Flow Duration Error;",
									},
								},
							},
						},
					},
				},
			},
			"debug_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_ip_addr": {
							Type: schema.TypeString, Required: true, Description: "Specify source/dest ip addr",
						},
						"debug_port": {
							Type: schema.TypeInt, Required: true, Description: "Specify port",
						},
						"debug_protocol": {
							Type: schema.TypeString, Required: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"delete_debug_file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify source/dest ip addr",
						},
						"debug_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify port",
						},
						"debug_protocol": {
							Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
						},
					},
				},
			},
			"index_sessions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start indexing associated sessions",
			},
			"index_sessions_type": {
				Type: schema.TypeString, Optional: true, Description: "'per-cpu': Use per cpu list;",
			},
			"mon_entity_topk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable topk for primary entities",
			},
			"monitor_key": {
				Type: schema.TypeString, Optional: true, Description: "'source': Monitor traffic from all sources; 'dest': Monitor traffic to any destination; 'service': Monitor traffic to any service; 'source-nat-ip': Monitor traffic to all source nat IPs;",
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
			"primary_monitor": {
				Type: schema.TypeString, Required: true, Description: "'traffic': Mointor traffic; 'xflow': Monitor xflow samples;",
			},
			"replay_debug_file": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"debug_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify source/dest ip addr",
						},
						"debug_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify port",
						},
						"debug_protocol": {
							Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
						},
					},
				},
			},
			"secondary_monitor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"secondary_monitoring_key": {
							Type: schema.TypeString, Optional: true, Description: "'service': Monitor traffic to any service;",
						},
						"mon_entity_topk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable topk for secondary entities",
						},
						"source_entity_topk": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable topk for sources to secondary-entities",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"debug_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"debug_ip_addr": {
										Type: schema.TypeString, Required: true, Description: "Specify source/dest ip addr",
									},
									"debug_port": {
										Type: schema.TypeInt, Required: true, Description: "Specify port",
									},
									"debug_protocol": {
										Type: schema.TypeString, Required: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"delete_debug_file": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"debug_ip_addr": {
										Type: schema.TypeString, Optional: true, Description: "Specify source/dest ip addr",
									},
									"debug_port": {
										Type: schema.TypeInt, Optional: true, Description: "Specify port",
									},
									"debug_protocol": {
										Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
									},
								},
							},
						},
						"replay_debug_file": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"debug_ip_addr": {
										Type: schema.TypeString, Optional: true, Description: "Specify source/dest ip addr",
									},
									"debug_port": {
										Type: schema.TypeInt, Optional: true, Description: "Specify port",
									},
									"debug_protocol": {
										Type: schema.TypeString, Optional: true, Description: "'TCP': TCP; 'UDP': UDP; 'ICMP': ICMP;",
									},
								},
							},
						},
					},
				},
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
			"source_entity_topk": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable topk for sources to primary-entities",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notif_template_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
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
func resourceVisibilityMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVisibilityMonitorAgentList(d []interface{}) []edpt.VisibilityMonitorAgentList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorAgentList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorAgentList
		oi.AgentName = in["agent_name"].(string)
		oi.AgentV4Addr = in["agent_v4_addr"].(string)
		oi.AgentV6Addr = in["agent_v6_addr"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceVisibilityMonitorAgentListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitorAgentListSamplingEnable(d []interface{}) []edpt.VisibilityMonitorAgentListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorAgentListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorAgentListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVisibilityMonitorDebugList(d []interface{}) []edpt.VisibilityMonitorDebugList {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorDebugList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorDebugList
		oi.DebugIpAddr = in["debug_ip_addr"].(string)
		oi.DebugPort = in["debug_port"].(int)
		oi.DebugProtocol = in["debug_protocol"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitorDeleteDebugFile1916(d []interface{}) edpt.VisibilityMonitorDeleteDebugFile1916 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorDeleteDebugFile1916
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugIpAddr = in["debug_ip_addr"].(string)
		ret.DebugPort = in["debug_port"].(int)
		ret.DebugProtocol = in["debug_protocol"].(string)
	}
	return ret
}

func getObjectVisibilityMonitorNetflow1917(d []interface{}) edpt.VisibilityMonitorNetflow1917 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorNetflow1917
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		ret.TemplateActiveTimeout = in["template_active_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityMonitorReplayDebugFile1918(d []interface{}) edpt.VisibilityMonitorReplayDebugFile1918 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorReplayDebugFile1918
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugIpAddr = in["debug_ip_addr"].(string)
		ret.DebugPort = in["debug_port"].(int)
		ret.DebugProtocol = in["debug_protocol"].(string)
	}
	return ret
}

func getObjectVisibilityMonitorSecondaryMonitor1919(d []interface{}) edpt.VisibilityMonitorSecondaryMonitor1919 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorSecondaryMonitor1919
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SecondaryMonitoringKey = in["secondary_monitoring_key"].(string)
		ret.MonEntityTopk = in["mon_entity_topk"].(int)
		ret.SourceEntityTopk = in["source_entity_topk"].(int)
		//omit uuid
		ret.DebugList = getSliceVisibilityMonitorSecondaryMonitorDebugList1920(in["debug_list"].([]interface{}))
		ret.DeleteDebugFile = getObjectVisibilityMonitorSecondaryMonitorDeleteDebugFile1921(in["delete_debug_file"].([]interface{}))
		ret.ReplayDebugFile = getObjectVisibilityMonitorSecondaryMonitorReplayDebugFile1922(in["replay_debug_file"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitorSecondaryMonitorDebugList1920(d []interface{}) []edpt.VisibilityMonitorSecondaryMonitorDebugList1920 {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorSecondaryMonitorDebugList1920, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorSecondaryMonitorDebugList1920
		oi.DebugIpAddr = in["debug_ip_addr"].(string)
		oi.DebugPort = in["debug_port"].(int)
		oi.DebugProtocol = in["debug_protocol"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVisibilityMonitorSecondaryMonitorDeleteDebugFile1921(d []interface{}) edpt.VisibilityMonitorSecondaryMonitorDeleteDebugFile1921 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorSecondaryMonitorDeleteDebugFile1921
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugIpAddr = in["debug_ip_addr"].(string)
		ret.DebugPort = in["debug_port"].(int)
		ret.DebugProtocol = in["debug_protocol"].(string)
	}
	return ret
}

func getObjectVisibilityMonitorSecondaryMonitorReplayDebugFile1922(d []interface{}) edpt.VisibilityMonitorSecondaryMonitorReplayDebugFile1922 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorSecondaryMonitorReplayDebugFile1922
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DebugIpAddr = in["debug_ip_addr"].(string)
		ret.DebugPort = in["debug_port"].(int)
		ret.DebugProtocol = in["debug_protocol"].(string)
	}
	return ret
}

func getObjectVisibilityMonitorSflow1923(d []interface{}) edpt.VisibilityMonitorSflow1923 {

	count1 := len(d)
	var ret edpt.VisibilityMonitorSflow1923
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ListeningPort = in["listening_port"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityMonitorTemplate(d []interface{}) edpt.VisibilityMonitorTemplate {

	count1 := len(d)
	var ret edpt.VisibilityMonitorTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Notification = getSliceVisibilityMonitorTemplateNotification(in["notification"].([]interface{}))
	}
	return ret
}

func getSliceVisibilityMonitorTemplateNotification(d []interface{}) []edpt.VisibilityMonitorTemplateNotification {

	count1 := len(d)
	ret := make([]edpt.VisibilityMonitorTemplateNotification, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VisibilityMonitorTemplateNotification
		oi.NotifTemplateName = in["notif_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVisibilityMonitor(d *schema.ResourceData) edpt.VisibilityMonitor {
	var ret edpt.VisibilityMonitor
	ret.Inst.AgentList = getSliceVisibilityMonitorAgentList(d.Get("agent_list").([]interface{}))
	ret.Inst.DebugList = getSliceVisibilityMonitorDebugList(d.Get("debug_list").([]interface{}))
	ret.Inst.DeleteDebugFile = getObjectVisibilityMonitorDeleteDebugFile1916(d.Get("delete_debug_file").([]interface{}))
	ret.Inst.IndexSessions = d.Get("index_sessions").(int)
	ret.Inst.IndexSessionsType = d.Get("index_sessions_type").(string)
	ret.Inst.MonEntityTopk = d.Get("mon_entity_topk").(int)
	ret.Inst.MonitorKey = d.Get("monitor_key").(string)
	ret.Inst.Netflow = getObjectVisibilityMonitorNetflow1917(d.Get("netflow").([]interface{}))
	ret.Inst.PrimaryMonitor = d.Get("primary_monitor").(string)
	ret.Inst.ReplayDebugFile = getObjectVisibilityMonitorReplayDebugFile1918(d.Get("replay_debug_file").([]interface{}))
	ret.Inst.SecondaryMonitor = getObjectVisibilityMonitorSecondaryMonitor1919(d.Get("secondary_monitor").([]interface{}))
	ret.Inst.Sflow = getObjectVisibilityMonitorSflow1923(d.Get("sflow").([]interface{}))
	ret.Inst.SourceEntityTopk = d.Get("source_entity_topk").(int)
	ret.Inst.Template = getObjectVisibilityMonitorTemplate(d.Get("template").([]interface{}))
	//omit uuid
	return ret
}
