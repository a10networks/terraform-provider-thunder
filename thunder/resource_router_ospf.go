package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ospf`: Open Shortest Path First (OSPF)\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterOspfCreate,
		UpdateContext: resourceRouterOspfUpdate,
		ReadContext:   resourceRouterOspfRead,
		DeleteContext: resourceRouterOspfDelete,

		Schema: map[string]*schema.Schema{
			"area_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_ipv4": {
							Type: schema.TypeString, Required: true, Description: "OSPF area ID in IP address format",
						},
						"area_num": {
							Type: schema.TypeInt, Required: true, Description: "OSPF area ID as a decimal value",
						},
						"auth_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
									},
									"message_digest": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use message-digest authentication",
									},
								},
							},
						},
						"filter_lists": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filter_list": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Filter networks between OSPF areas",
									},
									"acl_name": {
										Type: schema.TypeString, Optional: true, Description: "Filter networks by access-list (Name of an access-list)",
									},
									"acl_direction": {
										Type: schema.TypeString, Optional: true, Description: "'in': Filter networks sent to this area; 'out': Filter networks sent from this area;",
									},
									"plist_name": {
										Type: schema.TypeString, Optional: true, Description: "Filter networks by prefix-list (Name of an IP prefix-list)",
									},
									"plist_direction": {
										Type: schema.TypeString, Optional: true, Description: "'in': Filter networks sent to this area; 'out': Filter networks sent from this area;",
									},
								},
							},
						},
						"nssa_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"nssa": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a NSSA area",
									},
									"no_redistribution": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "No redistribution into this NSSA area",
									},
									"no_summary": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not send summary LSA into NSSA",
									},
									"translator_role": {
										Type: schema.TypeString, Optional: true, Default: "candidate", Description: "'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;",
									},
									"default_information_originate": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Originate Type 7 default into NSSA area",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "OSPF default metric (OSPF metric)",
									},
									"metric_type": {
										Type: schema.TypeInt, Optional: true, Default: 2, Description: "OSPF metric type (OSPF metric type for default routes)",
									},
								},
							},
						},
						"default_cost": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)",
						},
						"range_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_range_prefix": {
										Type: schema.TypeString, Optional: true, Description: "Area range for IPv4 prefix",
									},
									"option": {
										Type: schema.TypeString, Optional: true, Default: "advertise", Description: "'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;",
									},
								},
							},
						},
						"shortcut": {
							Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;",
						},
						"stub_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stub": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure OSPF area as stub",
									},
									"no_summary": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not inject inter-area routes into area",
									},
								},
							},
						},
						"virtual_link_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_link_ip_addr": {
										Type: schema.TypeString, Optional: true, Description: "ID (IP addr) associated with virtual link neighbor",
									},
									"bfd": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
									},
									"hello_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Hello packet interval (Seconds)",
									},
									"dead_interval": {
										Type: schema.TypeInt, Optional: true, Description: "Dead router detection time (Seconds)",
									},
									"retransmit_interval": {
										Type: schema.TypeInt, Optional: true, Description: "LSA retransmit interval (Seconds)",
									},
									"transmit_delay": {
										Type: schema.TypeInt, Optional: true, Default: 1, Description: "LSA transmission delay (Seconds)",
									},
									"virtual_link_authentication": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable authentication",
									},
									"virtual_link_auth_type": {
										Type: schema.TypeString, Optional: true, Description: "'message-digest': Use message-digest authentication; 'null': Use null authentication;",
									},
									"authentication_key": {
										Type: schema.TypeString, Optional: true, Description: "Set authentication key (Authentication key (8 chars))",
									},
									"message_digest_key": {
										Type: schema.TypeInt, Optional: true, Description: "Set message digest key (Key ID)",
									},
									"md5": {
										Type: schema.TypeString, Optional: true, Description: "Use MD5 algorithm (Authentication key (16 chars))",
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
			"auto_cost_reference_bandwidth": {
				Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Use reference bandwidth method to assign OSPF cost (The reference bandwidth in terms of Mbits per second)",
			},
			"bfd_all_interfaces": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BFD on all interfaces",
			},
			"default_information": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"originate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Distribute a default route",
						},
						"always": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always advertise default route",
						},
						"metric": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type": {
							Type: schema.TypeInt, Optional: true, Default: 2, Description: "OSPF metric type for default routes",
						},
						"route_map": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"default_metric": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Set metric of redistributed routes (Default metric)",
			},
			"distance": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance_value": {
							Type: schema.TypeInt, Optional: true, Default: 110, Description: "OSPF Administrative distance",
						},
						"distance_ospf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"distance_external": {
										Type: schema.TypeInt, Optional: true, Description: "External routes (Distance for external routes)",
									},
									"distance_inter_area": {
										Type: schema.TypeInt, Optional: true, Description: "Inter-area routes (Distance for inter-area routes)",
									},
									"distance_intra_area": {
										Type: schema.TypeInt, Optional: true, Description: "Intra-area routes (Distance for intra-area routes)",
									},
								},
							},
						},
					},
				},
			},
			"distribute_internal_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"di_type": {
							Type: schema.TypeString, Optional: true, Description: "'lw4o6': LW4O6 Prefix; 'floating-ip': Floating IP; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'vip': Only not flagged Virtual IP (VIP); 'vip-only-flagged': Selected Virtual IP (VIP);",
						},
						"di_area_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "OSPF area ID as a IP address format",
						},
						"di_area_num": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
						},
						"di_cost": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cost of route",
						},
					},
				},
			},
			"distribute_lists": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "Access-list name",
						},
						"direction": {
							Type: schema.TypeString, Optional: true, Description: "'in': Filter incoming routing updates; 'out': Filter outgoing routing updates;",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'lw4o6': LW4O6 Prefix; 'ip-nat': IP NAT; 'ip-nat-list': IP NAT list; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'ospf': Open Shortest Path First (OSPF); 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
						},
						"ospf_id": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF process ID",
						},
						"option": {
							Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
						},
					},
				},
			},
			"extern_lsa_equivalence_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "external LSA equivalance check",
			},
			"ha_standby_extra_cost": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extra_cost": {
							Type: schema.TypeInt, Optional: true, Description: "The extra cost value",
						},
						"group": {
							Type: schema.TypeInt, Optional: true, Description: "Group (Group ID)",
						},
					},
				},
			},
			"host_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_address": {
							Type: schema.TypeString, Optional: true, Description: "Host address",
						},
						"area_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"area_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "OSPF area ID in IP address format",
									},
									"area_num": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
									},
									"cost": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cost of host",
									},
								},
							},
						},
					},
				},
			},
			"log_adjacency_changes_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Description: "'detail': Log changes in adjacency state; 'disable': Disable logging;",
						},
					},
				},
			},
			"max_concurrent_dd": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Maximum number allowed to process DD concurrently (Number of DD process)",
			},
			"maximum_area": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of non-backbone areas (OSPF area limit)",
			},
			"neighbor_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type: schema.TypeString, Optional: true, Description: "Neighbor address",
						},
						"cost": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF cost for point-to-multipoint neighbor (Metric)",
						},
						"poll_interval": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF dead-router polling interval (Seconds)",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF priority of non-broadcast neighbor",
						},
					},
				},
			},
			"network_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_ipv4": {
							Type: schema.TypeString, Optional: true, Description: "Network number",
						},
						"network_ipv4_mask": {
							Type: schema.TypeString, Optional: true, Description: "OSPF wild card bits",
						},
						"network_ipv4_cidr": {
							Type: schema.TypeString, Optional: true, Description: "OSPF network prefix",
						},
						"network_area": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_area_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "OSPF area ID in IP address format",
									},
									"network_area_num": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
									},
									"instance_value": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Instance ID",
									},
								},
							},
						},
					},
				},
			},
			"ospf_1": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"abr_type": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"option": {
										Type: schema.TypeString, Optional: true, Default: "cisco", Description: "'cisco': Alternative ABR, Cisco implementation (RFC3509); 'ibm': Alternative ABR, IBM implementation (RFC3509); 'shortcut': Shortcut ABR; 'standard': Standard behavior (RFC2328);",
									},
								},
							},
						},
					},
				},
			},
			"overflow": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"database": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count1": {
										Type: schema.TypeInt, Optional: true, Default: 4294967294, Description: "Maximum number of LSAs",
									},
									"limit": {
										Type: schema.TypeString, Optional: true, Default: "hard", Description: "'hard': Hard limit: Instance will be shutdown if exceeded; 'soft': Soft limit: Warning will be given if exceeded;",
									},
									"db_external": {
										Type: schema.TypeInt, Optional: true, Default: 2147483647, Description: "Maximum number of LSAs",
									},
									"recovery_time": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Time to recover (0 not recover)",
									},
								},
							},
						},
					},
				},
			},
			"passive_interface": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"loopback_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"loopback": {
										Type: schema.TypeInt, Optional: true, Description: "Loopback interface (Port number)",
									},
									"loopback_address": {
										Type: schema.TypeString, Optional: true, Description: "Address of Interface",
									},
								},
							},
						},
						"trunk_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trunk": {
										Type: schema.TypeInt, Optional: true, Description: "Trunk interface (Trunk interface number)",
									},
									"trunk_address": {
										Type: schema.TypeString, Optional: true, Description: "Address of Interface",
									},
								},
							},
						},
						"ve_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve": {
										Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet interface (Virtual ethernet interface number)",
									},
									"ve_address": {
										Type: schema.TypeString, Optional: true, Description: "Address of Interface",
									},
								},
							},
						},
						"tunnel_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Tunnel interface (Tunnel interface number)",
									},
									"tunnel_address": {
										Type: schema.TypeString, Optional: true, Description: "Address of Interface",
									},
								},
							},
						},
						"lif_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lif": {
										Type: schema.TypeString, Optional: true, Description: "Logical interface (Lif interface name)",
									},
									"lif_address": {
										Type: schema.TypeString, Optional: true, Description: "Address of Interface",
									},
								},
							},
						},
						"eth_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ethernet": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet interface (Port number)",
									},
									"eth_address": {
										Type: schema.TypeString, Optional: true, Description: "Address of Interface",
									},
								},
							},
						},
					},
				},
			},
			"process_id": {
				Type: schema.TypeInt, Required: true, Description: "OSPF process ID",
			},
			"redistribute": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"redist_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type: schema.TypeString, Optional: true, Description: "'bgp': Border Gateway Protocol (BGP); 'connected': Connected; 'floating-ip': Floating IP; 'ip-nat-list': IP NAT list; 'lw4o6': LW4O6 Prefix; 'nat-map': NAT MAP Prefix; 'static-nat': Static NAT; 'isis': ISO IS-IS; 'rip': Routing Information Protocol (RIP); 'static': Static routes;",
									},
									"metric": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
									},
									"metric_type": {
										Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
									},
									"route_map": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
									"tag": {
										Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
									},
								},
							},
						},
						"ospf_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Open Shortest Path First (OSPF)",
									},
									"process_id": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF process ID",
									},
									"metric_ospf": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
									},
									"metric_type_ospf": {
										Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
									},
									"route_map_ospf": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
									"tag_ospf": {
										Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
									},
								},
							},
						},
						"ip_nat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP-NAT",
						},
						"metric_ip_nat": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
						},
						"metric_type_ip_nat": {
							Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
						},
						"route_map_ip_nat": {
							Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
						},
						"tag_ip_nat": {
							Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
						},
						"ip_nat_floating_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_nat_prefix": {
										Type: schema.TypeString, Optional: true, Description: "Address",
									},
									"ip_nat_floating_ip_forward": {
										Type: schema.TypeString, Optional: true, Description: "Floating-IP as forward address",
									},
								},
							},
						},
						"vip_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type_vip": {
										Type: schema.TypeString, Optional: true, Description: "'only-flagged': Selected Virtual IP (VIP); 'only-not-flagged': Only not flagged;",
									},
									"metric_vip": {
										Type: schema.TypeInt, Optional: true, Description: "OSPF default metric (OSPF metric)",
									},
									"metric_type_vip": {
										Type: schema.TypeString, Optional: true, Default: "2", Description: "'1': Set OSPF External Type 1 metrics; '2': Set OSPF External Type 2 metrics;",
									},
									"route_map_vip": {
										Type: schema.TypeString, Optional: true, Description: "Route map reference (Pointer to route-map entries)",
									},
									"tag_vip": {
										Type: schema.TypeInt, Optional: true, Description: "Set tag for routes redistributed into OSPF (32-bit tag value)",
									},
								},
							},
						},
						"vip_floating_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vip_address": {
										Type: schema.TypeString, Optional: true, Description: "Address",
									},
									"vip_floating_ip_forward": {
										Type: schema.TypeString, Optional: true, Description: "Floating-IP as forward address",
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
			"rfc1583_compatible": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Compatible with RFC 1583",
			},
			"router_id": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "OSPF router-id in IPv4 address format",
						},
					},
				},
			},
			"summary_address_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"summary_address": {
							Type: schema.TypeString, Optional: true, Description: "Configure IP address summaries (Summary prefix)",
						},
						"not_advertise": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Suppress routes that match the prefix",
						},
						"tag": {
							Type: schema.TypeInt, Optional: true, Description: "Set tag (32-bit tag value)",
						},
					},
				},
			},
			"timers": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spf": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exp": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"min_delay": {
													Type: schema.TypeInt, Optional: true, Default: 500, Description: "Minimum delay between receiving a change to SPF calculation in milliseconds",
												},
												"max_delay": {
													Type: schema.TypeInt, Optional: true, Default: 50000, Description: "Maximum delay between receiving a change to SPF calculation in milliseconds",
												},
											},
										},
									},
								},
							},
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
func resourceRouterOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterOspfRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterOspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterOspf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterOspfAreaList(d []interface{}) []edpt.RouterOspfAreaList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaList
		oi.AreaIpv4 = in["area_ipv4"].(string)
		oi.AreaNum = in["area_num"].(int)
		oi.AuthCfg = getObjectRouterOspfAreaListAuthCfg(in["auth_cfg"].([]interface{}))
		oi.FilterLists = getSliceRouterOspfAreaListFilterLists(in["filter_lists"].([]interface{}))
		oi.NssaCfg = getObjectRouterOspfAreaListNssaCfg(in["nssa_cfg"].([]interface{}))
		oi.DefaultCost = in["default_cost"].(int)
		oi.RangeList = getSliceRouterOspfAreaListRangeList(in["range_list"].([]interface{}))
		oi.Shortcut = in["shortcut"].(string)
		oi.StubCfg = getObjectRouterOspfAreaListStubCfg(in["stub_cfg"].([]interface{}))
		oi.VirtualLinkList = getSliceRouterOspfAreaListVirtualLinkList(in["virtual_link_list"].([]interface{}))
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfAreaListAuthCfg(d []interface{}) edpt.RouterOspfAreaListAuthCfg {

	count1 := len(d)
	var ret edpt.RouterOspfAreaListAuthCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Authentication = in["authentication"].(int)
		ret.MessageDigest = in["message_digest"].(int)
	}
	return ret
}

func getSliceRouterOspfAreaListFilterLists(d []interface{}) []edpt.RouterOspfAreaListFilterLists {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaListFilterLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaListFilterLists
		oi.FilterList = in["filter_list"].(int)
		oi.AclName = in["acl_name"].(string)
		oi.AclDirection = in["acl_direction"].(string)
		oi.PlistName = in["plist_name"].(string)
		oi.PlistDirection = in["plist_direction"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfAreaListNssaCfg(d []interface{}) edpt.RouterOspfAreaListNssaCfg {

	count1 := len(d)
	var ret edpt.RouterOspfAreaListNssaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nssa = in["nssa"].(int)
		ret.NoRedistribution = in["no_redistribution"].(int)
		ret.NoSummary = in["no_summary"].(int)
		ret.TranslatorRole = in["translator_role"].(string)
		ret.DefaultInformationOriginate = in["default_information_originate"].(int)
		ret.Metric = in["metric"].(int)
		ret.MetricType = in["metric_type"].(int)
	}
	return ret
}

func getSliceRouterOspfAreaListRangeList(d []interface{}) []edpt.RouterOspfAreaListRangeList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaListRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaListRangeList
		oi.AreaRangePrefix = in["area_range_prefix"].(string)
		oi.Option = in["option"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfAreaListStubCfg(d []interface{}) edpt.RouterOspfAreaListStubCfg {

	count1 := len(d)
	var ret edpt.RouterOspfAreaListStubCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stub = in["stub"].(int)
		ret.NoSummary = in["no_summary"].(int)
	}
	return ret
}

func getSliceRouterOspfAreaListVirtualLinkList(d []interface{}) []edpt.RouterOspfAreaListVirtualLinkList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfAreaListVirtualLinkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfAreaListVirtualLinkList
		oi.VirtualLinkIpAddr = in["virtual_link_ip_addr"].(string)
		oi.Bfd = in["bfd"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.VirtualLinkAuthentication = in["virtual_link_authentication"].(int)
		oi.VirtualLinkAuthType = in["virtual_link_auth_type"].(string)
		oi.AuthenticationKey = in["authentication_key"].(string)
		oi.MessageDigestKey = in["message_digest_key"].(int)
		oi.Md5 = in["md5"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfDefaultInformation1297(d []interface{}) edpt.RouterOspfDefaultInformation1297 {

	count1 := len(d)
	var ret edpt.RouterOspfDefaultInformation1297
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Originate = in["originate"].(int)
		ret.Always = in["always"].(int)
		ret.Metric = in["metric"].(int)
		ret.MetricType = in["metric_type"].(int)
		ret.RouteMap = in["route_map"].(string)
		//omit uuid
	}
	return ret
}

func getObjectRouterOspfDistance(d []interface{}) edpt.RouterOspfDistance {

	count1 := len(d)
	var ret edpt.RouterOspfDistance
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DistanceValue = in["distance_value"].(int)
		ret.DistanceOspf = getObjectRouterOspfDistanceDistanceOspf(in["distance_ospf"].([]interface{}))
	}
	return ret
}

func getObjectRouterOspfDistanceDistanceOspf(d []interface{}) edpt.RouterOspfDistanceDistanceOspf {

	count1 := len(d)
	var ret edpt.RouterOspfDistanceDistanceOspf
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DistanceExternal = in["distance_external"].(int)
		ret.DistanceInterArea = in["distance_inter_area"].(int)
		ret.DistanceIntraArea = in["distance_intra_area"].(int)
	}
	return ret
}

func getSliceRouterOspfDistributeInternalList(d []interface{}) []edpt.RouterOspfDistributeInternalList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfDistributeInternalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfDistributeInternalList
		oi.DiType = in["di_type"].(string)
		oi.DiAreaIpv4 = in["di_area_ipv4"].(string)
		oi.DiAreaNum = in["di_area_num"].(int)
		oi.DiCost = in["di_cost"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfDistributeLists(d []interface{}) []edpt.RouterOspfDistributeLists {

	count1 := len(d)
	ret := make([]edpt.RouterOspfDistributeLists, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfDistributeLists
		oi.Value = in["value"].(string)
		oi.Direction = in["direction"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.OspfId = in["ospf_id"].(int)
		oi.Option = in["option"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfHaStandbyExtraCost(d []interface{}) []edpt.RouterOspfHaStandbyExtraCost {

	count1 := len(d)
	ret := make([]edpt.RouterOspfHaStandbyExtraCost, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfHaStandbyExtraCost
		oi.ExtraCost = in["extra_cost"].(int)
		oi.Group = in["group"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfHostList(d []interface{}) []edpt.RouterOspfHostList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfHostList
		oi.HostAddress = in["host_address"].(string)
		oi.AreaCfg = getObjectRouterOspfHostListAreaCfg(in["area_cfg"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfHostListAreaCfg(d []interface{}) edpt.RouterOspfHostListAreaCfg {

	count1 := len(d)
	var ret edpt.RouterOspfHostListAreaCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AreaIpv4 = in["area_ipv4"].(string)
		ret.AreaNum = in["area_num"].(int)
		ret.Cost = in["cost"].(int)
	}
	return ret
}

func getObjectRouterOspfLogAdjacencyChangesCfg(d []interface{}) edpt.RouterOspfLogAdjacencyChangesCfg {

	count1 := len(d)
	var ret edpt.RouterOspfLogAdjacencyChangesCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func getSliceRouterOspfNeighborList(d []interface{}) []edpt.RouterOspfNeighborList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfNeighborList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfNeighborList
		oi.Address = in["address"].(string)
		oi.Cost = in["cost"].(int)
		oi.PollInterval = in["poll_interval"].(int)
		oi.Priority = in["priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfNetworkList(d []interface{}) []edpt.RouterOspfNetworkList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfNetworkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfNetworkList
		oi.NetworkIpv4 = in["network_ipv4"].(string)
		oi.NetworkIpv4Mask = in["network_ipv4_mask"].(string)
		oi.NetworkIpv4Cidr = in["network_ipv4_cidr"].(string)
		oi.NetworkArea = getObjectRouterOspfNetworkListNetworkArea(in["network_area"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfNetworkListNetworkArea(d []interface{}) edpt.RouterOspfNetworkListNetworkArea {

	count1 := len(d)
	var ret edpt.RouterOspfNetworkListNetworkArea
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetworkAreaIpv4 = in["network_area_ipv4"].(string)
		ret.NetworkAreaNum = in["network_area_num"].(int)
		ret.InstanceValue = in["instance_value"].(int)
	}
	return ret
}

func getObjectRouterOspfOspf1(d []interface{}) edpt.RouterOspfOspf1 {

	count1 := len(d)
	var ret edpt.RouterOspfOspf1
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AbrType = getObjectRouterOspfOspf1AbrType(in["abr_type"].([]interface{}))
	}
	return ret
}

func getObjectRouterOspfOspf1AbrType(d []interface{}) edpt.RouterOspfOspf1AbrType {

	count1 := len(d)
	var ret edpt.RouterOspfOspf1AbrType
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Option = in["option"].(string)
	}
	return ret
}

func getObjectRouterOspfOverflow(d []interface{}) edpt.RouterOspfOverflow {

	count1 := len(d)
	var ret edpt.RouterOspfOverflow
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Database = getObjectRouterOspfOverflowDatabase(in["database"].([]interface{}))
	}
	return ret
}

func getObjectRouterOspfOverflowDatabase(d []interface{}) edpt.RouterOspfOverflowDatabase {

	count1 := len(d)
	var ret edpt.RouterOspfOverflowDatabase
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Count1 = in["count1"].(int)
		ret.Limit = in["limit"].(string)
		ret.DbExternal = in["db_external"].(int)
		ret.RecoveryTime = in["recovery_time"].(int)
	}
	return ret
}

func getObjectRouterOspfPassiveInterface(d []interface{}) edpt.RouterOspfPassiveInterface {

	count1 := len(d)
	var ret edpt.RouterOspfPassiveInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LoopbackCfg = getSliceRouterOspfPassiveInterfaceLoopbackCfg(in["loopback_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceRouterOspfPassiveInterfaceTrunkCfg(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceRouterOspfPassiveInterfaceVeCfg(in["ve_cfg"].([]interface{}))
		ret.TunnelCfg = getSliceRouterOspfPassiveInterfaceTunnelCfg(in["tunnel_cfg"].([]interface{}))
		ret.LifCfg = getSliceRouterOspfPassiveInterfaceLifCfg(in["lif_cfg"].([]interface{}))
		ret.EthCfg = getSliceRouterOspfPassiveInterfaceEthCfg(in["eth_cfg"].([]interface{}))
	}
	return ret
}

func getSliceRouterOspfPassiveInterfaceLoopbackCfg(d []interface{}) []edpt.RouterOspfPassiveInterfaceLoopbackCfg {

	count1 := len(d)
	ret := make([]edpt.RouterOspfPassiveInterfaceLoopbackCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfPassiveInterfaceLoopbackCfg
		oi.Loopback = in["loopback"].(int)
		oi.LoopbackAddress = in["loopback_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfPassiveInterfaceTrunkCfg(d []interface{}) []edpt.RouterOspfPassiveInterfaceTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.RouterOspfPassiveInterfaceTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfPassiveInterfaceTrunkCfg
		oi.Trunk = in["trunk"].(int)
		oi.TrunkAddress = in["trunk_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfPassiveInterfaceVeCfg(d []interface{}) []edpt.RouterOspfPassiveInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.RouterOspfPassiveInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfPassiveInterfaceVeCfg
		oi.Ve = in["ve"].(int)
		oi.VeAddress = in["ve_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfPassiveInterfaceTunnelCfg(d []interface{}) []edpt.RouterOspfPassiveInterfaceTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.RouterOspfPassiveInterfaceTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfPassiveInterfaceTunnelCfg
		oi.Tunnel = in["tunnel"].(int)
		oi.TunnelAddress = in["tunnel_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfPassiveInterfaceLifCfg(d []interface{}) []edpt.RouterOspfPassiveInterfaceLifCfg {

	count1 := len(d)
	ret := make([]edpt.RouterOspfPassiveInterfaceLifCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfPassiveInterfaceLifCfg
		oi.Lif = in["lif"].(string)
		oi.LifAddress = in["lif_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfPassiveInterfaceEthCfg(d []interface{}) []edpt.RouterOspfPassiveInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.RouterOspfPassiveInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfPassiveInterfaceEthCfg
		oi.Ethernet = in["ethernet"].(int)
		oi.EthAddress = in["eth_address"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfRedistribute1298(d []interface{}) edpt.RouterOspfRedistribute1298 {

	count1 := len(d)
	var ret edpt.RouterOspfRedistribute1298
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RedistList = getSliceRouterOspfRedistributeRedistList1299(in["redist_list"].([]interface{}))
		ret.OspfList = getSliceRouterOspfRedistributeOspfList1300(in["ospf_list"].([]interface{}))
		ret.IpNat = in["ip_nat"].(int)
		ret.MetricIpNat = in["metric_ip_nat"].(int)
		ret.MetricTypeIpNat = in["metric_type_ip_nat"].(string)
		ret.RouteMapIpNat = in["route_map_ip_nat"].(string)
		ret.TagIpNat = in["tag_ip_nat"].(int)
		ret.IpNatFloatingList = getSliceRouterOspfRedistributeIpNatFloatingList1301(in["ip_nat_floating_list"].([]interface{}))
		ret.VipList = getSliceRouterOspfRedistributeVipList1302(in["vip_list"].([]interface{}))
		ret.VipFloatingList = getSliceRouterOspfRedistributeVipFloatingList1303(in["vip_floating_list"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceRouterOspfRedistributeRedistList1299(d []interface{}) []edpt.RouterOspfRedistributeRedistList1299 {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeRedistList1299, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeRedistList1299
		oi.Type = in["type"].(string)
		oi.Metric = in["metric"].(int)
		oi.MetricType = in["metric_type"].(string)
		oi.RouteMap = in["route_map"].(string)
		oi.Tag = in["tag"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeOspfList1300(d []interface{}) []edpt.RouterOspfRedistributeOspfList1300 {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeOspfList1300, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeOspfList1300
		oi.Ospf = in["ospf"].(int)
		oi.ProcessId = in["process_id"].(int)
		oi.MetricOspf = in["metric_ospf"].(int)
		oi.MetricTypeOspf = in["metric_type_ospf"].(string)
		oi.RouteMapOspf = in["route_map_ospf"].(string)
		oi.TagOspf = in["tag_ospf"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeIpNatFloatingList1301(d []interface{}) []edpt.RouterOspfRedistributeIpNatFloatingList1301 {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeIpNatFloatingList1301, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeIpNatFloatingList1301
		oi.IpNatPrefix = in["ip_nat_prefix"].(string)
		oi.IpNatFloatingIpForward = in["ip_nat_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeVipList1302(d []interface{}) []edpt.RouterOspfRedistributeVipList1302 {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeVipList1302, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeVipList1302
		oi.TypeVip = in["type_vip"].(string)
		oi.MetricVip = in["metric_vip"].(int)
		oi.MetricTypeVip = in["metric_type_vip"].(string)
		oi.RouteMapVip = in["route_map_vip"].(string)
		oi.TagVip = in["tag_vip"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterOspfRedistributeVipFloatingList1303(d []interface{}) []edpt.RouterOspfRedistributeVipFloatingList1303 {

	count1 := len(d)
	ret := make([]edpt.RouterOspfRedistributeVipFloatingList1303, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfRedistributeVipFloatingList1303
		oi.VipAddress = in["vip_address"].(string)
		oi.VipFloatingIpForward = in["vip_floating_ip_forward"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfRouterId(d []interface{}) edpt.RouterOspfRouterId {

	count1 := len(d)
	var ret edpt.RouterOspfRouterId
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(string)
	}
	return ret
}

func getSliceRouterOspfSummaryAddressList(d []interface{}) []edpt.RouterOspfSummaryAddressList {

	count1 := len(d)
	ret := make([]edpt.RouterOspfSummaryAddressList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterOspfSummaryAddressList
		oi.SummaryAddress = in["summary_address"].(string)
		oi.NotAdvertise = in["not_advertise"].(int)
		oi.Tag = in["tag"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRouterOspfTimers(d []interface{}) edpt.RouterOspfTimers {

	count1 := len(d)
	var ret edpt.RouterOspfTimers
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Spf = getObjectRouterOspfTimersSpf(in["spf"].([]interface{}))
	}
	return ret
}

func getObjectRouterOspfTimersSpf(d []interface{}) edpt.RouterOspfTimersSpf {

	count1 := len(d)
	var ret edpt.RouterOspfTimersSpf
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Exp = getObjectRouterOspfTimersSpfExp(in["exp"].([]interface{}))
	}
	return ret
}

func getObjectRouterOspfTimersSpfExp(d []interface{}) edpt.RouterOspfTimersSpfExp {

	count1 := len(d)
	var ret edpt.RouterOspfTimersSpfExp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinDelay = in["min_delay"].(int)
		ret.MaxDelay = in["max_delay"].(int)
	}
	return ret
}

func dataToEndpointRouterOspf(d *schema.ResourceData) edpt.RouterOspf {
	var ret edpt.RouterOspf
	ret.Inst.AreaList = getSliceRouterOspfAreaList(d.Get("area_list").([]interface{}))
	ret.Inst.AutoCostReferenceBandwidth = d.Get("auto_cost_reference_bandwidth").(int)
	ret.Inst.BfdAllInterfaces = d.Get("bfd_all_interfaces").(int)
	ret.Inst.DefaultInformation = getObjectRouterOspfDefaultInformation1297(d.Get("default_information").([]interface{}))
	ret.Inst.DefaultMetric = d.Get("default_metric").(int)
	ret.Inst.Distance = getObjectRouterOspfDistance(d.Get("distance").([]interface{}))
	ret.Inst.DistributeInternalList = getSliceRouterOspfDistributeInternalList(d.Get("distribute_internal_list").([]interface{}))
	ret.Inst.DistributeLists = getSliceRouterOspfDistributeLists(d.Get("distribute_lists").([]interface{}))
	ret.Inst.ExternLsaEquivalenceCheck = d.Get("extern_lsa_equivalence_check").(int)
	ret.Inst.HaStandbyExtraCost = getSliceRouterOspfHaStandbyExtraCost(d.Get("ha_standby_extra_cost").([]interface{}))
	ret.Inst.HostList = getSliceRouterOspfHostList(d.Get("host_list").([]interface{}))
	ret.Inst.LogAdjacencyChangesCfg = getObjectRouterOspfLogAdjacencyChangesCfg(d.Get("log_adjacency_changes_cfg").([]interface{}))
	ret.Inst.MaxConcurrentDd = d.Get("max_concurrent_dd").(int)
	ret.Inst.MaximumArea = d.Get("maximum_area").(int)
	ret.Inst.NeighborList = getSliceRouterOspfNeighborList(d.Get("neighbor_list").([]interface{}))
	ret.Inst.NetworkList = getSliceRouterOspfNetworkList(d.Get("network_list").([]interface{}))
	ret.Inst.Ospf1 = getObjectRouterOspfOspf1(d.Get("ospf_1").([]interface{}))
	ret.Inst.Overflow = getObjectRouterOspfOverflow(d.Get("overflow").([]interface{}))
	ret.Inst.PassiveInterface = getObjectRouterOspfPassiveInterface(d.Get("passive_interface").([]interface{}))
	ret.Inst.ProcessId = d.Get("process_id").(int)
	ret.Inst.Redistribute = getObjectRouterOspfRedistribute1298(d.Get("redistribute").([]interface{}))
	ret.Inst.Rfc1583Compatible = d.Get("rfc1583_compatible").(int)
	ret.Inst.RouterId = getObjectRouterOspfRouterId(d.Get("router_id").([]interface{}))
	ret.Inst.SummaryAddressList = getSliceRouterOspfSummaryAddressList(d.Get("summary_address_list").([]interface{}))
	ret.Inst.Timers = getObjectRouterOspfTimers(d.Get("timers").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
