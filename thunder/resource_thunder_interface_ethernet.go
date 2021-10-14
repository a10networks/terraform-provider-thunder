package thunder

//Thunder resource InterfaceEthernet

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceEthernetCreate,
		UpdateContext: resourceInterfaceEthernetUpdate,
		ReadContext:   resourceInterfaceEthernetRead,
		DeleteContext: resourceInterfaceEthernetDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port_scan_detection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ping_sweep_detection": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"l3_vlan_fwd_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"load_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"media_type_copper": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_neg_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fec_forced_on": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fec_forced_off": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"port_breakout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"speed_forced_1g": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"speed_forced_10g": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"speed_forced_40g": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ipg_bit_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"remove_vlan_tag": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"trap_source": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"duplexity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"flow_control": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"icmp_rate_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"normal": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lockup": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lockup_period": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"icmpv6_rate_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"normal_v6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lockup_v6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"lockup_period_v6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"monitor_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"mirror_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"monitor_vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cpu_process": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cpu_process_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"traffic_distribution_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"virtual_wire": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"update_l2_info": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vlan_learning": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"mac_learning": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"access_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sampling_enable": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"packet_capture_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"lldp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rt_enable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rx": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tx": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"notification_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"notification": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"notif_enable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"tx_dot1_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tx_dot1_tlvs": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"link_aggregation": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"vlan": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"tx_tlvs_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tx_tlvs": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"exclude": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"management_address": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_description": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"system_capabilities": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"system_description": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"system_name": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ddos": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dhcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"address_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"address_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"allow_promiscuous_vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"cache_spoofing_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"helper_address_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"helper_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ttl_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"syn_cookie": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"slb_partition_redirect": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"generate_membership_query": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"query_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_resp_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"client": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"server": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"stateful_firewall": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"class_list": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"outside": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"access_list": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"router": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"isis": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tag": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"rip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"str": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"string": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"mode": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"key_chain": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key_chain": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
											},
										},
									},
									"send_packet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"receive_packet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"send_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"version": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"receive_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"receive": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"version": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"split_horizon_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"state": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf_global": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"authentication": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"value": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"authentication_key": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"bfd_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bfd": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"disable": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"database_filter_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"database_filter": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"out": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"dead_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"disable": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"message_digest_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message_digest_key": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"md5_value": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"mtu": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"mtu_ignore": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"network": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"broadcast": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"non_broadcast": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"point_to_point": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"point_to_multipoint": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"p2mp_nbma": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"retransmit_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"transmit_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"ospf_ip_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_addr": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"authentication": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"authentication_key": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"database_filter": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"out": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dead_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"message_digest_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"message_digest_key": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"md5_value": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"mtu_ignore": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"retransmit_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"transmit_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
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
			"ipv6": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6_addr": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"address_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipv6_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ttl_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"access_list_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"v6_acl_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"inbound": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"router_adver": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"hop_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"max_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"min_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"default_lifetime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"reachable_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_timer": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"adver_mtu_disable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"adver_mtu": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"prefix_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"prefix": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"not_autonomous": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"not_on_link": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"preferred_lifetime": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"valid_lifetime": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"managed_config_action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"other_config_action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"adver_vrid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"use_floating_ip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"floating_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"adver_vrid_default": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"use_floating_ip_default_vrid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"floating_ip_default_vrid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"stateful_firewall": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inside": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"class_list": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"outside": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"access_list": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"acl_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"router": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ripng": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"rip": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"ospf": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"area_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"area_id_num": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"area_id_addr": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"tag": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"instance_id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"isis": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tag": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"rip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"split_horizon_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"state": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"network_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"broadcast_type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"p2mp_nbma": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"network_instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"bfd": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"cost_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"dead_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dead_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"hello_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mtu_ignore_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mtu_ignore": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"neighbor_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"neighbor": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"neig_inst": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_poll_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"priority_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"retransmit_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"retransmit_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"transmit_delay_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"transmit_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"nptv6": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"domain_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"bind_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"map": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"map_t_inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"map_t_outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"lw_4o6": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"trunk_group_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"admin_key": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"udld_timeout_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fast": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"slow": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"bfd": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"echo": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"demand": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"min_rx": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"multiplier": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"isis": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send_only_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"send_only": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"level": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"mode_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"level": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"key_chain_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_chain": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"level": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"bfd_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bfd": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"circuit_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"csnp_interval_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"csnp_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"padding": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"hello_interval_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"hello_interval_minimal_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_interval_minimal": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"hello_multiplier_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hello_multiplier": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"lsp_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mesh_group": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"blocked": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"metric_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"password_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"priority_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"retransmit_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"wide_metric_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"wide_metric": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"spanning_tree": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_edge": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"admin_edge": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"instance_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_start": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"mstp_path_cost": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"path_cost": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceInterfaceEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernet (Inside resourceInterfaceEthernetCreate) ")
		name1 := d.Get("ifnum").(int)
		data := dataToInterfaceEthernet(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceEthernet --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostInterfaceEthernet(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceEthernetRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceEthernetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceEthernet (Inside resourceInterfaceEthernetRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceEthernet(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return diags
}

func resourceInterfaceEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying InterfaceEthernet   (Inside resourceInterfaceEthernetUpdate) ")
		data := dataToInterfaceEthernet(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceEthernet ")
		err := go_thunder.PutInterfaceEthernet(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceEthernetRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceEthernetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceEthernetDelete) " + name1)
		err := go_thunder.DeleteInterfaceEthernet(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToInterfaceEthernet(d *schema.ResourceData) go_thunder.InterfaceEthernet {
	var vc go_thunder.InterfaceEthernet
	var c go_thunder.InterfaceEthernetInstance
	c.InterfaceEthernetInstanceIfnum = d.Get("ifnum").(int)
	c.InterfaceEthernetInstanceName = d.Get("name").(string)
	c.InterfaceEthernetInstancePortScanDetection = d.Get("port_scan_detection").(string)
	c.InterfaceEthernetInstancePingSweepDetection = d.Get("ping_sweep_detection").(string)
	c.InterfaceEthernetInstanceL3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	c.InterfaceEthernetInstanceLoadInterval = d.Get("load_interval").(int)
	c.InterfaceEthernetInstanceMediaTypeCopper = d.Get("media_type_copper").(int)
	c.InterfaceEthernetInstanceAutoNegEnable = d.Get("auto_neg_enable").(int)
	c.InterfaceEthernetInstanceFecForcedOn = d.Get("fec_forced_on").(int)
	c.InterfaceEthernetInstanceFecForcedOff = d.Get("fec_forced_off").(int)
	c.InterfaceEthernetInstancePortBreakout = d.Get("port_breakout").(string)
	c.InterfaceEthernetInstanceSpeedForced1G = d.Get("speed_forced_1g").(int)
	c.InterfaceEthernetInstanceSpeedForced10G = d.Get("speed_forced_10g").(int)
	c.InterfaceEthernetInstanceSpeedForced40G = d.Get("speed_forced_40g").(int)
	c.InterfaceEthernetInstanceIpgBitTime = d.Get("ipg_bit_time").(int)
	c.InterfaceEthernetInstanceRemoveVlanTag = d.Get("remove_vlan_tag").(int)
	c.InterfaceEthernetInstanceMtu = d.Get("mtu").(int)
	c.InterfaceEthernetInstanceTrapSource = d.Get("trap_source").(int)
	c.InterfaceEthernetInstanceDuplexity = d.Get("duplexity").(string)
	c.InterfaceEthernetInstanceSpeed = d.Get("speed").(string)
	c.InterfaceEthernetInstanceFlowControl = d.Get("flow_control").(int)
	c.InterfaceEthernetInstanceAction = d.Get("action").(string)

	var obj1 go_thunder.InterfaceEthernetInstanceIcmpRateLimit
	prefix1 := "icmp_rate_limit.0."
	obj1.InterfaceEthernetInstanceIcmpRateLimitNormal = d.Get(prefix1 + "normal").(int)
	obj1.InterfaceEthernetInstanceIcmpRateLimitLockup = d.Get(prefix1 + "lockup").(int)
	obj1.InterfaceEthernetInstanceIcmpRateLimitLockupPeriod = d.Get(prefix1 + "lockup_period").(int)

	c.InterfaceEthernetInstanceIcmpRateLimitNormal = obj1

	var obj2 go_thunder.InterfaceEthernetInstanceIcmpv6RateLimit
	prefix2 := "icmpv6_rate_limit.0."
	obj2.InterfaceEthernetInstanceIcmpv6RateLimitNormalV6 = d.Get(prefix2 + "normal_v6").(int)
	obj2.InterfaceEthernetInstanceIcmpv6RateLimitLockupV6 = d.Get(prefix2 + "lockup_v6").(int)
	obj2.InterfaceEthernetInstanceIcmpv6RateLimitLockupPeriodV6 = d.Get(prefix2 + "lockup_period_v6").(int)

	c.InterfaceEthernetInstanceIcmpv6RateLimitNormalV6 = obj2

	InterfaceEthernetInstanceMonitorListCount := d.Get("monitor_list.#").(int)
	c.InterfaceEthernetInstanceMonitorListMonitor = make([]go_thunder.InterfaceEthernetInstanceMonitorList, 0, InterfaceEthernetInstanceMonitorListCount)

	for i := 0; i < InterfaceEthernetInstanceMonitorListCount; i++ {
		var obj3 go_thunder.InterfaceEthernetInstanceMonitorList
		prefix3 := fmt.Sprintf("monitor_list.%d.", i)
		obj3.InterfaceEthernetInstanceMonitorListMonitor = d.Get(prefix3 + "monitor").(string)
		obj3.InterfaceEthernetInstanceMonitorListMirrorIndex = d.Get(prefix3 + "mirror_index").(int)
		obj3.InterfaceEthernetInstanceMonitorListMonitorVlan = d.Get(prefix3 + "monitor_vlan").(int)
		c.InterfaceEthernetInstanceMonitorListMonitor = append(c.InterfaceEthernetInstanceMonitorListMonitor, obj3)
	}

	c.InterfaceEthernetInstanceCPUProcess = d.Get("cpu_process").(int)
	c.InterfaceEthernetInstanceCPUProcessDir = d.Get("cpu_process_dir").(string)
	c.InterfaceEthernetInstanceTrafficDistributionMode = d.Get("traffic_distribution_mode").(string)
	c.InterfaceEthernetInstanceVirtualWire = d.Get("virtual_wire").(int)
	c.InterfaceEthernetInstanceUpdateL2Info = d.Get("update_l2_info").(int)
	c.InterfaceEthernetInstanceVlanLearning = d.Get("vlan_learning").(string)
	c.InterfaceEthernetInstanceMacLearning = d.Get("mac_learning").(string)

	var obj4 go_thunder.InterfaceEthernetInstanceAccessList
	prefix4 := "access_list.0."
	obj4.InterfaceEthernetInstanceAccessListAclID = d.Get(prefix4 + "acl_id").(int)
	obj4.InterfaceEthernetInstanceAccessListAclName = d.Get(prefix4 + "acl_name").(string)

	c.InterfaceEthernetInstanceAccessListAclID = obj4

	c.InterfaceEthernetInstanceUserTag = d.Get("user_tag").(string)

	InterfaceEthernetInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.InterfaceEthernetInstanceSamplingEnableCounters1 = make([]go_thunder.InterfaceEthernetInstanceSamplingEnable, 0, InterfaceEthernetInstanceSamplingEnableCount)

	for i := 0; i < InterfaceEthernetInstanceSamplingEnableCount; i++ {
		var obj5 go_thunder.InterfaceEthernetInstanceSamplingEnable
		prefix5 := fmt.Sprintf("sampling_enable.%d.", i)
		obj5.InterfaceEthernetInstanceSamplingEnableCounters1 = d.Get(prefix5 + "counters1").(string)
		c.InterfaceEthernetInstanceSamplingEnableCounters1 = append(c.InterfaceEthernetInstanceSamplingEnableCounters1, obj5)
	}

	c.InterfaceEthernetInstancePacketCaptureTemplate = d.Get("packet_capture_template").(string)

	var obj6 go_thunder.InterfaceEthernetInstanceLldp
	prefix6 := "lldp.0."

	var obj6_1 go_thunder.InterfaceEthernetInstanceLldpEnableCfg
	prefix6_1 := prefix6 + "enable_cfg.0."
	obj6_1.InterfaceEthernetInstanceLldpEnableCfgRtEnable = d.Get(prefix6_1 + "rt_enable").(int)
	obj6_1.InterfaceEthernetInstanceLldpEnableCfgRx = d.Get(prefix6_1 + "rx").(int)
	obj6_1.InterfaceEthernetInstanceLldpEnableCfgTx = d.Get(prefix6_1 + "tx").(int)

	obj6.InterfaceEthernetInstanceLldpEnableCfgRtEnable = obj6_1

	var obj6_2 go_thunder.InterfaceEthernetInstanceLldpNotificationCfg
	prefix6_2 := prefix6 + "notification_cfg.0."
	obj6_2.InterfaceEthernetInstanceLldpNotificationCfgNotification = d.Get(prefix6_2 + "notification").(int)
	obj6_2.InterfaceEthernetInstanceLldpNotificationCfgNotifEnable = d.Get(prefix6_2 + "notif_enable").(int)

	obj6.InterfaceEthernetInstanceLldpNotificationCfgNotification = obj6_2

	var obj6_3 go_thunder.InterfaceEthernetInstanceLldpTxDot1Cfg
	prefix6_3 := prefix6 + "tx_dot1_cfg.0."
	obj6_3.InterfaceEthernetInstanceLldpTxDot1CfgTxDot1Tlvs = d.Get(prefix6_3 + "tx_dot1_tlvs").(int)
	obj6_3.InterfaceEthernetInstanceLldpTxDot1CfgLinkAggregation = d.Get(prefix6_3 + "link_aggregation").(int)
	obj6_3.InterfaceEthernetInstanceLldpTxDot1CfgVlan = d.Get(prefix6_3 + "vlan").(int)

	obj6.InterfaceEthernetInstanceLldpTxDot1CfgTxDot1Tlvs = obj6_3

	var obj6_4 go_thunder.InterfaceEthernetInstanceLldpTxTlvsCfg
	prefix6_4 := prefix6 + "tx_tlvs_cfg.0."
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgTxTlvs = d.Get(prefix6_4 + "tx_tlvs").(int)
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgExclude = d.Get(prefix6_4 + "exclude").(int)
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgManagementAddress = d.Get(prefix6_4 + "management_address").(int)
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgPortDescription = d.Get(prefix6_4 + "port_description").(int)
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgSystemCapabilities = d.Get(prefix6_4 + "system_capabilities").(int)
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgSystemDescription = d.Get(prefix6_4 + "system_description").(int)
	obj6_4.InterfaceEthernetInstanceLldpTxTlvsCfgSystemName = d.Get(prefix6_4 + "system_name").(int)

	obj6.InterfaceEthernetInstanceLldpTxTlvsCfgTxTlvs = obj6_4

	c.InterfaceEthernetInstanceLldpEnableCfg = obj6

	var obj7 go_thunder.InterfaceEthernetInstanceDdos
	prefix7 := "ddos.0."
	obj7.InterfaceEthernetInstanceDdosOutside = d.Get(prefix7 + "outside").(int)
	obj7.InterfaceEthernetInstanceDdosInside = d.Get(prefix7 + "inside").(int)

	c.InterfaceEthernetInstanceDdosOutside = obj7

	var obj8 go_thunder.InterfaceEthernetInstanceIP
	prefix8 := "ip.0."
	obj8.InterfaceEthernetInstanceIPDhcp = d.Get(prefix8 + "dhcp").(int)

	InterfaceEthernetInstanceIPAddressListCount := d.Get(prefix8 + "address_list.#").(int)
	obj8.InterfaceEthernetInstanceIPAddressListIpv6Addr = make([]go_thunder.InterfaceEthernetInstanceIPAddressList, 0, InterfaceEthernetInstanceIPAddressListCount)

	for i := 0; i < InterfaceEthernetInstanceIPAddressListCount; i++ {
		var obj8_1 go_thunder.InterfaceEthernetInstanceIPAddressList
		prefix8_1 := prefix8 + fmt.Sprintf("address_list.%d.", i)
		obj8_1.InterfaceEthernetInstanceIPAddressListIpv6Addr = d.Get(prefix8_1 + "ipv6_addr").(string)
		obj8_1.InterfaceEthernetInstanceIPAddressListAddressType = d.Get(prefix8_1 + "address_type").(string)
		obj8.InterfaceEthernetInstanceIPAddressListIpv6Addr = append(obj8.InterfaceEthernetInstanceIPAddressListIpv6Addr, obj8_1)
	}

	obj8.InterfaceEthernetInstanceIPAllowPromiscuousVip = d.Get(prefix8 + "allow_promiscuous_vip").(int)
	obj8.InterfaceEthernetInstanceIPCacheSpoofingPort = d.Get(prefix8 + "cache_spoofing_port").(int)

	InterfaceEthernetInstanceIPHelperAddressListCount := d.Get(prefix8 + "helper_address_list.#").(int)
	obj8.InterfaceEthernetInstanceIPHelperAddressListHelperAddress = make([]go_thunder.InterfaceEthernetInstanceIPHelperAddressList, 0, InterfaceEthernetInstanceIPHelperAddressListCount)

	for i := 0; i < InterfaceEthernetInstanceIPHelperAddressListCount; i++ {
		var obj8_2 go_thunder.InterfaceEthernetInstanceIPHelperAddressList
		prefix8_2 := prefix8 + fmt.Sprintf("helper_address_list.%d.", i)
		obj8_2.InterfaceEthernetInstanceIPHelperAddressListHelperAddress = d.Get(prefix8_2 + "helper_address").(string)
		obj8.InterfaceEthernetInstanceIPHelperAddressListHelperAddress = append(obj8.InterfaceEthernetInstanceIPHelperAddressListHelperAddress, obj8_2)
	}

	obj8.InterfaceEthernetInstanceIPInside = d.Get(prefix8 + "inside").(int)
	obj8.InterfaceEthernetInstanceIPOutside = d.Get(prefix8 + "outside").(int)
	obj8.InterfaceEthernetInstanceIPTTLIgnore = d.Get(prefix8 + "ttl_ignore").(int)
	obj8.InterfaceEthernetInstanceIPSynCookie = d.Get(prefix8 + "syn_cookie").(int)
	obj8.InterfaceEthernetInstanceIPSlbPartitionRedirect = d.Get(prefix8 + "slb_partition_redirect").(int)
	obj8.InterfaceEthernetInstanceIPGenerateMembershipQuery = d.Get(prefix8 + "generate_membership_query").(int)
	obj8.InterfaceEthernetInstanceIPQueryInterval = d.Get(prefix8 + "query_interval").(int)
	obj8.InterfaceEthernetInstanceIPMaxRespTime = d.Get(prefix8 + "max_resp_time").(int)
	obj8.InterfaceEthernetInstanceIPClient = d.Get(prefix8 + "client").(int)
	obj8.InterfaceEthernetInstanceIPServer = d.Get(prefix8 + "server").(int)

	var obj8_3 go_thunder.InterfaceEthernetInstanceIPStatefulFirewall
	prefix8_3 := prefix8 + "stateful_firewall.0."
	obj8_3.InterfaceEthernetInstanceIPStatefulFirewallInside = d.Get(prefix8_3 + "inside").(int)
	obj8_3.InterfaceEthernetInstanceIPStatefulFirewallClassList = d.Get(prefix8_3 + "class_list").(string)
	obj8_3.InterfaceEthernetInstanceIPStatefulFirewallOutside = d.Get(prefix8_3 + "outside").(int)
	obj8_3.InterfaceEthernetInstanceIPStatefulFirewallAccessList = d.Get(prefix8_3 + "access_list").(int)
	obj8_3.InterfaceEthernetInstanceIPStatefulFirewallAclID = d.Get(prefix8_3 + "acl_id").(int)

	obj8.InterfaceEthernetInstanceIPStatefulFirewallInside = obj8_3

	var obj8_4 go_thunder.InterfaceEthernetInstanceIPRouter
	prefix8_4 := prefix8 + "router.0."

	var obj8_4_1 go_thunder.InterfaceEthernetInstanceIPRouterIsis
	prefix8_4_1 := prefix8_4 + "isis.0."
	obj8_4_1.InterfaceEthernetInstanceIPRouterIsisTag = d.Get(prefix8_4_1 + "tag").(string)

	obj8_4.InterfaceEthernetInstanceIPRouterIsisTag = obj8_4_1

	obj8.InterfaceEthernetInstanceIPRouterIsis = obj8_4

	var obj8_5 go_thunder.InterfaceEthernetInstanceIPRip
	prefix8_5 := prefix8 + "rip.0."

	var obj8_5_1 go_thunder.InterfaceEthernetInstanceIPRipAuthentication
	prefix8_5_1 := prefix8_5 + "authentication.0."

	var obj8_5_1_1 go_thunder.InterfaceEthernetInstanceIPRipAuthenticationStr
	prefix8_5_1_1 := prefix8_5_1 + "str.0."
	obj8_5_1_1.InterfaceEthernetInstanceIPRipAuthenticationStrString = d.Get(prefix8_5_1_1 + "string").(string)

	obj8_5_1.InterfaceEthernetInstanceIPRipAuthenticationStrString = obj8_5_1_1

	var obj8_5_1_2 go_thunder.InterfaceEthernetInstanceIPRipAuthenticationMode
	prefix8_5_1_2 := prefix8_5_1 + "mode.0."
	obj8_5_1_2.InterfaceEthernetInstanceIPRipAuthenticationModeMode = d.Get(prefix8_5_1_2 + "mode").(string)

	obj8_5_1.InterfaceEthernetInstanceIPRipAuthenticationModeMode = obj8_5_1_2

	var obj8_5_1_3 go_thunder.InterfaceEthernetInstanceIPRipAuthenticationKeyChain
	prefix8_5_1_3 := prefix8_5_1 + "key_chain.0."
	obj8_5_1_3.InterfaceEthernetInstanceIPRipAuthenticationKeyChainKeyChain = d.Get(prefix8_5_1_3 + "key_chain").(string)

	obj8_5_1.InterfaceEthernetInstanceIPRipAuthenticationKeyChainKeyChain = obj8_5_1_3

	obj8_5.InterfaceEthernetInstanceIPRipAuthenticationStr = obj8_5_1

	obj8_5.InterfaceEthernetInstanceIPRipSendPacket = d.Get(prefix8_5 + "send_packet").(int)
	obj8_5.InterfaceEthernetInstanceIPRipReceivePacket = d.Get(prefix8_5 + "receive_packet").(int)

	var obj8_5_2 go_thunder.InterfaceEthernetInstanceIPRipSendCfg
	prefix8_5_2 := prefix8_5 + "send_cfg.0."
	obj8_5_2.InterfaceEthernetInstanceIPRipSendCfgSend = d.Get(prefix8_5_2 + "send").(int)
	obj8_5_2.InterfaceEthernetInstanceIPRipSendCfgVersion = d.Get(prefix8_5_2 + "version").(string)

	obj8_5.InterfaceEthernetInstanceIPRipSendCfgSend = obj8_5_2

	var obj8_5_3 go_thunder.InterfaceEthernetInstanceIPRipReceiveCfg
	prefix8_5_3 := prefix8_5 + "receive_cfg.0."
	obj8_5_3.InterfaceEthernetInstanceIPRipReceiveCfgReceive = d.Get(prefix8_5_3 + "receive").(int)
	obj8_5_3.InterfaceEthernetInstanceIPRipReceiveCfgVersion = d.Get(prefix8_5_3 + "version").(string)

	obj8_5.InterfaceEthernetInstanceIPRipReceiveCfgReceive = obj8_5_3

	var obj8_5_4 go_thunder.InterfaceEthernetInstanceIPRipSplitHorizonCfg
	prefix8_5_4 := prefix8_5 + "split_horizon_cfg.0."
	obj8_5_4.InterfaceEthernetInstanceIPRipSplitHorizonCfgState = d.Get(prefix8_5_4 + "state").(string)

	obj8_5.InterfaceEthernetInstanceIPRipSplitHorizonCfgState = obj8_5_4

	obj8.InterfaceEthernetInstanceIPRipAuthentication = obj8_5

	var obj8_6 go_thunder.InterfaceEthernetInstanceIPOspf
	prefix8_6 := prefix8 + "ospf.0."

	var obj8_6_1 go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobal
	prefix8_6_1 := prefix8_6 + "ospf_global.0."

	var obj8_6_1_1 go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfg
	prefix8_6_1_1 := prefix8_6_1 + "authentication_cfg.0."
	obj8_6_1_1.InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication = d.Get(prefix8_6_1_1 + "authentication").(int)
	obj8_6_1_1.InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfgValue = d.Get(prefix8_6_1_1 + "value").(string)

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication = obj8_6_1_1

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationKey = d.Get(prefix8_6_1 + "authentication_key").(string)

	var obj8_6_1_2 go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfg
	prefix8_6_1_2 := prefix8_6_1 + "bfd_cfg.0."
	obj8_6_1_2.InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfgBfd = d.Get(prefix8_6_1_2 + "bfd").(int)
	obj8_6_1_2.InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfgDisable = d.Get(prefix8_6_1_2 + "disable").(int)

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalBfdCfgBfd = obj8_6_1_2

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalCost = d.Get(prefix8_6_1 + "cost").(int)

	var obj8_6_1_3 go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfg
	prefix8_6_1_3 := prefix8_6_1 + "database_filter_cfg.0."
	obj8_6_1_3.InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = d.Get(prefix8_6_1_3 + "database_filter").(string)
	obj8_6_1_3.InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfgOut = d.Get(prefix8_6_1_3 + "out").(int)

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = obj8_6_1_3

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalDeadInterval = d.Get(prefix8_6_1 + "dead_interval").(int)
	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalDisable = d.Get(prefix8_6_1 + "disable").(string)
	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalHelloInterval = d.Get(prefix8_6_1 + "hello_interval").(int)

	InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgCount := d.Get(prefix8_6_1 + "message_digest_cfg.#").(int)
	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfg, 0, InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgCount; i++ {
		var obj8_6_1_4 go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfg
		prefix8_6_1_4 := prefix8_6_1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj8_6_1_4.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = d.Get(prefix8_6_1_4 + "message_digest_key").(int)
		obj8_6_1_4.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMd5Value = d.Get(prefix8_6_1_4 + "md5_value").(string)
		obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = append(obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey, obj8_6_1_4)
	}

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalMtu = d.Get(prefix8_6_1 + "mtu").(int)
	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalMtuIgnore = d.Get(prefix8_6_1 + "mtu_ignore").(int)

	var obj8_6_1_5 go_thunder.InterfaceEthernetInstanceIPOspfOspfGlobalNetwork
	prefix8_6_1_5 := prefix8_6_1 + "network.0."
	obj8_6_1_5.InterfaceEthernetInstanceIPOspfOspfGlobalNetworkBroadcast = d.Get(prefix8_6_1_5 + "broadcast").(int)
	obj8_6_1_5.InterfaceEthernetInstanceIPOspfOspfGlobalNetworkNonBroadcast = d.Get(prefix8_6_1_5 + "non_broadcast").(int)
	obj8_6_1_5.InterfaceEthernetInstanceIPOspfOspfGlobalNetworkPointToPoint = d.Get(prefix8_6_1_5 + "point_to_point").(int)
	obj8_6_1_5.InterfaceEthernetInstanceIPOspfOspfGlobalNetworkPointToMultipoint = d.Get(prefix8_6_1_5 + "point_to_multipoint").(int)
	obj8_6_1_5.InterfaceEthernetInstanceIPOspfOspfGlobalNetworkP2MpNbma = d.Get(prefix8_6_1_5 + "p2mp_nbma").(int)

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalNetworkBroadcast = obj8_6_1_5

	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalPriority = d.Get(prefix8_6_1 + "priority").(int)
	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalRetransmitInterval = d.Get(prefix8_6_1 + "retransmit_interval").(int)
	obj8_6_1.InterfaceEthernetInstanceIPOspfOspfGlobalTransmitDelay = d.Get(prefix8_6_1 + "transmit_delay").(int)

	obj8_6.InterfaceEthernetInstanceIPOspfOspfGlobalAuthenticationCfg = obj8_6_1

	InterfaceEthernetInstanceIPOspfOspfIPListCount := d.Get(prefix8_6 + "ospf_ip_list.#").(int)
	obj8_6.InterfaceEthernetInstanceIPOspfOspfIPListIPAddr = make([]go_thunder.InterfaceEthernetInstanceIPOspfOspfIPList, 0, InterfaceEthernetInstanceIPOspfOspfIPListCount)

	for i := 0; i < InterfaceEthernetInstanceIPOspfOspfIPListCount; i++ {
		var obj8_6_2 go_thunder.InterfaceEthernetInstanceIPOspfOspfIPList
		prefix8_6_2 := prefix8_6 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListIPAddr = d.Get(prefix8_6_2 + "ip_addr").(string)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListAuthentication = d.Get(prefix8_6_2 + "authentication").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListValue = d.Get(prefix8_6_2 + "value").(string)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListAuthenticationKey = d.Get(prefix8_6_2 + "authentication_key").(string)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListCost = d.Get(prefix8_6_2 + "cost").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListDatabaseFilter = d.Get(prefix8_6_2 + "database_filter").(string)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListOut = d.Get(prefix8_6_2 + "out").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListDeadInterval = d.Get(prefix8_6_2 + "dead_interval").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListHelloInterval = d.Get(prefix8_6_2 + "hello_interval").(int)

		InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgCount := d.Get(prefix8_6_2 + "message_digest_cfg.#").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfg, 0, InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgCount)

		for i := 0; i < InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgCount; i++ {
			var obj8_6_2_1 go_thunder.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfg
			prefix8_6_2_1 := prefix8_6_2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj8_6_2_1.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = d.Get(prefix8_6_2_1 + "message_digest_key").(int)
			obj8_6_2_1.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMd5Value = d.Get(prefix8_6_2_1 + "md5_value").(string)
			obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = append(obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey, obj8_6_2_1)
		}

		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListMtuIgnore = d.Get(prefix8_6_2 + "mtu_ignore").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListPriority = d.Get(prefix8_6_2 + "priority").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListRetransmitInterval = d.Get(prefix8_6_2 + "retransmit_interval").(int)
		obj8_6_2.InterfaceEthernetInstanceIPOspfOspfIPListTransmitDelay = d.Get(prefix8_6_2 + "transmit_delay").(int)
		obj8_6.InterfaceEthernetInstanceIPOspfOspfIPListIPAddr = append(obj8_6.InterfaceEthernetInstanceIPOspfOspfIPListIPAddr, obj8_6_2)
	}

	obj8.InterfaceEthernetInstanceIPOspfOspfGlobal = obj8_6

	c.InterfaceEthernetInstanceIPDhcp = obj8

	var obj9 go_thunder.InterfaceEthernetInstanceIpv6
	prefix9 := "ipv6.0."

	InterfaceEthernetInstanceIpv6AddressListCount := d.Get(prefix9 + "address_list.#").(int)
	obj9.InterfaceEthernetInstanceIpv6AddressListIpv6Addr = make([]go_thunder.InterfaceEthernetInstanceIpv6AddressList, 0, InterfaceEthernetInstanceIpv6AddressListCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6AddressListCount; i++ {
		var obj9_1 go_thunder.InterfaceEthernetInstanceIpv6AddressList
		prefix9_1 := prefix9 + fmt.Sprintf("address_list.%d.", i)
		obj9_1.InterfaceEthernetInstanceIpv6AddressListIpv6Addr = d.Get(prefix9_1 + "ipv6_addr").(string)
		obj9_1.InterfaceEthernetInstanceIpv6AddressListAddressType = d.Get(prefix9_1 + "address_type").(string)
		obj9.InterfaceEthernetInstanceIpv6AddressListIpv6Addr = append(obj9.InterfaceEthernetInstanceIpv6AddressListIpv6Addr, obj9_1)
	}

	obj9.InterfaceEthernetInstanceIpv6Inside = d.Get(prefix9 + "inside").(int)
	obj9.InterfaceEthernetInstanceIpv6Outside = d.Get(prefix9 + "outside").(int)
	obj9.InterfaceEthernetInstanceIpv6Ipv6Enable = d.Get(prefix9 + "ipv6_enable").(int)
	obj9.InterfaceEthernetInstanceIpv6TTLIgnore = d.Get(prefix9 + "ttl_ignore").(int)

	var obj9_2 go_thunder.InterfaceEthernetInstanceIpv6AccessListCfg
	prefix9_2 := prefix9 + "access_list_cfg.0."
	obj9_2.InterfaceEthernetInstanceIpv6AccessListCfgV6AclName = d.Get(prefix9_2 + "v6_acl_name").(string)
	obj9_2.InterfaceEthernetInstanceIpv6AccessListCfgInbound = d.Get(prefix9_2 + "inbound").(int)

	obj9.InterfaceEthernetInstanceIpv6AccessListCfgV6AclName = obj9_2

	var obj9_3 go_thunder.InterfaceEthernetInstanceIpv6RouterAdver
	prefix9_3 := prefix9 + "router_adver.0."
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverAction = d.Get(prefix9_3 + "action").(string)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverHopLimit = d.Get(prefix9_3 + "hop_limit").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverMaxInterval = d.Get(prefix9_3 + "max_interval").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverMinInterval = d.Get(prefix9_3 + "min_interval").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverDefaultLifetime = d.Get(prefix9_3 + "default_lifetime").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverRateLimit = d.Get(prefix9_3 + "rate_limit").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverReachableTime = d.Get(prefix9_3 + "reachable_time").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverRetransmitTimer = d.Get(prefix9_3 + "retransmit_timer").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverAdverMtuDisable = d.Get(prefix9_3 + "adver_mtu_disable").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverAdverMtu = d.Get(prefix9_3 + "adver_mtu").(int)

	InterfaceEthernetInstanceIpv6RouterAdverPrefixListCount := d.Get(prefix9_3 + "prefix_list.#").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverPrefixListPrefix = make([]go_thunder.InterfaceEthernetInstanceIpv6RouterAdverPrefixList, 0, InterfaceEthernetInstanceIpv6RouterAdverPrefixListCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6RouterAdverPrefixListCount; i++ {
		var obj9_3_1 go_thunder.InterfaceEthernetInstanceIpv6RouterAdverPrefixList
		prefix9_3_1 := prefix9_3 + fmt.Sprintf("prefix_list.%d.", i)
		obj9_3_1.InterfaceEthernetInstanceIpv6RouterAdverPrefixListPrefix = d.Get(prefix9_3_1 + "prefix").(string)
		obj9_3_1.InterfaceEthernetInstanceIpv6RouterAdverPrefixListNotAutonomous = d.Get(prefix9_3_1 + "not_autonomous").(int)
		obj9_3_1.InterfaceEthernetInstanceIpv6RouterAdverPrefixListNotOnLink = d.Get(prefix9_3_1 + "not_on_link").(int)
		obj9_3_1.InterfaceEthernetInstanceIpv6RouterAdverPrefixListPreferredLifetime = d.Get(prefix9_3_1 + "preferred_lifetime").(int)
		obj9_3_1.InterfaceEthernetInstanceIpv6RouterAdverPrefixListValidLifetime = d.Get(prefix9_3_1 + "valid_lifetime").(int)
		obj9_3.InterfaceEthernetInstanceIpv6RouterAdverPrefixListPrefix = append(obj9_3.InterfaceEthernetInstanceIpv6RouterAdverPrefixListPrefix, obj9_3_1)
	}

	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverManagedConfigAction = d.Get(prefix9_3 + "managed_config_action").(string)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverOtherConfigAction = d.Get(prefix9_3 + "other_config_action").(string)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverAdverVrid = d.Get(prefix9_3 + "adver_vrid").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverUseFloatingIP = d.Get(prefix9_3 + "use_floating_ip").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverFloatingIP = d.Get(prefix9_3 + "floating_ip").(string)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverAdverVridDefault = d.Get(prefix9_3 + "adver_vrid_default").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverUseFloatingIPDefaultVrid = d.Get(prefix9_3 + "use_floating_ip_default_vrid").(int)
	obj9_3.InterfaceEthernetInstanceIpv6RouterAdverFloatingIPDefaultVrid = d.Get(prefix9_3 + "floating_ip_default_vrid").(string)

	obj9.InterfaceEthernetInstanceIpv6RouterAdverAction = obj9_3

	var obj9_4 go_thunder.InterfaceEthernetInstanceIpv6StatefulFirewall
	prefix9_4 := prefix9 + "stateful_firewall.0."
	obj9_4.InterfaceEthernetInstanceIpv6StatefulFirewallInside = d.Get(prefix9_4 + "inside").(int)
	obj9_4.InterfaceEthernetInstanceIpv6StatefulFirewallClassList = d.Get(prefix9_4 + "class_list").(string)
	obj9_4.InterfaceEthernetInstanceIpv6StatefulFirewallOutside = d.Get(prefix9_4 + "outside").(int)
	obj9_4.InterfaceEthernetInstanceIpv6StatefulFirewallAccessList = d.Get(prefix9_4 + "access_list").(int)
	obj9_4.InterfaceEthernetInstanceIpv6StatefulFirewallAclName = d.Get(prefix9_4 + "acl_name").(string)

	obj9.InterfaceEthernetInstanceIpv6StatefulFirewallInside = obj9_4

	var obj9_5 go_thunder.InterfaceEthernetInstanceIpv6Router
	prefix9_5 := prefix9 + "router.0."

	var obj9_5_1 go_thunder.InterfaceEthernetInstanceIpv6RouterRipng
	prefix9_5_1 := prefix9_5 + "ripng.0."
	obj9_5_1.InterfaceEthernetInstanceIpv6RouterRipngRip = d.Get(prefix9_5_1 + "rip").(int)

	obj9_5.InterfaceEthernetInstanceIpv6RouterRipngRip = obj9_5_1

	var obj9_5_2 go_thunder.InterfaceEthernetInstanceIpv6RouterOspf
	prefix9_5_2 := prefix9_5 + "ospf.0."

	InterfaceEthernetInstanceIpv6RouterOspfAreaListCount := d.Get(prefix9_5_2 + "area_list.#").(int)
	obj9_5_2.InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDNum = make([]go_thunder.InterfaceEthernetInstanceIpv6RouterOspfAreaList, 0, InterfaceEthernetInstanceIpv6RouterOspfAreaListCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6RouterOspfAreaListCount; i++ {
		var obj9_5_2_1 go_thunder.InterfaceEthernetInstanceIpv6RouterOspfAreaList
		prefix9_5_2_1 := prefix9_5_2 + fmt.Sprintf("area_list.%d.", i)
		obj9_5_2_1.InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDNum = d.Get(prefix9_5_2_1 + "area_id_num").(int)
		obj9_5_2_1.InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDAddr = d.Get(prefix9_5_2_1 + "area_id_addr").(string)
		obj9_5_2_1.InterfaceEthernetInstanceIpv6RouterOspfAreaListTag = d.Get(prefix9_5_2_1 + "tag").(string)
		obj9_5_2_1.InterfaceEthernetInstanceIpv6RouterOspfAreaListInstanceID = d.Get(prefix9_5_2_1 + "instance_id").(int)
		obj9_5_2.InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDNum = append(obj9_5_2.InterfaceEthernetInstanceIpv6RouterOspfAreaListAreaIDNum, obj9_5_2_1)
	}

	obj9_5.InterfaceEthernetInstanceIpv6RouterOspfAreaList = obj9_5_2

	var obj9_5_3 go_thunder.InterfaceEthernetInstanceIpv6RouterIsis
	prefix9_5_3 := prefix9_5 + "isis.0."
	obj9_5_3.InterfaceEthernetInstanceIpv6RouterIsisTag = d.Get(prefix9_5_3 + "tag").(string)

	obj9_5.InterfaceEthernetInstanceIpv6RouterIsisTag = obj9_5_3

	obj9.InterfaceEthernetInstanceIpv6RouterRipng = obj9_5

	var obj9_6 go_thunder.InterfaceEthernetInstanceIpv6Rip
	prefix9_6 := prefix9 + "rip.0."

	var obj9_6_1 go_thunder.InterfaceEthernetInstanceIpv6RipSplitHorizonCfg
	prefix9_6_1 := prefix9_6 + "split_horizon_cfg.0."
	obj9_6_1.InterfaceEthernetInstanceIpv6RipSplitHorizonCfgState = d.Get(prefix9_6_1 + "state").(string)

	obj9_6.InterfaceEthernetInstanceIpv6RipSplitHorizonCfgState = obj9_6_1

	obj9.InterfaceEthernetInstanceIpv6RipSplitHorizonCfg = obj9_6

	var obj9_7 go_thunder.InterfaceEthernetInstanceIpv6Ospf
	prefix9_7 := prefix9 + "ospf.0."

	InterfaceEthernetInstanceIpv6OspfNetworkListCount := d.Get(prefix9_7 + "network_list.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfNetworkListBroadcastType = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfNetworkList, 0, InterfaceEthernetInstanceIpv6OspfNetworkListCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfNetworkListCount; i++ {
		var obj9_7_1 go_thunder.InterfaceEthernetInstanceIpv6OspfNetworkList
		prefix9_7_1 := prefix9_7 + fmt.Sprintf("network_list.%d.", i)
		obj9_7_1.InterfaceEthernetInstanceIpv6OspfNetworkListBroadcastType = d.Get(prefix9_7_1 + "broadcast_type").(string)
		obj9_7_1.InterfaceEthernetInstanceIpv6OspfNetworkListP2MpNbma = d.Get(prefix9_7_1 + "p2mp_nbma").(int)
		obj9_7_1.InterfaceEthernetInstanceIpv6OspfNetworkListNetworkInstanceID = d.Get(prefix9_7_1 + "network_instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfNetworkListBroadcastType = append(obj9_7.InterfaceEthernetInstanceIpv6OspfNetworkListBroadcastType, obj9_7_1)
	}

	obj9_7.InterfaceEthernetInstanceIpv6OspfBfd = d.Get(prefix9_7 + "bfd").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfDisable = d.Get(prefix9_7 + "disable").(int)

	InterfaceEthernetInstanceIpv6OspfCostCfgCount := d.Get(prefix9_7 + "cost_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfCostCfgCost = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfCostCfg, 0, InterfaceEthernetInstanceIpv6OspfCostCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfCostCfgCount; i++ {
		var obj9_7_2 go_thunder.InterfaceEthernetInstanceIpv6OspfCostCfg
		prefix9_7_2 := prefix9_7 + fmt.Sprintf("cost_cfg.%d.", i)
		obj9_7_2.InterfaceEthernetInstanceIpv6OspfCostCfgCost = d.Get(prefix9_7_2 + "cost").(int)
		obj9_7_2.InterfaceEthernetInstanceIpv6OspfCostCfgInstanceID = d.Get(prefix9_7_2 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfCostCfgCost = append(obj9_7.InterfaceEthernetInstanceIpv6OspfCostCfgCost, obj9_7_2)
	}

	InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgCount := d.Get(prefix9_7 + "dead_interval_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgDeadInterval = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfg, 0, InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgCount; i++ {
		var obj9_7_3 go_thunder.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfg
		prefix9_7_3 := prefix9_7 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj9_7_3.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgDeadInterval = d.Get(prefix9_7_3 + "dead_interval").(int)
		obj9_7_3.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgInstanceID = d.Get(prefix9_7_3 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgDeadInterval = append(obj9_7.InterfaceEthernetInstanceIpv6OspfDeadIntervalCfgDeadInterval, obj9_7_3)
	}

	InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgCount := d.Get(prefix9_7 + "hello_interval_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgHelloInterval = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfg, 0, InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgCount; i++ {
		var obj9_7_4 go_thunder.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfg
		prefix9_7_4 := prefix9_7 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj9_7_4.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgHelloInterval = d.Get(prefix9_7_4 + "hello_interval").(int)
		obj9_7_4.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgInstanceID = d.Get(prefix9_7_4 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgHelloInterval = append(obj9_7.InterfaceEthernetInstanceIpv6OspfHelloIntervalCfgHelloInterval, obj9_7_4)
	}

	InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgCount := d.Get(prefix9_7 + "mtu_ignore_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfg, 0, InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgCount; i++ {
		var obj9_7_5 go_thunder.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfg
		prefix9_7_5 := prefix9_7 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj9_7_5.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = d.Get(prefix9_7_5 + "mtu_ignore").(int)
		obj9_7_5.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgInstanceID = d.Get(prefix9_7_5 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = append(obj9_7.InterfaceEthernetInstanceIpv6OspfMtuIgnoreCfgMtuIgnore, obj9_7_5)
	}

	InterfaceEthernetInstanceIpv6OspfNeighborCfgCount := d.Get(prefix9_7 + "neighbor_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighbor = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfNeighborCfg, 0, InterfaceEthernetInstanceIpv6OspfNeighborCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfNeighborCfgCount; i++ {
		var obj9_7_6 go_thunder.InterfaceEthernetInstanceIpv6OspfNeighborCfg
		prefix9_7_6 := prefix9_7 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj9_7_6.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighbor = d.Get(prefix9_7_6 + "neighbor").(string)
		obj9_7_6.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeigInst = d.Get(prefix9_7_6 + "neig_inst").(int)
		obj9_7_6.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighborCost = d.Get(prefix9_7_6 + "neighbor_cost").(int)
		obj9_7_6.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighborPollInterval = d.Get(prefix9_7_6 + "neighbor_poll_interval").(int)
		obj9_7_6.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighborPriority = d.Get(prefix9_7_6 + "neighbor_priority").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighbor = append(obj9_7.InterfaceEthernetInstanceIpv6OspfNeighborCfgNeighbor, obj9_7_6)
	}

	InterfaceEthernetInstanceIpv6OspfPriorityCfgCount := d.Get(prefix9_7 + "priority_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfPriorityCfgPriority = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfPriorityCfg, 0, InterfaceEthernetInstanceIpv6OspfPriorityCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfPriorityCfgCount; i++ {
		var obj9_7_7 go_thunder.InterfaceEthernetInstanceIpv6OspfPriorityCfg
		prefix9_7_7 := prefix9_7 + fmt.Sprintf("priority_cfg.%d.", i)
		obj9_7_7.InterfaceEthernetInstanceIpv6OspfPriorityCfgPriority = d.Get(prefix9_7_7 + "priority").(int)
		obj9_7_7.InterfaceEthernetInstanceIpv6OspfPriorityCfgInstanceID = d.Get(prefix9_7_7 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfPriorityCfgPriority = append(obj9_7.InterfaceEthernetInstanceIpv6OspfPriorityCfgPriority, obj9_7_7)
	}

	InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgCount := d.Get(prefix9_7 + "retransmit_interval_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfg, 0, InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgCount; i++ {
		var obj9_7_8 go_thunder.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfg
		prefix9_7_8 := prefix9_7 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj9_7_8.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = d.Get(prefix9_7_8 + "retransmit_interval").(int)
		obj9_7_8.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgInstanceID = d.Get(prefix9_7_8 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = append(obj9_7.InterfaceEthernetInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval, obj9_7_8)
	}

	InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgCount := d.Get(prefix9_7 + "transmit_delay_cfg.#").(int)
	obj9_7.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgTransmitDelay = make([]go_thunder.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfg, 0, InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgCount)

	for i := 0; i < InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgCount; i++ {
		var obj9_7_9 go_thunder.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfg
		prefix9_7_9 := prefix9_7 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj9_7_9.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgTransmitDelay = d.Get(prefix9_7_9 + "transmit_delay").(int)
		obj9_7_9.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgInstanceID = d.Get(prefix9_7_9 + "instance_id").(int)
		obj9_7.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgTransmitDelay = append(obj9_7.InterfaceEthernetInstanceIpv6OspfTransmitDelayCfgTransmitDelay, obj9_7_9)
	}

	obj9.InterfaceEthernetInstanceIpv6OspfNetworkList = obj9_7

	c.InterfaceEthernetInstanceIpv6AddressList = obj9

	var obj10 go_thunder.InterfaceEthernetInstanceNptv6
	prefix10 := "nptv6.0."

	InterfaceEthernetInstanceNptv6DomainListCount := d.Get(prefix10 + "domain_list.#").(int)
	obj10.InterfaceEthernetInstanceNptv6DomainListDomainName = make([]go_thunder.InterfaceEthernetInstanceNptv6DomainList, 0, InterfaceEthernetInstanceNptv6DomainListCount)

	for i := 0; i < InterfaceEthernetInstanceNptv6DomainListCount; i++ {
		var obj10_1 go_thunder.InterfaceEthernetInstanceNptv6DomainList
		prefix10_1 := prefix10 + fmt.Sprintf("domain_list.%d.", i)
		obj10_1.InterfaceEthernetInstanceNptv6DomainListDomainName = d.Get(prefix10_1 + "domain_name").(string)
		obj10_1.InterfaceEthernetInstanceNptv6DomainListBindType = d.Get(prefix10_1 + "bind_type").(string)
		obj10.InterfaceEthernetInstanceNptv6DomainListDomainName = append(obj10.InterfaceEthernetInstanceNptv6DomainListDomainName, obj10_1)
	}

	c.InterfaceEthernetInstanceNptv6DomainList = obj10

	var obj11 go_thunder.InterfaceEthernetInstanceMap
	prefix11 := "map.0."
	obj11.InterfaceEthernetInstanceMapInside = d.Get(prefix11 + "inside").(int)
	obj11.InterfaceEthernetInstanceMapOutside = d.Get(prefix11 + "outside").(int)
	obj11.InterfaceEthernetInstanceMapMapTInside = d.Get(prefix11 + "map_t_inside").(int)
	obj11.InterfaceEthernetInstanceMapMapTOutside = d.Get(prefix11 + "map_t_outside").(int)

	c.InterfaceEthernetInstanceMapInside = obj11

	var obj12 go_thunder.InterfaceEthernetInstanceLw4O6
	prefix12 := "lw_4o6.0."
	obj12.InterfaceEthernetInstanceLw4O6Outside = d.Get(prefix12 + "outside").(int)
	obj12.InterfaceEthernetInstanceLw4O6Inside = d.Get(prefix12 + "inside").(int)

	c.InterfaceEthernetInstanceLw4O6Outside = obj12

	InterfaceEthernetInstanceTrunkGroupListCount := d.Get("trunk_group_list.#").(int)
	c.InterfaceEthernetInstanceTrunkGroupListTrunkNumber = make([]go_thunder.InterfaceEthernetInstanceTrunkGroupList, 0, InterfaceEthernetInstanceTrunkGroupListCount)

	for i := 0; i < InterfaceEthernetInstanceTrunkGroupListCount; i++ {
		var obj13 go_thunder.InterfaceEthernetInstanceTrunkGroupList
		prefix13 := fmt.Sprintf("trunk_group_list.%d.", i)
		obj13.InterfaceEthernetInstanceTrunkGroupListTrunkNumber = d.Get(prefix13 + "trunk_number").(int)
		obj13.InterfaceEthernetInstanceTrunkGroupListType = d.Get(prefix13 + "type").(string)
		obj13.InterfaceEthernetInstanceTrunkGroupListAdminKey = d.Get(prefix13 + "admin_key").(int)
		obj13.InterfaceEthernetInstanceTrunkGroupListPortPriority = d.Get(prefix13 + "port_priority").(int)

		var obj13_1 go_thunder.InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfg
		prefix13_1 := prefix13 + "udld_timeout_cfg.0."
		obj13_1.InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfgFast = d.Get(prefix13_1 + "fast").(int)
		obj13_1.InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfgSlow = d.Get(prefix13_1 + "slow").(int)

		obj13.InterfaceEthernetInstanceTrunkGroupListUdldTimeoutCfgFast = obj13_1

		obj13.InterfaceEthernetInstanceTrunkGroupListMode = d.Get(prefix13 + "mode").(string)
		obj13.InterfaceEthernetInstanceTrunkGroupListTimeout = d.Get(prefix13 + "timeout").(string)
		obj13.InterfaceEthernetInstanceTrunkGroupListUserTag = d.Get(prefix13 + "user_tag").(string)
		c.InterfaceEthernetInstanceTrunkGroupListTrunkNumber = append(c.InterfaceEthernetInstanceTrunkGroupListTrunkNumber, obj13)
	}

	var obj14 go_thunder.InterfaceEthernetInstanceBfd
	prefix14 := "bfd.0."

	var obj14_1 go_thunder.InterfaceEthernetInstanceBfdAuthentication
	prefix14_1 := prefix14 + "authentication.0."
	obj14_1.InterfaceEthernetInstanceBfdAuthenticationKeyID = d.Get(prefix14_1 + "key_id").(int)
	obj14_1.InterfaceEthernetInstanceBfdAuthenticationMethod = d.Get(prefix14_1 + "method").(string)
	obj14_1.InterfaceEthernetInstanceBfdAuthenticationPassword = d.Get(prefix14_1 + "password").(string)

	obj14.InterfaceEthernetInstanceBfdAuthenticationKeyID = obj14_1

	obj14.InterfaceEthernetInstanceBfdEcho = d.Get(prefix14 + "echo").(int)
	obj14.InterfaceEthernetInstanceBfdDemand = d.Get(prefix14 + "demand").(int)

	var obj14_2 go_thunder.InterfaceEthernetInstanceBfdIntervalCfg
	prefix14_2 := prefix14 + "interval_cfg.0."
	obj14_2.InterfaceEthernetInstanceBfdIntervalCfgInterval = d.Get(prefix14_2 + "interval").(int)
	obj14_2.InterfaceEthernetInstanceBfdIntervalCfgMinRx = d.Get(prefix14_2 + "min_rx").(int)
	obj14_2.InterfaceEthernetInstanceBfdIntervalCfgMultiplier = d.Get(prefix14_2 + "multiplier").(int)

	obj14.InterfaceEthernetInstanceBfdIntervalCfgInterval = obj14_2

	c.InterfaceEthernetInstanceBfdAuthentication = obj14

	var obj15 go_thunder.InterfaceEthernetInstanceIsis
	prefix15 := "isis.0."

	var obj15_1 go_thunder.InterfaceEthernetInstanceIsisAuthentication
	prefix15_1 := prefix15 + "authentication.0."

	InterfaceEthernetInstanceIsisAuthenticationSendOnlyListCount := d.Get(prefix15_1 + "send_only_list.#").(int)
	obj15_1.InterfaceEthernetInstanceIsisAuthenticationSendOnlyListSendOnly = make([]go_thunder.InterfaceEthernetInstanceIsisAuthenticationSendOnlyList, 0, InterfaceEthernetInstanceIsisAuthenticationSendOnlyListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisAuthenticationSendOnlyListCount; i++ {
		var obj15_1_1 go_thunder.InterfaceEthernetInstanceIsisAuthenticationSendOnlyList
		prefix15_1_1 := prefix15_1 + fmt.Sprintf("send_only_list.%d.", i)
		obj15_1_1.InterfaceEthernetInstanceIsisAuthenticationSendOnlyListSendOnly = d.Get(prefix15_1_1 + "send_only").(int)
		obj15_1_1.InterfaceEthernetInstanceIsisAuthenticationSendOnlyListLevel = d.Get(prefix15_1_1 + "level").(string)
		obj15_1.InterfaceEthernetInstanceIsisAuthenticationSendOnlyListSendOnly = append(obj15_1.InterfaceEthernetInstanceIsisAuthenticationSendOnlyListSendOnly, obj15_1_1)
	}

	InterfaceEthernetInstanceIsisAuthenticationModeListCount := d.Get(prefix15_1 + "mode_list.#").(int)
	obj15_1.InterfaceEthernetInstanceIsisAuthenticationModeListMode = make([]go_thunder.InterfaceEthernetInstanceIsisAuthenticationModeList, 0, InterfaceEthernetInstanceIsisAuthenticationModeListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisAuthenticationModeListCount; i++ {
		var obj15_1_2 go_thunder.InterfaceEthernetInstanceIsisAuthenticationModeList
		prefix15_1_2 := prefix15_1 + fmt.Sprintf("mode_list.%d.", i)
		obj15_1_2.InterfaceEthernetInstanceIsisAuthenticationModeListMode = d.Get(prefix15_1_2 + "mode").(string)
		obj15_1_2.InterfaceEthernetInstanceIsisAuthenticationModeListLevel = d.Get(prefix15_1_2 + "level").(string)
		obj15_1.InterfaceEthernetInstanceIsisAuthenticationModeListMode = append(obj15_1.InterfaceEthernetInstanceIsisAuthenticationModeListMode, obj15_1_2)
	}

	InterfaceEthernetInstanceIsisAuthenticationKeyChainListCount := d.Get(prefix15_1 + "key_chain_list.#").(int)
	obj15_1.InterfaceEthernetInstanceIsisAuthenticationKeyChainListKeyChain = make([]go_thunder.InterfaceEthernetInstanceIsisAuthenticationKeyChainList, 0, InterfaceEthernetInstanceIsisAuthenticationKeyChainListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisAuthenticationKeyChainListCount; i++ {
		var obj15_1_3 go_thunder.InterfaceEthernetInstanceIsisAuthenticationKeyChainList
		prefix15_1_3 := prefix15_1 + fmt.Sprintf("key_chain_list.%d.", i)
		obj15_1_3.InterfaceEthernetInstanceIsisAuthenticationKeyChainListKeyChain = d.Get(prefix15_1_3 + "key_chain").(string)
		obj15_1_3.InterfaceEthernetInstanceIsisAuthenticationKeyChainListLevel = d.Get(prefix15_1_3 + "level").(string)
		obj15_1.InterfaceEthernetInstanceIsisAuthenticationKeyChainListKeyChain = append(obj15_1.InterfaceEthernetInstanceIsisAuthenticationKeyChainListKeyChain, obj15_1_3)
	}

	obj15.InterfaceEthernetInstanceIsisAuthenticationSendOnlyList = obj15_1

	var obj15_2 go_thunder.InterfaceEthernetInstanceIsisBfdCfg
	prefix15_2 := prefix15 + "bfd_cfg.0."
	obj15_2.InterfaceEthernetInstanceIsisBfdCfgBfd = d.Get(prefix15_2 + "bfd").(int)
	obj15_2.InterfaceEthernetInstanceIsisBfdCfgDisable = d.Get(prefix15_2 + "disable").(int)

	obj15.InterfaceEthernetInstanceIsisBfdCfgBfd = obj15_2

	obj15.InterfaceEthernetInstanceIsisCircuitType = d.Get(prefix15 + "circuit_type").(string)

	InterfaceEthernetInstanceIsisCsnpIntervalListCount := d.Get(prefix15 + "csnp_interval_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisCsnpIntervalListCsnpInterval = make([]go_thunder.InterfaceEthernetInstanceIsisCsnpIntervalList, 0, InterfaceEthernetInstanceIsisCsnpIntervalListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisCsnpIntervalListCount; i++ {
		var obj15_3 go_thunder.InterfaceEthernetInstanceIsisCsnpIntervalList
		prefix15_3 := prefix15 + fmt.Sprintf("csnp_interval_list.%d.", i)
		obj15_3.InterfaceEthernetInstanceIsisCsnpIntervalListCsnpInterval = d.Get(prefix15_3 + "csnp_interval").(int)
		obj15_3.InterfaceEthernetInstanceIsisCsnpIntervalListLevel = d.Get(prefix15_3 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisCsnpIntervalListCsnpInterval = append(obj15.InterfaceEthernetInstanceIsisCsnpIntervalListCsnpInterval, obj15_3)
	}

	obj15.InterfaceEthernetInstanceIsisPadding = d.Get(prefix15 + "padding").(int)

	InterfaceEthernetInstanceIsisHelloIntervalListCount := d.Get(prefix15 + "hello_interval_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisHelloIntervalListHelloInterval = make([]go_thunder.InterfaceEthernetInstanceIsisHelloIntervalList, 0, InterfaceEthernetInstanceIsisHelloIntervalListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisHelloIntervalListCount; i++ {
		var obj15_4 go_thunder.InterfaceEthernetInstanceIsisHelloIntervalList
		prefix15_4 := prefix15 + fmt.Sprintf("hello_interval_list.%d.", i)
		obj15_4.InterfaceEthernetInstanceIsisHelloIntervalListHelloInterval = d.Get(prefix15_4 + "hello_interval").(int)
		obj15_4.InterfaceEthernetInstanceIsisHelloIntervalListLevel = d.Get(prefix15_4 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisHelloIntervalListHelloInterval = append(obj15.InterfaceEthernetInstanceIsisHelloIntervalListHelloInterval, obj15_4)
	}

	InterfaceEthernetInstanceIsisHelloIntervalMinimalListCount := d.Get(prefix15 + "hello_interval_minimal_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = make([]go_thunder.InterfaceEthernetInstanceIsisHelloIntervalMinimalList, 0, InterfaceEthernetInstanceIsisHelloIntervalMinimalListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisHelloIntervalMinimalListCount; i++ {
		var obj15_5 go_thunder.InterfaceEthernetInstanceIsisHelloIntervalMinimalList
		prefix15_5 := prefix15 + fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj15_5.InterfaceEthernetInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = d.Get(prefix15_5 + "hello_interval_minimal").(int)
		obj15_5.InterfaceEthernetInstanceIsisHelloIntervalMinimalListLevel = d.Get(prefix15_5 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = append(obj15.InterfaceEthernetInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal, obj15_5)
	}

	InterfaceEthernetInstanceIsisHelloMultiplierListCount := d.Get(prefix15 + "hello_multiplier_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisHelloMultiplierListHelloMultiplier = make([]go_thunder.InterfaceEthernetInstanceIsisHelloMultiplierList, 0, InterfaceEthernetInstanceIsisHelloMultiplierListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisHelloMultiplierListCount; i++ {
		var obj15_6 go_thunder.InterfaceEthernetInstanceIsisHelloMultiplierList
		prefix15_6 := prefix15 + fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj15_6.InterfaceEthernetInstanceIsisHelloMultiplierListHelloMultiplier = d.Get(prefix15_6 + "hello_multiplier").(int)
		obj15_6.InterfaceEthernetInstanceIsisHelloMultiplierListLevel = d.Get(prefix15_6 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisHelloMultiplierListHelloMultiplier = append(obj15.InterfaceEthernetInstanceIsisHelloMultiplierListHelloMultiplier, obj15_6)
	}

	obj15.InterfaceEthernetInstanceIsisLspInterval = d.Get(prefix15 + "lsp_interval").(int)

	var obj15_7 go_thunder.InterfaceEthernetInstanceIsisMeshGroup
	prefix15_7 := prefix15 + "mesh_group.0."
	obj15_7.InterfaceEthernetInstanceIsisMeshGroupValue = d.Get(prefix15_7 + "value").(int)
	obj15_7.InterfaceEthernetInstanceIsisMeshGroupBlocked = d.Get(prefix15_7 + "blocked").(int)

	obj15.InterfaceEthernetInstanceIsisMeshGroupValue = obj15_7

	InterfaceEthernetInstanceIsisMetricListCount := d.Get(prefix15 + "metric_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisMetricListMetric = make([]go_thunder.InterfaceEthernetInstanceIsisMetricList, 0, InterfaceEthernetInstanceIsisMetricListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisMetricListCount; i++ {
		var obj15_8 go_thunder.InterfaceEthernetInstanceIsisMetricList
		prefix15_8 := prefix15 + fmt.Sprintf("metric_list.%d.", i)
		obj15_8.InterfaceEthernetInstanceIsisMetricListMetric = d.Get(prefix15_8 + "metric").(int)
		obj15_8.InterfaceEthernetInstanceIsisMetricListLevel = d.Get(prefix15_8 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisMetricListMetric = append(obj15.InterfaceEthernetInstanceIsisMetricListMetric, obj15_8)
	}

	obj15.InterfaceEthernetInstanceIsisNetwork = d.Get(prefix15 + "network").(string)

	InterfaceEthernetInstanceIsisPasswordListCount := d.Get(prefix15 + "password_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisPasswordListPassword = make([]go_thunder.InterfaceEthernetInstanceIsisPasswordList, 0, InterfaceEthernetInstanceIsisPasswordListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisPasswordListCount; i++ {
		var obj15_9 go_thunder.InterfaceEthernetInstanceIsisPasswordList
		prefix15_9 := prefix15 + fmt.Sprintf("password_list.%d.", i)
		obj15_9.InterfaceEthernetInstanceIsisPasswordListPassword = d.Get(prefix15_9 + "password").(string)
		obj15_9.InterfaceEthernetInstanceIsisPasswordListLevel = d.Get(prefix15_9 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisPasswordListPassword = append(obj15.InterfaceEthernetInstanceIsisPasswordListPassword, obj15_9)
	}

	InterfaceEthernetInstanceIsisPriorityListCount := d.Get(prefix15 + "priority_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisPriorityListPriority = make([]go_thunder.InterfaceEthernetInstanceIsisPriorityList, 0, InterfaceEthernetInstanceIsisPriorityListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisPriorityListCount; i++ {
		var obj15_10 go_thunder.InterfaceEthernetInstanceIsisPriorityList
		prefix15_10 := prefix15 + fmt.Sprintf("priority_list.%d.", i)
		obj15_10.InterfaceEthernetInstanceIsisPriorityListPriority = d.Get(prefix15_10 + "priority").(int)
		obj15_10.InterfaceEthernetInstanceIsisPriorityListLevel = d.Get(prefix15_10 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisPriorityListPriority = append(obj15.InterfaceEthernetInstanceIsisPriorityListPriority, obj15_10)
	}

	obj15.InterfaceEthernetInstanceIsisRetransmitInterval = d.Get(prefix15 + "retransmit_interval").(int)

	InterfaceEthernetInstanceIsisWideMetricListCount := d.Get(prefix15 + "wide_metric_list.#").(int)
	obj15.InterfaceEthernetInstanceIsisWideMetricListWideMetric = make([]go_thunder.InterfaceEthernetInstanceIsisWideMetricList, 0, InterfaceEthernetInstanceIsisWideMetricListCount)

	for i := 0; i < InterfaceEthernetInstanceIsisWideMetricListCount; i++ {
		var obj15_11 go_thunder.InterfaceEthernetInstanceIsisWideMetricList
		prefix15_11 := prefix15 + fmt.Sprintf("wide_metric_list.%d.", i)
		obj15_11.InterfaceEthernetInstanceIsisWideMetricListWideMetric = d.Get(prefix15_11 + "wide_metric").(int)
		obj15_11.InterfaceEthernetInstanceIsisWideMetricListLevel = d.Get(prefix15_11 + "level").(string)
		obj15.InterfaceEthernetInstanceIsisWideMetricListWideMetric = append(obj15.InterfaceEthernetInstanceIsisWideMetricListWideMetric, obj15_11)
	}

	c.InterfaceEthernetInstanceIsisAuthentication = obj15

	var obj16 go_thunder.InterfaceEthernetInstanceSpanningTree
	prefix16 := "spanning_tree.0."
	obj16.InterfaceEthernetInstanceSpanningTreeAutoEdge = d.Get(prefix16 + "auto_edge").(int)
	obj16.InterfaceEthernetInstanceSpanningTreeAdminEdge = d.Get(prefix16 + "admin_edge").(int)

	InterfaceEthernetInstanceSpanningTreeInstanceListCount := d.Get(prefix16 + "instance_list.#").(int)
	obj16.InterfaceEthernetInstanceSpanningTreeInstanceListInstanceStart = make([]go_thunder.InterfaceEthernetInstanceSpanningTreeInstanceList, 0, InterfaceEthernetInstanceSpanningTreeInstanceListCount)

	for i := 0; i < InterfaceEthernetInstanceSpanningTreeInstanceListCount; i++ {
		var obj16_1 go_thunder.InterfaceEthernetInstanceSpanningTreeInstanceList
		prefix16_1 := prefix16 + fmt.Sprintf("instance_list.%d.", i)
		obj16_1.InterfaceEthernetInstanceSpanningTreeInstanceListInstanceStart = d.Get(prefix16_1 + "instance_start").(int)
		obj16_1.InterfaceEthernetInstanceSpanningTreeInstanceListMstpPathCost = d.Get(prefix16_1 + "mstp_path_cost").(int)
		obj16.InterfaceEthernetInstanceSpanningTreeInstanceListInstanceStart = append(obj16.InterfaceEthernetInstanceSpanningTreeInstanceListInstanceStart, obj16_1)
	}

	obj16.InterfaceEthernetInstanceSpanningTreePathCost = d.Get(prefix16 + "path_cost").(int)

	c.InterfaceEthernetInstanceSpanningTreeAutoEdge = obj16

	vc.InterfaceEthernetInstanceIfnum = c
	return vc
}
