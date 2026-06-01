package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_ip`: DDOS Network-object IPv4 Subnet\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectIpCreate,
		UpdateContext: resourceDdosNetworkObjectIpUpdate,
		ReadContext:   resourceDdosNetworkObjectIpRead,
		DeleteContext: resourceDdosNetworkObjectIpDelete,

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
			"subnet_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "IP Subnet, supported prefix range is from 8 to 32",
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
func resourceDdosNetworkObjectIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectIpPrefixAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpPrefixAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpPrefixAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrefixPktRate = in["prefix_pkt_rate"].(int)
		ret.PrefixBitRate = in["prefix_bit_rate"].(int)
	}
	return ret
}

func getSliceDdosNetworkObjectIpSamplingEnable(d []interface{}) []edpt.DdosNetworkObjectIpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosNetworkObjectIpSrcPortList(d []interface{}) []edpt.DdosNetworkObjectIpSrcPortList {

	count1 := len(d)
	ret := make([]edpt.DdosNetworkObjectIpSrcPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosNetworkObjectIpSrcPortList
		oi.PortNum = in["port_num"].(int)
		oi.Protocol = in["protocol"].(string)
		oi.HostSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpSrcPortListHostSrcPortAnomalyThreshold(in["host_src_port_anomaly_threshold"].([]interface{}))
		oi.SubnetSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpSrcPortListSubnetSrcPortAnomalyThreshold(in["subnet_src_port_anomaly_threshold"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosNetworkObjectIpSrcPortListHostSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpSrcPortListHostSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpSrcPortListHostSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostSrcPortPktRate = in["host_src_port_pkt_rate"].(int)
		ret.HostSrcPortBitRate = in["host_src_port_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectIpSrcPortListSubnetSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpSrcPortListSubnetSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpSrcPortListSubnetSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SubnetSrcPortPktRate = in["subnet_src_port_pkt_rate"].(int)
		ret.SubnetSrcPortBitRate = in["subnet_src_port_bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectIp(d *schema.ResourceData) edpt.DdosNetworkObjectIp {
	var ret edpt.DdosNetworkObjectIp
	ret.Inst.PrefixAnomalyThreshold = getObjectDdosNetworkObjectIpPrefixAnomalyThreshold(d.Get("prefix_anomaly_threshold").([]interface{}))
	ret.Inst.SamplingEnable = getSliceDdosNetworkObjectIpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SrcPortList = getSliceDdosNetworkObjectIpSrcPortList(d.Get("src_port_list").([]interface{}))
	ret.Inst.SubnetIpAddr = d.Get("subnet_ip_addr").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
