package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRuleActionGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rule_set_rule_action_group`: Configure action-group\n\n__PLACEHOLDER__",
		CreateContext: resourceRuleSetRuleActionGroupCreate,
		UpdateContext: resourceRuleSetRuleActionGroupUpdate,
		ReadContext:   resourceRuleSetRuleActionGroupRead,
		DeleteContext: resourceRuleSetRuleActionGroupDelete,

		Schema: map[string]*schema.Schema{
			"cgnv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply CGNv6 policy",
			},
			"cgnv6_ds_lite": {
				Type: schema.TypeString, Optional: true, Description: "'lsn-lid': Apply specified CGNv6 LSN LID;",
			},
			"cgnv6_ds_lite_lsn_lid": {
				Type: schema.TypeInt, Optional: true, Description: "LSN LID",
			},
			"cgnv6_lsn_lid": {
				Type: schema.TypeInt, Optional: true, Description: "LSN LID",
			},
			"cgnv6_policy": {
				Type: schema.TypeString, Optional: true, Description: "'lsn-lid': Apply specified CGNv6 LSN LID; 'fixed-nat': Apply CGNv6 Fixed NAT; 'ds-lite': Apply CGNv6 DS-Lite;",
			},
			"deny_fw_log": {
				Type: schema.TypeString, Optional: true, Description: "Logging template name",
			},
			"deny_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"deny_log_template_type": {
				Type: schema.TypeString, Optional: true, Description: "'fw-logging-template': Logging with specified fw template;",
			},
			"dscp_number": {
				Type: schema.TypeInt, Optional: true, Description: "DSCP Number",
			},
			"dscp_value": {
				Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110);",
			},
			"forward": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward packet",
			},
			"inspect_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DS-Lite tunnel inspection",
			},
			"ipsec": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply IPsec encapsulation",
			},
			"ipsec_group": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply IPsec Group encapsulation",
			},
			"listen_on_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Listen on port",
			},
			"logging_template_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permit_log_template_type": {
							Type: schema.TypeString, Optional: true, Description: "'fw-logging-template': Logging with specified fw template; 'cgnv6-logging-template': Logging with specified cgnv6 template; 'netflow-monitor': Logging with specified netflow/ipfix monitor;",
						},
						"permit_fw_log": {
							Type: schema.TypeString, Optional: true, Description: "Logging template name",
						},
						"permit_cgnv6_log": {
							Type: schema.TypeString, Optional: true, Description: "Logging template name",
						},
						"permit_netflow_log": {
							Type: schema.TypeString, Optional: true, Description: "Name of netflow monitor",
						},
					},
				},
			},
			"permit_limit_policy": {
				Type: schema.TypeInt, Optional: true, Description: "Limit policy Template",
			},
			"permit_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"permit_respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default:off)",
			},
			"reset_fw_log": {
				Type: schema.TypeString, Optional: true, Description: "Logging template name",
			},
			"reset_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"reset_log_template_type": {
				Type: schema.TypeString, Optional: true, Description: "'fw-logging-template': Logging with specified fw template;",
			},
			"reset_respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default:off)",
			},
			"set_dscp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "DSCP setting",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'permit': permit; 'deny': deny; 'reset': reset;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vpn_ipsec_group_name": {
				Type: schema.TypeString, Optional: true, Description: "VPN IPsec Group name",
			},
			"vpn_ipsec_name": {
				Type: schema.TypeString, Optional: true, Description: "VPN IPsec name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceRuleSetRuleActionGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleActionGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleActionGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRuleActionGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceRuleSetRuleActionGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleActionGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleActionGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRuleActionGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceRuleSetRuleActionGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleActionGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleActionGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRuleSetRuleActionGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleActionGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleActionGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRuleSetRuleActionGroupLoggingTemplateList(d []interface{}) []edpt.RuleSetRuleActionGroupLoggingTemplateList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleActionGroupLoggingTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleActionGroupLoggingTemplateList
		oi.PermitLogTemplateType = in["permit_log_template_type"].(string)
		oi.PermitFwLog = in["permit_fw_log"].(string)
		oi.PermitCgnv6Log = in["permit_cgnv6_log"].(string)
		oi.PermitNetflowLog = in["permit_netflow_log"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetRuleActionGroup(d *schema.ResourceData) edpt.RuleSetRuleActionGroup {
	var ret edpt.RuleSetRuleActionGroup
	ret.Inst.Cgnv6 = d.Get("cgnv6").(int)
	ret.Inst.Cgnv6DsLite = d.Get("cgnv6_ds_lite").(string)
	ret.Inst.Cgnv6DsLiteLsnLid = d.Get("cgnv6_ds_lite_lsn_lid").(int)
	ret.Inst.Cgnv6LsnLid = d.Get("cgnv6_lsn_lid").(int)
	ret.Inst.Cgnv6Policy = d.Get("cgnv6_policy").(string)
	ret.Inst.DenyFwLog = d.Get("deny_fw_log").(string)
	ret.Inst.DenyLog = d.Get("deny_log").(int)
	ret.Inst.DenyLogTemplateType = d.Get("deny_log_template_type").(string)
	ret.Inst.DscpNumber = d.Get("dscp_number").(int)
	ret.Inst.DscpValue = d.Get("dscp_value").(string)
	ret.Inst.Forward = d.Get("forward").(int)
	ret.Inst.InspectPayload = d.Get("inspect_payload").(int)
	ret.Inst.Ipsec = d.Get("ipsec").(int)
	ret.Inst.IpsecGroup = d.Get("ipsec_group").(int)
	ret.Inst.ListenOnPort = d.Get("listen_on_port").(int)
	ret.Inst.LoggingTemplateList = getSliceRuleSetRuleActionGroupLoggingTemplateList(d.Get("logging_template_list").([]interface{}))
	ret.Inst.PermitLimitPolicy = d.Get("permit_limit_policy").(int)
	ret.Inst.PermitLog = d.Get("permit_log").(int)
	ret.Inst.PermitRespondToUserMac = d.Get("permit_respond_to_user_mac").(int)
	ret.Inst.ResetFwLog = d.Get("reset_fw_log").(string)
	ret.Inst.ResetLog = d.Get("reset_log").(int)
	ret.Inst.ResetLogTemplateType = d.Get("reset_log_template_type").(string)
	ret.Inst.ResetRespondToUserMac = d.Get("reset_respond_to_user_mac").(int)
	ret.Inst.SetDscp = d.Get("set_dscp").(int)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.VpnIpsecGroupName = d.Get("vpn_ipsec_group_name").(string)
	ret.Inst.VpnIpsecName = d.Get("vpn_ipsec_name").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
