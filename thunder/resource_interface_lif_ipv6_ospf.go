package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpv6Ospf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ipv6_ospf`: Open Shortest Path First for IPv6 (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpv6OspfCreate,
		UpdateContext: resourceInterfaceLifIpv6OspfUpdate,
		ReadContext:   resourceInterfaceLifIpv6OspfRead,
		DeleteContext: resourceInterfaceLifIpv6OspfDelete,

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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpv6OspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6OspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6Ospf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6OspfRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpv6OspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6OspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6Ospf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6OspfRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpv6OspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6OspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6Ospf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpv6OspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6OspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6Ospf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLifIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfCostCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfCostCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfDeadIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfDeadIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfHelloIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfHelloIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfMtuIgnoreCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfMtuIgnoreCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNeighborCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfNeighborCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNeighborCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNeighborCfg
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNetworkList(d []interface{}) []edpt.InterfaceLifIpv6OspfNetworkList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNetworkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNetworkList
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfPriorityCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfPriorityCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceLifIpv6OspfTransmitDelayCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfTransmitDelayCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLifIpv6Ospf(d *schema.ResourceData) edpt.InterfaceLifIpv6Ospf {
	var ret edpt.InterfaceLifIpv6Ospf
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.CostCfg = getSliceInterfaceLifIpv6OspfCostCfg(d.Get("cost_cfg").([]interface{}))
	ret.Inst.DeadIntervalCfg = getSliceInterfaceLifIpv6OspfDeadIntervalCfg(d.Get("dead_interval_cfg").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.HelloIntervalCfg = getSliceInterfaceLifIpv6OspfHelloIntervalCfg(d.Get("hello_interval_cfg").([]interface{}))
	ret.Inst.MtuIgnoreCfg = getSliceInterfaceLifIpv6OspfMtuIgnoreCfg(d.Get("mtu_ignore_cfg").([]interface{}))
	ret.Inst.NeighborCfg = getSliceInterfaceLifIpv6OspfNeighborCfg(d.Get("neighbor_cfg").([]interface{}))
	ret.Inst.NetworkList = getSliceInterfaceLifIpv6OspfNetworkList(d.Get("network_list").([]interface{}))
	ret.Inst.PriorityCfg = getSliceInterfaceLifIpv6OspfPriorityCfg(d.Get("priority_cfg").([]interface{}))
	ret.Inst.RetransmitIntervalCfg = getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg(d.Get("retransmit_interval_cfg").([]interface{}))
	ret.Inst.TransmitDelayCfg = getSliceInterfaceLifIpv6OspfTransmitDelayCfg(d.Get("transmit_delay_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
