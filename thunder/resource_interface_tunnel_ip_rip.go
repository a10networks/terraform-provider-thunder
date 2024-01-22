package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTunnelIpRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_tunnel_ip_rip`: RIP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTunnelIpRipCreate,
		UpdateContext: resourceInterfaceTunnelIpRipUpdate,
		ReadContext:   resourceInterfaceTunnelIpRipRead,
		DeleteContext: resourceInterfaceTunnelIpRipDelete,

		Schema: map[string]*schema.Schema{
			"authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"string": {
										Type: schema.TypeString, Optional: true, Description: "The RIP authentication string",
									},
								},
							},
						},
						"mode": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type: schema.TypeString, Optional: true, Default: "text", Description: "'md5': Keyed message digest; 'text': Clear text authentication;",
									},
								},
							},
						},
						"key_chain": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_chain": {
										Type: schema.TypeString, Optional: true, Description: "Authentication key-chain (Name of key-chain)",
									},
								},
							},
						},
					},
				},
			},
			"receive_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"receive": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement reception",
						},
						"version": {
							Type: schema.TypeString, Optional: true, Description: "'2': RIP version 2;",
						},
					},
				},
			},
			"receive_packet": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable receiving packet through the specified interface",
			},
			"send_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Advertisement transmission",
						},
						"version": {
							Type: schema.TypeString, Optional: true, Description: "'2': RIP version 2;",
						},
					},
				},
			},
			"send_packet": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Enable sending packets through the specified interface",
			},
			"split_horizon_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Default: "poisoned", Description: "'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTunnelIpRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpRipRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTunnelIpRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTunnelIpRipRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTunnelIpRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTunnelIpRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTunnelIpRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTunnelIpRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTunnelIpRipAuthentication(d []interface{}) edpt.InterfaceTunnelIpRipAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceTunnelIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceTunnelIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceTunnelIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationStr(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationStr {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationStr
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationMode(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationMode {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationMode
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceTunnelIpRipAuthenticationKeyChain {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipAuthenticationKeyChain
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipReceiveCfg(d []interface{}) edpt.InterfaceTunnelIpRipReceiveCfg {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipReceiveCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipSendCfg(d []interface{}) edpt.InterfaceTunnelIpRipSendCfg {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipSendCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTunnelIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceTunnelIpRipSplitHorizonCfg {

	count1 := len(d)
	var ret edpt.InterfaceTunnelIpRipSplitHorizonCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceTunnelIpRip(d *schema.ResourceData) edpt.InterfaceTunnelIpRip {
	var ret edpt.InterfaceTunnelIpRip
	ret.Inst.Authentication = getObjectInterfaceTunnelIpRipAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.ReceiveCfg = getObjectInterfaceTunnelIpRipReceiveCfg(d.Get("receive_cfg").([]interface{}))
	ret.Inst.ReceivePacket = d.Get("receive_packet").(int)
	ret.Inst.SendCfg = getObjectInterfaceTunnelIpRipSendCfg(d.Get("send_cfg").([]interface{}))
	ret.Inst.SendPacket = d.Get("send_packet").(int)
	ret.Inst.SplitHorizonCfg = getObjectInterfaceTunnelIpRipSplitHorizonCfg(d.Get("split_horizon_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
