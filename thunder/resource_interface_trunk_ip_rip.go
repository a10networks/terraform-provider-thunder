package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ip_rip`: RIP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpRipCreate,
		UpdateContext: resourceInterfaceTrunkIpRipUpdate,
		ReadContext:   resourceInterfaceTrunkIpRipRead,
		DeleteContext: resourceInterfaceTrunkIpRipDelete,

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
func resourceInterfaceTrunkIpRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpRipRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpRipRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTrunkIpRipAuthentication(d []interface{}) edpt.InterfaceTrunkIpRipAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceTrunkIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceTrunkIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceTrunkIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationStr(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationStr {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationStr
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationMode(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationMode {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationMode
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceTrunkIpRipAuthenticationKeyChain {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipAuthenticationKeyChain
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipReceiveCfg(d []interface{}) edpt.InterfaceTrunkIpRipReceiveCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipReceiveCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipSendCfg(d []interface{}) edpt.InterfaceTrunkIpRipSendCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipSendCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceTrunkIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceTrunkIpRipSplitHorizonCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpRipSplitHorizonCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceTrunkIpRip(d *schema.ResourceData) edpt.InterfaceTrunkIpRip {
	var ret edpt.InterfaceTrunkIpRip
	ret.Inst.Authentication = getObjectInterfaceTrunkIpRipAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.ReceiveCfg = getObjectInterfaceTrunkIpRipReceiveCfg(d.Get("receive_cfg").([]interface{}))
	ret.Inst.ReceivePacket = d.Get("receive_packet").(int)
	ret.Inst.SendCfg = getObjectInterfaceTrunkIpRipSendCfg(d.Get("send_cfg").([]interface{}))
	ret.Inst.SendPacket = d.Get("send_packet").(int)
	ret.Inst.SplitHorizonCfg = getObjectInterfaceTrunkIpRipSplitHorizonCfg(d.Get("split_horizon_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
