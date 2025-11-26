package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelIpv6Ospf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_tunnel_ipv6_ospf`: Open Shortest Path First for IPv6 (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTunnelIpv6OspfCreate,
		UpdateContext: resourceInterfaceTunnelIpv6OspfUpdate,
		ReadContext:   resourceInterfaceTunnelIpv6OspfRead,
		DeleteContext: resourceInterfaceTunnelIpv6OspfDelete,

		Schema: map[string]*schema.Schema{
			"bfd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
			},
			"cost_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": {
							Type: schema.TypeInt, Optional: true, Description: "Interface cost",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"dead_interval_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dead_interval": {
							Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
			},
			"hello_interval_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"mtu_ignore_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mtu_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"neighbor_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"neighbor": {
							Type: schema.TypeString, Optional: true, Default: "::", Description: "OSPFv3 neighbor (Neighbor IPv6 address)",
						},
						"neig_inst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
						"neighbor_cost": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF cost for point-to-multipoint neighbor (metric)",
						},
						"neighbor_poll_interval": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF dead-router polling interval (Seconds)",
						},
						"neighbor_priority": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF priority of non-broadcast neighbor",
						},
					},
				},
			},
			"network_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"broadcast_type": {
							Type: schema.TypeString, Optional: true, Description: "'broadcast': Specify OSPF broadcast multi-access network; 'non-broadcast': Specify OSPF NBMA network; 'point-to-point': Specify OSPF point-to-point network; 'point-to-multipoint': Specify OSPF point-to-multipoint network;",
						},
						"p2mp_nbma": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify non-broadcast point-to-multipoint network",
						},
						"network_instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"priority_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"retransmit_interval_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"transmit_delay_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transmit_delay": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTunnelIpv6OspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6OspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6Ospf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpv6OspfRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTunnelIpv6OspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6OspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6Ospf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpv6OspfRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTunnelIpv6OspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6OspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6Ospf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTunnelIpv6OspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6OspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6Ospf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceTunnelIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfCostCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfCostCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfNeighborCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfNeighborCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfNeighborCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfNeighborCfg
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfNetworkList(d []interface{}) []edpt.InterfaceTunnelIpv6OspfNetworkList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfNetworkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfNetworkList
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfPriorityCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfPriorityCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTunnelIpv6Ospf(d *schema.ResourceData) edpt.InterfaceTunnelIpv6Ospf {
	var ret edpt.InterfaceTunnelIpv6Ospf
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.CostCfg = getSliceInterfaceTunnelIpv6OspfCostCfg(d.Get("cost_cfg").([]interface{}))
	ret.Inst.DeadIntervalCfg = getSliceInterfaceTunnelIpv6OspfDeadIntervalCfg(d.Get("dead_interval_cfg").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.HelloIntervalCfg = getSliceInterfaceTunnelIpv6OspfHelloIntervalCfg(d.Get("hello_interval_cfg").([]interface{}))
	ret.Inst.MtuIgnoreCfg = getSliceInterfaceTunnelIpv6OspfMtuIgnoreCfg(d.Get("mtu_ignore_cfg").([]interface{}))
	ret.Inst.NeighborCfg = getSliceInterfaceTunnelIpv6OspfNeighborCfg(d.Get("neighbor_cfg").([]interface{}))
	ret.Inst.NetworkList = getSliceInterfaceTunnelIpv6OspfNetworkList(d.Get("network_list").([]interface{}))
	ret.Inst.PriorityCfg = getSliceInterfaceTunnelIpv6OspfPriorityCfg(d.Get("priority_cfg").([]interface{}))
	ret.Inst.RetransmitIntervalCfg = getSliceInterfaceTunnelIpv6OspfRetransmitIntervalCfg(d.Get("retransmit_interval_cfg").([]interface{}))
	ret.Inst.TransmitDelayCfg = getSliceInterfaceTunnelIpv6OspfTransmitDelayCfg(d.Get("transmit_delay_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
