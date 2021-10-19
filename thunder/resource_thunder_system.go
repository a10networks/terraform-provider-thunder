package thunder

//Thunder resource System

import (
	"context"
	"fmt"
	"github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"util"
)

func resourceSystem() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemCreate,
		UpdateContext: resourceSystemUpdate,
		ReadContext:   resourceSystemRead,
		DeleteContext: resourceSystemDelete,
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
			"system_chassis_port_split_enable": {
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
			"probe_network_devices": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"management_interface_mode": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dedicated": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"non_dedicated": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"set_tcp_syn_per_sec": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_syn_value": {
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
						"uuid": {
							Type:        schema.TypeString,
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
			"lro": {
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
			"tso": {
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
						"link_down_on_restart": {
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
						"class_list_entry_count": {
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
						"waf_template_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"auth_session_count": {
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
			"table_integrity": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"table": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"audit_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"auto_sync_action": {
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
						"qat": {
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
						"tcp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"udp": {
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
			"ssl_scv_verify_host": {
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
			"ssl_scv_verify_crl_sign": {
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
			"ssl_set_compatible_cipher": {
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
			"high_memory_l4_session": {
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
						"enable_shell_privileges": {
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
			"shm_logging": {
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
									"custom_attribute_name": {
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
			"tcp_syn_per_sec": {
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
						"ipv4_internet_host_list": {
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
						"ipv6_internet_host_list": {
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
			"mfa_management": {
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
			"mfa_validation_type": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_cert": {
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
			"mfa_cert_store": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"cert_store_path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"passwd_string": {
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
			"mfa_auth": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"second_factor": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"q_in_q": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inner_tpid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"outer_tpid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"enable_all_ports": {
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
			"port_count": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_count_kernel": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_count_hm": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_count_logging": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"port_count_alg": {
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
			"health_check_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l2hm_hc_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"method_l2bfd": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l2bfd_tx_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l2bfd_rx_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l2bfd_multiplier": {
							Type:        schema.TypeInt,
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
			"path_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l2hm_path_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"l2hm_vlan": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l2hm_setup_test_api": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ifpair_eth_start": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ifpair_eth_end": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ifpair_trunk_start": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ifpair_trunk_end": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"l2hm_attach": {
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
			"job_offload": {
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

func resourceSystemCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Creating System (Inside resourceSystemCreate) ")

		data := dataToSystem(d)
		logger.Println("[INFO] received formatted data from method data to System --")
		d.SetId("1")
		err := go_thunder.PostSystem(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSystemRead(ctx, d, meta)

	}
	return diags
}

func resourceSystemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading System (Inside resourceSystemRead)")

	var diags diag.Diagnostics
	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSystem(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return diags
}

func resourceSystemUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSystemRead(ctx, d, meta)
}

func resourceSystemDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSystemRead(ctx, d, meta)
}
func dataToSystem(d *schema.ResourceData) go_thunder.System {
	var vc go_thunder.System
	var c go_thunder.SystemInstance
	c.SystemInstanceAnomalyLog = d.Get("anomaly_log").(int)
	c.SystemInstanceAttackLog = d.Get("attack_log").(int)
	c.SystemInstanceDdosAttack = d.Get("ddos_attack").(int)
	c.SystemInstanceDdosLog = d.Get("ddos_log").(int)
	c.SystemInstanceSockstressDisable = d.Get("sockstress_disable").(int)
	c.SystemInstancePromiscuousMode = d.Get("promiscuous_mode").(int)
	c.SystemInstanceGlid = d.Get("glid").(int)
	c.SystemInstanceModuleCtrlCPU = d.Get("module_ctrl_cpu").(string)
	c.SystemInstanceSrcIPHashEnable = d.Get("src_ip_hash_enable").(int)
	c.SystemInstanceClassListHitcountEnable = d.Get("class_list_hitcount_enable").(int)
	c.SystemInstanceGeoDbHitcountEnable = d.Get("geo_db_hitcount_enable").(int)
	c.SystemInstanceDomainListHitcountEnable = d.Get("domain_list_hitcount_enable").(int)
	c.SystemInstanceDynamicServiceDNSSocketPool = d.Get("dynamic_service_dns_socket_pool").(int)
	c.SystemInstanceSystemChassisPortSplitEnable = d.Get("system_chassis_port_split_enable").(int)

	var obj1 go_thunder.SystemInstanceTimeoutValue
	prefix1 := "timeout_value.0."
	obj1.SystemInstanceTimeoutValueFtp = d.Get(prefix1 + "ftp").(int)
	obj1.SystemInstanceTimeoutValueScp = d.Get(prefix1 + "scp").(int)
	obj1.SystemInstanceTimeoutValueSftp = d.Get(prefix1 + "sftp").(int)
	obj1.SystemInstanceTimeoutValueTftp = d.Get(prefix1 + "tftp").(int)
	obj1.SystemInstanceTimeoutValueHTTP = d.Get(prefix1 + "http").(int)
	obj1.SystemInstanceTimeoutValueHTTPS = d.Get(prefix1 + "https").(int)

	c.SystemInstanceTimeoutValueFtp = obj1

	var obj2 go_thunder.SystemInstanceBandwidth
	prefix2 := "bandwidth.0."

	SystemInstanceBandwidthSamplingEnableCount := d.Get(prefix2 + "sampling_enable.#").(int)
	obj2.SystemInstanceBandwidthSamplingEnableCounters1 = make([]go_thunder.SystemInstanceBandwidthSamplingEnable, 0, SystemInstanceBandwidthSamplingEnableCount)

	for i := 0; i < SystemInstanceBandwidthSamplingEnableCount; i++ {
		var obj2_1 go_thunder.SystemInstanceBandwidthSamplingEnable
		prefix2_1 := prefix2 + fmt.Sprintf("sampling_enable.%d.", i)
		obj2_1.SystemInstanceBandwidthSamplingEnableCounters1 = d.Get(prefix2_1 + "counters1").(string)
		obj2.SystemInstanceBandwidthSamplingEnableCounters1 = append(obj2.SystemInstanceBandwidthSamplingEnableCounters1, obj2_1)
	}

	c.SystemInstanceBandwidthUUID = obj2

	var obj3 go_thunder.SystemInstanceMgmtPort
	prefix3 := "mgmt_port.0."
	obj3.SystemInstanceMgmtPortPortIndex = d.Get(prefix3 + "port_index").(int)
	obj3.SystemInstanceMgmtPortMacAddress = d.Get(prefix3 + "mac_address").(string)
	obj3.SystemInstanceMgmtPortPciAddress = d.Get(prefix3 + "pci_address").(string)

	c.SystemInstanceMgmtPortPortIndex = obj3

	var obj4 go_thunder.SystemInstanceSharedPollMode
	prefix4 := "shared_poll_mode.0."
	obj4.SystemInstanceSharedPollModeEnable = d.Get(prefix4 + "enable").(int)
	obj4.SystemInstanceSharedPollModeDisable = d.Get(prefix4 + "disable").(int)

	c.SystemInstanceSharedPollModeEnable = obj4

	var obj6 go_thunder.SystemInstanceManagementInterfaceMode
	prefix6 := "management_interface_mode.0."
	obj6.SystemInstanceManagementInterfaceModeDedicated = d.Get(prefix6 + "dedicated").(int)
	obj6.SystemInstanceManagementInterfaceModeNonDedicated = d.Get(prefix6 + "non_dedicated").(int)

	c.SystemInstanceManagementInterfaceModeDedicated = obj6

	var obj7 go_thunder.SystemInstanceSetTCPSynPerSec
	prefix7 := "set_tcp_syn_per_sec.0."
	obj7.SystemInstanceSetTCPSynPerSecTCPSynValue = d.Get(prefix7 + "tcp_syn_value").(int)

	c.SystemInstanceSetTCPSynPerSecTCPSynValue = obj7

	var obj8 go_thunder.SystemInstanceAddPort
	prefix8 := "add_port.0."
	obj8.SystemInstanceAddPortPortIndex = d.Get(prefix8 + "port_index").(int)

	c.SystemInstanceAddPortPortIndex = obj8

	var obj9 go_thunder.SystemInstanceDelPort
	prefix9 := "del_port.0."
	obj9.SystemInstanceDelPortPortIndex = d.Get(prefix9 + "port_index").(int)

	c.SystemInstanceDelPortPortIndex = obj9

	var obj10 go_thunder.SystemInstanceModifyPort
	prefix10 := "modify_port.0."
	obj10.SystemInstanceModifyPortPortIndex = d.Get(prefix10 + "port_index").(int)
	obj10.SystemInstanceModifyPortPortNumber = d.Get(prefix10 + "port_number").(int)

	c.SystemInstanceModifyPortPortIndex = obj10

	var obj11 go_thunder.SystemInstanceTLS13Mgmt
	prefix11 := "tls_1_3_mgmt.0."
	obj11.SystemInstanceTLS13MgmtEnable = d.Get(prefix11 + "enable").(int)

	c.SystemInstanceTLS13MgmtEnable = obj11

	var obj12 go_thunder.SystemInstanceMultiQueueSupport
	prefix12 := "multi_queue_support.0."
	obj12.SystemInstanceMultiQueueSupportEnable = d.Get(prefix12 + "enable").(int)

	c.SystemInstanceMultiQueueSupportEnable = obj12

	var obj13 go_thunder.SystemInstanceAddCPUCore
	prefix13 := "add_cpu_core.0."
	obj13.SystemInstanceAddCPUCoreCoreIndex = d.Get(prefix13 + "core_index").(int)

	c.SystemInstanceAddCPUCoreCoreIndex = obj13

	var obj14 go_thunder.SystemInstanceDeleteCPUCore
	prefix14 := "delete_cpu_core.0."
	obj14.SystemInstanceDeleteCPUCoreCoreIndex = d.Get(prefix14 + "core_index").(int)

	c.SystemInstanceDeleteCPUCoreCoreIndex = obj14

	var obj15 go_thunder.SystemInstanceCPUHyperThread
	prefix15 := "cpu_hyper_thread.0."
	obj15.SystemInstanceCPUHyperThreadEnable = d.Get(prefix15 + "enable").(int)
	obj15.SystemInstanceCPUHyperThreadDisable = d.Get(prefix15 + "disable").(int)

	c.SystemInstanceCPUHyperThreadEnable = obj15

	var obj16 go_thunder.SystemInstanceIoCPU
	prefix16 := "io_cpu.0."
	obj16.SystemInstanceIoCPUMaxCores = d.Get(prefix16 + "max_cores").(int)

	c.SystemInstanceIoCPUMaxCores = obj16

	var obj17 go_thunder.SystemInstanceLinkMonitor
	prefix17 := "link_monitor.0."
	obj17.SystemInstanceLinkMonitorEnable = d.Get(prefix17 + "enable").(int)
	obj17.SystemInstanceLinkMonitorDisable = d.Get(prefix17 + "disable").(int)

	c.SystemInstanceLinkMonitorEnable = obj17

	var obj18 go_thunder.SystemInstanceLro
	prefix18 := "lro.0."
	obj18.SystemInstanceLroEnable = d.Get(prefix18 + "enable").(int)
	obj18.SystemInstanceLroDisable = d.Get(prefix18 + "disable").(int)

	c.SystemInstanceLroEnable = obj18

	var obj19 go_thunder.SystemInstanceTso
	prefix19 := "tso.0."
	obj19.SystemInstanceTsoEnable = d.Get(prefix19 + "enable").(int)
	obj19.SystemInstanceTsoDisable = d.Get(prefix19 + "disable").(int)

	c.SystemInstanceTsoEnable = obj19

	var obj20 go_thunder.SystemInstanceSetRxtxDescSize
	prefix20 := "set_rxtx_desc_size.0."
	obj20.SystemInstanceSetRxtxDescSizePortIndex = d.Get(prefix20 + "port_index").(int)
	obj20.SystemInstanceSetRxtxDescSizeRxdSize = d.Get(prefix20 + "rxd_size").(int)
	obj20.SystemInstanceSetRxtxDescSizeTxdSize = d.Get(prefix20 + "txd_size").(int)

	c.SystemInstanceSetRxtxDescSizePortIndex = obj20

	var obj21 go_thunder.SystemInstanceSetRxtxQueue
	prefix21 := "set_rxtx_queue.0."
	obj21.SystemInstanceSetRxtxQueuePortIndex = d.Get(prefix21 + "port_index").(int)
	obj21.SystemInstanceSetRxtxQueueRxqSize = d.Get(prefix21 + "rxq_size").(int)
	obj21.SystemInstanceSetRxtxQueueTxqSize = d.Get(prefix21 + "txq_size").(int)

	c.SystemInstanceSetRxtxQueuePortIndex = obj21

	var obj22 go_thunder.SystemInstanceTemplate
	prefix22 := "template.0."
	obj22.SystemInstanceTemplateTemplatePolicy = d.Get(prefix22 + "template_policy").(string)

	c.SystemInstanceTemplateTemplatePolicy = obj22

	var obj23 go_thunder.SystemInstanceTemplateBind
	prefix23 := "template_bind.0."

	SystemInstanceTemplateBindMonitorListCount := d.Get(prefix23 + "monitor_list.#").(int)
	obj23.SystemInstanceTemplateBindMonitorListID = make([]go_thunder.SystemInstanceTemplateBindMonitorList, 0, SystemInstanceTemplateBindMonitorListCount)

	for i := 0; i < SystemInstanceTemplateBindMonitorListCount; i++ {
		var obj23_1 go_thunder.SystemInstanceTemplateBindMonitorList
		prefix23_1 := prefix23 + fmt.Sprintf("monitor_list.%d.", i)
		obj23_1.SystemInstanceTemplateBindMonitorListID = d.Get(prefix23_1 + "id").(int)

		SystemInstanceTemplateBindMonitorListClearCfgCount := d.Get(prefix23_1 + "clear_cfg.#").(int)
		obj23_1.SystemInstanceTemplateBindMonitorListClearCfgSessions = make([]go_thunder.SystemInstanceTemplateBindMonitorListClearCfg, 0, SystemInstanceTemplateBindMonitorListClearCfgCount)

		for i := 0; i < SystemInstanceTemplateBindMonitorListClearCfgCount; i++ {
			var obj23_1_1 go_thunder.SystemInstanceTemplateBindMonitorListClearCfg
			prefix23_1_1 := prefix23_1 + fmt.Sprintf("clear_cfg.%d.", i)
			obj23_1_1.SystemInstanceTemplateBindMonitorListClearCfgSessions = d.Get(prefix23_1_1 + "sessions").(string)
			obj23_1_1.SystemInstanceTemplateBindMonitorListClearCfgClearAllSequence = d.Get(prefix23_1_1 + "clear_all_sequence").(int)
			obj23_1_1.SystemInstanceTemplateBindMonitorListClearCfgClearSequence = d.Get(prefix23_1_1 + "clear_sequence").(int)
			obj23_1.SystemInstanceTemplateBindMonitorListClearCfgSessions = append(obj23_1.SystemInstanceTemplateBindMonitorListClearCfgSessions, obj23_1_1)
		}

		SystemInstanceTemplateBindMonitorListLinkDisableCfgCount := d.Get(prefix23_1 + "link_disable_cfg.#").(int)
		obj23_1.SystemInstanceTemplateBindMonitorListLinkDisableCfgDiseth = make([]go_thunder.SystemInstanceTemplateBindMonitorListLinkDisableCfg, 0, SystemInstanceTemplateBindMonitorListLinkDisableCfgCount)

		for i := 0; i < SystemInstanceTemplateBindMonitorListLinkDisableCfgCount; i++ {
			var obj23_1_2 go_thunder.SystemInstanceTemplateBindMonitorListLinkDisableCfg
			prefix23_1_2 := prefix23_1 + fmt.Sprintf("link_disable_cfg.%d.", i)
			obj23_1_2.SystemInstanceTemplateBindMonitorListLinkDisableCfgDiseth = d.Get(prefix23_1_2 + "diseth").(int)
			obj23_1_2.SystemInstanceTemplateBindMonitorListLinkDisableCfgDisSequence = d.Get(prefix23_1_2 + "dis_sequence").(int)
			obj23_1.SystemInstanceTemplateBindMonitorListLinkDisableCfgDiseth = append(obj23_1.SystemInstanceTemplateBindMonitorListLinkDisableCfgDiseth, obj23_1_2)
		}

		SystemInstanceTemplateBindMonitorListLinkEnableCfgCount := d.Get(prefix23_1 + "link_enable_cfg.#").(int)
		obj23_1.SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaeth = make([]go_thunder.SystemInstanceTemplateBindMonitorListLinkEnableCfg, 0, SystemInstanceTemplateBindMonitorListLinkEnableCfgCount)

		for i := 0; i < SystemInstanceTemplateBindMonitorListLinkEnableCfgCount; i++ {
			var obj23_1_3 go_thunder.SystemInstanceTemplateBindMonitorListLinkEnableCfg
			prefix23_1_3 := prefix23_1 + fmt.Sprintf("link_enable_cfg.%d.", i)
			obj23_1_3.SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaeth = d.Get(prefix23_1_3 + "enaeth").(int)
			obj23_1_3.SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaSequence = d.Get(prefix23_1_3 + "ena_sequence").(int)
			obj23_1.SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaeth = append(obj23_1.SystemInstanceTemplateBindMonitorListLinkEnableCfgEnaeth, obj23_1_3)
		}

		obj23_1.SystemInstanceTemplateBindMonitorListMonitorRelation = d.Get(prefix23_1 + "monitor_relation").(string)

		SystemInstanceTemplateBindMonitorListLinkUpCfgCount := d.Get(prefix23_1 + "link_up_cfg.#").(int)
		obj23_1.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet1 = make([]go_thunder.SystemInstanceTemplateBindMonitorListLinkUpCfg, 0, SystemInstanceTemplateBindMonitorListLinkUpCfgCount)

		for i := 0; i < SystemInstanceTemplateBindMonitorListLinkUpCfgCount; i++ {
			var obj23_1_4 go_thunder.SystemInstanceTemplateBindMonitorListLinkUpCfg
			prefix23_1_4 := prefix23_1 + fmt.Sprintf("link_up_cfg.%d.", i)
			obj23_1_4.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet1 = d.Get(prefix23_1_4 + "linkup_ethernet1").(int)
			obj23_1_4.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkUpSequence1 = d.Get(prefix23_1_4 + "link_up_sequence1").(int)
			obj23_1_4.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet2 = d.Get(prefix23_1_4 + "linkup_ethernet2").(int)
			obj23_1_4.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkUpSequence2 = d.Get(prefix23_1_4 + "link_up_sequence2").(int)
			obj23_1_4.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet3 = d.Get(prefix23_1_4 + "linkup_ethernet3").(int)
			obj23_1_4.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkUpSequence3 = d.Get(prefix23_1_4 + "link_up_sequence3").(int)
			obj23_1.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet1 = append(obj23_1.SystemInstanceTemplateBindMonitorListLinkUpCfgLinkupEthernet1, obj23_1_4)
		}

		SystemInstanceTemplateBindMonitorListLinkDownCfgCount := d.Get(prefix23_1 + "link_down_cfg.#").(int)
		obj23_1.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet1 = make([]go_thunder.SystemInstanceTemplateBindMonitorListLinkDownCfg, 0, SystemInstanceTemplateBindMonitorListLinkDownCfgCount)

		for i := 0; i < SystemInstanceTemplateBindMonitorListLinkDownCfgCount; i++ {
			var obj23_1_5 go_thunder.SystemInstanceTemplateBindMonitorListLinkDownCfg
			prefix23_1_5 := prefix23_1 + fmt.Sprintf("link_down_cfg.%d.", i)
			obj23_1_5.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet1 = d.Get(prefix23_1_5 + "linkdown_ethernet1").(int)
			obj23_1_5.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkDownSequence1 = d.Get(prefix23_1_5 + "link_down_sequence1").(int)
			obj23_1_5.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet2 = d.Get(prefix23_1_5 + "linkdown_ethernet2").(int)
			obj23_1_5.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkDownSequence2 = d.Get(prefix23_1_5 + "link_down_sequence2").(int)
			obj23_1_5.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet3 = d.Get(prefix23_1_5 + "linkdown_ethernet3").(int)
			obj23_1_5.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkDownSequence3 = d.Get(prefix23_1_5 + "link_down_sequence3").(int)
			obj23_1.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet1 = append(obj23_1.SystemInstanceTemplateBindMonitorListLinkDownCfgLinkdownEthernet1, obj23_1_5)
		}

		obj23_1.SystemInstanceTemplateBindMonitorListUserTag = d.Get(prefix23_1 + "user_tag").(string)
		obj23.SystemInstanceTemplateBindMonitorListID = append(obj23.SystemInstanceTemplateBindMonitorListID, obj23_1)
	}

	c.SystemInstanceTemplateBindMonitorList = obj23

	var obj24 go_thunder.SystemInstanceMonTemplate
	prefix24 := "mon_template.0."

	SystemInstanceMonTemplateMonitorListCount := d.Get(prefix24 + "monitor_list.#").(int)
	obj24.SystemInstanceMonTemplateMonitorListID = make([]go_thunder.SystemInstanceMonTemplateMonitorList, 0, SystemInstanceMonTemplateMonitorListCount)

	for i := 0; i < SystemInstanceMonTemplateMonitorListCount; i++ {
		var obj24_1 go_thunder.SystemInstanceMonTemplateMonitorList
		prefix24_1 := prefix24 + fmt.Sprintf("monitor_list.%d.", i)
		obj24_1.SystemInstanceMonTemplateMonitorListID = d.Get(prefix24_1 + "id").(int)

		SystemInstanceMonTemplateMonitorListClearCfgCount := d.Get(prefix24_1 + "clear_cfg.#").(int)
		obj24_1.SystemInstanceMonTemplateMonitorListClearCfgSessions = make([]go_thunder.SystemInstanceMonTemplateMonitorListClearCfg, 0, SystemInstanceMonTemplateMonitorListClearCfgCount)

		for i := 0; i < SystemInstanceMonTemplateMonitorListClearCfgCount; i++ {
			var obj24_1_1 go_thunder.SystemInstanceMonTemplateMonitorListClearCfg
			prefix24_1_1 := prefix24_1 + fmt.Sprintf("clear_cfg.%d.", i)
			obj24_1_1.SystemInstanceMonTemplateMonitorListClearCfgSessions = d.Get(prefix24_1_1 + "sessions").(string)
			obj24_1_1.SystemInstanceMonTemplateMonitorListClearCfgClearAllSequence = d.Get(prefix24_1_1 + "clear_all_sequence").(int)
			obj24_1_1.SystemInstanceMonTemplateMonitorListClearCfgClearSequence = d.Get(prefix24_1_1 + "clear_sequence").(int)
			obj24_1.SystemInstanceMonTemplateMonitorListClearCfgSessions = append(obj24_1.SystemInstanceMonTemplateMonitorListClearCfgSessions, obj24_1_1)
		}

		SystemInstanceMonTemplateMonitorListLinkDisableCfgCount := d.Get(prefix24_1 + "link_disable_cfg.#").(int)
		obj24_1.SystemInstanceMonTemplateMonitorListLinkDisableCfgDiseth = make([]go_thunder.SystemInstanceMonTemplateMonitorListLinkDisableCfg, 0, SystemInstanceMonTemplateMonitorListLinkDisableCfgCount)

		for i := 0; i < SystemInstanceMonTemplateMonitorListLinkDisableCfgCount; i++ {
			var obj24_1_2 go_thunder.SystemInstanceMonTemplateMonitorListLinkDisableCfg
			prefix24_1_2 := prefix24_1 + fmt.Sprintf("link_disable_cfg.%d.", i)
			obj24_1_2.SystemInstanceMonTemplateMonitorListLinkDisableCfgDiseth = d.Get(prefix24_1_2 + "diseth").(int)
			obj24_1_2.SystemInstanceMonTemplateMonitorListLinkDisableCfgDisSequence = d.Get(prefix24_1_2 + "dis_sequence").(int)
			obj24_1.SystemInstanceMonTemplateMonitorListLinkDisableCfgDiseth = append(obj24_1.SystemInstanceMonTemplateMonitorListLinkDisableCfgDiseth, obj24_1_2)
		}

		SystemInstanceMonTemplateMonitorListLinkEnableCfgCount := d.Get(prefix24_1 + "link_enable_cfg.#").(int)
		obj24_1.SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaeth = make([]go_thunder.SystemInstanceMonTemplateMonitorListLinkEnableCfg, 0, SystemInstanceMonTemplateMonitorListLinkEnableCfgCount)

		for i := 0; i < SystemInstanceMonTemplateMonitorListLinkEnableCfgCount; i++ {
			var obj24_1_3 go_thunder.SystemInstanceMonTemplateMonitorListLinkEnableCfg
			prefix24_1_3 := prefix24_1 + fmt.Sprintf("link_enable_cfg.%d.", i)
			obj24_1_3.SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaeth = d.Get(prefix24_1_3 + "enaeth").(int)
			obj24_1_3.SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaSequence = d.Get(prefix24_1_3 + "ena_sequence").(int)
			obj24_1.SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaeth = append(obj24_1.SystemInstanceMonTemplateMonitorListLinkEnableCfgEnaeth, obj24_1_3)
		}

		obj24_1.SystemInstanceMonTemplateMonitorListMonitorRelation = d.Get(prefix24_1 + "monitor_relation").(string)

		SystemInstanceMonTemplateMonitorListLinkUpCfgCount := d.Get(prefix24_1 + "link_up_cfg.#").(int)
		obj24_1.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet1 = make([]go_thunder.SystemInstanceMonTemplateMonitorListLinkUpCfg, 0, SystemInstanceMonTemplateMonitorListLinkUpCfgCount)

		for i := 0; i < SystemInstanceMonTemplateMonitorListLinkUpCfgCount; i++ {
			var obj24_1_4 go_thunder.SystemInstanceMonTemplateMonitorListLinkUpCfg
			prefix24_1_4 := prefix24_1 + fmt.Sprintf("link_up_cfg.%d.", i)
			obj24_1_4.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet1 = d.Get(prefix24_1_4 + "linkup_ethernet1").(int)
			obj24_1_4.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkUpSequence1 = d.Get(prefix24_1_4 + "link_up_sequence1").(int)
			obj24_1_4.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet2 = d.Get(prefix24_1_4 + "linkup_ethernet2").(int)
			obj24_1_4.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkUpSequence2 = d.Get(prefix24_1_4 + "link_up_sequence2").(int)
			obj24_1_4.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet3 = d.Get(prefix24_1_4 + "linkup_ethernet3").(int)
			obj24_1_4.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkUpSequence3 = d.Get(prefix24_1_4 + "link_up_sequence3").(int)
			obj24_1.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet1 = append(obj24_1.SystemInstanceMonTemplateMonitorListLinkUpCfgLinkupEthernet1, obj24_1_4)
		}

		SystemInstanceMonTemplateMonitorListLinkDownCfgCount := d.Get(prefix24_1 + "link_down_cfg.#").(int)
		obj24_1.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet1 = make([]go_thunder.SystemInstanceMonTemplateMonitorListLinkDownCfg, 0, SystemInstanceMonTemplateMonitorListLinkDownCfgCount)

		for i := 0; i < SystemInstanceMonTemplateMonitorListLinkDownCfgCount; i++ {
			var obj24_1_5 go_thunder.SystemInstanceMonTemplateMonitorListLinkDownCfg
			prefix24_1_5 := prefix24_1 + fmt.Sprintf("link_down_cfg.%d.", i)
			obj24_1_5.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet1 = d.Get(prefix24_1_5 + "linkdown_ethernet1").(int)
			obj24_1_5.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkDownSequence1 = d.Get(prefix24_1_5 + "link_down_sequence1").(int)
			obj24_1_5.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet2 = d.Get(prefix24_1_5 + "linkdown_ethernet2").(int)
			obj24_1_5.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkDownSequence2 = d.Get(prefix24_1_5 + "link_down_sequence2").(int)
			obj24_1_5.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet3 = d.Get(prefix24_1_5 + "linkdown_ethernet3").(int)
			obj24_1_5.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkDownSequence3 = d.Get(prefix24_1_5 + "link_down_sequence3").(int)
			obj24_1.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet1 = append(obj24_1.SystemInstanceMonTemplateMonitorListLinkDownCfgLinkdownEthernet1, obj24_1_5)
		}

		obj24_1.SystemInstanceMonTemplateMonitorListUserTag = d.Get(prefix24_1 + "user_tag").(string)
		obj24.SystemInstanceMonTemplateMonitorListID = append(obj24.SystemInstanceMonTemplateMonitorListID, obj24_1)
	}

	var obj24_2 go_thunder.SystemInstanceMonTemplateLinkBlockAsDown
	prefix24_2 := prefix24 + "link_block_as_down.0."
	obj24_2.SystemInstanceMonTemplateLinkBlockAsDownEnable = d.Get(prefix24_2 + "enable").(int)

	obj24.SystemInstanceMonTemplateLinkBlockAsDownEnable = obj24_2

	var obj24_3 go_thunder.SystemInstanceMonTemplateLinkDownOnRestart
	prefix24_3 := prefix24 + "link_down_on_restart.0."
	obj24_3.SystemInstanceMonTemplateLinkDownOnRestartEnable = d.Get(prefix24_3 + "enable").(int)

	obj24.SystemInstanceMonTemplateLinkDownOnRestartEnable = obj24_3

	c.SystemInstanceMonTemplateMonitorList = obj24

	var obj25 go_thunder.SystemInstanceMemory
	prefix25 := "memory.0."

	SystemInstanceMemorySamplingEnableCount := d.Get(prefix25 + "sampling_enable.#").(int)
	obj25.SystemInstanceMemorySamplingEnableCounters1 = make([]go_thunder.SystemInstanceMemorySamplingEnable, 0, SystemInstanceMemorySamplingEnableCount)

	for i := 0; i < SystemInstanceMemorySamplingEnableCount; i++ {
		var obj25_1 go_thunder.SystemInstanceMemorySamplingEnable
		prefix25_1 := prefix25 + fmt.Sprintf("sampling_enable.%d.", i)
		obj25_1.SystemInstanceMemorySamplingEnableCounters1 = d.Get(prefix25_1 + "counters1").(string)
		obj25.SystemInstanceMemorySamplingEnableCounters1 = append(obj25.SystemInstanceMemorySamplingEnableCounters1, obj25_1)
	}

	c.SystemInstanceMemoryUUID = obj25

	var obj26 go_thunder.SystemInstanceResourceUsage
	prefix26 := "resource_usage.0."
	obj26.SystemInstanceResourceUsageSslContextMemory = d.Get(prefix26 + "ssl_context_memory").(int)
	obj26.SystemInstanceResourceUsageSslDmaMemory = d.Get(prefix26 + "ssl_dma_memory").(int)
	obj26.SystemInstanceResourceUsageNatPoolAddrCount = d.Get(prefix26 + "nat_pool_addr_count").(int)
	obj26.SystemInstanceResourceUsageL4SessionCount = d.Get(prefix26 + "l4_session_count").(int)
	obj26.SystemInstanceResourceUsageAuthPortalHTMLFileSize = d.Get(prefix26 + "auth_portal_html_file_size").(int)
	obj26.SystemInstanceResourceUsageAuthPortalImageFileSize = d.Get(prefix26 + "auth_portal_image_file_size").(int)
	obj26.SystemInstanceResourceUsageMaxAflexFileSize = d.Get(prefix26 + "max_aflex_file_size").(int)
	obj26.SystemInstanceResourceUsageAflexTableEntryCount = d.Get(prefix26 + "aflex_table_entry_count").(int)
	obj26.SystemInstanceResourceUsageClassListIpv6AddrCount = d.Get(prefix26 + "class_list_ipv6_addr_count").(int)
	obj26.SystemInstanceResourceUsageClassListAcEntryCount = d.Get(prefix26 + "class_list_ac_entry_count").(int)
	obj26.SystemInstanceResourceUsageClassListEntryCount = d.Get(prefix26 + "class_list_entry_count").(int)
	obj26.SystemInstanceResourceUsageMaxAflexAuthzCollectionNumber = d.Get(prefix26 + "max_aflex_authz_collection_number").(int)
	obj26.SystemInstanceResourceUsageRadiusTableSize = d.Get(prefix26 + "radius_table_size").(int)
	obj26.SystemInstanceResourceUsageAuthzPolicyNumber = d.Get(prefix26 + "authz_policy_number").(int)
	obj26.SystemInstanceResourceUsageIpsecSaNumber = d.Get(prefix26 + "ipsec_sa_number").(int)
	obj26.SystemInstanceResourceUsageRAMCacheMemoryLimit = d.Get(prefix26 + "ram_cache_memory_limit").(int)
	obj26.SystemInstanceResourceUsageWafTemplateCount = d.Get(prefix26 + "waf_template_count").(int)
	obj26.SystemInstanceResourceUsageAuthSessionCount = d.Get(prefix26 + "auth_session_count").(int)

	var obj26_1 go_thunder.SystemInstanceResourceUsageVisibility
	prefix26_1 := prefix26 + "visibility.0."
	obj26_1.SystemInstanceResourceUsageVisibilityMonitoredEntityCount = d.Get(prefix26_1 + "monitored_entity_count").(int)

	obj26.SystemInstanceResourceUsageVisibilityMonitoredEntityCount = obj26_1

	c.SystemInstanceResourceUsageSslContextMemory = obj26

	var obj27 go_thunder.SystemInstanceLinkCapability
	prefix27 := "link_capability.0."
	obj27.SystemInstanceLinkCapabilityEnable = d.Get(prefix27 + "enable").(int)

	c.SystemInstanceLinkCapabilityEnable = obj27

	var obj28 go_thunder.SystemInstanceResourceAccounting
	prefix28 := "resource_accounting.0."

	SystemInstanceResourceAccountingTemplateListCount := d.Get(prefix28 + "template_list.#").(int)
	obj28.SystemInstanceResourceAccountingTemplateListName = make([]go_thunder.SystemInstanceResourceAccountingTemplateList, 0, SystemInstanceResourceAccountingTemplateListCount)

	for i := 0; i < SystemInstanceResourceAccountingTemplateListCount; i++ {
		var obj28_1 go_thunder.SystemInstanceResourceAccountingTemplateList
		prefix28_1 := prefix28 + fmt.Sprintf("template_list.%d.", i)
		obj28_1.SystemInstanceResourceAccountingTemplateListName = d.Get(prefix28_1 + "name").(string)
		obj28_1.SystemInstanceResourceAccountingTemplateListUserTag = d.Get(prefix28_1 + "user_tag").(string)

		var obj28_1_1 go_thunder.SystemInstanceResourceAccountingTemplateListAppResources
		prefix28_1_1 := prefix28_1 + "app_resources.0."

		var obj28_1_1_1 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfg
		prefix28_1_1_1 := prefix28_1_1 + "gslb_device_cfg.0."
		obj28_1_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfgGslbDeviceMax = d.Get(prefix28_1_1_1 + "gslb_device_max").(int)
		obj28_1_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfgGslbDeviceMinGuarantee = d.Get(prefix28_1_1_1 + "gslb_device_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfgGslbDeviceMax = obj28_1_1_1

		var obj28_1_1_2 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfg
		prefix28_1_1_2 := prefix28_1_1 + "gslb_geo_location_cfg.0."
		obj28_1_1_2.SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfgGslbGeoLocationMax = d.Get(prefix28_1_1_2 + "gslb_geo_location_max").(int)
		obj28_1_1_2.SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfgGslbGeoLocationMinGuarantee = d.Get(prefix28_1_1_2 + "gslb_geo_location_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbGeoLocationCfgGslbGeoLocationMax = obj28_1_1_2

		var obj28_1_1_3 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfg
		prefix28_1_1_3 := prefix28_1_1 + "gslb_ip_list_cfg.0."
		obj28_1_1_3.SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfgGslbIPListMax = d.Get(prefix28_1_1_3 + "gslb_ip_list_max").(int)
		obj28_1_1_3.SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfgGslbIPListMinGuarantee = d.Get(prefix28_1_1_3 + "gslb_ip_list_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbIPListCfgGslbIPListMax = obj28_1_1_3

		var obj28_1_1_4 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfg
		prefix28_1_1_4 := prefix28_1_1 + "gslb_policy_cfg.0."
		obj28_1_1_4.SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfgGslbPolicyMax = d.Get(prefix28_1_1_4 + "gslb_policy_max").(int)
		obj28_1_1_4.SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfgGslbPolicyMinGuarantee = d.Get(prefix28_1_1_4 + "gslb_policy_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbPolicyCfgGslbPolicyMax = obj28_1_1_4

		var obj28_1_1_5 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfg
		prefix28_1_1_5 := prefix28_1_1 + "gslb_service_cfg.0."
		obj28_1_1_5.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfgGslbServiceMax = d.Get(prefix28_1_1_5 + "gslb_service_max").(int)
		obj28_1_1_5.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfgGslbServiceMinGuarantee = d.Get(prefix28_1_1_5 + "gslb_service_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceCfgGslbServiceMax = obj28_1_1_5

		var obj28_1_1_6 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfg
		prefix28_1_1_6 := prefix28_1_1 + "gslb_service_ip_cfg.0."
		obj28_1_1_6.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfgGslbServiceIPMax = d.Get(prefix28_1_1_6 + "gslb_service_ip_max").(int)
		obj28_1_1_6.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfgGslbServiceIPMinGuarantee = d.Get(prefix28_1_1_6 + "gslb_service_ip_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServiceIPCfgGslbServiceIPMax = obj28_1_1_6

		var obj28_1_1_7 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfg
		prefix28_1_1_7 := prefix28_1_1 + "gslb_service_port_cfg.0."
		obj28_1_1_7.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfgGslbServicePortMax = d.Get(prefix28_1_1_7 + "gslb_service_port_max").(int)
		obj28_1_1_7.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfgGslbServicePortMinGuarantee = d.Get(prefix28_1_1_7 + "gslb_service_port_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbServicePortCfgGslbServicePortMax = obj28_1_1_7

		var obj28_1_1_8 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfg
		prefix28_1_1_8 := prefix28_1_1 + "gslb_site_cfg.0."
		obj28_1_1_8.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfgGslbSiteMax = d.Get(prefix28_1_1_8 + "gslb_site_max").(int)
		obj28_1_1_8.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfgGslbSiteMinGuarantee = d.Get(prefix28_1_1_8 + "gslb_site_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSiteCfgGslbSiteMax = obj28_1_1_8

		var obj28_1_1_9 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfg
		prefix28_1_1_9 := prefix28_1_1 + "gslb_svc_group_cfg.0."
		obj28_1_1_9.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfgGslbSvcGroupMax = d.Get(prefix28_1_1_9 + "gslb_svc_group_max").(int)
		obj28_1_1_9.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfgGslbSvcGroupMinGuarantee = d.Get(prefix28_1_1_9 + "gslb_svc_group_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbSvcGroupCfgGslbSvcGroupMax = obj28_1_1_9

		var obj28_1_1_10 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfg
		prefix28_1_1_10 := prefix28_1_1 + "gslb_template_cfg.0."
		obj28_1_1_10.SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfgGslbTemplateMax = d.Get(prefix28_1_1_10 + "gslb_template_max").(int)
		obj28_1_1_10.SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfgGslbTemplateMinGuarantee = d.Get(prefix28_1_1_10 + "gslb_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbTemplateCfgGslbTemplateMax = obj28_1_1_10

		var obj28_1_1_11 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfg
		prefix28_1_1_11 := prefix28_1_1 + "gslb_zone_cfg.0."
		obj28_1_1_11.SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfgGslbZoneMax = d.Get(prefix28_1_1_11 + "gslb_zone_max").(int)
		obj28_1_1_11.SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfgGslbZoneMinGuarantee = d.Get(prefix28_1_1_11 + "gslb_zone_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbZoneCfgGslbZoneMax = obj28_1_1_11

		var obj28_1_1_12 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfg
		prefix28_1_1_12 := prefix28_1_1 + "health_monitor_cfg.0."
		obj28_1_1_12.SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfgHealthMonitorMax = d.Get(prefix28_1_1_12 + "health_monitor_max").(int)
		obj28_1_1_12.SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfgHealthMonitorMinGuarantee = d.Get(prefix28_1_1_12 + "health_monitor_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesHealthMonitorCfgHealthMonitorMax = obj28_1_1_12

		var obj28_1_1_13 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfg
		prefix28_1_1_13 := prefix28_1_1 + "real_port_cfg.0."
		obj28_1_1_13.SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfgRealPortMax = d.Get(prefix28_1_1_13 + "real_port_max").(int)
		obj28_1_1_13.SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfgRealPortMinGuarantee = d.Get(prefix28_1_1_13 + "real_port_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesRealPortCfgRealPortMax = obj28_1_1_13

		var obj28_1_1_14 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfg
		prefix28_1_1_14 := prefix28_1_1 + "real_server_cfg.0."
		obj28_1_1_14.SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfgRealServerMax = d.Get(prefix28_1_1_14 + "real_server_max").(int)
		obj28_1_1_14.SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfgRealServerMinGuarantee = d.Get(prefix28_1_1_14 + "real_server_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesRealServerCfgRealServerMax = obj28_1_1_14

		var obj28_1_1_15 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfg
		prefix28_1_1_15 := prefix28_1_1 + "service_group_cfg.0."
		obj28_1_1_15.SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfgServiceGroupMax = d.Get(prefix28_1_1_15 + "service_group_max").(int)
		obj28_1_1_15.SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfgServiceGroupMinGuarantee = d.Get(prefix28_1_1_15 + "service_group_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesServiceGroupCfgServiceGroupMax = obj28_1_1_15

		var obj28_1_1_16 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfg
		prefix28_1_1_16 := prefix28_1_1 + "virtual_server_cfg.0."
		obj28_1_1_16.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfgVirtualServerMax = d.Get(prefix28_1_1_16 + "virtual_server_max").(int)
		obj28_1_1_16.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfgVirtualServerMinGuarantee = d.Get(prefix28_1_1_16 + "virtual_server_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualServerCfgVirtualServerMax = obj28_1_1_16

		var obj28_1_1_17 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfg
		prefix28_1_1_17 := prefix28_1_1 + "virtual_port_cfg.0."
		obj28_1_1_17.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfgVirtualPortMax = d.Get(prefix28_1_1_17 + "virtual_port_max").(int)
		obj28_1_1_17.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfgVirtualPortMinGuarantee = d.Get(prefix28_1_1_17 + "virtual_port_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesVirtualPortCfgVirtualPortMax = obj28_1_1_17

		var obj28_1_1_18 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfg
		prefix28_1_1_18 := prefix28_1_1 + "cache_template_cfg.0."
		obj28_1_1_18.SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfgCacheTemplateMax = d.Get(prefix28_1_1_18 + "cache_template_max").(int)
		obj28_1_1_18.SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfgCacheTemplateMinGuarantee = d.Get(prefix28_1_1_18 + "cache_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesCacheTemplateCfgCacheTemplateMax = obj28_1_1_18

		var obj28_1_1_19 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfg
		prefix28_1_1_19 := prefix28_1_1 + "client_ssl_template_cfg.0."
		obj28_1_1_19.SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfgClientSslTemplateMax = d.Get(prefix28_1_1_19 + "client_ssl_template_max").(int)
		obj28_1_1_19.SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfgClientSslTemplateMinGuarantee = d.Get(prefix28_1_1_19 + "client_ssl_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesClientSslTemplateCfgClientSslTemplateMax = obj28_1_1_19

		var obj28_1_1_20 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfg
		prefix28_1_1_20 := prefix28_1_1 + "conn_reuse_template_cfg.0."
		obj28_1_1_20.SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfgConnReuseTemplateMax = d.Get(prefix28_1_1_20 + "conn_reuse_template_max").(int)
		obj28_1_1_20.SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfgConnReuseTemplateMinGuarantee = d.Get(prefix28_1_1_20 + "conn_reuse_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesConnReuseTemplateCfgConnReuseTemplateMax = obj28_1_1_20

		var obj28_1_1_21 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfg
		prefix28_1_1_21 := prefix28_1_1 + "fast_tcp_template_cfg.0."
		obj28_1_1_21.SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfgFastTCPTemplateMax = d.Get(prefix28_1_1_21 + "fast_tcp_template_max").(int)
		obj28_1_1_21.SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfgFastTCPTemplateMinGuarantee = d.Get(prefix28_1_1_21 + "fast_tcp_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesFastTCPTemplateCfgFastTCPTemplateMax = obj28_1_1_21

		var obj28_1_1_22 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfg
		prefix28_1_1_22 := prefix28_1_1 + "fast_udp_template_cfg.0."
		obj28_1_1_22.SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfgFastUDPTemplateMax = d.Get(prefix28_1_1_22 + "fast_udp_template_max").(int)
		obj28_1_1_22.SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfgFastUDPTemplateMinGuarantee = d.Get(prefix28_1_1_22 + "fast_udp_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesFastUDPTemplateCfgFastUDPTemplateMax = obj28_1_1_22

		var obj28_1_1_23 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfg
		prefix28_1_1_23 := prefix28_1_1 + "fix_template_cfg.0."
		obj28_1_1_23.SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfgFixTemplateMax = d.Get(prefix28_1_1_23 + "fix_template_max").(int)
		obj28_1_1_23.SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfgFixTemplateMinGuarantee = d.Get(prefix28_1_1_23 + "fix_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesFixTemplateCfgFixTemplateMax = obj28_1_1_23

		var obj28_1_1_24 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfg
		prefix28_1_1_24 := prefix28_1_1 + "http_template_cfg.0."
		obj28_1_1_24.SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfgHTTPTemplateMax = d.Get(prefix28_1_1_24 + "http_template_max").(int)
		obj28_1_1_24.SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfgHTTPTemplateMinGuarantee = d.Get(prefix28_1_1_24 + "http_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesHTTPTemplateCfgHTTPTemplateMax = obj28_1_1_24

		var obj28_1_1_25 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfg
		prefix28_1_1_25 := prefix28_1_1 + "link_cost_template_cfg.0."
		obj28_1_1_25.SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfgLinkCostTemplateMax = d.Get(prefix28_1_1_25 + "link_cost_template_max").(int)
		obj28_1_1_25.SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfgLinkCostTemplateMinGuarantee = d.Get(prefix28_1_1_25 + "link_cost_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesLinkCostTemplateCfgLinkCostTemplateMax = obj28_1_1_25

		var obj28_1_1_26 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfg
		prefix28_1_1_26 := prefix28_1_1 + "persist_cookie_template_cfg.0."
		obj28_1_1_26.SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfgPersistCookieTemplateMax = d.Get(prefix28_1_1_26 + "persist_cookie_template_max").(int)
		obj28_1_1_26.SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfgPersistCookieTemplateMinGuarantee = d.Get(prefix28_1_1_26 + "persist_cookie_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesPersistCookieTemplateCfgPersistCookieTemplateMax = obj28_1_1_26

		var obj28_1_1_27 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfg
		prefix28_1_1_27 := prefix28_1_1 + "persist_srcip_template_cfg.0."
		obj28_1_1_27.SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfgPersistSrcipTemplateMax = d.Get(prefix28_1_1_27 + "persist_srcip_template_max").(int)
		obj28_1_1_27.SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfgPersistSrcipTemplateMinGuarantee = d.Get(prefix28_1_1_27 + "persist_srcip_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesPersistSrcipTemplateCfgPersistSrcipTemplateMax = obj28_1_1_27

		var obj28_1_1_28 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfg
		prefix28_1_1_28 := prefix28_1_1 + "server_ssl_template_cfg.0."
		obj28_1_1_28.SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfgServerSslTemplateMax = d.Get(prefix28_1_1_28 + "server_ssl_template_max").(int)
		obj28_1_1_28.SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfgServerSslTemplateMinGuarantee = d.Get(prefix28_1_1_28 + "server_ssl_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesServerSslTemplateCfgServerSslTemplateMax = obj28_1_1_28

		var obj28_1_1_29 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfg
		prefix28_1_1_29 := prefix28_1_1 + "proxy_template_cfg.0."
		obj28_1_1_29.SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfgProxyTemplateMax = d.Get(prefix28_1_1_29 + "proxy_template_max").(int)
		obj28_1_1_29.SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfgProxyTemplateMinGuarantee = d.Get(prefix28_1_1_29 + "proxy_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesProxyTemplateCfgProxyTemplateMax = obj28_1_1_29

		var obj28_1_1_30 go_thunder.SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfg
		prefix28_1_1_30 := prefix28_1_1 + "stream_template_cfg.0."
		obj28_1_1_30.SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfgStreamTemplateMax = d.Get(prefix28_1_1_30 + "stream_template_max").(int)
		obj28_1_1_30.SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfgStreamTemplateMinGuarantee = d.Get(prefix28_1_1_30 + "stream_template_min_guarantee").(int)

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesStreamTemplateCfgStreamTemplateMax = obj28_1_1_30

		obj28_1_1.SystemInstanceResourceAccountingTemplateListAppResourcesThreshold = d.Get(prefix28_1_1 + "threshold").(int)

		obj28_1.SystemInstanceResourceAccountingTemplateListAppResourcesGslbDeviceCfg = obj28_1_1

		var obj28_1_2 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResources
		prefix28_1_2 := prefix28_1 + "network_resources.0."

		var obj28_1_2_1 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg
		prefix28_1_2_1 := prefix28_1_2 + "static_ipv4_route_cfg.0."
		obj28_1_2_1.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfgStaticIpv4RouteMax = d.Get(prefix28_1_2_1 + "static_ipv4_route_max").(int)
		obj28_1_2_1.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfgStaticIpv4RouteMinGuarantee = d.Get(prefix28_1_2_1 + "static_ipv4_route_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfgStaticIpv4RouteMax = obj28_1_2_1

		var obj28_1_2_2 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfg
		prefix28_1_2_2 := prefix28_1_2 + "static_ipv6_route_cfg.0."
		obj28_1_2_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfgStaticIpv6RouteMax = d.Get(prefix28_1_2_2 + "static_ipv6_route_max").(int)
		obj28_1_2_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfgStaticIpv6RouteMinGuarantee = d.Get(prefix28_1_2_2 + "static_ipv6_route_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv6RouteCfgStaticIpv6RouteMax = obj28_1_2_2

		var obj28_1_2_3 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfg
		prefix28_1_2_3 := prefix28_1_2 + "ipv4_acl_line_cfg.0."
		obj28_1_2_3.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfgIpv4AclLineMax = d.Get(prefix28_1_2_3 + "ipv4_acl_line_max").(int)
		obj28_1_2_3.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfgIpv4AclLineMinGuarantee = d.Get(prefix28_1_2_3 + "ipv4_acl_line_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv4AclLineCfgIpv4AclLineMax = obj28_1_2_3

		var obj28_1_2_4 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfg
		prefix28_1_2_4 := prefix28_1_2 + "ipv6_acl_line_cfg.0."
		obj28_1_2_4.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfgIpv6AclLineMax = d.Get(prefix28_1_2_4 + "ipv6_acl_line_max").(int)
		obj28_1_2_4.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfgIpv6AclLineMinGuarantee = d.Get(prefix28_1_2_4 + "ipv6_acl_line_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesIpv6AclLineCfgIpv6AclLineMax = obj28_1_2_4

		var obj28_1_2_5 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfg
		prefix28_1_2_5 := prefix28_1_2 + "static_arp_cfg.0."
		obj28_1_2_5.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfgStaticArpMax = d.Get(prefix28_1_2_5 + "static_arp_max").(int)
		obj28_1_2_5.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfgStaticArpMinGuarantee = d.Get(prefix28_1_2_5 + "static_arp_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticArpCfgStaticArpMax = obj28_1_2_5

		var obj28_1_2_6 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfg
		prefix28_1_2_6 := prefix28_1_2 + "static_neighbor_cfg.0."
		obj28_1_2_6.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfgStaticNeighborMax = d.Get(prefix28_1_2_6 + "static_neighbor_max").(int)
		obj28_1_2_6.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfgStaticNeighborMinGuarantee = d.Get(prefix28_1_2_6 + "static_neighbor_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticNeighborCfgStaticNeighborMax = obj28_1_2_6

		var obj28_1_2_7 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfg
		prefix28_1_2_7 := prefix28_1_2 + "static_mac_cfg.0."
		obj28_1_2_7.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfgStaticMacMax = d.Get(prefix28_1_2_7 + "static_mac_max").(int)
		obj28_1_2_7.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfgStaticMacMinGuarantee = d.Get(prefix28_1_2_7 + "static_mac_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticMacCfgStaticMacMax = obj28_1_2_7

		var obj28_1_2_8 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfg
		prefix28_1_2_8 := prefix28_1_2 + "object_group_cfg.0."
		obj28_1_2_8.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfgObjectGroupMax = d.Get(prefix28_1_2_8 + "object_group_max").(int)
		obj28_1_2_8.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfgObjectGroupMinGuarantee = d.Get(prefix28_1_2_8 + "object_group_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupCfgObjectGroupMax = obj28_1_2_8

		var obj28_1_2_9 go_thunder.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfg
		prefix28_1_2_9 := prefix28_1_2 + "object_group_clause_cfg.0."
		obj28_1_2_9.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfgObjectGroupClauseMax = d.Get(prefix28_1_2_9 + "object_group_clause_max").(int)
		obj28_1_2_9.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfgObjectGroupClauseMinGuarantee = d.Get(prefix28_1_2_9 + "object_group_clause_min_guarantee").(int)

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesObjectGroupClauseCfgObjectGroupClauseMax = obj28_1_2_9

		obj28_1_2.SystemInstanceResourceAccountingTemplateListNetworkResourcesThreshold = d.Get(prefix28_1_2 + "threshold").(int)

		obj28_1.SystemInstanceResourceAccountingTemplateListNetworkResourcesStaticIpv4RouteCfg = obj28_1_2

		var obj28_1_3 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResources
		prefix28_1_3 := prefix28_1 + "system_resources.0."

		var obj28_1_3_1 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfg
		prefix28_1_3_1 := prefix28_1_3 + "bw_limit_cfg.0."
		obj28_1_3_1.SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfgBwLimitMax = d.Get(prefix28_1_3_1 + "bw_limit_max").(int)
		obj28_1_3_1.SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfgBwLimitWatermarkDisable = d.Get(prefix28_1_3_1 + "bw_limit_watermark_disable").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfgBwLimitMax = obj28_1_3_1

		var obj28_1_3_2 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfg
		prefix28_1_3_2 := prefix28_1_3 + "concurrent_session_limit_cfg.0."
		obj28_1_3_2.SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfgConcurrentSessionLimitMax = d.Get(prefix28_1_3_2 + "concurrent_session_limit_max").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesConcurrentSessionLimitCfgConcurrentSessionLimitMax = obj28_1_3_2

		var obj28_1_3_3 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfg
		prefix28_1_3_3 := prefix28_1_3 + "l4_session_limit_cfg.0."
		obj28_1_3_3.SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfgL4SessionLimitMax = d.Get(prefix28_1_3_3 + "l4_session_limit_max").(string)
		obj28_1_3_3.SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfgL4SessionLimitMinGuarantee = d.Get(prefix28_1_3_3 + "l4_session_limit_min_guarantee").(string)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesL4SessionLimitCfgL4SessionLimitMax = obj28_1_3_3

		var obj28_1_3_4 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfg
		prefix28_1_3_4 := prefix28_1_3 + "l4cps_limit_cfg.0."
		obj28_1_3_4.SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfgL4CpsLimitMax = d.Get(prefix28_1_3_4 + "l4cps_limit_max").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesL4CpsLimitCfgL4CpsLimitMax = obj28_1_3_4

		var obj28_1_3_5 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfg
		prefix28_1_3_5 := prefix28_1_3 + "l7cps_limit_cfg.0."
		obj28_1_3_5.SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfgL7CpsLimitMax = d.Get(prefix28_1_3_5 + "l7cps_limit_max").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesL7CpsLimitCfgL7CpsLimitMax = obj28_1_3_5

		var obj28_1_3_6 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfg
		prefix28_1_3_6 := prefix28_1_3 + "natcps_limit_cfg.0."
		obj28_1_3_6.SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfgNatcpsLimitMax = d.Get(prefix28_1_3_6 + "natcps_limit_max").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesNatcpsLimitCfgNatcpsLimitMax = obj28_1_3_6

		var obj28_1_3_7 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfg
		prefix28_1_3_7 := prefix28_1_3 + "fwcps_limit_cfg.0."
		obj28_1_3_7.SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfgFwcpsLimitMax = d.Get(prefix28_1_3_7 + "fwcps_limit_max").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesFwcpsLimitCfgFwcpsLimitMax = obj28_1_3_7

		var obj28_1_3_8 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfg
		prefix28_1_3_8 := prefix28_1_3 + "ssl_throughput_limit_cfg.0."
		obj28_1_3_8.SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfgSslThroughputLimitMax = d.Get(prefix28_1_3_8 + "ssl_throughput_limit_max").(int)
		obj28_1_3_8.SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfgSslThroughputLimitWatermarkDisable = d.Get(prefix28_1_3_8 + "ssl_throughput_limit_watermark_disable").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesSslThroughputLimitCfgSslThroughputLimitMax = obj28_1_3_8

		var obj28_1_3_9 go_thunder.SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfg
		prefix28_1_3_9 := prefix28_1_3 + "sslcps_limit_cfg.0."
		obj28_1_3_9.SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfgSslcpsLimitMax = d.Get(prefix28_1_3_9 + "sslcps_limit_max").(int)

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesSslcpsLimitCfgSslcpsLimitMax = obj28_1_3_9

		obj28_1_3.SystemInstanceResourceAccountingTemplateListSystemResourcesThreshold = d.Get(prefix28_1_3 + "threshold").(int)

		obj28_1.SystemInstanceResourceAccountingTemplateListSystemResourcesBwLimitCfg = obj28_1_3

		obj28.SystemInstanceResourceAccountingTemplateListName = append(obj28.SystemInstanceResourceAccountingTemplateListName, obj28_1)
	}

	c.SystemInstanceResourceAccountingUUID = obj28

	var obj29 go_thunder.SystemInstanceTrunk
	prefix29 := "trunk.0."

	var obj29_1 go_thunder.SystemInstanceTrunkLoadBalance
	prefix29_1 := prefix29 + "load_balance.0."
	obj29_1.SystemInstanceTrunkLoadBalanceUseL3 = d.Get(prefix29_1 + "use_l3").(int)
	obj29_1.SystemInstanceTrunkLoadBalanceUseL4 = d.Get(prefix29_1 + "use_l4").(int)

	obj29.SystemInstanceTrunkLoadBalanceUseL3 = obj29_1

	c.SystemInstanceTrunkLoadBalance = obj29

	var obj30 go_thunder.SystemInstancePorts
	prefix30 := "ports.0."
	obj30.SystemInstancePortsLinkDetectionInterval = d.Get(prefix30 + "link_detection_interval").(int)

	c.SystemInstancePortsLinkDetectionInterval = obj30

	var obj31 go_thunder.SystemInstanceTableIntegrity
	prefix31 := "table_integrity.0."
	obj31.SystemInstanceTableIntegrityTable = d.Get(prefix31 + "table").(string)
	obj31.SystemInstanceTableIntegrityAuditAction = d.Get(prefix31 + "audit_action").(string)
	obj31.SystemInstanceTableIntegrityAutoSyncAction = d.Get(prefix31 + "auto_sync_action").(string)

	SystemInstanceTableIntegritySamplingEnableCount := d.Get(prefix31 + "sampling_enable.#").(int)
	obj31.SystemInstanceTableIntegritySamplingEnableCounters1 = make([]go_thunder.SystemInstanceTableIntegritySamplingEnable, 0, SystemInstanceTableIntegritySamplingEnableCount)

	for i := 0; i < SystemInstanceTableIntegritySamplingEnableCount; i++ {
		var obj31_1 go_thunder.SystemInstanceTableIntegritySamplingEnable
		prefix31_1 := prefix31 + fmt.Sprintf("sampling_enable.%d.", i)
		obj31_1.SystemInstanceTableIntegritySamplingEnableCounters1 = d.Get(prefix31_1 + "counters1").(string)
		obj31.SystemInstanceTableIntegritySamplingEnableCounters1 = append(obj31.SystemInstanceTableIntegritySamplingEnableCounters1, obj31_1)
	}

	c.SystemInstanceTableIntegrityTable = obj31

	var obj32 go_thunder.SystemInstanceIpsec
	prefix32 := "ipsec.0."
	obj32.SystemInstanceIpsecPacketRoundRobin = d.Get(prefix32 + "packet_round_robin").(int)
	obj32.SystemInstanceIpsecCryptoCore = d.Get(prefix32 + "crypto_core").(int)
	obj32.SystemInstanceIpsecCryptoMem = d.Get(prefix32 + "crypto_mem").(int)
	obj32.SystemInstanceIpsecQAT = d.Get(prefix32 + "qat").(int)

	var obj32_1 go_thunder.SystemInstanceIpsecFpgaDecrypt
	prefix32_1 := prefix32 + "fpga_decrypt.0."
	obj32_1.SystemInstanceIpsecFpgaDecryptAction = d.Get(prefix32_1 + "action").(string)

	obj32.SystemInstanceIpsecFpgaDecryptAction = obj32_1

	c.SystemInstanceIpsecPacketRoundRobin = obj32

	var obj33 go_thunder.SystemInstanceSpeProfile
	prefix33 := "spe_profile.0."
	obj33.SystemInstanceSpeProfileAction = d.Get(prefix33 + "action").(string)

	c.SystemInstanceSpeProfileAction = obj33

	var obj34 go_thunder.SystemInstanceDeepHrxq
	prefix34 := "deep_hrxq.0."
	obj34.SystemInstanceDeepHrxqEnable = d.Get(prefix34 + "enable").(int)

	c.SystemInstanceDeepHrxqEnable = obj34

	var obj35 go_thunder.SystemInstanceCPULoadSharing
	prefix35 := "cpu_load_sharing.0."
	obj35.SystemInstanceCPULoadSharingDisable = d.Get(prefix35 + "disable").(int)

	var obj35_1 go_thunder.SystemInstanceCPULoadSharingPacketsPerSecond
	prefix35_1 := prefix35 + "packets_per_second.0."
	obj35_1.SystemInstanceCPULoadSharingPacketsPerSecondMin = d.Get(prefix35_1 + "min").(int)

	obj35.SystemInstanceCPULoadSharingPacketsPerSecondMin = obj35_1

	var obj35_2 go_thunder.SystemInstanceCPULoadSharingCPUUsage
	prefix35_2 := prefix35 + "cpu_usage.0."
	obj35_2.SystemInstanceCPULoadSharingCPUUsageLow = d.Get(prefix35_2 + "low").(int)
	obj35_2.SystemInstanceCPULoadSharingCPUUsageHigh = d.Get(prefix35_2 + "high").(int)

	obj35.SystemInstanceCPULoadSharingCPUUsageLow = obj35_2

	obj35.SystemInstanceCPULoadSharingTCP = d.Get(prefix35 + "tcp").(int)
	obj35.SystemInstanceCPULoadSharingUDP = d.Get(prefix35 + "udp").(int)

	c.SystemInstanceCPULoadSharingDisable = obj35

	var obj36 go_thunder.SystemInstancePerVlanLimit
	prefix36 := "per_vlan_limit.0."
	obj36.SystemInstancePerVlanLimitBcast = d.Get(prefix36 + "bcast").(int)
	obj36.SystemInstancePerVlanLimitIpmcast = d.Get(prefix36 + "ipmcast").(int)
	obj36.SystemInstancePerVlanLimitMcast = d.Get(prefix36 + "mcast").(int)
	obj36.SystemInstancePerVlanLimitUnknownUcast = d.Get(prefix36 + "unknown_ucast").(int)

	c.SystemInstancePerVlanLimitBcast = obj36

	var obj37 go_thunder.SystemInstanceAllVlanLimit
	prefix37 := "all_vlan_limit.0."
	obj37.SystemInstanceAllVlanLimitBcast = d.Get(prefix37 + "bcast").(int)
	obj37.SystemInstanceAllVlanLimitIpmcast = d.Get(prefix37 + "ipmcast").(int)
	obj37.SystemInstanceAllVlanLimitMcast = d.Get(prefix37 + "mcast").(int)
	obj37.SystemInstanceAllVlanLimitUnknownUcast = d.Get(prefix37 + "unknown_ucast").(int)

	c.SystemInstanceAllVlanLimitBcast = obj37

	var obj38 go_thunder.SystemInstanceVeMacScheme
	prefix38 := "ve_mac_scheme.0."
	obj38.SystemInstanceVeMacSchemeVeMacSchemeVal = d.Get(prefix38 + "ve_mac_scheme_val").(string)

	c.SystemInstanceVeMacSchemeVeMacSchemeVal = obj38

	var obj39 go_thunder.SystemInstanceSessionReclaimLimit
	prefix39 := "session_reclaim_limit.0."
	obj39.SystemInstanceSessionReclaimLimitNscanLimit = d.Get(prefix39 + "nscan_limit").(int)
	obj39.SystemInstanceSessionReclaimLimitScanFreq = d.Get(prefix39 + "scan_freq").(int)

	c.SystemInstanceSessionReclaimLimitNscanLimit = obj39

	var obj40 go_thunder.SystemInstanceSslScv
	prefix40 := "ssl_scv.0."
	obj40.SystemInstanceSslScvEnable = d.Get(prefix40 + "enable").(int)

	c.SystemInstanceSslScvEnable = obj40

	var obj41 go_thunder.SystemInstanceSslScvVerifyHost
	prefix41 := "ssl_scv_verify_host.0."
	obj41.SystemInstanceSslScvVerifyHostDisable = d.Get(prefix41 + "disable").(int)

	c.SystemInstanceSslScvVerifyHostDisable = obj41

	var obj42 go_thunder.SystemInstanceSslScvVerifyCrlSign
	prefix42 := "ssl_scv_verify_crl_sign.0."
	obj42.SystemInstanceSslScvVerifyCrlSignEnable = d.Get(prefix42 + "enable").(int)

	c.SystemInstanceSslScvVerifyCrlSignEnable = obj42

	var obj43 go_thunder.SystemInstanceSslSetCompatibleCipher
	prefix43 := "ssl_set_compatible_cipher.0."
	obj43.SystemInstanceSslSetCompatibleCipherDisable = d.Get(prefix43 + "disable").(int)

	c.SystemInstanceSslSetCompatibleCipherDisable = obj43

	var obj44 go_thunder.SystemInstanceHardwareForward
	prefix44 := "hardware_forward.0."

	SystemInstanceHardwareForwardSamplingEnableCount := d.Get(prefix44 + "sampling_enable.#").(int)
	obj44.SystemInstanceHardwareForwardSamplingEnableCounters1 = make([]go_thunder.SystemInstanceHardwareForwardSamplingEnable, 0, SystemInstanceHardwareForwardSamplingEnableCount)

	for i := 0; i < SystemInstanceHardwareForwardSamplingEnableCount; i++ {
		var obj44_1 go_thunder.SystemInstanceHardwareForwardSamplingEnable
		prefix44_1 := prefix44 + fmt.Sprintf("sampling_enable.%d.", i)
		obj44_1.SystemInstanceHardwareForwardSamplingEnableCounters1 = d.Get(prefix44_1 + "counters1").(string)
		obj44.SystemInstanceHardwareForwardSamplingEnableCounters1 = append(obj44.SystemInstanceHardwareForwardSamplingEnableCounters1, obj44_1)
	}

	var obj44_2 go_thunder.SystemInstanceHardwareForwardSlb
	prefix44_2 := prefix44 + "slb.0."

	SystemInstanceHardwareForwardSlbSamplingEnableCount := d.Get(prefix44_2 + "sampling_enable.#").(int)
	obj44_2.SystemInstanceHardwareForwardSlbSamplingEnableCounters1 = make([]go_thunder.SystemInstanceHardwareForwardSlbSamplingEnable, 0, SystemInstanceHardwareForwardSlbSamplingEnableCount)

	for i := 0; i < SystemInstanceHardwareForwardSlbSamplingEnableCount; i++ {
		var obj44_2_1 go_thunder.SystemInstanceHardwareForwardSlbSamplingEnable
		prefix44_2_1 := prefix44_2 + fmt.Sprintf("sampling_enable.%d.", i)
		obj44_2_1.SystemInstanceHardwareForwardSlbSamplingEnableCounters1 = d.Get(prefix44_2_1 + "counters1").(string)
		obj44_2.SystemInstanceHardwareForwardSlbSamplingEnableCounters1 = append(obj44_2.SystemInstanceHardwareForwardSlbSamplingEnableCounters1, obj44_2_1)
	}

	obj44.SystemInstanceHardwareForwardSlbUUID = obj44_2

	c.SystemInstanceHardwareForwardUUID = obj44

	var obj45 go_thunder.SystemInstanceThroughput
	prefix45 := "throughput.0."

	SystemInstanceThroughputSamplingEnableCount := d.Get(prefix45 + "sampling_enable.#").(int)
	obj45.SystemInstanceThroughputSamplingEnableCounters1 = make([]go_thunder.SystemInstanceThroughputSamplingEnable, 0, SystemInstanceThroughputSamplingEnableCount)

	for i := 0; i < SystemInstanceThroughputSamplingEnableCount; i++ {
		var obj45_1 go_thunder.SystemInstanceThroughputSamplingEnable
		prefix45_1 := prefix45 + fmt.Sprintf("sampling_enable.%d.", i)
		obj45_1.SystemInstanceThroughputSamplingEnableCounters1 = d.Get(prefix45_1 + "counters1").(string)
		obj45.SystemInstanceThroughputSamplingEnableCounters1 = append(obj45.SystemInstanceThroughputSamplingEnableCounters1, obj45_1)
	}

	c.SystemInstanceThroughputUUID = obj45

	var obj46 go_thunder.SystemInstanceIpmi
	prefix46 := "ipmi.0."
	obj46.SystemInstanceIpmiReset = d.Get(prefix46 + "reset").(int)

	var obj46_1 go_thunder.SystemInstanceIpmiIP
	prefix46_1 := prefix46 + "ip.0."
	obj46_1.SystemInstanceIpmiIPIpv4Address = d.Get(prefix46_1 + "ipv4_address").(string)
	obj46_1.SystemInstanceIpmiIPIpv4Netmask = d.Get(prefix46_1 + "ipv4_netmask").(string)
	obj46_1.SystemInstanceIpmiIPDefaultGateway = d.Get(prefix46_1 + "default_gateway").(string)

	obj46.SystemInstanceIpmiIPIpv4Address = obj46_1

	var obj46_2 go_thunder.SystemInstanceIpmiIpsrc
	prefix46_2 := prefix46 + "ipsrc.0."
	obj46_2.SystemInstanceIpmiIpsrcDhcp = d.Get(prefix46_2 + "dhcp").(int)
	obj46_2.SystemInstanceIpmiIpsrcStatic = d.Get(prefix46_2 + "static").(int)

	obj46.SystemInstanceIpmiIpsrcDhcp = obj46_2

	var obj46_3 go_thunder.SystemInstanceIpmiUser
	prefix46_3 := prefix46 + "user.0."
	obj46_3.SystemInstanceIpmiUserAdd = d.Get(prefix46_3 + "add").(string)
	obj46_3.SystemInstanceIpmiUserPassword = d.Get(prefix46_3 + "password").(string)
	obj46_3.SystemInstanceIpmiUserAdministrator = d.Get(prefix46_3 + "administrator").(int)
	obj46_3.SystemInstanceIpmiUserCallback = d.Get(prefix46_3 + "callback").(int)
	obj46_3.SystemInstanceIpmiUserOperator = d.Get(prefix46_3 + "operator").(int)
	obj46_3.SystemInstanceIpmiUserUser = d.Get(prefix46_3 + "user").(int)
	obj46_3.SystemInstanceIpmiUserDisable = d.Get(prefix46_3 + "disable").(string)
	obj46_3.SystemInstanceIpmiUserPrivilege = d.Get(prefix46_3 + "privilege").(string)
	obj46_3.SystemInstanceIpmiUserSetname = d.Get(prefix46_3 + "setname").(string)
	obj46_3.SystemInstanceIpmiUserNewname = d.Get(prefix46_3 + "newname").(string)
	obj46_3.SystemInstanceIpmiUserSetpass = d.Get(prefix46_3 + "setpass").(string)
	obj46_3.SystemInstanceIpmiUserNewpass = d.Get(prefix46_3 + "newpass").(string)

	obj46.SystemInstanceIpmiUserAdd = obj46_3

	var obj46_4 go_thunder.SystemInstanceIpmiTool
	prefix46_4 := prefix46 + "tool.0."
	obj46_4.SystemInstanceIpmiToolCmd = d.Get(prefix46_4 + "cmd").(string)

	obj46.SystemInstanceIpmiToolCmd = obj46_4

	c.SystemInstanceIpmiReset = obj46

	var obj47 go_thunder.SystemInstanceQueuingBuffer
	prefix47 := "queuing_buffer.0."
	obj47.SystemInstanceQueuingBufferEnable = d.Get(prefix47 + "enable").(int)

	c.SystemInstanceQueuingBufferEnable = obj47

	var obj48 go_thunder.SystemInstanceHighMemoryL4Session
	prefix48 := "high_memory_l4_session.0."
	obj48.SystemInstanceHighMemoryL4SessionEnable = d.Get(prefix48 + "enable").(int)

	c.SystemInstanceHighMemoryL4SessionEnable = obj48

	var obj49 go_thunder.SystemInstanceTrunkHwHash
	prefix49 := "trunk_hw_hash.0."
	obj49.SystemInstanceTrunkHwHashMode = d.Get(prefix49 + "mode").(int)

	c.SystemInstanceTrunkHwHashMode = obj49

	var obj50 go_thunder.SystemInstanceTrunkXauiHwHash
	prefix50 := "trunk_xaui_hw_hash.0."
	obj50.SystemInstanceTrunkXauiHwHashMode = d.Get(prefix50 + "mode").(int)

	c.SystemInstanceTrunkXauiHwHashMode = obj50

	var obj51 go_thunder.SystemInstanceCmUpdateFileNameRef
	prefix51 := "cm_update_file_name_ref.0."
	obj51.SystemInstanceCmUpdateFileNameRefSourceName = d.Get(prefix51 + "source_name").(string)
	obj51.SystemInstanceCmUpdateFileNameRefDestName = d.Get(prefix51 + "dest_name").(string)
	obj51.SystemInstanceCmUpdateFileNameRefID = d.Get(prefix51 + "id").(int)

	c.SystemInstanceCmUpdateFileNameRefSourceName = obj51

	var obj52 go_thunder.SystemInstanceAppsGlobal
	prefix52 := "apps_global.0."
	obj52.SystemInstanceAppsGlobalLogSessionOnEstablished = d.Get(prefix52 + "log_session_on_established").(int)
	obj52.SystemInstanceAppsGlobalMslTime = d.Get(prefix52 + "msl_time").(int)

	c.SystemInstanceAppsGlobalLogSessionOnEstablished = obj52

	var obj53 go_thunder.SystemInstanceShellPrivileges
	prefix53 := "shell_privileges.0."
	obj53.SystemInstanceShellPrivilegesEnableShellPrivileges = d.Get(prefix53 + "enable_shell_privileges").(int)

	c.SystemInstanceShellPrivilegesEnableShellPrivileges = obj53

	var obj54 go_thunder.SystemInstanceShmLogging
	prefix54 := "shm_logging.0."
	obj54.SystemInstanceShmLoggingEnable = d.Get(prefix54 + "enable").(int)

	c.SystemInstanceShmLoggingEnable = obj54

	var obj55 go_thunder.SystemInstanceFw
	prefix55 := "fw.0."
	obj55.SystemInstanceFwApplicationMempool = d.Get(prefix55 + "application_mempool").(int)
	obj55.SystemInstanceFwApplicationFlow = d.Get(prefix55 + "application_flow").(int)
	obj55.SystemInstanceFwBasicDpiEnable = d.Get(prefix55 + "basic_dpi_enable").(int)

	c.SystemInstanceFwApplicationMempool = obj55

	var obj56 go_thunder.SystemInstancePasswordPolicy
	prefix56 := "password_policy.0."
	obj56.SystemInstancePasswordPolicyComplexity = d.Get(prefix56 + "complexity").(string)
	obj56.SystemInstancePasswordPolicyAging = d.Get(prefix56 + "aging").(string)
	obj56.SystemInstancePasswordPolicyHistory = d.Get(prefix56 + "history").(string)
	obj56.SystemInstancePasswordPolicyMinPswdLen = d.Get(prefix56 + "min_pswd_len").(int)

	c.SystemInstancePasswordPolicyComplexity = obj56

	var obj57 go_thunder.SystemInstanceRadius
	prefix57 := "radius.0."

	var obj57_1 go_thunder.SystemInstanceRadiusServer
	prefix57_1 := prefix57 + "server.0."
	obj57_1.SystemInstanceRadiusServerListenPort = d.Get(prefix57_1 + "listen_port").(int)

	var obj57_1_1 go_thunder.SystemInstanceRadiusServerRemote
	prefix57_1_1 := prefix57_1 + "remote.0."

	SystemInstanceRadiusServerRemoteIPListCount := d.Get(prefix57_1_1 + "ip_list.#").(int)
	obj57_1_1.SystemInstanceRadiusServerRemoteIPListIPListName = make([]go_thunder.SystemInstanceRadiusServerRemoteIPList, 0, SystemInstanceRadiusServerRemoteIPListCount)

	for i := 0; i < SystemInstanceRadiusServerRemoteIPListCount; i++ {
		var obj57_1_1_1 go_thunder.SystemInstanceRadiusServerRemoteIPList
		prefix57_1_1_1 := prefix57_1_1 + fmt.Sprintf("ip_list.%d.", i)
		obj57_1_1_1.SystemInstanceRadiusServerRemoteIPListIPListName = d.Get(prefix57_1_1_1 + "ip_list_name").(string)
		obj57_1_1_1.SystemInstanceRadiusServerRemoteIPListIPListSecret = d.Get(prefix57_1_1_1 + "ip_list_secret").(int)
		obj57_1_1_1.SystemInstanceRadiusServerRemoteIPListIPListSecretString = d.Get(prefix57_1_1_1 + "ip_list_secret_string").(string)
		obj57_1_1.SystemInstanceRadiusServerRemoteIPListIPListName = append(obj57_1_1.SystemInstanceRadiusServerRemoteIPListIPListName, obj57_1_1_1)
	}

	obj57_1.SystemInstanceRadiusServerRemoteIPList = obj57_1_1

	obj57_1.SystemInstanceRadiusServerSecret = d.Get(prefix57_1 + "secret").(int)
	obj57_1.SystemInstanceRadiusServerSecretString = d.Get(prefix57_1 + "secret_string").(string)
	obj57_1.SystemInstanceRadiusServerVrid = d.Get(prefix57_1 + "vrid").(int)

	SystemInstanceRadiusServerAttributeCount := d.Get(prefix57_1 + "attribute.#").(int)
	obj57_1.SystemInstanceRadiusServerAttributeAttributeValue = make([]go_thunder.SystemInstanceRadiusServerAttribute, 0, SystemInstanceRadiusServerAttributeCount)

	for i := 0; i < SystemInstanceRadiusServerAttributeCount; i++ {
		var obj57_1_2 go_thunder.SystemInstanceRadiusServerAttribute
		prefix57_1_2 := prefix57_1 + fmt.Sprintf("attribute.%d.", i)
		obj57_1_2.SystemInstanceRadiusServerAttributeAttributeValue = d.Get(prefix57_1_2 + "attribute_value").(string)
		obj57_1_2.SystemInstanceRadiusServerAttributePrefixLength = d.Get(prefix57_1_2 + "prefix_length").(string)
		obj57_1_2.SystemInstanceRadiusServerAttributePrefixVendor = d.Get(prefix57_1_2 + "prefix_vendor").(int)
		obj57_1_2.SystemInstanceRadiusServerAttributePrefixNumber = d.Get(prefix57_1_2 + "prefix_number").(int)
		obj57_1_2.SystemInstanceRadiusServerAttributeName = d.Get(prefix57_1_2 + "name").(string)
		obj57_1_2.SystemInstanceRadiusServerAttributeValue = d.Get(prefix57_1_2 + "value").(string)
		obj57_1_2.SystemInstanceRadiusServerAttributeCustomVendor = d.Get(prefix57_1_2 + "custom_vendor").(int)
		obj57_1_2.SystemInstanceRadiusServerAttributeCustomNumber = d.Get(prefix57_1_2 + "custom_number").(int)
		obj57_1_2.SystemInstanceRadiusServerAttributeVendor = d.Get(prefix57_1_2 + "vendor").(int)
		obj57_1_2.SystemInstanceRadiusServerAttributeNumber = d.Get(prefix57_1_2 + "number").(int)
		obj57_1.SystemInstanceRadiusServerAttributeAttributeValue = append(obj57_1.SystemInstanceRadiusServerAttributeAttributeValue, obj57_1_2)
	}

	obj57_1.SystemInstanceRadiusServerDisableReply = d.Get(prefix57_1 + "disable_reply").(int)
	obj57_1.SystemInstanceRadiusServerAccountingStart = d.Get(prefix57_1 + "accounting_start").(string)
	obj57_1.SystemInstanceRadiusServerAccountingStop = d.Get(prefix57_1 + "accounting_stop").(string)
	obj57_1.SystemInstanceRadiusServerAccountingInterimUpdate = d.Get(prefix57_1 + "accounting_interim_update").(string)
	obj57_1.SystemInstanceRadiusServerAccountingOn = d.Get(prefix57_1 + "accounting_on").(string)
	obj57_1.SystemInstanceRadiusServerAttributeName = d.Get(prefix57_1 + "attribute_name").(string)
	obj57_1.SystemInstanceRadiusServerCustomAttributeName = d.Get(prefix57_1 + "custom_attribute_name").(string)

	SystemInstanceRadiusServerSamplingEnableCount := d.Get(prefix57_1 + "sampling_enable.#").(int)
	obj57_1.SystemInstanceRadiusServerSamplingEnableCounters1 = make([]go_thunder.SystemInstanceRadiusServerSamplingEnable, 0, SystemInstanceRadiusServerSamplingEnableCount)

	for i := 0; i < SystemInstanceRadiusServerSamplingEnableCount; i++ {
		var obj57_1_3 go_thunder.SystemInstanceRadiusServerSamplingEnable
		prefix57_1_3 := prefix57_1 + fmt.Sprintf("sampling_enable.%d.", i)
		obj57_1_3.SystemInstanceRadiusServerSamplingEnableCounters1 = d.Get(prefix57_1_3 + "counters1").(string)
		obj57_1.SystemInstanceRadiusServerSamplingEnableCounters1 = append(obj57_1.SystemInstanceRadiusServerSamplingEnableCounters1, obj57_1_3)
	}

	obj57.SystemInstanceRadiusServerListenPort = obj57_1

	c.SystemInstanceRadiusServer = obj57

	SystemInstanceGeolocListListCount := d.Get("geoloc_list_list.#").(int)
	c.SystemInstanceGeolocListListName = make([]go_thunder.SystemInstanceGeolocListList, 0, SystemInstanceGeolocListListCount)

	for i := 0; i < SystemInstanceGeolocListListCount; i++ {
		var obj58 go_thunder.SystemInstanceGeolocListList
		prefix58 := fmt.Sprintf("geoloc_list_list.%d.", i)
		obj58.SystemInstanceGeolocListListName = d.Get(prefix58 + "name").(string)
		obj58.SystemInstanceGeolocListListShared = d.Get(prefix58 + "shared").(int)

		SystemInstanceGeolocListListIncludeGeolocNameListCount := d.Get(prefix58 + "include_geoloc_name_list.#").(int)
		obj58.SystemInstanceGeolocListListIncludeGeolocNameListIncludeGeolocNameVal = make([]go_thunder.SystemInstanceGeolocListListIncludeGeolocNameList, 0, SystemInstanceGeolocListListIncludeGeolocNameListCount)

		for i := 0; i < SystemInstanceGeolocListListIncludeGeolocNameListCount; i++ {
			var obj58_1 go_thunder.SystemInstanceGeolocListListIncludeGeolocNameList
			prefix58_1 := prefix58 + fmt.Sprintf("include_geoloc_name_list.%d.", i)
			obj58_1.SystemInstanceGeolocListListIncludeGeolocNameListIncludeGeolocNameVal = d.Get(prefix58_1 + "include_geoloc_name_val").(string)
			obj58.SystemInstanceGeolocListListIncludeGeolocNameListIncludeGeolocNameVal = append(obj58.SystemInstanceGeolocListListIncludeGeolocNameListIncludeGeolocNameVal, obj58_1)
		}

		SystemInstanceGeolocListListExcludeGeolocNameListCount := d.Get(prefix58 + "exclude_geoloc_name_list.#").(int)
		obj58.SystemInstanceGeolocListListExcludeGeolocNameListExcludeGeolocNameVal = make([]go_thunder.SystemInstanceGeolocListListExcludeGeolocNameList, 0, SystemInstanceGeolocListListExcludeGeolocNameListCount)

		for i := 0; i < SystemInstanceGeolocListListExcludeGeolocNameListCount; i++ {
			var obj58_2 go_thunder.SystemInstanceGeolocListListExcludeGeolocNameList
			prefix58_2 := prefix58 + fmt.Sprintf("exclude_geoloc_name_list.%d.", i)
			obj58_2.SystemInstanceGeolocListListExcludeGeolocNameListExcludeGeolocNameVal = d.Get(prefix58_2 + "exclude_geoloc_name_val").(string)
			obj58.SystemInstanceGeolocListListExcludeGeolocNameListExcludeGeolocNameVal = append(obj58.SystemInstanceGeolocListListExcludeGeolocNameListExcludeGeolocNameVal, obj58_2)
		}

		obj58.SystemInstanceGeolocListListUserTag = d.Get(prefix58 + "user_tag").(string)

		SystemInstanceGeolocListListSamplingEnableCount := d.Get(prefix58 + "sampling_enable.#").(int)
		obj58.SystemInstanceGeolocListListSamplingEnableCounters1 = make([]go_thunder.SystemInstanceGeolocListListSamplingEnable, 0, SystemInstanceGeolocListListSamplingEnableCount)

		for i := 0; i < SystemInstanceGeolocListListSamplingEnableCount; i++ {
			var obj58_3 go_thunder.SystemInstanceGeolocListListSamplingEnable
			prefix58_3 := prefix58 + fmt.Sprintf("sampling_enable.%d.", i)
			obj58_3.SystemInstanceGeolocListListSamplingEnableCounters1 = d.Get(prefix58_3 + "counters1").(string)
			obj58.SystemInstanceGeolocListListSamplingEnableCounters1 = append(obj58.SystemInstanceGeolocListListSamplingEnableCounters1, obj58_3)
		}

		c.SystemInstanceGeolocListListName = append(c.SystemInstanceGeolocListListName, obj58)
	}

	var obj59 go_thunder.SystemInstanceGeolocNameHelper
	prefix59 := "geoloc_name_helper.0."

	SystemInstanceGeolocNameHelperSamplingEnableCount := d.Get(prefix59 + "sampling_enable.#").(int)
	obj59.SystemInstanceGeolocNameHelperSamplingEnableCounters1 = make([]go_thunder.SystemInstanceGeolocNameHelperSamplingEnable, 0, SystemInstanceGeolocNameHelperSamplingEnableCount)

	for i := 0; i < SystemInstanceGeolocNameHelperSamplingEnableCount; i++ {
		var obj59_1 go_thunder.SystemInstanceGeolocNameHelperSamplingEnable
		prefix59_1 := prefix59 + fmt.Sprintf("sampling_enable.%d.", i)
		obj59_1.SystemInstanceGeolocNameHelperSamplingEnableCounters1 = d.Get(prefix59_1 + "counters1").(string)
		obj59.SystemInstanceGeolocNameHelperSamplingEnableCounters1 = append(obj59.SystemInstanceGeolocNameHelperSamplingEnableCounters1, obj59_1)
	}

	c.SystemInstanceGeolocNameHelperUUID = obj59

	var obj61 go_thunder.SystemInstanceGeoloc
	prefix61 := "geoloc.0."

	SystemInstanceGeolocSamplingEnableCount := d.Get(prefix61 + "sampling_enable.#").(int)
	obj61.SystemInstanceGeolocSamplingEnableCounters1 = make([]go_thunder.SystemInstanceGeolocSamplingEnable, 0, SystemInstanceGeolocSamplingEnableCount)

	for i := 0; i < SystemInstanceGeolocSamplingEnableCount; i++ {
		var obj61_1 go_thunder.SystemInstanceGeolocSamplingEnable
		prefix61_1 := prefix61 + fmt.Sprintf("sampling_enable.%d.", i)
		obj61_1.SystemInstanceGeolocSamplingEnableCounters1 = d.Get(prefix61_1 + "counters1").(string)
		obj61.SystemInstanceGeolocSamplingEnableCounters1 = append(obj61.SystemInstanceGeolocSamplingEnableCounters1, obj61_1)
	}

	c.SystemInstanceGeolocUUID = obj61

	var obj62 go_thunder.SystemInstanceGeoLocation
	prefix62 := "geo_location.0."
	obj62.SystemInstanceGeoLocationGeoLocationIana = d.Get(prefix62 + "geo_location_iana").(int)
	obj62.SystemInstanceGeoLocationGeoLocationGeolite2City = d.Get(prefix62 + "geo_location_geolite2_city").(int)
	obj62.SystemInstanceGeoLocationGeolite2CityIncludeIpv6 = d.Get(prefix62 + "geolite2_city_include_ipv6").(int)
	obj62.SystemInstanceGeoLocationGeoLocationGeolite2Country = d.Get(prefix62 + "geo_location_geolite2_country").(int)
	obj62.SystemInstanceGeoLocationGeolite2CountryIncludeIpv6 = d.Get(prefix62 + "geolite2_country_include_ipv6").(int)

	SystemInstanceGeoLocationGeolocLoadFileListCount := d.Get(prefix62 + "geoloc_load_file_list.#").(int)
	obj62.SystemInstanceGeoLocationGeolocLoadFileListGeoLocationLoadFilename = make([]go_thunder.SystemInstanceGeoLocationGeolocLoadFileList, 0, SystemInstanceGeoLocationGeolocLoadFileListCount)

	for i := 0; i < SystemInstanceGeoLocationGeolocLoadFileListCount; i++ {
		var obj62_1 go_thunder.SystemInstanceGeoLocationGeolocLoadFileList
		prefix62_1 := prefix62 + fmt.Sprintf("geoloc_load_file_list.%d.", i)
		obj62_1.SystemInstanceGeoLocationGeolocLoadFileListGeoLocationLoadFilename = d.Get(prefix62_1 + "geo_location_load_filename").(string)
		obj62_1.SystemInstanceGeoLocationGeolocLoadFileListTemplateName = d.Get(prefix62_1 + "template_name").(string)
		obj62.SystemInstanceGeoLocationGeolocLoadFileListGeoLocationLoadFilename = append(obj62.SystemInstanceGeoLocationGeolocLoadFileListGeoLocationLoadFilename, obj62_1)
	}

	SystemInstanceGeoLocationEntryListCount := d.Get(prefix62 + "entry_list.#").(int)
	obj62.SystemInstanceGeoLocationEntryListGeoLocnObjName = make([]go_thunder.SystemInstanceGeoLocationEntryList, 0, SystemInstanceGeoLocationEntryListCount)

	for i := 0; i < SystemInstanceGeoLocationEntryListCount; i++ {
		var obj62_2 go_thunder.SystemInstanceGeoLocationEntryList
		prefix62_2 := prefix62 + fmt.Sprintf("entry_list.%d.", i)
		obj62_2.SystemInstanceGeoLocationEntryListGeoLocnObjName = d.Get(prefix62_2 + "geo_locn_obj_name").(string)

		SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesCount := d.Get(prefix62_2 + "geo_locn_multiple_addresses.#").(int)
		obj62_2.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIPAddress = make([]go_thunder.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddresses, 0, SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesCount)

		for i := 0; i < SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesCount; i++ {
			var obj62_2_1 go_thunder.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddresses
			prefix62_2_1 := prefix62_2 + fmt.Sprintf("geo_locn_multiple_addresses.%d.", i)
			obj62_2_1.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIPAddress = d.Get(prefix62_2_1 + "first_ip_address").(string)
			obj62_2_1.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesGeolIpv4Mask = d.Get(prefix62_2_1 + "geol_ipv4_mask").(string)
			obj62_2_1.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesIPAddr2 = d.Get(prefix62_2_1 + "ip_addr2").(string)
			obj62_2_1.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIpv6Address = d.Get(prefix62_2_1 + "first_ipv6_address").(string)
			obj62_2_1.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesGeolIpv6Mask = d.Get(prefix62_2_1 + "geol_ipv6_mask").(int)
			obj62_2_1.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesIpv6Addr2 = d.Get(prefix62_2_1 + "ipv6_addr2").(string)
			obj62_2.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIPAddress = append(obj62_2.SystemInstanceGeoLocationEntryListGeoLocnMultipleAddressesFirstIPAddress, obj62_2_1)
		}

		obj62_2.SystemInstanceGeoLocationEntryListUserTag = d.Get(prefix62_2 + "user_tag").(string)
		obj62.SystemInstanceGeoLocationEntryListGeoLocnObjName = append(obj62.SystemInstanceGeoLocationEntryListGeoLocnObjName, obj62_2)
	}

	c.SystemInstanceGeoLocationGeoLocationIana = obj62

	var obj63 go_thunder.SystemInstanceIPThreatList
	prefix63 := "ip_threat_list.0."

	SystemInstanceIPThreatListSamplingEnableCount := d.Get(prefix63 + "sampling_enable.#").(int)
	obj63.SystemInstanceIPThreatListSamplingEnableCounters1 = make([]go_thunder.SystemInstanceIPThreatListSamplingEnable, 0, SystemInstanceIPThreatListSamplingEnableCount)

	for i := 0; i < SystemInstanceIPThreatListSamplingEnableCount; i++ {
		var obj63_1 go_thunder.SystemInstanceIPThreatListSamplingEnable
		prefix63_1 := prefix63 + fmt.Sprintf("sampling_enable.%d.", i)
		obj63_1.SystemInstanceIPThreatListSamplingEnableCounters1 = d.Get(prefix63_1 + "counters1").(string)
		obj63.SystemInstanceIPThreatListSamplingEnableCounters1 = append(obj63.SystemInstanceIPThreatListSamplingEnableCounters1, obj63_1)
	}

	var obj63_2 go_thunder.SystemInstanceIPThreatListIpv4SourceList
	prefix63_2 := prefix63 + "ipv4_source_list.0."

	SystemInstanceIPThreatListIpv4SourceListClassListCfgCount := d.Get(prefix63_2 + "class_list_cfg.#").(int)
	obj63_2.SystemInstanceIPThreatListIpv4SourceListClassListCfgClassList = make([]go_thunder.SystemInstanceIPThreatListIpv4SourceListClassListCfg, 0, SystemInstanceIPThreatListIpv4SourceListClassListCfgCount)

	for i := 0; i < SystemInstanceIPThreatListIpv4SourceListClassListCfgCount; i++ {
		var obj63_2_1 go_thunder.SystemInstanceIPThreatListIpv4SourceListClassListCfg
		prefix63_2_1 := prefix63_2 + fmt.Sprintf("class_list_cfg.%d.", i)
		obj63_2_1.SystemInstanceIPThreatListIpv4SourceListClassListCfgClassList = d.Get(prefix63_2_1 + "class_list").(string)
		obj63_2_1.SystemInstanceIPThreatListIpv4SourceListClassListCfgIPThreatActionTmpl = d.Get(prefix63_2_1 + "ip_threat_action_tmpl").(int)
		obj63_2.SystemInstanceIPThreatListIpv4SourceListClassListCfgClassList = append(obj63_2.SystemInstanceIPThreatListIpv4SourceListClassListCfgClassList, obj63_2_1)
	}

	obj63.SystemInstanceIPThreatListIpv4SourceListClassListCfg = obj63_2

	var obj63_3 go_thunder.SystemInstanceIPThreatListIpv4DestList
	prefix63_3 := prefix63 + "ipv4_dest_list.0."

	SystemInstanceIPThreatListIpv4DestListClassListCfgCount := d.Get(prefix63_3 + "class_list_cfg.#").(int)
	obj63_3.SystemInstanceIPThreatListIpv4DestListClassListCfgClassList = make([]go_thunder.SystemInstanceIPThreatListIpv4DestListClassListCfg, 0, SystemInstanceIPThreatListIpv4DestListClassListCfgCount)

	for i := 0; i < SystemInstanceIPThreatListIpv4DestListClassListCfgCount; i++ {
		var obj63_3_1 go_thunder.SystemInstanceIPThreatListIpv4DestListClassListCfg
		prefix63_3_1 := prefix63_3 + fmt.Sprintf("class_list_cfg.%d.", i)
		obj63_3_1.SystemInstanceIPThreatListIpv4DestListClassListCfgClassList = d.Get(prefix63_3_1 + "class_list").(string)
		obj63_3_1.SystemInstanceIPThreatListIpv4DestListClassListCfgIPThreatActionTmpl = d.Get(prefix63_3_1 + "ip_threat_action_tmpl").(int)
		obj63_3.SystemInstanceIPThreatListIpv4DestListClassListCfgClassList = append(obj63_3.SystemInstanceIPThreatListIpv4DestListClassListCfgClassList, obj63_3_1)
	}

	obj63.SystemInstanceIPThreatListIpv4DestListClassListCfg = obj63_3

	var obj63_4 go_thunder.SystemInstanceIPThreatListIpv6SourceList
	prefix63_4 := prefix63 + "ipv6_source_list.0."

	SystemInstanceIPThreatListIpv6SourceListClassListCfgCount := d.Get(prefix63_4 + "class_list_cfg.#").(int)
	obj63_4.SystemInstanceIPThreatListIpv6SourceListClassListCfgClassList = make([]go_thunder.SystemInstanceIPThreatListIpv6SourceListClassListCfg, 0, SystemInstanceIPThreatListIpv6SourceListClassListCfgCount)

	for i := 0; i < SystemInstanceIPThreatListIpv6SourceListClassListCfgCount; i++ {
		var obj63_4_1 go_thunder.SystemInstanceIPThreatListIpv6SourceListClassListCfg
		prefix63_4_1 := prefix63_4 + fmt.Sprintf("class_list_cfg.%d.", i)
		obj63_4_1.SystemInstanceIPThreatListIpv6SourceListClassListCfgClassList = d.Get(prefix63_4_1 + "class_list").(string)
		obj63_4_1.SystemInstanceIPThreatListIpv6SourceListClassListCfgIPThreatActionTmpl = d.Get(prefix63_4_1 + "ip_threat_action_tmpl").(int)
		obj63_4.SystemInstanceIPThreatListIpv6SourceListClassListCfgClassList = append(obj63_4.SystemInstanceIPThreatListIpv6SourceListClassListCfgClassList, obj63_4_1)
	}

	obj63.SystemInstanceIPThreatListIpv6SourceListClassListCfg = obj63_4

	var obj63_5 go_thunder.SystemInstanceIPThreatListIpv6DestList
	prefix63_5 := prefix63 + "ipv6_dest_list.0."

	SystemInstanceIPThreatListIpv6DestListClassListCfgCount := d.Get(prefix63_5 + "class_list_cfg.#").(int)
	obj63_5.SystemInstanceIPThreatListIpv6DestListClassListCfgClassList = make([]go_thunder.SystemInstanceIPThreatListIpv6DestListClassListCfg, 0, SystemInstanceIPThreatListIpv6DestListClassListCfgCount)

	for i := 0; i < SystemInstanceIPThreatListIpv6DestListClassListCfgCount; i++ {
		var obj63_5_1 go_thunder.SystemInstanceIPThreatListIpv6DestListClassListCfg
		prefix63_5_1 := prefix63_5 + fmt.Sprintf("class_list_cfg.%d.", i)
		obj63_5_1.SystemInstanceIPThreatListIpv6DestListClassListCfgClassList = d.Get(prefix63_5_1 + "class_list").(string)
		obj63_5_1.SystemInstanceIPThreatListIpv6DestListClassListCfgIPThreatActionTmpl = d.Get(prefix63_5_1 + "ip_threat_action_tmpl").(int)
		obj63_5.SystemInstanceIPThreatListIpv6DestListClassListCfgClassList = append(obj63_5.SystemInstanceIPThreatListIpv6DestListClassListCfgClassList, obj63_5_1)
	}

	obj63.SystemInstanceIPThreatListIpv6DestListClassListCfg = obj63_5

	var obj63_6 go_thunder.SystemInstanceIPThreatListIpv4InternetHostList
	prefix63_6 := prefix63 + "ipv4_internet_host_list.0."

	SystemInstanceIPThreatListIpv4InternetHostListClassListCfgCount := d.Get(prefix63_6 + "class_list_cfg.#").(int)
	obj63_6.SystemInstanceIPThreatListIpv4InternetHostListClassListCfgClassList = make([]go_thunder.SystemInstanceIPThreatListIpv4InternetHostListClassListCfg, 0, SystemInstanceIPThreatListIpv4InternetHostListClassListCfgCount)

	for i := 0; i < SystemInstanceIPThreatListIpv4InternetHostListClassListCfgCount; i++ {
		var obj63_6_1 go_thunder.SystemInstanceIPThreatListIpv4InternetHostListClassListCfg
		prefix63_6_1 := prefix63_6 + fmt.Sprintf("class_list_cfg.%d.", i)
		obj63_6_1.SystemInstanceIPThreatListIpv4InternetHostListClassListCfgClassList = d.Get(prefix63_6_1 + "class_list").(string)
		obj63_6_1.SystemInstanceIPThreatListIpv4InternetHostListClassListCfgIPThreatActionTmpl = d.Get(prefix63_6_1 + "ip_threat_action_tmpl").(int)
		obj63_6.SystemInstanceIPThreatListIpv4InternetHostListClassListCfgClassList = append(obj63_6.SystemInstanceIPThreatListIpv4InternetHostListClassListCfgClassList, obj63_6_1)
	}

	obj63.SystemInstanceIPThreatListIpv4InternetHostListClassListCfg = obj63_6

	var obj63_7 go_thunder.SystemInstanceIPThreatListIpv6InternetHostList
	prefix63_7 := prefix63 + "ipv6_internet_host_list.0."

	SystemInstanceIPThreatListIpv6InternetHostListClassListCfgCount := d.Get(prefix63_7 + "class_list_cfg.#").(int)
	obj63_7.SystemInstanceIPThreatListIpv6InternetHostListClassListCfgClassList = make([]go_thunder.SystemInstanceIPThreatListIpv6InternetHostListClassListCfg, 0, SystemInstanceIPThreatListIpv6InternetHostListClassListCfgCount)

	for i := 0; i < SystemInstanceIPThreatListIpv6InternetHostListClassListCfgCount; i++ {
		var obj63_7_1 go_thunder.SystemInstanceIPThreatListIpv6InternetHostListClassListCfg
		prefix63_7_1 := prefix63_7 + fmt.Sprintf("class_list_cfg.%d.", i)
		obj63_7_1.SystemInstanceIPThreatListIpv6InternetHostListClassListCfgClassList = d.Get(prefix63_7_1 + "class_list").(string)
		obj63_7_1.SystemInstanceIPThreatListIpv6InternetHostListClassListCfgIPThreatActionTmpl = d.Get(prefix63_7_1 + "ip_threat_action_tmpl").(int)
		obj63_7.SystemInstanceIPThreatListIpv6InternetHostListClassListCfgClassList = append(obj63_7.SystemInstanceIPThreatListIpv6InternetHostListClassListCfgClassList, obj63_7_1)
	}

	obj63.SystemInstanceIPThreatListIpv6InternetHostListClassListCfg = obj63_7

	c.SystemInstanceIPThreatListUUID = obj63

	var obj64 go_thunder.SystemInstanceFpgaDrop
	prefix64 := "fpga_drop.0."

	SystemInstanceFpgaDropSamplingEnableCount := d.Get(prefix64 + "sampling_enable.#").(int)
	obj64.SystemInstanceFpgaDropSamplingEnableCounters1 = make([]go_thunder.SystemInstanceFpgaDropSamplingEnable, 0, SystemInstanceFpgaDropSamplingEnableCount)

	for i := 0; i < SystemInstanceFpgaDropSamplingEnableCount; i++ {
		var obj64_1 go_thunder.SystemInstanceFpgaDropSamplingEnable
		prefix64_1 := prefix64 + fmt.Sprintf("sampling_enable.%d.", i)
		obj64_1.SystemInstanceFpgaDropSamplingEnableCounters1 = d.Get(prefix64_1 + "counters1").(string)
		obj64.SystemInstanceFpgaDropSamplingEnableCounters1 = append(obj64.SystemInstanceFpgaDropSamplingEnableCounters1, obj64_1)
	}

	c.SystemInstanceFpgaDropUUID = obj64

	var obj65 go_thunder.SystemInstanceDpdkStats
	prefix65 := "dpdk_stats.0."

	SystemInstanceDpdkStatsSamplingEnableCount := d.Get(prefix65 + "sampling_enable.#").(int)
	obj65.SystemInstanceDpdkStatsSamplingEnableCounters1 = make([]go_thunder.SystemInstanceDpdkStatsSamplingEnable, 0, SystemInstanceDpdkStatsSamplingEnableCount)

	for i := 0; i < SystemInstanceDpdkStatsSamplingEnableCount; i++ {
		var obj65_1 go_thunder.SystemInstanceDpdkStatsSamplingEnable
		prefix65_1 := prefix65 + fmt.Sprintf("sampling_enable.%d.", i)
		obj65_1.SystemInstanceDpdkStatsSamplingEnableCounters1 = d.Get(prefix65_1 + "counters1").(string)
		obj65.SystemInstanceDpdkStatsSamplingEnableCounters1 = append(obj65.SystemInstanceDpdkStatsSamplingEnableCounters1, obj65_1)
	}

	c.SystemInstanceDpdkStatsUUID = obj65

	var obj66 go_thunder.SystemInstanceFpgaCoreCrc
	prefix66 := "fpga_core_crc.0."
	obj66.SystemInstanceFpgaCoreCrcMonitorDisable = d.Get(prefix66 + "monitor_disable").(int)
	obj66.SystemInstanceFpgaCoreCrcRebootEnable = d.Get(prefix66 + "reboot_enable").(int)

	c.SystemInstanceFpgaCoreCrcMonitorDisable = obj66

	var obj67 go_thunder.SystemInstanceMfaManagement
	prefix67 := "mfa_management.0."
	obj67.SystemInstanceMfaManagementEnable = d.Get(prefix67 + "enable").(int)

	c.SystemInstanceMfaManagementEnable = obj67

	var obj68 go_thunder.SystemInstanceMfaValidationType
	prefix68 := "mfa_validation_type.0."
	obj68.SystemInstanceMfaValidationTypeCaCert = d.Get(prefix68 + "ca_cert").(string)

	c.SystemInstanceMfaValidationTypeCaCert = obj68

	var obj69 go_thunder.SystemInstanceMfaCertStore
	prefix69 := "mfa_cert_store.0."
	obj69.SystemInstanceMfaCertStoreCertHost = d.Get(prefix69 + "cert_host").(string)
	obj69.SystemInstanceMfaCertStoreProtocol = d.Get(prefix69 + "protocol").(string)
	obj69.SystemInstanceMfaCertStoreCertStorePath = d.Get(prefix69 + "cert_store_path").(string)
	obj69.SystemInstanceMfaCertStoreUsername = d.Get(prefix69 + "username").(string)
	obj69.SystemInstanceMfaCertStorePasswdString = d.Get(prefix69 + "passwd_string").(string)

	c.SystemInstanceMfaCertStoreCertHost = obj69

	var obj70 go_thunder.SystemInstanceMfaAuth
	prefix70 := "mfa_auth.0."
	obj70.SystemInstanceMfaAuthUsername = d.Get(prefix70 + "username").(string)
	obj70.SystemInstanceMfaAuthSecondFactor = d.Get(prefix70 + "second_factor").(string)

	c.SystemInstanceMfaAuthUsername = obj70

	var obj71 go_thunder.SystemInstanceQInQ
	prefix71 := "q_in_q.0."
	obj71.SystemInstanceQInQInnerTpid = d.Get(prefix71 + "inner_tpid").(string)
	obj71.SystemInstanceQInQOuterTpid = d.Get(prefix71 + "outer_tpid").(string)
	obj71.SystemInstanceQInQEnableAllPorts = d.Get(prefix71 + "enable_all_ports").(int)

	c.SystemInstanceQInQInnerTpid = obj71

	var obj72 go_thunder.SystemInstancePortCount
	prefix72 := "port_count.0."
	obj72.SystemInstancePortCountPortCountKernel = d.Get(prefix72 + "port_count_kernel").(int)
	obj72.SystemInstancePortCountPortCountHm = d.Get(prefix72 + "port_count_hm").(int)
	obj72.SystemInstancePortCountPortCountLogging = d.Get(prefix72 + "port_count_logging").(int)
	obj72.SystemInstancePortCountPortCountAlg = d.Get(prefix72 + "port_count_alg").(int)

	c.SystemInstancePortCountPortCountKernel = obj72

	SystemInstanceHealthCheckListCount := d.Get("health_check_list.#").(int)
	c.SystemInstanceHealthCheckListL2HmHcName = make([]go_thunder.SystemInstanceHealthCheckList, 0, SystemInstanceHealthCheckListCount)

	for i := 0; i < SystemInstanceHealthCheckListCount; i++ {
		var obj73 go_thunder.SystemInstanceHealthCheckList
		prefix73 := fmt.Sprintf("health_check_list.%d.", i)
		obj73.SystemInstanceHealthCheckListL2HmHcName = d.Get(prefix73 + "l2hm_hc_name").(string)
		obj73.SystemInstanceHealthCheckListMethodL2Bfd = d.Get(prefix73 + "method_l2bfd").(int)
		obj73.SystemInstanceHealthCheckListL2BfdTxInterval = d.Get(prefix73 + "l2bfd_tx_interval").(int)
		obj73.SystemInstanceHealthCheckListL2BfdRxInterval = d.Get(prefix73 + "l2bfd_rx_interval").(int)
		obj73.SystemInstanceHealthCheckListL2BfdMultiplier = d.Get(prefix73 + "l2bfd_multiplier").(int)
		obj73.SystemInstanceHealthCheckListUserTag = d.Get(prefix73 + "user_tag").(string)
		c.SystemInstanceHealthCheckListL2HmHcName = append(c.SystemInstanceHealthCheckListL2HmHcName, obj73)
	}

	SystemInstancePathListCount := d.Get("path_list.#").(int)
	c.SystemInstancePathListL2HmPathName = make([]go_thunder.SystemInstancePathList, 0, SystemInstancePathListCount)

	for i := 0; i < SystemInstancePathListCount; i++ {
		var obj74 go_thunder.SystemInstancePathList
		prefix74 := fmt.Sprintf("path_list.%d.", i)
		obj74.SystemInstancePathListL2HmPathName = d.Get(prefix74 + "l2hm_path_name").(string)
		obj74.SystemInstancePathListL2HmVlan = d.Get(prefix74 + "l2hm_vlan").(int)
		obj74.SystemInstancePathListL2HmSetupTestAPI = d.Get(prefix74 + "l2hm_setup_test_api").(int)
		obj74.SystemInstancePathListIfpairEthStart = d.Get(prefix74 + "ifpair_eth_start").(int)
		obj74.SystemInstancePathListIfpairEthEnd = d.Get(prefix74 + "ifpair_eth_end").(int)
		obj74.SystemInstancePathListIfpairTrunkStart = d.Get(prefix74 + "ifpair_trunk_start").(int)
		obj74.SystemInstancePathListIfpairTrunkEnd = d.Get(prefix74 + "ifpair_trunk_end").(int)
		obj74.SystemInstancePathListL2HmAttach = d.Get(prefix74 + "l2hm_attach").(string)
		obj74.SystemInstancePathListUserTag = d.Get(prefix74 + "user_tag").(string)
		c.SystemInstancePathListL2HmPathName = append(c.SystemInstancePathListL2HmPathName, obj74)
	}

	var obj75 go_thunder.SystemInstanceSyslogTimeMsec
	prefix75 := "syslog_time_msec.0."
	obj75.SystemInstanceSyslogTimeMsecEnableFlag = d.Get(prefix75 + "enable_flag").(int)

	c.SystemInstanceSyslogTimeMsecEnableFlag = obj75

	var obj76 go_thunder.SystemInstanceIpmiService
	prefix76 := "ipmi_service.0."
	obj76.SystemInstanceIpmiServiceDisable = d.Get(prefix76 + "disable").(int)

	c.SystemInstanceIpmiServiceDisable = obj76

	var obj77 go_thunder.SystemInstanceAppPerformance
	prefix77 := "app_performance.0."

	SystemInstanceAppPerformanceSamplingEnableCount := d.Get(prefix77 + "sampling_enable.#").(int)
	obj77.SystemInstanceAppPerformanceSamplingEnableCounters1 = make([]go_thunder.SystemInstanceAppPerformanceSamplingEnable, 0, SystemInstanceAppPerformanceSamplingEnableCount)

	for i := 0; i < SystemInstanceAppPerformanceSamplingEnableCount; i++ {
		var obj77_1 go_thunder.SystemInstanceAppPerformanceSamplingEnable
		prefix77_1 := prefix77 + fmt.Sprintf("sampling_enable.%d.", i)
		obj77_1.SystemInstanceAppPerformanceSamplingEnableCounters1 = d.Get(prefix77_1 + "counters1").(string)
		obj77.SystemInstanceAppPerformanceSamplingEnableCounters1 = append(obj77.SystemInstanceAppPerformanceSamplingEnableCounters1, obj77_1)
	}

	c.SystemInstanceAppPerformanceUUID = obj77

	var obj78 go_thunder.SystemInstanceSslReqQ
	prefix78 := "ssl_req_q.0."

	SystemInstanceSslReqQSamplingEnableCount := d.Get(prefix78 + "sampling_enable.#").(int)
	obj78.SystemInstanceSslReqQSamplingEnableCounters1 = make([]go_thunder.SystemInstanceSslReqQSamplingEnable, 0, SystemInstanceSslReqQSamplingEnableCount)

	for i := 0; i < SystemInstanceSslReqQSamplingEnableCount; i++ {
		var obj78_1 go_thunder.SystemInstanceSslReqQSamplingEnable
		prefix78_1 := prefix78 + fmt.Sprintf("sampling_enable.%d.", i)
		obj78_1.SystemInstanceSslReqQSamplingEnableCounters1 = d.Get(prefix78_1 + "counters1").(string)
		obj78.SystemInstanceSslReqQSamplingEnableCounters1 = append(obj78.SystemInstanceSslReqQSamplingEnableCounters1, obj78_1)
	}

	c.SystemInstanceSslReqQUUID = obj78

	var obj79 go_thunder.SystemInstanceTCP
	prefix79 := "tcp.0."

	SystemInstanceTCPSamplingEnableCount := d.Get(prefix79 + "sampling_enable.#").(int)
	obj79.SystemInstanceTCPSamplingEnableCounters1 = make([]go_thunder.SystemInstanceTCPSamplingEnable, 0, SystemInstanceTCPSamplingEnableCount)

	for i := 0; i < SystemInstanceTCPSamplingEnableCount; i++ {
		var obj79_1 go_thunder.SystemInstanceTCPSamplingEnable
		prefix79_1 := prefix79 + fmt.Sprintf("sampling_enable.%d.", i)
		obj79_1.SystemInstanceTCPSamplingEnableCounters1 = d.Get(prefix79_1 + "counters1").(string)
		obj79.SystemInstanceTCPSamplingEnableCounters1 = append(obj79.SystemInstanceTCPSamplingEnableCounters1, obj79_1)
	}

	var obj79_2 go_thunder.SystemInstanceTCPRateLimitResetUnknownConn
	prefix79_2 := prefix79 + "rate_limit_reset_unknown_conn.0."
	obj79_2.SystemInstanceTCPRateLimitResetUnknownConnPktRateForResetUnknownConn = d.Get(prefix79_2 + "pkt_rate_for_reset_unknown_conn").(int)
	obj79_2.SystemInstanceTCPRateLimitResetUnknownConnLogForResetUnknownConn = d.Get(prefix79_2 + "log_for_reset_unknown_conn").(int)

	obj79.SystemInstanceTCPRateLimitResetUnknownConnPktRateForResetUnknownConn = obj79_2

	c.SystemInstanceTCPUUID = obj79

	var obj80 go_thunder.SystemInstanceIcmp
	prefix80 := "icmp.0."

	SystemInstanceIcmpSamplingEnableCount := d.Get(prefix80 + "sampling_enable.#").(int)
	obj80.SystemInstanceIcmpSamplingEnableCounters1 = make([]go_thunder.SystemInstanceIcmpSamplingEnable, 0, SystemInstanceIcmpSamplingEnableCount)

	for i := 0; i < SystemInstanceIcmpSamplingEnableCount; i++ {
		var obj80_1 go_thunder.SystemInstanceIcmpSamplingEnable
		prefix80_1 := prefix80 + fmt.Sprintf("sampling_enable.%d.", i)
		obj80_1.SystemInstanceIcmpSamplingEnableCounters1 = d.Get(prefix80_1 + "counters1").(string)
		obj80.SystemInstanceIcmpSamplingEnableCounters1 = append(obj80.SystemInstanceIcmpSamplingEnableCounters1, obj80_1)
	}

	c.SystemInstanceIcmpUUID = obj80

	var obj81 go_thunder.SystemInstanceIcmp6
	prefix81 := "icmp6.0."

	SystemInstanceIcmp6SamplingEnableCount := d.Get(prefix81 + "sampling_enable.#").(int)
	obj81.SystemInstanceIcmp6SamplingEnableCounters1 = make([]go_thunder.SystemInstanceIcmp6SamplingEnable, 0, SystemInstanceIcmp6SamplingEnableCount)

	for i := 0; i < SystemInstanceIcmp6SamplingEnableCount; i++ {
		var obj81_1 go_thunder.SystemInstanceIcmp6SamplingEnable
		prefix81_1 := prefix81 + fmt.Sprintf("sampling_enable.%d.", i)
		obj81_1.SystemInstanceIcmp6SamplingEnableCounters1 = d.Get(prefix81_1 + "counters1").(string)
		obj81.SystemInstanceIcmp6SamplingEnableCounters1 = append(obj81.SystemInstanceIcmp6SamplingEnableCounters1, obj81_1)
	}

	c.SystemInstanceIcmp6UUID = obj81

	var obj82 go_thunder.SystemInstanceIPStats
	prefix82 := "ip_stats.0."

	SystemInstanceIPStatsSamplingEnableCount := d.Get(prefix82 + "sampling_enable.#").(int)
	obj82.SystemInstanceIPStatsSamplingEnableCounters1 = make([]go_thunder.SystemInstanceIPStatsSamplingEnable, 0, SystemInstanceIPStatsSamplingEnableCount)

	for i := 0; i < SystemInstanceIPStatsSamplingEnableCount; i++ {
		var obj82_1 go_thunder.SystemInstanceIPStatsSamplingEnable
		prefix82_1 := prefix82 + fmt.Sprintf("sampling_enable.%d.", i)
		obj82_1.SystemInstanceIPStatsSamplingEnableCounters1 = d.Get(prefix82_1 + "counters1").(string)
		obj82.SystemInstanceIPStatsSamplingEnableCounters1 = append(obj82.SystemInstanceIPStatsSamplingEnableCounters1, obj82_1)
	}

	c.SystemInstanceIPStatsUUID = obj82

	var obj83 go_thunder.SystemInstanceIP6Stats
	prefix83 := "ip6_stats.0."

	SystemInstanceIP6StatsSamplingEnableCount := d.Get(prefix83 + "sampling_enable.#").(int)
	obj83.SystemInstanceIP6StatsSamplingEnableCounters1 = make([]go_thunder.SystemInstanceIP6StatsSamplingEnable, 0, SystemInstanceIP6StatsSamplingEnableCount)

	for i := 0; i < SystemInstanceIP6StatsSamplingEnableCount; i++ {
		var obj83_1 go_thunder.SystemInstanceIP6StatsSamplingEnable
		prefix83_1 := prefix83 + fmt.Sprintf("sampling_enable.%d.", i)
		obj83_1.SystemInstanceIP6StatsSamplingEnableCounters1 = d.Get(prefix83_1 + "counters1").(string)
		obj83.SystemInstanceIP6StatsSamplingEnableCounters1 = append(obj83.SystemInstanceIP6StatsSamplingEnableCounters1, obj83_1)
	}

	c.SystemInstanceIP6StatsUUID = obj83

	var obj84 go_thunder.SystemInstanceBfd
	prefix84 := "bfd.0."

	SystemInstanceBfdSamplingEnableCount := d.Get(prefix84 + "sampling_enable.#").(int)
	obj84.SystemInstanceBfdSamplingEnableCounters1 = make([]go_thunder.SystemInstanceBfdSamplingEnable, 0, SystemInstanceBfdSamplingEnableCount)

	for i := 0; i < SystemInstanceBfdSamplingEnableCount; i++ {
		var obj84_1 go_thunder.SystemInstanceBfdSamplingEnable
		prefix84_1 := prefix84 + fmt.Sprintf("sampling_enable.%d.", i)
		obj84_1.SystemInstanceBfdSamplingEnableCounters1 = d.Get(prefix84_1 + "counters1").(string)
		obj84.SystemInstanceBfdSamplingEnableCounters1 = append(obj84.SystemInstanceBfdSamplingEnableCounters1, obj84_1)
	}

	c.SystemInstanceBfdUUID = obj84

	var obj85 go_thunder.SystemInstanceIcmpRate
	prefix85 := "icmp_rate.0."

	SystemInstanceIcmpRateSamplingEnableCount := d.Get(prefix85 + "sampling_enable.#").(int)
	obj85.SystemInstanceIcmpRateSamplingEnableCounters1 = make([]go_thunder.SystemInstanceIcmpRateSamplingEnable, 0, SystemInstanceIcmpRateSamplingEnableCount)

	for i := 0; i < SystemInstanceIcmpRateSamplingEnableCount; i++ {
		var obj85_1 go_thunder.SystemInstanceIcmpRateSamplingEnable
		prefix85_1 := prefix85 + fmt.Sprintf("sampling_enable.%d.", i)
		obj85_1.SystemInstanceIcmpRateSamplingEnableCounters1 = d.Get(prefix85_1 + "counters1").(string)
		obj85.SystemInstanceIcmpRateSamplingEnableCounters1 = append(obj85.SystemInstanceIcmpRateSamplingEnableCounters1, obj85_1)
	}

	c.SystemInstanceIcmpRateUUID = obj85

	var obj86 go_thunder.SystemInstanceJobOffload
	prefix86 := "job_offload.0."

	SystemInstanceJobOffloadSamplingEnableCount := d.Get(prefix86 + "sampling_enable.#").(int)
	obj86.SystemInstanceJobOffloadSamplingEnableCounters1 = make([]go_thunder.SystemInstanceJobOffloadSamplingEnable, 0, SystemInstanceJobOffloadSamplingEnableCount)

	for i := 0; i < SystemInstanceJobOffloadSamplingEnableCount; i++ {
		var obj86_1 go_thunder.SystemInstanceJobOffloadSamplingEnable
		prefix86_1 := prefix86 + fmt.Sprintf("sampling_enable.%d.", i)
		obj86_1.SystemInstanceJobOffloadSamplingEnableCounters1 = d.Get(prefix86_1 + "counters1").(string)
		obj86.SystemInstanceJobOffloadSamplingEnableCounters1 = append(obj86.SystemInstanceJobOffloadSamplingEnableCounters1, obj86_1)
	}

	c.SystemInstanceJobOffloadUUID = obj86

	var obj87 go_thunder.SystemInstanceDNS
	prefix87 := "dns.0."

	SystemInstanceDNSSamplingEnableCount := d.Get(prefix87 + "sampling_enable.#").(int)
	obj87.SystemInstanceDNSSamplingEnableCounters1 = make([]go_thunder.SystemInstanceDNSSamplingEnable, 0, SystemInstanceDNSSamplingEnableCount)

	for i := 0; i < SystemInstanceDNSSamplingEnableCount; i++ {
		var obj87_1 go_thunder.SystemInstanceDNSSamplingEnable
		prefix87_1 := prefix87 + fmt.Sprintf("sampling_enable.%d.", i)
		obj87_1.SystemInstanceDNSSamplingEnableCounters1 = d.Get(prefix87_1 + "counters1").(string)
		obj87.SystemInstanceDNSSamplingEnableCounters1 = append(obj87.SystemInstanceDNSSamplingEnableCounters1, obj87_1)
	}

	c.SystemInstanceDNSUUID = obj87

	var obj88 go_thunder.SystemInstanceDNSCache
	prefix88 := "dns_cache.0."

	SystemInstanceDNSCacheSamplingEnableCount := d.Get(prefix88 + "sampling_enable.#").(int)
	obj88.SystemInstanceDNSCacheSamplingEnableCounters1 = make([]go_thunder.SystemInstanceDNSCacheSamplingEnable, 0, SystemInstanceDNSCacheSamplingEnableCount)

	for i := 0; i < SystemInstanceDNSCacheSamplingEnableCount; i++ {
		var obj88_1 go_thunder.SystemInstanceDNSCacheSamplingEnable
		prefix88_1 := prefix88 + fmt.Sprintf("sampling_enable.%d.", i)
		obj88_1.SystemInstanceDNSCacheSamplingEnableCounters1 = d.Get(prefix88_1 + "counters1").(string)
		obj88.SystemInstanceDNSCacheSamplingEnableCounters1 = append(obj88.SystemInstanceDNSCacheSamplingEnableCounters1, obj88_1)
	}

	c.SystemInstanceDNSCacheUUID = obj88

	var obj89 go_thunder.SystemInstanceSession
	prefix89 := "session.0."

	SystemInstanceSessionSamplingEnableCount := d.Get(prefix89 + "sampling_enable.#").(int)
	obj89.SystemInstanceSessionSamplingEnableCounters1 = make([]go_thunder.SystemInstanceSessionSamplingEnable, 0, SystemInstanceSessionSamplingEnableCount)

	for i := 0; i < SystemInstanceSessionSamplingEnableCount; i++ {
		var obj89_1 go_thunder.SystemInstanceSessionSamplingEnable
		prefix89_1 := prefix89 + fmt.Sprintf("sampling_enable.%d.", i)
		obj89_1.SystemInstanceSessionSamplingEnableCounters1 = d.Get(prefix89_1 + "counters1").(string)
		obj89.SystemInstanceSessionSamplingEnableCounters1 = append(obj89.SystemInstanceSessionSamplingEnableCounters1, obj89_1)
	}

	c.SystemInstanceSessionUUID = obj89

	var obj90 go_thunder.SystemInstanceNdiscRa
	prefix90 := "ndisc_ra.0."

	SystemInstanceNdiscRaSamplingEnableCount := d.Get(prefix90 + "sampling_enable.#").(int)
	obj90.SystemInstanceNdiscRaSamplingEnableCounters1 = make([]go_thunder.SystemInstanceNdiscRaSamplingEnable, 0, SystemInstanceNdiscRaSamplingEnableCount)

	for i := 0; i < SystemInstanceNdiscRaSamplingEnableCount; i++ {
		var obj90_1 go_thunder.SystemInstanceNdiscRaSamplingEnable
		prefix90_1 := prefix90 + fmt.Sprintf("sampling_enable.%d.", i)
		obj90_1.SystemInstanceNdiscRaSamplingEnableCounters1 = d.Get(prefix90_1 + "counters1").(string)
		obj90.SystemInstanceNdiscRaSamplingEnableCounters1 = append(obj90.SystemInstanceNdiscRaSamplingEnableCounters1, obj90_1)
	}

	c.SystemInstanceNdiscRaUUID = obj90

	var obj91 go_thunder.SystemInstanceTCPStats
	prefix91 := "tcp_stats.0."

	SystemInstanceTCPStatsSamplingEnableCount := d.Get(prefix91 + "sampling_enable.#").(int)
	obj91.SystemInstanceTCPStatsSamplingEnableCounters1 = make([]go_thunder.SystemInstanceTCPStatsSamplingEnable, 0, SystemInstanceTCPStatsSamplingEnableCount)

	for i := 0; i < SystemInstanceTCPStatsSamplingEnableCount; i++ {
		var obj91_1 go_thunder.SystemInstanceTCPStatsSamplingEnable
		prefix91_1 := prefix91 + fmt.Sprintf("sampling_enable.%d.", i)
		obj91_1.SystemInstanceTCPStatsSamplingEnableCounters1 = d.Get(prefix91_1 + "counters1").(string)
		obj91.SystemInstanceTCPStatsSamplingEnableCounters1 = append(obj91.SystemInstanceTCPStatsSamplingEnableCounters1, obj91_1)
	}

	c.SystemInstanceTCPStatsUUID = obj91

	vc.SystemInstanceAnomalyLog = c
	return vc
}
