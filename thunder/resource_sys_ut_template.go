package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template`: Packet config template\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateCreate,
		UpdateContext: resourceSysUtTemplateUpdate,
		ReadContext:   resourceSysUtTemplateRead,
		DeleteContext: resourceSysUtTemplateDelete,

		Schema: map[string]*schema.Schema{
			"ignore_validation": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l1": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate TX descriptor. This includes Tx port, Len & vlan",
						},
						"l2": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L2 header",
						},
						"l3": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L3 header",
						},
						"l4": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dont validate L4 header",
						},
						"all": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Skip validation",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
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
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
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
									"ipv4_start_address": {
										Type: schema.TypeString, Optional: true, Description: "IP address",
									},
									"ipv4_end_address": {
										Type: schema.TypeString, Optional: true, Description: "IP end address",
									},
									"ipv6_start_address": {
										Type: schema.TypeString, Optional: true, Description: "Ipv6 address",
									},
									"ipv6_end_address": {
										Type: schema.TypeString, Optional: true, Description: "Ipv6 end address",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "template name",
			},
			"tcp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port_range": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Source port value",
									},
									"src_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Src port end value",
									},
								},
							},
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
			"udp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port_range": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"src_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Source port value",
									},
									"src_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Src port end value",
									},
								},
							},
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSysUtTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSysUtTemplateIgnoreValidation1555(d []interface{}) edpt.SysUtTemplateIgnoreValidation1555 {

	count1 := len(d)
	var ret edpt.SysUtTemplateIgnoreValidation1555
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L1 = in["l1"].(int)
		ret.L2 = in["l2"].(int)
		ret.L3 = in["l3"].(int)
		ret.L4 = in["l4"].(int)
		ret.All = in["all"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtTemplateL11556(d []interface{}) edpt.SysUtTemplateL11556 {

	count1 := len(d)
	var ret edpt.SysUtTemplateL11556
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthList = getSliceSysUtTemplateL1EthList1557(in["eth_list"].([]interface{}))
		ret.Trunk_list = getSliceSysUtTemplateL1Trunk_list1558(in["trunk_list"].([]interface{}))
		ret.Drop = in["drop"].(int)
		ret.Length = in["length"].(int)
		ret.Value = in["value"].(int)
		ret.Auto = in["auto"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtTemplateL1EthList1557(d []interface{}) []edpt.SysUtTemplateL1EthList1557 {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL1EthList1557, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL1EthList1557
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtTemplateL1Trunk_list1558(d []interface{}) []edpt.SysUtTemplateL1Trunk_list1558 {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL1Trunk_list1558, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL1Trunk_list1558
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateL21559(d []interface{}) edpt.SysUtTemplateL21559 {

	count1 := len(d)
	var ret edpt.SysUtTemplateL21559
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ethertype = in["ethertype"].(int)
		ret.Protocol = in["protocol"].(string)
		ret.Value = in["value"].(int)
		ret.Vlan = in["vlan"].(int)
		//omit uuid
		ret.MacList = getSliceSysUtTemplateL2MacList1560(in["mac_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtTemplateL2MacList1560(d []interface{}) []edpt.SysUtTemplateL2MacList1560 {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL2MacList1560, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL2MacList1560
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

func getObjectSysUtTemplateL31561(d []interface{}) edpt.SysUtTemplateL31561 {

	count1 := len(d)
	var ret edpt.SysUtTemplateL31561
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Protocol = in["protocol"].(int)
		ret.Type = in["type"].(string)
		ret.Value = in["value"].(int)
		ret.Checksum = in["checksum"].(string)
		ret.Ttl = in["ttl"].(int)
		//omit uuid
		ret.IpList = getSliceSysUtTemplateL3IpList1562(in["ip_list"].([]interface{}))
	}
	return ret
}

func getSliceSysUtTemplateL3IpList1562(d []interface{}) []edpt.SysUtTemplateL3IpList1562 {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL3IpList1562, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL3IpList1562
		oi.SrcDst = in["src_dst"].(string)
		oi.Ipv4StartAddress = in["ipv4_start_address"].(string)
		oi.Ipv4EndAddress = in["ipv4_end_address"].(string)
		oi.Ipv6StartAddress = in["ipv6_start_address"].(string)
		oi.Ipv6EndAddress = in["ipv6_end_address"].(string)
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

func getObjectSysUtTemplateTcp1563(d []interface{}) edpt.SysUtTemplateTcp1563 {

	count1 := len(d)
	var ret edpt.SysUtTemplateTcp1563
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPortRange = getSliceSysUtTemplateTcpSrcPortRange1564(in["src_port_range"].([]interface{}))
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.SeqNumber = in["seq_number"].(string)
		ret.AckSeqNumber = in["ack_seq_number"].(string)
		ret.Checksum = in["checksum"].(string)
		ret.Urgent = in["urgent"].(string)
		ret.Window = in["window"].(string)
		//omit uuid
		ret.Flags = getObjectSysUtTemplateTcpFlags1565(in["flags"].([]interface{}))
		ret.Options = getObjectSysUtTemplateTcpOptions1566(in["options"].([]interface{}))
	}
	return ret
}

func getSliceSysUtTemplateTcpSrcPortRange1564(d []interface{}) []edpt.SysUtTemplateTcpSrcPortRange1564 {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateTcpSrcPortRange1564, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateTcpSrcPortRange1564
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSysUtTemplateTcpFlags1565(d []interface{}) edpt.SysUtTemplateTcpFlags1565 {

	count1 := len(d)
	var ret edpt.SysUtTemplateTcpFlags1565
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

func getObjectSysUtTemplateTcpOptions1566(d []interface{}) edpt.SysUtTemplateTcpOptions1566 {

	count1 := len(d)
	var ret edpt.SysUtTemplateTcpOptions1566
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

func getObjectSysUtTemplateUdp1567(d []interface{}) edpt.SysUtTemplateUdp1567 {

	count1 := len(d)
	var ret edpt.SysUtTemplateUdp1567
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcPortRange = getSliceSysUtTemplateUdpSrcPortRange1568(in["src_port_range"].([]interface{}))
		ret.DestPort = in["dest_port"].(int)
		ret.DestPortValue = in["dest_port_value"].(int)
		ret.NatPool = in["nat_pool"].(string)
		ret.Length = in["length"].(int)
		ret.Checksum = in["checksum"].(string)
		//omit uuid
	}
	return ret
}

func getSliceSysUtTemplateUdpSrcPortRange1568(d []interface{}) []edpt.SysUtTemplateUdpSrcPortRange1568 {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateUdpSrcPortRange1568, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateUdpSrcPortRange1568
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtTemplate(d *schema.ResourceData) edpt.SysUtTemplate {
	var ret edpt.SysUtTemplate
	ret.Inst.IgnoreValidation = getObjectSysUtTemplateIgnoreValidation1555(d.Get("ignore_validation").([]interface{}))
	ret.Inst.L1 = getObjectSysUtTemplateL11556(d.Get("l1").([]interface{}))
	ret.Inst.L2 = getObjectSysUtTemplateL21559(d.Get("l2").([]interface{}))
	ret.Inst.L3 = getObjectSysUtTemplateL31561(d.Get("l3").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Tcp = getObjectSysUtTemplateTcp1563(d.Get("tcp").([]interface{}))
	ret.Inst.Udp = getObjectSysUtTemplateUdp1567(d.Get("udp").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
