package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_ipv6`: DDOS Network-object IPv6 Subnet\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectIpv6Create,
		UpdateContext: resourceDdosNetworkObjectIpv6Update,
		ReadContext:   resourceDdosNetworkObjectIpv6Read,
		DeleteContext: resourceDdosNetworkObjectIpv6Delete,

		Schema: map[string]*schema.Schema{
			"prefix_anomaly_threshold": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix_pkt_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate of the prefix subnet",
						},
						"prefix_bit_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Bit rate of the prefix subnet",
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
			"src_port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Port Number",
						},
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
						},
						"host_src_port_anomaly_threshold": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host_src_port_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Forward packet rate of per-host source port entries",
									},
									"host_src_port_bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Forward bit rate of per-host source port entries",
									},
								},
							},
						},
						"subnet_src_port_anomaly_threshold": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet_src_port_pkt_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Forward packet rate of per-subnet source port entries",
									},
									"subnet_src_port_bit_rate": {
										Type: schema.TypeInt, Optional: true, Description: "Forward bit rate of per-subnet source port entries",
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
			"subnet_ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV6 Subnet, supported prefix range is from 40 to 64",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceDdosNetworkObjectIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectIpv6PrefixAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpv6PrefixAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6PrefixAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixPktRate = in["prefix_pkt_rate"].(int)
		ret.PrefixBitRate = in["prefix_bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectIpv6SamplingEnable(d []interface{}) []edpt.DdosNetworkObjectIpv6SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpv6SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpv6SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectIpv6SrcPortList(d []interface{}) []edpt.DdosNetworkObjectIpv6SrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpv6SrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpv6SrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.HostSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpv6SrcPortListHostSrcPortAnomalyThreshold(in["host_src_port_anomaly_threshold"].([]interface{}))
		oi.SubnetSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpv6SrcPortListSubnetSrcPortAnomalyThreshold(in["subnet_src_port_anomaly_threshold"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectIpv6SrcPortListHostSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpv6SrcPortListHostSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6SrcPortListHostSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostSrcPortPktRate = in["host_src_port_pkt_rate"].(int)
		ret.HostSrcPortBitRate = in["host_src_port_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectIpv6SrcPortListSubnetSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpv6SrcPortListSubnetSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6SrcPortListSubnetSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SubnetSrcPortPktRate = in["subnet_src_port_pkt_rate"].(int)
		ret.SubnetSrcPortBitRate = in["subnet_src_port_bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectIpv6(d *schema.ResourceData) edpt.DdosNetworkObjectIpv6 {
	var ret edpt.DdosNetworkObjectIpv6
	ret.Inst.PrefixAnomalyThreshold = getObjectDdosNetworkObjectIpv6PrefixAnomalyThreshold(d.Get("prefix_anomaly_threshold").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosNetworkObjectIpv6SamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SrcPortList = getSliceDdosNetworkObjectIpv6SrcPortList(d.Get("src_port_list").([]interface{}))
	ret.Inst.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
