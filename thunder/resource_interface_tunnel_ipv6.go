package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_tunnel_ipv6`: Global IPv6 configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTunnelIpv6Create,
		UpdateContext: resourceInterfaceTunnelIpv6Update,
		ReadContext:   resourceInterfaceTunnelIpv6Read,
		DeleteContext: resourceInterfaceTunnelIpv6Delete,

		Schema: map[string]*schema.Schema{
			"address_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Set the IPv6 address of an interface",
						},
						"address_type": {
							Type: schema.TypeString, Optional: true, Description: "'anycast': Configure an IPv6 anycast address; 'link-local': Configure an IPv6 link local address;",
						},
					},
				},
			},
			"inside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as inside",
			},
			"ipv6_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
			},
			"ospf": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
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
					},
				},
			},
			"outside": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure interface as outside",
			},
			"router": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ripng": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rip": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIP Routing for IPv6",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ospf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"area_id_num": {
													Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
												},
												"area_id_addr": {
													Type: schema.TypeString, Optional: true, Description: "OSPF area ID in IP address format",
												},
												"tag": {
													Type: schema.TypeString, Optional: true, Description: "Set the OSPFv3 process tag",
												},
												"instance_id": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface instance ID",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
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
func resourceInterfaceTunnelIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTunnelIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTunnelIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTunnelIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceTunnelIpv6AddressCfg(d []interface{}) []edpt.InterfaceTunnelIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6AddressCfg
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.AddressType = in["address_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpv6Ospf953(d []interface{}) edpt.InterfaceTunnelIpv6Ospf953 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6Ospf953
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceTunnelIpv6OspfNetworkList954(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceTunnelIpv6OspfCostCfg955(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceTunnelIpv6OspfDeadIntervalCfg956(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceTunnelIpv6OspfHelloIntervalCfg957(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceTunnelIpv6OspfMtuIgnoreCfg958(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceTunnelIpv6OspfNeighborCfg959(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceTunnelIpv6OspfPriorityCfg960(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceTunnelIpv6OspfRetransmitIntervalCfg961(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceTunnelIpv6OspfTransmitDelayCfg962(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfNetworkList954(d []interface{}) []edpt.InterfaceTunnelIpv6OspfNetworkList954 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfNetworkList954, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfNetworkList954
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfCostCfg955(d []interface{}) []edpt.InterfaceTunnelIpv6OspfCostCfg955 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfCostCfg955, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfCostCfg955
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfDeadIntervalCfg956(d []interface{}) []edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg956 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg956, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfDeadIntervalCfg956
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfHelloIntervalCfg957(d []interface{}) []edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg957 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg957, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfHelloIntervalCfg957
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfMtuIgnoreCfg958(d []interface{}) []edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg958 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg958, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfMtuIgnoreCfg958
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfNeighborCfg959(d []interface{}) []edpt.InterfaceTunnelIpv6OspfNeighborCfg959 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfNeighborCfg959, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfNeighborCfg959
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfPriorityCfg960(d []interface{}) []edpt.InterfaceTunnelIpv6OspfPriorityCfg960 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfPriorityCfg960, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfPriorityCfg960
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfRetransmitIntervalCfg961(d []interface{}) []edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg961 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg961, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfRetransmitIntervalCfg961
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTunnelIpv6OspfTransmitDelayCfg962(d []interface{}) []edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg962 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg962, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6OspfTransmitDelayCfg962
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTunnelIpv6Router963(d []interface{}) edpt.InterfaceTunnelIpv6Router963 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6Router963
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceTunnelIpv6RouterRipng964(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceTunnelIpv6RouterOspf965(in["ospf"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpv6RouterRipng964(d []interface{}) edpt.InterfaceTunnelIpv6RouterRipng964 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6RouterRipng964
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceTunnelIpv6RouterOspf965(d []interface{}) edpt.InterfaceTunnelIpv6RouterOspf965 {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpv6RouterOspf965
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceTunnelIpv6RouterOspfAreaList966(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceTunnelIpv6RouterOspfAreaList966(d []interface{}) []edpt.InterfaceTunnelIpv6RouterOspfAreaList966 {

	count1 := len(d)
	ret := make([]edpt.InterfaceTunnelIpv6RouterOspfAreaList966, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTunnelIpv6RouterOspfAreaList966
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTunnelIpv6(d *schema.ResourceData) edpt.InterfaceTunnelIpv6 {
	var ret edpt.InterfaceTunnelIpv6
	ret.Inst.AddressCfg = getSliceInterfaceTunnelIpv6AddressCfg(d.Get("address_cfg").([]interface{}))
	ret.Inst.Inside = d.Get("inside").(int)
	ret.Inst.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Inst.Ospf = getObjectInterfaceTunnelIpv6Ospf953(d.Get("ospf").([]interface{}))
	ret.Inst.Outside = d.Get("outside").(int)
	ret.Inst.Router = getObjectInterfaceTunnelIpv6Router963(d.Get("router").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
