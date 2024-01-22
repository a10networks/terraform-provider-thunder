package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortTcpPortDisableInterface() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_tcp_port_disable_interface`: Disable a port from an interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortTcpPortDisableInterfaceCreate,
		UpdateContext: resourceIpAppProtocolPortTcpPortDisableInterfaceUpdate,
		ReadContext:   resourceIpAppProtocolPortTcpPortDisableInterfaceRead,
		DeleteContext: resourceIpAppProtocolPortTcpPortDisableInterfaceDelete,

		Schema: map[string]*schema.Schema{
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
			"port": {
				Type: schema.TypeString, Required: true, Description: "Port",
			},
		},
	}
}
func resourceIpAppProtocolPortTcpPortDisableInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableInterfaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisableInterface(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortDisableInterfaceRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortDisableInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableInterfaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisableInterface(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortDisableInterfaceRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortTcpPortDisableInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableInterfaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisableInterface(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortDisableInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortDisableInterfaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortDisableInterface(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAppProtocolPortTcpPortDisableInterfaceEthCfg(d []interface{}) []edpt.IpAppProtocolPortTcpPortDisableInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortDisableInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortDisableInterfaceEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortDisableInterfaceVeCfg(d []interface{}) []edpt.IpAppProtocolPortTcpPortDisableInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortDisableInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortDisableInterfaceVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortTcpPortDisableInterface(d *schema.ResourceData) edpt.IpAppProtocolPortTcpPortDisableInterface {
	var ret edpt.IpAppProtocolPortTcpPortDisableInterface
	ret.Inst.EthCfg = getSliceIpAppProtocolPortTcpPortDisableInterfaceEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	ret.Inst.VeCfg = getSliceIpAppProtocolPortTcpPortDisableInterfaceVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
