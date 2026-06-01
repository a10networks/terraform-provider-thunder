package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectIpv6SrcPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_ipv6_src_port`: Configure per-host source port in a DDos Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectIpv6SrcPortCreate,
		UpdateContext: resourceDdosNetworkObjectIpv6SrcPortUpdate,
		ReadContext:   resourceDdosNetworkObjectIpv6SrcPortRead,
		DeleteContext: resourceDdosNetworkObjectIpv6SrcPortDelete,

		Schema: map[string]*schema.Schema{
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
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'udp': UDP port; 'tcp': TCP Port;",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"subnet_ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "SubnetIpv6Addr",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectIpv6SrcPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6SrcPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6SrcPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpv6SrcPortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectIpv6SrcPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6SrcPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6SrcPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpv6SrcPortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectIpv6SrcPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6SrcPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6SrcPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectIpv6SrcPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpv6SrcPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpv6SrcPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectIpv6SrcPortHostSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpv6SrcPortHostSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6SrcPortHostSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostSrcPortPktRate = in["host_src_port_pkt_rate"].(int)
		ret.HostSrcPortBitRate = in["host_src_port_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectIpv6SrcPortSubnetSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpv6SrcPortSubnetSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpv6SrcPortSubnetSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SubnetSrcPortPktRate = in["subnet_src_port_pkt_rate"].(int)
		ret.SubnetSrcPortBitRate = in["subnet_src_port_bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectIpv6SrcPort(d *schema.ResourceData) edpt.DdosNetworkObjectIpv6SrcPort {
	var ret edpt.DdosNetworkObjectIpv6SrcPort
	ret.Inst.HostSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpv6SrcPortHostSrcPortAnomalyThreshold(d.Get("host_src_port_anomaly_threshold").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SubnetSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpv6SrcPortSubnetSrcPortAnomalyThreshold(d.Get("subnet_src_port_anomaly_threshold").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
