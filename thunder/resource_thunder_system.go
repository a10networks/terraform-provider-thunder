package thunder

//Thunder resource System

import (
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"util"
)

func resourceSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCreate,
		Update: resourceSystemUpdate,
		Read:   resourceSystemRead,
		Delete: resourceSystemDelete,
		Schema: map[string]*schema.Schema{
			"anomaly_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"attack_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ddos_attack": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"ddos_log": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"sockstress_disable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"promiscuous_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"glid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"module_ctrl_cpu": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"src_ip_hash_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"class_list_hitcount_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"geo_db_hitcount_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"domain_list_hitcount_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"dynamic_service_dns_socket_pool": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"timeout_value": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ftp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"scp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"sftp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"tftp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"http": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"https": {
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
			"bandwidth": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"counter_lib_accounting": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"control_cpu": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"data_cpu": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"mgmt_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mac_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"pci_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"shared_poll_mode": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
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
			"add_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"del_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"modify_port": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"tls_1_3_mgmt": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"multi_queue_support": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"add_cpu_core": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"core_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"delete_cpu_core": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"core_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cpu_hyper_thread": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
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
			"io_cpu": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_cores": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"link_monitor": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
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
			"port_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"port_info": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"inuse_port_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cpu_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cpu_map": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"inuse_cpu_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"set_rxtx_desc_size": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rxd_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"txd_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"set_rxtx_queue": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rxq_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"txq_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_policy": {
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
			"template_bind": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"clear_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sessions": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"clear_all_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"clear_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"link_disable_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"diseth": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dis_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"link_enable_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enaeth": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"ena_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"monitor_relation": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"link_up_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkup_ethernet1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_up_sequence1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkup_ethernet2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_up_sequence2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkup_ethernet3": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_up_sequence3": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"link_down_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkdown_ethernet1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_down_sequence1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkdown_ethernet2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_down_sequence2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkdown_ethernet3": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_down_sequence3": {
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
									"user_tag": {
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
			"mon_template": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"clear_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sessions": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"clear_all_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"clear_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"link_disable_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"diseth": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"dis_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"link_enable_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enaeth": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"ena_sequence": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"monitor_relation": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"link_up_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkup_ethernet1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_up_sequence1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkup_ethernet2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_up_sequence2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkup_ethernet3": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_up_sequence3": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"link_down_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"linkdown_ethernet1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_down_sequence1": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkdown_ethernet2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_down_sequence2": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"linkdown_ethernet3": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"link_down_sequence3": {
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
									"user_tag": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"link_block_as_down": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
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
			"memory": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"resource_usage": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ssl_context_memory": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ssl_dma_memory": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"nat_pool_addr_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l4_session_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auth_portal_html_file_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auth_portal_image_file_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_aflex_file_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"aflex_table_entry_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"class_list_ipv6_addr_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"class_list_ac_entry_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"max_aflex_authz_collection_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"radius_table_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"authz_policy_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipsec_sa_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ram_cache_memory_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"visibility": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"monitored_entity_count": {
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
			"link_capability": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
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
			"resource_accounting": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"template_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
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
									"app_resources": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gslb_device_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_device_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_device_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_geo_location_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_geo_location_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_geo_location_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_ip_list_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_ip_list_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_ip_list_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_policy_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_policy_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_policy_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_service_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_service_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_service_ip_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_ip_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_service_ip_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_service_port_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_service_port_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_service_port_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_site_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_site_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_site_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_svc_group_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_svc_group_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_svc_group_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"gslb_zone_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"gslb_zone_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"gslb_zone_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"health_monitor_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"health_monitor_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"health_monitor_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"real_port_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"real_port_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"real_port_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"real_server_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"real_server_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"real_server_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"service_group_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"service_group_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"service_group_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"virtual_server_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"virtual_server_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"virtual_server_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"virtual_port_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"virtual_port_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"virtual_port_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"cache_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cache_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"cache_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"client_ssl_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"client_ssl_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"client_ssl_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"conn_reuse_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"conn_reuse_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"conn_reuse_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"fast_tcp_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fast_tcp_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"fast_tcp_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"fast_udp_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fast_udp_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"fast_udp_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"fix_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fix_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"fix_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"http_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"http_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"http_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"link_cost_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"link_cost_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"link_cost_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"persist_cookie_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"persist_cookie_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"persist_cookie_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"persist_srcip_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"persist_srcip_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"persist_srcip_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"server_ssl_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server_ssl_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"server_ssl_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"proxy_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"proxy_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"proxy_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"stream_template_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"stream_template_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"stream_template_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"threshold": {
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
									"network_resources": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"static_ipv4_route_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_ipv4_route_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"static_ipv4_route_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"static_ipv6_route_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_ipv6_route_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"static_ipv6_route_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ipv4_acl_line_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv4_acl_line_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"ipv4_acl_line_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ipv6_acl_line_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ipv6_acl_line_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"ipv6_acl_line_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"static_arp_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_arp_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"static_arp_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"static_neighbor_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_neighbor_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"static_neighbor_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"static_mac_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"static_mac_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"static_mac_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"object_group_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_group_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"object_group_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"object_group_clause_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"object_group_clause_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"object_group_clause_min_guarantee": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"threshold": {
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
									"system_resources": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bw_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"bw_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"bw_limit_watermark_disable": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"concurrent_session_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"concurrent_session_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"l4_session_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l4_session_limit_max": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"l4_session_limit_min_guarantee": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"l4cps_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l4cps_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"l7cps_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"l7cps_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"natcps_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"natcps_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"fwcps_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fwcps_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"ssl_throughput_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ssl_throughput_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"ssl_throughput_limit_watermark_disable": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"sslcps_limit_cfg": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sslcps_limit_max": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
														},
													},
												},
												"threshold": {
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
			"trunk": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balance": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"use_l3": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"use_l4": {
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
			"ports": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_detection_interval": {
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
			"ipsec": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_round_robin": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"crypto_core": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"crypto_mem": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"fpga_decrypt": {
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
								},
							},
						},
					},
				},
			},
			"spe_profile": {
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
					},
				},
			},
			"spe_status": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ssl_status": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"deep_hrxq": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"hrxq_status": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cpu_load_sharing": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"packets_per_second": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"min": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"cpu_usage": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"low": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"high": {
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
			"per_vlan_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipmcast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mcast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"unknown_ucast": {
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
			"all_vlan_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bcast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ipmcast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"mcast": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"unknown_ucast": {
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
			"ve_mac_scheme": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ve_mac_scheme_val": {
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
			"session_reclaim_limit": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nscan_limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"scan_freq": {
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
			"ssl_scv": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
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
			"hardware": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"platformtype": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"reboot": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"shutdown": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"environment": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"hardware_forward": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
						"slb": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
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
								},
							},
						},
					},
				},
			},
			"throughput": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"ipmi": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reset": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ip": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
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
									"default_gateway": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"ipsrc": {
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
									"static": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"user": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"add": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"administrator": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"callback": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"operator": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"user": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"disable": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"privilege": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"setname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"newname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"setpass": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"newpass": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"tool": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmd": {
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
			"queuing_buffer": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
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
			"trunk_hw_hash": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
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
			"trunk_xaui_hw_hash": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
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
			"upgrade_status": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"guest_file": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cm_update_file_name_ref": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"dest_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"core": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"apps_global": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_session_on_established": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"msl_time": {
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
			"shell_privileges": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cosq_stats": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"cosq_show": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"fw": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_mempool": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"application_flow": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"basic_dpi_enable": {
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
			"password_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complexity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"aging": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"history": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"min_pswd_len": {
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
			"radius": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listen_port": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"remote": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip_list": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ip_list_name": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "",
															},
															"ip_list_secret": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "",
															},
															"ip_list_secret_string": {
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
									"secret": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"secret_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"vrid": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"attribute": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"prefix_length": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"prefix_vendor": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"prefix_number": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"custom_vendor": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"custom_number": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"vendor": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"number": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"disable_reply": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"accounting_start": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"accounting_stop": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"accounting_interim_update": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"accounting_on": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"attribute_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"uuid": {
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
								},
							},
						},
					},
				},
			},
			"geoloc_list_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"shared": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"include_geoloc_name_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include_geoloc_name_val": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"exclude_geoloc_name_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exclude_geoloc_name_val": {
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
					},
				},
			},
			"geoloc_name_helper": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"geolocation_file": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"error_info": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
			"geoloc": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"geo_location": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_iana": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"geo_location_geolite2_city": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"geolite2_city_include_ipv6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"geo_location_geolite2_country": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"geolite2_country_include_ipv6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"geoloc_load_file_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_location_load_filename": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"template_name": {
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
						"entry_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"geo_locn_obj_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"geo_locn_multiple_addresses": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"first_ip_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"geol_ipv4_mask": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ip_addr2": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"first_ipv6_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"geol_ipv6_mask": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
												"ipv6_addr2": {
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
								},
							},
						},
					},
				},
			},
			"ip_threat_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
						"ipv4_source_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ip_threat_action_tmpl": {
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
						"ipv4_dest_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ip_threat_action_tmpl": {
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
						"ipv6_source_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ip_threat_action_tmpl": {
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
						"ipv6_dest_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"class_list_cfg": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"class_list": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"ip_threat_action_tmpl": {
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
			"fpga_drop": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"dpdk_stats": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"fpga_core_crc": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitor_disable": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"reboot_enable": {
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
			"psu_info": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"gui_image_list": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"syslog_time_msec": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_flag": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ipmi_service": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable": {
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
			"app_performance": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"ssl_req_q": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"tcp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
						"rate_limit_reset_unknown_conn": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pkt_rate_for_reset_unknown_conn": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
									"log_for_reset_unknown_conn": {
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
			"icmp": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"icmp6": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"ip_stats": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"ip6_stats": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"domain_list_info": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"ip_dns_cache": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"uuid": {
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
					},
				},
			},
			"icmp_rate": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"dns": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"dns_cache": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"session": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"ndisc_ra": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"tcp_stats": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
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
					},
				},
			},
			"telemetry_log": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top_k_source_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"top_k_app_svc_list": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"device_status": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"environment": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"partition_metrics": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
	}
}

func resourceSystemCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating System (Inside resourceSystemCreate) ")

		data := dataToSystem(d)
		logger.Println("[INFO] received formatted data from method data to System --")
		d.SetId("1")
		go_thunder.PostSystem(client.Token, data, client.Host)

		return resourceSystemRead(d, meta)

	}
	return nil
}

func resourceSystemRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading System (Inside resourceSystemRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSystem(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return err
	}
	return nil
}

func resourceSystemUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceSystemRead(d, meta)
}

func resourceSystemDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceSystemRead(d, meta)
}
func dataToSystem(d *schema.ResourceData) go_thunder.System {
	var vc go_thunder.System
	var c go_thunder.SystemInstance
	c.AnomalyLog = d.Get("anomaly_log").(int)
	c.AttackLog = d.Get("attack_log").(int)
	c.DdosAttack = d.Get("ddos_attack").(int)
	c.DdosLog = d.Get("ddos_log").(int)
	c.SockstressDisable = d.Get("sockstress_disable").(int)
	c.PromiscuousMode = d.Get("promiscuous_mode").(int)
	c.Glid = d.Get("glid").(int)
	c.ModuleCtrlCPU = d.Get("module_ctrl_cpu").(string)
	c.SrcIPHashEnable = d.Get("src_ip_hash_enable").(int)
	c.ClassListHitcountEnable = d.Get("class_list_hitcount_enable").(int)
	c.GeoDbHitcountEnable = d.Get("geo_db_hitcount_enable").(int)
	c.DomainListHitcountEnable = d.Get("domain_list_hitcount_enable").(int)
	c.DynamicServiceDNSSocketPool = d.Get("dynamic_service_dns_socket_pool").(int)
	/*
		var obj1 go_thunder.TimeoutValue
		prefix1 := "timeout_value.0."
		obj1.Ftp = d.Get(prefix1 + "ftp").(int)
		obj1.Scp = d.Get(prefix1 + "scp").(int)
		obj1.Sftp = d.Get(prefix1 + "sftp").(int)
		obj1.Tftp = d.Get(prefix1 + "tftp").(int)
		obj1.HTTP = d.Get(prefix1 + "http").(int)
		obj1.HTTPS = d.Get(prefix1 + "https").(int)

		c.Ftp = obj1

		// var obj2 go_thunder.Bandwidth
		// prefix2 := "bandwidth.0."

		// SamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
		// obj2.Counters1 = make([]go_thunder.SystemSamplingEnable, 0, SamplingEnableCount)

		// for i := 0; i < SamplingEnableCount; i++ {
		// var obj2_1 go_thunder.SystemSamplingEnable
		// prefix2_1 := fmt.Sprintf(prefix2 + "sampling_enable.%d.",i)
		// obj2_1.Counters1 = d.Get(prefix2_1+"counters1").(string)
		// obj2.Counters1 = append(obj2.Counters1, obj2_1)
		// }

		//c.SystemSamplingEnable = obj2

		// var obj3 go_thunder.CounterLibAccounting
		// prefix := "counter_lib_accounting.0."

		// c. = obj3

		// var obj4 go_thunder.ControlCpu
		// prefix := "control_cpu.0."

		// c. = obj4

		// var obj5 go_thunder.DataCpu
		// prefix := "data_cpu.0."

		// c. = obj5

		var obj6 go_thunder.MgmtPort
		prefix6 := "mgmt_port.0."
		obj6.PortIndex = d.Get(prefix6 + "port_index").(int)
		obj6.MacAddress = d.Get(prefix6 + "mac_address").(string)
		obj6.PciAddress = d.Get(prefix6 + "pci_address").(string)

		c.PortIndex = obj6

		var obj7 go_thunder.CPUHyperThread2
		prefix7 := "shared_poll_mode.0."
		obj7.Enable2 = d.Get(prefix7 + "enable").(int)
		obj7.Disable2 = d.Get(prefix7 + "disable").(int)

		c.Enable = obj7

		var obj8 go_thunder.AddPort
		prefix8 := "add_port.0."
		obj8.PortIndex = d.Get(prefix8 + "port_index").(int)

		c.PortIndex = obj8

		var obj9 go_thunder.AddPort1
		prefix9 := "del_port.0."
		obj9.PortIndex1 = d.Get(prefix9 + "port_index").(int)

		c.PortIndex = obj9

		var obj10 go_thunder.ModifyPort
		prefix10 := "modify_port.0."
		obj10.PortIndex = d.Get(prefix10 + "port_index").(int)
		obj10.PortNumber = d.Get(prefix10 + "port_number").(int)

		c.PortIndex = obj10

		var obj11 go_thunder.DeepHrxq2
		prefix11 := "tls_1_3_mgmt.0."
		obj11.DeepHrxq2Enable2 = d.Get(prefix11 + "enable").(int)

		c.Enable = obj11

		var obj12 go_thunder.DeepHrxq1
		prefix12 := "multi_queue_support.0."
		obj12.DeepHrxq1Enable1 = d.Get(prefix12 + "enable").(int)

		c.Enable = obj12

		var obj13 go_thunder.AddCpuCore
		prefix13 := "add_cpu_core.0."
		obj13.CoreIndex = d.Get(prefix13 + "core_index").(int)

		c.CoreIndex = obj13

		var obj14 go_thunder.AddCPUCore1
		prefix14 := "delete_cpu_core.0."
		obj14.CoreIndex1 = d.Get(prefix14 + "core_index").(int)

		c.CoreIndex = obj14

		var obj15 go_thunder.CpuHyperThread
		prefix15 := "cpu_hyper_thread.0."
		obj15.Enable = d.Get(prefix15 + "enable").(int)
		obj15.Disable = d.Get(prefix15 + "disable").(int)

		c.Enable = obj15

		var obj16 go_thunder.IoCpu
		prefix16 := "io_cpu.0."
		obj16.MaxCores = d.Get(prefix16 + "max_cores").(int)

		c.MaxCores = obj16

		var obj17 go_thunder.CPUHyperThread1
		prefix17 := "link_monitor.0."
		obj17.Enable1 = d.Get(prefix17 + "enable").(int)
		obj17.Disable1 = d.Get(prefix17 + "disable").(int)

		c.Enable = obj17

		// var obj18 go_thunder.PortList
		// prefix := "port_list.0."

		// c. = obj18

		// var obj19 go_thunder.PortInfo
		// prefix := "port_info.0."

		// c. = obj19

		// var obj20 go_thunder.InusePortList
		// prefix := "inuse_port_list.0."

		// c. = obj20

		// var obj21 go_thunder.CpuList
		// prefix := "cpu_list.0."

		// c. = obj21

		// var obj22 go_thunder.CpuMap
		// prefix := "cpu_map.0."

		// c. = obj22

		// var obj23 go_thunder.InuseCpuList
		// prefix := "inuse_cpu_list.0."

		// c. = obj23

		var obj24 go_thunder.SetRxtxDescSize
		prefix24 := "set_rxtx_desc_size.0."
		obj24.PortIndex = d.Get(prefix24 + "port_index").(int)
		obj24.RxdSize = d.Get(prefix24 + "rxd_size").(int)
		obj24.TxdSize = d.Get(prefix24 + "txd_size").(int)

		c.PortIndex = obj24

		var obj25 go_thunder.SetRxtxQueue
		prefix25 := "set_rxtx_queue.0."
		obj25.PortIndex = d.Get(prefix25 + "port_index").(int)
		obj25.RxqSize = d.Get(prefix25 + "rxq_size").(int)
		obj25.TxqSize = d.Get(prefix25 + "txq_size").(int)

		c.PortIndex = obj25

		var obj26 go_thunder.SystemTemplate
		prefix26 := "template.0."
		obj26.TemplatePolicy = d.Get(prefix26 + "template_policy").(string)

		c.TemplatePolicy = obj26

		var obj27 go_thunder.TemplateBind
		prefix27 := "template_bind.0."

		MonitorListCount := d.Get(prefix27 + "monitor_list.#").(int)
		obj27.Id = make([]go_thunder.MonitorList, 0, MonitorListCount)

		for i := 0; i < MonitorListCount; i++ {
			var obj27_1 go_thunder.MonitorList
			prefix27_1 := fmt.Sprintf(prefix27+"monitor_list.%d.", i)
			obj27_1.Id = d.Get(prefix27_1 + "id").(int)

			ClearCfgCount := d.Get(prefix27_1 + "clear_cfg.#").(int)
			obj27_1.Sessions = make([]go_thunder.ClearCfg, 0, ClearCfgCount)

			for i := 0; i < ClearCfgCount; i++ {
				var obj27_1_1 go_thunder.ClearCfg
				prefix27_1_1 := fmt.Sprintf(prefix27_1+"clear_cfg.%d.", i)
				obj27_1_1.Sessions = d.Get(prefix27_1_1 + "sessions").(string)
				obj27_1_1.ClearAllSequence = d.Get(prefix27_1_1 + "clear_all_sequence").(int)
				obj27_1_1.ClearSequence = d.Get(prefix27_1_1 + "clear_sequence").(int)
				obj27_1_1.Sessions = append(obj27_1.Sessions, obj27_1_1)
			}

			LinkDisableCfgCount := d.Get(prefix27_1 + "link_disable_cfg.#").(int)
			obj27_1.Diseth = make([]go_thunder.SystemLinkDisableCfg, 0, LinkDisableCfgCount)

			for i := 0; i < LinkDisableCfgCount; i++ {
				var obj27_1_2 go_thunder.SystemLinkDisableCfg
				prefix27_1_2 := fmt.Sprintf(prefix27_1_2+"link_disable_cfg.%d.", i)
				obj27_1_2.Diseth = d.Get(prefix27_1_2 + "diseth").(int)
				obj27_1_2.DisSequence = d.Get(prefix27_1_2 + "dis_sequence").(int)
				obj27_1.Diseth = append(obj27_1.Diseth, obj27_1_2)
			}

			LinkEnableCfgCount := d.Get(prefix27_1 + "link_enable_cfg.#").(int)
			obj27_1.Enaeth = make([]go_thunder.SystemLinkEnableCfg, 0, LinkEnableCfgCount)

			for i := 0; i < LinkEnableCfgCount; i++ {
				var obj27_1_3 go_thunder.SystemLinkEnableCfg
				prefix27_1_3 := fmt.Sprintf(prefix27_1_3+"link_enable_cfg.%d.", i)
				obj27_1_3.Enaeth = d.Get(prefix27_1_3 + "enaeth").(int)
				obj27_1_3.EnaSequence = d.Get(prefix27_1_3 + "ena_sequence").(int)
				obj27_1.Enaeth = append(obj27_1.Enaeth, obj27_1_3)
			}

			obj27_1.MonitorRelation = d.Get(prefix27_1 + "monitor_relation").(string)

			LinkUpCfgCount := d.Get(prefix27_1 + "link_up_cfg.#").(int)
			obj27_1.LinkupEthernet1 = make([]go_thunder.SystemLinkUpCfg, 0, LinkUpCfgCount)

			for i := 0; i < LinkUpCfgCount; i++ {
				var obj27_1_4 go_thunder.SystemLinkUpCfg
				prefix27_1_4 := fmt.Sprintf(prefix27_1+"link_up_cfg.%d.", i)
				obj27_1_4.LinkupEthernet1 = d.Get(prefix27_1_4 + "linkup_ethernet1").(int)
				obj27_1_4.LinkUpSequence1 = d.Get(prefix27_1_4 + "link_up_sequence1").(int)
				obj27_1_4.LinkupEthernet2 = d.Get(prefix27_1_4 + "linkup_ethernet2").(int)
				obj27_1_4.LinkUpSequence2 = d.Get(prefix27_1_4 + "link_up_sequence2").(int)
				obj27_1_4.LinkupEthernet3 = d.Get(prefix27_1_4 + "linkup_ethernet3").(int)
				obj27_1_4.LinkUpSequence3 = d.Get(prefix27_1_4 + "link_up_sequence3").(int)
				obj27_1.LinkupEthernet1 = append(obj27_1.LinkupEthernet1, obj27_1_4)
			}

			LinkDownCfgCount := d.Get(prefix27_1 + "link_down_cfg.#").(int)
			obj27_1.LinkdownEthernet1 = make([]go_thunder.SystemLinkDownCfg, 0, LinkDownCfgCount)

			for i := 0; i < LinkDownCfgCount; i++ {
				var obj27_1_5 go_thunder.SystemLinkDownCfg
				prefix27_1_5 := fmt.Sprintf(prefix27_1+"link_down_cfg.%d.", i)
				obj27_1_5.LinkdownEthernet1 = d.Get(prefix27_1_5 + "linkdown_ethernet1").(int)
				obj27_1_5.LinkDownSequence1 = d.Get(prefix27_1_5 + "link_down_sequence1").(int)
				obj27_1_5.LinkdownEthernet2 = d.Get(prefix27_1_5 + "linkdown_ethernet2").(int)
				obj27_1_5.LinkDownSequence2 = d.Get(prefix27_1_5 + "link_down_sequence2").(int)
				obj27_1_5.LinkdownEthernet3 = d.Get(prefix27_1_5 + "linkdown_ethernet3").(int)
				obj27_1_5.LinkDownSequence3 = d.Get(prefix27_1_5 + "link_down_sequence3").(int)
				obj27_1.LinkdownEthernet1 = append(obj27_1.LinkdownEthernet1, obj27_1_5)
			}

			obj27_1.UserTag = d.Get(prefix27_1 + "user_tag").(string)
			obj27.Id = append(obj27.Id, obj27_1)
		}

		c.MonitorList = obj27

		var obj28 go_thunder.MonTemplate
		prefix28 := "mon_template.0."

		MonitorListCount = d.Get(prefix28 + "monitor_list.#").(int)
		obj28.Id = make([]go_thunder.MonitorList, 0, MonitorListCount)

		for i := 0; i < MonitorListCount; i++ {
			var obj28_1 go_thunder.MonitorList
			prefix28_1 := fmt.Sprintf(prefix28+"monitor_list.%d.", i)
			obj28_1.Id = d.Get(prefix28_1 + "id").(int)

			ClearCfgCount := d.Get(prefix28_1 + "clear_cfg.#").(int)
			obj28_1.Sessions = make([]go_thunder.SystemClearCfg, 0, ClearCfgCount)

			for i := 0; i < ClearCfgCount; i++ {
				var obj28_1_1 go_thunder.SystemClearCfg
				prefix28_1_1 := fmt.Sprintf(prefix28_1_1+"clear_cfg.%d.", i)
				obj28_1_1.Sessions = d.Get(prefix28_1_1 + "sessions").(string)
				obj28_1_1.ClearAllSequence = d.Get(prefix28_1_1 + "clear_all_sequence").(int)
				obj28_1_1.ClearSequence = d.Get(prefix28_1_1 + "clear_sequence").(int)
				obj28_1.Sessions = append(obj28_1.Sessions, obj28_1_1)
			}

			LinkDisableCfgCount := d.Get(prefix28_1 + "link_disable_cfg.#").(int)
			obj28_1.Diseth = make([]go_thunder.SystemLinkDisableCfg, 0, LinkDisableCfgCount)

			for i := 0; i < LinkDisableCfgCount; i++ {
				var obj28_1_2 go_thunder.SystemLinkDisableCfg
				prefix28_1_2 := fmt.Sprintf(prefix28_1_2+"link_disable_cfg.%d.", i)
				obj28_1_2.Diseth = d.Get(prefix28_1_2 + "diseth").(int)
				obj28_1_2.DisSequence = d.Get(prefix28_1_2 + "dis_sequence").(int)
				obj28_1.Diseth = append(obj28_1.Diseth, obj28_1_2)
			}

			LinkEnableCfgCount := d.Get(prefix28_1 + "link_enable_cfg.#").(int)
			obj28_1.Enaeth = make([]go_thunder.SystemLinkEnableCfg, 0, LinkEnableCfgCount)

			for i := 0; i < LinkEnableCfgCount; i++ {
				var obj28_1_3 go_thunder.SystemLinkEnableCfg
				prefix28_1_3 := fmt.Sprintf(prefix28_1+"link_enable_cfg.%d.", i)
				obj28_1_3.Enaeth = d.Get(prefix28_1_3 + "enaeth").(int)
				obj28_1_3.EnaSequence = d.Get(prefix28_1_3 + "ena_sequence").(int)
				obj28_1.Enaeth = append(obj28_1.Enaeth, obj28_1_3)
			}

			obj28_1.MonitorRelation = d.Get(prefix28_1 + "monitor_relation").(string)

			LinkUpCfgCount := d.Get(prefix28_1 + "link_up_cfg.#").(int)
			obj28_1.LinkupEthernet1 = make([]go_thunder.SystemLinkUpCfg, 0, LinkUpCfgCount)

			for i := 0; i < LinkUpCfgCount; i++ {
				var obj28_1_4 go_thunder.SystemLinkUpCfg
				prefix28_1_4 := fmt.Sprintf(prefix28_1+"link_up_cfg.%d.", i)
				obj28_1_4.LinkupEthernet1 = d.Get(prefix28_1_4 + "linkup_ethernet1").(int)
				obj28_1_4.LinkUpSequence1 = d.Get(prefix28_1_4 + "link_up_sequence1").(int)
				obj28_1_4.LinkupEthernet2 = d.Get(prefix28_1_4 + "linkup_ethernet2").(int)
				obj28_1_4.LinkUpSequence2 = d.Get(prefix28_1_4 + "link_up_sequence2").(int)
				obj28_1_4.LinkupEthernet3 = d.Get(prefix28_1_4 + "linkup_ethernet3").(int)
				obj28_1_4.LinkUpSequence3 = d.Get(prefix28_1_4 + "link_up_sequence3").(int)
				obj28_1.LinkupEthernet1 = append(obj1.LinkupEthernet1, obj28_1_4)
			}

			LinkDownCfgCount := d.Get(prefix28_1 + "link_down_cfg.#").(int)
			obj28_1.LinkdownEthernet1 = make([]go_thunder.SystemLinkDownCfg, 0, LinkDownCfgCount)

			for i := 0; i < LinkDownCfgCount; i++ {
				var obj28_1_5 go_thunder.SystemLinkDownCfg
				prefix28_1_5 := fmt.Sprintf(prefix28_1+"link_down_cfg.%d.", i)
				obj28_1_5.LinkdownEthernet1 = d.Get(prefix28_1_5 + "linkdown_ethernet1").(int)
				obj28_1_5.LinkDownSequence1 = d.Get(prefix28_1_5 + "link_down_sequence1").(int)
				obj28_1_5.LinkdownEthernet2 = d.Get(prefix28_1_5 + "linkdown_ethernet2").(int)
				obj28_1_5.LinkDownSequence2 = d.Get(prefix28_1_5 + "link_down_sequence2").(int)
				obj28_1_5.LinkdownEthernet3 = d.Get(prefix28_1_5 + "linkdown_ethernet3").(int)
				obj28_1_5.LinkDownSequence3 = d.Get(prefix28_1_5 + "link_down_sequence3").(int)
				obj28_1.LinkdownEthernet1 = append(obj28_1.LinkdownEthernet1, obj28_1_5)
			}

			obj28_1.UserTag = d.Get(prefix28_1 + "user_tag").(string)
			obj28.Id = append(obj28.Id, obj28_1)
		}

		var obj28_2 go_thunder.LinkBlockAsDown
		prefix28_2 := prefix28 + "link_block_as_down.0."
		obj28_2.Enable = d.Get(prefix28_2 + "enable").(int)

		obj28.Enable = obj28_2

		c.MonitorList = obj28

		var obj29 go_thunder.Memory
		prefix29 := "memory.0."

		SamplingEnableCount := d.Get(prefix29 + "sampling_enable.#").(int)
		obj29.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		for i := 0; i < SamplingEnableCount; i++ {
			var obj29_1 go_thunder.SamplingEnable
			prefix29_1 := fmt.Sprintf(prefix29_1+"sampling_enable.%d.", i)
			obj29_1.Counters1 = d.Get(prefix29_1 + "counters1").(string)
			obj29.Counters1 = append(obj29.Counters1, obj29_1)
		}

		c.SamplingEnable = obj29

		var obj30 go_thunder.SystemResourceUsage
		prefix30 := "resource_usage.0."
		obj30.SslContextMemory = d.Get(prefix30 + "ssl_context_memory").(int)
		obj30.SslDmaMemory = d.Get(prefix30 + "ssl_dma_memory").(int)
		obj30.NatPoolAddrCount = d.Get(prefix30 + "nat_pool_addr_count").(int)
		obj30.L4SessionCount = d.Get(prefix30 + "l4_session_count").(int)
		obj30.AuthPortalHtmlFileSize = d.Get(prefix30 + "auth_portal_html_file_size").(int)
		obj30.AuthPortalImageFileSize = d.Get(prefix30 + "auth_portal_image_file_size").(int)
		obj30.MaxAflexFileSize = d.Get(prefix30 + "max_aflex_file_size").(int)
		obj30.AflexTableEntryCount = d.Get(prefix30 + "aflex_table_entry_count").(int)
		obj30.ClassListIpv6AddrCount = d.Get(prefix30 + "class_list_ipv6_addr_count").(int)
		obj30.ClassListAcEntryCount = d.Get(prefix30 + "class_list_ac_entry_count").(int)
		obj30.MaxAflexAuthzCollectionNumber = d.Get(prefix30 + "max_aflex_authz_collection_number").(int)
		obj30.RadiusTableSize = d.Get(prefix30 + "radius_table_size").(int)
		obj30.AuthzPolicyNumber = d.Get(prefix30 + "authz_policy_number").(int)
		obj30.IpsecSaNumber = d.Get(prefix30 + "ipsec_sa_number").(int)
		obj30.RamCacheMemoryLimit = d.Get(prefix30 + "ram_cache_memory_limit").(int)

		var obj1 go_thunder.Visibility
		prefix30_1 := prefix30 + "visibility.0."
		obj30_1.MonitoredEntityCount = d.Get(prefix30_1 + "monitored_entity_count").(int)

		obj30.MonitoredEntityCount = obj30_1

		c.SslContextMemory = obj30

		var obj31 go_thunder.LinkCapability
		prefix31 := "link_capability.0."
		obj31.Enable1 = d.Get(prefix31 + "enable").(int)

		c.Enable = obj31

		var obj32 go_thunder.ResourceAccounting
		prefix32 := "resource_accounting.0."

		TemplateListCount := d.Get(prefix32 + "template_list.#").(int)
		obj32.Name = make([]go_thunder.TemplateList, 0, TemplateListCount)

		for i := 0; i < TemplateListCount; i++ {
			var obj32_1 go_thunder.TemplateList
			prefix32_1 := fmt.Sprintf(prefix32+"template_list.%d.", i)
			obj32_1.Name = d.Get(prefix32 + "name").(string)
			obj32_1.UserTag = d.Get(prefix32 + "user_tag").(string)

			var obj32_1_1 go_thunder.AppResources
			prefix32_1_1 := prefix32_1 + "app_resources.0."

			var obj32_1_1_1 go_thunder.GslbDeviceCfg
			prefix32_1_1_1 := prefix32_1_1 + "gslb_device_cfg.0."
			obj32_1_1_1.GslbDeviceMax = d.Get(prefix32_1_1_1 + "gslb_device_max").(int)
			obj32_1_1_1.GslbDeviceMinGuarantee = d.Get(prefix32_1_1_1 + "gslb_device_min_guarantee").(int)

			obj32_1_1.GslbDeviceMax = obj32_1_1_1

			var obj32_1_1_2 go_thunder.GslbGeoLocationCfg
			prefix32_1_1_2 := prefix32_1_1 + "gslb_geo_location_cfg.0."
			obj32_1_1_2.GslbGeoLocationMax = d.Get(prefix32_1_1_2 + "gslb_geo_location_max").(int)
			obj32_1_1_2.GslbGeoLocationMinGuarantee = d.Get(prefix32_1_1_2 + "gslb_geo_location_min_guarantee").(int)

			obj1.GslbGeoLocationMax = obj32_1_1_2

			var obj32_1_1_3 go_thunder.GslbIpListCfg
			prefix32_1_1_3 := prefix32_1_1 + "gslb_ip_list_cfg.0."
			obj32_1_1_3.GslbIpListMax = d.Get(prefix32_1_1_3 + "gslb_ip_list_max").(int)
			obj32_1_1_3.GslbIpListMinGuarantee = d.Get(prefix32_1_1_3 + "gslb_ip_list_min_guarantee").(int)

			obj32_1_1.GslbIpListMax = obj32_1_1_3

			var obj32_1_1_4 go_thunder.GslbPolicyCfg
			prefix32_1_1_4 := prefix32_1_1 + "gslb_policy_cfg.0."
			obj32_1_1_4.GslbPolicyMax = d.Get(prefix32_1_1_4 + "gslb_policy_max").(int)
			obj32_1_1_4.GslbPolicyMinGuarantee = d.Get(prefix32_1_1_4 + "gslb_policy_min_guarantee").(int)

			obj32_1_1.GslbPolicyMax = obj32_1_1_4

			var obj32_1_1_5 go_thunder.GslbServiceCfg
			prefix32_1_1_5 := prefix32_1_1 + "gslb_service_cfg.0."
			obj32_1_1_5.GslbServiceMax = d.Get(prefix32_1_1_5 + "gslb_service_max").(int)
			obj32_1_1_5.GslbServiceMinGuarantee = d.Get(prefix32_1_1_5 + "gslb_service_min_guarantee").(int)

			obj32_1_1.GslbServiceMax = obj32_1_1_5

			var obj32_1_1_6 go_thunder.GslbServiceIpCfg
			prefix32_1_1_6 := prefix32_1_1 + "gslb_service_ip_cfg.0."
			obj32_1_1_6.GslbServiceIpMax = d.Get(prefix32_1_1_6 + "gslb_service_ip_max").(int)
			obj32_1_1_6.GslbServiceIpMinGuarantee = d.Get(prefix32_1_1_6 + "gslb_service_ip_min_guarantee").(int)

			obj32_1_1.GslbServiceIpMax = obj32_1_1_6

			var obj32_1_1_7 go_thunder.GslbServicePortCfg
			prefix32_1_1_7 := prefix32_1_1 + "gslb_service_port_cfg.0."
			obj32_1_1_7.GslbServicePortMax = d.Get(prefix32_1_1_7 + "gslb_service_port_max").(int)
			obj32_1_1_7.GslbServicePortMinGuarantee = d.Get(prefix32_1_1_7 + "gslb_service_port_min_guarantee").(int)

			obj32_1_1.GslbServicePortMax = obj32_1_1_7

			var obj32_1_1_8 go_thunder.GslbSiteCfg
			prefix32_1_1_8 := prefix32_1_1 + "gslb_site_cfg.0."
			obj32_1_1_8.GslbSiteMax = d.Get(prefix32_1_1_8 + "gslb_site_max").(int)
			obj32_1_1_8.GslbSiteMinGuarantee = d.Get(prefix32_1_1_8 + "gslb_site_min_guarantee").(int)

			obj32_1_1.GslbSiteMax = obj32_1_1_8

			var obj32_1_1_9 go_thunder.GslbSvcGroupCfg
			prefix32_1_1_9 := prefix32_1_1 + "gslb_svc_group_cfg.0."
			obj32_1_1_9.GslbSvcGroupMax = d.Get(prefix32_1_1_9 + "gslb_svc_group_max").(int)
			obj32_1_1_9.GslbSvcGroupMinGuarantee = d.Get(prefix32_1_1_9 + "gslb_svc_group_min_guarantee").(int)

			obj32_1_1.GslbSvcGroupMax = obj32_1_1_9

			var obj32_1_1_10 go_thunder.GslbTemplateCfg
			prefix32_1_1_10 := prefix32_1_1 + "gslb_template_cfg.0."
			obj32_1_1_10.GslbTemplateMax = d.Get(prefix32_1_1_10 + "gslb_template_max").(int)
			obj32_1_1_10.GslbTemplateMinGuarantee = d.Get(prefix32_1_1_10 + "gslb_template_min_guarantee").(int)

			obj32_1_1.GslbTemplateMax = obj32_1_1_10

			var obj32_1_1_11 go_thunder.GslbZoneCfg
			prefix32_1_1_11 := prefix32_1_1 + "gslb_zone_cfg.0."
			obj32_1_1_11.GslbZoneMax = d.Get(obj32_1_1_11 + "gslb_zone_max").(int)
			obj32_1_1_11.GslbZoneMinGuarantee = d.Get(obj32_1_1_11 + "gslb_zone_min_guarantee").(int)

			obj32_1_1.GslbZoneMax = obj32_1_1_11

			var obj32_1_1_12 go_thunder.HealthMonitorCfg
			prefix32_1_1_12 := prefix32_1_1 + "health_monitor_cfg.0."
			obj32_1_1_12.HealthMonitorMax = d.Get(prefix32_1_1_12 + "health_monitor_max").(int)
			obj32_1_1_12.HealthMonitorMinGuarantee = d.Get(prefix32_1_1_12 + "health_monitor_min_guarantee").(int)

			obj32_1_1.HealthMonitorMax = obj32_1_1_12

			var obj32_1_1_13 go_thunder.RealPortCfg
			prefix32_1_1_13 := prefix32_1_1 + "real_port_cfg.0."
			obj32_1_1_13.RealPortMax = d.Get(prefix32_1_1_13 + "real_port_max").(int)
			obj32_1_1_13.RealPortMinGuarantee = d.Get(prefix32_1_1_13 + "real_port_min_guarantee").(int)

			obj32_1_1.RealPortMax = obj32_1_1_13

			var obj32_1_1_14 go_thunder.RealServerCfg
			prefix32_1_1_14 := prefix32_1_1 + "real_server_cfg.0."
			obj32_1_1_14.RealServerMax = d.Get(prefix32_1_1_14 + "real_server_max").(int)
			obj32_1_1_14.RealServerMinGuarantee = d.Get(prefix32_1_1_14 + "real_server_min_guarantee").(int)

			obj32_1_1.RealServerMax = obj32_1_1_14

			var obj32_1_1_15 go_thunder.ServiceGroupCfg
			prefix32_1_1_15 := prefix32_1_1 + "service_group_cfg.0."
			obj32_1_1_15.ServiceGroupMax = d.Get(prefix32_1_1_15 + "service_group_max").(int)
			obj32_1_1_15.ServiceGroupMinGuarantee = d.Get(prefix32_1_1_15 + "service_group_min_guarantee").(int)

			obj32_1_1.ServiceGroupMax = obj32_1_1_15

			var obj32_1_1_16 go_thunder.VirtualServerCfg
			prefix32_1_1_16 := prefix32_1_1 + "virtual_server_cfg.0."
			obj32_1_1_16.VirtualServerMax = d.Get(prefix32_1_1_16 + "virtual_server_max").(int)
			obj32_1_1_16.VirtualServerMinGuarantee = d.Get(prefix32_1_1_16 + "virtual_server_min_guarantee").(int)

			obj32_1_1.VirtualServerMax = obj32_1_1_16

			var obj32_1_1_17 go_thunder.VirtualPortCfg
			prefix32_1_1_17 := prefix32_1_1 + "virtual_port_cfg.0."
			obj32_1_1_17.VirtualPortMax = d.Get(prefix32_1_1_17 + "virtual_port_max").(int)
			obj32_1_1_17.VirtualPortMinGuarantee = d.Get(prefix32_1_1_17 + "virtual_port_min_guarantee").(int)

			obj32_1_1.VirtualPortMax = obj32_1_1_17

			var obj32_1_1_18 go_thunder.CacheTemplateCfg
			prefix32_1_1_18 := prefix32_1_1 + "cache_template_cfg.0."
			obj32_1_1_18.CacheTemplateMax = d.Get(prefix32_1_1_18 + "cache_template_max").(int)
			obj32_1_1_18.CacheTemplateMinGuarantee = d.Get(prefix32_1_1_18 + "cache_template_min_guarantee").(int)

			obj32_1_1.CacheTemplateMax = obj32_1_1_18

			var obj32_1_1_19 go_thunder.ClientSslTemplateCfg
			prefix32_1_1_19 := prefix32_1_1 + "client_ssl_template_cfg.0."
			obj32_1_1_19.ClientSslTemplateMax = d.Get(prefix32_1_1_19 + "client_ssl_template_max").(int)
			obj32_1_1_19.ClientSslTemplateMinGuarantee = d.Get(prefix32_1_1_19 + "client_ssl_template_min_guarantee").(int)

			obj32_1_1.ClientSslTemplateMax = obj32_1_1_19

			var obj32_1_1_20 go_thunder.ConnReuseTemplateCfg
			prefix32_1_1_19 := prefix32_1_1 + "conn_reuse_template_cfg.0."
			obj32_1_1_20.ConnReuseTemplateMax = d.Get(prefix32_1_1_19 + "conn_reuse_template_max").(int)
			obj32_1_1_20.ConnReuseTemplateMinGuarantee = d.Get(prefix32_1_1_19 + "conn_reuse_template_min_guarantee").(int)

			obj32_1_1.ConnReuseTemplateMax = obj32_1_1_20

			var obj32_1_1_21 go_thunder.FastTcpTemplateCfg
			prefix32_1_1_21 := prefix32_1_1 + "fast_tcp_template_cfg.0."
			obj32_1_1_21.FastTcpTemplateMax = d.Get(prefix32_1_1_21 + "fast_tcp_template_max").(int)
			obj32_1_1_21.FastTcpTemplateMinGuarantee = d.Get(prefix32_1_1_21 + "fast_tcp_template_min_guarantee").(int)

			obj32_1_1.FastTcpTemplateMax = obj32_1_1_21

			var obj32_1_1_22 go_thunder.FastUdpTemplateCfg
			prefix32_1_1_22 := prefix32_1_1 + "fast_udp_template_cfg.0."
			obj32_1_1_22.FastUdpTemplateMax = d.Get(prefix32_1_1_22 + "fast_udp_template_max").(int)
			obj32_1_1_22.FastUdpTemplateMinGuarantee = d.Get(prefix32_1_1_22 + "fast_udp_template_min_guarantee").(int)

			obj32_1_1.FastUdpTemplateMax = obj32_1_1_22

			var obj32_1_1_23 go_thunder.FixTemplateCfg
			prefix32_1_1_23 := prefix32_1_1 + "fix_template_cfg.0."
			obj32_1_1_23.FixTemplateMax = d.Get(prefix32_1_1_23 + "fix_template_max").(int)
			obj32_1_1_23.FixTemplateMinGuarantee = d.Get(prefix32_1_1_23 + "fix_template_min_guarantee").(int)

			obj32_1_1.FixTemplateMax = obj32_1_1_23

			var obj32_1_1_24 go_thunder.HttpTemplateCfg
			prefix32_1_1_24 := prefix32_1_1 + "http_template_cfg.0."
			obj32_1_1_24.HttpTemplateMax = d.Get(prefix32_1_1_24 + "http_template_max").(int)
			obj32_1_1_24.HttpTemplateMinGuarantee = d.Get(prefix32_1_1_24 + "http_template_min_guarantee").(int)

			obj32_1_1.HttpTemplateMax = obj32_1_1_24

			var obj32_1_1_25 go_thunder.LinkCostTemplateCfg
			prefix32_1_1_25 := prefix32_1_1 + "link_cost_template_cfg.0."
			obj32_1_1_25.LinkCostTemplateMax = d.Get(prefix32_1_1_25 + "link_cost_template_max").(int)
			obj32_1_1_25.LinkCostTemplateMinGuarantee = d.Get(prefix32_1_1_25 + "link_cost_template_min_guarantee").(int)

			obj32_1_1.LinkCostTemplateMax = obj32_1_1_25

			var obj32_1_1_26 go_thunder.PersistCookieTemplateCfg
			prefix32_1_1_26 := prefix32_1_1 + "persist_cookie_template_cfg.0."
			obj32_1_1_26.PersistCookieTemplateMax = d.Get(prefix32_1_1_26 + "persist_cookie_template_max").(int)
			obj32_1_1_26.PersistCookieTemplateMinGuarantee = d.Get(prefix32_1_1_26 + "persist_cookie_template_min_guarantee").(int)

			obj32_1_1.PersistCookieTemplateMax = obj32_1_1_26

			var obj32_1_1_27 go_thunder.PersistSrcipTemplateCfg
			prefix32_1_1_27 := prefix32_1_1 + "persist_srcip_template_cfg.0."
			obj32_1_1_27.PersistSrcipTemplateMax = d.Get(prefix32_1_1_27 + "persist_srcip_template_max").(int)
			obj32_1_1_27.PersistSrcipTemplateMinGuarantee = d.Get(prefix32_1_1_27 + "persist_srcip_template_min_guarantee").(int)

			obj32_1_1.PersistSrcipTemplateMax = obj32_1_1_27

			var obj32_1_1_28 go_thunder.ServerSslTemplateCfg
			prefix32_1_1_28 := prefix32_1_1 + "server_ssl_template_cfg.0."
			obj32_1_1_28.ServerSslTemplateMax = d.Get(prefix32_1_1_28 + "server_ssl_template_max").(int)
			obj32_1_1_28.ServerSslTemplateMinGuarantee = d.Get(prefix32_1_1_28 + "server_ssl_template_min_guarantee").(int)

			obj32_1_1.ServerSslTemplateMax = obj32_1_1_28

			var obj32_1_1_29 go_thunder.ProxyTemplateCfg
			prefix32_1_1_29 := prefix32_1_1 + "proxy_template_cfg.0."
			obj32_1_1_29.ProxyTemplateMax = d.Get(prefix32_1_1_29 + "proxy_template_max").(int)
			obj32_1_1_29.ProxyTemplateMinGuarantee = d.Get(prefix32_1_1_29 + "proxy_template_min_guarantee").(int)

			obj32_1_1.ProxyTemplateMax = obj32_1_1_29

			var obj32_1_1_30 go_thunder.StreamTemplateCfg
			prefix32_1_1_30 := prefix32_1_1 + "stream_template_cfg.0."
			obj32_1_1_30.StreamTemplateMax = d.Get(prefix32_1_1_30 + "stream_template_max").(int)
			obj32_1_1_30.StreamTemplateMinGuarantee = d.Get(prefix32_1_1_30 + "stream_template_min_guarantee").(int)

			obj32_1_1.StreamTemplateMax = obj32_1_1_30

			obj32_1_1.Threshold = d.Get(prefix32_1_1 + "threshold").(int)

			obj32_1.GslbDeviceCfg = obj32_1_1

			var obj32_1_2 go_thunder.NetworkResources
			prefix32_1_2 := prefix32_1 + "network_resources.0."

			var obj32_1_2_1 go_thunder.StaticIpv4RouteCfg
			prefix32_1_2_1 := prefix32_1_2 + "static_ipv4_route_cfg.0."
			obj32_1_2_1.StaticIpv4RouteMax = d.Get(prefix32_1_2_1 + "static_ipv4_route_max").(int)
			obj32_1_2_1.StaticIpv4RouteMinGuarantee = d.Get(prefix32_1_2_1 + "static_ipv4_route_min_guarantee").(int)

			obj32_1_2.StaticIpv4RouteMax = obj32_1_2_1

			var obj32_1_2_2 go_thunder.StaticIpv6RouteCfg
			prefix32_1_2_2 := prefix32_1_2 + "static_ipv6_route_cfg.0."
			obj32_1_2_2.StaticIpv6RouteMax = d.Get(prefix32_1_2_2 + "static_ipv6_route_max").(int)
			obj32_1_2_2.StaticIpv6RouteMinGuarantee = d.Get(prefix32_1_2_2 + "static_ipv6_route_min_guarantee").(int)

			obj32_1_2.StaticIpv6RouteMax = obj32_1_2_2

			var obj32_1_2_3 go_thunder.Ipv4AclLineCfg
			prefix32_1_2_3 := prefix32_1_2 + "ipv4_acl_line_cfg.0."
			obj32_1_2_3.Ipv4AclLineMax = d.Get(prefix32_1_2_3 + "ipv4_acl_line_max").(int)
			obj32_1_2_3.Ipv4AclLineMinGuarantee = d.Get(prefix32_1_2_3 + "ipv4_acl_line_min_guarantee").(int)

			obj32_1_2.Ipv4AclLineMax = obj32_1_2_3

			var obj32_1_2_4 go_thunder.Ipv6AclLineCfg
			prefix32_1_2_4 := prefix32_1_2 + "ipv6_acl_line_cfg.0."
			obj32_1_2_4.Ipv6AclLineMax = d.Get(prefix32_1_2_4 + "ipv6_acl_line_max").(int)
			obj32_1_2_4.Ipv6AclLineMinGuarantee = d.Get(prefix32_1_2_4 + "ipv6_acl_line_min_guarantee").(int)

			obj32_1_2.Ipv6AclLineMax = obj32_1_2_4

			var obj32_1_2_5 go_thunder.StaticArpCfg
			prefix32_1_2_5 := prefix32_1_2 + "static_arp_cfg.0."
			obj32_1_2_5.StaticArpMax = d.Get(prefix32_1_2_5 + "static_arp_max").(int)
			obj32_1_2_5.StaticArpMinGuarantee = d.Get(prefix32_1_2_5 + "static_arp_min_guarantee").(int)

			obj32_1_2.StaticArpMax = obj32_1_2_5

			var obj32_1_2_6 go_thunder.StaticNeighborCfg
			prefix32_1_2_6 := prefix32_1_2 + "static_neighbor_cfg.0."
			obj32_1_2_6.StaticNeighborMax = d.Get(prefix32_1_2_6 + "static_neighbor_max").(int)
			obj32_1_2_6.StaticNeighborMinGuarantee = d.Get(prefix32_1_2_6 + "static_neighbor_min_guarantee").(int)

			obj32_1_2_6.StaticNeighborMax = obj32_1_2_6

			var obj32_1_2_7 go_thunder.StaticMacCfg
			prefix32_1_2_7 := prefix32_1_2 + "static_mac_cfg.0."
			obj32_1_2_7.StaticMacMax = d.Get(prefix32_1_2_7 + "static_mac_max").(int)
			obj32_1_2_7.StaticMacMinGuarantee = d.Get(prefix32_1_2_7 + "static_mac_min_guarantee").(int)

			obj32_1_2.StaticMacMax = obj32_1_2_7

			var obj32_1_2_8 go_thunder.ObjectGroupCfg
			prefix32_1_2_8 := prefix32_1_2 + "object_group_cfg.0."
			obj32_1_2_8.ObjectGroupMax = d.Get(prefix32_1_2_8 + "object_group_max").(int)
			obj32_1_2_8.ObjectGroupMinGuarantee = d.Get(prefix32_1_2_8 + "object_group_min_guarantee").(int)

			obj32_1_2.ObjectGroupMax = obj32_1_2_8

			var obj32_1_2_9 go_thunder.ObjectGroupClauseCfg
			prefix32_1_2_9 := prefix32_1_2 + "object_group_clause_cfg.0."
			obj32_1_2_9.ObjectGroupClauseMax = d.Get(prefix32_1_2_9 + "object_group_clause_max").(int)
			obj32_1_2_9.ObjectGroupClauseMinGuarantee = d.Get(prefix32_1_2_9 + "object_group_clause_min_guarantee").(int)

			obj32_1_2.ObjectGroupClauseMax = obj32_1_2_9

			obj32_1_2.Threshold = d.Get(prefix32_1_2 + "threshold").(int)

			obj32_1.StaticIpv4RouteCfg = obj32_1_2

			var obj32_1_3 go_thunder.SystemResources
			prefix32_1_3 := prefix32_1 + "system_resources.0."

			var obj32_1_3_1 go_thunder.BwLimitCfg
			prefix32_1_3_1 := prefix32_1_3 + "bw_limit_cfg.0."
			obj32_1_3_1.BwLimitMax = d.Get(prefix32_1_3_1 + "bw_limit_max").(int)
			obj32_1_3_1.BwLimitWatermarkDisable = d.Get(prefix32_1_3_1 + "bw_limit_watermark_disable").(int)

			obj32_1_3.BwLimitMax = obj32_1_3_1

			var obj32_1_3_2 go_thunder.ConcurrentSessionLimitCfg
			prefix32_1_3_2 := prefix32_1_3 + "concurrent_session_limit_cfg.0."
			obj32_1_3_2.ConcurrentSessionLimitMax = d.Get(prefix32_1_3_2 + "concurrent_session_limit_max").(int)

			obj32_1_3.ConcurrentSessionLimitMax = obj32_1_3_2

			var obj32_1_3_3 go_thunder.L4SessionLimitCfg
			prefix32_1_3_3 := prefix32_1_3 + "l4_session_limit_cfg.0."
			obj32_1_3_3.L4SessionLimitMax = d.Get(prefix32_1_3_3 + "l4_session_limit_max").(string)
			obj32_1_3_3.L4SessionLimitMinGuarantee = d.Get(prefix32_1_3_3 + "l4_session_limit_min_guarantee").(string)

			obj32_1_3.L4SessionLimitMax = obj32_1_3_3

			var obj32_1_3_4 go_thunder.L4cpsLimitCfg
			prefix32_1_3_4 := prefix32_1_3 + "l4cps_limit_cfg.0."
			obj32_1_3_4.L4cpsLimitMax = d.Get(prefix32_1_3_4 + "l4cps_limit_max").(int)

			obj32_1_3.L4cpsLimitMax = obj32_1_3_4

			var obj32_1_3_5 go_thunder.L7cpsLimitCfg
			prefix32_1_3_5 := prefix32_1_3 + "l7cps_limit_cfg.0."
			obj32_1_3_5.L7cpsLimitMax = d.Get(prefix32_1_3_5 + "l7cps_limit_max").(int)

			obj32_1_3.L7cpsLimitMax = obj32_1_3_5

			var obj32_1_3_6 go_thunder.NatcpsLimitCfg
			prefix32_1_3_6 := prefix32_1_3 + "natcps_limit_cfg.0."
			obj32_1_3_6.NatcpsLimitMax = d.Get(prefix32_1_3_6 + "natcps_limit_max").(int)

			obj32_1_3.NatcpsLimitMax = obj32_1_3_6

			var obj32_1_3_7 go_thunder.FwcpsLimitCfg
			prefix32_1_3_7 := prefix32_1_3 + "fwcps_limit_cfg.0."
			obj32_1_3_7.FwcpsLimitMax = d.Get(prefix32_1_3_7 + "fwcps_limit_max").(int)

			obj32_1_3.FwcpsLimitMax = obj32_1_3_7

			var obj32_1_3_8 go_thunder.SslThroughputLimitCfg
			prefix32_1_3_8 := prefix32_1_3 + "ssl_throughput_limit_cfg.0."
			obj32_1_3_8.SslThroughputLimitMax = d.Get(prefix32_1_3_8 + "ssl_throughput_limit_max").(int)
			obj32_1_3_8.SslThroughputLimitWatermarkDisable = d.Get(prefix32_1_3_8 + "ssl_throughput_limit_watermark_disable").(int)

			obj32_1_3.SslThroughputLimitMax = obj32_1_3_8

			var obj32_1_3_9 go_thunder.SslcpsLimitCfg
			prefix32_1_3_9 := prefix32_1_3 + "sslcps_limit_cfg.0."
			obj32_1_3_9.SslcpsLimitMax = d.Get(prefix32_1_3_9 + "sslcps_limit_max").(int)

			obj32_1_3.SslcpsLimitMax = obj32_1_3_9

			obj32_1_3.Threshold = d.Get(prefix32_1_3 + "threshold").(int)

			obj32_1.BwLimitCfg = obj32_1_3

			obj32.Name = append(obj32.Name, obj32_1)
		}

		c.TemplateList = obj32

		/*
		   var obj33 go_thunder.Trunk
		   prefix33 := "trunk.0."

		   var obj33_1 go_thunder.LoadBalance
		   prefix33_1 :=prefix33 + "load_balance.0."
		   obj33_1.UseL3 = d.Get(prefix33_1+"use_l3").(int)
		   obj33_1.UseL4 = d.Get(prefix33_1+"use_l4").(int)

		   obj33.UseL3 = obj33_1


		   c.LoadBalance = obj33


		   var obj34 go_thunder.Ports
		   prefix34 := "ports.0."
		   obj34.LinkDetectionInterval = d.Get(prefix34+"link_detection_interval").(int)

		   c.LinkDetectionInterval = obj34


		   var obj35 go_thunder.Ipsec
		   prefix35 := "ipsec.0."
		   obj35.PacketRoundRobin = d.Get(prefix35+"packet_round_robin").(int)
		   obj35.CryptoCore = d.Get(prefix35+"crypto_core").(int)
		   obj35.CryptoMem = d.Get(prefix35+"crypto_mem").(int)

		   var obj35_1 go_thunder.FpgaDecrypt
		   prefix35_1 := prefix35 + "fpga_decrypt.0."
		   obj35_1.Action = d.Get(prefix35_1 +"action").(string)

		   obj35.Action = obj35_1


		   c.PacketRoundRobin = obj35


		   var obj36 go_thunder.SpeProfile
		   prefix36 := "spe_profile.0."
		   obj36.Action = d.Get(prefix36+"action").(string)

		   c.Action = obj36


		   var obj37 go_thunder.SpeStatus
		   prefix := "spe_status.0."


		   c. = obj37


		   var obj38 go_thunder.SslStatus
		   prefix := "ssl_status.0."


		   c. = obj38


		   var obj39 go_thunder.DeepHrxq
		   prefix := "deep_hrxq.0."
		   obj39.Enable = d.Get(prefix+"enable").(int)

		   c.Enable = obj39


		   var obj40 go_thunder.HrxqStatus
		   prefix := "hrxq_status.0."


		   c. = obj40


		   var obj41 go_thunder.CpuLoadSharing
		   prefix := "cpu_load_sharing.0."
		   obj41.Disable = d.Get(prefix+"disable").(int)

		   var obj1 go_thunder.PacketsPerSecond
		   prefix := "packets_per_second.0."
		   obj1.Min = d.Get(prefix+"min").(int)

		   obj41.Min = obj1


		   var obj2 go_thunder.CpuUsage
		   prefix := "cpu_usage.0."
		   obj2.Low = d.Get(prefix+"low").(int)
		   obj2.High = d.Get(prefix+"high").(int)

		   obj41.Low = obj2


		   c.Disable = obj41


		   var obj42 go_thunder.PerVlanLimit
		   prefix := "per_vlan_limit.0."
		   obj42.Bcast = d.Get(prefix+"bcast").(int)
		   obj42.Ipmcast = d.Get(prefix+"ipmcast").(int)
		   obj42.Mcast = d.Get(prefix+"mcast").(int)
		   obj42.UnknownUcast = d.Get(prefix+"unknown_ucast").(int)

		   c.Bcast = obj42


		   var obj43 go_thunder.AllVlanLimit
		   prefix := "all_vlan_limit.0."
		   obj43.Bcast = d.Get(prefix+"bcast").(int)
		   obj43.Ipmcast = d.Get(prefix+"ipmcast").(int)
		   obj43.Mcast = d.Get(prefix+"mcast").(int)
		   obj43.UnknownUcast = d.Get(prefix+"unknown_ucast").(int)

		   c.Bcast = obj43


		   var obj44 go_thunder.VeMacScheme
		   prefix := "ve_mac_scheme.0."
		   obj44.VeMacSchemeVal = d.Get(prefix+"ve_mac_scheme_val").(string)

		   c.VeMacSchemeVal = obj44


		   var obj45 go_thunder.SessionReclaimLimit
		   prefix := "session_reclaim_limit.0."
		   obj45.NscanLimit = d.Get(prefix+"nscan_limit").(int)
		   obj45.ScanFreq = d.Get(prefix+"scan_freq").(int)

		   c.NscanLimit = obj45


		   var obj46 go_thunder.SslScv
		   prefix := "ssl_scv.0."
		   obj46.Enable = d.Get(prefix+"enable").(int)

		   c.Enable = obj46


		   var obj47 go_thunder.Hardware
		   prefix := "hardware.0."


		   c. = obj47


		   var obj48 go_thunder.Platformtype
		   prefix := "platformtype.0."


		   c. = obj48


		   var obj49 go_thunder.Reboot
		   prefix := "reboot.0."


		   c. = obj49


		   var obj50 go_thunder.Shutdown
		   prefix := "shutdown.0."


		   c. = obj50


		   var obj51 go_thunder.Environment
		   prefix := "environment.0."


		   c. = obj51


		   var obj52 go_thunder.HardwareForward
		   prefix := "hardware_forward.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj52.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj52.Counters1 = append(obj52.Counters1, obj1)
		   }


		   var obj2 go_thunder.Slb
		   prefix := "slb.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj2.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj2.Counters1 = append(obj2.Counters1, obj1)
		   }


		   obj52. = obj2


		   c.SamplingEnable = obj52


		   var obj53 go_thunder.Throughput
		   prefix := "throughput.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj53.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj53.Counters1 = append(obj53.Counters1, obj1)
		   }


		   c.SamplingEnable = obj53


		   var obj54 go_thunder.Ipmi
		   prefix := "ipmi.0."
		   obj54.Reset = d.Get(prefix+"reset").(int)

		   var obj1 go_thunder.Ip
		   prefix := "ip.0."
		   obj1.Ipv4Address = d.Get(prefix+"ipv4_address").(string)
		   obj1.Ipv4Netmask = d.Get(prefix+"ipv4_netmask").(string)
		   obj1.DefaultGateway = d.Get(prefix+"default_gateway").(string)

		   obj54.Ipv4Address = obj1


		   var obj2 go_thunder.Ipsrc
		   prefix := "ipsrc.0."
		   obj2.Dhcp = d.Get(prefix+"dhcp").(int)
		   obj2.Static = d.Get(prefix+"static").(int)

		   obj54.Dhcp = obj2


		   var obj3 go_thunder.User
		   prefix := "user.0."
		   obj3.Add = d.Get(prefix+"add").(string)
		   obj3.Password = d.Get(prefix+"password").(string)
		   obj3.Administrator = d.Get(prefix+"administrator").(int)
		   obj3.Callback = d.Get(prefix+"callback").(int)
		   obj3.Operator = d.Get(prefix+"operator").(int)
		   obj3.User = d.Get(prefix+"user").(int)
		   obj3.Disable = d.Get(prefix+"disable").(string)
		   obj3.Privilege = d.Get(prefix+"privilege").(string)
		   obj3.Setname = d.Get(prefix+"setname").(string)
		   obj3.Newname = d.Get(prefix+"newname").(string)
		   obj3.Setpass = d.Get(prefix+"setpass").(string)
		   obj3.Newpass = d.Get(prefix+"newpass").(string)

		   obj54.Add = obj3


		   var obj4 go_thunder.Tool
		   prefix := "tool.0."
		   obj4.Cmd = d.Get(prefix+"cmd").(string)

		   obj54.Cmd = obj4


		   c.Reset = obj54


		   var obj55 go_thunder.QueuingBuffer
		   prefix := "queuing_buffer.0."
		   obj55.Enable = d.Get(prefix+"enable").(int)

		   c.Enable = obj55


		   var obj56 go_thunder.TrunkHwHash
		   prefix := "trunk_hw_hash.0."
		   obj56.Mode = d.Get(prefix+"mode").(int)

		   c.Mode = obj56


		   var obj57 go_thunder.TrunkXauiHwHash
		   prefix := "trunk_xaui_hw_hash.0."
		   obj57.Mode = d.Get(prefix+"mode").(int)

		   c.Mode = obj57


		   var obj58 go_thunder.UpgradeStatus
		   prefix := "upgrade_status.0."


		   c. = obj58


		   var obj59 go_thunder.GuestFile
		   prefix := "guest_file.0."


		   c. = obj59


		   var obj60 go_thunder.CmUpdateFileNameRef
		   prefix := "cm_update_file_name_ref.0."
		   obj60.Source_name = d.Get(prefix+"source_name").(string)
		   obj60.Dest_name = d.Get(prefix+"dest_name").(string)
		   obj60.Id = d.Get(prefix+"id").(int)

		   c.Source_name = obj60


		   var obj61 go_thunder.Core
		   prefix := "core.0."


		   c. = obj61


		   var obj62 go_thunder.AppsGlobal
		   prefix := "apps_global.0."
		   obj62.LogSessionOnEstablished = d.Get(prefix+"log_session_on_established").(int)
		   obj62.MslTime = d.Get(prefix+"msl_time").(int)

		   c.LogSessionOnEstablished = obj62


		   var obj63 go_thunder.ShellPrivileges
		   prefix := "shell_privileges.0."


		   c. = obj63


		   var obj64 go_thunder.CosqStats
		   prefix := "cosq_stats.0."


		   c. = obj64


		   var obj65 go_thunder.CosqShow
		   prefix := "cosq_show.0."


		   c. = obj65


		   var obj66 go_thunder.Fw
		   prefix := "fw.0."
		   obj66.ApplicationMempool = d.Get(prefix+"application_mempool").(int)
		   obj66.ApplicationFlow = d.Get(prefix+"application_flow").(int)
		   obj66.BasicDpiEnable = d.Get(prefix+"basic_dpi_enable").(int)

		   c.ApplicationMempool = obj66


		   var obj67 go_thunder.PasswordPolicy
		   prefix := "password_policy.0."
		   obj67.Complexity = d.Get(prefix+"complexity").(string)
		   obj67.Aging = d.Get(prefix+"aging").(string)
		   obj67.History = d.Get(prefix+"history").(string)
		   obj67.MinPswdLen = d.Get(prefix+"min_pswd_len").(int)

		   c.Complexity = obj67


		   var obj68 go_thunder.Radius
		   prefix := "radius.0."

		   var obj1 go_thunder.Server
		   prefix := "server.0."
		   obj1.ListenPort = d.Get(prefix+"listen_port").(int)

		   var obj1 go_thunder.Remote
		   prefix := "remote.0."

		   IpListCount := d.Get("ip_list.#").(int)
		   obj1.IpListName = make([]go_thunder.IpList, 0, IpListCount)

		   for i := 0; i < IpListCount; i++ {
		   var obj1 go_thunder.IpList
		   prefix := fmt.Sprintf("ip_list.%d.",i)
		   obj1.IpListName = d.Get(prefix+"ip_list_name").(string)
		   obj1.IpListSecret = d.Get(prefix+"ip_list_secret").(int)
		   obj1.IpListSecretString = d.Get(prefix+"ip_list_secret_string").(string)
		   obj1.IpListName = append(obj1.IpListName, obj1)
		   }


		   obj1.IpList = obj1

		   obj1.Secret = d.Get(prefix+"secret").(int)
		   obj1.SecretString = d.Get(prefix+"secret_string").(string)
		   obj1.Vrid = d.Get(prefix+"vrid").(int)

		   AttributeCount := d.Get("attribute.#").(int)
		   obj1.AttributeValue = make([]go_thunder.Attribute, 0, AttributeCount)

		   for i := 0; i < AttributeCount; i++ {
		   var obj2 go_thunder.Attribute
		   prefix := fmt.Sprintf("attribute.%d.",i)
		   obj2.AttributeValue = d.Get(prefix+"attribute_value").(string)
		   obj2.PrefixLength = d.Get(prefix+"prefix_length").(string)
		   obj2.PrefixVendor = d.Get(prefix+"prefix_vendor").(int)
		   obj2.PrefixNumber = d.Get(prefix+"prefix_number").(int)
		   obj2.Name = d.Get(prefix+"name").(string)
		   obj2.Value = d.Get(prefix+"value").(string)
		   obj2.CustomVendor = d.Get(prefix+"custom_vendor").(int)
		   obj2.CustomNumber = d.Get(prefix+"custom_number").(int)
		   obj2.Vendor = d.Get(prefix+"vendor").(int)
		   obj2.Number = d.Get(prefix+"number").(int)
		   obj1.AttributeValue = append(obj1.AttributeValue, obj2)
		   }

		   obj1.DisableReply = d.Get(prefix+"disable_reply").(int)
		   obj1.AccountingStart = d.Get(prefix+"accounting_start").(string)
		   obj1.AccountingStop = d.Get(prefix+"accounting_stop").(string)
		   obj1.AccountingInterimUpdate = d.Get(prefix+"accounting_interim_update").(string)
		   obj1.AccountingOn = d.Get(prefix+"accounting_on").(string)
		   obj1.AttributeName = d.Get(prefix+"attribute_name").(string)

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj1.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj3 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj3.Counters1 = d.Get(prefix+"counters1").(string)
		   obj1.Counters1 = append(obj1.Counters1, obj3)
		   }


		   obj68.ListenPort = obj1


		   c.Server = obj68


		   GeolocListListCount := d.Get("geoloc_list_list.#").(int)
		   c.Name = make([]go_thunder.GeolocListList, 0, GeolocListListCount)

		   for i := 0; i < GeolocListListCount; i++ {
		   var obj69 go_thunder.GeolocListList
		   prefix := fmt.Sprintf("geoloc_list_list.%d.",i)
		   obj69.Name = d.Get(prefix+"name").(string)
		   obj69.Shared = d.Get(prefix+"shared").(int)

		   IncludeGeolocNameListCount := d.Get("include_geoloc_name_list.#").(int)
		   obj69.IncludeGeolocNameVal = make([]go_thunder.IncludeGeolocNameList, 0, IncludeGeolocNameListCount)

		   for i := 0; i < IncludeGeolocNameListCount; i++ {
		   var obj1 go_thunder.IncludeGeolocNameList
		   prefix := fmt.Sprintf("include_geoloc_name_list.%d.",i)
		   obj1.IncludeGeolocNameVal = d.Get(prefix+"include_geoloc_name_val").(string)
		   obj69.IncludeGeolocNameVal = append(obj69.IncludeGeolocNameVal, obj1)
		   }


		   ExcludeGeolocNameListCount := d.Get("exclude_geoloc_name_list.#").(int)
		   obj69.ExcludeGeolocNameVal = make([]go_thunder.ExcludeGeolocNameList, 0, ExcludeGeolocNameListCount)

		   for i := 0; i < ExcludeGeolocNameListCount; i++ {
		   var obj2 go_thunder.ExcludeGeolocNameList
		   prefix := fmt.Sprintf("exclude_geoloc_name_list.%d.",i)
		   obj2.ExcludeGeolocNameVal = d.Get(prefix+"exclude_geoloc_name_val").(string)
		   obj69.ExcludeGeolocNameVal = append(obj69.ExcludeGeolocNameVal, obj2)
		   }

		   obj69.UserTag = d.Get(prefix+"user_tag").(string)

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj69.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj3 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj3.Counters1 = d.Get(prefix+"counters1").(string)
		   obj69.Counters1 = append(obj69.Counters1, obj3)
		   }

		   c.Name = append(c.Name, obj69)
		   }


		   var obj70 go_thunder.GeolocNameHelper
		   prefix := "geoloc_name_helper.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj70.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj70.Counters1 = append(obj70.Counters1, obj1)
		   }


		   c.SamplingEnable = obj70


		   var obj71 go_thunder.GeolocationFile
		   prefix := "geolocation_file.0."

		   var obj1 go_thunder.ErrorInfo
		   prefix := "error_info.0."


		   obj71. = obj1


		   c.ErrorInfo = obj71


		   var obj72 go_thunder.Geoloc
		   prefix := "geoloc.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj72.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj72.Counters1 = append(obj72.Counters1, obj1)
		   }


		   c.SamplingEnable = obj72


		   var obj73 go_thunder.GeoLocation
		   prefix := "geo_location.0."
		   obj73.GeoLocationIana = d.Get(prefix+"geo_location_iana").(int)
		   obj73.GeoLocationGeolite2City = d.Get(prefix+"geo_location_geolite2_city").(int)
		   obj73.Geolite2CityIncludeIpv6 = d.Get(prefix+"geolite2_city_include_ipv6").(int)
		   obj73.GeoLocationGeolite2Country = d.Get(prefix+"geo_location_geolite2_country").(int)
		   obj73.Geolite2CountryIncludeIpv6 = d.Get(prefix+"geolite2_country_include_ipv6").(int)

		   GeolocLoadFileListCount := d.Get("geoloc_load_file_list.#").(int)
		   obj73.GeoLocationLoadFilename = make([]go_thunder.GeolocLoadFileList, 0, GeolocLoadFileListCount)

		   for i := 0; i < GeolocLoadFileListCount; i++ {
		   var obj1 go_thunder.GeolocLoadFileList
		   prefix := fmt.Sprintf("geoloc_load_file_list.%d.",i)
		   obj1.GeoLocationLoadFilename = d.Get(prefix+"geo_location_load_filename").(string)
		   obj1.TemplateName = d.Get(prefix+"template_name").(string)
		   obj73.GeoLocationLoadFilename = append(obj73.GeoLocationLoadFilename, obj1)
		   }


		   EntryListCount := d.Get("entry_list.#").(int)
		   obj73.GeoLocnObjName = make([]go_thunder.EntryList, 0, EntryListCount)

		   for i := 0; i < EntryListCount; i++ {
		   var obj2 go_thunder.EntryList
		   prefix := fmt.Sprintf("entry_list.%d.",i)
		   obj2.GeoLocnObjName = d.Get(prefix+"geo_locn_obj_name").(string)

		   GeoLocnMultipleAddressesCount := d.Get("geo_locn_multiple_addresses.#").(int)
		   obj2.FirstIpAddress = make([]go_thunder.GeoLocnMultipleAddresses, 0, GeoLocnMultipleAddressesCount)

		   for i := 0; i < GeoLocnMultipleAddressesCount; i++ {
		   var obj1 go_thunder.GeoLocnMultipleAddresses
		   prefix := fmt.Sprintf("geo_locn_multiple_addresses.%d.",i)
		   obj1.FirstIpAddress = d.Get(prefix+"first_ip_address").(string)
		   obj1.GeolIpv4Mask = d.Get(prefix+"geol_ipv4_mask").(string)
		   obj1.IpAddr2 = d.Get(prefix+"ip_addr2").(string)
		   obj1.FirstIpv6Address = d.Get(prefix+"first_ipv6_address").(string)
		   obj1.GeolIpv6Mask = d.Get(prefix+"geol_ipv6_mask").(int)
		   obj1.Ipv6Addr2 = d.Get(prefix+"ipv6_addr2").(string)
		   obj2.FirstIpAddress = append(obj2.FirstIpAddress, obj1)
		   }

		   obj2.UserTag = d.Get(prefix+"user_tag").(string)
		   obj73.GeoLocnObjName = append(obj73.GeoLocnObjName, obj2)
		   }


		   c.GeoLocationIana = obj73


		   var obj74 go_thunder.IpThreatList
		   prefix := "ip_threat_list.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj74.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj74.Counters1 = append(obj74.Counters1, obj1)
		   }


		   var obj2 go_thunder.Ipv4SourceList
		   prefix := "ipv4_source_list.0."

		   ClassListCfgCount := d.Get("class_list_cfg.#").(int)
		   obj2.ClassList = make([]go_thunder.ClassListCfg, 0, ClassListCfgCount)

		   for i := 0; i < ClassListCfgCount; i++ {
		   var obj1 go_thunder.ClassListCfg
		   prefix := fmt.Sprintf("class_list_cfg.%d.",i)
		   obj1.ClassList = d.Get(prefix+"class_list").(string)
		   obj1.IpThreatActionTmpl = d.Get(prefix+"ip_threat_action_tmpl").(int)
		   obj2.ClassList = append(obj2.ClassList, obj1)
		   }


		   obj74.ClassListCfg = obj2


		   var obj3 go_thunder.Ipv4DestList
		   prefix := "ipv4_dest_list.0."

		   ClassListCfgCount := d.Get("class_list_cfg.#").(int)
		   obj3.ClassList = make([]go_thunder.ClassListCfg, 0, ClassListCfgCount)

		   for i := 0; i < ClassListCfgCount; i++ {
		   var obj1 go_thunder.ClassListCfg
		   prefix := fmt.Sprintf("class_list_cfg.%d.",i)
		   obj1.ClassList = d.Get(prefix+"class_list").(string)
		   obj1.IpThreatActionTmpl = d.Get(prefix+"ip_threat_action_tmpl").(int)
		   obj3.ClassList = append(obj3.ClassList, obj1)
		   }


		   obj74.ClassListCfg = obj3


		   var obj4 go_thunder.Ipv6SourceList
		   prefix := "ipv6_source_list.0."

		   ClassListCfgCount := d.Get("class_list_cfg.#").(int)
		   obj4.ClassList = make([]go_thunder.ClassListCfg, 0, ClassListCfgCount)

		   for i := 0; i < ClassListCfgCount; i++ {
		   var obj1 go_thunder.ClassListCfg
		   prefix := fmt.Sprintf("class_list_cfg.%d.",i)
		   obj1.ClassList = d.Get(prefix+"class_list").(string)
		   obj1.IpThreatActionTmpl = d.Get(prefix+"ip_threat_action_tmpl").(int)
		   obj4.ClassList = append(obj4.ClassList, obj1)
		   }


		   obj74.ClassListCfg = obj4


		   var obj5 go_thunder.Ipv6DestList
		   prefix := "ipv6_dest_list.0."

		   ClassListCfgCount := d.Get("class_list_cfg.#").(int)
		   obj5.ClassList = make([]go_thunder.ClassListCfg, 0, ClassListCfgCount)

		   for i := 0; i < ClassListCfgCount; i++ {
		   var obj1 go_thunder.ClassListCfg
		   prefix := fmt.Sprintf("class_list_cfg.%d.",i)
		   obj1.ClassList = d.Get(prefix+"class_list").(string)
		   obj1.IpThreatActionTmpl = d.Get(prefix+"ip_threat_action_tmpl").(int)
		   obj5.ClassList = append(obj5.ClassList, obj1)
		   }


		   obj74.ClassListCfg = obj5


		   c.SamplingEnable = obj74


		   var obj75 go_thunder.FpgaDrop
		   prefix := "fpga_drop.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj75.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj75.Counters1 = append(obj75.Counters1, obj1)
		   }


		   c.SamplingEnable = obj75


		   var obj76 go_thunder.DpdkStats
		   prefix := "dpdk_stats.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj76.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj76.Counters1 = append(obj76.Counters1, obj1)
		   }


		   c.SamplingEnable = obj76


		   var obj77 go_thunder.FpgaCoreCrc
		   prefix := "fpga_core_crc.0."
		   obj77.MonitorDisable = d.Get(prefix+"monitor_disable").(int)
		   obj77.RebootEnable = d.Get(prefix+"reboot_enable").(int)

		   c.MonitorDisable = obj77


		   var obj78 go_thunder.PsuInfo
		   prefix := "psu_info.0."


		   c. = obj78


		   var obj79 go_thunder.GuiImageList
		   prefix := "gui_image_list.0."


		   c. = obj79


		   var obj80 go_thunder.SyslogTimeMsec
		   prefix := "syslog_time_msec.0."
		   obj80.EnableFlag = d.Get(prefix+"enable_flag").(int)

		   c.EnableFlag = obj80


		   var obj81 go_thunder.IpmiService
		   prefix := "ipmi_service.0."
		   obj81.Disable = d.Get(prefix+"disable").(int)

		   c.Disable = obj81


		   var obj82 go_thunder.AppPerformance
		   prefix := "app_performance.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj82.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj82.Counters1 = append(obj82.Counters1, obj1)
		   }


		   c.SamplingEnable = obj82


		   var obj83 go_thunder.SslReqQ
		   prefix := "ssl_req_q.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj83.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj83.Counters1 = append(obj83.Counters1, obj1)
		   }


		   c.SamplingEnable = obj83


		   var obj84 go_thunder.Tcp
		   prefix := "tcp.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj84.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj84.Counters1 = append(obj84.Counters1, obj1)
		   }


		   var obj2 go_thunder.RateLimitResetUnknownConn
		   prefix := "rate_limit_reset_unknown_conn.0."
		   obj2.PktRateForResetUnknownConn = d.Get(prefix+"pkt_rate_for_reset_unknown_conn").(int)
		   obj2.LogForResetUnknownConn = d.Get(prefix+"log_for_reset_unknown_conn").(int)

		   obj84.PktRateForResetUnknownConn = obj2


		   c.SamplingEnable = obj84


		   var obj85 go_thunder.Icmp
		   prefix := "icmp.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj85.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj85.Counters1 = append(obj85.Counters1, obj1)
		   }


		   c.SamplingEnable = obj85


		   var obj86 go_thunder.Icmp6
		   prefix := "icmp6.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj86.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj86.Counters1 = append(obj86.Counters1, obj1)
		   }


		   c.SamplingEnable = obj86


		   var obj87 go_thunder.IpStats
		   prefix := "ip_stats.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj87.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj87.Counters1 = append(obj87.Counters1, obj1)
		   }


		   c.SamplingEnable = obj87


		   var obj88 go_thunder.Ip6Stats
		   prefix := "ip6_stats.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj88.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj88.Counters1 = append(obj88.Counters1, obj1)
		   }


		   c.SamplingEnable = obj88


		   var obj89 go_thunder.DomainListInfo
		   prefix := "domain_list_info.0."


		   c. = obj89


		   var obj90 go_thunder.IpDnsCache
		   prefix := "ip_dns_cache.0."


		   c. = obj90


		   var obj91 go_thunder.Bfd
		   prefix := "bfd.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj91.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj91.Counters1 = append(obj91.Counters1, obj1)
		   }


		   c.SamplingEnable = obj91


		   var obj92 go_thunder.IcmpRate
		   prefix := "icmp_rate.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj92.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj92.Counters1 = append(obj92.Counters1, obj1)
		   }


		   c.SamplingEnable = obj92


		   var obj93 go_thunder.Dns
		   prefix := "dns.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj93.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj93.Counters1 = append(obj93.Counters1, obj1)
		   }


		   c.SamplingEnable = obj93


		   var obj94 go_thunder.DnsCache
		   prefix := "dns_cache.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj94.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj94.Counters1 = append(obj94.Counters1, obj1)
		   }


		   c.SamplingEnable = obj94


		   var obj95 go_thunder.Session
		   prefix := "session.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj95.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj95.Counters1 = append(obj95.Counters1, obj1)
		   }


		   c.SamplingEnable = obj95


		   var obj96 go_thunder.NdiscRa
		   prefix := "ndisc_ra.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj96.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj96.Counters1 = append(obj96.Counters1, obj1)
		   }


		   c.SamplingEnable = obj96


		   var obj97 go_thunder.TcpStats
		   prefix := "tcp_stats.0."

		   SamplingEnableCount := d.Get("sampling_enable.#").(int)
		   obj97.Counters1 = make([]go_thunder.SamplingEnable, 0, SamplingEnableCount)

		   for i := 0; i < SamplingEnableCount; i++ {
		   var obj1 go_thunder.SamplingEnable
		   prefix := fmt.Sprintf("sampling_enable.%d.",i)
		   obj1.Counters1 = d.Get(prefix+"counters1").(string)
		   obj97.Counters1 = append(obj97.Counters1, obj1)
		   }


		   c.SamplingEnable = obj97


		   var obj98 go_thunder.TelemetryLog
		   prefix := "telemetry_log.0."

		   var obj1 go_thunder.TopKSourceList
		   prefix := "top_k_source_list.0."


		   obj98.Uuid = obj1


		   var obj2 go_thunder.TopKAppSvcList
		   prefix := "top_k_app_svc_list.0."


		   obj98.Uuid = obj2


		   var obj3 go_thunder.DeviceStatus
		   prefix := "device_status.0."


		   obj98.Uuid = obj3


		   var obj4 go_thunder.Environment
		   prefix := "environment.0."


		   obj98.Uuid = obj4


		   var obj5 go_thunder.PartitionMetrics
		   prefix := "partition_metrics.0."


		   obj98.Uuid = obj5


		   c.TopKSourceList = obj98

	*/
	vc.AnomalyLog = c
	return vc
}
