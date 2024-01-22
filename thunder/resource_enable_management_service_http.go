package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_http`: HTTP service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceHttpCreate,
		UpdateContext: resourceEnableManagementServiceHttpUpdate,
		ReadContext:   resourceEnableManagementServiceHttpRead,
		DeleteContext: resourceEnableManagementServiceHttpDelete,

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
func resourceEnableManagementServiceHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceHttpAclV4List(d []interface{}) []edpt.EnableManagementServiceHttpAclV4List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV4List
		oi.AclId = in["acl_id"].(int)
		oi.EthCfg = getSliceEnableManagementServiceHttpAclV4ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceHttpAclV4ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceHttpAclV4ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV4ListEthCfg(d []interface{}) []edpt.EnableManagementServiceHttpAclV4ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV4ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV4ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV4ListVeCfg(d []interface{}) []edpt.EnableManagementServiceHttpAclV4ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV4ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV4ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV4ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceHttpAclV4ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV4ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV4ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV6List(d []interface{}) []edpt.EnableManagementServiceHttpAclV6List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV6List
		oi.AclName = in["acl_name"].(string)
		oi.EthCfg = getSliceEnableManagementServiceHttpAclV6ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceHttpAclV6ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceHttpAclV6ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV6ListEthCfg(d []interface{}) []edpt.EnableManagementServiceHttpAclV6ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV6ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV6ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV6ListVeCfg(d []interface{}) []edpt.EnableManagementServiceHttpAclV6ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV6ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV6ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpAclV6ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceHttpAclV6ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpAclV6ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpAclV6ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpEthCfg(d []interface{}) []edpt.EnableManagementServiceHttpEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpTunnelCfg(d []interface{}) []edpt.EnableManagementServiceHttpTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceHttpVeCfg(d []interface{}) []edpt.EnableManagementServiceHttpVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceHttpVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceHttpVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceHttp(d *schema.ResourceData) edpt.EnableManagementServiceHttp {
	var ret edpt.EnableManagementServiceHttp
	ret.Inst.AclV4List = getSliceEnableManagementServiceHttpAclV4List(d.Get("acl_v4_list").([]interface{}))
	ret.Inst.AclV6List = getSliceEnableManagementServiceHttpAclV6List(d.Get("acl_v6_list").([]interface{}))
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceHttpEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceHttpTunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceHttpVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
