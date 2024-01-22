package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlowspec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_flowspec`: Configure Flowspec\n\n__PLACEHOLDER__",
		CreateContext: resourceFlowspecCreate,
		UpdateContext: resourceFlowspecUpdate,
		ReadContext:   resourceFlowspecRead,
		DeleteContext: resourceFlowspecDelete,

		Schema: map[string]*schema.Schema{
			"dest_addr_type": {
				Type: schema.TypeString, Optional: true, Description: "'ip': IPv4 Address; 'ipv6': IPv6 Address;",
			},
			"dest_ip_host": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 host address",
			},
			"dest_ip_subnet": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Subnet address",
			},
			"dest_ipv6_host": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 host address",
			},
			"dest_ipv6_subnet": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Subnet address",
			},
			"destination_port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given destination port; 'gt': Match only packets with a greater port number; 'lt': Match only packets with a lower port number; 'range': match only packets in the range of port numbers;",
						},
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Specify the port number",
						},
						"port_num_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dscp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given DSCP; 'gt': Match only packets with a greater DSCP; 'lt': Match only packets with a lower DSCP; 'range': match only packets in the range of DSCPs;",
						},
						"dscp_val": {
							Type: schema.TypeInt, Required: true, Description: "Specify the DSCP value",
						},
						"dscp_val_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the DSCP value",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"filtering_action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"terminal_action": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Evaluation stops after this rule if not set",
						},
						"sample_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable traffic sampling and logging",
						},
						"traffic_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Type 0x8006 - Apply rate (in Bytes per second) for this class of traffic",
						},
						"traffic_marking": {
							Type: schema.TypeString, Optional: true, Description: "'dscp': IPv4 DSCP; 'ipv6-traffic-class': IPv6 Traffic Class;",
						},
						"dscp_val": {
							Type: schema.TypeInt, Optional: true, Description: "Set DSCP value",
						},
						"traffic_class": {
							Type: schema.TypeInt, Optional: true, Description: "Set IPv6 Traffic Class value",
						},
						"redirect": {
							Type: schema.TypeString, Optional: true, Description: "'next-hop-nlri': Type 0x0800 - IP encoded in MP_REACH_NLRI Next-hop network; 'next-hop': Type 0x0800 - Extended community Next-hop (Per v2 dated Feb 2015); 'vrf-route-target': Type 0x8008 - Redirect to VRF Route Target;",
						},
						"next_hop_nlri_type": {
							Type: schema.TypeString, Optional: true, Description: "'ip': Type 0x0800 - IPv4 Address; 'ipv6': Type 0x0800 - IPv6 Address;",
						},
						"ip_host_nlri": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 host address",
						},
						"copy_ip_host_nlri": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
						},
						"ipv6_host_nlri": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 host address",
						},
						"copy_ipv6_host_nlri": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
						},
						"next_hop_type": {
							Type: schema.TypeString, Optional: true, Description: "'ip': Type 0x0800 - IPv4 Address; 'ipv6': Type 0x0800 - IPv6 Address;",
						},
						"ip_host": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 host address",
						},
						"copy_ip_host": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
						},
						"ipv6_host": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 host address",
						},
						"copy_ipv6_host": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Copy bit",
						},
						"vrf_target_string": {
							Type: schema.TypeString, Optional: true, Description: "Type 0x8008(ASN-2:Index), 0x8208(ASN-4:Index) - Route Target AS",
						},
						"vrf_target_ip": {
							Type: schema.TypeString, Optional: true, Description: "'ip': Type 0x8108 - Redirect to route-target IP;",
						},
						"ip_host_rt": {
							Type: schema.TypeString, Optional: true, Description: "Type 0x8108 - Route Target IPv4",
						},
						"value_ip_host": {
							Type: schema.TypeInt, Optional: true, Description: "2-byte decimal value(local-administrator)",
						},
						"ecomm_custom_hex": {
							Type: schema.TypeString, Optional: true, Description: "Custom Extended Community in Hex",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"fragmentation_option_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"frag_attribute": {
							Type: schema.TypeString, Required: true, Description: "'is-fragment': Is fragmented packet; 'first-fragment': Is the first fragment packet; 'last-fragment': Is the last fragment; 'dont-fragment': Is DF bit set;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"icmp_code_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_code_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given ICMP Code; 'gt': Match only packets with a greater ICMP Code; 'lt': Match only packets with a lower ICMP Code; 'range': match only packets in the range of ICMP Codes;",
						},
						"code": {
							Type: schema.TypeInt, Required: true, Description: "Specify the ICMP Code",
						},
						"code_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP Code",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"icmp_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_type_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given ICMP Type; 'gt': Match only packets with a greater ICMP Type; 'lt': Match only packets with a lower ICMP Type; 'range': match only packets in the range of ICMP Types;",
						},
						"type": {
							Type: schema.TypeInt, Required: true, Description: "Specify the ICMP Type",
						},
						"type_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the ICMP Type",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Flowspec name",
			},
			"operational_mode": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'enabled': Enable the flowspec and send the prefix to BGP; 'disabled': Disable the flowspec and remove the prefix from BGP;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"packet_length_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_length_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given Packet Length; 'gt': Match only packets with a greater Packet Length; 'lt': Match only packets with a lower Packet Length; 'range': match only packets in the range of Packet Lengths;",
						},
						"length": {
							Type: schema.TypeInt, Required: true, Description: "Specify the Packet Length",
						},
						"length_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the Packet Length",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given port; 'gt': Match only packets with a greater port number; 'lt': Match only packets with a lower port number; 'range': match only packets in the range of port numbers;",
						},
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Specify the port number",
						},
						"port_num_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"protocol_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given protocol; 'gt': Match only packets with a greater protocol number; 'lt': Match only packets with a lower protocol number; 'range': match only packets in the range of protocol numbers;",
						},
						"proto_num": {
							Type: schema.TypeInt, Required: true, Description: "Specify the protocol number(6 for TCP and 17 for UDP)",
						},
						"proto_num_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the protocol number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"source_port_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_attribute": {
							Type: schema.TypeString, Required: true, Description: "'eq': Match only packets on a given source port; 'gt': Match only packets with a greater port number; 'lt': Match only packets with a lower port number; 'range': match only packets in the range of port numbers;",
						},
						"port_num": {
							Type: schema.TypeInt, Required: true, Description: "Specify the port number",
						},
						"port_num_end": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the port number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"src_addr_type": {
				Type: schema.TypeString, Optional: true, Description: "'ip': IPv4 Address; 'ipv6': IPv6 Address;",
			},
			"src_ip_host": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 host address",
			},
			"src_ip_subnet": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Subnet address",
			},
			"src_ipv6_host": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 host address",
			},
			"src_ipv6_subnet": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Subnet address",
			},
			"tcp_flags": {
				Type: schema.TypeString, Optional: true, Description: "'match-all': not = 0 match = 1; 'none-of': not = 1 match = 0; 'not-match': not = 1 match = 1; 'match-any': not = 0 match = 0;",
			},
			"tcp_flags_bitmask": {
				Type: schema.TypeString, Optional: true, Description: "Bitmask in Hex",
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
func resourceFlowspecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecRead(ctx, d, meta)
	}
	return diags
}

func resourceFlowspecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFlowspecRead(ctx, d, meta)
	}
	return diags
}
func resourceFlowspecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFlowspecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFlowspecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFlowspec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFlowspecDestinationPortList(d []interface{}) []edpt.FlowspecDestinationPortList {

	count1 := len(d)
	ret := make([]edpt.FlowspecDestinationPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecDestinationPortList
		oi.PortAttribute = in["port_attribute"].(string)
		oi.PortNum = in["port_num"].(int)
		oi.PortNumEnd = in["port_num_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFlowspecDscpList(d []interface{}) []edpt.FlowspecDscpList {

	count1 := len(d)
	ret := make([]edpt.FlowspecDscpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecDscpList
		oi.DscpAttribute = in["dscp_attribute"].(string)
		oi.DscpVal = in["dscp_val"].(int)
		oi.DscpValEnd = in["dscp_val_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFlowspecFilteringAction353(d []interface{}) edpt.FlowspecFilteringAction353 {

	count1 := len(d)
	var ret edpt.FlowspecFilteringAction353
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TerminalAction = in["terminal_action"].(int)
		ret.SampleLog = in["sample_log"].(int)
		ret.TrafficRate = in["traffic_rate"].(int)
		ret.TrafficMarking = in["traffic_marking"].(string)
		ret.DscpVal = in["dscp_val"].(int)
		ret.TrafficClass = in["traffic_class"].(int)
		ret.Redirect = in["redirect"].(string)
		ret.NextHopNlriType = in["next_hop_nlri_type"].(string)
		ret.IpHostNlri = in["ip_host_nlri"].(string)
		ret.CopyIpHostNlri = in["copy_ip_host_nlri"].(int)
		ret.Ipv6HostNlri = in["ipv6_host_nlri"].(string)
		ret.CopyIpv6HostNlri = in["copy_ipv6_host_nlri"].(int)
		ret.NextHopType = in["next_hop_type"].(string)
		ret.IpHost = in["ip_host"].(string)
		ret.CopyIpHost = in["copy_ip_host"].(int)
		ret.Ipv6Host = in["ipv6_host"].(string)
		ret.CopyIpv6Host = in["copy_ipv6_host"].(int)
		ret.VrfTargetString = in["vrf_target_string"].(string)
		ret.VrfTargetIp = in["vrf_target_ip"].(string)
		ret.IpHostRt = in["ip_host_rt"].(string)
		ret.ValueIpHost = in["value_ip_host"].(int)
		ret.EcommCustomHex = in["ecomm_custom_hex"].(string)
		//omit uuid
	}
	return ret
}

func getSliceFlowspecFragmentationOptionList(d []interface{}) []edpt.FlowspecFragmentationOptionList {

	count1 := len(d)
	ret := make([]edpt.FlowspecFragmentationOptionList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecFragmentationOptionList
		oi.FragAttribute = in["frag_attribute"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFlowspecIcmpCodeList(d []interface{}) []edpt.FlowspecIcmpCodeList {

	count1 := len(d)
	ret := make([]edpt.FlowspecIcmpCodeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecIcmpCodeList
		oi.IcmpCodeAttribute = in["icmp_code_attribute"].(string)
		oi.Code = in["code"].(int)
		oi.CodeEnd = in["code_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFlowspecIcmpTypeList(d []interface{}) []edpt.FlowspecIcmpTypeList {

	count1 := len(d)
	ret := make([]edpt.FlowspecIcmpTypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecIcmpTypeList
		oi.IcmpTypeAttribute = in["icmp_type_attribute"].(string)
		oi.Type = in["type"].(int)
		oi.TypeEnd = in["type_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectFlowspecOperationalMode354(d []interface{}) edpt.FlowspecOperationalMode354 {

	count1 := len(d)
	var ret edpt.FlowspecOperationalMode354
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
		//omit uuid
	}
	return ret
}

func getSliceFlowspecPacketLengthList(d []interface{}) []edpt.FlowspecPacketLengthList {

	count1 := len(d)
	ret := make([]edpt.FlowspecPacketLengthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecPacketLengthList
		oi.PacketLengthAttribute = in["packet_length_attribute"].(string)
		oi.Length = in["length"].(int)
		oi.LengthEnd = in["length_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFlowspecPortList(d []interface{}) []edpt.FlowspecPortList {

	count1 := len(d)
	ret := make([]edpt.FlowspecPortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecPortList
		oi.PortAttribute = in["port_attribute"].(string)
		oi.PortNum = in["port_num"].(int)
		oi.PortNumEnd = in["port_num_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFlowspecProtocolList(d []interface{}) []edpt.FlowspecProtocolList {

	count1 := len(d)
	ret := make([]edpt.FlowspecProtocolList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecProtocolList
		oi.ProtoAttribute = in["proto_attribute"].(string)
		oi.ProtoNum = in["proto_num"].(int)
		oi.ProtoNumEnd = in["proto_num_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFlowspecSourcePortList(d []interface{}) []edpt.FlowspecSourcePortList {

	count1 := len(d)
	ret := make([]edpt.FlowspecSourcePortList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FlowspecSourcePortList
		oi.PortAttribute = in["port_attribute"].(string)
		oi.PortNum = in["port_num"].(int)
		oi.PortNumEnd = in["port_num_end"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFlowspec(d *schema.ResourceData) edpt.Flowspec {
	var ret edpt.Flowspec
	ret.Inst.DestAddrType = d.Get("dest_addr_type").(string)
	ret.Inst.DestIpHost = d.Get("dest_ip_host").(string)
	ret.Inst.DestIpSubnet = d.Get("dest_ip_subnet").(string)
	ret.Inst.DestIpv6Host = d.Get("dest_ipv6_host").(string)
	ret.Inst.DestIpv6Subnet = d.Get("dest_ipv6_subnet").(string)
	ret.Inst.DestinationPortList = getSliceFlowspecDestinationPortList(d.Get("destination_port_list").([]interface{}))
	ret.Inst.DscpList = getSliceFlowspecDscpList(d.Get("dscp_list").([]interface{}))
	ret.Inst.FilteringAction = getObjectFlowspecFilteringAction353(d.Get("filtering_action").([]interface{}))
	ret.Inst.FragmentationOptionList = getSliceFlowspecFragmentationOptionList(d.Get("fragmentation_option_list").([]interface{}))
	ret.Inst.IcmpCodeList = getSliceFlowspecIcmpCodeList(d.Get("icmp_code_list").([]interface{}))
	ret.Inst.IcmpTypeList = getSliceFlowspecIcmpTypeList(d.Get("icmp_type_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OperationalMode = getObjectFlowspecOperationalMode354(d.Get("operational_mode").([]interface{}))
	ret.Inst.PacketLengthList = getSliceFlowspecPacketLengthList(d.Get("packet_length_list").([]interface{}))
	ret.Inst.PortList = getSliceFlowspecPortList(d.Get("port_list").([]interface{}))
	ret.Inst.ProtocolList = getSliceFlowspecProtocolList(d.Get("protocol_list").([]interface{}))
	ret.Inst.SourcePortList = getSliceFlowspecSourcePortList(d.Get("source_port_list").([]interface{}))
	ret.Inst.SrcAddrType = d.Get("src_addr_type").(string)
	ret.Inst.SrcIpHost = d.Get("src_ip_host").(string)
	ret.Inst.SrcIpSubnet = d.Get("src_ip_subnet").(string)
	ret.Inst.SrcIpv6Host = d.Get("src_ipv6_host").(string)
	ret.Inst.SrcIpv6Subnet = d.Get("src_ipv6_subnet").(string)
	ret.Inst.TcpFlags = d.Get("tcp_flags").(string)
	ret.Inst.TcpFlagsBitmask = d.Get("tcp_flags_bitmask").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
