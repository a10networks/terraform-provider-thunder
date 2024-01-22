package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceVeIpRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ve_ip_rip`: RIP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceVeIpRipCreate,
		UpdateContext: resourceInterfaceVeIpRipUpdate,
		ReadContext:   resourceInterfaceVeIpRipRead,
		DeleteContext: resourceInterfaceVeIpRipDelete,

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
func resourceInterfaceVeIpRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeIpRipRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceVeIpRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceVeIpRipRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceVeIpRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceVeIpRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceVeIpRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceVeIpRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceVeIpRipAuthentication(d []interface{}) edpt.InterfaceVeIpRipAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceVeIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceVeIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceVeIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationStr(d []interface{}) edpt.InterfaceVeIpRipAuthenticationStr {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthenticationStr
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationMode(d []interface{}) edpt.InterfaceVeIpRipAuthenticationMode {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthenticationMode
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceVeIpRipAuthenticationKeyChain {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipAuthenticationKeyChain
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipReceiveCfg(d []interface{}) edpt.InterfaceVeIpRipReceiveCfg {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipReceiveCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipSendCfg(d []interface{}) edpt.InterfaceVeIpRipSendCfg {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipSendCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceVeIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceVeIpRipSplitHorizonCfg {

	count1 := len(d)
	var ret edpt.InterfaceVeIpRipSplitHorizonCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceVeIpRip(d *schema.ResourceData) edpt.InterfaceVeIpRip {
	var ret edpt.InterfaceVeIpRip
	ret.Inst.Authentication = getObjectInterfaceVeIpRipAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.ReceiveCfg = getObjectInterfaceVeIpRipReceiveCfg(d.Get("receive_cfg").([]interface{}))
	ret.Inst.ReceivePacket = d.Get("receive_packet").(int)
	ret.Inst.SendCfg = getObjectInterfaceVeIpRipSendCfg(d.Get("send_cfg").([]interface{}))
	ret.Inst.SendPacket = d.Get("send_packet").(int)
	ret.Inst.SplitHorizonCfg = getObjectInterfaceVeIpRipSplitHorizonCfg(d.Get("split_horizon_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
