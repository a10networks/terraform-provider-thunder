package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnIpsec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn_ipsec`: IPsec settings\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnIpsecCreate,
		UpdateContext: resourceVpnIpsecUpdate,
		ReadContext:   resourceVpnIpsecRead,
		DeleteContext: resourceVpnIpsecDelete,

		Schema: map[string]*schema.Schema{
			"anti_replay_window": {
				Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': Disable Anti-Replay Window Check; '32': Window size of 32; '64': Window size of 64; '128': Window size of 128; '256': Window size of 256; '512': Window size of 512; '1024': Window size of 1024; '2048': Window size of 2048; '3072': Window size of 3072; '4096': Window size of 4096; '8192': Window size of 8192;",
			},
			"bind_tunnel": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Tunnel interface index",
						},
						"next_hop": {
							Type: schema.TypeString, Optional: true, Description: "IPsec Next Hop IP Address",
						},
						"next_hop_v6": {
							Type: schema.TypeString, Optional: true, Description: "IPsec Next Hop IPv6 Address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dh_group": {
				Type: schema.TypeString, Optional: true, Default: "0", Description: "'0': Diffie-Hellman group 0 (Default); '1': Diffie-Hellman group 1 - 768-bits; '2': Diffie-Hellman group 2 - 1024-bits; '5': Diffie-Hellman group 5 - 1536-bits; '14': Diffie-Hellman group 14 - 2048-bits; '15': Diffie-Hellman group 15 - 3072-bits; '16': Diffie-Hellman group 16 - 4096-bits; '18': Diffie-Hellman group 18 - 8192-bits; '19': Diffie-Hellman group 19 - 256-bit Elliptic Curve; '20': Diffie-Hellman group 20 - 384-bit Elliptic Curve;",
			},
			"dscp": {
				Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
			},
			"enc_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encryption": {
							Type: schema.TypeString, Optional: true, Description: "'des': Data Encryption Standard algorithm; '3des': Triple Data Encryption Standard algorithm; 'aes-128': Advanced Encryption Standard algorithm CBC Mode(key size: 128 bits); 'aes-192': Advanced Encryption Standard algorithm CBC Mode(key size: 192 bits); 'aes-256': Advanced Encryption Standard algorithm CBC Mode(key size: 256 bits); 'aes-gcm-128': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 128 bits, ICV size: 16 bytes); 'aes-gcm-192': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 192 bits, ICV size: 16 bytes); 'aes-gcm-256': Advanced Encryption Standard algorithm Galois/Counter Mode(key size: 256 bits, ICV size: 16 bytes); 'null': No encryption algorithm;",
						},
						"hash": {
							Type: schema.TypeString, Optional: true, Description: "'md5': MD5 Dessage-Digest Algorithm; 'sha1': Secure Hash Algorithm 1; 'sha256': Secure Hash Algorithm 256; 'sha384': Secure Hash Algorithm 384; 'sha512': Secure Hash Algorithm 512; 'null': No hash algorithm;",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Prioritizes (1-10) security protocol, least value has highest priority",
						},
						"gcm_priority": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Prioritizes (1-10) security protocol, least value has highest priority",
						},
					},
				},
			},
			"enforce_traffic_selector": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enforce Traffic Selector",
			},
			"ipsec_gateway": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ike_gateway": {
							Type: schema.TypeString, Optional: true, Description: "Gateway to use for IPsec SA",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"lifebytes": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPsec SA age in megabytes (0 indicates unlimited bytes)",
			},
			"lifetime": {
				Type: schema.TypeInt, Optional: true, Default: 28800, Description: "IPsec SA age in seconds",
			},
			"mode": {
				Type: schema.TypeString, Optional: true, Default: "tunnel", Description: "'tunnel': Encapsulating the packet in IPsec tunnel mode (Default);",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "IPsec name",
			},
			"proto": {
				Type: schema.TypeString, Optional: true, Default: "esp", Description: "'esp': Encapsulating security protocol (Default);",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets-encrypted': Encrypted Packets; 'packets-decrypted': Decrypted Packets; 'anti-replay-num': Anti-Replay Failure; 'rekey-num': Rekey Times; 'packets-err-inactive': Inactive Error; 'packets-err-encryption': Encryption Error; 'packets-err-pad-check': Pad Check Error; 'packets-err-pkt-sanity': Packets Sanity Error; 'packets-err-icv-check': ICV Check Error; 'packets-err-lifetime-lifebytes': Lifetime Lifebytes Error; 'bytes-encrypted': Encrypted Bytes; 'bytes-decrypted': Decrypted Bytes; 'prefrag-success': Pre-frag Success; 'prefrag-error': Pre-frag Error; 'cavium-bytes-encrypted': CAVIUM Encrypted Bytes; 'cavium-bytes-decrypted': CAVIUM Decrypted Bytes; 'cavium-packets-encrypted': CAVIUM Encrypted Packets; 'cavium-packets-decrypted': CAVIUM Decrypted Packets; 'qat-bytes-encrypted': QAT Encrypted Bytes; 'qat-bytes-decrypted': QAT Decrypted Bytes; 'qat-packets-encrypted': QAT Encrypted Packets; 'qat-packets-decrypted': QAT Decrypted Packets; 'tunnel-intf-down': Packet dropped: Tunnel Interface Down; 'pkt-fail-prep-to-send': Packet dropped: Failed in prepare to send; 'no-next-hop': Packet dropped: No next hop; 'invalid-tunnel-id': Packet dropped: Invalid tunnel ID; 'no-tunnel-found': Packet dropped: No tunnel found; 'pkt-fail-to-send': Packet dropped: Failed to send; 'frag-after-encap-frag-packets': Frag-after-encap Fragment Generated; 'frag-received': Fragment Received; 'sequence-num': Sequence Number; 'sequence-num-rollover': Sequence Number Rollover; 'packets-err-nh-check': Next Header Check Error;",
						},
					},
				},
			},
			"sequence_number_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not use incremental sequence number in the ESP header",
			},
			"traffic_selector": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local": {
										Type: schema.TypeString, Optional: true, Description: "Local Traffic Selector",
									},
									"local_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 Address Network Mask",
									},
									"local_port": {
										Type: schema.TypeInt, Optional: true, Description: "Port Number",
									},
									"remote_ipv4_assigned": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remote IP address assigned",
									},
									"remote_ip": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 Address",
									},
									"remote_netmask": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 Address Network Mask",
									},
									"remote_port": {
										Type: schema.TypeInt, Optional: true, Description: "Port Number",
									},
									"protocol": {
										Type: schema.TypeInt, Optional: true, Description: "IP Protocol Number (0-255)",
									},
								},
							},
						},
						"ipv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"localv6": {
										Type: schema.TypeString, Optional: true, Description: "Local Traffic Selector",
									},
									"local_portv6": {
										Type: schema.TypeInt, Optional: true, Description: "Port Number",
									},
									"remote_ipv6_assigned": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Remote IPv6 address assigned",
									},
									"remote_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 Address",
									},
									"remote_portv6": {
										Type: schema.TypeInt, Optional: true, Description: "Port Number",
									},
									"protocolv6": {
										Type: schema.TypeInt, Optional: true, Description: "IP Protocol Number (0-255)",
									},
								},
							},
						},
					},
				},
			},
			"up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiates SA negotiation to bring the IPsec connection up",
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
func resourceVpnIpsecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnIpsecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnIpsecRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnIpsecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnIpsecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnIpsecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnIpsec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVpnIpsecBindTunnel3612(d []interface{}) edpt.VpnIpsecBindTunnel3612 {

	count1 := len(d)
	var ret edpt.VpnIpsecBindTunnel3612
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tunnel = in["tunnel"].(int)
		ret.NextHop = in["next_hop"].(string)
		ret.NextHopV6 = in["next_hop_v6"].(string)
		//omit uuid
	}
	return ret
}

func getSliceVpnIpsecEncCfg(d []interface{}) []edpt.VpnIpsecEncCfg {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecEncCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecEncCfg
		oi.Encryption = in["encryption"].(string)
		oi.Hash = in["hash"].(string)
		oi.Priority = in["priority"].(int)
		oi.Gcm_priority = in["gcm_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIpsecIpsecGateway3613(d []interface{}) edpt.VpnIpsecIpsecGateway3613 {

	count1 := len(d)
	var ret edpt.VpnIpsecIpsecGateway3613
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IkeGateway = in["ike_gateway"].(string)
		//omit uuid
	}
	return ret
}

func getSliceVpnIpsecSamplingEnable(d []interface{}) []edpt.VpnIpsecSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.VpnIpsecSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VpnIpsecSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectVpnIpsecTrafficSelector(d []interface{}) edpt.VpnIpsecTrafficSelector {

	count1 := len(d)
	var ret edpt.VpnIpsecTrafficSelector
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4 = getObjectVpnIpsecTrafficSelectorIpv4(in["ipv4"].([]interface{}))
		ret.Ipv6 = getObjectVpnIpsecTrafficSelectorIpv6(in["ipv6"].([]interface{}))
	}
	return ret
}

func getObjectVpnIpsecTrafficSelectorIpv4(d []interface{}) edpt.VpnIpsecTrafficSelectorIpv4 {

	count1 := len(d)
	var ret edpt.VpnIpsecTrafficSelectorIpv4
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Local = in["local"].(string)
		ret.Local_netmask = in["local_netmask"].(string)
		ret.Local_port = in["local_port"].(int)
		ret.RemoteIpv4Assigned = in["remote_ipv4_assigned"].(int)
		ret.RemoteIp = in["remote_ip"].(string)
		ret.Remote_netmask = in["remote_netmask"].(string)
		ret.Remote_port = in["remote_port"].(int)
		ret.Protocol = in["protocol"].(int)
	}
	return ret
}

func getObjectVpnIpsecTrafficSelectorIpv6(d []interface{}) edpt.VpnIpsecTrafficSelectorIpv6 {

	count1 := len(d)
	var ret edpt.VpnIpsecTrafficSelectorIpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Localv6 = in["localv6"].(string)
		ret.Local_portv6 = in["local_portv6"].(int)
		ret.RemoteIpv6Assigned = in["remote_ipv6_assigned"].(int)
		ret.RemoteIpv6 = in["remote_ipv6"].(string)
		ret.Remote_portv6 = in["remote_portv6"].(int)
		ret.Protocolv6 = in["protocolv6"].(int)
	}
	return ret
}

func dataToEndpointVpnIpsec(d *schema.ResourceData) edpt.VpnIpsec {
	var ret edpt.VpnIpsec
	ret.Inst.AntiReplayWindow = d.Get("anti_replay_window").(string)
	ret.Inst.BindTunnel = getObjectVpnIpsecBindTunnel3612(d.Get("bind_tunnel").([]interface{}))
	ret.Inst.DhGroup = d.Get("dh_group").(string)
	ret.Inst.Dscp = d.Get("dscp").(string)
	ret.Inst.EncCfg = getSliceVpnIpsecEncCfg(d.Get("enc_cfg").([]interface{}))
	ret.Inst.EnforceTrafficSelector = d.Get("enforce_traffic_selector").(int)
	ret.Inst.IpsecGateway = getObjectVpnIpsecIpsecGateway3613(d.Get("ipsec_gateway").([]interface{}))
	ret.Inst.Lifebytes = d.Get("lifebytes").(int)
	ret.Inst.Lifetime = d.Get("lifetime").(int)
	ret.Inst.Mode = d.Get("mode").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Proto = d.Get("proto").(string)
	ret.Inst.SamplingEnable = getSliceVpnIpsecSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.SequenceNumberDisable = d.Get("sequence_number_disable").(int)
	ret.Inst.TrafficSelector = getObjectVpnIpsecTrafficSelector(d.Get("traffic_selector").([]interface{}))
	ret.Inst.Up = d.Get("up").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
