package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServicePing() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_ping`: Ping service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServicePingCreate,
		UpdateContext: resourceEnableManagementServicePingUpdate,
		ReadContext:   resourceEnableManagementServicePingRead,
		DeleteContext: resourceEnableManagementServicePingDelete,

		Schema: map[string]*schema.Schema{
			"acl_v4_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_id": {
							Type: schema.TypeInt, Required: true, Description: "ACL id",
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
						"management": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Management Interface",
						},
						"all_data_intf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All Data Interfaces",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"acl_v6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"acl_name": {
							Type: schema.TypeString, Required: true, Description: "ACL name",
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
						"management": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Management Interface",
						},
						"all_data_intf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "All Data Interfaces",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceEnableManagementServicePingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServicePingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServicePing(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServicePingRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServicePingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServicePingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServicePing(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServicePingRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServicePingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServicePingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServicePing(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServicePingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServicePingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServicePing(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServicePingAclV4List(d []interface{}) []edpt.EnableManagementServicePingAclV4List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV4List
		oi.AclId = in["acl_id"].(int)
		oi.EthCfg = getSliceEnableManagementServicePingAclV4ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServicePingAclV4ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServicePingAclV4ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV4ListEthCfg(d []interface{}) []edpt.EnableManagementServicePingAclV4ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV4ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV4ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV4ListVeCfg(d []interface{}) []edpt.EnableManagementServicePingAclV4ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV4ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV4ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV4ListTunnelCfg(d []interface{}) []edpt.EnableManagementServicePingAclV4ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV4ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV4ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV6List(d []interface{}) []edpt.EnableManagementServicePingAclV6List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV6List
		oi.AclName = in["acl_name"].(string)
		oi.EthCfg = getSliceEnableManagementServicePingAclV6ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServicePingAclV6ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServicePingAclV6ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV6ListEthCfg(d []interface{}) []edpt.EnableManagementServicePingAclV6ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV6ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV6ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV6ListVeCfg(d []interface{}) []edpt.EnableManagementServicePingAclV6ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV6ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV6ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServicePingAclV6ListTunnelCfg(d []interface{}) []edpt.EnableManagementServicePingAclV6ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServicePingAclV6ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServicePingAclV6ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServicePing(d *schema.ResourceData) edpt.EnableManagementServicePing {
	var ret edpt.EnableManagementServicePing
	ret.Inst.AclV4List = getSliceEnableManagementServicePingAclV4List(d.Get("acl_v4_list").([]interface{}))
	ret.Inst.AclV6List = getSliceEnableManagementServicePingAclV6List(d.Get("acl_v6_list").([]interface{}))
	//omit uuid
	return ret
}
