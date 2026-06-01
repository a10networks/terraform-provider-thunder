package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_ipv6`: Global IPv6 configuration subcommands\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIpv6Create,
		UpdateContext: resourceInterfaceLoopbackIpv6Update,
		ReadContext:   resourceInterfaceLoopbackIpv6Read,
		DeleteContext: resourceInterfaceLoopbackIpv6Delete,

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
			"ipv6_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable IPv6 processing",
			},
			"ospf": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
			"rip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"split_horizon_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"state": {
										Type: schema.TypeString, Optional: true, Default: "poisoned", Description: "'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;",
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
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceLoopbackIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLoopbackIpv6AddressList(d []interface{}) []edpt.InterfaceLoopbackIpv6AddressList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6AddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6AddressList
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.Anycast = in["anycast"].(int)
		oi.LinkLocal = in["link_local"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Ospf750(d []interface{}) edpt.InterfaceLoopbackIpv6Ospf750 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6Ospf750
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
		ret.CostCfg = getSliceInterfaceLoopbackIpv6OspfCostCfg751(in["cost_cfg"].([]interface{}))
		ret.DeadIntervalCfg = getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg752(in["dead_interval_cfg"].([]interface{}))
		ret.HelloIntervalCfg = getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg753(in["hello_interval_cfg"].([]interface{}))
		ret.MtuIgnoreCfg = getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg754(in["mtu_ignore_cfg"].([]interface{}))
		ret.PriorityCfg = getSliceInterfaceLoopbackIpv6OspfPriorityCfg755(in["priority_cfg"].([]interface{}))
		ret.RetransmitIntervalCfg = getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg756(in["retransmit_interval_cfg"].([]interface{}))
		ret.TransmitDelayCfg = getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg757(in["transmit_delay_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfCostCfg751(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfCostCfg751 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfCostCfg751, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfCostCfg751
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg752(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg752 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg752, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg752
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg753(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg753 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg753, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg753
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg754(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg754 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg754, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg754
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfPriorityCfg755(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfPriorityCfg755 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfPriorityCfg755, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfPriorityCfg755
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg756(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg756 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg756, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg756
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg757(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg757 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg757, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg757
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Rip758(d []interface{}) edpt.InterfaceLoopbackIpv6Rip758 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6Rip758
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SplitHorizonCfg = getObjectInterfaceLoopbackIpv6RipSplitHorizonCfg759(in["split_horizon_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RipSplitHorizonCfg759(d []interface{}) edpt.InterfaceLoopbackIpv6RipSplitHorizonCfg759 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RipSplitHorizonCfg759
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6Router760(d []interface{}) edpt.InterfaceLoopbackIpv6Router760 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6Router760
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ripng = getObjectInterfaceLoopbackIpv6RouterRipng761(in["ripng"].([]interface{}))
		ret.Ospf = getObjectInterfaceLoopbackIpv6RouterOspf762(in["ospf"].([]interface{}))
		ret.Isis = getObjectInterfaceLoopbackIpv6RouterIsis764(in["isis"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterRipng761(d []interface{}) edpt.InterfaceLoopbackIpv6RouterRipng761 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RouterRipng761
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rip = in["rip"].(int)
		//omit uuid
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterOspf762(d []interface{}) edpt.InterfaceLoopbackIpv6RouterOspf762 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RouterOspf762
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaList = getSliceInterfaceLoopbackIpv6RouterOspfAreaList763(in["area_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6RouterOspfAreaList763(d []interface{}) []edpt.InterfaceLoopbackIpv6RouterOspfAreaList763 {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6RouterOspfAreaList763, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6RouterOspfAreaList763
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIpv6RouterIsis764(d []interface{}) edpt.InterfaceLoopbackIpv6RouterIsis764 {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpv6RouterIsis764
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tag = in["tag"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointInterfaceLoopbackIpv6(d *schema.ResourceData) edpt.InterfaceLoopbackIpv6 {
	var ret edpt.InterfaceLoopbackIpv6
	ret.Inst.AddressList = getSliceInterfaceLoopbackIpv6AddressList(d.Get("address_list").([]interface{}))
	ret.Inst.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Inst.Ospf = getObjectInterfaceLoopbackIpv6Ospf750(d.Get("ospf").([]interface{}))
	ret.Inst.Rip = getObjectInterfaceLoopbackIpv6Rip758(d.Get("rip").([]interface{}))
	ret.Inst.Router = getObjectInterfaceLoopbackIpv6Router760(d.Get("router").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
