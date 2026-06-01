package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosNetworkObjectIpSrcPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_network_object_ip_src_port`: Configure per-host source port in a DDos Network Object\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosNetworkObjectIpSrcPortCreate,
		UpdateContext: resourceDdosNetworkObjectIpSrcPortUpdate,
		ReadContext:   resourceDdosNetworkObjectIpSrcPortRead,
		DeleteContext: resourceDdosNetworkObjectIpSrcPortDelete,

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
			"subnet_ip_addr": {
				Type: schema.TypeString, Required: true, Description: "SubnetIpAddr",
			},
			"object_name": {
				Type: schema.TypeString, Required: true, Description: "ObjectName",
			},
		},
	}
}
func resourceDdosNetworkObjectIpSrcPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpSrcPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpSrcPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpSrcPortRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosNetworkObjectIpSrcPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpSrcPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpSrcPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosNetworkObjectIpSrcPortRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosNetworkObjectIpSrcPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpSrcPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpSrcPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosNetworkObjectIpSrcPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosNetworkObjectIpSrcPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosNetworkObjectIpSrcPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosNetworkObjectIpSrcPortHostSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpSrcPortHostSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpSrcPortHostSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HostSrcPortPktRate = in["host_src_port_pkt_rate"].(int)
		ret.HostSrcPortBitRate = in["host_src_port_bit_rate"].(int)
	}
	return ret
}

func getObjectDdosNetworkObjectIpSrcPortSubnetSrcPortAnomalyThreshold(d []interface{}) edpt.DdosNetworkObjectIpSrcPortSubnetSrcPortAnomalyThreshold {

	count1 := len(d)
	var ret edpt.DdosNetworkObjectIpSrcPortSubnetSrcPortAnomalyThreshold
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SubnetSrcPortPktRate = in["subnet_src_port_pkt_rate"].(int)
		ret.SubnetSrcPortBitRate = in["subnet_src_port_bit_rate"].(int)
	}
	return ret
}

func dataToEndpointDdosNetworkObjectIpSrcPort(d *schema.ResourceData) edpt.DdosNetworkObjectIpSrcPort {
	var ret edpt.DdosNetworkObjectIpSrcPort
	ret.Inst.HostSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpSrcPortHostSrcPortAnomalyThreshold(d.Get("host_src_port_anomaly_threshold").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.SubnetSrcPortAnomalyThreshold = getObjectDdosNetworkObjectIpSrcPortSubnetSrcPortAnomalyThreshold(d.Get("subnet_src_port_anomaly_threshold").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SubnetIpAddr = d.Get("subnet_ip_addr").(string)
	ret.Inst.ObjectName = d.Get("object_name").(string)
	return ret
}
