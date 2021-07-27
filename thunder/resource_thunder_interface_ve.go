package thunder

//Thunder resource InterfaceVE

import (
	"context"
	"fmt"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVE() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceVECreate,
		UpdateContext: resourceInterfaceVEUpdate,
		ReadContext:   resourceInterfaceVERead,
		DeleteContext: resourceInterfaceVEDelete,
		Schema: map[string]*schema.Schema{
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
						"query_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"client": {
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
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"mtu": {
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
			"trap_source": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_tag": {
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
						"inbound": {
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
						"v6_acl_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
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
						"ttl_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
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
			"ifnum": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceInterfaceVECreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceVE (Inside resourceInterfaceVECreate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceVE(d)
		logger.Println("[INFO] received V from method data to InterfaceVE --")
		d.SetId(strconv.Itoa(name))
		err := go_thunder.PostInterfaceVE(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceVERead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceVERead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading InterfaceVE (Inside resourceInterfaceVERead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetInterfaceVE(client.Token, name, client.Host)
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

func resourceInterfaceVEUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Modifying InterfaceVE   (Inside resourceInterfaceVEUpdate) ")
		name := d.Get("ifnum").(int)
		data := dataToInterfaceVE(d)
		logger.Println("[INFO] received V from method data to InterfaceVE ")
		d.SetId(strconv.Itoa(name))
		err := go_thunder.PutInterfaceVE(client.Token, strconv.Itoa(name), data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceVERead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceVEDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceVEDelete) " + name)
		err := go_thunder.DeleteInterfaceVE(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return diags
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToInterfaceVE(d *schema.ResourceData) go_thunder.InterfaceVE {
	var vc go_thunder.InterfaceVE
	var c go_thunder.InterfaceVEInstance
	var obj1 go_thunder.VeMap
	prefix := "map.0."
	obj1.Inside = d.Get(prefix + "inside").(int)
	obj1.MapTInside = d.Get(prefix + "map_t_inside").(int)
	obj1.MapTOutside = d.Get(prefix + "map_t_outside").(int)
	obj1.Outside = d.Get(prefix + "outside").(int)
	c.Inside = obj1

	var obj2 go_thunder.VeNptv6
	prefix = "nptv6.0."

	DomainListCount := d.Get(prefix + "domain_list.#").(int)
	obj2.DomainName = make([]go_thunder.VeDomainList, 0, DomainListCount)

	for i := 0; i < DomainListCount; i++ {
		var obj2_1 go_thunder.VeDomainList
		prefix1 := prefix + fmt.Sprintf("domain_list.%d.", i)
		obj2_1.DomainName = d.Get(prefix1 + "domain_name").(string)
		obj2_1.BindType = d.Get(prefix1 + "bind_type").(string)
		obj2.DomainName = append(obj2.DomainName, obj2_1)
	}

	c.DomainName = obj2

	var obj3 go_thunder.VeIsis
	prefix = "isis.0."

	PriorityListCount := d.Get(prefix + "priority_list.#").(int)
	obj3.Priority = make([]go_thunder.VePriorityList, 0, PriorityListCount)

	for i := 0; i < PriorityListCount; i++ {
		var obj3_1 go_thunder.VePriorityList
		prefix1 := prefix + fmt.Sprintf("priority_list.%d.", i)
		obj3_1.Priority = d.Get(prefix1 + "priority").(int)
		obj3_1.Level = d.Get(prefix1 + "level").(string)
		obj3.Priority = append(obj3.Priority, obj3_1)
	}

	obj3.Padding = d.Get(prefix + "padding").(int)

	HelloIntervalMinimalListCount := d.Get(prefix + "hello_interval_minimal_list.#").(int)
	obj3.HelloIntervalMinimal = make([]go_thunder.VeHelloIntervalMinimalList, 0, HelloIntervalMinimalListCount)

	for i := 0; i < HelloIntervalMinimalListCount; i++ {
		var obj3_2 go_thunder.VeHelloIntervalMinimalList
		prefix1 := prefix + fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj3_2.HelloIntervalMinimal = d.Get(prefix1 + "hello_interval_minimal").(int)
		obj3_2.Level = d.Get(prefix1 + "level").(string)
		obj3.HelloIntervalMinimal = append(obj3.HelloIntervalMinimal, obj3_2)
	}

	var obj3_3 go_thunder.VeMeshGroup
	prefix1 := prefix + "mesh_group.0."
	obj3_3.Value = d.Get(prefix1 + "value").(int)
	obj3_3.Blocked = d.Get(prefix1 + "blocked").(int)
	obj3.Value = obj3_3

	obj3.Network = d.Get(prefix + "network").(string)

	var obj3_4 go_thunder.VeAuthentication
	prefix1 = prefix + "authentication.0."

	SendOnlyListCount := d.Get(prefix1 + "send_only_list.#").(int)
	obj3_4.SendOnly = make([]go_thunder.VeSendOnlyList, 0, SendOnlyListCount)

	for i := 0; i < SendOnlyListCount; i++ {
		var obj3_4_1 go_thunder.VeSendOnlyList
		prefix2 := prefix1 + fmt.Sprintf("send_only_list.%d.", i)
		obj3_4_1.SendOnly = d.Get(prefix2 + "send_only").(int)
		obj3_4_1.Level = d.Get(prefix2 + "level").(string)
		obj3_4.SendOnly = append(obj3_4.SendOnly, obj3_4_1)
	}

	ModeListCount := d.Get(prefix1 + "mode_list.#").(int)
	obj3_4.Mode = make([]go_thunder.VeModeList, 0, ModeListCount)

	for i := 0; i < ModeListCount; i++ {
		var obj3_4_2 go_thunder.VeModeList
		prefix2 := prefix1 + fmt.Sprintf("mode_list.%d.", i)
		obj3_4_2.Mode = d.Get(prefix2 + "mode").(string)
		obj3_4_2.Level = d.Get(prefix2 + "level").(string)
		obj3_4.Mode = append(obj3_4.Mode, obj3_4_2)
	}

	KeyChainListCount := d.Get(prefix1 + "key_chain_list.#").(int)
	obj3_4.KeyChain = make([]go_thunder.VeKeyChainList, 0, KeyChainListCount)

	for i := 0; i < KeyChainListCount; i++ {
		var obj3_4_3 go_thunder.VeKeyChainList
		prefix2 := prefix1 + fmt.Sprintf("key_chain_list.%d.", i)
		obj3_4_3.KeyChain = d.Get(prefix2 + "key_chain").(string)
		obj3_4_3.Level = d.Get(prefix2 + "level").(string)
		obj3_4.KeyChain = append(obj3_4.KeyChain, obj3_4_3)
	}

	obj3.SendOnly = obj3_4

	CsnpIntervalListCount := d.Get(prefix + "csnp_interval_list.#").(int)
	obj3.CsnpInterval = make([]go_thunder.VeCsnpIntervalList, 0, CsnpIntervalListCount)

	for i := 0; i < CsnpIntervalListCount; i++ {
		var obj3_5 go_thunder.VeCsnpIntervalList
		prefix1 = prefix + fmt.Sprintf("csnp_interval_list.%d.", i)
		obj3_5.CsnpInterval = d.Get(prefix1 + "csnp_interval").(int)
		obj3_5.Level = d.Get(prefix1 + "level").(string)
		obj3.CsnpInterval = append(obj3.CsnpInterval, obj3_5)
	}

	obj3.RetransmitInterval = d.Get(prefix + "retransmit_interval").(int)

	PasswordListCount := d.Get(prefix + "password_list.#").(int)
	obj3.Password = make([]go_thunder.VePasswordList, 0, PasswordListCount)

	for i := 0; i < PasswordListCount; i++ {
		var obj3_6 go_thunder.VePasswordList
		prefix1 = prefix + fmt.Sprintf("password_list.%d.", i)
		obj3_6.Password = d.Get(prefix1 + "password").(string)
		obj3_6.Level = d.Get(prefix1 + "level").(string)
		obj3.Password = append(obj3.Password, obj3_6)
	}

	var obj3_7 go_thunder.VeBfdCfg
	prefix1 = prefix + "bfd_cfg.0."
	obj3_7.Disable = d.Get(prefix1 + "disable").(int)
	obj3_7.Bfd = d.Get(prefix1 + "bfd").(int)
	obj3.Disable = obj3_7

	WideMetricListCount := d.Get(prefix + "wide_metric_list.#").(int)
	obj3.WideMetric = make([]go_thunder.VeWideMetricList, 0, WideMetricListCount)

	for i := 0; i < WideMetricListCount; i++ {
		var obj3_8 go_thunder.VeWideMetricList
		prefix1 = prefix + fmt.Sprintf("wide_metric_list.%d.", i)
		obj3_8.WideMetric = d.Get(prefix1 + "wide_metric").(int)
		obj3_8.Level = d.Get(prefix1 + "level").(string)
		obj3.WideMetric = append(obj3.WideMetric, obj3_8)
	}

	HelloIntervalListCount := d.Get(prefix + "hello_interval_list.#").(int)
	obj3.HelloInterval = make([]go_thunder.VeHelloIntervalList, 0, HelloIntervalListCount)

	for i := 0; i < HelloIntervalListCount; i++ {
		var obj3_9 go_thunder.VeHelloIntervalList
		prefix1 = prefix + fmt.Sprintf("hello_interval_list.%d.", i)
		obj3_9.HelloInterval = d.Get(prefix1 + "hello_interval").(int)
		obj3_9.Level = d.Get(prefix1 + "level").(string)
		obj3.HelloInterval = append(obj3.HelloInterval, obj3_9)
	}

	obj3.CircuitType = d.Get(prefix + "circuit_type").(string)

	HelloMultiplierListCount := d.Get(prefix + "hello_multiplier_list.#").(int)
	obj3.HelloMultiplier = make([]go_thunder.VeHelloMultiplierList, 0, HelloMultiplierListCount)

	for i := 0; i < HelloMultiplierListCount; i++ {
		var obj3_10 go_thunder.VeHelloMultiplierList
		prefix1 = prefix + fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj3_10.HelloMultiplier = d.Get(prefix1 + "hello_multiplier").(int)
		obj3_10.Level = d.Get(prefix1 + "level").(string)
		obj3.HelloMultiplier = append(obj3.HelloMultiplier, obj3_10)
	}

	MetricListCount := d.Get(prefix + "metric_list.#").(int)
	obj3.Metric = make([]go_thunder.VeMetricList, 0, MetricListCount)

	for i := 0; i < MetricListCount; i++ {
		var obj3_11 go_thunder.VeMetricList
		prefix1 = prefix + fmt.Sprintf("metric_list.%d.", i)
		obj3_11.Metric = d.Get(prefix1 + "metric").(int)
		obj3_11.Level = d.Get(prefix1 + "level").(string)
		obj3.Metric = append(obj3.Metric, obj3_11)
	}

	obj3.LspInterval = d.Get(prefix + "lsp_interval").(int)
	c.Priority = obj3

	c.Name = d.Get("name").(string)
	c.TrapSource = d.Get("trap_source").(int)

	var obj4 go_thunder.VeBfd
	prefix = "bfd.0."

	var obj4_1 go_thunder.VeIntervalCfg
	prefix1 = prefix + "interval_cfg.0."
	obj4_1.Interval = d.Get(prefix1 + "interval").(int)
	obj4_1.MinRx = d.Get(prefix1 + "min_rx").(int)
	obj4_1.Multiplier = d.Get(prefix1 + "multiplier").(int)
	obj4.Interval = obj4_1

	var obj4_2 go_thunder.VeAuthentication1
	prefix1 = prefix + "authentication.0."
	obj4_2.Encrypted = d.Get(prefix1 + "encrypted").(string)
	obj4_2.Password = d.Get(prefix1 + "password").(string)
	obj4_2.Method = d.Get(prefix1 + "method").(string)
	obj4_2.KeyID = d.Get(prefix1 + "key_id").(int)
	obj4.Encrypted = obj4_2

	obj4.Echo = d.Get(prefix + "echo").(int)
	obj4.Demand = d.Get(prefix + "demand").(int)
	c.Interval = obj4

	var obj5 go_thunder.VeIP1
	prefix = "ip.0."
	obj5.GenerateMembershipQuery = d.Get(prefix + "generate_membership_query").(int)

	AddressListCount := d.Get(prefix + "address_list.#").(int)
	obj5.AddressType = make([]go_thunder.VeAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj5_1 go_thunder.VeAddressList
		prefix1 = prefix + fmt.Sprintf("address_list.%d.", i)
		obj5_1.AddressType = d.Get(prefix1 + "address_type").(string)
		obj5_1.Ipv6Addr = d.Get(prefix1 + "ipv6_addr").(string)
		obj5.AddressType = append(obj5.AddressType, obj5_1)
	}

	obj5.Inside = d.Get(prefix + "inside").(int)
	obj5.AllowPromiscuousVip = d.Get(prefix + "allow_promiscuous_vip").(int)

	HelperAddressListCount := d.Get(prefix + "helper_address_list.#").(int)
	obj5.HelperAddress = make([]go_thunder.VeHelperAddressList, 0, HelperAddressListCount)

	for i := 0; i < HelperAddressListCount; i++ {
		var obj5_2 go_thunder.VeHelperAddressList
		prefix1 = prefix + fmt.Sprintf("helper_address_list.%d.", i)
		obj5_2.HelperAddress = d.Get(prefix1 + "helper_address").(string)
		obj5.HelperAddress = append(obj5.HelperAddress, obj5_2)
	}

	obj5.MaxRespTime = d.Get(prefix + "max_resp_time").(int)
	obj5.QueryInterval = d.Get(prefix + "query_interval").(int)
	obj5.Outside = d.Get(prefix + "outside").(int)
	obj5.Client = d.Get(prefix + "client").(int)

	var obj5_3 go_thunder.VeStatefulFirewall
	prefix1 = prefix + "stateful_firewall.0."
	obj5_3.ClassList = d.Get(prefix1 + "class_list").(string)
	obj5_3.Inside = d.Get(prefix1 + "inside").(int)
	obj5_3.Outside = d.Get(prefix1 + "outside").(int)
	obj5_3.ACLID = d.Get(prefix1 + "acl_id").(int)
	obj5_3.AccessList = d.Get(prefix1 + "access_list").(int)
	obj5.ClassList = obj5_3

	var obj5_4 go_thunder.VeRip
	prefix1 = prefix + "rip.0."

	var obj5_4_1 go_thunder.VeReceiveCfg
	prefix2 := prefix1 + "receive_cfg.0."
	obj5_4_1.Receive = d.Get(prefix2 + "receive").(int)
	obj5_4_1.Version = d.Get(prefix2 + "version").(string)
	obj5_4.Receive = obj5_4_1

	obj5_4.ReceivePacket = d.Get(prefix1 + "receive_packet").(int)

	var obj5_4_2 go_thunder.VeSplitHorizonCfg
	prefix2 = prefix1 + "split_horizon_cfg.0."
	obj5_4_2.State = d.Get(prefix2 + "state").(string)
	obj5_4.State = obj5_4_2

	var obj5_4_3 go_thunder.VeAuthentication2
	prefix2 = prefix1 + "authentication.0."

	var obj5_4_3_1 go_thunder.VeKeyChain
	prefix3 := prefix2 + "key_chain.0."
	obj5_4_3_1.KeyChain = d.Get(prefix3 + "key_chain").(string)
	obj5_4_3.KeyChain = obj5_4_3_1

	var obj5_4_3_2 go_thunder.VeMode
	prefix3 = prefix2 + "mode.0."
	obj5_4_3_2.Mode = d.Get(prefix3 + "mode").(string)
	obj5_4_3.Mode = obj5_4_3_2

	var obj5_4_3_3 go_thunder.VeStr
	prefix3 = prefix2 + "str.0."
	obj5_4_3_3.String = d.Get(prefix3 + "string").(string)
	obj5_4_3.String = obj5_4_3_3

	obj5_4.KeyChain = obj5_4_3

	var obj5_4_4 go_thunder.VeSendCfg
	prefix2 = prefix1 + "send_cfg.0."
	obj5_4_4.Version = d.Get(prefix2 + "version").(string)
	obj5_4_4.Send = d.Get(prefix2 + "send").(int)
	obj5_4.Version = obj5_4_4

	obj5_4.SendPacket = d.Get(prefix1 + "send_packet").(int)
	obj5.Receive = obj5_4

	obj5.TTLIgnore = d.Get(prefix + "ttl_ignore").(int)

	var obj5_5 go_thunder.VeRouter
	prefix1 = prefix + "router.0."

	var obj5_5_1 go_thunder.VeIsis1
	prefix2 = prefix1 + "isis.0."
	obj5_5_1.Tag = d.Get(prefix2 + "tag").(string)
	obj5_5.Tag = obj5_5_1

	obj5.Tag = obj5_5

	obj5.Dhcp = d.Get(prefix + "dhcp").(int)
	obj5.Server = d.Get(prefix + "server").(int)

	var obj5_6 go_thunder.VeOspf
	prefix1 = prefix + "ospf.0."

	OspfIpListCount := d.Get(prefix1 + "ospf_ip_list.#").(int)
	obj5_6.DeadInterval = make([]go_thunder.VeOspfIPList, 0, OspfIpListCount)

	for i := 0; i < OspfIpListCount; i++ {
		var obj5_6_1 go_thunder.VeOspfIPList
		prefix2 = prefix1 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj5_6_1.DeadInterval = d.Get(prefix2 + "dead_interval").(int)
		obj5_6_1.AuthenticationKey = d.Get(prefix2 + "authentication_key").(string)
		obj5_6_1.MtuIgnore = d.Get(prefix2 + "mtu_ignore").(int)
		obj5_6_1.TransmitDelay = d.Get(prefix2 + "transmit_delay").(int)
		obj5_6_1.Value = d.Get(prefix2 + "value").(string)
		obj5_6_1.Priority = d.Get(prefix2 + "priority").(int)
		obj5_6_1.Authentication = d.Get(prefix2 + "authentication").(int)
		obj5_6_1.Cost = d.Get(prefix2 + "cost").(int)
		obj5_6_1.DatabaseFilter = d.Get(prefix2 + "database_filter").(string)
		obj5_6_1.HelloInterval = d.Get(prefix2 + "hello_interval").(int)
		obj5_6_1.IPAddr = d.Get(prefix2 + "ip_addr").(string)
		obj5_6_1.RetransmitInterval = d.Get(prefix2 + "retransmit_interval").(int)

		MessageDigestCfgCount := d.Get(prefix2 + "message_digest_cfg.#").(int)
		obj5_6_1.MessageDigestKey = make([]go_thunder.VeMessageDigestCfg, 0, MessageDigestCfgCount)

		for i := 0; i < MessageDigestCfgCount; i++ {
			var obj5_6_1_1 go_thunder.VeMessageDigestCfg
			prefix3 = prefix2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj5_6_1_1.MessageDigestKey = d.Get(prefix3 + "message_digest_key").(int)

			var obj5_6_1_1_1 go_thunder.VeMd5
			prefix4 := prefix3 + "md5.0."
			obj5_6_1_1_1.Md5Value = d.Get(prefix4 + "md5_value").(string)
			obj5_6_1_1_1.Encrypted = d.Get(prefix4 + "encrypted").(string)
			obj5_6_1_1.Md5Value = obj5_6_1_1_1

			obj5_6_1.MessageDigestKey = append(obj5_6_1.MessageDigestKey, obj5_6_1_1)
		}

		obj5_6_1.Out = d.Get(prefix2 + "out").(int)
		obj5_6.DeadInterval = append(obj5_6.DeadInterval, obj5_6_1)
	}

	var obj5_6_2 go_thunder.VeOspfGlobal
	prefix2 = prefix1 + "ospf_global.0."
	obj5_6_2.Cost = d.Get(prefix2 + "cost").(int)
	obj5_6_2.DeadInterval = d.Get(prefix2 + "dead_interval").(int)
	obj5_6_2.AuthenticationKey = d.Get(prefix2 + "authentication_key").(string)

	var obj5_6_2_1 go_thunder.VeNetwork
	prefix3 = prefix2 + "network.0."
	obj5_6_2_1.Broadcast = d.Get(prefix3 + "broadcast").(int)
	obj5_6_2_1.PointToMultipoint = d.Get(prefix3 + "point_to_multipoint").(int)
	obj5_6_2_1.NonBroadcast = d.Get(prefix3 + "non_broadcast").(int)
	obj5_6_2_1.PointToPoint = d.Get(prefix3 + "point_to_point").(int)
	obj5_6_2_1.P2MpNbma = d.Get(prefix3 + "p2mp_nbma").(int)
	obj5_6_2.Broadcast = obj5_6_2_1

	obj5_6_2.MtuIgnore = d.Get(prefix2 + "mtu_ignore").(int)
	obj5_6_2.TransmitDelay = d.Get(prefix2 + "transmit_delay").(int)

	var obj5_6_2_2 go_thunder.VeAuthenticationCfg
	prefix3 = prefix2 + "authentication_cfg.0."
	obj5_6_2_2.Authentication = d.Get(prefix3 + "authentication").(int)
	obj5_6_2_2.Value = d.Get(prefix3 + "value").(string)
	obj5_6_2.Authentication = obj5_6_2_2

	obj5_6_2.RetransmitInterval = d.Get(prefix2 + "retransmit_interval").(int)

	var obj5_6_2_3 go_thunder.VeBfdCfg
	prefix3 = prefix2 + "bfd_cfg.0."
	obj5_6_2_3.Disable = d.Get(prefix3 + "disable").(int)
	obj5_6_2_3.Bfd = d.Get(prefix3 + "bfd").(int)
	obj5_6_2.Bfd = obj5_6_2_3

	obj5_6_2.Disable = d.Get(prefix2 + "disable").(string)
	obj5_6_2.HelloInterval = d.Get(prefix2 + "hello_interval").(int)

	var obj5_6_2_4 go_thunder.VeDatabaseFilterCfg
	prefix3 = prefix2 + "database_filter_cfg.0."
	obj5_6_2_4.DatabaseFilter = d.Get(prefix3 + "database_filter").(string)
	obj5_6_2_4.Out = d.Get(prefix3 + "out").(int)
	obj5_6_2.DatabaseFilter = obj5_6_2_4

	obj5_6_2.Priority = d.Get(prefix2 + "priority").(int)
	obj5_6_2.Mtu = d.Get(prefix2 + "mtu").(int)

	MessageDigestCfgCount := d.Get(prefix2 + "message_digest_cfg.#").(int)
	obj5_6_2.MessageDigestKey = make([]go_thunder.VeMessageDigestCfg, 0, MessageDigestCfgCount)

	for i := 0; i < MessageDigestCfgCount; i++ {
		var obj5_6_2_5 go_thunder.VeMessageDigestCfg
		prefix3 = prefix2 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj5_6_2_5.MessageDigestKey = d.Get(prefix3 + "message_digest_key").(int)

		var obj5_6_2_5_1 go_thunder.VeMd5
		prefix4 := prefix3 + "md5.0."
		obj5_6_2_5_1.Md5Value = d.Get(prefix4 + "md5_value").(string)
		obj5_6_2_5_1.Encrypted = d.Get(prefix4 + "encrypted").(string)
		obj5_6_2_5.Md5Value = obj5_6_2_5_1

		obj5_6_2.MessageDigestKey = append(obj5_6_2.MessageDigestKey, obj5_6_2_5)
	}

	obj5_6.Cost = obj5_6_2

	obj5.DeadInterval = obj5_6

	obj5.SlbPartitionRedirect = d.Get(prefix + "slb_partition_redirect").(int)
	c.GenerateMembershipQuery = obj5

	var obj6 go_thunder.VeIcmpv6RateLimit
	prefix = "icmpv6_rate_limit.0."
	obj6.LockupPeriodV6 = d.Get(prefix + "lockup_period_v6").(int)
	obj6.NormalV6 = d.Get(prefix + "normal_v6").(int)
	obj6.LockupV6 = d.Get(prefix + "lockup_v6").(int)
	c.LockupPeriodV6 = obj6

	c.UserTag = d.Get("user_tag").(string)
	c.Mtu = d.Get("mtu").(int)
	c.Action = d.Get("action").(string)
	c.Ifnum = d.Get("ifnum").(int)

	SamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.Counters1 = make([]go_thunder.VeSamplingEnable, 0, SamplingEnableCount)

	for i := 0; i < SamplingEnableCount; i++ {
		var obj7 go_thunder.VeSamplingEnable
		prefix = fmt.Sprintf("sampling_enable.%d.", i)
		obj7.Counters1 = d.Get(prefix + "counters1").(string)
		c.Counters1 = append(c.Counters1, obj7)
	}

	var obj8 go_thunder.VeLw4O6
	prefix = "lw_4o6.0."
	obj8.Outside = d.Get(prefix + "outside").(int)
	obj8.Inside = d.Get(prefix + "inside").(int)
	c.Outside = obj8

	var obj9 go_thunder.VeIpv61
	prefix = "ipv6.0."
	obj9.Inbound = d.Get(prefix + "inbound").(int)

	AddressListCount = d.Get(prefix + "address_list.#").(int)
	obj9.AddressType = make([]go_thunder.VeAddressList, 0, AddressListCount)

	for i := 0; i < AddressListCount; i++ {
		var obj9_1 go_thunder.VeAddressList
		prefix1 = prefix + fmt.Sprintf("address_list.%d.", i)
		obj9_1.AddressType = d.Get(prefix1 + "address_type").(string)
		obj9_1.Ipv6Addr = d.Get(prefix1 + "ipv6_addr").(string)
		obj9.AddressType = append(obj9.AddressType, obj9_1)
	}

	obj9.Inside = d.Get(prefix + "inside").(int)
	obj9.Ipv6Enable = d.Get(prefix + "ipv6_enable").(int)

	var obj9_2 go_thunder.VeRip1
	prefix1 = prefix + "rip.0."

	var obj9_2_1 go_thunder.VeSplitHorizonCfg
	prefix2 = prefix1 + "split_horizon_cfg.0."
	obj9_2_1.State = d.Get(prefix2 + "state").(string)
	obj9_2.State = obj9_2_1

	obj9.State = obj9_2

	obj9.Outside = d.Get(prefix + "outside").(int)

	var obj9_3 go_thunder.VeStatefulFirewall1
	prefix1 = prefix + "stateful_firewall.0."
	obj9_3.ClassList = d.Get(prefix1 + "class_list").(string)
	obj9_3.ACLName = d.Get(prefix1 + "acl_name").(string)
	obj9_3.Inside = d.Get(prefix1 + "inside").(int)
	obj9_3.Outside = d.Get(prefix1 + "outside").(int)
	obj9_3.AccessList = d.Get(prefix1 + "access_list").(int)
	obj9.ClassList = obj9_3

	obj9.V6ACLName = d.Get(prefix + "v6_acl_name").(string)
	obj9.TTLIgnore = d.Get(prefix + "ttl_ignore").(int)

	var obj9_4 go_thunder.VeRouter1
	prefix1 = prefix + "router.0."

	var obj9_4_1 go_thunder.VeRipng
	prefix2 = prefix1 + "ripng.0."
	obj9_4_1.Rip = d.Get(prefix2 + "rip").(int)
	obj9_4.Rip = obj9_4_1

	var obj9_4_2 go_thunder.VeOspf1
	prefix2 = prefix1 + "ospf.0."

	AreaListCount := d.Get(prefix2 + "area_list.#").(int)
	obj9_4_2.AreaIDAddr = make([]go_thunder.VeAreaList, 0, AreaListCount)

	for i := 0; i < AreaListCount; i++ {
		var obj9_4_2_1 go_thunder.VeAreaList
		prefix3 = prefix2 + fmt.Sprintf("area_list.%d.", i)
		obj9_4_2_1.AreaIDAddr = d.Get(prefix3 + "area_id_addr").(string)
		obj9_4_2_1.Tag = d.Get(prefix3 + "tag").(string)
		obj9_4_2_1.InstanceID = d.Get(prefix3 + "instance_id").(int)
		obj9_4_2_1.AreaIDNum = d.Get(prefix3 + "area_id_num").(int)
		obj9_4_2.AreaIDAddr = append(obj9_4_2.AreaIDAddr, obj9_4_2_1)
	}

	obj9_4.AreaIDAddr = obj9_4_2

	var obj9_4_3 go_thunder.VeIsis1
	prefix2 = prefix1 + "isis.0."
	obj9_4_3.Tag = d.Get(prefix2 + "tag").(string)
	obj9_4.Tag = obj9_4_3

	obj9.Rip = obj9_4

	var obj9_5 go_thunder.VeOspf2
	prefix1 = prefix + "ospf.0."
	obj9_5.Bfd = d.Get(prefix1 + "bfd").(int)

	CostCfgCount := d.Get(prefix1 + "cost_cfg.#").(int)
	obj9_5.Cost = make([]go_thunder.VeCostCfg, 0, CostCfgCount)

	for i := 0; i < CostCfgCount; i++ {
		var obj9_5_1 go_thunder.VeCostCfg
		prefix2 = prefix1 + fmt.Sprintf("cost_cfg.%d.", i)
		obj9_5_1.Cost = d.Get(prefix2 + "cost").(int)
		obj9_5_1.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.Cost = append(obj9_5.Cost, obj9_5_1)
	}

	PriorityCfgCount := d.Get(prefix1 + "priority_cfg.#").(int)
	obj9_5.Priority = make([]go_thunder.VePriorityCfg, 0, PriorityCfgCount)

	for i := 0; i < PriorityCfgCount; i++ {
		var obj9_5_2 go_thunder.VePriorityCfg
		prefix2 = prefix1 + fmt.Sprintf("priority_cfg.%d.", i)
		obj9_5_2.Priority = d.Get(prefix2 + "priority").(int)
		obj9_5_2.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.Priority = append(obj9_5.Priority, obj9_5_2)
	}

	HelloIntervalCfgCount := d.Get(prefix1 + "hello_interval_cfg.#").(int)
	obj9_5.HelloInterval = make([]go_thunder.VeHelloIntervalCfg, 0, HelloIntervalCfgCount)

	for i := 0; i < HelloIntervalCfgCount; i++ {
		var obj9_5_3 go_thunder.VeHelloIntervalCfg
		prefix2 = prefix1 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj9_5_3.HelloInterval = d.Get(prefix2 + "hello_interval").(int)
		obj9_5_3.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.HelloInterval = append(obj9_5.HelloInterval, obj9_5_3)
	}

	MtuIgnoreCfgCount := d.Get(prefix1 + "mtu_ignore_cfg.#").(int)
	obj9_5.MtuIgnore = make([]go_thunder.VeMtuIgnoreCfg, 0, MtuIgnoreCfgCount)

	for i := 0; i < MtuIgnoreCfgCount; i++ {
		var obj9_5_4 go_thunder.VeMtuIgnoreCfg
		prefix2 = prefix1 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj9_5_4.MtuIgnore = d.Get(prefix2 + "mtu_ignore").(int)
		obj9_5_4.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.MtuIgnore = append(obj9_5.MtuIgnore, obj9_5_4)
	}

	RetransmitIntervalCfgCount := d.Get(prefix1 + "retransmit_interval_cfg.#").(int)
	obj9_5.RetransmitInterval = make([]go_thunder.VeRetransmitIntervalCfg, 0, RetransmitIntervalCfgCount)

	for i := 0; i < RetransmitIntervalCfgCount; i++ {
		var obj9_5_5 go_thunder.VeRetransmitIntervalCfg
		prefix2 = prefix1 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj9_5_5.RetransmitInterval = d.Get(prefix2 + "retransmit_interval").(int)
		obj9_5_5.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.RetransmitInterval = append(obj9_5.RetransmitInterval, obj9_5_5)
	}

	obj9_5.Disable = d.Get(prefix1 + "disable").(int)

	TransmitDelayCfgCount := d.Get(prefix1 + "transmit_delay_cfg.#").(int)
	obj9_5.TransmitDelay = make([]go_thunder.VeTransmitDelayCfg, 0, TransmitDelayCfgCount)

	for i := 0; i < TransmitDelayCfgCount; i++ {
		var obj9_5_6 go_thunder.VeTransmitDelayCfg
		prefix2 = prefix1 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj9_5_6.TransmitDelay = d.Get(prefix2 + "transmit_delay").(int)
		obj9_5_6.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.TransmitDelay = append(obj9_5.TransmitDelay, obj9_5_6)
	}

	NeighborCfgCount := d.Get(prefix1 + "neighbor_cfg.#").(int)
	obj9_5.NeighborPriority = make([]go_thunder.VeNeighborCfg, 0, NeighborCfgCount)

	for i := 0; i < NeighborCfgCount; i++ {
		var obj9_5_7 go_thunder.VeNeighborCfg
		prefix2 = prefix1 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj9_5_7.NeighborPriority = d.Get(prefix2 + "neighbor_priority").(int)
		obj9_5_7.NeigInst = d.Get(prefix2 + "neig_inst").(int)
		obj9_5_7.NeighborPollInterval = d.Get(prefix2 + "neighbor_poll_interval").(int)
		obj9_5_7.NeighborCost = d.Get(prefix2 + "neighbor_cost").(int)
		obj9_5_7.Neighbor = d.Get(prefix2 + "neighbor").(string)
		obj9_5.NeighborPriority = append(obj9_5.NeighborPriority, obj9_5_7)
	}

	NetworkListCount := d.Get(prefix1 + "network_list.#").(int)
	obj9_5.BroadcastType = make([]go_thunder.VeNetworkList, 0, NetworkListCount)

	for i := 0; i < NetworkListCount; i++ {
		var obj9_5_8 go_thunder.VeNetworkList
		prefix2 = prefix1 + fmt.Sprintf("network_list.%d.", i)
		obj9_5_8.BroadcastType = d.Get(prefix2 + "broadcast_type").(string)
		obj9_5_8.P2MpNbma = d.Get(prefix2 + "p2mp_nbma").(int)
		obj9_5_8.NetworkInstanceID = d.Get(prefix2 + "network_instance_id").(int)
		obj9_5.BroadcastType = append(obj9_5.BroadcastType, obj9_5_8)
	}

	DeadIntervalCfgCount := d.Get(prefix1 + "dead_interval_cfg.#").(int)
	obj9_5.DeadInterval = make([]go_thunder.VeDeadIntervalCfg, 0, DeadIntervalCfgCount)

	for i := 0; i < DeadIntervalCfgCount; i++ {
		var obj9_5_9 go_thunder.VeDeadIntervalCfg
		prefix2 = prefix1 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj9_5_9.DeadInterval = d.Get(prefix2 + "dead_interval").(int)
		obj9_5_9.InstanceID = d.Get(prefix2 + "instance_id").(int)
		obj9_5.DeadInterval = append(obj9_5.DeadInterval, obj9_5_9)
	}

	obj9.Bfd = obj9_5

	var obj9_6 go_thunder.VeRouterAdver
	prefix1 = prefix + "router_adver.0."
	obj9_6.MaxInterval = d.Get(prefix1 + "max_interval").(int)
	obj9_6.DefaultLifetime = d.Get(prefix1 + "default_lifetime").(int)
	obj9_6.ReachableTime = d.Get(prefix1 + "reachable_time").(int)
	obj9_6.OtherConfigAction = d.Get(prefix1 + "other_config_action").(string)
	obj9_6.FloatingIPDefaultVrid = d.Get(prefix1 + "floating_ip_default_vrid").(string)
	obj9_6.ManagedConfigAction = d.Get(prefix1 + "managed_config_action").(string)
	obj9_6.MinInterval = d.Get(prefix1 + "min_interval").(int)
	obj9_6.RateLimit = d.Get(prefix1 + "rate_limit").(int)
	obj9_6.AdverMtuDisable = d.Get(prefix1 + "adver_mtu_disable").(int)

	PrefixListCount := d.Get(prefix1 + "prefix_list.#").(int)
	obj9_6.NotAutonomous = make([]go_thunder.VePrefixList, 0, PrefixListCount)

	for i := 0; i < PrefixListCount; i++ {
		var obj9_6_1 go_thunder.VePrefixList
		prefix2 = prefix1 + fmt.Sprintf("prefix_list.%d.", i)
		obj9_6_1.NotAutonomous = d.Get(prefix2 + "not_autonomous").(int)
		obj9_6_1.ValidLifetime = d.Get(prefix2 + "valid_lifetime").(int)
		obj9_6_1.NotOnLink = d.Get(prefix2 + "not_on_link").(int)
		obj9_6_1.Prefix = d.Get(prefix2 + "prefix").(string)
		obj9_6_1.PreferredLifetime = d.Get(prefix2 + "preferred_lifetime").(int)
		obj9_6.NotAutonomous = append(obj9_6.NotAutonomous, obj9_6_1)
	}

	obj9_6.FloatingIP = d.Get(prefix1 + "floating_ip").(string)
	obj9_6.AdverVrid = d.Get(prefix1 + "adver_vrid").(int)
	obj9_6.UseFloatingIPDefaultVrid = d.Get(prefix1 + "use_floating_ip_default_vrid").(int)
	obj9_6.Action = d.Get(prefix1 + "action").(string)
	obj9_6.AdverVridDefault = d.Get(prefix1 + "adver_vrid_default").(int)
	obj9_6.AdverMtu = d.Get(prefix1 + "adver_mtu").(int)
	obj9_6.RetransmitTimer = d.Get(prefix1 + "retransmit_timer").(int)
	obj9_6.HopLimit = d.Get(prefix1 + "hop_limit").(int)
	obj9_6.UseFloatingIP = d.Get(prefix1 + "use_floating_ip").(int)
	obj9.MaxInterval = obj9_6

	c.Inbound = obj9

	var obj10 go_thunder.VeAccessList
	prefix = "access_list.0."
	obj10.ACLName = d.Get(prefix + "acl_name").(string)
	obj10.ACLID = d.Get(prefix + "acl_id").(int)
	c.ACLName = obj10

	c.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)

	var obj11 go_thunder.VeIcmpRateLimit
	prefix = "icmp_rate_limit.0."
	obj11.Lockup = d.Get(prefix + "lockup").(int)
	obj11.LockupPeriod = d.Get(prefix + "lockup_period").(int)
	obj11.Normal = d.Get(prefix + "normal").(int)
	c.Lockup = obj11

	var obj12 go_thunder.VeDdos
	prefix = "ddos.0."
	obj12.Outside1 = d.Get(prefix + "outside").(int)
	obj12.Inside = d.Get(prefix + "inside").(int)
	c.Outside1 = obj12

	vc.UUID = c
	return vc
}
