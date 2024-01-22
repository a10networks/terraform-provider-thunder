package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_policy`: Configure class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairPolicyCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairPolicyUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairPolicyRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairPolicyDelete,

		Schema: map[string]*schema.Schema{
			"policy_class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"class_list_name": {
							Type: schema.TypeString, Required: true, Description: "Class-list name",
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
						"max_dynamic_entry_count": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src-dst entry under class-list",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
					},
				},
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "Src-based-policy name",
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
func resourceDdosDstEntrySrcDstPairPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListList
		oi.ClassListName = in["class_list_name"].(string)
		oi.Bypass = in["bypass"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListTemplate(in["template"].([]interface{}))
		oi.Glid = in["glid"].(string)
		oi.MaxDynamicEntryCount = in["max_dynamic_entry_count"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.SamplingEnable = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable(in["sampling_enable"].([]interface{}))
		oi.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList(in["l4_type_src_dst_list"].([]interface{}))
		oi.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList(in["app_type_src_dst_list"].([]interface{}))
		oi.ClassListOverflowPolicyList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList(in["class_list_overflow_policy_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListL4TypeSrcDstListTemplate
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

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyList
		oi.DummyName = in["dummy_name"].(string)
		oi.Bypass = in["bypass"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListTemplate(in["template"].([]interface{}))
		oi.Glid = in["glid"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList(in["l4_type_src_dst_overflow_list"].([]interface{}))
		oi.AppTypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList(in["app_type_src_dst_overflow_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListL4TypeSrcDstOverflowListTemplate
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

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListListClassListOverflowPolicyListAppTypeSrcDstOverflowListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairPolicy(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairPolicy {
	var ret edpt.DdosDstEntrySrcDstPairPolicy
	ret.Inst.PolicyClassListList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListList(d.Get("policy_class_list_list").([]interface{}))
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
