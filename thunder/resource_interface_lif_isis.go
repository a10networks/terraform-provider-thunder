package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_isis`: ISIS\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIsisCreate,
		UpdateContext: resourceInterfaceLifIsisUpdate,
		ReadContext:   resourceInterfaceLifIsisRead,
		DeleteContext: resourceInterfaceLifIsisDelete,

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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceLifIsisAuthentication(d []interface{}) edpt.InterfaceLifIsisAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceLifIsisAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceLifIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceLifIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceLifIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceLifIsisAuthenticationSendOnlyList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationSendOnlyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceLifIsisAuthenticationModeList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceLifIsisAuthenticationKeyChainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisAuthenticationKeyChainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsisBfdCfg(d []interface{}) edpt.InterfaceLifIsisBfdCfg {

	count1 := len(d)
	var ret edpt.InterfaceLifIsisBfdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceLifIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceLifIsisCsnpIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisCsnpIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloIntervalList(d []interface{}) []edpt.InterfaceLifIsisHelloIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceLifIsisHelloIntervalMinimalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloIntervalMinimalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceLifIsisHelloMultiplierList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisHelloMultiplierList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceLifIsisMeshGroup(d []interface{}) edpt.InterfaceLifIsisMeshGroup {

	count1 := len(d)
	var ret edpt.InterfaceLifIsisMeshGroup
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceLifIsisMetricList(d []interface{}) []edpt.InterfaceLifIsisMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisPasswordList(d []interface{}) []edpt.InterfaceLifIsisPasswordList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisPasswordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisPriorityList(d []interface{}) []edpt.InterfaceLifIsisPriorityList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisPriorityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLifIsisWideMetricList(d []interface{}) []edpt.InterfaceLifIsisWideMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIsisWideMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLifIsis(d *schema.ResourceData) edpt.InterfaceLifIsis {
	var ret edpt.InterfaceLifIsis
	ret.Inst.Authentication = getObjectInterfaceLifIsisAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.BfdCfg = getObjectInterfaceLifIsisBfdCfg(d.Get("bfd_cfg").([]interface{}))
	ret.Inst.CircuitType = d.Get("circuit_type").(string)
	ret.Inst.CsnpIntervalList = getSliceInterfaceLifIsisCsnpIntervalList(d.Get("csnp_interval_list").([]interface{}))
	ret.Inst.HelloIntervalList = getSliceInterfaceLifIsisHelloIntervalList(d.Get("hello_interval_list").([]interface{}))
	ret.Inst.HelloIntervalMinimalList = getSliceInterfaceLifIsisHelloIntervalMinimalList(d.Get("hello_interval_minimal_list").([]interface{}))
	ret.Inst.HelloMultiplierList = getSliceInterfaceLifIsisHelloMultiplierList(d.Get("hello_multiplier_list").([]interface{}))
	ret.Inst.LspInterval = d.Get("lsp_interval").(int)
	ret.Inst.MeshGroup = getObjectInterfaceLifIsisMeshGroup(d.Get("mesh_group").([]interface{}))
	ret.Inst.MetricList = getSliceInterfaceLifIsisMetricList(d.Get("metric_list").([]interface{}))
	ret.Inst.Network = d.Get("network").(string)
	ret.Inst.Padding = d.Get("padding").(int)
	ret.Inst.PasswordList = getSliceInterfaceLifIsisPasswordList(d.Get("password_list").([]interface{}))
	ret.Inst.PriorityList = getSliceInterfaceLifIsisPriorityList(d.Get("priority_list").([]interface{}))
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	//omit uuid
	ret.Inst.WideMetricList = getSliceInterfaceLifIsisWideMetricList(d.Get("wide_metric_list").([]interface{}))
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
