package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisableManagementServiceNtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_disable_management_service_ntp`: NTP service\n\n__PLACEHOLDER__",
		CreateContext: resourceDisableManagementServiceNtpCreate,
		UpdateContext: resourceDisableManagementServiceNtpUpdate,
		ReadContext:   resourceDisableManagementServiceNtpRead,
		DeleteContext: resourceDisableManagementServiceNtpDelete,

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
func resourceDisableManagementServiceNtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceNtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceNtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceNtpRead(ctx, d, meta)
	}
	return diags
}

func resourceDisableManagementServiceNtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceNtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceNtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDisableManagementServiceNtpRead(ctx, d, meta)
	}
	return diags
}
func resourceDisableManagementServiceNtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceNtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceNtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDisableManagementServiceNtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDisableManagementServiceNtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDisableManagementServiceNtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDisableManagementServiceNtpEthCfg(d []interface{}) []edpt.DisableManagementServiceNtpEthCfg {

	count1 := len(d)
	ret := make([]edpt.DisableManagementServiceNtpEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DisableManagementServiceNtpEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDisableManagementServiceNtpTunnelCfg(d []interface{}) []edpt.DisableManagementServiceNtpTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.DisableManagementServiceNtpTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DisableManagementServiceNtpTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDisableManagementServiceNtpVeCfg(d []interface{}) []edpt.DisableManagementServiceNtpVeCfg {

	count1 := len(d)
	ret := make([]edpt.DisableManagementServiceNtpVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DisableManagementServiceNtpVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDisableManagementServiceNtp(d *schema.ResourceData) edpt.DisableManagementServiceNtp {
	var ret edpt.DisableManagementServiceNtp
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceDisableManagementServiceNtpEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	ret.Inst.TunnelCfg = getSliceDisableManagementServiceNtpTunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceDisableManagementServiceNtpVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
