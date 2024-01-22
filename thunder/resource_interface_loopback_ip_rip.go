package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIpRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_ip_rip`: RIP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIpRipCreate,
		UpdateContext: resourceInterfaceLoopbackIpRipUpdate,
		ReadContext:   resourceInterfaceLoopbackIpRipRead,
		DeleteContext: resourceInterfaceLoopbackIpRipDelete,

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
							Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-2': RIP version 1 & 2;",
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
							Type: schema.TypeString, Optional: true, Description: "'1': RIP version 1; '2': RIP version 2; '1-compatible': RIPv1-compatible; '1-2': RIP version 1 & 2;",
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
func resourceInterfaceLoopbackIpRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpRipRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIpRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpRipRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIpRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIpRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLoopbackIpRipAuthentication(d []interface{}) edpt.InterfaceLoopbackIpRipAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceLoopbackIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLoopbackIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLoopbackIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationStr(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationStr {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationStr
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationMode(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationMode {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationMode
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceLoopbackIpRipAuthenticationKeyChain {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipAuthenticationKeyChain
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipReceiveCfg(d []interface{}) edpt.InterfaceLoopbackIpRipReceiveCfg {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipReceiveCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSendCfg(d []interface{}) edpt.InterfaceLoopbackIpRipSendCfg {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipSendCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLoopbackIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceLoopbackIpRipSplitHorizonCfg {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIpRipSplitHorizonCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceLoopbackIpRip(d *schema.ResourceData) edpt.InterfaceLoopbackIpRip {
	var ret edpt.InterfaceLoopbackIpRip
	ret.Inst.Authentication = getObjectInterfaceLoopbackIpRipAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.ReceiveCfg = getObjectInterfaceLoopbackIpRipReceiveCfg(d.Get("receive_cfg").([]interface{}))
	ret.Inst.ReceivePacket = d.Get("receive_packet").(int)
	ret.Inst.SendCfg = getObjectInterfaceLoopbackIpRipSendCfg(d.Get("send_cfg").([]interface{}))
	ret.Inst.SendPacket = d.Get("send_packet").(int)
	ret.Inst.SplitHorizonCfg = getObjectInterfaceLoopbackIpRipSplitHorizonCfg(d.Get("split_horizon_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
