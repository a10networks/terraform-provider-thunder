package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcDynamicEntryOverflowPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_dynamic_entry_overflow_policy`: Configure IP/IPv6 Policy Used When Src Dynamic Entry Count Overflows\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcDynamicEntryOverflowPolicyCreate,
		UpdateContext: resourceDdosSrcDynamicEntryOverflowPolicyUpdate,
		ReadContext:   resourceDdosSrcDynamicEntryOverflowPolicyRead,
		DeleteContext: resourceDdosSrcDynamicEntryOverflowPolicyDelete,

		Schema: map[string]*schema.Schema{
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
			"default_address_type": {
				Type: schema.TypeString, Required: true, Description: "'ip': ip; 'ipv6': ipv6;",
			},
			"exceed_log_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging of limit exceed drop's",
						},
						"with_sflow_sample": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn on sflow sample with log",
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
func resourceDdosSrcDynamicEntryOverflowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcDynamicEntryOverflowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcDynamicEntryOverflowPolicyRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcDynamicEntryOverflowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcDynamicEntryOverflowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcDynamicEntryOverflowPolicyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcDynamicEntryOverflowPolicy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcDynamicEntryOverflowPolicyAppTypeList(d []interface{}) []edpt.DdosSrcDynamicEntryOverflowPolicyAppTypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcDynamicEntryOverflowPolicyAppTypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcDynamicEntryOverflowPolicyAppTypeList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosSrcDynamicEntryOverflowPolicyAppTypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcDynamicEntryOverflowPolicyAppTypeListTemplate(d []interface{}) edpt.DdosSrcDynamicEntryOverflowPolicyAppTypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDynamicEntryOverflowPolicyAppTypeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosSrcDynamicEntryOverflowPolicyExceedLogCfg(d []interface{}) edpt.DdosSrcDynamicEntryOverflowPolicyExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosSrcDynamicEntryOverflowPolicyExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
		ret.WithSflowSample = in["with_sflow_sample"].(int)
	}
	return ret
}

func getSliceDdosSrcDynamicEntryOverflowPolicyL4TypeList(d []interface{}) []edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Deny = in["deny"].(int)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosSrcDynamicEntryOverflowPolicyL4TypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcDynamicEntryOverflowPolicyL4TypeListTemplate(d []interface{}) edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDynamicEntryOverflowPolicyL4TypeListTemplate
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

func getObjectDdosSrcDynamicEntryOverflowPolicyTemplate(d []interface{}) edpt.DdosSrcDynamicEntryOverflowPolicyTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcDynamicEntryOverflowPolicyTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosSrcDynamicEntryOverflowPolicy(d *schema.ResourceData) edpt.DdosSrcDynamicEntryOverflowPolicy {
	var ret edpt.DdosSrcDynamicEntryOverflowPolicy
	ret.Inst.AppTypeList = getSliceDdosSrcDynamicEntryOverflowPolicyAppTypeList(d.Get("app_type_list").([]interface{}))
	ret.Inst.DefaultAddressType = d.Get("default_address_type").(string)
	ret.Inst.ExceedLogCfg = getObjectDdosSrcDynamicEntryOverflowPolicyExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.L4TypeList = getSliceDdosSrcDynamicEntryOverflowPolicyL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.Template = getObjectDdosSrcDynamicEntryOverflowPolicyTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
