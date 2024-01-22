package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortUdpPortDisableInterface() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_udp_port_disable_interface`: Disable a port from an interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortUdpPortDisableInterfaceCreate,
		UpdateContext: resourceIpAppProtocolPortUdpPortDisableInterfaceUpdate,
		ReadContext:   resourceIpAppProtocolPortUdpPortDisableInterfaceRead,
		DeleteContext: resourceIpAppProtocolPortUdpPortDisableInterfaceDelete,

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
func resourceIpAppProtocolPortUdpPortDisableInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableInterfaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisableInterface(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortDisableInterfaceRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortDisableInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableInterfaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisableInterface(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortDisableInterfaceRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortUdpPortDisableInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableInterfaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisableInterface(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortDisableInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableInterfaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisableInterface(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAppProtocolPortUdpPortDisableInterfaceEthCfg(d []interface{}) []edpt.IpAppProtocolPortUdpPortDisableInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortDisableInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortDisableInterfaceEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortDisableInterfaceVeCfg(d []interface{}) []edpt.IpAppProtocolPortUdpPortDisableInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortDisableInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortDisableInterfaceVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortUdpPortDisableInterface(d *schema.ResourceData) edpt.IpAppProtocolPortUdpPortDisableInterface {
	var ret edpt.IpAppProtocolPortUdpPortDisableInterface
	ret.Inst.EthCfg = getSliceIpAppProtocolPortUdpPortDisableInterfaceEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	ret.Inst.VeCfg = getSliceIpAppProtocolPortUdpPortDisableInterfaceVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
