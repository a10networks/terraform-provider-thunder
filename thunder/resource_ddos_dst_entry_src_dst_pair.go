package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPair() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair`: per-Source-Destination pair Source-limiting configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairDelete,

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
			"default": {
				Type: schema.TypeInt, Required: true, Description: "Configure default",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dst_entry_name": {
				Type: schema.TypeString, Required: true, Description: "DstEntryName",
			},
		},
	}
}
func resourceDdosDstEntrySrcDstPairCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPair(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPair(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPair(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPair(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairAppTypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairAppTypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairAppTypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairAppTypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairAppTypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairAppTypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairAppTypeSrcDstListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairL4TypeSrcDstList(d []interface{}) []edpt.DdosDstEntrySrcDstPairL4TypeSrcDstList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairL4TypeSrcDstList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairL4TypeSrcDstList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairL4TypeSrcDstListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairL4TypeSrcDstListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairL4TypeSrcDstListTemplate
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

func getObjectDdosDstEntrySrcDstPairTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosDstEntrySrcDstPair(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPair {
	var ret edpt.DdosDstEntrySrcDstPair
	ret.Inst.AppTypeSrcDstList = getSliceDdosDstEntrySrcDstPairAppTypeSrcDstList(d.Get("app_type_src_dst_list").([]interface{}))
	ret.Inst.Bypass = d.Get("bypass").(int)
	ret.Inst.Default = d.Get("default").(int)
	ret.Inst.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.L4TypeSrcDstList = getSliceDdosDstEntrySrcDstPairL4TypeSrcDstList(d.Get("l4_type_src_dst_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.Template = getObjectDdosDstEntrySrcDstPairTemplate(d.Get("template").([]interface{}))
	//omit uuid
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
