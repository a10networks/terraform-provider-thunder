package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleListDefault() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_rule_list_default`: Configure the Generic Rule-Set\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnRuleListDefaultCreate,
		UpdateContext: resourceCgnv6LsnRuleListDefaultUpdate,
		ReadContext:   resourceCgnv6LsnRuleListDefaultRead,
		DeleteContext: resourceCgnv6LsnRuleListDefaultDelete,

		Schema: map[string]*schema.Schema{
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
										Type: schema.TypeString, Optional: true, Description: "'dnat': Apply Dest NAT; 'drop': Drop the Packets; 'one-to-one-snat': Apply one-to-one source NAT for the packets; 'pass-through': Pass the Packets Through; 'snat': Redirect the Packets to a Different Source NAT Pool; 'set-dscp': To set dscp value for the packets; 'template': Template;",
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
									"pool": {
										Type: schema.TypeString, Optional: true, Description: "NAT Pool (NAT Pool or Pool Group)",
									},
									"shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "The pool is a shared pool",
									},
									"http_alg": {
										Type: schema.TypeString, Optional: true, Description: "HTTP-ALG Template (Template Name)",
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
										Type: schema.TypeString, Optional: true, Description: "'dnat': Apply Dest NAT; 'drop': Drop the Packets; 'one-to-one-snat': Apply one-to-one source NAT for the packets; 'pass-through': Pass the Packets Through; 'snat': Redirect the Packets to a Different Source NAT Pool; 'set-dscp': To set dscp value for the packets;",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceCgnv6LsnRuleListDefaultCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDefaultCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDefault(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRuleListDefaultRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnRuleListDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDefaultUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDefault(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRuleListDefaultRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnRuleListDefaultDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDefaultDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDefault(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnRuleListDefaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDefaultRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleListDefault(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceCgnv6LsnRuleListDefaultRuleCfg(d []interface{}) []edpt.Cgnv6LsnRuleListDefaultRuleCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDefaultRuleCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDefaultRuleCfg
		oi.Proto = in["proto"].(string)
		oi.TcpCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgTcpCfg(in["tcp_cfg"].([]interface{}))
		oi.UdpCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgUdpCfg(in["udp_cfg"].([]interface{}))
		oi.IcmpOthersCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg(in["icmp_others_cfg"].([]interface{}))
		oi.DscpCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgDscpCfg(in["dscp_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDefaultRuleCfgTcpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgTcpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgTcpCfg
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
		ret.Pool = in["pool"].(string)
		ret.Shared = in["shared"].(int)
		ret.HttpAlg = in["http_alg"].(string)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDefaultRuleCfgUdpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgUdpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgUdpCfg
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
		ret.Pool = in["pool"].(string)
		ret.Shared = in["shared"].(int)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ActionCfg = in["action_cfg"].(string)
		ret.ActionType = in["action_type"].(string)
		ret.Ipv4List = in["ipv4_list"].(string)
		ret.NoSnat = in["no_snat"].(int)
		ret.Vrid = in["vrid"].(int)
		ret.Pool = in["pool"].(string)
		ret.Shared = in["shared"].(int)
		ret.DscpDirection = in["dscp_direction"].(string)
		ret.DscpValue = in["dscp_value"].(string)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDefaultRuleCfgDscpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgDscpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgDscpCfg
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

func getSliceCgnv6LsnRuleListDefaultSamplingEnable(d []interface{}) []edpt.Cgnv6LsnRuleListDefaultSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDefaultSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDefaultSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRuleListDefault(d *schema.ResourceData) edpt.Cgnv6LsnRuleListDefault {
	var ret edpt.Cgnv6LsnRuleListDefault
	ret.Inst.RuleCfg = getSliceCgnv6LsnRuleListDefaultRuleCfg(d.Get("rule_cfg").([]interface{}))
	ret.Inst.SamplingEnable = getSliceCgnv6LsnRuleListDefaultSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
