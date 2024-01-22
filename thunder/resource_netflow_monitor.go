package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetflowMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_netflow_monitor`: Configure NetFlow Monitor\n\n__PLACEHOLDER__",
		CreateContext: resourceNetflowMonitorCreate,
		UpdateContext: resourceNetflowMonitorUpdate,
		ReadContext:   resourceNetflowMonitorRead,
		DeleteContext: resourceNetflowMonitorDelete,

		Schema: map[string]*schema.Schema{
			"counter_polling_interval": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Configure the interval to export global counters (Number of seconds: default is 60)",
			},
			"custom_record": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_cfg": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"event": {
										Type: schema.TypeString, Optional: true, Description: "'sesn-event-nat44-creation': Export NAT44 session creation events; 'sesn-event-nat44-deletion': Export NAT44 session deletion events; 'sesn-event-nat64-creation': Export NAT64 session creation events; 'sesn-event-nat64-deletion': Export NAT64 session deletion events; 'sesn-event-dslite-creation': Export Dslite session creation events; 'sesn-event-dslite-deletion': Export Dslite session deletion events; 'sesn-event-fw4-creation': Export FW4 session creation events; 'sesn-event-fw4-deletion': Export FW4 session deletion events; 'sesn-event-fw6-creation': Export FW6 session creation events; 'sesn-event-fw6-deletion': Export FW6 session deletion events; 'deny-reset-event-fw4': Export FW4 Deny Reset events; 'deny-reset-event-fw6': Export FW6 Deny Reset events; 'port-mapping-nat44-creation': Export NAT44 Port Mapping Creation Event; 'port-mapping-nat44-deletion': Export NAT44 Port Mapping Deletion Event; 'port-mapping-nat64-creation': Export NAT64 Port Mapping Creation Event; 'port-mapping-nat64-deletion': Export NAT64 Port Mapping Deletion Event; 'port-mapping-dslite-creation': Export Dslite Port Mapping Creation Event; 'port-mapping-dslite-deletion': Export Dslite Port Mapping Deletion Event; 'port-batch-nat44-creation': Export NAT44 Port Batch Creation Event; 'port-batch-nat44-deletion': Export NAT44 Port Batch Deletion Event; 'port-batch-nat64-creation': Export NAT64 Port Batch Creation Event; 'port-batch-nat64-deletion': Export NAT64 Port Batch Deletion Event; 'port-batch-dslite-creation': Export Dslite Port Batch Creation Event; 'port-batch-dslite-deletion': Export Dslite Port Batch Deletion Event; 'port-batch-v2-nat44-creation': Export NAT44 Port Batch v2 Creation Event; 'port-batch-v2-nat44-deletion': Export NAT44 Port Batch v2 Deletion Event; 'port-batch-v2-nat64-creation': Export NAT64 Port Batch v2 Creation Event; 'port-batch-v2-nat64-deletion': Export NAT64 Port Batch v2 Deletion Event; 'port-batch-v2-dslite-creation': Export Dslite Port Batch v2 Creation Event; 'port-batch-v2-dslite-deletion': Export Dslite Port Batch v2 Deletion Event; 'gtp-c-tunnel-event': Export GTP Control Tunnel Creation or Deletion Events; 'gtp-u-tunnel-event': Export GTP User Tunnel Creation or Deletion Events; 'gtp-deny-event': Export GTP Deny events on GTP C/U Tunnels; 'gtp-info-event': Export GTP Info events on GTP C/U Tunnels; 'fw-ddos-entry-creation': Export FW iDDoS Entry Created Record; 'fw-ddos-entry-deletion': Export FW iDDoS Entry Deleted Record; 'fw-session-limit-exceeded': Export FW Session Limit Exceeded Record; 'cgn-ddos-l3-entry-creation': Export CGN iDDoS L3 Entry Creation; 'cgn-ddos-l3-entry-deletion': Export CGN iDDoS L3 Entry Deletion; 'cgn-ddos-l4-entry-creation': Export CGN iDDoS L4 Entry Creation; 'cgn-ddos-l4-entry-deletion': Export CGN iDDoS L4 Entry Deletion; 'gtp-rate-limit-periodic': Export GTP Rate Limit Periodic;",
									},
									"ipfix_template": {
										Type: schema.TypeString, Optional: true, Description: "Custom IPFIX Template",
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
			"destination": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_group": {
							Type: schema.TypeString, Optional: true, Description: "Service-group for load balancing between multiple collector servers",
						},
						"ip_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "IP address of netflow collector",
									},
									"port4": {
										Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Port number, default is 9996",
									},
								},
							},
						},
						"ipv6_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address of netflow collector",
									},
									"port6": {
										Type: schema.TypeInt, Optional: true, Default: 9996, Description: "Port number, default is 9996",
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
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable this netflow monitor",
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
			"flow_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure timeout value to export flow records periodically for long-live session ( Number of minutes: default is 10, 0 means only send flow record when session is deleted)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name of netflow monitor",
			},
			"packet_capture_template": {
				Type: schema.TypeString, Optional: true, Description: "Name of the packet capture template to be bind with this object",
			},
			"protocol": {
				Type: schema.TypeString, Optional: true, Default: "v9", Description: "'v9': Netflow version 9; 'v10': Netflow version 10 (IPFIX);",
			},
			"record": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"netflow_v5": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NetFlow V5 Flow Record Template",
						},
						"netflow_v5_ext": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Extended NetFlow V5 Flow Record Template, supports ipv6",
						},
						"nat44": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT44 Flow Record Template",
						},
						"nat64": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NAT64 Flow Record Template",
						},
						"dslite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DS-Lite Flow Record Template",
						},
						"sesn_event_nat44": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"sesn_event_nat64": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"sesn_event_dslite": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"sesn_event_fw4": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"sesn_event_fw6": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_mapping_nat44": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_mapping_nat64": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_mapping_dslite": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_batch_nat44": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_batch_nat64": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_batch_dslite": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_batch_v2_nat44": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_batch_v2_nat64": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"port_batch_v2_dslite": {
							Type: schema.TypeString, Optional: true, Description: "'both': Export both creation and deletion events; 'creation': Export only creation events; 'deletion': Export only deletion events;",
						},
						"ddos_general_stat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "General DDOS statistics",
						},
						"ddos_http_stat": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP DDOS statistics",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"resend_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 1800, Description: "To set time interval to resend template (number of seconds: default is 1800, 0 means disable template resend based on timeout)",
						},
						"records": {
							Type: schema.TypeInt, Optional: true, Default: 1000, Description: "To resend template once for each number of records (Number of records: default is 1000, 0 means disable template resend based on record-count)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"sample": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ifindex": {
										Type: schema.TypeInt, Required: true, Description: "Ethernet interface number",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"ve_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ve_num": {
										Type: schema.TypeInt, Required: true, Description: "VE interface number",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"nat_pool_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool_name": {
										Type: schema.TypeString, Required: true, Description: "Name of nat pool",
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
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packets-sent': Sent Packets Count; 'bytes-sent': Sent Bytes Count; 'nat44-records-sent': NAT44 Flow Records Sent; 'nat44-records-sent-failure': NAT44 Flow Records Failed; 'nat64-records-sent': NAT64 Flow Records Sent; 'nat64-records-sent-failure': NAT64 Flow Records Failed; 'dslite-records-sent': Dslite Flow Records Sent; 'dslite-records-sent-failure': Dslite Flow Records Failed; 'session-event-nat44-records-sent': Nat44 Session Event Records Sent; 'session-event-nat44-records-sent-failure': Nat44 Session Event Records Failed; 'session-event-nat64-records-sent': Nat64 Session Event Records Sent; 'session-event-nat64-records-sent-failure': Nat64 Session Event Records Falied; 'session-event-dslite-records-sent': Dslite Session Event Records Sent; 'session-event-dslite-records-sent-failure': Dslite Session Event Records Failed; 'session-event-fw4-records-sent': FW4 Session Event Records Sent; 'session-event-fw4-records-sent-failure': FW4 Session Event Records Failed; 'session-event-fw6-records-sent': FW6 Session Event Records Sent; 'session-event-fw6-records-sent-failure': FW6 Session Event Records Failed; 'port-mapping-nat44-records-sent': Port Mapping Nat44 Event Records Sent; 'port-mapping-nat44-records-sent-failure': Port Mapping Nat44 Event Records Failed; 'port-mapping-nat64-records-sent': Port Mapping Nat64 Event Records Sent; 'port-mapping-nat64-records-sent-failure': Port Mapping Nat64 Event Records Failed; 'port-mapping-dslite-records-sent': Port Mapping Dslite Event Records Sent; 'port-mapping-dslite-records-sent-failure': Port Mapping Dslite Event Records failed; 'netflow-v5-records-sent': Netflow v5 Records Sent; 'netflow-v5-records-sent-failure': Netflow v5 Records Failed; 'netflow-v5-ext-records-sent': Netflow v5 Ext Records Sent; 'netflow-v5-ext-records-sent-failure': Netflow v5 Ext Records Failed; 'port-batching-nat44-records-sent': Port Batching Nat44 Records Sent; 'port-batching-nat44-records-sent-failure': Port Batching Nat44 Records Failed; 'port-batching-nat64-records-sent': Port Batching Nat64 Records Sent; 'port-batching-nat64-records-sent-failure': Port Batching Nat64 Records Failed; 'port-batching-dslite-records-sent': Port Batching Dslite Records Sent; 'port-batching-dslite-records-sent-failure': Port Batching Dslite Records Failed; 'port-batching-v2-nat44-records-sent': Port Batching V2 Nat44 Records Sent; 'port-batching-v2-nat44-records-sent-failure': Port Batching V2 Nat44 Records Failed; 'port-batching-v2-nat64-records-sent': Port Batching V2 Nat64 Records Sent; 'port-batching-v2-nat64-records-sent-failure': Port Batching V2 Nat64 Records Failed; 'port-batching-v2-dslite-records-sent': Port Batching V2 Dslite Records Sent; 'port-batching-v2-dslite-records-sent-failure': Port Batching V2 Dslite Records Falied; 'custom-session-event-nat44-creation-records-sent': Custom Nat44 Session Creation Records Sent; 'custom-session-event-nat44-creation-records-sent-failure': Custom Nat44 Session Creation Records Failed; 'custom-session-event-nat64-creation-records-sent': Custom Nat64 Session Creation Records Sent; 'custom-session-event-nat64-creation-records-sent-failure': Custom Nat64 Session Creation Records Failed; 'custom-session-event-dslite-creation-records-sent': Custom Dslite Session Creation Records Sent; 'custom-session-event-dslite-creation-records-sent-failure': Custom Dslite Session Creation Records Failed; 'custom-session-event-nat44-deletion-records-sent': Custom Nat44 Session Deletion Records Sent; 'custom-session-event-nat44-deletion-records-sent-failure': Custom Nat44 Session Deletion Records Failed; 'custom-session-event-nat64-deletion-records-sent': Custom Nat64 Session Deletion Records Sent; 'custom-session-event-nat64-deletion-records-sent-failure': Custom Nat64 Session Deletion Records Failed; 'custom-session-event-dslite-deletion-records-sent': Custom Dslite Session Deletion Records Sent; 'custom-session-event-dslite-deletion-records-sent-failure': Custom Dslite Session Deletion Records Failed; 'custom-session-event-fw4-creation-records-sent': Custom FW4 Session Creation Records Sent; 'custom-session-event-fw4-creation-records-sent-failure': Custom FW4 Session Creation Records Failed; 'custom-session-event-fw6-creation-records-sent': Custom FW6 Session Creation Records Sent; 'custom-session-event-fw6-creation-records-sent-failure': Custom FW6 Session Creation Records Failed; 'custom-session-event-fw4-deletion-records-sent': Custom FW4 Session Deletion Records Sent; 'custom-session-event-fw4-deletion-records-sent-failure': Custom FW4 Session Deletion Records Failed; 'custom-session-event-fw6-deletion-records-sent': Custom FW6 Session Deletion Records Sent; 'custom-session-event-fw6-deletion-records-sent-failure': Custom FW6 Session Deletion Records Failed; 'custom-deny-reset-event-fw4-records-sent': Custom FW4 Deny/Reset Event Records Sent; 'custom-deny-reset-event-fw4-records-sent-failure': Custom FW4 Deny/Reset Event Records Failed; 'custom-deny-reset-event-fw6-records-sent': Custom FW6 Deny/Reset Event Records Sent; 'custom-deny-reset-event-fw6-records-sent-failure': Custom FW6 Deny/Reset Event Records Failed; 'custom-port-mapping-nat44-creation-records-sent': Custom Nat44 Port Map Creation Records Sent; 'custom-port-mapping-nat44-creation-records-sent-failure': Custom Nat44 Port Map Creation Records Failed; 'custom-port-mapping-nat64-creation-records-sent': Custom Nat64 Port Map Creation Records Sent; 'custom-port-mapping-nat64-creation-records-sent-failure': Custom Nat64 Port Map Creation Records Failed; 'custom-port-mapping-dslite-creation-records-sent': Custom Dslite Port Map Creation Records Sent; 'custom-port-mapping-dslite-creation-records-sent-failure': Custom Dslite Port Map Creation Records Failed; 'custom-port-mapping-nat44-deletion-records-sent': Custom Nat44 Port Map Deletion Records Sent; 'custom-port-mapping-nat44-deletion-records-sent-failure': Custom Nat44 Port Map Deletion Records Failed; 'custom-port-mapping-nat64-deletion-records-sent': Custom Nat64 Port Map Deletion Records Sent; 'custom-port-mapping-nat64-deletion-records-sent-failure': Custom Nat64 Port Map Deletion Records Failed; 'custom-port-mapping-dslite-deletion-records-sent': Custom Dslite Port Map Deletion Records Sent; 'custom-port-mapping-dslite-deletion-records-sent-failure': Custom Dslite Port Map Deletion Records Failed; 'custom-port-batching-nat44-creation-records-sent': Custom Nat44 Port Batch Creation Records Sent; 'custom-port-batching-nat44-creation-records-sent-failure': Custom Nat44 Port Batch Creation Records Failed; 'custom-port-batching-nat64-creation-records-sent': Custom Nat64 Port Batch Creation Records Sent; 'custom-port-batching-nat64-creation-records-sent-failure': Custom Nat64 Port Batch Creation Records Failed; 'custom-port-batching-dslite-creation-records-sent': Custom Dslite Port Batch Creation Records Sent; 'custom-port-batching-dslite-creation-records-sent-failure': Custom Dslite Port Batch Creation Records Failed; 'custom-port-batching-nat44-deletion-records-sent': Custom Nat44 Port Batch Deletion Records Sent; 'custom-port-batching-nat44-deletion-records-sent-failure': Custom Nat44 Port Batch Deletion Records Failed; 'custom-port-batching-nat64-deletion-records-sent': Custom Nat64 Port Batch Deletion Records Sent; 'custom-port-batching-nat64-deletion-records-sent-failure': Custom Nat64 Port Batch Deletion Records Failed; 'custom-port-batching-dslite-deletion-records-sent': Custom Dslite Port Batch Deletion Records Sent; 'custom-port-batching-dslite-deletion-records-sent-failure': Custom Dslite Port Batch Deletion Records Failed; 'custom-port-batching-v2-nat44-creation-records-sent': Custom Nat44 Port Batch V2 Creation Records Sent;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'custom-port-batching-v2-nat44-creation-records-sent-failure': Custom Nat44 Port Batch V2 Creation Records Failed; 'custom-port-batching-v2-nat64-creation-records-sent': Custom Nat64 Port Batch V2 Creation Records Sent; 'custom-port-batching-v2-nat64-creation-records-sent-failure': Custom Nat64 Port Batch V2 Creation Records Failed; 'custom-port-batching-v2-dslite-creation-records-sent': Custom Dslite Port Batch V2 Creation Records Sent; 'custom-port-batching-v2-dslite-creation-records-sent-failure': Custom Dslite Port Batch V2 Creation Records Failed; 'custom-port-batching-v2-nat44-deletion-records-sent': Custom Nat44 Port Batch V2 Deletion Records Sent; 'custom-port-batching-v2-nat44-deletion-records-sent-failure': Custom Nat44 Port Batch V2 Deletion Records Failed; 'custom-port-batching-v2-nat64-deletion-records-sent': Custom Nat64 Port Batch V2 Deletion Records Sent; 'custom-port-batching-v2-nat64-deletion-records-sent-failure': Custom Nat64 Port Batch V2 Deletion Records Failed; 'custom-port-batching-v2-dslite-deletion-records-sent': Custom Dslite Port Batch V2 Deletion Records Sent; 'custom-port-batching-v2-dslite-deletion-records-sent-failure': Custom Dslite Port Batch V2 Deletion Records Failed; 'reduced-logs-by-destination': Reduced Logs by Destination Protocol and Port; 'custom-gtp-c-tunnel-event-records-sent': Custom GTP C Tunnel Records Sent; 'custom-gtp-c-tunnel-event-records-sent-failure': Custom GTP C Tunnel Records Sent Failure; 'custom-gtp-u-tunnel-event-records-sent': Custom GTP U Tunnel Records Sent; 'custom-gtp-u-tunnel-event-records-sent-failure': Custom GTP U Tunnel Records Sent Failure; 'custom-gtp-deny-event-records-sent': Custom GTP Deny Records Sent; 'custom-gtp-deny-event-records-sent-failure': Custom GTP Deny Records Sent Failure; 'custom-gtp-info-event-records-sent': Custom GTP Info Records Sent; 'custom-gtp-info-event-records-sent-failure': Custom GTP Info Records Sent Failure; 'custom-fw-iddos-entry-created-records-sent': Custom FW iDDoS Entry Created Records Sent; 'custom-fw-iddos-entry-created-records-sent-failure': Custom FW iDDoS Entry Created Records Sent Failure; 'custom-fw-iddos-entry-deleted-records-sent': Custom FW iDDoS Entry Deleted Records Sent; 'custom-fw-iddos-entry-deleted-records-sent-failure': Custom FW iDDoS Entry Deleted Records Sent Failure; 'custom-fw-sesn-limit-exceeded-records-sent': Custom FW Session Limit Exceeded Records Sent; 'custom-fw-sesn-limit-exceeded-records-sent-failure': Custom FW Session Limit Exceeded Records Sent Failure; 'custom-nat-iddos-l3-entry-created-records-sent': Custom NAT iDDoS L3 Entry Created Records Sent; 'custom-nat-iddos-l3-entry-created-records-sent-failure': Custom NAT iDDoS L3 Entry Created Records Sent Failure; 'custom-nat-iddos-l3-entry-deleted-records-sent': Custom NAT iDDoS L3 Entry Deleted Records Sent; 'custom-nat-iddos-l3-entry-deleted-records-sent-failure': Custom NAT iDDoS L3 Entry Deleted Records Sent Failure; 'custom-nat-iddos-l4-entry-created-records-sent': Custom NAT iDDoS L4 Entry Created Records Sent; 'custom-nat-iddos-l4-entry-created-records-sent-failure': Custom NAT iDDoS L4 Entry Created Records Sent Failure; 'custom-nat-iddos-l4-entry-deleted-records-sent': Custom NAT iDDoS L4 Entry Deleted Records Sent; 'custom-nat-iddos-l4-entry-deleted-records-sent-failure': Custom NAT iDDoS L4 Entry Deleted Records Sent Failure; 'custom-gtp-rate-limit-periodic-records-sent': Custom GTP Rate Limit Periodic Records Sent; 'custom-gtp-rate-limit-periodic-records-sent-failure': Custom GTP Rate Limit Periodic Records Sent Failure; 'ddos-general-stat-records-sent': ddos-general-stat-records-sent; 'ddos-general-stat-records-sent-failure': ddos-general-stat-records-sent-failure; 'ddos-http-stat-records-sent': ddos-http-stat-records-sent; 'ddos-http-stat-records-sent-failure': ddos-http-stat-records-sent-failure;",
						},
					},
				},
			},
			"scope": {
				Type: schema.TypeString, Optional: true, Default: "global", Description: "'global': Netflow monitor is activated globally (Default); 'firewall-rule': Netflow monitor is only activated when referenced by a firewall rule;",
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
			"source_ip_use_mgmt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface's IP address for source ip of netflow packets",
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
func resourceNetflowMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceNetflowMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetflowMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceNetflowMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetflowMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetflowMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetflowMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNetflowMonitorCustomRecord1055(d []interface{}) edpt.NetflowMonitorCustomRecord1055 {

	count1 := len(d)
	var ret edpt.NetflowMonitorCustomRecord1055
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CustomCfg = getSliceNetflowMonitorCustomRecordCustomCfg1056(in["custom_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceNetflowMonitorCustomRecordCustomCfg1056(d []interface{}) []edpt.NetflowMonitorCustomRecordCustomCfg1056 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorCustomRecordCustomCfg1056, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorCustomRecordCustomCfg1056
		oi.Event = in["event"].(string)
		oi.IpfixTemplate = in["ipfix_template"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetflowMonitorDestination1057(d []interface{}) edpt.NetflowMonitorDestination1057 {

	count1 := len(d)
	var ret edpt.NetflowMonitorDestination1057
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ServiceGroup = in["service_group"].(string)
		ret.IpCfg = getObjectNetflowMonitorDestinationIpCfg1058(in["ip_cfg"].([]interface{}))
		ret.Ipv6Cfg = getObjectNetflowMonitorDestinationIpv6Cfg1059(in["ipv6_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getObjectNetflowMonitorDestinationIpCfg1058(d []interface{}) edpt.NetflowMonitorDestinationIpCfg1058 {

	count1 := len(d)
	var ret edpt.NetflowMonitorDestinationIpCfg1058
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(string)
		ret.Port4 = in["port4"].(int)
	}
	return ret
}

func getObjectNetflowMonitorDestinationIpv6Cfg1059(d []interface{}) edpt.NetflowMonitorDestinationIpv6Cfg1059 {

	count1 := len(d)
	var ret edpt.NetflowMonitorDestinationIpv6Cfg1059
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6 = in["ipv6"].(string)
		ret.Port6 = in["port6"].(int)
	}
	return ret
}

func getObjectNetflowMonitorDisableLogByDestination1060(d []interface{}) edpt.NetflowMonitorDisableLogByDestination1060 {

	count1 := len(d)
	var ret edpt.NetflowMonitorDisableLogByDestination1060
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpList = getSliceNetflowMonitorDisableLogByDestinationTcpList1061(in["tcp_list"].([]interface{}))
		ret.UdpList = getSliceNetflowMonitorDisableLogByDestinationUdpList1062(in["udp_list"].([]interface{}))
		ret.Icmp = in["icmp"].(int)
		ret.Others = in["others"].(int)
		//omit uuid
		ret.IpList = getSliceNetflowMonitorDisableLogByDestinationIpList1063(in["ip_list"].([]interface{}))
		ret.Ip6List = getSliceNetflowMonitorDisableLogByDestinationIp6List1066(in["ip6_list"].([]interface{}))
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationTcpList1061(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationTcpList1061 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationTcpList1061, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationTcpList1061
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationUdpList1062(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationUdpList1062 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationUdpList1062, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationUdpList1062
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIpList1063(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIpList1063 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIpList1063, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIpList1063
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.TcpList = getSliceNetflowMonitorDisableLogByDestinationIpListTcpList1064(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceNetflowMonitorDisableLogByDestinationIpListUdpList1065(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIpListTcpList1064(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIpListTcpList1064 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIpListTcpList1064, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIpListTcpList1064
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIpListUdpList1065(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIpListUdpList1065 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIpListUdpList1065, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIpListUdpList1065
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIp6List1066(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIp6List1066 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIp6List1066, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIp6List1066
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.TcpList = getSliceNetflowMonitorDisableLogByDestinationIp6ListTcpList1067(in["tcp_list"].([]interface{}))
		oi.UdpList = getSliceNetflowMonitorDisableLogByDestinationIp6ListUdpList1068(in["udp_list"].([]interface{}))
		oi.Icmp = in["icmp"].(int)
		oi.Others = in["others"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIp6ListTcpList1067(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIp6ListTcpList1067 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIp6ListTcpList1067, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIp6ListTcpList1067
		oi.TcpPortStart = in["tcp_port_start"].(int)
		oi.TcpPortEnd = in["tcp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorDisableLogByDestinationIp6ListUdpList1068(d []interface{}) []edpt.NetflowMonitorDisableLogByDestinationIp6ListUdpList1068 {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorDisableLogByDestinationIp6ListUdpList1068, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorDisableLogByDestinationIp6ListUdpList1068
		oi.UdpPortStart = in["udp_port_start"].(int)
		oi.UdpPortEnd = in["udp_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetflowMonitorRecord1069(d []interface{}) edpt.NetflowMonitorRecord1069 {

	count1 := len(d)
	var ret edpt.NetflowMonitorRecord1069
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NetflowV5 = in["netflow_v5"].(int)
		ret.NetflowV5Ext = in["netflow_v5_ext"].(int)
		ret.Nat44 = in["nat44"].(int)
		ret.Nat64 = in["nat64"].(int)
		ret.Dslite = in["dslite"].(int)
		ret.SesnEventNat44 = in["sesn_event_nat44"].(string)
		ret.SesnEventNat64 = in["sesn_event_nat64"].(string)
		ret.SesnEventDslite = in["sesn_event_dslite"].(string)
		ret.SesnEventFw4 = in["sesn_event_fw4"].(string)
		ret.SesnEventFw6 = in["sesn_event_fw6"].(string)
		ret.PortMappingNat44 = in["port_mapping_nat44"].(string)
		ret.PortMappingNat64 = in["port_mapping_nat64"].(string)
		ret.PortMappingDslite = in["port_mapping_dslite"].(string)
		ret.PortBatchNat44 = in["port_batch_nat44"].(string)
		ret.PortBatchNat64 = in["port_batch_nat64"].(string)
		ret.PortBatchDslite = in["port_batch_dslite"].(string)
		ret.PortBatchV2Nat44 = in["port_batch_v2_nat44"].(string)
		ret.PortBatchV2Nat64 = in["port_batch_v2_nat64"].(string)
		ret.PortBatchV2Dslite = in["port_batch_v2_dslite"].(string)
		ret.DdosGeneralStat = in["ddos_general_stat"].(int)
		ret.DdosHttpStat = in["ddos_http_stat"].(int)
		//omit uuid
	}
	return ret
}

func getObjectNetflowMonitorResendTemplate1070(d []interface{}) edpt.NetflowMonitorResendTemplate1070 {

	count1 := len(d)
	var ret edpt.NetflowMonitorResendTemplate1070
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Timeout = in["timeout"].(int)
		ret.Records = in["records"].(int)
		//omit uuid
	}
	return ret
}

func getObjectNetflowMonitorSample1071(d []interface{}) edpt.NetflowMonitorSample1071 {

	count1 := len(d)
	var ret edpt.NetflowMonitorSample1071
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthernetList = getSliceNetflowMonitorSampleEthernetList(in["ethernet_list"].([]interface{}))
		ret.VeList = getSliceNetflowMonitorSampleVeList(in["ve_list"].([]interface{}))
		ret.NatPoolList = getSliceNetflowMonitorSampleNatPoolList(in["nat_pool_list"].([]interface{}))
	}
	return ret
}

func getSliceNetflowMonitorSampleEthernetList(d []interface{}) []edpt.NetflowMonitorSampleEthernetList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorSampleEthernetList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorSampleEthernetList
		oi.Ifindex = in["ifindex"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorSampleVeList(d []interface{}) []edpt.NetflowMonitorSampleVeList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorSampleVeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorSampleVeList
		oi.VeNum = in["ve_num"].(int)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorSampleNatPoolList(d []interface{}) []edpt.NetflowMonitorSampleNatPoolList {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorSampleNatPoolList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorSampleNatPoolList
		oi.PoolName = in["pool_name"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetflowMonitorSamplingEnable(d []interface{}) []edpt.NetflowMonitorSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.NetflowMonitorSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetflowMonitorSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetflowMonitorSourceAddress1072(d []interface{}) edpt.NetflowMonitorSourceAddress1072 {

	count1 := len(d)
	var ret edpt.NetflowMonitorSourceAddress1072
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ip = in["ip"].(string)
		ret.Ipv6 = in["ipv6"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointNetflowMonitor(d *schema.ResourceData) edpt.NetflowMonitor {
	var ret edpt.NetflowMonitor
	ret.Inst.CounterPollingInterval = d.Get("counter_polling_interval").(int)
	ret.Inst.CustomRecord = getObjectNetflowMonitorCustomRecord1055(d.Get("custom_record").([]interface{}))
	ret.Inst.Destination = getObjectNetflowMonitorDestination1057(d.Get("destination").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DisableLogByDestination = getObjectNetflowMonitorDisableLogByDestination1060(d.Get("disable_log_by_destination").([]interface{}))
	ret.Inst.FlowTimeout = d.Get("flow_timeout").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Record = getObjectNetflowMonitorRecord1069(d.Get("record").([]interface{}))
	ret.Inst.ResendTemplate = getObjectNetflowMonitorResendTemplate1070(d.Get("resend_template").([]interface{}))
	ret.Inst.Sample = getObjectNetflowMonitorSample1071(d.Get("sample").([]interface{}))
	ret.Inst.SamplingEnable = getSliceNetflowMonitorSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Scope = d.Get("scope").(string)
	ret.Inst.SourceAddress = getObjectNetflowMonitorSourceAddress1072(d.Get("source_address").([]interface{}))
	ret.Inst.SourceIpUseMgmt = d.Get("source_ip_use_mgmt").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
