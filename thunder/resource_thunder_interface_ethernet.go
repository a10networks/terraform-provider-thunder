package thunder

//Thunder resource InterfaceEthernet

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"strconv"
	"util"
)

func resourceInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceEthernetCreate,
		UpdateContext: resourceInterfaceEthernetUpdate,
		ReadContext:   resourceInterfaceEthernetRead,
		DeleteContext: resourceInterfaceEthernetDelete,
		Schema: map[string]*schema.Schema{
			"ifnum": {
				Type: schema.TypeInt, Required: true, ForceNew: true,
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
									"ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv4_netmask": {
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
		data := dataToEndpointInterfaceEthernet(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceEthernet --")
		d.SetId(strconv.Itoa(name1))
		err := PostInterfaceEthernet(client.Token, data, client.Host)
		if err != nil {
			logger.Println("Ethernet Object Axapi through error ---> " + string(err.Error()))
			logger.Println("checking for Ethernet Instance AxAPI")
			err1 := PostInterfaceEthernetIfnum(client.Token, strconv.Itoa(name1), data, client.Host)
			if err1 != nil {
				err := errors.New("Ethernet object error -->" + string(err.Error()) + "\n Trying Ethernet instance error -->" + string(err1.Error()))
				return diag.FromErr(err)
			}
			return nil
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
		data, err := GetInterfaceEthernet(client.Token, name1, client.Host)
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
		data := dataToEndpointInterfaceEthernet(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceEthernet ")
		err := PutInterfaceEthernet(client.Token, name1, data, client.Host)
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
		err := DeleteInterfaceEthernet(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToEndpointInterfaceEthernet(d *schema.ResourceData) EndpointInterfaceEthernet {
	var c EndpointInterfaceEthernet
	c.Inst.Ifnum = d.Get("ifnum").(int)
	c.Inst.Name = d.Get("name").(string)
	c.Inst.PortScanDetection = d.Get("port_scan_detection").(string)
	c.Inst.PingSweepDetection = d.Get("ping_sweep_detection").(string)
	c.Inst.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	c.Inst.LoadInterval = d.Get("load_interval").(int)
	c.Inst.MediaTypeCopper = d.Get("media_type_copper").(int)
	c.Inst.AutoNegEnable = d.Get("auto_neg_enable").(int)
	c.Inst.FecForcedOn = d.Get("fec_forced_on").(int)
	c.Inst.FecForcedOff = d.Get("fec_forced_off").(int)
	c.Inst.PortBreakout = d.Get("port_breakout").(string)
	c.Inst.SpeedForced1g = d.Get("speed_forced_1g").(int)
	c.Inst.SpeedForced10g = d.Get("speed_forced_10g").(int)
	c.Inst.SpeedForced40g = d.Get("speed_forced_40g").(int)
	c.Inst.IpgBitTime = d.Get("ipg_bit_time").(int)
	c.Inst.RemoveVlanTag = d.Get("remove_vlan_tag").(int)
	c.Inst.Mtu = d.Get("mtu").(int)
	c.Inst.TrapSource = d.Get("trap_source").(int)
	c.Inst.Duplexity = d.Get("duplexity").(string)
	c.Inst.Speed = d.Get("speed").(string)
	c.Inst.FlowControl = d.Get("flow_control").(int)
	c.Inst.Action = d.Get("action").(string)

	var obj1 InterfaceEthernetIcmpRateLimit
	prefix1 := "icmp_rate_limit.0."
	obj1.Normal = d.Get(prefix1 + "normal").(int)
	obj1.Lockup = d.Get(prefix1 + "lockup").(int)
	obj1.LockupPeriod = d.Get(prefix1 + "lockup_period").(int)
	c.Inst.IcmpRateLimit = obj1

	var obj2 InterfaceEthernetIcmpv6RateLimit
	prefix2 := "icmpv6_rate_limit.0."
	obj2.NormalV6 = d.Get(prefix2 + "normal_v6").(int)
	obj2.LockupV6 = d.Get(prefix2 + "lockup_v6").(int)
	obj2.LockupPeriodV6 = d.Get(prefix2 + "lockup_period_v6").(int)
	c.Inst.Icmpv6RateLimit = obj2

	InterfaceEthernetMonitorListCount := d.Get("monitor_list.#").(int)
	c.Inst.MonitorList = make([]InterfaceEthernetMonitorList, 0, InterfaceEthernetMonitorListCount)
	for i := 0; i < InterfaceEthernetMonitorListCount; i++ {
		var obj3 InterfaceEthernetMonitorList
		prefix3 := fmt.Sprintf("monitor_list.%d.", i)
		obj3.Monitor = d.Get(prefix3 + "monitor").(string)
		obj3.MirrorIndex = d.Get(prefix3 + "mirror_index").(int)
		obj3.MonitorVlan = d.Get(prefix3 + "monitor_vlan").(int)
		c.Inst.MonitorList = append(c.Inst.MonitorList, obj3)
	}

	c.Inst.CpuProcess = d.Get("cpu_process").(int)
	c.Inst.CpuProcessDir = d.Get("cpu_process_dir").(string)
	c.Inst.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)
	c.Inst.VirtualWire = d.Get("virtual_wire").(int)
	c.Inst.UpdateL2Info = d.Get("update_l2_info").(int)
	c.Inst.VlanLearning = d.Get("vlan_learning").(string)
	c.Inst.MacLearning = d.Get("mac_learning").(string)

	var obj4 InterfaceEthernetAccessList
	prefix4 := "access_list.0."
	obj4.AclId = d.Get(prefix4 + "acl_id").(int)
	obj4.AclName = d.Get(prefix4 + "acl_name").(string)
	c.Inst.AccessList = obj4

	c.Inst.UserTag = d.Get("user_tag").(string)

	InterfaceEthernetSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Inst.SamplingEnable = make([]InterfaceEthernetSamplingEnable, 0, InterfaceEthernetSamplingEnableCount)
	for i := 0; i < InterfaceEthernetSamplingEnableCount; i++ {
		var obj5 InterfaceEthernetSamplingEnable
		prefix5 := fmt.Sprintf("sampling_enable.%d.", i)
		obj5.Counters1 = d.Get(prefix5 + "counters1").(string)
		c.Inst.SamplingEnable = append(c.Inst.SamplingEnable, obj5)
	}

	c.Inst.PacketCaptureTemplate = d.Get("packet_capture_template").(string)

	var obj6 InterfaceEthernetLldp
	prefix6 := "lldp.0."

	var obj6_1 InterfaceEthernetLldpEnableCfg
	prefix6_1 := prefix6 + "enable_cfg.0."
	obj6_1.RtEnable = d.Get(prefix6_1 + "rt_enable").(int)
	obj6_1.Rx = d.Get(prefix6_1 + "rx").(int)
	obj6_1.Tx = d.Get(prefix6_1 + "tx").(int)
	obj6.EnableCfg = obj6_1

	var obj6_2 InterfaceEthernetLldpNotificationCfg
	prefix6_2 := prefix6 + "notification_cfg.0."
	obj6_2.Notification = d.Get(prefix6_2 + "notification").(int)
	obj6_2.NotifEnable = d.Get(prefix6_2 + "notif_enable").(int)
	obj6.NotificationCfg = obj6_2

	var obj6_3 InterfaceEthernetLldpTxDot1Cfg
	prefix6_3 := prefix6 + "tx_dot1_cfg.0."
	obj6_3.TxDot1Tlvs = d.Get(prefix6_3 + "tx_dot1_tlvs").(int)
	obj6_3.LinkAggregation = d.Get(prefix6_3 + "link_aggregation").(int)
	obj6_3.Vlan = d.Get(prefix6_3 + "vlan").(int)
	obj6.TxDot1Cfg = obj6_3

	var obj6_4 InterfaceEthernetLldpTxTlvsCfg
	prefix6_4 := prefix6 + "tx_tlvs_cfg.0."
	obj6_4.TxTlvs = d.Get(prefix6_4 + "tx_tlvs").(int)
	obj6_4.Exclude = d.Get(prefix6_4 + "exclude").(int)
	obj6_4.ManagementAddress = d.Get(prefix6_4 + "management_address").(int)
	obj6_4.PortDescription = d.Get(prefix6_4 + "port_description").(int)
	obj6_4.SystemCapabilities = d.Get(prefix6_4 + "system_capabilities").(int)
	obj6_4.SystemDescription = d.Get(prefix6_4 + "system_description").(int)
	obj6_4.SystemName = d.Get(prefix6_4 + "system_name").(int)
	obj6.TxTlvsCfg = obj6_4

	c.Inst.Lldp = obj6

	var obj7 InterfaceEthernetDdos
	prefix7 := "ddos.0."
	obj7.Outside = d.Get(prefix7 + "outside").(int)
	obj7.Inside = d.Get(prefix7 + "inside").(int)
	c.Inst.Ddos = obj7

	var obj8 InterfaceEthernetIp
	prefix8 := "ip.0."
	obj8.Dhcp = d.Get(prefix8 + "dhcp").(int)

	InterfaceEthernetIpAddressListCount := d.Get(prefix8 + "address_list.#").(int)
	obj8.AddressList = make([]InterfaceEthernetIpAddressList, 0, InterfaceEthernetIpAddressListCount)
	for i := 0; i < InterfaceEthernetIpAddressListCount; i++ {
		var obj8_1 InterfaceEthernetIpAddressList
		prefix8_1 := prefix8 + fmt.Sprintf("address_list.%d.", i)
		obj8_1.Ipv4Address = d.Get(prefix8_1 + "ipv4_address").(string)
		obj8_1.Ipv4Netmask = d.Get(prefix8_1 + "ipv4_netmask").(string)
		obj8.AddressList = append(obj8.AddressList, obj8_1)
	}

	obj8.AllowPromiscuousVip = d.Get(prefix8 + "allow_promiscuous_vip").(int)
	obj8.CacheSpoofingPort = d.Get(prefix8 + "cache_spoofing_port").(int)

	InterfaceEthernetIpHelperAddressListCount := d.Get(prefix8 + "helper_address_list.#").(int)
	obj8.HelperAddressList = make([]InterfaceEthernetIpHelperAddressList, 0, InterfaceEthernetIpHelperAddressListCount)

	for i := 0; i < InterfaceEthernetIpHelperAddressListCount; i++ {
		var obj8_2 InterfaceEthernetIpHelperAddressList
		prefix8_2 := prefix8 + fmt.Sprintf("helper_address_list.%d.", i)
		obj8_2.HelperAddress = d.Get(prefix8_2 + "helper_address").(string)
		obj8.HelperAddressList = append(obj8.HelperAddressList, obj8_2)
	}

	obj8.Inside = d.Get(prefix8 + "inside").(int)
	obj8.Outside = d.Get(prefix8 + "outside").(int)
	obj8.TtlIgnore = d.Get(prefix8 + "ttl_ignore").(int)
	obj8.SynCookie = d.Get(prefix8 + "syn_cookie").(int)
	obj8.SlbPartitionRedirect = d.Get(prefix8 + "slb_partition_redirect").(int)
	obj8.GenerateMembershipQuery = d.Get(prefix8 + "generate_membership_query").(int)
	obj8.QueryInterval = d.Get(prefix8 + "query_interval").(int)
	obj8.MaxRespTime = d.Get(prefix8 + "max_resp_time").(int)
	obj8.Client = d.Get(prefix8 + "client").(int)
	obj8.Server = d.Get(prefix8 + "server").(int)

	var obj8_3 InterfaceEthernetIpStatefulFirewall
	prefix8_3 := prefix8 + "stateful_firewall.0."
	obj8_3.Inside = d.Get(prefix8_3 + "inside").(int)
	obj8_3.ClassList = d.Get(prefix8_3 + "class_list").(string)
	obj8_3.Outside = d.Get(prefix8_3 + "outside").(int)
	obj8_3.AccessList = d.Get(prefix8_3 + "access_list").(int)
	obj8_3.AclId = d.Get(prefix8_3 + "acl_id").(int)
	obj8.StatefulFirewall = obj8_3

	var obj8_4 InterfaceEthernetIpRouter
	prefix8_4 := prefix8 + "router.0."

	var obj8_4_1 InterfaceEthernetIpRouterIsis
	prefix8_4_1 := prefix8_4 + "isis.0."
	obj8_4_1.Tag = d.Get(prefix8_4_1 + "tag").(string)
	obj8_4.Isis = obj8_4_1

	obj8.Router = obj8_4

	var obj8_5 InterfaceEthernetIpRip
	prefix8_5 := prefix8 + "rip.0."

	var obj8_5_1 InterfaceEthernetIpRipAuthentication
	prefix8_5_1 := prefix8_5 + "authentication.0."

	var obj8_5_1_1 InterfaceEthernetIpRipAuthenticationStr
	prefix8_5_1_1 := prefix8_5_1 + "str.0."
	obj8_5_1_1.String = d.Get(prefix8_5_1_1 + "string").(string)
	obj8_5_1.Str = obj8_5_1_1

	var obj8_5_1_2 InterfaceEthernetIpRipAuthenticationMode
	prefix8_5_1_2 := prefix8_5_1 + "mode.0."
	obj8_5_1_2.Mode = d.Get(prefix8_5_1_2 + "mode").(string)
	obj8_5_1.Mode = obj8_5_1_2

	var obj8_5_1_3 InterfaceEthernetIpRipAuthenticationKeyChain
	prefix8_5_1_3 := prefix8_5_1 + "key_chain.0."
	obj8_5_1_3.KeyChain = d.Get(prefix8_5_1_3 + "key_chain").(string)
	obj8_5_1.KeyChain = obj8_5_1_3

	obj8_5.Authentication = obj8_5_1

	obj8_5.SendPacket = d.Get(prefix8_5 + "send_packet").(int)
	obj8_5.ReceivePacket = d.Get(prefix8_5 + "receive_packet").(int)

	var obj8_5_2 InterfaceEthernetIpRipSendCfg
	prefix8_5_2 := prefix8_5 + "send_cfg.0."
	obj8_5_2.Send = d.Get(prefix8_5_2 + "send").(int)
	obj8_5_2.Version = d.Get(prefix8_5_2 + "version").(string)
	obj8_5.SendCfg = obj8_5_2

	var obj8_5_3 InterfaceEthernetIpRipReceiveCfg
	prefix8_5_3 := prefix8_5 + "receive_cfg.0."
	obj8_5_3.Receive = d.Get(prefix8_5_3 + "receive").(int)
	obj8_5_3.Version = d.Get(prefix8_5_3 + "version").(string)
	obj8_5.ReceiveCfg = obj8_5_3

	var obj8_5_4 InterfaceEthernetIpRipSplitHorizonCfg
	prefix8_5_4 := prefix8_5 + "split_horizon_cfg.0."
	obj8_5_4.State = d.Get(prefix8_5_4 + "state").(string)
	obj8_5.SplitHorizonCfg = obj8_5_4

	obj8.Rip = obj8_5

	var obj8_6 InterfaceEthernetIpOspf
	prefix8_6 := prefix8 + "ospf.0."

	var obj8_6_1 InterfaceEthernetIpOspfOspfGlobal
	prefix8_6_1 := prefix8_6 + "ospf_global.0."

	var obj8_6_1_1 InterfaceEthernetIpOspfOspfGlobalAuthenticationCfg
	prefix8_6_1_1 := prefix8_6_1 + "authentication_cfg.0."
	obj8_6_1_1.Authentication = d.Get(prefix8_6_1_1 + "authentication").(int)
	obj8_6_1_1.Value = d.Get(prefix8_6_1_1 + "value").(string)
	obj8_6_1.AuthenticationCfg = obj8_6_1_1

	obj8_6_1.AuthenticationKey = d.Get(prefix8_6_1 + "authentication_key").(string)

	var obj8_6_1_2 InterfaceEthernetIpOspfOspfGlobalBfdCfg
	prefix8_6_1_2 := prefix8_6_1 + "bfd_cfg.0."
	obj8_6_1_2.Bfd = d.Get(prefix8_6_1_2 + "bfd").(int)
	obj8_6_1_2.Disable = d.Get(prefix8_6_1_2 + "disable").(int)
	obj8_6_1.BfdCfg = obj8_6_1_2

	obj8_6_1.Cost = d.Get(prefix8_6_1 + "cost").(int)

	var obj8_6_1_3 InterfaceEthernetIpOspfOspfGlobalDatabaseFilterCfg
	prefix8_6_1_3 := prefix8_6_1 + "database_filter_cfg.0."
	obj8_6_1_3.DatabaseFilter = d.Get(prefix8_6_1_3 + "database_filter").(string)
	obj8_6_1_3.Out = d.Get(prefix8_6_1_3 + "out").(int)
	obj8_6_1.DatabaseFilterCfg = obj8_6_1_3

	obj8_6_1.DeadInterval = d.Get(prefix8_6_1 + "dead_interval").(int)
	obj8_6_1.Disable = d.Get(prefix8_6_1 + "disable").(string)
	obj8_6_1.HelloInterval = d.Get(prefix8_6_1 + "hello_interval").(int)

	InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgCount := d.Get(prefix8_6_1 + "message_digest_cfg.#").(int)
	obj8_6_1.MessageDigestCfg = make([]InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg, 0, InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgCount)
	for i := 0; i < InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgCount; i++ {
		var obj8_6_1_4 InterfaceEthernetIpOspfOspfGlobalMessageDigestCfg
		prefix8_6_1_4 := prefix8_6_1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj8_6_1_4.MessageDigestKey = d.Get(prefix8_6_1_4 + "message_digest_key").(int)
		obj8_6_1_4.Md5 = InterfaceEthernetIpOspfOspfGlobalMessageDigestCfgMd5{Md5Value: d.Get(prefix8_6_1_4 + "md5_value").(string)}
		obj8_6_1.MessageDigestCfg = append(obj8_6_1.MessageDigestCfg, obj8_6_1_4)
	}

	obj8_6_1.Mtu = d.Get(prefix8_6_1 + "mtu").(int)
	obj8_6_1.MtuIgnore = d.Get(prefix8_6_1 + "mtu_ignore").(int)

	var obj8_6_1_5 InterfaceEthernetIpOspfOspfGlobalNetwork
	prefix8_6_1_5 := prefix8_6_1 + "network.0."
	obj8_6_1_5.Broadcast = d.Get(prefix8_6_1_5 + "broadcast").(int)
	obj8_6_1_5.NonBroadcast = d.Get(prefix8_6_1_5 + "non_broadcast").(int)
	obj8_6_1_5.PointToPoint = d.Get(prefix8_6_1_5 + "point_to_point").(int)
	obj8_6_1_5.PointToMultipoint = d.Get(prefix8_6_1_5 + "point_to_multipoint").(int)
	obj8_6_1_5.P2mpNbma = d.Get(prefix8_6_1_5 + "p2mp_nbma").(int)
	obj8_6_1.Network = obj8_6_1_5

	obj8_6_1.Priority = d.Get(prefix8_6_1 + "priority").(int)
	obj8_6_1.RetransmitInterval = d.Get(prefix8_6_1 + "retransmit_interval").(int)
	obj8_6_1.TransmitDelay = d.Get(prefix8_6_1 + "transmit_delay").(int)

	obj8_6.OspfGlobal = obj8_6_1

	InterfaceEthernetIpOspfOspfIpListCount := d.Get(prefix8_6 + "ospf_ip_list.#").(int)
	obj8_6.OspfIpList = make([]InterfaceEthernetIpOspfOspfIpList, 0, InterfaceEthernetIpOspfOspfIpListCount)
	for i := 0; i < InterfaceEthernetIpOspfOspfIpListCount; i++ {
		var obj8_6_2 InterfaceEthernetIpOspfOspfIpList
		prefix8_6_2 := prefix8_6 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj8_6_2.IpAddr = d.Get(prefix8_6_2 + "ip_addr").(string)
		obj8_6_2.Authentication = d.Get(prefix8_6_2 + "authentication").(int)
		obj8_6_2.Value = d.Get(prefix8_6_2 + "value").(string)
		obj8_6_2.AuthenticationKey = d.Get(prefix8_6_2 + "authentication_key").(string)
		obj8_6_2.Cost = d.Get(prefix8_6_2 + "cost").(int)
		obj8_6_2.DatabaseFilter = d.Get(prefix8_6_2 + "database_filter").(string)
		obj8_6_2.Out = d.Get(prefix8_6_2 + "out").(int)
		obj8_6_2.DeadInterval = d.Get(prefix8_6_2 + "dead_interval").(int)
		obj8_6_2.HelloInterval = d.Get(prefix8_6_2 + "hello_interval").(int)

		InterfaceEthernetIpOspfOspfIpListMessageDigestCfgCount := d.Get(prefix8_6_2 + "message_digest_cfg.#").(int)
		obj8_6_2.MessageDigestCfg = make([]InterfaceEthernetIpOspfOspfIpListMessageDigestCfg, 0, InterfaceEthernetIpOspfOspfIpListMessageDigestCfgCount)
		for i := 0; i < InterfaceEthernetIpOspfOspfIpListMessageDigestCfgCount; i++ {
			var obj8_6_2_1 InterfaceEthernetIpOspfOspfIpListMessageDigestCfg
			prefix8_6_2_1 := prefix8_6_2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj8_6_2_1.MessageDigestKey = d.Get(prefix8_6_2_1 + "message_digest_key").(int)
			obj8_6_2_1.Md5Value = d.Get(prefix8_6_2_1 + "md5_value").(string)
			obj8_6_2.MessageDigestCfg = append(obj8_6_2.MessageDigestCfg, obj8_6_2_1)
		}

		obj8_6_2.MtuIgnore = d.Get(prefix8_6_2 + "mtu_ignore").(int)
		obj8_6_2.Priority = d.Get(prefix8_6_2 + "priority").(int)
		obj8_6_2.RetransmitInterval = d.Get(prefix8_6_2 + "retransmit_interval").(int)
		obj8_6_2.TransmitDelay = d.Get(prefix8_6_2 + "transmit_delay").(int)
		obj8_6.OspfIpList = append(obj8_6.OspfIpList, obj8_6_2)
	}

	obj8.Ospf = obj8_6

	c.Inst.Ip = obj8

	var obj9 InterfaceEthernetIpv6
	prefix9 := "ipv6.0."

	InterfaceEthernetIpv6AddressListCount := d.Get(prefix9 + "address_list.#").(int)
	obj9.AddressList = make([]InterfaceEthernetIpv6AddressList, 0, InterfaceEthernetIpv6AddressListCount)

	for i := 0; i < InterfaceEthernetIpv6AddressListCount; i++ {
		var obj9_1 InterfaceEthernetIpv6AddressList
		prefix9_1 := prefix9 + fmt.Sprintf("address_list.%d.", i)
		obj9_1.Ipv6Addr = d.Get(prefix9_1 + "ipv6_addr").(string)
		obj9_1.AddressType = d.Get(prefix9_1 + "address_type").(string)
		obj9.AddressList = append(obj9.AddressList, obj9_1)
	}

	obj9.Inside = d.Get(prefix9 + "inside").(int)
	obj9.Outside = d.Get(prefix9 + "outside").(int)
	obj9.Ipv6Enable = d.Get(prefix9 + "ipv6_enable").(int)
	obj9.TtlIgnore = d.Get(prefix9 + "ttl_ignore").(int)

	var obj9_2 InterfaceEthernetIpv6AccessListCfg
	prefix9_2 := prefix9 + "access_list_cfg.0."
	obj9_2.V6AclName = d.Get(prefix9_2 + "v6_acl_name").(string)
	obj9_2.Inbound = d.Get(prefix9_2 + "inbound").(int)

	obj9.AccessListCfg = obj9_2

	var obj9_3 InterfaceEthernetIpv6RouterAdver
	prefix9_3 := prefix9 + "router_adver.0."
	obj9_3.Action = d.Get(prefix9_3 + "action").(string)
	obj9_3.HopLimit = d.Get(prefix9_3 + "hop_limit").(int)
	obj9_3.MaxInterval = d.Get(prefix9_3 + "max_interval").(int)
	obj9_3.MinInterval = d.Get(prefix9_3 + "min_interval").(int)
	obj9_3.DefaultLifetime = d.Get(prefix9_3 + "default_lifetime").(int)
	obj9_3.RateLimit = d.Get(prefix9_3 + "rate_limit").(int)
	obj9_3.ReachableTime = d.Get(prefix9_3 + "reachable_time").(int)
	obj9_3.RetransmitTimer = d.Get(prefix9_3 + "retransmit_timer").(int)
	obj9_3.AdverMtuDisable = d.Get(prefix9_3 + "adver_mtu_disable").(int)
	obj9_3.AdverMtu = d.Get(prefix9_3 + "adver_mtu").(int)

	InterfaceEthernetIpv6RouterAdverPrefixListCount := d.Get(prefix9_3 + "prefix_list.#").(int)
	obj9_3.PrefixList = make([]InterfaceEthernetIpv6RouterAdverPrefixList, 0, InterfaceEthernetIpv6RouterAdverPrefixListCount)
	for i := 0; i < InterfaceEthernetIpv6RouterAdverPrefixListCount; i++ {
		var obj9_3_1 InterfaceEthernetIpv6RouterAdverPrefixList
		prefix9_3_1 := prefix9_3 + fmt.Sprintf("prefix_list.%d.", i)
		obj9_3_1.Prefix = d.Get(prefix9_3_1 + "prefix").(string)
		obj9_3_1.NotAutonomous = d.Get(prefix9_3_1 + "not_autonomous").(int)
		obj9_3_1.NotOnLink = d.Get(prefix9_3_1 + "not_on_link").(int)
		obj9_3_1.PreferredLifetime = d.Get(prefix9_3_1 + "preferred_lifetime").(int)
		obj9_3_1.ValidLifetime = d.Get(prefix9_3_1 + "valid_lifetime").(int)
		obj9_3.PrefixList = append(obj9_3.PrefixList, obj9_3_1)
	}

	obj9_3.ManagedConfigAction = d.Get(prefix9_3 + "managed_config_action").(string)
	obj9_3.OtherConfigAction = d.Get(prefix9_3 + "other_config_action").(string)
	obj9_3.AdverVrid = d.Get(prefix9_3 + "adver_vrid").(int)
	obj9_3.UseFloatingIp = d.Get(prefix9_3 + "use_floating_ip").(int)
	obj9_3.FloatingIp = d.Get(prefix9_3 + "floating_ip").(string)
	obj9_3.AdverVridDefault = d.Get(prefix9_3 + "adver_vrid_default").(int)
	obj9_3.UseFloatingIpDefaultVrid = d.Get(prefix9_3 + "use_floating_ip_default_vrid").(int)
	obj9_3.FloatingIpDefaultVrid = d.Get(prefix9_3 + "floating_ip_default_vrid").(string)

	obj9.RouterAdver = obj9_3

	var obj9_4 InterfaceEthernetIpv6StatefulFirewall
	prefix9_4 := prefix9 + "stateful_firewall.0."
	obj9_4.Inside = d.Get(prefix9_4 + "inside").(int)
	obj9_4.ClassList = d.Get(prefix9_4 + "class_list").(string)
	obj9_4.Outside = d.Get(prefix9_4 + "outside").(int)
	obj9_4.AccessList = d.Get(prefix9_4 + "access_list").(int)
	obj9_4.AclName = d.Get(prefix9_4 + "acl_name").(string)
	obj9.StatefulFirewall = obj9_4

	var obj9_5 InterfaceEthernetIpv6Router
	prefix9_5 := prefix9 + "router.0."

	var obj9_5_1 InterfaceEthernetIpv6RouterRipng
	prefix9_5_1 := prefix9_5 + "ripng.0."
	obj9_5_1.Rip = d.Get(prefix9_5_1 + "rip").(int)
	obj9_5.Ripng = obj9_5_1

	var obj9_5_2 InterfaceEthernetIpv6RouterOspf
	prefix9_5_2 := prefix9_5 + "ospf.0."

	InterfaceEthernetIpv6RouterOspfAreaListCount := d.Get(prefix9_5_2 + "area_list.#").(int)
	obj9_5_2.AreaList = make([]InterfaceEthernetIpv6RouterOspfAreaList, 0, InterfaceEthernetIpv6RouterOspfAreaListCount)
	for i := 0; i < InterfaceEthernetIpv6RouterOspfAreaListCount; i++ {
		var obj9_5_2_1 InterfaceEthernetIpv6RouterOspfAreaList
		prefix9_5_2_1 := prefix9_5_2 + fmt.Sprintf("area_list.%d.", i)
		obj9_5_2_1.AreaIdNum = d.Get(prefix9_5_2_1 + "area_id_num").(int)
		obj9_5_2_1.AreaIdAddr = d.Get(prefix9_5_2_1 + "area_id_addr").(string)
		obj9_5_2_1.Tag = d.Get(prefix9_5_2_1 + "tag").(string)
		obj9_5_2_1.InstanceId = d.Get(prefix9_5_2_1 + "instance_id").(int)
		obj9_5_2.AreaList = append(obj9_5_2.AreaList, obj9_5_2_1)
	}

	obj9_5.Ospf = obj9_5_2

	var obj9_5_3 InterfaceEthernetIpv6RouterIsis
	prefix9_5_3 := prefix9_5 + "isis.0."
	obj9_5_3.Tag = d.Get(prefix9_5_3 + "tag").(string)
	obj9_5.Isis = obj9_5_3

	obj9.Router = obj9_5

	var obj9_6 InterfaceEthernetIpv6Rip
	prefix9_6 := prefix9 + "rip.0."

	var obj9_6_1 InterfaceEthernetIpv6RipSplitHorizonCfg
	prefix9_6_1 := prefix9_6 + "split_horizon_cfg.0."
	obj9_6_1.State = d.Get(prefix9_6_1 + "state").(string)
	obj9_6.SplitHorizonCfg = obj9_6_1

	obj9.Rip = obj9_6

	var obj9_7 InterfaceEthernetIpv6Ospf
	prefix9_7 := prefix9 + "ospf.0."

	InterfaceEthernetIpv6OspfNetworkListCount := d.Get(prefix9_7 + "network_list.#").(int)
	obj9_7.NetworkList = make([]InterfaceEthernetIpv6OspfNetworkList, 0, InterfaceEthernetIpv6OspfNetworkListCount)
	for i := 0; i < InterfaceEthernetIpv6OspfNetworkListCount; i++ {
		var obj9_7_1 InterfaceEthernetIpv6OspfNetworkList
		prefix9_7_1 := prefix9_7 + fmt.Sprintf("network_list.%d.", i)
		obj9_7_1.BroadcastType = d.Get(prefix9_7_1 + "broadcast_type").(string)
		obj9_7_1.P2mpNbma = d.Get(prefix9_7_1 + "p2mp_nbma").(int)
		obj9_7_1.NetworkInstanceId = d.Get(prefix9_7_1 + "network_instance_id").(int)
		obj9_7.NetworkList = append(obj9_7.NetworkList, obj9_7_1)
	}

	obj9_7.Bfd = d.Get(prefix9_7 + "bfd").(int)
	obj9_7.Disable = d.Get(prefix9_7 + "disable").(int)

	InterfaceEthernetIpv6OspfCostCfgCount := d.Get(prefix9_7 + "cost_cfg.#").(int)
	obj9_7.CostCfg = make([]InterfaceEthernetIpv6OspfCostCfg, 0, InterfaceEthernetIpv6OspfCostCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfCostCfgCount; i++ {
		var obj9_7_2 InterfaceEthernetIpv6OspfCostCfg
		prefix9_7_2 := prefix9_7 + fmt.Sprintf("cost_cfg.%d.", i)
		obj9_7_2.Cost = d.Get(prefix9_7_2 + "cost").(int)
		obj9_7_2.InstanceId = d.Get(prefix9_7_2 + "instance_id").(int)
		obj9_7.CostCfg = append(obj9_7.CostCfg, obj9_7_2)
	}

	InterfaceEthernetIpv6OspfDeadIntervalCfgCount := d.Get(prefix9_7 + "dead_interval_cfg.#").(int)
	obj9_7.DeadIntervalCfg = make([]InterfaceEthernetIpv6OspfDeadIntervalCfg, 0, InterfaceEthernetIpv6OspfDeadIntervalCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfDeadIntervalCfgCount; i++ {
		var obj9_7_3 InterfaceEthernetIpv6OspfDeadIntervalCfg
		prefix9_7_3 := prefix9_7 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj9_7_3.DeadInterval = d.Get(prefix9_7_3 + "dead_interval").(int)
		obj9_7_3.InstanceId = d.Get(prefix9_7_3 + "instance_id").(int)
		obj9_7.DeadIntervalCfg = append(obj9_7.DeadIntervalCfg, obj9_7_3)
	}

	InterfaceEthernetIpv6OspfHelloIntervalCfgCount := d.Get(prefix9_7 + "hello_interval_cfg.#").(int)
	obj9_7.HelloIntervalCfg = make([]InterfaceEthernetIpv6OspfHelloIntervalCfg, 0, InterfaceEthernetIpv6OspfHelloIntervalCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfHelloIntervalCfgCount; i++ {
		var obj9_7_4 InterfaceEthernetIpv6OspfHelloIntervalCfg
		prefix9_7_4 := prefix9_7 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj9_7_4.HelloInterval = d.Get(prefix9_7_4 + "hello_interval").(int)
		obj9_7_4.InstanceId = d.Get(prefix9_7_4 + "instance_id").(int)
		obj9_7.HelloIntervalCfg = append(obj9_7.HelloIntervalCfg, obj9_7_4)
	}

	InterfaceEthernetIpv6OspfMtuIgnoreCfgCount := d.Get(prefix9_7 + "mtu_ignore_cfg.#").(int)
	obj9_7.MtuIgnoreCfg = make([]InterfaceEthernetIpv6OspfMtuIgnoreCfg, 0, InterfaceEthernetIpv6OspfMtuIgnoreCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfMtuIgnoreCfgCount; i++ {
		var obj9_7_5 InterfaceEthernetIpv6OspfMtuIgnoreCfg
		prefix9_7_5 := prefix9_7 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj9_7_5.MtuIgnore = d.Get(prefix9_7_5 + "mtu_ignore").(int)
		obj9_7_5.InstanceId = d.Get(prefix9_7_5 + "instance_id").(int)
		obj9_7.MtuIgnoreCfg = append(obj9_7.MtuIgnoreCfg, obj9_7_5)
	}

	InterfaceEthernetIpv6OspfNeighborCfgCount := d.Get(prefix9_7 + "neighbor_cfg.#").(int)
	obj9_7.NeighborCfg = make([]InterfaceEthernetIpv6OspfNeighborCfg, 0, InterfaceEthernetIpv6OspfNeighborCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfNeighborCfgCount; i++ {
		var obj9_7_6 InterfaceEthernetIpv6OspfNeighborCfg
		prefix9_7_6 := prefix9_7 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj9_7_6.Neighbor = d.Get(prefix9_7_6 + "neighbor").(string)
		obj9_7_6.NeigInst = d.Get(prefix9_7_6 + "neig_inst").(int)
		obj9_7_6.NeighborCost = d.Get(prefix9_7_6 + "neighbor_cost").(int)
		obj9_7_6.NeighborPollInterval = d.Get(prefix9_7_6 + "neighbor_poll_interval").(int)
		obj9_7_6.NeighborPriority = d.Get(prefix9_7_6 + "neighbor_priority").(int)
		obj9_7.NeighborCfg = append(obj9_7.NeighborCfg, obj9_7_6)
	}

	InterfaceEthernetIpv6OspfPriorityCfgCount := d.Get(prefix9_7 + "priority_cfg.#").(int)
	obj9_7.PriorityCfg = make([]InterfaceEthernetIpv6OspfPriorityCfg, 0, InterfaceEthernetIpv6OspfPriorityCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfPriorityCfgCount; i++ {
		var obj9_7_7 InterfaceEthernetIpv6OspfPriorityCfg
		prefix9_7_7 := prefix9_7 + fmt.Sprintf("priority_cfg.%d.", i)
		obj9_7_7.Priority = d.Get(prefix9_7_7 + "priority").(int)
		obj9_7_7.InstanceId = d.Get(prefix9_7_7 + "instance_id").(int)
		obj9_7.PriorityCfg = append(obj9_7.PriorityCfg, obj9_7_7)
	}

	InterfaceEthernetIpv6OspfRetransmitIntervalCfgCount := d.Get(prefix9_7 + "retransmit_interval_cfg.#").(int)
	obj9_7.RetransmitIntervalCfg = make([]InterfaceEthernetIpv6OspfRetransmitIntervalCfg, 0, InterfaceEthernetIpv6OspfRetransmitIntervalCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfRetransmitIntervalCfgCount; i++ {
		var obj9_7_8 InterfaceEthernetIpv6OspfRetransmitIntervalCfg
		prefix9_7_8 := prefix9_7 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj9_7_8.RetransmitInterval = d.Get(prefix9_7_8 + "retransmit_interval").(int)
		obj9_7_8.InstanceId = d.Get(prefix9_7_8 + "instance_id").(int)
		obj9_7.RetransmitIntervalCfg = append(obj9_7.RetransmitIntervalCfg, obj9_7_8)
	}

	InterfaceEthernetIpv6OspfTransmitDelayCfgCount := d.Get(prefix9_7 + "transmit_delay_cfg.#").(int)
	obj9_7.TransmitDelayCfg = make([]InterfaceEthernetIpv6OspfTransmitDelayCfg, 0, InterfaceEthernetIpv6OspfTransmitDelayCfgCount)
	for i := 0; i < InterfaceEthernetIpv6OspfTransmitDelayCfgCount; i++ {
		var obj9_7_9 InterfaceEthernetIpv6OspfTransmitDelayCfg
		prefix9_7_9 := prefix9_7 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj9_7_9.TransmitDelay = d.Get(prefix9_7_9 + "transmit_delay").(int)
		obj9_7_9.InstanceId = d.Get(prefix9_7_9 + "instance_id").(int)
		obj9_7.TransmitDelayCfg = append(obj9_7.TransmitDelayCfg, obj9_7_9)
	}

	obj9.Ospf = obj9_7

	c.Inst.Ipv6 = obj9

	var obj10 InterfaceEthernetNptv6
	prefix10 := "nptv6.0."

	InterfaceEthernetNptv6DomainListCount := d.Get(prefix10 + "domain_list.#").(int)
	obj10.DomainList = make([]InterfaceEthernetNptv6DomainList, 0, InterfaceEthernetNptv6DomainListCount)
	for i := 0; i < InterfaceEthernetNptv6DomainListCount; i++ {
		var obj10_1 InterfaceEthernetNptv6DomainList
		prefix10_1 := prefix10 + fmt.Sprintf("domain_list.%d.", i)
		obj10_1.DomainName = d.Get(prefix10_1 + "domain_name").(string)
		obj10_1.BindType = d.Get(prefix10_1 + "bind_type").(string)
		obj10.DomainList = append(obj10.DomainList, obj10_1)
	}

	c.Inst.Nptv6 = obj10

	var obj11 InterfaceEthernetMap
	prefix11 := "map.0."
	obj11.Inside = d.Get(prefix11 + "inside").(int)
	obj11.Outside = d.Get(prefix11 + "outside").(int)
	obj11.MapTInside = d.Get(prefix11 + "map_t_inside").(int)
	obj11.MapTOutside = d.Get(prefix11 + "map_t_outside").(int)

	c.Inst.Map = obj11

	var obj12 InterfaceEthernetLw4o6
	prefix12 := "lw_4o6.0."
	obj12.Outside = d.Get(prefix12 + "outside").(int)
	obj12.Inside = d.Get(prefix12 + "inside").(int)

	c.Inst.Lw4o6 = obj12

	InterfaceEthernetTrunkGroupListCount := d.Get("trunk_group_list.#").(int)
	c.Inst.TrunkGroupList = make([]InterfaceEthernetTrunkGroupList, 0, InterfaceEthernetTrunkGroupListCount)
	for i := 0; i < InterfaceEthernetTrunkGroupListCount; i++ {
		var obj13 InterfaceEthernetTrunkGroupList
		prefix13 := fmt.Sprintf("trunk_group_list.%d.", i)
		obj13.TrunkNumber = d.Get(prefix13 + "trunk_number").(int)
		obj13.Type = d.Get(prefix13 + "type").(string)
		obj13.AdminKey = d.Get(prefix13 + "admin_key").(int)
		obj13.PortPriority = d.Get(prefix13 + "port_priority").(int)

		var obj13_1 InterfaceEthernetTrunkGroupListUdldTimeoutCfg
		prefix13_1 := prefix13 + "udld_timeout_cfg.0."
		obj13_1.Fast = d.Get(prefix13_1 + "fast").(int)
		obj13_1.Slow = d.Get(prefix13_1 + "slow").(int)

		obj13.UdldTimeoutCfg = obj13_1

		obj13.Mode = d.Get(prefix13 + "mode").(string)
		obj13.Timeout = d.Get(prefix13 + "timeout").(string)
		obj13.UserTag = d.Get(prefix13 + "user_tag").(string)
		c.Inst.TrunkGroupList = append(c.Inst.TrunkGroupList, obj13)
	}

	var obj14 InterfaceEthernetBfd
	prefix14 := "bfd.0."

	var obj14_1 InterfaceEthernetBfdAuthentication
	prefix14_1 := prefix14 + "authentication.0."
	obj14_1.KeyId = d.Get(prefix14_1 + "key_id").(int)
	obj14_1.Method = d.Get(prefix14_1 + "method").(string)
	obj14_1.Password = d.Get(prefix14_1 + "password").(string)

	obj14.Authentication = obj14_1

	obj14.Echo = d.Get(prefix14 + "echo").(int)
	obj14.Demand = d.Get(prefix14 + "demand").(int)

	var obj14_2 InterfaceEthernetBfdIntervalCfg
	prefix14_2 := prefix14 + "interval_cfg.0."
	obj14_2.Interval = d.Get(prefix14_2 + "interval").(int)
	obj14_2.MinRx = d.Get(prefix14_2 + "min_rx").(int)
	obj14_2.Multiplier = d.Get(prefix14_2 + "multiplier").(int)

	obj14.IntervalCfg = obj14_2

	c.Inst.Bfd = obj14

	var obj15 InterfaceEthernetIsis
	prefix15 := "isis.0."

	var obj15_1 InterfaceEthernetIsisAuthentication
	prefix15_1 := prefix15 + "authentication.0."

	InterfaceEthernetIsisAuthenticationSendOnlyListCount := d.Get(prefix15_1 + "send_only_list.#").(int)
	obj15_1.SendOnlyList = make([]InterfaceEthernetIsisAuthenticationSendOnlyList, 0, InterfaceEthernetIsisAuthenticationSendOnlyListCount)
	for i := 0; i < InterfaceEthernetIsisAuthenticationSendOnlyListCount; i++ {
		var obj15_1_1 InterfaceEthernetIsisAuthenticationSendOnlyList
		prefix15_1_1 := prefix15_1 + fmt.Sprintf("send_only_list.%d.", i)
		obj15_1_1.SendOnly = d.Get(prefix15_1_1 + "send_only").(int)
		obj15_1_1.Level = d.Get(prefix15_1_1 + "level").(string)
		obj15_1.SendOnlyList = append(obj15_1.SendOnlyList, obj15_1_1)
	}

	InterfaceEthernetIsisAuthenticationModeListCount := d.Get(prefix15_1 + "mode_list.#").(int)
	obj15_1.ModeList = make([]InterfaceEthernetIsisAuthenticationModeList, 0, InterfaceEthernetIsisAuthenticationModeListCount)
	for i := 0; i < InterfaceEthernetIsisAuthenticationModeListCount; i++ {
		var obj15_1_2 InterfaceEthernetIsisAuthenticationModeList
		prefix15_1_2 := prefix15_1 + fmt.Sprintf("mode_list.%d.", i)
		obj15_1_2.Mode = d.Get(prefix15_1_2 + "mode").(string)
		obj15_1_2.Level = d.Get(prefix15_1_2 + "level").(string)
		obj15_1.ModeList = append(obj15_1.ModeList, obj15_1_2)
	}

	InterfaceEthernetIsisAuthenticationKeyChainListCount := d.Get(prefix15_1 + "key_chain_list.#").(int)
	obj15_1.KeyChainList = make([]InterfaceEthernetIsisAuthenticationKeyChainList, 0, InterfaceEthernetIsisAuthenticationKeyChainListCount)
	for i := 0; i < InterfaceEthernetIsisAuthenticationKeyChainListCount; i++ {
		var obj15_1_3 InterfaceEthernetIsisAuthenticationKeyChainList
		prefix15_1_3 := prefix15_1 + fmt.Sprintf("key_chain_list.%d.", i)
		obj15_1_3.KeyChain = d.Get(prefix15_1_3 + "key_chain").(string)
		obj15_1_3.Level = d.Get(prefix15_1_3 + "level").(string)
		obj15_1.KeyChainList = append(obj15_1.KeyChainList, obj15_1_3)
	}

	obj15.Authentication = obj15_1

	var obj15_2 InterfaceEthernetIsisBfdCfg
	prefix15_2 := prefix15 + "bfd_cfg.0."
	obj15_2.Bfd = d.Get(prefix15_2 + "bfd").(int)
	obj15_2.Disable = d.Get(prefix15_2 + "disable").(int)
	obj15.BfdCfg = obj15_2

	obj15.CircuitType = d.Get(prefix15 + "circuit_type").(string)

	InterfaceEthernetIsisCsnpIntervalListCount := d.Get(prefix15 + "csnp_interval_list.#").(int)
	obj15.CsnpIntervalList = make([]InterfaceEthernetIsisCsnpIntervalList, 0, InterfaceEthernetIsisCsnpIntervalListCount)
	for i := 0; i < InterfaceEthernetIsisCsnpIntervalListCount; i++ {
		var obj15_3 InterfaceEthernetIsisCsnpIntervalList
		prefix15_3 := prefix15 + fmt.Sprintf("csnp_interval_list.%d.", i)
		obj15_3.CsnpInterval = d.Get(prefix15_3 + "csnp_interval").(int)
		obj15_3.Level = d.Get(prefix15_3 + "level").(string)
		obj15.CsnpIntervalList = append(obj15.CsnpIntervalList, obj15_3)
	}

	obj15.Padding = d.Get(prefix15 + "padding").(int)

	InterfaceEthernetIsisHelloIntervalListCount := d.Get(prefix15 + "hello_interval_list.#").(int)
	obj15.HelloIntervalList = make([]InterfaceEthernetIsisHelloIntervalList, 0, InterfaceEthernetIsisHelloIntervalListCount)
	for i := 0; i < InterfaceEthernetIsisHelloIntervalListCount; i++ {
		var obj15_4 InterfaceEthernetIsisHelloIntervalList
		prefix15_4 := prefix15 + fmt.Sprintf("hello_interval_list.%d.", i)
		obj15_4.HelloInterval = d.Get(prefix15_4 + "hello_interval").(int)
		obj15_4.Level = d.Get(prefix15_4 + "level").(string)
		obj15.HelloIntervalList = append(obj15.HelloIntervalList, obj15_4)
	}

	InterfaceEthernetIsisHelloIntervalMinimalListCount := d.Get(prefix15 + "hello_interval_minimal_list.#").(int)
	obj15.HelloIntervalMinimalList = make([]InterfaceEthernetIsisHelloIntervalMinimalList, 0, InterfaceEthernetIsisHelloIntervalMinimalListCount)
	for i := 0; i < InterfaceEthernetIsisHelloIntervalMinimalListCount; i++ {
		var obj15_5 InterfaceEthernetIsisHelloIntervalMinimalList
		prefix15_5 := prefix15 + fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj15_5.HelloIntervalMinimal = d.Get(prefix15_5 + "hello_interval_minimal").(int)
		obj15_5.Level = d.Get(prefix15_5 + "level").(string)
		obj15.HelloIntervalMinimalList = append(obj15.HelloIntervalMinimalList, obj15_5)
	}

	InterfaceEthernetIsisHelloMultiplierListCount := d.Get(prefix15 + "hello_multiplier_list.#").(int)
	obj15.HelloMultiplierList = make([]InterfaceEthernetIsisHelloMultiplierList, 0, InterfaceEthernetIsisHelloMultiplierListCount)
	for i := 0; i < InterfaceEthernetIsisHelloMultiplierListCount; i++ {
		var obj15_6 InterfaceEthernetIsisHelloMultiplierList
		prefix15_6 := prefix15 + fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj15_6.HelloMultiplier = d.Get(prefix15_6 + "hello_multiplier").(int)
		obj15_6.Level = d.Get(prefix15_6 + "level").(string)
		obj15.HelloMultiplierList = append(obj15.HelloMultiplierList, obj15_6)
	}

	obj15.LspInterval = d.Get(prefix15 + "lsp_interval").(int)

	var obj15_7 InterfaceEthernetIsisMeshGroup
	prefix15_7 := prefix15 + "mesh_group.0."
	obj15_7.Value = d.Get(prefix15_7 + "value").(int)
	obj15_7.Blocked = d.Get(prefix15_7 + "blocked").(int)

	obj15.MeshGroup = obj15_7

	InterfaceEthernetIsisMetricListCount := d.Get(prefix15 + "metric_list.#").(int)
	obj15.MetricList = make([]InterfaceEthernetIsisMetricList, 0, InterfaceEthernetIsisMetricListCount)
	for i := 0; i < InterfaceEthernetIsisMetricListCount; i++ {
		var obj15_8 InterfaceEthernetIsisMetricList
		prefix15_8 := prefix15 + fmt.Sprintf("metric_list.%d.", i)
		obj15_8.Metric = d.Get(prefix15_8 + "metric").(int)
		obj15_8.Level = d.Get(prefix15_8 + "level").(string)
		obj15.MetricList = append(obj15.MetricList, obj15_8)
	}

	obj15.Network = d.Get(prefix15 + "network").(string)

	InterfaceEthernetIsisPasswordListCount := d.Get(prefix15 + "password_list.#").(int)
	obj15.PasswordList = make([]InterfaceEthernetIsisPasswordList, 0, InterfaceEthernetIsisPasswordListCount)
	for i := 0; i < InterfaceEthernetIsisPasswordListCount; i++ {
		var obj15_9 InterfaceEthernetIsisPasswordList
		prefix15_9 := prefix15 + fmt.Sprintf("password_list.%d.", i)
		obj15_9.Password = d.Get(prefix15_9 + "password").(string)
		obj15_9.Level = d.Get(prefix15_9 + "level").(string)
		obj15.PasswordList = append(obj15.PasswordList, obj15_9)
	}

	InterfaceEthernetIsisPriorityListCount := d.Get(prefix15 + "priority_list.#").(int)
	obj15.PriorityList = make([]InterfaceEthernetIsisPriorityList, 0, InterfaceEthernetIsisPriorityListCount)
	for i := 0; i < InterfaceEthernetIsisPriorityListCount; i++ {
		var obj15_10 InterfaceEthernetIsisPriorityList
		prefix15_10 := prefix15 + fmt.Sprintf("priority_list.%d.", i)
		obj15_10.Priority = d.Get(prefix15_10 + "priority").(int)
		obj15_10.Level = d.Get(prefix15_10 + "level").(string)
		obj15.PriorityList = append(obj15.PriorityList, obj15_10)
	}

	obj15.RetransmitInterval = d.Get(prefix15 + "retransmit_interval").(int)

	InterfaceEthernetIsisWideMetricListCount := d.Get(prefix15 + "wide_metric_list.#").(int)
	obj15.WideMetricList = make([]InterfaceEthernetIsisWideMetricList, 0, InterfaceEthernetIsisWideMetricListCount)
	for i := 0; i < InterfaceEthernetIsisWideMetricListCount; i++ {
		var obj15_11 InterfaceEthernetIsisWideMetricList
		prefix15_11 := prefix15 + fmt.Sprintf("wide_metric_list.%d.", i)
		obj15_11.WideMetric = d.Get(prefix15_11 + "wide_metric").(int)
		obj15_11.Level = d.Get(prefix15_11 + "level").(string)
		obj15.WideMetricList = append(obj15.WideMetricList, obj15_11)
	}

	c.Inst.Isis = obj15

	var obj16 InterfaceEthernetSpanningTree
	prefix16 := "spanning_tree.0."
	obj16.AutoEdge = d.Get(prefix16 + "auto_edge").(int)
	obj16.AdminEdge = d.Get(prefix16 + "admin_edge").(int)

	InterfaceEthernetSpanningTreeInstanceListCount := d.Get(prefix16 + "instance_list.#").(int)
	obj16.InstanceList = make([]InterfaceEthernetSpanningTreeInstanceList, 0, InterfaceEthernetSpanningTreeInstanceListCount)
	for i := 0; i < InterfaceEthernetSpanningTreeInstanceListCount; i++ {
		var obj16_1 InterfaceEthernetSpanningTreeInstanceList
		prefix16_1 := prefix16 + fmt.Sprintf("instance_list.%d.", i)
		obj16_1.InstanceStart = d.Get(prefix16_1 + "instance_start").(int)
		obj16_1.MstpPathCost = d.Get(prefix16_1 + "mstp_path_cost").(int)
		obj16.InstanceList = append(obj16.InstanceList, obj16_1)
	}

	obj16.PathCost = d.Get(prefix16 + "path_cost").(int)

	c.Inst.SpanningTree = obj16
	return c
}
