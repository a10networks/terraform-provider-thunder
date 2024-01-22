package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceTrafficRedirection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_traffic_redirection`: Redirect interfaces and options\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceTrafficRedirectionRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionDelete,

		Schema: map[string]*schema.Schema{
			"follow_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Follow shared partition for redirection",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterLocalDeviceTrafficRedirectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceTrafficRedirectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1333(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1333 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1333
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.EthCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334(in["eth_cfg"].([]interface{}))
		ret.TrunkCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335(in["trunk_cfg"].([]interface{}))
		ret.VeCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336(in["ve_cfg"].([]interface{}))
		ret.LoopbackCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337(in["loopback_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg1334
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg1335
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg1336
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337 {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg1337
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1338(d []interface{}) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1338 {

	count1 := len(d)
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1338
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SkipDefaultRoute = in["skip_default_route"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceTrafficRedirection(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceTrafficRedirection {
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirection
	ret.Inst.FollowShared = d.Get("follow_shared").(int)
	ret.Inst.Interfaces = getObjectScaleoutClusterLocalDeviceTrafficRedirectionInterfaces1333(d.Get("interfaces").([]interface{}))
	ret.Inst.ReachabilityOptions = getObjectScaleoutClusterLocalDeviceTrafficRedirectionReachabilityOptions1338(d.Get("reachability_options").([]interface{}))
	//omit uuid
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
