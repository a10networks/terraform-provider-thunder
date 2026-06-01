package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstEntrySrcDstPairClassListCid() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_entry_src_dst_pair_class_list_cid`: Class-list id\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstEntrySrcDstPairClassListCidCreate,
		UpdateContext: resourceDdosDstEntrySrcDstPairClassListCidUpdate,
		ReadContext:   resourceDdosDstEntrySrcDstPairClassListCidRead,
		DeleteContext: resourceDdosDstEntrySrcDstPairClassListCidDelete,

		Schema: map[string]*schema.Schema{
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
			"log_periodic": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable periodic log while event is continuing",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceDdosDstEntrySrcDstPairClassListCidCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCid(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListCidRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListCidUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCid(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstEntrySrcDstPairClassListCidRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstEntrySrcDstPairClassListCidDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCid(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstEntrySrcDstPairClassListCidRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstEntrySrcDstPairClassListCidRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstEntrySrcDstPairClassListCid(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListCidExceedLogCfg(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidList(d []interface{}) []edpt.DdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidList {

	count1 := len(d)
	ret := make([]edpt.DdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidListTemplate(d []interface{}) edpt.DdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidListTemplate {

	count1 := len(d)
	var ret edpt.DdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidListTemplate
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

func dataToEndpointDdosDstEntrySrcDstPairClassListCid(d *schema.ResourceData) edpt.DdosDstEntrySrcDstPairClassListCid {
	var ret edpt.DdosDstEntrySrcDstPairClassListCid
	ret.Inst.AppTypeSrcDstCidList = getSliceDdosDstEntrySrcDstPairClassListCidAppTypeSrcDstCidList(d.Get("app_type_src_dst_cid_list").([]interface{}))
	ret.Inst.CidNum = d.Get("cid_num").(int)
	ret.Inst.ExceedLogCfg = getObjectDdosDstEntrySrcDstPairClassListCidExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.L4TypeSrcDstCidList = getSliceDdosDstEntrySrcDstPairClassListCidL4TypeSrcDstCidList(d.Get("l4_type_src_dst_cid_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ClassListName = d.Get("class_list_name").(string)
	ret.Inst.DstEntryName = d.Get("dst_entry_name").(string)
	return ret
}
