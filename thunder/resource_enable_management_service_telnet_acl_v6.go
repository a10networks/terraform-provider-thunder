package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceTelnetAclV6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_telnet_acl_v6`: IPv6 ACL for Telnet service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceTelnetAclV6Create,
		UpdateContext: resourceEnableManagementServiceTelnetAclV6Update,
		ReadContext:   resourceEnableManagementServiceTelnetAclV6Read,
		DeleteContext: resourceEnableManagementServiceTelnetAclV6Delete,

		Schema: map[string]*schema.Schema{
			"acl_name": {
				Type: schema.TypeString, Required: true, Description: "ACL name",
			},
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceEnableManagementServiceTelnetAclV6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceTelnetAclV6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceTelnetAclV6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceTelnetAclV6Read(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceTelnetAclV6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceTelnetAclV6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceTelnetAclV6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceTelnetAclV6Read(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceTelnetAclV6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceTelnetAclV6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceTelnetAclV6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceTelnetAclV6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceTelnetAclV6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceTelnetAclV6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceTelnetAclV6EthCfg(d []interface{}) []edpt.EnableManagementServiceTelnetAclV6EthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceTelnetAclV6EthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceTelnetAclV6EthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceTelnetAclV6TunnelCfg(d []interface{}) []edpt.EnableManagementServiceTelnetAclV6TunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceTelnetAclV6TunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceTelnetAclV6TunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceTelnetAclV6VeCfg(d []interface{}) []edpt.EnableManagementServiceTelnetAclV6VeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceTelnetAclV6VeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceTelnetAclV6VeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceTelnetAclV6(d *schema.ResourceData) edpt.EnableManagementServiceTelnetAclV6 {
	var ret edpt.EnableManagementServiceTelnetAclV6
	ret.Inst.AclName = d.Get("acl_name").(string)
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceTelnetAclV6EthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceTelnetAclV6TunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceTelnetAclV6VeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
