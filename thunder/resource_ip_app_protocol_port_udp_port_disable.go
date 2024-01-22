package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpAppProtocolPortUdpPortDisable() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_app_protocol_port_udp_port_disable`: Disable application port\n\n__PLACEHOLDER__",
		CreateContext: resourceIpAppProtocolPortUdpPortDisableCreate,
		UpdateContext: resourceIpAppProtocolPortUdpPortDisableUpdate,
		ReadContext:   resourceIpAppProtocolPortUdpPortDisableRead,
		DeleteContext: resourceIpAppProtocolPortUdpPortDisableDelete,

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
func resourceIpAppProtocolPortUdpPortDisableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisable(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortDisableRead(ctx, d, meta)
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortDisableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisable(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpAppProtocolPortUdpPortDisableRead(ctx, d, meta)
	}
	return diags
}
func resourceIpAppProtocolPortUdpPortDisableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisable(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpAppProtocolPortUdpPortDisableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpAppProtocolPortUdpPortDisableRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpAppProtocolPortUdpPortDisable(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpAppProtocolPortUdpPortDisableInterface1039(d []interface{}) edpt.IpAppProtocolPortUdpPortDisableInterface1039 {

	count1 := len(d)
	var ret edpt.IpAppProtocolPortUdpPortDisableInterface1039
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Management = in["management"].(int)
		ret.VeCfg = getSliceIpAppProtocolPortUdpPortDisableInterfaceVeCfg1040(in["ve_cfg"].([]interface{}))
		ret.EthCfg = getSliceIpAppProtocolPortUdpPortDisableInterfaceEthCfg1041(in["eth_cfg"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortDisableInterfaceVeCfg1040(d []interface{}) []edpt.IpAppProtocolPortUdpPortDisableInterfaceVeCfg1040 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortDisableInterfaceVeCfg1040, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortDisableInterfaceVeCfg1040
		oi.VeStart = in["ve_start"].(int)
		oi.VeEnd = in["ve_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpAppProtocolPortUdpPortDisableInterfaceEthCfg1041(d []interface{}) []edpt.IpAppProtocolPortUdpPortDisableInterfaceEthCfg1041 {

	count1 := len(d)
	ret := make([]edpt.IpAppProtocolPortUdpPortDisableInterfaceEthCfg1041, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpAppProtocolPortUdpPortDisableInterfaceEthCfg1041
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpAppProtocolPortUdpPortDisable(d *schema.ResourceData) edpt.IpAppProtocolPortUdpPortDisable {
	var ret edpt.IpAppProtocolPortUdpPortDisable
	ret.Inst.Interface = getObjectIpAppProtocolPortUdpPortDisableInterface1039(d.Get("interface").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
