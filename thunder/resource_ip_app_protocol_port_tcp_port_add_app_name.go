package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortTcpPortAddAppName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_tcp_port_add_app_name`: Application protocol name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortTcpPortAddAppNameCreate,
		UpdateContext: resourceIpAppProtocolPortTcpPortAddAppNameUpdate,
		ReadContext:   resourceIpAppProtocolPortTcpPortAddAppNameRead,
		DeleteContext: resourceIpAppProtocolPortTcpPortAddAppNameDelete,

		Schema: map[string]*schema.Schema{
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Application Protocol Port Name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"port": {
				Type: schema.TypeString, Required: true, Description: "Port",
			},
		},
	}
}
func resourceIpAppProtocolPortTcpPortAddAppNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortAddAppNameRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortAddAppNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortAddAppNameRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortTcpPortAddAppNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortAddAppNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpAppProtocolPortTcpPortAddAppNameInterface1030(d []interface{}) edpt.IpAppProtocolPortTcpPortAddAppNameInterface1030 {

	count1 := len(d)
	var ret edpt.IpAppProtocolPortTcpPortAddAppNameInterface1030
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Management = in["management"].(int)
		ret.VeCfg = getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg1031(in["ve_cfg"].([]interface{}))
		ret.EthCfg = getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg1032(in["eth_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg1031(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg1031 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg1031, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg1031
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg1032(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg1032 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg1032, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg1032
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortTcpPortAddAppName(d *schema.ResourceData) edpt.IpAppProtocolPortTcpPortAddAppName {
	var ret edpt.IpAppProtocolPortTcpPortAddAppName
	ret.Inst.Interface = getObjectIpAppProtocolPortTcpPortAddAppNameInterface1030(d.Get("interface").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
