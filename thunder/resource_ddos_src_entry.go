package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSrcEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_src_entry`: Configure IP/IPv6 static Entry\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSrcEntryCreate,
		UpdateContext: resourceDdosSrcEntryUpdate,
		ReadContext:   resourceDdosSrcEntryRead,
		DeleteContext: resourceDdosSrcEntryDelete,

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
			"description": {
				Type: schema.TypeString, Optional: true, Description: "Description for this Source Entry",
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
			"hw_blacklist_blocking": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Src side hardware blocking",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"l4_type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": {
							Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'permit': Whitelist incoming packets for protocol; 'deny': Blacklist incoming packets for protocol;",
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
			"src_entry_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
			"subnet_ip_addr": {
				Type: schema.TypeString, Optional: true, Description: "IP Subnet",
			},
			"subnet_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 Subnet",
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
func resourceDdosSrcEntryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntry(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcEntryRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSrcEntryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntry(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSrcEntryRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSrcEntryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntry(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSrcEntryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSrcEntryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSrcEntry(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSrcEntryAppTypeList(d []interface{}) []edpt.DdosSrcEntryAppTypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcEntryAppTypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcEntryAppTypeList
		oi.Protocol = in["protocol"].(string)
		oi.Template = getObjectDdosSrcEntryAppTypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcEntryAppTypeListTemplate(d []interface{}) edpt.DdosSrcEntryAppTypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcEntryAppTypeListTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.Sip = in["sip"].(string)
	}
	return ret
}

func getObjectDdosSrcEntryExceedLogCfg(d []interface{}) edpt.DdosSrcEntryExceedLogCfg {

	count1 := len(d)
	var ret edpt.DdosSrcEntryExceedLogCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogEnable = in["log_enable"].(int)
	}
	return ret
}

func getObjectDdosSrcEntryHwBlacklistBlocking293(d []interface{}) edpt.DdosSrcEntryHwBlacklistBlocking293 {

	count1 := len(d)
	var ret edpt.DdosSrcEntryHwBlacklistBlocking293
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SrcEnable = in["src_enable"].(int)
		//omit uuid
	}
	return ret
}

func getSliceDdosSrcEntryL4TypeList(d []interface{}) []edpt.DdosSrcEntryL4TypeList {

	count1 := len(d)
	ret := make([]edpt.DdosSrcEntryL4TypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSrcEntryL4TypeList
		oi.Protocol = in["protocol"].(string)
		oi.Action = in["action"].(string)
		oi.Glid = in["glid"].(string)
		oi.Template = getObjectDdosSrcEntryL4TypeListTemplate(in["template"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSrcEntryL4TypeListTemplate(d []interface{}) edpt.DdosSrcEntryL4TypeListTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcEntryL4TypeListTemplate
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

func getObjectDdosSrcEntryTemplate(d []interface{}) edpt.DdosSrcEntryTemplate {

	count1 := len(d)
	var ret edpt.DdosSrcEntryTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
	}
	return ret
}

func dataToEndpointDdosSrcEntry(d *schema.ResourceData) edpt.DdosSrcEntry {
	var ret edpt.DdosSrcEntry
	ret.Inst.AppTypeList = getSliceDdosSrcEntryAppTypeList(d.Get("app_type_list").([]interface{}))
	ret.Inst.Bypass = d.Get("bypass").(int)
	ret.Inst.Description = d.Get("description").(string)
	ret.Inst.ExceedLogCfg = getObjectDdosSrcEntryExceedLogCfg(d.Get("exceed_log_cfg").([]interface{}))
	ret.Inst.Glid = d.Get("glid").(string)
	ret.Inst.HwBlacklistBlocking = getObjectDdosSrcEntryHwBlacklistBlocking293(d.Get("hw_blacklist_blocking").([]interface{}))
	ret.Inst.IpAddr = d.Get("ip_addr").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.L4TypeList = getSliceDdosSrcEntryL4TypeList(d.Get("l4_type_list").([]interface{}))
	ret.Inst.LogPeriodic = d.Get("log_periodic").(int)
	ret.Inst.SrcEntryName = d.Get("src_entry_name").(string)
	ret.Inst.SubnetIpAddr = d.Get("subnet_ip_addr").(string)
	ret.Inst.SubnetIpv6Addr = d.Get("subnet_ipv6_addr").(string)
	ret.Inst.Template = getObjectDdosSrcEntryTemplate(d.Get("template").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
