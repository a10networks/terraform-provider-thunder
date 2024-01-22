package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortUdpPortAddAppNameInterface() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_udp_port_add_app_name_interface`: Bind a port to an interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortUdpPortAddAppNameInterfaceCreate,
		UpdateContext: resourceIpAppProtocolPortUdpPortAddAppNameInterfaceUpdate,
		ReadContext:   resourceIpAppProtocolPortUdpPortAddAppNameInterfaceRead,
		DeleteContext: resourceIpAppProtocolPortUdpPortAddAppNameInterfaceDelete,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"port": {
				Type: schema.TypeString, Required: true, Description: "Port",
			},
		},
	}
}
func resourceIpAppProtocolPortUdpPortAddAppNameInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameInterfaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppNameInterface(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortAddAppNameInterfaceRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortAddAppNameInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameInterfaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppNameInterface(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortAddAppNameInterfaceRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortUdpPortAddAppNameInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameInterfaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppNameInterface(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortAddAppNameInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameInterfaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppNameInterface(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortUdpPortAddAppNameInterface(d *schema.ResourceData) edpt.IpAppProtocolPortUdpPortAddAppNameInterface {
	var ret edpt.IpAppProtocolPortUdpPortAddAppNameInterface
	ret.Inst.EthCfg = getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	ret.Inst.VeCfg = getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
