package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceExcludeInterfaces() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_exclude_interfaces`: Interfaces for peers to exclude from scaleout redirection and session sync\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceExcludeInterfacesCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceExcludeInterfacesUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceExcludeInterfacesRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceExcludeInterfacesDelete,

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
							Type: schema.TypeInt, Optional: true, Description: "Loopback Interface (Loopback interface number)",
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
func resourceScaleoutClusterLocalDeviceExcludeInterfacesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceExcludeInterfacesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceExcludeInterfaces(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceExcludeInterfacesRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceExcludeInterfacesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceExcludeInterfacesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceExcludeInterfaces(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceExcludeInterfacesRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceExcludeInterfacesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceExcludeInterfacesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceExcludeInterfaces(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceExcludeInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceExcludeInterfacesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceExcludeInterfaces(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesEthCfg
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceExcludeInterfacesVeCfg
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceExcludeInterfaces(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceExcludeInterfaces {
	var ret edpt.ScaleoutClusterLocalDeviceExcludeInterfaces
	ret.Inst.EthCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.LoopbackCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesLoopbackCfg(d.Get("loopback_cfg").([]interface{}))
	ret.Inst.TrunkCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceScaleoutClusterLocalDeviceExcludeInterfacesVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
