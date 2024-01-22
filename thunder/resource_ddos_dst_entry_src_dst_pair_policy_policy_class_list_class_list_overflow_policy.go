package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy`: Configure src dst dynamic entry count overflow policy for class-list\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyDelete,

		Schema: map[string]*schema.Schema{
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
			"bypass": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Always permit for the Source to bypass all feature & limit checks",
			},
			"dummy_name": {
				Type: schema.TypeString, Required: true, Description: "'configuration': Configure src dst dynamic entry count overflow policy for class-list;",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"src_based_policy_name": {
				Type: schema.TypeString, Required: true, Description: "SrcBasedPolicyName",
			},
			"class_list_name": {
				Type: schema.TypeString, Required: true, Description: "ClassListName",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList(d []interface{}) []edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowListTemplate
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

func getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy {
	var ret edpt.DdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicy
	ret.Inst.AppTypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyAppTypeSrcDstOverflowList(d.Get("app_type_src_dst_overflow_list").([]interface{}))
	ret.Inst.Bypass = d.Get("bypass").(int)
	ret.Inst.DummyName = d.Get("dummy_name").(string)
	ret.Inst.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.L4TypeSrcDstOverflowList = getSliceDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyL4TypeSrcDstOverflowList(d.Get("l4_type_src_dst_overflow_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.Template = getObjectDdosDstEntrySrcDstPairPolicyPolicyClassListClassListOverflowPolicyTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.SrcBasedPolicyName = d.Get("src_based_policy_name").(string)
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
