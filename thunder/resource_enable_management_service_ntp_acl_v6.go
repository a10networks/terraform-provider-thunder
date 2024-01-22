package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceNtpAclV6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_ntp_acl_v6`: IPv6 ACL for NTP service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceNtpAclV6Create,
		UpdateContext: resourceEnableManagementServiceNtpAclV6Update,
		ReadContext:   resourceEnableManagementServiceNtpAclV6Read,
		DeleteContext: resourceEnableManagementServiceNtpAclV6Delete,

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
func resourceEnableManagementServiceNtpAclV6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceNtpAclV6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceNtpAclV6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceNtpAclV6Read(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceNtpAclV6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceNtpAclV6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceNtpAclV6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceNtpAclV6Read(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceNtpAclV6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceNtpAclV6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceNtpAclV6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceNtpAclV6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceNtpAclV6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceNtpAclV6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceNtpAclV6EthCfg(d []interface{}) []edpt.EnableManagementServiceNtpAclV6EthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceNtpAclV6EthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceNtpAclV6EthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceNtpAclV6TunnelCfg(d []interface{}) []edpt.EnableManagementServiceNtpAclV6TunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceNtpAclV6TunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceNtpAclV6TunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceNtpAclV6VeCfg(d []interface{}) []edpt.EnableManagementServiceNtpAclV6VeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceNtpAclV6VeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceNtpAclV6VeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceNtpAclV6(d *schema.ResourceData) edpt.EnableManagementServiceNtpAclV6 {
	var ret edpt.EnableManagementServiceNtpAclV6
	ret.Inst.AclName = d.Get("acl_name").(string)
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceNtpAclV6EthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceNtpAclV6TunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceNtpAclV6VeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
