package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_isis`: ISIS\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIsisCreate,
		UpdateContext: resourceInterfaceTrunkIsisUpdate,
		ReadContext:   resourceInterfaceTrunkIsisRead,
		DeleteContext: resourceInterfaceTrunkIsisDelete,

		Schema: map[string]*schema.Schema{
			"authentication": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"send_only_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"send_only": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Authentication send-only",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication send-only for level-1 PDUs; 'level-2': Specify authentication send-only for level-2 PDUs;",
									},
								},
							},
						},
						"mode_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type: schema.TypeString, Optional: true, Description: "'md5': Keyed message digest;",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication mode for level-1 PDUs; 'level-2': Specify authentication mode for level-2 PDUs;",
									},
								},
							},
						},
						"key_chain_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_chain": {
										Type: schema.TypeString, Optional: true, Description: "Authentication key-chain (Name of key-chain)",
									},
									"level": {
										Type: schema.TypeString, Optional: true, Description: "'level-1': Specify authentication key-chain for level-1 PDUs; 'level-2': Specify authentication key-chain for level-2 PDUs;",
									},
								},
							},
						},
					},
				},
			},
			"bfd_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
						},
						"disable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
						},
					},
				},
			},
			"circuit_type": {
				Type: schema.TypeString, Optional: true, Default: "level-1-2", Description: "'level-1': Level-1 only adjacencies are formed; 'level-1-2': Level-1-2 adjacencies are formed; 'level-2-only': Level-2 only adjacencies are formed;",
			},
			"csnp_interval_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"csnp_interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Set CSNP interval in seconds (CSNP interval value)",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Speficy interval for level-1 CSNPs; 'level-2': Specify interval for level-2 CSNPs;",
						},
					},
				},
			},
			"hello_interval_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Set Hello interval in seconds (Hello interval value)",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;",
						},
					},
				},
			},
			"hello_interval_minimal_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_interval_minimal": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set Hello holdtime 1 second, interval depends on multiplier",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello-interval for level-1 IIHs; 'level-2': Specify hello-interval for level-2 IIHs;",
						},
					},
				},
			},
			"hello_multiplier_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_multiplier": {
							Type: schema.TypeInt, Optional: true, Default: 3, Description: "Set multiplier for Hello holding time (Hello multiplier value)",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Specify hello multiplier for level-1 IIHs; 'level-2': Specify hello multiplier for level-2 IIHs;",
						},
					},
				},
			},
			"lsp_interval": {
				Type: schema.TypeInt, Optional: true, Default: 33, Description: "Set LSP transmission interval (LSP transmission interval (milliseconds))",
			},
			"mesh_group": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeInt, Optional: true, Description: "Mesh group number",
						},
						"blocked": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Block LSPs on this interface",
						},
					},
				},
			},
			"metric_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the metric for interface (Default metric)",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;",
						},
					},
				},
			},
			"network": {
				Type: schema.TypeString, Optional: true, Description: "'broadcast': Specify IS-IS broadcast multi-access network; 'point-to-point': Specify IS-IS point-to-point network;",
			},
			"padding": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Add padding to IS-IS hello packets",
			},
			"password_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type: schema.TypeString, Optional: true, Description: "Configure the authentication password for interface",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Specify password for level-1 PDUs; 'level-2': Specify password for level-2 PDUs;",
						},
					},
				},
			},
			"priority_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 64, Description: "Set priority for Designated Router election (Priority value)",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Specify priority for level-1 routing; 'level-2': Specify priority for level-2 routing;",
						},
					},
				},
			},
			"retransmit_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set per-LSP retransmission interval (Interval between retransmissions of the same LSP (seconds))",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"wide_metric_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"wide_metric": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Configure the wide metric for interface",
						},
						"level": {
							Type: schema.TypeString, Optional: true, Description: "'level-1': Apply metric to level-1 links; 'level-2': Apply metric to level-2 links;",
						},
					},
				},
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTrunkIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTrunkIsisAuthentication(d []interface{}) edpt.InterfaceTrunkIsisAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsisAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceTrunkIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceTrunkIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceTrunkIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceTrunkIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceTrunkIsisAuthenticationSendOnlyList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisAuthenticationSendOnlyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceTrunkIsisAuthenticationModeList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisAuthenticationModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceTrunkIsisAuthenticationKeyChainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisAuthenticationKeyChainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIsisBfdCfg(d []interface{}) edpt.InterfaceTrunkIsisBfdCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsisBfdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceTrunkIsisCsnpIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisCsnpIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisHelloIntervalList(d []interface{}) []edpt.InterfaceTrunkIsisHelloIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisHelloIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceTrunkIsisHelloIntervalMinimalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisHelloIntervalMinimalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceTrunkIsisHelloMultiplierList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisHelloMultiplierList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceTrunkIsisMeshGroup(d []interface{}) edpt.InterfaceTrunkIsisMeshGroup {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIsisMeshGroup
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceTrunkIsisMetricList(d []interface{}) []edpt.InterfaceTrunkIsisMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisPasswordList(d []interface{}) []edpt.InterfaceTrunkIsisPasswordList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisPasswordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisPriorityList(d []interface{}) []edpt.InterfaceTrunkIsisPriorityList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisPriorityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceTrunkIsisWideMetricList(d []interface{}) []edpt.InterfaceTrunkIsisWideMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIsisWideMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTrunkIsis(d *schema.ResourceData) edpt.InterfaceTrunkIsis {
	var ret edpt.InterfaceTrunkIsis
	ret.Inst.Authentication = getObjectInterfaceTrunkIsisAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.BfdCfg = getObjectInterfaceTrunkIsisBfdCfg(d.Get("bfd_cfg").([]interface{}))
	ret.Inst.CircuitType = d.Get("circuit_type").(string)
	ret.Inst.CsnpIntervalList = getSliceInterfaceTrunkIsisCsnpIntervalList(d.Get("csnp_interval_list").([]interface{}))
	ret.Inst.HelloIntervalList = getSliceInterfaceTrunkIsisHelloIntervalList(d.Get("hello_interval_list").([]interface{}))
	ret.Inst.HelloIntervalMinimalList = getSliceInterfaceTrunkIsisHelloIntervalMinimalList(d.Get("hello_interval_minimal_list").([]interface{}))
	ret.Inst.HelloMultiplierList = getSliceInterfaceTrunkIsisHelloMultiplierList(d.Get("hello_multiplier_list").([]interface{}))
	ret.Inst.LspInterval = d.Get("lsp_interval").(int)
	ret.Inst.MeshGroup = getObjectInterfaceTrunkIsisMeshGroup(d.Get("mesh_group").([]interface{}))
	ret.Inst.MetricList = getSliceInterfaceTrunkIsisMetricList(d.Get("metric_list").([]interface{}))
	ret.Inst.Network = d.Get("network").(string)
	ret.Inst.Padding = d.Get("padding").(int)
	ret.Inst.PasswordList = getSliceInterfaceTrunkIsisPasswordList(d.Get("password_list").([]interface{}))
	ret.Inst.PriorityList = getSliceInterfaceTrunkIsisPriorityList(d.Get("priority_list").([]interface{}))
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	//omit uuid
	ret.Inst.WideMetricList = getSliceInterfaceTrunkIsisWideMetricList(d.Get("wide_metric_list").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
