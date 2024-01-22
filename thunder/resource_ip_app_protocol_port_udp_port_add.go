package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortUdpPortAdd() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_udp_port_add`: Add application port\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortUdpPortAddCreate,
		UpdateContext: resourceIpAppProtocolPortUdpPortAddUpdate,
		ReadContext:   resourceIpAppProtocolPortUdpPortAddRead,
		DeleteContext: resourceIpAppProtocolPortUdpPortAddDelete,

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
func resourceIpAppProtocolPortUdpPortAddCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAdd(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortAddRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortAddUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAdd(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortAddRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortUdpPortAddDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAdd(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortAddRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAdd(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAppProtocolPortUdpPortAddAppNameList(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameList {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.Interface = getObjectIpAppProtocolPortUdpPortAddAppNameListInterface(in["interface"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectIpAppProtocolPortUdpPortAddAppNameListInterface(d []interface{}) edpt.IpAppProtocolPortUdpPortAddAppNameListInterface {

	count1 := len(d)
	var ret edpt.IpAppProtocolPortUdpPortAddAppNameListInterface
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Management = in["management"].(int)
		ret.VeCfg = getSliceIpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg(in["ve_cfg"].([]interface{}))
		ret.EthCfg = getSliceIpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg(in["eth_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortUdpPortAdd(d *schema.ResourceData) edpt.IpAppProtocolPortUdpPortAdd {
	var ret edpt.IpAppProtocolPortUdpPortAdd
	ret.Inst.AppNameList = getSliceIpAppProtocolPortUdpPortAddAppNameList(d.Get("app_name_list").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
