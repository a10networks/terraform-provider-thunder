package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObject() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object`: Configure DDoS a static Monitor Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectCreate,
		UpdateContext: resourceDdosNetworkObjectUpdate,
		ReadContext:   resourceDdosNetworkObjectRead,
		DeleteContext: resourceDdosNetworkObjectDelete,

		Schema: map[string]*schema.Schema{
			"anomaly_detection_trigger": {
				Type: schema.TypeString, Optional: true, Default: "all", Description: "'all': Use both learned and static thresholds (static thresholds take precedence); 'static-threshold-only': Use static thresholds only;",
			},
			"histogram_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable histogram statistics (Default: Disabled)",
			},
			"host_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
						},
						"host_byte_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Byte rate of per host",
						},
					},
				},
			},
			"ip": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ip_addr": {
							Type: schema.TypeString, Optional: true, Description: "IP Subnet, supported prefix range is from 8 to 31",
						},
					},
				},
			},
			"ipv6": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 Subnet, supported prefix range is from 40 to 63",
						},
					},
				},
			},
			"network_object_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_object_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of the network-object",
						},
						"network_object_byte_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Byte rate of the network-object",
						},
					},
				},
			},
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
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"operational_mode": {
				Type: schema.TypeString, Optional: true, Default: "learning", Description: "'monitor': Monitor mode; 'learning': Learning mode;",
			},
			"relative_auto_break_down_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "percentage of parent node",
						},
						"permil": {
							Type: schema.TypeInt, Optional: true, Description: "permil of root node",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'subnet_learned': Subnet Entry Learned; 'subnet_aged': Subnet Entry Aged; 'subnet_create_fail': Subnet Entry Create Failures; 'ip_learned': IP Entry Learned; 'ip_aged': IP Entry Aged; 'ip_create_fail': IP Entry Create Failures; 'service_learned': Service Entry Learned; 'service_aged': Service Entry Aged; 'service_create_fail': Service Entry Create Failures;",
						},
					},
				},
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
			"static_auto_break_down_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "packet rate of current node",
						},
					},
				},
			},
			"sub_network_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_ip_addr": {
							Type: schema.TypeString, Required: true, Description: "IPv4 Subnet/host, supported prefix range is from 24 to 32",
						},
						"host_anomaly_threshold": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_pkt_rate_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
									},
									"static_byte_rate_threshold": {
										Type: schema.TypeInt, Optional: true, Description: "Byte rate of per host",
									},
								},
							},
						},
						"sub_network_anomaly_threshold": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"static_sub_network_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Packet rate of the sub-network",
									},
									"static_sub_network_byte_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Byte rate of the sub-network",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosNetworkObjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObject(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostPktRate = in["host_pkt_rate"].(int)
		ret.HostByteRate = in["host_byte_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectIp(d []interface{}) []edpt.DdosNetworkObjectIp {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIp, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIp
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectIpv6(d []interface{}) []edpt.DdosNetworkObjectIpv6 {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpv6, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpv6
		oi.SubnetIpv6Addr = in["subnet_ipv6_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectNetworkObjectAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectNetworkObjectAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectNetworkObjectAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkObjectPktRate = in["network_object_pkt_rate"].(int)
		ret.NetworkObjectByteRate = in["network_object_byte_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectNotification282(d []interface{}) edpt.DdosNetworkObjectNotification282 {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectNotification282
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Configuration = in["configuration"].(string)
		ret.Notification = getSliceDdosNetworkObjectNotificationNotification283(in["notification"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceDdosNetworkObjectNotificationNotification283(d []interface{}) []edpt.DdosNetworkObjectNotificationNotification283 {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectNotificationNotification283, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectNotificationNotification283
		oi.NotificationTemplateName = in["notification_template_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectRelativeAutoBreakDownThreshold(d []interface{}) edpt.DdosNetworkObjectRelativeAutoBreakDownThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectRelativeAutoBreakDownThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkPercentage = in["network_percentage"].(int)
		ret.Permil = in["permil"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectServiceBreakDownThresholdLocal(d []interface{}) edpt.DdosNetworkObjectServiceBreakDownThresholdLocal {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectServiceBreakDownThresholdLocal
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SvcPercentage = in["svc_percentage"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectStaticAutoBreakDownThreshold(d []interface{}) edpt.DdosNetworkObjectStaticAutoBreakDownThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectStaticAutoBreakDownThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkPktRate = in["network_pkt_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSubNetworkList(d []interface{}) []edpt.DdosNetworkObjectSubNetworkList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkList
		oi.SubnetIpAddr = in["subnet_ip_addr"].(string)
		oi.HostAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkListHostAnomalyThreshold(in["host_anomaly_threshold"].([]interface{}))
		oi.SubNetworkAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkListSubNetworkAnomalyThreshold(in["sub_network_anomaly_threshold"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkListHostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkListHostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkListHostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticPktRateThreshold = in["static_pkt_rate_threshold"].(int)
		ret.StaticByteRateThreshold = in["static_byte_rate_threshold"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkListSubNetworkAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkListSubNetworkAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkListSubNetworkAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticSubNetworkPktRate = in["static_sub_network_pkt_rate"].(int)
		ret.StaticSubNetworkByteRate = in["static_sub_network_byte_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObject(d *schema.ResourceData) edpt.DdosNetworkObject {
	var ret edpt.DdosNetworkObject
	ret.Inst.AnomalyDetectionTrigger = d.Get("anomaly_detection_trigger").(string)
	ret.Inst.HistogramEnable = d.Get("histogram_enable").(int)
	ret.Inst.HostAnomalyThreshold = getObjectDdosNetworkObjectHostAnomalyThreshold(d.Get("host_anomaly_threshold").([]interface{}))
	ret.Inst.Ip = getSliceDdosNetworkObjectIp(d.Get("ip").([]interface{}))
	ret.Inst.Ipv6 = getSliceDdosNetworkObjectIpv6(d.Get("ipv6").([]interface{}))
	ret.Inst.NetworkObjectAnomalyThreshold = getObjectDdosNetworkObjectNetworkObjectAnomalyThreshold(d.Get("network_object_anomaly_threshold").([]interface{}))
	ret.Inst.Notification = getObjectDdosNetworkObjectNotification282(d.Get("notification").([]interface{}))
	ret.Inst.ObjectName = d.Get("object_name").(string)
	ret.Inst.OperationalMode = d.Get("operational_mode").(string)
	ret.Inst.RelativeAutoBreakDownThreshold = getObjectDdosNetworkObjectRelativeAutoBreakDownThreshold(d.Get("relative_auto_break_down_threshold").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosNetworkObjectSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceBreakDownThresholdLocal = getObjectDdosNetworkObjectServiceBreakDownThresholdLocal(d.Get("service_break_down_threshold_local").([]interface{}))
	ret.Inst.ServiceDiscovery = d.Get("service_discovery").(string)
	ret.Inst.StaticAutoBreakDownThreshold = getObjectDdosNetworkObjectStaticAutoBreakDownThreshold(d.Get("static_auto_break_down_threshold").([]interface{}))
	ret.Inst.SubNetworkList = getSliceDdosNetworkObjectSubNetworkList(d.Get("sub_network_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
