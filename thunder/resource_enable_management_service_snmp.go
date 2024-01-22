package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableManagementServiceSnmp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_management_service_snmp`: SNMP service\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableManagementServiceSnmpCreate,
		UpdateContext: resourceEnableManagementServiceSnmpUpdate,
		ReadContext:   resourceEnableManagementServiceSnmpRead,
		DeleteContext: resourceEnableManagementServiceSnmpDelete,

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
func resourceEnableManagementServiceSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceSnmpRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableManagementServiceSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableManagementServiceSnmpRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableManagementServiceSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableManagementServiceSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableManagementServiceSnmpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableManagementServiceSnmp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceEnableManagementServiceSnmpAclV4List(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4List
		oi.AclId = in["acl_id"].(int)
		oi.EthCfg = getSliceEnableManagementServiceSnmpAclV4ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceSnmpAclV4ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceSnmpAclV4ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV4ListEthCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV4ListVeCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV4ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV4ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV4ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV4ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV6List(d []interface{}) []edpt.EnableManagementServiceSnmpAclV6List {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV6List
		oi.AclName = in["acl_name"].(string)
		oi.EthCfg = getSliceEnableManagementServiceSnmpAclV6ListEthCfg(in["eth_cfg"].([]interface{}))
		oi.VeCfg = getSliceEnableManagementServiceSnmpAclV6ListVeCfg(in["ve_cfg"].([]interface{}))
		oi.TunnelCfg = getSliceEnableManagementServiceSnmpAclV6ListTunnelCfg(in["tunnel_cfg"].([]interface{}))
		oi.Management = in["management"].(int)
		oi.AllDataIntf = in["all_data_intf"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV6ListEthCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV6ListEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV6ListEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV6ListEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV6ListVeCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV6ListVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV6ListVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV6ListVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpAclV6ListTunnelCfg(d []interface{}) []edpt.EnableManagementServiceSnmpAclV6ListTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpAclV6ListTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpAclV6ListTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpEthCfg(d []interface{}) []edpt.EnableManagementServiceSnmpEthCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpTunnelCfg(d []interface{}) []edpt.EnableManagementServiceSnmpTunnelCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpTunnelCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpTunnelCfg
		oi.TunnelStart = in["tunnel_start"].(int)
		oi.TunnelEnd = in["tunnel_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceEnableManagementServiceSnmpVeCfg(d []interface{}) []edpt.EnableManagementServiceSnmpVeCfg {

	count1 := len(d)
	ret := make([]edpt.EnableManagementServiceSnmpVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.EnableManagementServiceSnmpVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointEnableManagementServiceSnmp(d *schema.ResourceData) edpt.EnableManagementServiceSnmp {
	var ret edpt.EnableManagementServiceSnmp
	ret.Inst.AclV4List = getSliceEnableManagementServiceSnmpAclV4List(d.Get("acl_v4_list").([]interface{}))
	ret.Inst.AclV6List = getSliceEnableManagementServiceSnmpAclV6List(d.Get("acl_v6_list").([]interface{}))
	ret.Inst.AllDataIntf = d.Get("all_data_intf").(int)
	ret.Inst.EthCfg = getSliceEnableManagementServiceSnmpEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.TunnelCfg = getSliceEnableManagementServiceSnmpTunnelCfg(d.Get("tunnel_cfg").([]interface{}))
	//omit uuid
	ret.Inst.VeCfg = getSliceEnableManagementServiceSnmpVeCfg(d.Get("ve_cfg").([]interface{}))
	return ret
}
