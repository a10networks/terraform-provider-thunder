package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOverlayTunnelVtep() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_overlay_tunnel_vtep`: Virtual Tunnel end point Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceOverlayTunnelVtepCreate,
		UpdateContext: resourceOverlayTunnelVtepUpdate,
		ReadContext:   resourceOverlayTunnelVtepRead,
		DeleteContext: resourceOverlayTunnelVtepDelete,

		Schema: map[string]*schema.Schema{
			"dest_port": {
				Type: schema.TypeInt, Optional: true, Description: "Layer-4 Destination Port (Port Number)",
			},
			"encap": {
				Type: schema.TypeString, Optional: true, Description: "'ip-encap': Tunnel encapsulation type is IP; 'gre': Tunnel encapsulation type is GRE; 'nvgre': Tunnel Encapsulation Type is NVGRE; 'vxlan': Tunnel Encapsulation Type is VXLAN;",
			},
			"host_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_addr": {
							Type: schema.TypeString, Required: true, Description: "IPv4 address of the overlay host",
						},
						"overlay_mac_addr": {
							Type: schema.TypeString, Required: true, Description: "MAC Address of the overlay host",
						},
						"vni": {
							Type: schema.TypeInt, Required: true, Description: "Configure the segment id ( VNI of the remote host)",
						},
						"remote_vtep": {
							Type: schema.TypeString, Required: true, Description: "Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"id1": {
				Type: schema.TypeInt, Required: true, Description: "VTEP Identifier",
			},
			"local_ip_address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type: schema.TypeString, Optional: true, Description: "Source Tunnel End Point IPv4 address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"vni_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"segment": {
										Type: schema.TypeInt, Required: true, Description: "Id of the segment that is being extended",
									},
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
									},
									"gateway": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "This is a Gateway segment id",
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
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
			"local_ipv6_address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "Source Tunnel End Point IPv6 address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"remote_ip_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type: schema.TypeString, Required: true, Description: "IP Address of the remote VTEP",
						},
						"encap": {
							Type: schema.TypeString, Optional: true, Description: "'nvgre': Tunnel Encapsulation Type is NVGRE; 'vxlan': Tunnel Encapsulation Type is VXLAN;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"use_lif": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"gre_keepalive": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retry_time": {
										Type: schema.TypeInt, Optional: true, Default: 10, Description: "Keepalive retry interval in seconds",
									},
									"retry_count": {
										Type: schema.TypeInt, Optional: true, Description: "Keepalive multiplier",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"use_gre_key": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gre_key": {
										Type: schema.TypeInt, Optional: true, Description: "key",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"vni_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"segment": {
										Type: schema.TypeInt, Required: true, Description: "VNI configured for the remote VTEP",
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
			"remote_ipv6_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": {
							Type: schema.TypeString, Required: true, Description: "IPv6 Address of the remote VTEP",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"use_lif": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition": {
										Type: schema.TypeString, Optional: true, Description: "Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)",
									},
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (logical interface name)",
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
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'cfg_err_count': Config errors; 'flooded_pkt_count': Flooded packet count; 'encap_unresolved_count': Encap unresolved failures; 'unknown_encap_rx_pkt': Encap miss rx pkts; 'unknown_encap_tx_pkt': Encap miss tx pkts; 'arp_req_sent': Arp request sent; 'vtep_host_learned': Hosts learned; 'vtep_host_learn_error': Host learn error; 'invalid_lif_rx': Invalid Lif pkts in; 'invalid_lif_tx': Invalid Lif pkts out; 'unknown_vtep_tx': Vtep unknown tx; 'unknown_vtep_rx': Vtep Unkown rx; 'unhandled_pkt_rx': Unhandled packets in; 'unhandled_pkt_tx': Unhandled packets out; 'total_pkts_rx': Total packets out; 'total_bytes_rx': Total packet bytes in; 'unicast_pkt_rx': Total unicast packets in; 'bcast_pkt_rx': Total broadcast packets in; 'mcast_pkt_rx': Total multicast packets in; 'dropped_pkt_rx': Dropped received packets; 'encap_miss_pkts_rx': Encap missed in received packets; 'bad_chksum_pks_rx': Bad checksum in received packets; 'requeue_pkts_in': Requeued packets in; 'pkts_out': Packets out; 'total_bytes_tx': Packet bytes out; 'unicast_pkt_tx': Unicast packets out; 'bcast_pkt_tx': Broadcast packets out; 'mcast_pkt_tx': Multicast packets out; 'dropped_pkts_tx': Dropped packets out; 'large_pkts_rx': Too large packets in; 'dot1q_pkts_rx': Dot1q packets in; 'frag_pkts_tx': Frag packets out; 'reassembled_pkts_rx': Reassembled packets in; 'bad_inner_ipv4_len_rx': bad inner ipv4 packet len; 'bad_inner_ipv6_len_rx': Bad inner ipv6 packet len; 'frag_drop_pkts_tx': Frag dropped packets out; 'lif_un_init_rx': Lif uninitialized packets in;",
						},
					},
				},
			},
			"src_port_range": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_port": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Minimum Port Number",
						},
						"max_port": {
							Type: schema.TypeInt, Optional: true, Default: 65535, Description: "Maximum Port Number",
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
func resourceOverlayTunnelVtepCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtep(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRead(ctx, d, meta)
	}
	return diags
}

func resourceOverlayTunnelVtepUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtep(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceOverlayTunnelVtepRead(ctx, d, meta)
	}
	return diags
}
func resourceOverlayTunnelVtepDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtep(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceOverlayTunnelVtepRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceOverlayTunnelVtepRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointOverlayTunnelVtep(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceOverlayTunnelVtepHostList(d []interface{}) []edpt.OverlayTunnelVtepHostList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepHostList
		oi.IpAddr = in["ip_addr"].(string)
		oi.OverlayMacAddr = in["overlay_mac_addr"].(string)
		oi.Vni = in["vni"].(int)
		oi.RemoteVtep = in["remote_vtep"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectOverlayTunnelVtepLocalIpAddress1085(d []interface{}) edpt.OverlayTunnelVtepLocalIpAddress1085 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepLocalIpAddress1085
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpAddress = in["ip_address"].(string)
		//omit uuid
		ret.VniList = getSliceOverlayTunnelVtepLocalIpAddressVniList1086(in["vni_list"].([]interface{}))
	}
	return ret
}

func getSliceOverlayTunnelVtepLocalIpAddressVniList1086(d []interface{}) []edpt.OverlayTunnelVtepLocalIpAddressVniList1086 {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepLocalIpAddressVniList1086, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepLocalIpAddressVniList1086
		oi.Segment = in["segment"].(int)
		oi.Partition = in["partition"].(string)
		oi.Gateway = in["gateway"].(int)
		oi.Lif = in["lif"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectOverlayTunnelVtepLocalIpv6Address1087(d []interface{}) edpt.OverlayTunnelVtepLocalIpv6Address1087 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepLocalIpv6Address1087
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6Address = in["ipv6_address"].(string)
		//omit uuid
	}
	return ret
}

func getSliceOverlayTunnelVtepRemoteIpAddressList(d []interface{}) []edpt.OverlayTunnelVtepRemoteIpAddressList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepRemoteIpAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepRemoteIpAddressList
		oi.IpAddress = in["ip_address"].(string)
		oi.Encap = in["encap"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.UseLif = getObjectOverlayTunnelVtepRemoteIpAddressListUseLif(in["use_lif"].([]interface{}))
		oi.GreKeepalive = getObjectOverlayTunnelVtepRemoteIpAddressListGreKeepalive(in["gre_keepalive"].([]interface{}))
		oi.UseGreKey = getObjectOverlayTunnelVtepRemoteIpAddressListUseGreKey(in["use_gre_key"].([]interface{}))
		oi.VniList = getSliceOverlayTunnelVtepRemoteIpAddressListVniList(in["vni_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpAddressListUseLif(d []interface{}) edpt.OverlayTunnelVtepRemoteIpAddressListUseLif {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpAddressListUseLif
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(string)
		ret.Lif = in["lif"].(string)
		//omit uuid
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpAddressListGreKeepalive(d []interface{}) edpt.OverlayTunnelVtepRemoteIpAddressListGreKeepalive {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpAddressListGreKeepalive
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetryTime = in["retry_time"].(int)
		ret.RetryCount = in["retry_count"].(int)
		//omit uuid
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpAddressListUseGreKey(d []interface{}) edpt.OverlayTunnelVtepRemoteIpAddressListUseGreKey {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpAddressListUseGreKey
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreKey = in["gre_key"].(int)
		//omit uuid
	}
	return ret
}

func getSliceOverlayTunnelVtepRemoteIpAddressListVniList(d []interface{}) []edpt.OverlayTunnelVtepRemoteIpAddressListVniList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepRemoteIpAddressListVniList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepRemoteIpAddressListVniList
		oi.Segment = in["segment"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceOverlayTunnelVtepRemoteIpv6AddressList(d []interface{}) []edpt.OverlayTunnelVtepRemoteIpv6AddressList {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepRemoteIpv6AddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepRemoteIpv6AddressList
		oi.Ipv6Address = in["ipv6_address"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.UseLif = getObjectOverlayTunnelVtepRemoteIpv6AddressListUseLif(in["use_lif"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectOverlayTunnelVtepRemoteIpv6AddressListUseLif(d []interface{}) edpt.OverlayTunnelVtepRemoteIpv6AddressListUseLif {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepRemoteIpv6AddressListUseLif
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Partition = in["partition"].(string)
		ret.Lif = in["lif"].(string)
		//omit uuid
	}
	return ret
}

func getSliceOverlayTunnelVtepSamplingEnable(d []interface{}) []edpt.OverlayTunnelVtepSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.OverlayTunnelVtepSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.OverlayTunnelVtepSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectOverlayTunnelVtepSrcPortRange1088(d []interface{}) edpt.OverlayTunnelVtepSrcPortRange1088 {

	count1 := len(d)
	var ret edpt.OverlayTunnelVtepSrcPortRange1088
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinPort = in["min_port"].(int)
		ret.MaxPort = in["max_port"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointOverlayTunnelVtep(d *schema.ResourceData) edpt.OverlayTunnelVtep {
	var ret edpt.OverlayTunnelVtep
	ret.Inst.DestPort = d.Get("dest_port").(int)
	ret.Inst.Encap = d.Get("encap").(string)
	ret.Inst.HostList = getSliceOverlayTunnelVtepHostList(d.Get("host_list").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.LocalIpAddress = getObjectOverlayTunnelVtepLocalIpAddress1085(d.Get("local_ip_address").([]interface{}))
	ret.Inst.LocalIpv6Address = getObjectOverlayTunnelVtepLocalIpv6Address1087(d.Get("local_ipv6_address").([]interface{}))
	ret.Inst.RemoteIpAddressList = getSliceOverlayTunnelVtepRemoteIpAddressList(d.Get("remote_ip_address_list").([]interface{}))
	ret.Inst.RemoteIpv6AddressList = getSliceOverlayTunnelVtepRemoteIpv6AddressList(d.Get("remote_ipv6_address_list").([]interface{}))
	ret.Inst.SamplingEnable = getSliceOverlayTunnelVtepSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SrcPortRange = getObjectOverlayTunnelVtepSrcPortRange1088(d.Get("src_port_range").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
