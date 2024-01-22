package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_isis`: ISIS\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIsisCreate,
		UpdateContext: resourceInterfaceLoopbackIsisUpdate,
		ReadContext:   resourceInterfaceLoopbackIsisRead,
		DeleteContext: resourceInterfaceLoopbackIsisDelete,

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
func resourceInterfaceLoopbackIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLoopbackIsisAuthentication(d []interface{}) edpt.InterfaceLoopbackIsisAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsisAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceLoopbackIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceLoopbackIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceLoopbackIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationModeList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceLoopbackIsisAuthenticationKeyChainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisAuthenticationKeyChainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsisBfdCfg(d []interface{}) edpt.InterfaceLoopbackIsisBfdCfg {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsisBfdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceLoopbackIsisCsnpIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisCsnpIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloIntervalList(d []interface{}) []edpt.InterfaceLoopbackIsisHelloIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceLoopbackIsisHelloIntervalMinimalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloIntervalMinimalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceLoopbackIsisHelloMultiplierList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisHelloMultiplierList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLoopbackIsisMeshGroup(d []interface{}) edpt.InterfaceLoopbackIsisMeshGroup {

	count1 := len(d)
	var ret edpt.InterfaceLoopbackIsisMeshGroup
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisMetricList(d []interface{}) []edpt.InterfaceLoopbackIsisMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisPasswordList(d []interface{}) []edpt.InterfaceLoopbackIsisPasswordList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisPasswordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisPriorityList(d []interface{}) []edpt.InterfaceLoopbackIsisPriorityList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisPriorityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIsisWideMetricList(d []interface{}) []edpt.InterfaceLoopbackIsisWideMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIsisWideMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLoopbackIsis(d *schema.ResourceData) edpt.InterfaceLoopbackIsis {
	var ret edpt.InterfaceLoopbackIsis
	ret.Inst.Authentication = getObjectInterfaceLoopbackIsisAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.BfdCfg = getObjectInterfaceLoopbackIsisBfdCfg(d.Get("bfd_cfg").([]interface{}))
	ret.Inst.CircuitType = d.Get("circuit_type").(string)
	ret.Inst.CsnpIntervalList = getSliceInterfaceLoopbackIsisCsnpIntervalList(d.Get("csnp_interval_list").([]interface{}))
	ret.Inst.HelloIntervalList = getSliceInterfaceLoopbackIsisHelloIntervalList(d.Get("hello_interval_list").([]interface{}))
	ret.Inst.HelloIntervalMinimalList = getSliceInterfaceLoopbackIsisHelloIntervalMinimalList(d.Get("hello_interval_minimal_list").([]interface{}))
	ret.Inst.HelloMultiplierList = getSliceInterfaceLoopbackIsisHelloMultiplierList(d.Get("hello_multiplier_list").([]interface{}))
	ret.Inst.LspInterval = d.Get("lsp_interval").(int)
	ret.Inst.MeshGroup = getObjectInterfaceLoopbackIsisMeshGroup(d.Get("mesh_group").([]interface{}))
	ret.Inst.MetricList = getSliceInterfaceLoopbackIsisMetricList(d.Get("metric_list").([]interface{}))
	ret.Inst.Padding = d.Get("padding").(int)
	ret.Inst.PasswordList = getSliceInterfaceLoopbackIsisPasswordList(d.Get("password_list").([]interface{}))
	ret.Inst.PriorityList = getSliceInterfaceLoopbackIsisPriorityList(d.Get("priority_list").([]interface{}))
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	//omit uuid
	ret.Inst.WideMetricList = getSliceInterfaceLoopbackIsisWideMetricList(d.Get("wide_metric_list").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
