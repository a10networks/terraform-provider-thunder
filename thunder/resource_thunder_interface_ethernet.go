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

func resourceInterfaceEthernetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceEthernet (Inside resourceInterfaceEthernetCreate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernet(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernet --")
		d.SetId(strconv.Itoa(name))
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

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading InterfaceEthernet (Inside resourceInterfaceEthernetRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceEthernet(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceInterfaceEthernetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying InterfaceEthernet   (Inside resourceInterfaceEthernetUpdate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceEthernet(d)
		logger.Println("[INFO] received V from method data to InterfaceEthernet ")
		d.SetId(strconv.Itoa(name))
		err := go_thunder.PutInterfaceEthernet(client.Token, strconv.Itoa(name), data, client.Host)
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
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceEthernetDelete) " + name)
		err := go_thunder.DeleteInterfaceEthernet(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToInterfaceEthernet(d *schema.ResourceData) go_thunder.InterfaceEthernet {
	var vc go_thunder.InterfaceEthernet
	var c go_thunder.InterfaceEthernetInstance
	c.FecForcedOn = d.Get("fec_forced_on").(int)
	c.TrapSource = d.Get("trap_source").(int)

	var obj1 go_thunder.EthernetIP
	prefix := "ip.0."

	AddressListCount := d.Get(prefix + "address_list.#").(int)
	obj1.AddressType = make([]go_thunder.EthernetAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj1_1 go_thunder.EthernetAddressList
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
	obj1.HelperAddress = make([]go_thunder.EthernetHelperAddressList, 0, HelperAddressListCount)

	for i := 0; i < HelperAddressListCount; i++ {
		var obj1_2 go_thunder.EthernetHelperAddressList
		prefix1 := prefix + fmt.Sprintf("helper_address_list.%d.", i)
		obj1_2.HelperAddress = d.Get(prefix1 + "helper_address").(string)
		obj1.HelperAddress = append(obj1.HelperAddress, obj1_2)
	}

	var obj1_3 go_thunder.EthernetStatefulFirewallId
	prefix1 := prefix + "stateful_firewall.0."
	obj1_3.ClassList = d.Get(prefix1 + "class_list").(string)
	obj1_3.Inside = d.Get(prefix1 + "inside").(int)
	obj1_3.Outside = d.Get(prefix1 + "outside").(int)
	obj1_3.ACLID = d.Get(prefix1 + "acl_id").(int)
	obj1_3.AccessList = d.Get(prefix1 + "access_list").(int)
	obj1.ClassList = obj1_3

	var obj1_4 go_thunder.EthernetRip
	prefix1 = prefix + "rip.0."

	var obj1_4_1 go_thunder.EthernetReceiveCfg
	prefix2 := prefix1 + "receive_cfg.0."
	obj1_4_1.Receive = d.Get(prefix2 + "receive").(int)
	obj1_4_1.Version = d.Get(prefix2 + "version").(string)
	obj1_4.Receive = obj1_4_1

	obj1_4.ReceivePacket = d.Get(prefix1 + "receive_packet").(int)

	var obj1_4_2 go_thunder.EthernetSplitHorizonCfg
	prefix2 = prefix1 + "split_horizon_cfg.0."
	obj1_4_2.State = d.Get(prefix2 + "state").(string)
	obj1_4.State = obj1_4_2

	var obj1_4_3 go_thunder.EthernetAuthentication
	prefix2 = prefix1 + "authentication.0."

	var obj1_4_3_1 go_thunder.EthernetKeyChain
	prefix3 := prefix2 + "key_chain.0."
	obj1_4_3_1.KeyChain = d.Get(prefix3 + "key_chain").(string)
	obj1_4_3.KeyChain = obj1_4_3_1

	var obj1_4_3_2 go_thunder.EthernetMode
	prefix3 = prefix2 + "mode.0."
	obj1_4_3_2.Mode = d.Get(prefix3 + "mode").(string)
	obj1_4_3.Mode = obj1_4_3_2

	var obj1_4_3_3 go_thunder.EthernetStr
	prefix3 = prefix2 + "str.0."
	obj1_4_3_3.String = d.Get(prefix3 + "string").(string)
	obj1_4_3.String = obj1_4_3_3

	obj1_4.KeyChain = obj1_4_3

	var obj1_4_4 go_thunder.EthernetSendCfg
	prefix2 = prefix1 + "send_cfg.0."
	obj1_4_4.Version = d.Get(prefix2 + "version").(string)
	obj1_4_4.Send = d.Get(prefix2 + "send").(int)
	obj1_4.Version = obj1_4_4

	obj1_4.SendPacket = d.Get(prefix1 + "send_packet").(int)
	obj1.ReceiveCfg = obj1_4

	obj1.TTLIgnore = d.Get(prefix + "ttl_ignore").(int)

	var obj1_5 go_thunder.EthernetRouter
	prefix1 = prefix + "router.0."

	var obj1_5_1 go_thunder.EthernetIsis
	prefix2 = prefix1 + "isis.0."
	obj1_5_1.Tag = d.Get(prefix2 + "tag").(string)
	obj1_5.Tag = obj1_5_1

	obj1.Tag = obj1_5

	obj1.Dhcp = d.Get(prefix + "dhcp").(int)
	obj1.Server = d.Get(prefix + "server").(int)

	var obj1_6 go_thunder.EthernetOspf
	prefix1 = prefix + "ospf.0."

	OspfIpListCount := d.Get(prefix1 + "ospf_ip_list.#").(int)
	obj1_6.DeadInterval = make([]go_thunder.EthernetOspfIPList, 0, OspfIpListCount)

	for i := 0; i < OspfIpListCount; i++ {
		var obj1_6_1 go_thunder.EthernetOspfIPList
		prefix2 = prefix1 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj1_6_1.DeadInterval = d.Get(prefix2 + "dead_interval").(int)
		obj1_6_1.AuthenticationKey = d.Get(prefix2 + "authentication_key").(string)
		obj1_6_1.MtuIgnore = d.Get(prefix2 + "mtu_ignore").(int)
		obj1_6_1.TransmitDelay = d.Get(prefix2 + "transmit_delay").(int)
		obj1_6_1.Value = d.Get(prefix2 + "value").(string)
		obj1_6_1.Priority = d.Get(prefix2 + "priority").(int)
		obj1_6_1.Authentication = d.Get(prefix2 + "authentication").(int)
		obj1_6_1.Cost = d.Get(prefix2 + "cost").(int)
		obj1_6_1.DatabaseFilter = d.Get(prefix2 + "database_filter").(string)
		obj1_6_1.HelloInterval = d.Get(prefix2 + "hello_interval").(int)
		obj1_6_1.IPAddr = d.Get(prefix2 + "ip_addr").(string)
		obj1_6_1.RetransmitInterval = d.Get(prefix2 + "retransmit_interval").(int)

		MessageDigestCfgCount := d.Get(prefix2 + "message_digest_cfg.#").(int)
		obj1_6_1.MessageDigestKey = make([]go_thunder.EthernetMessageDigestCfg, 0, MessageDigestCfgCount)

		for i := 0; i < MessageDigestCfgCount; i++ {
			var obj1_6_1_1 go_thunder.EthernetMessageDigestCfg
			prefix3 = prefix2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj1_6_1_1.MessageDigestKey = d.Get(prefix3 + "message_digest_key").(int)

			var obj1_6_1_1_1 go_thunder.EthernetMd5
			prefix4 := prefix3 + "md5.0."
			obj1_6_1_1_1.Md5Value = d.Get(prefix4 + "md5_value").(string)
			obj1_6_1_1_1.Encrypted = d.Get(prefix4 + "encrypted").(string)
			obj1_6_1_1.Md5Value = obj1_6_1_1_1

			obj1_6_1.MessageDigestKey = append(obj1_6_1.MessageDigestKey, obj1_6_1_1)
		}

		obj1_6_1.Out = d.Get(prefix2 + "out").(int)
		obj1_6.DeadInterval = append(obj1_6.DeadInterval, obj1_6_1)
	}

	var obj1_6_2 go_thunder.EthernetOspfGlobal
	prefix2 = prefix1 + "ospf_global.0."
	obj1_6_2.Cost = d.Get(prefix2 + "cost").(int)
	obj1_6_2.DeadInterval = d.Get(prefix2 + "dead_interval").(int)
	obj1_6_2.AuthenticationKey = d.Get(prefix2 + "authentication_key").(string)

	var obj1_6_2_1 go_thunder.EthernetNetwork
	prefix3 = prefix2 + "network.0."
	obj1_6_2_1.Broadcast = d.Get(prefix3 + "broadcast").(int)
	obj1_6_2_1.PointToMultipoint = d.Get(prefix3 + "point_to_multipoint").(int)
	obj1_6_2_1.NonBroadcast = d.Get(prefix3 + "non_broadcast").(int)
	obj1_6_2_1.PointToPoint = d.Get(prefix3 + "point_to_point").(int)
	obj1_6_2_1.P2MpNbma = d.Get(prefix3 + "p2mp_nbma").(int)
	obj1_6_2.Broadcast = obj1_6_2_1

	obj1_6_2.MtuIgnore = d.Get(prefix2 + "mtu_ignore").(int)
	obj1_6_2.TransmitDelay = d.Get(prefix2 + "transmit_delay").(int)

	var obj1_6_2_2 go_thunder.EthernetAuthenticationCfg
	prefix3 = prefix2 + "authentication_cfg.0."
	obj1_6_2_2.Authentication = d.Get(prefix3 + "authentication").(int)
	obj1_6_2_2.Value = d.Get(prefix3 + "value").(string)
	obj1_6_2.Authentication = obj1_6_2_2

	obj1_6_2.RetransmitInterval = d.Get(prefix2 + "retransmit_interval").(int)

	var obj1_6_2_3 go_thunder.EthernetBfdCfg
	prefix3 = prefix2 + "bfd_cfg.0."
	obj1_6_2_3.Disable = d.Get(prefix3 + "disable").(int)
	obj1_6_2_3.Bfd = d.Get(prefix3 + "bfd").(int)
	obj1_6_2.Bfd = obj1_6_2_3

	obj1_6_2.Disable = d.Get(prefix2 + "disable").(string)
	obj1_6_2.HelloInterval = d.Get(prefix2 + "hello_interval").(int)

	var obj1_6_2_4 go_thunder.EthernetDatabaseFilterCfg
	prefix3 = prefix2 + "database_filter_cfg.0."
	obj1_6_2_4.DatabaseFilter = d.Get(prefix3 + "database_filter").(string)
	obj1_6_2_4.Out = d.Get(prefix3 + "out").(int)
	obj1_6_2.DatabaseFilter = obj1_6_2_4

	obj1_6_2.Priority = d.Get(prefix2 + "priority").(int)
	obj1_6_2.Mtu = d.Get(prefix2 + "mtu").(int)

	MessageDigestCfgCount := d.Get(prefix2 + "message_digest_cfg.#").(int)
	obj1_6_2.MessageDigestKey = make([]go_thunder.EthernetMessageDigestCfg, 0, MessageDigestCfgCount)

	for i := 0; i < MessageDigestCfgCount; i++ {
		var obj1_6_2_5 go_thunder.EthernetMessageDigestCfg
		prefix3 = prefix2 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj1_6_2_5.MessageDigestKey = d.Get(prefix3 + "message_digest_key").(int)

		var obj1_6_2_5_1 go_thunder.EthernetMd5
		prefix4 := prefix3 + "md5.0."
		obj1_6_2_5_1.Md5Value = d.Get(prefix4 + "md5_value").(string)
		obj1_6_2_5_1.Encrypted = d.Get(prefix4 + "encrypted").(string)
		obj1_6_2_5.Md5Value = obj1_6_2_5_1

		obj1_6_2.MessageDigestKey = append(obj1_6_2.MessageDigestKey, obj1_6_2_5)
	}

	obj1_6.Cost = obj1_6_2

	obj1.DeadInterval = obj1_6

	obj1.SlbPartitionRedirect = d.Get(prefix + "slb_partition_redirect").(int)
	c.AddressType = obj1

	var obj2 go_thunder.EthernetDdos
	prefix = "ddos.0."
	obj2.Outside = d.Get(prefix + "outside").(int)
	obj2.Inside = d.Get(prefix + "inside").(int)
	c.Outside = obj2

	c.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)

	var obj3 go_thunder.EthernetAccessList
	prefix = "access_list.0."
	obj3.ACLName = d.Get(prefix + "acl_name").(string)
	obj3.ACLID = d.Get(prefix + "acl_id").(int)
	c.ACLName = obj3

	c.Speed = d.Get("speed").(string)
	c.SpeedForced40G = d.Get("speed_forced_40g").(int)

	var obj4 go_thunder.EthernetLldp
	prefix = "lldp.0."

	var obj4_1 go_thunder.EthernetTxDot1Cfg
	prefix1 = prefix + "tx_dot1_cfg.0."
	obj4_1.LinkAggregation = d.Get(prefix1 + "link_aggregation").(int)
	obj4_1.Vlan = d.Get(prefix1 + "vlan").(int)
	obj4_1.TxDot1Tlvs = d.Get(prefix1 + "tx_dot1_tlvs").(int)
	obj4.LinkAggregation = obj4_1

	var obj4_2 go_thunder.EthernetNotificationCfg
	prefix1 = prefix + "notification_cfg.0."
	obj4_2.Notification = d.Get(prefix1 + "notification").(int)
	obj4_2.NotifEnable = d.Get(prefix1 + "notif_enable").(int)
	obj4.Notification = obj4_2

	var obj4_3 go_thunder.EthernetEnableCfg
	prefix1 = prefix + "enable_cfg.0."
	obj4_3.Rx = d.Get(prefix1 + "rx").(int)
	obj4_3.Tx = d.Get(prefix1 + "tx").(int)
	obj4_3.RtEnable = d.Get(prefix1 + "rt_enable").(int)
	obj4.Rx = obj4_3

	var obj4_4 go_thunder.EthernetTxTlvsCfg
	prefix1 = prefix + "tx_tlvs_cfg.0."
	obj4_4.SystemCapabilities = d.Get(prefix1 + "system_capabilities").(int)
	obj4_4.SystemDescription = d.Get(prefix1 + "system_description").(int)
	obj4_4.ManagementAddress = d.Get(prefix1 + "management_address").(int)
	obj4_4.TxTlvs = d.Get(prefix1 + "tx_tlvs").(int)
	obj4_4.Exclude = d.Get(prefix1 + "exclude").(int)
	obj4_4.PortDescription = d.Get(prefix1 + "port_description").(int)
	obj4_4.SystemName = d.Get(prefix1 + "system_name").(int)
	obj4.SystemCapabilities = obj4_4

	c.LinkAggregation = obj4

	var obj5 go_thunder.EthernetBfd
	prefix = "bfd.0."

	var obj5_1 go_thunder.EthernetIntervalCfg
	prefix1 = prefix + "interval_cfg.0."
	obj5_1.Interval = d.Get(prefix1 + "interval").(int)
	obj5_1.MinRx = d.Get(prefix1 + "min_rx").(int)
	obj5_1.Multiplier = d.Get(prefix1 + "multiplier").(int)
	obj5.Interval = obj5_1

	var obj5_2 go_thunder.EthernetAuthentication2
	prefix1 = prefix + "authentication.0."
	obj5_2.Encrypted = d.Get(prefix1 + "encrypted").(string)
	obj5_2.Password = d.Get(prefix1 + "password").(string)
	obj5_2.Method = d.Get(prefix1 + "method").(string)
	obj5_2.KeyID = d.Get(prefix1 + "key_id").(int)
	obj5.Encrypted = obj5_2

	obj5.Echo = d.Get(prefix + "echo").(int)
	obj5.Demand = d.Get(prefix + "demand").(int)
	c.Interval = obj5

	c.MediaTypeCopper = d.Get("media_type_copper").(int)
	c.Ifnum = d.Get("ifnum").(int)
	c.RemoveVlanTag = d.Get("remove_vlan_tag").(int)

	MonitorListCount := d.Get("monitor_list.#").(int)
	c.MonitorVlan = make([]go_thunder.EthernetMonitorList, 0, MonitorListCount)

	for i := 0; i < MonitorListCount; i++ {
		var obj6 go_thunder.EthernetMonitorList
		prefix = fmt.Sprintf("monitor_list.%d.", i)
		obj6.MonitorVlan = d.Get(prefix + "monitor_vlan").(int)
		obj6.Monitor = d.Get(prefix + "monitor").(string)
		obj6.MirrorIndex = d.Get(prefix + "mirror_index").(int)
		c.MonitorVlan = append(c.MonitorVlan, obj6)
	}

	c.CPUProcess = d.Get("cpu_process").(int)
	c.AutoNegEnable = d.Get("auto_neg_enable").(int)

	var obj7 go_thunder.EthernetMap
	prefix = "map.0."
	obj7.Inside = d.Get(prefix + "inside").(int)
	obj7.MapTInside = d.Get(prefix + "map_t_inside").(int)
	obj7.MapTOutside = d.Get(prefix + "map_t_outside").(int)
	obj7.Outside = d.Get(prefix + "outside").(int)
	c.Inside = obj7

	c.TrafficDistributionMode = d.Get("traffic_distribution_mode").(string)

	TrunkGroupListCount := d.Get("trunk_group_list.#").(int)
	c.TrunkNumber = make([]go_thunder.EthernetTrunkGroupList, 0, TrunkGroupListCount)

	for i := 0; i < TrunkGroupListCount; i++ {
		var obj8 go_thunder.EthernetTrunkGroupList
		prefix = fmt.Sprintf("trunk_group_list.%d.", i)
		obj8.TrunkNumber = d.Get(prefix + "trunk_number").(int)
		obj8.UserTag = d.Get(prefix + "user_tag").(string)

		var obj8_1 go_thunder.EthernetUdldTimeoutCfg
		prefix1 = prefix + "udld_timeout_cfg.0."
		obj8_1.Slow = d.Get(prefix1 + "slow").(int)
		obj8_1.Fast = d.Get(prefix1 + "fast").(int)
		obj8.Slow = obj8_1

		obj8.Mode = d.Get(prefix + "mode").(string)
		obj8.Timeout = d.Get(prefix + "timeout").(string)
		obj8.Type = d.Get(prefix + "type").(string)
		obj8.AdminKey = d.Get(prefix + "admin_key").(int)
		obj8.PortPriority = d.Get(prefix + "port_priority").(int)
		c.TrunkNumber = append(c.TrunkNumber, obj8)
	}

	var obj9 go_thunder.EthernetNptv6
	prefix = "nptv6.0."

	DomainListCount := d.Get(prefix + "domain_list.#").(int)
	obj9.DomainName = make([]go_thunder.EthernetDomainList, 0, DomainListCount)

	for i := 0; i < DomainListCount; i++ {
		var obj9_1 go_thunder.EthernetDomainList
		prefix1 := prefix + fmt.Sprintf("domain_list.%d.", i)
		obj9_1.DomainName = d.Get(prefix1 + "domain_name").(string)
		obj9_1.BindType = d.Get(prefix1 + "bind_type").(string)
		obj9.DomainName = append(obj9.DomainName, obj9_1)
	}

	c.Nptv6 = obj9

	c.CPUProcessDir = d.Get("cpu_process_dir").(string)

	var obj10 go_thunder.EthernetIsis2
	prefix = "isis.0."

	PriorityListCount := d.Get(prefix + "priority_list.#").(int)
	obj10.Priority = make([]go_thunder.EthernetPriorityList, 0, PriorityListCount)

	for i := 0; i < PriorityListCount; i++ {
		var obj10_1 go_thunder.EthernetPriorityList
		prefix1 = prefix + fmt.Sprintf("priority_list.%d.", i)
		obj10_1.Priority = d.Get(prefix1 + "priority").(int)
		obj10_1.Level = d.Get(prefix1 + "level").(string)
		obj10.Priority = append(obj10.Priority, obj10_1)
	}

	obj10.Padding = d.Get(prefix + "padding").(int)

	HelloIntervalMinimalListCount := d.Get(prefix + "hello_interval_minimal_list.#").(int)
	obj10.HelloIntervalMinimal = make([]go_thunder.EthernetHelloIntervalMinimalList, 0, HelloIntervalMinimalListCount)

	for i := 0; i < HelloIntervalMinimalListCount; i++ {
		var obj10_2 go_thunder.EthernetHelloIntervalMinimalList
		prefix1 = prefix + fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj10_2.HelloIntervalMinimal = d.Get(prefix1 + "hello_interval_minimal").(int)
		obj10_2.Level = d.Get(prefix1 + "level").(string)
		obj10.HelloIntervalMinimal = append(obj10.HelloIntervalMinimal, obj10_2)
	}

	var obj10_3 go_thunder.EthernetMeshGroup
	prefix1 = prefix + "mesh_group.0."
	obj10_3.Value = d.Get(prefix1 + "value").(int)
	obj10_3.Blocked = d.Get(prefix1 + "blocked").(int)
	obj10.Value = obj10_3

	obj10.Network = d.Get(prefix + "network").(string)

	var obj10_4 go_thunder.EthernetAuthentication3
	prefix1 = prefix + "authentication.0."

	SendOnlyListCount := d.Get(prefix1 + "send_only_list.#").(int)
	obj10_4.SendOnly = make([]go_thunder.EthernetSendOnlyList, 0, SendOnlyListCount)

	for i := 0; i < SendOnlyListCount; i++ {
		var obj10_4_1 go_thunder.EthernetSendOnlyList
		prefix2 = prefix1 + fmt.Sprintf("send_only_list.%d.", i)
		obj10_4_1.SendOnly = d.Get(prefix2 + "send_only").(int)
		obj10_4_1.Level = d.Get(prefix2 + "level").(string)
		obj10_4.SendOnly = append(obj10_4.SendOnly, obj10_4_1)
	}

	ModeListCount := d.Get(prefix1 + "mode_list.#").(int)
	obj10_4.Mode = make([]go_thunder.EthernetModeList, 0, ModeListCount)

	for i := 0; i < ModeListCount; i++ {
		var obj10_4_2 go_thunder.EthernetModeList
		prefix2 = prefix1 + fmt.Sprintf("mode_list.%d.", i)
		obj10_4_2.Mode = d.Get(prefix2 + "mode").(string)
		obj10_4_2.Level = d.Get(prefix2 + "level").(string)
		obj10_4.Mode = append(obj10_4.Mode, obj10_4_2)
	}

	KeyChainListCount := d.Get(prefix1 + "key_chain_list.#").(int)
	obj10_4.KeyChain = make([]go_thunder.EthernetKeyChainList, 0, KeyChainListCount)

	for i := 0; i < KeyChainListCount; i++ {
		var obj10_4_3 go_thunder.EthernetKeyChainList
		prefix2 = prefix1 + fmt.Sprintf("key_chain_list.%d.", i)
		obj10_4_3.KeyChain = d.Get(prefix2 + "key_chain").(string)
		obj10_4_3.Level = d.Get(prefix2 + "level").(string)
		obj10_4.KeyChain = append(obj10_4.KeyChain, obj10_4_3)
	}

	obj10.SendOnly = obj10_4

	CsnpIntervalListCount := d.Get(prefix + "csnp_interval_list.#").(int)
	obj10.CsnpInterval = make([]go_thunder.EthernetCsnpIntervalList, 0, CsnpIntervalListCount)

	for i := 0; i < CsnpIntervalListCount; i++ {
		var obj10_5 go_thunder.EthernetCsnpIntervalList
		prefix1 = prefix + fmt.Sprintf("csnp_interval_list.%d.", i)
		obj10_5.CsnpInterval = d.Get(prefix1 + "csnp_interval").(int)
		obj10_5.Level = d.Get(prefix1 + "level").(string)
		obj10.CsnpInterval = append(obj10.CsnpInterval, obj10_5)
	}

	obj10.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)

	PasswordListCount := d.Get(prefix + "password_list.#").(int)
	obj10.Password = make([]go_thunder.EthernetPasswordList, 0, PasswordListCount)

	for i := 0; i < PasswordListCount; i++ {
		var obj10_6 go_thunder.EthernetPasswordList
		prefix1 = prefix + fmt.Sprintf("password_list.%d.", i)
		obj10_6.Password = d.Get(prefix1 + "password").(string)
		obj10_6.Level = d.Get(prefix1 + "level").(string)
		obj10.Password = append(obj10.Password, obj10_6)
	}

	var obj10_7 go_thunder.EthernetBfdCfg
	prefix1 = prefix + "bfd_cfg.0."
	obj10_7.Disable = d.Get(prefix1 + "disable").(int)
	obj10_7.Bfd = d.Get(prefix1 + "bfd").(int)
	obj10.Disable = obj10_7

	WideMetricListCount := d.Get(prefix + "wide_metric_list.#").(int)
	obj10.WideMetric = make([]go_thunder.EthernetWideMetricList, 0, WideMetricListCount)

	for i := 0; i < WideMetricListCount; i++ {
		var obj10_8 go_thunder.EthernetWideMetricList
		prefix1 = prefix + fmt.Sprintf("wide_metric_list.%d.", i)
		obj10_8.WideMetric = d.Get(prefix1 + "wide_metric").(int)
		obj10_8.Level = d.Get(prefix1 + "level").(string)
		obj10.WideMetric = append(obj10.WideMetric, obj10_8)
	}

	HelloIntervalListCount := d.Get(prefix + "hello_interval_list.#").(int)
	obj10.HelloInterval = make([]go_thunder.EthernetHelloIntervalList, 0, HelloIntervalListCount)

	for i := 0; i < HelloIntervalListCount; i++ {
		var obj10_9 go_thunder.EthernetHelloIntervalList
		prefix1 = prefix + fmt.Sprintf("hello_interval_list.%d.", i)
		obj10_9.HelloInterval = d.Get(prefix1 + "hello_interval").(int)
		obj10_9.Level = d.Get(prefix1 + "level").(string)
		obj10.HelloInterval = append(obj10.HelloInterval, obj10_9)
	}

	obj10.CircuitType = d.Get(prefix + "circuit_type").(string)

	HelloMultiplierListCount := d.Get(prefix + "hello_multiplier_list.#").(int)
	obj10.HelloMultiplier = make([]go_thunder.EthernetHelloMultiplierList, 0, HelloMultiplierListCount)

	for i := 0; i < HelloMultiplierListCount; i++ {
		var obj10_10 go_thunder.EthernetHelloMultiplierList
		prefix1 = prefix + fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj10_10.HelloMultiplier = d.Get(prefix1 + "hello_multiplier").(int)
		obj10_10.Level = d.Get(prefix1 + "level").(string)
		obj10.HelloMultiplier = append(obj10.HelloMultiplier, obj10_10)
	}

	MetricListCount := d.Get(prefix + "metric_list.#").(int)
	obj10.Metric = make([]go_thunder.EthernetMetricList, 0, MetricListCount)

	for i := 0; i < MetricListCount; i++ {
		var obj10_11 go_thunder.EthernetMetricList
		prefix1 = prefix + fmt.Sprintf("metric_list.%d.", i)
		obj10_11.Metric = d.Get(prefix1 + "metric").(int)
		obj10_11.Level = d.Get(prefix1 + "level").(string)
		obj10.Metric = append(obj10.Metric, obj10_11)
	}

	obj10.LspInterval = d.Get(prefix + "lsp_interval").(int)
	c.Priority = obj10

	c.Name = d.Get("name").(string)
	c.Duplexity = d.Get("duplexity").(string)

	var obj11 go_thunder.EthernetIcmpv6RateLimit
	prefix = "icmpv6_rate_limit.0."
	obj11.LockupPeriodV6 = d.Get(prefix + "lockup_period_v6").(int)
	obj11.NormalV6 = d.Get(prefix + "normal_v6").(int)
	obj11.LockupV6 = d.Get(prefix + "lockup_v6").(int)
	c.LockupPeriodV6 = obj11

	c.UserTag = d.Get("user_tag").(string)
	c.Mtu = d.Get("mtu").(int)

	var obj12 go_thunder.EthernetIpv6
	prefix = "ipv6.0."

	AddressListCount = d.Get(prefix + "address_list.#").(int)
	obj12.AddressType = make([]go_thunder.EthernetAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj12_1 go_thunder.EthernetAddressList
		prefix1 = prefix + fmt.Sprintf("address_list.%d.", i)
		obj12_1.AddressType = d.Get(prefix1 + "address_type").(string)
		obj12_1.Ipv6Addr = d.Get(prefix1 + "ipv6_addr").(string)
		obj12.AddressType = append(obj12.AddressType, obj12_1)
	}

	obj12.Inside = d.Get(prefix + "inside").(int)
	obj12.Ipv6Enable = d.Get(prefix + "ipv6_enable").(int)

	var obj12_2 go_thunder.EthernetRip2
	prefix1 = prefix + "rip.0."

	var obj12_2_1 go_thunder.EthernetSplitHorizonCfg
	prefix2 = prefix1 + "split_horizon_cfg.0."
	obj12_2_1.State = d.Get(prefix2 + "state").(string)
	obj12_2.State = obj12_2_1

	obj12.State = obj12_2

	obj12.Outside = d.Get(prefix + "outside").(int)

	var obj12_3 go_thunder.EthernetStatefulFirewall
	prefix1 = prefix + "stateful_firewall.0."
	obj12_3.ClassList = d.Get(prefix1 + "class_list").(string)
	obj12_3.ACLName = d.Get(prefix1 + "acl_name").(string)
	obj12_3.Inside = d.Get(prefix1 + "inside").(int)
	obj12_3.Outside = d.Get(prefix1 + "outside").(int)
	obj12_3.AccessList = d.Get(prefix1 + "access_list").(int)
	obj12.ClassList = obj12_3

	obj12.TTLIgnore = d.Get(prefix + "ttl_ignore").(int)

	var obj12_4 go_thunder.EthernetRouter2
	prefix1 = prefix + "router.0."

	var obj12_4_1 go_thunder.EthernetRipng
	prefix2 = prefix1 + "ripng.0."
	obj12_4_1.Rip = d.Get(prefix2 + "rip").(int)
	obj12_4.UUID = obj12_4_1

	var obj12_4_2 go_thunder.EthernetOspf2
	prefix2 = prefix1 + "ospf.0."

	AreaListCount := d.Get(prefix2 + "area_list.#").(int)
	obj12_4_2.AreaIDAddr = make([]go_thunder.EthernetAreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj12_4_2_1 go_thunder.EthernetAreaList
		prefix3 = prefix2 + fmt.Sprintf("area_list.%d.", i)
		obj12_4_2_1.AreaIDAddr = d.Get(prefix3 + "area_id_addr").(string)
		obj12_4_2_1.Tag = d.Get(prefix3 + "tag").(string)
		obj12_4_2_1.InstanceID = d.Get(prefix3 + "instance_id").(int)
		obj12_4_2_1.AreaIDNum = d.Get(prefix3 + "area_id_num").(int)
		obj12_4_2.AreaIDAddr = append(obj12_4_2.AreaIDAddr, obj12_4_2_1)
	}

	obj12_4.AreaIDAddr = obj12_4_2

	var obj12_4_3 go_thunder.EthernetIsis
	prefix2 = prefix1 + "isis.0."
	obj12_4_3.Tag = d.Get(prefix2 + "tag").(string)
	obj12_4.Tag = obj12_4_3

	obj12.AreaIDAddr = obj12_4

	var obj12_5 go_thunder.EthernetAccessListCfg
	prefix1 = prefix + "access_list_cfg.0."
	obj12_5.Inbound = d.Get(prefix1 + "inbound").(int)
	obj12_5.V6ACLName = d.Get(prefix1 + "v6_acl_name").(string)
	obj12.Inbound = obj12_5

	var obj12_6 go_thunder.EthernetOspf3
	prefix1 = prefix + "ospf.0."
	obj12_6.Bfd = d.Get(prefix1 + "bfd").(int)

	CostCfgCount := d.Get(prefix1 + "cost_cfg.#").(int)
	obj12_6.Cost = make([]go_thunder.EthernetCostCfg, 0, CostCfgCount)

	for i := 0; i < CostCfgCount; i++ {
		var obj12_6_1 go_thunder.EthernetCostCfg
		prefix2 = prefix1 + fmt.Sprintf("cost_cfg.%d.", i)
		obj12_6_1.Cost = d.Get(prefix2 + "cost").(int)
		obj12_6_1.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.Cost = append(obj12_6.Cost, obj12_6_1)
	}

	PriorityCfgCount := d.Get(prefix1 + "priority_cfg.#").(int)
	obj12_6.Priority = make([]go_thunder.EthernetPriorityCfg, 0, PriorityCfgCount)

	for i := 0; i < PriorityCfgCount; i++ {
		var obj12_6_2 go_thunder.EthernetPriorityCfg
		prefix2 = prefix1 + fmt.Sprintf("priority_cfg.%d.", i)
		obj12_6_2.Priority = d.Get(prefix2 + "priority").(int)
		obj12_6_2.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.Priority = append(obj12_6.Priority, obj12_6_2)
	}

	HelloIntervalCfgCount := d.Get(prefix1 + "hello_interval_cfg.#").(int)
	obj12_6.HelloInterval = make([]go_thunder.EthernetHelloIntervalCfg, 0, HelloIntervalCfgCount)

	for i := 0; i < HelloIntervalCfgCount; i++ {
		var obj12_6_3 go_thunder.EthernetHelloIntervalCfg
		prefix2 = prefix1 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj12_6_3.HelloInterval = d.Get(prefix2 + "hello_interval").(int)
		obj12_6_3.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.HelloInterval = append(obj12_6.HelloInterval, obj12_6_3)
	}

	MtuIgnoreCfgCount := d.Get(prefix1 + "mtu_ignore_cfg.#").(int)
	obj12_6.MtuIgnore = make([]go_thunder.EthernetMtuIgnoreCfg, 0, MtuIgnoreCfgCount)

	for i := 0; i < MtuIgnoreCfgCount; i++ {
		var obj12_6_4 go_thunder.EthernetMtuIgnoreCfg
		prefix2 := prefix1 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj12_6_4.MtuIgnore = d.Get(prefix2 + "mtu_ignore").(int)
		obj12_6_4.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.MtuIgnore = append(obj12_6.MtuIgnore, obj12_6_4)
	}

	RetransmitIntervalCfgCount := d.Get(prefix1 + "retransmit_interval_cfg.#").(int)
	obj12_6.RetransmitInterval = make([]go_thunder.EthernetRetransmitIntervalCfg, 0, RetransmitIntervalCfgCount)

	for i := 0; i < RetransmitIntervalCfgCount; i++ {
		var obj12_6_5 go_thunder.EthernetRetransmitIntervalCfg
		prefix2 = prefix1 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj12_6_5.RetransmitInterval = d.Get(prefix2 + "retransmit_interval").(int)
		obj12_6_5.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.RetransmitInterval = append(obj12_6.RetransmitInterval, obj12_6_5)
	}

	obj12_6.Disable = d.Get(prefix1 + "disable").(int)

	TransmitDelayCfgCount := d.Get(prefix1 + "transmit_delay_cfg.#").(int)
	obj12_6.TransmitDelay = make([]go_thunder.EthernetTransmitDelayCfg, 0, TransmitDelayCfgCount)

	for i := 0; i < TransmitDelayCfgCount; i++ {
		var obj12_6_6 go_thunder.EthernetTransmitDelayCfg
		prefix2 = prefix1 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj12_6_6.TransmitDelay = d.Get(prefix2 + "transmit_delay").(int)
		obj12_6_6.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.TransmitDelay = append(obj12_6.TransmitDelay, obj12_6_6)
	}

	NeighborCfgCount := d.Get(prefix1 + "neighbor_cfg.#").(int)
	obj12_6.NeighborPriority = make([]go_thunder.EthernetNeighborCfg, 0, NeighborCfgCount)

	for i := 0; i < NeighborCfgCount; i++ {
		var obj12_6_7 go_thunder.EthernetNeighborCfg
		prefix2 = prefix1 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj12_6_7.NeighborPriority = d.Get(prefix2 + "neighbor_priority").(int)
		obj12_6_7.NeigInst = d.Get(prefix2 + "neig_inst").(int)
		obj12_6_7.NeighborPollInterval = d.Get(prefix2 + "neighbor_poll_interval").(int)
		obj12_6_7.NeighborCost = d.Get(prefix2 + "neighbor_cost").(int)
		obj12_6_7.Neighbor = d.Get(prefix2 + "neighbor").(string)
		obj12_6.NeighborPriority = append(obj12_6.NeighborPriority, obj12_6_7)
	}

	NetworkListCount := d.Get(prefix1 + "network_list.#").(int)
	obj12_6.BroadcastType = make([]go_thunder.EthernetNetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj12_6_8 go_thunder.EthernetNetworkList
		prefix2 = prefix1 + fmt.Sprintf("network_list.%d.", i)
		obj12_6_8.BroadcastType = d.Get(prefix2 + "broadcast_type").(string)
		obj12_6_8.P2MpNbma = d.Get(prefix2 + "p2mp_nbma").(int)
		obj12_6_8.NetworkInstanceID = d.Get(prefix2 + "network_instance_id").(int)
		obj12_6.BroadcastType = append(obj12_6.BroadcastType, obj12_6_8)
	}

	DeadIntervalCfgCount := d.Get(prefix1 + "dead_interval_cfg.#").(int)
	obj12_6.DeadInterval = make([]go_thunder.EthernetDeadIntervalCfg, 0, DeadIntervalCfgCount)

	for i := 0; i < DeadIntervalCfgCount; i++ {
		var obj12_6_9 go_thunder.EthernetDeadIntervalCfg
		prefix2 = prefix1 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj12_6_9.DeadInterval = d.Get(prefix2 + "dead_interval").(int)
		obj12_6_9.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj12_6.DeadInterval = append(obj12_6.DeadInterval, obj12_6_9)
	}

	obj12.Bfd = obj12_6

	var obj12_7 go_thunder.EthernetRouterAdver
	prefix1 = prefix + "router_adver.0."
	obj12_7.MaxInterval = d.Get(prefix1 + "max_interval").(int)
	obj12_7.DefaultLifetime = d.Get(prefix1 + "default_lifetime").(int)
	obj12_7.ReachableTime = d.Get(prefix1 + "reachable_time").(int)
	obj12_7.OtherConfigAction = d.Get(prefix1 + "other_config_action").(string)
	obj12_7.FloatingIPDefaultVrid = d.Get(prefix1 + "floating_ip_default_vrid").(string)
	obj12_7.ManagedConfigAction = d.Get(prefix1 + "managed_config_action").(string)
	obj12_7.MinInterval = d.Get(prefix1 + "min_interval").(int)
	obj12_7.RateLimit = d.Get(prefix1 + "rate_limit").(int)
	obj12_7.AdverMtuDisable = d.Get(prefix1 + "adver_mtu_disable").(int)

	PrefixListCount := d.Get(prefix1 + "prefix_list.#").(int)
	obj12_7.NotAutonomous = make([]go_thunder.EthernetPrefixList, 0, PrefixListCount)

	for i := 0; i < PrefixListCount; i++ {
		var obj12_7_1 go_thunder.EthernetPrefixList
		prefix2 = prefix1 + fmt.Sprintf("prefix_list.%d.", i)
		obj12_7_1.NotAutonomous = d.Get(prefix2 + "not_autonomous").(int)
		obj12_7_1.ValidLifetime = d.Get(prefix2 + "valid_lifetime").(int)
		obj12_7_1.NotOnLink = d.Get(prefix2 + "not_on_link").(int)
		obj12_7_1.Prefix = d.Get(prefix2 + "prefix").(string)
		obj12_7_1.PreferredLifetime = d.Get(prefix2 + "preferred_lifetime").(int)
		obj12_7.NotAutonomous = append(obj12_7.NotAutonomous, obj12_7_1)
	}

	obj12_7.FloatingIP = d.Get(prefix1 + "floating_ip").(string)
	obj12_7.AdverVrid = d.Get(prefix1 + "adver_vrid").(int)
	obj12_7.UseFloatingIPDefaultVrid = d.Get(prefix1 + "use_floating_ip_default_vrid").(int)
	obj12_7.Action = d.Get(prefix1 + "action").(string)
	obj12_7.AdverVridDefault = d.Get(prefix1 + "adver_vrid_default").(int)
	obj12_7.AdverMtu = d.Get(prefix1 + "adver_mtu").(int)
	obj12_7.RetransmitTimer = d.Get(prefix1 + "retransmit_timer").(int)
	obj12_7.HopLimit = d.Get(prefix1 + "hop_limit").(int)
	obj12_7.UseFloatingIP = d.Get(prefix1 + "use_floating_ip").(int)
	obj12.MaxInterval = obj12_7

	c.Ipv6Enable = obj12

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.EthernetSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj13 go_thunder.EthernetSamplingEnable
		prefix := fmt.Sprintf("sampling_enable.%d.", i)
		obj13.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj13)
	}

	c.LoadInterval = d.Get("load_interval").(int)

	var obj14 go_thunder.EthernetLw4O6
	prefix = "lw_4o6.0."
	obj14.Outside1 = d.Get(prefix + "outside").(int)
	obj14.Inside = d.Get(prefix + "inside").(int)
	c.Outside1 = obj14

	c.Action = d.Get("action").(string)
	c.FecForcedOff = d.Get("fec_forced_off").(int)

	var obj15 go_thunder.EthernetIcmpRateLimit
	prefix = "icmp_rate_limit.0."
	obj15.Lockup = d.Get(prefix + "lockup").(int)
	obj15.LockupPeriod = d.Get(prefix + "lockup_period").(int)
	obj15.Normal = d.Get(prefix + "normal").(int)
	c.Lockup = obj15

	c.FlowControl = d.Get("flow_control").(int)

	vc.UUID = c
	return vc
}
