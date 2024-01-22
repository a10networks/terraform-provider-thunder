package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDynamicService() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dynamic_service`: Dynamic service template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDynamicServiceCreate,
		UpdateContext: resourceSlbTemplateDynamicServiceUpdate,
		ReadContext:   resourceSlbTemplateDynamicServiceRead,
		DeleteContext: resourceSlbTemplateDynamicServiceDelete,

		Schema: map[string]*schema.Schema{
			"class_list_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_class_list": {
							Type: schema.TypeString, Required: true, Description: "Name of Aho-Corasick class-list",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Description: "Priority of the class-list(the larger number, the higher priority)",
						},
						"dns_server": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_dns_server": {
										Type: schema.TypeString, Optional: true, Description: "DNS Server IPv4 Address",
									},
									"ipv6_dns_server": {
										Type: schema.TypeString, Optional: true, Description: "DNS Server IPv6 Address",
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
			"dns_server": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_dns_server": {
							Type: schema.TypeString, Optional: true, Description: "DNS Server IPv4 Address",
						},
						"ipv6_dns_server": {
							Type: schema.TypeString, Optional: true, Description: "DNS Server IPv6 Address",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Dynamic Service Template Name",
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
func resourceSlbTemplateDynamicServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicService(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDynamicServiceRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDynamicServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicService(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDynamicServiceRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDynamicServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicService(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDynamicServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDynamicServiceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDynamicService(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDynamicServiceClassListList(d []interface{}) []edpt.SlbTemplateDynamicServiceClassListList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDynamicServiceClassListList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDynamicServiceClassListList
		oi.DnsClassList = in["dns_class_list"].(string)
		oi.Priority = in["priority"].(int)
		oi.DnsServer = getSliceSlbTemplateDynamicServiceClassListListDnsServer(in["dns_server"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDynamicServiceClassListListDnsServer(d []interface{}) []edpt.SlbTemplateDynamicServiceClassListListDnsServer {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDynamicServiceClassListListDnsServer, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDynamicServiceClassListListDnsServer
		oi.Ipv4DnsServer = in["ipv4_dns_server"].(string)
		oi.Ipv6DnsServer = in["ipv6_dns_server"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateDynamicServiceDnsServer(d []interface{}) []edpt.SlbTemplateDynamicServiceDnsServer {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDynamicServiceDnsServer, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDynamicServiceDnsServer
		oi.Ipv4DnsServer = in["ipv4_dns_server"].(string)
		oi.Ipv6DnsServer = in["ipv6_dns_server"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDynamicService(d *schema.ResourceData) edpt.SlbTemplateDynamicService {
	var ret edpt.SlbTemplateDynamicService
	ret.Inst.ClassListList = getSliceSlbTemplateDynamicServiceClassListList(d.Get("class_list_list").([]interface{}))
	ret.Inst.DnsServer = getSliceSlbTemplateDynamicServiceDnsServer(d.Get("dns_server").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
