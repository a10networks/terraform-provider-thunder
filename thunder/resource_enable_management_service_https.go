package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceHttps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_https`: HTTPS service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceHttpsCreate,
		UpdateContext: resourceEnableManagementServiceHttpsUpdate,
		ReadContext:   resourceEnableManagementServiceHttpsRead,
		DeleteContext: resourceEnableManagementServiceHttpsDelete,

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
func resourceEnableManagementServiceHttpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceHttpsRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceHttpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceHttpsRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceHttpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceHttpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceHttpsAclV4List(d []interface{}) []edpt.EnableManagementServiceHttpsAclV4List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV4List
		oi.AclId = in["acl_id"].(int)
		oi.EthCfg = getSliceEnableManagementServiceHttpsAclV4ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceHttpsAclV4ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceHttpsAclV4ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV4ListEthCfg(d []interface{}) []edpt.EnableManagementServiceHttpsAclV4ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV4ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV4ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV4ListVeCfg(d []interface{}) []edpt.EnableManagementServiceHttpsAclV4ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV4ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV4ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV4ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceHttpsAclV4ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV4ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV4ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV6List(d []interface{}) []edpt.EnableManagementServiceHttpsAclV6List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV6List
		oi.AclName = in["acl_name"].(string)
		oi.EthCfg = getSliceEnableManagementServiceHttpsAclV6ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceHttpsAclV6ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceHttpsAclV6ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV6ListEthCfg(d []interface{}) []edpt.EnableManagementServiceHttpsAclV6ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV6ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV6ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV6ListVeCfg(d []interface{}) []edpt.EnableManagementServiceHttpsAclV6ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV6ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV6ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsAclV6ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceHttpsAclV6ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsAclV6ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsAclV6ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsEthCfg(d []interface{}) []edpt.EnableManagementServiceHttpsEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsTunnelCfg(d []interface{}) []edpt.EnableManagementServiceHttpsTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpsVeCfg(d []interface{}) []edpt.EnableManagementServiceHttpsVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpsVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpsVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceHttps(d *schema.ResourceData) edpt.EnableManagementServiceHttps {
	var ret edpt.EnableManagementServiceHttps
	ret.Inst.AclV4List = getSliceEnableManagementServiceHttpsAclV4List(d.Get("acl_v4_list").([]interface{}))
	ret.Inst.AclV6List = getSliceEnableManagementServiceHttpsAclV6List(d.Get("acl_v6_list").([]interface{}))
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceHttpsEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceHttpsTunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceHttpsVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
