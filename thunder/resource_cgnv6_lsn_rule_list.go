package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LsnRuleList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_lsn_rule_list`: Configure LSN Rule-List\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LsnRuleListCreate,
		UpdateContext: resourceCgnv6LsnRuleListUpdate,
		ReadContext:   resourceCgnv6LsnRuleListRead,
		DeleteContext: resourceCgnv6LsnRuleListDelete,

		Schema: map[string]*schema.Schema{
			"default": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
					},
				},
			},
			"domain_ip": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
					},
				},
			},
			"domain_list_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_domain_list": {
							Type: schema.TypeString, Required: true, Description: "Configure a Specific Rule-Set (Domain List Name)",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"domain_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_domain": {
							Type: schema.TypeString, Required: true, Description: "Configure a Specific Rule-Set (Domain Name)",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"http_match_domain_name": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable match domain name in http request",
			},
			"ip_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "LSN Rule-List Name",
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
func resourceCgnv6LsnRuleListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRuleListRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LsnRuleListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LsnRuleListRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LsnRuleListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LsnRuleListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LsnRuleListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LsnRuleList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6LsnRuleListDefault88(d []interface{}) edpt.Cgnv6LsnRuleListDefault88 {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefault88
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleCfg = getSliceCgnv6LsnRuleListDefaultRuleCfg89(in["rule_cfg"].([]interface{}))
		//omit uuid
		ret.SamplingEnable = getSliceCgnv6LsnRuleListDefaultSamplingEnable94(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6LsnRuleListDefaultRuleCfg89(d []interface{}) []edpt.Cgnv6LsnRuleListDefaultRuleCfg89 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDefaultRuleCfg89, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDefaultRuleCfg89
		oi.Proto = in["proto"].(string)
		oi.TcpCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgTcpCfg90(in["tcp_cfg"].([]interface{}))
		oi.UdpCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgUdpCfg91(in["udp_cfg"].([]interface{}))
		oi.IcmpOthersCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg92(in["icmp_others_cfg"].([]interface{}))
		oi.DscpCfg = getObjectCgnv6LsnRuleListDefaultRuleCfgDscpCfg93(in["dscp_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDefaultRuleCfgTcpCfg90(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgTcpCfg90 {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgTcpCfg90
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

func getObjectCgnv6LsnRuleListDefaultRuleCfgUdpCfg91(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgUdpCfg91 {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgUdpCfg91
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

func getObjectCgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg92(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg92 {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg92
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

func getObjectCgnv6LsnRuleListDefaultRuleCfgDscpCfg93(d []interface{}) edpt.Cgnv6LsnRuleListDefaultRuleCfgDscpCfg93 {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDefaultRuleCfgDscpCfg93
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

func getSliceCgnv6LsnRuleListDefaultSamplingEnable94(d []interface{}) []edpt.Cgnv6LsnRuleListDefaultSamplingEnable94 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDefaultSamplingEnable94, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDefaultSamplingEnable94
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDomainIp95(d []interface{}) edpt.Cgnv6LsnRuleListDomainIp95 {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainIp95
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		//omit uuid
		ret.SamplingEnable = getSliceCgnv6LsnRuleListDomainIpSamplingEnable96(in["sampling_enable"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6LsnRuleListDomainIpSamplingEnable96(d []interface{}) []edpt.Cgnv6LsnRuleListDomainIpSamplingEnable96 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainIpSamplingEnable96, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainIpSamplingEnable96
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnRuleListDomainListNameList(d []interface{}) []edpt.Cgnv6LsnRuleListDomainListNameList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainListNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainListNameList
		oi.NameDomainList = in["name_domain_list"].(string)
		oi.RuleCfg = getSliceCgnv6LsnRuleListDomainListNameListRuleCfg(in["rule_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceCgnv6LsnRuleListDomainListNameListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnRuleListDomainListNameListRuleCfg(d []interface{}) []edpt.Cgnv6LsnRuleListDomainListNameListRuleCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainListNameListRuleCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainListNameListRuleCfg
		oi.Proto = in["proto"].(string)
		oi.TcpCfg = getObjectCgnv6LsnRuleListDomainListNameListRuleCfgTcpCfg(in["tcp_cfg"].([]interface{}))
		oi.UdpCfg = getObjectCgnv6LsnRuleListDomainListNameListRuleCfgUdpCfg(in["udp_cfg"].([]interface{}))
		oi.IcmpOthersCfg = getObjectCgnv6LsnRuleListDomainListNameListRuleCfgIcmpOthersCfg(in["icmp_others_cfg"].([]interface{}))
		oi.DscpCfg = getObjectCgnv6LsnRuleListDomainListNameListRuleCfgDscpCfg(in["dscp_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDomainListNameListRuleCfgTcpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgTcpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgTcpCfg
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

func getObjectCgnv6LsnRuleListDomainListNameListRuleCfgUdpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgUdpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgUdpCfg
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

func getObjectCgnv6LsnRuleListDomainListNameListRuleCfgIcmpOthersCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgIcmpOthersCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgIcmpOthersCfg
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

func getObjectCgnv6LsnRuleListDomainListNameListRuleCfgDscpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgDscpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainListNameListRuleCfgDscpCfg
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

func getSliceCgnv6LsnRuleListDomainListNameListSamplingEnable(d []interface{}) []edpt.Cgnv6LsnRuleListDomainListNameListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainListNameListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainListNameListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnRuleListDomainNameList(d []interface{}) []edpt.Cgnv6LsnRuleListDomainNameList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainNameList
		oi.NameDomain = in["name_domain"].(string)
		oi.RuleCfg = getSliceCgnv6LsnRuleListDomainNameListRuleCfg(in["rule_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceCgnv6LsnRuleListDomainNameListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnRuleListDomainNameListRuleCfg(d []interface{}) []edpt.Cgnv6LsnRuleListDomainNameListRuleCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainNameListRuleCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainNameListRuleCfg
		oi.Proto = in["proto"].(string)
		oi.TcpCfg = getObjectCgnv6LsnRuleListDomainNameListRuleCfgTcpCfg(in["tcp_cfg"].([]interface{}))
		oi.UdpCfg = getObjectCgnv6LsnRuleListDomainNameListRuleCfgUdpCfg(in["udp_cfg"].([]interface{}))
		oi.IcmpOthersCfg = getObjectCgnv6LsnRuleListDomainNameListRuleCfgIcmpOthersCfg(in["icmp_others_cfg"].([]interface{}))
		oi.DscpCfg = getObjectCgnv6LsnRuleListDomainNameListRuleCfgDscpCfg(in["dscp_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListDomainNameListRuleCfgTcpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainNameListRuleCfgTcpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainNameListRuleCfgTcpCfg
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

func getObjectCgnv6LsnRuleListDomainNameListRuleCfgUdpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainNameListRuleCfgUdpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainNameListRuleCfgUdpCfg
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

func getObjectCgnv6LsnRuleListDomainNameListRuleCfgIcmpOthersCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainNameListRuleCfgIcmpOthersCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainNameListRuleCfgIcmpOthersCfg
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

func getObjectCgnv6LsnRuleListDomainNameListRuleCfgDscpCfg(d []interface{}) edpt.Cgnv6LsnRuleListDomainNameListRuleCfgDscpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListDomainNameListRuleCfgDscpCfg
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

func getSliceCgnv6LsnRuleListDomainNameListSamplingEnable(d []interface{}) []edpt.Cgnv6LsnRuleListDomainNameListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListDomainNameListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListDomainNameListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnRuleListIpList(d []interface{}) []edpt.Cgnv6LsnRuleListIpList {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListIpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListIpList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.RuleCfg = getSliceCgnv6LsnRuleListIpListRuleCfg(in["rule_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceCgnv6LsnRuleListIpListSamplingEnable(in["sampling_enable"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6LsnRuleListIpListRuleCfg(d []interface{}) []edpt.Cgnv6LsnRuleListIpListRuleCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListIpListRuleCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListIpListRuleCfg
		oi.Proto = in["proto"].(string)
		oi.TcpCfg = getObjectCgnv6LsnRuleListIpListRuleCfgTcpCfg(in["tcp_cfg"].([]interface{}))
		oi.UdpCfg = getObjectCgnv6LsnRuleListIpListRuleCfgUdpCfg(in["udp_cfg"].([]interface{}))
		oi.IcmpOthersCfg = getObjectCgnv6LsnRuleListIpListRuleCfgIcmpOthersCfg(in["icmp_others_cfg"].([]interface{}))
		oi.DscpCfg = getObjectCgnv6LsnRuleListIpListRuleCfgDscpCfg(in["dscp_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6LsnRuleListIpListRuleCfgTcpCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpListRuleCfgTcpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpListRuleCfgTcpCfg
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

func getObjectCgnv6LsnRuleListIpListRuleCfgUdpCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpListRuleCfgUdpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpListRuleCfgUdpCfg
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

func getObjectCgnv6LsnRuleListIpListRuleCfgIcmpOthersCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpListRuleCfgIcmpOthersCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpListRuleCfgIcmpOthersCfg
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

func getObjectCgnv6LsnRuleListIpListRuleCfgDscpCfg(d []interface{}) edpt.Cgnv6LsnRuleListIpListRuleCfgDscpCfg {

	count1 := len(d)
	var ret edpt.Cgnv6LsnRuleListIpListRuleCfgDscpCfg
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

func getSliceCgnv6LsnRuleListIpListSamplingEnable(d []interface{}) []edpt.Cgnv6LsnRuleListIpListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.Cgnv6LsnRuleListIpListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6LsnRuleListIpListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointCgnv6LsnRuleList(d *schema.ResourceData) edpt.Cgnv6LsnRuleList {
	var ret edpt.Cgnv6LsnRuleList
	ret.Inst.Default = getObjectCgnv6LsnRuleListDefault88(d.Get("default").([]interface{}))
	ret.Inst.DomainIp = getObjectCgnv6LsnRuleListDomainIp95(d.Get("domain_ip").([]interface{}))
	ret.Inst.DomainListNameList = getSliceCgnv6LsnRuleListDomainListNameList(d.Get("domain_list_name_list").([]interface{}))
	ret.Inst.DomainNameList = getSliceCgnv6LsnRuleListDomainNameList(d.Get("domain_name_list").([]interface{}))
	ret.Inst.HttpMatchDomainName = d.Get("http_match_domain_name").(int)
	ret.Inst.IpList = getSliceCgnv6LsnRuleListIpList(d.Get("ip_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
