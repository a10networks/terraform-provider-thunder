package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceSessionSyncInterfaces() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_session_sync_interfaces`: Session sync interfaces\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceSessionSyncInterfacesCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceSessionSyncInterfacesUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceSessionSyncInterfacesRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceSessionSyncInterfacesDelete,

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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"cluster_id": {
				Type: schema.TypeString, Required: true, Description: "ClusterId",
			},
		},
	}
}
func resourceScaleoutClusterLocalDeviceSessionSyncInterfacesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncInterfacesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncInterfaces(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceSessionSyncInterfacesRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceSessionSyncInterfacesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncInterfacesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncInterfaces(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceSessionSyncInterfacesRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceSessionSyncInterfacesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncInterfacesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncInterfaces(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceSessionSyncInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceSessionSyncInterfacesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceSessionSyncInterfaces(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceSessionSyncInterfaces(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces {
	var ret edpt.ScaleoutClusterLocalDeviceSessionSyncInterfaces
	ret.Inst.EthCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.LoopbackCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesLoopbackCfg(d.Get("loopback_cfg").([]interface{}))
	ret.Inst.TrunkCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceScaleoutClusterLocalDeviceSessionSyncInterfacesVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
