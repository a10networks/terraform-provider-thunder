package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntryIpProto() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_ip_proto`: DDOS IP protocol configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntryIpProtoCreate,
		UpdateContext: resourceDdosDstEntryIpProtoUpdate,
		ReadContext:   resourceDdosDstEntryIpProtoRead,
		DeleteContext: resourceDdosDstEntryIpProtoDelete,

		Schema: map[string]*schema.Schema{
			"deny": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
			},
			"esp_inspect": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_algorithm": {
							Type: schema.TypeString, Optional: true, Description: "'AUTH_NULL': No Integrity Check Value; 'HMAC-SHA-1-96': 96 bit Auth Algo; 'HMAC-SHA-256-96': 96 bit Auth Algo; 'HMAC-SHA-256-128': 128 bit Auth Algo; 'HMAC-SHA-384-192': 192 bit Auth Algo; 'HMAC-SHA-512-256': 256 bit Auth Algo; 'HMAC-MD5-96': 96 bit Auth Algo; 'MAC-RIPEMD-160-96': 96 bit Auth Algo;",
						},
						"encrypt_algorithm": {
							Type: schema.TypeString, Optional: true, Description: "'NULL': Null Encryption Algorithm;",
						},
						"mode": {
							Type: schema.TypeString, Optional: true, Description: "'transport': Transport mode;",
						},
					},
				},
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"glid_exceed_action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stateless_encap_action_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stateless_encap_action": {
										Type: schema.TypeString, Optional: true, Description: "'stateless-tunnel-encap': Encapsulate all packets; 'stateless-tunnel-encap-scrubbed': Encapsulate all packets and allow packets to go through other DDoS checks before sent (conn-limit exceeded packet can not be scrubbed, it will default to stateless-tunnel-encap);",
									},
									"encap_template": {
										Type: schema.TypeString, Optional: true, Description: "Apply legacy encap template for encap action",
									},
								},
							},
						},
					},
				},
			},
			"ip_filtering_policy": {
				Type: schema.TypeString, Optional: true, Description: "Configure IP Filter",
			},
			"ip_filtering_policy_oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Protocol Number",
			},
			"set_counter_base_val": {
				Type: schema.TypeInt, Optional: true, Description: "Set T2 counter value of current context to specified value",
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"other": {
							Type: schema.TypeString, Optional: true, Description: "DDOS other template",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntryIpProtoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryIpProtoCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryIpProto(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryIpProtoRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntryIpProtoUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryIpProtoUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryIpProto(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntryIpProtoRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntryIpProtoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryIpProtoDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryIpProto(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntryIpProtoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntryIpProtoRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntryIpProto(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstEntryIpProtoEspInspect(d []interface{}) edpt.DdosDstEntryIpProtoEspInspect {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoEspInspect
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.AuthAlgorithm = in["auth_algorithm"].(string)
		ret.EncryptAlgorithm = in["encrypt_algorithm"].(string)
		ret.Mode = in["mode"].(string)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoGlidExceedAction(d []interface{}) edpt.DdosDstEntryIpProtoGlidExceedAction {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoGlidExceedAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapActionCfg = getObjectDdosDstEntryIpProtoGlidExceedActionStatelessEncapActionCfg(in["stateless_encap_action_cfg"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstEntryIpProtoGlidExceedActionStatelessEncapActionCfg(d []interface{}) edpt.DdosDstEntryIpProtoGlidExceedActionStatelessEncapActionCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoGlidExceedActionStatelessEncapActionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StatelessEncapAction = in["stateless_encap_action"].(string)
		ret.EncapTemplate = in["encap_template"].(string)
	}
	return ret
}

func getObjectDdosDstEntryIpProtoIpFilteringPolicyOper157(d []interface{}) edpt.DdosDstEntryIpProtoIpFilteringPolicyOper157 {

	var ret edpt.DdosDstEntryIpProtoIpFilteringPolicyOper157
	return ret
}

func getObjectDdosDstEntryIpProtoTemplate(d []interface{}) edpt.DdosDstEntryIpProtoTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntryIpProtoTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Other = in["other"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntryIpProto(d *schema.ResourceData) edpt.DdosDstEntryIpProto {
	var ret edpt.DdosDstEntryIpProto
	ret.Inst.Deny = d.Get("deny").(int)
	ret.Inst.EspInspect = getObjectDdosDstEntryIpProtoEspInspect(d.Get("esp_inspect").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.GlidExceedAction = getObjectDdosDstEntryIpProtoGlidExceedAction(d.Get("glid_exceed_action").([]interface{}))
	ret.Inst.IpFilteringPolicy = d.Get("ip_filtering_policy").(string)
	ret.Inst.IpFilteringPolicyOper = getObjectDdosDstEntryIpProtoIpFilteringPolicyOper157(d.Get("ip_filtering_policy_oper").([]interface{}))
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.SetCounterBaseVal = d.Get("set_counter_base_val").(int)
	ret.Inst.Template = getObjectDdosDstEntryIpProtoTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
