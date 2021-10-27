package thunder

//Thunder resource InterfaceVe

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"util"
)

func resourceInterfaceVe() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceVeCreate,
		UpdateContext: resourceInterfaceVeUpdate,
		ReadContext:   resourceInterfaceVeRead,
		DeleteContext: resourceInterfaceVeDelete,
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
						"ipv6_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
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
									"default_lifetime": {
										Type:        schema.TypeInt,
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
		},
	}
}

func resourceInterfaceVeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceVe (Inside resourceInterfaceVeCreate) ")
		name1 := d.Get("ifnum").(int)
		data := dataToInterfaceVe(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceVe --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostInterfaceVe(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceVeRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceVeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceVe (Inside resourceInterfaceVeRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceVe(client.Token, name1, client.Host)
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

func resourceInterfaceVeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying InterfaceVe   (Inside resourceInterfaceVeUpdate) ")
		data := dataToInterfaceVe(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceVe ")
		err := go_thunder.PutInterfaceVe(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceVeRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceVeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceVeDelete) " + name1)
		err := go_thunder.DeleteInterfaceVe(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToInterfaceVe(d *schema.ResourceData) go_thunder.InterfaceVe {
	var vc go_thunder.InterfaceVe
	var c go_thunder.InterfaceVeInstance
	c.InterfaceVeInstanceIfnum = d.Get("ifnum").(int)
	c.InterfaceVeInstanceName = d.Get("name").(string)
	c.InterfaceVeInstancePortScanDetection = d.Get("port_scan_detection").(string)
	c.InterfaceVeInstancePingSweepDetection = d.Get("ping_sweep_detection").(string)
	c.InterfaceVeInstanceL3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	c.InterfaceVeInstanceMtu = d.Get("mtu").(int)
	c.InterfaceVeInstanceTrapSource = d.Get("trap_source").(int)
	c.InterfaceVeInstanceAction = d.Get("action").(string)

	var obj1 go_thunder.InterfaceVeInstanceIcmpRateLimit
	prefix1 := "icmp_rate_limit.0."
	obj1.InterfaceVeInstanceIcmpRateLimitNormal = d.Get(prefix1 + "normal").(int)
	obj1.InterfaceVeInstanceIcmpRateLimitLockup = d.Get(prefix1 + "lockup").(int)
	obj1.InterfaceVeInstanceIcmpRateLimitLockupPeriod = d.Get(prefix1 + "lockup_period").(int)

	c.InterfaceVeInstanceIcmpRateLimitNormal = obj1

	var obj2 go_thunder.InterfaceVeInstanceIcmpv6RateLimit
	prefix2 := "icmpv6_rate_limit.0."
	obj2.InterfaceVeInstanceIcmpv6RateLimitNormalV6 = d.Get(prefix2 + "normal_v6").(int)
	obj2.InterfaceVeInstanceIcmpv6RateLimitLockupV6 = d.Get(prefix2 + "lockup_v6").(int)
	obj2.InterfaceVeInstanceIcmpv6RateLimitLockupPeriodV6 = d.Get(prefix2 + "lockup_period_v6").(int)

	c.InterfaceVeInstanceIcmpv6RateLimitNormalV6 = obj2

	var obj3 go_thunder.InterfaceVeInstanceAccessList
	prefix3 := "access_list.0."
	obj3.InterfaceVeInstanceAccessListAclID = d.Get(prefix3 + "acl_id").(int)
	obj3.InterfaceVeInstanceAccessListAclName = d.Get(prefix3 + "acl_name").(string)

	c.InterfaceVeInstanceAccessListAclID = obj3

	c.InterfaceVeInstanceUserTag = d.Get("user_tag").(string)

	InterfaceVeInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.InterfaceVeInstanceSamplingEnableCounters1 = make([]go_thunder.InterfaceVeInstanceSamplingEnable, 0, InterfaceVeInstanceSamplingEnableCount)

	for i := 0; i < InterfaceVeInstanceSamplingEnableCount; i++ {
		var obj4 go_thunder.InterfaceVeInstanceSamplingEnable
		prefix4 := fmt.Sprintf("sampling_enable.%d.", i)
		obj4.InterfaceVeInstanceSamplingEnableCounters1 = d.Get(prefix4 + "counters1").(string)
		c.InterfaceVeInstanceSamplingEnableCounters1 = append(c.InterfaceVeInstanceSamplingEnableCounters1, obj4)
	}

	var obj5 go_thunder.InterfaceVeInstanceIP
	prefix5 := "ip.0."
	obj5.InterfaceVeInstanceIPDhcp = d.Get(prefix5 + "dhcp").(int)

	InterfaceVeInstanceIPAddressListCount := d.Get(prefix5 + "address_list.#").(int)
	obj5.InterfaceVeInstanceIPAddressListIpv6Addr = make([]go_thunder.InterfaceVeInstanceIPAddressList, 0, InterfaceVeInstanceIPAddressListCount)

	for i := 0; i < InterfaceVeInstanceIPAddressListCount; i++ {
		var obj5_1 go_thunder.InterfaceVeInstanceIPAddressList
		prefix5_1 := prefix5 + fmt.Sprintf("address_list.%d.", i)
		obj5_1.InterfaceVeInstanceIPAddressListIpv6Addr = d.Get(prefix5_1 + "ipv6_addr").(string)
		obj5_1.InterfaceVeInstanceIPAddressListAddressType = d.Get(prefix5_1 + "address_type").(string)
		obj5.InterfaceVeInstanceIPAddressListIpv6Addr = append(obj5.InterfaceVeInstanceIPAddressListIpv6Addr, obj5_1)
	}

	obj5.InterfaceVeInstanceIPAllowPromiscuousVip = d.Get(prefix5 + "allow_promiscuous_vip").(int)
	obj5.InterfaceVeInstanceIPClient = d.Get(prefix5 + "client").(int)
	obj5.InterfaceVeInstanceIPServer = d.Get(prefix5 + "server").(int)

	InterfaceVeInstanceIPHelperAddressListCount := d.Get(prefix5 + "helper_address_list.#").(int)
	obj5.InterfaceVeInstanceIPHelperAddressListHelperAddress = make([]go_thunder.InterfaceVeInstanceIPHelperAddressList, 0, InterfaceVeInstanceIPHelperAddressListCount)

	for i := 0; i < InterfaceVeInstanceIPHelperAddressListCount; i++ {
		var obj5_2 go_thunder.InterfaceVeInstanceIPHelperAddressList
		prefix5_2 := prefix5 + fmt.Sprintf("helper_address_list.%d.", i)
		obj5_2.InterfaceVeInstanceIPHelperAddressListHelperAddress = d.Get(prefix5_2 + "helper_address").(string)
		obj5.InterfaceVeInstanceIPHelperAddressListHelperAddress = append(obj5.InterfaceVeInstanceIPHelperAddressListHelperAddress, obj5_2)
	}

	obj5.InterfaceVeInstanceIPInside = d.Get(prefix5 + "inside").(int)
	obj5.InterfaceVeInstanceIPOutside = d.Get(prefix5 + "outside").(int)
	obj5.InterfaceVeInstanceIPTTLIgnore = d.Get(prefix5 + "ttl_ignore").(int)
	obj5.InterfaceVeInstanceIPSynCookie = d.Get(prefix5 + "syn_cookie").(int)
	obj5.InterfaceVeInstanceIPSlbPartitionRedirect = d.Get(prefix5 + "slb_partition_redirect").(int)
	obj5.InterfaceVeInstanceIPGenerateMembershipQuery = d.Get(prefix5 + "generate_membership_query").(int)
	obj5.InterfaceVeInstanceIPQueryInterval = d.Get(prefix5 + "query_interval").(int)
	obj5.InterfaceVeInstanceIPMaxRespTime = d.Get(prefix5 + "max_resp_time").(int)

	var obj5_3 go_thunder.InterfaceVeInstanceIPStatefulFirewall
	prefix5_3 := prefix5 + "stateful_firewall.0."
	obj5_3.InterfaceVeInstanceIPStatefulFirewallInside = d.Get(prefix5_3 + "inside").(int)
	obj5_3.InterfaceVeInstanceIPStatefulFirewallClassList = d.Get(prefix5_3 + "class_list").(string)
	obj5_3.InterfaceVeInstanceIPStatefulFirewallOutside = d.Get(prefix5_3 + "outside").(int)
	obj5_3.InterfaceVeInstanceIPStatefulFirewallAccessList = d.Get(prefix5_3 + "access_list").(int)
	obj5_3.InterfaceVeInstanceIPStatefulFirewallAclID = d.Get(prefix5_3 + "acl_id").(int)

	obj5.InterfaceVeInstanceIPStatefulFirewallInside = obj5_3

	var obj5_4 go_thunder.InterfaceVeInstanceIPRouter
	prefix5_4 := prefix5 + "router.0."

	var obj5_4_1 go_thunder.InterfaceVeInstanceIPRouterIsis
	prefix5_4_1 := prefix5_4 + "isis.0."
	obj5_4_1.InterfaceVeInstanceIPRouterIsisTag = d.Get(prefix5_4_1 + "tag").(string)

	obj5_4.InterfaceVeInstanceIPRouterIsisTag = obj5_4_1

	obj5.InterfaceVeInstanceIPRouterIsis = obj5_4

	var obj5_5 go_thunder.InterfaceVeInstanceIPRip
	prefix5_5 := prefix5 + "rip.0."

	var obj5_5_1 go_thunder.InterfaceVeInstanceIPRipAuthentication
	prefix5_5_1 := prefix5_5 + "authentication.0."

	var obj5_5_1_1 go_thunder.InterfaceVeInstanceIPRipAuthenticationStr
	prefix5_5_1_1 := prefix5_5_1 + "str.0."
	obj5_5_1_1.InterfaceVeInstanceIPRipAuthenticationStrString = d.Get(prefix5_5_1_1 + "string").(string)

	obj5_5_1.InterfaceVeInstanceIPRipAuthenticationStrString = obj5_5_1_1

	var obj5_5_1_2 go_thunder.InterfaceVeInstanceIPRipAuthenticationMode
	prefix5_5_1_2 := prefix5_5_1 + "mode.0."
	obj5_5_1_2.InterfaceVeInstanceIPRipAuthenticationModeMode = d.Get(prefix5_5_1_2 + "mode").(string)

	obj5_5_1.InterfaceVeInstanceIPRipAuthenticationModeMode = obj5_5_1_2

	var obj5_5_1_3 go_thunder.InterfaceVeInstanceIPRipAuthenticationKeyChain
	prefix5_5_1_3 := prefix5_5_1 + "key_chain.0."
	obj5_5_1_3.InterfaceVeInstanceIPRipAuthenticationKeyChainKeyChain = d.Get(prefix5_5_1_3 + "key_chain").(string)

	obj5_5_1.InterfaceVeInstanceIPRipAuthenticationKeyChainKeyChain = obj5_5_1_3

	obj5_5.InterfaceVeInstanceIPRipAuthenticationStr = obj5_5_1

	obj5_5.InterfaceVeInstanceIPRipSendPacket = d.Get(prefix5_5 + "send_packet").(int)
	obj5_5.InterfaceVeInstanceIPRipReceivePacket = d.Get(prefix5_5 + "receive_packet").(int)

	var obj5_5_2 go_thunder.InterfaceVeInstanceIPRipSendCfg
	prefix5_5_2 := prefix5_5 + "send_cfg.0."
	obj5_5_2.InterfaceVeInstanceIPRipSendCfgSend = d.Get(prefix5_5_2 + "send").(int)
	obj5_5_2.InterfaceVeInstanceIPRipSendCfgVersion = d.Get(prefix5_5_2 + "version").(string)

	obj5_5.InterfaceVeInstanceIPRipSendCfgSend = obj5_5_2

	var obj5_5_3 go_thunder.InterfaceVeInstanceIPRipReceiveCfg
	prefix5_5_3 := prefix5_5 + "receive_cfg.0."
	obj5_5_3.InterfaceVeInstanceIPRipReceiveCfgReceive = d.Get(prefix5_5_3 + "receive").(int)
	obj5_5_3.InterfaceVeInstanceIPRipReceiveCfgVersion = d.Get(prefix5_5_3 + "version").(string)

	obj5_5.InterfaceVeInstanceIPRipReceiveCfgReceive = obj5_5_3

	var obj5_5_4 go_thunder.InterfaceVeInstanceIPRipSplitHorizonCfg
	prefix5_5_4 := prefix5_5 + "split_horizon_cfg.0."
	obj5_5_4.InterfaceVeInstanceIPRipSplitHorizonCfgState = d.Get(prefix5_5_4 + "state").(string)

	obj5_5.InterfaceVeInstanceIPRipSplitHorizonCfgState = obj5_5_4

	obj5.InterfaceVeInstanceIPRipAuthentication = obj5_5

	var obj5_6 go_thunder.InterfaceVeInstanceIPOspf
	prefix5_6 := prefix5 + "ospf.0."

	var obj5_6_1 go_thunder.InterfaceVeInstanceIPOspfOspfGlobal
	prefix5_6_1 := prefix5_6 + "ospf_global.0."

	var obj5_6_1_1 go_thunder.InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfg
	prefix5_6_1_1 := prefix5_6_1 + "authentication_cfg.0."
	obj5_6_1_1.InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication = d.Get(prefix5_6_1_1 + "authentication").(int)
	obj5_6_1_1.InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfgValue = d.Get(prefix5_6_1_1 + "value").(string)

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication = obj5_6_1_1

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalAuthenticationKey = d.Get(prefix5_6_1 + "authentication_key").(string)

	var obj5_6_1_2 go_thunder.InterfaceVeInstanceIPOspfOspfGlobalBfdCfg
	prefix5_6_1_2 := prefix5_6_1 + "bfd_cfg.0."
	obj5_6_1_2.InterfaceVeInstanceIPOspfOspfGlobalBfdCfgBfd = d.Get(prefix5_6_1_2 + "bfd").(int)
	obj5_6_1_2.InterfaceVeInstanceIPOspfOspfGlobalBfdCfgDisable = d.Get(prefix5_6_1_2 + "disable").(int)

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalBfdCfgBfd = obj5_6_1_2

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalCost = d.Get(prefix5_6_1 + "cost").(int)

	var obj5_6_1_3 go_thunder.InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfg
	prefix5_6_1_3 := prefix5_6_1 + "database_filter_cfg.0."
	obj5_6_1_3.InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = d.Get(prefix5_6_1_3 + "database_filter").(string)
	obj5_6_1_3.InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfgOut = d.Get(prefix5_6_1_3 + "out").(int)

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = obj5_6_1_3

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalDeadInterval = d.Get(prefix5_6_1 + "dead_interval").(int)
	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalDisable = d.Get(prefix5_6_1 + "disable").(string)
	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalHelloInterval = d.Get(prefix5_6_1 + "hello_interval").(int)

	InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgCount := d.Get(prefix5_6_1 + "message_digest_cfg.#").(int)
	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfg, 0, InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgCount)

	for i := 0; i < InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgCount; i++ {
		var obj5_6_1_4 go_thunder.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfg
		prefix5_6_1_4 := prefix5_6_1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj5_6_1_4.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = d.Get(prefix5_6_1_4 + "message_digest_key").(int)
		obj5_6_1_4.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMd5Value = d.Get(prefix5_6_1_4 + "md5_value").(string)
		obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = append(obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey, obj5_6_1_4)
	}

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalMtu = d.Get(prefix5_6_1 + "mtu").(int)
	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalMtuIgnore = d.Get(prefix5_6_1 + "mtu_ignore").(int)

	var obj5_6_1_5 go_thunder.InterfaceVeInstanceIPOspfOspfGlobalNetwork
	prefix5_6_1_5 := prefix5_6_1 + "network.0."
	obj5_6_1_5.InterfaceVeInstanceIPOspfOspfGlobalNetworkBroadcast = d.Get(prefix5_6_1_5 + "broadcast").(int)
	obj5_6_1_5.InterfaceVeInstanceIPOspfOspfGlobalNetworkNonBroadcast = d.Get(prefix5_6_1_5 + "non_broadcast").(int)
	obj5_6_1_5.InterfaceVeInstanceIPOspfOspfGlobalNetworkPointToPoint = d.Get(prefix5_6_1_5 + "point_to_point").(int)
	obj5_6_1_5.InterfaceVeInstanceIPOspfOspfGlobalNetworkPointToMultipoint = d.Get(prefix5_6_1_5 + "point_to_multipoint").(int)
	obj5_6_1_5.InterfaceVeInstanceIPOspfOspfGlobalNetworkP2MpNbma = d.Get(prefix5_6_1_5 + "p2mp_nbma").(int)

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalNetworkBroadcast = obj5_6_1_5

	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalPriority = d.Get(prefix5_6_1 + "priority").(int)
	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalRetransmitInterval = d.Get(prefix5_6_1 + "retransmit_interval").(int)
	obj5_6_1.InterfaceVeInstanceIPOspfOspfGlobalTransmitDelay = d.Get(prefix5_6_1 + "transmit_delay").(int)

	obj5_6.InterfaceVeInstanceIPOspfOspfGlobalAuthenticationCfg = obj5_6_1

	InterfaceVeInstanceIPOspfOspfIPListCount := d.Get(prefix5_6 + "ospf_ip_list.#").(int)
	obj5_6.InterfaceVeInstanceIPOspfOspfIPListIPAddr = make([]go_thunder.InterfaceVeInstanceIPOspfOspfIPList, 0, InterfaceVeInstanceIPOspfOspfIPListCount)

	for i := 0; i < InterfaceVeInstanceIPOspfOspfIPListCount; i++ {
		var obj5_6_2 go_thunder.InterfaceVeInstanceIPOspfOspfIPList
		prefix5_6_2 := prefix5_6 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListIPAddr = d.Get(prefix5_6_2 + "ip_addr").(string)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListAuthentication = d.Get(prefix5_6_2 + "authentication").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListValue = d.Get(prefix5_6_2 + "value").(string)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListAuthenticationKey = d.Get(prefix5_6_2 + "authentication_key").(string)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListCost = d.Get(prefix5_6_2 + "cost").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListDatabaseFilter = d.Get(prefix5_6_2 + "database_filter").(string)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListOut = d.Get(prefix5_6_2 + "out").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListDeadInterval = d.Get(prefix5_6_2 + "dead_interval").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListHelloInterval = d.Get(prefix5_6_2 + "hello_interval").(int)

		InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgCount := d.Get(prefix5_6_2 + "message_digest_cfg.#").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfg, 0, InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgCount)

		for i := 0; i < InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgCount; i++ {
			var obj5_6_2_1 go_thunder.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfg
			prefix5_6_2_1 := prefix5_6_2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj5_6_2_1.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = d.Get(prefix5_6_2_1 + "message_digest_key").(int)
			obj5_6_2_1.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMd5Value = d.Get(prefix5_6_2_1 + "md5_value").(string)
			obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = append(obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey, obj5_6_2_1)
		}

		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListMtuIgnore = d.Get(prefix5_6_2 + "mtu_ignore").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListPriority = d.Get(prefix5_6_2 + "priority").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListRetransmitInterval = d.Get(prefix5_6_2 + "retransmit_interval").(int)
		obj5_6_2.InterfaceVeInstanceIPOspfOspfIPListTransmitDelay = d.Get(prefix5_6_2 + "transmit_delay").(int)
		obj5_6.InterfaceVeInstanceIPOspfOspfIPListIPAddr = append(obj5_6.InterfaceVeInstanceIPOspfOspfIPListIPAddr, obj5_6_2)
	}

	obj5.InterfaceVeInstanceIPOspfOspfGlobal = obj5_6

	c.InterfaceVeInstanceIPDhcp = obj5

	var obj6 go_thunder.InterfaceVeInstanceIpv6
	prefix6 := "ipv6.0."

	InterfaceVeInstanceIpv6AddressListCount := d.Get(prefix6 + "address_list.#").(int)
	obj6.InterfaceVeInstanceIpv6AddressListIpv6Addr = make([]go_thunder.InterfaceVeInstanceIpv6AddressList, 0, InterfaceVeInstanceIpv6AddressListCount)

	for i := 0; i < InterfaceVeInstanceIpv6AddressListCount; i++ {
		var obj6_1 go_thunder.InterfaceVeInstanceIpv6AddressList
		prefix6_1 := prefix6 + fmt.Sprintf("address_list.%d.", i)
		obj6_1.InterfaceVeInstanceIpv6AddressListIpv6Addr = d.Get(prefix6_1 + "ipv6_addr").(string)
		obj6_1.InterfaceVeInstanceIpv6AddressListAddressType = d.Get(prefix6_1 + "address_type").(string)
		obj6.InterfaceVeInstanceIpv6AddressListIpv6Addr = append(obj6.InterfaceVeInstanceIpv6AddressListIpv6Addr, obj6_1)
	}

	obj6.InterfaceVeInstanceIpv6Ipv6Enable = d.Get(prefix6 + "ipv6_enable").(int)
	obj6.InterfaceVeInstanceIpv6V6AclName = d.Get(prefix6 + "v6_acl_name").(string)
	obj6.InterfaceVeInstanceIpv6Inbound = d.Get(prefix6 + "inbound").(int)
	obj6.InterfaceVeInstanceIpv6Inside = d.Get(prefix6 + "inside").(int)
	obj6.InterfaceVeInstanceIpv6Outside = d.Get(prefix6 + "outside").(int)
	obj6.InterfaceVeInstanceIpv6TTLIgnore = d.Get(prefix6 + "ttl_ignore").(int)

	var obj6_2 go_thunder.InterfaceVeInstanceIpv6RouterAdver
	prefix6_2 := prefix6 + "router_adver.0."
	obj6_2.InterfaceVeInstanceIpv6RouterAdverAction = d.Get(prefix6_2 + "action").(string)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverDefaultLifetime = d.Get(prefix6_2 + "default_lifetime").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverHopLimit = d.Get(prefix6_2 + "hop_limit").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverMaxInterval = d.Get(prefix6_2 + "max_interval").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverMinInterval = d.Get(prefix6_2 + "min_interval").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverRateLimit = d.Get(prefix6_2 + "rate_limit").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverReachableTime = d.Get(prefix6_2 + "reachable_time").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverRetransmitTimer = d.Get(prefix6_2 + "retransmit_timer").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverAdverMtuDisable = d.Get(prefix6_2 + "adver_mtu_disable").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverAdverMtu = d.Get(prefix6_2 + "adver_mtu").(int)

	InterfaceVeInstanceIpv6RouterAdverPrefixListCount := d.Get(prefix6_2 + "prefix_list.#").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverPrefixListPrefix = make([]go_thunder.InterfaceVeInstanceIpv6RouterAdverPrefixList, 0, InterfaceVeInstanceIpv6RouterAdverPrefixListCount)

	for i := 0; i < InterfaceVeInstanceIpv6RouterAdverPrefixListCount; i++ {
		var obj6_2_1 go_thunder.InterfaceVeInstanceIpv6RouterAdverPrefixList
		prefix6_2_1 := prefix6_2 + fmt.Sprintf("prefix_list.%d.", i)
		obj6_2_1.InterfaceVeInstanceIpv6RouterAdverPrefixListPrefix = d.Get(prefix6_2_1 + "prefix").(string)
		obj6_2_1.InterfaceVeInstanceIpv6RouterAdverPrefixListNotAutonomous = d.Get(prefix6_2_1 + "not_autonomous").(int)
		obj6_2_1.InterfaceVeInstanceIpv6RouterAdverPrefixListNotOnLink = d.Get(prefix6_2_1 + "not_on_link").(int)
		obj6_2_1.InterfaceVeInstanceIpv6RouterAdverPrefixListPreferredLifetime = d.Get(prefix6_2_1 + "preferred_lifetime").(int)
		obj6_2_1.InterfaceVeInstanceIpv6RouterAdverPrefixListValidLifetime = d.Get(prefix6_2_1 + "valid_lifetime").(int)
		obj6_2.InterfaceVeInstanceIpv6RouterAdverPrefixListPrefix = append(obj6_2.InterfaceVeInstanceIpv6RouterAdverPrefixListPrefix, obj6_2_1)
	}

	obj6_2.InterfaceVeInstanceIpv6RouterAdverManagedConfigAction = d.Get(prefix6_2 + "managed_config_action").(string)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverOtherConfigAction = d.Get(prefix6_2 + "other_config_action").(string)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverAdverVrid = d.Get(prefix6_2 + "adver_vrid").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverUseFloatingIP = d.Get(prefix6_2 + "use_floating_ip").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverFloatingIP = d.Get(prefix6_2 + "floating_ip").(string)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverAdverVridDefault = d.Get(prefix6_2 + "adver_vrid_default").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverUseFloatingIPDefaultVrid = d.Get(prefix6_2 + "use_floating_ip_default_vrid").(int)
	obj6_2.InterfaceVeInstanceIpv6RouterAdverFloatingIPDefaultVrid = d.Get(prefix6_2 + "floating_ip_default_vrid").(string)

	obj6.InterfaceVeInstanceIpv6RouterAdverAction = obj6_2

	var obj6_3 go_thunder.InterfaceVeInstanceIpv6StatefulFirewall
	prefix6_3 := prefix6 + "stateful_firewall.0."
	obj6_3.InterfaceVeInstanceIpv6StatefulFirewallInside = d.Get(prefix6_3 + "inside").(int)
	obj6_3.InterfaceVeInstanceIpv6StatefulFirewallClassList = d.Get(prefix6_3 + "class_list").(string)
	obj6_3.InterfaceVeInstanceIpv6StatefulFirewallOutside = d.Get(prefix6_3 + "outside").(int)
	obj6_3.InterfaceVeInstanceIpv6StatefulFirewallAccessList = d.Get(prefix6_3 + "access_list").(int)
	obj6_3.InterfaceVeInstanceIpv6StatefulFirewallAclName = d.Get(prefix6_3 + "acl_name").(string)

	obj6.InterfaceVeInstanceIpv6StatefulFirewallInside = obj6_3

	var obj6_4 go_thunder.InterfaceVeInstanceIpv6Router
	prefix6_4 := prefix6 + "router.0."

	var obj6_4_1 go_thunder.InterfaceVeInstanceIpv6RouterRipng
	prefix6_4_1 := prefix6_4 + "ripng.0."
	obj6_4_1.InterfaceVeInstanceIpv6RouterRipngRip = d.Get(prefix6_4_1 + "rip").(int)

	obj6_4.InterfaceVeInstanceIpv6RouterRipngRip = obj6_4_1

	var obj6_4_2 go_thunder.InterfaceVeInstanceIpv6RouterOspf
	prefix6_4_2 := prefix6_4 + "ospf.0."

	InterfaceVeInstanceIpv6RouterOspfAreaListCount := d.Get(prefix6_4_2 + "area_list.#").(int)
	obj6_4_2.InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDNum = make([]go_thunder.InterfaceVeInstanceIpv6RouterOspfAreaList, 0, InterfaceVeInstanceIpv6RouterOspfAreaListCount)

	for i := 0; i < InterfaceVeInstanceIpv6RouterOspfAreaListCount; i++ {
		var obj6_4_2_1 go_thunder.InterfaceVeInstanceIpv6RouterOspfAreaList
		prefix6_4_2_1 := prefix6_4_2 + fmt.Sprintf("area_list.%d.", i)
		obj6_4_2_1.InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDNum = d.Get(prefix6_4_2_1 + "area_id_num").(int)
		obj6_4_2_1.InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDAddr = d.Get(prefix6_4_2_1 + "area_id_addr").(string)
		obj6_4_2_1.InterfaceVeInstanceIpv6RouterOspfAreaListTag = d.Get(prefix6_4_2_1 + "tag").(string)
		obj6_4_2_1.InterfaceVeInstanceIpv6RouterOspfAreaListInstanceID = d.Get(prefix6_4_2_1 + "instance_id").(int)
		obj6_4_2.InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDNum = append(obj6_4_2.InterfaceVeInstanceIpv6RouterOspfAreaListAreaIDNum, obj6_4_2_1)
	}

	obj6_4.InterfaceVeInstanceIpv6RouterOspfAreaList = obj6_4_2

	var obj6_4_3 go_thunder.InterfaceVeInstanceIpv6RouterIsis
	prefix6_4_3 := prefix6_4 + "isis.0."
	obj6_4_3.InterfaceVeInstanceIpv6RouterIsisTag = d.Get(prefix6_4_3 + "tag").(string)

	obj6_4.InterfaceVeInstanceIpv6RouterIsisTag = obj6_4_3

	obj6.InterfaceVeInstanceIpv6RouterRipng = obj6_4

	var obj6_5 go_thunder.InterfaceVeInstanceIpv6Rip
	prefix6_5 := prefix6 + "rip.0."

	var obj6_5_1 go_thunder.InterfaceVeInstanceIpv6RipSplitHorizonCfg
	prefix6_5_1 := prefix6_5 + "split_horizon_cfg.0."
	obj6_5_1.InterfaceVeInstanceIpv6RipSplitHorizonCfgState = d.Get(prefix6_5_1 + "state").(string)

	obj6_5.InterfaceVeInstanceIpv6RipSplitHorizonCfgState = obj6_5_1

	obj6.InterfaceVeInstanceIpv6RipSplitHorizonCfg = obj6_5

	var obj6_6 go_thunder.InterfaceVeInstanceIpv6Ospf
	prefix6_6 := prefix6 + "ospf.0."

	InterfaceVeInstanceIpv6OspfNetworkListCount := d.Get(prefix6_6 + "network_list.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfNetworkListBroadcastType = make([]go_thunder.InterfaceVeInstanceIpv6OspfNetworkList, 0, InterfaceVeInstanceIpv6OspfNetworkListCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfNetworkListCount; i++ {
		var obj6_6_1 go_thunder.InterfaceVeInstanceIpv6OspfNetworkList
		prefix6_6_1 := prefix6_6 + fmt.Sprintf("network_list.%d.", i)
		obj6_6_1.InterfaceVeInstanceIpv6OspfNetworkListBroadcastType = d.Get(prefix6_6_1 + "broadcast_type").(string)
		obj6_6_1.InterfaceVeInstanceIpv6OspfNetworkListP2MpNbma = d.Get(prefix6_6_1 + "p2mp_nbma").(int)
		obj6_6_1.InterfaceVeInstanceIpv6OspfNetworkListNetworkInstanceID = d.Get(prefix6_6_1 + "network_instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfNetworkListBroadcastType = append(obj6_6.InterfaceVeInstanceIpv6OspfNetworkListBroadcastType, obj6_6_1)
	}

	obj6_6.InterfaceVeInstanceIpv6OspfBfd = d.Get(prefix6_6 + "bfd").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfDisable = d.Get(prefix6_6 + "disable").(int)

	InterfaceVeInstanceIpv6OspfCostCfgCount := d.Get(prefix6_6 + "cost_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfCostCfgCost = make([]go_thunder.InterfaceVeInstanceIpv6OspfCostCfg, 0, InterfaceVeInstanceIpv6OspfCostCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfCostCfgCount; i++ {
		var obj6_6_2 go_thunder.InterfaceVeInstanceIpv6OspfCostCfg
		prefix6_6_2 := prefix6_6 + fmt.Sprintf("cost_cfg.%d.", i)
		obj6_6_2.InterfaceVeInstanceIpv6OspfCostCfgCost = d.Get(prefix6_6_2 + "cost").(int)
		obj6_6_2.InterfaceVeInstanceIpv6OspfCostCfgInstanceID = d.Get(prefix6_6_2 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfCostCfgCost = append(obj6_6.InterfaceVeInstanceIpv6OspfCostCfgCost, obj6_6_2)
	}

	InterfaceVeInstanceIpv6OspfDeadIntervalCfgCount := d.Get(prefix6_6 + "dead_interval_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfDeadIntervalCfgDeadInterval = make([]go_thunder.InterfaceVeInstanceIpv6OspfDeadIntervalCfg, 0, InterfaceVeInstanceIpv6OspfDeadIntervalCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfDeadIntervalCfgCount; i++ {
		var obj6_6_3 go_thunder.InterfaceVeInstanceIpv6OspfDeadIntervalCfg
		prefix6_6_3 := prefix6_6 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj6_6_3.InterfaceVeInstanceIpv6OspfDeadIntervalCfgDeadInterval = d.Get(prefix6_6_3 + "dead_interval").(int)
		obj6_6_3.InterfaceVeInstanceIpv6OspfDeadIntervalCfgInstanceID = d.Get(prefix6_6_3 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfDeadIntervalCfgDeadInterval = append(obj6_6.InterfaceVeInstanceIpv6OspfDeadIntervalCfgDeadInterval, obj6_6_3)
	}

	InterfaceVeInstanceIpv6OspfHelloIntervalCfgCount := d.Get(prefix6_6 + "hello_interval_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfHelloIntervalCfgHelloInterval = make([]go_thunder.InterfaceVeInstanceIpv6OspfHelloIntervalCfg, 0, InterfaceVeInstanceIpv6OspfHelloIntervalCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfHelloIntervalCfgCount; i++ {
		var obj6_6_4 go_thunder.InterfaceVeInstanceIpv6OspfHelloIntervalCfg
		prefix6_6_4 := prefix6_6 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj6_6_4.InterfaceVeInstanceIpv6OspfHelloIntervalCfgHelloInterval = d.Get(prefix6_6_4 + "hello_interval").(int)
		obj6_6_4.InterfaceVeInstanceIpv6OspfHelloIntervalCfgInstanceID = d.Get(prefix6_6_4 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfHelloIntervalCfgHelloInterval = append(obj6_6.InterfaceVeInstanceIpv6OspfHelloIntervalCfgHelloInterval, obj6_6_4)
	}

	InterfaceVeInstanceIpv6OspfMtuIgnoreCfgCount := d.Get(prefix6_6 + "mtu_ignore_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = make([]go_thunder.InterfaceVeInstanceIpv6OspfMtuIgnoreCfg, 0, InterfaceVeInstanceIpv6OspfMtuIgnoreCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfMtuIgnoreCfgCount; i++ {
		var obj6_6_5 go_thunder.InterfaceVeInstanceIpv6OspfMtuIgnoreCfg
		prefix6_6_5 := prefix6_6 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj6_6_5.InterfaceVeInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = d.Get(prefix6_6_5 + "mtu_ignore").(int)
		obj6_6_5.InterfaceVeInstanceIpv6OspfMtuIgnoreCfgInstanceID = d.Get(prefix6_6_5 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = append(obj6_6.InterfaceVeInstanceIpv6OspfMtuIgnoreCfgMtuIgnore, obj6_6_5)
	}

	InterfaceVeInstanceIpv6OspfNeighborCfgCount := d.Get(prefix6_6 + "neighbor_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighbor = make([]go_thunder.InterfaceVeInstanceIpv6OspfNeighborCfg, 0, InterfaceVeInstanceIpv6OspfNeighborCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfNeighborCfgCount; i++ {
		var obj6_6_6 go_thunder.InterfaceVeInstanceIpv6OspfNeighborCfg
		prefix6_6_6 := prefix6_6 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj6_6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighbor = d.Get(prefix6_6_6 + "neighbor").(string)
		obj6_6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeigInst = d.Get(prefix6_6_6 + "neig_inst").(int)
		obj6_6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighborCost = d.Get(prefix6_6_6 + "neighbor_cost").(int)
		obj6_6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighborPollInterval = d.Get(prefix6_6_6 + "neighbor_poll_interval").(int)
		obj6_6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighborPriority = d.Get(prefix6_6_6 + "neighbor_priority").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighbor = append(obj6_6.InterfaceVeInstanceIpv6OspfNeighborCfgNeighbor, obj6_6_6)
	}

	InterfaceVeInstanceIpv6OspfPriorityCfgCount := d.Get(prefix6_6 + "priority_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfPriorityCfgPriority = make([]go_thunder.InterfaceVeInstanceIpv6OspfPriorityCfg, 0, InterfaceVeInstanceIpv6OspfPriorityCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfPriorityCfgCount; i++ {
		var obj6_6_7 go_thunder.InterfaceVeInstanceIpv6OspfPriorityCfg
		prefix6_6_7 := prefix6_6 + fmt.Sprintf("priority_cfg.%d.", i)
		obj6_6_7.InterfaceVeInstanceIpv6OspfPriorityCfgPriority = d.Get(prefix6_6_7 + "priority").(int)
		obj6_6_7.InterfaceVeInstanceIpv6OspfPriorityCfgInstanceID = d.Get(prefix6_6_7 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfPriorityCfgPriority = append(obj6_6.InterfaceVeInstanceIpv6OspfPriorityCfgPriority, obj6_6_7)
	}

	InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgCount := d.Get(prefix6_6 + "retransmit_interval_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = make([]go_thunder.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfg, 0, InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgCount; i++ {
		var obj6_6_8 go_thunder.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfg
		prefix6_6_8 := prefix6_6 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj6_6_8.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = d.Get(prefix6_6_8 + "retransmit_interval").(int)
		obj6_6_8.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgInstanceID = d.Get(prefix6_6_8 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = append(obj6_6.InterfaceVeInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval, obj6_6_8)
	}

	InterfaceVeInstanceIpv6OspfTransmitDelayCfgCount := d.Get(prefix6_6 + "transmit_delay_cfg.#").(int)
	obj6_6.InterfaceVeInstanceIpv6OspfTransmitDelayCfgTransmitDelay = make([]go_thunder.InterfaceVeInstanceIpv6OspfTransmitDelayCfg, 0, InterfaceVeInstanceIpv6OspfTransmitDelayCfgCount)

	for i := 0; i < InterfaceVeInstanceIpv6OspfTransmitDelayCfgCount; i++ {
		var obj6_6_9 go_thunder.InterfaceVeInstanceIpv6OspfTransmitDelayCfg
		prefix6_6_9 := prefix6_6 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj6_6_9.InterfaceVeInstanceIpv6OspfTransmitDelayCfgTransmitDelay = d.Get(prefix6_6_9 + "transmit_delay").(int)
		obj6_6_9.InterfaceVeInstanceIpv6OspfTransmitDelayCfgInstanceID = d.Get(prefix6_6_9 + "instance_id").(int)
		obj6_6.InterfaceVeInstanceIpv6OspfTransmitDelayCfgTransmitDelay = append(obj6_6.InterfaceVeInstanceIpv6OspfTransmitDelayCfgTransmitDelay, obj6_6_9)
	}

	obj6.InterfaceVeInstanceIpv6OspfNetworkList = obj6_6

	c.InterfaceVeInstanceIpv6AddressList = obj6

	var obj7 go_thunder.InterfaceVeInstanceDdos
	prefix7 := "ddos.0."
	obj7.InterfaceVeInstanceDdosOutside = d.Get(prefix7 + "outside").(int)
	obj7.InterfaceVeInstanceDdosInside = d.Get(prefix7 + "inside").(int)

	c.InterfaceVeInstanceDdosOutside = obj7

	var obj8 go_thunder.InterfaceVeInstanceNptv6
	prefix8 := "nptv6.0."

	InterfaceVeInstanceNptv6DomainListCount := d.Get(prefix8 + "domain_list.#").(int)
	obj8.InterfaceVeInstanceNptv6DomainListDomainName = make([]go_thunder.InterfaceVeInstanceNptv6DomainList, 0, InterfaceVeInstanceNptv6DomainListCount)

	for i := 0; i < InterfaceVeInstanceNptv6DomainListCount; i++ {
		var obj8_1 go_thunder.InterfaceVeInstanceNptv6DomainList
		prefix8_1 := prefix8 + fmt.Sprintf("domain_list.%d.", i)
		obj8_1.InterfaceVeInstanceNptv6DomainListDomainName = d.Get(prefix8_1 + "domain_name").(string)
		obj8_1.InterfaceVeInstanceNptv6DomainListBindType = d.Get(prefix8_1 + "bind_type").(string)
		obj8.InterfaceVeInstanceNptv6DomainListDomainName = append(obj8.InterfaceVeInstanceNptv6DomainListDomainName, obj8_1)
	}

	c.InterfaceVeInstanceNptv6DomainList = obj8

	var obj9 go_thunder.InterfaceVeInstanceMap
	prefix9 := "map.0."
	obj9.InterfaceVeInstanceMapInside = d.Get(prefix9 + "inside").(int)
	obj9.InterfaceVeInstanceMapOutside = d.Get(prefix9 + "outside").(int)
	obj9.InterfaceVeInstanceMapMapTInside = d.Get(prefix9 + "map_t_inside").(int)
	obj9.InterfaceVeInstanceMapMapTOutside = d.Get(prefix9 + "map_t_outside").(int)

	c.InterfaceVeInstanceMapInside = obj9

	var obj10 go_thunder.InterfaceVeInstanceLw4O6
	prefix10 := "lw_4o6.0."
	obj10.InterfaceVeInstanceLw4O6Outside = d.Get(prefix10 + "outside").(int)
	obj10.InterfaceVeInstanceLw4O6Inside = d.Get(prefix10 + "inside").(int)

	c.InterfaceVeInstanceLw4O6Outside = obj10

	var obj11 go_thunder.InterfaceVeInstanceBfd
	prefix11 := "bfd.0."

	var obj11_1 go_thunder.InterfaceVeInstanceBfdAuthentication
	prefix11_1 := prefix11 + "authentication.0."
	obj11_1.InterfaceVeInstanceBfdAuthenticationKeyID = d.Get(prefix11_1 + "key_id").(int)
	obj11_1.InterfaceVeInstanceBfdAuthenticationMethod = d.Get(prefix11_1 + "method").(string)
	obj11_1.InterfaceVeInstanceBfdAuthenticationPassword = d.Get(prefix11_1 + "password").(string)

	obj11.InterfaceVeInstanceBfdAuthenticationKeyID = obj11_1

	obj11.InterfaceVeInstanceBfdEcho = d.Get(prefix11 + "echo").(int)
	obj11.InterfaceVeInstanceBfdDemand = d.Get(prefix11 + "demand").(int)

	var obj11_2 go_thunder.InterfaceVeInstanceBfdIntervalCfg
	prefix11_2 := prefix11 + "interval_cfg.0."
	obj11_2.InterfaceVeInstanceBfdIntervalCfgInterval = d.Get(prefix11_2 + "interval").(int)
	obj11_2.InterfaceVeInstanceBfdIntervalCfgMinRx = d.Get(prefix11_2 + "min_rx").(int)
	obj11_2.InterfaceVeInstanceBfdIntervalCfgMultiplier = d.Get(prefix11_2 + "multiplier").(int)

	obj11.InterfaceVeInstanceBfdIntervalCfgInterval = obj11_2

	c.InterfaceVeInstanceBfdAuthentication = obj11

	var obj12 go_thunder.InterfaceVeInstanceIsis
	prefix12 := "isis.0."

	var obj12_1 go_thunder.InterfaceVeInstanceIsisAuthentication
	prefix12_1 := prefix12 + "authentication.0."

	InterfaceVeInstanceIsisAuthenticationSendOnlyListCount := d.Get(prefix12_1 + "send_only_list.#").(int)
	obj12_1.InterfaceVeInstanceIsisAuthenticationSendOnlyListSendOnly = make([]go_thunder.InterfaceVeInstanceIsisAuthenticationSendOnlyList, 0, InterfaceVeInstanceIsisAuthenticationSendOnlyListCount)

	for i := 0; i < InterfaceVeInstanceIsisAuthenticationSendOnlyListCount; i++ {
		var obj12_1_1 go_thunder.InterfaceVeInstanceIsisAuthenticationSendOnlyList
		prefix12_1_1 := prefix12_1 + fmt.Sprintf("send_only_list.%d.", i)
		obj12_1_1.InterfaceVeInstanceIsisAuthenticationSendOnlyListSendOnly = d.Get(prefix12_1_1 + "send_only").(int)
		obj12_1_1.InterfaceVeInstanceIsisAuthenticationSendOnlyListLevel = d.Get(prefix12_1_1 + "level").(string)
		obj12_1.InterfaceVeInstanceIsisAuthenticationSendOnlyListSendOnly = append(obj12_1.InterfaceVeInstanceIsisAuthenticationSendOnlyListSendOnly, obj12_1_1)
	}

	InterfaceVeInstanceIsisAuthenticationModeListCount := d.Get(prefix12_1 + "mode_list.#").(int)
	obj12_1.InterfaceVeInstanceIsisAuthenticationModeListMode = make([]go_thunder.InterfaceVeInstanceIsisAuthenticationModeList, 0, InterfaceVeInstanceIsisAuthenticationModeListCount)

	for i := 0; i < InterfaceVeInstanceIsisAuthenticationModeListCount; i++ {
		var obj12_1_2 go_thunder.InterfaceVeInstanceIsisAuthenticationModeList
		prefix12_1_2 := prefix12_1 + fmt.Sprintf("mode_list.%d.", i)
		obj12_1_2.InterfaceVeInstanceIsisAuthenticationModeListMode = d.Get(prefix12_1_2 + "mode").(string)
		obj12_1_2.InterfaceVeInstanceIsisAuthenticationModeListLevel = d.Get(prefix12_1_2 + "level").(string)
		obj12_1.InterfaceVeInstanceIsisAuthenticationModeListMode = append(obj12_1.InterfaceVeInstanceIsisAuthenticationModeListMode, obj12_1_2)
	}

	InterfaceVeInstanceIsisAuthenticationKeyChainListCount := d.Get(prefix12_1 + "key_chain_list.#").(int)
	obj12_1.InterfaceVeInstanceIsisAuthenticationKeyChainListKeyChain = make([]go_thunder.InterfaceVeInstanceIsisAuthenticationKeyChainList, 0, InterfaceVeInstanceIsisAuthenticationKeyChainListCount)

	for i := 0; i < InterfaceVeInstanceIsisAuthenticationKeyChainListCount; i++ {
		var obj12_1_3 go_thunder.InterfaceVeInstanceIsisAuthenticationKeyChainList
		prefix12_1_3 := prefix12_1 + fmt.Sprintf("key_chain_list.%d.", i)
		obj12_1_3.InterfaceVeInstanceIsisAuthenticationKeyChainListKeyChain = d.Get(prefix12_1_3 + "key_chain").(string)
		obj12_1_3.InterfaceVeInstanceIsisAuthenticationKeyChainListLevel = d.Get(prefix12_1_3 + "level").(string)
		obj12_1.InterfaceVeInstanceIsisAuthenticationKeyChainListKeyChain = append(obj12_1.InterfaceVeInstanceIsisAuthenticationKeyChainListKeyChain, obj12_1_3)
	}

	obj12.InterfaceVeInstanceIsisAuthenticationSendOnlyList = obj12_1

	var obj12_2 go_thunder.InterfaceVeInstanceIsisBfdCfg
	prefix12_2 := prefix12 + "bfd_cfg.0."
	obj12_2.InterfaceVeInstanceIsisBfdCfgBfd = d.Get(prefix12_2 + "bfd").(int)
	obj12_2.InterfaceVeInstanceIsisBfdCfgDisable = d.Get(prefix12_2 + "disable").(int)

	obj12.InterfaceVeInstanceIsisBfdCfgBfd = obj12_2

	obj12.InterfaceVeInstanceIsisCircuitType = d.Get(prefix12 + "circuit_type").(string)

	InterfaceVeInstanceIsisCsnpIntervalListCount := d.Get(prefix12 + "csnp_interval_list.#").(int)
	obj12.InterfaceVeInstanceIsisCsnpIntervalListCsnpInterval = make([]go_thunder.InterfaceVeInstanceIsisCsnpIntervalList, 0, InterfaceVeInstanceIsisCsnpIntervalListCount)

	for i := 0; i < InterfaceVeInstanceIsisCsnpIntervalListCount; i++ {
		var obj12_3 go_thunder.InterfaceVeInstanceIsisCsnpIntervalList
		prefix12_3 := prefix12 + fmt.Sprintf("csnp_interval_list.%d.", i)
		obj12_3.InterfaceVeInstanceIsisCsnpIntervalListCsnpInterval = d.Get(prefix12_3 + "csnp_interval").(int)
		obj12_3.InterfaceVeInstanceIsisCsnpIntervalListLevel = d.Get(prefix12_3 + "level").(string)
		obj12.InterfaceVeInstanceIsisCsnpIntervalListCsnpInterval = append(obj12.InterfaceVeInstanceIsisCsnpIntervalListCsnpInterval, obj12_3)
	}

	obj12.InterfaceVeInstanceIsisPadding = d.Get(prefix12 + "padding").(int)

	InterfaceVeInstanceIsisHelloIntervalListCount := d.Get(prefix12 + "hello_interval_list.#").(int)
	obj12.InterfaceVeInstanceIsisHelloIntervalListHelloInterval = make([]go_thunder.InterfaceVeInstanceIsisHelloIntervalList, 0, InterfaceVeInstanceIsisHelloIntervalListCount)

	for i := 0; i < InterfaceVeInstanceIsisHelloIntervalListCount; i++ {
		var obj12_4 go_thunder.InterfaceVeInstanceIsisHelloIntervalList
		prefix12_4 := prefix12 + fmt.Sprintf("hello_interval_list.%d.", i)
		obj12_4.InterfaceVeInstanceIsisHelloIntervalListHelloInterval = d.Get(prefix12_4 + "hello_interval").(int)
		obj12_4.InterfaceVeInstanceIsisHelloIntervalListLevel = d.Get(prefix12_4 + "level").(string)
		obj12.InterfaceVeInstanceIsisHelloIntervalListHelloInterval = append(obj12.InterfaceVeInstanceIsisHelloIntervalListHelloInterval, obj12_4)
	}

	InterfaceVeInstanceIsisHelloIntervalMinimalListCount := d.Get(prefix12 + "hello_interval_minimal_list.#").(int)
	obj12.InterfaceVeInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = make([]go_thunder.InterfaceVeInstanceIsisHelloIntervalMinimalList, 0, InterfaceVeInstanceIsisHelloIntervalMinimalListCount)

	for i := 0; i < InterfaceVeInstanceIsisHelloIntervalMinimalListCount; i++ {
		var obj12_5 go_thunder.InterfaceVeInstanceIsisHelloIntervalMinimalList
		prefix12_5 := prefix12 + fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj12_5.InterfaceVeInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = d.Get(prefix12_5 + "hello_interval_minimal").(int)
		obj12_5.InterfaceVeInstanceIsisHelloIntervalMinimalListLevel = d.Get(prefix12_5 + "level").(string)
		obj12.InterfaceVeInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = append(obj12.InterfaceVeInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal, obj12_5)
	}

	InterfaceVeInstanceIsisHelloMultiplierListCount := d.Get(prefix12 + "hello_multiplier_list.#").(int)
	obj12.InterfaceVeInstanceIsisHelloMultiplierListHelloMultiplier = make([]go_thunder.InterfaceVeInstanceIsisHelloMultiplierList, 0, InterfaceVeInstanceIsisHelloMultiplierListCount)

	for i := 0; i < InterfaceVeInstanceIsisHelloMultiplierListCount; i++ {
		var obj12_6 go_thunder.InterfaceVeInstanceIsisHelloMultiplierList
		prefix12_6 := prefix12 + fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj12_6.InterfaceVeInstanceIsisHelloMultiplierListHelloMultiplier = d.Get(prefix12_6 + "hello_multiplier").(int)
		obj12_6.InterfaceVeInstanceIsisHelloMultiplierListLevel = d.Get(prefix12_6 + "level").(string)
		obj12.InterfaceVeInstanceIsisHelloMultiplierListHelloMultiplier = append(obj12.InterfaceVeInstanceIsisHelloMultiplierListHelloMultiplier, obj12_6)
	}

	obj12.InterfaceVeInstanceIsisLspInterval = d.Get(prefix12 + "lsp_interval").(int)

	var obj12_7 go_thunder.InterfaceVeInstanceIsisMeshGroup
	prefix12_7 := prefix12 + "mesh_group.0."
	obj12_7.InterfaceVeInstanceIsisMeshGroupValue = d.Get(prefix12_7 + "value").(int)
	obj12_7.InterfaceVeInstanceIsisMeshGroupBlocked = d.Get(prefix12_7 + "blocked").(int)

	obj12.InterfaceVeInstanceIsisMeshGroupValue = obj12_7

	InterfaceVeInstanceIsisMetricListCount := d.Get(prefix12 + "metric_list.#").(int)
	obj12.InterfaceVeInstanceIsisMetricListMetric = make([]go_thunder.InterfaceVeInstanceIsisMetricList, 0, InterfaceVeInstanceIsisMetricListCount)

	for i := 0; i < InterfaceVeInstanceIsisMetricListCount; i++ {
		var obj12_8 go_thunder.InterfaceVeInstanceIsisMetricList
		prefix12_8 := prefix12 + fmt.Sprintf("metric_list.%d.", i)
		obj12_8.InterfaceVeInstanceIsisMetricListMetric = d.Get(prefix12_8 + "metric").(int)
		obj12_8.InterfaceVeInstanceIsisMetricListLevel = d.Get(prefix12_8 + "level").(string)
		obj12.InterfaceVeInstanceIsisMetricListMetric = append(obj12.InterfaceVeInstanceIsisMetricListMetric, obj12_8)
	}

	obj12.InterfaceVeInstanceIsisNetwork = d.Get(prefix12 + "network").(string)

	InterfaceVeInstanceIsisPasswordListCount := d.Get(prefix12 + "password_list.#").(int)
	obj12.InterfaceVeInstanceIsisPasswordListPassword = make([]go_thunder.InterfaceVeInstanceIsisPasswordList, 0, InterfaceVeInstanceIsisPasswordListCount)

	for i := 0; i < InterfaceVeInstanceIsisPasswordListCount; i++ {
		var obj12_9 go_thunder.InterfaceVeInstanceIsisPasswordList
		prefix12_9 := prefix12 + fmt.Sprintf("password_list.%d.", i)
		obj12_9.InterfaceVeInstanceIsisPasswordListPassword = d.Get(prefix12_9 + "password").(string)
		obj12_9.InterfaceVeInstanceIsisPasswordListLevel = d.Get(prefix12_9 + "level").(string)
		obj12.InterfaceVeInstanceIsisPasswordListPassword = append(obj12.InterfaceVeInstanceIsisPasswordListPassword, obj12_9)
	}

	InterfaceVeInstanceIsisPriorityListCount := d.Get(prefix12 + "priority_list.#").(int)
	obj12.InterfaceVeInstanceIsisPriorityListPriority = make([]go_thunder.InterfaceVeInstanceIsisPriorityList, 0, InterfaceVeInstanceIsisPriorityListCount)

	for i := 0; i < InterfaceVeInstanceIsisPriorityListCount; i++ {
		var obj12_10 go_thunder.InterfaceVeInstanceIsisPriorityList
		prefix12_10 := prefix12 + fmt.Sprintf("priority_list.%d.", i)
		obj12_10.InterfaceVeInstanceIsisPriorityListPriority = d.Get(prefix12_10 + "priority").(int)
		obj12_10.InterfaceVeInstanceIsisPriorityListLevel = d.Get(prefix12_10 + "level").(string)
		obj12.InterfaceVeInstanceIsisPriorityListPriority = append(obj12.InterfaceVeInstanceIsisPriorityListPriority, obj12_10)
	}

	obj12.InterfaceVeInstanceIsisRetransmitInterval = d.Get(prefix12 + "retransmit_interval").(int)

	InterfaceVeInstanceIsisWideMetricListCount := d.Get(prefix12 + "wide_metric_list.#").(int)
	obj12.InterfaceVeInstanceIsisWideMetricListWideMetric = make([]go_thunder.InterfaceVeInstanceIsisWideMetricList, 0, InterfaceVeInstanceIsisWideMetricListCount)

	for i := 0; i < InterfaceVeInstanceIsisWideMetricListCount; i++ {
		var obj12_11 go_thunder.InterfaceVeInstanceIsisWideMetricList
		prefix12_11 := prefix12 + fmt.Sprintf("wide_metric_list.%d.", i)
		obj12_11.InterfaceVeInstanceIsisWideMetricListWideMetric = d.Get(prefix12_11 + "wide_metric").(int)
		obj12_11.InterfaceVeInstanceIsisWideMetricListLevel = d.Get(prefix12_11 + "level").(string)
		obj12.InterfaceVeInstanceIsisWideMetricListWideMetric = append(obj12.InterfaceVeInstanceIsisWideMetricListWideMetric, obj12_11)
	}

	c.InterfaceVeInstanceIsisAuthentication = obj12

	vc.InterfaceVeInstanceIfnum = c
	return vc
}
