package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortTcpPortAddAppNameInterface() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_tcp_port_add_app_name_interface`: Bind a port to an interface\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortTcpPortAddAppNameInterfaceCreate,
		UpdateContext: resourceIpAppProtocolPortTcpPortAddAppNameInterfaceUpdate,
		ReadContext:   resourceIpAppProtocolPortTcpPortAddAppNameInterfaceRead,
		DeleteContext: resourceIpAppProtocolPortTcpPortAddAppNameInterfaceDelete,

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
func resourceIpAppProtocolPortTcpPortAddAppNameInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameInterfaceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppNameInterface(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortAddAppNameInterfaceRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortAddAppNameInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameInterfaceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppNameInterface(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortTcpPortAddAppNameInterfaceRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortTcpPortAddAppNameInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameInterfaceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppNameInterface(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortTcpPortAddAppNameInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortTcpPortAddAppNameInterfaceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortTcpPortAddAppNameInterface(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg(d []interface{}) []edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortTcpPortAddAppNameInterface(d *schema.ResourceData) edpt.IpAppProtocolPortTcpPortAddAppNameInterface {
	var ret edpt.IpAppProtocolPortTcpPortAddAppNameInterface
	ret.Inst.EthCfg = getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg(d.Get("eth_cfg").([]interface{}))
	ret.Inst.Management = d.Get("management").(int)
	//omit uuid
	ret.Inst.VeCfg = getSliceIpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg(d.Get("ve_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
