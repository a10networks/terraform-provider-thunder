package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ipv6`: Global IPv6 configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpv6Create,
		UpdateContext: resourceInterfaceLifIpv6Update,
		ReadContext:   resourceInterfaceLifIpv6Read,
		DeleteContext: resourceInterfaceLifIpv6Delete,

		Schema: map[string]*schema.Schema{
			"address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Set the IPv6 address of an interface",
						},
						"anycast": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 anycast address",
						},
						"link_local": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure an IPv6 link local address",
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
						"isis": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag": {
										Type: schema.TypeString, Optional: true, Description: "ISO routing area tag",
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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLifIpv6AddressList(d []interface{}) []edpt.InterfaceLifIpv6AddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6AddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Anycast = in["anycast"].(int)
		oi.LinkLocal = in["link_local"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6Ospf659(d []interface{}) edpt.InterfaceLifIpv6Ospf659 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6Ospf659
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkList = getSliceInterfaceLifIpv6OspfNetworkList660(in["network_list"].([]interface{}))
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceLifIpv6OspfCostCfg661(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceLifIpv6OspfDeadIntervalCfg662(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceLifIpv6OspfHelloIntervalCfg663(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceLifIpv6OspfMtuIgnoreCfg664(in["mtu_ignore_cfg"].([]interface{}))
		ret.NeighborCfg = getSliceInterfaceLifIpv6OspfNeighborCfg665(in["neighbor_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceLifIpv6OspfPriorityCfg666(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg667(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceLifIpv6OspfTransmitDelayCfg668(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNetworkList660(d []interface{}) []edpt.InterfaceLifIpv6OspfNetworkList660 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNetworkList660, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNetworkList660
		oi.BroadcastType = in["broadcast_type"].(string)
		oi.P2mpNbma = in["p2mp_nbma"].(int)
		oi.NetworkInstanceId = in["network_instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfCostCfg661(d []interface{}) []edpt.InterfaceLifIpv6OspfCostCfg661 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfCostCfg661, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfCostCfg661
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfDeadIntervalCfg662(d []interface{}) []edpt.InterfaceLifIpv6OspfDeadIntervalCfg662 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfDeadIntervalCfg662, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfDeadIntervalCfg662
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfHelloIntervalCfg663(d []interface{}) []edpt.InterfaceLifIpv6OspfHelloIntervalCfg663 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfHelloIntervalCfg663, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfHelloIntervalCfg663
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfMtuIgnoreCfg664(d []interface{}) []edpt.InterfaceLifIpv6OspfMtuIgnoreCfg664 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfMtuIgnoreCfg664, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfMtuIgnoreCfg664
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfNeighborCfg665(d []interface{}) []edpt.InterfaceLifIpv6OspfNeighborCfg665 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfNeighborCfg665, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfNeighborCfg665
		oi.Neighbor = in["neighbor"].(string)
		oi.NeigInst = in["neig_inst"].(int)
		oi.NeighborCost = in["neighbor_cost"].(int)
		oi.NeighborPollInterval = in["neighbor_poll_interval"].(int)
		oi.NeighborPriority = in["neighbor_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfPriorityCfg666(d []interface{}) []edpt.InterfaceLifIpv6OspfPriorityCfg666 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfPriorityCfg666, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfPriorityCfg666
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfRetransmitIntervalCfg667(d []interface{}) []edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg667 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg667, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfRetransmitIntervalCfg667
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIpv6OspfTransmitDelayCfg668(d []interface{}) []edpt.InterfaceLifIpv6OspfTransmitDelayCfg668 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6OspfTransmitDelayCfg668, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6OspfTransmitDelayCfg668
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6Router669(d []interface{}) edpt.InterfaceLifIpv6Router669 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6Router669
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceLifIpv6RouterRipng670(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceLifIpv6RouterOspf671(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceLifIpv6RouterIsis673(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterRipng670(d []interface{}) edpt.InterfaceLifIpv6RouterRipng670 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6RouterRipng670
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterOspf671(d []interface{}) edpt.InterfaceLifIpv6RouterOspf671 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6RouterOspf671
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceLifIpv6RouterOspfAreaList672(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLifIpv6RouterOspfAreaList672(d []interface{}) []edpt.InterfaceLifIpv6RouterOspfAreaList672 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6RouterOspfAreaList672, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6RouterOspfAreaList672
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIpv6RouterIsis673(d []interface{}) edpt.InterfaceLifIpv6RouterIsis673 {

	count1 := len(d)
	var ret edpt.InterfaceLifIpv6RouterIsis673
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceLifIpv6(d *schema.ResourceData) edpt.InterfaceLifIpv6 {
	var ret edpt.InterfaceLifIpv6
	ret.Ipv6.AddressList = getSliceInterfaceLifIpv6AddressList(
		d.Get("address_list").([]interface{}),
	)
	ret.Ipv6.Inside = d.Get("inside").(int)
	ret.Ipv6.Outside = d.Get("outside").(int)
	ret.Ipv6.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Ipv6.Ospf = getObjectInterfaceLifIpv6Ospf659(d.Get("ospf").([]interface{}))
	ret.Ipv6.Router = getObjectInterfaceLifIpv6Router669(d.Get("router").([]interface{}))

	ret.Ipv6.Ifname = d.Get("ifname").(string)

	return ret
}
