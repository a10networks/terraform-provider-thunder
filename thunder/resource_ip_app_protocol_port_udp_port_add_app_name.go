package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortUdpPortAddAppName() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_udp_port_add_app_name`: Application protocol name\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortUdpPortAddAppNameCreate,
		UpdateContext: resourceIpAppProtocolPortUdpPortAddAppNameUpdate,
		ReadContext:   resourceIpAppProtocolPortUdpPortAddAppNameRead,
		DeleteContext: resourceIpAppProtocolPortUdpPortAddAppNameDelete,

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
func resourceIpAppProtocolPortUdpPortAddAppNameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppName(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortAddAppNameRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortAddAppNameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppName(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortAddAppNameRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortUdpPortAddAppNameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppName(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortAddAppNameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortAddAppNameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortAddAppName(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpAppProtocolPortUdpPortAddAppNameInterface1036(d []interface{}) edpt.IpAppProtocolPortUdpPortAddAppNameInterface1036 {

	count1 := len(d)
	var ret edpt.IpAppProtocolPortUdpPortAddAppNameInterface1036
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Management = in["management"].(int)
		ret.VeCfg = getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037(in["ve_cfg"].([]interface{}))
		ret.EthCfg = getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038(in["eth_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038(d []interface{}) []edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortUdpPortAddAppName(d *schema.ResourceData) edpt.IpAppProtocolPortUdpPortAddAppName {
	var ret edpt.IpAppProtocolPortUdpPortAddAppName
	ret.Inst.Interface = getObjectIpAppProtocolPortUdpPortAddAppNameInterface1036(d.Get("interface").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Port = d.Get("port").(string)
	return ret
}
