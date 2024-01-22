package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortTcpPortAdd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_tcp_port_add`: Add application port\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortTcpPortAddCreate,
		UpdateContext: resourceIpAppProtocolPortTcpPortAddUpdate,
		ReadContext:   resourceIpAppProtocolPortTcpPortAddRead,
		DeleteContext: resourceIpAppProtocolPortTcpPortAddDelete,

		Schema: map[string]*schema.Schema{
			"app_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "Application Protocol Port Name",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"interface": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"management": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Management Interface",
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
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"port": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpAppProtocolPortTcpPortAddCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAdd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortAddRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortAddUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAdd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortAddRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortTcpPortAddDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAdd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortAddRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAdd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAppProtocolPortTcpPortAddAppNameList(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameList {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.Interface = getObjectIpAppProtocolPortTcpPortAddAppNameListInterface(in["interface"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectIpAppProtocolPortTcpPortAddAppNameListInterface(d []interface{}) edpt.IpAppProtocolPortTcpPortAddAppNameListInterface {

	count1 := len(d)
	var ret edpt.IpAppProtocolPortTcpPortAddAppNameListInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Management = in["management"].(int)
		ret.VeCfg = getSliceIpAppProtocolPortTcpPortAddAppNameListInterfaceVeCfg(in["ve_cfg"].([]interface{}))
		ret.EthCfg = getSliceIpAppProtocolPortTcpPortAddAppNameListInterfaceEthCfg(in["eth_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortAddAppNameListInterfaceVeCfg(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameListInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameListInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameListInterfaceVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortAddAppNameListInterfaceEthCfg(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameListInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameListInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameListInterfaceEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortTcpPortAdd(d *schema.ResourceData) edpt.IpAppProtocolPortTcpPortAdd {
	var ret edpt.IpAppProtocolPortTcpPortAdd
	ret.Inst.AppNameList = getSliceIpAppProtocolPortTcpPortAddAppNameList(d.Get("app_name_list").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
