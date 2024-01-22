package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6TemplateLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_template_logging`: Logging Template\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6TemplateLoggingCreate,
		UpdateContext: resourceCgnv6TemplateLoggingUpdate,
		ReadContext:   resourceCgnv6TemplateLoggingRead,
		DeleteContext: resourceCgnv6TemplateLoggingDelete,

		Schema: map[string]*schema.Schema{
			"batched_logging_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable multiple logs per packet",
			},
			"custom": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_header": {
							Type: schema.TypeString, Optional: true, Description: "'use-syslog-header': Use syslog header as custom log header;",
						},
						"custom_message": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_dhcpv6_map_prefix_assigned": {
										Type: schema.TypeString, Optional: true, Description: "MAP DHCPv6 prefix assigned",
									},
									"custom_dhcpv6_map_prefix_released": {
										Type: schema.TypeString, Optional: true, Description: "MAP DHCPv6 prefix released",
									},
									"custom_dhcpv6_map_prefix_renewed": {
										Type: schema.TypeString, Optional: true, Description: "MAP DHCPv6 prefix renewed",
									},
									"custom_fixed_nat_allocated": {
										Type: schema.TypeString, Optional: true, Description: "Fixed-NAT allocated (Custom message string)",
									},
									"custom_fixed_nat_interim_update": {
										Type: schema.TypeString, Optional: true, Description: "Fixed-NAT interim update (Custom message string)",
									},
									"custom_fixed_nat_freed": {
										Type: schema.TypeString, Optional: true, Description: "Fixed-NAT freed (Custom message string)",
									},
									"custom_http_request_got": {
										Type: schema.TypeString, Optional: true, Description: "HTTP request got (Custom message string)",
									},
									"custom_port_allocated": {
										Type: schema.TypeString, Optional: true, Description: "Port allocated (Custom message string)",
									},
									"custom_port_batch_allocated": {
										Type: schema.TypeString, Optional: true, Description: "Port Batch allocated (Custom message string)",
									},
									"custom_port_batch_freed": {
										Type: schema.TypeString, Optional: true, Description: "Port Batch freed (Custom message string)",
									},
									"custom_port_batch_v2_allocated": {
										Type: schema.TypeString, Optional: true, Description: "Port Batch v2 allocated (Custom message string)",
									},
									"custom_port_batch_v2_freed": {
										Type: schema.TypeString, Optional: true, Description: "Port Batch v2 freed (Custom message string)",
									},
									"custom_port_batch_v2_interim_update": {
										Type: schema.TypeString, Optional: true, Description: "Port Batch v2 interim update (Custom message string)",
									},
									"custom_port_freed": {
										Type: schema.TypeString, Optional: true, Description: "Port freed (Custom message string)",
									},
									"custom_session_created": {
										Type: schema.TypeString, Optional: true, Description: "Session created (Custom message string)",
									},
									"custom_session_deleted": {
										Type: schema.TypeString, Optional: true, Description: "Session deleted (Custom message string)",
									},
								},
							},
						},
						"custom_time_stamp_format": {
							Type: schema.TypeString, Optional: true, Description: "Customize the time stamp format (Customize the time-stamp format. Default:%Y%m%d%H%M%S)",
						},
					},
				},
			},
			"disable_log_by_destination": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
									},
									"tcp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Port Range End",
									},
								},
							},
						},
						"udp_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"udp_port_start": {
										Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
									},
									"udp_port_end": {
										Type: schema.TypeInt, Optional: true, Description: "Port Range End",
									},
								},
							},
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
						},
						"others": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"ip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_addr": {
										Type: schema.TypeString, Required: true, Description: "Configure an IP subnet",
									},
									"tcp_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
												},
												"tcp_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "Port Range End",
												},
											},
										},
									},
									"udp_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"udp_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
												},
												"udp_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "Port Range End",
												},
											},
										},
									},
									"icmp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
									},
									"others": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
						"ip6_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_addr": {
										Type: schema.TypeString, Required: true, Description: "Configure an IPv6 subnet",
									},
									"tcp_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
												},
												"tcp_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "Port Range End",
												},
											},
										},
									},
									"udp_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"udp_port_start": {
													Type: schema.TypeInt, Optional: true, Description: "Destination Port (Single Destination Port or Port Range Start)",
												},
												"udp_port_end": {
													Type: schema.TypeInt, Optional: true, Description: "Port Range End",
												},
											},
										},
									},
									"icmp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for icmp traffic",
									},
									"others": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable logging for other L4 protocols",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"facility": {
				Type: schema.TypeString, Optional: true, Default: "local0", Description: "'kernel': 0: Kernel; 'user': 1: User-level; 'mail': 2: Mail; 'daemon': 3: System daemons; 'security-authorization': 4: Security/authorization; 'syslog': 5: Syslog internal; 'line-printer': 6: Line printer; 'news': 7: Network news; 'uucp': 8: UUCP subsystem; 'cron': 9: Time-related; 'security-authorization-private': 10: Private security/authorization; 'ftp': 11: FTP; 'ntp': 12: NTP; 'audit': 13: Audit; 'alert': 14: Alert; 'clock': 15: Clock-related; 'local0': 16: Local use 0; 'local1': 17: Local use 1; 'local2': 18: Local use 2; 'local3': 19: Local use 3; 'local4': 20: Local use 4; 'local5': 21: Local use 5; 'local6': 22: Local use 6; 'local7': 23: Local use 7;",
			},
			"format": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'binary': Binary logging format; 'compact': Compact ASCII logging format (Hex format with compact representation); 'custom': Arbitrary custom logging format; 'default': Default A10 logging format (ASCII); 'rfc5424': RFC5424 compliant logging format; 'cef': Common Event Format for logging;",
			},
			"include_destination": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the destination IP and port in logs",
			},
			"include_http": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_header": {
										Type: schema.TypeString, Optional: true, Description: "'cookie': Log HTTP Cookie Header; 'referer': Log HTTP Referer Header; 'user-agent': Log HTTP User-Agent Header; 'header1': Log HTTP Header 1; 'header2': Log HTTP Header 2; 'header3': Log HTTP Header 3;",
									},
									"max_length": {
										Type: schema.TypeInt, Optional: true, Default: 100, Description: "Max length for a HTTP request log (Max header length (Default: 100 char))",
									},
									"custom_header_name": {
										Type: schema.TypeString, Optional: true, Description: "Header name",
									},
									"custom_max_length": {
										Type: schema.TypeInt, Optional: true, Default: 100, Description: "Max length for a HTTP request log (Max header length (Default: 100 char))",
									},
								},
							},
						},
						"l4_session_info": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the L4 session information of the HTTP request",
						},
						"method": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log the HTTP Request Method",
						},
						"request_number": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP Request Number",
						},
						"file_extension": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP file extension",
						},
					},
				},
			},
			"include_inside_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the inside user MAC address in logs",
			},
			"include_partition_name": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include partition name in logging events",
			},
			"include_port_block_account": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "include bytes accounting information in port-batch-v2 port-mapping and fixed-nat user-ports messages",
			},
			"include_radius_attribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attr": {
										Type: schema.TypeString, Optional: true, Description: "'imei': Include IMEI; 'imsi': Include IMSI; 'msisdn': Include MSISDN; 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;",
									},
									"attr_event": {
										Type: schema.TypeString, Optional: true, Description: "'http-requests': Include in HTTP request logs; 'port-mappings': Include in port-mapping logs; 'sessions': Include in session logs; 'user-data': Include in user-data logs;",
									},
								},
							},
						},
						"no_quote": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No quotation marks for RADIUS attributes in logs",
						},
						"insert_if_not_existing": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure what string is to be inserted for custom RADIUS attributes",
						},
						"zero_in_custom_attr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Insert 0000 for standard and custom attributes in log string",
						},
						"framed_ipv6_prefix": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include radius attributes for the prefix",
						},
						"prefix_length": {
							Type: schema.TypeString, Optional: true, Description: "'32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;",
						},
					},
				},
			},
			"include_session_byte_count": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "include byte count in session deletion logs",
			},
			"log": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fixed_nat": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fixed_nat_http_requests": {
										Type: schema.TypeString, Optional: true, Description: "'host': Log the HTTP Host Header; 'url': Log the HTTP Request URL;",
									},
									"fixed_nat_port_mappings": {
										Type: schema.TypeString, Optional: true, Description: "'both': Log creation and deletion of NAT mappings; 'creation': Log creation of NAT mappings;",
									},
									"fixed_nat_sessions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log all Fixed NAT sessions",
									},
									"fixed_nat_merged_style": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Merge creation and deletion of session logs to one",
									},
									"user_ports": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_ports": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log Fixed NAT User Ports Configured",
												},
												"days": {
													Type: schema.TypeInt, Optional: true, Description: "Specify period in days",
												},
												"start_time": {
													Type: schema.TypeString, Optional: true, Description: "Time when periodic logging starts (Specify start time(hh:mm))",
												},
											},
										},
									},
								},
							},
						},
						"one_to_one_nat": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"one_to_one_nat_sessions": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log all One-to-One NAT sessions",
									},
									"one_to_one_merged_style": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Merge creation and deletion of session logs to one",
									},
								},
							},
						},
						"map_dhcpv6": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"map_dhcpv6_prefix_all": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log MAP DHCPv6 prefix assignment/renewal/release",
									},
									"map_dhcpv6_msg_type": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"map_dhcpv6_msg_type": {
													Type: schema.TypeString, Optional: true, Description: "'prefix-assignment': Log MAP DHCPv6 prefix assignment; 'prefix-renewal': Log MAP DHCPv6 prefix renewal; 'prefix-release': Log MAP DHCPv6 prefix release;",
												},
											},
										},
									},
								},
							},
						},
						"http_requests": {
							Type: schema.TypeString, Optional: true, Description: "'host': Log the HTTP Host Header; 'url': Log the HTTP Request URL;",
						},
						"port_mappings": {
							Type: schema.TypeString, Optional: true, Default: "both", Description: "'creation': Log only creation of NAT mappings; 'disable': Disable Log creation and deletion of NAT mappings; 'both': Log creation and deletion of NAT mappings;",
						},
						"port_overloading": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force logging of all port-overloading sessions",
						},
						"user_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log LSN Subscriber Information",
						},
						"sessions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log all data sessions created using NAT",
						},
						"merged_style": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Merge creation and deletion of session logs to one",
						},
					},
				},
			},
			"log_receiver": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"radius": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use RADIUS server for NAT logging",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "The RADIUS server's secret",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Logging template name",
			},
			"resolution": {
				Type: schema.TypeString, Optional: true, Default: "seconds", Description: "'seconds': Logging timestamp resolution in seconds (default); '10-milliseconds': Logging timestamp resolution in 10s of milli-seconds;",
			},
			"rfc_custom": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"use_alternate_timestamp": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use alternate non-RFC5424 compliant timestamp. Ex: 1990 Jan 15 12:30:30",
									},
								},
							},
						},
						"message": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_tech": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tech_type": {
													Type: schema.TypeString, Optional: true, Description: "'lsn': LSN; 'nat64': NAT64; 'ds-lite': DS-Lite; 'sixrd-nat64': 6rd-NAT64;",
												},
												"fixed_nat_allocated": {
													Type: schema.TypeString, Optional: true, Description: "Fixed-NAT allocated (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"fixed_nat_freed": {
													Type: schema.TypeString, Optional: true, Description: "Fixed-NAT freed (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"port_allocated": {
													Type: schema.TypeString, Optional: true, Description: "Port allocated (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"port_freed": {
													Type: schema.TypeString, Optional: true, Description: "Port freed (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"port_batch_allocated": {
													Type: schema.TypeString, Optional: true, Description: "Port Batch allocated (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"port_batch_freed": {
													Type: schema.TypeString, Optional: true, Description: "Port Batch freed (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"port_batch_v2_allocated": {
													Type: schema.TypeString, Optional: true, Description: "Port Batch v2 allocated (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
												"port_batch_v2_freed": {
													Type: schema.TypeString, Optional: true, Description: "Port Batch v2 freed (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
												},
											},
										},
									},
									"dhcpv6_map_prefix_assigned": {
										Type: schema.TypeString, Optional: true, Description: "MAP DHCPv6 prefix assigned",
									},
									"dhcpv6_map_prefix_released": {
										Type: schema.TypeString, Optional: true, Description: "MAP DHCPv6 prefix released",
									},
									"dhcpv6_map_prefix_renewed": {
										Type: schema.TypeString, Optional: true, Description: "MAP DHCPv6 prefix renewed",
									},
									"http_request_got": {
										Type: schema.TypeString, Optional: true, Description: "HTTP request got (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
									},
									"session_created": {
										Type: schema.TypeString, Optional: true, Description: "Session created (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
									},
									"session_deleted": {
										Type: schema.TypeString, Optional: true, Description: "Session deleted (Custom message string. Should be in the format of \"MSGID [STRUCTURED-DATA] MSG\")",
									},
								},
							},
						},
					},
				},
			},
			"rule": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_http_requests": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dest_port": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dest_port_number": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"include_byte_count": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include the byte count of HTTP Request/Response in CGN session deletion logs",
												},
											},
										},
									},
									"log_every_http_request": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log every HTTP request in an HTTP 1.1 session (Default: Log the first HTTP request in a session)",
									},
									"max_url_len": {
										Type: schema.TypeInt, Optional: true, Default: 128, Description: "Max length of URL log (Max URL length (Default: 128 char))",
									},
									"include_all_headers": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include all configured headers despite of absence in HTTP request",
									},
									"disable_sequence_check": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable http packet sequence check and don't drop out of order packets",
									},
								},
							},
						},
						"interim_update_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Log interim update of NAT mappings (Interim update interval in minutes(Interval is floored to a multiple of 5))",
						},
					},
				},
			},
			"service_group": {
				Type: schema.TypeString, Optional: true, Description: "Set NAT logging service-group",
			},
			"severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity_string": {
							Type: schema.TypeString, Optional: true, Default: "debug", Description: "'emergency': 0: Emergency; 'alert': 1: Alert; 'critical': 2: Critical; 'error': 3: Error; 'warning': 4: Warning; 'notice': 5: Notice; 'informational': 6: Informational; 'debug': 7: Debug;",
						},
						"severity_val": {
							Type: schema.TypeInt, Optional: true, Default: 7, Description: "Logging severity level",
						},
					},
				},
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Service group is in shared patition",
			},
			"source_address": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type: schema.TypeString, Optional: true, Description: "Specify source IP address",
						},
						"ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify source IPv6 address",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"source_port": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_port_num": {
							Type: schema.TypeInt, Optional: true, Default: 514, Description: "Set source port for sending NAT syslogs (default: 514)",
						},
						"any": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use any source port",
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
func resourceCgnv6TemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6TemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6TemplateLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6TemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6TemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6TemplateLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6TemplateLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6TemplateLoggingCustom(d []interface{}) edpt.Cgnv6TemplateLoggingCustom {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingCustom
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CustomHeader = in["custom_header"].(string)
		ret.CustomMessage = getObjectCgnv6TemplateLoggingCustomCustomMessage(in["custom_message"].([]interface{}))
		ret.CustomTimeStampFormat = in["custom_time_stamp_format"].(string)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingCustomCustomMessage(d []interface{}) edpt.Cgnv6TemplateLoggingCustomCustomMessage {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingCustomCustomMessage
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CustomDhcpv6MapPrefixAssigned = in["custom_dhcpv6_map_prefix_assigned"].(string)
		ret.CustomDhcpv6MapPrefixReleased = in["custom_dhcpv6_map_prefix_released"].(string)
		ret.CustomDhcpv6MapPrefixRenewed = in["custom_dhcpv6_map_prefix_renewed"].(string)
		ret.CustomFixedNatAllocated = in["custom_fixed_nat_allocated"].(string)
		ret.CustomFixedNatInterimUpdate = in["custom_fixed_nat_interim_update"].(string)
		ret.CustomFixedNatFreed = in["custom_fixed_nat_freed"].(string)
		ret.CustomHttpRequestGot = in["custom_http_request_got"].(string)
		ret.CustomPortAllocated = in["custom_port_allocated"].(string)
		ret.CustomPortBatchAllocated = in["custom_port_batch_allocated"].(string)
		ret.CustomPortBatchFreed = in["custom_port_batch_freed"].(string)
		ret.CustomPortBatchV2Allocated = in["custom_port_batch_v2_allocated"].(string)
		ret.CustomPortBatchV2Freed = in["custom_port_batch_v2_freed"].(string)
		ret.CustomPortBatchV2InterimUpdate = in["custom_port_batch_v2_interim_update"].(string)
		ret.CustomPortFreed = in["custom_port_freed"].(string)
		ret.CustomSessionCreated = in["custom_session_created"].(string)
		ret.CustomSessionDeleted = in["custom_session_deleted"].(string)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingDisableLogByDestination115(d []interface{}) edpt.Cgnv6TemplateLoggingDisableLogByDestination115 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingDisableLogByDestination115
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationTcpList116(in["tcp_list"].([]interface{}))
		ret.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationUdpList117(in["udp_list"].([]interface{}))
		ret.Icmp = in["icmp"].(int)
		ret.Others = in["others"].(int)
		//omit uuid
		ret.IpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpList118(in["ip_list"].([]interface{}))
		ret.Ip6List = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6List121(in["ip6_list"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationTcpList116(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationTcpList116 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationTcpList116, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationTcpList116
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationUdpList117(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationUdpList117 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationUdpList117, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationUdpList117
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpList118(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpList118 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpList118, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpList118
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListTcpList119
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIpListUdpList120
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6List121(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6List121 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6List121, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6List121
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListTcpList122
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123(d []interface{}) []edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123 {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingDisableLogByDestinationIp6ListUdpList123
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingIncludeHttp(d []interface{}) edpt.Cgnv6TemplateLoggingIncludeHttp {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingIncludeHttp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HeaderCfg = getSliceCgnv6TemplateLoggingIncludeHttpHeaderCfg(in["header_cfg"].([]interface{}))
		ret.L4SessionInfo = in["l4_session_info"].(int)
		ret.Method = in["method"].(int)
		ret.RequestNumber = in["request_number"].(int)
		ret.FileExtension = in["file_extension"].(int)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingIncludeHttpHeaderCfg(d []interface{}) []edpt.Cgnv6TemplateLoggingIncludeHttpHeaderCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingIncludeHttpHeaderCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingIncludeHttpHeaderCfg
		oi.HttpHeader = in["http_header"].(string)
		oi.MaxLength = in["max_length"].(int)
		oi.CustomHeaderName = in["custom_header_name"].(string)
		oi.CustomMaxLength = in["custom_max_length"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingIncludeRadiusAttribute(d []interface{}) edpt.Cgnv6TemplateLoggingIncludeRadiusAttribute {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingIncludeRadiusAttribute
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AttrCfg = getSliceCgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg(in["attr_cfg"].([]interface{}))
		ret.NoQuote = in["no_quote"].(int)
		ret.InsertIfNotExisting = in["insert_if_not_existing"].(int)
		ret.ZeroInCustomAttr = in["zero_in_custom_attr"].(int)
		ret.FramedIpv6Prefix = in["framed_ipv6_prefix"].(int)
		ret.PrefixLength = in["prefix_length"].(string)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg(d []interface{}) []edpt.Cgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingIncludeRadiusAttributeAttrCfg
		oi.Attr = in["attr"].(string)
		oi.AttrEvent = in["attr_event"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingLog(d []interface{}) edpt.Cgnv6TemplateLoggingLog {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingLog
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixedNat = getObjectCgnv6TemplateLoggingLogFixedNat(in["fixed_nat"].([]interface{}))
		ret.OneToOneNat = getObjectCgnv6TemplateLoggingLogOneToOneNat(in["one_to_one_nat"].([]interface{}))
		ret.MapDhcpv6 = getObjectCgnv6TemplateLoggingLogMapDhcpv6(in["map_dhcpv6"].([]interface{}))
		ret.HttpRequests = in["http_requests"].(string)
		ret.PortMappings = in["port_mappings"].(string)
		ret.PortOverloading = in["port_overloading"].(int)
		ret.UserData = in["user_data"].(int)
		ret.Sessions = in["sessions"].(int)
		ret.MergedStyle = in["merged_style"].(int)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingLogFixedNat(d []interface{}) edpt.Cgnv6TemplateLoggingLogFixedNat {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingLogFixedNat
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FixedNatHttpRequests = in["fixed_nat_http_requests"].(string)
		ret.FixedNatPortMappings = in["fixed_nat_port_mappings"].(string)
		ret.FixedNatSessions = in["fixed_nat_sessions"].(int)
		ret.FixedNatMergedStyle = in["fixed_nat_merged_style"].(int)
		ret.UserPorts = getObjectCgnv6TemplateLoggingLogFixedNatUserPorts(in["user_ports"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6TemplateLoggingLogFixedNatUserPorts(d []interface{}) edpt.Cgnv6TemplateLoggingLogFixedNatUserPorts {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingLogFixedNatUserPorts
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UserPorts = in["user_ports"].(int)
		ret.Days = in["days"].(int)
		ret.StartTime = in["start_time"].(string)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingLogOneToOneNat(d []interface{}) edpt.Cgnv6TemplateLoggingLogOneToOneNat {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingLogOneToOneNat
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OneToOneNatSessions = in["one_to_one_nat_sessions"].(int)
		ret.OneToOneMergedStyle = in["one_to_one_merged_style"].(int)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingLogMapDhcpv6(d []interface{}) edpt.Cgnv6TemplateLoggingLogMapDhcpv6 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingLogMapDhcpv6
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MapDhcpv6PrefixAll = in["map_dhcpv6_prefix_all"].(int)
		ret.MapDhcpv6MsgType = getSliceCgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType(in["map_dhcpv6_msg_type"].([]interface{}))
	}
	return ret
}

func getSliceCgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType(d []interface{}) []edpt.Cgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingLogMapDhcpv6MapDhcpv6MsgType
		oi.MapDhcpv6MsgType = in["map_dhcpv6_msg_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingLogReceiver(d []interface{}) edpt.Cgnv6TemplateLoggingLogReceiver {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingLogReceiver
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Radius = in["radius"].(int)
		ret.SecretString = in["secret_string"].(string)
		//omit encrypted
	}
	return ret
}

func getObjectCgnv6TemplateLoggingRfcCustom(d []interface{}) edpt.Cgnv6TemplateLoggingRfcCustom {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingRfcCustom
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Header = getObjectCgnv6TemplateLoggingRfcCustomHeader(in["header"].([]interface{}))
		ret.Message = getObjectCgnv6TemplateLoggingRfcCustomMessage(in["message"].([]interface{}))
	}
	return ret
}

func getObjectCgnv6TemplateLoggingRfcCustomHeader(d []interface{}) edpt.Cgnv6TemplateLoggingRfcCustomHeader {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingRfcCustomHeader
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UseAlternateTimestamp = in["use_alternate_timestamp"].(int)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingRfcCustomMessage(d []interface{}) edpt.Cgnv6TemplateLoggingRfcCustomMessage {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingRfcCustomMessage
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6Tech = getSliceCgnv6TemplateLoggingRfcCustomMessageIpv6Tech(in["ipv6_tech"].([]interface{}))
		ret.Dhcpv6MapPrefixAssigned = in["dhcpv6_map_prefix_assigned"].(string)
		ret.Dhcpv6MapPrefixReleased = in["dhcpv6_map_prefix_released"].(string)
		ret.Dhcpv6MapPrefixRenewed = in["dhcpv6_map_prefix_renewed"].(string)
		ret.HttpRequestGot = in["http_request_got"].(string)
		ret.SessionCreated = in["session_created"].(string)
		ret.SessionDeleted = in["session_deleted"].(string)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingRfcCustomMessageIpv6Tech(d []interface{}) []edpt.Cgnv6TemplateLoggingRfcCustomMessageIpv6Tech {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingRfcCustomMessageIpv6Tech, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingRfcCustomMessageIpv6Tech
		oi.TechType = in["tech_type"].(string)
		oi.FixedNatAllocated = in["fixed_nat_allocated"].(string)
		oi.FixedNatFreed = in["fixed_nat_freed"].(string)
		oi.PortAllocated = in["port_allocated"].(string)
		oi.PortFreed = in["port_freed"].(string)
		oi.PortBatchAllocated = in["port_batch_allocated"].(string)
		oi.PortBatchFreed = in["port_batch_freed"].(string)
		oi.PortBatchV2Allocated = in["port_batch_v2_allocated"].(string)
		oi.PortBatchV2Freed = in["port_batch_v2_freed"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingRule(d []interface{}) edpt.Cgnv6TemplateLoggingRule {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingRule
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RuleHttpRequests = getObjectCgnv6TemplateLoggingRuleRuleHttpRequests(in["rule_http_requests"].([]interface{}))
		ret.InterimUpdateInterval = in["interim_update_interval"].(int)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingRuleRuleHttpRequests(d []interface{}) edpt.Cgnv6TemplateLoggingRuleRuleHttpRequests {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingRuleRuleHttpRequests
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DestPort = getSliceCgnv6TemplateLoggingRuleRuleHttpRequestsDestPort(in["dest_port"].([]interface{}))
		ret.LogEveryHttpRequest = in["log_every_http_request"].(int)
		ret.MaxUrlLen = in["max_url_len"].(int)
		ret.IncludeAllHeaders = in["include_all_headers"].(int)
		ret.DisableSequenceCheck = in["disable_sequence_check"].(int)
	}
	return ret
}

func getSliceCgnv6TemplateLoggingRuleRuleHttpRequestsDestPort(d []interface{}) []edpt.Cgnv6TemplateLoggingRuleRuleHttpRequestsDestPort {

	count1 := len(d)
	ret := make([]edpt.Cgnv6TemplateLoggingRuleRuleHttpRequestsDestPort, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.Cgnv6TemplateLoggingRuleRuleHttpRequestsDestPort
		oi.DestPortNumber = in["dest_port_number"].(int)
		oi.IncludeByteCount = in["include_byte_count"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingSeverity(d []interface{}) edpt.Cgnv6TemplateLoggingSeverity {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingSeverity
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SeverityString = in["severity_string"].(string)
		ret.SeverityVal = in["severity_val"].(int)
	}
	return ret
}

func getObjectCgnv6TemplateLoggingSourceAddress124(d []interface{}) edpt.Cgnv6TemplateLoggingSourceAddress124 {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingSourceAddress124
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		//omit uuid
	}
	return ret
}

func getObjectCgnv6TemplateLoggingSourcePort(d []interface{}) edpt.Cgnv6TemplateLoggingSourcePort {

	count1 := len(d)
	var ret edpt.Cgnv6TemplateLoggingSourcePort
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SourcePortNum = in["source_port_num"].(int)
		ret.Any = in["any"].(int)
	}
	return ret
}

func dataToEndpointCgnv6TemplateLogging(d *schema.ResourceData) edpt.Cgnv6TemplateLogging {
	var ret edpt.Cgnv6TemplateLogging
	ret.Inst.BatchedLoggingDisable = d.Get("batched_logging_disable").(int)
	ret.Inst.Custom = getObjectCgnv6TemplateLoggingCustom(d.Get("custom").([]interface{}))
	ret.Inst.DisableLogByDestination = getObjectCgnv6TemplateLoggingDisableLogByDestination115(d.Get("disable_log_by_destination").([]interface{}))
	ret.Inst.Facility = d.Get("facility").(string)
	ret.Inst.Format = d.Get("format").(string)
	ret.Inst.IncludeDestination = d.Get("include_destination").(int)
	ret.Inst.IncludeHttp = getObjectCgnv6TemplateLoggingIncludeHttp(d.Get("include_http").([]interface{}))
	ret.Inst.IncludeInsideUserMac = d.Get("include_inside_user_mac").(int)
	ret.Inst.IncludePartitionName = d.Get("include_partition_name").(int)
	ret.Inst.IncludePortBlockAccount = d.Get("include_port_block_account").(int)
	ret.Inst.IncludeRadiusAttribute = getObjectCgnv6TemplateLoggingIncludeRadiusAttribute(d.Get("include_radius_attribute").([]interface{}))
	ret.Inst.IncludeSessionByteCount = d.Get("include_session_byte_count").(int)
	ret.Inst.Log = getObjectCgnv6TemplateLoggingLog(d.Get("log").([]interface{}))
	ret.Inst.LogReceiver = getObjectCgnv6TemplateLoggingLogReceiver(d.Get("log_receiver").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Resolution = d.Get("resolution").(string)
	ret.Inst.RfcCustom = getObjectCgnv6TemplateLoggingRfcCustom(d.Get("rfc_custom").([]interface{}))
	ret.Inst.Rule = getObjectCgnv6TemplateLoggingRule(d.Get("rule").([]interface{}))
	ret.Inst.ServiceGroup = d.Get("service_group").(string)
	ret.Inst.Severity = getObjectCgnv6TemplateLoggingSeverity(d.Get("severity").([]interface{}))
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.SourceAddress = getObjectCgnv6TemplateLoggingSourceAddress124(d.Get("source_address").([]interface{}))
	ret.Inst.SourcePort = getObjectCgnv6TemplateLoggingSourcePort(d.Get("source_port").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
