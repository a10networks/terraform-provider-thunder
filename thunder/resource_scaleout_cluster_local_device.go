package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDevice() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device`: Local device configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"cluster_mode": {
				Type: schema.TypeString, Optional: true, Default: "layer-2", Description: "'layer-2': Nodes in cluster are layer 2 connected (default mode); 'layer-3': Nodes in cluster are l3 connected;",
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
			"id1": {
				Type: schema.TypeInt, Optional: true, Description: "",
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
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "",
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
			"start_delay": {
				Type: schema.TypeInt, Optional: true, Description: "",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterLocalDeviceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDevice(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDevice(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDevice(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDevice(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectScaleoutClusterLocalDeviceExcludeInterfaces1339(d []interface{}) edpt.ScaleoutClusterLocalDeviceExcludeInterfaces1339 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceExcludeInterfaces1339
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg1340
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg1341
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg1342
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg1343
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceL2Redirect1344(d []interface{}) edpt.ScaleoutClusterLocalDeviceL2Redirect1344 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceL2Redirect1344
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

func getObjectScaleoutClusterLocalDeviceSessionSync1345(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSync1345 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSync1345
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		//omit uuid
		ret.Interfaces = getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1346(in["interfaces"].([]interface{}))
		ret.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1351(in["reachability_options"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1346(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1346 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1346
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1347
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1348
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1349
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1350
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1351(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1351 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1351
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrackingTemplate1352(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrackingTemplate1352 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrackingTemplate1352
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TemplateList = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateList(in["template_list"].([]interface{}))
		ret.MultiTemplateList = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList(in["multi_template_list"].([]interface{}))
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateList(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateList
		oi.Template = in["template"].(string)
		oi.ThresholdCfg = getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg(in["threshold_cfg"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateTemplateListThresholdCfg
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateList
		oi.MultiTemplate = in["multi_template"].(string)
		oi.Template = getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate(in["template"].([]interface{}))
		oi.Threshold = in["threshold"].(int)
		oi.Action = in["action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrackingTemplateMultiTemplateListTemplate
		oi.TemplateName = in["template_name"].(string)
		oi.PartitionName = in["partition_name"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirection1353(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirection1353 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirection1353
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.FollowShared = in["follow_shared"].(int)
		//omit uuid
		ret.Interfaces = getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1354(in["interfaces"].([]interface{}))
		ret.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1359(in["reachability_options"].([]interface{}))
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1354(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1354 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1354
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1355
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1356
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1357
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1358
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1359(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1359 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1359
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDevice(d *schema.ResourceData) edpt.ScaleoutClusterLocalDevice {
	var ret edpt.ScaleoutClusterLocalDevice
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ClusterMode = d.Get("cluster_mode").(string)
	ret.Inst.ExcludeInterfaces = getObjectScaleoutClusterLocalDeviceExcludeInterfaces1339(d.Get("exclude_interfaces").([]interface{}))
	ret.Inst.Id1 = d.Get("id1").(int)
	ret.Inst.L2Redirect = getObjectScaleoutClusterLocalDeviceL2Redirect1344(d.Get("l2_redirect").([]interface{}))
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.SessionSync = getObjectScaleoutClusterLocalDeviceSessionSync1345(d.Get("session_sync").([]interface{}))
	ret.Inst.StartDelay = d.Get("start_delay").(int)
	ret.Inst.TrackingTemplate = getObjectScaleoutClusterLocalDeviceTrackingTemplate1352(d.Get("tracking_template").([]interface{}))
	ret.Inst.TrafficRedirection = getObjectScaleoutClusterLocalDeviceTrafficRedirection1353(d.Get("traffic_redirection").([]interface{}))
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
