package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list`: Configure class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairPolicyPolicyClassListRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListDelete,

		Schema: map[string]*schema.Schema{
			"app_type_src_dst_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssl_l4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
									},
									"dns": {
										Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
									},
									"http": {
										Type: schema.TypeString, Optional: true, Description: "DDOS http template",
									},
									"sip": {
										Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
			},
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
			},
			"class_list_overflow_policy_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy_name": {
							Type: schema.TypeString, Required: true, Description: "'configuration': Configure src dst dynamic entry count overflow policy for class-list;",
						},
						"bypass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
						},
						"exceed_log_cfg": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_enable": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
									},
								},
							},
						},
						"log_periodic": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logging": {
										Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
									},
								},
							},
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"l4_type_src_dst_overflow_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
									},
									"deny": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
									},
									"glid": {
										Type: schema.TypeString, Optional: true, Description: "Global limit ID",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tcp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
												},
												"udp": {
													Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
												},
												"other": {
													Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
												},
												"template_icmp_v4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
												},
												"template_icmp_v6": {
													Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
						"app_type_src_dst_overflow_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type: schema.TypeString, Required: true, Description: "'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;",
									},
									"template": {
										Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ssl_l4": {
													Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
												},
												"dns": {
													Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
												},
												"http": {
													Type: schema.TypeString, Optional: true, Description: "DDOS http template",
												},
												"sip": {
													Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
												},
											},
										},
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
									"user_tag": {
										Type: schema.TypeString, Optional: true, Description: "Customized tag",
									},
								},
							},
						},
					},
				},
			},
			"exceed_log_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
						},
					},
				},
			},
			"glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"l4_type_src_dst_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
						},
						"deny": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Blacklist and Drop all incoming packets for protocol",
						},
						"glid": {
							Type: schema.TypeString, Optional: true, Description: "Global limit ID",
						},
						"template": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tcp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS TCP template",
									},
									"udp": {
										Type: schema.TypeString, Optional: true, Description: "DDOS UDP template",
									},
									"other": {
										Type: schema.TypeString, Optional: true, Description: "DDOS OTHER template",
									},
									"template_icmp_v4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v4 template",
									},
									"template_icmp_v6": {
										Type: schema.TypeString, Optional: true, Description: "DDOS icmp-v6 template",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
			},
			"max_dynamic_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry under class-list",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'packet_received': Packets Received; 'packet_dropped': Packets Dropped; 'entry_learned': Entry Learned; 'entry_count_overflow': Entry Count Overflow;",
						},
					},
				},
			},
			"template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging template",
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
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyPolicyClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyPolicyClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Bypass = in["bypass"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListTemplate(in["template"].([]interface{}))
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList(in["l4_type_src_dst_overflow_list"].([]interface{}))
		oi.AppTypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList(in["app_type_src_dst_overflow_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Other = in["other"].(string)
		ret.TemplateIcmpV4 = in["template_icmp_v4"].(string)
		ret.TemplateIcmpV6 = in["template_icmp_v6"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassList(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassList {
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassList
	ret.Inst.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListAppTypeSrcDstList(d.Get("app_type_src_dst_list").([]interface{}))
	ret.Inst.Bypass = d.Get("bypass").(int)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.ClassListOverflowPolicyList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyList(d.Get("class_list_overflow_policy_list").([]interface{}))
	ret.Inst.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListL4TypeSrcDstList(d.Get("l4_type_src_dst_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.SamplingEnable = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
