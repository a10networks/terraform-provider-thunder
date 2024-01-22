package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceSessionSync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_session_sync`: Interface for scaleout session sync\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceSessionSyncCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceSessionSyncUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceSessionSyncRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceSessionSyncDelete,

		Schema: map[string]*schema.Schema{
			"follow_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow shared partition for session sync",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterLocalDeviceSessionSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceSessionSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceSessionSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceSessionSyncRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceSessionSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceSessionSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1327(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1327 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces1327
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1328(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1329(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1330(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1331(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1328(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1328 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1328, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg1328
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1329(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1329 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1329, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg1329
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1330(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1330 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1330, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg1330
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1331(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1331 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1331, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg1331
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1332(d []interface{}) edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1332 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1332
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceSessionSync(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceSessionSync {
	var ret edpt.ScaleoutClusterLocalDeviceSessionSync
	ret.Inst.FollowShared = d.Get("follow_shared").(int)
	ret.Inst.Interfaces = getObjectScaleoutClusterLocalDeviceSessionSyncInterfaces1327(d.Get("interfaces").([]interface{}))
	ret.Inst.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceSessionSyncReachabilityOptions1332(d.Get("reachability_options").([]interface{}))
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
