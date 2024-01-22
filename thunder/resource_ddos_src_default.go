package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcDefault() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_default`: Configure IP/IPv6 default entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcDefaultCreate,
		UpdateContext: resourceDdosSrcDefaultUpdate,
		ReadContext:   resourceDdosSrcDefaultRead,
		DeleteContext: resourceDdosSrcDefaultDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle age for ip entry",
			},
			"app_type_list": {
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
									"dns": {
										Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
									},
									"http": {
										Type: schema.TypeString, Optional: true, Description: "DDOS http template",
									},
									"ssl_l4": {
										Type: schema.TypeString, Optional: true, Description: "DDOS SSL-L4 template",
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
			"apply_policy_on_overflow": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this flag to apply overflow policy when dynamic entry count overflows",
			},
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "'ip': ip; 'ipv6': ipv6;",
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable",
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
			"l4_type_list": {
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
				Type: schema.TypeInt, Optional: true, Description: "Maximum count for dynamic src entry",
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
		},
	}
}
func resourceDdosSrcDefaultCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDefaultCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDefault(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcDefaultRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDefaultUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDefault(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcDefaultRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcDefaultDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDefaultDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDefault(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcDefaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDefaultRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDefault(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcDefaultAppTypeList(d []interface{}) []edpt.DdosSrcDefaultAppTypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcDefaultAppTypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcDefaultAppTypeList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosSrcDefaultAppTypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcDefaultAppTypeListTemplate(d []interface{}) edpt.DdosSrcDefaultAppTypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDefaultAppTypeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosSrcDefaultExceedLogCfg(d []interface{}) edpt.DdosSrcDefaultExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosSrcDefaultExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getSliceDdosSrcDefaultL4TypeList(d []interface{}) []edpt.DdosSrcDefaultL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcDefaultL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcDefaultL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosSrcDefaultL4TypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcDefaultL4TypeListTemplate(d []interface{}) edpt.DdosSrcDefaultL4TypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDefaultL4TypeListTemplate
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

func getObjectDdosSrcDefaultTemplate(d []interface{}) edpt.DdosSrcDefaultTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDefaultTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosSrcDefault(d *schema.ResourceData) edpt.DdosSrcDefault {
	var ret edpt.DdosSrcDefault
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.AppTypeList = getSliceDdosSrcDefaultAppTypeList(d.Get("app_type_list").([]interface{}))
	ret.Inst.ApplyPolicyOnOverflow = d.Get("apply_policy_on_overflow").(int)
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.ExceedLogCfg = getObjectDdosSrcDefaultExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.L4TypeList = getSliceDdosSrcDefaultL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.MaxDynamicEntryCount = d.Get("max_dynamic_entry_count").(int)
	ret.Inst.Template = getObjectDdosSrcDefaultTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
