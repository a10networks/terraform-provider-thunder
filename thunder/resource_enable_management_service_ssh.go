package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceSsh() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_ssh`: SSH service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceSshCreate,
		UpdateContext: resourceEnableManagementServiceSshUpdate,
		ReadContext:   resourceEnableManagementServiceSshRead,
		DeleteContext: resourceEnableManagementServiceSshDelete,

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
func resourceEnableManagementServiceSshCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSshCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSsh(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceSshRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceSshUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSshUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSsh(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceSshRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceSshDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSshDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSsh(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceSshRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSshRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSsh(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceSshAclV4List(d []interface{}) []edpt.EnableManagementServiceSshAclV4List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV4List
		oi.AclId = in["acl_id"].(int)
		oi.EthCfg = getSliceEnableManagementServiceSshAclV4ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceSshAclV4ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceSshAclV4ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV4ListEthCfg(d []interface{}) []edpt.EnableManagementServiceSshAclV4ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV4ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV4ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV4ListVeCfg(d []interface{}) []edpt.EnableManagementServiceSshAclV4ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV4ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV4ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV4ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceSshAclV4ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV4ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV4ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV6List(d []interface{}) []edpt.EnableManagementServiceSshAclV6List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV6List
		oi.AclName = in["acl_name"].(string)
		oi.EthCfg = getSliceEnableManagementServiceSshAclV6ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceSshAclV6ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceSshAclV6ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV6ListEthCfg(d []interface{}) []edpt.EnableManagementServiceSshAclV6ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV6ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV6ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV6ListVeCfg(d []interface{}) []edpt.EnableManagementServiceSshAclV6ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV6ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV6ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshAclV6ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceSshAclV6ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshAclV6ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshAclV6ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshEthCfg(d []interface{}) []edpt.EnableManagementServiceSshEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshTunnelCfg(d []interface{}) []edpt.EnableManagementServiceSshTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSshVeCfg(d []interface{}) []edpt.EnableManagementServiceSshVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSshVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSshVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceSsh(d *schema.ResourceData) edpt.EnableManagementServiceSsh {
	var ret edpt.EnableManagementServiceSsh
	ret.Inst.AclV4List = getSliceEnableManagementServiceSshAclV4List(d.Get("acl_v4_list").([]interface{}))
	ret.Inst.AclV6List = getSliceEnableManagementServiceSshAclV6List(d.Get("acl_v6_list").([]interface{}))
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceSshEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceSshTunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceSshVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
