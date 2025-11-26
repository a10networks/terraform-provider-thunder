package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSubNetworkSubNetworkV4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sub_network_sub_network_v4`: Configure sub-network in a DDos Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSubNetworkSubNetworkV4Create,
		UpdateContext: resourceDdosNetworkObjectSubNetworkSubNetworkV4Update,
		ReadContext:   resourceDdosNetworkObjectSubNetworkSubNetworkV4Read,
		DeleteContext: resourceDdosNetworkObjectSubNetworkSubNetworkV4Delete,

		Schema: map[string]*schema.Schema{
			"breakdown_subnet_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"breakdown_subnet_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
						},
						"breakdown_subnet_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
						},
					},
				},
			},
			"host_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
						},
						"static_rev_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of per host",
						},
						"static_bit_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
						},
						"static_rev_bit_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of per host",
						},
						"static_undiscovered_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Undiscovered packet rate of per host",
						},
						"static_flow_count_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Flow count of per host",
						},
						"static_syn_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "SYN packet rate of per host",
						},
						"static_fin_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "FIN packet rate of per host",
						},
						"static_rst_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "RST packet rate of per host",
						},
						"static_tcp_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "TCP packet rate of per host",
						},
						"static_udp_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "UDP packet rate of per host",
						},
						"static_icmp_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP packet rate of per host",
						},
						"static_undiscovered_host_pkt_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "packet rate of per undiscovered host",
						},
						"static_undiscovered_host_bit_rate_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of per undiscovered host",
						},
					},
				},
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_rate': PPS; 'bit_rate': B(bits)PS;",
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
						"static_sub_network_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of the sub-network",
						},
					},
				},
			},
			"subnet_breakdown": {
				Type: schema.TypeInt, Optional: true, Description: "additional layer of breakdown subnet",
			},
			"subnet_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "IPv4 Subnet/host, supported prefix range is from 24 to 32",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectSubNetworkSubNetworkV4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSubNetworkSubNetworkV4Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSubNetworkSubNetworkV4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSubNetworkSubNetworkV4Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSubNetworkSubNetworkV4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSubNetworkSubNetworkV4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4BreakdownSubnetThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4BreakdownSubnetThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4BreakdownSubnetThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.BreakdownSubnetPktRate = in["breakdown_subnet_pkt_rate"].(int)
		ret.BreakdownSubnetBitRate = in["breakdown_subnet_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4HostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4HostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4HostAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticPktRateThreshold = in["static_pkt_rate_threshold"].(int)
		ret.StaticRevPktRateThreshold = in["static_rev_pkt_rate_threshold"].(int)
		ret.StaticBitRateThreshold = in["static_bit_rate_threshold"].(int)
		ret.StaticRevBitRateThreshold = in["static_rev_bit_rate_threshold"].(int)
		ret.StaticUndiscoveredPktRateThreshold = in["static_undiscovered_pkt_rate_threshold"].(int)
		ret.StaticFlowCountThreshold = in["static_flow_count_threshold"].(int)
		ret.StaticSynRateThreshold = in["static_syn_rate_threshold"].(int)
		ret.StaticFinRateThreshold = in["static_fin_rate_threshold"].(int)
		ret.StaticRstRateThreshold = in["static_rst_rate_threshold"].(int)
		ret.StaticTcpPktRateThreshold = in["static_tcp_pkt_rate_threshold"].(int)
		ret.StaticUdpPktRateThreshold = in["static_udp_pkt_rate_threshold"].(int)
		ret.StaticIcmpPktRateThreshold = in["static_icmp_pkt_rate_threshold"].(int)
		ret.StaticUndiscoveredHostPktRateThreshold = in["static_undiscovered_host_pkt_rate_threshold"].(int)
		ret.StaticUndiscoveredHostBitRateThreshold = in["static_undiscovered_host_bit_rate_threshold"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable(d []interface{}) []edpt.DdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV4SubNetworkAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV4SubNetworkAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4SubNetworkAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticSubNetworkPktRate = in["static_sub_network_pkt_rate"].(int)
		ret.StaticSubNetworkBitRate = in["static_sub_network_bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV4(d *schema.ResourceData) edpt.DdosNetworkObjectSubNetworkSubNetworkV4 {
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV4
	ret.Inst.BreakdownSubnetThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV4BreakdownSubnetThreshold(d.Get("breakdown_subnet_threshold").([]interface{}))
	ret.Inst.HostAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV4HostAnomalyThreshold(d.Get("host_anomaly_threshold").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosNetworkObjectSubNetworkSubNetworkV4SamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SubNetworkAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV4SubNetworkAnomalyThreshold(d.Get("sub_network_anomaly_threshold").([]interface{}))
	ret.Inst.SubnetBreakdown = d.Get("subnet_breakdown").(int)
	ret.Inst.SubnetIpAddr = d.Get("subnet_ip_addr").(string)
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
