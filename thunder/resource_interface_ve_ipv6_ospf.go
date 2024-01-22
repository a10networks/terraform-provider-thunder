package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeIpv6Ospf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve_ipv6_ospf`: Open Shortest Path First for IPv6 (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceVeIpv6OspfCreate,
		UpdateContext: resourceInterfaceVeIpv6OspfUpdate,
		ReadContext:   resourceInterfaceVeIpv6OspfRead,
		DeleteContext: resourceInterfaceVeIpv6OspfDelete,

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
func resourceInterfaceVeIpv6OspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6OspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6Ospf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeIpv6OspfRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeIpv6OspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6OspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6Ospf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeIpv6OspfRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceVeIpv6OspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6OspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6Ospf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceVeIpv6OspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpv6OspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpv6Ospf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceVeIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfCostCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfCostCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfDeadIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfDeadIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfHelloIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfHelloIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfMtuIgnoreCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfMtuIgnoreCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfNeighborCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfNeighborCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfNeighborCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfNeighborCfg
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfNetworkList(d []interface{}) []edpt.InterfaceVeIpv6OspfNetworkList {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfNetworkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfNetworkList
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfPriorityCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfPriorityCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceVeIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceVeIpv6OspfTransmitDelayCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceVeIpv6OspfTransmitDelayCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceVeIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceVeIpv6Ospf(d *schema.ResourceData) edpt.InterfaceVeIpv6Ospf {
	var ret edpt.InterfaceVeIpv6Ospf
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.CostCfg = getSliceInterfaceVeIpv6OspfCostCfg(d.Get("cost_cfg").([]interface{}))
	ret.Inst.DeadIntervalCfg = getSliceInterfaceVeIpv6OspfDeadIntervalCfg(d.Get("dead_interval_cfg").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.HelloIntervalCfg = getSliceInterfaceVeIpv6OspfHelloIntervalCfg(d.Get("hello_interval_cfg").([]interface{}))
	ret.Inst.MtuIgnoreCfg = getSliceInterfaceVeIpv6OspfMtuIgnoreCfg(d.Get("mtu_ignore_cfg").([]interface{}))
	ret.Inst.NeighborCfg = getSliceInterfaceVeIpv6OspfNeighborCfg(d.Get("neighbor_cfg").([]interface{}))
	ret.Inst.NetworkList = getSliceInterfaceVeIpv6OspfNetworkList(d.Get("network_list").([]interface{}))
	ret.Inst.PriorityCfg = getSliceInterfaceVeIpv6OspfPriorityCfg(d.Get("priority_cfg").([]interface{}))
	ret.Inst.RetransmitIntervalCfg = getSliceInterfaceVeIpv6OspfRetransmitIntervalCfg(d.Get("retransmit_interval_cfg").([]interface{}))
	ret.Inst.TransmitDelayCfg = getSliceInterfaceVeIpv6OspfTransmitDelayCfg(d.Get("transmit_delay_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
