package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortTcpPortDisable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_tcp_port_disable`: Disable application port\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortTcpPortDisableCreate,
		UpdateContext: resourceIpAppProtocolPortTcpPortDisableUpdate,
		ReadContext:   resourceIpAppProtocolPortTcpPortDisableRead,
		DeleteContext: resourceIpAppProtocolPortTcpPortDisableDelete,

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
func resourceIpAppProtocolPortTcpPortDisableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortDisableRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortDisableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortDisableRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortTcpPortDisableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortDisableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpAppProtocolPortTcpPortDisableInterface1033(d []interface{}) edpt.IpAppProtocolPortTcpPortDisableInterface1033 {

	count1 := len(d)
	var ret edpt.IpAppProtocolPortTcpPortDisableInterface1033
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Management = in["management"].(int)
		ret.VeCfg = getSliceIpAppProtocolPortTcpPortDisableInterfaceVeCfg1034(in["ve_cfg"].([]interface{}))
		ret.EthCfg = getSliceIpAppProtocolPortTcpPortDisableInterfaceEthCfg1035(in["eth_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortDisableInterfaceVeCfg1034(d []interface{}) []edpt.IpAppProtocolPortTcpPortDisableInterfaceVeCfg1034 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortDisableInterfaceVeCfg1034, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortDisableInterfaceVeCfg1034
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortDisableInterfaceEthCfg1035(d []interface{}) []edpt.IpAppProtocolPortTcpPortDisableInterfaceEthCfg1035 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortDisableInterfaceEthCfg1035, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortDisableInterfaceEthCfg1035
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortTcpPortDisable(d *schema.ResourceData) edpt.IpAppProtocolPortTcpPortDisable {
	var ret edpt.IpAppProtocolPortTcpPortDisable
	ret.Inst.Interface = getObjectIpAppProtocolPortTcpPortDisableInterface1033(d.Get("interface").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
