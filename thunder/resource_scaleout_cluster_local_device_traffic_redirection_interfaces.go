package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfaces() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_scaleout_cluster_local_device_traffic_redirection_interfaces`: Redirect interfaces\n\n__PLACEHOLDER__",
		CreateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesCreate,
		UpdateContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesUpdate,
		ReadContext:   resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesRead,
		DeleteContext: resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesDelete,

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
func resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionInterfaces(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesRead(ctx, d, meta)
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionInterfaces(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesRead(ctx, d, meta)
	}
	return diags
}
func resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionInterfaces(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionInterfaces(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg
		oi.Ethernet = in["ethernet"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg
		oi.Loopback = in["loopback"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg
		oi.Trunk = in["trunk"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg(d []interface{}) []edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg {

	count1 := len(d)
	ret := make([]edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg
		oi.Ve = in["ve"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointScaleoutClusterLocalDeviceTrafficRedirectionInterfaces(d *schema.ResourceData) edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces {
	var ret edpt.ScaleoutClusterLocalDeviceTrafficRedirectionInterfaces
	ret.Inst.EthCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.LoopbackCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesLoopbackCfg(d.Get("loopback_cfg").([]interface{}))
	ret.Inst.TrunkCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesTrunkCfg(d.Get("trunk_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceScaleoutClusterLocalDeviceTrafficRedirectionInterfacesVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.ClusterId = d.Get("cluster_id").(string)
	return ret
}
