package thunder

//Thunder resource InterfaceLif

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceInterfaceLif() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceLifCreate,
		UpdateContext: resourceInterfaceLifUpdate,
		ReadContext:   resourceInterfaceLifRead,
		DeleteContext: resourceInterfaceLifDelete,
		Schema: map[string]*schema.Schema{
			"ifname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"name": {
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
									"anycast": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"link_local": {
										Type:        schema.TypeInt,
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
									"anycast": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"link_local": {
										Type:        schema.TypeInt,
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
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
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

func resourceInterfaceLifCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating InterfaceLif (Inside resourceInterfaceLifCreate) ")
		name1 := d.Get("ifname").(string)
		data := dataToInterfaceLif(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceLif --")
		d.SetId(name1)
		err := go_thunder.PostInterfaceLif(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceLifRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceLifRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading InterfaceLif (Inside resourceInterfaceLifRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetInterfaceLif(client.Token, name1, client.Host)
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

func resourceInterfaceLifUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying InterfaceLif   (Inside resourceInterfaceLifUpdate) ")
		data := dataToInterfaceLif(d)
		logger.Println("[INFO] received formatted data from method data to InterfaceLif ")
		err := go_thunder.PutInterfaceLif(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceInterfaceLifRead(ctx, d, meta)

	}
	return diags
}

func resourceInterfaceLifDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceInterfaceLifDelete) " + name1)
		err := go_thunder.DeleteInterfaceLif(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diags
		}
		return nil
	}
	return diags
}

func dataToInterfaceLif(d *schema.ResourceData) go_thunder.InterfaceLif {
	var vc go_thunder.InterfaceLif
	var c go_thunder.InterfaceLifInstance
	c.InterfaceLifInstanceIfname = d.Get("ifname").(string)
	c.InterfaceLifInstanceMtu = d.Get("mtu").(int)
	c.InterfaceLifInstanceAction = d.Get("action").(string)
	c.InterfaceLifInstanceName = d.Get("name").(string)

	var obj1 go_thunder.InterfaceLifInstanceAccessList
	prefix1 := "access_list.0."
	obj1.InterfaceLifInstanceAccessListAclID = d.Get(prefix1 + "acl_id").(int)
	obj1.InterfaceLifInstanceAccessListAclName = d.Get(prefix1 + "acl_name").(string)

	c.InterfaceLifInstanceAccessListAclID = obj1

	c.InterfaceLifInstanceUserTag = d.Get("user_tag").(string)

	InterfaceLifInstanceSamplingEnableCount := d.Get("sampling_enable.#").(int)
	c.InterfaceLifInstanceSamplingEnableCounters1 = make([]go_thunder.InterfaceLifInstanceSamplingEnable, 0, InterfaceLifInstanceSamplingEnableCount)

	for i := 0; i < InterfaceLifInstanceSamplingEnableCount; i++ {
		var obj2 go_thunder.InterfaceLifInstanceSamplingEnable
		prefix2 := fmt.Sprintf("sampling_enable.%d.", i)
		obj2.InterfaceLifInstanceSamplingEnableCounters1 = d.Get(prefix2 + "counters1").(string)
		c.InterfaceLifInstanceSamplingEnableCounters1 = append(c.InterfaceLifInstanceSamplingEnableCounters1, obj2)
	}

	var obj3 go_thunder.InterfaceLifInstanceIP
	prefix3 := "ip.0."
	obj3.InterfaceLifInstanceIPDhcp = d.Get(prefix3 + "dhcp").(int)

	InterfaceLifInstanceIPAddressListCount := d.Get(prefix3 + "address_list.#").(int)
	obj3.InterfaceLifInstanceIPAddressListIpv6Addr = make([]go_thunder.InterfaceLifInstanceIPAddressList, 0, InterfaceLifInstanceIPAddressListCount)

	for i := 0; i < InterfaceLifInstanceIPAddressListCount; i++ {
		var obj3_1 go_thunder.InterfaceLifInstanceIPAddressList
		prefix3_1 := prefix3 + fmt.Sprintf("address_list.%d.", i)
		obj3_1.InterfaceLifInstanceIPAddressListIpv6Addr = d.Get(prefix3_1 + "ipv6_addr").(string)
		obj3_1.InterfaceLifInstanceIPAddressListAnycast = d.Get(prefix3_1 + "anycast").(int)
		obj3_1.InterfaceLifInstanceIPAddressListLinkLocal = d.Get(prefix3_1 + "link_local").(int)
		obj3.InterfaceLifInstanceIPAddressListIpv6Addr = append(obj3.InterfaceLifInstanceIPAddressListIpv6Addr, obj3_1)
	}

	obj3.InterfaceLifInstanceIPAllowPromiscuousVip = d.Get(prefix3 + "allow_promiscuous_vip").(int)
	obj3.InterfaceLifInstanceIPCacheSpoofingPort = d.Get(prefix3 + "cache_spoofing_port").(int)
	obj3.InterfaceLifInstanceIPInside = d.Get(prefix3 + "inside").(int)
	obj3.InterfaceLifInstanceIPOutside = d.Get(prefix3 + "outside").(int)
	obj3.InterfaceLifInstanceIPGenerateMembershipQuery = d.Get(prefix3 + "generate_membership_query").(int)
	obj3.InterfaceLifInstanceIPQueryInterval = d.Get(prefix3 + "query_interval").(int)
	obj3.InterfaceLifInstanceIPMaxRespTime = d.Get(prefix3 + "max_resp_time").(int)

	var obj3_2 go_thunder.InterfaceLifInstanceIPRouter
	prefix3_2 := prefix3 + "router.0."

	var obj3_2_1 go_thunder.InterfaceLifInstanceIPRouterIsis
	prefix3_2_1 := prefix3_2 + "isis.0."
	obj3_2_1.InterfaceLifInstanceIPRouterIsisTag = d.Get(prefix3_2_1 + "tag").(string)

	obj3_2.InterfaceLifInstanceIPRouterIsisTag = obj3_2_1

	obj3.InterfaceLifInstanceIPRouterIsis = obj3_2

	var obj3_3 go_thunder.InterfaceLifInstanceIPRip
	prefix3_3 := prefix3 + "rip.0."

	var obj3_3_1 go_thunder.InterfaceLifInstanceIPRipAuthentication
	prefix3_3_1 := prefix3_3 + "authentication.0."

	var obj3_3_1_1 go_thunder.InterfaceLifInstanceIPRipAuthenticationStr
	prefix3_3_1_1 := prefix3_3_1 + "str.0."
	obj3_3_1_1.InterfaceLifInstanceIPRipAuthenticationStrString = d.Get(prefix3_3_1_1 + "string").(string)

	obj3_3_1.InterfaceLifInstanceIPRipAuthenticationStrString = obj3_3_1_1

	var obj3_3_1_2 go_thunder.InterfaceLifInstanceIPRipAuthenticationMode
	prefix3_3_1_2 := prefix3_3_1 + "mode.0."
	obj3_3_1_2.InterfaceLifInstanceIPRipAuthenticationModeMode = d.Get(prefix3_3_1_2 + "mode").(string)

	obj3_3_1.InterfaceLifInstanceIPRipAuthenticationModeMode = obj3_3_1_2

	var obj3_3_1_3 go_thunder.InterfaceLifInstanceIPRipAuthenticationKeyChain
	prefix3_3_1_3 := prefix3_3_1 + "key_chain.0."
	obj3_3_1_3.InterfaceLifInstanceIPRipAuthenticationKeyChainKeyChain = d.Get(prefix3_3_1_3 + "key_chain").(string)

	obj3_3_1.InterfaceLifInstanceIPRipAuthenticationKeyChainKeyChain = obj3_3_1_3

	obj3_3.InterfaceLifInstanceIPRipAuthenticationStr = obj3_3_1

	obj3_3.InterfaceLifInstanceIPRipSendPacket = d.Get(prefix3_3 + "send_packet").(int)
	obj3_3.InterfaceLifInstanceIPRipReceivePacket = d.Get(prefix3_3 + "receive_packet").(int)

	var obj3_3_2 go_thunder.InterfaceLifInstanceIPRipSendCfg
	prefix3_3_2 := prefix3_3 + "send_cfg.0."
	obj3_3_2.InterfaceLifInstanceIPRipSendCfgSend = d.Get(prefix3_3_2 + "send").(int)
	obj3_3_2.InterfaceLifInstanceIPRipSendCfgVersion = d.Get(prefix3_3_2 + "version").(string)

	obj3_3.InterfaceLifInstanceIPRipSendCfgSend = obj3_3_2

	var obj3_3_3 go_thunder.InterfaceLifInstanceIPRipReceiveCfg
	prefix3_3_3 := prefix3_3 + "receive_cfg.0."
	obj3_3_3.InterfaceLifInstanceIPRipReceiveCfgReceive = d.Get(prefix3_3_3 + "receive").(int)
	obj3_3_3.InterfaceLifInstanceIPRipReceiveCfgVersion = d.Get(prefix3_3_3 + "version").(string)

	obj3_3.InterfaceLifInstanceIPRipReceiveCfgReceive = obj3_3_3

	var obj3_3_4 go_thunder.InterfaceLifInstanceIPRipSplitHorizonCfg
	prefix3_3_4 := prefix3_3 + "split_horizon_cfg.0."
	obj3_3_4.InterfaceLifInstanceIPRipSplitHorizonCfgState = d.Get(prefix3_3_4 + "state").(string)

	obj3_3.InterfaceLifInstanceIPRipSplitHorizonCfgState = obj3_3_4

	obj3.InterfaceLifInstanceIPRipAuthentication = obj3_3

	var obj3_4 go_thunder.InterfaceLifInstanceIPOspf
	prefix3_4 := prefix3 + "ospf.0."

	var obj3_4_1 go_thunder.InterfaceLifInstanceIPOspfOspfGlobal
	prefix3_4_1 := prefix3_4 + "ospf_global.0."

	var obj3_4_1_1 go_thunder.InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfg
	prefix3_4_1_1 := prefix3_4_1 + "authentication_cfg.0."
	obj3_4_1_1.InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication = d.Get(prefix3_4_1_1 + "authentication").(int)
	obj3_4_1_1.InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfgValue = d.Get(prefix3_4_1_1 + "value").(string)

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfgAuthentication = obj3_4_1_1

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalAuthenticationKey = d.Get(prefix3_4_1 + "authentication_key").(string)

	var obj3_4_1_2 go_thunder.InterfaceLifInstanceIPOspfOspfGlobalBfdCfg
	prefix3_4_1_2 := prefix3_4_1 + "bfd_cfg.0."
	obj3_4_1_2.InterfaceLifInstanceIPOspfOspfGlobalBfdCfgBfd = d.Get(prefix3_4_1_2 + "bfd").(int)
	obj3_4_1_2.InterfaceLifInstanceIPOspfOspfGlobalBfdCfgDisable = d.Get(prefix3_4_1_2 + "disable").(int)

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalBfdCfgBfd = obj3_4_1_2

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalCost = d.Get(prefix3_4_1 + "cost").(int)

	var obj3_4_1_3 go_thunder.InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfg
	prefix3_4_1_3 := prefix3_4_1 + "database_filter_cfg.0."
	obj3_4_1_3.InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = d.Get(prefix3_4_1_3 + "database_filter").(string)
	obj3_4_1_3.InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfgOut = d.Get(prefix3_4_1_3 + "out").(int)

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalDatabaseFilterCfgDatabaseFilter = obj3_4_1_3

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalDeadInterval = d.Get(prefix3_4_1 + "dead_interval").(int)
	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalDisable = d.Get(prefix3_4_1 + "disable").(string)
	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalHelloInterval = d.Get(prefix3_4_1 + "hello_interval").(int)

	InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgCount := d.Get(prefix3_4_1 + "message_digest_cfg.#").(int)
	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfg, 0, InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgCount)

	for i := 0; i < InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgCount; i++ {
		var obj3_4_1_4 go_thunder.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfg
		prefix3_4_1_4 := prefix3_4_1 + fmt.Sprintf("message_digest_cfg.%d.", i)
		obj3_4_1_4.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = d.Get(prefix3_4_1_4 + "message_digest_key").(int)
		obj3_4_1_4.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMd5Value = d.Get(prefix3_4_1_4 + "md5_value").(string)
		obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey = append(obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalMessageDigestCfgMessageDigestKey, obj3_4_1_4)
	}

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalMtu = d.Get(prefix3_4_1 + "mtu").(int)
	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalMtuIgnore = d.Get(prefix3_4_1 + "mtu_ignore").(int)

	var obj3_4_1_5 go_thunder.InterfaceLifInstanceIPOspfOspfGlobalNetwork
	prefix3_4_1_5 := prefix3_4_1 + "network.0."
	obj3_4_1_5.InterfaceLifInstanceIPOspfOspfGlobalNetworkBroadcast = d.Get(prefix3_4_1_5 + "broadcast").(int)
	obj3_4_1_5.InterfaceLifInstanceIPOspfOspfGlobalNetworkNonBroadcast = d.Get(prefix3_4_1_5 + "non_broadcast").(int)
	obj3_4_1_5.InterfaceLifInstanceIPOspfOspfGlobalNetworkPointToPoint = d.Get(prefix3_4_1_5 + "point_to_point").(int)
	obj3_4_1_5.InterfaceLifInstanceIPOspfOspfGlobalNetworkPointToMultipoint = d.Get(prefix3_4_1_5 + "point_to_multipoint").(int)
	obj3_4_1_5.InterfaceLifInstanceIPOspfOspfGlobalNetworkP2MpNbma = d.Get(prefix3_4_1_5 + "p2mp_nbma").(int)

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalNetworkBroadcast = obj3_4_1_5

	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalPriority = d.Get(prefix3_4_1 + "priority").(int)
	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalRetransmitInterval = d.Get(prefix3_4_1 + "retransmit_interval").(int)
	obj3_4_1.InterfaceLifInstanceIPOspfOspfGlobalTransmitDelay = d.Get(prefix3_4_1 + "transmit_delay").(int)

	obj3_4.InterfaceLifInstanceIPOspfOspfGlobalAuthenticationCfg = obj3_4_1

	InterfaceLifInstanceIPOspfOspfIPListCount := d.Get(prefix3_4 + "ospf_ip_list.#").(int)
	obj3_4.InterfaceLifInstanceIPOspfOspfIPListIPAddr = make([]go_thunder.InterfaceLifInstanceIPOspfOspfIPList, 0, InterfaceLifInstanceIPOspfOspfIPListCount)

	for i := 0; i < InterfaceLifInstanceIPOspfOspfIPListCount; i++ {
		var obj3_4_2 go_thunder.InterfaceLifInstanceIPOspfOspfIPList
		prefix3_4_2 := prefix3_4 + fmt.Sprintf("ospf_ip_list.%d.", i)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListIPAddr = d.Get(prefix3_4_2 + "ip_addr").(string)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListAuthentication = d.Get(prefix3_4_2 + "authentication").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListValue = d.Get(prefix3_4_2 + "value").(string)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListAuthenticationKey = d.Get(prefix3_4_2 + "authentication_key").(string)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListCost = d.Get(prefix3_4_2 + "cost").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListDatabaseFilter = d.Get(prefix3_4_2 + "database_filter").(string)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListOut = d.Get(prefix3_4_2 + "out").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListDeadInterval = d.Get(prefix3_4_2 + "dead_interval").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListHelloInterval = d.Get(prefix3_4_2 + "hello_interval").(int)

		InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgCount := d.Get(prefix3_4_2 + "message_digest_cfg.#").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = make([]go_thunder.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfg, 0, InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgCount)

		for i := 0; i < InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgCount; i++ {
			var obj3_4_2_1 go_thunder.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfg
			prefix3_4_2_1 := prefix3_4_2 + fmt.Sprintf("message_digest_cfg.%d.", i)
			obj3_4_2_1.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = d.Get(prefix3_4_2_1 + "message_digest_key").(int)
			obj3_4_2_1.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMd5Value = d.Get(prefix3_4_2_1 + "md5_value").(string)
			obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey = append(obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListMessageDigestCfgMessageDigestKey, obj3_4_2_1)
		}

		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListMtuIgnore = d.Get(prefix3_4_2 + "mtu_ignore").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListPriority = d.Get(prefix3_4_2 + "priority").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListRetransmitInterval = d.Get(prefix3_4_2 + "retransmit_interval").(int)
		obj3_4_2.InterfaceLifInstanceIPOspfOspfIPListTransmitDelay = d.Get(prefix3_4_2 + "transmit_delay").(int)
		obj3_4.InterfaceLifInstanceIPOspfOspfIPListIPAddr = append(obj3_4.InterfaceLifInstanceIPOspfOspfIPListIPAddr, obj3_4_2)
	}

	obj3.InterfaceLifInstanceIPOspfOspfGlobal = obj3_4

	c.InterfaceLifInstanceIPDhcp = obj3

	var obj4 go_thunder.InterfaceLifInstanceIpv6
	prefix4 := "ipv6.0."

	InterfaceLifInstanceIpv6AddressListCount := d.Get(prefix4 + "address_list.#").(int)
	obj4.InterfaceLifInstanceIpv6AddressListIpv6Addr = make([]go_thunder.InterfaceLifInstanceIpv6AddressList, 0, InterfaceLifInstanceIpv6AddressListCount)

	for i := 0; i < InterfaceLifInstanceIpv6AddressListCount; i++ {
		var obj4_1 go_thunder.InterfaceLifInstanceIpv6AddressList
		prefix4_1 := prefix4 + fmt.Sprintf("address_list.%d.", i)
		obj4_1.InterfaceLifInstanceIpv6AddressListIpv6Addr = d.Get(prefix4_1 + "ipv6_addr").(string)
		obj4_1.InterfaceLifInstanceIpv6AddressListAnycast = d.Get(prefix4_1 + "anycast").(int)
		obj4_1.InterfaceLifInstanceIpv6AddressListLinkLocal = d.Get(prefix4_1 + "link_local").(int)
		obj4.InterfaceLifInstanceIpv6AddressListIpv6Addr = append(obj4.InterfaceLifInstanceIpv6AddressListIpv6Addr, obj4_1)
	}

	obj4.InterfaceLifInstanceIpv6Ipv6Enable = d.Get(prefix4 + "ipv6_enable").(int)
	obj4.InterfaceLifInstanceIpv6Inside = d.Get(prefix4 + "inside").(int)
	obj4.InterfaceLifInstanceIpv6Outside = d.Get(prefix4 + "outside").(int)

	var obj4_2 go_thunder.InterfaceLifInstanceIpv6Router
	prefix4_2 := prefix4 + "router.0."

	var obj4_2_1 go_thunder.InterfaceLifInstanceIpv6RouterRipng
	prefix4_2_1 := prefix4_2 + "ripng.0."
	obj4_2_1.InterfaceLifInstanceIpv6RouterRipngRip = d.Get(prefix4_2_1 + "rip").(int)

	obj4_2.InterfaceLifInstanceIpv6RouterRipngRip = obj4_2_1

	var obj4_2_2 go_thunder.InterfaceLifInstanceIpv6RouterOspf
	prefix4_2_2 := prefix4_2 + "ospf.0."

	InterfaceLifInstanceIpv6RouterOspfAreaListCount := d.Get(prefix4_2_2 + "area_list.#").(int)
	obj4_2_2.InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDNum = make([]go_thunder.InterfaceLifInstanceIpv6RouterOspfAreaList, 0, InterfaceLifInstanceIpv6RouterOspfAreaListCount)

	for i := 0; i < InterfaceLifInstanceIpv6RouterOspfAreaListCount; i++ {
		var obj4_2_2_1 go_thunder.InterfaceLifInstanceIpv6RouterOspfAreaList
		prefix4_2_2_1 := prefix4_2_2 + fmt.Sprintf("area_list.%d.", i)
		obj4_2_2_1.InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDNum = d.Get(prefix4_2_2_1 + "area_id_num").(int)
		obj4_2_2_1.InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDAddr = d.Get(prefix4_2_2_1 + "area_id_addr").(string)
		obj4_2_2_1.InterfaceLifInstanceIpv6RouterOspfAreaListTag = d.Get(prefix4_2_2_1 + "tag").(string)
		obj4_2_2_1.InterfaceLifInstanceIpv6RouterOspfAreaListInstanceID = d.Get(prefix4_2_2_1 + "instance_id").(int)
		obj4_2_2.InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDNum = append(obj4_2_2.InterfaceLifInstanceIpv6RouterOspfAreaListAreaIDNum, obj4_2_2_1)
	}

	obj4_2.InterfaceLifInstanceIpv6RouterOspfAreaList = obj4_2_2

	var obj4_2_3 go_thunder.InterfaceLifInstanceIpv6RouterIsis
	prefix4_2_3 := prefix4_2 + "isis.0."
	obj4_2_3.InterfaceLifInstanceIpv6RouterIsisTag = d.Get(prefix4_2_3 + "tag").(string)

	obj4_2.InterfaceLifInstanceIpv6RouterIsisTag = obj4_2_3

	obj4.InterfaceLifInstanceIpv6RouterRipng = obj4_2

	var obj4_3 go_thunder.InterfaceLifInstanceIpv6Ospf
	prefix4_3 := prefix4 + "ospf.0."

	InterfaceLifInstanceIpv6OspfNetworkListCount := d.Get(prefix4_3 + "network_list.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfNetworkListBroadcastType = make([]go_thunder.InterfaceLifInstanceIpv6OspfNetworkList, 0, InterfaceLifInstanceIpv6OspfNetworkListCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfNetworkListCount; i++ {
		var obj4_3_1 go_thunder.InterfaceLifInstanceIpv6OspfNetworkList
		prefix4_3_1 := prefix4_3 + fmt.Sprintf("network_list.%d.", i)
		obj4_3_1.InterfaceLifInstanceIpv6OspfNetworkListBroadcastType = d.Get(prefix4_3_1 + "broadcast_type").(string)
		obj4_3_1.InterfaceLifInstanceIpv6OspfNetworkListP2MpNbma = d.Get(prefix4_3_1 + "p2mp_nbma").(int)
		obj4_3_1.InterfaceLifInstanceIpv6OspfNetworkListNetworkInstanceID = d.Get(prefix4_3_1 + "network_instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfNetworkListBroadcastType = append(obj4_3.InterfaceLifInstanceIpv6OspfNetworkListBroadcastType, obj4_3_1)
	}

	obj4_3.InterfaceLifInstanceIpv6OspfBfd = d.Get(prefix4_3 + "bfd").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfDisable = d.Get(prefix4_3 + "disable").(int)

	InterfaceLifInstanceIpv6OspfCostCfgCount := d.Get(prefix4_3 + "cost_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfCostCfgCost = make([]go_thunder.InterfaceLifInstanceIpv6OspfCostCfg, 0, InterfaceLifInstanceIpv6OspfCostCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfCostCfgCount; i++ {
		var obj4_3_2 go_thunder.InterfaceLifInstanceIpv6OspfCostCfg
		prefix4_3_2 := prefix4_3 + fmt.Sprintf("cost_cfg.%d.", i)
		obj4_3_2.InterfaceLifInstanceIpv6OspfCostCfgCost = d.Get(prefix4_3_2 + "cost").(int)
		obj4_3_2.InterfaceLifInstanceIpv6OspfCostCfgInstanceID = d.Get(prefix4_3_2 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfCostCfgCost = append(obj4_3.InterfaceLifInstanceIpv6OspfCostCfgCost, obj4_3_2)
	}

	InterfaceLifInstanceIpv6OspfDeadIntervalCfgCount := d.Get(prefix4_3 + "dead_interval_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfDeadIntervalCfgDeadInterval = make([]go_thunder.InterfaceLifInstanceIpv6OspfDeadIntervalCfg, 0, InterfaceLifInstanceIpv6OspfDeadIntervalCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfDeadIntervalCfgCount; i++ {
		var obj4_3_3 go_thunder.InterfaceLifInstanceIpv6OspfDeadIntervalCfg
		prefix4_3_3 := prefix4_3 + fmt.Sprintf("dead_interval_cfg.%d.", i)
		obj4_3_3.InterfaceLifInstanceIpv6OspfDeadIntervalCfgDeadInterval = d.Get(prefix4_3_3 + "dead_interval").(int)
		obj4_3_3.InterfaceLifInstanceIpv6OspfDeadIntervalCfgInstanceID = d.Get(prefix4_3_3 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfDeadIntervalCfgDeadInterval = append(obj4_3.InterfaceLifInstanceIpv6OspfDeadIntervalCfgDeadInterval, obj4_3_3)
	}

	InterfaceLifInstanceIpv6OspfHelloIntervalCfgCount := d.Get(prefix4_3 + "hello_interval_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfHelloIntervalCfgHelloInterval = make([]go_thunder.InterfaceLifInstanceIpv6OspfHelloIntervalCfg, 0, InterfaceLifInstanceIpv6OspfHelloIntervalCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfHelloIntervalCfgCount; i++ {
		var obj4_3_4 go_thunder.InterfaceLifInstanceIpv6OspfHelloIntervalCfg
		prefix4_3_4 := prefix4_3 + fmt.Sprintf("hello_interval_cfg.%d.", i)
		obj4_3_4.InterfaceLifInstanceIpv6OspfHelloIntervalCfgHelloInterval = d.Get(prefix4_3_4 + "hello_interval").(int)
		obj4_3_4.InterfaceLifInstanceIpv6OspfHelloIntervalCfgInstanceID = d.Get(prefix4_3_4 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfHelloIntervalCfgHelloInterval = append(obj4_3.InterfaceLifInstanceIpv6OspfHelloIntervalCfgHelloInterval, obj4_3_4)
	}

	InterfaceLifInstanceIpv6OspfMtuIgnoreCfgCount := d.Get(prefix4_3 + "mtu_ignore_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = make([]go_thunder.InterfaceLifInstanceIpv6OspfMtuIgnoreCfg, 0, InterfaceLifInstanceIpv6OspfMtuIgnoreCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfMtuIgnoreCfgCount; i++ {
		var obj4_3_5 go_thunder.InterfaceLifInstanceIpv6OspfMtuIgnoreCfg
		prefix4_3_5 := prefix4_3 + fmt.Sprintf("mtu_ignore_cfg.%d.", i)
		obj4_3_5.InterfaceLifInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = d.Get(prefix4_3_5 + "mtu_ignore").(int)
		obj4_3_5.InterfaceLifInstanceIpv6OspfMtuIgnoreCfgInstanceID = d.Get(prefix4_3_5 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfMtuIgnoreCfgMtuIgnore = append(obj4_3.InterfaceLifInstanceIpv6OspfMtuIgnoreCfgMtuIgnore, obj4_3_5)
	}

	InterfaceLifInstanceIpv6OspfNeighborCfgCount := d.Get(prefix4_3 + "neighbor_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfNeighborCfgNeighbor = make([]go_thunder.InterfaceLifInstanceIpv6OspfNeighborCfg, 0, InterfaceLifInstanceIpv6OspfNeighborCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfNeighborCfgCount; i++ {
		var obj4_3_6 go_thunder.InterfaceLifInstanceIpv6OspfNeighborCfg
		prefix4_3_6 := prefix4_3 + fmt.Sprintf("neighbor_cfg.%d.", i)
		obj4_3_6.InterfaceLifInstanceIpv6OspfNeighborCfgNeighbor = d.Get(prefix4_3_6 + "neighbor").(string)
		obj4_3_6.InterfaceLifInstanceIpv6OspfNeighborCfgNeigInst = d.Get(prefix4_3_6 + "neig_inst").(int)
		obj4_3_6.InterfaceLifInstanceIpv6OspfNeighborCfgNeighborCost = d.Get(prefix4_3_6 + "neighbor_cost").(int)
		obj4_3_6.InterfaceLifInstanceIpv6OspfNeighborCfgNeighborPollInterval = d.Get(prefix4_3_6 + "neighbor_poll_interval").(int)
		obj4_3_6.InterfaceLifInstanceIpv6OspfNeighborCfgNeighborPriority = d.Get(prefix4_3_6 + "neighbor_priority").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfNeighborCfgNeighbor = append(obj4_3.InterfaceLifInstanceIpv6OspfNeighborCfgNeighbor, obj4_3_6)
	}

	InterfaceLifInstanceIpv6OspfPriorityCfgCount := d.Get(prefix4_3 + "priority_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfPriorityCfgPriority = make([]go_thunder.InterfaceLifInstanceIpv6OspfPriorityCfg, 0, InterfaceLifInstanceIpv6OspfPriorityCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfPriorityCfgCount; i++ {
		var obj4_3_7 go_thunder.InterfaceLifInstanceIpv6OspfPriorityCfg
		prefix4_3_7 := prefix4_3 + fmt.Sprintf("priority_cfg.%d.", i)
		obj4_3_7.InterfaceLifInstanceIpv6OspfPriorityCfgPriority = d.Get(prefix4_3_7 + "priority").(int)
		obj4_3_7.InterfaceLifInstanceIpv6OspfPriorityCfgInstanceID = d.Get(prefix4_3_7 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfPriorityCfgPriority = append(obj4_3.InterfaceLifInstanceIpv6OspfPriorityCfgPriority, obj4_3_7)
	}

	InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgCount := d.Get(prefix4_3 + "retransmit_interval_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = make([]go_thunder.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfg, 0, InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgCount; i++ {
		var obj4_3_8 go_thunder.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfg
		prefix4_3_8 := prefix4_3 + fmt.Sprintf("retransmit_interval_cfg.%d.", i)
		obj4_3_8.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = d.Get(prefix4_3_8 + "retransmit_interval").(int)
		obj4_3_8.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgInstanceID = d.Get(prefix4_3_8 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval = append(obj4_3.InterfaceLifInstanceIpv6OspfRetransmitIntervalCfgRetransmitInterval, obj4_3_8)
	}

	InterfaceLifInstanceIpv6OspfTransmitDelayCfgCount := d.Get(prefix4_3 + "transmit_delay_cfg.#").(int)
	obj4_3.InterfaceLifInstanceIpv6OspfTransmitDelayCfgTransmitDelay = make([]go_thunder.InterfaceLifInstanceIpv6OspfTransmitDelayCfg, 0, InterfaceLifInstanceIpv6OspfTransmitDelayCfgCount)

	for i := 0; i < InterfaceLifInstanceIpv6OspfTransmitDelayCfgCount; i++ {
		var obj4_3_9 go_thunder.InterfaceLifInstanceIpv6OspfTransmitDelayCfg
		prefix4_3_9 := prefix4_3 + fmt.Sprintf("transmit_delay_cfg.%d.", i)
		obj4_3_9.InterfaceLifInstanceIpv6OspfTransmitDelayCfgTransmitDelay = d.Get(prefix4_3_9 + "transmit_delay").(int)
		obj4_3_9.InterfaceLifInstanceIpv6OspfTransmitDelayCfgInstanceID = d.Get(prefix4_3_9 + "instance_id").(int)
		obj4_3.InterfaceLifInstanceIpv6OspfTransmitDelayCfgTransmitDelay = append(obj4_3.InterfaceLifInstanceIpv6OspfTransmitDelayCfgTransmitDelay, obj4_3_9)
	}

	obj4.InterfaceLifInstanceIpv6OspfNetworkList = obj4_3

	c.InterfaceLifInstanceIpv6AddressList = obj4

	var obj5 go_thunder.InterfaceLifInstanceBfd
	prefix5 := "bfd.0."

	var obj5_1 go_thunder.InterfaceLifInstanceBfdAuthentication
	prefix5_1 := prefix5 + "authentication.0."
	obj5_1.InterfaceLifInstanceBfdAuthenticationKeyID = d.Get(prefix5_1 + "key_id").(int)
	obj5_1.InterfaceLifInstanceBfdAuthenticationMethod = d.Get(prefix5_1 + "method").(string)
	obj5_1.InterfaceLifInstanceBfdAuthenticationPassword = d.Get(prefix5_1 + "password").(string)

	obj5.InterfaceLifInstanceBfdAuthenticationKeyID = obj5_1

	obj5.InterfaceLifInstanceBfdEcho = d.Get(prefix5 + "echo").(int)
	obj5.InterfaceLifInstanceBfdDemand = d.Get(prefix5 + "demand").(int)

	var obj5_2 go_thunder.InterfaceLifInstanceBfdIntervalCfg
	prefix5_2 := prefix5 + "interval_cfg.0."
	obj5_2.InterfaceLifInstanceBfdIntervalCfgInterval = d.Get(prefix5_2 + "interval").(int)
	obj5_2.InterfaceLifInstanceBfdIntervalCfgMinRx = d.Get(prefix5_2 + "min_rx").(int)
	obj5_2.InterfaceLifInstanceBfdIntervalCfgMultiplier = d.Get(prefix5_2 + "multiplier").(int)

	obj5.InterfaceLifInstanceBfdIntervalCfgInterval = obj5_2

	c.InterfaceLifInstanceBfdAuthentication = obj5

	var obj6 go_thunder.InterfaceLifInstanceIsis
	prefix6 := "isis.0."

	var obj6_1 go_thunder.InterfaceLifInstanceIsisAuthentication
	prefix6_1 := prefix6 + "authentication.0."

	InterfaceLifInstanceIsisAuthenticationSendOnlyListCount := d.Get(prefix6_1 + "send_only_list.#").(int)
	obj6_1.InterfaceLifInstanceIsisAuthenticationSendOnlyListSendOnly = make([]go_thunder.InterfaceLifInstanceIsisAuthenticationSendOnlyList, 0, InterfaceLifInstanceIsisAuthenticationSendOnlyListCount)

	for i := 0; i < InterfaceLifInstanceIsisAuthenticationSendOnlyListCount; i++ {
		var obj6_1_1 go_thunder.InterfaceLifInstanceIsisAuthenticationSendOnlyList
		prefix6_1_1 := prefix6_1 + fmt.Sprintf("send_only_list.%d.", i)
		obj6_1_1.InterfaceLifInstanceIsisAuthenticationSendOnlyListSendOnly = d.Get(prefix6_1_1 + "send_only").(int)
		obj6_1_1.InterfaceLifInstanceIsisAuthenticationSendOnlyListLevel = d.Get(prefix6_1_1 + "level").(string)
		obj6_1.InterfaceLifInstanceIsisAuthenticationSendOnlyListSendOnly = append(obj6_1.InterfaceLifInstanceIsisAuthenticationSendOnlyListSendOnly, obj6_1_1)
	}

	InterfaceLifInstanceIsisAuthenticationModeListCount := d.Get(prefix6_1 + "mode_list.#").(int)
	obj6_1.InterfaceLifInstanceIsisAuthenticationModeListMode = make([]go_thunder.InterfaceLifInstanceIsisAuthenticationModeList, 0, InterfaceLifInstanceIsisAuthenticationModeListCount)

	for i := 0; i < InterfaceLifInstanceIsisAuthenticationModeListCount; i++ {
		var obj6_1_2 go_thunder.InterfaceLifInstanceIsisAuthenticationModeList
		prefix6_1_2 := prefix6_1 + fmt.Sprintf("mode_list.%d.", i)
		obj6_1_2.InterfaceLifInstanceIsisAuthenticationModeListMode = d.Get(prefix6_1_2 + "mode").(string)
		obj6_1_2.InterfaceLifInstanceIsisAuthenticationModeListLevel = d.Get(prefix6_1_2 + "level").(string)
		obj6_1.InterfaceLifInstanceIsisAuthenticationModeListMode = append(obj6_1.InterfaceLifInstanceIsisAuthenticationModeListMode, obj6_1_2)
	}

	InterfaceLifInstanceIsisAuthenticationKeyChainListCount := d.Get(prefix6_1 + "key_chain_list.#").(int)
	obj6_1.InterfaceLifInstanceIsisAuthenticationKeyChainListKeyChain = make([]go_thunder.InterfaceLifInstanceIsisAuthenticationKeyChainList, 0, InterfaceLifInstanceIsisAuthenticationKeyChainListCount)

	for i := 0; i < InterfaceLifInstanceIsisAuthenticationKeyChainListCount; i++ {
		var obj6_1_3 go_thunder.InterfaceLifInstanceIsisAuthenticationKeyChainList
		prefix6_1_3 := prefix6_1 + fmt.Sprintf("key_chain_list.%d.", i)
		obj6_1_3.InterfaceLifInstanceIsisAuthenticationKeyChainListKeyChain = d.Get(prefix6_1_3 + "key_chain").(string)
		obj6_1_3.InterfaceLifInstanceIsisAuthenticationKeyChainListLevel = d.Get(prefix6_1_3 + "level").(string)
		obj6_1.InterfaceLifInstanceIsisAuthenticationKeyChainListKeyChain = append(obj6_1.InterfaceLifInstanceIsisAuthenticationKeyChainListKeyChain, obj6_1_3)
	}

	obj6.InterfaceLifInstanceIsisAuthenticationSendOnlyList = obj6_1

	var obj6_2 go_thunder.InterfaceLifInstanceIsisBfdCfg
	prefix6_2 := prefix6 + "bfd_cfg.0."
	obj6_2.InterfaceLifInstanceIsisBfdCfgBfd = d.Get(prefix6_2 + "bfd").(int)
	obj6_2.InterfaceLifInstanceIsisBfdCfgDisable = d.Get(prefix6_2 + "disable").(int)

	obj6.InterfaceLifInstanceIsisBfdCfgBfd = obj6_2

	obj6.InterfaceLifInstanceIsisCircuitType = d.Get(prefix6 + "circuit_type").(string)

	InterfaceLifInstanceIsisCsnpIntervalListCount := d.Get(prefix6 + "csnp_interval_list.#").(int)
	obj6.InterfaceLifInstanceIsisCsnpIntervalListCsnpInterval = make([]go_thunder.InterfaceLifInstanceIsisCsnpIntervalList, 0, InterfaceLifInstanceIsisCsnpIntervalListCount)

	for i := 0; i < InterfaceLifInstanceIsisCsnpIntervalListCount; i++ {
		var obj6_3 go_thunder.InterfaceLifInstanceIsisCsnpIntervalList
		prefix6_3 := prefix6 + fmt.Sprintf("csnp_interval_list.%d.", i)
		obj6_3.InterfaceLifInstanceIsisCsnpIntervalListCsnpInterval = d.Get(prefix6_3 + "csnp_interval").(int)
		obj6_3.InterfaceLifInstanceIsisCsnpIntervalListLevel = d.Get(prefix6_3 + "level").(string)
		obj6.InterfaceLifInstanceIsisCsnpIntervalListCsnpInterval = append(obj6.InterfaceLifInstanceIsisCsnpIntervalListCsnpInterval, obj6_3)
	}

	obj6.InterfaceLifInstanceIsisPadding = d.Get(prefix6 + "padding").(int)

	InterfaceLifInstanceIsisHelloIntervalListCount := d.Get(prefix6 + "hello_interval_list.#").(int)
	obj6.InterfaceLifInstanceIsisHelloIntervalListHelloInterval = make([]go_thunder.InterfaceLifInstanceIsisHelloIntervalList, 0, InterfaceLifInstanceIsisHelloIntervalListCount)

	for i := 0; i < InterfaceLifInstanceIsisHelloIntervalListCount; i++ {
		var obj6_4 go_thunder.InterfaceLifInstanceIsisHelloIntervalList
		prefix6_4 := prefix6 + fmt.Sprintf("hello_interval_list.%d.", i)
		obj6_4.InterfaceLifInstanceIsisHelloIntervalListHelloInterval = d.Get(prefix6_4 + "hello_interval").(int)
		obj6_4.InterfaceLifInstanceIsisHelloIntervalListLevel = d.Get(prefix6_4 + "level").(string)
		obj6.InterfaceLifInstanceIsisHelloIntervalListHelloInterval = append(obj6.InterfaceLifInstanceIsisHelloIntervalListHelloInterval, obj6_4)
	}

	InterfaceLifInstanceIsisHelloIntervalMinimalListCount := d.Get(prefix6 + "hello_interval_minimal_list.#").(int)
	obj6.InterfaceLifInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = make([]go_thunder.InterfaceLifInstanceIsisHelloIntervalMinimalList, 0, InterfaceLifInstanceIsisHelloIntervalMinimalListCount)

	for i := 0; i < InterfaceLifInstanceIsisHelloIntervalMinimalListCount; i++ {
		var obj6_5 go_thunder.InterfaceLifInstanceIsisHelloIntervalMinimalList
		prefix6_5 := prefix6 + fmt.Sprintf("hello_interval_minimal_list.%d.", i)
		obj6_5.InterfaceLifInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = d.Get(prefix6_5 + "hello_interval_minimal").(int)
		obj6_5.InterfaceLifInstanceIsisHelloIntervalMinimalListLevel = d.Get(prefix6_5 + "level").(string)
		obj6.InterfaceLifInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal = append(obj6.InterfaceLifInstanceIsisHelloIntervalMinimalListHelloIntervalMinimal, obj6_5)
	}

	InterfaceLifInstanceIsisHelloMultiplierListCount := d.Get(prefix6 + "hello_multiplier_list.#").(int)
	obj6.InterfaceLifInstanceIsisHelloMultiplierListHelloMultiplier = make([]go_thunder.InterfaceLifInstanceIsisHelloMultiplierList, 0, InterfaceLifInstanceIsisHelloMultiplierListCount)

	for i := 0; i < InterfaceLifInstanceIsisHelloMultiplierListCount; i++ {
		var obj6_6 go_thunder.InterfaceLifInstanceIsisHelloMultiplierList
		prefix6_6 := prefix6 + fmt.Sprintf("hello_multiplier_list.%d.", i)
		obj6_6.InterfaceLifInstanceIsisHelloMultiplierListHelloMultiplier = d.Get(prefix6_6 + "hello_multiplier").(int)
		obj6_6.InterfaceLifInstanceIsisHelloMultiplierListLevel = d.Get(prefix6_6 + "level").(string)
		obj6.InterfaceLifInstanceIsisHelloMultiplierListHelloMultiplier = append(obj6.InterfaceLifInstanceIsisHelloMultiplierListHelloMultiplier, obj6_6)
	}

	obj6.InterfaceLifInstanceIsisLspInterval = d.Get(prefix6 + "lsp_interval").(int)

	var obj6_7 go_thunder.InterfaceLifInstanceIsisMeshGroup
	prefix6_7 := prefix6 + "mesh_group.0."
	obj6_7.InterfaceLifInstanceIsisMeshGroupValue = d.Get(prefix6_7 + "value").(int)
	obj6_7.InterfaceLifInstanceIsisMeshGroupBlocked = d.Get(prefix6_7 + "blocked").(int)

	obj6.InterfaceLifInstanceIsisMeshGroupValue = obj6_7

	InterfaceLifInstanceIsisMetricListCount := d.Get(prefix6 + "metric_list.#").(int)
	obj6.InterfaceLifInstanceIsisMetricListMetric = make([]go_thunder.InterfaceLifInstanceIsisMetricList, 0, InterfaceLifInstanceIsisMetricListCount)

	for i := 0; i < InterfaceLifInstanceIsisMetricListCount; i++ {
		var obj6_8 go_thunder.InterfaceLifInstanceIsisMetricList
		prefix6_8 := prefix6 + fmt.Sprintf("metric_list.%d.", i)
		obj6_8.InterfaceLifInstanceIsisMetricListMetric = d.Get(prefix6_8 + "metric").(int)
		obj6_8.InterfaceLifInstanceIsisMetricListLevel = d.Get(prefix6_8 + "level").(string)
		obj6.InterfaceLifInstanceIsisMetricListMetric = append(obj6.InterfaceLifInstanceIsisMetricListMetric, obj6_8)
	}

	obj6.InterfaceLifInstanceIsisNetwork = d.Get(prefix6 + "network").(string)

	InterfaceLifInstanceIsisPasswordListCount := d.Get(prefix6 + "password_list.#").(int)
	obj6.InterfaceLifInstanceIsisPasswordListPassword = make([]go_thunder.InterfaceLifInstanceIsisPasswordList, 0, InterfaceLifInstanceIsisPasswordListCount)

	for i := 0; i < InterfaceLifInstanceIsisPasswordListCount; i++ {
		var obj6_9 go_thunder.InterfaceLifInstanceIsisPasswordList
		prefix6_9 := prefix6 + fmt.Sprintf("password_list.%d.", i)
		obj6_9.InterfaceLifInstanceIsisPasswordListPassword = d.Get(prefix6_9 + "password").(string)
		obj6_9.InterfaceLifInstanceIsisPasswordListLevel = d.Get(prefix6_9 + "level").(string)
		obj6.InterfaceLifInstanceIsisPasswordListPassword = append(obj6.InterfaceLifInstanceIsisPasswordListPassword, obj6_9)
	}

	InterfaceLifInstanceIsisPriorityListCount := d.Get(prefix6 + "priority_list.#").(int)
	obj6.InterfaceLifInstanceIsisPriorityListPriority = make([]go_thunder.InterfaceLifInstanceIsisPriorityList, 0, InterfaceLifInstanceIsisPriorityListCount)

	for i := 0; i < InterfaceLifInstanceIsisPriorityListCount; i++ {
		var obj6_10 go_thunder.InterfaceLifInstanceIsisPriorityList
		prefix6_10 := prefix6 + fmt.Sprintf("priority_list.%d.", i)
		obj6_10.InterfaceLifInstanceIsisPriorityListPriority = d.Get(prefix6_10 + "priority").(int)
		obj6_10.InterfaceLifInstanceIsisPriorityListLevel = d.Get(prefix6_10 + "level").(string)
		obj6.InterfaceLifInstanceIsisPriorityListPriority = append(obj6.InterfaceLifInstanceIsisPriorityListPriority, obj6_10)
	}

	obj6.InterfaceLifInstanceIsisRetransmitInterval = d.Get(prefix6 + "retransmit_interval").(int)

	InterfaceLifInstanceIsisWideMetricListCount := d.Get(prefix6 + "wide_metric_list.#").(int)
	obj6.InterfaceLifInstanceIsisWideMetricListWideMetric = make([]go_thunder.InterfaceLifInstanceIsisWideMetricList, 0, InterfaceLifInstanceIsisWideMetricListCount)

	for i := 0; i < InterfaceLifInstanceIsisWideMetricListCount; i++ {
		var obj6_11 go_thunder.InterfaceLifInstanceIsisWideMetricList
		prefix6_11 := prefix6 + fmt.Sprintf("wide_metric_list.%d.", i)
		obj6_11.InterfaceLifInstanceIsisWideMetricListWideMetric = d.Get(prefix6_11 + "wide_metric").(int)
		obj6_11.InterfaceLifInstanceIsisWideMetricListLevel = d.Get(prefix6_11 + "level").(string)
		obj6.InterfaceLifInstanceIsisWideMetricListWideMetric = append(obj6.InterfaceLifInstanceIsisWideMetricListWideMetric, obj6_11)
	}

	c.InterfaceLifInstanceIsisAuthentication = obj6

	vc.InterfaceLifInstanceIfname = c
	return vc
}
