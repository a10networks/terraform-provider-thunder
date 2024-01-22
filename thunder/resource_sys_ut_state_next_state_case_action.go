package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCaseAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case_action`: Specify event parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseActionCreate,
		UpdateContext: resourceSysUtStateNextStateCaseActionUpdate,
		ReadContext:   resourceSysUtStateNextStateCaseActionRead,
		DeleteContext: resourceSysUtStateNextStateCaseActionDelete,

		Schema: map[string]*schema.Schema{
			"delay": {
				Type: schema.TypeInt, Optional: true, Description: "Delay in seconds",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "'send': Test event; 'expect': Expected result; 'wait': Introduce a delay;",
			},
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
			},
			"l1": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eth_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet_start": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
									},
									"ethernet_end": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
									},
								},
							},
						},
						"trunk_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk_start": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
									},
									"trunk_end": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
									},
								},
							},
						},
						"length": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet length",
						},
						"value": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at L2 header",
						},
						"auto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"l2": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethertype": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "L2 frame type",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Default: "ipv4", Description: "'arp': arp; 'ipv4': ipv4; 'ipv6': ipv6;",
						},
						"value": {
							Type: schema.TypeInt, Optional: true, Description: "ethertype number",
						},
						"vlan": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Vlan ID on the packet. 0 is untagged",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"mac_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_dst": {
										Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
									},
									"address_type": {
										Type: schema.TypeString, Optional: true, Description: "'broadcast': broadcast; 'multicast': multicast;",
									},
									"virtual_server": {
										Type: schema.TypeString, Optional: true, Description: "vip",
									},
									"nat_pool": {
										Type: schema.TypeString, Optional: true, Description: "Nat pool",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk number",
									},
									"value": {
										Type: schema.TypeString, Optional: true, Description: "Mac Address",
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
			"l3": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "L4 Protocol",
						},
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp;",
						},
						"value": {
							Type: schema.TypeInt, Optional: true, Description: "protocol number",
						},
						"checksum": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"ttl": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_dst": {
										Type: schema.TypeString, Required: true, Description: "'dest': dest; 'src': src;",
									},
									"ipv4_address": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"ipv6_address": {
										Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
									},
									"virtual_server": {
										Type: schema.TypeString, Optional: true, Description: "vip",
									},
									"nat_pool": {
										Type: schema.TypeString, Optional: true, Description: "Nat pool",
									},
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface",
									},
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual Ethernet interface",
									},
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk number",
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
			"tcp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Source port value",
						},
						"dest_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
						},
						"dest_port_value": {
							Type: schema.TypeInt, Optional: true, Description: "Dest port value",
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Nat pool port",
						},
						"seq_number": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"ack_seq_number": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"checksum": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"urgent": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"window": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"flags": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"syn": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
									},
									"ack": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
									},
									"fin": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
									},
									"rst": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
									},
									"psh": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
									},
									"ece": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
									},
									"urg": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
									},
									"cwr": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"options": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mss": {
										Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
									},
									"wscale": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sack_type": {
										Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
									},
									"time_stamp_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
									},
									"nop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
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
			"template": {
				Type: schema.TypeString, Optional: true, Description: "Packet template",
			},
			"udp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Source port value",
						},
						"dest_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
						},
						"dest_port_value": {
							Type: schema.TypeInt, Optional: true, Description: "Dest port value",
						},
						"nat_pool": {
							Type: schema.TypeString, Optional: true, Description: "Nat pool port",
						},
						"length": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at UDP header",
						},
						"checksum": {
							Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"case_number": {
				Type: schema.TypeString, Required: true, Description: "CaseNumber",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtStateNextStateCaseActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSysUtStateNextStateCaseActionL11542(d []interface{}) edpt.SysUtStateNextStateCaseActionL11542 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionL11542
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthList = getSliceSysUtStateNextStateCaseActionL1EthList1543(in["eth_list"].([]interface{}))
		ret.Trunk_list = getSliceSysUtStateNextStateCaseActionL1Trunk_list1544(in["trunk_list"].([]interface{}))
		ret.Length = in["length"].(int)
		ret.Value = in["value"].(int)
		ret.Auto = in["auto"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionL1EthList1543(d []interface{}) []edpt.SysUtStateNextStateCaseActionL1EthList1543 {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionL1EthList1543, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionL1EthList1543
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionL1Trunk_list1544(d []interface{}) []edpt.SysUtStateNextStateCaseActionL1Trunk_list1544 {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionL1Trunk_list1544, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionL1Trunk_list1544
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionL21545(d []interface{}) edpt.SysUtStateNextStateCaseActionL21545 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionL21545
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethertype = in["ethertype"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Value = in["value"].(int)
		ret.Vlan = in["vlan"].(int)
		//omit uuid
		ret.MacList = getSliceSysUtStateNextStateCaseActionL2MacList1546(in["mac_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionL2MacList1546(d []interface{}) []edpt.SysUtStateNextStateCaseActionL2MacList1546 {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionL2MacList1546, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionL2MacList1546
		oi.SrcDst = in["src_dst"].(string)
		oi.AddressType = in["address_type"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		oi.Value = in["value"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionL31547(d []interface{}) edpt.SysUtStateNextStateCaseActionL31547 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionL31547
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(int)
		ret.Type = in["type"].(string)
		ret.Value = in["value"].(int)
		ret.Checksum = in["checksum"].(string)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.IpList = getSliceSysUtStateNextStateCaseActionL3IpList1548(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionL3IpList1548(d []interface{}) []edpt.SysUtStateNextStateCaseActionL3IpList1548 {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionL3IpList1548, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionL3IpList1548
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.VirtualServer = in["virtual_server"].(string)
		oi.NatPool = in["nat_pool"].(string)
		oi.Ethernet = in["ethernet"].(int)
		oi.Ve = in["ve"].(int)
		oi.Trunk = in["trunk"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionTcp1549(d []interface{}) edpt.SysUtStateNextStateCaseActionTcp1549 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionTcp1549
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.SeqNumber = in["seq_number"].(string)
		ret.AckSeqNumber = in["ack_seq_number"].(string)
		ret.Checksum = in["checksum"].(string)
		ret.Urgent = in["urgent"].(string)
		ret.Window = in["window"].(string)
		//omit uuid
		ret.Flags = getObjectSysUtStateNextStateCaseActionTcpFlags1550(in["flags"].([]interface{}))
		ret.Options = getObjectSysUtStateNextStateCaseActionTcpOptions1551(in["options"].([]interface{}))
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionTcpFlags1550(d []interface{}) edpt.SysUtStateNextStateCaseActionTcpFlags1550 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionTcpFlags1550
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn = in["syn"].(int)
		ret.Ack = in["ack"].(int)
		ret.Fin = in["fin"].(int)
		ret.Rst = in["rst"].(int)
		ret.Psh = in["psh"].(int)
		ret.Ece = in["ece"].(int)
		ret.Urg = in["urg"].(int)
		ret.Cwr = in["cwr"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionTcpOptions1551(d []interface{}) edpt.SysUtStateNextStateCaseActionTcpOptions1551 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionTcpOptions1551
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mss = in["mss"].(int)
		ret.Wscale = in["wscale"].(int)
		ret.SackType = in["sack_type"].(string)
		ret.TimeStampEnable = in["time_stamp_enable"].(int)
		ret.Nop = in["nop"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtStateNextStateCaseActionUdp1552(d []interface{}) edpt.SysUtStateNextStateCaseActionUdp1552 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionUdp1552
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPort = in["src_port"].(int)
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.Length = in["length"].(int)
		ret.Checksum = in["checksum"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointSysUtStateNextStateCaseAction(d *schema.ResourceData) edpt.SysUtStateNextStateCaseAction {
	var ret edpt.SysUtStateNextStateCaseAction
	ret.Inst.Delay = d.Get("delay").(int)
	ret.Inst.Direction = d.Get("direction").(string)
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.L1 = getObjectSysUtStateNextStateCaseActionL11542(d.Get("l1").([]interface{}))
	ret.Inst.L2 = getObjectSysUtStateNextStateCaseActionL21545(d.Get("l2").([]interface{}))
	ret.Inst.L3 = getObjectSysUtStateNextStateCaseActionL31547(d.Get("l3").([]interface{}))
	ret.Inst.Tcp = getObjectSysUtStateNextStateCaseActionTcp1549(d.Get("tcp").([]interface{}))
	ret.Inst.Template = d.Get("template").(string)
	ret.Inst.Udp = getObjectSysUtStateNextStateCaseActionUdp1552(d.Get("udp").([]interface{}))
	//omit uuid
	ret.Inst.CaseNumber = d.Get("case_number").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
