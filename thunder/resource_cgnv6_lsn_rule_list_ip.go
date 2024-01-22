package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListIp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_rule_list_ip`: Configure a Specific Rule-Set (IP Network Address)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnRuleListIpCreate,
		UpdateContext: resourceCgnv6LsnRuleListIpUpdate,
		ReadContext:   resourceCgnv6LsnRuleListIpRead,
		DeleteContext: resourceCgnv6LsnRuleListIpDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "Configure a Specific Rule-Set (IP Network Address)",
			},
			"rule_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': TCP L4 Protocol; 'udp': UDP L4 Protocol; 'icmp': ICMP L4 Protocol; 'others': Other L4 Protocol; 'dscp': Match DSCP Value; 'default': Default Action;",
						},
						"tcp_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_port": {
										Type: schema.TypeInt, Optional: true, Description: "Single Port or Start of Port Range (inclusive), Port 0 is Match Any Port",
									},
									"end_port": {
										Type: schema.TypeInt, Optional: true, Description: "End of Port Range (inclusive)",
									},
									"action_cfg": {
										Type: schema.TypeString, Optional: true, Description: "'action': LSN Rule-List Action; 'no-action': Exclude LSN Rule-List Action;",
									},
									"action_type": {
										Type: schema.TypeString, Optional: true, Description: "'dnat': Apply Dest NAT; 'drop': Drop the Packets; 'one-to-one-snat': Apply one-to-one source NAT for the packets; 'pass-through': Pass the Packets Through; 'snat': Redirect the Packets to a Different Source NAT Pool; 'set-dscp': To set dscp value for the packets; 'template': Template; 'idle-timeout': Configure idle timeout;",
									},
									"ipv4_list": {
										Type: schema.TypeString, Optional: true, Description: "IP-List (IP-List Name)",
									},
									"port_list": {
										Type: schema.TypeString, Optional: true, Description: "Port-List Name",
									},
									"no_snat": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable source NAT",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid",
									},
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "DNAT Using IP of a Domain (Domain Name)",
									},
									"pool": {
										Type: schema.TypeString, Optional: true, Description: "NAT Pool (NAT Pool or Pool Group)",
									},
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "The pool is a shared pool",
									},
									"http_alg": {
										Type: schema.TypeString, Optional: true, Description: "HTTP-ALG Template (Template Name)",
									},
									"timeout_val": {
										Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds",
									},
									"fast": {
										Type: schema.TypeString, Optional: true, Description: "'fast': Fast Aging;",
									},
									"dscp_direction": {
										Type: schema.TypeString, Optional: true, Description: "'inbound': To set dscp value for inbound packets; 'outbound': To set dscp value for outbound packets;",
									},
									"dscp_value": {
										Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
									},
								},
							},
						},
						"udp_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_port": {
										Type: schema.TypeInt, Optional: true, Description: "Single Port or Start of Port Range (inclusive), Port 0 is Match Any Port",
									},
									"end_port": {
										Type: schema.TypeInt, Optional: true, Description: "End of Port Range (inclusive)",
									},
									"action_cfg": {
										Type: schema.TypeString, Optional: true, Description: "'action': LSN Rule-List Action; 'no-action': Exclude LSN Rule-List Action;",
									},
									"action_type": {
										Type: schema.TypeString, Optional: true, Description: "'dnat': Apply Dest NAT; 'drop': Drop the Packets; 'one-to-one-snat': Apply one-to-one source NAT for the packets; 'pass-through': Pass the Packets Through; 'snat': Redirect the Packets to a Different Source NAT Pool; 'set-dscp': To set dscp value for the packets; 'idle-timeout': Configure idle timeout;",
									},
									"ipv4_list": {
										Type: schema.TypeString, Optional: true, Description: "IP-List (IP-List Name)",
									},
									"port_list": {
										Type: schema.TypeString, Optional: true, Description: "Port-List Name",
									},
									"no_snat": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable source NAT",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid",
									},
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "DNAT Using IP of a Domain (Domain Name)",
									},
									"pool": {
										Type: schema.TypeString, Optional: true, Description: "NAT Pool (NAT Pool or Pool Group)",
									},
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "The pool is a shared pool",
									},
									"timeout_val": {
										Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds",
									},
									"fast": {
										Type: schema.TypeString, Optional: true, Description: "'fast': Fast Aging;",
									},
									"dscp_direction": {
										Type: schema.TypeString, Optional: true, Description: "'inbound': To set dscp value for inbound packets; 'outbound': To set dscp value for outbound packets;",
									},
									"dscp_value": {
										Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
									},
								},
							},
						},
						"icmp_others_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_cfg": {
										Type: schema.TypeString, Optional: true, Description: "'action': LSN Rule-List Action; 'no-action': Exclude LSN Rule-List Action;",
									},
									"action_type": {
										Type: schema.TypeString, Optional: true, Description: "'dnat': Apply Dest NAT; 'drop': Drop the Packets; 'one-to-one-snat': Apply one-to-one source NAT for the packets; 'pass-through': Pass the Packets Through; 'snat': Redirect the Packets to a Different Source NAT Pool; 'set-dscp': To set dscp value for the packets;",
									},
									"ipv4_list": {
										Type: schema.TypeString, Optional: true, Description: "IP-List (IP-List Name)",
									},
									"no_snat": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable source NAT",
									},
									"vrid": {
										Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid",
									},
									"domain": {
										Type: schema.TypeString, Optional: true, Description: "DNAT Using IP of a Domain (Domain Name)",
									},
									"pool": {
										Type: schema.TypeString, Optional: true, Description: "NAT Pool (NAT Pool or Pool Group)",
									},
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "The pool is a shared pool",
									},
									"dscp_direction": {
										Type: schema.TypeString, Optional: true, Description: "'inbound': To set dscp value for inbound packets; 'outbound': To set dscp value for outbound packets;",
									},
									"dscp_value": {
										Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
									},
								},
							},
						},
						"dscp_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dscp_match": {
										Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); 'any': Match any dscp value; '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
									},
									"action_cfg": {
										Type: schema.TypeString, Optional: true, Description: "'action': LSN Rule-List Action;",
									},
									"action_type": {
										Type: schema.TypeString, Optional: true, Description: "'set-dscp': To set dscp value for the packets;",
									},
									"dscp_direction": {
										Type: schema.TypeString, Optional: true, Description: "'inbound': To set dscp value for inbound packets; 'outbound': To set dscp value for outbound packets;",
									},
									"dscp_value": {
										Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110); '0': 000000; '1': 000001; '2': 000010; '3': 000011; '4': 000100; '5': 000101; '6': 000110; '7': 000111; '8': 001000; '9': 001001; '10': 001010; '11': 001011; '12': 001100; '13': 001101; '14': 001110; '15': 001111; '16': 010000; '17': 010001; '18': 010010; '19': 010011; '20': 010100; '21': 010101; '22': 010110; '23': 010111; '24': 011000; '25': 011001; '26': 011010; '27': 011011; '28': 011100; '29': 011101; '30': 011110; '31': 011111; '32': 100000; '33': 100001; '34': 100010; '35': 100011; '36': 100100; '37': 100101; '38': 100110; '39': 100111; '40': 101000; '41': 101001; '42': 101010; '43': 101011; '44': 101100; '45': 101101; '46': 101110; '47': 101111; '48': 110000; '49': 110001; '50': 110010; '51': 110011; '52': 110100; '53': 110101; '54': 110110; '55': 110111; '56': 111000; '57': 111001; '58': 111010; '59': 111011; '60': 111100; '61': 111101; '62': 111110; '63': 111111;",
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
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'placeholder': placeholder;",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6LsnRuleListIpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListIpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListIp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRuleListIpRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnRuleListIpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListIpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListIp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRuleListIpRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnRuleListIpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListIpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListIp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnRuleListIpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListIpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListIp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnRuleListIpRuleCfg(d []interface{}) []edpt.Cgnv6LsnRuleListIpRuleCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListIpRuleCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListIpRuleCfg
		oi.Proto = in["proto"].(string)
		oi.TcpCfg = getObjectCgnv6LsnRuleListIpRuleCfgTcpCfg(in["tcp_cfg"].([]interface{}))
		oi.UdpCfg = getObjectCgnv6LsnRuleListIpRuleCfgUdpCfg(in["udp_cfg"].([]interface{}))
		oi.IcmpOthersCfg = getObjectCgnv6LsnRuleListIpRuleCfgIcmpOthersCfg(in["icmp_others_cfg"].([]interface{}))
		oi.DscpCfg = getObjectCgnv6LsnRuleListIpRuleCfgDscpCfg(in["dscp_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListIpRuleCfgTcpCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpRuleCfgTcpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpRuleCfgTcpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartPort = in["start_port"].(int)
		ret.EndPort = in["end_port"].(int)
		ret.ActionCfg = in["action_cfg"].(string)
		ret.ActionType = in["action_type"].(string)
		ret.Ipv4List = in["ipv4_list"].(string)
		ret.PortList = in["port_list"].(string)
		ret.NoSnat = in["no_snat"].(int)
		ret.Vrid = in["vrid"].(int)
		ret.Domain = in["domain"].(string)
		ret.Pool = in["pool"].(string)
		ret.Shared = in["shared"].(int)
		ret.HttpAlg = in["http_alg"].(string)
		ret.TimeoutVal = in["timeout_val"].(int)
		ret.Fast = in["fast"].(string)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getObjectCgnv6LsnRuleListIpRuleCfgUdpCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpRuleCfgUdpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpRuleCfgUdpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StartPort = in["start_port"].(int)
		ret.EndPort = in["end_port"].(int)
		ret.ActionCfg = in["action_cfg"].(string)
		ret.ActionType = in["action_type"].(string)
		ret.Ipv4List = in["ipv4_list"].(string)
		ret.PortList = in["port_list"].(string)
		ret.NoSnat = in["no_snat"].(int)
		ret.Vrid = in["vrid"].(int)
		ret.Domain = in["domain"].(string)
		ret.Pool = in["pool"].(string)
		ret.Shared = in["shared"].(int)
		ret.TimeoutVal = in["timeout_val"].(int)
		ret.Fast = in["fast"].(string)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getObjectCgnv6LsnRuleListIpRuleCfgIcmpOthersCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpRuleCfgIcmpOthersCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpRuleCfgIcmpOthersCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActionCfg = in["action_cfg"].(string)
		ret.ActionType = in["action_type"].(string)
		ret.Ipv4List = in["ipv4_list"].(string)
		ret.NoSnat = in["no_snat"].(int)
		ret.Vrid = in["vrid"].(int)
		ret.Domain = in["domain"].(string)
		ret.Pool = in["pool"].(string)
		ret.Shared = in["shared"].(int)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getObjectCgnv6LsnRuleListIpRuleCfgDscpCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpRuleCfgDscpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpRuleCfgDscpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DscpMatch = in["dscp_match"].(string)
		ret.ActionCfg = in["action_cfg"].(string)
		ret.ActionType = in["action_type"].(string)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getSliceCgnv6LsnRuleListIpSamplingEnable(d []interface{}) []edpt.Cgnv6LsnRuleListIpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListIpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListIpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRuleListIp(d *schema.ResourceData) edpt.Cgnv6LsnRuleListIp {
	var ret edpt.Cgnv6LsnRuleListIp
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.RuleCfg = getSliceCgnv6LsnRuleListIpRuleCfg(d.Get("rule_cfg").([]interface{}))
	ret.Inst.SamplingEnable = getSliceCgnv6LsnRuleListIpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
