package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection`: DDOS Detection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionCreate,
		UpdateContext: resourceDdosDstZoneDetectionUpdate,
		ReadContext:   resourceDdosDstZoneDetectionRead,
		DeleteContext: resourceDdosDstZoneDetectionDelete,

		Schema: map[string]*schema.Schema{
			"notification": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
						},
						"notification": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification_template_name": {
										Type: schema.TypeString, Optional: true, Description: "Specify the notification template name",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"outbound_detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
						},
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable outbound detection; 'disable': Disable outbound detection;",
						},
						"discovery_method": {
							Type: schema.TypeString, Optional: true, Description: "'asn': Autonomous Systems number; 'country': Country;",
						},
						"discovery_record": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Maximum number of top locations",
						},
						"enable_top_k": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topk_type": {
										Type: schema.TypeString, Optional: true, Description: "'source-subnet': Topk source subnet;",
									},
									"topk_netmask": {
										Type: schema.TypeInt, Optional: true, Default: 128, Description: "Subnet mask. The value should be less than or equal to the minimum zone subnet mask + 8 (IPv6 Subnet mask)",
									},
									"topk_num_records": {
										Type: schema.TypeInt, Optional: true, Default: 20, Description: "Maximum number of records to show in topk",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"indicator_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets;",
									},
									"tcp_window_size": {
										Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
									},
									"data_packet_size": {
										Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
									},
									"threshold_num": {
										Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
									},
									"threshold_large_num": {
										Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
									},
									"threshold_str": {
										Type: schema.TypeString, Optional: true, Description: "Threshold for each geo-location (Non-zero floating point)",
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
						"topk_source_subnet": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"packet_anomaly_detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
						},
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable packet anomaly; 'disable': Disable packet anomaly;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"indicator_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Required: true, Description: "'port-zero-pkt-rate': Port Zero Packet Rate (default 100 packet per second);",
									},
									"threshold_num": {
										Type: schema.TypeInt, Optional: true, Default: 100, Description: "Threshold for each indicator",
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
					},
				},
			},
			"service_discovery": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
						},
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable service discovery; 'disable': Disable service discovery;",
						},
						"pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "packet rate threshold for discovery (default 10 packets per second)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"settings": {
				Type: schema.TypeString, Required: true, Description: "'settings': settings;",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable detection; 'disable': Disable detection;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"victim_ip_detection": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration": {
							Type: schema.TypeString, Optional: true, Description: "'configuration': configuration;",
						},
						"toggle": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable victim IP detection; 'disable': Disable victim IP detection;",
						},
						"histogram_toggle": {
							Type: schema.TypeString, Optional: true, Default: "histogram-disable", Description: "'histogram-enable': Enable histogram statistics of victim IP detection; 'histogram-disable': Disable histogram statistics of victim IP detection;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"indicator_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'reverse-pkt-rate': rate of reverse coming packets; 'fwd-byte-rate': rate of incoming bytes; 'rev-byte-rate': rate of reverse coming bytes;",
									},
									"ip_threshold_num": {
										Type: schema.TypeInt, Optional: true, Description: "Threshold for IP",
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
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZoneDetectionNotification188(d []interface{}) edpt.DdosDstZoneDetectionNotification188 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionNotification188
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Notification = getSliceDdosDstZoneDetectionNotificationNotification189(in["notification"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosDstZoneDetectionNotificationNotification189(d []interface{}) []edpt.DdosDstZoneDetectionNotificationNotification189 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionNotificationNotification189, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionNotificationNotification189
		oi.NotificationTemplateName = in["notification_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetection190(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetection190 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionOutboundDetection190
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		ret.DiscoveryMethod = in["discovery_method"].(string)
		ret.DiscoveryRecord = in["discovery_record"].(int)
		ret.EnableTopK = getSliceDdosDstZoneDetectionOutboundDetectionEnableTopK191(in["enable_top_k"].([]interface{}))
		//omit uuid
		ret.IndicatorList = getSliceDdosDstZoneDetectionOutboundDetectionIndicatorList192(in["indicator_list"].([]interface{}))
		ret.TopkSourceSubnet = getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnet193(in["topk_source_subnet"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionEnableTopK191(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK191 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK191, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionEnableTopK191
		oi.TopkType = in["topk_type"].(string)
		oi.TopkNetmask = in["topk_netmask"].(int)
		oi.TopkNumRecords = in["topk_num_records"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstZoneDetectionOutboundDetectionIndicatorList192(d []interface{}) []edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList192 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList192, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionOutboundDetectionIndicatorList192
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.ThresholdNum = in["threshold_num"].(int)
		oi.ThresholdLargeNum = in["threshold_large_num"].(int)
		oi.ThresholdStr = in["threshold_str"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionOutboundDetectionTopkSourceSubnet193(d []interface{}) edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet193 {

	var ret edpt.DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet193
	return ret
}

func getObjectDdosDstZoneDetectionPacketAnomalyDetection194(d []interface{}) edpt.DdosDstZoneDetectionPacketAnomalyDetection194 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionPacketAnomalyDetection194
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		//omit uuid
		ret.IndicatorList = getSliceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195(in["indicator_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195(d []interface{}) []edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList195
		oi.Type = in["type"].(string)
		oi.ThresholdNum = in["threshold_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneDetectionServiceDiscovery196(d []interface{}) edpt.DdosDstZoneDetectionServiceDiscovery196 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionServiceDiscovery196
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		ret.PktRateThreshold = in["pkt_rate_threshold"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosDstZoneDetectionVictimIpDetection197(d []interface{}) edpt.DdosDstZoneDetectionVictimIpDetection197 {

	count1 := len(d)
	var ret edpt.DdosDstZoneDetectionVictimIpDetection197
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Toggle = in["toggle"].(string)
		ret.HistogramToggle = in["histogram_toggle"].(string)
		//omit uuid
		ret.IndicatorList = getSliceDdosDstZoneDetectionVictimIpDetectionIndicatorList198(in["indicator_list"].([]interface{}))
	}
	return ret
}

func getSliceDdosDstZoneDetectionVictimIpDetectionIndicatorList198(d []interface{}) []edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList198 {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList198, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList198
		oi.Type = in["type"].(string)
		oi.IpThresholdNum = in["ip_threshold_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetection(d *schema.ResourceData) edpt.DdosDstZoneDetection {
	var ret edpt.DdosDstZoneDetection
	ret.Inst.Notification = getObjectDdosDstZoneDetectionNotification188(d.Get("notification").([]interface{}))
	ret.Inst.OutboundDetection = getObjectDdosDstZoneDetectionOutboundDetection190(d.Get("outbound_detection").([]interface{}))
	ret.Inst.PacketAnomalyDetection = getObjectDdosDstZoneDetectionPacketAnomalyDetection194(d.Get("packet_anomaly_detection").([]interface{}))
	ret.Inst.ServiceDiscovery = getObjectDdosDstZoneDetectionServiceDiscovery196(d.Get("service_discovery").([]interface{}))
	ret.Inst.Settings = d.Get("settings").(string)
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	ret.Inst.VictimIpDetection = getObjectDdosDstZoneDetectionVictimIpDetection197(d.Get("victim_ip_detection").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
