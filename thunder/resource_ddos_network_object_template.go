package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_template`: Network-object Template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectTemplateCreate,
		UpdateContext: resourceDdosNetworkObjectTemplateUpdate,
		ReadContext:   resourceDdosNetworkObjectTemplateRead,
		DeleteContext: resourceDdosNetworkObjectTemplateDelete,

		Schema: map[string]*schema.Schema{
			"anomaly_detection_trigger": {
				Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': Use both learned and static thresholds (static thresholds take precedence); 'static-threshold-only': Use static thresholds only;",
			},
			"flooding_multiplier": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "multiplier for flooding detection threshold in network objects (default 2x threshold)",
			},
			"histogram_mode": {
				Type: schema.TypeString, Optional: true, Default: "observe", Description: "'off': histogram feature disabled; 'monitor': histogram feature enabled with anomaly escalation; 'observe': histogram feature enabled and observe only;",
			},
			"host_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Forward packet rate of per host",
						},
						"host_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Forward bit rate of per host",
						},
						"host_rev_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse packet rate of per host",
						},
						"host_rev_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse bit rate of per host",
						},
						"host_undiscovered_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Undiscovered forward packet rate of per host",
						},
						"host_flow_count": {
							Type: schema.TypeInt, Optional: true, Description: "Flow count of per host",
						},
						"host_syn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "SYN packet rate of per host",
						},
						"host_fin_rate": {
							Type: schema.TypeInt, Optional: true, Description: "FIN packet rate of per host",
						},
						"host_rst_rate": {
							Type: schema.TypeInt, Optional: true, Description: "RST packet rate of per host",
						},
						"host_tcp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Tcp packet rate of per host",
						},
						"host_udp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Udp packet rate of per host",
						},
						"host_icmp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP packet rate of per host",
						},
						"host_undiscovered_host_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "forward packet rate of per undiscovered host",
						},
						"host_undiscovered_host_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Forward bit rate of per undiscovered host",
						},
					},
				},
			},
			"indicators_to_monitor": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"monitor_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward packet rate",
						},
						"monitor_bit_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward bit rate",
						},
						"monitor_rev_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse packet rate",
						},
						"monitor_rev_bit_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Reverse bit rate",
						},
						"monitor_undiscovered_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Undiscovered forward packet rate",
						},
						"monitor_flow_count": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Flow count",
						},
						"monitor_syn_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "SYN packet rate",
						},
						"monitor_fin_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "FIN packet rate",
						},
						"monitor_rst_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "RST packet rate",
						},
						"monitor_tcp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP packet rate",
						},
						"monitor_udp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "UDP packet rate",
						},
						"monitor_icmp_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ICMP packet rate",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS network-object-template name",
			},
			"network_object_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_object_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of the network-object",
						},
						"network_object_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of the network-object",
						},
					},
				},
			},
			"operational_mode": {
				Type: schema.TypeString, Optional: true, Default: "learning", Description: "'monitor': Monitor mode; 'learning': Learning mode;",
			},
			"service_break_down_threshold_local": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"svc_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "percentage of parent ip node",
						},
					},
				},
			},
			"service_discovery": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable service discovery for hosts (default: enabled);",
			},
			"sport_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of a source port entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"packet_rate_percentage": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"bit_rate": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Bit rate of a source port entry",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"bit_rate_percentage": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type: schema.TypeInt, Optional: true, Description: "Percentage of source port entry's parent entry",
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
			"threshold_sensitivity": {
				Type: schema.TypeString, Optional: true, Default: "OFF", Description: "tune threshold ranges with levels LOW/MEDIUM/HIGH/OFF (default) or multiplier of threshold value (available options are LOW=5x/MEDIUM=3x/HIGH=1.5x/OFF=1x, or float value between 1.0-10.0)",
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
func resourceDdosNetworkObjectTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectTemplateHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectTemplateHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostPktRate = in["host_pkt_rate"].(int)
		ret.HostBitRate = in["host_bit_rate"].(int)
		ret.HostRevPktRate = in["host_rev_pkt_rate"].(int)
		ret.HostRevBitRate = in["host_rev_bit_rate"].(int)
		ret.HostUndiscoveredPktRate = in["host_undiscovered_pkt_rate"].(int)
		ret.HostFlowCount = in["host_flow_count"].(int)
		ret.HostSynRate = in["host_syn_rate"].(int)
		ret.HostFinRate = in["host_fin_rate"].(int)
		ret.HostRstRate = in["host_rst_rate"].(int)
		ret.HostTcpPktRate = in["host_tcp_pkt_rate"].(int)
		ret.HostUdpPktRate = in["host_udp_pkt_rate"].(int)
		ret.HostIcmpPktRate = in["host_icmp_pkt_rate"].(int)
		ret.HostUndiscoveredHostPktRate = in["host_undiscovered_host_pkt_rate"].(int)
		ret.HostUndiscoveredHostBitRate = in["host_undiscovered_host_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateIndicatorsToMonitor306(d []interface{}) edpt.DdosNetworkObjectTemplateIndicatorsToMonitor306 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateIndicatorsToMonitor306
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.MonitorPktRate = in["monitor_pkt_rate"].(int)
		ret.MonitorBitRate = in["monitor_bit_rate"].(int)
		ret.MonitorRevPktRate = in["monitor_rev_pkt_rate"].(int)
		ret.MonitorRevBitRate = in["monitor_rev_bit_rate"].(int)
		ret.MonitorUndiscoveredPktRate = in["monitor_undiscovered_pkt_rate"].(int)
		ret.MonitorFlowCount = in["monitor_flow_count"].(int)
		ret.MonitorSynRate = in["monitor_syn_rate"].(int)
		ret.MonitorFinRate = in["monitor_fin_rate"].(int)
		ret.MonitorRstRate = in["monitor_rst_rate"].(int)
		ret.MonitorTcpPktRate = in["monitor_tcp_pkt_rate"].(int)
		ret.MonitorUdpPktRate = in["monitor_udp_pkt_rate"].(int)
		ret.MonitorIcmpPktRate = in["monitor_icmp_pkt_rate"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateNetworkObjectAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectTemplateNetworkObjectAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateNetworkObjectAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkObjectPktRate = in["network_object_pkt_rate"].(int)
		ret.NetworkObjectBitRate = in["network_object_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateServiceBreakDownThresholdLocal(d []interface{}) edpt.DdosNetworkObjectTemplateServiceBreakDownThresholdLocal {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateServiceBreakDownThresholdLocal
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SvcPercentage = in["svc_percentage"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateSportAnomalyThreshold307(d []interface{}) edpt.DdosNetworkObjectTemplateSportAnomalyThreshold307 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThreshold307
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PacketRate = getObjectDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate308(in["packet_rate"].([]interface{}))
		ret.PacketRatePercentage = getObjectDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage309(in["packet_rate_percentage"].([]interface{}))
		ret.BitRate = getObjectDdosNetworkObjectTemplateSportAnomalyThresholdBitRate310(in["bit_rate"].([]interface{}))
		ret.BitRatePercentage = getObjectDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage311(in["bit_rate_percentage"].([]interface{}))
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateSportAnomalyThresholdPacketRate308(d []interface{}) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate308 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRate308
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage309(d []interface{}) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage309 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdPacketRatePercentage309
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateSportAnomalyThresholdBitRate310(d []interface{}) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRate310 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRate310
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage311(d []interface{}) edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage311 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectTemplateSportAnomalyThresholdBitRatePercentage311
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosNetworkObjectTemplate(d *schema.ResourceData) edpt.DdosNetworkObjectTemplate {
	var ret edpt.DdosNetworkObjectTemplate
	ret.Inst.AnomalyDetectionTrigger = d.Get("anomaly_detection_trigger").(string)
	ret.Inst.FloodingMultiplier = d.Get("flooding_multiplier").(int)
	ret.Inst.HistogramMode = d.Get("histogram_mode").(string)
	ret.Inst.HostAnomalyThreshold = getObjectDdosNetworkObjectTemplateHostAnomalyThreshold(d.Get("host_anomaly_threshold").([]interface{}))
	ret.Inst.IndicatorsToMonitor = getObjectDdosNetworkObjectTemplateIndicatorsToMonitor306(d.Get("indicators_to_monitor").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NetworkObjectAnomalyThreshold = getObjectDdosNetworkObjectTemplateNetworkObjectAnomalyThreshold(d.Get("network_object_anomaly_threshold").([]interface{}))
	ret.Inst.OperationalMode = d.Get("operational_mode").(string)
	ret.Inst.ServiceBreakDownThresholdLocal = getObjectDdosNetworkObjectTemplateServiceBreakDownThresholdLocal(d.Get("service_break_down_threshold_local").([]interface{}))
	ret.Inst.ServiceDiscovery = d.Get("service_discovery").(string)
	ret.Inst.SportAnomalyThreshold = getObjectDdosNetworkObjectTemplateSportAnomalyThreshold307(d.Get("sport_anomaly_threshold").([]interface{}))
	ret.Inst.ThresholdSensitivity = d.Get("threshold_sensitivity").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
