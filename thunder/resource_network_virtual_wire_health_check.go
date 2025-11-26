package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVirtualWireHealthCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_virtual_wire_health_check`: Virtual Wire Health Check Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVirtualWireHealthCheckCreate,
		UpdateContext: resourceNetworkVirtualWireHealthCheckUpdate,
		ReadContext:   resourceNetworkVirtualWireHealthCheckRead,
		DeleteContext: resourceNetworkVirtualWireHealthCheckDelete,

		Schema: map[string]*schema.Schema{
			"active_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold(Packet Per Second) for entering active mode",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the health check",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"garp_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Interval (Seconds), default is 3, 0 for interface event only",
			},
			"inner_vlan": {
				Type: schema.TypeInt, Optional: true, Description: "inner VLAN for 802.1QinQ",
			},
			"inner_vlan_packet_count": {
				Type: schema.TypeInt, Optional: true, Description: "Only count packets with configured inner VLAN ID",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "Interval (Seconds), default is 3",
			},
			"l3_packet": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only count L3 packets",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Description: "'ping': Ping; 'garp': GARP; 'packet-count': Packet Count;",
			},
			"nexthop_ip": {
				Type: schema.TypeString, Optional: true, Description: "Nexthop address",
			},
			"nexthop_mac": {
				Type: schema.TypeString, Optional: true, Description: "Nexthop mac address",
			},
			"partition_health_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Health check state stands for the partition",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'act-event': Active Event Count; 'sby-event': Standby Event Count; 'packet-count': Packet Count;",
						},
					},
				},
			},
			"sby_ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
			},
			"sby_trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk interface",
			},
			"source_ip": {
				Type: schema.TypeString, Optional: true, Description: "source ip for ping method",
			},
			"source_mac": {
				Type: schema.TypeString, Optional: true, Description: "source mac for ping method",
			},
			"trunk": {
				Type: schema.TypeInt, Optional: true, Description: "Trunk interface",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vlan": {
				Type: schema.TypeInt, Required: true, Description: "VLAN ID, specify 1 for untagged traffic",
			},
		},
	}
}
func resourceNetworkVirtualWireHealthCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireHealthCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireHealthCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireHealthCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVirtualWireHealthCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireHealthCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireHealthCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVirtualWireHealthCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVirtualWireHealthCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireHealthCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireHealthCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVirtualWireHealthCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVirtualWireHealthCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVirtualWireHealthCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVirtualWireHealthCheckSamplingEnable(d []interface{}) []edpt.NetworkVirtualWireHealthCheckSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.NetworkVirtualWireHealthCheckSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVirtualWireHealthCheckSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVirtualWireHealthCheck(d *schema.ResourceData) edpt.NetworkVirtualWireHealthCheck {
	var ret edpt.NetworkVirtualWireHealthCheck
	ret.Inst.ActiveThreshold = d.Get("active_threshold").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.GarpInterval = d.Get("garp_interval").(int)
	ret.Inst.InnerVlan = d.Get("inner_vlan").(int)
	ret.Inst.InnerVlanPacketCount = d.Get("inner_vlan_packet_count").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.L3Packet = d.Get("l3_packet").(int)
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.NexthopIp = d.Get("nexthop_ip").(string)
	ret.Inst.NexthopMac = d.Get("nexthop_mac").(string)
	ret.Inst.PartitionHealthCheck = d.Get("partition_health_check").(int)
	ret.Inst.SamplingEnable = getSliceNetworkVirtualWireHealthCheckSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SbyEthernet = d.Get("sby_ethernet").(int)
	ret.Inst.SbyTrunk = d.Get("sby_trunk").(int)
	ret.Inst.SourceIp = d.Get("source_ip").(string)
	ret.Inst.SourceMac = d.Get("source_mac").(string)
	ret.Inst.Trunk = d.Get("trunk").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vlan = d.Get("vlan").(int)
	return ret
}
