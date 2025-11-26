package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectSubNetworkSubNetworkV6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_sub_network_sub_network_v6`: Configure sub-network in a DDos Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectSubNetworkSubNetworkV6Create,
		UpdateContext: resourceDdosNetworkObjectSubNetworkSubNetworkV6Update,
		ReadContext:   resourceDdosNetworkObjectSubNetworkSubNetworkV6Read,
		DeleteContext: resourceDdosNetworkObjectSubNetworkSubNetworkV6Delete,

		Schema: map[string]*schema.Schema{
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
			"subnet_ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "IPv6 Subnet/host, supported prefix range is from 56 to 64",
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
func resourceDdosNetworkObjectSubNetworkSubNetworkV6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSubNetworkSubNetworkV6Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectSubNetworkSubNetworkV6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectSubNetworkSubNetworkV6Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectSubNetworkSubNetworkV6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectSubNetworkSubNetworkV6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectSubNetworkSubNetworkV6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV6HostAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV6HostAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6HostAnomalyThreshold
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

func getSliceDdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable(d []interface{}) []edpt.DdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectSubNetworkSubNetworkV6SubNetworkAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectSubNetworkSubNetworkV6SubNetworkAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6SubNetworkAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticSubNetworkPktRate = in["static_sub_network_pkt_rate"].(int)
		ret.StaticSubNetworkBitRate = in["static_sub_network_bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectSubNetworkSubNetworkV6(d *schema.ResourceData) edpt.DdosNetworkObjectSubNetworkSubNetworkV6 {
	var ret edpt.DdosNetworkObjectSubNetworkSubNetworkV6
	ret.Inst.HostAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV6HostAnomalyThreshold(d.Get("host_anomaly_threshold").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosNetworkObjectSubNetworkSubNetworkV6SamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SubNetworkAnomalyThreshold = getObjectDdosNetworkObjectSubNetworkSubNetworkV6SubNetworkAnomalyThreshold(d.Get("sub_network_anomaly_threshold").([]interface{}))
	ret.Inst.SubnetBreakdown = d.Get("subnet_breakdown").(int)
	ret.Inst.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
