package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutCluster() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster`: Configure scaleout cluster\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterCreate,
		UpdateContext: resourceScaleoutClusterUpdate,
		ReadContext:   resourceScaleoutClusterRead,
		DeleteContext: resourceScaleoutClusterDelete,

		Schema: map[string]*schema.Schema{
			"cluster_devices": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"minimum_nodes": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"minimum_nodes_num": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the minimum number of the node required to start service",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"cluster_discovery_timeout": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"device_id_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"action": {
										Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
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
			"cluster_id": {
				Type: schema.TypeInt, Required: true, Description: "Scaleout cluster-id",
			},
			"db_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ticktime": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"initlimit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"synclimit": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"minsessiontimeout": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "",
						},
						"maxsessiontimeout": {
							Type: schema.TypeInt, Optional: true, Default: 30000, Description: "",
						},
						"client_recv_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 13000, Description: "",
						},
						"clientport": {
							Type: schema.TypeInt, Optional: true, Description: "client session port",
						},
						"loopback_intf_support": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "support loopback interface for scaleout database (enabled by default)",
						},
						"broken_detect_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 12000, Description: "database connection broken detection timeout (mseconds) (12000 mseconds for default)",
						},
						"more_election_packet": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "send more election packet in election period (enabled by default)",
						},
						"elect_conn_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 1200, Description: "election connection timeout (mseconds) (1200 for default)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"device_groups": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"device_group_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_group": {
										Type: schema.TypeInt, Required: true, Description: "scaleout device group",
									},
									"device_id_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"device_id_start": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
												"device_id_end": {
													Type: schema.TypeInt, Optional: true, Description: "",
												},
											},
										},
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
			"local_device": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"id1": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
						},
						"start_delay": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cluster_mode": {
							Type: schema.TypeString, Optional: true, Default: "layer-2", Description: "'layer-2': Nodes in cluster are layer 2 connected (default mode); 'layer-3': Nodes in cluster are l3 connected;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"l2_redirect": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redirect_eth": {
										Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Port Value)",
									},
									"ethernet_vlan": {
										Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
									},
									"redirect_trunk": {
										Type: schema.TypeInt, Optional: true, Description: "L2 Trunk group",
									},
									"trunk_vlan": {
										Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"traffic_redirection": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"follow_shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow shared partition for redirection",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"interfaces": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"eth_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ethernet": {
																Type: schema.TypeInt, Optional: true, Description: "Ethernet Interface (Ethernet interface number)",
															},
														},
													},
												},
												"trunk_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trunk": {
																Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk interface number)",
															},
														},
													},
												},
												"ve_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ve": {
																Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet Interface (Virtual ethernet interface number)",
															},
														},
													},
												},
												"loopback_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"loopback": {
																Type: schema.TypeInt, Optional: true, Description: "Loopback Interface (Loopback interface number)",
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
									"reachability_options": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"skip_default_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not choose default route for redirection",
												},
												"uuid": {
													Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
												},
											},
										},
									},
									"encap": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type: schema.TypeString, Optional: true, Default: "vxlan", Description: "'vxlan': Use vxlan for encapsulation;",
												},
												"use_v4_vxlan": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always use IPv4 VxLAN for redirection",
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
						"session_sync": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"follow_shared": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow shared partition for session sync",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"interfaces": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"eth_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ethernet": {
																Type: schema.TypeInt, Optional: true, Description: "Ethernet Interface (Ethernet interface number)",
															},
														},
													},
												},
												"trunk_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trunk": {
																Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk interface number)",
															},
														},
													},
												},
												"ve_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ve": {
																Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet Interface (Virtual ethernet interface number)",
															},
														},
													},
												},
												"loopback_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"loopback": {
																Type: schema.TypeInt, Optional: true, Description: "Loopback Interface(Not applicable in 'layer-2' mode) (Loopback interface number)",
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
									"reachability_options": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"skip_default_route": {
													Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not choose default route for redirection(Not applicable in 'layer-2' mode)",
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
						"exclude_interfaces": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eth_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ethernet": {
													Type: schema.TypeInt, Optional: true, Description: "Ethernet Interface (Ethernet interface number)",
												},
											},
										},
									},
									"trunk_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trunk": {
													Type: schema.TypeInt, Optional: true, Description: "Trunk Interface (Trunk interface number)",
												},
											},
										},
									},
									"ve_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ve": {
													Type: schema.TypeInt, Optional: true, Description: "Virtual ethernet Interface (Virtual ethernet interface number)",
												},
											},
										},
									},
									"loopback_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"loopback": {
													Type: schema.TypeInt, Optional: true, Description: "Loopback Interface (Loopback interface number)",
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
						"tracking_template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"template": {
													Type: schema.TypeString, Required: true, Description: "bind tracking template name",
												},
												"ip_version": {
													Type: schema.TypeString, Required: true, Description: "'ipv4': take action for IPv4 traffic-only; 'ipv6': take action for IPv6 traffic-only;",
												},
												"threshold_cfg": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"threshold": {
																Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
															},
															"action": {
																Type: schema.TypeString, Optional: true, Description: "'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;",
															},
														},
													},
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
									"multi_template_list": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"multi_template": {
													Type: schema.TypeString, Required: true, Description: "bind multi tracking template name",
												},
												"template": {
													Type: schema.TypeList, Optional: true, Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"template_name": {
																Type: schema.TypeString, Optional: true, Description: "bind tracking template name",
															},
															"partition_name": {
																Type: schema.TypeString, Optional: true, Description: "Partition name",
															},
														},
													},
												},
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;",
												},
												"ip_version": {
													Type: schema.TypeString, Required: true, Description: "'ipv4': take action for IPv4 traffic-only; 'ipv6': take action for IPv6 traffic-only;",
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
					},
				},
			},
			"service_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_user_group_count": {
							Type: schema.TypeInt, Optional: true, Description: "Number of default traffic buckets",
						},
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"template_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type: schema.TypeString, Required: true, Description: "Scaleout template Name",
									},
									"user_group_count": {
										Type: schema.TypeInt, Optional: true, Description: "Number of traffic buckets",
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
			"slog_level": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the level of slog for Scaleout",
			},
			"tracking_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template": {
										Type: schema.TypeString, Required: true, Description: "bind tracking template name",
									},
									"threshold_cfg": {
										Type: schema.TypeList, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"threshold": {
													Type: schema.TypeInt, Optional: true, Description: "action triggering threshold",
												},
												"action": {
													Type: schema.TypeString, Optional: true, Description: "'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;",
												},
											},
										},
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceScaleoutClusterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutCluster(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutCluster(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutCluster(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutCluster(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectScaleoutClusterClusterDevices1449(d []interface{}) edpt.ScaleoutClusterClusterDevices1449 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterClusterDevices1449
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
		ret.MinimumNodes = getObjectScaleoutClusterClusterDevicesMinimumNodes1450(in["minimum_nodes"].([]interface{}))
		ret.ClusterDiscoveryTimeout = getObjectScaleoutClusterClusterDevicesClusterDiscoveryTimeout1451(in["cluster_discovery_timeout"].([]interface{}))
		ret.DeviceIdList = getSliceScaleoutClusterClusterDevicesDeviceIdList1452(in["device_id_list"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterClusterDevicesMinimumNodes1450(d []interface{}) edpt.ScaleoutClusterClusterDevicesMinimumNodes1450 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterClusterDevicesMinimumNodes1450
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinimumNodesNum = in["minimum_nodes_num"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterClusterDevicesClusterDiscoveryTimeout1451(d []interface{}) edpt.ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1451 {

	var ret edpt.ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1451
	return ret
}

func getSliceScaleoutClusterClusterDevicesDeviceIdList1452(d []interface{}) []edpt.ScaleoutClusterClusterDevicesDeviceIdList1452 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterClusterDevicesDeviceIdList1452, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterClusterDevicesDeviceIdList1452
		oi.Ip = in["ip"].(string)
		oi.Action = in["action"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterDbConfig1453(d []interface{}) edpt.ScaleoutClusterDbConfig1453 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterDbConfig1453
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ticktime = in["ticktime"].(int)
		ret.Initlimit = in["initlimit"].(int)
		ret.Synclimit = in["synclimit"].(int)
		ret.Minsessiontimeout = in["minsessiontimeout"].(int)
		ret.Maxsessiontimeout = in["maxsessiontimeout"].(int)
		ret.ClientRecvTimeout = in["client_recv_timeout"].(int)
		ret.Clientport = in["clientport"].(int)
		ret.LoopbackIntfSupport = in["loopback_intf_support"].(int)
		ret.BrokenDetectTimeout = in["broken_detect_timeout"].(int)
		ret.MoreElectionPacket = in["more_election_packet"].(int)
		ret.ElectConnTimeout = in["elect_conn_timeout"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterDeviceGroups1454(d []interface{}) edpt.ScaleoutClusterDeviceGroups1454 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterDeviceGroups1454
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
		ret.DeviceGroupList = getSliceScaleoutClusterDeviceGroupsDeviceGroupList1455(in["device_group_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupList1455(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupList1455 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupList1455, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupList1455
		oi.DeviceGroup = in["device_group"].(int)
		oi.DeviceIdList = getSliceScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456(in["device_id_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1456
		oi.DeviceIdStart = in["device_id_start"].(int)
		oi.DeviceIdEnd = in["device_id_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDevice1457(d []interface{}) edpt.ScaleoutClusterLocalDevice1457 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDevice1457
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Priority = in["priority"].(int)
		ret.Id1 = in["id1"].(int)
		ret.Action = in["action"].(string)
		ret.StartDelay = in["start_delay"].(int)
		ret.ClusterMode = in["cluster_mode"].(string)
		//omit uuid
		ret.L2Redirect = getObjectScaleoutClusterLocalDeviceL2Redirect1458(in["l2_redirect"].([]interface{}))
		ret.TrafficRedirection = getObjectScaleoutClusterLocalDeviceTrafficRedirection1459(in["traffic_redirection"].([]interface{}))
		ret.SessionSync = getObjectScaleoutClusterLocalDeviceSessionSync1467(in["session_sync"].([]interface{}))
		ret.ExcludeInterfaces = getObjectScaleoutClusterLocalDeviceExcludeInterfaces1474(in["exclude_interfaces"].([]interface{}))
		ret.TrackingTemplate = getObjectScaleoutClusterLocalDeviceTrackingTemplate1479(in["tracking_template"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceL2Redirect1458(d []interface{}) edpt.ScaleoutClusterLocalDeviceL2Redirect1458 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceL2Redirect1458
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RedirectEth = in["redirect_eth"].(int)
		ret.EthernetVlan = in["ethernet_vlan"].(int)
		ret.RedirectTrunk = in["redirect_trunk"].(int)
		ret.TrunkVlan = in["trunk_vlan"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirection1459(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirection1459 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirection1459
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		//omit uuid
		ret.Interfaces = getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1460(in["interfaces"].([]interface{}))
		ret.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1465(in["reachability_options"].([]interface{}))
		ret.Encap = getObjectScaleoutClusterLocalDeviceTrafficRedirectionEncap1466(in["encap"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1460(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1460 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1460
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1461
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1462
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1463
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1464
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1465(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1465 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1465
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionEncap1466(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionEncap1466 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionEncap1466
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(string)
		ret.UseV4Vxlan = in["use_v4_vxlan"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSync1467(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSync1467 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSync1467
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		//omit uuid
		ret.Interfaces = getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1468(in["interfaces"].([]interface{}))
		ret.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1473(in["reachability_options"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1468(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1468 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1468
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1469
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1470
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1471
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1472
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1473(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1473 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1473
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceExcludeInterfaces1474(d []interface{}) edpt.ScaleoutClusterLocalDeviceExcludeInterfaces1474 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceExcludeInterfaces1474
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1475
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1476
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1477
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1478
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrackingTemplate1479(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrackingTemplate1479 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrackingTemplate1479
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateList = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480(in["template_list"].([]interface{}))
		ret.MultiTemplateList = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482(in["multi_template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1480
		oi.Template = in["template"].(string)
		oi.IpVersion = in["ip_version"].(string)
		oi.ThresholdCfg = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481(in["threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1481
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1482
		oi.MultiTemplate = in["multi_template"].(string)
		oi.Template = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483(in["template"].([]interface{}))
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		oi.IpVersion = in["ip_version"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1483
		oi.TemplateName = in["template_name"].(string)
		oi.PartitionName = in["partition_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterServiceConfig1484(d []interface{}) edpt.ScaleoutClusterServiceConfig1484 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterServiceConfig1484
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DefaultUserGroupCount = in["default_user_group_count"].(int)
		ret.Enable = in["enable"].(int)
		//omit uuid
		ret.TemplateList = getSliceScaleoutClusterServiceConfigTemplateList1485(in["template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterServiceConfigTemplateList1485(d []interface{}) []edpt.ScaleoutClusterServiceConfigTemplateList1485 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterServiceConfigTemplateList1485, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterServiceConfigTemplateList1485
		oi.Name = in["name"].(string)
		oi.UserGroupCount = in["user_group_count"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterTrackingTemplate1486(d []interface{}) edpt.ScaleoutClusterTrackingTemplate1486 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterTrackingTemplate1486
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateList = getSliceScaleoutClusterTrackingTemplateTemplateList(in["template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterTrackingTemplateTemplateList(d []interface{}) []edpt.ScaleoutClusterTrackingTemplateTemplateList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterTrackingTemplateTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterTrackingTemplateTemplateList
		oi.Template = in["template"].(string)
		oi.ThresholdCfg = getSliceScaleoutClusterTrackingTemplateTemplateListThresholdCfg(in["threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterTrackingTemplateTemplateListThresholdCfg(d []interface{}) []edpt.ScaleoutClusterTrackingTemplateTemplateListThresholdCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterTrackingTemplateTemplateListThresholdCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterTrackingTemplateTemplateListThresholdCfg
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutCluster(d *schema.ResourceData) edpt.ScaleoutCluster {
	var ret edpt.ScaleoutCluster
	ret.Inst.ClusterDevices = getObjectScaleoutClusterClusterDevices1449(d.Get("cluster_devices").([]interface{}))
	ret.Inst.ClusterId = d.Get("cluster_id").(int)
	ret.Inst.DbConfig = getObjectScaleoutClusterDbConfig1453(d.Get("db_config").([]interface{}))
	ret.Inst.DeviceGroups = getObjectScaleoutClusterDeviceGroups1454(d.Get("device_groups").([]interface{}))
	ret.Inst.LocalDevice = getObjectScaleoutClusterLocalDevice1457(d.Get("local_device").([]interface{}))
	ret.Inst.ServiceConfig = getObjectScaleoutClusterServiceConfig1484(d.Get("service_config").([]interface{}))
	ret.Inst.SlogLevel = d.Get("slog_level").(int)
	ret.Inst.TrackingTemplate = getObjectScaleoutClusterTrackingTemplate1486(d.Get("tracking_template").([]interface{}))
	//omit uuid
	return ret
}
