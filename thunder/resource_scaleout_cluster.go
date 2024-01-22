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
									"bucket_count": {
										Type: schema.TypeInt, Optional: true, Default: 256, Description: "Number of traffic buckets",
									},
									"device_group": {
										Type: schema.TypeInt, Optional: true, Description: "Device group id",
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

func getObjectScaleoutClusterClusterDevices1360(d []interface{}) edpt.ScaleoutClusterClusterDevices1360 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterClusterDevices1360
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
		ret.MinimumNodes = getObjectScaleoutClusterClusterDevicesMinimumNodes1361(in["minimum_nodes"].([]interface{}))
		ret.ClusterDiscoveryTimeout = getObjectScaleoutClusterClusterDevicesClusterDiscoveryTimeout1362(in["cluster_discovery_timeout"].([]interface{}))
		ret.DeviceIdList = getSliceScaleoutClusterClusterDevicesDeviceIdList1363(in["device_id_list"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterClusterDevicesMinimumNodes1361(d []interface{}) edpt.ScaleoutClusterClusterDevicesMinimumNodes1361 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterClusterDevicesMinimumNodes1361
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinimumNodesNum = in["minimum_nodes_num"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterClusterDevicesClusterDiscoveryTimeout1362(d []interface{}) edpt.ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1362 {

	var ret edpt.ScaleoutClusterClusterDevicesClusterDiscoveryTimeout1362
	return ret
}

func getSliceScaleoutClusterClusterDevicesDeviceIdList1363(d []interface{}) []edpt.ScaleoutClusterClusterDevicesDeviceIdList1363 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterClusterDevicesDeviceIdList1363, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterClusterDevicesDeviceIdList1363
		oi.Ip = in["ip"].(string)
		oi.Action = in["action"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterDbConfig1364(d []interface{}) edpt.ScaleoutClusterDbConfig1364 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterDbConfig1364
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

func getObjectScaleoutClusterDeviceGroups1365(d []interface{}) edpt.ScaleoutClusterDeviceGroups1365 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterDeviceGroups1365
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
		ret.DeviceGroupList = getSliceScaleoutClusterDeviceGroupsDeviceGroupList1366(in["device_group_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupList1366(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupList1366 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupList1366, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupList1366
		oi.DeviceGroup = in["device_group"].(int)
		oi.DeviceIdList = getSliceScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367(in["device_id_list"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367(d []interface{}) []edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterDeviceGroupsDeviceGroupListDeviceIdList1367
		oi.DeviceIdStart = in["device_id_start"].(int)
		oi.DeviceIdEnd = in["device_id_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDevice1368(d []interface{}) edpt.ScaleoutClusterLocalDevice1368 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDevice1368
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Priority = in["priority"].(int)
		ret.Id1 = in["id1"].(int)
		ret.Action = in["action"].(string)
		ret.StartDelay = in["start_delay"].(int)
		ret.ClusterMode = in["cluster_mode"].(string)
		//omit uuid
		ret.L2Redirect = getObjectScaleoutClusterLocalDeviceL2Redirect1369(in["l2_redirect"].([]interface{}))
		ret.TrafficRedirection = getObjectScaleoutClusterLocalDeviceTrafficRedirection1370(in["traffic_redirection"].([]interface{}))
		ret.SessionSync = getObjectScaleoutClusterLocalDeviceSessionSync1377(in["session_sync"].([]interface{}))
		ret.ExcludeInterfaces = getObjectScaleoutClusterLocalDeviceExcludeInterfaces1384(in["exclude_interfaces"].([]interface{}))
		ret.TrackingTemplate = getObjectScaleoutClusterLocalDeviceTrackingTemplate1389(in["tracking_template"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceL2Redirect1369(d []interface{}) edpt.ScaleoutClusterLocalDeviceL2Redirect1369 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceL2Redirect1369
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

func getObjectScaleoutClusterLocalDeviceTrafficRedirection1370(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirection1370 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirection1370
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		//omit uuid
		ret.Interfaces = getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1371(in["interfaces"].([]interface{}))
		ret.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1376(in["reachability_options"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1371(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1371 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1371
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1372
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1373
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1374
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1375
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1376(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1376 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1376
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSync1377(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSync1377 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSync1377
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		//omit uuid
		ret.Interfaces = getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1378(in["interfaces"].([]interface{}))
		ret.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1383(in["reachability_options"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1378(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1378 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1378
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1379
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1380
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1381
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1382
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1383(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1383 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1383
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceExcludeInterfaces1384(d []interface{}) edpt.ScaleoutClusterLocalDeviceExcludeInterfaces1384 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceExcludeInterfaces1384
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1385
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1386
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1387
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1388
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrackingTemplate1389(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrackingTemplate1389 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrackingTemplate1389
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateList = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390(in["template_list"].([]interface{}))
		ret.MultiTemplateList = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392(in["multi_template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList1390
		oi.Template = in["template"].(string)
		oi.ThresholdCfg = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391(in["threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg1391
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList1392
		oi.MultiTemplate = in["multi_template"].(string)
		oi.Template = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393(in["template"].([]interface{}))
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate1393
		oi.TemplateName = in["template_name"].(string)
		oi.PartitionName = in["partition_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterServiceConfig1394(d []interface{}) edpt.ScaleoutClusterServiceConfig1394 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterServiceConfig1394
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		//omit uuid
		ret.TemplateList = getSliceScaleoutClusterServiceConfigTemplateList1395(in["template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterServiceConfigTemplateList1395(d []interface{}) []edpt.ScaleoutClusterServiceConfigTemplateList1395 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterServiceConfigTemplateList1395, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterServiceConfigTemplateList1395
		oi.Name = in["name"].(string)
		oi.BucketCount = in["bucket_count"].(int)
		oi.DeviceGroup = in["device_group"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterTrackingTemplate1396(d []interface{}) edpt.ScaleoutClusterTrackingTemplate1396 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterTrackingTemplate1396
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
	ret.Inst.ClusterDevices = getObjectScaleoutClusterClusterDevices1360(d.Get("cluster_devices").([]interface{}))
	ret.Inst.ClusterId = d.Get("cluster_id").(int)
	ret.Inst.DbConfig = getObjectScaleoutClusterDbConfig1364(d.Get("db_config").([]interface{}))
	ret.Inst.DeviceGroups = getObjectScaleoutClusterDeviceGroups1365(d.Get("device_groups").([]interface{}))
	ret.Inst.LocalDevice = getObjectScaleoutClusterLocalDevice1368(d.Get("local_device").([]interface{}))
	ret.Inst.ServiceConfig = getObjectScaleoutClusterServiceConfig1394(d.Get("service_config").([]interface{}))
	ret.Inst.SlogLevel = d.Get("slog_level").(int)
	ret.Inst.TrackingTemplate = getObjectScaleoutClusterTrackingTemplate1396(d.Get("tracking_template").([]interface{}))
	//omit uuid
	return ret
}
