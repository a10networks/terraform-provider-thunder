package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_class_list`: Configure class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairClassListCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairClassListUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairClassListRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairClassListDelete,

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
			"cid_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cid_num": {
							Type: schema.TypeInt, Required: true, Description: "Class-list id",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
						"l4_type_src_dst_cid_list": {
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
						"app_type_src_dst_cid_list": {
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
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "Class-list name",
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
func resourceDdosDstEntrySrcDstPairClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairClassListAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListCidList
		oi.CidNum = in["cid_num"].(int)
		oi.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairClassListCidListExceedLogCfg(in["exceed_log_cfg"].([]interface{}))
		oi.LogPeriodic = in["log_periodic"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		oi.L4TypeSrcDstCidList = getSliceDdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList(in["l4_type_src_dst_cid_list"].([]interface{}))
		oi.AppTypeSrcDstCidList = getSliceDdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList(in["app_type_src_dst_cid_list"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListCidListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidListL4TypeSrcDstCidListTemplate
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

func getSliceDdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidListAppTypeSrcDstCidListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListL4TypeSrcDstListTemplate
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

func dataToEndpointDdosDstEntrySrcDstPairClassList(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairClassList {
	var ret edpt.DdosDstEntrySrcDstPairClassList
	ret.Inst.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairClassListAppTypeSrcDstList(d.Get("app_type_src_dst_list").([]interface{}))
	ret.Inst.CidList = getSliceDdosDstEntrySrcDstPairClassListCidList(d.Get("cid_list").([]interface{}))
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairClassListExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairClassListL4TypeSrcDstList(d.Get("l4_type_src_dst_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
