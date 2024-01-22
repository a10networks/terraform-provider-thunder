package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceSnmpAclV4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_snmp_acl_v4`: IPv4 ACL for SNMP service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceSnmpAclV4Create,
		UpdateContext: resourceEnableManagementServiceSnmpAclV4Update,
		ReadContext:   resourceEnableManagementServiceSnmpAclV4Read,
		DeleteContext: resourceEnableManagementServiceSnmpAclV4Delete,

		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type: schema.TypeInt, Required: true, Description: "ACL id",
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
func resourceEnableManagementServiceSnmpAclV4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpAclV4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmpAclV4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceSnmpAclV4Read(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceSnmpAclV4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpAclV4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmpAclV4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceSnmpAclV4Read(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceSnmpAclV4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpAclV4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmpAclV4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceSnmpAclV4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpAclV4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmpAclV4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceSnmpAclV4EthCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4EthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4EthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4EthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV4TunnelCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4TunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4TunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4TunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV4VeCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4VeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4VeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4VeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceSnmpAclV4(d *schema.ResourceData) edpt.EnableManagementServiceSnmpAclV4 {
	var ret edpt.EnableManagementServiceSnmpAclV4
	ret.Inst.AclId = d.Get("acl_id").(int)
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceSnmpAclV4EthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceSnmpAclV4TunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceSnmpAclV4VeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
