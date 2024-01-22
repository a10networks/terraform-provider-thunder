package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceEthernetIsis() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_ethernet_isis`: ISIS\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceEthernetIsisCreate,
		UpdateContext: resourceInterfaceEthernetIsisUpdate,
		ReadContext:   resourceInterfaceEthernetIsisRead,
		DeleteContext: resourceInterfaceEthernetIsisDelete,

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
func resourceInterfaceEthernetIsisCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIsisCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIsis(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIsisRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceEthernetIsisUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIsisUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIsis(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceEthernetIsisRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceEthernetIsisDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIsisDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIsis(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceEthernetIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceEthernetIsisRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceEthernetIsis(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceEthernetIsisAuthentication(d []interface{}) edpt.InterfaceEthernetIsisAuthentication {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsisAuthentication
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SendOnlyList = getSliceInterfaceEthernetIsisAuthenticationSendOnlyList(in["send_only_list"].([]interface{}))
		ret.ModeList = getSliceInterfaceEthernetIsisAuthenticationModeList(in["mode_list"].([]interface{}))
		ret.KeyChainList = getSliceInterfaceEthernetIsisAuthenticationKeyChainList(in["key_chain_list"].([]interface{}))
	}
	return ret
}

func getSliceInterfaceEthernetIsisAuthenticationSendOnlyList(d []interface{}) []edpt.InterfaceEthernetIsisAuthenticationSendOnlyList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisAuthenticationSendOnlyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisAuthenticationSendOnlyList
		oi.SendOnly = in["send_only"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisAuthenticationModeList(d []interface{}) []edpt.InterfaceEthernetIsisAuthenticationModeList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisAuthenticationModeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisAuthenticationModeList
		oi.Mode = in["mode"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisAuthenticationKeyChainList(d []interface{}) []edpt.InterfaceEthernetIsisAuthenticationKeyChainList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisAuthenticationKeyChainList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisAuthenticationKeyChainList
		oi.KeyChain = in["key_chain"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIsisBfdCfg(d []interface{}) edpt.InterfaceEthernetIsisBfdCfg {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsisBfdCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Bfd = in["bfd"].(int)
		ret.Disable = in["disable"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIsisCsnpIntervalList(d []interface{}) []edpt.InterfaceEthernetIsisCsnpIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisCsnpIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisCsnpIntervalList
		oi.CsnpInterval = in["csnp_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisHelloIntervalList(d []interface{}) []edpt.InterfaceEthernetIsisHelloIntervalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisHelloIntervalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisHelloIntervalList
		oi.HelloInterval = in["hello_interval"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisHelloIntervalMinimalList(d []interface{}) []edpt.InterfaceEthernetIsisHelloIntervalMinimalList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisHelloIntervalMinimalList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisHelloIntervalMinimalList
		oi.HelloIntervalMinimal = in["hello_interval_minimal"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisHelloMultiplierList(d []interface{}) []edpt.InterfaceEthernetIsisHelloMultiplierList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisHelloMultiplierList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisHelloMultiplierList
		oi.HelloMultiplier = in["hello_multiplier"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectInterfaceEthernetIsisMeshGroup(d []interface{}) edpt.InterfaceEthernetIsisMeshGroup {

	count1 := len(d)
	var ret edpt.InterfaceEthernetIsisMeshGroup
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Value = in["value"].(int)
		ret.Blocked = in["blocked"].(int)
	}
	return ret
}

func getSliceInterfaceEthernetIsisMetricList(d []interface{}) []edpt.InterfaceEthernetIsisMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisMetricList
		oi.Metric = in["metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisPasswordList(d []interface{}) []edpt.InterfaceEthernetIsisPasswordList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisPasswordList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisPasswordList
		oi.Password = in["password"].(string)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisPriorityList(d []interface{}) []edpt.InterfaceEthernetIsisPriorityList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisPriorityList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisPriorityList
		oi.Priority = in["priority"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceEthernetIsisWideMetricList(d []interface{}) []edpt.InterfaceEthernetIsisWideMetricList {

	count1 := len(d)
	ret := make([]edpt.InterfaceEthernetIsisWideMetricList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceEthernetIsisWideMetricList
		oi.WideMetric = in["wide_metric"].(int)
		oi.Level = in["level"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceEthernetIsis(d *schema.ResourceData) edpt.InterfaceEthernetIsis {
	var ret edpt.InterfaceEthernetIsis
	ret.Inst.Authentication = getObjectInterfaceEthernetIsisAuthentication(d.Get("authentication").([]interface{}))
	ret.Inst.BfdCfg = getObjectInterfaceEthernetIsisBfdCfg(d.Get("bfd_cfg").([]interface{}))
	ret.Inst.CircuitType = d.Get("circuit_type").(string)
	ret.Inst.CsnpIntervalList = getSliceInterfaceEthernetIsisCsnpIntervalList(d.Get("csnp_interval_list").([]interface{}))
	ret.Inst.HelloIntervalList = getSliceInterfaceEthernetIsisHelloIntervalList(d.Get("hello_interval_list").([]interface{}))
	ret.Inst.HelloIntervalMinimalList = getSliceInterfaceEthernetIsisHelloIntervalMinimalList(d.Get("hello_interval_minimal_list").([]interface{}))
	ret.Inst.HelloMultiplierList = getSliceInterfaceEthernetIsisHelloMultiplierList(d.Get("hello_multiplier_list").([]interface{}))
	ret.Inst.LspInterval = d.Get("lsp_interval").(int)
	ret.Inst.MeshGroup = getObjectInterfaceEthernetIsisMeshGroup(d.Get("mesh_group").([]interface{}))
	ret.Inst.MetricList = getSliceInterfaceEthernetIsisMetricList(d.Get("metric_list").([]interface{}))
	ret.Inst.Network = d.Get("network").(string)
	ret.Inst.Padding = d.Get("padding").(int)
	ret.Inst.PasswordList = getSliceInterfaceEthernetIsisPasswordList(d.Get("password_list").([]interface{}))
	ret.Inst.PriorityList = getSliceInterfaceEthernetIsisPriorityList(d.Get("priority_list").([]interface{}))
	ret.Inst.RetransmitInterval = d.Get("retransmit_interval").(int)
	//omit uuid
	ret.Inst.WideMetricList = getSliceInterfaceEthernetIsisWideMetricList(d.Get("wide_metric_list").([]interface{}))
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
