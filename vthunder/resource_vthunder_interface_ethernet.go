package vthunder

//vThunder resource InterfaceEthernet

import (
	"fmt"
	"strconv"
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceEthernetCreate,
		Update: resourceInterfaceEthernetUpdate,
		Read:   resourceInterfaceEthernetRead,
		Delete: resourceInterfaceEthernetDelete,
		Schema: map[string]*schema.Schema{
			"speed_forced_40g": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"fec_forced_off": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"flow_control": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"speed": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"monitor_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mirror_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"monitor": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"traffic_distribution_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
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
			"icmpv6_rate_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"normal_v6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
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
						"stateful_firewall": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_list": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
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
									"class_list": {
										Type:        schema.TypeString,
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
						"access_list_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"inbound": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"v6_acl_name": {
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
															"instance_id": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"tag": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"area_id_addr": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"area_id_num": {
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
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
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
						"router_adver": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"managed_config_action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"hop_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"other_config_action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"rate_limit": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"prefix_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"valid_lifetime": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"not_autonomous": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"prefix": {
													Type:        schema.TypeString,
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
											},
										},
									},
									"use_floating_ip": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"retransmit_timer": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"min_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"adver_vrid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"floating_ip_default_vrid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"adver_vrid_default": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"max_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"use_floating_ip_default_vrid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"adver_mtu_disable": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"reachable_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"floating_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"adver_mtu": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"default_lifetime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"address_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv6_addr": {
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipv6_enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ospf": {
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
									"neighbor_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"neighbor_priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neig_inst": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_poll_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor_cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"neighbor": {
													Type:        schema.TypeString,
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
									"dead_interval_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dead_interval": {
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
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"retransmit_interval": {
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
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
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
									"priority_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"priority": {
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
									"network_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"network_instance_id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
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
											},
										},
									},
								},
							},
						},
						"ttl_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"map": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"map_t_outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"map_t_inside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
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
			"remove_vlan_tag": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"duplexity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"load_interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bfd": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interval_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multiplier": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
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
								},
							},
						},
						"echo": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"demand": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"encrypted": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"key_id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
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
						"mesh_group": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"blocked": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"value": {
										Type:        schema.TypeInt,
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
						"priority_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"priority": {
										Type:        schema.TypeInt,
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
						"csnp_interval_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"csnp_interval": {
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
						"network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"hello_interval_minimal_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"hello_interval_minimal": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"hello_interval_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"hello_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
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
						"circuit_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
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
						"authentication": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
						"lsp_interval": {
							Type:        schema.TypeInt,
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
						"stateful_firewall": {
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
									"access_list": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
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
									"class_list": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"server": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"slb_partition_redirect": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"allow_promiscuous_vip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_resp_time": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"address_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"ipv6_addr": {
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ospf": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ospf_ip_list": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"ip_addr": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"mtu_ignore": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"database_filter": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"retransmit_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dead_interval": {
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
															"md5": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"encrypted": {
																			Type:        schema.TypeString,
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
														},
													},
												},
												"priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"out": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"transmit_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"authentication_key": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"authentication": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"ospf_global": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cost": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"mtu_ignore": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"retransmit_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dead_interval": {
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
															"md5": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"encrypted": {
																			Type:        schema.TypeString,
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
														},
													},
												},
												"priority": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"uuid": {
													Type:        schema.TypeString,
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
															"point_to_multipoint": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"point_to_point": {
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
												"mtu": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"transmit_delay": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"disable": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"authentication_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"value": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"authentication": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"hello_interval": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"authentication_key": {
													Type:        schema.TypeString,
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
											},
										},
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
						"cache_spoofing_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"outside": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"receive_packet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
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
									"send_packet": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
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
									"send_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"version": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"send": {
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
									"authentication": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
								},
							},
						},
						"client": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"query_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"generate_membership_query": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"dhcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ttl_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"media_type_copper": {
				Type:        schema.TypeInt,
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
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"auto_neg_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
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
			"l3_vlan_fwd_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cpu_process": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"trap_source": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"cpu_process_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
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
			"fec_forced_on": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"name": {
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
						"lockup": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"normal": {
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
			"trunk_group_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"user_tag": {
							Type:        schema.TypeString,
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
						"port_priority": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"admin_key": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"trunk_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"lldp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tx_tlvs_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"system_capabilities": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"tx_tlvs": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"port_description": {
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
								},
							},
						},
						"tx_dot1_cfg": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"tx_dot1_tlvs": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
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
									"tx": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"rx": {
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
					},
				},
			},
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceEthernetCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernet (Inside resourceInterfaceEthernetCreate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernet(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernet --")
		d.SetId(strconv.Itoa(name))
		go_vthunder.PostInterfaceEthernet(client.Token, data, client.Host)

		return resourceInterfaceEthernetRead(d, meta)

	}
	return nil
}

func resourceInterfaceEthernetRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading InterfaceEthernet (Inside resourceInterfaceEthernetRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetInterfaceEthernet(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceInterfaceEthernetUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying InterfaceEthernet   (Inside resourceInterfaceEthernetUpdate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernet(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernet ")
		d.SetId(strconv.Itoa(name))
		go_vthunder.PutInterfaceEthernet(client.Token, strconv.Itoa(name), data, client.Host)

		return resourceInterfaceEthernetRead(d, meta)

	}
	return nil
}

func resourceInterfaceEthernetDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceEthernetDelete) " + name)
		err := go_vthunder.DeleteInterfaceEthernet(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToInterfaceEthernet(d *schema.ResourceData) go_vthunder.InterfaceEthernet {
	var vc go_vthunder.InterfaceEthernet
	var c go_vthunder.InterfaceEthernetInstance
	c.FecForcedOn = d.Get("fec_forced_on").(int)
	c.TrapSource = d.Get("trap_source").(int)

	var obj1 go_vthunder.EthernetIp
	prefix := "ip.0."

	AddressListCount := d.Get(prefix + "address_list.#").(int)
	obj1.AddressType = make([]go_vthunder.EthernetAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1_1 go_vthunder.EthernetAddressList
		prefix1 := prefix + fmt.Sprintf("address_list.%d.", i)
		obj1_1.AddressType = d.Get(prefix1 + "address_type").(string)
		obj1_1.Ipv6Addr = d.Get(prefix1 + "ipv6_addr").(string)
		obj1.AddressType = append(obj1.AddressType, obj1_1)
	}

	obj1.GenerateMembershipQuery = d.Get(prefix + "generate_membership_query").(int)
	obj1.CacheSpoofingPort = d.Get(prefix + "cache_spoofing_port").(int)
	obj1.Inside = d.Get(prefix + "inside").(int)
	obj1.AllowPromiscuousVip = d.Get(prefix + "allow_promiscuous_vip").(int)
	obj1.Client = d.Get(prefix + "client").(int)
	obj1.MaxRespTime = d.Get(prefix + "max_resp_time").(int)
	obj1.QueryInterval = d.Get(prefix + "query_interval").(int)
	obj1.Outside = d.Get(prefix + "outside").(int)

	HelperAddressListCount := d.Get(prefix + "helper_address_list.#").(int)
	obj1.HelperAddress = make([]go_vthunder.EthernetHelperAddressList, 0, HelperAddressListCount)

	for i := 0; i < HelperAddressListCount; i++ {
		var obj1_2 go_vthunder.EthernetHelperAddressList
		prefix1 := prefix1 + fmt.Sprintf("helper_address_list.%d.", i)
		obj1_2.HelperAddress = d.Get(prefix + "helper_address").(string)
		obj1.HelperAddress = append(obj1.HelperAddress, obj1_2)
	}

	var obj1_3 go_vthunder.EthernetStatefulFirewall
	prefix1 = prefix + "stateful_firewall.0."
	obj1_3.ClassList = d.Get(prefix1 + "class_list").(string)
	obj1_3.Inside = d.Get(prefix1 + "inside").(int)
	obj1_3.Outside = d.Get(prefix1 + "outside").(int)
	obj1_3.AclId = d.Get(prefix1 + "acl_id").(int)
	obj1_3.AccessList = d.Get(prefix1 + "access_list").(int)
	obj1.ClassList = obj1_3

	var obj1_4 go_vthunder.EthernetRip
	prefix1 = prefix + "rip.0."

	var obj1_4_1 go_vthunder.EthernetReceiveCfg
	prefix2 = prefix1 + "receive_cfg.0."
	obj1_4_1.Receive = d.Get(prefix2 + "receive").(int)
	obj1_4_1.Version = d.Get(prefix2 + "version").(string)
	obj1_4.Receive = obj1_4_1

	obj1_4.ReceivePacket = d.Get(prefix1 + "receive_packet").(int)

	var obj1_4_2 go_vthunder.EthernetSplitHorizonCfg
	prefix2 = prefix1 + "split_horizon_cfg.0."
	obj1_4_2.State = d.Get(prefix2 + "state").(string)
	obj1_4.State = obj1_4_2

	var obj1_4_3 go_vthunder.EthernetAuthentication
	prefix3 := prefix2 + "authentication.0."

	var obj1_4_3_1 go_vthunder.EthernetKeyChain
	prefix4 := prefix3 + "key_chain.0."
	obj1_4_3_1.KeyChain = d.Get(prefix4 + "key_chain").(string)
	obj1_4_3.KeyChain = obj1_4_3_1

	var obj1_4_3_2 go_vthunder.EthernetMode
	prefix4 = prefix3 + "mode.0."
	obj1_4_3_2.Mode = d.Get(prefix4 + "mode").(string)
	obj1_4_3.Mode = obj1_4_3_2

	var obj1_4_3_3 go_vthunder.EthernetStr
	prefix4 = prefix3 + "str.0."
	obj1_4_3_3.String = d.Get(prefix4 + "string").(string)
	obj1_4_3.String = obj1_4_3_3

	obj1_4.KeyChain = obj1_4_3

	var obj1_4_4 go_vthunder.EthernetSendCfg
	prefix2 = prefix1 + "send_cfg.0."
	obj1_4_4.Version = d.Get(prefix2 + "version").(string)
	obj1_4_4.Send = d.Get(prefix2 + "send").(int)
	obj1_4.Version = obj1_4_4

	obj1_4.SendPacket = d.Get(prefix1 + "send_packet").(int)
	obj1.ReceiveCfg = obj1_4

	obj1.TtlIgnore = d.Get(prefix + "ttl_ignore").(int)

	var obj1_5 go_vthunder.EthernetRouter
	prefix1 = "router.0."

	var obj1_5_1 go_vthunder.EthernetIsis
	prefix2 = prefix1 + "isis.0."
	obj1_5_1.Tag = d.Get(prefix2 + "tag").(string)
	obj1_5.Tag = obj1_5_1

	obj1.Isis = obj5

	obj1.Dhcp = d.Get(prefix1 + "dhcp").(int)
	obj1.Server = d.Get(prefix1 + "server").(int)

	var obj1_6 go_vthunder.EthernetOspf
	prefix1 = "ospf.0."

	OspfIpListCount := d.Get("ospf_ip_list.#").(int)
	obj1_6.DeadInterval = make([]go_vthunder.EthernetOspfIpList, 0, OspfIpListCount)

	for i := 0; i < OspfIpListCount; i++ {
		var obj1_6_1 go_vthunder.EthernetOspfIpList
		prefix := fmt.Sprintf("ospf_ip_list.%d.", i)
		obj1_6_1.DeadInterval = d.Get(prefix + "dead_interval").(int)
		obj1_6_1.AuthenticationKey = d.Get(prefix + "authentication_key").(string)
		obj1_6_1.MtuIgnore = d.Get(prefix + "mtu_ignore").(int)
		obj1_6_1.TransmitDelay = d.Get(prefix + "transmit_delay").(int)
		obj1_6_1.Value = d.Get(prefix + "value").(string)
		obj1_6_1.Priority = d.Get(prefix + "priority").(int)
		obj1_6_1.Authentication = d.Get(prefix + "authentication").(int)
		obj1_6_1.Cost = d.Get(prefix + "cost").(int)
		obj1_6_1.DatabaseFilter = d.Get(prefix + "database_filter").(string)
		obj1_6_1.HelloInterval = d.Get(prefix + "hello_interval").(int)
		obj1_6_1.IpAddr = d.Get(prefix + "ip_addr").(string)
		obj1_6_1.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)

		MessageDigestCfgCount := d.Get("message_digest_cfg.#").(int)
		obj1_6_1.MessageDigestKey = make([]go_vthunder.EthernetMessageDigestCfg, 0, MessageDigestCfgCount)

		for i := 0; i < MessageDigestCfgCount; i++ {
			var obj1_6_1_1 go_vthunder.EthernetMessageDigestCfg
			prefix = fmt.Sprintf("message_digest_cfg.%d.", i)
			obj1_6_1_1.MessageDigestKey = d.Get(prefix + "message_digest_key").(int)

			var obj1_6_1_1_1 go_vthunder.EthernetMd5
			prefix = "md5.0."
			obj1_6_1_1_1.Md5Value = d.Get(prefix + "md5_value").(string)
			obj1_6_1_1_1.Encrypted = d.Get(prefix + "encrypted").(string)
			obj1_6_1_1.Md5Value = obj1_6_1_1_1

			obj1_6_1.MessageDigestKey = append(obj1_6_1.MessageDigestKey, obj1_6_1_1)
		}

		obj1_6_1.Out = d.Get(prefix + "out").(int)
		obj6.DeadInterval = append(obj6.DeadInterval, obj1)
	}

	var obj2 go_vthunder.EthernetOspfGlobal
	prefix := "ospf_global.0."
	obj2.Cost = d.Get(prefix + "cost").(int)
	obj2.DeadInterval = d.Get(prefix + "dead_interval").(int)
	obj2.AuthenticationKey = d.Get(prefix + "authentication_key").(string)

	var obj1 go_vthunder.EthernetNetwork
	prefix := "network.0."
	obj1.Broadcast = d.Get(prefix + "broadcast").(int)
	obj1.PointToMultipoint = d.Get(prefix + "point_to_multipoint").(int)
	obj1.NonBroadcast = d.Get(prefix + "non_broadcast").(int)
	obj1.PointToPoint = d.Get(prefix + "point_to_point").(int)
	obj1.P2mpNbma = d.Get(prefix + "p2mp_nbma").(int)
	obj2.Broadcast = obj1

	obj2.MtuIgnore = d.Get(prefix + "mtu_ignore").(int)
	obj2.TransmitDelay = d.Get(prefix + "transmit_delay").(int)

	var obj2 go_vthunder.EthernetAuthenticationCfg
	prefix := "authentication_cfg.0."
	obj2.Authentication = d.Get(prefix + "authentication").(int)
	obj2.Value = d.Get(prefix + "value").(string)
	obj2.Authentication = obj2

	obj2.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)

	var obj3 go_vthunder.EthernetBfdCfg
	prefix := "bfd_cfg.0."
	obj3.Disable = d.Get(prefix + "disable").(int)
	obj3.Bfd = d.Get(prefix + "bfd").(int)
	obj2.Bfd = obj3

	obj2.Disable = d.Get(prefix + "disable").(string)
	obj2.HelloInterval = d.Get(prefix + "hello_interval").(int)

	var obj4 go_vthunder.EthernetDatabaseFilterCfg
	prefix := "database_filter_cfg.0."
	obj4.DatabaseFilter = d.Get(prefix + "database_filter").(string)
	obj4.Out = d.Get(prefix + "out").(int)
	obj2.DatabaseFilter = obj4

	obj2.Priority = d.Get(prefix + "priority").(int)
	obj2.Mtu = d.Get(prefix + "mtu").(int)

	MessageDigestCfgCount := d.Get("message_digest_cfg.#").(int)
	obj2.MessageDigestKey = make([]go_vthunder.EthernetMessageDigestCfg, 0, MessageDigestCfgCount)

	for i := 0; i < MessageDigestCfgCount; i++ {
		var obj5 go_vthunder.EthernetMessageDigestCfg
		prefix := fmt.Sprintf("message_digest_cfg.%d.", i)
		obj5.MessageDigestKey = d.Get(prefix + "message_digest_key").(int)

		var obj1 go_vthunder.EthernetMd5
		prefix := "md5.0."
		obj1.Md5Value = d.Get(prefix + "md5_value").(string)
		obj1.Encrypted = d.Get(prefix + "encrypted").(string)
		obj5.Md5Value = obj1

		obj2.MessageDigestKey = append(obj2.MessageDigestKey, obj5)
	}

	obj6.Cost = obj2

	obj1.OspfIpList = obj6

	obj1.SlbPartitionRedirect = d.Get(prefix + "slb_partition_redirect").(int)
	c.AddressList = obj1

	var obj2 go_vthunder.EthernetDdos
	prefix := "ddos.0."
	obj2.Outside = d.Get(prefix + "outside").(int)
	obj2.Inside = d.Get(prefix + "inside").(int)
	c.Outside = obj2

	c.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)

	var obj3 go_vthunder.EthernetAccessList
	prefix := "access_list.0."
	obj3.AclName = d.Get(prefix + "acl_name").(string)
	obj3.AclId = d.Get(prefix + "acl_id").(int)
	c.AclName = obj3

	c.Speed = d.Get("speed").(string)
	c.SpeedForced40g = d.Get("speed_forced_40g").(int)

	var obj4 go_vthunder.EthernetLldp
	prefix := "lldp.0."

	var obj1 go_vthunder.EthernetTxDot1Cfg
	prefix := "tx_dot1_cfg.0."
	obj1.LinkAggregation = d.Get(prefix + "link_aggregation").(int)
	obj1.Vlan = d.Get(prefix + "vlan").(int)
	obj1.TxDot1Tlvs = d.Get(prefix + "tx_dot1_tlvs").(int)
	obj4.LinkAggregation = obj1

	var obj2 go_vthunder.EthernetNotificationCfg
	prefix := "notification_cfg.0."
	obj2.Notification = d.Get(prefix + "notification").(int)
	obj2.NotifEnable = d.Get(prefix + "notif_enable").(int)
	obj4.Notification = obj2

	var obj3 go_vthunder.EthernetEnableCfg
	prefix := "enable_cfg.0."
	obj3.Rx = d.Get(prefix + "rx").(int)
	obj3.Tx = d.Get(prefix + "tx").(int)
	obj3.RtEnable = d.Get(prefix + "rt_enable").(int)
	obj4.Rx = obj3

	var obj4 go_vthunder.EthernetTxTlvsCfg
	prefix := "tx_tlvs_cfg.0."
	obj4.SystemCapabilities = d.Get(prefix + "system_capabilities").(int)
	obj4.SystemDescription = d.Get(prefix + "system_description").(int)
	obj4.ManagementAddress = d.Get(prefix + "management_address").(int)
	obj4.TxTlvs = d.Get(prefix + "tx_tlvs").(int)
	obj4.Exclude = d.Get(prefix + "exclude").(int)
	obj4.PortDescription = d.Get(prefix + "port_description").(int)
	obj4.SystemName = d.Get(prefix + "system_name").(int)
	obj4.SystemCapabilities = obj4

	c.TxDot1Cfg = obj4

	var obj5 go_vthunder.EthernetBfd
	prefix := "bfd.0."

	var obj1 go_vthunder.EthernetIntervalCfg
	prefix := "interval_cfg.0."
	obj1.Interval = d.Get(prefix + "interval").(int)
	obj1.MinRx = d.Get(prefix + "min_rx").(int)
	obj1.Multiplier = d.Get(prefix + "multiplier").(int)
	obj5.Interval = obj1

	var obj2 go_vthunder.EthernetAuthentication
	prefix := "authentication.0."
	obj2.Encrypted = d.Get(prefix + "encrypted").(string)
	obj2.Password = d.Get(prefix + "password").(string)
	obj2.Method = d.Get(prefix + "method").(string)
	obj2.KeyId = d.Get(prefix + "key_id").(int)
	obj5.Encrypted = obj2

	obj5.Echo = d.Get(prefix + "echo").(int)
	obj5.Demand = d.Get(prefix + "demand").(int)
	c.IntervalCfg = obj5

	c.MediaTypeCopper = d.Get("media_type_copper").(int)
	c.Ifnum = d.Get("ifnum").(int)
	c.RemoveVlanTag = d.Get("remove_vlan_tag").(int)

	MonitorListCount := d.Get("monitor_list.#").(int)
	c.MonitorVlan = make([]go_vthunder.EthernetMonitorList, 0, MonitorListCount)

	for i := 0; i < MonitorListCount; i++ {
		var obj6 go_vthunder.EthernetMonitorList
		prefix := fmt.Sprintf("monitor_list.%d.", i)
		obj6.MonitorVlan = d.Get(prefix + "monitor_vlan").(int)
		obj6.Monitor = d.Get(prefix + "monitor").(string)
		obj6.MirrorIndex = d.Get(prefix + "mirror_index").(int)
		c.MonitorVlan = append(c.MonitorVlan, obj6)
	}

	c.CpuProcess = d.Get("cpu_process").(int)
	c.AutoNegEnable = d.Get("auto_neg_enable").(int)

	var obj7 go_vthunder.EthernetMap
	prefix := "map.0."
	obj7.Inside = d.Get(prefix + "inside").(int)
	obj7.MapTInside = d.Get(prefix + "map_t_inside").(int)
	obj7.MapTOutside = d.Get(prefix + "map_t_outside").(int)
	obj7.Outside = d.Get(prefix + "outside").(int)
	c.Inside = obj7

	c.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)

	TrunkGroupListCount := d.Get("trunk_group_list.#").(int)
	c.TrunkNumber = make([]go_vthunder.EthernetTrunkGroupList, 0, TrunkGroupListCount)

	for i := 0; i < TrunkGroupListCount; i++ {
		var obj8 go_vthunder.EthernetTrunkGroupList
		prefix := fmt.Sprintf("trunk_group_list.%d.", i)
		obj8.TrunkNumber = d.Get(prefix + "trunk_number").(int)
		obj8.UserTag = d.Get(prefix + "user_tag").(string)

		var obj1 go_vthunder.EthernetUdldTimeoutCfg
		prefix := "udld_timeout_cfg.0."
		obj1.Slow = d.Get(prefix + "slow").(int)
		obj1.Fast = d.Get(prefix + "fast").(int)
		obj8.Slow = obj1

		obj8.Mode = d.Get(prefix + "mode").(string)
		obj8.Timeout = d.Get(prefix + "timeout").(string)
		obj8.Type = d.Get(prefix + "type").(string)
		obj8.AdminKey = d.Get(prefix + "admin_key").(int)
		obj8.PortPriority = d.Get(prefix + "port_priority").(int)
		c.TrunkNumber = append(c.TrunkNumber, obj8)
	}

	var obj9 go_vthunder.EthernetNptv6
	prefix := "nptv6.0."

	DomainListCount := d.Get("domain_list.#").(int)
	obj9.DomainName = make([]go_vthunder.EthernetDomainList, 0, DomainListCount)

	for i := 0; i < DomainListCount; i++ {
		var obj1 go_vthunder.EthernetDomainList
		prefix := fmt.Sprintf("domain_list.%d.", i)
		obj1.DomainName = d.Get(prefix + "domain_name").(string)
		obj1.BindType = d.Get(prefix + "bind_type").(string)
		obj9.DomainName = append(obj9.DomainName, obj1)
	}

	c.DomainList = obj9

	c.CpuProcessDir = d.Get("cpu_process_dir").(string)

	var obj10 go_vthunder.EthernetIsis
	prefix := "isis.0."

	PriorityListCount := d.Get("priority_list.#").(int)
	obj10.Priority = make([]go_vthunder.EthernetPriorityList, 0, PriorityListCount)

	for i := 0; i < PriorityListCount; i++ {
		var obj1 go_vthunder.EthernetPriorityList
		prefix := fmt.Sprintf("priority_list.%d.", i)
		obj1.Priority = d.Get(prefix + "priority").(int)
		obj1.Level = d.Get(prefix + "level").(string)
		obj10.Priority = append(obj10.Priority, obj1)
	}

	obj10.Padding = d.Get(prefix + "padding").(int)

	HelloIntervalMinimalListCount := d.Get("hello_interval_minimal_list.#").(int)
	obj10.HelloIntervalMinimal = make([]go_vthunder.EthernetHelloIntervalMinimalList, 0, HelloIntervalMinimalListCount)

	for i := 0; i < HelloIntervalMinimalListCount; i++ {
		var obj2 go_vthunder.EthernetHelloIntervalMinimalList
		prefix := fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj2.HelloIntervalMinimal = d.Get(prefix + "hello_interval_minimal").(int)
		obj2.Level = d.Get(prefix + "level").(string)
		obj10.HelloIntervalMinimal = append(obj10.HelloIntervalMinimal, obj2)
	}

	var obj3 go_vthunder.EthernetMeshGroup
	prefix := "mesh_group.0."
	obj3.Value = d.Get(prefix + "value").(int)
	obj3.Blocked = d.Get(prefix + "blocked").(int)
	obj10.Value = obj3

	obj10.Network = d.Get(prefix + "network").(string)

	var obj4 go_vthunder.EthernetAuthentication
	prefix := "authentication.0."

	SendOnlyListCount := d.Get("send_only_list.#").(int)
	obj4.SendOnly = make([]go_vthunder.EthernetSendOnlyList, 0, SendOnlyListCount)

	for i := 0; i < SendOnlyListCount; i++ {
		var obj1 go_vthunder.EthernetSendOnlyList
		prefix := fmt.Sprintf("send_only_list.%d.", i)
		obj1.SendOnly = d.Get(prefix + "send_only").(int)
		obj1.Level = d.Get(prefix + "level").(string)
		obj4.SendOnly = append(obj4.SendOnly, obj1)
	}

	ModeListCount := d.Get("mode_list.#").(int)
	obj4.Mode = make([]go_vthunder.EthernetModeList, 0, ModeListCount)

	for i := 0; i < ModeListCount; i++ {
		var obj2 go_vthunder.EthernetModeList
		prefix := fmt.Sprintf("mode_list.%d.", i)
		obj2.Mode = d.Get(prefix + "mode").(string)
		obj2.Level = d.Get(prefix + "level").(string)
		obj4.Mode = append(obj4.Mode, obj2)
	}

	KeyChainListCount := d.Get("key_chain_list.#").(int)
	obj4.KeyChain = make([]go_vthunder.EthernetKeyChainList, 0, KeyChainListCount)

	for i := 0; i < KeyChainListCount; i++ {
		var obj3 go_vthunder.EthernetKeyChainList
		prefix := fmt.Sprintf("key_chain_list.%d.", i)
		obj3.KeyChain = d.Get(prefix + "key_chain").(string)
		obj3.Level = d.Get(prefix + "level").(string)
		obj4.KeyChain = append(obj4.KeyChain, obj3)
	}

	obj10.SendOnlyList = obj4

	CsnpIntervalListCount := d.Get("csnp_interval_list.#").(int)
	obj10.CsnpInterval = make([]go_vthunder.EthernetCsnpIntervalList, 0, CsnpIntervalListCount)

	for i := 0; i < CsnpIntervalListCount; i++ {
		var obj5 go_vthunder.EthernetCsnpIntervalList
		prefix := fmt.Sprintf("csnp_interval_list.%d.", i)
		obj5.CsnpInterval = d.Get(prefix + "csnp_interval").(int)
		obj5.Level = d.Get(prefix + "level").(string)
		obj10.CsnpInterval = append(obj10.CsnpInterval, obj5)
	}

	obj10.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)

	PasswordListCount := d.Get("password_list.#").(int)
	obj10.Password = make([]go_vthunder.EthernetPasswordList, 0, PasswordListCount)

	for i := 0; i < PasswordListCount; i++ {
		var obj6 go_vthunder.EthernetPasswordList
		prefix := fmt.Sprintf("password_list.%d.", i)
		obj6.Password = d.Get(prefix + "password").(string)
		obj6.Level = d.Get(prefix + "level").(string)
		obj10.Password = append(obj10.Password, obj6)
	}

	var obj7 go_vthunder.EthernetBfdCfg
	prefix := "bfd_cfg.0."
	obj7.Disable = d.Get(prefix + "disable").(int)
	obj7.Bfd = d.Get(prefix + "bfd").(int)
	obj10.Disable = obj7

	WideMetricListCount := d.Get("wide_metric_list.#").(int)
	obj10.WideMetric = make([]go_vthunder.EthernetWideMetricList, 0, WideMetricListCount)

	for i := 0; i < WideMetricListCount; i++ {
		var obj8 go_vthunder.EthernetWideMetricList
		prefix := fmt.Sprintf("wide_metric_list.%d.", i)
		obj8.WideMetric = d.Get(prefix + "wide_metric").(int)
		obj8.Level = d.Get(prefix + "level").(string)
		obj10.WideMetric = append(obj10.WideMetric, obj8)
	}

	HelloIntervalListCount := d.Get("hello_interval_list.#").(int)
	obj10.HelloInterval = make([]go_vthunder.EthernetHelloIntervalList, 0, HelloIntervalListCount)

	for i := 0; i < HelloIntervalListCount; i++ {
		var obj9 go_vthunder.EthernetHelloIntervalList
		prefix := fmt.Sprintf("hello_interval_list.%d.", i)
		obj9.HelloInterval = d.Get(prefix + "hello_interval").(int)
		obj9.Level = d.Get(prefix + "level").(string)
		obj10.HelloInterval = append(obj10.HelloInterval, obj9)
	}

	obj10.CircuitType = d.Get(prefix + "circuit_type").(string)

	HelloMultiplierListCount := d.Get("hello_multiplier_list.#").(int)
	obj10.HelloMultiplier = make([]go_vthunder.EthernetHelloMultiplierList, 0, HelloMultiplierListCount)

	for i := 0; i < HelloMultiplierListCount; i++ {
		var obj10 go_vthunder.EthernetHelloMultiplierList
		prefix := fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj10.HelloMultiplier = d.Get(prefix + "hello_multiplier").(int)
		obj10.Level = d.Get(prefix + "level").(string)
		obj10.HelloMultiplier = append(obj10.HelloMultiplier, obj10)
	}

	MetricListCount := d.Get("metric_list.#").(int)
	obj10.Metric = make([]go_vthunder.EthernetMetricList, 0, MetricListCount)

	for i := 0; i < MetricListCount; i++ {
		var obj11 go_vthunder.EthernetMetricList
		prefix := fmt.Sprintf("metric_list.%d.", i)
		obj11.Metric = d.Get(prefix + "metric").(int)
		obj11.Level = d.Get(prefix + "level").(string)
		obj10.Metric = append(obj10.Metric, obj11)
	}

	obj10.LspInterval = d.Get(prefix + "lsp_interval").(int)
	c.PriorityList = obj10

	c.Name = d.Get("name").(string)
	c.Duplexity = d.Get("duplexity").(string)

	var obj11 go_vthunder.EthernetIcmpv6RateLimit
	prefix := "icmpv6_rate_limit.0."
	obj11.LockupPeriodV6 = d.Get(prefix + "lockup_period_v6").(int)
	obj11.NormalV6 = d.Get(prefix + "normal_v6").(int)
	obj11.LockupV6 = d.Get(prefix + "lockup_v6").(int)
	c.LockupPeriodV6 = obj11

	c.UserTag = d.Get("user_tag").(string)
	c.Mtu = d.Get("mtu").(int)

	var obj12 go_vthunder.EthernetIpv6
	prefix := "ipv6.0."

	AddressListCount := d.Get("address_list.#").(int)
	obj12.AddressType = make([]go_vthunder.EthernetAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1 go_vthunder.EthernetAddressList
		prefix := fmt.Sprintf("address_list.%d.", i)
		obj1.AddressType = d.Get(prefix + "address_type").(string)
		obj1.Ipv6Addr = d.Get(prefix + "ipv6_addr").(string)
		obj12.AddressType = append(obj12.AddressType, obj1)
	}

	obj12.Inside = d.Get(prefix + "inside").(int)
	obj12.Ipv6Enable = d.Get(prefix + "ipv6_enable").(int)

	var obj2 go_vthunder.EthernetRip
	prefix := "rip.0."

	var obj1 go_vthunder.EthernetSplitHorizonCfg
	prefix := "split_horizon_cfg.0."
	obj1.State = d.Get(prefix + "state").(string)
	obj2.State = obj1

	obj12.SplitHorizonCfg = obj2

	obj12.Outside = d.Get(prefix + "outside").(int)

	var obj3 go_vthunder.EthernetStatefulFirewall
	prefix := "stateful_firewall.0."
	obj3.ClassList = d.Get(prefix + "class_list").(string)
	obj3.AclName = d.Get(prefix + "acl_name").(string)
	obj3.Inside = d.Get(prefix + "inside").(int)
	obj3.Outside = d.Get(prefix + "outside").(int)
	obj3.AccessList = d.Get(prefix + "access_list").(int)
	obj12.ClassList = obj3

	obj12.TtlIgnore = d.Get(prefix + "ttl_ignore").(int)

	var obj4 go_vthunder.EthernetRouter
	prefix := "router.0."

	var obj1 go_vthunder.EthernetRipng
	prefix := "ripng.0."
	obj1.Rip = d.Get(prefix + "rip").(int)
	obj4.Uuid = obj1

	var obj2 go_vthunder.EthernetOspf
	prefix := "ospf.0."

	AreaListCount := d.Get("area_list.#").(int)
	obj2.AreaIdAddr = make([]go_vthunder.EthernetAreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj1 go_vthunder.EthernetAreaList
		prefix := fmt.Sprintf("area_list.%d.", i)
		obj1.AreaIdAddr = d.Get(prefix + "area_id_addr").(string)
		obj1.Tag = d.Get(prefix + "tag").(string)
		obj1.InstanceId = d.Get(prefix + "instance_id").(int)
		obj1.AreaIdNum = d.Get(prefix + "area_id_num").(int)
		obj2.AreaIdAddr = append(obj2.AreaIdAddr, obj1)
	}

	obj4.AreaList = obj2

	var obj3 go_vthunder.EthernetIsis
	prefix := "isis.0."
	obj3.Tag = d.Get(prefix + "tag").(string)
	obj4.Tag = obj3

	obj12.Ripng = obj4

	var obj5 go_vthunder.EthernetAccessListCfg
	prefix := "access_list_cfg.0."
	obj5.Inbound = d.Get(prefix + "inbound").(int)
	obj5.V6AclName = d.Get(prefix + "v6_acl_name").(string)
	obj12.Inbound = obj5

	var obj6 go_vthunder.EthernetOspf
	prefix := "ospf.0."
	obj6.Bfd = d.Get(prefix + "bfd").(int)

	CostCfgCount := d.Get("cost_cfg.#").(int)
	obj6.Cost = make([]go_vthunder.EthernetCostCfg, 0, CostCfgCount)

	for i := 0; i < CostCfgCount; i++ {
		var obj1 go_vthunder.EthernetCostCfg
		prefix := fmt.Sprintf("cost_cfg.%d.", i)
		obj1.Cost = d.Get(prefix + "cost").(int)
		obj1.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.Cost = append(obj6.Cost, obj1)
	}

	PriorityCfgCount := d.Get("priority_cfg.#").(int)
	obj6.Priority = make([]go_vthunder.EthernetPriorityCfg, 0, PriorityCfgCount)

	for i := 0; i < PriorityCfgCount; i++ {
		var obj2 go_vthunder.EthernetPriorityCfg
		prefix := fmt.Sprintf("priority_cfg.%d.", i)
		obj2.Priority = d.Get(prefix + "priority").(int)
		obj2.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.Priority = append(obj6.Priority, obj2)
	}

	HelloIntervalCfgCount := d.Get("hello_interval_cfg.#").(int)
	obj6.HelloInterval = make([]go_vthunder.EthernetHelloIntervalCfg, 0, HelloIntervalCfgCount)

	for i := 0; i < HelloIntervalCfgCount; i++ {
		var obj3 go_vthunder.EthernetHelloIntervalCfg
		prefix := fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj3.HelloInterval = d.Get(prefix + "hello_interval").(int)
		obj3.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.HelloInterval = append(obj6.HelloInterval, obj3)
	}

	MtuIgnoreCfgCount := d.Get("mtu_ignore_cfg.#").(int)
	obj6.MtuIgnore = make([]go_vthunder.EthernetMtuIgnoreCfg, 0, MtuIgnoreCfgCount)

	for i := 0; i < MtuIgnoreCfgCount; i++ {
		var obj4 go_vthunder.EthernetMtuIgnoreCfg
		prefix := fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj4.MtuIgnore = d.Get(prefix + "mtu_ignore").(int)
		obj4.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.MtuIgnore = append(obj6.MtuIgnore, obj4)
	}

	RetransmitIntervalCfgCount := d.Get("retransmit_interval_cfg.#").(int)
	obj6.RetransmitInterval = make([]go_vthunder.EthernetRetransmitIntervalCfg, 0, RetransmitIntervalCfgCount)

	for i := 0; i < RetransmitIntervalCfgCount; i++ {
		var obj5 go_vthunder.EthernetRetransmitIntervalCfg
		prefix := fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj5.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)
		obj5.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.RetransmitInterval = append(obj6.RetransmitInterval, obj5)
	}

	obj6.Disable = d.Get(prefix + "disable").(int)

	TransmitDelayCfgCount := d.Get("transmit_delay_cfg.#").(int)
	obj6.TransmitDelay = make([]go_vthunder.EthernetTransmitDelayCfg, 0, TransmitDelayCfgCount)

	for i := 0; i < TransmitDelayCfgCount; i++ {
		var obj6 go_vthunder.EthernetTransmitDelayCfg
		prefix := fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj6.TransmitDelay = d.Get(prefix + "transmit_delay").(int)
		obj6.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.TransmitDelay = append(obj6.TransmitDelay, obj6)
	}

	NeighborCfgCount := d.Get("neighbor_cfg.#").(int)
	obj6.NeighborPriority = make([]go_vthunder.EthernetNeighborCfg, 0, NeighborCfgCount)

	for i := 0; i < NeighborCfgCount; i++ {
		var obj7 go_vthunder.EthernetNeighborCfg
		prefix := fmt.Sprintf("neighbor_cfg.%d.", i)
		obj7.NeighborPriority = d.Get(prefix + "neighbor_priority").(int)
		obj7.NeigInst = d.Get(prefix + "neig_inst").(int)
		obj7.NeighborPollInterval = d.Get(prefix + "neighbor_poll_interval").(int)
		obj7.NeighborCost = d.Get(prefix + "neighbor_cost").(int)
		obj7.Neighbor = d.Get(prefix + "neighbor").(string)
		obj6.NeighborPriority = append(obj6.NeighborPriority, obj7)
	}

	NetworkListCount := d.Get("network_list.#").(int)
	obj6.BroadcastType = make([]go_vthunder.EthernetNetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj8 go_vthunder.EthernetNetworkList
		prefix := fmt.Sprintf("network_list.%d.", i)
		obj8.BroadcastType = d.Get(prefix + "broadcast_type").(string)
		obj8.P2mpNbma = d.Get(prefix + "p2mp_nbma").(int)
		obj8.NetworkInstanceId = d.Get(prefix + "network_instance_id").(int)
		obj6.BroadcastType = append(obj6.BroadcastType, obj8)
	}

	DeadIntervalCfgCount := d.Get("dead_interval_cfg.#").(int)
	obj6.DeadInterval = make([]go_vthunder.EthernetDeadIntervalCfg, 0, DeadIntervalCfgCount)

	for i := 0; i < DeadIntervalCfgCount; i++ {
		var obj9 go_vthunder.EthernetDeadIntervalCfg
		prefix := fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj9.DeadInterval = d.Get(prefix + "dead_interval").(int)
		obj9.InstanceId = d.Get(prefix + "instance_id").(int)
		obj6.DeadInterval = append(obj6.DeadInterval, obj9)
	}

	obj12.Bfd = obj6

	var obj7 go_vthunder.EthernetRouterAdver
	prefix := "router_adver.0."
	obj7.MaxInterval = d.Get(prefix + "max_interval").(int)
	obj7.DefaultLifetime = d.Get(prefix + "default_lifetime").(int)
	obj7.ReachableTime = d.Get(prefix + "reachable_time").(int)
	obj7.OtherConfigAction = d.Get(prefix + "other_config_action").(string)
	obj7.FloatingIpDefaultVrid = d.Get(prefix + "floating_ip_default_vrid").(string)
	obj7.ManagedConfigAction = d.Get(prefix + "managed_config_action").(string)
	obj7.MinInterval = d.Get(prefix + "min_interval").(int)
	obj7.RateLimit = d.Get(prefix + "rate_limit").(int)
	obj7.AdverMtuDisable = d.Get(prefix + "adver_mtu_disable").(int)

	PrefixListCount := d.Get("prefix_list.#").(int)
	obj7.NotAutonomous = make([]go_vthunder.EthernetPrefixList, 0, PrefixListCount)

	for i := 0; i < PrefixListCount; i++ {
		var obj1 go_vthunder.EthernetPrefixList
		prefix := fmt.Sprintf("prefix_list.%d.", i)
		obj1.NotAutonomous = d.Get(prefix + "not_autonomous").(int)
		obj1.ValidLifetime = d.Get(prefix + "valid_lifetime").(int)
		obj1.NotOnLink = d.Get(prefix + "not_on_link").(int)
		obj1.Prefix = d.Get(prefix + "prefix").(string)
		obj1.PreferredLifetime = d.Get(prefix + "preferred_lifetime").(int)
		obj7.NotAutonomous = append(obj7.NotAutonomous, obj1)
	}

	obj7.FloatingIp = d.Get(prefix + "floating_ip").(string)
	obj7.AdverVrid = d.Get(prefix + "adver_vrid").(int)
	obj7.UseFloatingIpDefaultVrid = d.Get(prefix + "use_floating_ip_default_vrid").(int)
	obj7.Action = d.Get(prefix + "action").(string)
	obj7.AdverVridDefault = d.Get(prefix + "adver_vrid_default").(int)
	obj7.AdverMtu = d.Get(prefix + "adver_mtu").(int)
	obj7.RetransmitTimer = d.Get(prefix + "retransmit_timer").(int)
	obj7.HopLimit = d.Get(prefix + "hop_limit").(int)
	obj7.UseFloatingIp = d.Get(prefix + "use_floating_ip").(int)
	obj12.MaxInterval = obj7

	c.AddressList = obj12

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_vthunder.EthernetSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj13 go_vthunder.EthernetSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj13.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj13)
	}

	c.LoadInterval = d.Get("load_interval").(int)

	var obj14 go_vthunder.EthernetLw4o6
	prefix := "lw_4o6.0."
	obj14.Outside = d.Get(prefix + "outside").(int)
	obj14.Inside = d.Get(prefix + "inside").(int)
	c.Outside = obj14

	c.Action = d.Get("action").(string)
	c.FecForcedOff = d.Get("fec_forced_off").(int)

	var obj15 go_vthunder.EthernetIcmpRateLimit
	prefix := "icmp_rate_limit.0."
	obj15.Lockup = d.Get(prefix + "lockup").(int)
	obj15.LockupPeriod = d.Get(prefix + "lockup_period").(int)
	obj15.Normal = d.Get(prefix + "normal").(int)
	c.Lockup = obj15

	c.FlowControl = d.Get("flow_control").(int)

	vc.UUID = c
	return vc
}
