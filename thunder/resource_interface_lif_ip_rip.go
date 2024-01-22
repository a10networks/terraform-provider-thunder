package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpRip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ip_rip`: RIP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpRipCreate,
		UpdateContext: resourceInterfaceLifIpRipUpdate,
		ReadContext:   resourceInterfaceLifIpRipRead,
		DeleteContext: resourceInterfaceLifIpRipDelete,

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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpRipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpRipRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpRipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpRipRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpRipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpRipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpRipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpRip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLifIpRipAuthentication(d []interface{}) edpt.InterfaceLifIpRipAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Str = getObjectInterfaceLifIpRipAuthenticationStr(in["str"].([]interface{}))
		ret.Mode = getObjectInterfaceLifIpRipAuthenticationMode(in["mode"].([]interface{}))
		ret.KeyChain = getObjectInterfaceLifIpRipAuthenticationKeyChain(in["key_chain"].([]interface{}))
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationStr(d []interface{}) edpt.InterfaceLifIpRipAuthenticationStr {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationStr
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.String = in["string"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationMode(d []interface{}) edpt.InterfaceLifIpRipAuthenticationMode {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationMode
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipAuthenticationKeyChain(d []interface{}) edpt.InterfaceLifIpRipAuthenticationKeyChain {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipAuthenticationKeyChain
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyChain = in["key_chain"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipReceiveCfg(d []interface{}) edpt.InterfaceLifIpRipReceiveCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipReceiveCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Receive = in["receive"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSendCfg(d []interface{}) edpt.InterfaceLifIpRipSendCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipSendCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Send = in["send"].(int)
		ret.Version = in["version"].(string)
	}
	return ret
}

func getObjectInterfaceLifIpRipSplitHorizonCfg(d []interface{}) edpt.InterfaceLifIpRipSplitHorizonCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIpRipSplitHorizonCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceLifIpRip(d *schema.ResourceData) edpt.InterfaceLifIpRip {
	var ret edpt.InterfaceLifIpRip
	ret.Inst.Authentication = getObjectInterfaceLifIpRipAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.ReceiveCfg = getObjectInterfaceLifIpRipReceiveCfg(d.Get("receive_cfg").([]interface{}))
	ret.Inst.ReceivePacket = d.Get("receive_packet").(int)
	ret.Inst.SendCfg = getObjectInterfaceLifIpRipSendCfg(d.Get("send_cfg").([]interface{}))
	ret.Inst.SendPacket = d.Get("send_packet").(int)
	ret.Inst.SplitHorizonCfg = getObjectInterfaceLifIpRipSplitHorizonCfg(d.Get("split_horizon_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
