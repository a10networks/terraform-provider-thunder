package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisableManagementServicePing() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_disable_management_service_ping`: Ping service\n\n__PLACEHOLDER__",
		CreateContext: resourceDisableManagementServicePingCreate,
		UpdateContext: resourceDisableManagementServicePingUpdate,
		ReadContext:   resourceDisableManagementServicePingRead,
		DeleteContext: resourceDisableManagementServicePingDelete,

		Schema: map[string]*schema.Schema{
			"all_data_intf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All Data Interfaces",
			},
			"eth_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Ethernet Interface number)",
						},
						"ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
						},
					},
				},
			},
			"management": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Management Interface",
			},
			"tunnel_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel_start": {
							Type: schema.TypeInt, Optional: true, Description: "tunnel port (tunnel Interface number)",
						},
						"tunnel_end": {
							Type: schema.TypeInt, Optional: true, Description: "tunnel port",
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
						"ve_start": {
							Type: schema.TypeInt, Optional: true, Description: "VE port (VE Interface number)",
						},
						"ve_end": {
							Type: schema.TypeInt, Optional: true, Description: "VE port",
						},
					},
				},
			},
		},
	}
}
func resourceDisableManagementServicePingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServicePingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServicePing(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServicePingRead(ctx, d, meta)
	}
	return diags
}

func resourceDisableManagementServicePingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServicePingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServicePing(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServicePingRead(ctx, d, meta)
	}
	return diags
}
func resourceDisableManagementServicePingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServicePingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServicePing(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDisableManagementServicePingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServicePingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServicePing(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDisableManagementServicePingEthCfg(d []interface{}) []edpt.DisableManagementServicePingEthCfg {

	count1 := len(d)
	ret := make([]edpt.DisableManagementServicePingEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DisableManagementServicePingEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDisableManagementServicePingTunnelCfg(d []interface{}) []edpt.DisableManagementServicePingTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.DisableManagementServicePingTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DisableManagementServicePingTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDisableManagementServicePingVeCfg(d []interface{}) []edpt.DisableManagementServicePingVeCfg {

	count1 := len(d)
	ret := make([]edpt.DisableManagementServicePingVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DisableManagementServicePingVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDisableManagementServicePing(d *schema.ResourceData) edpt.DisableManagementServicePing {
	var ret edpt.DisableManagementServicePing
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceDisableManagementServicePingEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	ret.Inst.TunnelCfg = getSliceDisableManagementServicePingTunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceDisableManagementServicePingVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
